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

	"github.com/google/gnxi/utils/xpath"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndd-yang/pkg/yparser"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *reconciler) ReconcileTransaction(ctx context.Context, t *systemv1alpha1.Transaction) error {
	log := r.log.WithValues("target", r.target.Config.Name, "address", r.target.Config.Address)
	log.Debug("reconciling transaction device config", "transaction", t.Name)

	switch t.Action {
	case systemv1alpha1.E_TransactionAction_Delete:
		delResources := make([]*systemv1alpha1.Gvk, 0)
		delPaths := make([]*gnmi.Path, 0)
		for _, gvkName := range t.Gvk {
			resource, err := r.getResource(gvkName)
			if err != nil {
				return err
			}
			delResources = append(delResources, resource)
			// append delete path to the list
			path, err := xpath.ToGNMIPath(resource.Rootpath)
			if err != nil {
				return err
			}
			delPaths = append(delPaths, path)

		}

		// we delete all the paths on the device in a single transaction
		// we use the change notification to update the cache
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
				// we dont set transaction status to failed otherwise it will never be deleted
				/*
					if err := r.updateTransactionStatus(t.Name, systemv1alpha1.E_TransactionStatus_Failed); err != nil {
						return err
					}
				*/
			} else {
				//log.Debug("gnmi delete success", "Paths", delPaths)

				// delete resources from the system cache
				for _, resource := range delResources {
					r.deleteResource(resource.Name)
				}
				// delete the transaction from the system cache
				r.deleteTransaction(t.Name)
			}
		}
	case systemv1alpha1.E_TransactionAction_Create:
		updResources := make([]*systemv1alpha1.Gvk, 0)
		updates := make([]*gnmi.Update, 0)
		for _, gvkName := range t.Gvk {
			resource, err := r.getResource(gvkName)
			if err != nil {
				return err
			}
			// append to the resource list of resources needed an update
			updResources = append(updResources, resource)
			// get updates which are part of the transaction
			resUpdate, err := r.getUpdates(resource)
			if err != nil {
				return err
			}
			updates = append(updates, resUpdate)
		}

		if len(updates) > 0 {
			// debug
			/*
				for _, upd := range updates {
					fmt.Printf("Updates: path: %s, data: %v\n", yparser.GnmiPath2XPath(upd.GetPath(), true), upd.GetVal())
				}
			*/
			// update the device
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
				// set transaction status to failed
				if err := r.updateTransactionStatus(t.Name, systemv1alpha1.E_TransactionStatus_Failed); err != nil {
					return err
				}
			}

			// set resource status to success
			for _, resource := range updResources {
				if err := r.updateResourceStatus(resource.Name, systemv1alpha1.E_GvkStatus_Success); err != nil {
					return err
				}
			}
			// set transaction status to success
			if err := r.updateTransactionStatus(t.Name, systemv1alpha1.E_TransactionStatus_Success); err != nil {
				return err
			}
		}
	}

	return nil
}
