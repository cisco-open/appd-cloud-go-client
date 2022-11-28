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

// AgentClientRequest Request object for requests to manage Agent client identities. Extends BaseClientRequest.
type AgentClientRequest struct {
	// The display name for the client.
	DisplayName *string `json:"displayName,omitempty"`
	// A user provided description of the client.
	Description *string `json:"description,omitempty"`
	// Supported authentication methods used to request oAuth tokens: * `client_secret_basic` - The client credentials will be sent in the authorization header. * `client_secret_post` - The client credentials will be sent in the request body.
	AuthType *string `json:"authType,omitempty"`
}

// NewAgentClientRequest instantiates a new AgentClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentClientRequest() *AgentClientRequest {
	this := AgentClientRequest{}
	return &this
}

// NewAgentClientRequestWithDefaults instantiates a new AgentClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentClientRequestWithDefaults() *AgentClientRequest {
	this := AgentClientRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AgentClientRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AgentClientRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AgentClientRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AgentClientRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AgentClientRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AgentClientRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *AgentClientRequest) GetAuthType() string {
	if o == nil || isNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientRequest) GetAuthTypeOk() (*string, bool) {
	if o == nil || isNil(o.AuthType) {
    return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *AgentClientRequest) HasAuthType() bool {
	if o != nil && !isNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *AgentClientRequest) SetAuthType(v string) {
	o.AuthType = &v
}

func (o AgentClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
	}
	return json.Marshal(toSerialize)
}

type NullableAgentClientRequest struct {
	value *AgentClientRequest
	isSet bool
}

func (v NullableAgentClientRequest) Get() *AgentClientRequest {
	return v.value
}

func (v *NullableAgentClientRequest) Set(val *AgentClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentClientRequest(val *AgentClientRequest) *NullableAgentClientRequest {
	return &NullableAgentClientRequest{value: val, isSet: true}
}

func (v NullableAgentClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


