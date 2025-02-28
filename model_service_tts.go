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

// checks if the ServiceTTS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTTS{}

// ServiceTTS struct for ServiceTTS
type ServiceTTS struct {
	Text string `json:"text"`
	Voice *string `json:"voice,omitempty"`
}

type _ServiceTTS ServiceTTS

// NewServiceTTS instantiates a new ServiceTTS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTTS(text string) *ServiceTTS {
	this := ServiceTTS{}
	this.Text = text
	return &this
}

// NewServiceTTSWithDefaults instantiates a new ServiceTTS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTTSWithDefaults() *ServiceTTS {
	this := ServiceTTS{}
	return &this
}

// GetText returns the Text field value
func (o *ServiceTTS) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *ServiceTTS) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *ServiceTTS) SetText(v string) {
	o.Text = v
}

// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *ServiceTTS) GetVoice() string {
	if o == nil || IsNil(o.Voice) {
		var ret string
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTTS) GetVoiceOk() (*string, bool) {
	if o == nil || IsNil(o.Voice) {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *ServiceTTS) HasVoice() bool {
	if o != nil && !IsNil(o.Voice) {
		return true
	}

	return false
}

// SetVoice gets a reference to the given string and assigns it to the Voice field.
func (o *ServiceTTS) SetVoice(v string) {
	o.Voice = &v
}

func (o ServiceTTS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTTS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if !IsNil(o.Voice) {
		toSerialize["voice"] = o.Voice
	}
	return toSerialize, nil
}

func (o *ServiceTTS) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
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

	varServiceTTS := _ServiceTTS{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTTS)

	if err != nil {
		return err
	}

	*o = ServiceTTS(varServiceTTS)

	return err
}

type NullableServiceTTS struct {
	value *ServiceTTS
	isSet bool
}

func (v NullableServiceTTS) Get() *ServiceTTS {
	return v.value
}

func (v *NullableServiceTTS) Set(val *ServiceTTS) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTTS) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTTS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTTS(val *ServiceTTS) *NullableServiceTTS {
	return &NullableServiceTTS{value: val, isSet: true}
}

func (v NullableServiceTTS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTTS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


