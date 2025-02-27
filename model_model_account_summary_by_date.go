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

// checks if the ModelAccountSummaryByDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAccountSummaryByDate{}

// ModelAccountSummaryByDate struct for ModelAccountSummaryByDate
type ModelAccountSummaryByDate struct {
	Cost *float32 `json:"cost,omitempty"`
	Date *string `json:"date,omitempty"`
	TransactionCount *int32 `json:"transaction_count,omitempty"`
}

// NewModelAccountSummaryByDate instantiates a new ModelAccountSummaryByDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAccountSummaryByDate() *ModelAccountSummaryByDate {
	this := ModelAccountSummaryByDate{}
	return &this
}

// NewModelAccountSummaryByDateWithDefaults instantiates a new ModelAccountSummaryByDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAccountSummaryByDateWithDefaults() *ModelAccountSummaryByDate {
	this := ModelAccountSummaryByDate{}
	return &this
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *ModelAccountSummaryByDate) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAccountSummaryByDate) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *ModelAccountSummaryByDate) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *ModelAccountSummaryByDate) SetCost(v float32) {
	o.Cost = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ModelAccountSummaryByDate) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAccountSummaryByDate) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ModelAccountSummaryByDate) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *ModelAccountSummaryByDate) SetDate(v string) {
	o.Date = &v
}

// GetTransactionCount returns the TransactionCount field value if set, zero value otherwise.
func (o *ModelAccountSummaryByDate) GetTransactionCount() int32 {
	if o == nil || IsNil(o.TransactionCount) {
		var ret int32
		return ret
	}
	return *o.TransactionCount
}

// GetTransactionCountOk returns a tuple with the TransactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAccountSummaryByDate) GetTransactionCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TransactionCount) {
		return nil, false
	}
	return o.TransactionCount, true
}

// HasTransactionCount returns a boolean if a field has been set.
func (o *ModelAccountSummaryByDate) HasTransactionCount() bool {
	if o != nil && !IsNil(o.TransactionCount) {
		return true
	}

	return false
}

// SetTransactionCount gets a reference to the given int32 and assigns it to the TransactionCount field.
func (o *ModelAccountSummaryByDate) SetTransactionCount(v int32) {
	o.TransactionCount = &v
}

func (o ModelAccountSummaryByDate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAccountSummaryByDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.TransactionCount) {
		toSerialize["transaction_count"] = o.TransactionCount
	}
	return toSerialize, nil
}

type NullableModelAccountSummaryByDate struct {
	value *ModelAccountSummaryByDate
	isSet bool
}

func (v NullableModelAccountSummaryByDate) Get() *ModelAccountSummaryByDate {
	return v.value
}

func (v *NullableModelAccountSummaryByDate) Set(val *ModelAccountSummaryByDate) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAccountSummaryByDate) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAccountSummaryByDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAccountSummaryByDate(val *ModelAccountSummaryByDate) *NullableModelAccountSummaryByDate {
	return &NullableModelAccountSummaryByDate{value: val, isSet: true}
}

func (v NullableModelAccountSummaryByDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAccountSummaryByDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


