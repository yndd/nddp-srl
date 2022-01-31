/*
Copyright 2022 NDD.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package srl

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/karimra/gnmic/target"
	gnmitypes "github.com/karimra/gnmic/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/pkg/errors"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	"github.com/yndd/ndd-runtime/pkg/event"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/reconciler/managed"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-runtime/pkg/utils"
	"github.com/yndd/ndd-yang/pkg/leafref"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/ndd-yang/pkg/yresource"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	cevent "sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
	"github.com/yndd/nddp-srl/internal/shared"
)

const (
	// Errors
	errUnexpectedNetworkinstanceProtocolsBgpevpn       = "the managed resource is not a NetworkinstanceProtocolsBgpevpn resource"
	errKubeUpdateFailedNetworkinstanceProtocolsBgpevpn = "cannot update NetworkinstanceProtocolsBgpevpn"
	errReadNetworkinstanceProtocolsBgpevpn             = "cannot read NetworkinstanceProtocolsBgpevpn"
	errCreateNetworkinstanceProtocolsBgpevpn           = "cannot create NetworkinstanceProtocolsBgpevpn"
	errUpdateNetworkinstanceProtocolsBgpevpn           = "cannot update NetworkinstanceProtocolsBgpevpn"
	errDeleteNetworkinstanceProtocolsBgpevpn           = "cannot delete NetworkinstanceProtocolsBgpevpn"
)

// SetupNetworkinstanceProtocolsBgpevpn adds a controller that reconciles NetworkinstanceProtocolsBgpevpns.
//func SetupInterface(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) (string, chan cevent.GenericEvent, error) {
func SetupNetworkinstanceProtocolsBgpevpn(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) error {

	name := managed.ControllerName(srlv1alpha1.NetworkinstanceProtocolsBgpevpnGroupKind)

	events := make(chan cevent.GenericEvent)

	y := initYangNetworkinstanceProtocolsBgpevpn()

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(srlv1alpha1.NetworkinstanceProtocolsBgpevpnGroupVersionKind),
		managed.WithExternalConnecter(&connectorNetworkinstanceProtocolsBgpevpn{
			log:          nddcopts.Logger,
			kube:         mgr.GetClient(),
			usage:        resource.NewNetworkNodeUsageTracker(mgr.GetClient(), &ndrv1.NetworkNodeUsage{}),
			deviceSchema: nddcopts.DeviceSchema,
			nddpSchema:   nddcopts.NddpSchema,
			y:            y,
			newClientFn:  target.NewTarget,
			gnmiAddress:  nddcopts.GnmiAddress},
		),
		managed.WithValidator(&validatorNetworkinstanceProtocolsBgpevpn{
			log:          nddcopts.Logger,
			deviceSchema: nddcopts.DeviceSchema,
			y:            y},
		),
		managed.WithLogger(nddcopts.Logger.WithValues("NetworkinstanceProtocolsBgpevpn-controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	//return srlv1alpha1.NetworkinstanceProtocolsBgpevpnGroupKind, events, ctrl.NewControllerManagedBy(mgr).
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn{}).
		Owns(&srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn{}).
		WithEventFilter(resource.IgnoreUpdateWithoutGenerationChangePredicate()).
		Watches(
			&source.Channel{Source: events},
			&handler.EnqueueRequestForObject{},
		).
		Complete(r)
}

type NetworkinstanceProtocolsBgpevpn struct {
	*yresource.Resource
}

func initYangNetworkinstanceProtocolsBgpevpn(opts ...yresource.Option) yresource.Handler {
	return &NetworkinstanceProtocolsBgpevpn{&yresource.Resource{}}

}

// GetRootPath returns the rootpath of the resource
func (r *NetworkinstanceProtocolsBgpevpn) GetRootPath(mg resource.Managed) []*gnmi.Path {

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn)
	if !ok {
		fmt.Printf("wrong cr: %v\n", cr)
		return []*gnmi.Path{}
	}

	return []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "protocols"},
				{Name: "bgp-evpn"},
			},
		},
	}
}

// GetParentDependency returns the parent dependency of the resource
func (r *NetworkinstanceProtocolsBgpevpn) GetParentDependency(mg resource.Managed) []*leafref.LeafRef {
	rootPath := r.GetRootPath(mg)
	// if the path is not bigger than 1 element there is no parent dependency
	if len(rootPath[0].GetElem()) < 2 {
		return []*leafref.LeafRef{}
	}
	// the dependency path is the rootPath except for the last element
	dependencyPathElem := rootPath[0].GetElem()[:(len(rootPath[0].GetElem()) - 1)]
	// check for keys present, if no keys present we return an empty list
	keysPresent := false
	for _, pathElem := range dependencyPathElem {
		if len(pathElem.GetKey()) != 0 {
			keysPresent = true
		}
	}
	if !keysPresent {
		return []*leafref.LeafRef{}
	}

	// return the rootPath except the last entry
	return []*leafref.LeafRef{
		{
			RemotePath: &gnmi.Path{Elem: dependencyPathElem},
		},
	}
}

type validatorNetworkinstanceProtocolsBgpevpn struct {
	log          logging.Logger
	deviceSchema *yentry.Entry
	y            yresource.Handler
}

func (v *validatorNetworkinstanceProtocolsBgpevpn) ValidateLocalleafRef(ctx context.Context, mg resource.Managed) (managed.ValidateLocalleafRefObservation, error) {
	return managed.ValidateLocalleafRefObservation{
		Success:          true,
		ResolvedLeafRefs: []*leafref.ResolvedLeafRef{}}, nil
}

func (v *validatorNetworkinstanceProtocolsBgpevpn) ValidateExternalleafRef(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateExternalleafRefObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateExternalleafRef...")

	// json unmarshal the resource
	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn)
	if !ok {
		return managed.ValidateExternalleafRefObservation{}, errors.New(errUnexpectedNetworkinstanceProtocolsBgpevpn)
	}
	d, err := json.Marshal(&cr.Spec.NetworkinstanceProtocolsBgpevpn)
	if err != nil {
		return managed.ValidateExternalleafRefObservation{}, errors.Wrap(err, errJSONMarshal)
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	// json unmarshal the external data
	var x2 interface{}
	json.Unmarshal(cfg, &x2)

	rootPath := v.y.GetRootPath(cr)

	leafRefs := v.deviceSchema.GetLeafRefsLocal(true, rootPath[0], &gnmi.Path{}, make([]*leafref.LeafRef, 0))
	log.Debug("Validate leafRefs ...", "Path", yparser.GnmiPath2XPath(rootPath[0], false), "leafRefs", leafRefs)

	// For local external leafref validation we need to supply the external
	// data to validate the remote leafref, we use x2 for this
	success, resultValidation, err := yparser.ValidateLeafRef(
		rootPath[0], x1, x2, leafRefs, v.deviceSchema)
	if err != nil {
		return managed.ValidateExternalleafRefObservation{
			Success: false,
		}, nil
	}
	if !success {
		for _, r := range resultValidation {
			log.Debug("ValidateExternalleafRef failed",
				"localPath", yparser.GnmiPath2XPath(r.LeafRef.LocalPath, true),
				"RemotePath", yparser.GnmiPath2XPath(r.LeafRef.RemotePath, true),
				"Resolved", r.Resolved,
				"Value", r.Value,
			)
		}
		return managed.ValidateExternalleafRefObservation{
			Success:          false,
			ResolvedLeafRefs: resultValidation}, nil
	}
	for _, r := range resultValidation {
		log.Debug("ValidateExternalleafRef success",
			"localPath", yparser.GnmiPath2XPath(r.LeafRef.LocalPath, true),
			"RemotePath", yparser.GnmiPath2XPath(r.LeafRef.RemotePath, true),
			"Resolved", r.Resolved,
			"Value", r.Value,
		)
	}
	return managed.ValidateExternalleafRefObservation{
		Success:          true,
		ResolvedLeafRefs: resultValidation}, nil
}

func (v *validatorNetworkinstanceProtocolsBgpevpn) ValidateParentDependency(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateParentDependencyObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateParentDependency...")

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn)
	if !ok {
		return managed.ValidateParentDependencyObservation{}, errors.New(errUnexpectedNetworkinstanceProtocolsBgpevpn)
	}

	dependencyLeafRef := v.y.GetParentDependency(cr)

	// unmarshal the config
	var x1 interface{}
	json.Unmarshal(cfg, &x1)
	//log.Debug("Latest Config", "data", x1)

	success, resultValidation, err := yparser.ValidateParentDependency(
		x1, dependencyLeafRef, v.deviceSchema)
	if err != nil {
		return managed.ValidateParentDependencyObservation{
			Success: false,
		}, nil
	}
	if !success {
		log.Debug("ValidateParentDependency failed", "resultParentValidation", resultValidation)
		return managed.ValidateParentDependencyObservation{
			Success:          false,
			ResolvedLeafRefs: resultValidation}, nil
	}
	log.Debug("ValidateParentDependency success", "resultParentValidation", resultValidation)
	return managed.ValidateParentDependencyObservation{
		Success:          true,
		ResolvedLeafRefs: resultValidation}, nil
}

// ValidateResourceIndexes validates if the indexes of a resource got changed
// if so we need to delete the original resource, because it will be dangling if we dont delete it
func (v *validatorNetworkinstanceProtocolsBgpevpn) ValidateResourceIndexes(ctx context.Context, mg resource.Managed) (managed.ValidateResourceIndexesObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn)
	if !ok {
		return managed.ValidateResourceIndexesObservation{}, errors.New(errUnexpectedNetworkinstanceProtocolsBgpevpn)
	}

	log.Debug("ValidateResourceIndexes", "Spec", cr.Spec)

	return managed.ValidateResourceIndexesObservation{Changed: false, ResourceIndexes: map[string]string{}}, nil
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connectorNetworkinstanceProtocolsBgpevpn struct {
	log          logging.Logger
	kube         client.Client
	usage        resource.Tracker
	deviceSchema *yentry.Entry
	nddpSchema   *yentry.Entry
	y            yresource.Handler
	newClientFn  func(c *gnmitypes.TargetConfig) *target.Target
	gnmiAddress  string
}

// Connect produces an ExternalClient by:
// 1. Tracking that the managed resource is using a NetworkNode.
// 2. Getting the managed resource's NetworkNode with connection details
// A resource is mapped to a single target
func (c *connectorNetworkinstanceProtocolsBgpevpn) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	log := c.log.WithValues("resource", mg.GetName())
	log.Debug("Connect")
	cfg := &gnmitypes.TargetConfig{
		Name:       "dummy",
		Address:    c.gnmiAddress,
		Username:   utils.StringPtr("admin"),
		Password:   utils.StringPtr("admin"),
		Timeout:    10 * time.Second,
		SkipVerify: utils.BoolPtr(true),
		Insecure:   utils.BoolPtr(true),
		TLSCA:      utils.StringPtr(""), //TODO TLS
		TLSCert:    utils.StringPtr(""), //TODO TLS
		TLSKey:     utils.StringPtr(""),
		Gzip:       utils.BoolPtr(false),
	}

	cl := target.NewTarget(cfg)
	if err := cl.CreateGNMIClient(ctx); err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	tns := []string{"localGNMIServer"}

	return &externalNetworkinstanceProtocolsBgpevpn{client: cl, targets: tns, log: log, deviceSchema: c.deviceSchema, nddpSchema: c.nddpSchema, y: c.y}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type externalNetworkinstanceProtocolsBgpevpn struct {
	client       *target.Target
	targets      []string
	log          logging.Logger
	deviceSchema *yentry.Entry
	nddpSchema   *yentry.Entry
	y            yresource.Handler
}

func (e *externalNetworkinstanceProtocolsBgpevpn) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Observing ...")

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedNetworkinstanceProtocolsBgpevpn)
	}

	// rootpath of the resource
	rootPath := e.y.GetRootPath(cr)
	hierElements := e.deviceSchema.GetHierarchicalResourcesLocal(true, rootPath[0], &gnmi.Path{}, make([]*gnmi.Path, 0))
	log.Debug("Observing hierElements ...", "Path", yparser.GnmiPath2XPath(rootPath[0], false), "hierElements", hierElements)

	gvkName := gvkresource.GetGvkName(mg)

	// gnmi get request
	req := &gnmi.GetRequest{
		//Prefix:   &gnmi.Path{Target: GnmiTarget, Origin: GnmiOrigin},
		Prefix:   &gnmi.Path{Target: shared.GetCrDeviceName(mg.GetNamespace(), cr.GetNetworkNodeReference().Name)},
		Path:     rootPath,
		Encoding: gnmi.Encoding_JSON,
		//Type:     gnmi.GetRequest_DataType(gnmi.GetRequest_STATE),
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gvkName)}}},
		},
	}

	// gnmi get response
	exists := true
	resp, err := e.client.Get(ctx, req)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Unavailable, codes.ResourceExhausted:
				// we use this to signal not ready
				return managed.ExternalObservation{
					Ready:            false,
					ResourceExists:   false,
					ResourceHasData:  false,
					ResourceUpToDate: false,
				}, nil
			case codes.NotFound:
				// the k8s resource does not exists but the data can still exist
				// if data exists it means we go from UMR -> MR
				log.Debug("observing when using gnmic: resource does not exist")
				exists = false
			}
		} else {
			// WORKAROUND WAITING FOR KARIM TO REMOVE THE ERROR WRAP In GNMIC
			switch {
			case strings.Contains(err.Error(), "Unavailable"):
				// we use this to signal not ready
				return managed.ExternalObservation{
					Ready:            false,
					ResourceExists:   false,
					ResourceHasData:  false,
					ResourceUpToDate: false,
				}, nil
			case strings.Contains(err.Error(), "NotFound"):
				log.Debug("observing: resource does not exist")
				exists = false
			default:
				return managed.ExternalObservation{}, errors.Wrap(err, errReadInterfaceSubinterface)
			}
		}
	}

	// processObserve
	// o. marshal/unmarshal data
	// 1. check if resource exists
	// 2. remove parent hierarchical elements from spec
	// 3. remove resource hierarchical elements from gnmi response
	// 4. remove state
	// 5. transform the data in gnmi to process the delta
	// 6. find the resource delta: updates and/or deletes in gnmi
	//exists, deletes, updates, b, err := processObserve(rootPath[0], hierElements, &cr.Spec, resp, e.deviceSchema)
	e.log.Debug("processObserve", "notification", resp.GetNotification())
	observe, err := processObserve(rootPath[0], hierElements, &cr.Spec, resp, e.deviceSchema)
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	if !observe.hasData {
		// No Data exists -> Create it or Delete is complete
		log.Debug("Observing Response:", "observe", observe, "exists", exists, "Response", resp)
		return managed.ExternalObservation{
			Ready:            true,
			ResourceExists:   exists,
			ResourceHasData:  false,
			ResourceUpToDate: false,
		}, nil
	}
	// Data exists

	if len(observe.deletes) != 0 || len(observe.updates) != 0 {
		// resource is NOT up to date
		log.Debug("Observing Response: resource NOT up to date", "Observe", observe, "exists", exists, "Response", resp)
		for _, del := range observe.deletes {
			log.Debug("Observing Response: resource NOT up to date, deletes", "path", yparser.GnmiPath2XPath(del, true))
		}
		for _, upd := range observe.updates {
			val, _ := yparser.GetValue(upd.GetVal())
			log.Debug("Observing Response: resource NOT up to date, updates", "path", yparser.GnmiPath2XPath(upd.GetPath(), true), "data", val)
		}
		return managed.ExternalObservation{
			Ready:            true,
			ResourceExists:   exists,
			ResourceHasData:  true,
			ResourceUpToDate: false,
			ResourceDeletes:  observe.deletes,
			ResourceUpdates:  observe.updates,
		}, nil
	}
	// resource is up to date
	log.Debug("Observing Response: resource up to date", "Observe", observe, "Response", resp)
	return managed.ExternalObservation{
		Ready:            true,
		ResourceExists:   exists,
		ResourceHasData:  true,
		ResourceUpToDate: true,
	}, nil
}

func (e *externalNetworkinstanceProtocolsBgpevpn) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Creating ...")

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedNetworkinstanceProtocolsBgpevpn)
	}

	// get the rootpath of the resource
	rootPath := e.y.GetRootPath(mg)

	// create k8s object
	// processCreate
	// 0. marshal/unmarshal data
	// 1. transform the spec data to gnmi updates
	updates, err := processCreateK8s(mg, rootPath[0], &cr.Spec, e.deviceSchema, e.nddpSchema)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateObject)
	}
	for _, update := range updates {
		log.Debug("Create Fine Grane Updates", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
	}

	if len(updates) == 0 {
		log.Debug("cannot create object since there are no updates present")
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateObject)
	}

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := &gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, req)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateInterfaceSubinterface)
	}

	return managed.ExternalCreation{}, nil
}

func (e *externalNetworkinstanceProtocolsBgpevpn) Update(ctx context.Context, mg resource.Managed, obs managed.ExternalObservation) (managed.ExternalUpdate, error) {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Updating ...")

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedNetworkinstanceProtocolsBgpevpn)
	}

	// get the rootpath of the resource
	rootPath := e.y.GetRootPath(mg)

	updates, err := processUpdateK8s(mg, rootPath[0], &cr.Spec, e.deviceSchema, e.nddpSchema)
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errUpdateInterfaceSubinterface)
	}
	for _, update := range updates {
		log.Debug("Update Fine Grane Updates", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
	}

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, &req)
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errUpdateInterfaceSubinterface)
	}

	return managed.ExternalUpdate{}, nil
}

func (e *externalNetworkinstanceProtocolsBgpevpn) Delete(ctx context.Context, mg resource.Managed) error {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Deleting ...")

	// get the rootpath of the resource
	rootPath := e.y.GetRootPath(mg)

	updates, err := processDeleteK8sResource(mg, rootPath[0], e.nddpSchema)
	if err != nil {
		return errors.Wrap(err, errDeleteInterfaceSubinterface)
	}
	for _, update := range updates {
		log.Debug("Delete Fine Grane Updates", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
	}

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, &req)
	if err != nil {
		return errors.Wrap(err, errDeleteInterfaceSubinterface)
	}

	return nil
}

func (e *externalNetworkinstanceProtocolsBgpevpn) GetTarget() []string {
	return e.targets
}

func (e *externalNetworkinstanceProtocolsBgpevpn) GetConfig(ctx context.Context, mg resource.Managed) ([]byte, error) {
	e.log.Debug("Get Config ...")
	req := &gnmi.GetRequest{
		Prefix: &gnmi.Path{Target: shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name)},
		Path: []*gnmi.Path{
			{
				Elem: []*gnmi.PathElem{},
			},
		},
		Encoding: gnmi.Encoding_JSON,
		//Type:     gnmi.GetRequest_DataType(gnmi.GetRequest_CONFIG),
	}

	resp, err := e.client.Get(ctx, req)
	if err != nil {
		return make([]byte, 0), errors.Wrap(err, errGetConfig)
	}

	if len(resp.GetNotification()) != 0 {
		if len(resp.GetNotification()[0].GetUpdate()) != 0 {
			x2, err := yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
			if err != nil {
				return make([]byte, 0), errors.Wrap(err, errGetConfig)
			}

			data, err := json.Marshal(x2)
			if err != nil {
				return make([]byte, 0), errors.Wrap(err, errJSONMarshal)
			}
			return data, nil
		}
	}
	e.log.Debug("Get Config Empty response")
	return nil, nil
}

func (e *externalNetworkinstanceProtocolsBgpevpn) GetResourceName(ctx context.Context, path []*gnmi.Path) (string, error) {
	return "", nil
}
