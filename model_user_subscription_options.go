/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.2.2
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// UserSubscriptionOptions struct for UserSubscriptionOptions
type UserSubscriptionOptions struct {
	RetainPreviousOwner *bool `json:"retain_previous_owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSubscriptionOptions UserSubscriptionOptions

// NewUserSubscriptionOptions instantiates a new UserSubscriptionOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSubscriptionOptions() *UserSubscriptionOptions {
	this := UserSubscriptionOptions{}
	return &this
}

// NewUserSubscriptionOptionsWithDefaults instantiates a new UserSubscriptionOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSubscriptionOptionsWithDefaults() *UserSubscriptionOptions {
	this := UserSubscriptionOptions{}
	return &this
}

// GetRetainPreviousOwner returns the RetainPreviousOwner field value if set, zero value otherwise.
func (o *UserSubscriptionOptions) GetRetainPreviousOwner() bool {
	if o == nil || o.RetainPreviousOwner == nil {
		var ret bool
		return ret
	}
	return *o.RetainPreviousOwner
}

// GetRetainPreviousOwnerOk returns a tuple with the RetainPreviousOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscriptionOptions) GetRetainPreviousOwnerOk() (*bool, bool) {
	if o == nil || o.RetainPreviousOwner == nil {
		return nil, false
	}
	return o.RetainPreviousOwner, true
}

// HasRetainPreviousOwner returns a boolean if a field has been set.
func (o *UserSubscriptionOptions) HasRetainPreviousOwner() bool {
	if o != nil && o.RetainPreviousOwner != nil {
		return true
	}

	return false
}

// SetRetainPreviousOwner gets a reference to the given bool and assigns it to the RetainPreviousOwner field.
func (o *UserSubscriptionOptions) SetRetainPreviousOwner(v bool) {
	o.RetainPreviousOwner = &v
}

func (o UserSubscriptionOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RetainPreviousOwner != nil {
		toSerialize["retain_previous_owner"] = o.RetainPreviousOwner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSubscriptionOptions) UnmarshalJSON(bytes []byte) (err error) {
	varUserSubscriptionOptions := _UserSubscriptionOptions{}

	if err = json.Unmarshal(bytes, &varUserSubscriptionOptions); err == nil {
		*o = UserSubscriptionOptions(varUserSubscriptionOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "retain_previous_owner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSubscriptionOptions struct {
	value *UserSubscriptionOptions
	isSet bool
}

func (v NullableUserSubscriptionOptions) Get() *UserSubscriptionOptions {
	return v.value
}

func (v *NullableUserSubscriptionOptions) Set(val *UserSubscriptionOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSubscriptionOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSubscriptionOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSubscriptionOptions(val *UserSubscriptionOptions) *NullableUserSubscriptionOptions {
	return &NullableUserSubscriptionOptions{value: val, isSet: true}
}

func (v NullableUserSubscriptionOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSubscriptionOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


