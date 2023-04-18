/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.2.1
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// DeletePlayerSuccessResponse struct for DeletePlayerSuccessResponse
type DeletePlayerSuccessResponse struct {
	Success *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeletePlayerSuccessResponse DeletePlayerSuccessResponse

// NewDeletePlayerSuccessResponse instantiates a new DeletePlayerSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePlayerSuccessResponse() *DeletePlayerSuccessResponse {
	this := DeletePlayerSuccessResponse{}
	return &this
}

// NewDeletePlayerSuccessResponseWithDefaults instantiates a new DeletePlayerSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePlayerSuccessResponseWithDefaults() *DeletePlayerSuccessResponse {
	this := DeletePlayerSuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeletePlayerSuccessResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletePlayerSuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeletePlayerSuccessResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeletePlayerSuccessResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o DeletePlayerSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeletePlayerSuccessResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDeletePlayerSuccessResponse := _DeletePlayerSuccessResponse{}

	if err = json.Unmarshal(bytes, &varDeletePlayerSuccessResponse); err == nil {
		*o = DeletePlayerSuccessResponse(varDeletePlayerSuccessResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeletePlayerSuccessResponse struct {
	value *DeletePlayerSuccessResponse
	isSet bool
}

func (v NullableDeletePlayerSuccessResponse) Get() *DeletePlayerSuccessResponse {
	return v.value
}

func (v *NullableDeletePlayerSuccessResponse) Set(val *DeletePlayerSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePlayerSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePlayerSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePlayerSuccessResponse(val *DeletePlayerSuccessResponse) *NullableDeletePlayerSuccessResponse {
	return &NullableDeletePlayerSuccessResponse{value: val, isSet: true}
}

func (v NullableDeletePlayerSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePlayerSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


