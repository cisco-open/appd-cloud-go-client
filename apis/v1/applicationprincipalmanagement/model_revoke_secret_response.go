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

// RevokeSecretResponse struct for RevokeSecretResponse
type RevokeSecretResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

// NewRevokeSecretResponse instantiates a new RevokeSecretResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokeSecretResponse(status string, message string) *RevokeSecretResponse {
	this := RevokeSecretResponse{}
	this.Status = status
	this.Message = message
	return &this
}

// NewRevokeSecretResponseWithDefaults instantiates a new RevokeSecretResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeSecretResponseWithDefaults() *RevokeSecretResponse {
	this := RevokeSecretResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *RevokeSecretResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RevokeSecretResponse) GetStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RevokeSecretResponse) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *RevokeSecretResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RevokeSecretResponse) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RevokeSecretResponse) SetMessage(v string) {
	o.Message = v
}

func (o RevokeSecretResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableRevokeSecretResponse struct {
	value *RevokeSecretResponse
	isSet bool
}

func (v NullableRevokeSecretResponse) Get() *RevokeSecretResponse {
	return v.value
}

func (v *NullableRevokeSecretResponse) Set(val *RevokeSecretResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokeSecretResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokeSecretResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokeSecretResponse(val *RevokeSecretResponse) *NullableRevokeSecretResponse {
	return &NullableRevokeSecretResponse{value: val, isSet: true}
}

func (v NullableRevokeSecretResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokeSecretResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


