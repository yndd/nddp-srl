package ndda

import (
	"context"

	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/nddo-runtime/pkg/resource"
	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func New(opts ...Option) Handler {
	s := &handler{
		newItfceList: func() srlv1alpha1.IFSrlInterfaceList { return &srlv1alpha1.SrlInterfaceList{} },
		newSubInterfaceList: func() srlv1alpha1.IFSrlInterfaceSubinterfaceList {
			return &srlv1alpha1.SrlInterfaceSubinterfaceList{}
		},
		newNiList: func() srlv1alpha1.IFSrlNetworkinstanceList {
			return &srlv1alpha1.SrlNetworkinstanceList{}
		},
		ctx: context.Background(),
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
	client client.Client
	ctx    context.Context

	newItfceList        func() srlv1alpha1.IFSrlInterfaceList
	newSubInterfaceList func() srlv1alpha1.IFSrlInterfaceSubinterfaceList
	newNiList           func() srlv1alpha1.IFSrlNetworkinstanceList
}
