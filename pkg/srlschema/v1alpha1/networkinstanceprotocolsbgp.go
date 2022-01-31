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
	"github.com/yndd/ndd-runtime/pkg/meta"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
)

const (
	errCreateNetworkinstanceProtocolsBgp = "cannot create NetworkinstanceProtocolsBgp"
	errDeleteNetworkinstanceProtocolsBgp = "cannot delete NetworkinstanceProtocolsBgp"
	errGetNetworkinstanceProtocolsBgp    = "cannot get NetworkinstanceProtocolsBgp"
)

type NetworkinstanceProtocolsBgp interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.NetworkinstanceProtocolsBgp
	Update(x *srlv1alpha1.NetworkinstanceProtocolsBgp)
	AddNetworkinstanceProtocolsBgpGroup(ai *srlv1alpha1.NetworkinstanceProtocolsBgpGroup)
	AddNetworkinstanceProtocolsBgpNeighbor(ai *srlv1alpha1.NetworkinstanceProtocolsBgpNeighbor)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewNetworkinstanceProtocolsBgp(c resource.ClientApplicator, p Networkinstance, key string) NetworkinstanceProtocolsBgp {
	newNetworkinstanceProtocolsBgpList := func() srlv1alpha1.IFSrlNetworkinstanceProtocolsBgpList {
		return &srlv1alpha1.SrlNetworkinstanceProtocolsBgpList{}
	}
	return &networkinstanceprotocolsbgp{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//NetworkinstanceProtocolsBgp: &srlv1alpha1.NetworkinstanceProtocolsBgp{
		//	Name: &name,
		//},
		newNetworkinstanceProtocolsBgpList: newNetworkinstanceProtocolsBgpList,
	}
}

type networkinstanceprotocolsbgp struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Networkinstance
	// children
	// Data
	NetworkinstanceProtocolsBgp        *srlv1alpha1.NetworkinstanceProtocolsBgp
	newNetworkinstanceProtocolsBgpList func() srlv1alpha1.IFSrlNetworkinstanceProtocolsBgpList
}

// key type/method

type NetworkinstanceProtocolsBgpKey struct {
}

func WithNetworkinstanceProtocolsBgpKey(key *NetworkinstanceProtocolsBgpKey) string {
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
func (x *networkinstanceprotocolsbgp) Update(d *srlv1alpha1.NetworkinstanceProtocolsBgp) {
	x.NetworkinstanceProtocolsBgp = d
}

// methods data
func (x *networkinstanceprotocolsbgp) Get() *srlv1alpha1.NetworkinstanceProtocolsBgp {
	return x.NetworkinstanceProtocolsBgp
}

func (x *networkinstanceprotocolsbgp) GetKey() []string {
	return strings.Split(x.key, ".")
}

// NetworkinstanceProtocolsBgp group bgp Bgp []
func (x *networkinstanceprotocolsbgp) AddNetworkinstanceProtocolsBgpGroup(ai *srlv1alpha1.NetworkinstanceProtocolsBgpGroup) {
	x.NetworkinstanceProtocolsBgp.Group = append(x.NetworkinstanceProtocolsBgp.Group, ai)
}

// NetworkinstanceProtocolsBgp neighbor bgp Bgp []
func (x *networkinstanceprotocolsbgp) AddNetworkinstanceProtocolsBgpNeighbor(ai *srlv1alpha1.NetworkinstanceProtocolsBgpNeighbor) {
	x.NetworkinstanceProtocolsBgp.Neighbor = append(x.NetworkinstanceProtocolsBgp.Neighbor, ai)
}

// methods schema

func (x *networkinstanceprotocolsbgp) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.NetworkinstanceProtocolsBgp)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s NetworkinstanceProtocolsBgp: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s NetworkinstanceProtocolsBgp: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *networkinstanceprotocolsbgp) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateNetworkinstanceProtocolsBgp)
		}
	}

	return nil
}
func (x *networkinstanceprotocolsbgp) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlNetworkinstanceProtocolsBgp {
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
	return &srlv1alpha1.SrlNetworkinstanceProtocolsBgp{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.NetworkinstanceProtocolsBgpSpec{
			NetworkInstanceName: &parent0Key0,
			//1
			NetworkinstanceProtocolsBgp: x.NetworkinstanceProtocolsBgp,
		},
	}
}

func (x *networkinstanceprotocolsbgp) InitializeDummySchema() {
}

func (x *networkinstanceprotocolsbgp) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newNetworkinstanceProtocolsBgpList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetNetworkinstanceProtocolsBgps() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *networkinstanceprotocolsbgp) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		parent0Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
		//1

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(parent0Key0),
				//1
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.NetworkinstanceProtocolsBgpKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *networkinstanceprotocolsbgp) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.NetworkinstanceProtocolsBgpKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlNetworkinstanceProtocolsBgp{
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
