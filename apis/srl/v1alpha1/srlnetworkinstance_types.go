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

// Networkinstance struct
type Networkinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Bridgetable *NetworkinstanceBridgetable `json:"bridge-table,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface       []*NetworkinstanceInterface     `json:"interface,omitempty"`
	Ipforwarding    *NetworkinstanceIpforwarding    `json:"ip-forwarding,omitempty"`
	Iploadbalancing *NetworkinstanceIploadbalancing `json:"ip-load-balancing,omitempty"`
	Mpls            *NetworkinstanceMpls            `json:"mpls,omitempty"`
	Mtu             *NetworkinstanceMtu             `json:"mtu,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name      *string                   `json:"name"`
	Protocols *NetworkinstanceProtocols `json:"protocols,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Routerid           *string                            `json:"router-id,omitempty"`
	Segmentrouting     *NetworkinstanceSegmentrouting     `json:"segment-routing,omitempty"`
	Tepolicies         *NetworkinstanceTepolicies         `json:"te-policies,omitempty"`
	Trafficengineering *NetworkinstanceTrafficengineering `json:"traffic-engineering,omitempty"`
	// +kubebuilder:default:="default"
	Type *string `json:"type,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Vxlaninterface []*NetworkinstanceVxlaninterface `json:"vxlan-interface,omitempty"`
}

// NetworkinstanceBridgetable struct
type NetworkinstanceBridgetable struct {
	// +kubebuilder:default:=false
	Discardunknowndestmac *bool                                     `json:"discard-unknown-dest-mac,omitempty"`
	Macduplication        *NetworkinstanceBridgetableMacduplication `json:"mac-duplication,omitempty"`
	Maclearning           *NetworkinstanceBridgetableMaclearning    `json:"mac-learning,omitempty"`
	Maclimit              *NetworkinstanceBridgetableMaclimit       `json:"mac-limit,omitempty"`
	// +kubebuilder:default:=false
	Protectanycastgwmac *bool                                `json:"protect-anycast-gw-mac,omitempty"`
	Staticmac           *NetworkinstanceBridgetableStaticmac `json:"static-mac,omitempty"`
}

// NetworkinstanceBridgetableMacduplication struct
type NetworkinstanceBridgetableMacduplication struct {
	// +kubebuilder:validation:Enum=`blackhole`;`oper-down`;`stop-learning`
	// +kubebuilder:default:="stop-learning"
	Action E_NetworkinstanceBridgetableMacduplicationAction `json:"action,omitempty"`
	//Action *string `json:"action,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceBridgetableMacduplicationAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=9
	Holddowntime *uint32 `json:"hold-down-time,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=15
	// +kubebuilder:default:=3
	Monitoringwindow *uint32 `json:"monitoring-window,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=10
	// +kubebuilder:default:=5
	Nummoves *uint32 `json:"num-moves,omitempty"`
}

// NetworkinstanceBridgetableMaclearning struct
type NetworkinstanceBridgetableMaclearning struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceBridgetableMaclearningAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Aging *NetworkinstanceBridgetableMaclearningAging `json:"aging,omitempty"`
}

// NetworkinstanceBridgetableMaclearningAging struct
type NetworkinstanceBridgetableMaclearningAging struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceBridgetableMaclearningAgingAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=86400
	// +kubebuilder:default:=300
	Agetime *int32 `json:"age-time,omitempty"`
}

// NetworkinstanceBridgetableMaclimit struct
type NetworkinstanceBridgetableMaclimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=250
	Maximumentries *int32 `json:"maximum-entries,omitempty"`
	// kubebuilder:validation:Minimum=6
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=95
	Warningthresholdpct *int32 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceBridgetableStaticmac struct
type NetworkinstanceBridgetableStaticmac struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Mac []*NetworkinstanceBridgetableStaticmacMac `json:"mac,omitempty"`
}

// NetworkinstanceBridgetableStaticmacMac struct
type NetworkinstanceBridgetableStaticmacMac struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Address     *string `json:"address"`
	Destination *string `json:"destination"`
}

