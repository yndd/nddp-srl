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

package transaction

import (
	"context"
	"fmt"

	//"strings"
	"time"

	"github.com/karimra/gnmic/target"
	gnmitypes "github.com/karimra/gnmic/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	"github.com/yndd/ndd-runtime/pkg/event"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/reconciler/transaction"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-runtime/pkg/tresource"
	"github.com/yndd/ndd-runtime/pkg/utils"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	nddpsystransaction "github.com/yndd/nddp-system/pkg/transaction"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	//"sigs.k8s.io/controller-runtime/pkg/handler"

	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
	"github.com/yndd/nddp-srl/internal/shared"
	"github.com/yndd/nddp-srl/internal/srlhandler"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
)

const (
	// Errors
	errObserveTransaction = "cannot observe Transaction"
	errCreateTransaction  = "cannot create Transaction"
	errDeleteTransaction  = "cannot delete Transaction"
	errGetNetworkNode     = "cannot get NetworkNode"
	errDeviceNotReady     = "device not ready"
	errNewClient          = "cannot create new client"
)

// Setup adds a controller that reconciles the Lock.
func Setup(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) error {
	//func SetupSystemNtp(mgr ctrl.Manager, o controller.Options, nddcopts *shared.NddControllerOptions) error {

	name := transaction.ControllerName(srlv1alpha1.TransactionGroupKind)

	h := srlhandler.New(
		srlhandler.WithClient(resource.ClientApplicator{
			Client:     mgr.GetClient(),
			Applicator: resource.NewAPIPatchingApplicator(mgr.GetClient()),
		}),
		srlhandler.WithLogger(nddcopts.Logger.WithValues("srlhandler", name)))

	r := transaction.NewReconciler(mgr,
		tresource.TransactionKind(srlv1alpha1.TransactionGroupVersionKind),
		transaction.WithPollInterval(nddcopts.Poll),
		transaction.WithHandler(&handlerTransaction{
			handler: h,
		}),
		transaction.WithExternalConnecter(&connectorTransaction{
			log:         nddcopts.Logger,
			kube:        mgr.GetClient(),
			nddpSchema:  nddcopts.NddpSchema,
			newClientFn: target.NewTarget,
			gnmiAddress: nddcopts.GnmiAddress},
		),
		transaction.WithLogger(nddcopts.Logger.WithValues("transaction", name)),
		transaction.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&srlv1alpha1.SrlTransaction{}).
		Owns(&srlv1alpha1.SrlTransaction{}).
		WithEventFilter(resource.IgnoreUpdateWithoutGenerationChangePredicate()).
		Complete(r)
}

type handlerTransaction struct {
	handler srlhandler.Handler
}

