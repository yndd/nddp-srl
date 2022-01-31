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

/*
import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)
*/
// Networkinstance
type E_NetworkinstanceAdminstate string

const (
	E_NetworkinstanceAdminstate_Disable E_NetworkinstanceAdminstate = "disable"
	E_NetworkinstanceAdminstate_Enable  E_NetworkinstanceAdminstate = "enable"
)

// Networkinstance
type E_NetworkinstanceBridgetableMacduplicationAction string

const (
	E_NetworkinstanceBridgetableMacduplicationAction_Blackhole    E_NetworkinstanceBridgetableMacduplicationAction = "blackhole"
	E_NetworkinstanceBridgetableMacduplicationAction_OperDown     E_NetworkinstanceBridgetableMacduplicationAction = "oper-down"
	E_NetworkinstanceBridgetableMacduplicationAction_StopLearning E_NetworkinstanceBridgetableMacduplicationAction = "stop-learning"
)

// Networkinstance
type E_NetworkinstanceBridgetableMacduplicationAdminstate string

const (
	E_NetworkinstanceBridgetableMacduplicationAdminstate_Disable E_NetworkinstanceBridgetableMacduplicationAdminstate = "disable"
	E_NetworkinstanceBridgetableMacduplicationAdminstate_Enable  E_NetworkinstanceBridgetableMacduplicationAdminstate = "enable"
)

// Networkinstance
type E_NetworkinstanceBridgetableMaclearningAdminstate string

const (
	E_NetworkinstanceBridgetableMaclearningAdminstate_Disable E_NetworkinstanceBridgetableMaclearningAdminstate = "disable"
	E_NetworkinstanceBridgetableMaclearningAdminstate_Enable  E_NetworkinstanceBridgetableMaclearningAdminstate = "enable"
)

// Networkinstance
type E_NetworkinstanceBridgetableMaclearningAgingAdminstate string

const (
	E_NetworkinstanceBridgetableMaclearningAgingAdminstate_Disable E_NetworkinstanceBridgetableMaclearningAgingAdminstate = "disable"
	E_NetworkinstanceBridgetableMaclearningAgingAdminstate_Enable  E_NetworkinstanceBridgetableMaclearningAgingAdminstate = "enable"
)

// Networkinstance
type E_NetworkinstanceMplsAdminstate string

const (
	E_NetworkinstanceMplsAdminstate_Disable E_NetworkinstanceMplsAdminstate = "disable"
	E_NetworkinstanceMplsAdminstate_Enable  E_NetworkinstanceMplsAdminstate = "enable"
)

// Networkinstance
type E_NetworkinstanceMplsStaticentryAdminstate string

const (
	E_NetworkinstanceMplsStaticentryAdminstate_Disable E_NetworkinstanceMplsStaticentryAdminstate = "disable"
	E_NetworkinstanceMplsStaticentryAdminstate_Enable  E_NetworkinstanceMplsStaticentryAdminstate = "enable"
)

// Networkinstance
type E_NetworkinstanceMplsStaticentryOperation string

const (
	E_NetworkinstanceMplsStaticentryOperation_Pop  E_NetworkinstanceMplsStaticentryOperation = "pop"
	E_NetworkinstanceMplsStaticentryOperation_Swap E_NetworkinstanceMplsStaticentryOperation = "swap"
)

// Networkinstance
type E_NetworkinstanceProtocolsLdpAdminstate string

const (
	E_NetworkinstanceProtocolsLdpAdminstate_Disable E_NetworkinstanceProtocolsLdpAdminstate = "disable"
	E_NetworkinstanceProtocolsLdpAdminstate_Enable  E_NetworkinstanceProtocolsLdpAdminstate = "enable"
)

// Networkinstance
type E_NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Adminstate string

const (
	E_NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Adminstate_Disable E_NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Adminstate = "disable"
	E_NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Adminstate_Enable  E_NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4Adminstate = "enable"
)

// Networkinstance
type E_NetworkinstanceTepoliciesStaticpolicyAdminstate string

const (
	E_NetworkinstanceTepoliciesStaticpolicyAdminstate_Disable E_NetworkinstanceTepoliciesStaticpolicyAdminstate = "disable"
	E_NetworkinstanceTepoliciesStaticpolicyAdminstate_Enable  E_NetworkinstanceTepoliciesStaticpolicyAdminstate = "enable"
)

// Networkinstance
type E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType string

const (
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType1  E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-1"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType10 E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-10"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType11 E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-11"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType2  E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-2"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType3  E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-3"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType4  E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-4"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType5  E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-5"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType6  E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-6"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType7  E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-7"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType8  E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-8"
	E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType_SegmentType9  E_NetworkinstanceTepoliciesStaticpolicySegmentlistSegmentType = "segment-type-9"
)

func NewNetworkinstance() *Networkinstance {
	return &Networkinstance{}
}
