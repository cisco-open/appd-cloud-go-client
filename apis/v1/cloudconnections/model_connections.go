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

// Connections struct for Connections
type Connections struct {
	Items []ConnectionResponse `json:"items"`
}

// NewConnections instantiates a new Connections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnections(items []ConnectionResponse) *Connections {
	this := Connections{}
	this.Items = items
	return &this
}

// NewConnectionsWithDefaults instantiates a new Connections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionsWithDefaults() *Connections {
	this := Connections{}
	return &this
}

// GetItems returns the Items field value
func (o *Connections) GetItems() []ConnectionResponse {
	if o == nil {
		var ret []ConnectionResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Connections) GetItemsOk() ([]ConnectionResponse, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Connections) SetItems(v []ConnectionResponse) {
	o.Items = v
}

func (o Connections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableConnections struct {
	value *Connections
	isSet bool
}

func (v NullableConnections) Get() *Connections {
	return v.value
}

func (v *NullableConnections) Set(val *Connections) {
	v.value = val
	v.isSet = true
}

func (v NullableConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnections(val *Connections) *NullableConnections {
	return &NullableConnections{value: val, isSet: true}
}

func (v NullableConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