func (h *handlerTransaction) GetResources(ctx context.Context, tr tresource.Transaction) (map[string]map[string]map[string]interface{}, error) {
	resources, err := h.handler.ListResourcesByTransaction(ctx, tr)
	fmt.Printf("getresources: transaction: %s, resources: %v\n", tr.GetName(), resources)
	return resources, err
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connectorTransaction struct {
	log         logging.Logger
	kube        client.Client
	nddpSchema  *yentry.Entry
	newClientFn func(c *gnmitypes.TargetConfig) *target.Target
	gnmiAddress string
}

// Connect produces an ExternalClient by:
// 1. Tracking that the managed resource is using a NetworkNode.
// 2. Getting the managed resource's NetworkNode with connection details
// A resource is mapped to a single target
func (c *connectorTransaction) Connect(ctx context.Context, tr tresource.Transaction, deviceName string) (transaction.ExternalClient, error) {
	log := c.log.WithValues("resource", tr.GetName())
	//log.Debug("Connect")

	// find network node that is configured status
	nn := &ndrv1.NetworkNode{}
	if err := c.kube.Get(ctx, types.NamespacedName{Name: deviceName}, nn); err != nil {
		return nil, errors.Wrap(err, errGetNetworkNode)
	}

	if nn.GetCondition(ndrv1.ConditionKindDeviceDriverConfigured).Status != corev1.ConditionTrue {
		return nil, errors.New(errDeviceNotReady)
	}

	cfg := &gnmitypes.TargetConfig{
		Name:       deviceName,
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

	return &externalTransaction{client: cl, log: log, nddpSchema: c.nddpSchema}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type externalTransaction struct {
	client       *target.Target
	log          logging.Logger
	deviceSchema *yentry.Entry
	nddpSchema   *yentry.Entry
}

func (e *externalTransaction) Close() {
	e.client.Close()
}

func (e *externalTransaction) Observe(ctx context.Context, tr tresource.Transaction) (transaction.ExternalObservation, error) {
	log := e.log.WithValues("Transaction", tr.GetName())
	log.Debug("Observing ...")

	// gnmi get request
	req := nddpsystransaction.GetTransactionRequest(tr, e.client.Config.Name)

	// gnmi get response
	resp, err := e.client.Get(ctx, req)
	if err != nil {
		if er, ok := status.FromError(err); ok {
			switch er.Code() {
			case codes.ResourceExhausted:
				// device or cache is exhausted
				return transaction.ExternalObservation{
					Ready:     false,
					Exhausted: true,
					Exists:    false,
					Pending:   false,
					Success:   false,
				}, nil
			case codes.Unavailable:
				// cache is not ready
				return transaction.ExternalObservation{
					Ready:     false,
					Exhausted: false,
					Exists:    false,
					Pending:   false,
					Success:   false,
				}, nil
			}
		}
	}

	gvkResourceList, err := e.getDeviceGvkList(ctx, tr)
	if err != nil {
		return transaction.ExternalObservation{}, err
	}

	t, err := nddpsystransaction.GetTransactionFromGnmiResponse(resp)
	if err != nil {
		return transaction.ExternalObservation{}, err
	}
	if t == nil {
		// Transaction does not exist
		return transaction.ExternalObservation{
			Ready:           true,
			Exhausted:       false,
			Exists:          false,
			Pending:         false,
			Success:         false,
			GvkResourceList: gvkResourceList,
		}, nil
	}
	// Transaction was already created
	switch t.Status {
	case systemv1alpha1.E_TransactionStatus_Pending:
		// Transaction pending
		return transaction.ExternalObservation{
			Ready:           true,
			Exhausted:       false,
			Exists:          true,
			Pending:         true,
			Success:         false,
			GvkResourceList: gvkResourceList,
		}, nil
	case systemv1alpha1.E_TransactionStatus_Failed:
		// Transaction failed
		return transaction.ExternalObservation{
			Ready:           true,
			Exhausted:       false,
			Exists:          true,
			Pending:         false,
			Success:         false,
			GvkResourceList: gvkResourceList,
		}, nil
	case systemv1alpha1.E_TransactionStatus_Success:
		// Transaction success
		return transaction.ExternalObservation{
			Ready:           true,
			Exhausted:       false,
			Exists:          true,
			Pending:         false,
			Success:         false,
			GvkResourceList: gvkResourceList,
		}, nil
	default:
		return transaction.ExternalObservation{}, errors.New("unknown status of the transaction")
	}

}

func (e *externalTransaction) Create(ctx context.Context, tr tresource.Transaction, gvkList []string) error {
	log := e.log.WithValues("Transaction", tr.GetName())
	log.Debug("Create ...")
	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(tr.GetNamespace(), e.client.Config.Name))

	for _, gvk := range gvkList {
		log.Debug("Create ...", "gvk", gvk)
	}

	updates, err := nddpsystransaction.ProcessTransaction(
		&systemv1alpha1.Transaction{
			Action:           systemv1alpha1.E_TransactionAction_Create,
			Generation:       tr.GetOwnerGeneration(),
			ResourceRevision: tr.GetResourceVersion(),
			Name:             tr.GetName(),
			Gvk:              gvkList,
			Status:           systemv1alpha1.E_TransactionStatus_Pending,
		}, e.nddpSchema,
	)

	req := &gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (e *externalTransaction) Delete(ctx context.Context, tr tresource.Transaction, gvkList []string) error {
	log := e.log.WithValues("Transaction", tr.GetName())
	log.Debug("Delete ...")
	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(tr.GetNamespace(), e.client.Config.Name))

	updates, err := nddpsystransaction.ProcessTransaction(
		&systemv1alpha1.Transaction{
			Action:           systemv1alpha1.E_TransactionAction_Delete,
			Generation:       tr.GetOwnerGeneration(),
			ResourceRevision: tr.GetResourceVersion(),
			Name:             tr.GetName(),
			Gvk:              gvkList,
			Status:           systemv1alpha1.E_TransactionStatus_Pending,
		}, e.nddpSchema,
	)

	req := &gnmi.SetRequest{
		Prefix:  &gnmi.Path{Target: crSystemDeviceName},
		Replace: updates,
	}

	_, err = e.client.Set(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (e *externalTransaction) getDeviceGvkList(ctx context.Context, tr tresource.Transaction) ([]*systemv1alpha1.Gvk, error) {
	crSystemDeviceName := shared.GetCrSystemDeviceName(shared.GetCrDeviceName(tr.GetNamespace(), e.client.Config.Name))

	// gnmi get request
	req := &gnmi.GetRequest{
		Prefix:   &gnmi.Path{Target: crSystemDeviceName},
		Path:     []*gnmi.Path{{Elem: []*gnmi.PathElem{{Name: "gvk"}}}},
		Encoding: gnmi.Encoding_JSON,
	}

	// gnmi get response
	resp, err := e.client.Get(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get resourceList from gnmiserver")
	}

	rl, err := yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].GetVal())
	if err != nil {
		return nil, errors.Wrap(err, "cannot get value from response")
	}
	return gvkresource.GetResourceList(rl)
}
