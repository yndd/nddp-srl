package srlhandler

import (
	"context"
	"strings"
	"sync"

	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/nddo-runtime/pkg/resource"
	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
	schema "github.com/yndd/nddp-srl/pkg/srlschema/v1alpha1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func New(opts ...Option) Handler {
	s := &handler{
		schema: make(map[string]schema.Schema),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (r *handler) WithLogger(log logging.Logger) {
	r.log = log
}

func (r *handler) WithClient(c client.Client) {
	r.client = resource.ClientApplicator{
		Client:     c,
		Applicator: resource.NewAPIPatchingApplicator(c),
	}
}

type handler struct {
	log logging.Logger
	// kubernetes
	client resource.ClientApplicator

	mutex  sync.Mutex
	schema map[string]schema.Schema
}

func (r *handler) InitSchema(crName string) schema.Schema {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, ok := r.schema[crName]; !ok {
		r.schema[crName] = schema.NewSchema(r.client)
	}
	return r.schema[crName]
}

func (r *handler) DestroySchema(crName string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.schema, crName)
}

func (r *handler) ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction) (map[string]map[string]map[string]interface{}, error) {
	crName := getCrName(cr)
	ds := r.InitSchema(crName)
	ds.InitializeDummySchema()
	resources, err := ds.ListResourcesByTransaction(ctx, cr)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func getCrName(cr srlv1alpha1.IFSrlTransaction) string {
	return strings.Join([]string{cr.GetNamespace(), cr.GetName()}, ".")
}
