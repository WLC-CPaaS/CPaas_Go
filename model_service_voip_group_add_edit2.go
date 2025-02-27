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

// checks if the ServiceVOIPGroupAddEdit2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVOIPGroupAddEdit2{}

// ServiceVOIPGroupAddEdit2 struct for ServiceVOIPGroupAddEdit2
type ServiceVOIPGroupAddEdit2 struct {
	Endpoints map[string]ServiceEndpoint `json:"endpoints"`
	Name string `json:"name"`
}

type _ServiceVOIPGroupAddEdit2 ServiceVOIPGroupAddEdit2

// NewServiceVOIPGroupAddEdit2 instantiates a new ServiceVOIPGroupAddEdit2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVOIPGroupAddEdit2(endpoints map[string]ServiceEndpoint, name string) *ServiceVOIPGroupAddEdit2 {
	this := ServiceVOIPGroupAddEdit2{}
	this.Endpoints = endpoints
	this.Name = name
	return &this
}

// NewServiceVOIPGroupAddEdit2WithDefaults instantiates a new ServiceVOIPGroupAddEdit2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVOIPGroupAddEdit2WithDefaults() *ServiceVOIPGroupAddEdit2 {
	this := ServiceVOIPGroupAddEdit2{}
	return &this
}

// GetEndpoints returns the Endpoints field value
func (o *ServiceVOIPGroupAddEdit2) GetEndpoints() map[string]ServiceEndpoint {
	if o == nil {
		var ret map[string]ServiceEndpoint
		return ret
	}

	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value
// and a boolean to check if the value has been set.
func (o *ServiceVOIPGroupAddEdit2) GetEndpointsOk() (*map[string]ServiceEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoints, true
}

// SetEndpoints sets field value
func (o *ServiceVOIPGroupAddEdit2) SetEndpoints(v map[string]ServiceEndpoint) {
	o.Endpoints = v
}

// GetName returns the Name field value
func (o *ServiceVOIPGroupAddEdit2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceVOIPGroupAddEdit2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceVOIPGroupAddEdit2) SetName(v string) {
	o.Name = v
}

func (o ServiceVOIPGroupAddEdit2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVOIPGroupAddEdit2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpoints"] = o.Endpoints
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ServiceVOIPGroupAddEdit2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"endpoints",
		"name",
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

	varServiceVOIPGroupAddEdit2 := _ServiceVOIPGroupAddEdit2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceVOIPGroupAddEdit2)

	if err != nil {
		return err
	}

	*o = ServiceVOIPGroupAddEdit2(varServiceVOIPGroupAddEdit2)

	return err
}

type NullableServiceVOIPGroupAddEdit2 struct {
	value *ServiceVOIPGroupAddEdit2
	isSet bool
}

func (v NullableServiceVOIPGroupAddEdit2) Get() *ServiceVOIPGroupAddEdit2 {
	return v.value
}

func (v *NullableServiceVOIPGroupAddEdit2) Set(val *ServiceVOIPGroupAddEdit2) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVOIPGroupAddEdit2) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVOIPGroupAddEdit2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVOIPGroupAddEdit2(val *ServiceVOIPGroupAddEdit2) *NullableServiceVOIPGroupAddEdit2 {
	return &NullableServiceVOIPGroupAddEdit2{value: val, isSet: true}
}

func (v NullableServiceVOIPGroupAddEdit2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVOIPGroupAddEdit2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


