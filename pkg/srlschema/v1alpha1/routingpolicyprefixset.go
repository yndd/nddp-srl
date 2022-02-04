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
	errCreateRoutingpolicyPrefixset = "cannot create RoutingpolicyPrefixset"
	errDeleteRoutingpolicyPrefixset = "cannot delete RoutingpolicyPrefixset"
	errGetRoutingpolicyPrefixset    = "cannot get RoutingpolicyPrefixset"
)

type RoutingpolicyPrefixset interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.RoutingpolicyPrefixset
	Update(x *srlv1alpha1.RoutingpolicyPrefixset)
	AddRoutingpolicyPrefixsetPrefix(ai *srlv1alpha1.RoutingpolicyPrefixsetPrefix)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewRoutingpolicyPrefixset(c resource.ClientApplicator, p Device, key string) RoutingpolicyPrefixset {
	newRoutingpolicyPrefixsetList := func() srlv1alpha1.IFSrlRoutingpolicyPrefixsetList {
		return &srlv1alpha1.SrlRoutingpolicyPrefixsetList{}
	}
	return &routingpolicyprefixset{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//RoutingpolicyPrefixset: &srlv1alpha1.RoutingpolicyPrefixset{
		//	Name: &name,
		//},
		newRoutingpolicyPrefixsetList: newRoutingpolicyPrefixsetList,
	}
}

type routingpolicyprefixset struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	// Data
	RoutingpolicyPrefixset        *srlv1alpha1.RoutingpolicyPrefixset
	newRoutingpolicyPrefixsetList func() srlv1alpha1.IFSrlRoutingpolicyPrefixsetList
}

// key type/method

type RoutingpolicyPrefixsetKey struct {
	Name string
}

func WithRoutingpolicyPrefixsetKey(key *RoutingpolicyPrefixsetKey) string {
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
func (x *routingpolicyprefixset) Update(d *srlv1alpha1.RoutingpolicyPrefixset) {
	x.RoutingpolicyPrefixset = d
}

// methods data
func (x *routingpolicyprefixset) Get() *srlv1alpha1.RoutingpolicyPrefixset {
	return x.RoutingpolicyPrefixset
}

func (x *routingpolicyprefixset) GetKey() []string {
	return strings.Split(x.key, ".")
}

// RoutingpolicyPrefixset prefix prefixset Prefixset []
func (x *routingpolicyprefixset) AddRoutingpolicyPrefixsetPrefix(ai *srlv1alpha1.RoutingpolicyPrefixsetPrefix) {
	//x.RoutingpolicyPrefixset.Prefix = append(x.RoutingpolicyPrefixset.Prefix, ai)
	if len(x.RoutingpolicyPrefixset.Prefix) == 0 {
		x.RoutingpolicyPrefixset.Prefix = make([]*srlv1alpha1.RoutingpolicyPrefixsetPrefix, 0)
	}
	found := false
	for _, xx := range x.RoutingpolicyPrefixset.Prefix {

		// [ip-prefix mask-length-range]
		if *xx.Ipprefix == *ai.Ipprefix &&
			*xx.Masklengthrange == *ai.Masklengthrange {
			found = true
			xx = ai
		}
	}
	if !found {
		x.RoutingpolicyPrefixset.Prefix = append(x.RoutingpolicyPrefixset.Prefix, ai)
	}
}

// methods schema

func (x *routingpolicyprefixset) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.RoutingpolicyPrefixset)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s RoutingpolicyPrefixset: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s RoutingpolicyPrefixset: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *routingpolicyprefixset) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateRoutingpolicyPrefixset)
		}
	}

	return nil
}
func (x *routingpolicyprefixset) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlRoutingpolicyPrefixset {
	key0 := strings.ReplaceAll(*x.RoutingpolicyPrefixset.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(key0),
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName

	namespace := mg.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}

	return &srlv1alpha1.SrlRoutingpolicyPrefixset{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.RoutingpolicyPrefixsetSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			RoutingpolicyPrefixset: x.RoutingpolicyPrefixset,
		},
	}
}

func (x *routingpolicyprefixset) InitializeDummySchema() {
}

func (x *routingpolicyprefixset) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newRoutingpolicyPrefixsetList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetRoutingpolicyPrefixsets() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *routingpolicyprefixset) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		key0 := strings.ReplaceAll(*x.RoutingpolicyPrefixset.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.RoutingpolicyPrefixsetKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *routingpolicyprefixset) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.RoutingpolicyPrefixsetKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlRoutingpolicyPrefixset{
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
