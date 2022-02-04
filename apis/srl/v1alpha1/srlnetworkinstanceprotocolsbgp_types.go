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

// NetworkinstanceProtocolsBgp struct
type NetworkinstanceProtocolsBgp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsBgpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Aspathoptions  *NetworkinstanceProtocolsBgpAspathoptions  `json:"as-path-options,omitempty"`
	Authentication *NetworkinstanceProtocolsBgpAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Autonomoussystem  *uint32                                       `json:"autonomous-system"`
	Convergence       *NetworkinstanceProtocolsBgpConvergence       `json:"convergence,omitempty"`
	Dynamicneighbors  *NetworkinstanceProtocolsBgpDynamicneighbors  `json:"dynamic-neighbors,omitempty"`
	Ebgpdefaultpolicy *NetworkinstanceProtocolsBgpEbgpdefaultpolicy `json:"ebgp-default-policy,omitempty"`
	Evpn              *NetworkinstanceProtocolsBgpEvpn              `json:"evpn,omitempty"`
	Exportpolicy      *string                                       `json:"export-policy,omitempty"`
	Failuredetection  *NetworkinstanceProtocolsBgpFailuredetection  `json:"failure-detection,omitempty"`
	Gracefulrestart   *NetworkinstanceProtocolsBgpGracefulrestart   `json:"graceful-restart,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group        []*NetworkinstanceProtocolsBgpGroup     `json:"group,omitempty"`
	Importpolicy *string                                 `json:"import-policy,omitempty"`
	Ipv4unicast  *NetworkinstanceProtocolsBgpIpv4unicast `json:"ipv4-unicast,omitempty"`
	Ipv6unicast  *NetworkinstanceProtocolsBgpIpv6unicast `json:"ipv6-unicast,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=100
	Localpreference *uint32 `json:"local-preference,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Neighbor           []*NetworkinstanceProtocolsBgpNeighbor         `json:"neighbor,omitempty"`
	Preference         *NetworkinstanceProtocolsBgpPreference         `json:"preference,omitempty"`
	Routeadvertisement *NetworkinstanceProtocolsBgpRouteadvertisement `json:"route-advertisement,omitempty"`
	Routereflector     *NetworkinstanceProtocolsBgpRoutereflector     `json:"route-reflector,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Routerid      *string                                   `json:"router-id"`
	Sendcommunity *NetworkinstanceProtocolsBgpSendcommunity `json:"send-community,omitempty"`
	Traceoptions  *NetworkinstanceProtocolsBgpTraceoptions  `json:"trace-options,omitempty"`
	Transport     *NetworkinstanceProtocolsBgpTransport     `json:"transport,omitempty"`
}

// NetworkinstanceProtocolsBgpAspathoptions struct
type NetworkinstanceProtocolsBgpAspathoptions struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	Allowownas      *uint8                                                   `json:"allow-own-as,omitempty"`
	Removeprivateas *NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateas `json:"remove-private-as,omitempty"`
}

// NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateas struct
type NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateas struct {
	// +kubebuilder:default:=false
	Ignorepeeras *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	Leadingonly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	// +kubebuilder:default:="disabled"
	Mode E_NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateasMode `json:"mode,omitempty"`
	//Mode *string `json:"mode,omitempty"`
}

// NetworkinstanceProtocolsBgpAuthentication struct
type NetworkinstanceProtocolsBgpAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsBgpConvergence struct
type NetworkinstanceProtocolsBgpConvergence struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	Minwaittoadvertise *uint16 `json:"min-wait-to-advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpDynamicneighbors struct
type NetworkinstanceProtocolsBgpDynamicneighbors struct {
	Accept *NetworkinstanceProtocolsBgpDynamicneighborsAccept `json:"accept,omitempty"`
}

// NetworkinstanceProtocolsBgpDynamicneighborsAccept struct
type NetworkinstanceProtocolsBgpDynamicneighborsAccept struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Match []*NetworkinstanceProtocolsBgpDynamicneighborsAcceptMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=0
	Maxsessions *uint16 `json:"max-sessions,omitempty"`
}

// NetworkinstanceProtocolsBgpDynamicneighborsAcceptMatch struct
type NetworkinstanceProtocolsBgpDynamicneighborsAcceptMatch struct {
	Allowedpeeras *NetworkinstanceProtocolsBgpDynamicneighborsAcceptMatchAllowedpeeras `json:"allowed-peer-as,omitempty"`
	Peergroup     *string                                                              `json:"peer-group"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
}

