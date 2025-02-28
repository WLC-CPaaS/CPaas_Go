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

// checks if the ServiceE911StatusOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceE911StatusOutput{}

// ServiceE911StatusOutput struct for ServiceE911StatusOutput
type ServiceE911StatusOutput struct {
	Code *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewServiceE911StatusOutput instantiates a new ServiceE911StatusOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceE911StatusOutput() *ServiceE911StatusOutput {
	this := ServiceE911StatusOutput{}
	return &this
}

// NewServiceE911StatusOutputWithDefaults instantiates a new ServiceE911StatusOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceE911StatusOutputWithDefaults() *ServiceE911StatusOutput {
	this := ServiceE911StatusOutput{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ServiceE911StatusOutput) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911StatusOutput) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ServiceE911StatusOutput) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ServiceE911StatusOutput) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceE911StatusOutput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911StatusOutput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceE911StatusOutput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceE911StatusOutput) SetDescription(v string) {
	o.Description = &v
}

func (o ServiceE911StatusOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceE911StatusOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableServiceE911StatusOutput struct {
	value *ServiceE911StatusOutput
	isSet bool
}

func (v NullableServiceE911StatusOutput) Get() *ServiceE911StatusOutput {
	return v.value
}

func (v *NullableServiceE911StatusOutput) Set(val *ServiceE911StatusOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceE911StatusOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceE911StatusOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceE911StatusOutput(val *ServiceE911StatusOutput) *NullableServiceE911StatusOutput {
	return &NullableServiceE911StatusOutput{value: val, isSet: true}
}

func (v NullableServiceE911StatusOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceE911StatusOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


