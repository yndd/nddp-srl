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
// Interface
type E_InterfaceAdminstate string

const (
	E_InterfaceAdminstate_Disable E_InterfaceAdminstate = "disable"
	E_InterfaceAdminstate_Enable  E_InterfaceAdminstate = "enable"
)

// Interface
type E_InterfaceBreakoutmodeChannelspeed string

const (
	E_InterfaceBreakoutmodeChannelspeed_10g E_InterfaceBreakoutmodeChannelspeed = "10G"
	E_InterfaceBreakoutmodeChannelspeed_25g E_InterfaceBreakoutmodeChannelspeed = "25G"
)

// Interface
type E_InterfaceBreakoutmodeNumchannels string

const (
	E_InterfaceBreakoutmodeNumchannels_4 E_InterfaceBreakoutmodeNumchannels = "4"
)

// Interface
type E_InterfaceEthernetDuplexmode string

const (
	E_InterfaceEthernetDuplexmode_Full E_InterfaceEthernetDuplexmode = "full"
	E_InterfaceEthernetDuplexmode_Half E_InterfaceEthernetDuplexmode = "half"
)

// Interface
type E_InterfaceEthernetPortspeed string

const (
	E_InterfaceEthernetPortspeed_100g E_InterfaceEthernetPortspeed = "100G"
	E_InterfaceEthernetPortspeed_100m E_InterfaceEthernetPortspeed = "100M"
	E_InterfaceEthernetPortspeed_10g  E_InterfaceEthernetPortspeed = "10G"
	E_InterfaceEthernetPortspeed_10m  E_InterfaceEthernetPortspeed = "10M"
	E_InterfaceEthernetPortspeed_1g   E_InterfaceEthernetPortspeed = "1G"
	E_InterfaceEthernetPortspeed_1t   E_InterfaceEthernetPortspeed = "1T"
	E_InterfaceEthernetPortspeed_200g E_InterfaceEthernetPortspeed = "200G"
	E_InterfaceEthernetPortspeed_25g  E_InterfaceEthernetPortspeed = "25G"
	E_InterfaceEthernetPortspeed_400g E_InterfaceEthernetPortspeed = "400G"
	E_InterfaceEthernetPortspeed_40g  E_InterfaceEthernetPortspeed = "40G"
	E_InterfaceEthernetPortspeed_50g  E_InterfaceEthernetPortspeed = "50G"
)

// Interface
type E_InterfaceEthernetStandbysignaling string

const (
	E_InterfaceEthernetStandbysignaling_Lacp     E_InterfaceEthernetStandbysignaling = "lacp"
	E_InterfaceEthernetStandbysignaling_PowerOff E_InterfaceEthernetStandbysignaling = "power-off"
)

// Interface
type E_InterfaceEthernetStormcontrolUnits string

const (
	E_InterfaceEthernetStormcontrolUnits_Kbps       E_InterfaceEthernetStormcontrolUnits = "kbps"
	E_InterfaceEthernetStormcontrolUnits_Percentage E_InterfaceEthernetStormcontrolUnits = "percentage"
)

// Interface
type E_InterfaceLagLacpfallbackmode string

const (
	E_InterfaceLagLacpfallbackmode_Static E_InterfaceLagLacpfallbackmode = "static"
)

// Interface
type E_InterfaceLagLagtype string

const (
	E_InterfaceLagLagtype_Lacp   E_InterfaceLagLagtype = "lacp"
	E_InterfaceLagLagtype_Static E_InterfaceLagLagtype = "static"
)

// Interface
type E_InterfaceLagMemberspeed string

const (
	E_InterfaceLagMemberspeed_100g E_InterfaceLagMemberspeed = "100G"
	E_InterfaceLagMemberspeed_100m E_InterfaceLagMemberspeed = "100M"
	E_InterfaceLagMemberspeed_10g  E_InterfaceLagMemberspeed = "10G"
	E_InterfaceLagMemberspeed_10m  E_InterfaceLagMemberspeed = "10M"
	E_InterfaceLagMemberspeed_1g   E_InterfaceLagMemberspeed = "1G"
	E_InterfaceLagMemberspeed_25g  E_InterfaceLagMemberspeed = "25G"
	E_InterfaceLagMemberspeed_400g E_InterfaceLagMemberspeed = "400G"
	E_InterfaceLagMemberspeed_40g  E_InterfaceLagMemberspeed = "40G"
)

// Interface
type E_InterfaceLagLacpInterval string

const (
	E_InterfaceLagLacpInterval_Fast E_InterfaceLagLacpInterval = "FAST"
	E_InterfaceLagLacpInterval_Slow E_InterfaceLagLacpInterval = "SLOW"
)

// Interface
type E_InterfaceLagLacpLacpmode string

const (
	E_InterfaceLagLacpLacpmode_Active  E_InterfaceLagLacpLacpmode = "ACTIVE"
	E_InterfaceLagLacpLacpmode_Passive E_InterfaceLagLacpLacpmode = "PASSIVE"
)

// Interface
type E_InterfaceSflowAdminstate string

const (
	E_InterfaceSflowAdminstate_Disable E_InterfaceSflowAdminstate = "disable"
	E_InterfaceSflowAdminstate_Enable  E_InterfaceSflowAdminstate = "enable"
)

// Interface
type E_InterfaceTransceiverForwarderrorcorrection string

const (
	E_InterfaceTransceiverForwarderrorcorrection_BaseR    E_InterfaceTransceiverForwarderrorcorrection = "base-r"
	E_InterfaceTransceiverForwarderrorcorrection_Disabled E_InterfaceTransceiverForwarderrorcorrection = "disabled"
	E_InterfaceTransceiverForwarderrorcorrection_Rs108    E_InterfaceTransceiverForwarderrorcorrection = "rs-108"
	E_InterfaceTransceiverForwarderrorcorrection_Rs528    E_InterfaceTransceiverForwarderrorcorrection = "rs-528"
	E_InterfaceTransceiverForwarderrorcorrection_Rs544    E_InterfaceTransceiverForwarderrorcorrection = "rs-544"
)

func NewInterface() *Interface {
	return &Interface{}
}