// NetworkinstanceInterface struct
type NetworkinstanceInterface struct {
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0\.0|system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name"`
}

// NetworkinstanceIpforwarding struct
type NetworkinstanceIpforwarding struct {
	Receiveipv4check *bool `json:"receive-ipv4-check,omitempty"`
	Receiveipv6check *bool `json:"receive-ipv6-check,omitempty"`
}

// NetworkinstanceIploadbalancing struct
type NetworkinstanceIploadbalancing struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Resilienthashprefix []*NetworkinstanceIploadbalancingResilienthashprefix `json:"resilient-hash-prefix,omitempty"`
}

// NetworkinstanceIploadbalancingResilienthashprefix struct
type NetworkinstanceIploadbalancingResilienthashprefix struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=32
	// +kubebuilder:default:=1
	Hashbucketsperpath *uint8 `json:"hash-buckets-per-path,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefix *string `json:"ip-prefix"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpaths *uint8 `json:"max-paths,omitempty"`
}

// NetworkinstanceMpls struct
type NetworkinstanceMpls struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_NetworkinstanceMplsAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Staticentry      []*NetworkinstanceMplsStaticentry `json:"static-entry,omitempty"`
	Staticlabelblock *string                           `json:"static-label-block,omitempty"`
}

// NetworkinstanceMplsStaticentry struct
type NetworkinstanceMplsStaticentry struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceMplsStaticentryAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Collectstats *bool   `json:"collect-stats,omitempty"`
	Nexthopgroup *string `json:"next-hop-group,omitempty"`
	// +kubebuilder:validation:Enum=`pop`;`swap`
	// +kubebuilder:default:="swap"
	Operation E_NetworkinstanceMplsStaticentryOperation `json:"operation,omitempty"`
	//Operation *string `json:"operation,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Preference *uint8  `json:"preference"`
	Toplabel   *string `json:"top-label"`
}

// NetworkinstanceMtu struct
type NetworkinstanceMtu struct {
	Pathmtudiscovery *bool `json:"path-mtu-discovery,omitempty"`
}

// NetworkinstanceProtocols struct
type NetworkinstanceProtocols struct {
	Directlyconnected *NetworkinstanceProtocolsDirectlyconnected `json:"directly-connected,omitempty"`
	Ldp               *NetworkinstanceProtocolsLdp               `json:"ldp,omitempty"`
}

// NetworkinstanceProtocolsDirectlyconnected struct
type NetworkinstanceProtocolsDirectlyconnected struct {
	Tedatabaseinstall *NetworkinstanceProtocolsDirectlyconnectedTedatabaseinstall `json:"te-database-install,omitempty"`
}

// NetworkinstanceProtocolsDirectlyconnectedTedatabaseinstall struct
type NetworkinstanceProtocolsDirectlyconnectedTedatabaseinstall struct {
	Bgpls *NetworkinstanceProtocolsDirectlyconnectedTedatabaseinstallBgpls `json:"bgp-ls,omitempty"`
}

// NetworkinstanceProtocolsDirectlyconnectedTedatabaseinstallBgpls struct
type NetworkinstanceProtocolsDirectlyconnectedTedatabaseinstallBgpls struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Bgplsidentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Igpidentifier *uint64 `json:"igp-identifier,omitempty"`
}

// NetworkinstanceProtocolsLdp struct
type NetworkinstanceProtocolsLdp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_NetworkinstanceProtocolsLdpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Discovery         *NetworkinstanceProtocolsLdpDiscovery       `json:"discovery,omitempty"`
	Dynamiclabelblock *string                                     `json:"dynamic-label-block"`
	Gracefulrestart   *NetworkinstanceProtocolsLdpGracefulrestart `json:"graceful-restart,omitempty"`
	Ipv4              *NetworkinstanceProtocolsLdpIpv4            `json:"ipv4,omitempty"`
	Multipath         *NetworkinstanceProtocolsLdpMultipath       `json:"multipath,omitempty"`
	Peers             *NetworkinstanceProtocolsLdpPeers           `json:"peers,omitempty"`
	Traceoptions      *NetworkinstanceProtocolsLdpTraceoptions    `json:"trace-options,omitempty"`
}

