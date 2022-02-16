/*
Copyright 2021 Wim Henderickx.

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

package srl

import (
	"context"
	"fmt"
	"strings"

	"github.com/karimra/gnmic/target"
	gutils "github.com/karimra/gnmic/utils"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/pkg/errors"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/utils"
	"github.com/yndd/ndd-yang/pkg/parser"
	"github.com/yndd/nddp-srl/internal/device"
	"github.com/yndd/nddp-srl/internal/gnmic"
)

const (
	DeviceType    = "nokia-srl"
	State         = "STATE"
	Configuration = "CONFIG"
	encoding      = "JSON_IETF"
	//errors
	errGetGnmiCapabilities     = "gnmi capabilities error"
	errGnmiCreateGetRequest    = "gnmi create get request error"
	errGnmiGet                 = "gnmi get error "
	errGnmiHandleGetResponse   = "gnmi get response error"
	errGnmiCreateSetRequest    = "gnmi create set request error"
	errGnmiSet                 = "gnmi set error "
	errGnmiCreateDeleteRequest = "gnmi create delete request error"
)

func init() {
	device.Register(DeviceType, func() device.Device {
		return new(srl)
	})
}

type srl struct {
	target *target.Target
	log    logging.Logger
	parser *parser.Parser
	//deviceDetails *ndddvrv1.DeviceDetails
}

func (d *srl) Init(opts ...device.DeviceOption) error {
	for _, o := range opts {
		o(d)
	}
	return nil
}

func (d *srl) WithTarget(target *target.Target) {
	d.target = target
}

func (d *srl) WithLogging(log logging.Logger) {
	d.log = log
}

func (d *srl) WithParser(log logging.Logger) {
	d.parser = parser.NewParser(parser.WithLogger((log)))
}

func (d *srl) Capabilities(ctx context.Context) ([]*gnmi.ModelData, error) {
	d.log.Debug("verifying capabilities ...")

	ext := new(gnmi_ext.Extension)
	resp, err := d.target.Capabilities(ctx, ext)
	if err != nil {
		return nil, errors.Wrap(err, errGetGnmiCapabilities)
	}
	//t.log.Debug("Gnmi Capability", "response", resp)

	return resp.SupportedModels, nil
}

func (d *srl) Discover(ctx context.Context) (*ndrv1.DeviceDetails, error) {
	d.log.Debug("Discover SRL details ...")
	var err error
	var p string
	devDetails := &ndrv1.DeviceDetails{
		Type: nddv1.DeviceTypePtr(DeviceType),
	}

	// get version
	p = "/system/app-management/application[name=idb_server]"
	data, err := d.GetFrom(ctx, &p, State)
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}

	version := data["application"].(map[string]interface{})["version"]

	d.log.Info("set sw version type...")
	devDetails.SwVersion = &strings.Split(fmt.Sprintf("%v", version), "-")[0]

	d.log.Debug("gnmi idb application information", "update response", data)
	d.log.Debug("Device details", "sw version", devDetails.SwVersion)

	// Get chassis details
	p = "/platform/chassis"
	data, err = d.GetFrom(ctx, &p, State)
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}

	chassis := data["chassis"].(map[string]interface{})

	chassisType := chassis["type"]
	chassisSerial := chassis["serial-number"]
	chassisMac := chassis["mac-address"]

	d.log.Debug("set hardware type...")
	devDetails.Kind = utils.StringPtr(fmt.Sprintf("%v", chassisType))
	d.log.Debug("set serial number...")
	devDetails.SerialNumber = utils.StringPtr(fmt.Sprintf("%v", chassisSerial))
	d.log.Debug("set mac address...")
	devDetails.MacAddress = utils.StringPtr(fmt.Sprintf("%v", chassisMac))

	d.log.Debug("gnmi platform information", "update response", data)
	d.log.Debug("Device details", "device details", devDetails)

	return devDetails, nil
}

// GetConfig gathers the entire config of the device
func (d *srl) GetConfig(ctx context.Context) (map[string]interface{}, error) {
	var p = "/"
	return d.GetFrom(ctx, &p, Configuration)
}

// Get gathers device config based on a *string path
func (d *srl) GetFrom(ctx context.Context, p *string, from string) (map[string]interface{}, error) {
	req, err := gnmic.CreateGetRequest(p, utils.StringPtr("CONFIG"), utils.StringPtr(encoding))
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}
	return d.gnmiGatherSingle(ctx, req)
}

// Get gathers device config based on a *string path
func (d *srl) Get(ctx context.Context, p *string) (map[string]interface{}, error) {
	return d.GetFrom(ctx, p, Configuration)
}

// GetGnmi gathers config data of a given *gnmi.Path
func (d *srl) GetGnmi(ctx context.Context, p []*gnmi.Path) (map[string]interface{}, error) {
	req, err := gnmic.CreateConfigGetRequest(p, utils.StringPtr("CONFIG"), utils.StringPtr(encoding))
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}
	return d.gnmiGatherSingle(ctx, req)
}

// gnmiGatherSingle takes a given *gnmi.GetRequest and the error variable from its creation. It evaluates the error and continues to gather the data.
func (d *srl) gnmiGatherSingle(ctx context.Context, req *gnmi.GetRequest) (map[string]interface{}, error) {
	u, err := d.gnmiExecGetRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	for _, update := range u {
		//d.log.Debug("GetConfig", "response", update)
		return update.Values, nil
	}
	return nil, nil
}

func (d *srl) gnmiExecGetRequest(ctx context.Context, req *gnmi.GetRequest) ([]gnmic.Update, error) {
	rsp, err := d.target.Get(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiGet, "error", err)
		return nil, errors.Wrap(err, errGnmiGet)
	}
	u, err := gnmic.HandleGetResponse(rsp)
	if err != nil {
		d.log.Debug(errGnmiHandleGetResponse, "error", err)
		return nil, err
	}
	return u, nil
}

func (d *srl) UpdateGnmi(ctx context.Context, u []*gnmi.Update) (*gnmi.SetResponse, error) {

	gnmiPrefix, err := gutils.CreatePrefix("", "")
	if err != nil {
		d.log.Debug(errGnmiSet, "error", err)
		return nil, errors.Wrap(err, "prefix parse error")
	}

	req := &gnmi.SetRequest{
		Prefix: gnmiPrefix,
		Update: u,
	}

	resp, err := d.target.Set(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiSet, "error", err)
		return nil, err
	}
	//d.log.Debug("update response:", "resp", resp)
	return resp, nil
}

func (d *srl) DeleteGnmi(ctx context.Context, p []*gnmi.Path) (*gnmi.SetResponse, error) {
	gnmiPrefix, err := gutils.CreatePrefix("", "")
	if err != nil {
		d.log.Debug(errGnmiSet, "error", err)
		return nil, errors.Wrap(err, "prefix parse error")
	}

	req := &gnmi.SetRequest{
		Prefix: gnmiPrefix,
		Delete: p,
	}

	resp, err := d.target.Set(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiSet, "error", err)
		return nil, err
	}
	//d.log.Debug("delete response:", "resp", resp)

	return resp, nil
}

func (d *srl) SetGnmi(ctx context.Context, u []*gnmi.Update, p []*gnmi.Path) (*gnmi.SetResponse, error) {

	gnmiPrefix, err := gutils.CreatePrefix("", "")
	if err != nil {
		d.log.Debug(errGnmiSet, "error", err)
		return nil, errors.Wrap(err, "prefix parse error")
	}

	req := &gnmi.SetRequest{
		Prefix: gnmiPrefix,
		Update: u,
		Delete: p,
	}

	resp, err := d.target.Set(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiSet, "error", err)
		return nil, err
	}
	//d.log.Debug("set response:", "resp", resp)
	return resp, nil
}
