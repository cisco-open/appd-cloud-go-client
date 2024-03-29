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
AppDynamics Cloud Query Service API

An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudquery

import (
	"encoding/json"
)

// ErrorDetailItem A detailed description of the error.
type ErrorDetailItem struct {
	// A short description on error cause.
	Message string `json:"message"`
	// A hint to resolve the error.
	FixSuggestion *string `json:"fixSuggestion,omitempty"`
	// A list of fix possibilities to resolve the error.
	FixPossibilities []ErrorDetailItemFixPossibilitiesInner `json:"fixPossibilities"`
	// The type of the error.
	ErrorType string `json:"errorType"`
	// The start position of the error in format 'lineNum:charIdx'.
	ErrorFrom string `json:"errorFrom"`
	// The end position of the error in format 'lineNum:charIdx'.
	ErrorTo string `json:"errorTo"`
}

// NewErrorDetailItem instantiates a new ErrorDetailItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetailItem(message string, fixPossibilities []ErrorDetailItemFixPossibilitiesInner, errorType string, errorFrom string, errorTo string) *ErrorDetailItem {
	this := ErrorDetailItem{}
	this.Message = message
	this.FixPossibilities = fixPossibilities
	this.ErrorType = errorType
	this.ErrorFrom = errorFrom
	this.ErrorTo = errorTo
	return &this
}

// NewErrorDetailItemWithDefaults instantiates a new ErrorDetailItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailItemWithDefaults() *ErrorDetailItem {
	this := ErrorDetailItem{}
	return &this
}

// GetMessage returns the Message field value
func (o *ErrorDetailItem) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailItem) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorDetailItem) SetMessage(v string) {
	o.Message = v
}

// GetFixSuggestion returns the FixSuggestion field value if set, zero value otherwise.
func (o *ErrorDetailItem) GetFixSuggestion() string {
	if o == nil || isNil(o.FixSuggestion) {
		var ret string
		return ret
	}
	return *o.FixSuggestion
}

// GetFixSuggestionOk returns a tuple with the FixSuggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetailItem) GetFixSuggestionOk() (*string, bool) {
	if o == nil || isNil(o.FixSuggestion) {
    return nil, false
	}
	return o.FixSuggestion, true
}

// HasFixSuggestion returns a boolean if a field has been set.
func (o *ErrorDetailItem) HasFixSuggestion() bool {
	if o != nil && !isNil(o.FixSuggestion) {
		return true
	}

	return false
}

// SetFixSuggestion gets a reference to the given string and assigns it to the FixSuggestion field.
func (o *ErrorDetailItem) SetFixSuggestion(v string) {
	o.FixSuggestion = &v
}

// GetFixPossibilities returns the FixPossibilities field value
func (o *ErrorDetailItem) GetFixPossibilities() []ErrorDetailItemFixPossibilitiesInner {
	if o == nil {
		var ret []ErrorDetailItemFixPossibilitiesInner
		return ret
	}

	return o.FixPossibilities
}

// GetFixPossibilitiesOk returns a tuple with the FixPossibilities field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailItem) GetFixPossibilitiesOk() ([]ErrorDetailItemFixPossibilitiesInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.FixPossibilities, true
}

// SetFixPossibilities sets field value
func (o *ErrorDetailItem) SetFixPossibilities(v []ErrorDetailItemFixPossibilitiesInner) {
	o.FixPossibilities = v
}

// GetErrorType returns the ErrorType field value
func (o *ErrorDetailItem) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailItem) GetErrorTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *ErrorDetailItem) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorFrom returns the ErrorFrom field value
func (o *ErrorDetailItem) GetErrorFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorFrom
}

// GetErrorFromOk returns a tuple with the ErrorFrom field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailItem) GetErrorFromOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ErrorFrom, true
}

// SetErrorFrom sets field value
func (o *ErrorDetailItem) SetErrorFrom(v string) {
	o.ErrorFrom = v
}

// GetErrorTo returns the ErrorTo field value
func (o *ErrorDetailItem) GetErrorTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorTo
}

// GetErrorToOk returns a tuple with the ErrorTo field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailItem) GetErrorToOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ErrorTo, true
}

// SetErrorTo sets field value
func (o *ErrorDetailItem) SetErrorTo(v string) {
	o.ErrorTo = v
}

func (o ErrorDetailItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.FixSuggestion) {
		toSerialize["fixSuggestion"] = o.FixSuggestion
	}
	if true {
		toSerialize["fixPossibilities"] = o.FixPossibilities
	}
	if true {
		toSerialize["errorType"] = o.ErrorType
	}
	if true {
		toSerialize["errorFrom"] = o.ErrorFrom
	}
	if true {
		toSerialize["errorTo"] = o.ErrorTo
	}
	return json.Marshal(toSerialize)
}

type NullableErrorDetailItem struct {
	value *ErrorDetailItem
	isSet bool
}

func (v NullableErrorDetailItem) Get() *ErrorDetailItem {
	return v.value
}

func (v *NullableErrorDetailItem) Set(val *ErrorDetailItem) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetailItem) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetailItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetailItem(val *ErrorDetailItem) *NullableErrorDetailItem {
	return &NullableErrorDetailItem{value: val, isSet: true}
}

func (v NullableErrorDetailItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetailItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


