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

// checks if the ServiceAdminUserEditData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAdminUserEditData{}

// ServiceAdminUserEditData struct for ServiceAdminUserEditData
type ServiceAdminUserEditData struct {
	Role string `json:"role"`
}

type _ServiceAdminUserEditData ServiceAdminUserEditData

// NewServiceAdminUserEditData instantiates a new ServiceAdminUserEditData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAdminUserEditData(role string) *ServiceAdminUserEditData {
	this := ServiceAdminUserEditData{}
	this.Role = role
	return &this
}

// NewServiceAdminUserEditDataWithDefaults instantiates a new ServiceAdminUserEditData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAdminUserEditDataWithDefaults() *ServiceAdminUserEditData {
	this := ServiceAdminUserEditData{}
	return &this
}

// GetRole returns the Role field value
func (o *ServiceAdminUserEditData) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ServiceAdminUserEditData) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ServiceAdminUserEditData) SetRole(v string) {
	o.Role = v
}

func (o ServiceAdminUserEditData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAdminUserEditData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *ServiceAdminUserEditData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role",
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

	varServiceAdminUserEditData := _ServiceAdminUserEditData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceAdminUserEditData)

	if err != nil {
		return err
	}

	*o = ServiceAdminUserEditData(varServiceAdminUserEditData)

	return err
}

type NullableServiceAdminUserEditData struct {
	value *ServiceAdminUserEditData
	isSet bool
}

func (v NullableServiceAdminUserEditData) Get() *ServiceAdminUserEditData {
	return v.value
}

func (v *NullableServiceAdminUserEditData) Set(val *ServiceAdminUserEditData) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAdminUserEditData) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAdminUserEditData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAdminUserEditData(val *ServiceAdminUserEditData) *NullableServiceAdminUserEditData {
	return &NullableServiceAdminUserEditData{value: val, isSet: true}
}

func (v NullableServiceAdminUserEditData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAdminUserEditData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


