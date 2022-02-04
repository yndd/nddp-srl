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
	errCreateSystemNtp = "cannot create SystemNtp"
	errDeleteSystemNtp = "cannot delete SystemNtp"
	errGetSystemNtp    = "cannot get SystemNtp"
)

type SystemNtp interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.SystemNtp
	Update(x *srlv1alpha1.SystemNtp)
	AddSystemNtpServer(ai *srlv1alpha1.SystemNtpServer)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewSystemNtp(c resource.ClientApplicator, p Device, key string) SystemNtp {
	newSystemNtpList := func() srlv1alpha1.IFSrlSystemNtpList { return &srlv1alpha1.SrlSystemNtpList{} }
	return &systemntp{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//SystemNtp: &srlv1alpha1.SystemNtp{
		//	Name: &name,
		//},
		newSystemNtpList: newSystemNtpList,
	}
}

type systemntp struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	// Data
	SystemNtp        *srlv1alpha1.SystemNtp
	newSystemNtpList func() srlv1alpha1.IFSrlSystemNtpList
}

// key type/method

type SystemNtpKey struct {
}

func WithSystemNtpKey(key *SystemNtpKey) string {
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
func (x *systemntp) Update(d *srlv1alpha1.SystemNtp) {
	x.SystemNtp = d
}

// methods data
func (x *systemntp) Get() *srlv1alpha1.SystemNtp {
	return x.SystemNtp
}

func (x *systemntp) GetKey() []string {
	return strings.Split(x.key, ".")
}

// SystemNtp server ntp Ntp []
func (x *systemntp) AddSystemNtpServer(ai *srlv1alpha1.SystemNtpServer) {
	//x.SystemNtp.Server = append(x.SystemNtp.Server, ai)
	if len(x.SystemNtp.Server) == 0 {
		x.SystemNtp.Server = make([]*srlv1alpha1.SystemNtpServer, 0)
	}
	found := false
	for _, xx := range x.SystemNtp.Server {

		// [address]
		if *xx.Address == *ai.Address {
			found = true
			xx = ai
		}
	}
	if !found {
		x.SystemNtp.Server = append(x.SystemNtp.Server, ai)
	}
}

// methods schema

func (x *systemntp) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.SystemNtp)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s SystemNtp: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s SystemNtp: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *systemntp) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateSystemNtp)
		}
	}

	return nil
}
func (x *systemntp) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlSystemNtp {

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

	return &srlv1alpha1.SrlSystemNtp{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.SystemNtpSpec{
			ResourceSpec: nddv1.ResourceSpec{
				NetworkNodeReference: &nddv1.Reference{
					Name: deviceName,
				},
			},
			SystemNtp: x.SystemNtp,
		},
	}
}

func (x *systemntp) InitializeDummySchema() {
}

func (x *systemntp) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newSystemNtpList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemNtps() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *systemntp) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.SystemNtpKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *systemntp) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.SystemNtpKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlSystemNtp{
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
