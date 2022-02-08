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

// TunnelinterfaceVxlaninterface struct
type TunnelinterfaceVxlaninterface struct {
	Bridgetable *TunnelinterfaceVxlaninterfaceBridgetable `json:"bridge-table,omitempty"`
	Egress      *TunnelinterfaceVxlaninterfaceEgress      `json:"egress,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=99999999
	Index   *uint32                               `json:"index"`
	Ingress *TunnelinterfaceVxlaninterfaceIngress `json:"ingress,omitempty"`
	Type    *string                               `json:"type"`
}

// TunnelinterfaceVxlaninterfaceBridgetable struct
type TunnelinterfaceVxlaninterfaceBridgetable struct {
}

// TunnelinterfaceVxlaninterfaceEgress struct
type TunnelinterfaceVxlaninterfaceEgress struct {
	Destinationgroups   *TunnelinterfaceVxlaninterfaceEgressDestinationgroups   `json:"destination-groups,omitempty"`
	Innerethernetheader *TunnelinterfaceVxlaninterfaceEgressInnerethernetheader `json:"inner-ethernet-header,omitempty"`
	// +kubebuilder:default:="use-system-ipv4-address"
	Sourceip *string `json:"source-ip,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationgroups struct
type TunnelinterfaceVxlaninterfaceEgressDestinationgroups struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group []*TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroup `json:"group,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroup struct
type TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroup struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=128
	Destination []*TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestination `json:"destination,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi *string `json:"esi,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestination struct
type TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestination struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestinationAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Index               *uint16                                                                                  `json:"index"`
	Innerethernetheader *TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestinationInnerethernetheader `json:"inner-ethernet-header,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestinationInnerethernetheader struct
type TunnelinterfaceVxlaninterfaceEgressDestinationgroupsGroupDestinationInnerethernetheader struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Destinationmac *string `json:"destination-mac,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressInnerethernetheader struct
type TunnelinterfaceVxlaninterfaceEgressInnerethernetheader struct {
	// +kubebuilder:default:="use-system-mac"
	Sourcemac *string `json:"source-mac,omitempty"`
}

// TunnelinterfaceVxlaninterfaceIngress struct
type TunnelinterfaceVxlaninterfaceIngress struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni"`
}

// A TunnelinterfaceVxlaninterfaceSpec defines the desired state of a TunnelinterfaceVxlaninterface.
type TunnelinterfaceVxlaninterfaceSpec struct {
	nddv1.ResourceSpec            `json:",inline"`
	TunnelInterfaceName           *string                        `json:"tunnel-interface-name"`
	TunnelinterfaceVxlaninterface *TunnelinterfaceVxlaninterface `json:"vxlan-interface,omitempty"`
}

// A TunnelinterfaceVxlaninterfaceStatus represents the observed state of a TunnelinterfaceVxlaninterface.
type TunnelinterfaceVxlaninterfaceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterfaceVxlaninterface is the Schema for the TunnelinterfaceVxlaninterface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlTunnelinterfaceVxlaninterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TunnelinterfaceVxlaninterfaceSpec   `json:"spec,omitempty"`
	Status TunnelinterfaceVxlaninterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterfaceVxlaninterfaceList contains a list of TunnelinterfaceVxlaninterfaces
type SrlTunnelinterfaceVxlaninterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlTunnelinterfaceVxlaninterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlTunnelinterfaceVxlaninterface{}, &SrlTunnelinterfaceVxlaninterfaceList{})
}

// TunnelinterfaceVxlaninterface type metadata.
var (
	TunnelinterfaceVxlaninterfaceKindKind         = reflect.TypeOf(SrlTunnelinterfaceVxlaninterface{}).Name()
	TunnelinterfaceVxlaninterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: TunnelinterfaceVxlaninterfaceKindKind}.String()
	TunnelinterfaceVxlaninterfaceKindAPIVersion   = TunnelinterfaceVxlaninterfaceKindKind + "." + GroupVersion.String()
	TunnelinterfaceVxlaninterfaceGroupVersionKind = GroupVersion.WithKind(TunnelinterfaceVxlaninterfaceKindKind)
)
