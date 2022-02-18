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
	errCreateSystemNetworkinstanceProtocolsEvpn = "cannot create SystemNetworkinstanceProtocolsEvpn"
	errDeleteSystemNetworkinstanceProtocolsEvpn = "cannot delete SystemNetworkinstanceProtocolsEvpn"
	errGetSystemNetworkinstanceProtocolsEvpn    = "cannot get SystemNetworkinstanceProtocolsEvpn"
)

type SystemNetworkinstanceProtocolsEvpn interface {
	// methods children
	NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstance(c resource.ClientApplicator, key string) SystemNetworkinstanceProtocolsEvpnEsisBgpinstance
	GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() map[string]SystemNetworkinstanceProtocolsEvpnEsisBgpinstance
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.SystemNetworkinstanceProtocolsEvpn
	Update(x *srlv1alpha1.SystemNetworkinstanceProtocolsEvpn)
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

func NewSystemNetworkinstanceProtocolsEvpn(c resource.ClientApplicator, p Device, key string) SystemNetworkinstanceProtocolsEvpn {
	newSystemNetworkinstanceProtocolsEvpnList := func() srlv1alpha1.IFSrlSystemNetworkinstanceProtocolsEvpnList {
		return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnList{}
	}
	return &systemnetworkinstanceprotocolsevpn{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		SystemNetworkinstanceProtocolsEvpnEsisBgpinstance: make(map[string]SystemNetworkinstanceProtocolsEvpnEsisBgpinstance),
		// data key
		//SystemNetworkinstanceProtocolsEvpn: &srlv1alpha1.SystemNetworkinstanceProtocolsEvpn{
		//	Name: &name,
		//},
		newSystemNetworkinstanceProtocolsEvpnList: newSystemNetworkinstanceProtocolsEvpnList,
	}
}

type systemnetworkinstanceprotocolsevpn struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstance map[string]SystemNetworkinstanceProtocolsEvpnEsisBgpinstance
	// Data
	SystemNetworkinstanceProtocolsEvpn        *srlv1alpha1.SystemNetworkinstanceProtocolsEvpn
	newSystemNetworkinstanceProtocolsEvpnList func() srlv1alpha1.IFSrlSystemNetworkinstanceProtocolsEvpnList
}

// key type/method

type SystemNetworkinstanceProtocolsEvpnKey struct {
}

func WithSystemNetworkinstanceProtocolsEvpnKey(key *SystemNetworkinstanceProtocolsEvpnKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
func (x *systemnetworkinstanceprotocolsevpn) NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstance(c resource.ClientApplicator, key string) SystemNetworkinstanceProtocolsEvpnEsisBgpinstance {
	if _, ok := x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance[key]; !ok {
		x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance[key] = NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstance(c, x, key)
	}
	return x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance[key]
}
func (x *systemnetworkinstanceprotocolsevpn) GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() map[string]SystemNetworkinstanceProtocolsEvpnEsisBgpinstance {
	return x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance
}

// Data methods
func (x *systemnetworkinstanceprotocolsevpn) Update(d *srlv1alpha1.SystemNetworkinstanceProtocolsEvpn) {
	x.SystemNetworkinstanceProtocolsEvpn = d
}

// methods data
func (x *systemnetworkinstanceprotocolsevpn) Get() *srlv1alpha1.SystemNetworkinstanceProtocolsEvpn {
	return x.SystemNetworkinstanceProtocolsEvpn
}

func (x *systemnetworkinstanceprotocolsevpn) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *systemnetworkinstanceprotocolsevpn) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.SystemNetworkinstanceProtocolsEvpn)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s SystemNetworkinstanceProtocolsEvpn: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s SystemNetworkinstanceProtocolsEvpn: %s\n", strings.Repeat(" ", n), key)
	}

	n++
	for key, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() {
		i.Print(key, n)
	}
}

func (x *systemnetworkinstanceprotocolsevpn) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateSystemNetworkinstanceProtocolsEvpn)
		}
	}
	for _, r := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}

	return nil
}

func (x *systemnetworkinstanceprotocolsevpn) DestroySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Delete(ctx, o); err != nil {
			return errors.Wrap(err, errCreateSystemNetworkinstanceProtocolsEvpn)
		}
	}
	for _, r := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() {
		if err := r.DestroySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}

	return nil
}
func (x *systemnetworkinstanceprotocolsevpn) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpn {

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
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

	return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpn{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.SystemNetworkinstanceProtocolsEvpnSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			SystemNetworkinstanceProtocolsEvpn: x.SystemNetworkinstanceProtocolsEvpn,
		},
	}
}

func (x *systemnetworkinstanceprotocolsevpn) InitializeDummySchema() {
	c0 := x.NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstance(x.client, "dummy")
	c0.InitializeDummySchema()
}

func (x *systemnetworkinstanceprotocolsevpn) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newSystemNetworkinstanceProtocolsEvpnList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemNetworkinstanceProtocolsEvpns() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *systemnetworkinstanceprotocolsevpn) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.SystemNetworkinstanceProtocolsEvpnKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *systemnetworkinstanceprotocolsevpn) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.SystemNetworkinstanceProtocolsEvpnKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpn{
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
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}

	return nil
}

func (x *systemnetworkinstanceprotocolsevpn) ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error {
	// options list all resources belonging to the transaction based on transaction owner and generation
	opts := []client.ListOption{
		client.InNamespace(cr.GetNamespace()),
		client.MatchingLabels(map[string]string{
			srlv1alpha1.LabelNddaOwner: cr.GetName(),
			//srlv1alpha1.LabelNddaOwnerGeneration: cr.GetOwnerGeneration(),
		}),
	}
	list := x.newSystemNetworkinstanceProtocolsEvpnList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemNetworkinstanceProtocolsEvpns() {
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
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() {
		if err := i.ListResourcesByTransaction(ctx, cr, resources); err != nil {
			return err
		}
	}
	return nil
}
