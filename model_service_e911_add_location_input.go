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

// checks if the ServiceE911AddLocationInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceE911AddLocationInput{}

// ServiceE911AddLocationInput struct for ServiceE911AddLocationInput
type ServiceE911AddLocationInput struct {
	Location ServiceE911LocationInput `json:"location"`
	Uri ServiceE911URIInput `json:"uri"`
}

type _ServiceE911AddLocationInput ServiceE911AddLocationInput

// NewServiceE911AddLocationInput instantiates a new ServiceE911AddLocationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceE911AddLocationInput(location ServiceE911LocationInput, uri ServiceE911URIInput) *ServiceE911AddLocationInput {
	this := ServiceE911AddLocationInput{}
	this.Location = location
	this.Uri = uri
	return &this
}

// NewServiceE911AddLocationInputWithDefaults instantiates a new ServiceE911AddLocationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceE911AddLocationInputWithDefaults() *ServiceE911AddLocationInput {
	this := ServiceE911AddLocationInput{}
	return &this
}

// GetLocation returns the Location field value
func (o *ServiceE911AddLocationInput) GetLocation() ServiceE911LocationInput {
	if o == nil {
		var ret ServiceE911LocationInput
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *ServiceE911AddLocationInput) GetLocationOk() (*ServiceE911LocationInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *ServiceE911AddLocationInput) SetLocation(v ServiceE911LocationInput) {
	o.Location = v
}

// GetUri returns the Uri field value
func (o *ServiceE911AddLocationInput) GetUri() ServiceE911URIInput {
	if o == nil {
		var ret ServiceE911URIInput
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ServiceE911AddLocationInput) GetUriOk() (*ServiceE911URIInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ServiceE911AddLocationInput) SetUri(v ServiceE911URIInput) {
	o.Uri = v
}

func (o ServiceE911AddLocationInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceE911AddLocationInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *ServiceE911AddLocationInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"location",
		"uri",
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

	varServiceE911AddLocationInput := _ServiceE911AddLocationInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceE911AddLocationInput)

	if err != nil {
		return err
	}

	*o = ServiceE911AddLocationInput(varServiceE911AddLocationInput)

	return err
}

type NullableServiceE911AddLocationInput struct {
	value *ServiceE911AddLocationInput
	isSet bool
}

func (v NullableServiceE911AddLocationInput) Get() *ServiceE911AddLocationInput {
	return v.value
}

func (v *NullableServiceE911AddLocationInput) Set(val *ServiceE911AddLocationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceE911AddLocationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceE911AddLocationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceE911AddLocationInput(val *ServiceE911AddLocationInput) *NullableServiceE911AddLocationInput {
	return &NullableServiceE911AddLocationInput{value: val, isSet: true}
}

func (v NullableServiceE911AddLocationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceE911AddLocationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


