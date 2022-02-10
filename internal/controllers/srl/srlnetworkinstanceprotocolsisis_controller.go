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
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	cevent "sigs.k8s.io/controller-runtime/pkg/event"

	//"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
	"github.com/yndd/nddp-srl/internal/shared"
)

const (
	// Errors
	errUnexpectedNetworkinstanceProtocolsIsis       = "the managed resource is not a NetworkinstanceProtocolsIsis resource"
	errKubeUpdateFailedNetworkinstanceProtocolsIsis = "cannot update NetworkinstanceProtocolsIsis"
	errReadNetworkinstanceProtocolsIsis             = "cannot read NetworkinstanceProtocolsIsis"
	errCreateNetworkinstanceProtocolsIsis           = "cannot create NetworkinstanceProtocolsIsis"
	errUpdateNetworkinstanceProtocolsIsis           = "cannot update NetworkinstanceProtocolsIsis"
	errDeleteNetworkinstanceProtocolsIsis           = "cannot delete NetworkinstanceProtocolsIsis"
)

// SetupNetworkinstanceProtocolsIsis adds a controller that reconciles NetworkinstanceProtocolsIsiss.
func SetupNetworkinstanceProtocolsIsis(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) (string, chan cevent.GenericEvent, error) {
	//func SetupNetworkinstanceProtocolsIsis(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) error {

	name := managed.ControllerName(srlv1alpha1.NetworkinstanceProtocolsIsisGroupKind)

	events := make(chan cevent.GenericEvent)

	y := initYangNetworkinstanceProtocolsIsis(
		nddcopts.DeviceSchema,
	)

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(srlv1alpha1.NetworkinstanceProtocolsIsisGroupVersionKind),
		managed.WithExternalConnecter(&connectorNetworkinstanceProtocolsIsis{
			log:          nddcopts.Logger,
			kube:         mgr.GetClient(),
			usage:        resource.NewNetworkNodeUsageTracker(mgr.GetClient(), &ndrv1.NetworkNodeUsage{}),
			deviceSchema: nddcopts.DeviceSchema,
			nddpSchema:   nddcopts.NddpSchema,
			y:            y,
			newClientFn:  target.NewTarget,
			gnmiAddress:  nddcopts.GnmiAddress},
		),
		managed.WithValidator(&validatorNetworkinstanceProtocolsIsis{
			log:          nddcopts.Logger,
			deviceSchema: nddcopts.DeviceSchema,
			y:            y},
		),
		managed.WithLogger(nddcopts.Logger.WithValues("NetworkinstanceProtocolsIsis-controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	NetworkinstanceProtocolsIsisHandler := &EnqueueRequestForAllNetworkinstanceProtocolsIsis{
		client: mgr.GetClient(),
		log:    nddcopts.Logger,
		ctx:    context.Background(),
	}

	//return ctrl.NewControllerManagedBy(mgr).
	return srlv1alpha1.NetworkinstanceProtocolsIsisGroupKind, events, ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&srlv1alpha1.SrlNetworkinstanceProtocolsIsis{}).
		Owns(&srlv1alpha1.SrlNetworkinstanceProtocolsIsis{}).
		WithEventFilter(resource.IgnoreUpdateWithoutGenerationChangePredicate()).
		/*
			Watches(
				&source.Channel{Source: events},
				&handler.EnqueueRequestForObject{},
			).
		*/
		Watches(&source.Kind{Type: &srlv1alpha1.SrlNetworkinstanceProtocolsIsis{}}, NetworkinstanceProtocolsIsisHandler).
		Watches(&source.Channel{Source: events}, NetworkinstanceProtocolsIsisHandler).
		Complete(r)
}

type NetworkinstanceProtocolsIsis struct {
	*yresource.Resource
}

func initYangNetworkinstanceProtocolsIsis(deviceSchema *yentry.Entry, opts ...yresource.Option) yresource.Handler {
	return &NetworkinstanceProtocolsIsis{&yresource.Resource{
		DeviceSchema: deviceSchema,
	}}

}

// GetRootPath returns the rootpath of the resource
func (r *NetworkinstanceProtocolsIsis) GetRootPath(mg resource.Managed) []*gnmi.Path {

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsIsis)
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
				{Name: "isis"},
			},
		},
	}
}

