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

// checks if the ServiceE911LocationURI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceE911LocationURI{}

// ServiceE911LocationURI struct for ServiceE911LocationURI
type ServiceE911LocationURI struct {
	ActivatedTime *string `json:"activatedTime,omitempty"`
	Address1 *string `json:"address1,omitempty"`
	Address2 *string `json:"address2,omitempty"`
	CallerName *string `json:"callerName,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Community *string `json:"community,omitempty"`
	CustomerOrderID *string `json:"customerOrderID,omitempty"`
	Latitude *float32 `json:"latitude,omitempty"`
	LegacyData *ServiceE911LocationURILegacyData `json:"legacyData,omitempty"`
	LocationID *string `json:"locationID,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
	PlusFour *string `json:"plusFour,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	State *string `json:"state,omitempty"`
	Status *ServiceE911LocationURIStatus `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
}

// NewServiceE911LocationURI instantiates a new ServiceE911LocationURI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceE911LocationURI() *ServiceE911LocationURI {
	this := ServiceE911LocationURI{}
	return &this
}

// NewServiceE911LocationURIWithDefaults instantiates a new ServiceE911LocationURI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceE911LocationURIWithDefaults() *ServiceE911LocationURI {
	this := ServiceE911LocationURI{}
	return &this
}

// GetActivatedTime returns the ActivatedTime field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetActivatedTime() string {
	if o == nil || IsNil(o.ActivatedTime) {
		var ret string
		return ret
	}
	return *o.ActivatedTime
}

// GetActivatedTimeOk returns a tuple with the ActivatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetActivatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ActivatedTime) {
		return nil, false
	}
	return o.ActivatedTime, true
}

// HasActivatedTime returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasActivatedTime() bool {
	if o != nil && !IsNil(o.ActivatedTime) {
		return true
	}

	return false
}

// SetActivatedTime gets a reference to the given string and assigns it to the ActivatedTime field.
func (o *ServiceE911LocationURI) SetActivatedTime(v string) {
	o.ActivatedTime = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *ServiceE911LocationURI) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *ServiceE911LocationURI) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCallerName returns the CallerName field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetCallerName() string {
	if o == nil || IsNil(o.CallerName) {
		var ret string
		return ret
	}
	return *o.CallerName
}

// GetCallerNameOk returns a tuple with the CallerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetCallerNameOk() (*string, bool) {
	if o == nil || IsNil(o.CallerName) {
		return nil, false
	}
	return o.CallerName, true
}

// HasCallerName returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasCallerName() bool {
	if o != nil && !IsNil(o.CallerName) {
		return true
	}

	return false
}

// SetCallerName gets a reference to the given string and assigns it to the CallerName field.
func (o *ServiceE911LocationURI) SetCallerName(v string) {
	o.CallerName = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ServiceE911LocationURI) SetComments(v string) {
	o.Comments = &v
}

// GetCommunity returns the Community field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetCommunity() string {
	if o == nil || IsNil(o.Community) {
		var ret string
		return ret
	}
	return *o.Community
}

// GetCommunityOk returns a tuple with the Community field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetCommunityOk() (*string, bool) {
	if o == nil || IsNil(o.Community) {
		return nil, false
	}
	return o.Community, true
}

// HasCommunity returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasCommunity() bool {
	if o != nil && !IsNil(o.Community) {
		return true
	}

	return false
}

// SetCommunity gets a reference to the given string and assigns it to the Community field.
func (o *ServiceE911LocationURI) SetCommunity(v string) {
	o.Community = &v
}

// GetCustomerOrderID returns the CustomerOrderID field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetCustomerOrderID() string {
	if o == nil || IsNil(o.CustomerOrderID) {
		var ret string
		return ret
	}
	return *o.CustomerOrderID
}

// GetCustomerOrderIDOk returns a tuple with the CustomerOrderID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetCustomerOrderIDOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderID) {
		return nil, false
	}
	return o.CustomerOrderID, true
}

// HasCustomerOrderID returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasCustomerOrderID() bool {
	if o != nil && !IsNil(o.CustomerOrderID) {
		return true
	}

	return false
}

