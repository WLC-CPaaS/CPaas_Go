/*
White Label Communications CPaas API Documentation

A CPaaS platform API

API version: 1.1
Contact: support@whitelabelcomm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ServiceDeviceOutputFullCalleridEmergency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDeviceOutputFullCalleridEmergency{}

// ServiceDeviceOutputFullCalleridEmergency struct for ServiceDeviceOutputFullCalleridEmergency
type ServiceDeviceOutputFullCalleridEmergency struct {
	Number *string `json:"number,omitempty"`
}

// NewServiceDeviceOutputFullCalleridEmergency instantiates a new ServiceDeviceOutputFullCalleridEmergency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDeviceOutputFullCalleridEmergency() *ServiceDeviceOutputFullCalleridEmergency {
	this := ServiceDeviceOutputFullCalleridEmergency{}
	return &this
}

// NewServiceDeviceOutputFullCalleridEmergencyWithDefaults instantiates a new ServiceDeviceOutputFullCalleridEmergency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDeviceOutputFullCalleridEmergencyWithDefaults() *ServiceDeviceOutputFullCalleridEmergency {
	this := ServiceDeviceOutputFullCalleridEmergency{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ServiceDeviceOutputFullCalleridEmergency) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeviceOutputFullCalleridEmergency) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ServiceDeviceOutputFullCalleridEmergency) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *ServiceDeviceOutputFullCalleridEmergency) SetNumber(v string) {
	o.Number = &v
}

func (o ServiceDeviceOutputFullCalleridEmergency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDeviceOutputFullCalleridEmergency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return toSerialize, nil
}

type NullableServiceDeviceOutputFullCalleridEmergency struct {
	value *ServiceDeviceOutputFullCalleridEmergency
	isSet bool
}

func (v NullableServiceDeviceOutputFullCalleridEmergency) Get() *ServiceDeviceOutputFullCalleridEmergency {
	return v.value
}

func (v *NullableServiceDeviceOutputFullCalleridEmergency) Set(val *ServiceDeviceOutputFullCalleridEmergency) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDeviceOutputFullCalleridEmergency) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDeviceOutputFullCalleridEmergency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDeviceOutputFullCalleridEmergency(val *ServiceDeviceOutputFullCalleridEmergency) *NullableServiceDeviceOutputFullCalleridEmergency {
	return &NullableServiceDeviceOutputFullCalleridEmergency{value: val, isSet: true}
}

func (v NullableServiceDeviceOutputFullCalleridEmergency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDeviceOutputFullCalleridEmergency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


