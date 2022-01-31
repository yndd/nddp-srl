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

var _ IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList = &SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList{}

// +k8s:deepcopy-gen=false
type IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList interface {
	client.ObjectList

	GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() []IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi
}

func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList) GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() []IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi {
	xs := make([]IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi = &SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}

// +k8s:deepcopy-gen=false
type IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi interface {
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
	GetEthernetsegmentAdminState() E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiAdminstate
	GetEthernetsegmentDfElection() *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelection
	GetEthernetsegmentEsi() string
	GetEthernetsegmentInterface() string
	GetEthernetsegmentMultiHomingMode() E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiMultihomingmode
	GetEthernetsegmentName() string
	GetEthernetsegmentRoutes() *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes
	// add based on type
}

// GetCondition
func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetEthernetsegmentAdminState() E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiAdminstate {
	if reflect.ValueOf(x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Adminstate).IsZero() {
		return ""
	}
	return E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiAdminstate(x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Adminstate)
}
func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetEthernetsegmentDfElection() *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelection {
	if reflect.ValueOf(x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Dfelection).IsZero() {
		return nil
	}
	return x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Dfelection
}
func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetEthernetsegmentEsi() string {
	if reflect.ValueOf(x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Esi).IsZero() {
		return ""
	}
	return *x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Esi
}
func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetEthernetsegmentInterface() string {
	if reflect.ValueOf(x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Interface).IsZero() {
		return ""
	}
	return *x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Interface
}
func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetEthernetsegmentMultiHomingMode() E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiMultihomingmode {
	if reflect.ValueOf(x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Multihomingmode).IsZero() {
		return ""
	}
	return E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiMultihomingmode(x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Multihomingmode)
}
func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetEthernetsegmentName() string {
	if reflect.ValueOf(x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Name).IsZero() {
		return ""
	}
	return *x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Name
}
func (x *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) GetEthernetsegmentRoutes() *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes {
	if reflect.ValueOf(x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Routes).IsZero() {
		return nil
	}
	return x.Spec.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Routes
}
