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

// checks if the ServiceVOIPDeviceAddEdit2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVOIPDeviceAddEdit2{}

// ServiceVOIPDeviceAddEdit2 struct for ServiceVOIPDeviceAddEdit2
type ServiceVOIPDeviceAddEdit2 struct {
	CallForward *ServiceCallForward `json:"call_forward,omitempty"`
	CallRecording *ServiceCallRecordingSettings `json:"call_recording,omitempty"`
	CallerId *ServiceVOIPDeviceAddEdit3c `json:"caller_id,omitempty"`
	DeviceType *string `json:"device_type,omitempty"`
	DoNotDisturb *ServiceVOIPSharedDoNotDisturb `json:"do_not_disturb,omitempty"`
	// cannot use required, else it has to be true and false is not allowed
	Enabled *bool `json:"enabled,omitempty"`
	// dont use mac, it enforces :, which voip does not like
	MacAddress *string `json:"mac_address,omitempty"`
	MusicOnHold *ServiceMusicOnHold `json:"music_on_hold,omitempty"`
	Name string `json:"name"`
	// json omitempty is needed else voip fails on \"\" for owner_id
	OwnerId *string `json:"owner_id,omitempty"`
	Sip ServiceVOIPDeviceAddEdit3a `json:"sip"`
}

type _ServiceVOIPDeviceAddEdit2 ServiceVOIPDeviceAddEdit2

// NewServiceVOIPDeviceAddEdit2 instantiates a new ServiceVOIPDeviceAddEdit2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVOIPDeviceAddEdit2(name string, sip ServiceVOIPDeviceAddEdit3a) *ServiceVOIPDeviceAddEdit2 {
	this := ServiceVOIPDeviceAddEdit2{}
	this.Name = name
	this.Sip = sip
	return &this
}

// NewServiceVOIPDeviceAddEdit2WithDefaults instantiates a new ServiceVOIPDeviceAddEdit2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVOIPDeviceAddEdit2WithDefaults() *ServiceVOIPDeviceAddEdit2 {
	this := ServiceVOIPDeviceAddEdit2{}
	return &this
}

// GetCallForward returns the CallForward field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit2) GetCallForward() ServiceCallForward {
	if o == nil || IsNil(o.CallForward) {
		var ret ServiceCallForward
		return ret
	}
	return *o.CallForward
}

// GetCallForwardOk returns a tuple with the CallForward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetCallForwardOk() (*ServiceCallForward, bool) {
	if o == nil || IsNil(o.CallForward) {
		return nil, false
	}
	return o.CallForward, true
}

// HasCallForward returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit2) HasCallForward() bool {
	if o != nil && !IsNil(o.CallForward) {
		return true
	}

	return false
}

// SetCallForward gets a reference to the given ServiceCallForward and assigns it to the CallForward field.
func (o *ServiceVOIPDeviceAddEdit2) SetCallForward(v ServiceCallForward) {
	o.CallForward = &v
}

// GetCallRecording returns the CallRecording field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit2) GetCallRecording() ServiceCallRecordingSettings {
	if o == nil || IsNil(o.CallRecording) {
		var ret ServiceCallRecordingSettings
		return ret
	}
	return *o.CallRecording
}

// GetCallRecordingOk returns a tuple with the CallRecording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetCallRecordingOk() (*ServiceCallRecordingSettings, bool) {
	if o == nil || IsNil(o.CallRecording) {
		return nil, false
	}
	return o.CallRecording, true
}

// HasCallRecording returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit2) HasCallRecording() bool {
	if o != nil && !IsNil(o.CallRecording) {
		return true
	}

	return false
}

// SetCallRecording gets a reference to the given ServiceCallRecordingSettings and assigns it to the CallRecording field.
func (o *ServiceVOIPDeviceAddEdit2) SetCallRecording(v ServiceCallRecordingSettings) {
	o.CallRecording = &v
}

// GetCallerId returns the CallerId field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit2) GetCallerId() ServiceVOIPDeviceAddEdit3c {
	if o == nil || IsNil(o.CallerId) {
		var ret ServiceVOIPDeviceAddEdit3c
		return ret
	}
	return *o.CallerId
}

// GetCallerIdOk returns a tuple with the CallerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetCallerIdOk() (*ServiceVOIPDeviceAddEdit3c, bool) {
	if o == nil || IsNil(o.CallerId) {
		return nil, false
	}
	return o.CallerId, true
}

// HasCallerId returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit2) HasCallerId() bool {
	if o != nil && !IsNil(o.CallerId) {
		return true
	}

	return false
}

// SetCallerId gets a reference to the given ServiceVOIPDeviceAddEdit3c and assigns it to the CallerId field.
func (o *ServiceVOIPDeviceAddEdit2) SetCallerId(v ServiceVOIPDeviceAddEdit3c) {
	o.CallerId = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit2) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit2) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *ServiceVOIPDeviceAddEdit2) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetDoNotDisturb returns the DoNotDisturb field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit2) GetDoNotDisturb() ServiceVOIPSharedDoNotDisturb {
	if o == nil || IsNil(o.DoNotDisturb) {
		var ret ServiceVOIPSharedDoNotDisturb
		return ret
	}
	return *o.DoNotDisturb
}

