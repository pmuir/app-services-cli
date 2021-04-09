/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amsclient

import (
	"encoding/json"
	"time"
)

// Metric struct for Metric
type Metric struct {
	Href           *string    `json:"href,omitempty"`
	Id             *string    `json:"id,omitempty"`
	Kind           *string    `json:"kind,omitempty"`
	ExternalId     *string    `json:"external_id,omitempty"`
	HealthState    *string    `json:"health_state,omitempty"`
	Metrics        *string    `json:"metrics,omitempty"`
	QueryTimestamp *time.Time `json:"query_timestamp,omitempty"`
}

// NewMetric instantiates a new Metric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetric() *Metric {
	this := Metric{}
	return &this
}

// NewMetricWithDefaults instantiates a new Metric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricWithDefaults() *Metric {
	this := Metric{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Metric) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Metric) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Metric) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Metric) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Metric) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Metric) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Metric) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Metric) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Metric) SetKind(v string) {
	o.Kind = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Metric) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Metric) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Metric) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetHealthState returns the HealthState field value if set, zero value otherwise.
func (o *Metric) GetHealthState() string {
	if o == nil || o.HealthState == nil {
		var ret string
		return ret
	}
	return *o.HealthState
}

// GetHealthStateOk returns a tuple with the HealthState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetHealthStateOk() (*string, bool) {
	if o == nil || o.HealthState == nil {
		return nil, false
	}
	return o.HealthState, true
}

// HasHealthState returns a boolean if a field has been set.
func (o *Metric) HasHealthState() bool {
	if o != nil && o.HealthState != nil {
		return true
	}

	return false
}

// SetHealthState gets a reference to the given string and assigns it to the HealthState field.
func (o *Metric) SetHealthState(v string) {
	o.HealthState = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *Metric) GetMetrics() string {
	if o == nil || o.Metrics == nil {
		var ret string
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetMetricsOk() (*string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *Metric) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given string and assigns it to the Metrics field.
func (o *Metric) SetMetrics(v string) {
	o.Metrics = &v
}

// GetQueryTimestamp returns the QueryTimestamp field value if set, zero value otherwise.
func (o *Metric) GetQueryTimestamp() time.Time {
	if o == nil || o.QueryTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.QueryTimestamp
}

// GetQueryTimestampOk returns a tuple with the QueryTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetQueryTimestampOk() (*time.Time, bool) {
	if o == nil || o.QueryTimestamp == nil {
		return nil, false
	}
	return o.QueryTimestamp, true
}

// HasQueryTimestamp returns a boolean if a field has been set.
func (o *Metric) HasQueryTimestamp() bool {
	if o != nil && o.QueryTimestamp != nil {
		return true
	}

	return false
}

// SetQueryTimestamp gets a reference to the given time.Time and assigns it to the QueryTimestamp field.
func (o *Metric) SetQueryTimestamp(v time.Time) {
	o.QueryTimestamp = &v
}

func (o Metric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.HealthState != nil {
		toSerialize["health_state"] = o.HealthState
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.QueryTimestamp != nil {
		toSerialize["query_timestamp"] = o.QueryTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableMetric struct {
	value *Metric
	isSet bool
}

func (v NullableMetric) Get() *Metric {
	return v.value
}

func (v *NullableMetric) Set(val *Metric) {
	v.value = val
	v.isSet = true
}

func (v NullableMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetric(val *Metric) *NullableMetric {
	return &NullableMetric{value: val, isSet: true}
}

func (v NullableMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
