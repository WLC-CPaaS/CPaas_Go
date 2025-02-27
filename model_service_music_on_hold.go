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

// checks if the ServiceMusicOnHold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceMusicOnHold{}

// ServiceMusicOnHold struct for ServiceMusicOnHold
type ServiceMusicOnHold struct {
	MediaId *string `json:"media_id,omitempty"`
}

// NewServiceMusicOnHold instantiates a new ServiceMusicOnHold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceMusicOnHold() *ServiceMusicOnHold {
	this := ServiceMusicOnHold{}
	return &this
}

// NewServiceMusicOnHoldWithDefaults instantiates a new ServiceMusicOnHold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceMusicOnHoldWithDefaults() *ServiceMusicOnHold {
	this := ServiceMusicOnHold{}
	return &this
}

// GetMediaId returns the MediaId field value if set, zero value otherwise.
func (o *ServiceMusicOnHold) GetMediaId() string {
	if o == nil || IsNil(o.MediaId) {
		var ret string
		return ret
	}
	return *o.MediaId
}

// GetMediaIdOk returns a tuple with the MediaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMusicOnHold) GetMediaIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaId) {
		return nil, false
	}
	return o.MediaId, true
}

// HasMediaId returns a boolean if a field has been set.
func (o *ServiceMusicOnHold) HasMediaId() bool {
	if o != nil && !IsNil(o.MediaId) {
		return true
	}

	return false
}

// SetMediaId gets a reference to the given string and assigns it to the MediaId field.
func (o *ServiceMusicOnHold) SetMediaId(v string) {
	o.MediaId = &v
}

func (o ServiceMusicOnHold) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceMusicOnHold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MediaId) {
		toSerialize["media_id"] = o.MediaId
	}
	return toSerialize, nil
}

type NullableServiceMusicOnHold struct {
	value *ServiceMusicOnHold
	isSet bool
}

func (v NullableServiceMusicOnHold) Get() *ServiceMusicOnHold {
	return v.value
}

func (v *NullableServiceMusicOnHold) Set(val *ServiceMusicOnHold) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceMusicOnHold) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceMusicOnHold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceMusicOnHold(val *ServiceMusicOnHold) *NullableServiceMusicOnHold {
	return &NullableServiceMusicOnHold{value: val, isSet: true}
}

func (v NullableServiceMusicOnHold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceMusicOnHold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


