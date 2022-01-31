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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// RoutingpolicyPolicy struct
type RoutingpolicyPolicy struct {
	Defaultaction *RoutingpolicyPolicyDefaultaction `json:"default-action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	Statement []*RoutingpolicyPolicyStatement `json:"statement,omitempty"`
}

// RoutingpolicyPolicyDefaultaction struct
type RoutingpolicyPolicyDefaultaction struct {
	Accept     *RoutingpolicyPolicyDefaultactionAccept     `json:"accept,omitempty"`
	Nextentry  *RoutingpolicyPolicyDefaultactionNextentry  `json:"next-entry,omitempty"`
	Nextpolicy *RoutingpolicyPolicyDefaultactionNextpolicy `json:"next-policy,omitempty"`
	Reject     *RoutingpolicyPolicyDefaultactionReject     `json:"reject,omitempty"`
}

// RoutingpolicyPolicyDefaultactionAccept struct
type RoutingpolicyPolicyDefaultactionAccept struct {
	Bgp *RoutingpolicyPolicyDefaultactionAcceptBgp `json:"bgp,omitempty"`
}

// RoutingpolicyPolicyDefaultactionAcceptBgp struct
type RoutingpolicyPolicyDefaultactionAcceptBgp struct {
	Aspath          *RoutingpolicyPolicyDefaultactionAcceptBgpAspath          `json:"as-path,omitempty"`
	Communities     *RoutingpolicyPolicyDefaultactionAcceptBgpCommunities     `json:"communities,omitempty"`
	Localpreference *RoutingpolicyPolicyDefaultactionAcceptBgpLocalpreference `json:"local-preference,omitempty"`
	Origin          *RoutingpolicyPolicyDefaultactionAcceptBgpOrigin          `json:"origin,omitempty"`
}

// RoutingpolicyPolicyDefaultactionAcceptBgpAspath struct
type RoutingpolicyPolicyDefaultactionAcceptBgpAspath struct {
	Prepend *RoutingpolicyPolicyDefaultactionAcceptBgpAspathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                   `json:"remove,omitempty"`
	Replace *RoutingpolicyPolicyDefaultactionAcceptBgpAspathReplace `json:"replace,omitempty"`
}

// RoutingpolicyPolicyDefaultactionAcceptBgpAspathPrepend struct
type RoutingpolicyPolicyDefaultactionAcceptBgpAspathPrepend struct {
	Asnumber *string `json:"as-number,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	Repeatn *uint8 `json:"repeat-n,omitempty"`
}

// RoutingpolicyPolicyDefaultactionAcceptBgpAspathReplace struct
type RoutingpolicyPolicyDefaultactionAcceptBgpAspathReplace struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// RoutingpolicyPolicyDefaultactionAcceptBgpCommunities struct
type RoutingpolicyPolicyDefaultactionAcceptBgpCommunities struct {
	Add     *string `json:"add,omitempty"`
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
}

// RoutingpolicyPolicyDefaultactionAcceptBgpLocalpreference struct
type RoutingpolicyPolicyDefaultactionAcceptBgpLocalpreference struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// RoutingpolicyPolicyDefaultactionAcceptBgpOrigin struct
type RoutingpolicyPolicyDefaultactionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set E_RoutingpolicyPolicyDefaultactionAcceptBgpOriginSet `json:"set,omitempty"`
	//Set *string `json:"set,omitempty"`
}

// RoutingpolicyPolicyDefaultactionNextentry struct
type RoutingpolicyPolicyDefaultactionNextentry struct {
}

// RoutingpolicyPolicyDefaultactionNextpolicy struct
type RoutingpolicyPolicyDefaultactionNextpolicy struct {
}

// RoutingpolicyPolicyDefaultactionReject struct
type RoutingpolicyPolicyDefaultactionReject struct {
}

