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
// Policy
type E_RoutingpolicyPolicyDefaultactionAcceptBgpOriginSet string

const (
	E_RoutingpolicyPolicyDefaultactionAcceptBgpOriginSet_Egp        E_RoutingpolicyPolicyDefaultactionAcceptBgpOriginSet = "egp"
	E_RoutingpolicyPolicyDefaultactionAcceptBgpOriginSet_Igp        E_RoutingpolicyPolicyDefaultactionAcceptBgpOriginSet = "igp"
	E_RoutingpolicyPolicyDefaultactionAcceptBgpOriginSet_Incomplete E_RoutingpolicyPolicyDefaultactionAcceptBgpOriginSet = "incomplete"
)

// Policy
type E_RoutingpolicyPolicyStatementActionAcceptBgpOriginSet string

const (
	E_RoutingpolicyPolicyStatementActionAcceptBgpOriginSet_Egp        E_RoutingpolicyPolicyStatementActionAcceptBgpOriginSet = "egp"
	E_RoutingpolicyPolicyStatementActionAcceptBgpOriginSet_Igp        E_RoutingpolicyPolicyStatementActionAcceptBgpOriginSet = "igp"
	E_RoutingpolicyPolicyStatementActionAcceptBgpOriginSet_Incomplete E_RoutingpolicyPolicyStatementActionAcceptBgpOriginSet = "incomplete"
)

// Policy
type E_RoutingpolicyPolicyStatementMatchBgpAspathlengthOperator string

const (
	E_RoutingpolicyPolicyStatementMatchBgpAspathlengthOperator_Eq E_RoutingpolicyPolicyStatementMatchBgpAspathlengthOperator = "eq"
	E_RoutingpolicyPolicyStatementMatchBgpAspathlengthOperator_Ge E_RoutingpolicyPolicyStatementMatchBgpAspathlengthOperator = "ge"
	E_RoutingpolicyPolicyStatementMatchBgpAspathlengthOperator_Le E_RoutingpolicyPolicyStatementMatchBgpAspathlengthOperator = "le"
)

// Policy
type E_RoutingpolicyPolicyStatementMatchIsisRoutetype string

const (
	E_RoutingpolicyPolicyStatementMatchIsisRoutetype_External E_RoutingpolicyPolicyStatementMatchIsisRoutetype = "external"
	E_RoutingpolicyPolicyStatementMatchIsisRoutetype_Internal E_RoutingpolicyPolicyStatementMatchIsisRoutetype = "internal"
)

func NewRoutingpolicyPolicy() *RoutingpolicyPolicy {
	return &RoutingpolicyPolicy{}
}
