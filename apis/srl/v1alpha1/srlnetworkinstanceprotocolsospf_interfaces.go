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

var _ IFSrlNetworkinstanceProtocolsOspfList = &SrlNetworkinstanceProtocolsOspfList{}

// +k8s:deepcopy-gen=false
type IFSrlNetworkinstanceProtocolsOspfList interface {
	client.ObjectList

	GetNetworkinstanceProtocolsOspfs() []IFSrlNetworkinstanceProtocolsOspf
}

func (x *SrlNetworkinstanceProtocolsOspfList) GetNetworkinstanceProtocolsOspfs() []IFSrlNetworkinstanceProtocolsOspf {
	xs := make([]IFSrlNetworkinstanceProtocolsOspf, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlNetworkinstanceProtocolsOspf = &SrlNetworkinstanceProtocolsOspf{}

// +k8s:deepcopy-gen=false
type IFSrlNetworkinstanceProtocolsOspf interface {
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
	GetOspfInstance() []*NetworkinstanceProtocolsOspfInstance
	// add based on type
	AddOspfInstance(a *NetworkinstanceProtocolsOspfInstance)
}

// GetCondition
func (x *SrlNetworkinstanceProtocolsOspf) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlNetworkinstanceProtocolsOspf) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlNetworkinstanceProtocolsOspf) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsOspf) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsOspf) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsOspf) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsOspf) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsOspf) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsOspf) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlNetworkinstanceProtocolsOspf) GetOspfInstance() []*NetworkinstanceProtocolsOspfInstance {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsOspf.Instance).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsOspf.Instance
}
func (x *SrlNetworkinstanceProtocolsOspf) AddOspfInstance(a *NetworkinstanceProtocolsOspfInstance) {
	x.Spec.NetworkinstanceProtocolsOspf.Instance = append(x.Spec.NetworkinstanceProtocolsOspf.Instance, a)
}
