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

package v1alpha1

import (
	"reflect"

	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// InterfaceSubinterface struct
type InterfaceSubinterface struct {
	Acl *InterfaceSubinterfaceAcl `json:"acl,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_InterfaceSubinterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Anycastgw   *InterfaceSubinterfaceAnycastgw   `json:"anycast-gw,omitempty"`
	Bridgetable *InterfaceSubinterfaceBridgetable `json:"bridge-table,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=9999
	Index *uint32 `json:"index"`
	// kubebuilder:validation:Minimum=1280
	// kubebuilder:validation:Maximum=9486
	Ipmtu *uint16                    `json:"ip-mtu,omitempty"`
	Ipv4  *InterfaceSubinterfaceIpv4 `json:"ipv4,omitempty"`
	Ipv6  *InterfaceSubinterfaceIpv6 `json:"ipv6,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	L2mtu                  *uint16                                      `json:"l2-mtu,omitempty"`
	Localmirrordestination *InterfaceSubinterfaceLocalmirrordestination `json:"local-mirror-destination,omitempty"`
	// kubebuilder:validation:Minimum=1284
	// kubebuilder:validation:Maximum=9496
	Mplsmtu *uint16                       `json:"mpls-mtu,omitempty"`
	Qos     *InterfaceSubinterfaceQos     `json:"qos,omitempty"`
	Raguard *InterfaceSubinterfaceRaguard `json:"ra-guard,omitempty"`
	Type    *string                       `json:"type,omitempty"`
	Vlan    *InterfaceSubinterfaceVlan    `json:"vlan,omitempty"`
}

// InterfaceSubinterfaceAcl struct
type InterfaceSubinterfaceAcl struct {
	Input  *InterfaceSubinterfaceAclInput  `json:"input,omitempty"`
	Output *InterfaceSubinterfaceAclOutput `json:"output,omitempty"`
}

// InterfaceSubinterfaceAclInput struct
type InterfaceSubinterfaceAclInput struct {
	Ipv4filter *string `json:"ipv4-filter,omitempty"`
	Ipv6filter *string `json:"ipv6-filter,omitempty"`
}

// InterfaceSubinterfaceAclOutput struct
type InterfaceSubinterfaceAclOutput struct {
	Ipv4filter *string `json:"ipv4-filter,omitempty"`
	Ipv6filter *string `json:"ipv6-filter,omitempty"`
}

// InterfaceSubinterfaceAnycastgw struct
type InterfaceSubinterfaceAnycastgw struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Anycastgwmac *string `json:"anycast-gw-mac,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	Virtualrouterid *uint8 `json:"virtual-router-id,omitempty"`
}

// InterfaceSubinterfaceBridgetable struct
type InterfaceSubinterfaceBridgetable struct {
	// +kubebuilder:default:=false
	Discardunknownsrcmac *bool                                           `json:"discard-unknown-src-mac,omitempty"`
	Macduplication       *InterfaceSubinterfaceBridgetableMacduplication `json:"mac-duplication,omitempty"`
	Maclearning          *InterfaceSubinterfaceBridgetableMaclearning    `json:"mac-learning,omitempty"`
	Maclimit             *InterfaceSubinterfaceBridgetableMaclimit       `json:"mac-limit,omitempty"`
}

// InterfaceSubinterfaceBridgetableMacduplication struct
type InterfaceSubinterfaceBridgetableMacduplication struct {
	// +kubebuilder:validation:Enum=`blackhole`;`oper-down`;`stop-learning`;`use-net-instance-action`
	// +kubebuilder:default:="use-net-instance-action"
	Action E_InterfaceSubinterfaceBridgetableMacduplicationAction `json:"action,omitempty"`
	//Action *string `json:"action,omitempty"`
}

// InterfaceSubinterfaceBridgetableMaclearning struct
type InterfaceSubinterfaceBridgetableMaclearning struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_InterfaceSubinterfaceBridgetableMaclearningAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Aging *InterfaceSubinterfaceBridgetableMaclearningAging `json:"aging,omitempty"`
}

