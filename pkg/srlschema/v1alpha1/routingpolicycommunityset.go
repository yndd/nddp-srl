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
	errCreateRoutingpolicyCommunityset = "cannot create RoutingpolicyCommunityset"
	errDeleteRoutingpolicyCommunityset = "cannot delete RoutingpolicyCommunityset"
	errGetRoutingpolicyCommunityset    = "cannot get RoutingpolicyCommunityset"
)

type RoutingpolicyCommunityset interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.RoutingpolicyCommunityset
	Update(x *srlv1alpha1.RoutingpolicyCommunityset)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewRoutingpolicyCommunityset(c resource.ClientApplicator, p Device, key string) RoutingpolicyCommunityset {
	newRoutingpolicyCommunitysetList := func() srlv1alpha1.IFSrlRoutingpolicyCommunitysetList {
		return &srlv1alpha1.SrlRoutingpolicyCommunitysetList{}
	}
	return &routingpolicycommunityset{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//RoutingpolicyCommunityset: &srlv1alpha1.RoutingpolicyCommunityset{
		//	Name: &name,
		//},
		newRoutingpolicyCommunitysetList: newRoutingpolicyCommunitysetList,
	}
}

type routingpolicycommunityset struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	// Data
	RoutingpolicyCommunityset        *srlv1alpha1.RoutingpolicyCommunityset
	newRoutingpolicyCommunitysetList func() srlv1alpha1.IFSrlRoutingpolicyCommunitysetList
}

// key type/method

type RoutingpolicyCommunitysetKey struct {
	Name string
}

func WithRoutingpolicyCommunitysetKey(key *RoutingpolicyCommunitysetKey) string {
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
func (x *routingpolicycommunityset) Update(d *srlv1alpha1.RoutingpolicyCommunityset) {
	x.RoutingpolicyCommunityset = d
}

// methods data
func (x *routingpolicycommunityset) Get() *srlv1alpha1.RoutingpolicyCommunityset {
	return x.RoutingpolicyCommunityset
}

func (x *routingpolicycommunityset) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *routingpolicycommunityset) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.RoutingpolicyCommunityset)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s RoutingpolicyCommunityset: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s RoutingpolicyCommunityset: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *routingpolicycommunityset) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateRoutingpolicyCommunityset)
		}
	}

	return nil
}
func (x *routingpolicycommunityset) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlRoutingpolicyCommunityset {
	key0 := strings.ReplaceAll(*x.RoutingpolicyCommunityset.Name, "/", "-")

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

	return &srlv1alpha1.SrlRoutingpolicyCommunityset{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.RoutingpolicyCommunitysetSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			RoutingpolicyCommunityset: x.RoutingpolicyCommunityset,
		},
	}
}

func (x *routingpolicycommunityset) InitializeDummySchema() {
}

func (x *routingpolicycommunityset) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newRoutingpolicyCommunitysetList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetRoutingpolicyCommunitysets() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *routingpolicycommunityset) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		key0 := strings.ReplaceAll(*x.RoutingpolicyCommunityset.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.RoutingpolicyCommunitysetKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *routingpolicycommunityset) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.RoutingpolicyCommunitysetKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlRoutingpolicyCommunityset{
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
