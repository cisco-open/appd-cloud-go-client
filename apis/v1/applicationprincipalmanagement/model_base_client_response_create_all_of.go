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
Application Principal Management Service

Handles all administrative requests to manage application identities, including both Agents and Service principals.

API version: 1.0
Contact: info@appdynamics.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package applicationprincipalmanagement

import (
	"encoding/json"
)

// BaseClientResponseCreateAllOf struct for BaseClientResponseCreateAllOf
type BaseClientResponseCreateAllOf struct {
	// The client's secret, used to authenticate during an oAuth token request.
	ClientSecret *string `json:"clientSecret,omitempty"`
}

// NewBaseClientResponseCreateAllOf instantiates a new BaseClientResponseCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseClientResponseCreateAllOf() *BaseClientResponseCreateAllOf {
	this := BaseClientResponseCreateAllOf{}
	return &this
}

// NewBaseClientResponseCreateAllOfWithDefaults instantiates a new BaseClientResponseCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseClientResponseCreateAllOfWithDefaults() *BaseClientResponseCreateAllOf {
	this := BaseClientResponseCreateAllOf{}
	return &this
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *BaseClientResponseCreateAllOf) GetClientSecret() string {
	if o == nil || isNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseClientResponseCreateAllOf) GetClientSecretOk() (*string, bool) {
	if o == nil || isNil(o.ClientSecret) {
    return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *BaseClientResponseCreateAllOf) HasClientSecret() bool {
	if o != nil && !isNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *BaseClientResponseCreateAllOf) SetClientSecret(v string) {
	o.ClientSecret = &v
}

func (o BaseClientResponseCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableBaseClientResponseCreateAllOf struct {
	value *BaseClientResponseCreateAllOf
	isSet bool
}

func (v NullableBaseClientResponseCreateAllOf) Get() *BaseClientResponseCreateAllOf {
	return v.value
}

func (v *NullableBaseClientResponseCreateAllOf) Set(val *BaseClientResponseCreateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseClientResponseCreateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseClientResponseCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseClientResponseCreateAllOf(val *BaseClientResponseCreateAllOf) *NullableBaseClientResponseCreateAllOf {
	return &NullableBaseClientResponseCreateAllOf{value: val, isSet: true}
}

func (v NullableBaseClientResponseCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseClientResponseCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


