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

// checks if the ServiceDocsWebhookGetSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDocsWebhookGetSingle{}

// ServiceDocsWebhookGetSingle struct for ServiceDocsWebhookGetSingle
type ServiceDocsWebhookGetSingle struct {
	Data *ModelAccountWebhook `json:"data,omitempty"`
	RequestId *string `json:"request_id,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewServiceDocsWebhookGetSingle instantiates a new ServiceDocsWebhookGetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDocsWebhookGetSingle() *ServiceDocsWebhookGetSingle {
	this := ServiceDocsWebhookGetSingle{}
	return &this
}

// NewServiceDocsWebhookGetSingleWithDefaults instantiates a new ServiceDocsWebhookGetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDocsWebhookGetSingleWithDefaults() *ServiceDocsWebhookGetSingle {
	this := ServiceDocsWebhookGetSingle{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceDocsWebhookGetSingle) GetData() ModelAccountWebhook {
	if o == nil || IsNil(o.Data) {
		var ret ModelAccountWebhook
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsWebhookGetSingle) GetDataOk() (*ModelAccountWebhook, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceDocsWebhookGetSingle) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ModelAccountWebhook and assigns it to the Data field.
func (o *ServiceDocsWebhookGetSingle) SetData(v ModelAccountWebhook) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ServiceDocsWebhookGetSingle) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsWebhookGetSingle) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ServiceDocsWebhookGetSingle) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ServiceDocsWebhookGetSingle) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ServiceDocsWebhookGetSingle) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsWebhookGetSingle) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ServiceDocsWebhookGetSingle) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ServiceDocsWebhookGetSingle) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ServiceDocsWebhookGetSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDocsWebhookGetSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableServiceDocsWebhookGetSingle struct {
	value *ServiceDocsWebhookGetSingle
	isSet bool
}

func (v NullableServiceDocsWebhookGetSingle) Get() *ServiceDocsWebhookGetSingle {
	return v.value
}

func (v *NullableServiceDocsWebhookGetSingle) Set(val *ServiceDocsWebhookGetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDocsWebhookGetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDocsWebhookGetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDocsWebhookGetSingle(val *ServiceDocsWebhookGetSingle) *NullableServiceDocsWebhookGetSingle {
	return &NullableServiceDocsWebhookGetSingle{value: val, isSet: true}
}

func (v NullableServiceDocsWebhookGetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDocsWebhookGetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