// GetParentDependency returns the parent dependency of the resource
func (r *NetworkinstanceProtocolsIsis) GetParentDependency(mg resource.Managed) []*leafref.LeafRef {
	rootPath := r.GetRootPath(mg)
	// if the path is not bigger than 1 element there is no parent dependency
	if len(rootPath[0].GetElem()) < 2 {
		return []*leafref.LeafRef{}
	}
	dependencyPath := r.DeviceSchema.GetParentDependency(rootPath[0], rootPath[0], "")
	// the dependency path is the rootPath except for the last element
	//dependencyPathElem := rootPath[0].GetElem()[:(len(rootPath[0].GetElem()) - 1)]
	// check for keys present, if no keys present we return an empty list
	keysPresent := false
	for _, pathElem := range dependencyPath.GetElem() {
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
			RemotePath: dependencyPath,
		},
	}
}

type validatorNetworkinstanceProtocolsIsis struct {
	log          logging.Logger
	deviceSchema *yentry.Entry
	y            yresource.Handler
}

func (v *validatorNetworkinstanceProtocolsIsis) ValidateLeafRef(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateLeafRefObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateLeafRef...")

	// json unmarshal the resource
	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsIsis)
	if !ok {
		return managed.ValidateLeafRefObservation{}, errors.New(errUnexpectedNetworkinstanceProtocolsIsis)
	}
	d, err := json.Marshal(&cr.Spec.NetworkinstanceProtocolsIsis)
	if err != nil {
		return managed.ValidateLeafRefObservation{}, errors.Wrap(err, errJSONMarshal)
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	// json unmarshal the external data
	var x2 interface{}
	json.Unmarshal(cfg, &x2)

	rootPath := v.y.GetRootPath(cr)

	leafRefs := v.deviceSchema.GetLeafRefsLocal(true, rootPath[0], &gnmi.Path{}, make([]*leafref.LeafRef, 0))
	//log.Debug("Validate leafRefs ...", "Path", yparser.GnmiPath2XPath(rootPath[0], false), "leafRefs", leafRefs)
	/*
		for _, leafRef := range leafRefs{
			log.Debug("Validate leafRefs ...",
				"rootPath", yparser.GnmiPath2XPath(rootPath[0], true),
				"localPath", yparser.GnmiPath2XPath(leafRef.LocalPath, true),
				"RemotePath", yparser.GnmiPath2XPath(leafRef.RemotePath, true))
		}
	*/

	// For local external leafref validation we need to supply the external
	// data to validate the remote leafref, we use x2 for this
	success, resultValidation, err := yparser.ValidateLeafRef(
		rootPath[0], x1, x2, leafRefs, v.deviceSchema)
	if err != nil {
		return managed.ValidateLeafRefObservation{
			Success: false,
		}, nil
	}
	if !success {
		for _, r := range resultValidation {
			log.Debug("ValidateLeafRef failed",
				"localPath", yparser.GnmiPath2XPath(r.LeafRef.LocalPath, true),
				"RemotePath", yparser.GnmiPath2XPath(r.LeafRef.RemotePath, true),
				"Resolved", r.Resolved,
				"External", r.External,
				"Value", r.Value,
			)
		}
		log.Debug("Input  Spec  ", "data", x1)
		log.Debug("Latest Config", "data", x2)

		return managed.ValidateLeafRefObservation{
			Success:          false,
			ResolvedLeafRefs: resultValidation}, nil
	}
	/*
		for _, r := range resultValidation {
			log.Debug("ValidateLeafRef success",
				"localPath", yparser.GnmiPath2XPath(r.LeafRef.LocalPath, true),
				"RemotePath", yparser.GnmiPath2XPath(r.LeafRef.RemotePath, true),
				"Resolved", r.Resolved,
				"External", r.External,
				"Value", r.Value,
			)
		}
	*/
	return managed.ValidateLeafRefObservation{
		Success:          true,
		ResolvedLeafRefs: resultValidation}, nil
	/*
		return managed.ValidateLeafRefObservation{
			Success:          true,
			ResolvedLeafRefs: []*leafref.ResolvedLeafRef{}}, nil
	*/
}

