/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.0.2
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// FilterNotificationTarget struct for FilterNotificationTarget
type FilterNotificationTarget struct {
	// relation = \">\" or \"<\" hours_ago = number of hours before or after the users last session. Example: \"1.1\" 
	LastSession *string `json:"last_session,omitempty"`
	// relation = \">\" or \"<\" hours_ago = number of hours before or after the users first session. Example: \"1.1\" 
	FirstSession *string `json:"first_session,omitempty"`
	// relation = \">\", \"<\", \"=\" or \"!=\" value = number sessions. Example: \"1\" 
	SessionCount *string `json:"session_count,omitempty"`
	// relation = \">\", \"<\", \"=\" or \"!=\" value = Time in seconds the user has been in your app. Example: \"3600\" 
	SessionTime *string `json:"session_time,omitempty"`
	// relation = \">\", \"<\", or \"=\" value = Amount in USD a user has spent on IAP (In App Purchases). Example: \"0.99\" 
	AmountSpent *string `json:"amount_spent,omitempty"`
	// relation = \">\", \"<\" or \"=\" key = SKU purchased in your app as an IAP (In App Purchases). Example: \"com.domain.100coinpack\" value = value of SKU to compare to. Example: \"0.99\" 
	BoughtSku *string `json:"bought_sku,omitempty"`
	// relation = \">\", \"<\", \"=\", \"!=\", \"exists\", \"not_exists\", \"time_elapsed_gt\" (paid plan only) or \"time_elapsed_lt\" (paid plan only) See Time Operators key = Tag key to compare. value = Tag value to compare. Not required for \"exists\" or \"not_exists\". Example: See Formatting Filters 
	Tag *string `json:"tag,omitempty"`
	// relation = \"=\" or \"!=\" value = 2 character language code. Example: \"en\". For a list of all language codes see Language & Localization. 
	Language *string `json:"language,omitempty"`
	// relation = \">\", \"<\", \"=\" or \"!=\" value = app version. Example: \"1.0.0\" 
	AppVersion *string `json:"app_version,omitempty"`
	// radius = in meters lat = latitude long = longitude 
	Location *string `json:"location,omitempty"`
	// value = email address Only for sending Push Notifications Use this for targeting push subscribers associated with an email set with all SDK setEmail methods To send emails to specific email addresses use include_email_tokens parameter 
	Email *string `json:"email,omitempty"`
	// relation = \"=\" value = 2-digit Country code Example: \"field\": \"country\", \"relation\": \"=\", \"value\", \"US\" 
	Country *string `json:"country,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FilterNotificationTarget FilterNotificationTarget

// NewFilterNotificationTarget instantiates a new FilterNotificationTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterNotificationTarget() *FilterNotificationTarget {
	this := FilterNotificationTarget{}
	return &this
}

// NewFilterNotificationTargetWithDefaults instantiates a new FilterNotificationTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterNotificationTargetWithDefaults() *FilterNotificationTarget {
	this := FilterNotificationTarget{}
	return &this
}

// GetLastSession returns the LastSession field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetLastSession() string {
	if o == nil || o.LastSession == nil {
		var ret string
		return ret
	}
	return *o.LastSession
}

// GetLastSessionOk returns a tuple with the LastSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetLastSessionOk() (*string, bool) {
	if o == nil || o.LastSession == nil {
		return nil, false
	}
	return o.LastSession, true
}

// HasLastSession returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasLastSession() bool {
	if o != nil && o.LastSession != nil {
		return true
	}

	return false
}

// SetLastSession gets a reference to the given string and assigns it to the LastSession field.
func (o *FilterNotificationTarget) SetLastSession(v string) {
	o.LastSession = &v
}

// GetFirstSession returns the FirstSession field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetFirstSession() string {
	if o == nil || o.FirstSession == nil {
		var ret string
		return ret
	}
	return *o.FirstSession
}

// GetFirstSessionOk returns a tuple with the FirstSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetFirstSessionOk() (*string, bool) {
	if o == nil || o.FirstSession == nil {
		return nil, false
	}
	return o.FirstSession, true
}

// HasFirstSession returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasFirstSession() bool {
	if o != nil && o.FirstSession != nil {
		return true
	}

	return false
}

// SetFirstSession gets a reference to the given string and assigns it to the FirstSession field.
func (o *FilterNotificationTarget) SetFirstSession(v string) {
	o.FirstSession = &v
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetSessionCount() string {
	if o == nil || o.SessionCount == nil {
		var ret string
		return ret
	}
	return *o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetSessionCountOk() (*string, bool) {
	if o == nil || o.SessionCount == nil {
		return nil, false
	}
	return o.SessionCount, true
}

// HasSessionCount returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasSessionCount() bool {
	if o != nil && o.SessionCount != nil {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given string and assigns it to the SessionCount field.
func (o *FilterNotificationTarget) SetSessionCount(v string) {
	o.SessionCount = &v
}

// GetSessionTime returns the SessionTime field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetSessionTime() string {
	if o == nil || o.SessionTime == nil {
		var ret string
		return ret
	}
	return *o.SessionTime
}

// GetSessionTimeOk returns a tuple with the SessionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetSessionTimeOk() (*string, bool) {
	if o == nil || o.SessionTime == nil {
		return nil, false
	}
	return o.SessionTime, true
}

// HasSessionTime returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasSessionTime() bool {
	if o != nil && o.SessionTime != nil {
		return true
	}

	return false
}

// SetSessionTime gets a reference to the given string and assigns it to the SessionTime field.
func (o *FilterNotificationTarget) SetSessionTime(v string) {
	o.SessionTime = &v
}

// GetAmountSpent returns the AmountSpent field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetAmountSpent() string {
	if o == nil || o.AmountSpent == nil {
		var ret string
		return ret
	}
	return *o.AmountSpent
}

// GetAmountSpentOk returns a tuple with the AmountSpent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetAmountSpentOk() (*string, bool) {
	if o == nil || o.AmountSpent == nil {
		return nil, false
	}
	return o.AmountSpent, true
}

// HasAmountSpent returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasAmountSpent() bool {
	if o != nil && o.AmountSpent != nil {
		return true
	}

	return false
}

// SetAmountSpent gets a reference to the given string and assigns it to the AmountSpent field.
func (o *FilterNotificationTarget) SetAmountSpent(v string) {
	o.AmountSpent = &v
}

// GetBoughtSku returns the BoughtSku field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetBoughtSku() string {
	if o == nil || o.BoughtSku == nil {
		var ret string
		return ret
	}
	return *o.BoughtSku
}

// GetBoughtSkuOk returns a tuple with the BoughtSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetBoughtSkuOk() (*string, bool) {
	if o == nil || o.BoughtSku == nil {
		return nil, false
	}
	return o.BoughtSku, true
}

// HasBoughtSku returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasBoughtSku() bool {
	if o != nil && o.BoughtSku != nil {
		return true
	}

	return false
}

// SetBoughtSku gets a reference to the given string and assigns it to the BoughtSku field.
func (o *FilterNotificationTarget) SetBoughtSku(v string) {
	o.BoughtSku = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *FilterNotificationTarget) SetTag(v string) {
	o.Tag = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *FilterNotificationTarget) SetLanguage(v string) {
	o.Language = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetAppVersion() string {
	if o == nil || o.AppVersion == nil {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetAppVersionOk() (*string, bool) {
	if o == nil || o.AppVersion == nil {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasAppVersion() bool {
	if o != nil && o.AppVersion != nil {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *FilterNotificationTarget) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *FilterNotificationTarget) SetLocation(v string) {
	o.Location = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *FilterNotificationTarget) SetEmail(v string) {
	o.Email = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *FilterNotificationTarget) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterNotificationTarget) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *FilterNotificationTarget) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *FilterNotificationTarget) SetCountry(v string) {
	o.Country = &v
}

func (o FilterNotificationTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastSession != nil {
		toSerialize["last_session"] = o.LastSession
	}
	if o.FirstSession != nil {
		toSerialize["first_session"] = o.FirstSession
	}
	if o.SessionCount != nil {
		toSerialize["session_count"] = o.SessionCount
	}
	if o.SessionTime != nil {
		toSerialize["session_time"] = o.SessionTime
	}
	if o.AmountSpent != nil {
		toSerialize["amount_spent"] = o.AmountSpent
	}
	if o.BoughtSku != nil {
		toSerialize["bought_sku"] = o.BoughtSku
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.AppVersion != nil {
		toSerialize["app_version"] = o.AppVersion
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FilterNotificationTarget) UnmarshalJSON(bytes []byte) (err error) {
	varFilterNotificationTarget := _FilterNotificationTarget{}

	if err = json.Unmarshal(bytes, &varFilterNotificationTarget); err == nil {
		*o = FilterNotificationTarget(varFilterNotificationTarget)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "last_session")
		delete(additionalProperties, "first_session")
		delete(additionalProperties, "session_count")
		delete(additionalProperties, "session_time")
		delete(additionalProperties, "amount_spent")
		delete(additionalProperties, "bought_sku")
		delete(additionalProperties, "tag")
		delete(additionalProperties, "language")
		delete(additionalProperties, "app_version")
		delete(additionalProperties, "location")
		delete(additionalProperties, "email")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilterNotificationTarget struct {
	value *FilterNotificationTarget
	isSet bool
}

func (v NullableFilterNotificationTarget) Get() *FilterNotificationTarget {
	return v.value
}

func (v *NullableFilterNotificationTarget) Set(val *FilterNotificationTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterNotificationTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterNotificationTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterNotificationTarget(val *FilterNotificationTarget) *NullableFilterNotificationTarget {
	return &NullableFilterNotificationTarget{value: val, isSet: true}
}

func (v NullableFilterNotificationTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterNotificationTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