// SetCustomerOrderID gets a reference to the given string and assigns it to the CustomerOrderID field.
func (o *ServiceE911LocationURI) SetCustomerOrderID(v string) {
	o.CustomerOrderID = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetLatitude() float32 {
	if o == nil || IsNil(o.Latitude) {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetLatitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *ServiceE911LocationURI) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLegacyData returns the LegacyData field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetLegacyData() ServiceE911LocationURILegacyData {
	if o == nil || IsNil(o.LegacyData) {
		var ret ServiceE911LocationURILegacyData
		return ret
	}
	return *o.LegacyData
}

// GetLegacyDataOk returns a tuple with the LegacyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetLegacyDataOk() (*ServiceE911LocationURILegacyData, bool) {
	if o == nil || IsNil(o.LegacyData) {
		return nil, false
	}
	return o.LegacyData, true
}

// HasLegacyData returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasLegacyData() bool {
	if o != nil && !IsNil(o.LegacyData) {
		return true
	}

	return false
}

// SetLegacyData gets a reference to the given ServiceE911LocationURILegacyData and assigns it to the LegacyData field.
func (o *ServiceE911LocationURI) SetLegacyData(v ServiceE911LocationURILegacyData) {
	o.LegacyData = &v
}

// GetLocationID returns the LocationID field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetLocationID() string {
	if o == nil || IsNil(o.LocationID) {
		var ret string
		return ret
	}
	return *o.LocationID
}

// GetLocationIDOk returns a tuple with the LocationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetLocationIDOk() (*string, bool) {
	if o == nil || IsNil(o.LocationID) {
		return nil, false
	}
	return o.LocationID, true
}

// HasLocationID returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasLocationID() bool {
	if o != nil && !IsNil(o.LocationID) {
		return true
	}

	return false
}

// SetLocationID gets a reference to the given string and assigns it to the LocationID field.
func (o *ServiceE911LocationURI) SetLocationID(v string) {
	o.LocationID = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetLongitude() float32 {
	if o == nil || IsNil(o.Longitude) {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetLongitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *ServiceE911LocationURI) SetLongitude(v float32) {
	o.Longitude = &v
}

// GetPlusFour returns the PlusFour field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetPlusFour() string {
	if o == nil || IsNil(o.PlusFour) {
		var ret string
		return ret
	}
	return *o.PlusFour
}

// GetPlusFourOk returns a tuple with the PlusFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetPlusFourOk() (*string, bool) {
	if o == nil || IsNil(o.PlusFour) {
		return nil, false
	}
	return o.PlusFour, true
}

// HasPlusFour returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasPlusFour() bool {
	if o != nil && !IsNil(o.PlusFour) {
		return true
	}

	return false
}

// SetPlusFour gets a reference to the given string and assigns it to the PlusFour field.
func (o *ServiceE911LocationURI) SetPlusFour(v string) {
	o.PlusFour = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *ServiceE911LocationURI) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ServiceE911LocationURI) SetState(v string) {
	o.State = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetStatus() ServiceE911LocationURIStatus {
	if o == nil || IsNil(o.Status) {
		var ret ServiceE911LocationURIStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetStatusOk() (*ServiceE911LocationURIStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ServiceE911LocationURIStatus and assigns it to the Status field.
func (o *ServiceE911LocationURI) SetStatus(v ServiceE911LocationURIStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceE911LocationURI) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *ServiceE911LocationURI) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURI) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *ServiceE911LocationURI) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *ServiceE911LocationURI) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

func (o ServiceE911LocationURI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceE911LocationURI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivatedTime) {
		toSerialize["activatedTime"] = o.ActivatedTime
	}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.CallerName) {
		toSerialize["callerName"] = o.CallerName
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Community) {
		toSerialize["community"] = o.Community
	}
	if !IsNil(o.CustomerOrderID) {
		toSerialize["customerOrderID"] = o.CustomerOrderID
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.LegacyData) {
		toSerialize["legacyData"] = o.LegacyData
	}
	if !IsNil(o.LocationID) {
		toSerialize["locationID"] = o.LocationID
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.PlusFour) {
		toSerialize["plusFour"] = o.PlusFour
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableServiceE911LocationURI struct {
	value *ServiceE911LocationURI
	isSet bool
}

func (v NullableServiceE911LocationURI) Get() *ServiceE911LocationURI {
	return v.value
}

func (v *NullableServiceE911LocationURI) Set(val *ServiceE911LocationURI) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceE911LocationURI) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceE911LocationURI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceE911LocationURI(val *ServiceE911LocationURI) *NullableServiceE911LocationURI {
	return &NullableServiceE911LocationURI{value: val, isSet: true}
}

func (v NullableServiceE911LocationURI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceE911LocationURI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


