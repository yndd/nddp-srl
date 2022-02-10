package devicecollector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
	"github.com/yndd/nddp-srl/internal/shared"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

const (
	unmanagedResource = "Unmanaged resource"
)

func (c *collector) handleSubscription(resp *gnmi.SubscribeResponse) error {
	log := c.log.WithValues("target", c.target.Config.Name, "address", c.target.Config.Address)
	//log.Debug("handle target update from device")

	switch resp.GetResponse().(type) {
	case *gnmi.SubscribeResponse_Update:
		//log.Debug("handle target update from device", "Prefix", resp.GetUpdate().GetPrefix())

		// check if the target cache exists
		crDeviceName := shared.GetCrDeviceName(c.namespace, c.target.Config.Name)

		if !c.cache.GetCache().HasTarget(crDeviceName) {
			log.Debug("handle target update target not found in cache")
			return errors.New("target cache does not exist")
		}

		resourceList, err := c.getResourceList(crDeviceName)
		if err != nil {
			return err
		}
		//log.Debug("resourceList", "list", resourceList)

		// handle deletes
		c.handleDeletes(crDeviceName, resourceList, resp.GetUpdate().Delete)

		c.handleUpdates(crDeviceName, resourceList, resp.GetUpdate().Update)

	case *gnmi.SubscribeResponse_SyncResponse:
		//log.Debug("SyncResponse")
	}

	return nil
}

func (c *collector) handleDeletes(crDeviceName string, resourceList []*systemv1alpha1.Gvk, delPaths []*gnmi.Path) error {
	for _, path := range delPaths {
		xpath := yparser.GnmiPath2XPath(path, true)
		resourceName, err := c.findManagedResource(xpath, resourceList)
		if err != nil {
			return err
		}

		// clean the path for now to remove the module information from the pathElem
		for _, pe := range path.GetElem() {
			pe.Name = strings.Split(pe.Name, ":")[len(strings.Split(pe.Name, ":"))-1]
		}

		// if a default is enabled on the path we should revert to default
		def := c.deviceSchema.GetPathDefault(path)
		var n *gnmi.Notification
		c.log.Debug("collector config delete", "path", xpath, "default", def)
		if def != "" {
			d, err := json.Marshal(def)
			if err != nil {
				return err
			}
			// if the data is empty, there is no need for an update
			if string(d) == "null" {
				return nil
			}

			n = &gnmi.Notification{
				Timestamp: time.Now().UnixNano(),
				Prefix:    &gnmi.Path{Target: crDeviceName},
				Update: []*gnmi.Update{
					{
						Path: path,
						Val: &gnmi.TypedValue{
							Value: &gnmi.TypedValue_JsonIetfVal{
								JsonIetfVal: bytes.Trim(d, " \r\n\t"),
							},
						},
					},
				},
			}
		} else {
			n = &gnmi.Notification{
				Timestamp: time.Now().UnixNano(),
				Prefix:    &gnmi.Path{Target: crDeviceName},
				Delete:    []*gnmi.Path{path},
			}
		}
		// update the cache with the latest config from the device
		if err := c.cache.GnmiUpdate(crDeviceName, n); err != nil {
			c.log.Debug("handle target update", "error", err, "Path", xpath)
			return errors.New("cache update failed")
		}

		if *resourceName != unmanagedResource {
			// TODO Trigger reconcile event
			c.triggerReconcileEvent(resourceName)
		}
	}
	return nil
}

