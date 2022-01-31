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
	"sync"
	"time"

	"github.com/karimra/gnmic/target"
	"github.com/karimra/gnmic/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-yang/pkg/cache"
	"github.com/yndd/ndd-yang/pkg/yentry"
	deviceschema "github.com/yndd/nddp-srl/pkg/yangschema"
	nddpschema "github.com/yndd/nddp-system/pkg/yangschema"
	"google.golang.org/grpc"
)

const (
	// timers
	defaultTimeout             = 5 * time.Second
	defaultTargetReceivebuffer = 1000
	defaultLockRetry           = 5 * time.Second
	defaultRetryTimer          = 10 * time.Second

	// errors
	errCreateGnmiClient          = "cannot create gnmi client"
	errCreateSubscriptionRequest = "cannot create subscription request"
)

// DeviceCollector defines the interfaces for the collector
type DeviceCollector interface {
	Start() error
	Stop() error
	WithLogger(log logging.Logger)
	WithCache(c *cache.Cache)
}

// Option can be used to manipulate Options.
type Option func(DeviceCollector)

// WithLogger specifies how the collector logs messages.
func WithLogger(log logging.Logger) Option {
	return func(d DeviceCollector) {
		d.WithLogger(log)
	}
}

func WithCache(c *cache.Cache) Option {
	return func(d DeviceCollector) {
		d.WithCache(c)
	}
}

// collector defines the parameters for the collector
type collector struct {
	namespace           string
	target              *target.Target
	cache               *cache.Cache
	subscriptions       []*Subscription
	ctx                 context.Context
	targetReceiveBuffer uint
	retryTimer          time.Duration

	nddpSchema   *yentry.Entry
	deviceSchema *yentry.Entry

	stopCh chan bool // used to stop the child go routines if the device gets deleted

	mutex sync.RWMutex
	log   logging.Logger
}

// NewCollector creates a new GNMI collector
func New(t *types.TargetConfig, namespace string, paths []*string, opts ...Option) (DeviceCollector, error) {
	c := &collector{
		namespace: namespace,
		subscriptions: []*Subscription{
			{
				name:   "device-config-collector",
				paths:  paths,
				stopCh: make(chan bool),
			},
		},

		stopCh:              make(chan bool),
		mutex:               sync.RWMutex{},
		targetReceiveBuffer: defaultTargetReceivebuffer,
		retryTimer:          defaultRetryTimer,
		ctx:                 context.Background(),
	}
	for _, opt := range opts {
		opt(c)
	}

	c.target = target.NewTarget(t)
	if err := c.target.CreateGNMIClient(c.ctx, grpc.WithBlock()); err != nil { // TODO add dialopts
		return nil, errors.Wrap(err, errCreateGnmiClient)
	}

	c.nddpSchema = nddpschema.InitRoot(nil,
		yentry.WithLogging(c.log))

	c.deviceSchema = deviceschema.InitRoot(nil,
		yentry.WithLogging(c.log))

	return c, nil
}

func (c *collector) WithLogger(log logging.Logger) {
	c.log = log
}

func (c *collector) WithCache(tc *cache.Cache) {
	c.cache = tc
}

func (c *collector) GetTarget() *target.Target {
	return c.target
}

func (c *collector) GetSubscriptions() []*Subscription {
	return c.subscriptions
}

func (c *collector) GetSubscription(subName string) *Subscription {
	for _, s := range c.GetSubscriptions() {
		if s.GetName() == subName {
			return s
		}
	}
	return nil
}

// Lock locks a gnmi collector
func (c *collector) Lock() {
	c.mutex.RLock()
}

// Unlock unlocks a gnmi collector
func (c *collector) Unlock() {
	c.mutex.RUnlock()
}

// StartGnmiSubscriptionHandler starts gnmi subscription
func (c *collector) Start() error {
	log := c.log.WithValues("target", c.target.Config.Name, "address", c.target.Config.Address)
	log.Debug("starting device collector...")

	errChannel := make(chan error)
	go func() {
		if err := c.run(); err != nil {
			errChannel <- errors.Wrap(err, "error starting metriccollector ")
		}
		errChannel <- nil
	}()
	return nil
}

// run metric collector
func (c *collector) run() error {
	log := c.log.WithValues("target", c.target.Config.Name, "address", c.target.Config.Address)
	log.Debug("running device collector...")

	c.ctx, c.subscriptions[0].cancelFn = context.WithCancel(c.ctx)

	c.Lock()
	// this subscription is a go routine that runs until you send a stop through the stopCh
	go c.startSubscription(c.ctx, &gnmi.Path{}, c.GetSubscriptions())
	c.Unlock()

	chanSubResp, chanSubErr := c.GetTarget().ReadSubscriptions()

	// run the response handler
	for {
		select {
		case resp := <-chanSubResp:
			c.handleSubscription(resp.Response)
		case tErr := <-chanSubErr:
			c.log.Debug("subscribe", "error", tErr)
			return errors.New("handle subscription error")
		case <-c.stopCh:
			c.log.Debug("stopping device colletcor process...")
			return nil
		}
	}
}

// StartSubscription starts a subscription
func (c *collector) startSubscription(ctx context.Context, prefix *gnmi.Path, s []*Subscription) error {
	log := c.log.WithValues("target", c.target.Config.Name, "address", c.target.Config.Address, "subscription", s[0].GetName())
	log.Debug("subscription start...")
	// initialize new subscription

	req, err := createSubscriptionRequest(prefix, s[0])
	if err != nil {
		c.log.Debug(errCreateSubscriptionRequest, "error", err)
		return errors.Wrap(err, errCreateSubscriptionRequest)
	}

	//log.Debug("Subscription", "Request", req)
	go func() {
		c.target.Subscribe(ctx, req, s[0].GetName())
	}()
	log.Debug("subscription started ...")

	for {
		select {
		case <-s[0].stopCh: // execute quit
			c.log.Debug("subscription cancelled")
			s[0].cancelFn()
			//c.mutex.Lock()
			// TODO delete subscription from list
			//delete(c.subscriptions, subName)
			//c.mutex.Unlock()

			return nil
		}
	}
}

// StartGnmiSubscriptionHandler starts gnmi subscription
func (c *collector) Stop() error {
	log := c.log.WithValues("target", c.target.Config.Name, "address", c.target.Config.Address)
	log.Debug("stop Collector...")

	c.stopSubscription(c.ctx, c.GetSubscriptions()[0])
	c.stopCh <- true

	return nil
}

// StopSubscription stops a subscription
func (c *collector) stopSubscription(ctx context.Context, s *Subscription) error {
	log := c.log.WithValues("target", c.target.Config.Name, "address", c.target.Config.Address)
	log.Debug("stop subscription...")
	//s.stopCh <- true // trigger quit
	s.cancelFn()
	c.log.Debug("subscription stopped")
	return nil
}
