/*
Copyright 2021 NDD.

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

package transaction

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/karimra/gnmic/target"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/event"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/utils"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"github.com/yndd/nddp-system/pkg/transaction"

	"github.com/yndd/nddp-srl/internal/shared"
	"github.com/yndd/nddp-srl/internal/srlhandler"
	"k8s.io/apimachinery/pkg/types"

	//"github.com/yndd/ndd-runtime/pkg/meta"
	"github.com/yndd/ndd-runtime/pkg/resource"
	//corev1 "k8s.io/api/core/v1"
	gnmitypes "github.com/karimra/gnmic/types"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	// Finalizer
	finalizer = "transaction.srl.ndda.yndd.io"

	// default
	//defaultGrpcPort = 9999

	// Timers
	defaultTimeout   = 5 * time.Second
	reconcileTimeout = 15 * time.Minute
	shortWait        = 30 * time.Second
	veryShortWait    = 1 * time.Second
	longWait         = 1 * time.Minute

	// Errors
	errGetTransaction        = "cannot get transaction"
	errGetManaged            = "cannot get managed resource"
	errUpdateStatus          = "cannot update transaction status"
	errGetTransactionUpdates = "cannot get transaction updates"

	errAddFinalizer    = "cannot add transaction finalizer"
	errRemoveFinalizer = "cannot remove transaction finalizer"

	errJSONMarshal   = "cannot marshal JSON object"
	errJSONUnMarshal = "cannot unmarshal JSON object"

	errGetNetworkNode = "cannot get NetworkNode"
	errDeviceNotReady = "networkNode is not ready"
	errNewGnmiClient  = "cannot create new gnmi client"

	// Event reasons
	reasonSync event.Reason = "SyncTransaction"
)

// ReconcilerOption is used to configure the Reconciler.
type ReconcilerOption func(*Reconciler)

// WithLogger specifies how the Reconciler should log messages.
func WithLogger(log logging.Logger) ReconcilerOption {
	return func(r *Reconciler) {
		r.log = log
	}
}

// WithRecorder specifies how the Reconciler should record Kubernetes events.
func WithRecorder(er event.Recorder) ReconcilerOption {
	return func(r *Reconciler) {
		r.record = er
	}
}

// WithNewTransactionFn determines the transaction being reconciled.
func WithNewTransactionFn(f func() srlv1alpha1.IFSrlTransaction) ReconcilerOption {
	return func(r *Reconciler) {
		r.newTransaction = f
	}
}

func WithSrlHandler(h srlhandler.Handler) ReconcilerOption {
	return func(r *Reconciler) {
		r.srlHandler = h
	}
}

func WithNddpSchema(nddpSchema *yentry.Entry) ReconcilerOption {
	return func(r *Reconciler) {
		r.nddpSchema = nddpSchema
	}
}

func WithGnmiAddress(a string) ReconcilerOption {
	return func(r *Reconciler) {
		r.gnmiAddress = a
	}
}

// Reconciler reconciles packages.
type Reconciler struct {
	client      resource.ClientApplicator
	tnFinalizer resource.Finalizer
	log         logging.Logger
	record      event.Recorder

	newTransaction func() srlv1alpha1.IFSrlTransaction

	srlHandler srlhandler.Handler
	//deviceSchema *yentry.Entry
	gnmiAddress string
	nddpSchema  *yentry.Entry
}

// Setup adds a controller that reconciles the Lock.
func Setup(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) error {
	name := "dvr/" + strings.ToLower(srlv1alpha1.TransactionKindKind)
	tn := func() srlv1alpha1.IFSrlTransaction { return &srlv1alpha1.SrlTransaction{} }

	r := NewReconciler(mgr,
		WithLogger(nddcopts.Logger.WithValues("controller", name)),
		WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		WithNewTransactionFn(tn),
		WithSrlHandler(srlhandler.New(
			srlhandler.WithClient(resource.ClientApplicator{
				Client:     mgr.GetClient(),
				Applicator: resource.NewAPIPatchingApplicator(mgr.GetClient()),
			}),
			srlhandler.WithLogger(nddcopts.Logger.WithValues("srlhandler", name))),
		),
		//WithDeviceSchema(nddcopts.DeviceSchema),
		WithGnmiAddress(nddcopts.GnmiAddress),
		WithNddpSchema(nddcopts.NddpSchema),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&srlv1alpha1.SrlTransaction{}).
		Owns(&srlv1alpha1.SrlTransaction{}).
		WithEventFilter(resource.IgnoreUpdateWithoutGenerationChangePredicate()).
		Complete(r)
}

// NewReconciler creates a new package revision reconciler.
func NewReconciler(mgr manager.Manager, opts ...ReconcilerOption) *Reconciler {
	r := &Reconciler{
		client: resource.ClientApplicator{
			Client:     mgr.GetClient(),
			Applicator: resource.NewAPIPatchingApplicator(mgr.GetClient()),
		},
		tnFinalizer: resource.NewAPIFinalizer(mgr.GetClient(), finalizer),
		log:         logging.NewNopLogger(),
		record:      event.NewNopRecorder(),
	}

	for _, f := range opts {
		f(r)
	}

	return r
}

// Reconcile network node.
func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) { // nolint:gocyclo
	log := r.log.WithValues("request", req)
	log.Debug("Transaction", "NameSpace", req.NamespacedName)

	t := r.newTransaction()
	if err := r.client.Get(ctx, req.NamespacedName, t); err != nil {
		// There's no need to requeue if we no longer exist. Otherwise we'll be
		// requeued implicitly because we return an error.
		log.Debug(errGetTransaction, "error", err)
		return reconcile.Result{}, errors.Wrap(resource.IgnoreNotFound(err), errGetTransaction)
	}

	resources, err := r.srlHandler.ListResourcesByTransaction(ctx, t)
	if err != nil {
		return reconcile.Result{RequeueAfter: veryShortWait}, err
	}

	log.Debug("Transaction", "resources", resources)

	// validate if the transaction is complete in the cache
	// we walk over the system resource cache and if the transaction macthes we delete the entries from the k8s api resourcelist
	// if the transaction is complete we should have no entries any longer in the k8s resourceList
	allDeviceTransactionsComplete := true
	gvkList := make([]string, 0)
	for deviceName, deviceCrs := range resources {
		// create a gnmi client per device to interact with the cache
		cl, err := r.getGnmiClient(ctx, deviceName)
		if err != nil {
			return reconcile.Result{RequeueAfter: shortWait}, err
		}
		defer cl.Close()

		/*
			for deviceResourceKind, deviceResourceByKind := range deviceResources {
				for deviceResourceName := range deviceResourceByKind {
					r.log.Debug("transaction resource list", "transaction", t.GetName(), "deviceName", deviceName, "kind", deviceResourceKind, "name", deviceResourceName)
				}
			}
		*/

		// get the resourceList from the system cache
		gvkResourceList, err := r.getDeviceGvkList(ctx, cl, t.GetNamespace(), deviceName)
		if err != nil {
			return reconcile.Result{RequeueAfter: shortWait}, err
		}

		// check if the transaction list is complete
		// loop over the gvk list from the cache and check if the resource exists; if so delete it from the k8s transaction resourceList
		for _, gvkResource := range gvkResourceList {
			gvk, err := gvkresource.String2Gvk(gvkResource.Name)
			if err != nil {
				return reconcile.Result{RequeueAfter: shortWait}, err
			}
			if gvkResource.Transaction == t.GetName() && gvkResource.Transactiongeneration == t.GetOwnerGeneration() && gvk.NameSpace == t.GetNamespace() {
				// check the kind
				if deviceCrNames, ok := deviceCrs[gvk.Kind]; ok {
					delete(deviceCrNames, gvk.Name)
					gvkList = append(gvkList, gvkResource.Name)
				}
			}
		}

		transactionComplete := true
		// when the transaction list is complete all the deviceCrPerKind entries should be 0
		for _, deviceCrPerKind := range deviceCrs {
			if len(deviceCrPerKind) != 0 {
				transactionComplete = false
				allDeviceTransactionsComplete = false
			}
		}

		if transactionComplete {
			// set the status in the system cache to true
			r.log.Debug("transaction COMPLETE", "device", deviceName)

			crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(t.GetNamespace(), deviceName))

			updates, err := transaction.ProcessTransaction(
				&systemv1alpha1.Transaction{
					Action:     systemv1alpha1.E_TransactionAction_Create,
					Generation: t.GetOwnerGeneration(),
					Name:       t.GetName(),
					Gvk:        gvkList,
					Status:     systemv1alpha1.E_TransactionStatus_Pending,
				}, r.nddpSchema,
			)

			req := &gnmi.SetRequest{
				Prefix:  &gnmi.Path{Target: crSystemDeviceName},
				Replace: updates,
			}

			_, err = cl.Set(ctx, req)
			if err != nil {
				return reconcile.Result{}, err
			}
		}
	}

	/*
		updates, err := r.getTransactionUpdates(ctx, t, resources)
		if err != nil {
			log.Debug(errGetTransactionUpdates, "error", err)
			r.record.Event(t, event.Warning(reasonSync, errors.Wrap(err, errGetTransactionUpdates)))
			return reconcile.Result{RequeueAfter: shortWait}, err
		}

		for _, u := range updates {
			fmt.Printf("transaction update: path: %s, value: %v\n", yparser.GnmiPath2XPath(u.GetPath(), true), u.GetVal())
		}
	*/
	// Add a finalizer to newly created objects and update the conditions
	/*
		if err := r.tnFinalizer.AddFinalizer(ctx, t); err != nil {
			log.Debug(errAddFinalizer, "error", err)
			r.record.Event(t, event.Warning(reasonSync, errors.Wrap(err, errAddFinalizer)))
			return reconcile.Result{RequeueAfter: shortWait}, nil
		}
	*/
	if !allDeviceTransactionsComplete {
		return reconcile.Result{RequeueAfter: veryShortWait}, errors.Wrap(r.client.Status().Update(ctx, t), errUpdateStatus)
	}

	return reconcile.Result{RequeueAfter: reconcileTimeout}, errors.Wrap(r.client.Status().Update(ctx, t), errUpdateStatus)
}

