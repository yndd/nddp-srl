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
	var req *gnmi.GetRequest
	var rsp *gnmi.GetResponse
	devDetails := &ndrv1.DeviceDetails{
		Type: nddv1.DeviceTypePtr(DeviceType),
	}

	p = "/system/app-management/application[name=idb_server]"
	req, err = gnmic.CreateGetRequest(&p, utils.StringPtr(State), utils.StringPtr(encoding))
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}
	rsp, err = d.target.Get(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiGet, "error", err)
		return nil, errors.Wrap(err, errGnmiGet)
	}
	u, err := gnmic.HandleGetResponse(rsp)
	if err != nil {
		d.log.Debug(errGnmiHandleGetResponse, "error", err)
		return nil, errors.Wrap(err, errGnmiHandleGetResponse)
	}
	for _, update := range u {
		// we expect a single response in the get since we target the explicit resource
		switch x := update.Values["application"].(type) {
		case map[string]interface{}:
			for k, v := range x {
				sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]
				switch sk {
				case "version":
					d.log.Info("set sw version type...")
					devDetails.SwVersion = &strings.Split(fmt.Sprintf("%v", v), "-")[0]
				}
			}
		}
		d.log.Debug("gnmi idb application information", "update response", update)
	}
	d.log.Debug("Device details", "sw version", devDetails.SwVersion)

	p = "/platform/chassis"
	req, err = gnmic.CreateGetRequest(&p, utils.StringPtr(State), utils.StringPtr(encoding))
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}
	rsp, err = d.target.Get(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiGet, "error", err)
		return nil, errors.Wrap(err, errGnmiGet)
	}

	u, err = gnmic.HandleGetResponse(rsp)
	if err != nil {
		d.log.Debug(errGnmiHandleGetResponse, "error", err)
		return nil, errors.Wrap(err, errGnmiHandleGetResponse)
	}
	for _, update := range u {
		// we expect a single response in the get since we target the explicit resource
		switch x := update.Values["chassis"].(type) {
		case map[string]interface{}:
			for k, v := range x {
				sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]
				switch sk {
				case "type":
					d.log.Debug("set hardware type...")
					devDetails.Kind = utils.StringPtr(fmt.Sprintf("%v", v))
				case "serial-number":
					d.log.Debug("set serial number...")
					devDetails.SerialNumber = utils.StringPtr(fmt.Sprintf("%v", v))
				case "mac-address":
					d.log.Debug("set mac address...")
					devDetails.MacAddress = utils.StringPtr(fmt.Sprintf("%v", v))
				default:
				}
			}
		}
		d.log.Debug("gnmi platform information", "update response", update)
	}
	d.log.Debug("Device details", "device details", devDetails)

	return devDetails, nil
}

func (d *srl) GetConfig(ctx context.Context) (map[string]interface{}, error) {
	var err error
	var p string
	var req *gnmi.GetRequest
	var rsp *gnmi.GetResponse

	p = "/"
	req, err = gnmic.CreateGetRequest(&p, utils.StringPtr(Configuration), utils.StringPtr("JSON_IETF"))
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}
	rsp, err = d.target.Get(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiGet, "error", err)
		return nil, errors.Wrap(err, errGnmiGet)
	}
	u, err := gnmic.HandleGetResponse(rsp)
	if err != nil {
		d.log.Debug(errGnmiHandleGetResponse, "error", err)
		return nil, errors.Wrap(err, errGnmiHandleGetResponse)
	}
	for _, update := range u {
		d.log.Debug("GetConfig", "response", update)
		return update.Values, nil
	}
	return nil, nil
}

func (d *srl) Get(ctx context.Context, p *string) (map[string]interface{}, error) {
	var err error
	var req *gnmi.GetRequest
	var rsp *gnmi.GetResponse

	req, err = gnmic.CreateGetRequest(p, utils.StringPtr("CONFIG"), utils.StringPtr("JSON_IETF"))
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}
	rsp, err = d.target.Get(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiGet, "error", err)
		return nil, errors.Wrap(err, errGnmiGet)
	}
	u, err := gnmic.HandleGetResponse(rsp)
	if err != nil {
		d.log.Debug(errGnmiHandleGetResponse, "error", err)
		return nil, errors.Wrap(err, errGnmiHandleGetResponse)
	}
	for _, update := range u {
		d.log.Debug("GetConfig", "response", update)
		return update.Values, nil
	}
	return nil, nil
}

func (d *srl) GetGnmi(ctx context.Context, p []*gnmi.Path) (map[string]interface{}, error) {
	var err error
	var req *gnmi.GetRequest
	var rsp *gnmi.GetResponse

	req, err = gnmic.CreateConfigGetRequest(p, utils.StringPtr("CONFIG"), utils.StringPtr("JSON_IETF"))
	if err != nil {
		d.log.Debug(errGnmiCreateGetRequest, "error", err)
		return nil, errors.Wrap(err, errGnmiCreateGetRequest)
	}
	rsp, err = d.target.Get(ctx, req)
	if err != nil {
		d.log.Debug(errGnmiGet, "error", err)
		return nil, errors.Wrap(err, errGnmiGet)
	}
	u, err := gnmic.HandleGetResponse(rsp)
	if err != nil {
		d.log.Debug(errGnmiHandleGetResponse, "error", err)
		return nil, errors.Wrap(err, errGnmiHandleGetResponse)
	}
	for _, update := range u {
		d.log.Debug("GetConfig", "response", update)
		return update.Values, nil
	}
	return nil, nil
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
		return nil, errors.Wrap(err, errGnmiSet)
	}
	d.log.Debug("update response:", "resp", resp)
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
		return nil, errors.Wrap(err, errGnmiSet)
	}
	d.log.Debug("delete response:", "resp", resp)

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
		return nil, errors.Wrap(err, errGnmiSet)
	}
	d.log.Debug("set response:", "resp", resp)
	return resp, nil
}