// NetworkinstanceProtocolsBgpDynamicneighborsAcceptMatchAllowedpeeras struct
type NetworkinstanceProtocolsBgpDynamicneighborsAcceptMatchAllowedpeeras struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([1-9][0-9]*)|([1-9][0-9]*)\.\.([1-9][0-9]*)`
	Allowedpeeras *string `json:"allowed-peer-as,omitempty"`
}

// NetworkinstanceProtocolsBgpEbgpdefaultpolicy struct
type NetworkinstanceProtocolsBgpEbgpdefaultpolicy struct {
	// +kubebuilder:default:=true
	Exportrejectall *bool `json:"export-reject-all,omitempty"`
	// +kubebuilder:default:=true
	Importrejectall *bool `json:"import-reject-all,omitempty"`
}

// NetworkinstanceProtocolsBgpEvpn struct
type NetworkinstanceProtocolsBgpEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_NetworkinstanceProtocolsBgpEvpnAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Advertiseipv6nexthops *bool `json:"advertise-ipv6-next-hops,omitempty"`
	Keepallroutes         *bool `json:"keep-all-routes,omitempty"`
	// +kubebuilder:default:=false
	Rapidupdate *bool `json:"rapid-update,omitempty"`
}

// NetworkinstanceProtocolsBgpFailuredetection struct
type NetworkinstanceProtocolsBgpFailuredetection struct {
	// +kubebuilder:default:=false
	Enablebfd *bool `json:"enable-bfd,omitempty"`
	// +kubebuilder:default:=true
	Fastfailover *bool `json:"fast-failover,omitempty"`
}

// NetworkinstanceProtocolsBgpGracefulrestart struct
type NetworkinstanceProtocolsBgpGracefulrestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_NetworkinstanceProtocolsBgpGracefulrestartAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=360
	Staleroutestime *uint16 `json:"stale-routes-time,omitempty"`
}

// NetworkinstanceProtocolsBgpGroup struct
type NetworkinstanceProtocolsBgpGroup struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsBgpGroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Aspathoptions  *NetworkinstanceProtocolsBgpGroupAspathoptions  `json:"as-path-options,omitempty"`
	Authentication *NetworkinstanceProtocolsBgpGroupAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description      *string                                           `json:"description,omitempty"`
	Evpn             *NetworkinstanceProtocolsBgpGroupEvpn             `json:"evpn,omitempty"`
	Exportpolicy     *string                                           `json:"export-policy,omitempty"`
	Failuredetection *NetworkinstanceProtocolsBgpGroupFailuredetection `json:"failure-detection,omitempty"`
	Gracefulrestart  *NetworkinstanceProtocolsBgpGroupGracefulrestart  `json:"graceful-restart,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Groupname    *string                                      `json:"group-name"`
	Importpolicy *string                                      `json:"import-policy,omitempty"`
	Ipv4unicast  *NetworkinstanceProtocolsBgpGroupIpv4unicast `json:"ipv4-unicast,omitempty"`
	Ipv6unicast  *NetworkinstanceProtocolsBgpGroupIpv6unicast `json:"ipv6-unicast,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Localas []*NetworkinstanceProtocolsBgpGroupLocalas `json:"local-as,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Localpreference *uint32 `json:"local-preference,omitempty"`
	// +kubebuilder:default:=false
	Nexthopself *bool `json:"next-hop-self,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Peeras           *uint32                                           `json:"peer-as,omitempty"`
	Routereflector   *NetworkinstanceProtocolsBgpGroupRoutereflector   `json:"route-reflector,omitempty"`
	Sendcommunity    *NetworkinstanceProtocolsBgpGroupSendcommunity    `json:"send-community,omitempty"`
	Senddefaultroute *NetworkinstanceProtocolsBgpGroupSenddefaultroute `json:"send-default-route,omitempty"`
	Timers           *NetworkinstanceProtocolsBgpGroupTimers           `json:"timers,omitempty"`
	Traceoptions     *NetworkinstanceProtocolsBgpGroupTraceoptions     `json:"trace-options,omitempty"`
	Transport        *NetworkinstanceProtocolsBgpGroupTransport        `json:"transport,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupAspathoptions struct
