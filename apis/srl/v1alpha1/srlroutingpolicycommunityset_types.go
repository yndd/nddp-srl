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

// RoutingpolicyCommunityset struct
type RoutingpolicyCommunityset struct {
	Member *RoutingpolicyCommunitysetMember `json:"member,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// RoutingpolicyCommunitysetMember struct
type RoutingpolicyCommunitysetMember struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|.*:.*|([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})|.*:.*:.*`
	Member *string `json:"member,omitempty"`
}

// A RoutingpolicyCommunitysetSpec defines the desired state of a RoutingpolicyCommunityset.
type RoutingpolicyCommunitysetSpec struct {
	nddv1.ResourceSpec        `json:",inline"`
	RoutingpolicyCommunityset *RoutingpolicyCommunityset `json:"community-set,omitempty"`
}

// A RoutingpolicyCommunitysetStatus represents the observed state of a RoutingpolicyCommunityset.
type RoutingpolicyCommunitysetStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyCommunityset is the Schema for the RoutingpolicyCommunityset API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlRoutingpolicyCommunityset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoutingpolicyCommunitysetSpec   `json:"spec,omitempty"`
	Status RoutingpolicyCommunitysetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyCommunitysetList contains a list of RoutingpolicyCommunitysets
type SrlRoutingpolicyCommunitysetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlRoutingpolicyCommunityset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlRoutingpolicyCommunityset{}, &SrlRoutingpolicyCommunitysetList{})
}

// RoutingpolicyCommunityset type metadata.
var (
	RoutingpolicyCommunitysetKindKind         = reflect.TypeOf(SrlRoutingpolicyCommunityset{}).Name()
	RoutingpolicyCommunitysetGroupKind        = schema.GroupKind{Group: Group, Kind: RoutingpolicyCommunitysetKindKind}.String()
	RoutingpolicyCommunitysetKindAPIVersion   = RoutingpolicyCommunitysetKindKind + "." + GroupVersion.String()
	RoutingpolicyCommunitysetGroupVersionKind = GroupVersion.WithKind(RoutingpolicyCommunitysetKindKind)
)
