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

// Bfd struct
type Bfd struct {
	Microbfdsessions *BfdMicrobfdsessions `json:"micro-bfd-sessions,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Subinterface []*BfdSubinterface `json:"subinterface,omitempty"`
}

// BfdMicrobfdsessions struct
type BfdMicrobfdsessions struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Laginterface []*BfdMicrobfdsessionsLaginterface `json:"lag-interface,omitempty"`
}

// BfdMicrobfdsessionsLaginterface struct
type BfdMicrobfdsessionsLaginterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_BfdMicrobfdsessionsLaginterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	Desiredminimumtransmitinterval *uint32 `json:"desired-minimum-transmit-interval,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=20
	// +kubebuilder:default:=3
	Detectionmultiplier *uint8 `json:"detection-multiplier,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Localaddress *string `json:"local-address,omitempty"`
	Name         *string `json:"name"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Remoteaddress *string `json:"remote-address,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	Requiredminimumreceive *uint32 `json:"required-minimum-receive,omitempty"`
}

// BfdSubinterface struct
type BfdSubinterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate E_BfdSubinterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	Desiredminimumtransmitinterval *uint32 `json:"desired-minimum-transmit-interval,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=20
	// +kubebuilder:default:=3
	Detectionmultiplier *uint8 `json:"detection-multiplier,omitempty"`
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Id *string `json:"id"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	// +kubebuilder:default:=0
	Minimumechoreceiveinterval *uint32 `json:"minimum-echo-receive-interval,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	Requiredminimumreceive *uint32 `json:"required-minimum-receive,omitempty"`
}

// A BfdSpec defines the desired state of a Bfd.
type BfdSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	Bfd                *Bfd `json:"bfd,omitempty"`
}

// A BfdStatus represents the observed state of a Bfd.
type BfdStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlBfd is the Schema for the Bfd API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlBfd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BfdSpec   `json:"spec,omitempty"`
	Status BfdStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlBfdList contains a list of Bfds
type SrlBfdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlBfd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlBfd{}, &SrlBfdList{})
}

// Bfd type metadata.
var (
	BfdKindKind         = reflect.TypeOf(SrlBfd{}).Name()
	BfdGroupKind        = schema.GroupKind{Group: Group, Kind: BfdKindKind}.String()
	BfdKindAPIVersion   = BfdKindKind + "." + GroupVersion.String()
	BfdGroupVersionKind = GroupVersion.WithKind(BfdKindKind)
)
