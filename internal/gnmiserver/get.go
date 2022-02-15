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

	"github.com/google/gnxi/utils/xpath"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
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

	prefix := req.GetPrefix()
	target := ""
	if prefix != nil {
		target = prefix.Target
	}
	log := s.log.WithValues("Type", req.GetType())
	if req.GetPath() != nil {
		log.Debug("Get...", "Target", target, "Path", yparser.GnmiPath2XPath(req.GetPath()[0], true))
	} else {
		log.Debug("Get...", "Target", target)
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
		// if the cache is exhausted we need to backoff
		exhausted, err := s.getExhausted(crSystemDeviceName)
		if err != nil {
			return nil, err
		}
		if exhausted != 0 {
			return nil, status.Errorf(codes.ResourceExhausted, "device exhausted")
		}
		gvkTransaction := req.GetExtension()[0].GetRegisteredExt().GetMsg()
		if string(gvkTransaction) == gvkresource.Operation_GetResourceNameFromPath {
			// procedure to get resource name
			updates, err := s.getResourceName(crSystemDeviceName, req.GetPath()[0])
			if err != nil {
				return nil, err
			}
			return updates, nil
		} else {
			// this is a regular get but we need to check in the systemDevice cache
			// if a managed resource exists or not. The outcome determines if the managed resource will be created
			gvkt, err := gvkresource.String2GvkTransaction(string(gvkTransaction))
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			s.log.Debug("gvkTransaction", "gvkname", gvkt.GVKName, "transaction", gvkt.Transaction, "transactionGeneration", gvkt.TransactionGeneration)

			if gvkt.Transaction != gvkresource.TransactionNone {
				// resource which is part of a transaction
				gvk, err := s.getResource(crSystemDeviceName, gvkt.GVKName)
				if err != nil {
					return nil, status.Error(codes.Internal, err.Error())
				}
				if gvk == nil {
					// transaction does not exist
					exists = false
				} else {
					if gvk.Transaction != gvkt.Transaction || gvk.Transactiongeneration != gvk.Transactiongeneration {
						// transaction does not exist
						exists = false
					} else {
						// dont validate the status when we delete -> we assume if you were able to create it you should be able to delete it
						if gvk.Action == systemv1alpha1.E_GvkAction_TransactionDelete {
							switch gvk.Status {
							case systemv1alpha1.E_GvkStatus_Transactionpending:
								// the transaction is not complete yet so we keep it in this status
								return nil, status.Error(codes.AlreadyExists, "")
							}
						} else {
							switch gvk.Status {
							case systemv1alpha1.E_GvkStatus_Failed:
								// we validate the failed status for upates and creates but not for deletes
								// resource exists, but failed, return the spec data
								x, err := s.getSpecdata(crSystemDeviceName, gvk)
								if err != nil {
									return nil, status.Error(codes.Internal, err.Error())
								}
								if updates, err = appendUpdateResponse(x, req.GetPath()[0], updates); err != nil {
									return nil, status.Error(codes.Internal, err.Error())
								}
								return updates, status.Error(codes.FailedPrecondition, "resource exist, but failed")
							case systemv1alpha1.E_GvkStatus_Transactionpending:
								// the transaction is not complete yet so we keep it in this status
								return nil, status.Error(codes.AlreadyExists, "")
							}
						}
					}
				}
			} else {
				// regular resource which is not part of a trancation
				gvk, err := s.getResource(crSystemDeviceName, gvkt.GVKName)
				if err != nil {
					return nil, status.Error(codes.Internal, err.Error())
				}
				if gvk == nil {
					exists = false
				} else {

					if gvk.Action != systemv1alpha1.E_GvkAction_Delete {
						switch gvk.Status {
						case systemv1alpha1.E_GvkStatus_Deletepending:
							// the action did not complete so far
							return nil, status.Error(codes.AlreadyExists, "")
						}
					} else {
						switch gvk.Status {
						case systemv1alpha1.E_GvkStatus_Failed:
							// resource exists, but failed, return the spec data
							x, err := s.getSpecdata(crSystemDeviceName, gvk)
							if err != nil {
								return nil, status.Error(codes.Internal, err.Error())
							}
							if updates, err = appendUpdateResponse(x, req.GetPath()[0], updates); err != nil {
								return nil, status.Error(codes.Internal, err.Error())
							}
							return updates, status.Error(codes.FailedPrecondition, "resource exist, but failed")
						case systemv1alpha1.E_GvkStatus_Createpending, systemv1alpha1.E_GvkStatus_Updatepending:
							// the action did not complete so far
							return nil, status.Error(codes.AlreadyExists, "")
						}
					}
				}
			}
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

func (s *server) getResourceName(crSystemDeviceName string, reqpath *gnmi.Path) ([]*gnmi.Update, error) {
	// provide a string from the gnmi Path, we expect a single path in the GetRequest
	reqPath := yparser.GnmiPath2XPath(reqpath, true)
	// initialize the variables which are used to keep track of the matched strings
	matchedResourceName := ""
	matchedResourcePath := ""
	// loop over the resourceList
	resourceList, err := s.getResourceList(crSystemDeviceName)
	if err != nil {
		return nil, err
	}
	for _, resource := range resourceList {
		if resource != nil {
			//s.log.Debug("K8sResource GetResourceName Match", "resource", resource)
			// check if the string is contained in the path
			if strings.Contains(reqPath, resource.Rootpath) {
				// if there is a better match we use the better match
				if len(resource.Rootpath) > len(matchedResourcePath) {
					matchedResourcePath = resource.Rootpath
					matchedResourceName = resource.Name
				}
			}
		}
	}
	//s.log.Debug("K8sResource GetResourceName Match", "ResourceName", matchedResourceName)

	d, err := json.Marshal(nddv1.ResourceName{
		Name: matchedResourceName,
	})
	if err != nil {
		return nil, err
	}
	updates := make([]*gnmi.Update, 0)
	upd := &gnmi.Update{
		Path: reqpath,
		Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: d}},
	}
	updates = append(updates, upd)
	return updates, nil
}

