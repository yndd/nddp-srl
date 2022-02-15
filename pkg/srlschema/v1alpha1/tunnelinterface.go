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
	errCreateTunnelinterface = "cannot create Tunnelinterface"
	errDeleteTunnelinterface = "cannot delete Tunnelinterface"
	errGetTunnelinterface    = "cannot get Tunnelinterface"
)

type Tunnelinterface interface {
	// methods children
	NewTunnelinterfaceVxlaninterface(c resource.ClientApplicator, key string) TunnelinterfaceVxlaninterface
	GetTunnelinterfaceVxlaninterfaces() map[string]TunnelinterfaceVxlaninterface
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.Tunnelinterface
	Update(x *srlv1alpha1.Tunnelinterface)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error
}

func NewTunnelinterface(c resource.ClientApplicator, p Device, key string) Tunnelinterface {
	newTunnelinterfaceList := func() srlv1alpha1.IFSrlTunnelinterfaceList { return &srlv1alpha1.SrlTunnelinterfaceList{} }
	return &tunnelinterface{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		TunnelinterfaceVxlaninterface: make(map[string]TunnelinterfaceVxlaninterface),
		// data key
		//Tunnelinterface: &srlv1alpha1.Tunnelinterface{
		//	Name: &name,
		//},
		newTunnelinterfaceList: newTunnelinterfaceList,
	}
}

type tunnelinterface struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	TunnelinterfaceVxlaninterface map[string]TunnelinterfaceVxlaninterface
	// Data
	Tunnelinterface        *srlv1alpha1.Tunnelinterface
	newTunnelinterfaceList func() srlv1alpha1.IFSrlTunnelinterfaceList
}

// key type/method

type TunnelinterfaceKey struct {
	Name string
}

func WithTunnelinterfaceKey(key *TunnelinterfaceKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
func (x *tunnelinterface) NewTunnelinterfaceVxlaninterface(c resource.ClientApplicator, key string) TunnelinterfaceVxlaninterface {
	if _, ok := x.TunnelinterfaceVxlaninterface[key]; !ok {
		x.TunnelinterfaceVxlaninterface[key] = NewTunnelinterfaceVxlaninterface(c, x, key)
	}
	return x.TunnelinterfaceVxlaninterface[key]
}
func (x *tunnelinterface) GetTunnelinterfaceVxlaninterfaces() map[string]TunnelinterfaceVxlaninterface {
	return x.TunnelinterfaceVxlaninterface
}

// Data methods
func (x *tunnelinterface) Update(d *srlv1alpha1.Tunnelinterface) {
	x.Tunnelinterface = d
}

// methods data
func (x *tunnelinterface) Get() *srlv1alpha1.Tunnelinterface {
	return x.Tunnelinterface
}

func (x *tunnelinterface) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *tunnelinterface) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.Tunnelinterface)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s Tunnelinterface: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s Tunnelinterface: %s\n", strings.Repeat(" ", n), key)
	}

	n++
	for key, i := range x.GetTunnelinterfaceVxlaninterfaces() {
		i.Print(key, n)
	}
}

func (x *tunnelinterface) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateTunnelinterface)
		}
	}
	for _, r := range x.GetTunnelinterfaceVxlaninterfaces() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}

	return nil
}
func (x *tunnelinterface) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlTunnelinterface {
	key0 := strings.ReplaceAll(*x.Tunnelinterface.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(key0),
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

	return &srlv1alpha1.SrlTunnelinterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.TunnelinterfaceSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			Tunnelinterface: x.Tunnelinterface,
		},
	}
}

func (x *tunnelinterface) InitializeDummySchema() {
	c0 := x.NewTunnelinterfaceVxlaninterface(x.client, "dummy")
	c0.InitializeDummySchema()
}

func (x *tunnelinterface) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newTunnelinterfaceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetTunnelinterfaces() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	for _, i := range x.GetTunnelinterfaceVxlaninterfaces() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *tunnelinterface) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		key0 := strings.ReplaceAll(*x.Tunnelinterface.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.TunnelinterfaceKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	for _, i := range x.GetTunnelinterfaceVxlaninterfaces() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *tunnelinterface) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.TunnelinterfaceKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlTunnelinterface{
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
	for _, i := range x.GetTunnelinterfaceVxlaninterfaces() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}

	return nil
}

func (x *tunnelinterface) ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error {
	// options list all resources belonging to the transaction based on transaction owner and generation
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: cr.GetName()},
		client.MatchingLabels{srlv1alpha1.LabelNddaOwnerGeneration: cr.GetOwnerGeneration()},
	}
	list := x.newTunnelinterfaceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetTunnelinterfaces() {
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
	for _, i := range x.GetTunnelinterfaceVxlaninterfaces() {
		if err := i.ListResourcesByTransaction(ctx, cr, resources); err != nil {
			return err
		}
	}
	return nil
}
