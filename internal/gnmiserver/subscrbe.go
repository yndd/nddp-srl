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
	"fmt"
	"io"
	"strings"

	"github.com/openconfig/gnmi/coalesce"
	"github.com/openconfig/gnmi/match"
	"github.com/openconfig/gnmi/path"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type syncMarker struct{}

func (s *server) Subscribe(stream gnmi.GNMI_SubscribeServer) error {
	sc := &streamClient{
		stream: stream,
	}
	var err error
	sc.req, err = stream.Recv()
	switch {
	case err == io.EOF:
		return nil
	case err != nil:
		return err
	case sc.req.GetSubscribe() == nil:
		return status.Errorf(codes.InvalidArgument, errGetSubscribe)
	}
	sc.target = sc.req.GetSubscribe().GetPrefix().GetTarget()

	peer, _ := peer.FromContext(stream.Context())
	s.log.Debug("received a subscribe request", "mode", sc.req.GetSubscribe().GetMode(), "from", peer.Addr, "target", sc.target)
	defer s.log.Debug("subscription terminated", "peer", peer.Addr)

	sc.queue = coalesce.NewQueue()
	errChan := make(chan error, 3)
	sc.errChan = errChan

	s.log.Debug("acquiring subscription spot", "Target", sc.target)
	ok := s.subscribeRPCsem.TryAcquire(1)
	if !ok {
		return status.Errorf(codes.ResourceExhausted, "could not acquire a subscription spot")
	}
	s.log.Debug("acquired subscription spot", "Target", sc.target)

	switch sc.req.GetSubscribe().GetMode() {
	case gnmi.SubscriptionList_ONCE:
		go func() {
			s.log.Debug("Handle Subscription", "Mode", "ONCE", "Target", sc.target)
			//s.handleSubscriptionRequest(sc)
			sc.queue.Close()
		}()
	case gnmi.SubscriptionList_POLL:
		go s.log.Debug("Handle Subscription", "Mode", "POLL", "Target", sc.target)
	case gnmi.SubscriptionList_STREAM:
		if sc.req.GetSubscribe().GetUpdatesOnly() {
			sc.queue.Insert(syncMarker{})
		}
		remove := addSubscription(s.m, sc.req.GetSubscribe(), &matchClient{queue: sc.queue})
		defer remove()
		if !sc.req.GetSubscribe().GetUpdatesOnly() {
			s.log.Debug("Handle Subscription", "Mode", "STREAM", "Target", sc.target)
			go s.handleSubscriptionRequest(sc)

		}
	default:
		return status.Errorf(codes.InvalidArgument, "unrecognized subscription mode: %v", sc.req.GetSubscribe().GetMode())
	}
	// send all nodes added to queue
	go s.sendStreamingResults(sc)

	var errs = make([]error, 0)
	for err := range errChan {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		sb := strings.Builder{}
		sb.WriteString("multiple errors occurred:\n")
		for _, err := range errs {
			sb.WriteString(fmt.Sprintf("- %v\n", err))
		}
		return fmt.Errorf("%v", sb)
	}
	return nil
}

type matchClient struct {
	queue *coalesce.Queue
	err   error
}

func (m *matchClient) Update(n interface{}) {
	if m.err != nil {
		return
	}
	_, m.err = m.queue.Insert(n)
}

func addSubscription(m *match.Match, s *gnmi.SubscriptionList, c *matchClient) func() {
	var removes []func()
	prefix := path.ToStrings(s.GetPrefix(), true)
	for _, p := range s.GetSubscription() {
		if p.GetPath() == nil {
			continue
		}

		path := append(prefix, path.ToStrings(p.GetPath(), false)...)
		removes = append(removes, m.AddQuery(path, c))
	}
	return func() {
		for _, remove := range removes {
			remove()
		}
	}
}
