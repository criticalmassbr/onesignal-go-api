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

// SubscriptionObject struct for SubscriptionObject
type SubscriptionObject struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Token *string `json:"token,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	NotificationTypes *float32 `json:"notification_types,omitempty"`
	SessionTime *float32 `json:"session_time,omitempty"`
	SessionCount *float32 `json:"session_count,omitempty"`
	Sdk *string `json:"sdk,omitempty"`
	DeviceModel *string `json:"device_model,omitempty"`
	DeviceOs *string `json:"device_os,omitempty"`
	Rooted *bool `json:"rooted,omitempty"`
	TestType *float32 `json:"test_type,omitempty"`
	AppVersion *string `json:"app_version,omitempty"`
	NetType *float32 `json:"net_type,omitempty"`
	Carrier *string `json:"carrier,omitempty"`
	WebAuth *string `json:"web_auth,omitempty"`
	WebP256 *string `json:"web_p256,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionObject SubscriptionObject

// NewSubscriptionObject instantiates a new SubscriptionObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionObject() *SubscriptionObject {
	this := SubscriptionObject{}
	return &this
}

// NewSubscriptionObjectWithDefaults instantiates a new SubscriptionObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionObjectWithDefaults() *SubscriptionObject {
	this := SubscriptionObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionObject) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionObject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionObject) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubscriptionObject) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubscriptionObject) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SubscriptionObject) SetType(v string) {
	o.Type = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SubscriptionObject) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SubscriptionObject) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SubscriptionObject) SetToken(v string) {
	o.Token = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SubscriptionObject) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SubscriptionObject) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SubscriptionObject) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNotificationTypes returns the NotificationTypes field value if set, zero value otherwise.
func (o *SubscriptionObject) GetNotificationTypes() float32 {
	if o == nil || o.NotificationTypes == nil {
		var ret float32
		return ret
	}
	return *o.NotificationTypes
}

// GetNotificationTypesOk returns a tuple with the NotificationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetNotificationTypesOk() (*float32, bool) {
	if o == nil || o.NotificationTypes == nil {
		return nil, false
	}
	return o.NotificationTypes, true
}

// HasNotificationTypes returns a boolean if a field has been set.
func (o *SubscriptionObject) HasNotificationTypes() bool {
	if o != nil && o.NotificationTypes != nil {
		return true
	}

	return false
}

// SetNotificationTypes gets a reference to the given float32 and assigns it to the NotificationTypes field.
func (o *SubscriptionObject) SetNotificationTypes(v float32) {
	o.NotificationTypes = &v
}

// GetSessionTime returns the SessionTime field value if set, zero value otherwise.
func (o *SubscriptionObject) GetSessionTime() float32 {
	if o == nil || o.SessionTime == nil {
		var ret float32
		return ret
	}
	return *o.SessionTime
}

// GetSessionTimeOk returns a tuple with the SessionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetSessionTimeOk() (*float32, bool) {
	if o == nil || o.SessionTime == nil {
		return nil, false
	}
	return o.SessionTime, true
}

// HasSessionTime returns a boolean if a field has been set.
func (o *SubscriptionObject) HasSessionTime() bool {
	if o != nil && o.SessionTime != nil {
		return true
	}

	return false
}

