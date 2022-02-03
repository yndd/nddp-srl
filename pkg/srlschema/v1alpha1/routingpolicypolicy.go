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

package srlschema

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/meta"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
)

const (
	errCreateRoutingpolicyPolicy = "cannot create RoutingpolicyPolicy"
	errDeleteRoutingpolicyPolicy = "cannot delete RoutingpolicyPolicy"
	errGetRoutingpolicyPolicy    = "cannot get RoutingpolicyPolicy"
)

type RoutingpolicyPolicy interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.RoutingpolicyPolicy
	Update(x *srlv1alpha1.RoutingpolicyPolicy)
	AddRoutingpolicyPolicyStatement(ai *srlv1alpha1.RoutingpolicyPolicyStatement)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewRoutingpolicyPolicy(c resource.ClientApplicator, p Device, key string) RoutingpolicyPolicy {
	newRoutingpolicyPolicyList := func() srlv1alpha1.IFSrlRoutingpolicyPolicyList { return &srlv1alpha1.SrlRoutingpolicyPolicyList{} }
	return &routingpolicypolicy{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//RoutingpolicyPolicy: &srlv1alpha1.RoutingpolicyPolicy{
		//	Name: &name,
		//},
		newRoutingpolicyPolicyList: newRoutingpolicyPolicyList,
	}
}

type routingpolicypolicy struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	// Data
	RoutingpolicyPolicy        *srlv1alpha1.RoutingpolicyPolicy
	newRoutingpolicyPolicyList func() srlv1alpha1.IFSrlRoutingpolicyPolicyList
}

// key type/method

type RoutingpolicyPolicyKey struct {
	Name string
}

func WithRoutingpolicyPolicyKey(key *RoutingpolicyPolicyKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
// Data methods
func (x *routingpolicypolicy) Update(d *srlv1alpha1.RoutingpolicyPolicy) {
	x.RoutingpolicyPolicy = d
}

// methods data
func (x *routingpolicypolicy) Get() *srlv1alpha1.RoutingpolicyPolicy {
	return x.RoutingpolicyPolicy
}

func (x *routingpolicypolicy) GetKey() []string {
	return strings.Split(x.key, ".")
}

// RoutingpolicyPolicy statement policy Policy []
func (x *routingpolicypolicy) AddRoutingpolicyPolicyStatement(ai *srlv1alpha1.RoutingpolicyPolicyStatement) {
	x.RoutingpolicyPolicy.Statement = append(x.RoutingpolicyPolicy.Statement, ai)
}

// methods schema

func (x *routingpolicypolicy) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.RoutingpolicyPolicy)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s RoutingpolicyPolicy: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s RoutingpolicyPolicy: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *routingpolicypolicy) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateRoutingpolicyPolicy)
		}
	}

	return nil
}
func (x *routingpolicypolicy) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlRoutingpolicyPolicy {
	key0 := strings.ReplaceAll(*x.RoutingpolicyPolicy.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(key0),
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName
	return &srlv1alpha1.SrlRoutingpolicyPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.RoutingpolicyPolicySpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			RoutingpolicyPolicy: x.RoutingpolicyPolicy,
		},
	}
}

func (x *routingpolicypolicy) InitializeDummySchema() {
}

func (x *routingpolicypolicy) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newRoutingpolicyPolicyList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetRoutingpolicyPolicys() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *routingpolicypolicy) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		key0 := strings.ReplaceAll(*x.RoutingpolicyPolicy.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.RoutingpolicyPolicyKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *routingpolicypolicy) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.RoutingpolicyPolicyKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlRoutingpolicyPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resName,
					Namespace: mg.GetNamespace(),
				},
			}
			if err := x.client.Delete(ctx, o); err != nil {
				return err
			}
		}
	}

	// children

	return nil
}
