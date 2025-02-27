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

// checks if the ServiceQueueRecipientOutputFullRecipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceQueueRecipientOutputFullRecipient{}

// ServiceQueueRecipientOutputFullRecipient struct for ServiceQueueRecipientOutputFullRecipient
type ServiceQueueRecipientOutputFullRecipient struct {
	Features *ServiceQueueRecipientOutputFullFeatures `json:"features,omitempty"`
}

// NewServiceQueueRecipientOutputFullRecipient instantiates a new ServiceQueueRecipientOutputFullRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceQueueRecipientOutputFullRecipient() *ServiceQueueRecipientOutputFullRecipient {
	this := ServiceQueueRecipientOutputFullRecipient{}
	return &this
}

// NewServiceQueueRecipientOutputFullRecipientWithDefaults instantiates a new ServiceQueueRecipientOutputFullRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceQueueRecipientOutputFullRecipientWithDefaults() *ServiceQueueRecipientOutputFullRecipient {
	this := ServiceQueueRecipientOutputFullRecipient{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ServiceQueueRecipientOutputFullRecipient) GetFeatures() ServiceQueueRecipientOutputFullFeatures {
	if o == nil || IsNil(o.Features) {
		var ret ServiceQueueRecipientOutputFullFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceQueueRecipientOutputFullRecipient) GetFeaturesOk() (*ServiceQueueRecipientOutputFullFeatures, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ServiceQueueRecipientOutputFullRecipient) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given ServiceQueueRecipientOutputFullFeatures and assigns it to the Features field.
func (o *ServiceQueueRecipientOutputFullRecipient) SetFeatures(v ServiceQueueRecipientOutputFullFeatures) {
	o.Features = &v
}

func (o ServiceQueueRecipientOutputFullRecipient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceQueueRecipientOutputFullRecipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return toSerialize, nil
}

type NullableServiceQueueRecipientOutputFullRecipient struct {
	value *ServiceQueueRecipientOutputFullRecipient
	isSet bool
}

func (v NullableServiceQueueRecipientOutputFullRecipient) Get() *ServiceQueueRecipientOutputFullRecipient {
	return v.value
}

func (v *NullableServiceQueueRecipientOutputFullRecipient) Set(val *ServiceQueueRecipientOutputFullRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceQueueRecipientOutputFullRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceQueueRecipientOutputFullRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceQueueRecipientOutputFullRecipient(val *ServiceQueueRecipientOutputFullRecipient) *NullableServiceQueueRecipientOutputFullRecipient {
	return &NullableServiceQueueRecipientOutputFullRecipient{value: val, isSet: true}
}

func (v NullableServiceQueueRecipientOutputFullRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceQueueRecipientOutputFullRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


