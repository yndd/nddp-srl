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

// SystemName struct
type SystemName struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Domainname *string `json:"domain-name,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])`
	Hostname *string `json:"host-name,omitempty"`
}

// A SystemNameSpec defines the desired state of a SystemName.
type SystemNameSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	SystemName         *SystemName `json:"name,omitempty"`
}

// A SystemNameStatus represents the observed state of a SystemName.
type SystemNameStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlSystemName is the Schema for the SystemName API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlSystemName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNameSpec   `json:"spec,omitempty"`
	Status SystemNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNameList contains a list of SystemNames
type SrlSystemNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemName `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemName{}, &SrlSystemNameList{})
}

// SystemName type metadata.
var (
	SystemNameKindKind         = reflect.TypeOf(SrlSystemName{}).Name()
	SystemNameGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNameKindKind}.String()
	SystemNameKindAPIVersion   = SystemNameKindKind + "." + GroupVersion.String()
	SystemNameGroupVersionKind = GroupVersion.WithKind(SystemNameKindKind)
)
