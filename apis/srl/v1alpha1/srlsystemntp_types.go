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

// SystemNtp struct
type SystemNtp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_SystemNtpAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	Networkinstance *string `json:"network-instance"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Server []*SystemNtpServer `json:"server,omitempty"`
}

// SystemNtpServer struct
type SystemNtpServer struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address"`
	// +kubebuilder:default:=false
	Iburst *bool `json:"iburst,omitempty"`
	// +kubebuilder:default:=false
	Prefer *bool `json:"prefer,omitempty"`
}

// A SystemNtpSpec defines the desired state of a SystemNtp.
type SystemNtpSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	SystemNtp          *SystemNtp `json:"ntp,omitempty"`
}

// A SystemNtpStatus represents the observed state of a SystemNtp.
type SystemNtpStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlSystemNtp is the Schema for the SystemNtp API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={ndd,srl2}
type SrlSystemNtp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNtpSpec   `json:"spec,omitempty"`
	Status SystemNtpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNtpList contains a list of SystemNtps
type SrlSystemNtpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNtp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNtp{}, &SrlSystemNtpList{})
}

// SystemNtp type metadata.
var (
	SystemNtpKindKind         = reflect.TypeOf(SrlSystemNtp{}).Name()
	SystemNtpGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNtpKindKind}.String()
	SystemNtpKindAPIVersion   = SystemNtpKindKind + "." + GroupVersion.String()
	SystemNtpGroupVersionKind = GroupVersion.WithKind(SystemNtpKindKind)
)
