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
package shared

import (
	"strings"
	"time"

	"github.com/karimra/gnmic/types"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-yang/pkg/yentry"
)

const (
	SystemNamespace    = "system"
	CandidateNamespace = "candidate"
)

func GetCrDeviceName(namespace, name string) string {
	if namespace == "" {
		namespace = "default"
	}
	return strings.Join([]string{namespace, name}, ".")
}

func GetCrSystemDeviceName(name string) string {
	return strings.Join([]string{SystemNamespace, name}, ".")
}

func GetCrCandidateDeviceName(name string) string {
	return strings.Join([]string{CandidateNamespace, name}, ".")
}

// A DeviceAction based on the device
type DeviceAction string

// DeviceAction.
const (
	DeviceStart  DeviceAction = "start"
	DeviceStop   DeviceAction = "stop"
	DeviceStatus DeviceAction = "status"
)

// TargetUpdate identifies the update actions on the target
type DeviceUpdate struct {
	Action       DeviceAction
	TargetConfig *types.TargetConfig
	Namespace    string
}

type DeviceResponse struct {
	Error         error
	Exists        bool
	DeviceDetails *ndrv1.DeviceDetails
	TargetConfig  *types.TargetConfig
}

type NddControllerOptions struct {
	Logger                 logging.Logger
	Poll                   time.Duration
	DeviceSchema           *yentry.Entry
	NddpSchema             *yentry.Entry
	Namespace              string
	GnmiAddress            string
	DeviceDriverRequestCh  chan DeviceUpdate
	DeviceDriverResponseCh chan DeviceResponse
}
