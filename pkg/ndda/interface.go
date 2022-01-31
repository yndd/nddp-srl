package ndda

import (
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndda-network/pkg/ndda/itfceinfo"
	nddov1 "github.com/yndd/nddo-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Option can be used to manipulate Options.
type Option func(Handler)

// WithLogger specifies how the Reconciler should log messages.
func WithLogger(log logging.Logger) Option {
	return func(s Handler) {
		s.WithLogger(log)
	}
}

func WithClient(c client.Client) Option {
	return func(s Handler) {
		s.WithClient(c)
	}
}

type Handler interface {
	WithLogger(log logging.Logger)
	WithClient(client.Client)
	GetSelectedNodeItfces(mg resource.Managed, epgSelectors []*nddov1.EpgInfo, nodeItfceSelectors map[string]*nddov1.ItfceInfo) (map[string][]itfceinfo.ItfceInfo, error)
	//GetSelectedNodeItfcesIrb(mg resource.Managed, s srlschema.Schema, niName string) (map[string][]itfceinfo.ItfceInfo, error)
	//GetSelectedNodeItfcesVxlan(mg resource.Managed, s srlschema.Schema, niName string) (map[string][]itfceinfo.ItfceInfo, error)
}
