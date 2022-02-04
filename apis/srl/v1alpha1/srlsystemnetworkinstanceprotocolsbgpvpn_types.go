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

// SystemNetworkinstanceProtocolsBgpvpn struct
type SystemNetworkinstanceProtocolsBgpvpn struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	Bgpinstance []*SystemNetworkinstanceProtocolsBgpvpnBgpinstance `json:"bgp-instance,omitempty"`
}

// SystemNetworkinstanceProtocolsBgpvpnBgpinstance struct
type SystemNetworkinstanceProtocolsBgpvpnBgpinstance struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Id                 *uint8                                                             `json:"id"`
	Routedistinguisher *SystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher `json:"route-distinguisher,omitempty"`
	Routetarget        *SystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget        `json:"route-target,omitempty"`
}

// SystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher struct
type SystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutedistinguisher struct {
}

// SystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget struct
type SystemNetworkinstanceProtocolsBgpvpnBgpinstanceRoutetarget struct {
}

// A SystemNetworkinstanceProtocolsBgpvpnSpec defines the desired state of a SystemNetworkinstanceProtocolsBgpvpn.
type SystemNetworkinstanceProtocolsBgpvpnSpec struct {
	nddv1.ResourceSpec                   `json:",inline"`
	SystemNetworkinstanceProtocolsBgpvpn *SystemNetworkinstanceProtocolsBgpvpn `json:"bgp-vpn,omitempty"`
}

// A SystemNetworkinstanceProtocolsBgpvpnStatus represents the observed state of a SystemNetworkinstanceProtocolsBgpvpn.
type SystemNetworkinstanceProtocolsBgpvpnStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsBgpvpn is the Schema for the SystemNetworkinstanceProtocolsBgpvpn API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={ndd,srl2}
type SrlSystemNetworkinstanceProtocolsBgpvpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNetworkinstanceProtocolsBgpvpnSpec   `json:"spec,omitempty"`
	Status SystemNetworkinstanceProtocolsBgpvpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsBgpvpnList contains a list of SystemNetworkinstanceProtocolsBgpvpns
type SrlSystemNetworkinstanceProtocolsBgpvpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsBgpvpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsBgpvpn{}, &SrlSystemNetworkinstanceProtocolsBgpvpnList{})
}

// SystemNetworkinstanceProtocolsBgpvpn type metadata.
var (
	SystemNetworkinstanceProtocolsBgpvpnKindKind         = reflect.TypeOf(SrlSystemNetworkinstanceProtocolsBgpvpn{}).Name()
	SystemNetworkinstanceProtocolsBgpvpnGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNetworkinstanceProtocolsBgpvpnKindKind}.String()
	SystemNetworkinstanceProtocolsBgpvpnKindAPIVersion   = SystemNetworkinstanceProtocolsBgpvpnKindKind + "." + GroupVersion.String()
	SystemNetworkinstanceProtocolsBgpvpnGroupVersionKind = GroupVersion.WithKind(SystemNetworkinstanceProtocolsBgpvpnKindKind)
)
