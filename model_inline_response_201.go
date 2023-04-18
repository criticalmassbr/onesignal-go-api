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

// InlineResponse201 struct for InlineResponse201
type InlineResponse201 struct {
	Subscription *SubscriptionObject `json:"subscription,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineResponse201 InlineResponse201

// NewInlineResponse201 instantiates a new InlineResponse201 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse201() *InlineResponse201 {
	this := InlineResponse201{}
	return &this
}

// NewInlineResponse201WithDefaults instantiates a new InlineResponse201 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse201WithDefaults() *InlineResponse201 {
	this := InlineResponse201{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *InlineResponse201) GetSubscription() SubscriptionObject {
	if o == nil || o.Subscription == nil {
		var ret SubscriptionObject
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse201) GetSubscriptionOk() (*SubscriptionObject, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *InlineResponse201) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given SubscriptionObject and assigns it to the Subscription field.
func (o *InlineResponse201) SetSubscription(v SubscriptionObject) {
	o.Subscription = &v
}

func (o InlineResponse201) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineResponse201) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse201 := _InlineResponse201{}

	if err = json.Unmarshal(bytes, &varInlineResponse201); err == nil {
		*o = InlineResponse201(varInlineResponse201)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "subscription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineResponse201 struct {
	value *InlineResponse201
	isSet bool
}

func (v NullableInlineResponse201) Get() *InlineResponse201 {
	return v.value
}

func (v *NullableInlineResponse201) Set(val *InlineResponse201) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse201) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse201) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse201(val *InlineResponse201) *NullableInlineResponse201 {
	return &NullableInlineResponse201{value: val, isSet: true}
}

func (v NullableInlineResponse201) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse201) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


