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
	"fmt"
)

// Configuration - struct for Configuration
type Configuration struct {
	AWSConfiguration *AWSConfiguration
	AzureConfiguration *AzureConfiguration
}

// AWSConfigurationAsConfiguration is a convenience function that returns AWSConfiguration wrapped in Configuration
func AWSConfigurationAsConfiguration(v *AWSConfiguration) Configuration {
	return Configuration{
		AWSConfiguration: v,
	}
}

// AzureConfigurationAsConfiguration is a convenience function that returns AzureConfiguration wrapped in Configuration
func AzureConfigurationAsConfiguration(v *AzureConfiguration) Configuration {
	return Configuration{
		AzureConfiguration: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Configuration) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSConfiguration
	err = newStrictDecoder(data).Decode(&dst.AWSConfiguration)
	if err == nil {
		jsonAWSConfiguration, _ := json.Marshal(dst.AWSConfiguration)
		if string(jsonAWSConfiguration) == "{}" { // empty struct
			dst.AWSConfiguration = nil
		} else {
			match++
		}
	} else {
		dst.AWSConfiguration = nil
	}

	// try to unmarshal data into AzureConfiguration
	err = newStrictDecoder(data).Decode(&dst.AzureConfiguration)
	if err == nil {
		jsonAzureConfiguration, _ := json.Marshal(dst.AzureConfiguration)
		if string(jsonAzureConfiguration) == "{}" { // empty struct
			dst.AzureConfiguration = nil
		} else {
			match++
		}
	} else {
		dst.AzureConfiguration = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AWSConfiguration = nil
		dst.AzureConfiguration = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Configuration)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Configuration)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Configuration) MarshalJSON() ([]byte, error) {
	if src.AWSConfiguration != nil {
		return json.Marshal(&src.AWSConfiguration)
	}

	if src.AzureConfiguration != nil {
		return json.Marshal(&src.AzureConfiguration)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Configuration) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AWSConfiguration != nil {
		return obj.AWSConfiguration
	}

	if obj.AzureConfiguration != nil {
		return obj.AzureConfiguration
	}

	// all schemas are nil
	return nil
}

type NullableConfiguration struct {
	value *Configuration
	isSet bool
}

func (v NullableConfiguration) Get() *Configuration {
	return v.value
}

func (v *NullableConfiguration) Set(val *Configuration) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguration(val *Configuration) *NullableConfiguration {
	return &NullableConfiguration{value: val, isSet: true}
}

func (v NullableConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

