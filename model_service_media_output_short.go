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

// checks if the ServiceMediaOutputShort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceMediaOutputShort{}

// ServiceMediaOutputShort struct for ServiceMediaOutputShort
type ServiceMediaOutputShort struct {
	Id *string `json:"id,omitempty"`
	IsPrompt *bool `json:"is_prompt,omitempty"`
	Language *string `json:"language,omitempty"`
	MediaSource *string `json:"media_source,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewServiceMediaOutputShort instantiates a new ServiceMediaOutputShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceMediaOutputShort() *ServiceMediaOutputShort {
	this := ServiceMediaOutputShort{}
	return &this
}

// NewServiceMediaOutputShortWithDefaults instantiates a new ServiceMediaOutputShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceMediaOutputShortWithDefaults() *ServiceMediaOutputShort {
	this := ServiceMediaOutputShort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceMediaOutputShort) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMediaOutputShort) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceMediaOutputShort) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceMediaOutputShort) SetId(v string) {
	o.Id = &v
}

// GetIsPrompt returns the IsPrompt field value if set, zero value otherwise.
func (o *ServiceMediaOutputShort) GetIsPrompt() bool {
	if o == nil || IsNil(o.IsPrompt) {
		var ret bool
		return ret
	}
	return *o.IsPrompt
}

// GetIsPromptOk returns a tuple with the IsPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMediaOutputShort) GetIsPromptOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrompt) {
		return nil, false
	}
	return o.IsPrompt, true
}

// HasIsPrompt returns a boolean if a field has been set.
func (o *ServiceMediaOutputShort) HasIsPrompt() bool {
	if o != nil && !IsNil(o.IsPrompt) {
		return true
	}

	return false
}

// SetIsPrompt gets a reference to the given bool and assigns it to the IsPrompt field.
func (o *ServiceMediaOutputShort) SetIsPrompt(v bool) {
	o.IsPrompt = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ServiceMediaOutputShort) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMediaOutputShort) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ServiceMediaOutputShort) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ServiceMediaOutputShort) SetLanguage(v string) {
	o.Language = &v
}

// GetMediaSource returns the MediaSource field value if set, zero value otherwise.
func (o *ServiceMediaOutputShort) GetMediaSource() string {
	if o == nil || IsNil(o.MediaSource) {
		var ret string
		return ret
	}
	return *o.MediaSource
}

// GetMediaSourceOk returns a tuple with the MediaSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMediaOutputShort) GetMediaSourceOk() (*string, bool) {
	if o == nil || IsNil(o.MediaSource) {
		return nil, false
	}
	return o.MediaSource, true
}

// HasMediaSource returns a boolean if a field has been set.
func (o *ServiceMediaOutputShort) HasMediaSource() bool {
	if o != nil && !IsNil(o.MediaSource) {
		return true
	}

	return false
}

// SetMediaSource gets a reference to the given string and assigns it to the MediaSource field.
func (o *ServiceMediaOutputShort) SetMediaSource(v string) {
	o.MediaSource = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceMediaOutputShort) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMediaOutputShort) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceMediaOutputShort) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceMediaOutputShort) SetName(v string) {
	o.Name = &v
}

func (o ServiceMediaOutputShort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceMediaOutputShort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsPrompt) {
		toSerialize["is_prompt"] = o.IsPrompt
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.MediaSource) {
		toSerialize["media_source"] = o.MediaSource
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableServiceMediaOutputShort struct {
	value *ServiceMediaOutputShort
	isSet bool
}

func (v NullableServiceMediaOutputShort) Get() *ServiceMediaOutputShort {
	return v.value
}

func (v *NullableServiceMediaOutputShort) Set(val *ServiceMediaOutputShort) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceMediaOutputShort) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceMediaOutputShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceMediaOutputShort(val *ServiceMediaOutputShort) *NullableServiceMediaOutputShort {
	return &NullableServiceMediaOutputShort{value: val, isSet: true}
}

func (v NullableServiceMediaOutputShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceMediaOutputShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


