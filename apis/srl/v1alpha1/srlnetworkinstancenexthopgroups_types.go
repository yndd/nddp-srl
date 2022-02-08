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

// NetworkinstanceNexthopgroups struct
type NetworkinstanceNexthopgroups struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Group []*NetworkinstanceNexthopgroupsGroup `json:"group,omitempty"`
}

// NetworkinstanceNexthopgroupsGroup struct
type NetworkinstanceNexthopgroupsGroup struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceNexthopgroupsGroupAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Blackhole *NetworkinstanceNexthopgroupsGroupBlackhole `json:"blackhole,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=128
	Nexthop []*NetworkinstanceNexthopgroupsGroupNexthop `json:"nexthop,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupBlackhole struct
type NetworkinstanceNexthopgroupsGroupBlackhole struct {
	// +kubebuilder:default:=false
	Generateicmp *bool `json:"generate-icmp,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthop struct
type NetworkinstanceNexthopgroupsGroupNexthop struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_NetworkinstanceNexthopgroupsGroupNexthopAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Failuredetection *NetworkinstanceNexthopgroupsGroupNexthopFailuredetection `json:"failure-detection,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Index *uint16 `json:"index"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipaddress            *string                                                       `json:"ip-address,omitempty"`
	Pushedmplslabelstack *NetworkinstanceNexthopgroupsGroupNexthopPushedmplslabelstack `json:"pushed-mpls-label-stack,omitempty"`
	// +kubebuilder:default:=true
	Resolve *bool `json:"resolve,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthopFailuredetection struct
type NetworkinstanceNexthopgroupsGroupNexthopFailuredetection struct {
	Enablebfd *NetworkinstanceNexthopgroupsGroupNexthopFailuredetectionEnablebfd `json:"enable-bfd,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthopFailuredetectionEnablebfd struct
type NetworkinstanceNexthopgroupsGroupNexthopFailuredetectionEnablebfd struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Localaddress *string `json:"local-address"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16384
	Localdiscriminator *uint32 `json:"local-discriminator,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16384
	Remotediscriminator *uint32 `json:"remote-discriminator,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthopPushedmplslabelstack struct
type NetworkinstanceNexthopgroupsGroupNexthopPushedmplslabelstack struct {
	Pushedmplslabelstack *string `json:"pushed-mpls-label-stack,omitempty"`
}

// A NetworkinstanceNexthopgroupsSpec defines the desired state of a NetworkinstanceNexthopgroups.
type NetworkinstanceNexthopgroupsSpec struct {
	nddv1.ResourceSpec           `json:",inline"`
	NetworkInstanceName          *string                       `json:"network-instance-name"`
	NetworkinstanceNexthopgroups *NetworkinstanceNexthopgroups `json:"next-hop-groups,omitempty"`
}

// A NetworkinstanceNexthopgroupsStatus represents the observed state of a NetworkinstanceNexthopgroups.
type NetworkinstanceNexthopgroupsStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceNexthopgroups is the Schema for the NetworkinstanceNexthopgroups API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlNetworkinstanceNexthopgroups struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceNexthopgroupsSpec   `json:"spec,omitempty"`
	Status NetworkinstanceNexthopgroupsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceNexthopgroupsList contains a list of NetworkinstanceNexthopgroupss
type SrlNetworkinstanceNexthopgroupsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceNexthopgroups `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceNexthopgroups{}, &SrlNetworkinstanceNexthopgroupsList{})
}

// NetworkinstanceNexthopgroups type metadata.
var (
	NetworkinstanceNexthopgroupsKindKind         = reflect.TypeOf(SrlNetworkinstanceNexthopgroups{}).Name()
	NetworkinstanceNexthopgroupsGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceNexthopgroupsKindKind}.String()
	NetworkinstanceNexthopgroupsKindAPIVersion   = NetworkinstanceNexthopgroupsKindKind + "." + GroupVersion.String()
	NetworkinstanceNexthopgroupsGroupVersionKind = GroupVersion.WithKind(NetworkinstanceNexthopgroupsKindKind)
)
