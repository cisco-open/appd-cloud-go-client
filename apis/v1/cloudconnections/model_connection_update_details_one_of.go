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

// ConnectionUpdateDetailsOneOf struct for ConnectionUpdateDetailsOneOf
type ConnectionUpdateDetailsOneOf struct {
	RoleName string `json:"roleName"`
}

// NewConnectionUpdateDetailsOneOf instantiates a new ConnectionUpdateDetailsOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionUpdateDetailsOneOf(roleName string) *ConnectionUpdateDetailsOneOf {
	this := ConnectionUpdateDetailsOneOf{}
	this.RoleName = roleName
	return &this
}

// NewConnectionUpdateDetailsOneOfWithDefaults instantiates a new ConnectionUpdateDetailsOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionUpdateDetailsOneOfWithDefaults() *ConnectionUpdateDetailsOneOf {
	this := ConnectionUpdateDetailsOneOf{}
	return &this
}

// GetRoleName returns the RoleName field value
func (o *ConnectionUpdateDetailsOneOf) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *ConnectionUpdateDetailsOneOf) GetRoleNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *ConnectionUpdateDetailsOneOf) SetRoleName(v string) {
	o.RoleName = v
}

func (o ConnectionUpdateDetailsOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionUpdateDetailsOneOf struct {
	value *ConnectionUpdateDetailsOneOf
	isSet bool
}

func (v NullableConnectionUpdateDetailsOneOf) Get() *ConnectionUpdateDetailsOneOf {
	return v.value
}

func (v *NullableConnectionUpdateDetailsOneOf) Set(val *ConnectionUpdateDetailsOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionUpdateDetailsOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionUpdateDetailsOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionUpdateDetailsOneOf(val *ConnectionUpdateDetailsOneOf) *NullableConnectionUpdateDetailsOneOf {
	return &NullableConnectionUpdateDetailsOneOf{value: val, isSet: true}
}

func (v NullableConnectionUpdateDetailsOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionUpdateDetailsOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

