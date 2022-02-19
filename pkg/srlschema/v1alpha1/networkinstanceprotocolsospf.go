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
	errCreateNetworkinstanceProtocolsOspf = "cannot create NetworkinstanceProtocolsOspf"
	errDeleteNetworkinstanceProtocolsOspf = "cannot delete NetworkinstanceProtocolsOspf"
	errGetNetworkinstanceProtocolsOspf    = "cannot get NetworkinstanceProtocolsOspf"
)

type NetworkinstanceProtocolsOspf interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.NetworkinstanceProtocolsOspf
	Update(x *srlv1alpha1.NetworkinstanceProtocolsOspf)
	AddNetworkinstanceProtocolsOspfInstance(ai *srlv1alpha1.NetworkinstanceProtocolsOspfInstance)
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

func NewNetworkinstanceProtocolsOspf(c resource.ClientApplicator, p Networkinstance, key string) NetworkinstanceProtocolsOspf {
	newNetworkinstanceProtocolsOspfList := func() srlv1alpha1.IFSrlNetworkinstanceProtocolsOspfList {
		return &srlv1alpha1.SrlNetworkinstanceProtocolsOspfList{}
	}
	return &networkinstanceprotocolsospf{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//NetworkinstanceProtocolsOspf: &srlv1alpha1.NetworkinstanceProtocolsOspf{
		//	Name: &name,
		//},
		newNetworkinstanceProtocolsOspfList: newNetworkinstanceProtocolsOspfList,
	}
}

type networkinstanceprotocolsospf struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Networkinstance
	// children
	// Data
	NetworkinstanceProtocolsOspf        *srlv1alpha1.NetworkinstanceProtocolsOspf
	newNetworkinstanceProtocolsOspfList func() srlv1alpha1.IFSrlNetworkinstanceProtocolsOspfList
}

// key type/method

type NetworkinstanceProtocolsOspfKey struct {
}

func WithNetworkinstanceProtocolsOspfKey(key *NetworkinstanceProtocolsOspfKey) string {
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
func (x *networkinstanceprotocolsospf) Update(d *srlv1alpha1.NetworkinstanceProtocolsOspf) {
	x.NetworkinstanceProtocolsOspf = d
}

// methods data
func (x *networkinstanceprotocolsospf) Get() *srlv1alpha1.NetworkinstanceProtocolsOspf {
	return x.NetworkinstanceProtocolsOspf
}

func (x *networkinstanceprotocolsospf) GetKey() []string {
	return strings.Split(x.key, ".")
}

// NetworkinstanceProtocolsOspf instance ospf Ospf []
func (x *networkinstanceprotocolsospf) AddNetworkinstanceProtocolsOspfInstance(ai *srlv1alpha1.NetworkinstanceProtocolsOspfInstance) {
	//x.NetworkinstanceProtocolsOspf.Instance = append(x.NetworkinstanceProtocolsOspf.Instance, ai)
	if len(x.NetworkinstanceProtocolsOspf.Instance) == 0 {
		x.NetworkinstanceProtocolsOspf.Instance = make([]*srlv1alpha1.NetworkinstanceProtocolsOspfInstance, 0)
	}
	found := false
	for _, xx := range x.NetworkinstanceProtocolsOspf.Instance {

		// [name]
		if *xx.Name == *ai.Name {
			found = true
			xx = ai
		}
	}
	if !found {
		x.NetworkinstanceProtocolsOspf.Instance = append(x.NetworkinstanceProtocolsOspf.Instance, ai)
	}
}

// methods schema

func (x *networkinstanceprotocolsospf) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.NetworkinstanceProtocolsOspf)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s NetworkinstanceProtocolsOspf: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s NetworkinstanceProtocolsOspf: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *networkinstanceprotocolsospf) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateNetworkinstanceProtocolsOspf)
		}
	}

	return nil
}

func (x *networkinstanceprotocolsospf) DestroySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Delete(ctx, o); err != nil {
			return errors.Wrap(resource.IgnoreNotFound(err), errDeleteNetworkinstanceProtocolsOspf)
		}
	}

	return nil
}
func (x *networkinstanceprotocolsospf) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlNetworkinstanceProtocolsOspf {
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

	return &srlv1alpha1.SrlNetworkinstanceProtocolsOspf{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.NetworkinstanceProtocolsOspfSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			NetworkInstanceName: &x.parent.GetKey()[0],
			//1
			NetworkinstanceProtocolsOspf: x.NetworkinstanceProtocolsOspf,
		},
	}
}

func (x *networkinstanceprotocolsospf) InitializeDummySchema() {
}

func (x *networkinstanceprotocolsospf) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newNetworkinstanceProtocolsOspfList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetNetworkinstanceProtocolsOspfs() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *networkinstanceprotocolsospf) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		parent0Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
		//1

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(parent0Key0),
				//1
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.NetworkinstanceProtocolsOspfKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *networkinstanceprotocolsospf) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.NetworkinstanceProtocolsOspfKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlNetworkinstanceProtocolsOspf{
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

func (x *networkinstanceprotocolsospf) ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction, resources map[string]map[string]map[string]interface{}) error {
	// options list all resources belonging to the transaction based on transaction owner and generation
	opts := []client.ListOption{
		client.InNamespace(cr.GetNamespace()),
		client.MatchingLabels(map[string]string{
			srlv1alpha1.LabelNddaOwner: cr.GetName(),
			//srlv1alpha1.LabelNddaOwnerGeneration: cr.GetOwnerGeneration(),
		}),
	}
	list := x.newNetworkinstanceProtocolsOspfList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetNetworkinstanceProtocolsOspfs() {
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
