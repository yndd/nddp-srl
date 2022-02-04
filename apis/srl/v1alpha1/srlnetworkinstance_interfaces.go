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

var _ IFSrlNetworkinstanceList = &SrlNetworkinstanceList{}

// +k8s:deepcopy-gen=false
type IFSrlNetworkinstanceList interface {
	client.ObjectList

	GetNetworkinstances() []IFSrlNetworkinstance
}

func (x *SrlNetworkinstanceList) GetNetworkinstances() []IFSrlNetworkinstance {
	xs := make([]IFSrlNetworkinstance, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlNetworkinstance = &SrlNetworkinstance{}

// +k8s:deepcopy-gen=false
type IFSrlNetworkinstance interface {
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
	GetNetworkinstanceAdminState() E_NetworkinstanceAdminstate
	GetNetworkinstanceBridgeTable() *NetworkinstanceBridgetable
	GetNetworkinstanceDescription() string
	GetNetworkinstanceInterface() []*NetworkinstanceInterface
	GetNetworkinstanceIpForwarding() *NetworkinstanceIpforwarding
	GetNetworkinstanceIpLoadBalancing() *NetworkinstanceIploadbalancing
	GetNetworkinstanceMpls() *NetworkinstanceMpls
	GetNetworkinstanceMtu() *NetworkinstanceMtu
	GetNetworkinstanceName() string
	GetNetworkinstanceProtocols() *NetworkinstanceProtocols
	GetNetworkinstanceRouterId() string
	GetNetworkinstanceSegmentRouting() *NetworkinstanceSegmentrouting
	GetNetworkinstanceTePolicies() *NetworkinstanceTepolicies
	GetNetworkinstanceTrafficEngineering() *NetworkinstanceTrafficengineering
	GetNetworkinstanceType() string
	GetNetworkinstanceVxlanInterface() []*NetworkinstanceVxlaninterface
	// add based on type
	AddNetworkinstanceInterface(a *NetworkinstanceInterface)
	AddNetworkinstanceVxlanInterface(a *NetworkinstanceVxlaninterface)
}

// GetCondition
func (x *SrlNetworkinstance) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlNetworkinstance) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlNetworkinstance) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstance) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstance) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstance) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstance) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstance) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstance) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlNetworkinstance) GetNetworkinstanceAdminState() E_NetworkinstanceAdminstate {
	if reflect.ValueOf(x.Spec.Networkinstance.Adminstate).IsZero() {
		return ""
	}
	return E_NetworkinstanceAdminstate(x.Spec.Networkinstance.Adminstate)
}
func (x *SrlNetworkinstance) GetNetworkinstanceBridgeTable() *NetworkinstanceBridgetable {
	if reflect.ValueOf(x.Spec.Networkinstance.Bridgetable).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Bridgetable
}
func (x *SrlNetworkinstance) GetNetworkinstanceDescription() string {
	if reflect.ValueOf(x.Spec.Networkinstance.Description).IsZero() {
		return ""
	}
	return *x.Spec.Networkinstance.Description
}
func (x *SrlNetworkinstance) GetNetworkinstanceInterface() []*NetworkinstanceInterface {
	if reflect.ValueOf(x.Spec.Networkinstance.Interface).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Interface
}
func (x *SrlNetworkinstance) GetNetworkinstanceIpForwarding() *NetworkinstanceIpforwarding {
	if reflect.ValueOf(x.Spec.Networkinstance.Ipforwarding).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Ipforwarding
}
func (x *SrlNetworkinstance) GetNetworkinstanceIpLoadBalancing() *NetworkinstanceIploadbalancing {
	if reflect.ValueOf(x.Spec.Networkinstance.Iploadbalancing).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Iploadbalancing
}
func (x *SrlNetworkinstance) GetNetworkinstanceMpls() *NetworkinstanceMpls {
	if reflect.ValueOf(x.Spec.Networkinstance.Mpls).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Mpls
}
func (x *SrlNetworkinstance) GetNetworkinstanceMtu() *NetworkinstanceMtu {
	if reflect.ValueOf(x.Spec.Networkinstance.Mtu).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Mtu
}
func (x *SrlNetworkinstance) GetNetworkinstanceName() string {
	if reflect.ValueOf(x.Spec.Networkinstance.Name).IsZero() {
		return ""
	}
	return *x.Spec.Networkinstance.Name
}
func (x *SrlNetworkinstance) GetNetworkinstanceProtocols() *NetworkinstanceProtocols {
	if reflect.ValueOf(x.Spec.Networkinstance.Protocols).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Protocols
}
func (x *SrlNetworkinstance) GetNetworkinstanceRouterId() string {
	if reflect.ValueOf(x.Spec.Networkinstance.Routerid).IsZero() {
		return ""
	}
	return *x.Spec.Networkinstance.Routerid
}
func (x *SrlNetworkinstance) GetNetworkinstanceSegmentRouting() *NetworkinstanceSegmentrouting {
	if reflect.ValueOf(x.Spec.Networkinstance.Segmentrouting).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Segmentrouting
}
func (x *SrlNetworkinstance) GetNetworkinstanceTePolicies() *NetworkinstanceTepolicies {
	if reflect.ValueOf(x.Spec.Networkinstance.Tepolicies).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Tepolicies
}
func (x *SrlNetworkinstance) GetNetworkinstanceTrafficEngineering() *NetworkinstanceTrafficengineering {
	if reflect.ValueOf(x.Spec.Networkinstance.Trafficengineering).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Trafficengineering
}
func (x *SrlNetworkinstance) GetNetworkinstanceType() string {
	if reflect.ValueOf(x.Spec.Networkinstance.Type).IsZero() {
		return ""
	}
	return *x.Spec.Networkinstance.Type
}
func (x *SrlNetworkinstance) GetNetworkinstanceVxlanInterface() []*NetworkinstanceVxlaninterface {
	if reflect.ValueOf(x.Spec.Networkinstance.Vxlaninterface).IsZero() {
		return nil
	}
	return x.Spec.Networkinstance.Vxlaninterface
}
func (x *SrlNetworkinstance) AddNetworkinstanceInterface(a *NetworkinstanceInterface) {
	if len(x.Spec.Networkinstance.Interface) == 0 {
		x.Spec.Networkinstance.Interface = make([]*NetworkinstanceInterface, 0)
	}
	found := false
	for _, xx := range x.Spec.Networkinstance.Interface {
		if xx.Name == a.Name {
			found = true
			xx = a
		}
	}
	if !found {
		x.Spec.Networkinstance.Interface = append(x.Spec.Networkinstance.Interface, a)
	}
}
func (x *SrlNetworkinstance) AddNetworkinstanceVxlanInterface(a *NetworkinstanceVxlaninterface) {
	if len(x.Spec.Networkinstance.Vxlaninterface) == 0 {
		x.Spec.Networkinstance.Vxlaninterface = make([]*NetworkinstanceVxlaninterface, 0)
	}
	found := false
	for _, xx := range x.Spec.Networkinstance.Vxlaninterface {
		if xx.Name == a.Name {
			found = true
			xx = a
		}
	}
	if !found {
		x.Spec.Networkinstance.Vxlaninterface = append(x.Spec.Networkinstance.Vxlaninterface, a)
	}

}
