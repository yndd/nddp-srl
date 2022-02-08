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

// NetworkinstanceProtocolsIsis struct
type NetworkinstanceProtocolsIsis struct {
	Dynamiclabelblock *string `json:"dynamic-label-block,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Instance []*NetworkinstanceProtocolsIsisInstance `json:"instance,omitempty"`
}

// NetworkinstanceProtocolsIsisInstance struct
type NetworkinstanceProtocolsIsisInstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_NetworkinstanceProtocolsIsisInstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Attachedbit                   *NetworkinstanceProtocolsIsisInstanceAttachedbit                   `json:"attached-bit,omitempty"`
	Authentication                *NetworkinstanceProtocolsIsisInstanceAuthentication                `json:"authentication,omitempty"`
	Autocost                      *NetworkinstanceProtocolsIsisInstanceAutocost                      `json:"auto-cost,omitempty"`
	Exportpolicy                  *string                                                            `json:"export-policy,omitempty"`
	Gracefulrestart               *NetworkinstanceProtocolsIsisInstanceGracefulrestart               `json:"graceful-restart,omitempty"`
	Interlevelpropagationpolicies *NetworkinstanceProtocolsIsisInstanceInterlevelpropagationpolicies `json:"inter-level-propagation-policies,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface          []*NetworkinstanceProtocolsIsisInstanceInterface        `json:"interface,omitempty"`
	Ipv4unicast        *NetworkinstanceProtocolsIsisInstanceIpv4unicast        `json:"ipv4-unicast,omitempty"`
	Ipv6unicast        *NetworkinstanceProtocolsIsisInstanceIpv6unicast        `json:"ipv6-unicast,omitempty"`
	Ldpsynchronization *NetworkinstanceProtocolsIsisInstanceLdpsynchronization `json:"ldp-synchronization,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=2
	Level []*NetworkinstanceProtocolsIsisInstanceLevel `json:"level,omitempty"`
	// +kubebuilder:validation:Enum=`L1`;`L1L2`;`L2`
	// +kubebuilder:default:="L2"
	Levelcapability E_NetworkinstanceProtocolsIsisInstanceLevelcapability `json:"level-capability,omitempty"`
	//Levelcapability *string `json:"level-capability,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxecmppaths *uint8 `json:"max-ecmp-paths,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name     *string                                       `json:"name"`
	Net      *NetworkinstanceProtocolsIsisInstanceNet      `json:"net,omitempty"`
	Overload *NetworkinstanceProtocolsIsisInstanceOverload `json:"overload,omitempty"`
	// +kubebuilder:default:=false
	Poitlv             *bool                                                   `json:"poi-tlv,omitempty"`
	Segmentrouting     *NetworkinstanceProtocolsIsisInstanceSegmentrouting     `json:"segment-routing,omitempty"`
	Tedatabaseinstall  *NetworkinstanceProtocolsIsisInstanceTedatabaseinstall  `json:"te-database-install,omitempty"`
	Timers             *NetworkinstanceProtocolsIsisInstanceTimers             `json:"timers,omitempty"`
	Traceoptions       *NetworkinstanceProtocolsIsisInstanceTraceoptions       `json:"trace-options,omitempty"`
	Trafficengineering *NetworkinstanceProtocolsIsisInstanceTrafficengineering `json:"traffic-engineering,omitempty"`
	Transport          *NetworkinstanceProtocolsIsisInstanceTransport          `json:"transport,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceAttachedbit struct
type NetworkinstanceProtocolsIsisInstanceAttachedbit struct {
	// +kubebuilder:default:=false
	Ignore *bool `json:"ignore,omitempty"`
	// +kubebuilder:default:=false
	Suppress *bool `json:"suppress,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceAuthentication struct
type NetworkinstanceProtocolsIsisInstanceAuthentication struct {
	Csnpauthentication  *bool   `json:"csnp-authentication,omitempty"`
	Helloauthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
	Psnpauthentication  *bool   `json:"psnp-authentication,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceAutocost struct
type NetworkinstanceProtocolsIsisInstanceAutocost struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8000000000
	Referencebandwidth *uint64 `json:"reference-bandwidth,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceGracefulrestart struct
type NetworkinstanceProtocolsIsisInstanceGracefulrestart struct {
	// +kubebuilder:default:=false
	Helpermode *bool `json:"helper-mode,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterlevelpropagationpolicies struct
type NetworkinstanceProtocolsIsisInstanceInterlevelpropagationpolicies struct {
	Level1tolevel2 *NetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2 `json:"level1-to-level2,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2 struct
type NetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2 struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Summaryaddress []*NetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2Summaryaddress `json:"summary-address,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2Summaryaddress struct
type NetworkinstanceProtocolsIsisInstanceInterlevelpropagationpoliciesLevel1tolevel2Summaryaddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefix *string `json:"ip-prefix"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Routetag *uint32 `json:"route-tag,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterface struct
type NetworkinstanceProtocolsIsisInstanceInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsIsisInstanceInterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Authentication *NetworkinstanceProtocolsIsisInstanceInterfaceAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	Circuittype E_NetworkinstanceProtocolsIsisInstanceInterfaceCircuittype `json:"circuit-type,omitempty"`
	//Circuittype *string `json:"circuit-type,omitempty"`
	// +kubebuilder:validation:Enum=`adaptive`;`disable`;`loose`;`strict`
	// +kubebuilder:default:="disable"
	Hellopadding E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding `json:"hello-padding,omitempty"`
	//Hellopadding *string `json:"hello-padding,omitempty"`
	Interfacename      *string                                                          `json:"interface-name"`
	Ipv4unicast        *NetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicast        `json:"ipv4-unicast,omitempty"`
	Ipv6unicast        *NetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicast        `json:"ipv6-unicast,omitempty"`
	Ldpsynchronization *NetworkinstanceProtocolsIsisInstanceInterfaceLdpsynchronization `json:"ldp-synchronization,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=2
	Level []*NetworkinstanceProtocolsIsisInstanceInterfaceLevel `json:"level,omitempty"`
	// +kubebuilder:default:=false
	Passive        *bool                                                        `json:"passive,omitempty"`
	Segmentrouting *NetworkinstanceProtocolsIsisInstanceInterfaceSegmentrouting `json:"segment-routing,omitempty"`
	Timers         *NetworkinstanceProtocolsIsisInstanceInterfaceTimers         `json:"timers,omitempty"`
	Traceoptions   *NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptions   `json:"trace-options,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceAuthentication struct
type NetworkinstanceProtocolsIsisInstanceInterfaceAuthentication struct {
	Helloauthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicast struct
type NetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Enablebfd *bool `json:"enable-bfd,omitempty"`
	// +kubebuilder:default:=false
	Includebfdtlv *bool `json:"include-bfd-tlv,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicast struct
type NetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Enablebfd *bool `json:"enable-bfd,omitempty"`
	// +kubebuilder:default:=false
	Includebfdtlv *bool `json:"include-bfd-tlv,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLdpsynchronization struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLdpsynchronization struct {
	Disable  *string `json:"disable,omitempty"`
	Endoflib *bool   `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	Holddowntimer *uint16 `json:"hold-down-timer,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLevel struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLevel struct {
	Authentication *NetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:default:=false
	Disable *bool `json:"disable,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Ipv6unicastmetric *uint32 `json:"ipv6-unicast-metric,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Levelnumber *uint8 `json:"level-number"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Metric *uint32 `json:"metric,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=64
	Priority *uint8                                                    `json:"priority,omitempty"`
	Timers   *NetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers `json:"timers,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=20000
	// +kubebuilder:default:=9
	Hellointerval *uint32 `json:"hello-interval,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=3
	Hellomultiplier *uint8 `json:"hello-multiplier,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceSegmentrouting struct
type NetworkinstanceProtocolsIsisInstanceInterfaceSegmentrouting struct {
	Mpls *NetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMpls `json:"mpls,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMpls struct
type NetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMpls struct {
	Ipv4nodesid        *NetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4nodesid        `json:"ipv4-node-sid,omitempty"`
	Staticadjacencysid *NetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsStaticadjacencysid `json:"static-adjacency-sid,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4nodesid struct
type NetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsIpv4nodesid struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Index *uint32 `json:"index"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsStaticadjacencysid struct
type NetworkinstanceProtocolsIsisInstanceInterfaceSegmentroutingMplsStaticadjacencysid struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1048575
	Labelvalue *uint32 `json:"label-value,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceTimers struct
type NetworkinstanceProtocolsIsisInstanceInterfaceTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10
	Csnpinterval *uint16 `json:"csnp-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=100
	Lsppacinginterval *uint64 `json:"lsp-pacing-interval,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptions struct
type NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptions struct {
	Trace *NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTrace `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTrace struct
type NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`
	Trace E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceIpv4unicast struct
type NetworkinstanceProtocolsIsisInstanceIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsIsisInstanceIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceIpv6unicast struct
type NetworkinstanceProtocolsIsisInstanceIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsIsisInstanceIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Multitopology *bool `json:"multi-topology,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLdpsynchronization struct
type NetworkinstanceProtocolsIsisInstanceLdpsynchronization struct {
	// +kubebuilder:default:=false
	Endoflib *bool `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	Holddowntimer *uint16 `json:"hold-down-timer,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevel struct
type NetworkinstanceProtocolsIsisInstanceLevel struct {
	Authentication *NetworkinstanceProtocolsIsisInstanceLevelAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:default:=false
	Bgplsexclude *bool `json:"bgp-ls-exclude,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Levelnumber *uint8 `json:"level-number"`
	// +kubebuilder:validation:Enum=`narrow`;`wide`
	// +kubebuilder:default:="wide"
	Metricstyle E_NetworkinstanceProtocolsIsisInstanceLevelMetricstyle `json:"metric-style,omitempty"`
	//Metricstyle *string `json:"metric-style,omitempty"`
	Routepreference *NetworkinstanceProtocolsIsisInstanceLevelRoutepreference `json:"route-preference,omitempty"`
	Traceoptions    *NetworkinstanceProtocolsIsisInstanceLevelTraceoptions    `json:"trace-options,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelAuthentication struct
type NetworkinstanceProtocolsIsisInstanceLevelAuthentication struct {
	Csnpauthentication  *bool   `json:"csnp-authentication,omitempty"`
	Helloauthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
	Psnpauthentication  *bool   `json:"psnp-authentication,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelRoutepreference struct
type NetworkinstanceProtocolsIsisInstanceLevelRoutepreference struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	External *uint8 `json:"external,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Internal *uint8 `json:"internal,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelTraceoptions struct
type NetworkinstanceProtocolsIsisInstanceLevelTraceoptions struct {
	Trace *NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTrace `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTrace struct
type NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`lsdb`;`routes`;`spf`
	Trace E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceNet struct
type NetworkinstanceProtocolsIsisInstanceNet struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[a-fA-F0-9]{2}(\.[a-fA-F0-9]{4}){3,9}\.[0]{2}`
	Net *string `json:"net,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceOverload struct
type NetworkinstanceProtocolsIsisInstanceOverload struct {
	// +kubebuilder:default:=false
	Advertiseexternal *bool `json:"advertise-external,omitempty"`
	// +kubebuilder:default:=false
	Advertiseinterlevel *bool                                                  `json:"advertise-interlevel,omitempty"`
	Immediate           *NetworkinstanceProtocolsIsisInstanceOverloadImmediate `json:"immediate,omitempty"`
	Onboot              *NetworkinstanceProtocolsIsisInstanceOverloadOnboot    `json:"on-boot,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceOverloadImmediate struct
type NetworkinstanceProtocolsIsisInstanceOverloadImmediate struct {
	// +kubebuilder:default:=false
	Maxmetric *bool `json:"max-metric,omitempty"`
	// +kubebuilder:default:=false
	Setbit *bool `json:"set-bit,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceOverloadOnboot struct
type NetworkinstanceProtocolsIsisInstanceOverloadOnboot struct {
	Maxmetric *bool `json:"max-metric,omitempty"`
	Setbit    *bool `json:"set-bit,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=1800
	Timeout *uint16 `json:"timeout,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceSegmentrouting struct
type NetworkinstanceProtocolsIsisInstanceSegmentrouting struct {
	Mpls *NetworkinstanceProtocolsIsisInstanceSegmentroutingMpls `json:"mpls,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceSegmentroutingMpls struct
type NetworkinstanceProtocolsIsisInstanceSegmentroutingMpls struct {
	// +kubebuilder:default:=15
	Adjacencysidholdtime *uint16 `json:"adjacency-sid-hold-time,omitempty"`
	Staticlabelblock     *string `json:"static-label-block,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTedatabaseinstall struct
type NetworkinstanceProtocolsIsisInstanceTedatabaseinstall struct {
	Bgpls *NetworkinstanceProtocolsIsisInstanceTedatabaseinstallBgpls `json:"bgp-ls,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTedatabaseinstallBgpls struct
type NetworkinstanceProtocolsIsisInstanceTedatabaseinstallBgpls struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Bgplsidentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Igpidentifier *uint64 `json:"igp-identifier,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimers struct
type NetworkinstanceProtocolsIsisInstanceTimers struct {
	Lspgeneration *NetworkinstanceProtocolsIsisInstanceTimersLspgeneration `json:"lsp-generation,omitempty"`
	// kubebuilder:validation:Minimum=350
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1200
	Lsplifetime *uint16                                               `json:"lsp-lifetime,omitempty"`
	Lsprefresh  *NetworkinstanceProtocolsIsisInstanceTimersLsprefresh `json:"lsp-refresh,omitempty"`
	Spf         *NetworkinstanceProtocolsIsisInstanceTimersSpf        `json:"spf,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimersLspgeneration struct
type NetworkinstanceProtocolsIsisInstanceTimersLspgeneration struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=10
	Initialwait *uint64 `json:"initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=5000
	Maxwait *uint64 `json:"max-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Secondwait *uint64 `json:"second-wait,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimersLsprefresh struct
type NetworkinstanceProtocolsIsisInstanceTimersLsprefresh struct {
	// +kubebuilder:default:=true
	Halflifetime *bool `json:"half-lifetime,omitempty"`
	// kubebuilder:validation:Minimum=150
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=600
	Interval *uint16 `json:"interval,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimersSpf struct
type NetworkinstanceProtocolsIsisInstanceTimersSpf struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Initialwait *uint64 `json:"initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=10000
	Maxwait *uint64 `json:"max-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Secondwait *uint64 `json:"second-wait,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTraceoptions struct
type NetworkinstanceProtocolsIsisInstanceTraceoptions struct {
	Trace *NetworkinstanceProtocolsIsisInstanceTraceoptionsTrace `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTraceoptionsTrace struct
type NetworkinstanceProtocolsIsisInstanceTraceoptionsTrace struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`graceful-restart`;`interfaces`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`;`routes`;`summary-addresses`
	Trace E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace `json:"trace,omitempty"`
	//Trace *string `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTrafficengineering struct
type NetworkinstanceProtocolsIsisInstanceTrafficengineering struct {
	// +kubebuilder:default:=false
	Advertisement *bool `json:"advertisement,omitempty"`
	// +kubebuilder:default:=true
	Legacylinkattributeadvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTransport struct
type NetworkinstanceProtocolsIsisInstanceTransport struct {
	// kubebuilder:validation:Minimum=490
	// kubebuilder:validation:Maximum=9490
	// +kubebuilder:default:=1492
	Lspmtusize *uint16 `json:"lsp-mtu-size,omitempty"`
}

// A NetworkinstanceProtocolsIsisSpec defines the desired state of a NetworkinstanceProtocolsIsis.
type NetworkinstanceProtocolsIsisSpec struct {
	nddv1.ResourceSpec           `json:",inline"`
	NetworkInstanceName          *string                       `json:"network-instance-name"`
	NetworkinstanceProtocolsIsis *NetworkinstanceProtocolsIsis `json:"isis,omitempty"`
}

// A NetworkinstanceProtocolsIsisStatus represents the observed state of a NetworkinstanceProtocolsIsis.
type NetworkinstanceProtocolsIsisStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsIsis is the Schema for the NetworkinstanceProtocolsIsis API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlNetworkinstanceProtocolsIsis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsIsisSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsIsisStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsIsisList contains a list of NetworkinstanceProtocolsIsiss
type SrlNetworkinstanceProtocolsIsisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsIsis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsIsis{}, &SrlNetworkinstanceProtocolsIsisList{})
}

// NetworkinstanceProtocolsIsis type metadata.
var (
	NetworkinstanceProtocolsIsisKindKind         = reflect.TypeOf(SrlNetworkinstanceProtocolsIsis{}).Name()
	NetworkinstanceProtocolsIsisGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsIsisKindKind}.String()
	NetworkinstanceProtocolsIsisKindAPIVersion   = NetworkinstanceProtocolsIsisKindKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsIsisGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsIsisKindKind)
)
