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

// RoutingpolicyAspathset struct
type RoutingpolicyAspathset struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=65535
	Expression *string `json:"expression,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// A RoutingpolicyAspathsetSpec defines the desired state of a RoutingpolicyAspathset.
type RoutingpolicyAspathsetSpec struct {
	nddv1.ResourceSpec     `json:",inline"`
	RoutingpolicyAspathset *RoutingpolicyAspathset `json:"as-path-set,omitempty"`
}

// A RoutingpolicyAspathsetStatus represents the observed state of a RoutingpolicyAspathset.
type RoutingpolicyAspathsetStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyAspathset is the Schema for the RoutingpolicyAspathset API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={ndd,srl2}
type SrlRoutingpolicyAspathset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoutingpolicyAspathsetSpec   `json:"spec,omitempty"`
	Status RoutingpolicyAspathsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyAspathsetList contains a list of RoutingpolicyAspathsets
type SrlRoutingpolicyAspathsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlRoutingpolicyAspathset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlRoutingpolicyAspathset{}, &SrlRoutingpolicyAspathsetList{})
}

// RoutingpolicyAspathset type metadata.
var (
	RoutingpolicyAspathsetKindKind         = reflect.TypeOf(SrlRoutingpolicyAspathset{}).Name()
	RoutingpolicyAspathsetGroupKind        = schema.GroupKind{Group: Group, Kind: RoutingpolicyAspathsetKindKind}.String()
	RoutingpolicyAspathsetKindAPIVersion   = RoutingpolicyAspathsetKindKind + "." + GroupVersion.String()
	RoutingpolicyAspathsetGroupVersionKind = GroupVersion.WithKind(RoutingpolicyAspathsetKindKind)
)