// RoutingpolicyPolicyStatement struct
type RoutingpolicyPolicyStatement struct {
	Action *RoutingpolicyPolicyStatementAction `json:"action,omitempty"`
	Match  *RoutingpolicyPolicyStatementMatch  `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Sequenceid *uint32 `json:"sequence-id"`
}

// RoutingpolicyPolicyStatementAction struct
type RoutingpolicyPolicyStatementAction struct {
	Accept     *RoutingpolicyPolicyStatementActionAccept     `json:"accept,omitempty"`
	Nextentry  *RoutingpolicyPolicyStatementActionNextentry  `json:"next-entry,omitempty"`
	Nextpolicy *RoutingpolicyPolicyStatementActionNextpolicy `json:"next-policy,omitempty"`
	Reject     *RoutingpolicyPolicyStatementActionReject     `json:"reject,omitempty"`
}

// RoutingpolicyPolicyStatementActionAccept struct
type RoutingpolicyPolicyStatementActionAccept struct {
	Bgp *RoutingpolicyPolicyStatementActionAcceptBgp `json:"bgp,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgp struct
type RoutingpolicyPolicyStatementActionAcceptBgp struct {
	Aspath          *RoutingpolicyPolicyStatementActionAcceptBgpAspath          `json:"as-path,omitempty"`
	Communities     *RoutingpolicyPolicyStatementActionAcceptBgpCommunities     `json:"communities,omitempty"`
	Localpreference *RoutingpolicyPolicyStatementActionAcceptBgpLocalpreference `json:"local-preference,omitempty"`
	Origin          *RoutingpolicyPolicyStatementActionAcceptBgpOrigin          `json:"origin,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpAspath struct
type RoutingpolicyPolicyStatementActionAcceptBgpAspath struct {
	Prepend *RoutingpolicyPolicyStatementActionAcceptBgpAspathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                     `json:"remove,omitempty"`
	Replace *RoutingpolicyPolicyStatementActionAcceptBgpAspathReplace `json:"replace,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpAspathPrepend struct
type RoutingpolicyPolicyStatementActionAcceptBgpAspathPrepend struct {
	Asnumber *string `json:"as-number,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	Repeatn *uint8 `json:"repeat-n,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpAspathReplace struct
type RoutingpolicyPolicyStatementActionAcceptBgpAspathReplace struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpCommunities struct
type RoutingpolicyPolicyStatementActionAcceptBgpCommunities struct {
	Add     *string `json:"add,omitempty"`
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpLocalpreference struct
type RoutingpolicyPolicyStatementActionAcceptBgpLocalpreference struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpOrigin struct
type RoutingpolicyPolicyStatementActionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set E_RoutingpolicyPolicyStatementActionAcceptBgpOriginSet `json:"set,omitempty"`
	//Set *string `json:"set,omitempty"`
}

// RoutingpolicyPolicyStatementActionNextentry struct
type RoutingpolicyPolicyStatementActionNextentry struct {
}

// RoutingpolicyPolicyStatementActionNextpolicy struct
type RoutingpolicyPolicyStatementActionNextpolicy struct {
}

// RoutingpolicyPolicyStatementActionReject struct
type RoutingpolicyPolicyStatementActionReject struct {
}

// RoutingpolicyPolicyStatementMatch struct
type RoutingpolicyPolicyStatementMatch struct {
	Bgp       *RoutingpolicyPolicyStatementMatchBgp  `json:"bgp,omitempty"`
	Family    *string                                `json:"family,omitempty"`
	Isis      *RoutingpolicyPolicyStatementMatchIsis `json:"isis,omitempty"`
	Ospf      *RoutingpolicyPolicyStatementMatchOspf `json:"ospf,omitempty"`
	Prefixset *string                                `json:"prefix-set,omitempty"`
	Protocol  *string                                `json:"protocol,omitempty"`
}

// RoutingpolicyPolicyStatementMatchBgp struct
type RoutingpolicyPolicyStatementMatchBgp struct {
	Aspathlength *RoutingpolicyPolicyStatementMatchBgpAspathlength `json:"as-path-length,omitempty"`
	Aspathset    *string                                           `json:"as-path-set,omitempty"`
	Communityset *string                                           `json:"community-set,omitempty"`
	Evpn         *RoutingpolicyPolicyStatementMatchBgpEvpn         `json:"evpn,omitempty"`
}

// RoutingpolicyPolicyStatementMatchBgpAspathlength struct
type RoutingpolicyPolicyStatementMatchBgpAspathlength struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	// +kubebuilder:default:="eq"
	Operator E_RoutingpolicyPolicyStatementMatchBgpAspathlengthOperator `json:"operator,omitempty"`
	//Operator *string `json:"operator,omitempty"`
	// +kubebuilder:default:=false
	Unique *bool `json:"unique,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Value *uint8 `json:"value"`
}

// RoutingpolicyPolicyStatementMatchBgpEvpn struct
type RoutingpolicyPolicyStatementMatchBgpEvpn struct {
	Routetype *RoutingpolicyPolicyStatementMatchBgpEvpnRoutetype `json:"route-type,omitempty"`
}

// RoutingpolicyPolicyStatementMatchBgpEvpnRoutetype struct
type RoutingpolicyPolicyStatementMatchBgpEvpnRoutetype struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=5
	Routetype *uint8 `json:"route-type,omitempty"`
}

// RoutingpolicyPolicyStatementMatchIsis struct
type RoutingpolicyPolicyStatementMatchIsis struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Level *uint8 `json:"level,omitempty"`
	// +kubebuilder:validation:Enum=`external`;`internal`
	Routetype E_RoutingpolicyPolicyStatementMatchIsisRoutetype `json:"route-type,omitempty"`
	//Routetype *string `json:"route-type,omitempty"`
}

// RoutingpolicyPolicyStatementMatchOspf struct
type RoutingpolicyPolicyStatementMatchOspf struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|[0-9\.]*|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])([\p{N}\p{L}]+)?`
	Areaid *string `json:"area-id,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Instanceid *uint32 `json:"instance-id,omitempty"`
	Routetype  *string `json:"route-type,omitempty"`
}

// A RoutingpolicyPolicySpec defines the desired state of a RoutingpolicyPolicy.
type RoutingpolicyPolicySpec struct {
	nddv1.ResourceSpec  `json:",inline"`
	RoutingpolicyPolicy *RoutingpolicyPolicy `json:"policy,omitempty"`
}

// A RoutingpolicyPolicyStatus represents the observed state of a RoutingpolicyPolicy.
type RoutingpolicyPolicyStatus struct {
	nddv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyPolicy is the Schema for the RoutingpolicyPolicy API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='LeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl2}
type SrlRoutingpolicyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoutingpolicyPolicySpec   `json:"spec,omitempty"`
	Status RoutingpolicyPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyPolicyList contains a list of RoutingpolicyPolicys
type SrlRoutingpolicyPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlRoutingpolicyPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlRoutingpolicyPolicy{}, &SrlRoutingpolicyPolicyList{})
}

// RoutingpolicyPolicy type metadata.
var (
	RoutingpolicyPolicyKindKind         = reflect.TypeOf(SrlRoutingpolicyPolicy{}).Name()
	RoutingpolicyPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: RoutingpolicyPolicyKindKind}.String()
	RoutingpolicyPolicyKindAPIVersion   = RoutingpolicyPolicyKindKind + "." + GroupVersion.String()
	RoutingpolicyPolicyGroupVersionKind = GroupVersion.WithKind(RoutingpolicyPolicyKindKind)
)
