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

// checks if the ServiceSystemStatusMessagingService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceSystemStatusMessagingService{}

// ServiceSystemStatusMessagingService struct for ServiceSystemStatusMessagingService
type ServiceSystemStatusMessagingService struct {
	MessagingServer *string `json:"messaging_server,omitempty"`
}

// NewServiceSystemStatusMessagingService instantiates a new ServiceSystemStatusMessagingService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSystemStatusMessagingService() *ServiceSystemStatusMessagingService {
	this := ServiceSystemStatusMessagingService{}
	return &this
}

// NewServiceSystemStatusMessagingServiceWithDefaults instantiates a new ServiceSystemStatusMessagingService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSystemStatusMessagingServiceWithDefaults() *ServiceSystemStatusMessagingService {
	this := ServiceSystemStatusMessagingService{}
	return &this
}

// GetMessagingServer returns the MessagingServer field value if set, zero value otherwise.
func (o *ServiceSystemStatusMessagingService) GetMessagingServer() string {
	if o == nil || IsNil(o.MessagingServer) {
		var ret string
		return ret
	}
	return *o.MessagingServer
}

// GetMessagingServerOk returns a tuple with the MessagingServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSystemStatusMessagingService) GetMessagingServerOk() (*string, bool) {
	if o == nil || IsNil(o.MessagingServer) {
		return nil, false
	}
	return o.MessagingServer, true
}

// HasMessagingServer returns a boolean if a field has been set.
func (o *ServiceSystemStatusMessagingService) HasMessagingServer() bool {
	if o != nil && !IsNil(o.MessagingServer) {
		return true
	}

	return false
}

// SetMessagingServer gets a reference to the given string and assigns it to the MessagingServer field.
func (o *ServiceSystemStatusMessagingService) SetMessagingServer(v string) {
	o.MessagingServer = &v
}

func (o ServiceSystemStatusMessagingService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceSystemStatusMessagingService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessagingServer) {
		toSerialize["messaging_server"] = o.MessagingServer
	}
	return toSerialize, nil
}

type NullableServiceSystemStatusMessagingService struct {
	value *ServiceSystemStatusMessagingService
	isSet bool
}

func (v NullableServiceSystemStatusMessagingService) Get() *ServiceSystemStatusMessagingService {
	return v.value
}

func (v *NullableServiceSystemStatusMessagingService) Set(val *ServiceSystemStatusMessagingService) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSystemStatusMessagingService) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSystemStatusMessagingService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSystemStatusMessagingService(val *ServiceSystemStatusMessagingService) *NullableServiceSystemStatusMessagingService {
	return &NullableServiceSystemStatusMessagingService{value: val, isSet: true}
}

func (v NullableServiceSystemStatusMessagingService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSystemStatusMessagingService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


