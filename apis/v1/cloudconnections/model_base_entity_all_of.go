// Copyright 2023 Cisco Systems, Inc.
// 
// Licensed under the MPL License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     https://www.mozilla.org/en-US/MPL/2.0/
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/*
AppDynamics Cloud Connections API

Enables you to manage cloud connections

API version: 1.0.0
Contact: support@appdynamics.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudconnections

import (
	"encoding/json"
)

// BaseEntityAllOf struct for BaseEntityAllOf
type BaseEntityAllOf struct {
	Type ProviderType `json:"type"`
}

// NewBaseEntityAllOf instantiates a new BaseEntityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEntityAllOf(type_ ProviderType) *BaseEntityAllOf {
	this := BaseEntityAllOf{}
	this.Type = type_
	return &this
}

// NewBaseEntityAllOfWithDefaults instantiates a new BaseEntityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEntityAllOfWithDefaults() *BaseEntityAllOf {
	this := BaseEntityAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *BaseEntityAllOf) GetType() ProviderType {
	if o == nil {
		var ret ProviderType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BaseEntityAllOf) GetTypeOk() (*ProviderType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BaseEntityAllOf) SetType(v ProviderType) {
	o.Type = v
}

func (o BaseEntityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEntityAllOf struct {
	value *BaseEntityAllOf
	isSet bool
}

func (v NullableBaseEntityAllOf) Get() *BaseEntityAllOf {
	return v.value
}

func (v *NullableBaseEntityAllOf) Set(val *BaseEntityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEntityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEntityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEntityAllOf(val *BaseEntityAllOf) *NullableBaseEntityAllOf {
	return &NullableBaseEntityAllOf{value: val, isSet: true}
}

func (v NullableBaseEntityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEntityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


