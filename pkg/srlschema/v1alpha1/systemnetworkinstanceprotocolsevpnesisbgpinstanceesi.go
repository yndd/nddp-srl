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
	errCreateSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi = "cannot create SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi"
	errDeleteSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi = "cannot delete SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi"
	errGetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi    = "cannot get SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi"
)

type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi
	Update(x *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error
}

func NewSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi(c resource.ClientApplicator, p SystemNetworkinstanceProtocolsEvpnEsisBgpinstance, key string) SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi {
	newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList := func() srlv1alpha1.IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList {
		return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList{}
	}
	return &systemnetworkinstanceprotocolsevpnesisbgpinstanceesi{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi: &srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{
		//	Name: &name,
		//},
		newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList: newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList,
	}
}

type systemnetworkinstanceprotocolsevpnesisbgpinstanceesi struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent SystemNetworkinstanceProtocolsEvpnEsisBgpinstance
	// children
	// Data
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi        *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi
	newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList func() srlv1alpha1.IFSrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList
}

// key type/method

type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKey struct {
	Name string
}

func WithSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKey(key *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKey) string {
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
func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) Update(d *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) {
	x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi = d
}

// methods data
func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) Get() *srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi {
	return x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi)
		}
	}

	return nil
}
func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi {
	parent5Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
	//1
	key0 := strings.ReplaceAll(*x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			// 0
			strings.ToLower(parent5Key0),
			//1
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

	return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			BgpInstanceId: &x.parent.GetKey()[0],
			//1
			SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi: x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi,
		},
	}
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) InitializeDummySchema() {
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		parent5Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
		//1
		key0 := strings.ReplaceAll(*x.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(parent5Key0),
				//1
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{
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

func (x *systemnetworkinstanceprotocolsevpnesisbgpinstanceesi) ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error {
	// options list all resources belonging to the transaction based on transaction owner and generation
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: cr.GetName()},
		client.MatchingLabels{srlv1alpha1.LabelNddaOwnerGeneration: cr.GetOwnerGeneration()},
	}
	list := x.newSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis() {
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