// SetSessionTime gets a reference to the given float32 and assigns it to the SessionTime field.
func (o *SubscriptionObject) SetSessionTime(v float32) {
	o.SessionTime = &v
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise.
func (o *SubscriptionObject) GetSessionCount() float32 {
	if o == nil || o.SessionCount == nil {
		var ret float32
		return ret
	}
	return *o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetSessionCountOk() (*float32, bool) {
	if o == nil || o.SessionCount == nil {
		return nil, false
	}
	return o.SessionCount, true
}

// HasSessionCount returns a boolean if a field has been set.
func (o *SubscriptionObject) HasSessionCount() bool {
	if o != nil && o.SessionCount != nil {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given float32 and assigns it to the SessionCount field.
func (o *SubscriptionObject) SetSessionCount(v float32) {
	o.SessionCount = &v
}

// GetSdk returns the Sdk field value if set, zero value otherwise.
func (o *SubscriptionObject) GetSdk() string {
	if o == nil || o.Sdk == nil {
		var ret string
		return ret
	}
	return *o.Sdk
}

// GetSdkOk returns a tuple with the Sdk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetSdkOk() (*string, bool) {
	if o == nil || o.Sdk == nil {
		return nil, false
	}
	return o.Sdk, true
}

// HasSdk returns a boolean if a field has been set.
func (o *SubscriptionObject) HasSdk() bool {
	if o != nil && o.Sdk != nil {
		return true
	}

	return false
}

// SetSdk gets a reference to the given string and assigns it to the Sdk field.
func (o *SubscriptionObject) SetSdk(v string) {
	o.Sdk = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *SubscriptionObject) GetDeviceModel() string {
	if o == nil || o.DeviceModel == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetDeviceModelOk() (*string, bool) {
	if o == nil || o.DeviceModel == nil {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *SubscriptionObject) HasDeviceModel() bool {
	if o != nil && o.DeviceModel != nil {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *SubscriptionObject) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetDeviceOs returns the DeviceOs field value if set, zero value otherwise.
func (o *SubscriptionObject) GetDeviceOs() string {
	if o == nil || o.DeviceOs == nil {
		var ret string
		return ret
	}
	return *o.DeviceOs
}

// GetDeviceOsOk returns a tuple with the DeviceOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetDeviceOsOk() (*string, bool) {
	if o == nil || o.DeviceOs == nil {
		return nil, false
	}
	return o.DeviceOs, true
}

// HasDeviceOs returns a boolean if a field has been set.
func (o *SubscriptionObject) HasDeviceOs() bool {
	if o != nil && o.DeviceOs != nil {
		return true
	}

	return false
}

// SetDeviceOs gets a reference to the given string and assigns it to the DeviceOs field.
func (o *SubscriptionObject) SetDeviceOs(v string) {
	o.DeviceOs = &v
}

// GetRooted returns the Rooted field value if set, zero value otherwise.
func (o *SubscriptionObject) GetRooted() bool {
	if o == nil || o.Rooted == nil {
		var ret bool
		return ret
	}
	return *o.Rooted
}

// GetRootedOk returns a tuple with the Rooted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetRootedOk() (*bool, bool) {
	if o == nil || o.Rooted == nil {
		return nil, false
	}
	return o.Rooted, true
}

// HasRooted returns a boolean if a field has been set.
func (o *SubscriptionObject) HasRooted() bool {
	if o != nil && o.Rooted != nil {
		return true
	}

	return false
}

// SetRooted gets a reference to the given bool and assigns it to the Rooted field.
func (o *SubscriptionObject) SetRooted(v bool) {
	o.Rooted = &v
}

// GetTestType returns the TestType field value if set, zero value otherwise.
func (o *SubscriptionObject) GetTestType() float32 {
	if o == nil || o.TestType == nil {
		var ret float32
		return ret
	}
	return *o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetTestTypeOk() (*float32, bool) {
	if o == nil || o.TestType == nil {
		return nil, false
	}
	return o.TestType, true
}

// HasTestType returns a boolean if a field has been set.
func (o *SubscriptionObject) HasTestType() bool {
	if o != nil && o.TestType != nil {
		return true
	}

	return false
}

// SetTestType gets a reference to the given float32 and assigns it to the TestType field.
func (o *SubscriptionObject) SetTestType(v float32) {
	o.TestType = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *SubscriptionObject) GetAppVersion() string {
	if o == nil || o.AppVersion == nil {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetAppVersionOk() (*string, bool) {
	if o == nil || o.AppVersion == nil {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *SubscriptionObject) HasAppVersion() bool {
	if o != nil && o.AppVersion != nil {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *SubscriptionObject) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetNetType returns the NetType field value if set, zero value otherwise.
func (o *SubscriptionObject) GetNetType() float32 {
	if o == nil || o.NetType == nil {
		var ret float32
		return ret
	}
	return *o.NetType
}

// GetNetTypeOk returns a tuple with the NetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetNetTypeOk() (*float32, bool) {
	if o == nil || o.NetType == nil {
		return nil, false
	}
	return o.NetType, true
}

// HasNetType returns a boolean if a field has been set.
func (o *SubscriptionObject) HasNetType() bool {
	if o != nil && o.NetType != nil {
		return true
	}

	return false
}

// SetNetType gets a reference to the given float32 and assigns it to the NetType field.
func (o *SubscriptionObject) SetNetType(v float32) {
	o.NetType = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *SubscriptionObject) GetCarrier() string {
	if o == nil || o.Carrier == nil {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetCarrierOk() (*string, bool) {
	if o == nil || o.Carrier == nil {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *SubscriptionObject) HasCarrier() bool {
	if o != nil && o.Carrier != nil {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *SubscriptionObject) SetCarrier(v string) {
	o.Carrier = &v
}

// GetWebAuth returns the WebAuth field value if set, zero value otherwise.
func (o *SubscriptionObject) GetWebAuth() string {
	if o == nil || o.WebAuth == nil {
		var ret string
		return ret
	}
	return *o.WebAuth
}

// GetWebAuthOk returns a tuple with the WebAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetWebAuthOk() (*string, bool) {
	if o == nil || o.WebAuth == nil {
		return nil, false
	}
	return o.WebAuth, true
}

// HasWebAuth returns a boolean if a field has been set.
func (o *SubscriptionObject) HasWebAuth() bool {
	if o != nil && o.WebAuth != nil {
		return true
	}

	return false
}

// SetWebAuth gets a reference to the given string and assigns it to the WebAuth field.
func (o *SubscriptionObject) SetWebAuth(v string) {
	o.WebAuth = &v
}

// GetWebP256 returns the WebP256 field value if set, zero value otherwise.
func (o *SubscriptionObject) GetWebP256() string {
	if o == nil || o.WebP256 == nil {
		var ret string
		return ret
	}
	return *o.WebP256
}

// GetWebP256Ok returns a tuple with the WebP256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionObject) GetWebP256Ok() (*string, bool) {
	if o == nil || o.WebP256 == nil {
		return nil, false
	}
	return o.WebP256, true
}

// HasWebP256 returns a boolean if a field has been set.
func (o *SubscriptionObject) HasWebP256() bool {
	if o != nil && o.WebP256 != nil {
		return true
	}

	return false
}

// SetWebP256 gets a reference to the given string and assigns it to the WebP256 field.
func (o *SubscriptionObject) SetWebP256(v string) {
	o.WebP256 = &v
}

func (o SubscriptionObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.NotificationTypes != nil {
		toSerialize["notification_types"] = o.NotificationTypes
	}
	if o.SessionTime != nil {
		toSerialize["session_time"] = o.SessionTime
	}
	if o.SessionCount != nil {
		toSerialize["session_count"] = o.SessionCount
	}
	if o.Sdk != nil {
		toSerialize["sdk"] = o.Sdk
	}
	if o.DeviceModel != nil {
		toSerialize["device_model"] = o.DeviceModel
	}
	if o.DeviceOs != nil {
		toSerialize["device_os"] = o.DeviceOs
	}
	if o.Rooted != nil {
		toSerialize["rooted"] = o.Rooted
	}
	if o.TestType != nil {
		toSerialize["test_type"] = o.TestType
	}
	if o.AppVersion != nil {
		toSerialize["app_version"] = o.AppVersion
	}
	if o.NetType != nil {
		toSerialize["net_type"] = o.NetType
	}
	if o.Carrier != nil {
		toSerialize["carrier"] = o.Carrier
	}
	if o.WebAuth != nil {
		toSerialize["web_auth"] = o.WebAuth
	}
	if o.WebP256 != nil {
		toSerialize["web_p256"] = o.WebP256
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SubscriptionObject) UnmarshalJSON(bytes []byte) (err error) {
	varSubscriptionObject := _SubscriptionObject{}

	if err = json.Unmarshal(bytes, &varSubscriptionObject); err == nil {
		*o = SubscriptionObject(varSubscriptionObject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "token")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "notification_types")
		delete(additionalProperties, "session_time")
		delete(additionalProperties, "session_count")
		delete(additionalProperties, "sdk")
		delete(additionalProperties, "device_model")
		delete(additionalProperties, "device_os")
		delete(additionalProperties, "rooted")
		delete(additionalProperties, "test_type")
		delete(additionalProperties, "app_version")
		delete(additionalProperties, "net_type")
		delete(additionalProperties, "carrier")
		delete(additionalProperties, "web_auth")
		delete(additionalProperties, "web_p256")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionObject struct {
	value *SubscriptionObject
	isSet bool
}

func (v NullableSubscriptionObject) Get() *SubscriptionObject {
	return v.value
}

func (v *NullableSubscriptionObject) Set(val *SubscriptionObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionObject(val *SubscriptionObject) *NullableSubscriptionObject {
	return &NullableSubscriptionObject{value: val, isSet: true}
}

func (v NullableSubscriptionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

