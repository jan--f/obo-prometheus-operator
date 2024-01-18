// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/rhobs/obo-prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1 "github.com/rhobs/obo-prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
)

// HTTPSDConfigApplyConfiguration represents an declarative configuration of the HTTPSDConfig type for use
// with apply.
type HTTPSDConfigApplyConfiguration struct {
	URL                           *string                                           `json:"url,omitempty"`
	RefreshInterval               *v1.Duration                                      `json:"refreshInterval,omitempty"`
	BasicAuth                     *monitoringv1.BasicAuthApplyConfiguration         `json:"basicAuth,omitempty"`
	Authorization                 *monitoringv1.SafeAuthorizationApplyConfiguration `json:"authorization,omitempty"`
	TLSConfig                     *monitoringv1.SafeTLSConfigApplyConfiguration     `json:"tlsConfig,omitempty"`
	ProxyConfigApplyConfiguration `json:",inline"`
}

// HTTPSDConfigApplyConfiguration constructs an declarative configuration of the HTTPSDConfig type for use with
// apply.
func HTTPSDConfig() *HTTPSDConfigApplyConfiguration {
	return &HTTPSDConfigApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *HTTPSDConfigApplyConfiguration) WithURL(value string) *HTTPSDConfigApplyConfiguration {
	b.URL = &value
	return b
}

// WithRefreshInterval sets the RefreshInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RefreshInterval field is set to the value of the last call.
func (b *HTTPSDConfigApplyConfiguration) WithRefreshInterval(value v1.Duration) *HTTPSDConfigApplyConfiguration {
	b.RefreshInterval = &value
	return b
}

// WithBasicAuth sets the BasicAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BasicAuth field is set to the value of the last call.
func (b *HTTPSDConfigApplyConfiguration) WithBasicAuth(value *monitoringv1.BasicAuthApplyConfiguration) *HTTPSDConfigApplyConfiguration {
	b.BasicAuth = value
	return b
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *HTTPSDConfigApplyConfiguration) WithAuthorization(value *monitoringv1.SafeAuthorizationApplyConfiguration) *HTTPSDConfigApplyConfiguration {
	b.Authorization = value
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *HTTPSDConfigApplyConfiguration) WithTLSConfig(value *monitoringv1.SafeTLSConfigApplyConfiguration) *HTTPSDConfigApplyConfiguration {
	b.TLSConfig = value
	return b
}
