/*
Copyright 2022 NDD.

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

package srl

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-yang/pkg/yentry"
	"github.com/yndd/ndd-yang/pkg/yparser"
	"github.com/yndd/nddp-system/pkg/gvkresource"
)

type adder interface {
	Add(item interface{})
}

type observe struct {
	hasData bool
	delta   bool
	//deletes []*gnmi.Path
	//updates []*gnmi.Update
}

// processObserve
// 0. marshal/unmarshal data
// 1. check if resource has data
// 2. remove parent hierarchical elements from spec
// 3. remove resource hierarchical elements from gnmi response
// 4. transform the data in gnmi to process the delta
// 5. find the resource delta: updates and/or deletes in gnmi
func processObserve(rootPath *gnmi.Path, hierPaths []*gnmi.Path, specData interface{}, resp *gnmi.GetResponse, deviceSchema *yentry.Entry) (*observe, error) {

	// prepare the input data to compare against the response data
	x1, err := processSpecData(rootPath, specData)
	if err != nil {
		//return false, nil, nil, nil, err
		return nil, err
	}

	// validate gnmi resp information
	var x2 interface{}
	if len(resp.GetNotification()) == 0 {
		return &observe{hasData: false}, nil
	}
	if len(resp.GetNotification()) != 0 && len(resp.GetNotification()[0].GetUpdate()) != 0 {
		// get value from gnmi get response
		x2, err = yparser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
		if err != nil {
			return nil, errors.Wrap(err, errJSONMarshal)
		}

		//fmt.Printf("processObserve: raw x2: %v\n", x2)

		switch x2.(type) {
		case nil:
			// resource has no data
			return &observe{hasData: false}, nil

		}
	}
	// resource has data

	//fmt.Printf("processObserve rootPath %s\n", yparser.GnmiPath2XPath(rootPath, true))
	// for the Spec data we remove the first element which is aligned with the last element of the rootPath
	// gnmi does not return this information hence to compare the spec data with the gnmi resp data we need to remove
	// the first element from the Spec
	switch x := x1.(type) {
	case map[string]interface{}:
		x1 = x[rootPath.GetElem()[len(rootPath.GetElem())-1].GetName()]
		fmt.Printf("processObserve x1 data %v\n", x1)
	}
	// the gnmi response already comes without the last element in the return data
	//fmt.Printf("processObserve x2 data %v\n", x2)

	// remove hierarchical resource elements from the data to be able to compare the gnmi response
	// with the k8s Spec
	switch x := x2.(type) {
	case map[string]interface{}:
		for _, hierPath := range hierPaths {
			x2 = removeHierarchicalResourceData(x, hierPath)
		}
	}

	fmt.Printf("processObserve x2 data %v\n", x2)

	// returns the data in gnmi Updates per container/list/leaflists/...
	updatesx1, err := yparser.GetUpdatesFromJSON(rootPath, x1, deviceSchema)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}
	/*
		for _, update := range updatesx1 {
			log.Debug("Observe Fine Grane Updates X1", "Path", yparser.GnmiPath2XPath(update.Path, true), "Value", update.GetVal())
		}
	*/

	updatesx2, err := yparser.GetUpdatesFromJSON(rootPath, x2, deviceSchema)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}
	/*
		for _, update := range updatesx2 {
			fmt.Printf("processObserve x2 update, Path %s, Value %v\n", yparser.GnmiPath2XPath(update.Path, true), update.GetVal())
		}
	*/
	// returns the deletes and updates that need to be performed to bring the spec object back to the desired state
	deletes, updates, err := yparser.FindResourceDelta(updatesx1, updatesx2)
	// check for defaults:
	delta := false // means all ok
	for _, u := range updates {
		delta, err = validateDefaults(u, rootPath, x2, deviceSchema)
		if err != nil {
			return nil, err
		}
		/*
			defVal := deviceSchema.GetPathDefault(u.GetPath())

			v, err := yparser.GetValue(u.GetVal())
			if err != nil {
				return nil, err
			}

			var updVal string
			switch val := v.(type) {
			case bool:
				updVal = strconv.FormatBool(val)
			case uint32:
				updVal = strconv.Itoa(int(val))
			case uint8:
				updVal = strconv.Itoa(int(val))
			case uint16:
				updVal = strconv.Itoa(int(val))
			case float64, float32:
				updVal = fmt.Sprintf("%.0f", val)
			case string:
				updVal = val
			case map[string]interface{}:
				delta = true
			case []interface{}:
				delta = true
			}
			fmt.Printf("processObserve delta check: %v, default: %s, update value: %v, path:%s\n", v, defVal, updVal, yparser.GnmiPath2XPath(u.GetPath(), true))

			// only perform the check on defaults if the data does not exist
			if !dataExists(u.GetPath().GetElem()[len(rootPath.GetElem()):], x2) {
				if defVal != "" && updVal != defVal {
					delta = true
					fmt.Printf("processObserve default check: path %s, deviceschema default: %s, update value: %v\n",
						yparser.GnmiPath2XPath(u.GetPath(), true),
						defVal,
						updVal)
				}
			} else {
				delta = true
			}
		*/

	}

	if len(deletes) != 0 {
		delta = true
	}
	return &observe{
		hasData: true,
		delta:   delta,
		//deletes: deletes,
		//updates: updates,
		//data:    b,
	}, err
}