// InterfaceSubinterfaceBridgetableMaclearningAging struct
type InterfaceSubinterfaceBridgetableMaclearningAging struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_InterfaceSubinterfaceBridgetableMaclearningAgingAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// InterfaceSubinterfaceBridgetableMaclimit struct
type InterfaceSubinterfaceBridgetableMaclimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=250
	Maximumentries *int32 `json:"maximum-entries,omitempty"`
	// kubebuilder:validation:Minimum=6
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=95
	Warningthresholdpct *int32 `json:"warning-threshold-pct,omitempty"`
}

// InterfaceSubinterfaceIpv4 struct
type InterfaceSubinterfaceIpv4 struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=64
	Address []*InterfaceSubinterfaceIpv4Address `json:"address,omitempty"`
	// +kubebuilder:default:=false
	Allowdirectedbroadcast *bool                                `json:"allow-directed-broadcast,omitempty"`
	Arp                    *InterfaceSubinterfaceIpv4Arp        `json:"arp,omitempty"`
	Dhcpclient             *InterfaceSubinterfaceIpv4Dhcpclient `json:"dhcp-client,omitempty"`
	Dhcprelay              *InterfaceSubinterfaceIpv4Dhcprelay  `json:"dhcp-relay,omitempty"`
	Dhcpserver             *InterfaceSubinterfaceIpv4Dhcpserver `json:"dhcp-server,omitempty"`
	Vrrp                   *InterfaceSubinterfaceIpv4Vrrp       `json:"vrrp,omitempty"`
}

// InterfaceSubinterfaceIpv4Address struct
type InterfaceSubinterfaceIpv4Address struct {
	Anycastgw *bool `json:"anycast-gw,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	Ipprefix *string `json:"ip-prefix"`
	Primary  *string `json:"primary,omitempty"`
}

// InterfaceSubinterfaceIpv4Arp struct
type InterfaceSubinterfaceIpv4Arp struct {
	Debug *InterfaceSubinterfaceIpv4ArpDebug `json:"debug,omitempty"`
	// +kubebuilder:default:=true
	Duplicateaddressdetection *bool                                  `json:"duplicate-address-detection,omitempty"`
	Evpn                      *InterfaceSubinterfaceIpv4ArpEvpn      `json:"evpn,omitempty"`
	Hostroute                 *InterfaceSubinterfaceIpv4ArpHostroute `json:"host-route,omitempty"`
	// +kubebuilder:default:=false
	Learnunsolicited *bool `json:"learn-unsolicited,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Neighbor []*InterfaceSubinterfaceIpv4ArpNeighbor `json:"neighbor,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=14400
	Timeout *uint16 `json:"timeout,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpDebug struct
type InterfaceSubinterfaceIpv4ArpDebug struct {
	// +kubebuilder:validation:Enum=`messages`
	Debug E_InterfaceSubinterfaceIpv4ArpDebugDebug `json:"debug,omitempty"`
	//Debug *string `json:"debug,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpEvpn struct
type InterfaceSubinterfaceIpv4ArpEvpn struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Advertise []*InterfaceSubinterfaceIpv4ArpEvpnAdvertise `json:"advertise,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpEvpnAdvertise struct
type InterfaceSubinterfaceIpv4ArpEvpnAdvertise struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	Admintag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`static`
	Routetype E_InterfaceSubinterfaceIpv4ArpEvpnAdvertiseRoutetype `json:"route-type,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpHostroute struct
type InterfaceSubinterfaceIpv4ArpHostroute struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Populate []*InterfaceSubinterfaceIpv4ArpHostroutePopulate `json:"populate,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpHostroutePopulate struct
type InterfaceSubinterfaceIpv4ArpHostroutePopulate struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Admintag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`evpn`;`static`
	Routetype E_InterfaceSubinterfaceIpv4ArpHostroutePopulateRoutetype `json:"route-type,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpNeighbor struct
type InterfaceSubinterfaceIpv4ArpNeighbor struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4address *string `json:"ipv4-address"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Linklayeraddress *string `json:"link-layer-address"`
}

// InterfaceSubinterfaceIpv4Dhcpclient struct
type InterfaceSubinterfaceIpv4Dhcpclient struct {
	Traceoptions *InterfaceSubinterfaceIpv4DhcpclientTraceoptions `json:"trace-options,omitempty"`
}

// InterfaceSubinterfaceIpv4DhcpclientTraceoptions struct
type InterfaceSubinterfaceIpv4DhcpclientTraceoptions struct {
	Trace *InterfaceSubinterfaceIpv4DhcpclientTraceoptionsTrace `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv4DhcpclientTraceoptionsTrace struct
type InterfaceSubinterfaceIpv4DhcpclientTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_InterfaceSubinterfaceIpv4DhcpclientTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv4Dhcprelay struct
type InterfaceSubinterfaceIpv4Dhcprelay struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_InterfaceSubinterfaceIpv4DhcprelayAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Giaddress    *string                                         `json:"gi-address,omitempty"`
	Option       *InterfaceSubinterfaceIpv4DhcprelayOption       `json:"option,omitempty"`
	Server       *InterfaceSubinterfaceIpv4DhcprelayServer       `json:"server,omitempty"`
	Traceoptions *InterfaceSubinterfaceIpv4DhcprelayTraceoptions `json:"trace-options,omitempty"`
	// +kubebuilder:default:=false
	Usegiaddrassrcipaddr *bool `json:"use-gi-addr-as-src-ip-addr,omitempty"`
}

