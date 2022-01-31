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
// Bfd
type E_BfdMicrobfdsessionsLaginterfaceAdminstate string

const (
	E_BfdMicrobfdsessionsLaginterfaceAdminstate_Disable E_BfdMicrobfdsessionsLaginterfaceAdminstate = "disable"
	E_BfdMicrobfdsessionsLaginterfaceAdminstate_Enable  E_BfdMicrobfdsessionsLaginterfaceAdminstate = "enable"
)

// Bfd
type E_BfdSubinterfaceAdminstate string

const (
	E_BfdSubinterfaceAdminstate_Disable E_BfdSubinterfaceAdminstate = "disable"
	E_BfdSubinterfaceAdminstate_Enable  E_BfdSubinterfaceAdminstate = "enable"
)

func NewBfd() *Bfd {
	return &Bfd{}
}
