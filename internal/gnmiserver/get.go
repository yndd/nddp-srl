/*
Copyright 2021 NDD.

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

package gnmiserver

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl/internal/shared"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Get(ctx context.Context, req *gnmi.GetRequest) (*gnmi.GetResponse, error) {
	ok := s.unaryRPCsem.TryAcquire(1)
	if !ok {
		return nil, status.Errorf(codes.ResourceExhausted, "max number of Unary RPC reached")
	}
	defer s.unaryRPCsem.Release(1)

	log := s.log.WithValues("Type", req.GetType())
	if req.GetPath() != nil {
		log.Debug("Get...", "Path", yparser.GnmiPath2XPath(req.GetPath()[0], true))
	} else {
		log.Debug("Get...")
	}

	// We dont act upon the error here, but we pass it on the response with updates
	updates, err := s.HandleGet(req)
	return &gnmi.GetResponse{
		Notification: []*gnmi.Notification{
			{
				Timestamp: time.Now().UnixNano(),
				Prefix:    req.GetPrefix(),
				Update:    updates,
			},
		},
	}, err
}

func (s *server) HandleGet(req *gnmi.GetRequest) ([]*gnmi.Update, error) {

	//var err error
	updates := make([]*gnmi.Update, 0)

	prefix := req.GetPrefix()
	crDeviceName := req.GetPrefix().GetTarget()
	if !s.cache.GetCache().HasTarget(crDeviceName) {
		return nil, status.Errorf(codes.Unavailable, "cache not ready")
	}

	var schema *yentry.Entry
	var crSystemDeviceName string
	// if the request is for the system resources per leaf we take the target/crDeviceName iso
	// adding the system part
	if strings.HasPrefix(crDeviceName, shared.SystemNamespace) {
		schema = s.nddpSchema
		crSystemDeviceName = crDeviceName
	} else {
		schema = s.deviceSchema
		crSystemDeviceName = shared.GetCrSystemDeviceName(crDeviceName)
	}

	// if the extension is set we check the resourcelist
	// this is needed for the device driver to know when a create should be triggered
	exists := true
	if len(req.GetExtension()) > 0 {
		gvkName := req.GetExtension()[0].GetRegisteredExt().GetMsg()
		gvk, err := s.getResource(crSystemDeviceName, string(gvkName))
		if err != nil {
			return nil, err
		}
		if gvk == nil {
			exists = false
		}
	}

	for _, path := range req.GetPath() {
		x, err := s.cache.GetJson(prefix.GetTarget(), prefix, path, schema)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if updates, err = appendUpdateResponse(x, path, updates); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if exists {
		// resource exists
		return updates, nil
	}
	// the resource does not exists
	return updates, status.Error(codes.NotFound, "resource does not exist")
}

func appendUpdateResponse(data interface{}, path *gnmi.Path, updates []*gnmi.Update) ([]*gnmi.Update, error) {
	var err error
	var d []byte
	//fmt.Printf("data1: %v\n", data)
	if data != nil {
		d, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	//fmt.Printf("data2: %v\n", string(d))

	upd := &gnmi.Update{
		Path: path,
		Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: d}},
	}
	updates = append(updates, upd)
	return updates, nil
}

/*
func (s *server) getResourceList(crSystemDeviceName string) ([]*systemv1alpha1.Gvk, error) {
	rl, err := s.cache.GetJson(crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{Elem: []*gnmi.PathElem{{Name: "gvk"}}},
		s.nddpSchema)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return gvkresource.GetResourceList(rl)
}
*/

func (s *server) getResource(crSystemDeviceName, gvkName string) (*systemv1alpha1.Gvk, error) {
	rl, err := s.cache.GetJson(crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{Elem: []*gnmi.PathElem{
			{Name: "gvk", Key: map[string]string{"name": gvkName}},
		}},
		s.nddpSchema)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return gvkresource.GetResource(rl)
}
