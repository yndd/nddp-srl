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

var _ IFSrlInterfaceSubinterfaceList = &SrlInterfaceSubinterfaceList{}

// +k8s:deepcopy-gen=false
type IFSrlInterfaceSubinterfaceList interface {
	client.ObjectList

	GetInterfaceSubinterfaces() []IFSrlInterfaceSubinterface
}

func (x *SrlInterfaceSubinterfaceList) GetInterfaceSubinterfaces() []IFSrlInterfaceSubinterface {
	xs := make([]IFSrlInterfaceSubinterface, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlInterfaceSubinterface = &SrlInterfaceSubinterface{}

// +k8s:deepcopy-gen=false
type IFSrlInterfaceSubinterface interface {
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
	GetSubinterfaceAcl() *InterfaceSubinterfaceAcl
	GetSubinterfaceAdminState() E_InterfaceSubinterfaceAdminstate
	GetSubinterfaceAnycastGw() *InterfaceSubinterfaceAnycastgw
	GetSubinterfaceBridgeTable() *InterfaceSubinterfaceBridgetable
	GetSubinterfaceDescription() string
	GetSubinterfaceIndex() uint32
	GetSubinterfaceIpMtu() uint16
	GetSubinterfaceIpv4() *InterfaceSubinterfaceIpv4
	GetSubinterfaceIpv6() *InterfaceSubinterfaceIpv6
	GetSubinterfaceL2Mtu() uint16
	GetSubinterfaceLocalMirrorDestination() *InterfaceSubinterfaceLocalmirrordestination
	GetSubinterfaceMplsMtu() uint16
	GetSubinterfaceQos() *InterfaceSubinterfaceQos
	GetSubinterfaceRaGuard() *InterfaceSubinterfaceRaguard
	GetSubinterfaceType() string
	GetSubinterfaceVlan() *InterfaceSubinterfaceVlan
	// add based on type
}

// GetCondition
func (x *SrlInterfaceSubinterface) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlInterfaceSubinterface) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlInterfaceSubinterface) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterfaceSubinterface) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterfaceSubinterface) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterfaceSubinterface) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterfaceSubinterface) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterfaceSubinterface) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlInterfaceSubinterface) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceAcl() *InterfaceSubinterfaceAcl {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Acl).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Acl
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceAdminState() E_InterfaceSubinterfaceAdminstate {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Adminstate).IsZero() {
		return ""
	}
	return E_InterfaceSubinterfaceAdminstate(x.Spec.InterfaceSubinterface.Adminstate)
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceAnycastGw() *InterfaceSubinterfaceAnycastgw {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Anycastgw).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Anycastgw
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceBridgeTable() *InterfaceSubinterfaceBridgetable {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Bridgetable).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Bridgetable
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceDescription() string {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Description).IsZero() {
		return ""
	}
	return *x.Spec.InterfaceSubinterface.Description
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceIndex() uint32 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Index).IsZero() {
		return 0
	}
	return *x.Spec.InterfaceSubinterface.Index
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceIpMtu() uint16 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Ipmtu).IsZero() {
		return 0
	}
	return *x.Spec.InterfaceSubinterface.Ipmtu
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceIpv4() *InterfaceSubinterfaceIpv4 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Ipv4).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Ipv4
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceIpv6() *InterfaceSubinterfaceIpv6 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Ipv6).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Ipv6
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceL2Mtu() uint16 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.L2mtu).IsZero() {
		return 0
	}
	return *x.Spec.InterfaceSubinterface.L2mtu
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceLocalMirrorDestination() *InterfaceSubinterfaceLocalmirrordestination {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Localmirrordestination).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Localmirrordestination
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceMplsMtu() uint16 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Mplsmtu).IsZero() {
		return 0
	}
	return *x.Spec.InterfaceSubinterface.Mplsmtu
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceQos() *InterfaceSubinterfaceQos {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Qos).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Qos
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceRaGuard() *InterfaceSubinterfaceRaguard {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Raguard).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Raguard
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceType() string {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Type).IsZero() {
		return ""
	}
	return *x.Spec.InterfaceSubinterface.Type
}
func (x *SrlInterfaceSubinterface) GetSubinterfaceVlan() *InterfaceSubinterfaceVlan {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Vlan).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Vlan
}
