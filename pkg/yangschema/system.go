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

func initSystem(p *yentry.Entry, opts ...yentry.EntryOption) *yentry.Entry {
	children := map[string]yentry.EntryInitFunc{
		"aaa":              initSystemAaa,
		"authentication":   initSystemAuthentication,
		"banner":           initSystemBanner,
		"boot":             initSystemBoot,
		"bridge-table":     initSystemBridgetable,
		"clock":            initSystemClock,
		"configuration":    initSystemConfiguration,
		"dhcp-server":      initSystemDhcpserver,
		"dns":              initSystemDns,
		"ftp-server":       initSystemFtpserver,
		"gnmi-server":      initSystemGnmiserver,
		"gribi-server":     initSystemGribiserver,
		"information":      initSystemInformation,
		"json-rpc-server":  initSystemJsonrpcserver,
		"lacp":             initSystemLacp,
		"lldp":             initSystemLldp,
		"load-balancing":   initSystemLoadbalancing,
		"logging":          initSystemLogging,
		"maintenance":      initSystemMaintenance,
		"mirroring":        initSystemMirroring,
		"mpls":             initSystemMpls,
		"mtu":              initSystemMtu,
		"name":             initSystemName,
		"network-instance": initSystemNetworkinstance,
		"ntp":              initSystemNtp,
		"ra-guard-policy":  initSystemRaguardpolicy,
		"sflow":            initSystemSflow,
		"snmp":             initSystemSnmp,
		"ssh-server":       initSystemSshserver,
		"tls":              initSystemTls,
		"trace-options":    initSystemTraceoptions,
		"warm-reboot":      initSystemWarmreboot,
	}
	e := &yentry.Entry{
		Name:             "system",
		Key:              []string{},
		Module:           "srl_nokia-system",
		Namespace:        "urn:srl_nokia/system",
		Prefix:           "srl_nokia-system",
		Parent:           p,
		Children:         make(map[string]*yentry.Entry),
		ResourceBoundary: false,
		LeafRefs:         []*leafref.LeafRef{},
		Defaults:         map[string]string{},
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
