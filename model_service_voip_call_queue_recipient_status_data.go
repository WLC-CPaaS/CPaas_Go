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

// checks if the ServiceVOIPCallQueueRecipientStatusData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVOIPCallQueueRecipientStatusData{}

// ServiceVOIPCallQueueRecipientStatusData struct for ServiceVOIPCallQueueRecipientStatusData
type ServiceVOIPCallQueueRecipientStatusData struct {
	Status string `json:"status"`
}

type _ServiceVOIPCallQueueRecipientStatusData ServiceVOIPCallQueueRecipientStatusData

// NewServiceVOIPCallQueueRecipientStatusData instantiates a new ServiceVOIPCallQueueRecipientStatusData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVOIPCallQueueRecipientStatusData(status string) *ServiceVOIPCallQueueRecipientStatusData {
	this := ServiceVOIPCallQueueRecipientStatusData{}
	this.Status = status
	return &this
}

// NewServiceVOIPCallQueueRecipientStatusDataWithDefaults instantiates a new ServiceVOIPCallQueueRecipientStatusData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVOIPCallQueueRecipientStatusDataWithDefaults() *ServiceVOIPCallQueueRecipientStatusData {
	this := ServiceVOIPCallQueueRecipientStatusData{}
	return &this
}

// GetStatus returns the Status field value
func (o *ServiceVOIPCallQueueRecipientStatusData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ServiceVOIPCallQueueRecipientStatusData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ServiceVOIPCallQueueRecipientStatusData) SetStatus(v string) {
	o.Status = v
}

func (o ServiceVOIPCallQueueRecipientStatusData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVOIPCallQueueRecipientStatusData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ServiceVOIPCallQueueRecipientStatusData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varServiceVOIPCallQueueRecipientStatusData := _ServiceVOIPCallQueueRecipientStatusData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceVOIPCallQueueRecipientStatusData)

	if err != nil {
		return err
	}

	*o = ServiceVOIPCallQueueRecipientStatusData(varServiceVOIPCallQueueRecipientStatusData)

	return err
}

type NullableServiceVOIPCallQueueRecipientStatusData struct {
	value *ServiceVOIPCallQueueRecipientStatusData
	isSet bool
}

func (v NullableServiceVOIPCallQueueRecipientStatusData) Get() *ServiceVOIPCallQueueRecipientStatusData {
	return v.value
}

func (v *NullableServiceVOIPCallQueueRecipientStatusData) Set(val *ServiceVOIPCallQueueRecipientStatusData) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVOIPCallQueueRecipientStatusData) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVOIPCallQueueRecipientStatusData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVOIPCallQueueRecipientStatusData(val *ServiceVOIPCallQueueRecipientStatusData) *NullableServiceVOIPCallQueueRecipientStatusData {
	return &NullableServiceVOIPCallQueueRecipientStatusData{value: val, isSet: true}
}

func (v NullableServiceVOIPCallQueueRecipientStatusData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVOIPCallQueueRecipientStatusData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


