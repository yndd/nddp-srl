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

package device

import (
	"context"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/karimra/gnmic/types"
	"github.com/pkg/errors"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	"github.com/yndd/ndd-runtime/pkg/event"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/meta"
	"github.com/yndd/ndd-runtime/pkg/utils"
	"github.com/yndd/nddp-srl/internal/devicedriver"
	"github.com/yndd/nddp-srl/internal/shared"

	//"github.com/yndd/ndd-runtime/pkg/meta"
	"github.com/yndd/nddo-runtime/pkg/resource"
	//corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	// Finalizer
	finalizer = "device.srl.ndda.yndd.io"

	// default
	//defaultGrpcPort = 9999

	// Timers
	defaultTimeout   = 5 * time.Second
	reconcileTimeout = 1 * time.Minute
	shortWait        = 30 * time.Second
	veryShortWait    = 5 * time.Second
	longWait         = 1 * time.Minute

	// Errors
	errGetNetworkNode     = "cannot get network node resource"
	errUpdateStatus       = "cannot update network node status"
	errCreateDeviceDriver = "cannot create device driver"
	errStopDeviceDriver   = "cannot stop device driver"

	errAddFinalizer    = "cannot add network node finalizer"
	errRemoveFinalizer = "cannot remove network node finalizer"

	errCredentials = "invalid credentials"

	//errDeleteObjects = "cannot delete configmap, servide or deployment"
	//errCreateObjects = "cannot create configmap, servide or deployment"

	// Event reasons
	reasonSync event.Reason = "SyncNetworkNode"
)

// ReconcilerOption is used to configure the Reconciler.
type ReconcilerOption func(*Reconciler)

// WithNewNetworkNodeFn determines the type of network node being reconciled.
func WithNewNetworkNodeFn(f func() ndrv1.Nn) ReconcilerOption {
	return func(r *Reconciler) {
		r.newNetworkNode = f
	}
}

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

// WithRecorder specifies how the Reconciler should record Kubernetes events.
func WithDeviceDriverChannel(reqCh chan shared.DeviceUpdate, respCh chan shared.DeviceResponse) ReconcilerOption {
	return func(r *Reconciler) {
		r.deviceDriverRequestCh = reqCh
		r.deviceDriverResponseCh = respCh
	}
}

// Reconciler reconciles packages.
type Reconciler struct {
	client      resource.ClientApplicator
	nnFinalizer resource.Finalizer
	log         logging.Logger
	record      event.Recorder

	newNetworkNode         func() ndrv1.Nn
	devices                map[string]devicedriver.DeviceDriver
	deviceDriverRequestCh  chan shared.DeviceUpdate
	deviceDriverResponseCh chan shared.DeviceResponse
}

// Setup adds a controller that reconciles the Lock.
func Setup(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) error {
	name := "dvr/" + strings.ToLower(ndrv1.NetworkNodeKind)
	nn := func() ndrv1.Nn { return &ndrv1.NetworkNode{} }

	r := NewReconciler(mgr,
		WithLogger(nddcopts.Logger.WithValues("controller", name)),
		WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		WithNewNetworkNodeFn(nn),
		WithDeviceDriverChannel(nddcopts.DeviceDriverRequestCh, nddcopts.DeviceDriverResponseCh),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&ndrv1.NetworkNode{}).
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
		nnFinalizer: resource.NewAPIFinalizer(mgr.GetClient(), finalizer),
		log:         logging.NewNopLogger(),
		record:      event.NewNopRecorder(),
		devices:     make(map[string]devicedriver.DeviceDriver),
	}

	for _, f := range opts {
		f(r)
	}

	return r
}

