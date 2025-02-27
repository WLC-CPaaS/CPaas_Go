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

// checks if the ServicePhonenumberOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePhonenumberOutput{}

// ServicePhonenumberOutput struct for ServicePhonenumberOutput
type ServicePhonenumberOutput struct {
	Id *string `json:"id,omitempty"`
}

// NewServicePhonenumberOutput instantiates a new ServicePhonenumberOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePhonenumberOutput() *ServicePhonenumberOutput {
	this := ServicePhonenumberOutput{}
	return &this
}

// NewServicePhonenumberOutputWithDefaults instantiates a new ServicePhonenumberOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePhonenumberOutputWithDefaults() *ServicePhonenumberOutput {
	this := ServicePhonenumberOutput{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServicePhonenumberOutput) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePhonenumberOutput) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServicePhonenumberOutput) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServicePhonenumberOutput) SetId(v string) {
	o.Id = &v
}

func (o ServicePhonenumberOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePhonenumberOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableServicePhonenumberOutput struct {
	value *ServicePhonenumberOutput
	isSet bool
}

func (v NullableServicePhonenumberOutput) Get() *ServicePhonenumberOutput {
	return v.value
}

func (v *NullableServicePhonenumberOutput) Set(val *ServicePhonenumberOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePhonenumberOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePhonenumberOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePhonenumberOutput(val *ServicePhonenumberOutput) *NullableServicePhonenumberOutput {
	return &NullableServicePhonenumberOutput{value: val, isSet: true}
}

func (v NullableServicePhonenumberOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePhonenumberOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


