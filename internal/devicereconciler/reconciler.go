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
	"sort"
	"strings"
	"time"

	"github.com/google/gnxi/utils/xpath"
	"github.com/karimra/gnmic/target"
	"github.com/karimra/gnmic/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-yang/pkg/cache"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl/internal/device"
	deviceschema "github.com/yndd/nddp-srl/internal/yangschema"
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
	if err := r.setUpdateStatus(true); err != nil {
		return err
	}
	for {
		select {
		case <-timeout:
			time.Sleep(reconcileTimer)
			timeout <- true

			// reconcile cache when:
			// -> new updates from k8s operator are received
			// else dont do anything since we need to wait for an update
			// TODO check if cache is set
			work, _ := r.getUpdateStatus()
			if work {
				if err := r.Reconcile(r.ctx); err != nil {
					log.Debug("error reconciler", "error", err)
				}
			}

		case <-r.stopCh:
			log.Debug("Stopping device reconciler")
			return nil
		}
	}
}

func (r *reconciler) Reconcile(ctx context.Context) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	log.Debug("reconciling device config")

	// gte the list of MR
	resourceList, err := r.getResourceList()
	if err != nil {
		return err
	}
	resourceListRaw, err := r.getResourceListRaw()
	if err != nil {
		return err
	}
	log.Debug("resourceList1", "raw", resourceListRaw)

	// sort the MR list based on the pathElements
	sort.SliceStable(resourceList, func(i, j int) bool {
		iPathElem := len(strings.Split(*resourceList[i].Rootpath, "/"))
		jPathElem := len(strings.Split(*resourceList[i].Rootpath, "/"))
		return iPathElem < jPathElem
	})

	// process updates
	// we go straight to the device, a repalce is dangerous as it can affect traffic
	// deletes needs to be handled directly to the cache and the device
	for _, resource := range resourceList {
		if resource.Status == systemv1alpha1.E_GvkStatus_Updatepending {
			// check deletes, updates
			deletes, updates, err := r.processUpdates(resource)
			if err != nil {
				return err
			}
			// debug
			for _, d := range deletes {
				log.Debug("Update deletes", "delPath", yparser.GnmiPath2XPath(d, true))
			}
			for _, u := range updates {
				log.Debug("Update updates", "updPath", yparser.GnmiPath2XPath(u.GetPath(), true), "val", u.GetVal())
			}

			// execute the deletes and updates in the cache and to the device
			_, err = r.device.SetGnmi(r.ctx, updates, deletes)
			if err != nil {
				// Set status to failed
				if err := r.updateResourceStatus(*resource.Name, systemv1alpha1.E_GvkStatus_Failed); err != nil {
					return err
				}
				return err
			}
			if err := r.updateResourceStatus(*resource.Name, systemv1alpha1.E_GvkStatus_Success); err != nil {
				return err
			}
			// TBD, could we use the on change notification to update the cache ???
			// Avoids updating the cache

		}
	}

	// initialize the candidate cache; delete/recreate if it exists
	r.initializeCandidateCache()

	// process deletes first,
	// create a list of resources and paths to be deletes
	delResources := make([]*systemv1alpha1.Gvk, 0)
	delPaths := make([]*gnmi.Path, 0)
	for _, resource := range resourceList {
		// all the dependencies should be taken care of with the leafref validations
		// in the provider
		// maybe aggregating some deletes if they have a parent dependency might be needed
		if resource.Status == systemv1alpha1.E_GvkStatus_Deletepending {
			delResources = append(delResources, resource)
			path, err := xpath.ToGNMIPath(*resource.Rootpath)
			if err != nil {
				return err
			}
			delPaths = append(delPaths, path)
		}
	}

	// delete the paths if there are present in a single transaction
	// if there is a path which is /, it would kill the ox; we protect this
	murder := false
	for _, delPath := range delPaths {
		log.Debug("Delete", "Path", delPath)
		if delPath == nil {
			murder = true
		}
	}
	// if we dont do suicide and len delete paths > 0, perform delete
	if !murder && len(delPaths) > 0 {
		// apply deletes on the device
		_, err := r.device.DeleteGnmi(ctx, delPaths)
		if err != nil {
			// TODO we should fail in certain consitions
			// we keep the status in DeletePending to retry
			log.Debug("gnmi delete failed", "Paths", delPaths, "Error", err)
		} else {
			log.Debug("gnmi delete success", "Paths", delPaths)
			// delete the paths from the device cache
			r.deletePathsFromCache(delPaths)

			// delete resources
			for _, resource := range delResources {
				r.deleteResource(*resource.Name)
			}
		}
	}

	// get copy from cache
	if err := r.copyRunning2Candidate(); err != nil {
		return err
	}

	// debug
	resourceListRaw, err = r.getResourceListRaw()
	if err != nil {
		return err
	}
	log.Debug("resourceList2a", "raw", resourceListRaw)

	doUpdate := false
	// create a list of resources to be updated
	updResources := make([]*systemv1alpha1.Gvk, 0)
	for _, resource := range resourceList {
		// only merge the resources that are NOT in failed and delete pending state
		if resource.Status != systemv1alpha1.E_GvkStatus_Deletepending && resource.Status != systemv1alpha1.E_GvkStatus_Failed {
			// append to the resource list of resources needed an update
			updResources = append(updResources, resource)
			// update the candidate cache with the resource data
			if err := r.updateCandidate(resource); err != nil {
				return err
			}
			doUpdate = true
		}
	}

	// debug
	resourceListRaw, err = r.getResourceListRaw()
	if err != nil {
		return err
	}
	log.Debug("resourceList2b", "raw", resourceListRaw)

	if doUpdate {
		// retrieve the config that will be applied to the device
		updates, ok, err := r.getCandidateUpdate()
		if err != nil {
			return err
		}
		if ok {
			if _, err := r.device.UpdateGnmi(ctx, updates); err != nil {
				// get candidate config to show which configuration failed
				candidateConfig, err := r.getCandidateConfig()
				if err != nil {
					// TODO set status to failed and stop reconciliation ???
					return err
				}
				// print failed candidate config
				log.Debug("candidate config failed", "config", candidateConfig)
				// set resource status to failed
				for _, resource := range updResources {
					if err := r.updateResourceStatus(*resource.Name, systemv1alpha1.E_GvkStatus_Failed); err != nil {
						return err
					}
				}
			}
			// debug
			resourceListRaw, err := r.getResourceListRaw()
			if err != nil {
				return err
			}
			log.Debug("resourceList3", "raw", resourceListRaw)

			// set resource status to success
			for _, resource := range updResources {
				if err := r.updateResourceStatus(*resource.Name, systemv1alpha1.E_GvkStatus_Success); err != nil {
					return err
				}
			}

			// debug
			resourceListRaw, err = r.getResourceListRaw()
			if err != nil {
				return err
			}
			log.Debug("resourceList4", "raw", resourceListRaw)

			if err := r.copyCandidate2Running(); err != nil {
				return err
			}

		} else {
			log.Debug("gnmi update empty")
		}
	}

	// debug
	resourceListRaw, err = r.getResourceListRaw()
	if err != nil {
		return err
	}
	log.Debug("resourceList5", "raw", resourceListRaw)

	// set reconcile flag to false to avoid a new reconciliation if there is no new work
	if err := r.setUpdateStatus(false); err != nil {
		return err
	}
	return nil
}
