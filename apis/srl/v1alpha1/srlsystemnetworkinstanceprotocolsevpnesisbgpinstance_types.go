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

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstance struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstance struct {
	Id *string `json:"id"`
}

// A SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec defines the desired state of a SystemNetworkinstanceProtocolsEvpnEsisBgpinstance.
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec struct {
	nddv1.ResourceSpec                                `json:",inline"`
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstance *SystemNetworkinstanceProtocolsEvpnEsisBgpinstance `json:"bgp-instance,omitempty"`
}

// A SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus represents the observed state of a SystemNetworkinstanceProtocolsEvpnEsisBgpinstance.
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance is the Schema for the SystemNetworkinstanceProtocolsEvpnEsisBgpinstance API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec   `json:"spec,omitempty"`
	Status SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList contains a list of SystemNetworkinstanceProtocolsEvpnEsisBgpinstances
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{}, &SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList{})
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstance type metadata.
var (
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKindKind         = reflect.TypeOf(SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{}).Name()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKindKind}.String()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKindAPIVersion   = SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKindKind + "." + GroupVersion.String()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceGroupVersionKind = GroupVersion.WithKind(SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKindKind)
)
