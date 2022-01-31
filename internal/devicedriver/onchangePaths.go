package devicedriver

import "github.com/yndd/ndd-runtime/pkg/utils"

var (
	subscriptions = []*string{
		utils.StringPtr("/acl"),
		utils.StringPtr("/bfd"),
		utils.StringPtr("/interface"),
		utils.StringPtr("/network-instance"),
		utils.StringPtr("/platform"),
		utils.StringPtr("/qos"),
		utils.StringPtr("/routing-policy"),
		utils.StringPtr("/tunnel"),
		utils.StringPtr("/tunnel-interface"),
		utils.StringPtr("/system/snmp"),
		utils.StringPtr("/system/sflow"),
		utils.StringPtr("/system/ntp"),
		utils.StringPtr("/system/network-instance"),
		utils.StringPtr("/system/name"),
		utils.StringPtr("/system/mtu"),
		utils.StringPtr("/system/maintenance"),
		utils.StringPtr("/system/lldp"),
		utils.StringPtr("/system/lacp"),
		utils.StringPtr("/system/authentication"),
		utils.StringPtr("/system/banner"),
		utils.StringPtr("/system/bridge-table"),
		utils.StringPtr("/system/ftp-server"),
		utils.StringPtr("/system/ip-load-balancing"),
		utils.StringPtr("/system/json-rpc-server"),
	}
)

/*
ExceptionPaths: []string{
	"/interface[name=mgmt0]",
	"/network-instance[name=mgmt]",
	"/system/gnmi-server",
	"/system/tls",
	"/system/ssh-server",
	"/system/aaa",
	"/system/logging",
	"/acl/cpm-filter",
},
ExplicitExceptionPaths: []string{
	"/acl",
	"/bfd",
	"/platform",
	"/qos",
	"/routing-policy",
	"/system",
	"/tunnel",
},
*/
