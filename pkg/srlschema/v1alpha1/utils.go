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

package srlschema

import (
	"sort"
	"strings"
)

func getKey(x1 interface{}) string {
	switch x := x1.(type) {
	case map[string]interface{}:
		keys := make(map[string]string)
		for k, v := range x {
			switch vv := v.(type) {
			case string:
				keys[k] = vv
			}
		}
		ssl := toStrings(keys)
		return toString(ssl)
	default:
		return ""
	}
}

func toStrings(keys map[string]string) []string {
	str := make([]string, 0)
	switch len(keys) {
	case 0:
		// No keys, don't do anything.
	case 1:
		// Special case single key lists, append the only value.
		for _, v := range keys {
			str = append(str, v)
		}
	default:
		str = append(str, sortedVals(keys)...)
	}
	return str
}

func sortedVals(m map[string]string) []string {
	// Return deterministic ordering of values from multi-key lists.
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(m))
	for _, k := range ks {
		vs = append(vs, m[k])
	}
	return vs
}

func toString(ssl []string) string {
	str := ""
	for i, s := range ssl {
		if i == 0 {
			str = s
		} else {
			str = strings.Join([]string{str, s}, ".")
		}
	}
	return str
}
