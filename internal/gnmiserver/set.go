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
	"strings"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl/internal/shared"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Set(ctx context.Context, req *gnmi.SetRequest) (*gnmi.SetResponse, error) {

	ok := s.unaryRPCsem.TryAcquire(1)
	if !ok {
		return nil, status.Errorf(codes.ResourceExhausted, errMaxNbrOfUnaryRPCReached)
	}
	defer s.unaryRPCsem.Release(1)

	numUpdates := len(req.GetUpdate())
	numReplaces := len(req.GetReplace())
	numDeletes := len(req.GetDelete())
	if numUpdates+numReplaces+numDeletes == 0 {
		return nil, status.Errorf(codes.InvalidArgument, errMissingPathsInGNMISet)
	}

	log := s.log.WithValues("numUpdates", numUpdates, "numReplaces", numReplaces, "numDeletes", numDeletes)
	prefix := req.GetPrefix()
	log.Debug("Set", "prefix", prefix)

	if numReplaces > 0 {
		for _, u := range req.GetReplace() {
			if err := s.UpdateCache(prefix, u); err != nil {
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
	}

	if numUpdates > 0 {
		for _, u := range req.GetUpdate() {
			if err := s.UpdateCache(prefix, u); err != nil {
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
	}

	if numDeletes > 0 {
		for _, p := range req.GetDelete() {
			if err := s.DeleteCache(prefix, p); err != nil {
				return nil, status.Errorf(codes.InvalidArgument, err.Error())
			}
		}
	}

	// set the status in the cache to indicate there is work for the reconciler
	if err := s.setUpdateStatus(req); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	return &gnmi.SetResponse{
		Response: []*gnmi.UpdateResult{
			{
				Timestamp: time.Now().UnixNano(),
			},
		}}, nil
}

func (s *server) UpdateCache(prefix *gnmi.Path, u *gnmi.Update) error {
	//v, _ := yparser.GetValue(u.GetVal())
	//s.log.Debug("UpdateCache", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "val", u.GetVal(), "type", reflect.TypeOf(v))
	// Validating in the device schema if a key is present
	hasKey, err := s.hasKey(prefix, u)
	if err != nil {
		return err
	}
	n, err := s.cache.GetNotificationFromUpdate(prefix, u, hasKey)
	if err != nil {
		//log.Debug("GetNotificationFromUpdate Error", "Notification", n, "Error", err)
		return err
	}
	s.log.Debug("UpdateCache", "notification", n)
	if n != nil {
		for _, u := range n.GetUpdate() {
			s.log.Debug("gnmiserver update cache", "notification path", yparser.GnmiPath2XPath(u.GetPath(), true), "val", u.GetVal())
		}

		if err := s.cache.GnmiUpdate(prefix.GetTarget(), n); err != nil {
			//log.Debug("GnmiUpdate Error", "Notification", n, "Error", err)
			return err
		}
	}
	return nil
}

func (s *server) DeleteCache(prefix *gnmi.Path, p *gnmi.Path) error {
	// delete from config cache
	n, err := s.cache.GetNotificationFromDelete(prefix, p)
	if err != nil {
		return err
	}
	for _, d := range n.GetDelete() {
		s.log.Debug("gnmiserver delete cache", "notification", yparser.GnmiPath2XPath(d, true))
	}
	if err := s.cache.GnmiUpdate(prefix.GetTarget(), n); err != nil {
		return err
	}
	return nil
}

func (s *server) setUpdateStatus(req *gnmi.SetRequest) error {
	crDeviceName := req.GetPrefix().GetTarget()
	s.log.Debug("setUpdateStatus", "cacheName", crDeviceName)

	if strings.HasPrefix(crDeviceName, shared.SystemNamespace) {
		crSystemDeviceName := crDeviceName

		if !s.cache.GetCache().HasTarget(crSystemDeviceName) {
			return status.Error(codes.Unavailable, "cache not ready")
		}

		n := &gnmi.Notification{
			Timestamp: time.Now().UnixNano(),
			Prefix:    &gnmi.Path{Target: crSystemDeviceName},
			Update: []*gnmi.Update{
				{
					Path: &gnmi.Path{
						Elem: []*gnmi.PathElem{{Name: "cache-update"}},
					},
					Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_BoolVal{BoolVal: true}},
				},
			},
		}

		if err := s.cache.GnmiUpdate(crSystemDeviceName, n); err != nil {
			return errors.New("cache update failed")
		}
	}

	return nil
}

func (s *server) hasKey(prefix *gnmi.Path, u *gnmi.Update) (bool, error) {
	// update is for the system cache
	crDeviceName := prefix.GetTarget()
	if strings.HasPrefix(crDeviceName, shared.SystemNamespace) {
		// only handle the cases where the data is updated to the cache
		if strings.HasPrefix(yparser.GnmiPath2XPath(u.GetPath(), false), "/gvk/data") {
			//p := yparser.DeepCopyGnmiPath(u.GetPath())
			p := &gnmi.Path{Elem: u.Path.GetElem()[2:]}
			// check the device schema if keys exist
			if len(s.deviceSchema.GetKeys(p)) == 0 {
				s.log.Debug("hasKey", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "Bool", false)
				return false, nil
			} else {
				s.log.Debug("hasKey", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "Bool", true)
				return true, nil
			}
		}
	}
	return false, nil
}
