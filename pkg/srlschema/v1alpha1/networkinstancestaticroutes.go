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
	"strconv"
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
	errCreateNetworkinstanceStaticroutes = "cannot create NetworkinstanceStaticroutes"
	errDeleteNetworkinstanceStaticroutes = "cannot delete NetworkinstanceStaticroutes"
	errGetNetworkinstanceStaticroutes    = "cannot get NetworkinstanceStaticroutes"
)

type NetworkinstanceStaticroutes interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.NetworkinstanceStaticroutes
	Update(x *srlv1alpha1.NetworkinstanceStaticroutes)
	AddNetworkinstanceStaticroutesRoute(ai *srlv1alpha1.NetworkinstanceStaticroutesRoute)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	DestroySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error
}

func NewNetworkinstanceStaticroutes(c resource.ClientApplicator, p Networkinstance, key string) NetworkinstanceStaticroutes {
	newNetworkinstanceStaticroutesList := func() srlv1alpha1.IFSrlNetworkinstanceStaticroutesList {
		return &srlv1alpha1.SrlNetworkinstanceStaticroutesList{}
	}
	return &networkinstancestaticroutes{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//NetworkinstanceStaticroutes: &srlv1alpha1.NetworkinstanceStaticroutes{
		//	Name: &name,
		//},
		newNetworkinstanceStaticroutesList: newNetworkinstanceStaticroutesList,
	}
}

type networkinstancestaticroutes struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Networkinstance
	// children
	// Data
	NetworkinstanceStaticroutes        *srlv1alpha1.NetworkinstanceStaticroutes
	newNetworkinstanceStaticroutesList func() srlv1alpha1.IFSrlNetworkinstanceStaticroutesList
}

// key type/method

type NetworkinstanceStaticroutesKey struct {
}

func WithNetworkinstanceStaticroutesKey(key *NetworkinstanceStaticroutesKey) string {
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
func (x *networkinstancestaticroutes) Update(d *srlv1alpha1.NetworkinstanceStaticroutes) {
	x.NetworkinstanceStaticroutes = d
}

// methods data
func (x *networkinstancestaticroutes) Get() *srlv1alpha1.NetworkinstanceStaticroutes {
	return x.NetworkinstanceStaticroutes
}

func (x *networkinstancestaticroutes) GetKey() []string {
	return strings.Split(x.key, ".")
}

// NetworkinstanceStaticroutes route staticroutes Staticroutes []
func (x *networkinstancestaticroutes) AddNetworkinstanceStaticroutesRoute(ai *srlv1alpha1.NetworkinstanceStaticroutesRoute) {
	//x.NetworkinstanceStaticroutes.Route = append(x.NetworkinstanceStaticroutes.Route, ai)
	if len(x.NetworkinstanceStaticroutes.Route) == 0 {
		x.NetworkinstanceStaticroutes.Route = make([]*srlv1alpha1.NetworkinstanceStaticroutesRoute, 0)
	}
	found := false
	for _, xx := range x.NetworkinstanceStaticroutes.Route {

		// [prefix]
		if *xx.Prefix == *ai.Prefix {
			found = true
			xx = ai
		}
	}
	if !found {
		x.NetworkinstanceStaticroutes.Route = append(x.NetworkinstanceStaticroutes.Route, ai)
	}
}

// methods schema

func (x *networkinstancestaticroutes) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.NetworkinstanceStaticroutes)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s NetworkinstanceStaticroutes: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s NetworkinstanceStaticroutes: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *networkinstancestaticroutes) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateNetworkinstanceStaticroutes)
		}
	}

	return nil
}

func (x *networkinstancestaticroutes) DestroySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Delete(ctx, o); err != nil {
			return errors.Wrap(err, errCreateNetworkinstanceStaticroutes)
		}
	}

	return nil
}
func (x *networkinstancestaticroutes) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlNetworkinstanceStaticroutes {
	parent0Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
	//1

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			// 0
			strings.ToLower(parent0Key0),
			//1
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaOwnerGeneration] = strconv.Itoa(int(mg.GetGeneration()))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName

	namespace := mg.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}

	return &srlv1alpha1.SrlNetworkinstanceStaticroutes{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.NetworkinstanceStaticroutesSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			NetworkInstanceName: &x.parent.GetKey()[0],
			//1
			NetworkinstanceStaticroutes: x.NetworkinstanceStaticroutes,
		},
	}
}

func (x *networkinstancestaticroutes) InitializeDummySchema() {
}

func (x *networkinstancestaticroutes) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newNetworkinstanceStaticroutesList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetNetworkinstanceStaticroutess() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *networkinstancestaticroutes) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		parent0Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
		//1

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(parent0Key0),
				//1
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.NetworkinstanceStaticroutesKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *networkinstancestaticroutes) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.NetworkinstanceStaticroutesKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlNetworkinstanceStaticroutes{
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

func (x *networkinstancestaticroutes) ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error {
	// options list all resources belonging to the transaction based on transaction owner and generation
	opts := []client.ListOption{
		client.InNamespace(cr.GetNamespace()),
		client.MatchingLabels(map[string]string{
			srlv1alpha1.LabelNddaOwner: cr.GetName(),
			//srlv1alpha1.LabelNddaOwnerGeneration: cr.GetOwnerGeneration(),
		}),
	}
	list := x.newNetworkinstanceStaticroutesList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetNetworkinstanceStaticroutess() {
		deviceName := i.GetLabels()[srlv1alpha1.LabelNddaDevice]
		if _, ok := resources[deviceName]; !ok {
			resources[deviceName] = make(map[string]map[string]interface{})
		}

		if _, ok := resources[deviceName][i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[deviceName][i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[deviceName][i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"
	}

	// children
	return nil
}
