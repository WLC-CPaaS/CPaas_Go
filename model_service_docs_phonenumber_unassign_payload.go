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

// checks if the ServiceDocsPhonenumberUnassignPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDocsPhonenumberUnassignPayload{}

// ServiceDocsPhonenumberUnassignPayload struct for ServiceDocsPhonenumberUnassignPayload
type ServiceDocsPhonenumberUnassignPayload struct {
	Phonenumber *string `json:"phonenumber,omitempty"`
}

// NewServiceDocsPhonenumberUnassignPayload instantiates a new ServiceDocsPhonenumberUnassignPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDocsPhonenumberUnassignPayload() *ServiceDocsPhonenumberUnassignPayload {
	this := ServiceDocsPhonenumberUnassignPayload{}
	return &this
}

// NewServiceDocsPhonenumberUnassignPayloadWithDefaults instantiates a new ServiceDocsPhonenumberUnassignPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDocsPhonenumberUnassignPayloadWithDefaults() *ServiceDocsPhonenumberUnassignPayload {
	this := ServiceDocsPhonenumberUnassignPayload{}
	return &this
}

// GetPhonenumber returns the Phonenumber field value if set, zero value otherwise.
func (o *ServiceDocsPhonenumberUnassignPayload) GetPhonenumber() string {
	if o == nil || IsNil(o.Phonenumber) {
		var ret string
		return ret
	}
	return *o.Phonenumber
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsPhonenumberUnassignPayload) GetPhonenumberOk() (*string, bool) {
	if o == nil || IsNil(o.Phonenumber) {
		return nil, false
	}
	return o.Phonenumber, true
}

// HasPhonenumber returns a boolean if a field has been set.
func (o *ServiceDocsPhonenumberUnassignPayload) HasPhonenumber() bool {
	if o != nil && !IsNil(o.Phonenumber) {
		return true
	}

	return false
}

// SetPhonenumber gets a reference to the given string and assigns it to the Phonenumber field.
func (o *ServiceDocsPhonenumberUnassignPayload) SetPhonenumber(v string) {
	o.Phonenumber = &v
}

func (o ServiceDocsPhonenumberUnassignPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDocsPhonenumberUnassignPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phonenumber) {
		toSerialize["phonenumber"] = o.Phonenumber
	}
	return toSerialize, nil
}

type NullableServiceDocsPhonenumberUnassignPayload struct {
	value *ServiceDocsPhonenumberUnassignPayload
	isSet bool
}

func (v NullableServiceDocsPhonenumberUnassignPayload) Get() *ServiceDocsPhonenumberUnassignPayload {
	return v.value
}

func (v *NullableServiceDocsPhonenumberUnassignPayload) Set(val *ServiceDocsPhonenumberUnassignPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDocsPhonenumberUnassignPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDocsPhonenumberUnassignPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDocsPhonenumberUnassignPayload(val *ServiceDocsPhonenumberUnassignPayload) *NullableServiceDocsPhonenumberUnassignPayload {
	return &NullableServiceDocsPhonenumberUnassignPayload{value: val, isSet: true}
}

func (v NullableServiceDocsPhonenumberUnassignPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDocsPhonenumberUnassignPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


