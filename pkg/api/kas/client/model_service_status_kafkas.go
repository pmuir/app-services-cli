/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtv1

import (
	"encoding/json"
)

// ServiceStatusKafkas The kafka resource api status
type ServiceStatusKafkas struct {
	// Indicates whether we have reached kafka maximum capacity
	MaxCapacityReached bool `json:"max_capacity_reached"`
}

// NewServiceStatusKafkas instantiates a new ServiceStatusKafkas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStatusKafkas(maxCapacityReached bool) *ServiceStatusKafkas {
	this := ServiceStatusKafkas{}
	this.MaxCapacityReached = maxCapacityReached
	return &this
}

// NewServiceStatusKafkasWithDefaults instantiates a new ServiceStatusKafkas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStatusKafkasWithDefaults() *ServiceStatusKafkas {
	this := ServiceStatusKafkas{}
	return &this
}

// GetMaxCapacityReached returns the MaxCapacityReached field value
func (o *ServiceStatusKafkas) GetMaxCapacityReached() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MaxCapacityReached
}

// GetMaxCapacityReachedOk returns a tuple with the MaxCapacityReached field value
// and a boolean to check if the value has been set.
func (o *ServiceStatusKafkas) GetMaxCapacityReachedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCapacityReached, true
}

// SetMaxCapacityReached sets field value
func (o *ServiceStatusKafkas) SetMaxCapacityReached(v bool) {
	o.MaxCapacityReached = v
}

func (o ServiceStatusKafkas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["max_capacity_reached"] = o.MaxCapacityReached
	}
	return json.Marshal(toSerialize)
}

type NullableServiceStatusKafkas struct {
	value *ServiceStatusKafkas
	isSet bool
}

func (v NullableServiceStatusKafkas) Get() *ServiceStatusKafkas {
	return v.value
}

func (v *NullableServiceStatusKafkas) Set(val *ServiceStatusKafkas) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStatusKafkas) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStatusKafkas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStatusKafkas(val *ServiceStatusKafkas) *NullableServiceStatusKafkas {
	return &NullableServiceStatusKafkas{value: val, isSet: true}
}

func (v NullableServiceStatusKafkas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStatusKafkas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