// NetworkinstanceProtocolsLdpDiscovery struct
type NetworkinstanceProtocolsLdpDiscovery struct {
	Interfaces *NetworkinstanceProtocolsLdpDiscoveryInterfaces `json:"interfaces,omitempty"`
}

// NetworkinstanceProtocolsLdpDiscoveryInterfaces struct
type NetworkinstanceProtocolsLdpDiscoveryInterfaces struct {
	// kubebuilder:validation:Minimum=15
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=15
	Helloholdtime *uint16 `json:"hello-holdtime,omitempty"`
	// kubebuilder:validation:Minimum=5
	// kubebuilder:validation:Maximum=1200
	// +kubebuilder:default:=5
	Hellointerval *uint16 `json:"hello-interval,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*NetworkinstanceProtocolsLdpDiscoveryInterfacesInterface `json:"interface,omitempty"`
}

// NetworkinstanceProtocolsLdpDiscoveryInterfacesInterface struct
type NetworkinstanceProtocolsLdpDiscoveryInterfacesInterface struct {
	// kubebuilder:validation:Minimum=15
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=15
	Helloholdtime *uint16 `json:"hello-holdtime,omitempty"`
	// kubebuilder:validation:Minimum=5
	// kubebuilder:validation:Maximum=1200
	// +kubebuilder:default:=5
	Hellointerval *uint16                                                      `json:"hello-interval,omitempty"`
	Ipv4          *NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4 `json:"ipv4,omitempty"`
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name"`
}

// NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4 struct
type NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4 struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Adminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Statistics *NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Statistics `json:"statistics,omitempty"`
}

// NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Statistics struct
type NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Statistics struct {
	Hellomessageerrors *NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4StatisticsHellomessageerrors `json:"hello-message-errors,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Helloreceived *uint64 `json:"hello-received,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Hellosent *uint64 `json:"hello-sent,omitempty"`
}

// NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4StatisticsHellomessageerrors struct
type NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4StatisticsHellomessageerrors struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Badmessagelength *uint64 `json:"bad-message-length,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Badpdulength *uint64 `json:"bad-pdu-length,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Badprotocolversion *uint64 `json:"bad-protocol-version,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Malformedtlvvalue *uint64 `json:"malformed-tlv-value,omitempty"`
}

