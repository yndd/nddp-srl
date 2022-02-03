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
	errCreateNetworkinstanceProtocolsIsis = "cannot create NetworkinstanceProtocolsIsis"
	errDeleteNetworkinstanceProtocolsIsis = "cannot delete NetworkinstanceProtocolsIsis"
	errGetNetworkinstanceProtocolsIsis    = "cannot get NetworkinstanceProtocolsIsis"
)

type NetworkinstanceProtocolsIsis interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.NetworkinstanceProtocolsIsis
	Update(x *srlv1alpha1.NetworkinstanceProtocolsIsis)
	AddNetworkinstanceProtocolsIsisInstance(ai *srlv1alpha1.NetworkinstanceProtocolsIsisInstance)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewNetworkinstanceProtocolsIsis(c resource.ClientApplicator, p Networkinstance, key string) NetworkinstanceProtocolsIsis {
	newNetworkinstanceProtocolsIsisList := func() srlv1alpha1.IFSrlNetworkinstanceProtocolsIsisList {
		return &srlv1alpha1.SrlNetworkinstanceProtocolsIsisList{}
	}
	return &networkinstanceprotocolsisis{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//NetworkinstanceProtocolsIsis: &srlv1alpha1.NetworkinstanceProtocolsIsis{
		//	Name: &name,
		//},
		newNetworkinstanceProtocolsIsisList: newNetworkinstanceProtocolsIsisList,
	}
}

type networkinstanceprotocolsisis struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Networkinstance
	// children
	// Data
	NetworkinstanceProtocolsIsis        *srlv1alpha1.NetworkinstanceProtocolsIsis
	newNetworkinstanceProtocolsIsisList func() srlv1alpha1.IFSrlNetworkinstanceProtocolsIsisList
}

// key type/method

type NetworkinstanceProtocolsIsisKey struct {
}

func WithNetworkinstanceProtocolsIsisKey(key *NetworkinstanceProtocolsIsisKey) string {
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
func (x *networkinstanceprotocolsisis) Update(d *srlv1alpha1.NetworkinstanceProtocolsIsis) {
	x.NetworkinstanceProtocolsIsis = d
}

// methods data
func (x *networkinstanceprotocolsisis) Get() *srlv1alpha1.NetworkinstanceProtocolsIsis {
	return x.NetworkinstanceProtocolsIsis
}

func (x *networkinstanceprotocolsisis) GetKey() []string {
	return strings.Split(x.key, ".")
}

// NetworkinstanceProtocolsIsis instance isis Isis []
func (x *networkinstanceprotocolsisis) AddNetworkinstanceProtocolsIsisInstance(ai *srlv1alpha1.NetworkinstanceProtocolsIsisInstance) {
	x.NetworkinstanceProtocolsIsis.Instance = append(x.NetworkinstanceProtocolsIsis.Instance, ai)
}

// methods schema

func (x *networkinstanceprotocolsisis) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.NetworkinstanceProtocolsIsis)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s NetworkinstanceProtocolsIsis: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s NetworkinstanceProtocolsIsis: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *networkinstanceprotocolsisis) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateNetworkinstanceProtocolsIsis)
		}
	}

	return nil
}
func (x *networkinstanceprotocolsisis) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlNetworkinstanceProtocolsIsis {
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
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName
	return &srlv1alpha1.SrlNetworkinstanceProtocolsIsis{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.NetworkinstanceProtocolsIsisSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			NetworkInstanceName: &parent0Key0,
			//1
			NetworkinstanceProtocolsIsis: x.NetworkinstanceProtocolsIsis,
		},
	}
}

func (x *networkinstanceprotocolsisis) InitializeDummySchema() {
}

func (x *networkinstanceprotocolsisis) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newNetworkinstanceProtocolsIsisList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetNetworkinstanceProtocolsIsiss() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *networkinstanceprotocolsisis) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		parent0Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
		//1

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(parent0Key0),
				//1
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.NetworkinstanceProtocolsIsisKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *networkinstanceprotocolsisis) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.NetworkinstanceProtocolsIsisKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlNetworkinstanceProtocolsIsis{
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
