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

// RoleBindingList struct for RoleBindingList
type RoleBindingList struct {
	Kind  string        `json:"kind"`
	Page  int32         `json:"page"`
	Size  int32         `json:"size"`
	Total int32         `json:"total"`
	Items []RoleBinding `json:"items"`
}

// NewRoleBindingList instantiates a new RoleBindingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBindingList(kind string, page int32, size int32, total int32, items []RoleBinding) *RoleBindingList {
	this := RoleBindingList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewRoleBindingListWithDefaults instantiates a new RoleBindingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBindingListWithDefaults() *RoleBindingList {
	this := RoleBindingList{}
	return &this
}

// GetKind returns the Kind field value
func (o *RoleBindingList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *RoleBindingList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *RoleBindingList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *RoleBindingList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *RoleBindingList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *RoleBindingList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *RoleBindingList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *RoleBindingList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *RoleBindingList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *RoleBindingList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *RoleBindingList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *RoleBindingList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *RoleBindingList) GetItems() []RoleBinding {
	if o == nil {
		var ret []RoleBinding
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RoleBindingList) GetItemsOk() (*[]RoleBinding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *RoleBindingList) SetItems(v []RoleBinding) {
	o.Items = v
}

func (o RoleBindingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableRoleBindingList struct {
	value *RoleBindingList
	isSet bool
}

func (v NullableRoleBindingList) Get() *RoleBindingList {
	return v.value
}

func (v *NullableRoleBindingList) Set(val *RoleBindingList) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleBindingList) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleBindingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleBindingList(val *RoleBindingList) *NullableRoleBindingList {
	return &NullableRoleBindingList{value: val, isSet: true}
}

func (v NullableRoleBindingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleBindingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
