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

// NetworkinstanceAggregateroutes struct
type NetworkinstanceAggregateroutes struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=16384
	Route []*NetworkinstanceAggregateroutesRoute `json:"route,omitempty"`
}

// NetworkinstanceAggregateroutesRoute struct
type NetworkinstanceAggregateroutesRoute struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceAggregateroutesRouteAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Aggregator   *NetworkinstanceAggregateroutesRouteAggregator  `json:"aggregator,omitempty"`
	Communities  *NetworkinstanceAggregateroutesRouteCommunities `json:"communities,omitempty"`
	Generateicmp *bool                                           `json:"generate-icmp,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
	// +kubebuilder:default:=false
	Summaryonly *bool `json:"summary-only,omitempty"`
}

// NetworkinstanceAggregateroutesRouteAggregator struct
type NetworkinstanceAggregateroutesRouteAggregator struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Asnumber *uint32 `json:"as-number,omitempty"`
}

// NetworkinstanceAggregateroutesRouteCommunities struct
type NetworkinstanceAggregateroutesRouteCommunities struct {
	Add *NetworkinstanceAggregateroutesRouteCommunitiesAdd `json:"add,omitempty"`
}

// NetworkinstanceAggregateroutesRouteCommunitiesAdd struct
type NetworkinstanceAggregateroutesRouteCommunitiesAdd struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|.*:.*|([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})|.*:.*:.*`
	Add *string `json:"add,omitempty"`
}

// A NetworkinstanceAggregateroutesSpec defines the desired state of a NetworkinstanceAggregateroutes.
type NetworkinstanceAggregateroutesSpec struct {
	nddv1.ResourceSpec             `json:",inline"`
	NetworkInstanceName            *string                         `json:"network-instance-name"`
	NetworkinstanceAggregateroutes *NetworkinstanceAggregateroutes `json:"aggregate-routes,omitempty"`
}

// A NetworkinstanceAggregateroutesStatus represents the observed state of a NetworkinstanceAggregateroutes.
type NetworkinstanceAggregateroutesStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceAggregateroutes is the Schema for the NetworkinstanceAggregateroutes API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl2}
type SrlNetworkinstanceAggregateroutes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceAggregateroutesSpec   `json:"spec,omitempty"`
	Status NetworkinstanceAggregateroutesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceAggregateroutesList contains a list of NetworkinstanceAggregateroutess
type SrlNetworkinstanceAggregateroutesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceAggregateroutes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceAggregateroutes{}, &SrlNetworkinstanceAggregateroutesList{})
}

// NetworkinstanceAggregateroutes type metadata.
var (
	NetworkinstanceAggregateroutesKindKind         = reflect.TypeOf(SrlNetworkinstanceAggregateroutes{}).Name()
	NetworkinstanceAggregateroutesGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceAggregateroutesKindKind}.String()
	NetworkinstanceAggregateroutesKindAPIVersion   = NetworkinstanceAggregateroutesKindKind + "." + GroupVersion.String()
	NetworkinstanceAggregateroutesGroupVersionKind = GroupVersion.WithKind(NetworkinstanceAggregateroutesKindKind)
)
