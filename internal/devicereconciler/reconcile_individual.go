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

	"github.com/google/gnxi/utils/xpath"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndd-yang/pkg/yparser"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *reconciler) Reconcile(ctx context.Context) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	log.Debug("reconciling device config")

	// get the list of MR
	resourceList, err := r.getResourceList()
	if err != nil {
		return err
	}

	//log.Debug("Reconcile resourceList", "resourceList", resourceList)

	// sort the MR list based on the pathElements
	sort.SliceStable(resourceList, func(i, j int) bool {
		iPathElem := len(strings.Split(resourceList[i].Rootpath, "/"))
		jPathElem := len(strings.Split(resourceList[i].Rootpath, "/"))
		return iPathElem < jPathElem
	})

	// process updates
	// we go straight to the device -> the cache is updated using the notifications
	// the updates are handled per resource and are NOT aggregated accross all resources
	for _, resource := range resourceList {
		if resource.Status == systemv1alpha1.E_GvkStatus_Updatepending {
			// check deletes, updates
			deletes, updates, err := r.processUpdates(resource)
			if err != nil {
				return err
			}
			// debug
			/*
				for _, d := range deletes {
					log.Debug("Update deletes", "delPath", yparser.GnmiPath2XPath(d, true))
				}
				for _, u := range updates {
					log.Debug("Update updates", "updPath", yparser.GnmiPath2XPath(u.GetPath(), true), "val", u.GetVal())
				}
			*/

			// execute the deletes and updates in the cache and to the device
			_, err = r.device.SetGnmi(r.ctx, updates, deletes)
			if err != nil {
				if e, ok := status.FromError(err); ok {
					switch e.Code() {
					case codes.ResourceExhausted:
						log.Debug("gnmi update failed exhausted", "error", err)
						if err := r.setExhausted(60); err != nil {
							return err
						}
					}
				}
				// Set status to failed
				if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Failed); err != nil {
					return err
				}
				return err
			}
			if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Success); err != nil {
				return err
			}
		}
	}

	// initialize the candidate cache; delete/recreate if it exists
	//r.initializeCandidateCache()

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
			path, err := xpath.ToGNMIPath(resource.Rootpath)
			if err != nil {
				return err
			}
			delPaths = append(delPaths, path)
		}
	}

	// we delete all the paths on the device in a single transaction
	// we use the chnage notification to update the cache
	murder := false
	for _, delPath := range delPaths {
		//log.Debug("Delete", "Path", delPath)
		if delPath == nil {
			murder = true
		}
	}
	// if we dont do suicide and len delete paths > 0, perform delete
	if !murder && len(delPaths) > 0 {
		// apply deletes on the device
		_, err := r.device.DeleteGnmi(ctx, delPaths)
		if err != nil {
			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.ResourceExhausted:
					log.Debug("gnmi delate failed exhausted", "error", err)
					if err := r.setExhausted(60); err != nil {
						return err
					}
					// we return and keep the status as is since we can retry once the device is back in normal state
					return nil
				}
			}
			log.Debug("gnmi delete failed", "Paths", delPaths, "Error", err)
			// update the status in the system cache
			for _, resource := range delResources {
				if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Failed); err != nil {
					return err
				}
			}
		} else {
			//log.Debug("gnmi delete success", "Paths", delPaths)

			// delete resources from the system cache
			for _, resource := range delResources {
				r.deleteResource(resource.Name)
			}
		}
	}

	// get copy from cache
	/*
		if err := r.copyRunning2Candidate(); err != nil {
			return err
		}
	*/

	// create a list of resources to be updated
	updResources := make([]*systemv1alpha1.Gvk, 0)
	updates := make([]*gnmi.Update, 0)
	for _, resource := range resourceList {
		// only merge the resources that are NOT in failed and delete pending state
		if resource.Status == systemv1alpha1.E_GvkStatus_Createpending {
			// append to the resource list of resources needed an update
			updResources = append(updResources, resource)
			// update the candidate cache with the resource data
			resUpdate, err := r.getUpdates(resource)
			if err != nil {
				return err
			}
			updates = append(updates, resUpdate)
		}
	}

	if len(updates) > 0 {
		// debug
		/*
			for _, upd := range updates {
				fmt.Printf("Updates: path: %s, data: %v\n", yparser.GnmiPath2XPath(upd.GetPath(), true), upd.GetVal())
			}
		*/
		// retrieve the config that will be applied to the device
		if _, err := r.device.UpdateGnmi(ctx, updates); err != nil {
			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.ResourceExhausted:
					log.Debug("gnmi update failed exhausted", "error", err)
					if err := r.setExhausted(60); err != nil {
						return err
					}
					// we return and keep the status as is since we can retry once the device is back in normal state
					return nil
				}
			}

			for _, upd := range updates {
				log.Debug("gnmi update failed", "Path", yparser.GnmiPath2XPath(upd.GetPath(), true), "Val", upd.GetVal(), "Error", err)
			}
			// set resource status to failed
			for _, resource := range updResources {
				if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Failed); err != nil {
					return err
				}
			}
		}

		// set resource status to success
		for _, resource := range updResources {
			if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Success); err != nil {
				return err
			}
		}
	}

	// set reconcile flag to false to avoid a new reconciliation if there is no new work
	if err := r.setUpdateStatus(false); err != nil {
		return err
	}
	return nil
}

func (r *reconciler) PrintResourceList(idx int) error {
	resourceListRaw, err := r.getResourceListRaw()
	if err != nil {
		return err
	}
	r.log.Debug("resourceList", "idx", idx, "raw", resourceListRaw)
	return nil
}
