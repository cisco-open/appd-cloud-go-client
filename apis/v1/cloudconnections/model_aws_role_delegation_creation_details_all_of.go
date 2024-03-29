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

// AWSRoleDelegationCreationDetailsAllOf struct for AWSRoleDelegationCreationDetailsAllOf
type AWSRoleDelegationCreationDetailsAllOf struct {
	AccessType ConnectionAccessType `json:"accessType"`
}

// NewAWSRoleDelegationCreationDetailsAllOf instantiates a new AWSRoleDelegationCreationDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSRoleDelegationCreationDetailsAllOf(accessType ConnectionAccessType) *AWSRoleDelegationCreationDetailsAllOf {
	this := AWSRoleDelegationCreationDetailsAllOf{}
	this.AccessType = accessType
	return &this
}

// NewAWSRoleDelegationCreationDetailsAllOfWithDefaults instantiates a new AWSRoleDelegationCreationDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSRoleDelegationCreationDetailsAllOfWithDefaults() *AWSRoleDelegationCreationDetailsAllOf {
	this := AWSRoleDelegationCreationDetailsAllOf{}
	var accessType ConnectionAccessType = ROLE_DELEGATION
	this.AccessType = accessType
	return &this
}

// GetAccessType returns the AccessType field value
func (o *AWSRoleDelegationCreationDetailsAllOf) GetAccessType() ConnectionAccessType {
	if o == nil {
		var ret ConnectionAccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *AWSRoleDelegationCreationDetailsAllOf) GetAccessTypeOk() (*ConnectionAccessType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *AWSRoleDelegationCreationDetailsAllOf) SetAccessType(v ConnectionAccessType) {
	o.AccessType = v
}

func (o AWSRoleDelegationCreationDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	return json.Marshal(toSerialize)
}

type NullableAWSRoleDelegationCreationDetailsAllOf struct {
	value *AWSRoleDelegationCreationDetailsAllOf
	isSet bool
}

func (v NullableAWSRoleDelegationCreationDetailsAllOf) Get() *AWSRoleDelegationCreationDetailsAllOf {
	return v.value
}

func (v *NullableAWSRoleDelegationCreationDetailsAllOf) Set(val *AWSRoleDelegationCreationDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSRoleDelegationCreationDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSRoleDelegationCreationDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSRoleDelegationCreationDetailsAllOf(val *AWSRoleDelegationCreationDetailsAllOf) *NullableAWSRoleDelegationCreationDetailsAllOf {
	return &NullableAWSRoleDelegationCreationDetailsAllOf{value: val, isSet: true}
}

func (v NullableAWSRoleDelegationCreationDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSRoleDelegationCreationDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