// Reconcile network node.
func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) { // nolint:gocyclo
	log := r.log.WithValues("request", req)
	log.Debug("Network Node", "NameSpace", req.NamespacedName)

	nn := r.newNetworkNode()
	if err := r.client.Get(ctx, req.NamespacedName, nn); err != nil {
		// There's no need to requeue if we no longer exist. Otherwise we'll be
		// requeued implicitly because we return an error.
		log.Debug(errGetNetworkNode, "error", err)
		return reconcile.Result{}, errors.Wrap(resource.IgnoreNotFound(err), errGetNetworkNode)
	}
	// TODO check if the network node is of the right Kind
	// if not ignore this update, otherwise process it

	log.Debug("Health status", "status", nn.GetCondition(ndrv1.ConditionKindDeviceDriverHealthy).Status)
	if meta.WasDeleted(nn) {

		// CHECK IF A DEVICEDRIVER WAS RUNNING -> IF SO DELETE IT
		r.deviceDriverRequestCh <- shared.DeviceUpdate{
			Action:    shared.DeviceStop,
			Namespace: nn.GetNamespace(),
			TargetConfig: &types.TargetConfig{
				Name: nn.GetName(),
			},
		}
		resp := <-r.deviceDriverResponseCh
		if resp.Error != nil {
			log.Debug(errStopDeviceDriver, "error", resp.Error)
			r.record.Event(nn, event.Warning(reasonSync, errors.Wrap(resp.Error, errStopDeviceDriver)))
			nn.SetConditions(ndrv1.Unhealthy(), ndrv1.NotConfigured(), ndrv1.NotDiscovered())
			return reconcile.Result{RequeueAfter: shortWait}, errors.Wrap(r.client.Status().Update(ctx, nn), errUpdateStatus)
		}

		// Delete finalizer after the object is deleted
		if err := r.nnFinalizer.RemoveFinalizer(ctx, nn); err != nil {
			log.Debug(errRemoveFinalizer, "error", err)
			r.record.Event(nn, event.Warning(reasonSync, errors.Wrap(err, errRemoveFinalizer)))
			return reconcile.Result{RequeueAfter: shortWait}, nil
		}
		return reconcile.Result{Requeue: false}, nil
	}

	// Add a finalizer to newly created objects and update the conditions
	if err := r.nnFinalizer.AddFinalizer(ctx, nn); err != nil {
		log.Debug(errAddFinalizer, "error", err)
		r.record.Event(nn, event.Warning(reasonSync, errors.Wrap(err, errAddFinalizer)))
		return reconcile.Result{RequeueAfter: shortWait}, nil
	}

	// Retrieve the Login details from the network node spec and validate
	// the network node details and build the credentials for communicating
	// to the network node.
	creds, err := r.validateCredentials(ctx, nn)
	//log.Debug("Network node creds", "creds", creds, "err", err)
	if err != nil || creds == nil {

		// CHECK IF A DEVICEDRIVER WAS ALREDY RUNNING -> IF SO STOP IT
		r.deviceDriverRequestCh <- shared.DeviceUpdate{
			Action:    shared.DeviceStop,
			Namespace: nn.GetNamespace(),
			TargetConfig: &types.TargetConfig{
				Name: nn.GetName(),
			},
		}
		resp := <-r.deviceDriverResponseCh
		if resp.Error != nil {
			log.Debug(errStopDeviceDriver, "error", resp.Error)
			r.record.Event(nn, event.Warning(reasonSync, errors.Wrap(resp.Error, errStopDeviceDriver)))
			nn.SetDeviceDetails(&ndrv1.DeviceDetails{
				HostName:     utils.StringPtr(nn.GetName()),
				Kind:         new(string),
				SwVersion:    new(string),
				MacAddress:   new(string),
				SerialNumber: new(string),
			})
			nn.SetConditions(ndrv1.Unhealthy(), ndrv1.NotConfigured(), ndrv1.NotDiscovered())
			return reconcile.Result{RequeueAfter: shortWait}, errors.Wrap(r.client.Status().Update(ctx, nn), errUpdateStatus)
		}

		log.Debug(errCredentials, "error", err)
		r.record.Event(nn, event.Warning(reasonSync, errors.Wrap(err, errCredentials)))
		nn.SetConditions(ndrv1.Unhealthy(), ndrv1.NotConfigured(), ndrv1.NotDiscovered())
		return reconcile.Result{RequeueAfter: shortWait}, errors.Wrap(r.client.Status().Update(ctx, nn), errUpdateStatus)
	}

	// CHECK IF A DEVICEDRIVER WAS ALREDY RUNNING -> IF NOT CREATE/START IT
	r.deviceDriverRequestCh <- shared.DeviceUpdate{
		Action:    shared.DeviceStatus,
		Namespace: nn.GetNamespace(),
		TargetConfig: &types.TargetConfig{
			Name: nn.GetName(),
		},
	}
	resp := <-r.deviceDriverResponseCh
	// success means exists in this context
	if !resp.Exists {
		// when the target does not exist, we have to create it
		r.deviceDriverRequestCh <- shared.DeviceUpdate{
			Action:       shared.DeviceStart,
			Namespace:    nn.GetNamespace(),
			TargetConfig: getTargetConfig(nn, creds),
		}
		resp := <-r.deviceDriverResponseCh
		if resp.Error != nil {
			log.Debug(errCreateDeviceDriver, "error", resp.Error)
			r.record.Event(nn, event.Warning(reasonSync, errors.Wrap(resp.Error, errCreateDeviceDriver)))
			nn.SetConditions(ndrv1.Unhealthy(), ndrv1.NotConfigured(), ndrv1.NotDiscovered())
			return reconcile.Result{RequeueAfter: shortWait}, errors.Wrap(r.client.Status().Update(ctx, nn), errUpdateStatus)
		}
		nn.SetConditions(ndrv1.Healthy(), ndrv1.Configured(), ndrv1.Discovered())
		nn.SetDeviceDetails(resp.DeviceDetails)
	} else {
		// check if the config changed
		if !cmp.Equal(getTargetConfig(nn, creds), resp.TargetConfig) {
			// CONFIG CHANGED
			r.deviceDriverRequestCh <- shared.DeviceUpdate{
				Action:    shared.DeviceStop,
				Namespace: nn.GetNamespace(),
				TargetConfig: &types.TargetConfig{
					Name: nn.GetName(),
				},
			}
			// STOP THE DEVCE DRIVER
			resp := <-r.deviceDriverResponseCh
			if resp.Error != nil {
				log.Debug(errStopDeviceDriver, "error", resp.Error)
				r.record.Event(nn, event.Warning(reasonSync, errors.Wrap(resp.Error, errStopDeviceDriver)))
				nn.SetDeviceDetails(&ndrv1.DeviceDetails{
					HostName:     utils.StringPtr(nn.GetName()),
					Kind:         new(string),
					SwVersion:    new(string),
					MacAddress:   new(string),
					SerialNumber: new(string),
				})
				nn.SetConditions(ndrv1.Unhealthy(), ndrv1.NotConfigured(), ndrv1.NotDiscovered())
				return reconcile.Result{RequeueAfter: shortWait}, errors.Wrap(r.client.Status().Update(ctx, nn), errUpdateStatus)
			}
			// START THE DEVCE DRIVER
			nn.SetDeviceDetails(&ndrv1.DeviceDetails{
				HostName:     utils.StringPtr(nn.GetName()),
				Kind:         new(string),
				SwVersion:    new(string),
				MacAddress:   new(string),
				SerialNumber: new(string),
			})
			r.deviceDriverRequestCh <- shared.DeviceUpdate{
				Action:       shared.DeviceStart,
				Namespace:    nn.GetNamespace(),
				TargetConfig: getTargetConfig(nn, creds),
			}
			resp = <-r.deviceDriverResponseCh
			if resp.Error != nil {
				log.Debug(errCreateDeviceDriver, "error", resp.Error)
				r.record.Event(nn, event.Warning(reasonSync, errors.Wrap(resp.Error, errCreateDeviceDriver)))
				nn.SetConditions(ndrv1.Unhealthy(), ndrv1.NotConfigured(), ndrv1.NotDiscovered())
				return reconcile.Result{RequeueAfter: shortWait}, errors.Wrap(r.client.Status().Update(ctx, nn), errUpdateStatus)
			}
			nn.SetConditions(ndrv1.Healthy(), ndrv1.Configured(), ndrv1.Discovered())
			nn.SetDeviceDetails(resp.DeviceDetails)
		} else {
			nn.SetConditions(ndrv1.Healthy(), ndrv1.Configured(), ndrv1.Discovered())
			nn.SetDeviceDetails(resp.DeviceDetails)
		}
	}
	return reconcile.Result{RequeueAfter: reconcileTimeout}, errors.Wrap(r.client.Status().Update(ctx, nn), errUpdateStatus)
}

func getTargetConfig(nn ndrv1.Nn, creds *Credentials) *types.TargetConfig {
	return &types.TargetConfig{
		Name:       nn.GetName(),
		Address:    nn.GetTargetAddress(),
		Username:   &creds.Username,
		Password:   &creds.Password,
		Timeout:    defaultTimeout,
		Insecure:   utils.BoolPtr(nn.GetTargetInsecure()),
		SkipVerify: utils.BoolPtr(nn.GetTargetSkipVerify()),
		TLSCA:      utils.StringPtr(""), //TODO TLS
		TLSCert:    utils.StringPtr(""), //TODO TLS
		TLSKey:     utils.StringPtr(""), //TODO TLS
		Gzip:       utils.BoolPtr(false),
	}
}
