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
		utils.StringPtr("/system"),
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
