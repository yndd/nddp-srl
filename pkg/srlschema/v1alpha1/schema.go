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
	"fmt"
	"strconv"
	"strings"

	"github.com/yndd/ndd-runtime/pkg/meta"
	"github.com/yndd/ndd-runtime/pkg/utils"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"

	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Schema interface {
	// methods children
	NewDevice(c resource.ClientApplicator, name string) Device
	GetDevices() map[string]Device
	PrintDevices(n string)
	// methods schema/data
	DeploySchema(ctx context.Context, mg resource.Managed, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed) (map[string]map[string]interface{}, error)
	ValidateResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) (map[string]map[string]interface{}, error)
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction) (map[string]map[string]map[string]interface{}, error)
}

func NewSchema(c resource.ClientApplicator) Schema {
	return &schema{
		// k8s client
		client: c,
		// parent is nil/root
		// children
		devices: make(map[string]Device),
		// data key
	}
}

type schema struct {
	// k8s client
	client resource.ClientApplicator
	// parent is nil/root
	// children
	devices map[string]Device
	// data is nil
}

func (x *schema) NewDevice(c resource.ClientApplicator, name string) Device {
	if _, ok := x.devices[name]; !ok {
		x.devices[name] = NewDevice(c, x, name)
	}
	return x.devices[name]
}

func (x *schema) GetDevices() map[string]Device {
	return x.devices
}

func (x *schema) PrintDevices(n string) {
	fmt.Printf("schema information: %s\n", n)
	for deviceName, d := range x.GetDevices() {
		d.Print(deviceName, 1)
	}
}

func (x *schema) DeploySchema(ctx context.Context, mg resource.Managed, labels map[string]string) error {
	for deviceName, d := range x.GetDevices() {
		if err := d.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	// create a transaction
	o := x.buildCR(mg)
	if err := x.client.Apply(ctx, o); err != nil {
		return err
	}
	return nil
}

func (x *schema) InitializeDummySchema() {
	d := x.NewDevice(x.client, "dummy")
	d.InitializeDummySchema()
}

func (x *schema) ListResources(ctx context.Context, mg resource.Managed) (map[string]map[string]interface{}, error) {
	resources := make(map[string]map[string]interface{})
	for _, d := range x.GetDevices() {
		if err := d.ListResources(ctx, mg, resources); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func (x *schema) ValidateResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) (map[string]map[string]interface{}, error) {
	for deviceName, d := range x.GetDevices() {
		if err := d.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func (x *schema) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	for _, d := range x.GetDevices() {
		if err := d.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

// first entry is device, 2nd: resourceKind, 3rd: resourceName
func (x *schema) ListResourcesByTransaction(ctx context.Context, cr srlv1alpha1.IFSrlTransaction) (map[string]map[string]map[string]interface{}, error) {
	resources := make(map[string]map[string]map[string]interface{})
	for _, d := range x.GetDevices() {
		if err := d.ListResourcesByTransaction(ctx, cr, resources); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func (x *schema) buildCR(mg resource.Managed) *srlv1alpha1.SrlTransaction {
	namespace := mg.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}

	return &srlv1alpha1.SrlTransaction{
		ObjectMeta: metav1.ObjectMeta{
			Name:            odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind)),
			Namespace:       namespace,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.TransactionSpec{
			OwnerGeneration: utils.StringPtr(strconv.Itoa(int(mg.GetGeneration()))),
		},
	}
}
