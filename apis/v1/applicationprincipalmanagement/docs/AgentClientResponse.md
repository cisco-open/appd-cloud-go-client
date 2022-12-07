# AgentClientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name for the client. | [optional] 
**Description** | Pointer to **string** | A user provided description of the client. | [optional] 
**AuthType** | Pointer to **string** | Supported authentication methods used to request oAuth tokens: * &#x60;client_secret_basic&#x60; - The client credentials will be sent in the authorization header. * &#x60;client_secret_post&#x60; - The client credentials will be sent in the request body. | [optional] 
**Id** | Pointer to **string** | The unique client identifier. | [optional] 
**HasRotatedSecrets** | Pointer to **bool** | Indicates if the client has rotated secrets. Rotated client secrets can be revoked. | [optional] 
**CreatedAt** | Pointer to **string** | The RFC3339 timestamp when the client was created. | [optional] 
**UpdatedAt** | Pointer to **string** | The RFC3339 timestamp when the client was last updated. | [optional] 

## Methods

### NewAgentClientResponse

`func NewAgentClientResponse() *AgentClientResponse`

NewAgentClientResponse instantiates a new AgentClientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentClientResponseWithDefaults

`func NewAgentClientResponseWithDefaults() *AgentClientResponse`

NewAgentClientResponseWithDefaults instantiates a new AgentClientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AgentClientResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AgentClientResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AgentClientResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AgentClientResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *AgentClientResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentClientResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentClientResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentClientResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthType

`func (o *AgentClientResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AgentClientResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AgentClientResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AgentClientResponse) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetId

`func (o *AgentClientResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentClientResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentClientResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentClientResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHasRotatedSecrets

`func (o *AgentClientResponse) GetHasRotatedSecrets() bool`

GetHasRotatedSecrets returns the HasRotatedSecrets field if non-nil, zero value otherwise.

### GetHasRotatedSecretsOk

`func (o *AgentClientResponse) GetHasRotatedSecretsOk() (*bool, bool)`

GetHasRotatedSecretsOk returns a tuple with the HasRotatedSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRotatedSecrets

`func (o *AgentClientResponse) SetHasRotatedSecrets(v bool)`

SetHasRotatedSecrets sets HasRotatedSecrets field to given value.

### HasHasRotatedSecrets

`func (o *AgentClientResponse) HasHasRotatedSecrets() bool`

HasHasRotatedSecrets returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentClientResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentClientResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentClientResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentClientResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentClientResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentClientResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentClientResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentClientResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


