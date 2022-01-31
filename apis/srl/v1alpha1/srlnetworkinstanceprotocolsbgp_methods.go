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
// Bgp
type E_NetworkinstanceProtocolsBgpAdminstate string

const (
	E_NetworkinstanceProtocolsBgpAdminstate_Disable E_NetworkinstanceProtocolsBgpAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpAdminstate_Enable  E_NetworkinstanceProtocolsBgpAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateasMode string

const (
	E_NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateasMode_Delete   E_NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateasMode = "delete"
	E_NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateasMode_Disabled E_NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateasMode = "disabled"
	E_NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateasMode_Replace  E_NetworkinstanceProtocolsBgpAspathoptionsRemoveprivateasMode = "replace"
)

// Bgp
type E_NetworkinstanceProtocolsBgpEvpnAdminstate string

const (
	E_NetworkinstanceProtocolsBgpEvpnAdminstate_Disable E_NetworkinstanceProtocolsBgpEvpnAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpEvpnAdminstate_Enable  E_NetworkinstanceProtocolsBgpEvpnAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpGracefulrestartAdminstate string

const (
	E_NetworkinstanceProtocolsBgpGracefulrestartAdminstate_Disable E_NetworkinstanceProtocolsBgpGracefulrestartAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpGracefulrestartAdminstate_Enable  E_NetworkinstanceProtocolsBgpGracefulrestartAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpGroupAdminstate string

const (
	E_NetworkinstanceProtocolsBgpGroupAdminstate_Disable E_NetworkinstanceProtocolsBgpGroupAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpGroupAdminstate_Enable  E_NetworkinstanceProtocolsBgpGroupAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateasMode string

const (
	E_NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateasMode_Delete   E_NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateasMode = "delete"
	E_NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateasMode_Disabled E_NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateasMode = "disabled"
	E_NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateasMode_Replace  E_NetworkinstanceProtocolsBgpGroupAspathoptionsRemoveprivateasMode = "replace"
)

// Bgp
type E_NetworkinstanceProtocolsBgpGroupEvpnAdminstate string

const (
	E_NetworkinstanceProtocolsBgpGroupEvpnAdminstate_Disable E_NetworkinstanceProtocolsBgpGroupEvpnAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpGroupEvpnAdminstate_Enable  E_NetworkinstanceProtocolsBgpGroupEvpnAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpGroupGracefulrestartAdminstate string

const (
	E_NetworkinstanceProtocolsBgpGroupGracefulrestartAdminstate_Disable E_NetworkinstanceProtocolsBgpGroupGracefulrestartAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpGroupGracefulrestartAdminstate_Enable  E_NetworkinstanceProtocolsBgpGroupGracefulrestartAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpGroupIpv4unicastAdminstate string

const (
	E_NetworkinstanceProtocolsBgpGroupIpv4unicastAdminstate_Disable E_NetworkinstanceProtocolsBgpGroupIpv4unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpGroupIpv4unicastAdminstate_Enable  E_NetworkinstanceProtocolsBgpGroupIpv4unicastAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpGroupIpv6unicastAdminstate string

const (
	E_NetworkinstanceProtocolsBgpGroupIpv6unicastAdminstate_Disable E_NetworkinstanceProtocolsBgpGroupIpv6unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpGroupIpv6unicastAdminstate_Enable  E_NetworkinstanceProtocolsBgpGroupIpv6unicastAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagModifier string

const (
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagModifier_Detail  E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagModifier = "detail"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagModifier_Receive E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagModifier = "receive"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagModifier_Send    E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagModifier = "send"
)

// Bgp
type E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName string

const (
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_Events          E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "events"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_GracefulRestart E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "graceful-restart"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_Keepalive       E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "keepalive"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_Notification    E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "notification"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_Open            E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "open"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_Packets         E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "packets"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_Route           E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "route"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_Socket          E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "socket"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_Timers          E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "timers"
	E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName_Update          E_NetworkinstanceProtocolsBgpGroupTraceoptionsFlagName = "update"
)

// Bgp
type E_NetworkinstanceProtocolsBgpIpv4unicastAdminstate string

const (
	E_NetworkinstanceProtocolsBgpIpv4unicastAdminstate_Disable E_NetworkinstanceProtocolsBgpIpv4unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpIpv4unicastAdminstate_Enable  E_NetworkinstanceProtocolsBgpIpv4unicastAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode string

const (
	E_NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode_Disabled E_NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode = "disabled"
	E_NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode_Prefer   E_NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode = "prefer"
	E_NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode_Require  E_NetworkinstanceProtocolsBgpIpv4unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode = "require"
)

// Bgp
type E_NetworkinstanceProtocolsBgpIpv6unicastAdminstate string

const (
	E_NetworkinstanceProtocolsBgpIpv6unicastAdminstate_Disable E_NetworkinstanceProtocolsBgpIpv6unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpIpv6unicastAdminstate_Enable  E_NetworkinstanceProtocolsBgpIpv6unicastAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode string

const (
	E_NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode_Disabled E_NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode = "disabled"
	E_NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode_Prefer   E_NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode = "prefer"
	E_NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode_Require  E_NetworkinstanceProtocolsBgpIpv6unicastNexthopresolutionIpv4nexthopsTunnelresolutionMode = "require"
)

// Bgp
type E_NetworkinstanceProtocolsBgpNeighborAdminstate string

const (
	E_NetworkinstanceProtocolsBgpNeighborAdminstate_Disable E_NetworkinstanceProtocolsBgpNeighborAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpNeighborAdminstate_Enable  E_NetworkinstanceProtocolsBgpNeighborAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateasMode string

const (
	E_NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateasMode_Delete   E_NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateasMode = "delete"
	E_NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateasMode_Disabled E_NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateasMode = "disabled"
	E_NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateasMode_Replace  E_NetworkinstanceProtocolsBgpNeighborAspathoptionsRemoveprivateasMode = "replace"
)

// Bgp
type E_NetworkinstanceProtocolsBgpNeighborEvpnAdminstate string

const (
	E_NetworkinstanceProtocolsBgpNeighborEvpnAdminstate_Disable E_NetworkinstanceProtocolsBgpNeighborEvpnAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpNeighborEvpnAdminstate_Enable  E_NetworkinstanceProtocolsBgpNeighborEvpnAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpNeighborGracefulrestartAdminstate string

const (
	E_NetworkinstanceProtocolsBgpNeighborGracefulrestartAdminstate_Disable E_NetworkinstanceProtocolsBgpNeighborGracefulrestartAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpNeighborGracefulrestartAdminstate_Enable  E_NetworkinstanceProtocolsBgpNeighborGracefulrestartAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpNeighborIpv4unicastAdminstate string

const (
	E_NetworkinstanceProtocolsBgpNeighborIpv4unicastAdminstate_Disable E_NetworkinstanceProtocolsBgpNeighborIpv4unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpNeighborIpv4unicastAdminstate_Enable  E_NetworkinstanceProtocolsBgpNeighborIpv4unicastAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpNeighborIpv6unicastAdminstate string

const (
	E_NetworkinstanceProtocolsBgpNeighborIpv6unicastAdminstate_Disable E_NetworkinstanceProtocolsBgpNeighborIpv6unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsBgpNeighborIpv6unicastAdminstate_Enable  E_NetworkinstanceProtocolsBgpNeighborIpv6unicastAdminstate = "enable"
)

// Bgp
type E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagModifier string

const (
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagModifier_Detail  E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagModifier = "detail"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagModifier_Receive E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagModifier = "receive"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagModifier_Send    E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagModifier = "send"
)

// Bgp
type E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName string

const (
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_Events          E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "events"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_GracefulRestart E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "graceful-restart"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_Keepalive       E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "keepalive"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_Notification    E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "notification"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_Open            E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "open"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_Packets         E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "packets"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_Route           E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "route"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_Socket          E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "socket"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_Timers          E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "timers"
	E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName_Update          E_NetworkinstanceProtocolsBgpNeighborTraceoptionsFlagName = "update"
)

// Bgp
type E_NetworkinstanceProtocolsBgpTraceoptionsFlagModifier string

const (
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagModifier_Detail  E_NetworkinstanceProtocolsBgpTraceoptionsFlagModifier = "detail"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagModifier_Receive E_NetworkinstanceProtocolsBgpTraceoptionsFlagModifier = "receive"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagModifier_Send    E_NetworkinstanceProtocolsBgpTraceoptionsFlagModifier = "send"
)

// Bgp
type E_NetworkinstanceProtocolsBgpTraceoptionsFlagName string

const (
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_Events          E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "events"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_GracefulRestart E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "graceful-restart"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_Keepalive       E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "keepalive"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_Notification    E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "notification"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_Open            E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "open"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_Packets         E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "packets"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_Route           E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "route"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_Socket          E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "socket"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_Timers          E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "timers"
	E_NetworkinstanceProtocolsBgpTraceoptionsFlagName_Update          E_NetworkinstanceProtocolsBgpTraceoptionsFlagName = "update"
)

func NewNetworkinstanceProtocolsBgp() *NetworkinstanceProtocolsBgp {
	return &NetworkinstanceProtocolsBgp{}
}
