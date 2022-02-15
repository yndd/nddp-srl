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
	"context"
	"time"

	"github.com/karimra/gnmic/target"
	"github.com/karimra/gnmic/types"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-yang/pkg/cache"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/nddp-srl/internal/device"
	deviceschema "github.com/yndd/nddp-srl/pkg/yangschema"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	nddpschema "github.com/yndd/nddp-system/pkg/yangschema"
	"google.golang.org/grpc"
)

const (
	// timers
	reconcileTimer = 1 * time.Second

	// errors
	errCreateGnmiClient = "cannot create gnmi client"
)

// DeviceCollector defines the interfaces for the collector
type DeviceReconciler interface {
	Start() error
	Stop() error
	WithLogger(log logging.Logger)
	WithCache(c *cache.Cache)
	WithDevice(d device.Device)
}

// Option can be used to manipulate Options.
type Option func(DeviceReconciler)

// WithLogger specifies how the collector logs messages.
func WithLogger(log logging.Logger) Option {
	return func(d DeviceReconciler) {
		d.WithLogger(log)
	}
}

func WithCache(c *cache.Cache) Option {
	return func(d DeviceReconciler) {
		d.WithCache(c)
	}
}

func WithDevice(dev device.Device) Option {
	return func(d DeviceReconciler) {
		d.WithDevice(dev)
	}
}

// reconciler defines the parameters for the collector
type reconciler struct {
	namespace string
	target    *target.Target
	device    device.Device
	cache     *cache.Cache
	ctx       context.Context

	nddpSchema   *yentry.Entry
	deviceSchema *yentry.Entry
	//mutex       sync.Mutex

	stopCh chan bool // used to stop the child go routines if the device gets deleted

	log logging.Logger
}

// NewCollector creates a new GNMI collector
func New(t *types.TargetConfig, namespace string, opts ...Option) (DeviceReconciler, error) {
	r := &reconciler{
		namespace: namespace,
		stopCh:    make(chan bool),
		ctx:       context.Background(),
	}
	for _, opt := range opts {
		opt(r)
	}

	r.target = target.NewTarget(t)
	if err := r.target.CreateGNMIClient(r.ctx, grpc.WithBlock()); err != nil { // TODO add dialopts
		return nil, errors.Wrap(err, errCreateGnmiClient)
	}

	r.nddpSchema = nddpschema.InitRoot(nil,
		yentry.WithLogging(r.log))

	r.deviceSchema = deviceschema.InitRoot(nil,
		yentry.WithLogging(r.log))

	return r, nil
}

func (r *reconciler) WithLogger(log logging.Logger) {
	r.log = log
}

func (r *reconciler) WithCache(tc *cache.Cache) {
	r.cache = tc
}

func (r *reconciler) WithDevice(d device.Device) {
	r.device = d
}

// Stop reconciler
func (r *reconciler) Stop() error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	log.Debug("stop device reconciler...")

	r.stopCh <- true

	return nil
}

// Start reconciler
func (r *reconciler) Start() error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	log.Debug("starting device reconciler...")

	errChannel := make(chan error)
	go func() {
		if err := r.run(); err != nil {
			errChannel <- errors.Wrap(err, "error starting device reconciler")
		}
		errChannel <- nil
	}()
	return nil
}

// run reconciler
func (r *reconciler) run() error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	log.Debug("running device reconciler...")

	timeout := make(chan bool, 1)
	timeout <- true

	// set cache status to up
	if err := r.initExhausted(); err != nil {
		return err
	}
	if err := r.setUpdateStatus(true); err != nil {
		return err
	}
	for {
		select {
		case <-timeout:
			time.Sleep(reconcileTimer)
			timeout <- true

			// reconcile cache when:
			// -> device is not exhausted
			// -> new updates from k8s operator are received
			// else dont do anything since we need to wait for an update
			// TODO check if cache is set
			exhausted, err := r.getExhausted()
			if err != nil {
				log.Debug("error getting exhausted", "error", err)
			} else {
				if exhausted == 0 {
					// handle transaction
					tList, err := r.getTransactionList()
					if err != nil {
						log.Debug("error get transaction list", "error", err)
					} else {
						for _, t := range tList {
							if t.Status == systemv1alpha1.E_TransactionStatus_Pending {
								// process the transaction
								if err := r.ReconcileTransaction(r.ctx, t); err != nil {
									log.Debug("transaction reconciler error", "trnsaction", t.Name, "error", err)
								}
								// we only process 1 tranaction at the time
								break
							}
						}
					}

					// handle regular update
					work, _ := r.getUpdateStatus()
					if work {
						if err := r.Reconcile(r.ctx); err != nil {
							log.Debug("reconciler error", "error", err)
						}
					}
				} else {
					exhausted--
					if exhausted < 0 {
						exhausted = 0
					}
					r.setExhausted(exhausted)
				}
			}

		case <-r.stopCh:
			log.Debug("Stopping device reconciler")
			return nil
		}
	}
}
