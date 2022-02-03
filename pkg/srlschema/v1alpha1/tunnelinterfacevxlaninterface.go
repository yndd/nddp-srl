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
	errCreateTunnelinterfaceVxlaninterface = "cannot create TunnelinterfaceVxlaninterface"
	errDeleteTunnelinterfaceVxlaninterface = "cannot delete TunnelinterfaceVxlaninterface"
	errGetTunnelinterfaceVxlaninterface    = "cannot get TunnelinterfaceVxlaninterface"
)

type TunnelinterfaceVxlaninterface interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.TunnelinterfaceVxlaninterface
	Update(x *srlv1alpha1.TunnelinterfaceVxlaninterface)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewTunnelinterfaceVxlaninterface(c resource.ClientApplicator, p Tunnelinterface, key string) TunnelinterfaceVxlaninterface {
	newTunnelinterfaceVxlaninterfaceList := func() srlv1alpha1.IFSrlTunnelinterfaceVxlaninterfaceList {
		return &srlv1alpha1.SrlTunnelinterfaceVxlaninterfaceList{}
	}
	return &tunnelinterfacevxlaninterface{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//TunnelinterfaceVxlaninterface: &srlv1alpha1.TunnelinterfaceVxlaninterface{
		//	Name: &name,
		//},
		newTunnelinterfaceVxlaninterfaceList: newTunnelinterfaceVxlaninterfaceList,
	}
}

type tunnelinterfacevxlaninterface struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Tunnelinterface
	// children
	// Data
	TunnelinterfaceVxlaninterface        *srlv1alpha1.TunnelinterfaceVxlaninterface
	newTunnelinterfaceVxlaninterfaceList func() srlv1alpha1.IFSrlTunnelinterfaceVxlaninterfaceList
}

// key type/method

type TunnelinterfaceVxlaninterfaceKey struct {
	Index string
}

func WithTunnelinterfaceVxlaninterfaceKey(key *TunnelinterfaceVxlaninterfaceKey) string {
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
func (x *tunnelinterfacevxlaninterface) Update(d *srlv1alpha1.TunnelinterfaceVxlaninterface) {
	x.TunnelinterfaceVxlaninterface = d
}

// methods data
func (x *tunnelinterfacevxlaninterface) Get() *srlv1alpha1.TunnelinterfaceVxlaninterface {
	return x.TunnelinterfaceVxlaninterface
}

func (x *tunnelinterfacevxlaninterface) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *tunnelinterfacevxlaninterface) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.TunnelinterfaceVxlaninterface)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s TunnelinterfaceVxlaninterface: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s TunnelinterfaceVxlaninterface: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *tunnelinterfacevxlaninterface) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateTunnelinterfaceVxlaninterface)
		}
	}

	return nil
}
func (x *tunnelinterfacevxlaninterface) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlTunnelinterfaceVxlaninterface {
	parent0Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
	//1
	key0 := strconv.Itoa(int(*x.TunnelinterfaceVxlaninterface.Index))

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			// 0
			strings.ToLower(parent0Key0),
			//1
			strings.ToLower(key0),
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName
	return &srlv1alpha1.SrlTunnelinterfaceVxlaninterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.TunnelinterfaceVxlaninterfaceSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			TunnelInterfaceName: &parent0Key0,
			//1
			TunnelinterfaceVxlaninterface: x.TunnelinterfaceVxlaninterface,
		},
	}
}

func (x *tunnelinterfacevxlaninterface) InitializeDummySchema() {
}

func (x *tunnelinterfacevxlaninterface) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newTunnelinterfaceVxlaninterfaceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetTunnelinterfaceVxlaninterfaces() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *tunnelinterfacevxlaninterface) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		parent0Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
		//1
		key0 := strconv.Itoa(int(*x.TunnelinterfaceVxlaninterface.Index))

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(parent0Key0),
				//1
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.TunnelinterfaceVxlaninterfaceKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *tunnelinterfacevxlaninterface) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.TunnelinterfaceVxlaninterfaceKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlTunnelinterfaceVxlaninterface{
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
