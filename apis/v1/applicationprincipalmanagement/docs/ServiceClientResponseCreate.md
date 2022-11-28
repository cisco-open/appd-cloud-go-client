# ServiceClientResponseCreate

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
**ClientSecret** | Pointer to **string** | The client&#39;s secret, used to authenticate during an oAuth token request. | [optional] 

## Methods

### NewServiceClientResponseCreate

`func NewServiceClientResponseCreate() *ServiceClientResponseCreate`

NewServiceClientResponseCreate instantiates a new ServiceClientResponseCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceClientResponseCreateWithDefaults

`func NewServiceClientResponseCreateWithDefaults() *ServiceClientResponseCreate`

NewServiceClientResponseCreateWithDefaults instantiates a new ServiceClientResponseCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ServiceClientResponseCreate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ServiceClientResponseCreate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ServiceClientResponseCreate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ServiceClientResponseCreate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceClientResponseCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceClientResponseCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceClientResponseCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceClientResponseCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthType

`func (o *ServiceClientResponseCreate) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ServiceClientResponseCreate) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ServiceClientResponseCreate) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *ServiceClientResponseCreate) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetId

`func (o *ServiceClientResponseCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceClientResponseCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceClientResponseCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceClientResponseCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHasRotatedSecrets

`func (o *ServiceClientResponseCreate) GetHasRotatedSecrets() bool`

GetHasRotatedSecrets returns the HasRotatedSecrets field if non-nil, zero value otherwise.

### GetHasRotatedSecretsOk

`func (o *ServiceClientResponseCreate) GetHasRotatedSecretsOk() (*bool, bool)`

GetHasRotatedSecretsOk returns a tuple with the HasRotatedSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRotatedSecrets

`func (o *ServiceClientResponseCreate) SetHasRotatedSecrets(v bool)`

SetHasRotatedSecrets sets HasRotatedSecrets field to given value.

### HasHasRotatedSecrets

`func (o *ServiceClientResponseCreate) HasHasRotatedSecrets() bool`

HasHasRotatedSecrets returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceClientResponseCreate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceClientResponseCreate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceClientResponseCreate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceClientResponseCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceClientResponseCreate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceClientResponseCreate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceClientResponseCreate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceClientResponseCreate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetClientSecret

`func (o *ServiceClientResponseCreate) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ServiceClientResponseCreate) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ServiceClientResponseCreate) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ServiceClientResponseCreate) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


