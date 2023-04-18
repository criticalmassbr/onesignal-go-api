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

// PlatformDeliveryDataEmailAllOf struct for PlatformDeliveryDataEmailAllOf
type PlatformDeliveryDataEmailAllOf struct {
	// Number of times an email has been opened.
	Opened NullableInt32 `json:"opened,omitempty"`
	// Number of unique recipients who have opened your email.
	UniqueOpens NullableInt32 `json:"unique_opens,omitempty"`
	// Number of clicked links from your email. This can include the recipient clicking email links multiple times.
	Clicks NullableInt32 `json:"clicks,omitempty"`
	// Number of unique clicks that your recipients have made on links from your email.
	UniqueClicks NullableInt32 `json:"unique_clicks,omitempty"`
	// Number of recipients who registered as a hard or soft bounce and didn't receive your email.
	Bounced NullableInt32 `json:"bounced,omitempty"`
	// Number of recipients who reported this email as spam.
	ReportedSpam NullableInt32 `json:"reported_spam,omitempty"`
	// Number of recipients who opted out of your emails using the unsubscribe link in this email.
	Unsubscribed NullableInt32 `json:"unsubscribed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlatformDeliveryDataEmailAllOf PlatformDeliveryDataEmailAllOf

// NewPlatformDeliveryDataEmailAllOf instantiates a new PlatformDeliveryDataEmailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformDeliveryDataEmailAllOf() *PlatformDeliveryDataEmailAllOf {
	this := PlatformDeliveryDataEmailAllOf{}
	return &this
}

// NewPlatformDeliveryDataEmailAllOfWithDefaults instantiates a new PlatformDeliveryDataEmailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformDeliveryDataEmailAllOfWithDefaults() *PlatformDeliveryDataEmailAllOf {
	this := PlatformDeliveryDataEmailAllOf{}
	return &this
}

// GetOpened returns the Opened field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataEmailAllOf) GetOpened() int32 {
	if o == nil || o.Opened.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Opened.Get()
}

// GetOpenedOk returns a tuple with the Opened field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataEmailAllOf) GetOpenedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Opened.Get(), o.Opened.IsSet()
}

// HasOpened returns a boolean if a field has been set.
func (o *PlatformDeliveryDataEmailAllOf) HasOpened() bool {
	if o != nil && o.Opened.IsSet() {
		return true
	}

	return false
}

// SetOpened gets a reference to the given NullableInt32 and assigns it to the Opened field.
func (o *PlatformDeliveryDataEmailAllOf) SetOpened(v int32) {
	o.Opened.Set(&v)
}
// SetOpenedNil sets the value for Opened to be an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) SetOpenedNil() {
	o.Opened.Set(nil)
}

// UnsetOpened ensures that no value is present for Opened, not even an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) UnsetOpened() {
	o.Opened.Unset()
}

// GetUniqueOpens returns the UniqueOpens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataEmailAllOf) GetUniqueOpens() int32 {
	if o == nil || o.UniqueOpens.Get() == nil {
		var ret int32
		return ret
	}
	return *o.UniqueOpens.Get()
}

// GetUniqueOpensOk returns a tuple with the UniqueOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataEmailAllOf) GetUniqueOpensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniqueOpens.Get(), o.UniqueOpens.IsSet()
}

// HasUniqueOpens returns a boolean if a field has been set.
func (o *PlatformDeliveryDataEmailAllOf) HasUniqueOpens() bool {
	if o != nil && o.UniqueOpens.IsSet() {
		return true
	}

	return false
}

// SetUniqueOpens gets a reference to the given NullableInt32 and assigns it to the UniqueOpens field.
func (o *PlatformDeliveryDataEmailAllOf) SetUniqueOpens(v int32) {
	o.UniqueOpens.Set(&v)
}
// SetUniqueOpensNil sets the value for UniqueOpens to be an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) SetUniqueOpensNil() {
	o.UniqueOpens.Set(nil)
}

// UnsetUniqueOpens ensures that no value is present for UniqueOpens, not even an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) UnsetUniqueOpens() {
	o.UniqueOpens.Unset()
}

// GetClicks returns the Clicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataEmailAllOf) GetClicks() int32 {
	if o == nil || o.Clicks.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Clicks.Get()
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataEmailAllOf) GetClicksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clicks.Get(), o.Clicks.IsSet()
}

// HasClicks returns a boolean if a field has been set.
func (o *PlatformDeliveryDataEmailAllOf) HasClicks() bool {
	if o != nil && o.Clicks.IsSet() {
		return true
	}

	return false
}

// SetClicks gets a reference to the given NullableInt32 and assigns it to the Clicks field.
func (o *PlatformDeliveryDataEmailAllOf) SetClicks(v int32) {
	o.Clicks.Set(&v)
}
// SetClicksNil sets the value for Clicks to be an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) SetClicksNil() {
	o.Clicks.Set(nil)
}

// UnsetClicks ensures that no value is present for Clicks, not even an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) UnsetClicks() {
	o.Clicks.Unset()
}

// GetUniqueClicks returns the UniqueClicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataEmailAllOf) GetUniqueClicks() int32 {
	if o == nil || o.UniqueClicks.Get() == nil {
		var ret int32
		return ret
	}
	return *o.UniqueClicks.Get()
}

// GetUniqueClicksOk returns a tuple with the UniqueClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataEmailAllOf) GetUniqueClicksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniqueClicks.Get(), o.UniqueClicks.IsSet()
}

// HasUniqueClicks returns a boolean if a field has been set.
func (o *PlatformDeliveryDataEmailAllOf) HasUniqueClicks() bool {
	if o != nil && o.UniqueClicks.IsSet() {
		return true
	}

	return false
}

// SetUniqueClicks gets a reference to the given NullableInt32 and assigns it to the UniqueClicks field.
func (o *PlatformDeliveryDataEmailAllOf) SetUniqueClicks(v int32) {
	o.UniqueClicks.Set(&v)
}
// SetUniqueClicksNil sets the value for UniqueClicks to be an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) SetUniqueClicksNil() {
	o.UniqueClicks.Set(nil)
}

// UnsetUniqueClicks ensures that no value is present for UniqueClicks, not even an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) UnsetUniqueClicks() {
	o.UniqueClicks.Unset()
}

// GetBounced returns the Bounced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataEmailAllOf) GetBounced() int32 {
	if o == nil || o.Bounced.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Bounced.Get()
}

// GetBouncedOk returns a tuple with the Bounced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataEmailAllOf) GetBouncedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bounced.Get(), o.Bounced.IsSet()
}

// HasBounced returns a boolean if a field has been set.
func (o *PlatformDeliveryDataEmailAllOf) HasBounced() bool {
	if o != nil && o.Bounced.IsSet() {
		return true
	}

	return false
}

// SetBounced gets a reference to the given NullableInt32 and assigns it to the Bounced field.
func (o *PlatformDeliveryDataEmailAllOf) SetBounced(v int32) {
	o.Bounced.Set(&v)
}
// SetBouncedNil sets the value for Bounced to be an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) SetBouncedNil() {
	o.Bounced.Set(nil)
}

// UnsetBounced ensures that no value is present for Bounced, not even an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) UnsetBounced() {
	o.Bounced.Unset()
}

// GetReportedSpam returns the ReportedSpam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataEmailAllOf) GetReportedSpam() int32 {
	if o == nil || o.ReportedSpam.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ReportedSpam.Get()
}

// GetReportedSpamOk returns a tuple with the ReportedSpam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataEmailAllOf) GetReportedSpamOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportedSpam.Get(), o.ReportedSpam.IsSet()
}

// HasReportedSpam returns a boolean if a field has been set.
func (o *PlatformDeliveryDataEmailAllOf) HasReportedSpam() bool {
	if o != nil && o.ReportedSpam.IsSet() {
		return true
	}

	return false
}

// SetReportedSpam gets a reference to the given NullableInt32 and assigns it to the ReportedSpam field.
func (o *PlatformDeliveryDataEmailAllOf) SetReportedSpam(v int32) {
	o.ReportedSpam.Set(&v)
}
// SetReportedSpamNil sets the value for ReportedSpam to be an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) SetReportedSpamNil() {
	o.ReportedSpam.Set(nil)
}

// UnsetReportedSpam ensures that no value is present for ReportedSpam, not even an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) UnsetReportedSpam() {
	o.ReportedSpam.Unset()
}

// GetUnsubscribed returns the Unsubscribed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataEmailAllOf) GetUnsubscribed() int32 {
	if o == nil || o.Unsubscribed.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Unsubscribed.Get()
}

// GetUnsubscribedOk returns a tuple with the Unsubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataEmailAllOf) GetUnsubscribedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unsubscribed.Get(), o.Unsubscribed.IsSet()
}

// HasUnsubscribed returns a boolean if a field has been set.
func (o *PlatformDeliveryDataEmailAllOf) HasUnsubscribed() bool {
	if o != nil && o.Unsubscribed.IsSet() {
		return true
	}

	return false
}

// SetUnsubscribed gets a reference to the given NullableInt32 and assigns it to the Unsubscribed field.
func (o *PlatformDeliveryDataEmailAllOf) SetUnsubscribed(v int32) {
	o.Unsubscribed.Set(&v)
}
// SetUnsubscribedNil sets the value for Unsubscribed to be an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) SetUnsubscribedNil() {
	o.Unsubscribed.Set(nil)
}

// UnsetUnsubscribed ensures that no value is present for Unsubscribed, not even an explicit nil
func (o *PlatformDeliveryDataEmailAllOf) UnsetUnsubscribed() {
	o.Unsubscribed.Unset()
}

func (o PlatformDeliveryDataEmailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Opened.IsSet() {
		toSerialize["opened"] = o.Opened.Get()
	}
	if o.UniqueOpens.IsSet() {
		toSerialize["unique_opens"] = o.UniqueOpens.Get()
	}
	if o.Clicks.IsSet() {
		toSerialize["clicks"] = o.Clicks.Get()
	}
	if o.UniqueClicks.IsSet() {
		toSerialize["unique_clicks"] = o.UniqueClicks.Get()
	}
	if o.Bounced.IsSet() {
		toSerialize["bounced"] = o.Bounced.Get()
	}
	if o.ReportedSpam.IsSet() {
		toSerialize["reported_spam"] = o.ReportedSpam.Get()
	}
	if o.Unsubscribed.IsSet() {
		toSerialize["unsubscribed"] = o.Unsubscribed.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PlatformDeliveryDataEmailAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformDeliveryDataEmailAllOf := _PlatformDeliveryDataEmailAllOf{}

	if err = json.Unmarshal(bytes, &varPlatformDeliveryDataEmailAllOf); err == nil {
		*o = PlatformDeliveryDataEmailAllOf(varPlatformDeliveryDataEmailAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "opened")
		delete(additionalProperties, "unique_opens")
		delete(additionalProperties, "clicks")
		delete(additionalProperties, "unique_clicks")
		delete(additionalProperties, "bounced")
		delete(additionalProperties, "reported_spam")
		delete(additionalProperties, "unsubscribed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlatformDeliveryDataEmailAllOf struct {
	value *PlatformDeliveryDataEmailAllOf
	isSet bool
}

func (v NullablePlatformDeliveryDataEmailAllOf) Get() *PlatformDeliveryDataEmailAllOf {
	return v.value
}

func (v *NullablePlatformDeliveryDataEmailAllOf) Set(val *PlatformDeliveryDataEmailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformDeliveryDataEmailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformDeliveryDataEmailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformDeliveryDataEmailAllOf(val *PlatformDeliveryDataEmailAllOf) *NullablePlatformDeliveryDataEmailAllOf {
	return &NullablePlatformDeliveryDataEmailAllOf{value: val, isSet: true}
}

func (v NullablePlatformDeliveryDataEmailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformDeliveryDataEmailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


