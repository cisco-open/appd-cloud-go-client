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

// AgentClientResponse Response object for requests to manage Agent client identities.
type AgentClientResponse struct {
	// The display name for the client.
	DisplayName *string `json:"displayName,omitempty"`
	// A user provided description of the client.
	Description *string `json:"description,omitempty"`
	// Supported authentication methods used to request oAuth tokens: * `client_secret_basic` - The client credentials will be sent in the authorization header. * `client_secret_post` - The client credentials will be sent in the request body.
	AuthType *string `json:"authType,omitempty"`
	// The unique client identifier.
	Id *string `json:"id,omitempty"`
	// Indicates if the client has rotated secrets. Rotated client secrets can be revoked.
	HasRotatedSecrets *bool `json:"hasRotatedSecrets,omitempty"`
	// The RFC3339 timestamp when the client was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The RFC3339 timestamp when the client was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewAgentClientResponse instantiates a new AgentClientResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentClientResponse() *AgentClientResponse {
	this := AgentClientResponse{}
	return &this
}

// NewAgentClientResponseWithDefaults instantiates a new AgentClientResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentClientResponseWithDefaults() *AgentClientResponse {
	this := AgentClientResponse{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AgentClientResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AgentClientResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AgentClientResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AgentClientResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AgentClientResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AgentClientResponse) SetDescription(v string) {
	o.Description = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *AgentClientResponse) GetAuthType() string {
	if o == nil || isNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponse) GetAuthTypeOk() (*string, bool) {
	if o == nil || isNil(o.AuthType) {
    return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *AgentClientResponse) HasAuthType() bool {
	if o != nil && !isNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *AgentClientResponse) SetAuthType(v string) {
	o.AuthType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentClientResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentClientResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentClientResponse) SetId(v string) {
	o.Id = &v
}

// GetHasRotatedSecrets returns the HasRotatedSecrets field value if set, zero value otherwise.
func (o *AgentClientResponse) GetHasRotatedSecrets() bool {
	if o == nil || isNil(o.HasRotatedSecrets) {
		var ret bool
		return ret
	}
	return *o.HasRotatedSecrets
}

// GetHasRotatedSecretsOk returns a tuple with the HasRotatedSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponse) GetHasRotatedSecretsOk() (*bool, bool) {
	if o == nil || isNil(o.HasRotatedSecrets) {
    return nil, false
	}
	return o.HasRotatedSecrets, true
}

// HasHasRotatedSecrets returns a boolean if a field has been set.
func (o *AgentClientResponse) HasHasRotatedSecrets() bool {
	if o != nil && !isNil(o.HasRotatedSecrets) {
		return true
	}

	return false
}

// SetHasRotatedSecrets gets a reference to the given bool and assigns it to the HasRotatedSecrets field.
func (o *AgentClientResponse) SetHasRotatedSecrets(v bool) {
	o.HasRotatedSecrets = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AgentClientResponse) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AgentClientResponse) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AgentClientResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AgentClientResponse) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AgentClientResponse) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AgentClientResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o AgentClientResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.HasRotatedSecrets) {
		toSerialize["hasRotatedSecrets"] = o.HasRotatedSecrets
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAgentClientResponse struct {
	value *AgentClientResponse
	isSet bool
}

func (v NullableAgentClientResponse) Get() *AgentClientResponse {
	return v.value
}

func (v *NullableAgentClientResponse) Set(val *AgentClientResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentClientResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentClientResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentClientResponse(val *AgentClientResponse) *NullableAgentClientResponse {
	return &NullableAgentClientResponse{value: val, isSet: true}
}

func (v NullableAgentClientResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentClientResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


