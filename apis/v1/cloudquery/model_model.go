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

// Model A description of the schema of the data to follow in the DataResultChunks.
type Model struct {
	// This name will be used by the following DataResultChunks as a reference to the description of their structure.
	Name *string `json:"name,omitempty"`
	// An array of type descriptors for each fetched expression in the query. The order of the array matches the order of values in the DataResultChunks with this schema model.'
	Fields []ModelFieldsInner `json:"fields,omitempty"`
}

// NewModel instantiates a new Model object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel() *Model {
	this := Model{}
	return &this
}

// NewModelWithDefaults instantiates a new Model object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWithDefaults() *Model {
	this := Model{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Model) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Model) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Model) SetName(v string) {
	o.Name = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *Model) GetFields() []ModelFieldsInner {
	if o == nil || isNil(o.Fields) {
		var ret []ModelFieldsInner
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetFieldsOk() ([]ModelFieldsInner, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *Model) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []ModelFieldsInner and assigns it to the Fields field.
func (o *Model) SetFields(v []ModelFieldsInner) {
	o.Fields = v
}

func (o Model) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableModel struct {
	value *Model
	isSet bool
}

func (v NullableModel) Get() *Model {
	return v.value
}

func (v *NullableModel) Set(val *Model) {
	v.value = val
	v.isSet = true
}

func (v NullableModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel(val *Model) *NullableModel {
	return &NullableModel{value: val, isSet: true}
}

func (v NullableModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

