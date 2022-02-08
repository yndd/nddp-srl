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

	"github.com/yndd/nddo-runtime/pkg/resource"
)

type Device interface {
	// methods children
	NewNetworkinstance(c resource.ClientApplicator, key string) Networkinstance
	NewBfd(c resource.ClientApplicator, key string) Bfd
	NewSystemNtp(c resource.ClientApplicator, key string) SystemNtp
	NewSystemNetworkinstanceProtocolsBgpvpn(c resource.ClientApplicator, key string) SystemNetworkinstanceProtocolsBgpvpn
	NewSystemNetworkinstanceProtocolsEvpn(c resource.ClientApplicator, key string) SystemNetworkinstanceProtocolsEvpn
	NewInterface(c resource.ClientApplicator, key string) Interface
	NewRoutingpolicyCommunityset(c resource.ClientApplicator, key string) RoutingpolicyCommunityset
	NewRoutingpolicyPrefixset(c resource.ClientApplicator, key string) RoutingpolicyPrefixset
	NewSystemName(c resource.ClientApplicator, key string) SystemName
	NewTunnelinterface(c resource.ClientApplicator, key string) Tunnelinterface
	NewRoutingpolicyAspathset(c resource.ClientApplicator, key string) RoutingpolicyAspathset
	NewRoutingpolicyPolicy(c resource.ClientApplicator, key string) RoutingpolicyPolicy
	GetNetworkinstances() map[string]Networkinstance
	GetBfds() map[string]Bfd
	GetSystemNtps() map[string]SystemNtp
	GetSystemNetworkinstanceProtocolsBgpvpns() map[string]SystemNetworkinstanceProtocolsBgpvpn
	GetSystemNetworkinstanceProtocolsEvpns() map[string]SystemNetworkinstanceProtocolsEvpn
	GetInterfaces() map[string]Interface
	GetRoutingpolicyCommunitysets() map[string]RoutingpolicyCommunityset
	GetRoutingpolicyPrefixsets() map[string]RoutingpolicyPrefixset
	GetSystemNames() map[string]SystemName
	GetTunnelinterfaces() map[string]Tunnelinterface
	GetRoutingpolicyAspathsets() map[string]RoutingpolicyAspathset
	GetRoutingpolicyPolicys() map[string]RoutingpolicyPolicy
	// methods data
	GetKey() []string
	Get() interface{}
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewDevice(c resource.ClientApplicator, p Schema, key string) Device {
	return &device{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		Networkinstance:                      make(map[string]Networkinstance),
		Bfd:                                  make(map[string]Bfd),
		SystemNtp:                            make(map[string]SystemNtp),
		SystemNetworkinstanceProtocolsBgpvpn: make(map[string]SystemNetworkinstanceProtocolsBgpvpn),
		SystemNetworkinstanceProtocolsEvpn:   make(map[string]SystemNetworkinstanceProtocolsEvpn),
		Interface:                            make(map[string]Interface),
		RoutingpolicyCommunityset:            make(map[string]RoutingpolicyCommunityset),
		RoutingpolicyPrefixset:               make(map[string]RoutingpolicyPrefixset),
		SystemName:                           make(map[string]SystemName),
		Tunnelinterface:                      make(map[string]Tunnelinterface),
		RoutingpolicyAspathset:               make(map[string]RoutingpolicyAspathset),
		RoutingpolicyPolicy:                  make(map[string]RoutingpolicyPolicy),
		// data key
		//Device: &srlv1alpha1.Device{
		//	Name: &name,
		//},
	}
}

type device struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Schema
	// children
	Networkinstance                      map[string]Networkinstance
	Bfd                                  map[string]Bfd
	SystemNtp                            map[string]SystemNtp
	SystemNetworkinstanceProtocolsBgpvpn map[string]SystemNetworkinstanceProtocolsBgpvpn
	SystemNetworkinstanceProtocolsEvpn   map[string]SystemNetworkinstanceProtocolsEvpn
	Interface                            map[string]Interface
	RoutingpolicyCommunityset            map[string]RoutingpolicyCommunityset
	RoutingpolicyPrefixset               map[string]RoutingpolicyPrefixset
	SystemName                           map[string]SystemName
	Tunnelinterface                      map[string]Tunnelinterface
	RoutingpolicyAspathset               map[string]RoutingpolicyAspathset
	RoutingpolicyPolicy                  map[string]RoutingpolicyPolicy
	// Data
}

// key type/method

type DeviceKey struct {
	Name string
}

