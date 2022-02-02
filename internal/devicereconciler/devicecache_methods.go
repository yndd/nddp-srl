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
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/gnxi/utils/xpath"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-srl/internal/shared"
	systemv1alpha1 "github.com/yndd/nddp-system/apis/system/v1alpha1"
)

func (r *reconciler) deletePathsFromCache(delPaths []*gnmi.Path) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	n := &gnmi.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix:    &gnmi.Path{Target: crDeviceName},
		Delete:    delPaths,
	}

	if err := r.cache.GnmiUpdate(crDeviceName, n); err != nil {
		return errors.Wrap(err, "cache update failed")
	}
	return nil
}

func (r *reconciler) copyRunning2Candidate() error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crCandidateDeviceName := shared.GetCrCandidateDeviceName(crDeviceName)

	ns, err := r.cache.QueryAll(crDeviceName,
		&gnmi.Path{Target: crDeviceName},
		&gnmi.Path{})
	if err != nil {
		return err
	}

	for _, n := range ns {
		newNotification := &gnmi.Notification{
			Timestamp: time.Now().UnixNano(),
			Prefix:    &gnmi.Path{Target: crCandidateDeviceName},
			Update:    n.GetUpdate(),
		}
		if err := r.cache.GnmiUpdate(crCandidateDeviceName, newNotification); err != nil {
			return err
		}
	}
	return nil
}

func (r *reconciler) copyCandidate2Running() error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crCandidateDeviceName := shared.GetCrCandidateDeviceName(crDeviceName)

	ns, err := r.cache.QueryAll(
		crCandidateDeviceName,
		&gnmi.Path{Target: crCandidateDeviceName},
		&gnmi.Path{})
	if err != nil {
		return err
	}

	for _, n := range ns {
		newNotification := &gnmi.Notification{
			Timestamp: time.Now().UnixNano(),
			Prefix:    &gnmi.Path{Target: crDeviceName},
			Update:    n.GetUpdate(),
		}
		/*
			for _, u := range newNotification.GetUpdate() {
				r.log.Debug("copyCandidate2Running", "path", yparser.GnmiPath2XPath(u.GetPath(), true), "val", u.GetVal())
			}
		*/
		if err := r.cache.GnmiUpdate(crDeviceName, newNotification); err != nil {
			return err
		}
	}
	return nil
}

func (r *reconciler) updateCandidate(resource *systemv1alpha1.Gvk) error {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)
	crCandidateDeviceName := shared.GetCrCandidateDeviceName(crDeviceName)

	ns, err := r.cache.QueryAll(
		crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "gvk", Key: map[string]string{"name": *resource.Name}},
				{Name: "data"},
			},
		})
	if err != nil {
		return err
	}

	for _, n := range ns {
		// create a new update where the path
		for _, u := range n.GetUpdate() {
			path := yparser.DeepCopyGnmiPath(u.GetPath())
			update := &gnmi.Update{
				Path: &gnmi.Path{
					// remove gvk and data from the path elem
					Elem: path.GetElem()[2:],
				},
				Val: u.GetVal(),
			}
			// create new notification and set the target to the candidate device name iso the system
			newNotification := &gnmi.Notification{
				Timestamp: time.Now().UnixNano(),
				Prefix:    &gnmi.Path{Target: crCandidateDeviceName},
				Update:    []*gnmi.Update{update},
			}
			r.log.Debug("updateCandidate notification",
				"path", yparser.GnmiPath2XPath(newNotification.GetUpdate()[0].GetPath(), true),
				"val", newNotification.GetUpdate()[0].GetVal(),
			)
			if err := r.cache.GnmiUpdate(crCandidateDeviceName, newNotification); err != nil {
				return errors.Wrap(err, "cache update failed")
			}
		}
	}
	return nil
}

func (r *reconciler) getCandidateUpdate() ([]*gnmi.Update, bool, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crCandidateDeviceName := shared.GetCrCandidateDeviceName(crDeviceName)

	candidatecacheData, err := r.cache.GetJson(crCandidateDeviceName,
		&gnmi.Path{Target: crCandidateDeviceName},
		&gnmi.Path{},
		r.deviceSchema)
	if err != nil {
		return nil, false, err
	}

	r.log.Debug("getCandidateUpdate", "config", candidatecacheData)

	d, err := json.Marshal(candidatecacheData)
	if err != nil {
		return nil, false, err
	}
	// if the data is empty, there is no need for an update
	if string(d) == "null" {
		return nil, false, nil
	}
	updates := make([]*gnmi.Update, 0)
	updates = append(updates, &gnmi.Update{
		Path: &gnmi.Path{}, // this is the roo path
		Val: &gnmi.TypedValue{
			Value: &gnmi.TypedValue_JsonIetfVal{
				JsonIetfVal: bytes.Trim(d, " \r\n\t"),
			},
		},
	})

	return updates, true, nil
}

func (r *reconciler) getCandidateConfig() (interface{}, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crCandidateDeviceName := shared.GetCrCandidateDeviceName(crDeviceName)

	candidatecacheData, err := r.cache.GetJson(crCandidateDeviceName,
		&gnmi.Path{Target: crCandidateDeviceName},
		&gnmi.Path{},
		r.deviceSchema)
	if err != nil {
		return nil, err
	}

	return candidatecacheData, nil
}