func (c *collector) handleUpdates(crDeviceName string, resourceList []*systemv1alpha1.Gvk, u []*gnmi.Update) error {
	for _, upd := range u {
		xpath := yparser.GnmiPath2XPath(upd.GetPath(), true)
		// check if this is a managed resource or unmanged resource
		// name == unmanagedResource is an unmanaged resource
		resourceName, err := c.findManagedResource(xpath, resourceList)
		if err != nil {
			return err
		}

		// clean the path for now to remove the module information from the pathElem
		for _, pe := range upd.GetPath().GetElem() {
			pe.Name = strings.Split(pe.Name, ":")[len(strings.Split(pe.Name, ":"))-1]
			keys := make(map[string]string)
			for k, v := range pe.GetKey() {
				if strings.Contains(v, "::") {
					keys[strings.Split(k, ":")[len(strings.Split(k, ":"))-1]] = v
				} else {
					keys[strings.Split(k, ":")[len(strings.Split(k, ":"))-1]] = strings.Split(v, ":")[len(strings.Split(v, ":"))-1]
				}
			}
			pe.Key = keys
		}

		switch upd.GetVal().Value.(type) {
		case *gnmi.TypedValue_JsonIetfVal:
			jsondata := upd.GetVal().GetJsonIetfVal()
			var v interface{}
			if len(jsondata) != 0 {
				err := json.Unmarshal(jsondata, &v)
				if err != nil {
					return err
				}
			}

			/*
				if strings.Contains(string(jsondata), "routed") {
					//fmt.Printf("type of data1: %v %v\n", string(jsondata), reflect.TypeOf(jsondata))
					fmt.Printf("type of data2: %v %v\n", string(jsondata), reflect.TypeOf(v))
				}
			*/
			switch vv := v.(type) {
			case map[string]interface{}:
				vv = yparser.CleanConfig(vv)
				b, err := json.Marshal(vv)
				if err != nil {
					return err
				}
				//fmt.Printf("string cleaned: %s\n", string(b))
				upd.Val = &gnmi.TypedValue{
					Value: &gnmi.TypedValue_JsonIetfVal{
						JsonIetfVal: bytes.Trim(b, " \r\n\t"),
					},
				}
			case string:
				// for string values there can be also a header in the values e.g. type, Value: srl_nokia-network-instance:ip-vrf
				if !strings.Contains(vv, "::") {
					// if there are more ":" in the string it is likely an esi or mac address
					if len(strings.Split(vv, ":")) <= 2 {
						//fmt.Printf("string to be cleaned: %s\n", vv)
						vv = strings.Split(vv, ":")[len(strings.Split(vv, ":"))-1]
						b, err := json.Marshal(vv)
						if err != nil {
							return err
						}
						//fmt.Printf("string cleaned: %s\n", string(b))
						upd.Val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_JsonIetfVal{
								JsonIetfVal: bytes.Trim(b, " \r\n\t"),
							},
						}
					}
				}
			}
		}

		// Validate if the path has a key using the device schema
		// Used to allow insertion of an empty container
		keys := c.deviceSchema.GetKeys(upd.GetPath())
		hashKey := false
		if len(keys) != 0 {
			hashKey = true
		}
		crDeviceName := shared.GetCrDeviceName(c.namespace, c.target.Config.Name)
		n, err := c.cache.GetNotificationFromUpdate(&gnmi.Path{Target: crDeviceName}, upd, hashKey)
		if err != nil {
			return err
		}

		// default handling
		defaults := c.deviceSchema.GetPathDefaults(upd.GetPath())
		for pathElemName, defValue := range defaults {
			c.log.Debug("collector config update defaults", "pathElemName", pathElemName, "defValue", defValue, "path", yparser.GnmiPath2XPath(upd.GetPath(), true))

			d, err := json.Marshal(defValue)
			if err != nil {
				return err
			}
			// if the data is empty, there is no need for an update
			if string(d) == "null" {
				return nil
			}

			// check if the element exists in the original notification
			// if not we add the default, if it is there we avoid adding the default
			found := false
			for _, nu := range n.GetUpdate() {
				if nu.GetPath().GetElem()[len(nu.GetPath().GetElem())-1].GetName() == pathElemName {
					found = true
				}
			}
			if !found {
				newPath := yparser.DeepCopyGnmiPath(upd.GetPath())
				newPath.Elem = append(newPath.GetElem(), &gnmi.PathElem{Name: pathElemName})
				u := &gnmi.Update{
					Path: newPath,
					Val: &gnmi.TypedValue{
						Value: &gnmi.TypedValue_JsonIetfVal{
							JsonIetfVal: bytes.Trim(d, " \r\n\t"),
						},
					},
				}
				n.Update = append(n.GetUpdate(), u)
			}

		}

		// update the cache with the latest config from the device
		if err := c.cache.GnmiUpdate(crDeviceName, n); err != nil {
			for _, u := range n.GetUpdate() {
				c.log.Debug("collector config update", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "value", u.GetVal(), "error", err)
			}
			return errors.Wrap(err, "cache update failed")
		}

		for _, u := range n.GetUpdate() {
			c.log.Debug("collector config update", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "value", u.GetVal(), "error", err)
		}

		if *resourceName != unmanagedResource {
			c.triggerReconcileEvent(resourceName)
		}
	}
	return nil
}

