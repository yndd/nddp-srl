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
package yangschema

import (
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndd-yang/pkg/leafref"
	"github.com/yndd/ndd-yang/pkg/yentry"
)

func initSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstance(p *yentry.Entry, opts ...yentry.EntryOption) *yentry.Entry {
	children := map[string]yentry.EntryInitFunc{
		"ethernet-segment": initSystemNetworkinstanceProtocolsEvpnEthernetsegmentsBgpinstanceEthernetsegment,
	}
	e := &yentry.Entry{
		Name: "bgp-instance",
		Key: []string{
			"id",
		},
		Module:           "",
		Namespace:        "",
		Prefix:           "srl_nokia-system-bgp-evpn-ethernet-segments",
		Parent:           p,
		Children:         make(map[string]*yentry.Entry),
		ResourceBoundary: true,
		LeafRefs: []*leafref.LeafRef{
			{
				LocalPath: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "id"},
					},
				},
				RemotePath: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "system"},
						{Name: "network-instance"},
						{Name: "protocols"},
						{Name: "bgp-vpn"},
						{Name: "bgp-instance", Key: map[string]string{"id": ""}},
					},
				},
			},
		},
		Defaults: map[string]string{},
	}

	for _, opt := range opts {
		opt(e)
	}

	for name, initFunc := range children {
		e.Children[name] = initFunc(e, yentry.WithLogging(e.Log))
	}

	//if e.ResourceBoundary {
	//	e.Register(&gnmi.Path{Elem: []*gnmi.PathElem{}})
	//}

	return e
}
