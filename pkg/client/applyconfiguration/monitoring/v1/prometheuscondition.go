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

package v1

import (
	v1 "github.com/rhobs/obo-prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrometheusConditionApplyConfiguration represents an declarative configuration of the PrometheusCondition type for use
// with apply.
type PrometheusConditionApplyConfiguration struct {
	Type               *v1.PrometheusConditionType   `json:"type,omitempty"`
	Status             *v1.PrometheusConditionStatus `json:"status,omitempty"`
	LastTransitionTime *metav1.Time                  `json:"lastTransitionTime,omitempty"`
	Reason             *string                       `json:"reason,omitempty"`
	Message            *string                       `json:"message,omitempty"`
	ObservedGeneration *int64                        `json:"observedGeneration,omitempty"`
}

// PrometheusConditionApplyConfiguration constructs an declarative configuration of the PrometheusCondition type for use with
// apply.
func PrometheusCondition() *PrometheusConditionApplyConfiguration {
	return &PrometheusConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *PrometheusConditionApplyConfiguration) WithType(value v1.PrometheusConditionType) *PrometheusConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *PrometheusConditionApplyConfiguration) WithStatus(value v1.PrometheusConditionStatus) *PrometheusConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *PrometheusConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *PrometheusConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *PrometheusConditionApplyConfiguration) WithReason(value string) *PrometheusConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *PrometheusConditionApplyConfiguration) WithMessage(value string) *PrometheusConditionApplyConfiguration {
	b.Message = &value
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *PrometheusConditionApplyConfiguration) WithObservedGeneration(value int64) *PrometheusConditionApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}
