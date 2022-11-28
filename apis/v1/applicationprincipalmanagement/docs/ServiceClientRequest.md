# ServiceClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name for the client. | [optional] 
**Description** | Pointer to **string** | A user provided description of the client. | [optional] 
**AuthType** | Pointer to **string** | Supported authentication methods used to request oAuth tokens: * &#x60;client_secret_basic&#x60; - The client credentials will be sent in the authorization header. * &#x60;client_secret_post&#x60; - The client credentials will be sent in the request body. | [optional] 

## Methods

### NewServiceClientRequest

`func NewServiceClientRequest() *ServiceClientRequest`

NewServiceClientRequest instantiates a new ServiceClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceClientRequestWithDefaults

`func NewServiceClientRequestWithDefaults() *ServiceClientRequest`

NewServiceClientRequestWithDefaults instantiates a new ServiceClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ServiceClientRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ServiceClientRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ServiceClientRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ServiceClientRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceClientRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceClientRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceClientRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceClientRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthType

`func (o *ServiceClientRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ServiceClientRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ServiceClientRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *ServiceClientRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


