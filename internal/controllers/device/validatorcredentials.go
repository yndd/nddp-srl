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

package device

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	ndddvrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	"github.com/yndd/ndd-runtime/pkg/resource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	// Errors
	errEmptyTargetSecretReference   = "empty target secret reference"
	errCredentialSecretDoesNotExist = "credential secret does not exist"
	errEmptyTargetAddress           = "empty target address"
	errMissingUsername              = "missing username in credentials"
	errMissingPassword              = "missing password in credentials"
)

// Credentials holds the information for authenticating with the Server.
type Credentials struct {
	Username string
	Password string
}

func (r *Reconciler) validateCredentials(ctx context.Context, nn ndddvrv1.Nn) (creds *Credentials, err error) {
	log := r.log.WithValues("namespace", nn.GetNamespace(), "credentialsName", nn.GetTargetCredentialsName(), "targetAddress", nn.GetTargetAddress())
	log.Debug("Credentials Validation")
	// Retrieve the secret from Kubernetes for this network node

	credsSecret, err := r.getSecret(ctx, nn)
	if err != nil {
		return nil, err
	}

	// Check if address is defined on the network node
	if nn.GetTargetAddress() == "" {
		return nil, errors.New(errEmptyTargetAddress)
	}

	creds = &Credentials{
		Username: strings.TrimSuffix(string(credsSecret.Data["username"]), "\n"),
		Password: strings.TrimSuffix(string(credsSecret.Data["password"]), "\n"),
	}

	log.Debug("Credentials", "creds", creds)

	if creds.Username == "" {
		return nil, errors.New(errMissingUsername)
	}
	if creds.Password == "" {
		return nil, errors.New(errMissingPassword)
	}

	return creds, nil
}

// Retrieve the secret containing the credentials for talking to the Network Node.
func (r *Reconciler) getSecret(ctx context.Context, nn ndddvrv1.Nn) (credsSecret *corev1.Secret, err error) {
	// if namespace
	// check if credentialName is specified
	if nn.GetTargetCredentialsName() == "" {
		return nil, errors.New(errEmptyTargetSecretReference)
	}
	namespace := nn.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}

	// check if credential secret exists
	secretKey := types.NamespacedName{
		Name:      nn.GetTargetCredentialsName(),
		Namespace: namespace,
	}
	credsSecret = &corev1.Secret{}
	if err := r.client.Get(ctx, secretKey, credsSecret); resource.IgnoreNotFound(err) != nil {
		return nil, errors.Wrap(err, errCredentialSecretDoesNotExist)
	}
	return credsSecret, nil
}
