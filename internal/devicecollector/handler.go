package devicecollector

import (
	"strings"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl/internal/shared"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		log.Debug("resourceList", "list", resourceList)

		// handle deletes
		c.handleDeletes(crDeviceName, resourceList, resp.GetUpdate().Delete)

		c.handleUpdates(crDeviceName, resourceList, resp.GetUpdate().Update)

	case *gnmi.SubscribeResponse_SyncResponse:
		log.Debug("SyncResponse")
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

		c.log.Debug("collector config delete", "path", xpath)

		// clean the path for now to remove the module information from the pathElem
		for _, pe := range path.GetElem() {
			pe.Name = strings.Split(pe.Name, ":")[len(strings.Split(pe.Name, ":"))-1]
		}

		n := &gnmi.Notification{
			Timestamp: time.Now().UnixNano(),
			Prefix:    &gnmi.Path{Target: crDeviceName},
			Delete:    []*gnmi.Path{path},
		}

		// update the cache with the latest config from the device
		if err := c.cache.GnmiUpdate(crDeviceName, n); err != nil {
			//log.Debug("handle target update", "error", err, "Path", yparser.GnmiPath2XPath(u.GetPath(), true), "Value", u.GetVal())
			//log.Debug("handle target update", "error", err, "Notification", *n)
			return errors.New("cache update failed")
		}

		if *resourceName != unmanagedResource {
			// TODO Trigger reconcile event
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
		}

		crDeviceName := shared.GetCrDeviceName(c.namespace, c.target.Config.Name)
		n, err := c.cache.GetNotificationFromUpdate(&gnmi.Path{Target: crDeviceName}, upd)
		if err != nil {
			return err
		}
		for _, u := range n.GetUpdate() {
			c.log.Debug("collector config update", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "value", u.GetVal())
		}

		// update the cache with the latest config from the device
		if err := c.cache.GnmiUpdate(crDeviceName, n); err != nil {
			//log.Debug("handle target update", "error", err, "Path", yparser.GnmiPath2XPath(u.GetPath(), true), "Value", u.GetVal())
			//log.Debug("handle target update", "error", err, "Notification", *n)
			return errors.New("cache update failed")
		}

		if *resourceName != unmanagedResource {
			// TODO Trigger reconcile event
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
