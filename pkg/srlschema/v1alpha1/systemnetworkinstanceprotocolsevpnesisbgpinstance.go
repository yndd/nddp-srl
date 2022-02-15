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
	errCreateSystemNetworkinstanceProtocolsEvpnEsisBgpinstance = "cannot create SystemNetworkinstanceProtocolsEvpnEsisBgpinstance"
	errDeleteSystemNetworkinstanceProtocolsEvpnEsisBgpinstance = "cannot delete SystemNetworkinstanceProtocolsEvpnEsisBgpinstance"
	errGetSystemNetworkinstanceProtocolsEvpnEsisBgpinstance    = "cannot get SystemNetworkinstanceProtocolsEvpnEsisBgpinstance"
)

type SystemNetworkinstanceProtocolsEvpnEsisBgpinstance interface {
	// methods children
	NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi(c resource.ClientApplicator, key string) SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi
	GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() map[string]SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance
	Update(x *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error
}

func NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstance(c resource.ClientApplicator, p SystemNetworkinstanceProtocolsEvpn, key string) SystemNetworkinstanceProtocolsEvpnEsisBgpinstance {
	newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList := func() srlv1alpha1.IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList {
		return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList{}
	}
	return &systemnetworkinstanceprotocolsevpnesisbgpinstance{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi: make(map[string]SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi),
		// data key
		//SystemNetworkinstanceProtocolsEvpnEsisBgpinstance: &srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance{
		//	Name: &name,
		//},
		newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList: newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList,
	}
}

type systemnetworkinstanceprotocolsevpnesisbgpinstance struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent SystemNetworkinstanceProtocolsEvpn
	// children
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi map[string]SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi
	// Data
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstance        *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance
	newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList func() srlv1alpha1.IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList
}

// key type/method

type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKey struct {
	Id string
}

func WithSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKey(key *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi(c resource.ClientApplicator, key string) SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi {
	if _, ok := x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi[key]; !ok {
		x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi[key] = NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi(c, x, key)
	}
	return x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi[key]
}
func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() map[string]SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi {
	return x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi
}

// Data methods
func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) Update(d *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance) {
	x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance = d
}

// methods data
func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) Get() *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance {
	return x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s SystemNetworkinstanceProtocolsEvpnEsisBgpinstance: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s SystemNetworkinstanceProtocolsEvpnEsisBgpinstance: %s\n", strings.Repeat(" ", n), key)
	}

	n++
	for key, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() {
		i.Print(key, n)
	}
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateSystemNetworkinstanceProtocolsEvpnEsisBgpinstance)
		}
	}
	for _, r := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}

	return nil
}
func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance {
	key0 := strings.ReplaceAll(*x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance.Id, "/", "-")

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

	return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			SystemNetworkinstanceProtocolsEvpnEsisBgpinstance: x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance,
		},
	}
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) InitializeDummySchema() {
	c0 := x.NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi(x.client, "dummy")
	c0.InitializeDummySchema()
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		key0 := strings.ReplaceAll(*x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstance.Id, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{
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
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}

	return nil
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstance) ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error {
	// options list all resources belonging to the transaction based on transaction owner and generation
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: cr.GetName()},
		client.MatchingLabels{srlv1alpha1.LabelNddaOwnerGeneration: cr.GetOwnerGeneration()},
	}
	list := x.newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstances() {
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
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() {
		if err := i.ListResourcesByTransaction(ctx, cr, resources); err != nil {
			return err
		}
	}
	return nil
}
