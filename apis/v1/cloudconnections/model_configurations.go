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

// Configurations struct for Configurations
type Configurations struct {
	Items []ConfigurationDetail `json:"items"`
}

// NewConfigurations instantiates a new Configurations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurations(items []ConfigurationDetail) *Configurations {
	this := Configurations{}
	this.Items = items
	return &this
}

// NewConfigurationsWithDefaults instantiates a new Configurations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationsWithDefaults() *Configurations {
	this := Configurations{}
	return &this
}

// GetItems returns the Items field value
func (o *Configurations) GetItems() []ConfigurationDetail {
	if o == nil {
		var ret []ConfigurationDetail
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Configurations) GetItemsOk() ([]ConfigurationDetail, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Configurations) SetItems(v []ConfigurationDetail) {
	o.Items = v
}

func (o Configurations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurations struct {
	value *Configurations
	isSet bool
}

func (v NullableConfigurations) Get() *Configurations {
	return v.value
}

func (v *NullableConfigurations) Set(val *Configurations) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurations) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurations(val *Configurations) *NullableConfigurations {
	return &NullableConfigurations{value: val, isSet: true}
}

func (v NullableConfigurations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


