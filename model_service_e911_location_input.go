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

// checks if the ServiceE911LocationInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceE911LocationInput{}

// ServiceE911LocationInput struct for ServiceE911LocationInput
type ServiceE911LocationInput struct {
	Address1 string `json:"address_1"`
	Address2 *string `json:"address_2,omitempty"`
	Community string `json:"community"`
	PlusFour *string `json:"plus_four,omitempty"`
	PostalCode string `json:"postal_code"`
	State string `json:"state"`
	Type *string `json:"type,omitempty"`
}

type _ServiceE911LocationInput ServiceE911LocationInput

// NewServiceE911LocationInput instantiates a new ServiceE911LocationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceE911LocationInput(address1 string, community string, postalCode string, state string) *ServiceE911LocationInput {
	this := ServiceE911LocationInput{}
	this.Address1 = address1
	this.Community = community
	this.PostalCode = postalCode
	this.State = state
	return &this
}

// NewServiceE911LocationInputWithDefaults instantiates a new ServiceE911LocationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceE911LocationInputWithDefaults() *ServiceE911LocationInput {
	this := ServiceE911LocationInput{}
	return &this
}

// GetAddress1 returns the Address1 field value
func (o *ServiceE911LocationInput) GetAddress1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationInput) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address1, true
}

// SetAddress1 sets field value
func (o *ServiceE911LocationInput) SetAddress1(v string) {
	o.Address1 = v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *ServiceE911LocationInput) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationInput) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *ServiceE911LocationInput) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *ServiceE911LocationInput) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCommunity returns the Community field value
func (o *ServiceE911LocationInput) GetCommunity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Community
}

// GetCommunityOk returns a tuple with the Community field value
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationInput) GetCommunityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Community, true
}

// SetCommunity sets field value
func (o *ServiceE911LocationInput) SetCommunity(v string) {
	o.Community = v
}

// GetPlusFour returns the PlusFour field value if set, zero value otherwise.
func (o *ServiceE911LocationInput) GetPlusFour() string {
	if o == nil || IsNil(o.PlusFour) {
		var ret string
		return ret
	}
	return *o.PlusFour
}

// GetPlusFourOk returns a tuple with the PlusFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationInput) GetPlusFourOk() (*string, bool) {
	if o == nil || IsNil(o.PlusFour) {
		return nil, false
	}
	return o.PlusFour, true
}

// HasPlusFour returns a boolean if a field has been set.
func (o *ServiceE911LocationInput) HasPlusFour() bool {
	if o != nil && !IsNil(o.PlusFour) {
		return true
	}

	return false
}

// SetPlusFour gets a reference to the given string and assigns it to the PlusFour field.
func (o *ServiceE911LocationInput) SetPlusFour(v string) {
	o.PlusFour = &v
}

// GetPostalCode returns the PostalCode field value
func (o *ServiceE911LocationInput) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationInput) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *ServiceE911LocationInput) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetState returns the State field value
func (o *ServiceE911LocationInput) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationInput) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ServiceE911LocationInput) SetState(v string) {
	o.State = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceE911LocationInput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationInput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceE911LocationInput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceE911LocationInput) SetType(v string) {
	o.Type = &v
}

func (o ServiceE911LocationInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceE911LocationInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address_1"] = o.Address1
	if !IsNil(o.Address2) {
		toSerialize["address_2"] = o.Address2
	}
	toSerialize["community"] = o.Community
	if !IsNil(o.PlusFour) {
		toSerialize["plus_four"] = o.PlusFour
	}
	toSerialize["postal_code"] = o.PostalCode
	toSerialize["state"] = o.State
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *ServiceE911LocationInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address_1",
		"community",
		"postal_code",
		"state",
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

	varServiceE911LocationInput := _ServiceE911LocationInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceE911LocationInput)

	if err != nil {
		return err
	}

	*o = ServiceE911LocationInput(varServiceE911LocationInput)

	return err
}

type NullableServiceE911LocationInput struct {
	value *ServiceE911LocationInput
	isSet bool
}

func (v NullableServiceE911LocationInput) Get() *ServiceE911LocationInput {
	return v.value
}

func (v *NullableServiceE911LocationInput) Set(val *ServiceE911LocationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceE911LocationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceE911LocationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceE911LocationInput(val *ServiceE911LocationInput) *NullableServiceE911LocationInput {
	return &NullableServiceE911LocationInput{value: val, isSet: true}
}

func (v NullableServiceE911LocationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceE911LocationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


