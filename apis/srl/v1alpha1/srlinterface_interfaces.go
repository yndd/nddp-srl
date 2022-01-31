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

var _ IFSrlInterfaceList = &SrlInterfaceList{}

// +k8s:deepcopy-gen=false
type IFSrlInterfaceList interface {
	client.ObjectList

	GetInterfaces() []IFSrlInterface
}

func (x *SrlInterfaceList) GetInterfaces() []IFSrlInterface {
	xs := make([]IFSrlInterface, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlInterface = &SrlInterface{}

// +k8s:deepcopy-gen=false
type IFSrlInterface interface {
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
	GetInterfaceAdminState() E_InterfaceAdminstate
	GetInterfaceBreakoutMode() *InterfaceBreakoutmode
	GetInterfaceDescription() string
	GetInterfaceEthernet() *InterfaceEthernet
	GetInterfaceLag() *InterfaceLag
	GetInterfaceLoopbackMode() bool
	GetInterfaceMtu() uint16
	GetInterfaceName() string
	GetInterfaceQos() *InterfaceQos
	GetInterfaceSflow() *InterfaceSflow
	GetInterfaceTransceiver() *InterfaceTransceiver
	GetInterfaceVlanTagging() bool
	// add based on type
}

// GetCondition
func (x *SrlInterface) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlInterface) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlInterface) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterface) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterface) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterface) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterface) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterface) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterface) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlInterface) GetInterfaceAdminState() E_InterfaceAdminstate {
	if reflect.ValueOf(x.Spec.Interface.Adminstate).IsZero() {
		return ""
	}
	return E_InterfaceAdminstate(x.Spec.Interface.Adminstate)
}
func (x *SrlInterface) GetInterfaceBreakoutMode() *InterfaceBreakoutmode {
	if reflect.ValueOf(x.Spec.Interface.Breakoutmode).IsZero() {
		return nil
	}
	return x.Spec.Interface.Breakoutmode
}
func (x *SrlInterface) GetInterfaceDescription() string {
	if reflect.ValueOf(x.Spec.Interface.Description).IsZero() {
		return ""
	}
	return *x.Spec.Interface.Description
}
func (x *SrlInterface) GetInterfaceEthernet() *InterfaceEthernet {
	if reflect.ValueOf(x.Spec.Interface.Ethernet).IsZero() {
		return nil
	}
	return x.Spec.Interface.Ethernet
}
func (x *SrlInterface) GetInterfaceLag() *InterfaceLag {
	if reflect.ValueOf(x.Spec.Interface.Lag).IsZero() {
		return nil
	}
	return x.Spec.Interface.Lag
}
func (x *SrlInterface) GetInterfaceLoopbackMode() bool {
	if reflect.ValueOf(x.Spec.Interface.Loopbackmode).IsZero() {
		return false
	}
	return *x.Spec.Interface.Loopbackmode
}
func (x *SrlInterface) GetInterfaceMtu() uint16 {
	if reflect.ValueOf(x.Spec.Interface.Mtu).IsZero() {
		return 0
	}
	return *x.Spec.Interface.Mtu
}
func (x *SrlInterface) GetInterfaceName() string {
	if reflect.ValueOf(x.Spec.Interface.Name).IsZero() {
		return ""
	}
	return *x.Spec.Interface.Name
}
func (x *SrlInterface) GetInterfaceQos() *InterfaceQos {
	if reflect.ValueOf(x.Spec.Interface.Qos).IsZero() {
		return nil
	}
	return x.Spec.Interface.Qos
}
func (x *SrlInterface) GetInterfaceSflow() *InterfaceSflow {
	if reflect.ValueOf(x.Spec.Interface.Sflow).IsZero() {
		return nil
	}
	return x.Spec.Interface.Sflow
}
func (x *SrlInterface) GetInterfaceTransceiver() *InterfaceTransceiver {
	if reflect.ValueOf(x.Spec.Interface.Transceiver).IsZero() {
		return nil
	}
	return x.Spec.Interface.Transceiver
}
func (x *SrlInterface) GetInterfaceVlanTagging() bool {
	if reflect.ValueOf(x.Spec.Interface.Vlantagging).IsZero() {
		return false
	}
	return *x.Spec.Interface.Vlantagging
}