func (c *collector) findManagedResource(xpath string, resourceList []*systemv1alpha1.Gvk) (*string, error) {
	matchedResourceName := unmanagedResource
	matchedResourcePath := ""
	for _, r := range resourceList {
		if strings.Contains(xpath, *r.Rootpath) {
			// if there is a better match we use the better match
			if len(*r.Rootpath) > len(matchedResourcePath) {
				matchedResourcePath = *r.Rootpath
				matchedResourceName = *r.Name
			}
		}
	}
	return &matchedResourceName, nil
}

func (c *collector) getResourceList(crDeviceName string) ([]*systemv1alpha1.Gvk, error) {
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	rl, err := c.cache.GetJson(crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{Elem: []*gnmi.PathElem{{Name: "gvk"}}},
		c.nddpSchema)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return gvkresource.GetResourceList(rl)
}

func (c *collector) triggerReconcileEvent(resourceName *string) error {

	gvk, err := gvkresource.String2Gvk(*resourceName)
	if err != nil {
		return err
	}
	kindgroup := strings.Join([]string{gvk.Kind, gvk.Group}, ".")

	object := getObject(gvk)

	//c.log.Debug("triggerReconcileEvent", "kindgroup", kindgroup, "gvk", gvk, "object", object)

	if eventCh, ok := c.eventChs[kindgroup]; ok {
		c.log.Debug("triggerReconcileEvent with channel lookup", "kindgroup", kindgroup, "gvk", gvk, "object", object)
		eventCh <- event.GenericEvent{
			Object: object,
		}
	}
	return nil
}

func getObject(gvk *gvkresource.GVK) client.Object {
	switch gvk.Kind {
	case "SrlBfd":
		return &srlv1alpha1.SrlBfd{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlInterface":
		return &srlv1alpha1.SrlInterface{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlInterfaceSubinterface":
		return &srlv1alpha1.SrlInterfaceSubinterface{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlRoutingpolicyPrefixset":
		return &srlv1alpha1.SrlRoutingpolicyPrefixset{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlRoutingpolicyPolicy":
		return &srlv1alpha1.SrlRoutingpolicyPolicy{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlRoutingpolicyAspathset":
		return &srlv1alpha1.SrlRoutingpolicyAspathset{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlRoutingpolicyCommunityset":
		return &srlv1alpha1.SrlRoutingpolicyCommunityset{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstance":
		return &srlv1alpha1.SrlNetworkinstance{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstanceAggregateroutes":
		return &srlv1alpha1.SrlNetworkinstanceAggregateroutes{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstanceNexthopgroups":
		return &srlv1alpha1.SrlNetworkinstanceNexthopgroups{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstanceProtocolsBgp":
		return &srlv1alpha1.SrlNetworkinstanceProtocolsBgp{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstanceProtocolsBgpevpn":
		return &srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstanceProtocolsBgpvpn":
		return &srlv1alpha1.SrlNetworkinstanceProtocolsBgpvpn{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstanceProtocolsIsis":
		return &srlv1alpha1.SrlNetworkinstanceProtocolsIsis{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstanceProtocolsLinux":
		return &srlv1alpha1.SrlNetworkinstanceProtocolsLinux{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstanceProtocolsOspf":
		return &srlv1alpha1.SrlNetworkinstanceProtocolsOspf{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlNetworkinstanceStaticroutes":
		return &srlv1alpha1.SrlNetworkinstanceStaticroutes{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlSystemName":
		return &srlv1alpha1.SrlSystemName{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlSystemNetworkinstanceProtocolsBgpvpn":
		return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsBgpvpn{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlSystemNetworkinstanceProtocolsEvpn":
		return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpn{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance":
		return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi":
		return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlSystemNtp":
		return &srlv1alpha1.SrlSystemNtp{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlTunnelinterface":
		return &srlv1alpha1.SrlTunnelinterface{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	case "SrlTunnelinterfaceVxlaninterface":
		return &srlv1alpha1.SrlTunnelinterfaceVxlaninterface{
			ObjectMeta: metav1.ObjectMeta{Name: gvk.Name, Namespace: gvk.NameSpace},
		}
	default:
		fmt.Printf("getObject not found gvk: %v\n", *gvk)
		return nil
	}
}
