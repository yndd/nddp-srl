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
// Subinterface
type E_InterfaceSubinterfaceAdminstate string

const (
	E_InterfaceSubinterfaceAdminstate_Disable E_InterfaceSubinterfaceAdminstate = "disable"
	E_InterfaceSubinterfaceAdminstate_Enable  E_InterfaceSubinterfaceAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceBridgetableMacduplicationAction string

const (
	E_InterfaceSubinterfaceBridgetableMacduplicationAction_Blackhole            E_InterfaceSubinterfaceBridgetableMacduplicationAction = "blackhole"
	E_InterfaceSubinterfaceBridgetableMacduplicationAction_OperDown             E_InterfaceSubinterfaceBridgetableMacduplicationAction = "oper-down"
	E_InterfaceSubinterfaceBridgetableMacduplicationAction_StopLearning         E_InterfaceSubinterfaceBridgetableMacduplicationAction = "stop-learning"
	E_InterfaceSubinterfaceBridgetableMacduplicationAction_UseNetInstanceAction E_InterfaceSubinterfaceBridgetableMacduplicationAction = "use-net-instance-action"
)

// Subinterface
type E_InterfaceSubinterfaceBridgetableMaclearningAdminstate string

const (
	E_InterfaceSubinterfaceBridgetableMaclearningAdminstate_Disable E_InterfaceSubinterfaceBridgetableMaclearningAdminstate = "disable"
	E_InterfaceSubinterfaceBridgetableMaclearningAdminstate_Enable  E_InterfaceSubinterfaceBridgetableMaclearningAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceBridgetableMaclearningAgingAdminstate string

const (
	E_InterfaceSubinterfaceBridgetableMaclearningAgingAdminstate_Disable E_InterfaceSubinterfaceBridgetableMaclearningAgingAdminstate = "disable"
	E_InterfaceSubinterfaceBridgetableMaclearningAgingAdminstate_Enable  E_InterfaceSubinterfaceBridgetableMaclearningAgingAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceIpv4ArpDebugDebug string

const (
	E_InterfaceSubinterfaceIpv4ArpDebugDebug_Messages E_InterfaceSubinterfaceIpv4ArpDebugDebug = "messages"
)

// Subinterface
type E_InterfaceSubinterfaceIpv4ArpEvpnAdvertiseRoutetype string

const (
	E_InterfaceSubinterfaceIpv4ArpEvpnAdvertiseRoutetype_Dynamic E_InterfaceSubinterfaceIpv4ArpEvpnAdvertiseRoutetype = "dynamic"
	E_InterfaceSubinterfaceIpv4ArpEvpnAdvertiseRoutetype_Static  E_InterfaceSubinterfaceIpv4ArpEvpnAdvertiseRoutetype = "static"
)

// Subinterface
type E_InterfaceSubinterfaceIpv4ArpHostroutePopulateRoutetype string

const (
	E_InterfaceSubinterfaceIpv4ArpHostroutePopulateRoutetype_Dynamic E_InterfaceSubinterfaceIpv4ArpHostroutePopulateRoutetype = "dynamic"
	E_InterfaceSubinterfaceIpv4ArpHostroutePopulateRoutetype_Evpn    E_InterfaceSubinterfaceIpv4ArpHostroutePopulateRoutetype = "evpn"
	E_InterfaceSubinterfaceIpv4ArpHostroutePopulateRoutetype_Static  E_InterfaceSubinterfaceIpv4ArpHostroutePopulateRoutetype = "static"
)

// Subinterface
type E_InterfaceSubinterfaceIpv4DhcpclientTraceoptionsTraceTrace string

const (
	E_InterfaceSubinterfaceIpv4DhcpclientTraceoptionsTraceTrace_Messages E_InterfaceSubinterfaceIpv4DhcpclientTraceoptionsTraceTrace = "messages"
)

// Subinterface
type E_InterfaceSubinterfaceIpv4DhcprelayAdminstate string

const (
	E_InterfaceSubinterfaceIpv4DhcprelayAdminstate_Disable E_InterfaceSubinterfaceIpv4DhcprelayAdminstate = "disable"
	E_InterfaceSubinterfaceIpv4DhcprelayAdminstate_Enable  E_InterfaceSubinterfaceIpv4DhcprelayAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceIpv4DhcprelayOptionOption string

const (
	E_InterfaceSubinterfaceIpv4DhcprelayOptionOption_CircuitId E_InterfaceSubinterfaceIpv4DhcprelayOptionOption = "circuit-id"
	E_InterfaceSubinterfaceIpv4DhcprelayOptionOption_RemoteId  E_InterfaceSubinterfaceIpv4DhcprelayOptionOption = "remote-id"
)

// Subinterface
type E_InterfaceSubinterfaceIpv4DhcprelayTraceoptionsTraceTrace string

const (
	E_InterfaceSubinterfaceIpv4DhcprelayTraceoptionsTraceTrace_Messages E_InterfaceSubinterfaceIpv4DhcprelayTraceoptionsTraceTrace = "messages"
)

// Subinterface
type E_InterfaceSubinterfaceIpv4DhcpserverAdminstate string

const (
	E_InterfaceSubinterfaceIpv4DhcpserverAdminstate_Disable E_InterfaceSubinterfaceIpv4DhcpserverAdminstate = "disable"
	E_InterfaceSubinterfaceIpv4DhcpserverAdminstate_Enable  E_InterfaceSubinterfaceIpv4DhcpserverAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceIpv4VrrpVrrpgroupAdminstate string

const (
	E_InterfaceSubinterfaceIpv4VrrpVrrpgroupAdminstate_Disable E_InterfaceSubinterfaceIpv4VrrpVrrpgroupAdminstate = "disable"
	E_InterfaceSubinterfaceIpv4VrrpVrrpgroupAdminstate_Enable  E_InterfaceSubinterfaceIpv4VrrpVrrpgroupAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6DhcpclientTraceoptionsTraceTrace string

const (
	E_InterfaceSubinterfaceIpv6DhcpclientTraceoptionsTraceTrace_Messages E_InterfaceSubinterfaceIpv6DhcpclientTraceoptionsTraceTrace = "messages"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6DhcprelayAdminstate string

const (
	E_InterfaceSubinterfaceIpv6DhcprelayAdminstate_Disable E_InterfaceSubinterfaceIpv6DhcprelayAdminstate = "disable"
	E_InterfaceSubinterfaceIpv6DhcprelayAdminstate_Enable  E_InterfaceSubinterfaceIpv6DhcprelayAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6DhcprelayOptionOption string

const (
	E_InterfaceSubinterfaceIpv6DhcprelayOptionOption_InterfaceId E_InterfaceSubinterfaceIpv6DhcprelayOptionOption = "interface-id"
	E_InterfaceSubinterfaceIpv6DhcprelayOptionOption_RemoteId    E_InterfaceSubinterfaceIpv6DhcprelayOptionOption = "remote-id"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6DhcprelayTraceoptionsTraceTrace string

const (
	E_InterfaceSubinterfaceIpv6DhcprelayTraceoptionsTraceTrace_Messages E_InterfaceSubinterfaceIpv6DhcprelayTraceoptionsTraceTrace = "messages"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6Dhcpv6serverAdminstate string

const (
	E_InterfaceSubinterfaceIpv6Dhcpv6serverAdminstate_Disable E_InterfaceSubinterfaceIpv6Dhcpv6serverAdminstate = "disable"
	E_InterfaceSubinterfaceIpv6Dhcpv6serverAdminstate_Enable  E_InterfaceSubinterfaceIpv6Dhcpv6serverAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited string

const (
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited_Both      E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited = "both"
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited_Global    E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited = "global"
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited_LinkLocal E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited = "link-local"
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited_None      E_InterfaceSubinterfaceIpv6NeighbordiscoveryLearnunsolicited = "none"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6NeighbordiscoveryDebugDebug string

const (
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryDebugDebug_Messages E_InterfaceSubinterfaceIpv6NeighbordiscoveryDebugDebug = "messages"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertiseRoutetype string

const (
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertiseRoutetype_Dynamic E_InterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertiseRoutetype = "dynamic"
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertiseRoutetype_Static  E_InterfaceSubinterfaceIpv6NeighbordiscoveryEvpnAdvertiseRoutetype = "static"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulateRoutetype string

const (
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulateRoutetype_Dynamic E_InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulateRoutetype = "dynamic"
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulateRoutetype_Evpn    E_InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulateRoutetype = "evpn"
	E_InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulateRoutetype_Static  E_InterfaceSubinterfaceIpv6NeighbordiscoveryHostroutePopulateRoutetype = "static"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6RouteradvertisementDebugDebug string

const (
	E_InterfaceSubinterfaceIpv6RouteradvertisementDebugDebug_Messages E_InterfaceSubinterfaceIpv6RouteradvertisementDebugDebug = "messages"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6RouteradvertisementRouterroleAdminstate string

const (
	E_InterfaceSubinterfaceIpv6RouteradvertisementRouterroleAdminstate_Disable E_InterfaceSubinterfaceIpv6RouteradvertisementRouterroleAdminstate = "disable"
	E_InterfaceSubinterfaceIpv6RouteradvertisementRouterroleAdminstate_Enable  E_InterfaceSubinterfaceIpv6RouteradvertisementRouterroleAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceIpv6VrrpVrrpgroupAdminstate string

const (
	E_InterfaceSubinterfaceIpv6VrrpVrrpgroupAdminstate_Disable E_InterfaceSubinterfaceIpv6VrrpVrrpgroupAdminstate = "disable"
	E_InterfaceSubinterfaceIpv6VrrpVrrpgroupAdminstate_Enable  E_InterfaceSubinterfaceIpv6VrrpVrrpgroupAdminstate = "enable"
)

// Subinterface
type E_InterfaceSubinterfaceLocalmirrordestinationAdminstate string

const (
	E_InterfaceSubinterfaceLocalmirrordestinationAdminstate_Disable E_InterfaceSubinterfaceLocalmirrordestinationAdminstate = "disable"
	E_InterfaceSubinterfaceLocalmirrordestinationAdminstate_Enable  E_InterfaceSubinterfaceLocalmirrordestinationAdminstate = "enable"
)

func NewInterfaceSubinterface() *InterfaceSubinterface {
	return &InterfaceSubinterface{}
}
