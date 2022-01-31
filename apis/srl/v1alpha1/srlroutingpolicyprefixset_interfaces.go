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

var _ IFSrlRoutingpolicyPrefixsetList = &SrlRoutingpolicyPrefixsetList{}

// +k8s:deepcopy-gen=false
type IFSrlRoutingpolicyPrefixsetList interface {
	client.ObjectList

	GetRoutingpolicyPrefixsets() []IFSrlRoutingpolicyPrefixset
}

func (x *SrlRoutingpolicyPrefixsetList) GetRoutingpolicyPrefixsets() []IFSrlRoutingpolicyPrefixset {
	xs := make([]IFSrlRoutingpolicyPrefixset, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlRoutingpolicyPrefixset = &SrlRoutingpolicyPrefixset{}

// +k8s:deepcopy-gen=false
type IFSrlRoutingpolicyPrefixset interface {
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
	GetPrefixsetName() string
	GetPrefixsetPrefix() []*RoutingpolicyPrefixsetPrefix
	// add based on type
	AddPrefixsetPrefix(a *RoutingpolicyPrefixsetPrefix)
}

// GetCondition
func (x *SrlRoutingpolicyPrefixset) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlRoutingpolicyPrefixset) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlRoutingpolicyPrefixset) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPrefixset) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPrefixset) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPrefixset) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPrefixset) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPrefixset) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlRoutingpolicyPrefixset) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlRoutingpolicyPrefixset) GetPrefixsetName() string {
	if reflect.ValueOf(x.Spec.RoutingpolicyPrefixset.Name).IsZero() {
		return ""
	}
	return *x.Spec.RoutingpolicyPrefixset.Name
}
func (x *SrlRoutingpolicyPrefixset) GetPrefixsetPrefix() []*RoutingpolicyPrefixsetPrefix {
	if reflect.ValueOf(x.Spec.RoutingpolicyPrefixset.Prefix).IsZero() {
		return nil
	}
	return x.Spec.RoutingpolicyPrefixset.Prefix
}
func (x *SrlRoutingpolicyPrefixset) AddPrefixsetPrefix(a *RoutingpolicyPrefixsetPrefix) {
	x.Spec.RoutingpolicyPrefixset.Prefix = append(x.Spec.RoutingpolicyPrefixset.Prefix, a)
}
