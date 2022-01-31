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

	"github.com/yndd/nddo-runtime/pkg/resource"
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
