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

// NetworkinstanceProtocolsBgpevpn struct
type NetworkinstanceProtocolsBgpevpn struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Bgpinstance []*NetworkinstanceProtocolsBgpevpnBgpinstance `json:"bgp-instance,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpinstance struct
type NetworkinstanceProtocolsBgpevpnBgpinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceProtocolsBgpevpnBgpinstanceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	Defaultadmintag *uint32 `json:"default-admin-tag,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8
	// +kubebuilder:default:=1
	Ecmp *uint8 `json:"ecmp,omitempty"`
	// +kubebuilder:validation:Enum=`vxlan`
	// +kubebuilder:default:="vxlan"
	Encapsulationtype E_NetworkinstanceProtocolsBgpevpnBgpinstanceEncapsulationtype `json:"encapsulation-type,omitempty"`
	//Encapsulationtype *string `json:"encapsulation-type,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Evi            *uint32                                           `json:"evi"`
	Id             *string                                           `json:"id"`
	Routes         *NetworkinstanceProtocolsBgpevpnBgpinstanceRoutes `json:"routes,omitempty"`
	Vxlaninterface *string                                           `json:"vxlan-interface,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpinstanceRoutes struct
type NetworkinstanceProtocolsBgpevpnBgpinstanceRoutes struct {
	Bridgetable *NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetable `json:"bridge-table,omitempty"`
	Routetable  *NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetable  `json:"route-table,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetable struct
type NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetable struct {
	Inclusivemcast *NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableInclusivemcast `json:"inclusive-mcast,omitempty"`
	Macip          *NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableMacip          `json:"mac-ip,omitempty"`
	// +kubebuilder:default:="use-system-ipv4-address"
	Nexthop *string `json:"next-hop,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableInclusivemcast struct
type NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableInclusivemcast struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Originatingip *string `json:"originating-ip,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableMacip struct
type NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesBridgetableMacip struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetable struct
type NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetable struct {
	Macip *NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetableMacip `json:"mac-ip,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetableMacip struct
type NetworkinstanceProtocolsBgpevpnBgpinstanceRoutesRoutetableMacip struct {
	// +kubebuilder:default:=false
	Advertisegatewaymac *bool `json:"advertise-gateway-mac,omitempty"`
}

// A NetworkinstanceProtocolsBgpevpnSpec defines the desired state of a NetworkinstanceProtocolsBgpevpn.
type NetworkinstanceProtocolsBgpevpnSpec struct {
	nddv1.ResourceSpec              `json:",inline"`
	NetworkInstanceName             *string                          `json:"network-instance-name"`
	NetworkinstanceProtocolsBgpevpn *NetworkinstanceProtocolsBgpevpn `json:"bgp-evpn,omitempty"`
}

// A NetworkinstanceProtocolsBgpevpnStatus represents the observed state of a NetworkinstanceProtocolsBgpevpn.
type NetworkinstanceProtocolsBgpevpnStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgpevpn is the Schema for the NetworkinstanceProtocolsBgpevpn API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlNetworkinstanceProtocolsBgpevpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsBgpevpnSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsBgpevpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgpevpnList contains a list of NetworkinstanceProtocolsBgpevpns
type SrlNetworkinstanceProtocolsBgpevpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsBgpevpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsBgpevpn{}, &SrlNetworkinstanceProtocolsBgpevpnList{})
}

// NetworkinstanceProtocolsBgpevpn type metadata.
var (
	NetworkinstanceProtocolsBgpevpnKindKind         = reflect.TypeOf(SrlNetworkinstanceProtocolsBgpevpn{}).Name()
	NetworkinstanceProtocolsBgpevpnGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsBgpevpnKindKind}.String()
	NetworkinstanceProtocolsBgpevpnKindAPIVersion   = NetworkinstanceProtocolsBgpevpnKindKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsBgpevpnGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsBgpevpnKindKind)
)
