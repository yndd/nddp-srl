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
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ IFSrlTransactionList = &SrlTransactionList{}

// +k8s:deepcopy-gen=false
type IFSrlTransactionList interface {
	client.ObjectList

	GetTransactions() []IFSrlTransaction
}

func (x *SrlTransactionList) GetTransactions() []IFSrlTransaction {
	xs := make([]IFSrlTransaction, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlTransaction = &SrlTransaction{}

// +k8s:deepcopy-gen=false
type IFSrlTransaction interface {
	resource.Object
	resource.Conditioned

	GetCondition(ct nddv1.ConditionKind) nddv1.Condition
	SetConditions(c ...nddv1.Condition)
	// getters
	GetOwnerGeneration() string
}

// GetCondition
func (x *SrlTransaction) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlTransaction) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlTransaction) GetOwnerGeneration() string {
	return *x.Spec.OwnerGeneration
}

func (x *SrlTransaction) GetDeviceCondition(d string, ct nddv1.ConditionKind) nddv1.Condition {
	if x.Status.Device == nil {
		return nddv1.Condition{}
	}
	dev := x.Status.Device[d]
	return dev.GetCondition(ct)
}

func (x *SrlTransaction) SetDeviceConditions(d string, c ...nddv1.Condition) {
	if x.Status.Device == nil {
		x.Status.Device = make(map[string]nddv1.ConditionedStatus)
	}
	dev := x.Status.Device[d]
	dev.SetConditions(c...)
}