func (v *validatorNetworkinstanceProtocolsIsis) ValidateParentDependency(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateParentDependencyObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateParentDependency...")

	dependencyLeafRef := v.y.GetParentDependency(mg)

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
		log.Debug("Latest Config", "data", x1)
		return managed.ValidateParentDependencyObservation{
			Success:          false,
			ResolvedLeafRefs: resultValidation}, nil
	}
	//log.Debug("ValidateParentDependency success", "resultParentValidation", resultValidation)
	return managed.ValidateParentDependencyObservation{
		Success:          true,
		ResolvedLeafRefs: resultValidation}, nil
}

// ValidateResourceIndexes validates if the indexes of a resource got changed
// if so we need to delete the original resource, because it will be dangling if we dont delete it
func (v *validatorNetworkinstanceProtocolsIsis) ValidateResourceIndexes(ctx context.Context, mg resource.Managed) (managed.ValidateResourceIndexesObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateResourceIndexes ...")

	rootPath := v.y.GetRootPath(mg)
	origResourceIndex := mg.GetResourceIndexes()

	// we call the CompareConfigPathsWithResourceKeys irrespective is the get resource index returns nil
	changed, deletPaths, newResourceIndex := yparser.CompareGnmiPathsWithResourceKeys(rootPath[0], origResourceIndex)
	if changed {
		log.Debug("ValidateResourceIndexes changed", "indexes", newResourceIndex, "deletPaths", deletPaths[0])
		return managed.ValidateResourceIndexesObservation{Changed: true, ResourceDeletes: deletPaths, ResourceIndexes: newResourceIndex}, nil
	}

	//log.Debug("ValidateResourceIndexes success", "indexes", newResourceIndex)
	return managed.ValidateResourceIndexesObservation{Changed: false, ResourceIndexes: newResourceIndex}, nil
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connectorNetworkinstanceProtocolsIsis struct {
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
func (c *connectorNetworkinstanceProtocolsIsis) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	log := c.log.WithValues("resource", mg.GetName())
	log.Debug("Connect")

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsIsis)
	if !ok {
		return nil, errors.New(errUnexpectedNetworkinstanceProtocolsIsis)
	}
	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackTCUsage)
	}

	// find network node that is configured status
	nn := &ndrv1.NetworkNode{}
	if err := c.kube.Get(ctx, types.NamespacedName{Name: cr.GetNetworkNodeReference().Name}, nn); err != nil {
		return nil, errors.Wrap(err, errGetNetworkNode)
	}

	if nn.GetCondition(ndrv1.ConditionKindDeviceDriverConfigured).Status != corev1.ConditionTrue {
		return nil, errors.New(targetNotConfigured)
	}

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

	tns := []string{nn.GetName()}

	return &externalNetworkinstanceProtocolsIsis{client: cl, targets: tns, log: log, deviceSchema: c.deviceSchema, nddpSchema: c.nddpSchema, y: c.y}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type externalNetworkinstanceProtocolsIsis struct {
	client       *target.Target
	targets      []string
	log          logging.Logger
	deviceSchema *yentry.Entry
	nddpSchema   *yentry.Entry
	y            yresource.Handler
}

func (e *externalNetworkinstanceProtocolsIsis) Close() {
	e.client.Close()
}

