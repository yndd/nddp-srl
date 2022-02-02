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

package devicereconciler

import (
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl/internal/shared"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"github.com/yndd/nddp-system/pkg/gvkresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *reconciler) setUpdateStatus(status bool) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	n := &gnmi.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix:    &gnmi.Path{Target: crSystemDeviceName},
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{
					Elem: []*gnmi.PathElem{{Name: "cache-update"}},
				},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_BoolVal{BoolVal: status}},
			},
		},
	}

	if err := r.cache.GnmiUpdate(crSystemDeviceName, n); err != nil {
		return errors.Wrap(err, "cache update failed")
	}
	return nil
}

func (r *reconciler) getUpdateStatus() (bool, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{{Name: "cache-update"}},
	}

	n, err := r.cache.Query(crSystemDeviceName, &gnmi.Path{Target: crSystemDeviceName}, path)
	if err != nil {
		return false, err
	}
	if n != nil {
		for _, u := range n.GetUpdate() {
			val, err := yparser.GetValue(u.GetVal())
			if err != nil {
				return false, err
			}
			switch v := val.(type) {
			case bool:
				return v, nil
			}
		}
		return false, errors.New("unknown type in cache")
	}

	return false, nil
}

func (r *reconciler) getResourceList() ([]*systemv1alpha1.Gvk, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	rl, err := r.cache.GetJson(crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{Elem: []*gnmi.PathElem{{Name: "gvk"}}},
		r.nddpSchema)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return gvkresource.GetResourceList(rl)
}

func (r *reconciler) getResourceListRaw() (interface{}, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	rl, err := r.cache.GetJson(crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{Elem: []*gnmi.PathElem{{Name: "gvk"}}},
		r.nddpSchema)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return rl, nil
}

func (r *reconciler) deleteResource(resourceGvkName string) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	n := &gnmi.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix:    &gnmi.Path{Target: crSystemDeviceName},
		Delete: []*gnmi.Path{
			{
				Elem: []*gnmi.PathElem{{Name: "gvk", Key: map[string]string{"name": resourceGvkName}}},
			},
		},
	}

	if err := r.cache.GnmiUpdate(crSystemDeviceName, n); err != nil {
		return errors.Wrap(err, "cache update failed")
	}
	return nil
}

func (r *reconciler) updateResourceStatus(resourceGvkName string, status systemv1alpha1.E_GvkStatus) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	n := &gnmi.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix:    &gnmi.Path{Target: crSystemDeviceName},
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{
					Elem: []*gnmi.PathElem{
						{Name: "gvk", Key: map[string]string{"name": resourceGvkName}},
						{Name: "status"},
					},
				},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: string(status)}},
			},
		},
	}

	if err := r.cache.GnmiUpdate(crSystemDeviceName, n); err != nil {
		return errors.Wrap(err, "cache update failed")
	}
	return nil
}
