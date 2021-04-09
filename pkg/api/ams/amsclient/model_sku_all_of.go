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
)

// SKUAllOf struct for SKUAllOf
type SKUAllOf struct {
	AvailabilityZoneType *string                   `json:"availability_zone_type,omitempty"`
	Byoc                 bool                      `json:"byoc"`
	Id                   *string                   `json:"id,omitempty"`
	ResourceName         *string                   `json:"resource_name,omitempty"`
	ResourceType         *string                   `json:"resource_type,omitempty"`
	Resources            *[]EphemeralResourceQuota `json:"resources,omitempty"`
}

// NewSKUAllOf instantiates a new SKUAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSKUAllOf(byoc bool) *SKUAllOf {
	this := SKUAllOf{}
	this.Byoc = byoc
	return &this
}

// NewSKUAllOfWithDefaults instantiates a new SKUAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSKUAllOfWithDefaults() *SKUAllOf {
	this := SKUAllOf{}
	return &this
}

// GetAvailabilityZoneType returns the AvailabilityZoneType field value if set, zero value otherwise.
func (o *SKUAllOf) GetAvailabilityZoneType() string {
	if o == nil || o.AvailabilityZoneType == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneType
}

// GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SKUAllOf) GetAvailabilityZoneTypeOk() (*string, bool) {
	if o == nil || o.AvailabilityZoneType == nil {
		return nil, false
	}
	return o.AvailabilityZoneType, true
}

// HasAvailabilityZoneType returns a boolean if a field has been set.
func (o *SKUAllOf) HasAvailabilityZoneType() bool {
	if o != nil && o.AvailabilityZoneType != nil {
		return true
	}

	return false
}

// SetAvailabilityZoneType gets a reference to the given string and assigns it to the AvailabilityZoneType field.
func (o *SKUAllOf) SetAvailabilityZoneType(v string) {
	o.AvailabilityZoneType = &v
}

// GetByoc returns the Byoc field value
func (o *SKUAllOf) GetByoc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Byoc
}

// GetByocOk returns a tuple with the Byoc field value
// and a boolean to check if the value has been set.
func (o *SKUAllOf) GetByocOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Byoc, true
}

// SetByoc sets field value
func (o *SKUAllOf) SetByoc(v bool) {
	o.Byoc = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SKUAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SKUAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SKUAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SKUAllOf) SetId(v string) {
	o.Id = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *SKUAllOf) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SKUAllOf) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *SKUAllOf) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *SKUAllOf) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *SKUAllOf) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SKUAllOf) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *SKUAllOf) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *SKUAllOf) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *SKUAllOf) GetResources() []EphemeralResourceQuota {
	if o == nil || o.Resources == nil {
		var ret []EphemeralResourceQuota
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SKUAllOf) GetResourcesOk() (*[]EphemeralResourceQuota, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *SKUAllOf) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []EphemeralResourceQuota and assigns it to the Resources field.
func (o *SKUAllOf) SetResources(v []EphemeralResourceQuota) {
	o.Resources = &v
}

func (o SKUAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailabilityZoneType != nil {
		toSerialize["availability_zone_type"] = o.AvailabilityZoneType
	}
	if true {
		toSerialize["byoc"] = o.Byoc
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableSKUAllOf struct {
	value *SKUAllOf
	isSet bool
}

func (v NullableSKUAllOf) Get() *SKUAllOf {
	return v.value
}

func (v *NullableSKUAllOf) Set(val *SKUAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSKUAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSKUAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSKUAllOf(val *SKUAllOf) *NullableSKUAllOf {
	return &NullableSKUAllOf{value: val, isSet: true}
}

func (v NullableSKUAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSKUAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
