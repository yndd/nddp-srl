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

// RoutingpolicyPrefixset struct
type RoutingpolicyPrefixset struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Prefix []*RoutingpolicyPrefixsetPrefix `json:"prefix,omitempty"`
}

// RoutingpolicyPrefixsetPrefix struct
type RoutingpolicyPrefixsetPrefix struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipprefix *string `json:"ip-prefix"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([0-9]+\.\.[0-9]+)|exact`
	Masklengthrange *string `json:"mask-length-range"`
}

// A RoutingpolicyPrefixsetSpec defines the desired state of a RoutingpolicyPrefixset.
type RoutingpolicyPrefixsetSpec struct {
	nddv1.ResourceSpec     `json:",inline"`
	RoutingpolicyPrefixset *RoutingpolicyPrefixset `json:"prefix-set,omitempty"`
}

// A RoutingpolicyPrefixsetStatus represents the observed state of a RoutingpolicyPrefixset.
type RoutingpolicyPrefixsetStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyPrefixset is the Schema for the RoutingpolicyPrefixset API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlRoutingpolicyPrefixset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoutingpolicyPrefixsetSpec   `json:"spec,omitempty"`
	Status RoutingpolicyPrefixsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyPrefixsetList contains a list of RoutingpolicyPrefixsets
type SrlRoutingpolicyPrefixsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlRoutingpolicyPrefixset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlRoutingpolicyPrefixset{}, &SrlRoutingpolicyPrefixsetList{})
}

// RoutingpolicyPrefixset type metadata.
var (
	RoutingpolicyPrefixsetKindKind         = reflect.TypeOf(SrlRoutingpolicyPrefixset{}).Name()
	RoutingpolicyPrefixsetGroupKind        = schema.GroupKind{Group: Group, Kind: RoutingpolicyPrefixsetKindKind}.String()
	RoutingpolicyPrefixsetKindAPIVersion   = RoutingpolicyPrefixsetKindKind + "." + GroupVersion.String()
	RoutingpolicyPrefixsetGroupVersionKind = GroupVersion.WithKind(RoutingpolicyPrefixsetKindKind)
)
