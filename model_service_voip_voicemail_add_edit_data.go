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

// checks if the ServiceVOIPVoicemailAddEditData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVOIPVoicemailAddEditData{}

// ServiceVOIPVoicemailAddEditData struct for ServiceVOIPVoicemailAddEditData
type ServiceVOIPVoicemailAddEditData struct {
	Mailbox string `json:"mailbox"`
	Media *ServiceVoicemailMedia `json:"media,omitempty"`
	MediaExtension *string `json:"media_extension,omitempty"`
	Name string `json:"name"`
	OwnerId *string `json:"owner_id,omitempty"`
	Pin *string `json:"pin,omitempty"`
	RequirePin *bool `json:"require_pin,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

type _ServiceVOIPVoicemailAddEditData ServiceVOIPVoicemailAddEditData

// NewServiceVOIPVoicemailAddEditData instantiates a new ServiceVOIPVoicemailAddEditData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVOIPVoicemailAddEditData(mailbox string, name string) *ServiceVOIPVoicemailAddEditData {
	this := ServiceVOIPVoicemailAddEditData{}
	this.Mailbox = mailbox
	this.Name = name
	return &this
}

// NewServiceVOIPVoicemailAddEditDataWithDefaults instantiates a new ServiceVOIPVoicemailAddEditData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVOIPVoicemailAddEditDataWithDefaults() *ServiceVOIPVoicemailAddEditData {
	this := ServiceVOIPVoicemailAddEditData{}
	return &this
}

// GetMailbox returns the Mailbox field value
func (o *ServiceVOIPVoicemailAddEditData) GetMailbox() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mailbox
}

// GetMailboxOk returns a tuple with the Mailbox field value
// and a boolean to check if the value has been set.
func (o *ServiceVOIPVoicemailAddEditData) GetMailboxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mailbox, true
}

// SetMailbox sets field value
func (o *ServiceVOIPVoicemailAddEditData) SetMailbox(v string) {
	o.Mailbox = v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *ServiceVOIPVoicemailAddEditData) GetMedia() ServiceVoicemailMedia {
	if o == nil || IsNil(o.Media) {
		var ret ServiceVoicemailMedia
		return ret
	}
	return *o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPVoicemailAddEditData) GetMediaOk() (*ServiceVoicemailMedia, bool) {
	if o == nil || IsNil(o.Media) {
		return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *ServiceVOIPVoicemailAddEditData) HasMedia() bool {
	if o != nil && !IsNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given ServiceVoicemailMedia and assigns it to the Media field.
func (o *ServiceVOIPVoicemailAddEditData) SetMedia(v ServiceVoicemailMedia) {
	o.Media = &v
}

// GetMediaExtension returns the MediaExtension field value if set, zero value otherwise.
func (o *ServiceVOIPVoicemailAddEditData) GetMediaExtension() string {
	if o == nil || IsNil(o.MediaExtension) {
		var ret string
		return ret
	}
	return *o.MediaExtension
}

// GetMediaExtensionOk returns a tuple with the MediaExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPVoicemailAddEditData) GetMediaExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.MediaExtension) {
		return nil, false
	}
	return o.MediaExtension, true
}

// HasMediaExtension returns a boolean if a field has been set.
func (o *ServiceVOIPVoicemailAddEditData) HasMediaExtension() bool {
	if o != nil && !IsNil(o.MediaExtension) {
		return true
	}

	return false
}

// SetMediaExtension gets a reference to the given string and assigns it to the MediaExtension field.
func (o *ServiceVOIPVoicemailAddEditData) SetMediaExtension(v string) {
	o.MediaExtension = &v
}

// GetName returns the Name field value
func (o *ServiceVOIPVoicemailAddEditData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceVOIPVoicemailAddEditData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceVOIPVoicemailAddEditData) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ServiceVOIPVoicemailAddEditData) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPVoicemailAddEditData) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ServiceVOIPVoicemailAddEditData) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ServiceVOIPVoicemailAddEditData) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *ServiceVOIPVoicemailAddEditData) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPVoicemailAddEditData) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *ServiceVOIPVoicemailAddEditData) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *ServiceVOIPVoicemailAddEditData) SetPin(v string) {
	o.Pin = &v
}

// GetRequirePin returns the RequirePin field value if set, zero value otherwise.
func (o *ServiceVOIPVoicemailAddEditData) GetRequirePin() bool {
	if o == nil || IsNil(o.RequirePin) {
		var ret bool
		return ret
	}
	return *o.RequirePin
}

// GetRequirePinOk returns a tuple with the RequirePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPVoicemailAddEditData) GetRequirePinOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirePin) {
		return nil, false
	}
	return o.RequirePin, true
}

// HasRequirePin returns a boolean if a field has been set.
func (o *ServiceVOIPVoicemailAddEditData) HasRequirePin() bool {
	if o != nil && !IsNil(o.RequirePin) {
		return true
	}

	return false
}

// SetRequirePin gets a reference to the given bool and assigns it to the RequirePin field.
func (o *ServiceVOIPVoicemailAddEditData) SetRequirePin(v bool) {
	o.RequirePin = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ServiceVOIPVoicemailAddEditData) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPVoicemailAddEditData) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ServiceVOIPVoicemailAddEditData) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ServiceVOIPVoicemailAddEditData) SetTimezone(v string) {
	o.Timezone = &v
}

func (o ServiceVOIPVoicemailAddEditData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVOIPVoicemailAddEditData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mailbox"] = o.Mailbox
	if !IsNil(o.Media) {
		toSerialize["media"] = o.Media
	}
	if !IsNil(o.MediaExtension) {
		toSerialize["media_extension"] = o.MediaExtension
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !IsNil(o.RequirePin) {
		toSerialize["require_pin"] = o.RequirePin
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	return toSerialize, nil
}

func (o *ServiceVOIPVoicemailAddEditData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mailbox",
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

	varServiceVOIPVoicemailAddEditData := _ServiceVOIPVoicemailAddEditData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceVOIPVoicemailAddEditData)

	if err != nil {
		return err
	}

	*o = ServiceVOIPVoicemailAddEditData(varServiceVOIPVoicemailAddEditData)

	return err
}

type NullableServiceVOIPVoicemailAddEditData struct {
	value *ServiceVOIPVoicemailAddEditData
	isSet bool
}

func (v NullableServiceVOIPVoicemailAddEditData) Get() *ServiceVOIPVoicemailAddEditData {
	return v.value
}

func (v *NullableServiceVOIPVoicemailAddEditData) Set(val *ServiceVOIPVoicemailAddEditData) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVOIPVoicemailAddEditData) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVOIPVoicemailAddEditData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVOIPVoicemailAddEditData(val *ServiceVOIPVoicemailAddEditData) *NullableServiceVOIPVoicemailAddEditData {
	return &NullableServiceVOIPVoicemailAddEditData{value: val, isSet: true}
}

func (v NullableServiceVOIPVoicemailAddEditData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVOIPVoicemailAddEditData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


