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
	"bytes"
	"fmt"
)

// checks if the ServiceVOIPCallQueueEnableMembershipData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVOIPCallQueueEnableMembershipData{}

// ServiceVOIPCallQueueEnableMembershipData struct for ServiceVOIPCallQueueEnableMembershipData
type ServiceVOIPCallQueueEnableMembershipData struct {
	AcceptCharges bool `json:"accept_charges"`
}

type _ServiceVOIPCallQueueEnableMembershipData ServiceVOIPCallQueueEnableMembershipData

// NewServiceVOIPCallQueueEnableMembershipData instantiates a new ServiceVOIPCallQueueEnableMembershipData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVOIPCallQueueEnableMembershipData(acceptCharges bool) *ServiceVOIPCallQueueEnableMembershipData {
	this := ServiceVOIPCallQueueEnableMembershipData{}
	this.AcceptCharges = acceptCharges
	return &this
}

// NewServiceVOIPCallQueueEnableMembershipDataWithDefaults instantiates a new ServiceVOIPCallQueueEnableMembershipData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVOIPCallQueueEnableMembershipDataWithDefaults() *ServiceVOIPCallQueueEnableMembershipData {
	this := ServiceVOIPCallQueueEnableMembershipData{}
	return &this
}

// GetAcceptCharges returns the AcceptCharges field value
func (o *ServiceVOIPCallQueueEnableMembershipData) GetAcceptCharges() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptCharges
}

// GetAcceptChargesOk returns a tuple with the AcceptCharges field value
// and a boolean to check if the value has been set.
func (o *ServiceVOIPCallQueueEnableMembershipData) GetAcceptChargesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptCharges, true
}

// SetAcceptCharges sets field value
func (o *ServiceVOIPCallQueueEnableMembershipData) SetAcceptCharges(v bool) {
	o.AcceptCharges = v
}

func (o ServiceVOIPCallQueueEnableMembershipData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVOIPCallQueueEnableMembershipData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accept_charges"] = o.AcceptCharges
	return toSerialize, nil
}

func (o *ServiceVOIPCallQueueEnableMembershipData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accept_charges",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServiceVOIPCallQueueEnableMembershipData := _ServiceVOIPCallQueueEnableMembershipData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceVOIPCallQueueEnableMembershipData)

	if err != nil {
		return err
	}

	*o = ServiceVOIPCallQueueEnableMembershipData(varServiceVOIPCallQueueEnableMembershipData)

	return err
}

type NullableServiceVOIPCallQueueEnableMembershipData struct {
	value *ServiceVOIPCallQueueEnableMembershipData
	isSet bool
}

func (v NullableServiceVOIPCallQueueEnableMembershipData) Get() *ServiceVOIPCallQueueEnableMembershipData {
	return v.value
}

func (v *NullableServiceVOIPCallQueueEnableMembershipData) Set(val *ServiceVOIPCallQueueEnableMembershipData) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVOIPCallQueueEnableMembershipData) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVOIPCallQueueEnableMembershipData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVOIPCallQueueEnableMembershipData(val *ServiceVOIPCallQueueEnableMembershipData) *NullableServiceVOIPCallQueueEnableMembershipData {
	return &NullableServiceVOIPCallQueueEnableMembershipData{value: val, isSet: true}
}

func (v NullableServiceVOIPCallQueueEnableMembershipData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVOIPCallQueueEnableMembershipData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


