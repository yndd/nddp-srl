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
	"github.com/yndd/nddo-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ IFSrlRoutingpolicyPolicyList = &SrlRoutingpolicyPolicyList{}

// +k8s:deepcopy-gen=false
type IFSrlRoutingpolicyPolicyList interface {
	client.ObjectList

	GetRoutingpolicyPolicys() []IFSrlRoutingpolicyPolicy
}

func (x *SrlRoutingpolicyPolicyList) GetRoutingpolicyPolicys() []IFSrlRoutingpolicyPolicy {
	xs := make([]IFSrlRoutingpolicyPolicy, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlRoutingpolicyPolicy = &SrlRoutingpolicyPolicy{}

// +k8s:deepcopy-gen=false
type IFSrlRoutingpolicyPolicy interface {
	resource.Object
	resource.Conditioned

	GetCondition(ct nddv1.ConditionKind) nddv1.Condition
	SetConditions(c ...nddv1.Condition)
	// getters based on labels
	GetOwner() string
	GetDeploymentPolicy() string
	GetDeviceName() string
	GetEndpointGroup() string
	GetOrganization() string
	GetDeployment() string
	GetAvailabilityZone() string
	// getters based on type
	GetPolicyDefaultAction() *RoutingpolicyPolicyDefaultaction
	GetPolicyName() string
	GetPolicyStatement() []*RoutingpolicyPolicyStatement
	// add based on type
	AddPolicyStatement(a *RoutingpolicyPolicyStatement)
}

// GetCondition
func (x *SrlRoutingpolicyPolicy) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlRoutingpolicyPolicy) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlRoutingpolicyPolicy) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPolicy) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPolicy) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPolicy) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPolicy) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPolicy) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPolicy) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlRoutingpolicyPolicy) GetPolicyDefaultAction() *RoutingpolicyPolicyDefaultaction {
	if reflect.ValueOf(x.Spec.RoutingpolicyPolicy.Defaultaction).IsZero() {
		return nil
	}
	return x.Spec.RoutingpolicyPolicy.Defaultaction
}
func (x *SrlRoutingpolicyPolicy) GetPolicyName() string {
	if reflect.ValueOf(x.Spec.RoutingpolicyPolicy.Name).IsZero() {
		return ""
	}
	return *x.Spec.RoutingpolicyPolicy.Name
}
func (x *SrlRoutingpolicyPolicy) GetPolicyStatement() []*RoutingpolicyPolicyStatement {
	if reflect.ValueOf(x.Spec.RoutingpolicyPolicy.Statement).IsZero() {
		return nil
	}
	return x.Spec.RoutingpolicyPolicy.Statement
}
func (x *SrlRoutingpolicyPolicy) AddPolicyStatement(a *RoutingpolicyPolicyStatement) {
	x.Spec.RoutingpolicyPolicy.Statement = append(x.Spec.RoutingpolicyPolicy.Statement, a)
}