// NetworkinstanceProtocolsLdpGracefulrestart struct
type NetworkinstanceProtocolsLdpGracefulrestart struct {
	// +kubebuilder:default:=false
	Helperenable *bool `json:"helper-enable,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=120
	Maxreconnecttime *uint16 `json:"max-reconnect-time,omitempty"`
	// kubebuilder:validation:Minimum=30
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=120
	Maxrecoverytime *uint16 `json:"max-recovery-time,omitempty"`
}

// NetworkinstanceProtocolsLdpIpv4 struct
type NetworkinstanceProtocolsLdpIpv4 struct {
	Fecresolution *NetworkinstanceProtocolsLdpIpv4Fecresolution `json:"fec-resolution,omitempty"`
}

// NetworkinstanceProtocolsLdpIpv4Fecresolution struct
type NetworkinstanceProtocolsLdpIpv4Fecresolution struct {
	// +kubebuilder:default:=false
	Longestprefix *bool `json:"longest-prefix,omitempty"`
}

// NetworkinstanceProtocolsLdpMultipath struct
type NetworkinstanceProtocolsLdpMultipath struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=32
	Maxpaths *uint8 `json:"max-paths,omitempty"`
}

// NetworkinstanceProtocolsLdpPeers struct
type NetworkinstanceProtocolsLdpPeers struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Peer []*NetworkinstanceProtocolsLdpPeersPeer `json:"peer,omitempty"`
	// kubebuilder:validation:Minimum=45
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=180
	Sessionkeepaliveholdtime *uint16 `json:"session-keepalive-holdtime,omitempty"`
	// kubebuilder:validation:Minimum=15
	// kubebuilder:validation:Maximum=1200
	// +kubebuilder:default:=60
	Sessionkeepaliveinterval *uint16                                       `json:"session-keepalive-interval,omitempty"`
	Tcptransport             *NetworkinstanceProtocolsLdpPeersTcptransport `json:"tcp-transport,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeer struct
type NetworkinstanceProtocolsLdpPeersPeer struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	Feclimit *uint32 `json:"fec-limit,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Labelspaceid *uint16 `json:"label-space-id"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Lsrid                *string                                                   `json:"lsr-id"`
	Overload             *NetworkinstanceProtocolsLdpPeersPeerOverload             `json:"overload,omitempty"`
	Receivedcapabilities *NetworkinstanceProtocolsLdpPeersPeerReceivedcapabilities `json:"received-capabilities,omitempty"`
	Tcptransport         *NetworkinstanceProtocolsLdpPeersPeerTcptransport         `json:"tcp-transport,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeerOverload struct
type NetworkinstanceProtocolsLdpPeersPeerOverload struct {
	Localrouterisoverloaded *bool `json:"local-router-is-overloaded,omitempty"`
	Peerisoverloaded        *bool `json:"peer-is-overloaded,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeerReceivedcapabilities struct
type NetworkinstanceProtocolsLdpPeersPeerReceivedcapabilities struct {
	Dualstackcapability                *bool                                                                              `json:"dual-stack-capability,omitempty"`
	Dynamiccapability                  *bool                                                                              `json:"dynamic-capability,omitempty"`
	Entropylabelcapability             *bool                                                                              `json:"entropy-label-capability,omitempty"`
	Gracefulrestartcapability          *bool                                                                              `json:"graceful-restart-capability,omitempty"`
	Makebeforebreakcapability          *bool                                                                              `json:"make-before-break-capability,omitempty"`
	Multipointtomultipointcapability   *bool                                                                              `json:"multipoint-to-multipoint-capability,omitempty"`
	Nokiavendoroverloadcapability      *bool                                                                              `json:"nokia-vendor-overload-capability,omitempty"`
	Pointtomultipointcapability        *bool                                                                              `json:"point-to-multipoint-capability,omitempty"`
	Stateadvertisementcontrol          *NetworkinstanceProtocolsLdpPeersPeerReceivedcapabilitiesStateadvertisementcontrol `json:"state-advertisement-control,omitempty"`
	Unrecognizednotificationcapability *bool                                                                              `json:"unrecognized-notification-capability,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeerReceivedcapabilitiesStateadvertisementcontrol struct
type NetworkinstanceProtocolsLdpPeersPeerReceivedcapabilitiesStateadvertisementcontrol struct {
	Ipv4prefixdisable          *bool `json:"ipv4-prefix-disable,omitempty"`
	Ipv6prefixdisable          *bool `json:"ipv6-prefix-disable,omitempty"`
	P2ppseudowirefec128disable *bool `json:"p2p-pseudowire-fec-128-disable,omitempty"`
	P2ppseudowirefec129disable *bool `json:"p2p-pseudowire-fec-129-disable,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeerTcptransport struct
type NetworkinstanceProtocolsLdpPeersPeerTcptransport struct {
	Authentication *NetworkinstanceProtocolsLdpPeersPeerTcptransportAuthentication `json:"authentication,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeerTcptransportAuthentication struct
type NetworkinstanceProtocolsLdpPeersPeerTcptransportAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersTcptransport struct
type NetworkinstanceProtocolsLdpPeersTcptransport struct {
	Authentication *NetworkinstanceProtocolsLdpPeersTcptransportAuthentication `json:"authentication,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersTcptransportAuthentication struct
type NetworkinstanceProtocolsLdpPeersTcptransportAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsLdpTraceoptions struct
type NetworkinstanceProtocolsLdpTraceoptions struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*NetworkinstanceProtocolsLdpTraceoptionsInterface `json:"interface,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Peer []*NetworkinstanceProtocolsLdpTraceoptionsPeer `json:"peer,omitempty"`
}

// NetworkinstanceProtocolsLdpTraceoptionsInterface struct
type NetworkinstanceProtocolsLdpTraceoptionsInterface struct {
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name"`
}

// NetworkinstanceProtocolsLdpTraceoptionsPeer struct
type NetworkinstanceProtocolsLdpTraceoptionsPeer struct {
	Labelspaceid *string `json:"label-space-id"`
	Lsrid        *string `json:"lsr-id"`
}

// NetworkinstanceSegmentrouting struct
type NetworkinstanceSegmentrouting struct {
	Mpls *NetworkinstanceSegmentroutingMpls `json:"mpls,omitempty"`
}

// NetworkinstanceSegmentroutingMpls struct
type NetworkinstanceSegmentroutingMpls struct {
	Globalblock *NetworkinstanceSegmentroutingMplsGlobalblock `json:"global-block,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=4
	Prefixsid []*NetworkinstanceSegmentroutingMplsPrefixsid `json:"prefix-sid,omitempty"`
}

// NetworkinstanceSegmentroutingMplsGlobalblock struct
type NetworkinstanceSegmentroutingMplsGlobalblock struct {
	Labelrange *string `json:"label-range,omitempty"`
}

// NetworkinstanceSegmentroutingMplsPrefixsid struct
type NetworkinstanceSegmentroutingMplsPrefixsid struct {
	Interface *string `json:"interface"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Ipv4labelindex *uint32 `json:"ipv4-label-index,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Ipv6labelindex *uint32 `json:"ipv6-label-index,omitempty"`
	// +kubebuilder:default:=true
	Nodesid *bool `json:"node-sid,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4
	Prefixsidindex *uint8 `json:"prefix-sid-index"`
}

// NetworkinstanceTepolicies struct
type NetworkinstanceTepolicies struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Staticpolicy []*NetworkinstanceTepoliciesStaticpolicy `json:"static-policy,omitempty"`
}

// NetworkinstanceTepoliciesStaticpolicy struct
type NetworkinstanceTepoliciesStaticpolicy struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceTepoliciesStaticpolicyAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Color *uint32 `json:"color,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Endpoint *string `json:"endpoint"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=100
	Preference *uint32 `json:"preference,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Segmentlist []*NetworkinstanceTepoliciesStaticpolicySegmentlist `json:"segment-list,omitempty"`
}

// NetworkinstanceTepoliciesStaticpolicySegmentlist struct
type NetworkinstanceTepoliciesStaticpolicySegmentlist struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Segment []*NetworkinstanceTepoliciesStaticpolicySegmentlistSegment `json:"segment,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Segmentlistindex *uint8 `json:"segment-list-index"`
}

