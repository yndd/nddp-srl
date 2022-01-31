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
	"github.com/openconfig/gnmi/coalesce"
	"github.com/openconfig/gnmi/ctree"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/subscribe"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type streamClient struct {
	target  string
	req     *gnmi.SubscribeRequest
	queue   *coalesce.Queue
	stream  gnmi.GNMI_SubscribeServer
	errChan chan<- error
}

func (s *server) handleSubscriptionRequest(sc *streamClient) {
	var err error
	s.log.Debug("processing subscription", "Target", sc.target)
	defer func() {
		if err != nil {
			s.log.Debug("error processing subscription", "Target", sc.target, "Error", err)
			sc.queue.Close()
			sc.errChan <- err
			return
		}
		s.log.Debug("subscription request processed", "Target", sc.target)
	}()
	/*
		if !sc.req.GetSubscribe().GetUpdatesOnly() {

			for _, sub := range sc.req.GetSubscribe().GetSubscription() {

				var fp []string
				fp, err = path.CompletePath(sc.req.GetSubscribe().GetPrefix(), sub.GetPath())
				if err != nil {
					return
				}

				err = s.cache.c.Query(sc.target, fp,
					func(_ []string, l *ctree.Leaf, _ interface{}) error {
						if err != nil {
							return err
						}
						_, err = sc.queue.Insert(l)
						return nil
					})
				if err != nil {
					s.log.Debug("failed internal cache query", "Target", sc.target, "Error", err)
					return
				}

			}

		}
	*/
	_, err = sc.queue.Insert(syncMarker{})

}

func (s *server) sendStreamingResults(sc *streamClient) {
	ctx := sc.stream.Context()
	peer, _ := peer.FromContext(ctx)
	s.log.Debug("sending streaming results", "Target", sc.target, "Peer", peer.Addr)
	defer s.subscribeRPCsem.Release(1)
	for {
		item, dup, err := sc.queue.Next(ctx)
		if coalesce.IsClosedQueue(err) {
			sc.errChan <- nil
			return
		}
		if err != nil {
			sc.errChan <- err
			return
		}
		if _, ok := item.(syncMarker); ok {
			err = sc.stream.Send(&gnmi.SubscribeResponse{
				Response: &gnmi.SubscribeResponse_SyncResponse{
					SyncResponse: true,
				}})
			if err != nil {
				sc.errChan <- err
				return
			}
			continue
		}
		node, ok := item.(*ctree.Leaf)
		if !ok || node == nil {
			sc.errChan <- status.Errorf(codes.Internal, "invalid cache node: %+v", item)
			return
		}
		err = s.sendSubscribeResponse(&resp{
			stream: sc.stream,
			n:      node,
			dup:    dup,
		}, sc)
		if err != nil {
			s.log.Debug("failed sending subscribeResponse", "target", sc.target, "error", err)
			sc.errChan <- err
			return
		}
		// TODO: check if target was deleted ? necessary ?
	}
}

type resp struct {
	stream gnmi.GNMI_SubscribeServer
	n      *ctree.Leaf
	dup    uint32
}

func (s *server) sendSubscribeResponse(r *resp, sc *streamClient) error {
	notif, err := subscribe.MakeSubscribeResponse(r.n.Value(), r.dup)
	if err != nil {
		return status.Errorf(codes.Unknown, "unknown error: %v", err)
	}
	// No acls
	return r.stream.Send(notif)

}