func WithDeviceKey(key *DeviceKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
func (x *device) NewNetworkinstance(c resource.ClientApplicator, key string) Networkinstance {
	if _, ok := x.Networkinstance[key]; !ok {
		x.Networkinstance[key] = NewNetworkinstance(c, x, key)
	}
	return x.Networkinstance[key]
}
func (x *device) NewBfd(c resource.ClientApplicator, key string) Bfd {
	if _, ok := x.Bfd[key]; !ok {
		x.Bfd[key] = NewBfd(c, x, key)
	}
	return x.Bfd[key]
}
func (x *device) NewSystemNtp(c resource.ClientApplicator, key string) SystemNtp {
	if _, ok := x.SystemNtp[key]; !ok {
		x.SystemNtp[key] = NewSystemNtp(c, x, key)
	}
	return x.SystemNtp[key]
}
func (x *device) NewSystemNetworkinstanceProtocolsBgpvpn(c resource.ClientApplicator, key string) SystemNetworkinstanceProtocolsBgpvpn {
	if _, ok := x.SystemNetworkinstanceProtocolsBgpvpn[key]; !ok {
		x.SystemNetworkinstanceProtocolsBgpvpn[key] = NewSystemNetworkinstanceProtocolsBgpvpn(c, x, key)
	}
	return x.SystemNetworkinstanceProtocolsBgpvpn[key]
}
func (x *device) NewSystemNetworkinstanceProtocolsEvpn(c resource.ClientApplicator, key string) SystemNetworkinstanceProtocolsEvpn {
	if _, ok := x.SystemNetworkinstanceProtocolsEvpn[key]; !ok {
		x.SystemNetworkinstanceProtocolsEvpn[key] = NewSystemNetworkinstanceProtocolsEvpn(c, x, key)
	}
	return x.SystemNetworkinstanceProtocolsEvpn[key]
}
func (x *device) NewInterface(c resource.ClientApplicator, key string) Interface {
	if _, ok := x.Interface[key]; !ok {
		x.Interface[key] = NewInterface(c, x, key)
	}
	return x.Interface[key]
}
func (x *device) NewRoutingpolicyCommunityset(c resource.ClientApplicator, key string) RoutingpolicyCommunityset {
	if _, ok := x.RoutingpolicyCommunityset[key]; !ok {
		x.RoutingpolicyCommunityset[key] = NewRoutingpolicyCommunityset(c, x, key)
	}
	return x.RoutingpolicyCommunityset[key]
}
func (x *device) NewRoutingpolicyPrefixset(c resource.ClientApplicator, key string) RoutingpolicyPrefixset {
	if _, ok := x.RoutingpolicyPrefixset[key]; !ok {
		x.RoutingpolicyPrefixset[key] = NewRoutingpolicyPrefixset(c, x, key)
	}
	return x.RoutingpolicyPrefixset[key]
}
func (x *device) NewSystemName(c resource.ClientApplicator, key string) SystemName {
	if _, ok := x.SystemName[key]; !ok {
		x.SystemName[key] = NewSystemName(c, x, key)
	}
	return x.SystemName[key]
}
func (x *device) NewTunnelinterface(c resource.ClientApplicator, key string) Tunnelinterface {
	if _, ok := x.Tunnelinterface[key]; !ok {
		x.Tunnelinterface[key] = NewTunnelinterface(c, x, key)
	}
	return x.Tunnelinterface[key]
}
func (x *device) NewRoutingpolicyAspathset(c resource.ClientApplicator, key string) RoutingpolicyAspathset {
	if _, ok := x.RoutingpolicyAspathset[key]; !ok {
		x.RoutingpolicyAspathset[key] = NewRoutingpolicyAspathset(c, x, key)
	}
	return x.RoutingpolicyAspathset[key]
}
func (x *device) NewRoutingpolicyPolicy(c resource.ClientApplicator, key string) RoutingpolicyPolicy {
	if _, ok := x.RoutingpolicyPolicy[key]; !ok {
		x.RoutingpolicyPolicy[key] = NewRoutingpolicyPolicy(c, x, key)
	}
	return x.RoutingpolicyPolicy[key]
}
func (x *device) GetNetworkinstances() map[string]Networkinstance {
	return x.Networkinstance
}
func (x *device) GetBfds() map[string]Bfd {
	return x.Bfd
}
func (x *device) GetSystemNtps() map[string]SystemNtp {
	return x.SystemNtp
}
func (x *device) GetSystemNetworkinstanceProtocolsBgpvpns() map[string]SystemNetworkinstanceProtocolsBgpvpn {
	return x.SystemNetworkinstanceProtocolsBgpvpn
}
func (x *device) GetSystemNetworkinstanceProtocolsEvpns() map[string]SystemNetworkinstanceProtocolsEvpn {
	return x.SystemNetworkinstanceProtocolsEvpn
}
func (x *device) GetInterfaces() map[string]Interface {
	return x.Interface
}
func (x *device) GetRoutingpolicyCommunitysets() map[string]RoutingpolicyCommunityset {
	return x.RoutingpolicyCommunityset
}
func (x *device) GetRoutingpolicyPrefixsets() map[string]RoutingpolicyPrefixset {
	return x.RoutingpolicyPrefixset
}
func (x *device) GetSystemNames() map[string]SystemName {
	return x.SystemName
}
func (x *device) GetTunnelinterfaces() map[string]Tunnelinterface {
	return x.Tunnelinterface
}
func (x *device) GetRoutingpolicyAspathsets() map[string]RoutingpolicyAspathset {
	return x.RoutingpolicyAspathset
}
func (x *device) GetRoutingpolicyPolicys() map[string]RoutingpolicyPolicy {
	return x.RoutingpolicyPolicy
}

// methods data
func (x *device) Get() interface{} {
	return nil
}

func (x *device) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *device) Print(key string, n int) {
	if x.Get() != nil {
		return
	} else {
		fmt.Printf("%s Device: %s\n", strings.Repeat(" ", n), key)
	}

	n++
	for key, i := range x.GetNetworkinstances() {
		i.Print(key, n)
	}
	for key, i := range x.GetBfds() {
		i.Print(key, n)
	}
	for key, i := range x.GetSystemNtps() {
		i.Print(key, n)
	}
	for key, i := range x.GetSystemNetworkinstanceProtocolsBgpvpns() {
		i.Print(key, n)
	}
	for key, i := range x.GetSystemNetworkinstanceProtocolsEvpns() {
		i.Print(key, n)
	}
	for key, i := range x.GetInterfaces() {
		i.Print(key, n)
	}
	for key, i := range x.GetRoutingpolicyCommunitysets() {
		i.Print(key, n)
	}
	for key, i := range x.GetRoutingpolicyPrefixsets() {
		i.Print(key, n)
	}
	for key, i := range x.GetSystemNames() {
		i.Print(key, n)
	}
	for key, i := range x.GetTunnelinterfaces() {
		i.Print(key, n)
	}
	for key, i := range x.GetRoutingpolicyAspathsets() {
		i.Print(key, n)
	}
	for key, i := range x.GetRoutingpolicyPolicys() {
		i.Print(key, n)
	}
}

