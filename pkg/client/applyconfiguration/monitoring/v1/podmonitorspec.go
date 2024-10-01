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
	monitoringv1 "github.com/rhobs/obo-prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// PodMonitorSpecApplyConfiguration represents a declarative configuration of the PodMonitorSpec type for use
// with apply.
type PodMonitorSpecApplyConfiguration struct {
	JobLabel              *string                                 `json:"jobLabel,omitempty"`
	PodTargetLabels       []string                                `json:"podTargetLabels,omitempty"`
	PodMetricsEndpoints   []PodMetricsEndpointApplyConfiguration  `json:"podMetricsEndpoints,omitempty"`
	Selector              *metav1.LabelSelectorApplyConfiguration `json:"selector,omitempty"`
	NamespaceSelector     *NamespaceSelectorApplyConfiguration    `json:"namespaceSelector,omitempty"`
	SampleLimit           *uint64                                 `json:"sampleLimit,omitempty"`
	TargetLimit           *uint64                                 `json:"targetLimit,omitempty"`
	ScrapeProtocols       []monitoringv1.ScrapeProtocol           `json:"scrapeProtocols,omitempty"`
	LabelLimit            *uint64                                 `json:"labelLimit,omitempty"`
	LabelNameLengthLimit  *uint64                                 `json:"labelNameLengthLimit,omitempty"`
	LabelValueLengthLimit *uint64                                 `json:"labelValueLengthLimit,omitempty"`
	KeepDroppedTargets    *uint64                                 `json:"keepDroppedTargets,omitempty"`
	AttachMetadata        *AttachMetadataApplyConfiguration       `json:"attachMetadata,omitempty"`
	ScrapeClassName       *string                                 `json:"scrapeClass,omitempty"`
	BodySizeLimit         *monitoringv1.ByteSize                  `json:"bodySizeLimit,omitempty"`
}

// PodMonitorSpecApplyConfiguration constructs a declarative configuration of the PodMonitorSpec type for use with
// apply.
func PodMonitorSpec() *PodMonitorSpecApplyConfiguration {
	return &PodMonitorSpecApplyConfiguration{}
}

// WithJobLabel sets the JobLabel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JobLabel field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithJobLabel(value string) *PodMonitorSpecApplyConfiguration {
	b.JobLabel = &value
	return b
}

// WithPodTargetLabels adds the given value to the PodTargetLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PodTargetLabels field.
func (b *PodMonitorSpecApplyConfiguration) WithPodTargetLabels(values ...string) *PodMonitorSpecApplyConfiguration {
	for i := range values {
		b.PodTargetLabels = append(b.PodTargetLabels, values[i])
	}
	return b
}

// WithPodMetricsEndpoints adds the given value to the PodMetricsEndpoints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PodMetricsEndpoints field.
func (b *PodMonitorSpecApplyConfiguration) WithPodMetricsEndpoints(values ...*PodMetricsEndpointApplyConfiguration) *PodMonitorSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPodMetricsEndpoints")
		}
		b.PodMetricsEndpoints = append(b.PodMetricsEndpoints, *values[i])
	}
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithSelector(value *metav1.LabelSelectorApplyConfiguration) *PodMonitorSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithNamespaceSelector sets the NamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NamespaceSelector field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithNamespaceSelector(value *NamespaceSelectorApplyConfiguration) *PodMonitorSpecApplyConfiguration {
	b.NamespaceSelector = value
	return b
}

// WithSampleLimit sets the SampleLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SampleLimit field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithSampleLimit(value uint64) *PodMonitorSpecApplyConfiguration {
	b.SampleLimit = &value
	return b
}

// WithTargetLimit sets the TargetLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetLimit field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithTargetLimit(value uint64) *PodMonitorSpecApplyConfiguration {
	b.TargetLimit = &value
	return b
}

// WithScrapeProtocols adds the given value to the ScrapeProtocols field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ScrapeProtocols field.
func (b *PodMonitorSpecApplyConfiguration) WithScrapeProtocols(values ...monitoringv1.ScrapeProtocol) *PodMonitorSpecApplyConfiguration {
	for i := range values {
		b.ScrapeProtocols = append(b.ScrapeProtocols, values[i])
	}
	return b
}

// WithLabelLimit sets the LabelLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelLimit field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithLabelLimit(value uint64) *PodMonitorSpecApplyConfiguration {
	b.LabelLimit = &value
	return b
}

// WithLabelNameLengthLimit sets the LabelNameLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelNameLengthLimit field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithLabelNameLengthLimit(value uint64) *PodMonitorSpecApplyConfiguration {
	b.LabelNameLengthLimit = &value
	return b
}

// WithLabelValueLengthLimit sets the LabelValueLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelValueLengthLimit field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithLabelValueLengthLimit(value uint64) *PodMonitorSpecApplyConfiguration {
	b.LabelValueLengthLimit = &value
	return b
}

// WithKeepDroppedTargets sets the KeepDroppedTargets field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeepDroppedTargets field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithKeepDroppedTargets(value uint64) *PodMonitorSpecApplyConfiguration {
	b.KeepDroppedTargets = &value
	return b
}

// WithAttachMetadata sets the AttachMetadata field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AttachMetadata field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithAttachMetadata(value *AttachMetadataApplyConfiguration) *PodMonitorSpecApplyConfiguration {
	b.AttachMetadata = value
	return b
}

// WithScrapeClassName sets the ScrapeClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeClassName field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithScrapeClassName(value string) *PodMonitorSpecApplyConfiguration {
	b.ScrapeClassName = &value
	return b
}

// WithBodySizeLimit sets the BodySizeLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BodySizeLimit field is set to the value of the last call.
func (b *PodMonitorSpecApplyConfiguration) WithBodySizeLimit(value monitoringv1.ByteSize) *PodMonitorSpecApplyConfiguration {
	b.BodySizeLimit = &value
	return b
}
