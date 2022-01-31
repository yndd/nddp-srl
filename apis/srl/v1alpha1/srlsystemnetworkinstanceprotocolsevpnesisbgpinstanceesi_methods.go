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
// Ethernetsegment
type E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiAdminstate string

const (
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiAdminstate_Disable E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiAdminstate = "disable"
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiAdminstate_Enable  E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiAdminstate = "enable"
)

// Ethernetsegment
type E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiMultihomingmode string

const (
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiMultihomingmode_AllActive    E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiMultihomingmode = "all-active"
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiMultihomingmode_SingleActive E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiMultihomingmode = "single-active"
)

// Ethernetsegment
type E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmType string

const (
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmType_Default    E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmType = "default"
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmType_Preference E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmType = "preference"
)

// Ethernetsegment
type E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealgCapabilitiesAcdf string

const (
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealgCapabilitiesAcdf_Exclude E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealgCapabilitiesAcdf = "exclude"
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealgCapabilitiesAcdf_Include E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealgCapabilitiesAcdf = "include"
)

// Ethernetsegment
type E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesNexthop string

const (
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesNexthop_UseSystemIpv4Address E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesNexthop = "use-system-ipv4-address"
)

// Ethernetsegment
type E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEthernetsegmentOriginatingip string

const (
	E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEthernetsegmentOriginatingip_UseSystemIpv4Address E_SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEthernetsegmentOriginatingip = "use-system-ipv4-address"
)

func NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi() *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi {
	return &SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}
}
