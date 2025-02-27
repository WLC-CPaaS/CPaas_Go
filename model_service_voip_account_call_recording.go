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

// checks if the ServiceVOIPAccountCallRecording type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVOIPAccountCallRecording{}

// ServiceVOIPAccountCallRecording struct for ServiceVOIPAccountCallRecording
type ServiceVOIPAccountCallRecording struct {
	Account *ServiceCallRecordingSettings `json:"account,omitempty"`
	Endpoint *ServiceCallRecordingSettings `json:"endpoint,omitempty"`
}

// NewServiceVOIPAccountCallRecording instantiates a new ServiceVOIPAccountCallRecording object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVOIPAccountCallRecording() *ServiceVOIPAccountCallRecording {
	this := ServiceVOIPAccountCallRecording{}
	return &this
}

// NewServiceVOIPAccountCallRecordingWithDefaults instantiates a new ServiceVOIPAccountCallRecording object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVOIPAccountCallRecordingWithDefaults() *ServiceVOIPAccountCallRecording {
	this := ServiceVOIPAccountCallRecording{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ServiceVOIPAccountCallRecording) GetAccount() ServiceCallRecordingSettings {
	if o == nil || IsNil(o.Account) {
		var ret ServiceCallRecordingSettings
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPAccountCallRecording) GetAccountOk() (*ServiceCallRecordingSettings, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ServiceVOIPAccountCallRecording) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ServiceCallRecordingSettings and assigns it to the Account field.
func (o *ServiceVOIPAccountCallRecording) SetAccount(v ServiceCallRecordingSettings) {
	o.Account = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ServiceVOIPAccountCallRecording) GetEndpoint() ServiceCallRecordingSettings {
	if o == nil || IsNil(o.Endpoint) {
		var ret ServiceCallRecordingSettings
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPAccountCallRecording) GetEndpointOk() (*ServiceCallRecordingSettings, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ServiceVOIPAccountCallRecording) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given ServiceCallRecordingSettings and assigns it to the Endpoint field.
func (o *ServiceVOIPAccountCallRecording) SetEndpoint(v ServiceCallRecordingSettings) {
	o.Endpoint = &v
}

func (o ServiceVOIPAccountCallRecording) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVOIPAccountCallRecording) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	return toSerialize, nil
}

type NullableServiceVOIPAccountCallRecording struct {
	value *ServiceVOIPAccountCallRecording
	isSet bool
}

func (v NullableServiceVOIPAccountCallRecording) Get() *ServiceVOIPAccountCallRecording {
	return v.value
}

func (v *NullableServiceVOIPAccountCallRecording) Set(val *ServiceVOIPAccountCallRecording) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVOIPAccountCallRecording) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVOIPAccountCallRecording) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVOIPAccountCallRecording(val *ServiceVOIPAccountCallRecording) *NullableServiceVOIPAccountCallRecording {
	return &NullableServiceVOIPAccountCallRecording{value: val, isSet: true}
}

func (v NullableServiceVOIPAccountCallRecording) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVOIPAccountCallRecording) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