func (r *Reconciler) getDeviceGvkList(ctx context.Context, cl *target.Target, namespace, deviceName string) ([]*systemv1alpha1.Gvk, error) {
	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(namespace, deviceName))

	// gnmi get request
	req := &gnmi.GetRequest{
		Prefix:   &gnmi.Path{Target: crSystemDeviceName},
		Path:     []*gnmi.Path{{Elem: []*gnmi.PathElem{{Name: "gvk"}}}},
		Encoding: gnmi.Encoding_JSON,
	}

	// gnmi get response
	resp, err := cl.Get(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get resourceList from gnmiserver")
	}

	fmt.Printf("resourceList per device: %s\n", deviceName)
	rl, err := yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].GetVal())
	if err != nil {
		return nil, errors.Wrap(err, "cannot get value from response")
	}
	return gvkresource.GetResourceList(rl)
}

func (r *Reconciler) getGnmiClient(ctx context.Context, deviceName string) (*target.Target, error) {
	nn := &ndrv1.NetworkNode{}
	if err := r.client.Get(ctx, types.NamespacedName{Name: deviceName}, nn); err != nil {
		return nil, errors.Wrap(err, errGetNetworkNode)
	}

	if nn.GetCondition(ndrv1.ConditionKindDeviceDriverConfigured).Status != corev1.ConditionTrue {
		return nil, errors.New(errDeviceNotReady)
	}

	cfg := &gnmitypes.TargetConfig{
		Name:       "dummy",
		Address:    r.gnmiAddress,
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
		return nil, errors.Wrap(err, errNewGnmiClient)
	}
	return cl, nil
}