func validateDefaults(u *gnmi.Update, rootPath *gnmi.Path, x2 interface{}, deviceSchema *yentry.Entry) (bool, error) {
	v, err := yparser.GetValue(u.GetVal())
	if err != nil {
		return false, err
	}

	// check if the value contains a map, if so validate the defaults per element in the map
	delta := false
	switch val := v.(type) {
	case map[string]interface{}:
		for k, vv := range val {
			path := yparser.DeepCopyGnmiPath(u.GetPath())
			path.Elem = append(path.GetElem(), &gnmi.PathElem{Name: k})
			delta := validateDefault(path, rootPath, vv, x2, deviceSchema)
			if delta { // if we find a diff we can already return
				return delta, nil
			}
		}
	default:
		delta = validateDefault(u.GetPath(), rootPath, v, x2, deviceSchema)
	}
	return delta, nil
}

func validateDefault(p, rootPath *gnmi.Path, v, x2 interface{}, deviceSchema *yentry.Entry) bool {
	defVal := deviceSchema.GetPathDefault(p)
	var updVal string
	switch val := v.(type) {
	case bool:
		updVal = strconv.FormatBool(val)
	case uint32:
		updVal = strconv.Itoa(int(val))
	case uint8:
		updVal = strconv.Itoa(int(val))
	case uint16:
		updVal = strconv.Itoa(int(val))
	case float64, float32:
		updVal = fmt.Sprintf("%.0f", val)
	case string:
		updVal = val
	}
	fmt.Printf("processObserve delta check: %v, default: %s, update value: %v, path:%s\n", v, defVal, updVal, yparser.GnmiPath2XPath(p, true))

	delta := false // means all ok
	// only perform the check on defaults if the data does not exist
	if !dataExists(p.GetElem()[len(rootPath.GetElem()):], x2) {
		if defVal != "" && updVal != defVal {
			delta = true
			fmt.Printf("processObserve default check: path %s, deviceschema default: %s, update value: %v\n",
				yparser.GnmiPath2XPath(p, true),
				defVal,
				updVal)
		}
	} else {
		delta = true
	}
	return delta
}

// given the updates are per container i dont expect we would ever come here
func dataExists(pe []*gnmi.PathElem, x interface{}) bool {
	fmt.Printf("dataExists: PathElem: %s\n", pe[0].GetName())
	if len(pe[0].Key) == 0 {
		switch xx := x.(type) {
		case map[string]interface{}:
			if xxx, ok := xx[pe[0].GetName()]; ok {
				if len(pe) > 1 {
					return dataExists(pe[1:], xxx)
				} else {
					return true
				}
			} else {
				return false
			}
		}
	}
	return false
}

// returns the update using group version kind namespace name
func processUpdateK8sResource(mg resource.Managed, rootPath *gnmi.Path, nddpSchema *yentry.Entry) ([]*gnmi.Update, error) {
	gvkData := gvkresource.GetK8sResourceUpdate(mg, rootPath)
	gvkPath := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: "gvk", Key: map[string]string{"name": *gvkData.Name}},
		},
	}
	gvkd, err := processGvkData(*gvkData)
	if err != nil {
		return nil, err
	}

	return yparser.GetUpdatesFromJSON(gvkPath, gvkd, nddpSchema)
}

// returns the delete using group version kind namespace name
func processDeleteK8sResource(mg resource.Managed, rootPath *gnmi.Path, nddpSchema *yentry.Entry) ([]*gnmi.Update, error) {
	gvkData := gvkresource.GetK8sResourceDelete(mg, rootPath)
	gvkPath := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: "gvk", Key: map[string]string{"name": *gvkData.Name}},
		},
	}
	gvkd, err := processGvkData(*gvkData)
	if err != nil {
		return nil, err
	}

	return yparser.GetUpdatesFromJSON(gvkPath, gvkd, nddpSchema)
}

// returns the create using group version kind namespace name
func processCreateK8sResource(mg resource.Managed, rootPath *gnmi.Path, nddpSchema *yentry.Entry) ([]*gnmi.Update, error) {
	gvkData := gvkresource.GetK8sResourceCreate(mg, rootPath)
	gvkPath := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: "gvk", Key: map[string]string{"name": *gvkData.Name}},
		},
	}
	gvkd, err := processGvkData(*gvkData)
	if err != nil {
		return nil, err
	}

	return yparser.GetUpdatesFromJSON(gvkPath, gvkd, nddpSchema)
}