func (x *device) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		return nil
	}
	for _, r := range x.GetNetworkinstances() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetBfds() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetSystemNtps() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetSystemNetworkinstanceProtocolsBgpvpns() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetSystemNetworkinstanceProtocolsEvpns() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetInterfaces() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetRoutingpolicyCommunitysets() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetRoutingpolicyPrefixsets() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetSystemNames() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetTunnelinterfaces() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetRoutingpolicyAspathsets() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetRoutingpolicyPolicys() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}

	return nil
}

func (x *device) InitializeDummySchema() {
	c0 := x.NewNetworkinstance(x.client, "dummy")
	c0.InitializeDummySchema()
	c1 := x.NewBfd(x.client, "dummy")
	c1.InitializeDummySchema()
	c2 := x.NewSystemNtp(x.client, "dummy")
	c2.InitializeDummySchema()
	c3 := x.NewSystemNetworkinstanceProtocolsBgpvpn(x.client, "dummy")
	c3.InitializeDummySchema()
	c4 := x.NewSystemNetworkinstanceProtocolsEvpn(x.client, "dummy")
	c4.InitializeDummySchema()
	c5 := x.NewInterface(x.client, "dummy")
	c5.InitializeDummySchema()
	c6 := x.NewRoutingpolicyCommunityset(x.client, "dummy")
	c6.InitializeDummySchema()
	c7 := x.NewRoutingpolicyPrefixset(x.client, "dummy")
	c7.InitializeDummySchema()
	c8 := x.NewSystemName(x.client, "dummy")
	c8.InitializeDummySchema()
	c9 := x.NewTunnelinterface(x.client, "dummy")
	c9.InitializeDummySchema()
	c10 := x.NewRoutingpolicyAspathset(x.client, "dummy")
	c10.InitializeDummySchema()
	c11 := x.NewRoutingpolicyPolicy(x.client, "dummy")
	c11.InitializeDummySchema()
}

func (x *device) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list

	// children
	for _, i := range x.GetNetworkinstances() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetBfds() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNtps() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNetworkinstanceProtocolsBgpvpns() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpns() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetInterfaces() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyCommunitysets() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyPrefixsets() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNames() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetTunnelinterfaces() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyAspathsets() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyPolicys() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *device) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation

	// children
	for _, i := range x.GetNetworkinstances() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetBfds() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNtps() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNetworkinstanceProtocolsBgpvpns() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpns() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetInterfaces() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyCommunitysets() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyPrefixsets() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNames() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetTunnelinterfaces() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyAspathsets() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyPolicys() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *device) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion

	// children
	for _, i := range x.GetNetworkinstances() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetBfds() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNtps() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNetworkinstanceProtocolsBgpvpns() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNetworkinstanceProtocolsEvpns() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetInterfaces() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyCommunitysets() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyPrefixsets() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemNames() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetTunnelinterfaces() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyAspathsets() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetRoutingpolicyPolicys() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}

	return nil
}