func (s *server) getSpecdata(crSystemDeviceName string, resource *systemv1alpha1.Gvk) (interface{}, error) {
	x1, err := s.cache.GetJson(
		crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "gvk", Key: map[string]string{"name": resource.Name}},
			},
		},
		s.nddpSchema)
	if err != nil {
		return nil, err
	}
	// remove the rootPath data
	rootPath, err := xpath.ToGNMIPath(resource.Rootpath)
	if err != nil {
		return nil, err
	}
	switch x := x1.(type) {
	case map[string]interface{}:
		x1 = x["data"]
		x1 = getDataFromRootPath(rootPath, x1)
	}
	return x1, nil
}

func getDataFromRootPath(path *gnmi.Path, x1 interface{}) interface{} {
	//fmt.Printf("gnmiserver getDataFromRootPath: %s, data: %v\n", yparser.GnmiPath2XPath(path, true), x1)
	p := yparser.DeepCopyGnmiPath(path)
	if len(p.GetElem()) > 0 {
		hasKey := false
		if len(p.GetElem()[0].GetKey()) > 0 {
			hasKey = true
		}
		switch x := x1.(type) {
		case map[string]interface{}:
			for k, v := range x {
				if k == p.GetElem()[0].GetName() {
					// when the spec data rootpath last element has a key
					// we need to return the first element of the list
					if hasKey {
						switch x := v.(type) {
						case []interface{}:
							x1 = x[0]
						}
					} else {
						x1 = x[p.GetElem()[0].GetName()]
					}
				}
			}

		}
		p.Elem = p.Elem[1:]
		x1 = getDataFromRootPath(p, x1)
	}
	return x1
}

func (s *server) getExhausted(crSystemDeviceName string) (int64, error) {

	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{{Name: "exhausted"}},
	}

	n, err := s.cache.Query(crSystemDeviceName, &gnmi.Path{Target: crSystemDeviceName}, path)
	if err != nil {
		return 0, err
	}
	if n != nil {
		for _, u := range n.GetUpdate() {
			val, err := yparser.GetValue(u.GetVal())
			if err != nil {
				return 0, err
			}
			switch v := val.(type) {
			case int64:
				return v, nil
			}
		}
		return 0, errors.New("unknown type in cache")
	}

	return 0, nil
}