func (r *reconciler) initializeCandidateCache() {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crCandidateDeviceName := shared.GetCrCandidateDeviceName(crDeviceName)

	if r.cache.GetCache().HasTarget(crCandidateDeviceName) {
		r.cache.GetCache().Remove(crCandidateDeviceName)
	}
	r.cache.GetCache().Add(crCandidateDeviceName)
}

func (r *reconciler) getSpecdata(resource *systemv1alpha1.Gvk) (interface{}, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)
	crSystemDeviceName := shared.GetCrSystemDeviceName(crDeviceName)

	x1, err := r.cache.GetJson(
		crSystemDeviceName,
		&gnmi.Path{Target: crSystemDeviceName},
		&gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "gvk", Key: map[string]string{"name": *resource.Name}},
			},
		},
		r.nddpSchema)
	if err != nil {
		return nil, err
	}
	// remove the rootPath data
	rootPath, err := xpath.ToGNMIPath(*resource.Rootpath)
	if err != nil {
		return nil, err
	}
	switch x := x1.(type) {
	case map[string]interface{}:
		x1 = x["data"]
		x1 = getDataFromRootPath(rootPath, x1)
	}
	return x1, nil
}

func getDataFromRootPath(path *gnmi.Path, x1 interface{}) interface{} {
	fmt.Printf("getDataFromRootPath: %s, data: %v\n", yparser.GnmiPath2XPath(path, true), x1)
	p := yparser.DeepCopyGnmiPath(path)
	if len(p.GetElem()) > 0 {
		hasKey := false
		if len(p.GetElem()[0].GetKey()) > 0 {
			hasKey = true
		}
		switch x := x1.(type) {
		case map[string]interface{}:
			for k, v := range x {
				if k == p.GetElem()[0].GetName() {
					// when the spec data rootpath last element has a key
					// we need to return the first element of the list
					if hasKey {
						switch x := v.(type) {
						case []interface{}:
							x1 = x[0]
						}
					} else {
						x1 = x[p.GetElem()[0].GetName()]
					}
				}
			}

		}
		p.Elem = p.Elem[1:]
		x1 = getDataFromRootPath(p, x1)
	}
	return x1
}

func (r *reconciler) getCachedata(resource *systemv1alpha1.Gvk) (interface{}, error) {
	crDeviceName := shared.GetCrDeviceName(r.namespace, r.target.Config.Name)

	rootPath, err := xpath.ToGNMIPath(*resource.Rootpath)
	if err != nil {
		return nil, err
	}
	x2, err := r.cache.GetJson(
		crDeviceName,
		&gnmi.Path{Target: crDeviceName},
		rootPath,
		r.deviceSchema)
	if err != nil {
		return nil, err
	}
	return x2, nil
}

func (r *reconciler) processUpdates(resource *systemv1alpha1.Gvk) ([]*gnmi.Path, []*gnmi.Update, error) {
	rootPath, err := xpath.ToGNMIPath(*resource.Rootpath)
	if err != nil {
		return nil, nil, err
	}
	x1, err := r.getSpecdata(resource)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("processUpdates x1 data %v\n", x1)
	x2, err := r.getCachedata(resource)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("processUpdates x2 data %v\n", x2)

	// remove hierarchical elements
	hierPaths := r.deviceSchema.GetHierarchicalResourcesLocal(true, rootPath, &gnmi.Path{}, make([]*gnmi.Path, 0))
	// remove hierarchical resource elements from the data to be able to compare the gnmi response
	// with the k8s Spec
	switch x := x2.(type) {
	case map[string]interface{}:
		for _, hierPath := range hierPaths {
			x2 = removeHierarchicalResourceData(x, hierPath)
		}
	}

	updatesx1, err := yparser.GetUpdatesFromJSON(rootPath, x1, r.deviceSchema)
	if err != nil {
		return nil, nil, err
	}
	updatesx2, err := yparser.GetUpdatesFromJSON(rootPath, x2, r.deviceSchema)
	if err != nil {
		return nil, nil, err
	}

	return yparser.FindResourceDelta(updatesx1, updatesx2)
}

func removeHierarchicalResourceData(x map[string]interface{}, hierPath *gnmi.Path) interface{} {
	// this is the last pathElem of the hierarchical path, which is to be deleted
	if len(hierPath.GetElem()) == 1 {
		delete(x, hierPath.GetElem()[0].GetName())
	} else {
		// there is more pathElem in the hierachical Path
		if xx, ok := x[hierPath.GetElem()[0].GetName()]; ok {
			switch x1 := xx.(type) {
			case map[string]interface{}:
				removeHierarchicalResourceData(x1, &gnmi.Path{Elem: hierPath.GetElem()[1:]})
			case []interface{}:
				for _, xxx := range x1 {
					switch x2 := xxx.(type) {
					case map[string]interface{}:
						removeHierarchicalResourceData(x2, &gnmi.Path{Elem: hierPath.GetElem()[1:]})
					}
				}
			default:
				// it can be that no data is present, so we ignore this
			}
		}
	}

	return x
}
