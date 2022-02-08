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

// Tunnelinterface struct
type Tunnelinterface struct {
	// kubebuilder:validation:MinLength=6
	// kubebuilder:validation:MaxLength=8
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9]))`
	Name *string `json:"name"`
}

// A TunnelinterfaceSpec defines the desired state of a Tunnelinterface.
type TunnelinterfaceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	Tunnelinterface    *Tunnelinterface `json:"tunnel-interface,omitempty"`
}

// A TunnelinterfaceStatus represents the observed state of a Tunnelinterface.
type TunnelinterfaceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterface is the Schema for the Tunnelinterface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlTunnelinterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TunnelinterfaceSpec   `json:"spec,omitempty"`
	Status TunnelinterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterfaceList contains a list of Tunnelinterfaces
type SrlTunnelinterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlTunnelinterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlTunnelinterface{}, &SrlTunnelinterfaceList{})
}

// Tunnelinterface type metadata.
var (
	TunnelinterfaceKindKind         = reflect.TypeOf(SrlTunnelinterface{}).Name()
	TunnelinterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: TunnelinterfaceKindKind}.String()
	TunnelinterfaceKindAPIVersion   = TunnelinterfaceKindKind + "." + GroupVersion.String()
	TunnelinterfaceGroupVersionKind = GroupVersion.WithKind(TunnelinterfaceKindKind)
)
