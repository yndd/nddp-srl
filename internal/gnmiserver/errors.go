/*
Copyright 2021 NDDO.

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

package gnmiserver

const (
	// errors
	errStartGRPCServer         = "cannot start GRPC server"
	errCreateTcpListener       = "cannot create TCP listener"
	errGrpcServer              = "cannot serve GRPC server"
	errMaxNbrOfUnaryRPCReached = "max number of Unary RPC reached"
	errMissingPathsInGNMISet   = "missing update/replace/delete path(s) in gnmi set"
	errKeyNotPresentinPath     = "missing key in path"
	errSchema                  = "cannot update Schema"
	errJSONMarshal             = "cannot Marshal json object"
	errJSONUnMarshal           = "canot Unmarshal json object"
	errDependencyCheckFailed   = "cannot create object since the dependency is not present"
	errInvalidPath             = "invalid Path"
	errGetValue                = "cannot get Value from update"
	errGetSubscribe            = "cannot get subscription definition, the subscribe request must contain a subscription definition"
)
