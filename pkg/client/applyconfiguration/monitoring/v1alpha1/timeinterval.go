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
	monitoringv1alpha1 "github.com/rhobs/obo-prometheus-operator/pkg/apis/monitoring/v1alpha1"
)

// TimeIntervalApplyConfiguration represents a declarative configuration of the TimeInterval type for use
// with apply.
type TimeIntervalApplyConfiguration struct {
	Times       []TimeRangeApplyConfiguration       `json:"times,omitempty"`
	Weekdays    []monitoringv1alpha1.WeekdayRange   `json:"weekdays,omitempty"`
	DaysOfMonth []DayOfMonthRangeApplyConfiguration `json:"daysOfMonth,omitempty"`
	Months      []monitoringv1alpha1.MonthRange     `json:"months,omitempty"`
	Years       []monitoringv1alpha1.YearRange      `json:"years,omitempty"`
}

// TimeIntervalApplyConfiguration constructs a declarative configuration of the TimeInterval type for use with
// apply.
func TimeInterval() *TimeIntervalApplyConfiguration {
	return &TimeIntervalApplyConfiguration{}
}

// WithTimes adds the given value to the Times field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Times field.
func (b *TimeIntervalApplyConfiguration) WithTimes(values ...*TimeRangeApplyConfiguration) *TimeIntervalApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTimes")
		}
		b.Times = append(b.Times, *values[i])
	}
	return b
}

// WithWeekdays adds the given value to the Weekdays field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Weekdays field.
func (b *TimeIntervalApplyConfiguration) WithWeekdays(values ...monitoringv1alpha1.WeekdayRange) *TimeIntervalApplyConfiguration {
	for i := range values {
		b.Weekdays = append(b.Weekdays, values[i])
	}
	return b
}

// WithDaysOfMonth adds the given value to the DaysOfMonth field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the DaysOfMonth field.
func (b *TimeIntervalApplyConfiguration) WithDaysOfMonth(values ...*DayOfMonthRangeApplyConfiguration) *TimeIntervalApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithDaysOfMonth")
		}
		b.DaysOfMonth = append(b.DaysOfMonth, *values[i])
	}
	return b
}

// WithMonths adds the given value to the Months field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Months field.
func (b *TimeIntervalApplyConfiguration) WithMonths(values ...monitoringv1alpha1.MonthRange) *TimeIntervalApplyConfiguration {
	for i := range values {
		b.Months = append(b.Months, values[i])
	}
	return b
}

// WithYears adds the given value to the Years field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Years field.
func (b *TimeIntervalApplyConfiguration) WithYears(values ...monitoringv1alpha1.YearRange) *TimeIntervalApplyConfiguration {
	for i := range values {
		b.Years = append(b.Years, values[i])
	}
	return b
}
