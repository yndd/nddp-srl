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

var _ IFSrlNetworkinstanceNexthopgroupsList = &SrlNetworkinstanceNexthopgroupsList{}

// +k8s:deepcopy-gen=false
type IFSrlNetworkinstanceNexthopgroupsList interface {
	client.ObjectList

	GetNetworkinstanceNexthopgroupss() []IFSrlNetworkinstanceNexthopgroups
}

func (x *SrlNetworkinstanceNexthopgroupsList) GetNetworkinstanceNexthopgroupss() []IFSrlNetworkinstanceNexthopgroups {
	xs := make([]IFSrlNetworkinstanceNexthopgroups, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlNetworkinstanceNexthopgroups = &SrlNetworkinstanceNexthopgroups{}

// +k8s:deepcopy-gen=false
type IFSrlNetworkinstanceNexthopgroups interface {
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
	GetNexthopgroupsGroup() []*NetworkinstanceNexthopgroupsGroup
	// add based on type
	AddNexthopgroupsGroup(a *NetworkinstanceNexthopgroupsGroup)
}

// GetCondition
func (x *SrlNetworkinstanceNexthopgroups) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlNetworkinstanceNexthopgroups) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlNetworkinstanceNexthopgroups) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceNexthopgroups) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceNexthopgroups) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceNexthopgroups) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceNexthopgroups) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceNexthopgroups) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceNexthopgroups) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlNetworkinstanceNexthopgroups) GetNexthopgroupsGroup() []*NetworkinstanceNexthopgroupsGroup {
	if reflect.ValueOf(x.Spec.NetworkinstanceNexthopgroups.Group).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceNexthopgroups.Group
}
func (x *SrlNetworkinstanceNexthopgroups) AddNexthopgroupsGroup(a *NetworkinstanceNexthopgroupsGroup) {
	x.Spec.NetworkinstanceNexthopgroups.Group = append(x.Spec.NetworkinstanceNexthopgroups.Group, a)
}
