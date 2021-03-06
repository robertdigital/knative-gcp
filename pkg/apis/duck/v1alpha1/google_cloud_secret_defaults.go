/*
Copyright 2019 Google LLC

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	defaultSecretName = "google-cloud-key"
	defaultSecretKey  = "key.json"
)

// DefaultGoogleCloudSecretSelector is the default secret selector used to load
// the creds for the objects that will auth with Google Cloud.
func DefaultGoogleCloudSecretSelector() *corev1.SecretKeySelector {
	return &corev1.SecretKeySelector{
		LocalObjectReference: corev1.LocalObjectReference{
			Name: defaultSecretName,
		},
		Key: defaultSecretKey,
	}
}
