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
	"github.com/yndd/ndd-yang/pkg/leafref"
	"github.com/yndd/ndd-yang/pkg/yentry"
)

func initInterfaceSubinterfaceIpv4VrrpVrrpgroup(p *yentry.Entry, opts ...yentry.EntryOption) *yentry.Entry {
	children := map[string]yentry.EntryInitFunc{
		"authentication":     initInterfaceSubinterfaceIpv4VrrpVrrpgroupAuthentication,
		"interface-tracking": initInterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetracking,
		"statistics":         initInterfaceSubinterfaceIpv4VrrpVrrpgroupStatistics,
		"virtual-address":    initInterfaceSubinterfaceIpv4VrrpVrrpgroupVirtualaddress,
	}
	e := &yentry.Entry{
		Name: "vrrp-group",
		Key: []string{
			"virtual-router-id",
		},
		Module:           "",
		Namespace:        "",
		Prefix:           "srl_nokia-if-ip-vrrp",
		Parent:           p,
		Children:         make(map[string]*yentry.Entry),
		ResourceBoundary: false,
		LeafRefs:         []*leafref.LeafRef{},
		Defaults: map[string]string{
			"admin-state":             "enable",
			"advertise-interval":      "1000",
			"master-inherit-interval": "false",
			"priority":                "100",
			"version":                 "2",
		},
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