func (e *externalNetworkinstanceProtocolsIsis) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Observing ...")

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsIsis)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedNetworkinstanceProtocolsIsis)
	}

	// rootpath of the resource
	rootPath := e.y.GetRootPath(cr)
	hierElements := e.deviceSchema.GetHierarchicalResourcesLocal(true, rootPath[0], &gnmi.Path{}, make([]*gnmi.Path, 0))
	//log.Debug("Observing hierElements ...", "Path", yparser.GnmiPath2XPath(rootPath[0], false), "hierElements", hierElements)

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
		if er, ok := status.FromError(err); ok {
			switch er.Code() {
			case codes.Unavailable, codes.ResourceExhausted:
				// we use this to signal not ready
				return managed.ExternalObservation{
					Ready:            false,
					ResourceExists:   false,
					ResourceSuccess:  false,
					ResourceHasData:  false,
					ResourceUpToDate: false,
				}, nil
			case codes.NotFound:
				// the k8s resource does not exists but the data can still exist
				// if data exists it means we go from UMR -> MR
				//log.Debug("observing when using gnmic: resource does not exist")
				exists = false
			case codes.FailedPrecondition:
				// the k8s resource exists but is in failed status, compare the response spec with current spec
				// if the specs are equal return observation.ResponseSuccess -> False
				// if the specs are not equal follow the regular procedure
				//log.Debug("observing when using gnmic: resource failed")
				failedObserve, err := processObserve(rootPath[0], hierElements, &cr.Spec, resp, e.deviceSchema)
				if err != nil {
					return managed.ExternalObservation{}, err
				}
				if !failedObserve.delta {
					// there is no difference between the previous spec and the current spec, so we dont retry
					// given the previous attempt failed
					return managed.ExternalObservation{
						Ready:            true,
						ResourceExists:   true,
						ResourceSuccess:  false,
						ResourceHasData:  false,
						ResourceUpToDate: false,
					}, nil
				} else {
					// this should trigger an update
					return managed.ExternalObservation{
						Ready:            true,
						ResourceExists:   true,
						ResourceSuccess:  true,
						ResourceHasData:  true,
						ResourceUpToDate: false,
					}, nil
				}
			}
		} else {
			// WORKAROUND WAITING FOR KARIM TO REMOVE THE ERROR WRAP In GNMIC
			switch {
			case strings.Contains(err.Error(), "Unavailable"):
				// we use this to signal not ready
				return managed.ExternalObservation{
					Ready:            false,
					ResourceExists:   false,
					ResourceSuccess:  false,
					ResourceHasData:  false,
					ResourceUpToDate: false,
				}, nil
			case strings.Contains(err.Error(), "NotFound"):
				//log.Debug("observing: resource does not exist")
				exists = false
			case strings.Contains(err.Error(), "Failed"):
				//log.Debug("observing: resource failed")
				// the k8s resource exists but is in failed status, compare the response spec with current spec
				// if the specs are equal return observation.ResponseSuccess -> False
				// if the specs are not equal follow the regular procedure
				failedObserve, err := processObserve(rootPath[0], hierElements, &cr.Spec, resp, e.deviceSchema)
				if err != nil {
					return managed.ExternalObservation{}, err
				}
				if !failedObserve.delta {
					// there is no difference between the previous spec and the current spec, so we dont retry
					// given the previous attempt failed
					return managed.ExternalObservation{
						Ready:            true,
						ResourceExists:   true,
						ResourceSuccess:  false,
						ResourceHasData:  false,
						ResourceUpToDate: false,
					}, nil
				} else {
					// this should trigger an update
					return managed.ExternalObservation{
						Ready:            true,
						ResourceExists:   true,
						ResourceSuccess:  true,
						ResourceHasData:  true,
						ResourceUpToDate: false,
					}, nil
				}
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
	//e.log.Debug("processObserve", "notification", resp.GetNotification())
	observe, err := processObserve(rootPath[0], hierElements, &cr.Spec, resp, e.deviceSchema)
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	if !observe.hasData {
		// No Data exists -> Create it or Delete is complete
		//log.Debug("Observing Response:", "observe", observe, "exists", exists, "Response", resp)
		return managed.ExternalObservation{
			Ready:            true,
			ResourceExists:   exists,
			ResourceSuccess:  true,
			ResourceHasData:  false,
			ResourceUpToDate: false,
		}, nil
	}
	// Data exists

	if observe.delta {
		// resource is NOT up to date
		log.Debug("Observing Response: resource NOT up to date", "Observe", observe, "exists", exists, "Response", resp)
		return managed.ExternalObservation{
			Ready:            true,
			ResourceExists:   exists,
			ResourceSuccess:  true,
			ResourceHasData:  true,
			ResourceUpToDate: false,
			//ResourceDeletes:  observe.deletes,
			//ResourceUpdates:  observe.updates,
		}, nil
	}
	// resource is up to date
	//log.Debug("Observing Response: resource up to date", "Observe", observe, "Response", resp)
	return managed.ExternalObservation{
		Ready:            true,
		ResourceExists:   exists,
		ResourceSuccess:  true,
		ResourceHasData:  true,
		ResourceUpToDate: true,
	}, nil
}

func (e *externalNetworkinstanceProtocolsIsis) Create(ctx context.Context, mg resource.Managed) error {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Creating ...")

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsIsis)
	if !ok {
		return errors.New(errUnexpectedNetworkinstanceProtocolsIsis)
	}

	// get the rootpath of the resource
	rootPath := e.y.GetRootPath(mg)

	// create k8s object
	// processCreate
	// 0. marshal/unmarshal data
	// 1. transform the spec data to gnmi updates
	updates, err := processCreateK8s(mg, rootPath[0], &cr.Spec, e.deviceSchema, e.nddpSchema)
	if err != nil {
		return errors.Wrap(err, errCreateObject)
	}
	/*
		for _, update := range updates {
			log.Debug("Create Fine Grane Updates", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
		}

		if len(updates) == 0 {
			log.Debug("cannot create object since there are no updates present")
			return errors.Wrap(err, errCreateObject)
		}
	*/

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := &gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, req)
	if err != nil {
		return errors.Wrap(err, errCreateInterfaceSubinterface)
	}

	return nil
}

