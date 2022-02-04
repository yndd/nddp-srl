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
	errCreateSystemNetworkinstanceProtocolsBgpvpn = "cannot create SystemNetworkinstanceProtocolsBgpvpn"
	errDeleteSystemNetworkinstanceProtocolsBgpvpn = "cannot delete SystemNetworkinstanceProtocolsBgpvpn"
	errGetSystemNetworkinstanceProtocolsBgpvpn    = "cannot get SystemNetworkinstanceProtocolsBgpvpn"
)

type SystemNetworkinstanceProtocolsBgpvpn interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpn
	Update(x *srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpn)
	AddSystemNetworkinstanceProtocolsBgpvpnBgpinstance(ai *srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpnBgpinstance)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewSystemNetworkinstanceProtocolsBgpvpn(c resource.ClientApplicator, p Device, key string) SystemNetworkinstanceProtocolsBgpvpn {
	newSystemNetworkinstanceProtocolsBgpvpnList := func() srlv1alpha1.IFSrlSystemNetworkinstanceProtocolsBgpvpnList {
		return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsBgpvpnList{}
	}
	return &systemnetworkinstanceprotocolsbgpvpn{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//SystemNetworkinstanceProtocolsBgpvpn: &srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpn{
		//	Name: &name,
		//},
		newSystemNetworkinstanceProtocolsBgpvpnList: newSystemNetworkinstanceProtocolsBgpvpnList,
	}
}

type systemnetworkinstanceprotocolsbgpvpn struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	// Data
	SystemNetworkinstanceProtocolsBgpvpn        *srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpn
	newSystemNetworkinstanceProtocolsBgpvpnList func() srlv1alpha1.IFSrlSystemNetworkinstanceProtocolsBgpvpnList
}

// key type/method

type SystemNetworkinstanceProtocolsBgpvpnKey struct {
}

func WithSystemNetworkinstanceProtocolsBgpvpnKey(key *SystemNetworkinstanceProtocolsBgpvpnKey) string {
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
func (x *systemnetworkinstanceprotocolsbgpvpn) Update(d *srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpn) {
	x.SystemNetworkinstanceProtocolsBgpvpn = d
}

// methods data
func (x *systemnetworkinstanceprotocolsbgpvpn) Get() *srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpn {
	return x.SystemNetworkinstanceProtocolsBgpvpn
}

func (x *systemnetworkinstanceprotocolsbgpvpn) GetKey() []string {
	return strings.Split(x.key, ".")
}

// SystemNetworkinstanceProtocolsBgpvpn bgp-instance bgpvpn Bgpvpn []
func (x *systemnetworkinstanceprotocolsbgpvpn) AddSystemNetworkinstanceProtocolsBgpvpnBgpinstance(ai *srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpnBgpinstance) {
	x.SystemNetworkinstanceProtocolsBgpvpn.Bgpinstance = append(x.SystemNetworkinstanceProtocolsBgpvpn.Bgpinstance, ai)
}

// methods schema

func (x *systemnetworkinstanceprotocolsbgpvpn) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.SystemNetworkinstanceProtocolsBgpvpn)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s SystemNetworkinstanceProtocolsBgpvpn: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s SystemNetworkinstanceProtocolsBgpvpn: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *systemnetworkinstanceprotocolsbgpvpn) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateSystemNetworkinstanceProtocolsBgpvpn)
		}
	}

	return nil
}
func (x *systemnetworkinstanceprotocolsbgpvpn) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlSystemNetworkinstanceProtocolsBgpvpn {

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName

	namespace := mg.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}

	return &srlv1alpha1.SrlSystemNetworkinstanceProtocolsBgpvpn{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpnSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			SystemNetworkinstanceProtocolsBgpvpn: x.SystemNetworkinstanceProtocolsBgpvpn,
		},
	}
}

func (x *systemnetworkinstanceprotocolsbgpvpn) InitializeDummySchema() {
}

func (x *systemnetworkinstanceprotocolsbgpvpn) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newSystemNetworkinstanceProtocolsBgpvpnList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemNetworkinstanceProtocolsBgpvpns() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *systemnetworkinstanceprotocolsbgpvpn) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpnKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *systemnetworkinstanceprotocolsbgpvpn) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.SystemNetworkinstanceProtocolsBgpvpnKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlSystemNetworkinstanceProtocolsBgpvpn{
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
