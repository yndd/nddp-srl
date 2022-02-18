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

// A TransactionSpec defines the desired state of a Transaction.
type TransactionSpec struct {
	OwnerKind            *string `json:"owner-kind,omitempty"`
	OwnerGeneration      *string `json:"owner-generation,omitempty"`
	OwnerResourceVersion *string `json:"owner-resource-version,omitempty"`
}

// A TransactionStatus represents the observed state of a Transaction.
type TransactionStatus struct {
	nddv1.TransactionResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlTransaction is the Schema for the Transaction API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndd,srl2}
type SrlTransaction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TransactionSpec   `json:"spec,omitempty"`
	Status TransactionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlTransactionList contains a list of Transactions
type SrlTransactionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlTransaction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlTransaction{}, &SrlTransactionList{})
}

// Transaction type metadata.
var (
	TransactionKindKind         = reflect.TypeOf(SrlTransaction{}).Name()
	TransactionGroupKind        = schema.GroupKind{Group: Group, Kind: TransactionKindKind}.String()
	TransactionKindAPIVersion   = TransactionKindKind + "." + GroupVersion.String()
	TransactionGroupVersionKind = GroupVersion.WithKind(TransactionKindKind)
)