func (e *externalNetworkinstanceProtocolsIsis) Update(ctx context.Context, mg resource.Managed, obs managed.ExternalObservation) error {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Updating ...")

	cr, ok := mg.(*srlv1alpha1.SrlNetworkinstanceProtocolsIsis)
	if !ok {
		return errors.New(errUnexpectedNetworkinstanceProtocolsIsis)
	}

	// get the rootpath of the resource
	rootPath := e.y.GetRootPath(mg)

	updates, err := processUpdateK8s(mg, rootPath[0], &cr.Spec, e.deviceSchema, e.nddpSchema)
	if err != nil {
		return errors.Wrap(err, errUpdateInterfaceSubinterface)
	}
	/*
		for _, update := range updates {
			log.Debug("Update Fine Grane Updates", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
		}
	*/

	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	req := gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, &req)
	if err != nil {
		return errors.Wrap(err, errUpdateInterfaceSubinterface)
	}

	return nil
}

func (e *externalNetworkinstanceProtocolsIsis) Delete(ctx context.Context, mg resource.Managed) error {
	log := e.log.WithValues("Resource", mg.GetName())
	log.Debug("Deleting ...")

	// get the rootpath of the resource
	rootPath := e.y.GetRootPath(mg)

	updates, err := processDeleteK8sResource(mg, rootPath[0], e.nddpSchema)
	if err != nil {
		return errors.Wrap(err, errDeleteInterfaceSubinterface)
	}
	/*
		for _, update := range updates {
			log.Debug("Delete Fine Grane Updates", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
		}
	*/

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

func (e *externalNetworkinstanceProtocolsIsis) GetTarget() []string {
	return e.targets
}

func (e *externalNetworkinstanceProtocolsIsis) GetConfig(ctx context.Context, mg resource.Managed) ([]byte, error) {
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
	//e.log.Debug("Get Config Empty response")
	return nil, nil
}

func (e *externalNetworkinstanceProtocolsIsis) GetResourceName(ctx context.Context, mg resource.Managed, path *gnmi.Path) (string, error) {
	//e.log.Debug("Get GetResourceName ...", "remotePath", yparser.GnmiPath2XPath(path, true))
	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(mg.GetNamespace(), mg.GetNetworkNodeReference().Name))

	// gnmi get request
	req := &gnmi.GetRequest{
		Prefix:   &gnmi.Path{Target: crSystemDeviceName},
		Path:     []*gnmi.Path{path},
		Encoding: gnmi.Encoding_JSON,
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gvkresource.Operation_GetResourceNameFromPath)}}},
		},
	}

	// gnmi get response
	resp, err := e.client.Get(ctx, req)
	if err != nil {
		return "", errors.Wrap(err, errGetResourceName)
	}

	x2, err := yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
	if err != nil {
		return "", errors.Wrap(err, errJSONMarshal)
	}

	d, err := json.Marshal(x2)
	if err != nil {
		return "", errors.Wrap(err, errJSONMarshal)
	}

	var resourceName nddv1.ResourceName
	if err := json.Unmarshal(d, &resourceName); err != nil {
		return "", errors.Wrap(err, errJSONUnMarshal)
	}

	e.log.Debug("Get ResourceName Response", "remotePath", yparser.GnmiPath2XPath(path, true), "ResourceName", resourceName)

	return resourceName.Name, nil
}
