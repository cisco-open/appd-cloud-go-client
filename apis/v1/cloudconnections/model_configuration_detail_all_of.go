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
	"time"
)

// ConfigurationDetailAllOf struct for ConfigurationDetailAllOf
type ConfigurationDetailAllOf struct {
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewConfigurationDetailAllOf instantiates a new ConfigurationDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationDetailAllOf(id string, createdAt time.Time, updatedAt time.Time) *ConfigurationDetailAllOf {
	this := ConfigurationDetailAllOf{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewConfigurationDetailAllOfWithDefaults instantiates a new ConfigurationDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationDetailAllOfWithDefaults() *ConfigurationDetailAllOf {
	this := ConfigurationDetailAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigurationDetailAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDetailAllOf) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigurationDetailAllOf) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ConfigurationDetailAllOf) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDetailAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ConfigurationDetailAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ConfigurationDetailAllOf) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDetailAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ConfigurationDetailAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ConfigurationDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationDetailAllOf struct {
	value *ConfigurationDetailAllOf
	isSet bool
}

func (v NullableConfigurationDetailAllOf) Get() *ConfigurationDetailAllOf {
	return v.value
}

func (v *NullableConfigurationDetailAllOf) Set(val *ConfigurationDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationDetailAllOf(val *ConfigurationDetailAllOf) *NullableConfigurationDetailAllOf {
	return &NullableConfigurationDetailAllOf{value: val, isSet: true}
}

func (v NullableConfigurationDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