// NetworkinstanceTepoliciesStaticpolicySegmentlistSegment struct
type NetworkinstanceTepoliciesStaticpolicySegmentlistSegment struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Segmentindex *uint8                                                               `json:"segment-index"`
	Segmenttype1 *NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentSegmenttype1 `json:"segment-type-1,omitempty"`
	// +kubebuilder:validation:Enum=`segment-type-1`;`segment-type-10`;`segment-type-11`;`segment-type-2`;`segment-type-3`;`segment-type-4`;`segment-type-5`;`segment-type-6`;`segment-type-7`;`segment-type-8`;`segment-type-9`
	Type E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentSegmenttype1 struct
type NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentSegmenttype1 struct {
	Sidvalue *string `json:"sid-value,omitempty"`
}

// NetworkinstanceTrafficengineering struct
type NetworkinstanceTrafficengineering struct {
	Admingroups *NetworkinstanceTrafficengineeringAdmingroups `json:"admin-groups,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Autonomoussystem *uint32 `json:"autonomous-system,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*NetworkinstanceTrafficengineeringInterface `json:"interface,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4terouterid *string `json:"ipv4-te-router-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6terouterid       *string                                                `json:"ipv6-te-router-id,omitempty"`
	Sharedrisklinkgroups *NetworkinstanceTrafficengineeringSharedrisklinkgroups `json:"shared-risk-link-groups,omitempty"`
}

// NetworkinstanceTrafficengineeringAdmingroups struct
type NetworkinstanceTrafficengineeringAdmingroups struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group []*NetworkinstanceTrafficengineeringAdmingroupsGroup `json:"group,omitempty"`
}

// NetworkinstanceTrafficengineeringAdmingroupsGroup struct
type NetworkinstanceTrafficengineeringAdmingroupsGroup struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=31
	Bitposition *uint32 `json:"bit-position"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// NetworkinstanceTrafficengineeringInterface struct
type NetworkinstanceTrafficengineeringInterface struct {
	Admingroup     *NetworkinstanceTrafficengineeringInterfaceAdmingroup     `json:"admin-group,omitempty"`
	Delay          *NetworkinstanceTrafficengineeringInterfaceDelay          `json:"delay,omitempty"`
	Interfacename  *string                                                   `json:"interface-name"`
	Srlgmembership *NetworkinstanceTrafficengineeringInterfaceSrlgmembership `json:"srlg-membership,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Temetric *uint32 `json:"te-metric,omitempty"`
}

// NetworkinstanceTrafficengineeringInterfaceAdmingroup struct
type NetworkinstanceTrafficengineeringInterfaceAdmingroup struct {
	Admingroup *string `json:"admin-group,omitempty"`
}

// NetworkinstanceTrafficengineeringInterfaceDelay struct
type NetworkinstanceTrafficengineeringInterfaceDelay struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Static *uint32 `json:"static,omitempty"`
}

// NetworkinstanceTrafficengineeringInterfaceSrlgmembership struct
type NetworkinstanceTrafficengineeringInterfaceSrlgmembership struct {
	Srlgmembership *string `json:"srlg-membership,omitempty"`
}

// NetworkinstanceTrafficengineeringSharedrisklinkgroups struct
type NetworkinstanceTrafficengineeringSharedrisklinkgroups struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group []*NetworkinstanceTrafficengineeringSharedrisklinkgroupsGroup `json:"group,omitempty"`
}

// NetworkinstanceTrafficengineeringSharedrisklinkgroupsGroup struct
type NetworkinstanceTrafficengineeringSharedrisklinkgroupsGroup struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Cost *uint32 `json:"cost,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Staticmember []*NetworkinstanceTrafficengineeringSharedrisklinkgroupsGroupStaticmember `json:"static-member,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Value *uint32 `json:"value"`
}

// NetworkinstanceTrafficengineeringSharedrisklinkgroupsGroupStaticmember struct
type NetworkinstanceTrafficengineeringSharedrisklinkgroupsGroupStaticmember struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Fromaddress *string `json:"from-address"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Toaddress *string `json:"to-address,omitempty"`
}

// NetworkinstanceVxlaninterface struct
type NetworkinstanceVxlaninterface struct {
	// kubebuilder:validation:MinLength=8
	// kubebuilder:validation:MaxLength=17
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,8}))`
	Name *string `json:"name"`
}

// A NetworkinstanceSpec defines the desired state of a Networkinstance.
type NetworkinstanceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	Networkinstance    *Networkinstance `json:"network-instance,omitempty"`
}

// A NetworkinstanceStatus represents the observed state of a Networkinstance.
type NetworkinstanceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstance is the Schema for the Networkinstance API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl2}
type SrlNetworkinstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceSpec   `json:"spec,omitempty"`
	Status NetworkinstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceList contains a list of Networkinstances
type SrlNetworkinstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstance{}, &SrlNetworkinstanceList{})
}

// Networkinstance type metadata.
var (
	NetworkinstanceKindKind         = reflect.TypeOf(SrlNetworkinstance{}).Name()
	NetworkinstanceGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceKindKind}.String()
	NetworkinstanceKindAPIVersion   = NetworkinstanceKindKind + "." + GroupVersion.String()
	NetworkinstanceGroupVersionKind = GroupVersion.WithKind(NetworkinstanceKindKind)
)