/*
func (r *Reconciler) getTransactionUpdates(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]interface{}) ([]*gnmi.Update, error) {
	updates := make([]*gnmi.Update, 0)
	for kind, res := range resources {
		for resName := range res {
			//r.log.Debug("transaction resources", "kind", kind, "resource name", resName)
			u, err := r.getUpdatesPerResource(ctx, kind, resName, cr.GetNamespace())
			if err != nil {
				return nil, err
			}
			updates = append(updates, u...)
		}
	}
	return updates, nil
}

func (r *Reconciler) getUpdatesPerResource(ctx context.Context, kind, resName, namespace string) ([]*gnmi.Update, error) {
	// get managed resource object
	rootPath, specData, err := r.getCrData(ctx, kind, resName, namespace)
	if err != nil {
		return nil, err
	}

	x1, err := processSpecData(rootPath, specData)
	if err != nil {
		return nil, err
	}
	fmt.Printf("getUpdatesPerResource kind: %s, resName: %s data: %v\n", kind, resName, x1)

	// we remove the first element which is aligned with the last element of the rootPath
	// gnmi does not return this information hence to compare the spec data with the gnmi resp data we need to remove
	// the first element from the Spec
	switch x := x1.(type) {
	case map[string]interface{}:
		x1 := x[rootPath.GetElem()[len(rootPath.GetElem())-1].GetName()]
		//fmt.Printf("processCreate data %v, rootPath: %s\n", x1, yparser.GnmiPath2XPath(rootPath, true))
		gnmiUpdate, err := yparser.GetUpdatesFromJSON(rootPath, x1, r.deviceSchema)
		if err != nil {
			return nil, err
		}
		return gnmiUpdate, nil
	}
	return nil, errors.New("getUpdatesPerResource wrong data: expected map[string]interface{}")

}

//process Spec data marshals the data and remove the prent hierarchical keys
func processSpecData(rootPath *gnmi.Path, specData interface{}) (interface{}, error) {
	// prepare the input data to compare against the response data
	d, err := json.Marshal(specData)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}
	var x1 interface{}
	if err := json.Unmarshal(d, &x1); err != nil {
		return nil, errors.Wrap(err, errJSONUnMarshal)
	}
	// removes the parent hierarchical ids; they are there to define the parent in k8s so
	// we can define the full path in gnmi
	return yparser.RemoveHierIDsFomData(yparser.GetHierIDsFromPath(rootPath), x1), nil
}

func (r *Reconciler) getCrData(ctx context.Context, kind, resName, namespace string) (*gnmi.Path, interface{}, error) {
	switch kind {
	case "SrlBfd":
		cr := &srlv1alpha1.SrlBfd{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "bfd"},
			},
		}, &cr.Spec, nil
	case "SrlInterface":
		cr := &srlv1alpha1.SrlInterface{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "interface", Key: map[string]string{
					"name": *cr.Spec.Interface.Name,
				}},
			},
		}, &cr.Spec, nil
	case "SrlInterfaceSubinterface":
		cr := &srlv1alpha1.SrlInterfaceSubinterface{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "interface", Key: map[string]string{
					"name": *cr.Spec.InterfaceName,
				}},
				{Name: "subinterface", Key: map[string]string{
					"index": strconv.Itoa(int(*cr.Spec.InterfaceSubinterface.Index)),
				}},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstance":
		cr := &srlv1alpha1.SrlNetworkinstance{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.Networkinstance.Name,
				}},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstanceAggregateroutes":
		cr := &srlv1alpha1.SrlNetworkinstanceAggregateroutes{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "aggregate-routes"},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstanceNexthopgroups":
		cr := &srlv1alpha1.SrlNetworkinstanceNexthopgroups{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "next-hop-groups"},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstanceProtocolsBgp":
		cr := &srlv1alpha1.SrlNetworkinstanceProtocolsBgp{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "protocols"},
				{Name: "bgp"},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstanceProtocolsBgpevpn":
		cr := &srlv1alpha1.SrlNetworkinstanceProtocolsBgpevpn{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "protocols"},
				{Name: "bgp-evpn"},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstanceProtocolsBgpvpn":
		cr := &srlv1alpha1.SrlNetworkinstanceProtocolsBgpvpn{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "protocols"},
				{Name: "bgp-vpn"},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstanceProtocolsIsis":
		cr := &srlv1alpha1.SrlNetworkinstanceProtocolsIsis{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "protocols"},
				{Name: "isis"},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstanceProtocolsLinux":
		cr := &srlv1alpha1.SrlNetworkinstanceProtocolsLinux{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "protocols"},
				{Name: "linux"},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstanceProtocolsOspf":
		cr := &srlv1alpha1.SrlNetworkinstanceProtocolsOspf{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "protocols"},
				{Name: "ospf"},
			},
		}, &cr.Spec, nil
	case "SrlNetworkinstanceStaticroutes":
		cr := &srlv1alpha1.SrlNetworkinstanceStaticroutes{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "network-instance", Key: map[string]string{
					"name": *cr.Spec.NetworkInstanceName,
				}},
				{Name: "static-routes"},
			},
		}, &cr.Spec, nil
	case "SrlRoutingpolicyAspathset":
		cr := &srlv1alpha1.SrlRoutingpolicyAspathset{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "routing-policy"},
				{Name: "as-path-set", Key: map[string]string{
					"name": *cr.Spec.RoutingpolicyAspathset.Name,
				}},
			},
		}, &cr.Spec, nil
	case "SrlRoutingpolicyCommunityset":
		cr := &srlv1alpha1.SrlRoutingpolicyCommunityset{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "routing-policy"},
				{Name: "community-set", Key: map[string]string{
					"name": *cr.Spec.RoutingpolicyCommunityset.Name,
				}},
			},
		}, &cr.Spec, nil
	case "SrlRoutingpolicyPolicy":
		cr := &srlv1alpha1.SrlRoutingpolicyPolicy{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "routing-policy"},
				{Name: "policy", Key: map[string]string{
					"name": *cr.Spec.RoutingpolicyPolicy.Name,
				}},
			},
		}, &cr.Spec, nil
	case "SrlRoutingpolicyPrefixset":
		cr := &srlv1alpha1.SrlRoutingpolicyPrefixset{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "routing-policy"},
				{Name: "prefix-set", Key: map[string]string{
					"name": *cr.Spec.RoutingpolicyPrefixset.Name,
				}},
			},
		}, &cr.Spec, nil
	case "SrlSystemName":
		cr := &srlv1alpha1.SrlSystemName{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "system"},
				{Name: "name"},
			},
		}, &cr.Spec, nil
	case "SrlSystemNetworkinstanceProtocolsBgpvpn":
		cr := &srlv1alpha1.SrlSystemNetworkinstanceProtocolsBgpvpn{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "system"},
				{Name: "network-instance"},
				{Name: "protocols"},
				{Name: "bgp-vpn"},
			},
		}, &cr.Spec, nil
	case "SrlSystemNetworkinstanceProtocolsEvpn":
		cr := &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpn{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "system"},
				{Name: "network-instance"},
				{Name: "protocols"},
				{Name: "evpn"},
			},
		}, &cr.Spec, nil
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance":
		cr := &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "system"},
				{Name: "network-instance"},
				{Name: "protocols"},
				{Name: "evpn"},
				{Name: "ethernet-segments"},
				{Name: "bgp-instance", Key: map[string]string{
					"id": *cr.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance.Id,
				}},
			}}, &cr.Spec, nil
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi":
		cr := &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "system"},
				{Name: "network-instance"},
				{Name: "protocols"},
				{Name: "evpn"},
				{Name: "ethernet-segments"},
				{Name: "bgp-instance", Key: map[string]string{
					"id": *cr.Spec.BgpInstanceId,
				}},
				{Name: "ethernet-segment", Key: map[string]string{
					"name": *cr.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Name,
				}},
			}}, &cr.Spec, nil
	case "SrlSystemNtp":
		cr := &srlv1alpha1.SrlSystemNtp{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "system"},
				{Name: "ntp"},
			}}, &cr.Spec, nil
	case "SrlTunnelinterface":
		cr := &srlv1alpha1.SrlTunnelinterface{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "tunnel-interface", Key: map[string]string{
					"name": *cr.Spec.Tunnelinterface.Name,
				}},
			}}, &cr.Spec, nil
	case "SrlTunnelinterfaceVxlaninterface":
		cr := &srlv1alpha1.SrlTunnelinterfaceVxlaninterface{}
		if err := r.client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      resName,
		}, cr); err != nil {
			// There's no need to requeue if we no longer exist. Otherwise we'll be
			// requeued implicitly because we return an error.
			r.log.Debug(errGetManaged, "error", err)
			return nil, nil, err
		}
		return &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "tunnel-interface", Key: map[string]string{
					"name": *cr.Spec.TunnelInterfaceName,
				}},
				{Name: "vxlan-interface", Key: map[string]string{
					"index": strconv.Itoa(int(*cr.Spec.TunnelinterfaceVxlaninterface.Index)),
				}},
			}}, &cr.Spec, nil
	default:
		fmt.Printf("getObject not found gvk: %s\n", kind)
		return nil, nil, fmt.Errorf("object not found: %s", kind)
	}
}
*/