//process Spec data marshals the data and remove the prent hierarchical keys
func processGvkData(gvkData interface{}) (interface{}, error) {
	// prepare the input data to compare against the response data
	d, err := json.Marshal(gvkData)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}
	var x1 interface{}
	if err := json.Unmarshal(d, &x1); err != nil {
		return nil, errors.Wrap(err, errJSONUnMarshal)
	}
	return x1, nil
}

// processCreate
// o. marshal/unmarshal data
// 1. transform the spec data to gnmi updates
// 2. merge gvk and spec data
func processCreateK8s(mg resource.Managed, rootPath *gnmi.Path, specData interface{}, deviceSchema, nddpSchema *yentry.Entry) ([]*gnmi.Update, error) {

	// prepare gvk resource data
	gvkUpdates, err := processCreateK8sResource(mg, rootPath, nddpSchema)
	if err != nil {
		return nil, err
	}
	fmt.Printf("gvkUpdates: %v\n", gvkUpdates)

	for _, gvku := range gvkUpdates {
		fmt.Printf("gvk update :%s, data: %v \n", yparser.GnmiPath2XPath(gvku.Path, true), gvku.GetVal())
	}

	// prepare the input data to compare against the response data
	x1, err := processSpecData(rootPath, specData)
	if err != nil {
		return nil, err
	}

	// we remove the first element which is aligned with the last element of the rootPath
	// gnmi does not return this information hence to compare the spec data with the gnmi resp data we need to remove
	// the first element from the Spec
	switch x := x1.(type) {
	case map[string]interface{}:
		x1 := x[rootPath.GetElem()[len(rootPath.GetElem())-1].GetName()]
		fmt.Printf("processCreate data %v, rootPath: %s\n", x1, yparser.GnmiPath2XPath(rootPath, true))
		gnmiUpdate, err := yparser.GetUpdatesFromJSON(rootPath, x1, deviceSchema)
		if err != nil {
			return nil, err
		}

		gvkName := gvkresource.GetGvkName(mg)

		for _, u := range gnmiUpdate {
			u.GetPath().Elem = append(
				[]*gnmi.PathElem{
					{Name: "gvk", Key: map[string]string{"name": gvkName}},
					{Name: "data"},
				}, u.GetPath().GetElem()...,
			)
		}
		gvkUpdates = append(gvkUpdates, gnmiUpdate...)
		return gvkUpdates, nil
	}
	return nil, errors.New("wrong data structure")

}

// processCreate
// o. marshal/unmarshal data
// 1. transform the spec data to gnmi updates
// 2. merge gvk and spec data
func processUpdateK8s(mg resource.Managed, rootPath *gnmi.Path, specData interface{}, deviceSchema, nddpSchema *yentry.Entry) ([]*gnmi.Update, error) {

	// prepare gvk resource data
	gvkUpdates, err := processUpdateK8sResource(mg, rootPath, nddpSchema)
	if err != nil {
		return nil, err
	}

	// prepare the input data to compare against the response data
	x1, err := processSpecData(rootPath, specData)
	if err != nil {
		return nil, err
	}

	// we remove the first element which is aligned with the last element of the rootPath
	// gnmi does not return this information hence to compare the spec data with the gnmi resp data we need to remove
	// the first element from the Spec
	switch x := x1.(type) {
	case map[string]interface{}:
		x1 := x[rootPath.GetElem()[len(rootPath.GetElem())-1].GetName()]
		//fmt.Printf("processCreate data %v\n", x1)
		gnmiUpdate, err := yparser.GetUpdatesFromJSON(rootPath, x1, deviceSchema)
		if err != nil {
			return nil, err
		}

		gvkName := gvkresource.GetGvkName(mg)

		for _, u := range gnmiUpdate {
			u.GetPath().Elem = append(
				[]*gnmi.PathElem{
					{Name: "gvk", Key: map[string]string{"name": gvkName}},
					{Name: "data"},
				}, u.GetPath().GetElem()...,
			)
		}
		gvkUpdates = append(gvkUpdates, gnmiUpdate...)
		return gvkUpdates, nil
	}
	return nil, errors.New("wrong data structure")

}

//process Spec data marshals the data and remove the prent hierarchical keys
func processSpecData(rootPath *gnmi.Path, specData interface{}) (interface{}, error) {
	// prepare the input data to compare against the response data
	d, err := json.Marshal(specData)
	if err != nil {
		return nil, errors.Wrap(err, errJSONMarshal)
	}
	var x1 interface{}
	if err := json.Unmarshal(d, &x1); err != nil {
		return nil, errors.Wrap(err, errJSONUnMarshal)
	}
	// removes the parent hierarchical ids; they are there to define the parent in k8s so
	// we can define the full path in gnmi
	return yparser.RemoveHierIDsFomData(yparser.GetHierIDsFromPath(rootPath), x1), nil
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
