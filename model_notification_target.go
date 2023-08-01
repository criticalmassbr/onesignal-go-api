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
	"fmt"
)

// NotificationTarget struct for NotificationTarget
type NotificationTarget struct {
	PlayerNotificationTarget *PlayerNotificationTarget
	SegmentNotificationTarget *SegmentNotificationTarget
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationTarget) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PlayerNotificationTarget
	err = json.Unmarshal(data, &dst.PlayerNotificationTarget);
	if err == nil {
		jsonPlayerNotificationTarget, _ := json.Marshal(dst.PlayerNotificationTarget)
		if string(jsonPlayerNotificationTarget) == "{}" { // empty struct
			dst.PlayerNotificationTarget = nil
		} else {
			return nil // data stored in dst.PlayerNotificationTarget, return on the first match
		}
	} else {
		dst.PlayerNotificationTarget = nil
	}

	// try to unmarshal JSON data into SegmentNotificationTarget
	err = json.Unmarshal(data, &dst.SegmentNotificationTarget);
	if err == nil {
		jsonSegmentNotificationTarget, _ := json.Marshal(dst.SegmentNotificationTarget)
		if string(jsonSegmentNotificationTarget) == "{}" { // empty struct
			dst.SegmentNotificationTarget = nil
		} else {
			return nil // data stored in dst.SegmentNotificationTarget, return on the first match
		}
	} else {
		dst.SegmentNotificationTarget = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(NotificationTarget)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationTarget) MarshalJSON() ([]byte, error) {
	if src.PlayerNotificationTarget != nil {
		return json.Marshal(&src.PlayerNotificationTarget)
	}

	if src.SegmentNotificationTarget != nil {
		return json.Marshal(&src.SegmentNotificationTarget)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationTarget struct {
	value *NotificationTarget
	isSet bool
}

func (v NullableNotificationTarget) Get() *NotificationTarget {
	return v.value
}

func (v *NullableNotificationTarget) Set(val *NotificationTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTarget(val *NotificationTarget) *NullableNotificationTarget {
	return &NullableNotificationTarget{value: val, isSet: true}
}

func (v NullableNotificationTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


