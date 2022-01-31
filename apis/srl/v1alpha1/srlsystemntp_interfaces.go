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

var _ IFSrlSystemNtpList = &SrlSystemNtpList{}

// +k8s:deepcopy-gen=false
type IFSrlSystemNtpList interface {
	client.ObjectList

	GetSystemNtps() []IFSrlSystemNtp
}

func (x *SrlSystemNtpList) GetSystemNtps() []IFSrlSystemNtp {
	xs := make([]IFSrlSystemNtp, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlSystemNtp = &SrlSystemNtp{}

// +k8s:deepcopy-gen=false
type IFSrlSystemNtp interface {
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
	GetNtpAdminState() E_SystemNtpAdminstate
	GetNtpNetworkInstance() string
	GetNtpServer() []*SystemNtpServer
	// add based on type
	AddNtpServer(a *SystemNtpServer)
}

// GetCondition
func (x *SrlSystemNtp) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlSystemNtp) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlSystemNtp) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNtp) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNtp) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNtp) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNtp) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNtp) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlSystemNtp) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlSystemNtp) GetNtpAdminState() E_SystemNtpAdminstate {
	if reflect.ValueOf(x.Spec.SystemNtp.Adminstate).IsZero() {
		return ""
	}
	return E_SystemNtpAdminstate(x.Spec.SystemNtp.Adminstate)
}
func (x *SrlSystemNtp) GetNtpNetworkInstance() string {
	if reflect.ValueOf(x.Spec.SystemNtp.Networkinstance).IsZero() {
		return ""
	}
	return *x.Spec.SystemNtp.Networkinstance
}
func (x *SrlSystemNtp) GetNtpServer() []*SystemNtpServer {
	if reflect.ValueOf(x.Spec.SystemNtp.Server).IsZero() {
		return nil
	}
	return x.Spec.SystemNtp.Server
}
func (x *SrlSystemNtp) AddNtpServer(a *SystemNtpServer) {
	x.Spec.SystemNtp.Server = append(x.Spec.SystemNtp.Server, a)
}
