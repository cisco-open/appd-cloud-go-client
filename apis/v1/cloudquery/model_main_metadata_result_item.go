/*
AppDynamics Cloud Query Service API

An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudquery

import (
	"encoding/json"
	"time"
)

// MainMetadataResultItem Metadata for the main data chunk.
type MainMetadataResultItem struct {
	// Specifies the start of the time range which was used for querying.
	Since *time.Time `json:"since,omitempty"`
	// Specifies the end of the time range which was used for querying.
	Until *time.Time `json:"until,omitempty"`
}

// NewMainMetadataResultItem instantiates a new MainMetadataResultItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMainMetadataResultItem() *MainMetadataResultItem {
	this := MainMetadataResultItem{}
	return &this
}

// NewMainMetadataResultItemWithDefaults instantiates a new MainMetadataResultItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMainMetadataResultItemWithDefaults() *MainMetadataResultItem {
	this := MainMetadataResultItem{}
	return &this
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *MainMetadataResultItem) GetSince() time.Time {
	if o == nil || isNil(o.Since) {
		var ret time.Time
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainMetadataResultItem) GetSinceOk() (*time.Time, bool) {
	if o == nil || isNil(o.Since) {
    return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *MainMetadataResultItem) HasSince() bool {
	if o != nil && !isNil(o.Since) {
		return true
	}

	return false
}

// SetSince gets a reference to the given time.Time and assigns it to the Since field.
func (o *MainMetadataResultItem) SetSince(v time.Time) {
	o.Since = &v
}

// GetUntil returns the Until field value if set, zero value otherwise.
func (o *MainMetadataResultItem) GetUntil() time.Time {
	if o == nil || isNil(o.Until) {
		var ret time.Time
		return ret
	}
	return *o.Until
}

// GetUntilOk returns a tuple with the Until field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainMetadataResultItem) GetUntilOk() (*time.Time, bool) {
	if o == nil || isNil(o.Until) {
    return nil, false
	}
	return o.Until, true
}

// HasUntil returns a boolean if a field has been set.
func (o *MainMetadataResultItem) HasUntil() bool {
	if o != nil && !isNil(o.Until) {
		return true
	}

	return false
}

// SetUntil gets a reference to the given time.Time and assigns it to the Until field.
func (o *MainMetadataResultItem) SetUntil(v time.Time) {
	o.Until = &v
}

func (o MainMetadataResultItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Since) {
		toSerialize["since"] = o.Since
	}
	if !isNil(o.Until) {
		toSerialize["until"] = o.Until
	}
	return json.Marshal(toSerialize)
}

type NullableMainMetadataResultItem struct {
	value *MainMetadataResultItem
	isSet bool
}

func (v NullableMainMetadataResultItem) Get() *MainMetadataResultItem {
	return v.value
}

func (v *NullableMainMetadataResultItem) Set(val *MainMetadataResultItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMainMetadataResultItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMainMetadataResultItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMainMetadataResultItem(val *MainMetadataResultItem) *NullableMainMetadataResultItem {
	return &NullableMainMetadataResultItem{value: val, isSet: true}
}

func (v NullableMainMetadataResultItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMainMetadataResultItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


