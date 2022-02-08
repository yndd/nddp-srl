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

package devicedriver

import (
	"context"
	"sync"
	"time"

	"github.com/karimra/gnmic/target"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-yang/pkg/cache"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl/internal/device"
	"github.com/yndd/nddp-srl/internal/device/srl"
	"github.com/yndd/nddp-srl/internal/devicecollector"
	"github.com/yndd/nddp-srl/internal/devicereconciler"
	"github.com/yndd/nddp-srl/internal/gnmiserver"
	"github.com/yndd/nddp-srl/internal/shared"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

const (
	// timers
	//defaultTimeout = 5 * time.Second
	//notReadyTimeout = 10 * time.Second

	// errors
	errCreateGnmiClient      = "cannot create gnmi client"
	errDeviceInitFailed      = "cannot initialize the device"
	errDeviceNotRegistered   = "the device type is not registered"
	errDeviceDiscoveryFailed = "cannot discover device"
)

type DeviceDriver interface {
	WithLogger(log logging.Logger)
	WithClient(c resource.ClientApplicator)
	WithCh(reqCh chan shared.DeviceUpdate, respCh chan shared.DeviceResponse)
	WithDeviceSchema(y *yentry.Entry)
	WithEventCh(eventChs map[string]chan event.GenericEvent)
	Start() error
	Stop() error
}

type Option func(DeviceDriver)

// WithLogger specifies how the collector logs messages.
func WithLogger(l logging.Logger) Option {
	return func(o DeviceDriver) {
		o.WithLogger(l)
	}
}

func WithClient(c resource.ClientApplicator) Option {
	return func(o DeviceDriver) {
		o.WithClient(c)
	}
}

func WithCh(reqCh chan shared.DeviceUpdate, respCh chan shared.DeviceResponse) Option {
	return func(o DeviceDriver) {
		o.WithCh(reqCh, respCh)
	}
}

func WithDeviceSchema(y *yentry.Entry) Option {
	return func(o DeviceDriver) {
		o.WithDeviceSchema(y)
	}
}

func WithEventCh(eventChs map[string]chan event.GenericEvent) Option {
	return func(o DeviceDriver) {
		o.WithEventCh(eventChs)
	}
}

type deviceInfo struct {
	// kubernetes
	ctx context.Context
	//client    resource.ClientApplicator

	// target info
	namespace string
	target    *target.Target
	paths     []*string
	// device info
	device     device.Device
	collector  devicecollector.DeviceCollector
	reconciler devicereconciler.DeviceReconciler
	// dynamic discovered data
	deviceDetails *ndrv1.DeviceDetails
	initialConfig map[string]interface{}
	// chan
	stopCh chan bool // used to stop the child go routines if the device gets deleted
	// logging
	log logging.Logger
}

type deviceDriver struct {
	// gnmi target
	devices      map[string]*deviceInfo
	cache        *cache.Cache
	deviceSchema *yentry.Entry
	// deviceUpdate
	deviceDriverRequestCh  chan shared.DeviceUpdate
	deviceDriverResponseCh chan shared.DeviceResponse

	// kubernetes
	client   resource.ClientApplicator
	eventChs map[string]chan event.GenericEvent
	// server
	server gnmiserver.Server

	ctx    context.Context
	stopCh chan bool
	mutex  sync.RWMutex

	log logging.Logger
}

func New(opts ...Option) DeviceDriver {
	d := &deviceDriver{
		devices: make(map[string]*deviceInfo),
		ctx:     context.Background(),
		stopCh:  make(chan bool),
		mutex:   sync.RWMutex{},
	}

	for _, opt := range opts {
		opt(d)
	}

	// initialize the device cache
	d.cache = cache.New(
		[]string{},
		cache.WithLogging(d.log),
	)

	return d
}

func (d *deviceDriver) WithLogger(l logging.Logger) {
	d.log = l
}

func (d *deviceDriver) WithClient(c resource.ClientApplicator) {
	d.client = c
}

func (d *deviceDriver) WithCh(reqCh chan shared.DeviceUpdate, respCh chan shared.DeviceResponse) {
	d.deviceDriverRequestCh = reqCh
	d.deviceDriverResponseCh = respCh
}

func (d *deviceDriver) WithDeviceSchema(y *yentry.Entry) {
	d.deviceSchema = y
}

func (d *deviceDriver) WithEventCh(eventChs map[string]chan event.GenericEvent) {
	d.eventChs = eventChs
}

func (d *deviceDriver) Start() error {
	d.log.Debug("starting deviceDriver...")

	// start gnmi server
	d.server = gnmiserver.New(
		gnmiserver.WithCache(d.cache),
		gnmiserver.WithLogger(d.log),
		gnmiserver.WithDeviceSchema(d.deviceSchema),
	)
	if err := d.server.Start(); err != nil {
		return err
	}

	// start device haandler, which enables crud operations for devices
	// create, delete, status requests
	errChannel := make(chan error)
	go func() {
		if err := d.startDeviceChangeHandler(); err != nil {
			errChannel <- errors.Wrap(err, "error starting devicedriver ")
		}
		errChannel <- nil
	}()
	return nil
}

func (d *deviceDriver) startDeviceChangeHandler() error {
	d.log.Debug("Starting deviceChangeHandler...")

	for {
		select {
		case du := <-d.deviceDriverRequestCh:
			d.log.Debug("device driver handler", "action", du.Action, "target", du.TargetConfig.Name, "address", du.TargetConfig.Address)
			crDeviceName := shared.GetCrDeviceName(du.Namespace, du.TargetConfig.Name)
			switch du.Action {
			case shared.DeviceStatus:
				d.log.Debug("status", "deviceName", crDeviceName, "devices", d.devices)
				if _, ok := d.devices[crDeviceName]; !ok {
					d.deviceDriverResponseCh <- shared.DeviceResponse{
						Exists: false,
						Error:  nil}
					d.log.Debug("device status does not exist")
				} else {
					d.deviceDriverResponseCh <- shared.DeviceResponse{
						Exists:        true,
						Error:         nil,
						TargetConfig:  d.devices[crDeviceName].target.Config,
						DeviceDetails: d.devices[crDeviceName].deviceDetails,
					}
					d.log.Debug("device status exists")
				}
			case shared.DeviceStart:
				if err := d.createDevice(du); err != nil {
					d.log.Debug("device init failed")
					d.deviceDriverResponseCh <- shared.DeviceResponse{
						Error: err,
					}
					// delete the context since it is not ok to connect to the device
					delete(d.devices, crDeviceName)
				} else {
					d.log.Debug("device init success")
					d.deviceDriverResponseCh <- shared.DeviceResponse{
						Error:         nil,
						DeviceDetails: d.devices[crDeviceName].deviceDetails,
					}
				}
			case shared.DeviceStop:
				// delete the device from the devicelist
				if err := d.deleteDevice(du); err != nil {
					d.log.Debug("device stop failed")
					delete(d.devices, crDeviceName)
					d.deviceDriverResponseCh <- shared.DeviceResponse{Error: err}

				} else {
					d.log.Debug("device stop success")
					delete(d.devices, crDeviceName)
					d.deviceDriverResponseCh <- shared.DeviceResponse{Error: nil}
				}
			}
		case <-d.stopCh:
			d.log.Debug("stopping subscription handler")
		}
	}
}

func (d *deviceDriver) Stop() error {
	d.log.Debug("stopping  deviceDriver...")

	d.stopCh <- true

	return nil
}

func (d *deviceDriver) createDevice(du shared.DeviceUpdate) error {
	crDeviceName := shared.GetCrDeviceName(du.Namespace, du.TargetConfig.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	d.devices[crDeviceName] = &deviceInfo{
		ctx: context.Background(),
		//client:    d.client,
		namespace: du.Namespace,
		paths:     subscriptions,
		stopCh:    make(chan bool),
		log:       d.log,
	}

	// reference the device driver device
	ddd := d.devices[crDeviceName]
	d.log.Debug("initDevice", "deviceName", crDeviceName)

	// create gnmi client
	ddd.target = target.NewTarget(du.TargetConfig)
	if err := ddd.target.CreateGNMIClient(d.ctx, grpc.WithBlock()); err != nil { // TODO add dialopts
		return errors.Wrap(err, errCreateGnmiClient)
	}

	// initialize the device
	if deviceInitializer, ok := device.Devices[srl.DeviceType]; !ok {
		// set the network node condition to not ready
		//ddd.notReady(errDeviceNotRegistered)
		return errors.New(errDeviceNotRegistered)
	} else {
		ddd.device = deviceInitializer()
		if err := ddd.device.Init(
			device.WithLogging(d.log.WithValues("device", du.TargetConfig.Name)),
			device.WithTarget(ddd.target),
		); err != nil {
			return err
		}
	}

	// discover the device
	cap, err := ddd.device.Capabilities(d.ctx)
	if err != nil {
		return err
	}
	d.printDeviceCapabilities(cap)

	// get device details through gnmi
	ddd.deviceDetails, err = ddd.device.Discover(d.ctx)
	if err != nil {
		return err
	}
	d.log.Debug("deviceDetails", "info", ddd.deviceDetails)

	// get initial config through gnmi
	ddd.initialConfig, err = ddd.device.GetConfig(d.ctx)
	if err != nil {
		return err
	}

	// removes the module from the entry names; removes the first entry from the map; return a map/string
	ddd.initialConfig, _, err = yparser.CleanConfig2String(ddd.initialConfig)
	if err != nil {
		return err
	}

	d.log.Debug("initial config", "config", ddd.initialConfig)
	for childName := range d.deviceSchema.GetChildren() {
		d.log.Debug("device schema", "child", childName)
	}

	// initialize cache with target
	if !d.cache.GetCache().HasTarget(crDeviceName) {
		d.cache.GetCache().Add(crDeviceName)
	}
	if !d.cache.GetCache().HasTarget(crSystemDeviceName) {
		d.cache.GetCache().Add(crSystemDeviceName)
	}

	u, err := yparser.GetGranularUpdatesFromJSON(&gnmi.Path{}, ddd.initialConfig, d.deviceSchema)
	if err != nil {
		return err
	}
	n := &gnmi.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix:    &gnmi.Path{Target: crDeviceName},
		Update:    u,
	}

	// update the cache with the latest config from the device
	if err := d.cache.GnmiUpdate(crDeviceName, n); err != nil {
		//log.Debug("handle target update", "error", err, "Path", yparser.GnmiPath2XPath(u.GetPath(), true), "Value", u.GetVal())
		//log.Debug("handle target update", "error", err, "Notification", *n)
		return errors.New("cache update failed")
	}

	// start per device reconciler
	ddd.reconciler, err = devicereconciler.New(du.TargetConfig, du.Namespace,
		devicereconciler.WithDevice(ddd.device),
		devicereconciler.WithCache(d.cache),
		devicereconciler.WithLogger(d.log),
	)
	if err != nil {
		return err
	}
	ddd.reconciler.Start()

	// start per device collector
	ddd.collector, err = devicecollector.New(du.TargetConfig, du.Namespace, ddd.paths,
		devicecollector.WithCache(d.cache),
		devicecollector.WithLogger(d.log),
		devicecollector.WithEventCh(d.eventChs),
	)
	if err != nil {
		return err
	}
	ddd.collector.Start()

	// TODO SET THE CACHE STATE TO READY

	return nil
}

// getDeviceType returns the devicetype using the registered data from the provider
func (d *deviceDriver) printDeviceCapabilities(gnmiCap []*gnmi.ModelData) {
	//for _, sm := range gnmiCap {
	d.log.Debug("device capabilities", "capability", gnmiCap[0])
	//}
	//return srl.DeviceType
}

func (d *deviceDriver) deleteDevice(du shared.DeviceUpdate) error {
	crDeviceName := shared.GetCrDeviceName(du.Namespace, du.TargetConfig.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)
	// stop the collector
	if ddd, ok := d.devices[crDeviceName]; ok {
		ddd.collector.Stop()
	}

	// delete the device from the cache
	d.cache.GetCache().Remove(crDeviceName)
	d.cache.GetCache().Remove(crSystemDeviceName)

	return nil
}
