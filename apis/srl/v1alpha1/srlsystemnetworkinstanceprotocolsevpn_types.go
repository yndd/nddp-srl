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

// SystemNetworkinstanceProtocolsEvpn struct
type SystemNetworkinstanceProtocolsEvpn struct {
	Ethernetsegments *SystemNetworkinstanceProtocolsEvpnEthernetsegments `json:"ethernet-segments,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEthernetsegments struct
type SystemNetworkinstanceProtocolsEvpnEthernetsegments struct {
	Timers *SystemNetworkinstanceProtocolsEvpnEthernetsegmentsTimers `json:"timers,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEthernetsegmentsTimers struct
type SystemNetworkinstanceProtocolsEvpnEthernetsegmentsTimers struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=3
	Activationtimer *uint32 `json:"activation-timer,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=6000
	// +kubebuilder:default:=10
	Boottimer *uint32 `json:"boot-timer,omitempty"`
}

// A SystemNetworkinstanceProtocolsEvpnSpec defines the desired state of a SystemNetworkinstanceProtocolsEvpn.
type SystemNetworkinstanceProtocolsEvpnSpec struct {
	nddv1.ResourceSpec                 `json:",inline"`
	SystemNetworkinstanceProtocolsEvpn *SystemNetworkinstanceProtocolsEvpn `json:"evpn,omitempty"`
}

// A SystemNetworkinstanceProtocolsEvpnStatus represents the observed state of a SystemNetworkinstanceProtocolsEvpn.
type SystemNetworkinstanceProtocolsEvpnStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpn is the Schema for the SystemNetworkinstanceProtocolsEvpn API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlSystemNetworkinstanceProtocolsEvpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNetworkinstanceProtocolsEvpnSpec   `json:"spec,omitempty"`
	Status SystemNetworkinstanceProtocolsEvpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnList contains a list of SystemNetworkinstanceProtocolsEvpns
type SrlSystemNetworkinstanceProtocolsEvpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsEvpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsEvpn{}, &SrlSystemNetworkinstanceProtocolsEvpnList{})
}

// SystemNetworkinstanceProtocolsEvpn type metadata.
var (
	SystemNetworkinstanceProtocolsEvpnKindKind         = reflect.TypeOf(SrlSystemNetworkinstanceProtocolsEvpn{}).Name()
	SystemNetworkinstanceProtocolsEvpnGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNetworkinstanceProtocolsEvpnKindKind}.String()
	SystemNetworkinstanceProtocolsEvpnKindAPIVersion   = SystemNetworkinstanceProtocolsEvpnKindKind + "." + GroupVersion.String()
	SystemNetworkinstanceProtocolsEvpnGroupVersionKind = GroupVersion.WithKind(SystemNetworkinstanceProtocolsEvpnKindKind)
)