// InterfaceSubinterfaceIpv4DhcprelayOption struct
type InterfaceSubinterfaceIpv4DhcprelayOption struct {
	// +kubebuilder:validation:Enum=`circuit-id`;`remote-id`
	Option E_InterfaceSubinterfaceIpv4DhcprelayOptionOption `json:"option,omitempty"`
	//Option *string `json:"option,omitempty"`
}

// InterfaceSubinterfaceIpv4DhcprelayServer struct
type InterfaceSubinterfaceIpv4DhcprelayServer struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Server *string `json:"server,omitempty"`
}

// InterfaceSubinterfaceIpv4DhcprelayTraceoptions struct
type InterfaceSubinterfaceIpv4DhcprelayTraceoptions struct {
	Trace *InterfaceSubinterfaceIpv4DhcprelayTraceoptionsTrace `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv4DhcprelayTraceoptionsTrace struct
type InterfaceSubinterfaceIpv4DhcprelayTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_InterfaceSubinterfaceIpv4DhcprelayTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv4Dhcpserver struct
type InterfaceSubinterfaceIpv4Dhcpserver struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_InterfaceSubinterfaceIpv4DhcpserverAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// InterfaceSubinterfaceIpv4Vrrp struct
type InterfaceSubinterfaceIpv4Vrrp struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Vrrpgroup []*InterfaceSubinterfaceIpv4VrrpVrrpgroup `json:"vrrp-group,omitempty"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpgroup struct
type InterfaceSubinterfaceIpv4VrrpVrrpgroup struct {
	Acceptmode *bool `json:"accept-mode,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_InterfaceSubinterfaceIpv4VrrpVrrpgroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1000
	Advertiseinterval *uint16                                               `json:"advertise-interval,omitempty"`
	Authentication    *InterfaceSubinterfaceIpv4VrrpVrrpgroupAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Initdelay         *uint16                                                  `json:"init-delay,omitempty"`
	Interfacetracking *InterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetracking `json:"interface-tracking,omitempty"`
	// +kubebuilder:default:=false
	Masterinheritinterval *bool `json:"master-inherit-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Operinterval *uint16 `json:"oper-interval,omitempty"`
	Preempt      *bool   `json:"preempt,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Preemptdelay *uint16 `json:"preempt-delay,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=100
	Priority   *uint8                                            `json:"priority,omitempty"`
	Statistics *InterfaceSubinterfaceIpv4VrrpVrrpgroupStatistics `json:"statistics,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=2
	Version        *uint8                                                `json:"version,omitempty"`
	Virtualaddress *InterfaceSubinterfaceIpv4VrrpVrrpgroupVirtualaddress `json:"virtual-address,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Virtualrouterid *uint8 `json:"virtual-router-id"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpgroupAuthentication struct
type InterfaceSubinterfaceIpv4VrrpVrrpgroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetracking struct
type InterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetracking struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Trackinterface []*InterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetrackingTrackinterface `json:"track-interface,omitempty"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetrackingTrackinterface struct
type InterfaceSubinterfaceIpv4VrrpVrrpgroupInterfacetrackingTrackinterface struct {
	Interface *string `json:"interface"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Prioritydecrement *uint8 `json:"priority-decrement,omitempty"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpgroupStatistics struct
type InterfaceSubinterfaceIpv4VrrpVrrpgroupStatistics struct {
}

// InterfaceSubinterfaceIpv4VrrpVrrpgroupVirtualaddress struct
type InterfaceSubinterfaceIpv4VrrpVrrpgroupVirtualaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Virtualaddress *string `json:"virtual-address,omitempty"`
}

// InterfaceSubinterfaceIpv6 struct
type InterfaceSubinterfaceIpv6 struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=18
	Address             []*InterfaceSubinterfaceIpv6Address           `json:"address,omitempty"`
	Dhcpclient          *InterfaceSubinterfaceIpv6Dhcpclient          `json:"dhcp-client,omitempty"`
	Dhcprelay           *InterfaceSubinterfaceIpv6Dhcprelay           `json:"dhcp-relay,omitempty"`
	Dhcpv6server        *InterfaceSubinterfaceIpv6Dhcpv6server        `json:"dhcpv6-server,omitempty"`
	Neighbordiscovery   *InterfaceSubinterfaceIpv6Neighbordiscovery   `json:"neighbor-discovery,omitempty"`
	Routeradvertisement *InterfaceSubinterfaceIpv6Routeradvertisement `json:"router-advertisement,omitempty"`
	Vrrp                *InterfaceSubinterfaceIpv6Vrrp                `json:"vrrp,omitempty"`
}

// InterfaceSubinterfaceIpv6Address struct
type InterfaceSubinterfaceIpv6Address struct {
	Anycastgw *bool `json:"anycast-gw,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefix *string `json:"ip-prefix"`
	Primary  *string `json:"primary,omitempty"`
}

// InterfaceSubinterfaceIpv6Dhcpclient struct
type InterfaceSubinterfaceIpv6Dhcpclient struct {
	Traceoptions *InterfaceSubinterfaceIpv6DhcpclientTraceoptions `json:"trace-options,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcpclientTraceoptions struct
type InterfaceSubinterfaceIpv6DhcpclientTraceoptions struct {
	Trace *InterfaceSubinterfaceIpv6DhcpclientTraceoptionsTrace `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcpclientTraceoptionsTrace struct
type InterfaceSubinterfaceIpv6DhcpclientTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_InterfaceSubinterfaceIpv6DhcpclientTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv6Dhcprelay struct
type InterfaceSubinterfaceIpv6Dhcprelay struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_InterfaceSubinterfaceIpv6DhcprelayAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Option *InterfaceSubinterfaceIpv6DhcprelayOption `json:"option,omitempty"`
	Server *InterfaceSubinterfaceIpv6DhcprelayServer `json:"server,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Sourceaddress *string                                         `json:"source-address,omitempty"`
	Traceoptions  *InterfaceSubinterfaceIpv6DhcprelayTraceoptions `json:"trace-options,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcprelayOption struct
type InterfaceSubinterfaceIpv6DhcprelayOption struct {
	// +kubebuilder:validation:Enum=`interface-id`;`remote-id`
	Option E_InterfaceSubinterfaceIpv6DhcprelayOptionOption `json:"option,omitempty"`
	//Option *string `json:"option,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcprelayServer struct
type InterfaceSubinterfaceIpv6DhcprelayServer struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))|((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Server *string `json:"server,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcprelayTraceoptions struct
type InterfaceSubinterfaceIpv6DhcprelayTraceoptions struct {
	Trace *InterfaceSubinterfaceIpv6DhcprelayTraceoptionsTrace `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcprelayTraceoptionsTrace struct
type InterfaceSubinterfaceIpv6DhcprelayTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace E_InterfaceSubinterfaceIpv6DhcprelayTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv6Dhcpv6server struct
type InterfaceSubinterfaceIpv6Dhcpv6server struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_InterfaceSubinterfaceIpv6Dhcpv6serverAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// InterfaceSubinterfaceIpv6Neighbordiscovery struct
type InterfaceSubinterfaceIpv6Neighbordiscovery struct {
	Debug *InterfaceSubinterfaceIpv6NeighbordiscoveryDebug `json:"debug,omitempty"`
	// +kubebuilder:default:=true
	Duplicateaddressdetection *bool                                                `json:"duplicate-address-detection,omitempty"`
	Evpn                      *InterfaceSubinterfaceIpv6NeighbordiscoveryEvpn      `json:"evpn,omitempty"`
	Hostroute                 *InterfaceSubinterfaceIpv6NeighbordiscoveryHostroute `json:"host-route,omitempty"`
	// +kubebuilder:validation:Enum=`both`;`global`;`link-local`;`none`
	// +kubebuilder:default:="none"
	Learnunsolicited E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited `json:"learn-unsolicited,omitempty"`
	//Learnunsolicited *string `json:"learn-unsolicited,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Neighbor []*InterfaceSubinterfaceIpv6NeighbordiscoveryNeighbor `json:"neighbor,omitempty"`
	// kubebuilder:validation:Minimum=30
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=30
	Reachabletime *uint32 `json:"reachable-time,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=14400
	Staletime *uint32 `json:"stale-time,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighbordiscoveryDebug struct
type InterfaceSubinterfaceIpv6NeighbordiscoveryDebug struct {
	// +kubebuilder:validation:Enum=`messages`
	Debug E_InterfaceSubinterfaceIpv6NeighbordiscoveryDebugDebug `json:"debug,omitempty"`
	//Debug *string `json:"debug,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighbordiscoveryEvpn struct
type InterfaceSubinterfaceIpv6NeighbordiscoveryEvpn struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Advertise []*InterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertise `json:"advertise,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertise struct
type InterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertise struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	Admintag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`static`
	Routetype E_InterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertiseRoutetype `json:"route-type,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighbordiscoveryHostroute struct
type InterfaceSubinterfaceIpv6NeighbordiscoveryHostroute struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Populate []*InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulate `json:"populate,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulate struct
type InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulate struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Admintag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`evpn`;`static`
	Routetype E_InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulateRoutetype `json:"route-type,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighbordiscoveryNeighbor struct
type InterfaceSubinterfaceIpv6NeighbordiscoveryNeighbor struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6address *string `json:"ipv6-address"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Linklayeraddress *string `json:"link-layer-address"`
}

// InterfaceSubinterfaceIpv6Routeradvertisement struct
type InterfaceSubinterfaceIpv6Routeradvertisement struct {
	Debug      *InterfaceSubinterfaceIpv6RouteradvertisementDebug      `json:"debug,omitempty"`
	Routerrole *InterfaceSubinterfaceIpv6RouteradvertisementRouterrole `json:"router-role,omitempty"`
}

// InterfaceSubinterfaceIpv6RouteradvertisementDebug struct
type InterfaceSubinterfaceIpv6RouteradvertisementDebug struct {
	// +kubebuilder:validation:Enum=`messages`
	Debug E_InterfaceSubinterfaceIpv6RouteradvertisementDebugDebug `json:"debug,omitempty"`
	//Debug *string `json:"debug,omitempty"`
}

// InterfaceSubinterfaceIpv6RouteradvertisementRouterrole struct
type InterfaceSubinterfaceIpv6RouteradvertisementRouterrole struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_InterfaceSubinterfaceIpv6RouteradvertisementRouterroleAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=64
	Currenthoplimit *uint8 `json:"current-hop-limit,omitempty"`
	// kubebuilder:validation:Minimum=1280
	// kubebuilder:validation:Maximum=9486
	Ipmtu *uint16 `json:"ip-mtu,omitempty"`
	// +kubebuilder:default:=false
	Managedconfigurationflag *bool `json:"managed-configuration-flag,omitempty"`
	// kubebuilder:validation:Minimum=4
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=600
	Maxadvertisementinterval *uint16 `json:"max-advertisement-interval,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=1350
	// +kubebuilder:default:=200
	Minadvertisementinterval *uint16 `json:"min-advertisement-interval,omitempty"`
	// +kubebuilder:default:=false
	Otherconfigurationflag *bool `json:"other-configuration-flag,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=16
	Prefix []*InterfaceSubinterfaceIpv6RouteradvertisementRouterrolePrefix `json:"prefix,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600000
	// +kubebuilder:default:=0
	Reachabletime *uint32 `json:"reachable-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1800000
	// +kubebuilder:default:=0
	Retransmittime *uint32 `json:"retransmit-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=9000
	// +kubebuilder:default:=1800
	Routerlifetime *uint16 `json:"router-lifetime,omitempty"`
}

// InterfaceSubinterfaceIpv6RouteradvertisementRouterrolePrefix struct
type InterfaceSubinterfaceIpv6RouteradvertisementRouterrolePrefix struct {
	// +kubebuilder:default:=true
	Autonomousflag *bool `json:"autonomous-flag,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipv6prefix *string `json:"ipv6-prefix"`
	// +kubebuilder:default:=true
	Onlinkflag *bool `json:"on-link-flag,omitempty"`
	// +kubebuilder:default:=604800
	Preferredlifetime *uint32 `json:"preferred-lifetime,omitempty"`
	// +kubebuilder:default:=2592000
	Validlifetime *uint32 `json:"valid-lifetime,omitempty"`
}

// InterfaceSubinterfaceIpv6Vrrp struct
type InterfaceSubinterfaceIpv6Vrrp struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Vrrpgroup []*InterfaceSubinterfaceIpv6VrrpVrrpgroup `json:"vrrp-group,omitempty"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpgroup struct
type InterfaceSubinterfaceIpv6VrrpVrrpgroup struct {
	Acceptmode *bool `json:"accept-mode,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_InterfaceSubinterfaceIpv6VrrpVrrpgroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1000
	Advertiseinterval *uint16                                               `json:"advertise-interval,omitempty"`
	Authentication    *InterfaceSubinterfaceIpv6VrrpVrrpgroupAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Initdelay         *uint16                                                  `json:"init-delay,omitempty"`
	Interfacetracking *InterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetracking `json:"interface-tracking,omitempty"`
	// +kubebuilder:default:=false
	Masterinheritinterval *bool `json:"master-inherit-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Operinterval *uint16 `json:"oper-interval,omitempty"`
	Preempt      *bool   `json:"preempt,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Preemptdelay *uint16 `json:"preempt-delay,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=100
	Priority   *uint8                                            `json:"priority,omitempty"`
	Statistics *InterfaceSubinterfaceIpv6VrrpVrrpgroupStatistics `json:"statistics,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=3
	Version        *uint8                                                `json:"version,omitempty"`
	Virtualaddress *InterfaceSubinterfaceIpv6VrrpVrrpgroupVirtualaddress `json:"virtual-address,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Virtualrouterid *uint8 `json:"virtual-router-id"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpgroupAuthentication struct
type InterfaceSubinterfaceIpv6VrrpVrrpgroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetracking struct
type InterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetracking struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Trackinterface []*InterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetrackingTrackinterface `json:"track-interface,omitempty"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetrackingTrackinterface struct
type InterfaceSubinterfaceIpv6VrrpVrrpgroupInterfacetrackingTrackinterface struct {
	Interface *string `json:"interface"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Prioritydecrement *uint8 `json:"priority-decrement,omitempty"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpgroupStatistics struct
type InterfaceSubinterfaceIpv6VrrpVrrpgroupStatistics struct {
}

// InterfaceSubinterfaceIpv6VrrpVrrpgroupVirtualaddress struct
type InterfaceSubinterfaceIpv6VrrpVrrpgroupVirtualaddress struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Virtualaddress *string `json:"virtual-address,omitempty"`
}

// InterfaceSubinterfaceLocalmirrordestination struct
type InterfaceSubinterfaceLocalmirrordestination struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_InterfaceSubinterfaceLocalmirrordestinationAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// InterfaceSubinterfaceQos struct
type InterfaceSubinterfaceQos struct {
	Input  *InterfaceSubinterfaceQosInput  `json:"input,omitempty"`
	Output *InterfaceSubinterfaceQosOutput `json:"output,omitempty"`
}

// InterfaceSubinterfaceQosInput struct
type InterfaceSubinterfaceQosInput struct {
	Classifiers *InterfaceSubinterfaceQosInputClassifiers `json:"classifiers,omitempty"`
}

// InterfaceSubinterfaceQosInputClassifiers struct
type InterfaceSubinterfaceQosInputClassifiers struct {
	Dscp             *string `json:"dscp,omitempty"`
	Ipv4dscp         *string `json:"ipv4-dscp,omitempty"`
	Ipv6dscp         *string `json:"ipv6-dscp,omitempty"`
	Mplstrafficclass *string `json:"mpls-traffic-class,omitempty"`
}

// InterfaceSubinterfaceQosOutput struct
type InterfaceSubinterfaceQosOutput struct {
	Rewriterules *InterfaceSubinterfaceQosOutputRewriterules `json:"rewrite-rules,omitempty"`
}

// InterfaceSubinterfaceQosOutputRewriterules struct
type InterfaceSubinterfaceQosOutputRewriterules struct {
	Dscp             *string `json:"dscp,omitempty"`
	Ipv4dscp         *string `json:"ipv4-dscp,omitempty"`
	Ipv6dscp         *string `json:"ipv6-dscp,omitempty"`
	Mplstrafficclass *string `json:"mpls-traffic-class,omitempty"`
}

// InterfaceSubinterfaceRaguard struct
type InterfaceSubinterfaceRaguard struct {
	Policy *string `json:"policy,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Vlanlist []*InterfaceSubinterfaceRaguardVlanlist `json:"vlan-list,omitempty"`
}

// InterfaceSubinterfaceRaguardVlanlist struct
type InterfaceSubinterfaceRaguardVlanlist struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4095
	Vlanid *uint16 `json:"vlan-id"`
}

// InterfaceSubinterfaceVlan struct
type InterfaceSubinterfaceVlan struct {
	Encap *InterfaceSubinterfaceVlanEncap `json:"encap,omitempty"`
}

// InterfaceSubinterfaceVlanEncap struct
type InterfaceSubinterfaceVlanEncap struct {
	Singletagged *InterfaceSubinterfaceVlanEncapSingletagged `json:"single-tagged,omitempty"`
	Untagged     *InterfaceSubinterfaceVlanEncapUntagged     `json:"untagged,omitempty"`
}

// InterfaceSubinterfaceVlanEncapSingletagged struct
type InterfaceSubinterfaceVlanEncapSingletagged struct {
	Vlanid *string `json:"vlan-id,omitempty"`
}

// InterfaceSubinterfaceVlanEncapUntagged struct
type InterfaceSubinterfaceVlanEncapUntagged struct {
}

// A InterfaceSubinterfaceSpec defines the desired state of a InterfaceSubinterface.
type InterfaceSubinterfaceSpec struct {
	nddv1.ResourceSpec    `json:",inline"`
	InterfaceName         *string                `json:"interface-name"`
	InterfaceSubinterface *InterfaceSubinterface `json:"subinterface,omitempty"`
}

// A InterfaceSubinterfaceStatus represents the observed state of a InterfaceSubinterface.
type InterfaceSubinterfaceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlInterfaceSubinterface is the Schema for the InterfaceSubinterface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlInterfaceSubinterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InterfaceSubinterfaceSpec   `json:"spec,omitempty"`
	Status InterfaceSubinterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlInterfaceSubinterfaceList contains a list of InterfaceSubinterfaces
type SrlInterfaceSubinterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlInterfaceSubinterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlInterfaceSubinterface{}, &SrlInterfaceSubinterfaceList{})
}

// InterfaceSubinterface type metadata.
var (
	InterfaceSubinterfaceKindKind         = reflect.TypeOf(SrlInterfaceSubinterface{}).Name()
	InterfaceSubinterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: InterfaceSubinterfaceKindKind}.String()
	InterfaceSubinterfaceKindAPIVersion   = InterfaceSubinterfaceKindKind + "." + GroupVersion.String()
	InterfaceSubinterfaceGroupVersionKind = GroupVersion.WithKind(InterfaceSubinterfaceKindKind)
)