// GetDoNotDisturbOk returns a tuple with the DoNotDisturb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetDoNotDisturbOk() (*ServiceVOIPSharedDoNotDisturb, bool) {
	if o == nil || IsNil(o.DoNotDisturb) {
		return nil, false
	}
	return o.DoNotDisturb, true
}

// HasDoNotDisturb returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit2) HasDoNotDisturb() bool {
	if o != nil && !IsNil(o.DoNotDisturb) {
		return true
	}

	return false
}

// SetDoNotDisturb gets a reference to the given ServiceVOIPSharedDoNotDisturb and assigns it to the DoNotDisturb field.
func (o *ServiceVOIPDeviceAddEdit2) SetDoNotDisturb(v ServiceVOIPSharedDoNotDisturb) {
	o.DoNotDisturb = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit2) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit2) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ServiceVOIPDeviceAddEdit2) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit2) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit2) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *ServiceVOIPDeviceAddEdit2) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMusicOnHold returns the MusicOnHold field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit2) GetMusicOnHold() ServiceMusicOnHold {
	if o == nil || IsNil(o.MusicOnHold) {
		var ret ServiceMusicOnHold
		return ret
	}
	return *o.MusicOnHold
}

// GetMusicOnHoldOk returns a tuple with the MusicOnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetMusicOnHoldOk() (*ServiceMusicOnHold, bool) {
	if o == nil || IsNil(o.MusicOnHold) {
		return nil, false
	}
	return o.MusicOnHold, true
}

// HasMusicOnHold returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit2) HasMusicOnHold() bool {
	if o != nil && !IsNil(o.MusicOnHold) {
		return true
	}

	return false
}

// SetMusicOnHold gets a reference to the given ServiceMusicOnHold and assigns it to the MusicOnHold field.
func (o *ServiceVOIPDeviceAddEdit2) SetMusicOnHold(v ServiceMusicOnHold) {
	o.MusicOnHold = &v
}

// GetName returns the Name field value
func (o *ServiceVOIPDeviceAddEdit2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceVOIPDeviceAddEdit2) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit2) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit2) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ServiceVOIPDeviceAddEdit2) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetSip returns the Sip field value
func (o *ServiceVOIPDeviceAddEdit2) GetSip() ServiceVOIPDeviceAddEdit3a {
	if o == nil {
		var ret ServiceVOIPDeviceAddEdit3a
		return ret
	}

	return o.Sip
}

// GetSipOk returns a tuple with the Sip field value
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit2) GetSipOk() (*ServiceVOIPDeviceAddEdit3a, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sip, true
}

// SetSip sets field value
func (o *ServiceVOIPDeviceAddEdit2) SetSip(v ServiceVOIPDeviceAddEdit3a) {
	o.Sip = v
}

func (o ServiceVOIPDeviceAddEdit2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVOIPDeviceAddEdit2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallForward) {
		toSerialize["call_forward"] = o.CallForward
	}
	if !IsNil(o.CallRecording) {
		toSerialize["call_recording"] = o.CallRecording
	}
	if !IsNil(o.CallerId) {
		toSerialize["caller_id"] = o.CallerId
	}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if !IsNil(o.DoNotDisturb) {
		toSerialize["do_not_disturb"] = o.DoNotDisturb
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MacAddress) {
		toSerialize["mac_address"] = o.MacAddress
	}
	if !IsNil(o.MusicOnHold) {
		toSerialize["music_on_hold"] = o.MusicOnHold
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	toSerialize["sip"] = o.Sip
	return toSerialize, nil
}

func (o *ServiceVOIPDeviceAddEdit2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"sip",
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

	varServiceVOIPDeviceAddEdit2 := _ServiceVOIPDeviceAddEdit2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceVOIPDeviceAddEdit2)

	if err != nil {
		return err
	}

	*o = ServiceVOIPDeviceAddEdit2(varServiceVOIPDeviceAddEdit2)

	return err
}

type NullableServiceVOIPDeviceAddEdit2 struct {
	value *ServiceVOIPDeviceAddEdit2
	isSet bool
}

func (v NullableServiceVOIPDeviceAddEdit2) Get() *ServiceVOIPDeviceAddEdit2 {
	return v.value
}

func (v *NullableServiceVOIPDeviceAddEdit2) Set(val *ServiceVOIPDeviceAddEdit2) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVOIPDeviceAddEdit2) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVOIPDeviceAddEdit2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVOIPDeviceAddEdit2(val *ServiceVOIPDeviceAddEdit2) *NullableServiceVOIPDeviceAddEdit2 {
	return &NullableServiceVOIPDeviceAddEdit2{value: val, isSet: true}
}

func (v NullableServiceVOIPDeviceAddEdit2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVOIPDeviceAddEdit2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


