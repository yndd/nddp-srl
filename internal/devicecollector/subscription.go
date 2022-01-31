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

package devicecollector

import (
	"context"
	"strings"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndd-yang/pkg/yparser"
)

// Subscription defines the parameters for the subscription
type Subscription struct {
	name  string
	paths []*string

	//ctx      context.Context
	stopCh   chan bool
	cancelFn context.CancelFunc
}

func (s *Subscription) GetName() string {
	return s.name
}

func (s *Subscription) GetPaths() []*gnmi.Path {
	paths := []*gnmi.Path{}
	for _, p := range s.paths {
		paths = append(paths, yparser.Xpath2GnmiPath(*p, 0))
	}

	return paths
}

func (s *Subscription) SetName(n string) {
	s.name = n
}

func (s *Subscription) SetPaths(p []*string) {
	s.paths = p
}

func (s *Subscription) SetStopChannel(c chan bool) {
	s.stopCh = c
}

func (s *Subscription) SetCancelFn(c context.CancelFunc) {
	s.cancelFn = c
}

// CreateSubscriptionRequest create a gnmi subscription
func createSubscriptionRequest(prefix *gnmi.Path, s *Subscription) (*gnmi.SubscribeRequest, error) {
	// create subscription

	modeVal := gnmi.SubscriptionList_Mode_value[strings.ToUpper("STREAM")]
	qos := &gnmi.QOSMarking{Marking: 21}

	subscriptions := make([]*gnmi.Subscription, len(s.GetPaths()))
	/*
		for i, p := range s.GetPaths() {
			subscriptions[i] = &gnmi.Subscription{Path: p}
			switch gnmi.SubscriptionList_Mode(modeVal) {
			case gnmi.SubscriptionList_STREAM:
				mode := gnmi.SubscriptionMode_value[strings.Replace(strings.ToUpper("SAMPLE"), "-", "_", -1)]
				subscriptions[i].Mode = gnmi.SubscriptionMode(mode)
				subscriptions[i].SampleInterval = uint64((5 * time.Second).Nanoseconds())

			}
		}
	*/
	for i, p := range s.GetPaths() {
		subscriptions[i] = &gnmi.Subscription{Path: p}
		switch gnmi.SubscriptionList_Mode(modeVal) {
		case gnmi.SubscriptionList_STREAM:
			mode := gnmi.SubscriptionMode_value[strings.Replace(strings.ToUpper("ON_CHANGE"), "-", "_", -1)]
			subscriptions[i].Mode = gnmi.SubscriptionMode(mode)
		}
	}
	req := &gnmi.SubscribeRequest{
		Request: &gnmi.SubscribeRequest_Subscribe{
			Subscribe: &gnmi.SubscriptionList{
				Prefix:       prefix,
				Mode:         gnmi.SubscriptionList_Mode(modeVal),
				Encoding:     46, // ON CHANGE NOTIFICATIONS
				Subscription: subscriptions,
				Qos:          qos,
			},
		},
	}
	return req, nil
}
