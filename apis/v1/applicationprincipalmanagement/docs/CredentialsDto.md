# CredentialsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The unique client identifier. | [optional] 
**ClientSecret** | Pointer to **string** | The client&#39;s secret, used to authenticate during an oAuth token request. | [optional] 
**RotatedSecretExpiresAt** | Pointer to **string** | The RFC3339 timestamp when the rotated client secret will expire. | [optional] 

## Methods

### NewCredentialsDto

`func NewCredentialsDto() *CredentialsDto`

NewCredentialsDto instantiates a new CredentialsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsDtoWithDefaults

`func NewCredentialsDtoWithDefaults() *CredentialsDto`

NewCredentialsDtoWithDefaults instantiates a new CredentialsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CredentialsDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CredentialsDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CredentialsDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CredentialsDto) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *CredentialsDto) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CredentialsDto) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CredentialsDto) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CredentialsDto) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetRotatedSecretExpiresAt

`func (o *CredentialsDto) GetRotatedSecretExpiresAt() string`

GetRotatedSecretExpiresAt returns the RotatedSecretExpiresAt field if non-nil, zero value otherwise.

### GetRotatedSecretExpiresAtOk

`func (o *CredentialsDto) GetRotatedSecretExpiresAtOk() (*string, bool)`

GetRotatedSecretExpiresAtOk returns a tuple with the RotatedSecretExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedSecretExpiresAt

`func (o *CredentialsDto) SetRotatedSecretExpiresAt(v string)`

SetRotatedSecretExpiresAt sets RotatedSecretExpiresAt field to given value.

### HasRotatedSecretExpiresAt

`func (o *CredentialsDto) HasRotatedSecretExpiresAt() bool`

HasRotatedSecretExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