type NetworkinstanceProtocolsBgpGroupAspathoptions struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Allowownas      *uint8                                                        `json:"allow-own-as,omitempty"`
	Removeprivateas *NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateas `json:"remove-private-as,omitempty"`
	Replacepeeras   *bool                                                         `json:"replace-peer-as,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateas struct
type NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateas struct {
	// +kubebuilder:default:=false
	Ignorepeeras *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	Leadingonly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode E_NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateasMode `json:"mode,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupAuthentication struct
type NetworkinstanceProtocolsBgpGroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupEvpn struct
type NetworkinstanceProtocolsBgpGroupEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_NetworkinstanceProtocolsBgpGroupEvpnAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Advertiseipv6nexthops *bool                                            `json:"advertise-ipv6-next-hops,omitempty"`
	Prefixlimit           *NetworkinstanceProtocolsBgpGroupEvpnPrefixlimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupEvpnPrefixlimit struct
type NetworkinstanceProtocolsBgpGroupEvpnPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupFailuredetection struct
type NetworkinstanceProtocolsBgpGroupFailuredetection struct {
	Enablebfd    *bool `json:"enable-bfd,omitempty"`
	Fastfailover *bool `json:"fast-failover,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupGracefulrestart struct
type NetworkinstanceProtocolsBgpGroupGracefulrestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_NetworkinstanceProtocolsBgpGroupGracefulrestartAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	Staleroutestime *uint16 `json:"stale-routes-time,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv4unicast struct
type NetworkinstanceProtocolsBgpGroupIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_NetworkinstanceProtocolsBgpGroupIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Advertiseipv6nexthops *bool                                                   `json:"advertise-ipv6-next-hops,omitempty"`
	Prefixlimit           *NetworkinstanceProtocolsBgpGroupIpv4unicastPrefixlimit `json:"prefix-limit,omitempty"`
	Receiveipv6nexthops   *bool                                                   `json:"receive-ipv6-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv4unicastPrefixlimit struct
type NetworkinstanceProtocolsBgpGroupIpv4unicastPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv6unicast struct
type NetworkinstanceProtocolsBgpGroupIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_NetworkinstanceProtocolsBgpGroupIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Prefixlimit *NetworkinstanceProtocolsBgpGroupIpv6unicastPrefixlimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv6unicastPrefixlimit struct
type NetworkinstanceProtocolsBgpGroupIpv6unicastPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupLocalas struct
type NetworkinstanceProtocolsBgpGroupLocalas struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Asnumber *uint32 `json:"as-number"`
	// +kubebuilder:default:=true
	Prependglobalas *bool `json:"prepend-global-as,omitempty"`
	// +kubebuilder:default:=true
	Prependlocalas *bool `json:"prepend-local-as,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupRoutereflector struct
type NetworkinstanceProtocolsBgpGroupRoutereflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Clusterid *string `json:"cluster-id,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupSendcommunity struct
type NetworkinstanceProtocolsBgpGroupSendcommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupSenddefaultroute struct
type NetworkinstanceProtocolsBgpGroupSenddefaultroute struct {
	Exportpolicy *string `json:"export-policy,omitempty"`
	// +kubebuilder:default:=false
	Ipv4unicast *bool `json:"ipv4-unicast,omitempty"`
	// +kubebuilder:default:=false
	Ipv6unicast *bool `json:"ipv6-unicast,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupTimers struct
type NetworkinstanceProtocolsBgpGroupTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=120
	Connectretry *uint16 `json:"connect-retry,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	// +kubebuilder:default:=90
	Holdtime *uint16 `json:"hold-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=21845
	Keepaliveinterval *uint16 `json:"keepalive-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=5
	Minimumadvertisementinterval *uint16 `json:"minimum-advertisement-interval,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupTraceoptions struct
type NetworkinstanceProtocolsBgpGroupTraceoptions struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Flag []*NetworkinstanceProtocolsBgpGroupTraceoptionsFlag `json:"flag,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupTraceoptionsFlag struct
type NetworkinstanceProtocolsBgpGroupTraceoptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName `json:"name,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupTransport struct
type NetworkinstanceProtocolsBgpGroupTransport struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Localaddress *string `json:"local-address,omitempty"`
	// +kubebuilder:default:=false
	Passivemode *bool `json:"passive-mode,omitempty"`
	// kubebuilder:validation:Minimum=536
	// kubebuilder:validation:Maximum=9446
	Tcpmss *uint16 `json:"tcp-mss,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4unicast struct
type NetworkinstanceProtocolsBgpIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsBgpIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	Advertiseipv6nexthops *bool                                                    `json:"advertise-ipv6-next-hops,omitempty"`
	Convergence           *NetworkinstanceProtocolsBgpIpv4unicastConvergence       `json:"convergence,omitempty"`
	Multipath             *NetworkinstanceProtocolsBgpIpv4unicastMultipath         `json:"multipath,omitempty"`
	Nexthopresolution     *NetworkinstanceProtocolsBgpIpv4unicastNexthopresolution `json:"next-hop-resolution,omitempty"`
	// +kubebuilder:default:=false
	Receiveipv6nexthops *bool `json:"receive-ipv6-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4unicastConvergence struct
type NetworkinstanceProtocolsBgpIpv4unicastConvergence struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	Maxwaittoadvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4unicastMultipath struct
type NetworkinstanceProtocolsBgpIpv4unicastMultipath struct {
	// +kubebuilder:default:=true
	Allowmultipleas *bool `json:"allow-multiple-as,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpathslevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpathslevel2 *uint32 `json:"max-paths-level-2,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4unicastNexthopresolution struct
type NetworkinstanceProtocolsBgpIpv4unicastNexthopresolution struct {
	Ipv4nexthops *NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthops `json:"ipv4-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthops struct
type NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthops struct {
	Tunnelresolution *NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolution `json:"tunnel-resolution,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolution struct
type NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolution struct {
	Allowedtunneltypes *NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes `json:"allowed-tunnel-types,omitempty"`
	// +kubebuilder:validation:Enum=`disabled`;`prefer`;`require`
	// +kubebuilder:default:="disabled"
	Mode E_NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode `json:"mode,omitempty"`
	//Mode *string `json:"mode,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes struct
type NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes struct {
	Allowedtunneltypes *string `json:"allowed-tunnel-types,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6unicast struct
type NetworkinstanceProtocolsBgpIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_NetworkinstanceProtocolsBgpIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Convergence       *NetworkinstanceProtocolsBgpIpv6unicastConvergence       `json:"convergence,omitempty"`
	Multipath         *NetworkinstanceProtocolsBgpIpv6unicastMultipath         `json:"multipath,omitempty"`
	Nexthopresolution *NetworkinstanceProtocolsBgpIpv6unicastNexthopresolution `json:"next-hop-resolution,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6unicastConvergence struct
type NetworkinstanceProtocolsBgpIpv6unicastConvergence struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	Maxwaittoadvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6unicastMultipath struct
type NetworkinstanceProtocolsBgpIpv6unicastMultipath struct {
	// +kubebuilder:default:=true
	Allowmultipleas *bool `json:"allow-multiple-as,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpathslevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Maxpathslevel2 *uint32 `json:"max-paths-level-2,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6unicastNexthopresolution struct
type NetworkinstanceProtocolsBgpIpv6unicastNexthopresolution struct {
	Ipv4nexthops *NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthops `json:"ipv4-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthops struct
type NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthops struct {
	Tunnelresolution *NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolution `json:"tunnel-resolution,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolution struct
type NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolution struct {
	Allowedtunneltypes *NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes `json:"allowed-tunnel-types,omitempty"`
	// +kubebuilder:validation:Enum=`disabled`;`prefer`;`require`
	// +kubebuilder:default:="disabled"
	Mode E_NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode `json:"mode,omitempty"`
	//Mode *string `json:"mode,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes struct
type NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionAllowedtunneltypes struct {
	Allowedtunneltypes *string `json:"allowed-tunnel-types,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighbor struct
type NetworkinstanceProtocolsBgpNeighbor struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsBgpNeighborAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Aspathoptions  *NetworkinstanceProtocolsBgpNeighborAspathoptions  `json:"as-path-options,omitempty"`
	Authentication *NetworkinstanceProtocolsBgpNeighborAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description      *string                                              `json:"description,omitempty"`
	Evpn             *NetworkinstanceProtocolsBgpNeighborEvpn             `json:"evpn,omitempty"`
	Exportpolicy     *string                                              `json:"export-policy,omitempty"`
	Failuredetection *NetworkinstanceProtocolsBgpNeighborFailuredetection `json:"failure-detection,omitempty"`
	Gracefulrestart  *NetworkinstanceProtocolsBgpNeighborGracefulrestart  `json:"graceful-restart,omitempty"`
	Importpolicy     *string                                              `json:"import-policy,omitempty"`
	Ipv4unicast      *NetworkinstanceProtocolsBgpNeighborIpv4unicast      `json:"ipv4-unicast,omitempty"`
	Ipv6unicast      *NetworkinstanceProtocolsBgpNeighborIpv6unicast      `json:"ipv6-unicast,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Localas []*NetworkinstanceProtocolsBgpNeighborLocalas `json:"local-as,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Localpreference *uint32 `json:"local-preference,omitempty"`
	Nexthopself     *bool   `json:"next-hop-self,omitempty"`
	// +kubebuilder:validation:Optional
	Peeraddress *string `json:"peer-address"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Peeras           *uint32                                              `json:"peer-as,omitempty"`
	Peergroup        *string                                              `json:"peer-group"`
	Routereflector   *NetworkinstanceProtocolsBgpNeighborRoutereflector   `json:"route-reflector,omitempty"`
	Sendcommunity    *NetworkinstanceProtocolsBgpNeighborSendcommunity    `json:"send-community,omitempty"`
	Senddefaultroute *NetworkinstanceProtocolsBgpNeighborSenddefaultroute `json:"send-default-route,omitempty"`
	Timers           *NetworkinstanceProtocolsBgpNeighborTimers           `json:"timers,omitempty"`
	Traceoptions     *NetworkinstanceProtocolsBgpNeighborTraceoptions     `json:"trace-options,omitempty"`
	Transport        *NetworkinstanceProtocolsBgpNeighborTransport        `json:"transport,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborAspathoptions struct
type NetworkinstanceProtocolsBgpNeighborAspathoptions struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Allowownas      *uint8                                                           `json:"allow-own-as,omitempty"`
	Removeprivateas *NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateas `json:"remove-private-as,omitempty"`
	Replacepeeras   *bool                                                            `json:"replace-peer-as,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateas struct
type NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateas struct {
	// +kubebuilder:default:=false
	Ignorepeeras *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	Leadingonly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode E_NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateasMode `json:"mode,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborAuthentication struct
type NetworkinstanceProtocolsBgpNeighborAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborEvpn struct
type NetworkinstanceProtocolsBgpNeighborEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_NetworkinstanceProtocolsBgpNeighborEvpnAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Advertiseipv6nexthops *bool                                               `json:"advertise-ipv6-next-hops,omitempty"`
	Prefixlimit           *NetworkinstanceProtocolsBgpNeighborEvpnPrefixlimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborEvpnPrefixlimit struct
type NetworkinstanceProtocolsBgpNeighborEvpnPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborFailuredetection struct
type NetworkinstanceProtocolsBgpNeighborFailuredetection struct {
	Enablebfd    *bool `json:"enable-bfd,omitempty"`
	Fastfailover *bool `json:"fast-failover,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborGracefulrestart struct
type NetworkinstanceProtocolsBgpNeighborGracefulrestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_NetworkinstanceProtocolsBgpNeighborGracefulrestartAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	Staleroutestime *uint16 `json:"stale-routes-time,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv4unicast struct
type NetworkinstanceProtocolsBgpNeighborIpv4unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_NetworkinstanceProtocolsBgpNeighborIpv4unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Advertiseipv6nexthops *bool                                                      `json:"advertise-ipv6-next-hops,omitempty"`
	Prefixlimit           *NetworkinstanceProtocolsBgpNeighborIpv4unicastPrefixlimit `json:"prefix-limit,omitempty"`
	Receiveipv6nexthops   *bool                                                      `json:"receive-ipv6-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv4unicastPrefixlimit struct
type NetworkinstanceProtocolsBgpNeighborIpv4unicastPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv6unicast struct
type NetworkinstanceProtocolsBgpNeighborIpv6unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_NetworkinstanceProtocolsBgpNeighborIpv6unicastAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Prefixlimit *NetworkinstanceProtocolsBgpNeighborIpv6unicastPrefixlimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv6unicastPrefixlimit struct
type NetworkinstanceProtocolsBgpNeighborIpv6unicastPrefixlimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Maxreceivedroutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	Warningthresholdpct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborLocalas struct
type NetworkinstanceProtocolsBgpNeighborLocalas struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Asnumber        *uint32 `json:"as-number"`
	Prependglobalas *bool   `json:"prepend-global-as,omitempty"`
	Prependlocalas  *bool   `json:"prepend-local-as,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborRoutereflector struct
type NetworkinstanceProtocolsBgpNeighborRoutereflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Clusterid *string `json:"cluster-id,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborSendcommunity struct
type NetworkinstanceProtocolsBgpNeighborSendcommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborSenddefaultroute struct
type NetworkinstanceProtocolsBgpNeighborSenddefaultroute struct {
	Exportpolicy *string `json:"export-policy,omitempty"`
	Ipv4unicast  *bool   `json:"ipv4-unicast,omitempty"`
	Ipv6unicast  *bool   `json:"ipv6-unicast,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTimers struct
type NetworkinstanceProtocolsBgpNeighborTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Connectretry *uint16 `json:"connect-retry,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	Holdtime *uint16 `json:"hold-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=21845
	Keepaliveinterval *uint16 `json:"keepalive-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Minimumadvertisementinterval *uint16 `json:"minimum-advertisement-interval,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTraceoptions struct
type NetworkinstanceProtocolsBgpNeighborTraceoptions struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Flag []*NetworkinstanceProtocolsBgpNeighborTraceoptionsFlag `json:"flag,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTraceoptionsFlag struct
type NetworkinstanceProtocolsBgpNeighborTraceoptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName `json:"name,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTransport struct
type NetworkinstanceProtocolsBgpNeighborTransport struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Localaddress *string `json:"local-address,omitempty"`
	Passivemode  *bool   `json:"passive-mode,omitempty"`
	// kubebuilder:validation:Minimum=536
	// kubebuilder:validation:Maximum=9446
	Tcpmss *uint16 `json:"tcp-mss,omitempty"`
}

// NetworkinstanceProtocolsBgpPreference struct
type NetworkinstanceProtocolsBgpPreference struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=170
	Ebgp *uint8 `json:"ebgp,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=170
	Ibgp *uint8 `json:"ibgp,omitempty"`
}

// NetworkinstanceProtocolsBgpRouteadvertisement struct
type NetworkinstanceProtocolsBgpRouteadvertisement struct {
	// +kubebuilder:default:=false
	Rapidwithdrawal *bool `json:"rapid-withdrawal,omitempty"`
	// +kubebuilder:default:=true
	Waitforfibinstall *bool `json:"wait-for-fib-install,omitempty"`
}

// NetworkinstanceProtocolsBgpRoutereflector struct
type NetworkinstanceProtocolsBgpRoutereflector struct {
	// +kubebuilder:default:=false
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Clusterid *string `json:"cluster-id,omitempty"`
}

// NetworkinstanceProtocolsBgpSendcommunity struct
type NetworkinstanceProtocolsBgpSendcommunity struct {
	// +kubebuilder:default:=true
	Large *bool `json:"large,omitempty"`
	// +kubebuilder:default:=true
	Standard *bool `json:"standard,omitempty"`
}

// NetworkinstanceProtocolsBgpTraceoptions struct
type NetworkinstanceProtocolsBgpTraceoptions struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Flag []*NetworkinstanceProtocolsBgpTraceoptionsFlag `json:"flag,omitempty"`
}

// NetworkinstanceProtocolsBgpTraceoptionsFlag struct
type NetworkinstanceProtocolsBgpTraceoptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier E_NetworkinstanceProtocolsBgpTraceoptionsFlagModifier `json:"modifier,omitempty"`
	//Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name E_NetworkinstanceProtocolsBgpTraceoptionsFlagName `json:"name,omitempty"`
}

// NetworkinstanceProtocolsBgpTransport struct
type NetworkinstanceProtocolsBgpTransport struct {
	// kubebuilder:validation:Minimum=536
	// kubebuilder:validation:Maximum=9446
	// +kubebuilder:default:=1024
	Tcpmss *uint16 `json:"tcp-mss,omitempty"`
}

// A NetworkinstanceProtocolsBgpSpec defines the desired state of a NetworkinstanceProtocolsBgp.
type NetworkinstanceProtocolsBgpSpec struct {
	nddv1.ResourceSpec          `json:",inline"`
	NetworkInstanceName         *string                      `json:"network-instance-name"`
	NetworkinstanceProtocolsBgp *NetworkinstanceProtocolsBgp `json:"bgp,omitempty"`
}

// A NetworkinstanceProtocolsBgpStatus represents the observed state of a NetworkinstanceProtocolsBgp.
type NetworkinstanceProtocolsBgpStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgp is the Schema for the NetworkinstanceProtocolsBgp API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={ndd,srl2}
type SrlNetworkinstanceProtocolsBgp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsBgpSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsBgpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgpList contains a list of NetworkinstanceProtocolsBgps
type SrlNetworkinstanceProtocolsBgpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsBgp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsBgp{}, &SrlNetworkinstanceProtocolsBgpList{})
}

// NetworkinstanceProtocolsBgp type metadata.
var (
	NetworkinstanceProtocolsBgpKindKind         = reflect.TypeOf(SrlNetworkinstanceProtocolsBgp{}).Name()
	NetworkinstanceProtocolsBgpGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsBgpKindKind}.String()
	NetworkinstanceProtocolsBgpKindAPIVersion   = NetworkinstanceProtocolsBgpKindKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsBgpGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsBgpKindKind)
)
