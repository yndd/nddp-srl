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

// NetworkinstanceProtocolsOspf struct
type NetworkinstanceProtocolsOspf struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=3
	Instance []*NetworkinstanceProtocolsOspfInstance `json:"instance,omitempty"`
}

// NetworkinstanceProtocolsOspfInstance struct
type NetworkinstanceProtocolsOspfInstance struct {
	Addressfamily *string `json:"address-family,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_NetworkinstanceProtocolsOspfInstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Enum=`area`;`as`;`false`;`link`
	Advertiseroutercapability E_NetworkinstanceProtocolsOspfInstanceAdvertiseroutercapability `json:"advertise-router-capability,omitempty"`
	//Advertiseroutercapability *string `json:"advertise-router-capability,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Area               []*NetworkinstanceProtocolsOspfInstanceArea             `json:"area,omitempty"`
	Asbr               *NetworkinstanceProtocolsOspfInstanceAsbr               `json:"asbr,omitempty"`
	Exportlimit        *NetworkinstanceProtocolsOspfInstanceExportlimit        `json:"export-limit,omitempty"`
	Exportpolicy       *string                                                 `json:"export-policy,omitempty"`
	Externaldboverflow *NetworkinstanceProtocolsOspfInstanceExternaldboverflow `json:"external-db-overflow,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=150
	Externalpreference *uint8                                               `json:"external-preference,omitempty"`
	Gracefulrestart    *NetworkinstanceProtocolsOspfInstanceGracefulrestart `json:"graceful-restart,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Instanceid         *uint32                                                 `json:"instance-id,omitempty"`
	Ldpsynchronization *NetworkinstanceProtocolsOspfInstanceLdpsynchronization `json:"ldp-synchronization,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxecmppaths *uint8 `json:"max-ecmp-paths,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name     *string                                       `json:"name"`
	Overload *NetworkinstanceProtocolsOspfInstanceOverload `json:"overload,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=10
	Preference *uint8 `json:"preference,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8000000000
	// +kubebuilder:default:=400000000
	Referencebandwidth *uint64 `json:"reference-bandwidth,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Routerid           *string                                                 `json:"router-id,omitempty"`
	Tedatabaseinstall  *NetworkinstanceProtocolsOspfInstanceTedatabaseinstall  `json:"te-database-install,omitempty"`
	Timers             *NetworkinstanceProtocolsOspfInstanceTimers             `json:"timers,omitempty"`
	Traceoptions       *NetworkinstanceProtocolsOspfInstanceTraceoptions       `json:"trace-options,omitempty"`
	Trafficengineering *NetworkinstanceProtocolsOspfInstanceTrafficengineering `json:"traffic-engineering,omitempty"`
	Version            *string                                                 `json:"version"`
}

// NetworkinstanceProtocolsOspfInstanceArea struct
type NetworkinstanceProtocolsOspfInstanceArea struct {
	// +kubebuilder:default:=true
	Advertiseroutercapability *bool `json:"advertise-router-capability,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|[0-9\.]*|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])([\p{N}\p{L}]+)?`
	Areaid *string `json:"area-id"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Arearange []*NetworkinstanceProtocolsOspfInstanceAreaArearange `json:"area-range,omitempty"`
	// +kubebuilder:default:=false
	Bgplsexclude *bool `json:"bgp-ls-exclude,omitempty"`
	// +kubebuilder:default:=true
	Blackholeaggregate *bool   `json:"blackhole-aggregate,omitempty"`
	Exportpolicy       *string `json:"export-policy,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Interface []*NetworkinstanceProtocolsOspfInstanceAreaInterface `json:"interface,omitempty"`
	Nssa      *NetworkinstanceProtocolsOspfInstanceAreaNssa        `json:"nssa,omitempty"`
	Stub      *NetworkinstanceProtocolsOspfInstanceAreaStub        `json:"stub,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaArearange struct
type NetworkinstanceProtocolsOspfInstanceAreaArearange struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefixmask *string `json:"ip-prefix-mask"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterface struct
type NetworkinstanceProtocolsOspfInstanceAreaInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsOspfInstanceAreaInterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=true
	Advertiseroutercapability *bool `json:"advertise-router-capability,omitempty"`
	// +kubebuilder:default:=true
	Advertisesubnet *bool                                                            `json:"advertise-subnet,omitempty"`
	Authentication  *NetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=40
	Deadinterval     *uint32                                                            `json:"dead-interval,omitempty"`
	Failuredetection *NetworkinstanceProtocolsOspfInstanceAreaInterfaceFailuredetection `json:"failure-detection,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10
	Hellointerval *uint32 `json:"hello-interval,omitempty"`
	Interfacename *string `json:"interface-name"`
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	Interfacetype E_NetworkinstanceProtocolsOspfInstanceAreaInterfaceInterfacetype `json:"interface-type,omitempty"`
	//Interfacetype *string `json:"interface-type,omitempty"`
	Ldpsynchronization *NetworkinstanceProtocolsOspfInstanceAreaInterfaceLdpsynchronization `json:"ldp-synchronization,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`except-own-rtrlsa`;`except-own-rtrlsa-and-defaults`;`none`
	// +kubebuilder:default:="none"
	Lsafilterout E_NetworkinstanceProtocolsOspfInstanceAreaInterfaceLsafilterout `json:"lsa-filter-out,omitempty"`
	//Lsafilterout *string `json:"lsa-filter-out,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Metric *uint16 `json:"metric,omitempty"`
	// kubebuilder:validation:Minimum=512
	// kubebuilder:validation:Maximum=9486
	Mtu     *uint32 `json:"mtu,omitempty"`
	Passive *bool   `json:"passive,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	Priority *uint16 `json:"priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=5
	Retransmitinterval *uint32                                                        `json:"retransmit-interval,omitempty"`
	Traceoptions       *NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptions `json:"trace-options,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=1
	Transitdelay *uint32 `json:"transit-delay,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceFailuredetection struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceFailuredetection struct {
	// +kubebuilder:default:=false
	Enablebfd *bool `json:"enable-bfd,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceLdpsynchronization struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceLdpsynchronization struct {
	Disable  *string `json:"disable,omitempty"`
	Endoflib *bool   `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	Holddowntimer *uint16 `json:"hold-down-timer,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptions struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptions struct {
	Trace *NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTrace `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTrace struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTrace struct {
	Adjacencies *string                                                                   `json:"adjacencies,omitempty"`
	Interfaces  *string                                                                   `json:"interfaces,omitempty"`
	Packet      *NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacket `json:"packet,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacket struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacket struct {
	Detail *string `json:"detail,omitempty"`
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier E_NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacketModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type E_NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceoptionsTracePacketType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaNssa struct
type NetworkinstanceProtocolsOspfInstanceAreaNssa struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Arearange             []*NetworkinstanceProtocolsOspfInstanceAreaNssaArearange           `json:"area-range,omitempty"`
	Originatedefaultroute *NetworkinstanceProtocolsOspfInstanceAreaNssaOriginatedefaultroute `json:"originate-default-route,omitempty"`
	// +kubebuilder:default:=true
	Redistributeexternal *bool `json:"redistribute-external,omitempty"`
	// +kubebuilder:default:=true
	Summaries *bool `json:"summaries,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaNssaArearange struct
type NetworkinstanceProtocolsOspfInstanceAreaNssaArearange struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefixmask *string `json:"ip-prefix-mask"`
}

// NetworkinstanceProtocolsOspfInstanceAreaNssaOriginatedefaultroute struct
type NetworkinstanceProtocolsOspfInstanceAreaNssaOriginatedefaultroute struct {
	// +kubebuilder:default:=true
	Adjacencycheck *bool `json:"adjacency-check,omitempty"`
	// +kubebuilder:default:=false
	Typenssa *bool `json:"type-nssa,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaStub struct
type NetworkinstanceProtocolsOspfInstanceAreaStub struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1
	Defaultmetric *uint16 `json:"default-metric,omitempty"`
	// +kubebuilder:default:=true
	Summaries *bool `json:"summaries,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAsbr struct
type NetworkinstanceProtocolsOspfInstanceAsbr struct {
	// +kubebuilder:default:="none"
	Tracepath *string `json:"trace-path,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceExportlimit struct
type NetworkinstanceProtocolsOspfInstanceExportlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	Logpercent *uint32 `json:"log-percent,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Number *uint32 `json:"number"`
}

// NetworkinstanceProtocolsOspfInstanceExternaldboverflow struct
type NetworkinstanceProtocolsOspfInstanceExternaldboverflow struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=2147483647
	// +kubebuilder:default:=0
	Interval *uint32 `json:"interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=2147483647
	// +kubebuilder:default:=0
	Limit *uint32 `json:"limit,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceGracefulrestart struct
type NetworkinstanceProtocolsOspfInstanceGracefulrestart struct {
	// +kubebuilder:default:=false
	Helpermode *bool `json:"helper-mode,omitempty"`
	// +kubebuilder:default:=false
	Strictlsachecking *bool `json:"strict-lsa-checking,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceLdpsynchronization struct
type NetworkinstanceProtocolsOspfInstanceLdpsynchronization struct {
	// +kubebuilder:default:=false
	Endoflib *bool `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	Holddowntimer *uint16 `json:"hold-down-timer,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceOverload struct
type NetworkinstanceProtocolsOspfInstanceOverload struct {
	// +kubebuilder:default:=false
	Active *bool `json:"active,omitempty"`
	// +kubebuilder:default:=false
	Overloadincludeext1 *bool `json:"overload-include-ext-1,omitempty"`
	// +kubebuilder:default:=false
	Overloadincludeext2 *bool `json:"overload-include-ext-2,omitempty"`
	// +kubebuilder:default:=false
	Overloadincludestub *bool                                                       `json:"overload-include-stub,omitempty"`
	Overloadonboot      *NetworkinstanceProtocolsOspfInstanceOverloadOverloadonboot `json:"overload-on-boot,omitempty"`
	Rtradvlsalimit      *NetworkinstanceProtocolsOspfInstanceOverloadRtradvlsalimit `json:"rtr-adv-lsa-limit,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceOverloadOverloadonboot struct
type NetworkinstanceProtocolsOspfInstanceOverloadOverloadonboot struct {
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	Timeout *uint32 `json:"timeout,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceOverloadRtradvlsalimit struct
type NetworkinstanceProtocolsOspfInstanceOverloadRtradvlsalimit struct {
	Logonly *bool `json:"log-only,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Maxlsacount *uint32 `json:"max-lsa-count,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	Overloadtimeout *uint16 `json:"overload-timeout,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=0
	Warningthreshold *uint8 `json:"warning-threshold,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTedatabaseinstall struct
type NetworkinstanceProtocolsOspfInstanceTedatabaseinstall struct {
	Bgpls *NetworkinstanceProtocolsOspfInstanceTedatabaseinstallBgpls `json:"bgp-ls,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTedatabaseinstallBgpls struct
type NetworkinstanceProtocolsOspfInstanceTedatabaseinstallBgpls struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Bgplsidentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	Igpidentifier *uint64 `json:"igp-identifier,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTimers struct
type NetworkinstanceProtocolsOspfInstanceTimers struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	Incrementalspfwait *uint32 `json:"incremental-spf-wait,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	Lsaaccumulate *uint32 `json:"lsa-accumulate,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=1000
	Lsaarrival  *uint32                                                `json:"lsa-arrival,omitempty"`
	Lsagenerate *NetworkinstanceProtocolsOspfInstanceTimersLsagenerate `json:"lsa-generate,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	Redistributedelay *uint32                                            `json:"redistribute-delay,omitempty"`
	Spfwait           *NetworkinstanceProtocolsOspfInstanceTimersSpfwait `json:"spf-wait,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTimersLsagenerate struct
type NetworkinstanceProtocolsOspfInstanceTimersLsagenerate struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=5000
	Lsainitialwait *uint32 `json:"lsa-initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=5000
	Lsasecondwait *uint32 `json:"lsa-second-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=5000
	Maxlsawait *uint32 `json:"max-lsa-wait,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTimersSpfwait struct
type NetworkinstanceProtocolsOspfInstanceTimersSpfwait struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Spfinitialwait *uint32 `json:"spf-initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=10000
	Spfmaxwait *uint32 `json:"spf-max-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	Spfsecondwait *uint32 `json:"spf-second-wait,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceoptions struct
type NetworkinstanceProtocolsOspfInstanceTraceoptions struct {
	Trace *NetworkinstanceProtocolsOspfInstanceTraceoptionsTrace `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceoptionsTrace struct
type NetworkinstanceProtocolsOspfInstanceTraceoptionsTrace struct {
	Adjacencies     *string                                                      `json:"adjacencies,omitempty"`
	Gracefulrestart *string                                                      `json:"graceful-restart,omitempty"`
	Interfaces      *string                                                      `json:"interfaces,omitempty"`
	Lsdb            *NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceLsdb   `json:"lsdb,omitempty"`
	Misc            *string                                                      `json:"misc,omitempty"`
	Packet          *NetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacket `json:"packet,omitempty"`
	Routes          *NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceRoutes `json:"routes,omitempty"`
	Spf             *NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceSpf    `json:"spf,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceLsdb struct
type NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceLsdb struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Linkstateid *string `json:"link-state-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Routerid *string `json:"router-id,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`external`;`inter-area-prefix`;`inter-area-router`;`intra-area-prefix`;`network`;`nssa`;`opaque`;`router`;`summary`
	Type E_NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceLsdbType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacket struct
type NetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacket struct {
	Detail *string `json:"detail,omitempty"`
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier E_NetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacketModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type E_NetworkinstanceProtocolsOspfInstanceTraceoptionsTracePacketType `json:"type,omitempty"`
	//Type *string `json:"type,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceRoutes struct
type NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceRoutes struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Destaddress *string `json:"dest-address,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceSpf struct
type NetworkinstanceProtocolsOspfInstanceTraceoptionsTraceSpf struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Destaddress *string `json:"dest-address,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTrafficengineering struct
type NetworkinstanceProtocolsOspfInstanceTrafficengineering struct {
	// +kubebuilder:default:=false
	Advertisement *bool `json:"advertisement,omitempty"`
	// +kubebuilder:default:=true
	Legacylinkattributeadvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// A NetworkinstanceProtocolsOspfSpec defines the desired state of a NetworkinstanceProtocolsOspf.
type NetworkinstanceProtocolsOspfSpec struct {
	nddv1.ResourceSpec           `json:",inline"`
	NetworkInstanceName          *string                       `json:"network-instance-name"`
	NetworkinstanceProtocolsOspf *NetworkinstanceProtocolsOspf `json:"ospf,omitempty"`
}

// A NetworkinstanceProtocolsOspfStatus represents the observed state of a NetworkinstanceProtocolsOspf.
type NetworkinstanceProtocolsOspfStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsOspf is the Schema for the NetworkinstanceProtocolsOspf API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={ndd,srl2}
type SrlNetworkinstanceProtocolsOspf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsOspfSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsOspfStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsOspfList contains a list of NetworkinstanceProtocolsOspfs
type SrlNetworkinstanceProtocolsOspfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsOspf `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsOspf{}, &SrlNetworkinstanceProtocolsOspfList{})
}

// NetworkinstanceProtocolsOspf type metadata.
var (
	NetworkinstanceProtocolsOspfKindKind         = reflect.TypeOf(SrlNetworkinstanceProtocolsOspf{}).Name()
	NetworkinstanceProtocolsOspfGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsOspfKindKind}.String()
	NetworkinstanceProtocolsOspfKindAPIVersion   = NetworkinstanceProtocolsOspfKindKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsOspfGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsOspfKindKind)
)
