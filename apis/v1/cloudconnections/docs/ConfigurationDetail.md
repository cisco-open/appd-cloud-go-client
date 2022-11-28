# ConfigurationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DisplayName** | **string** | Name of the connection or configuration | 
**Description** | Pointer to **string** | Description for this connection or configuration | [optional] 
**Type** | [**ProviderType**](ProviderType.md) |  | 
**Details** | [**AzureConfigurationDetails**](AzureConfigurationDetails.md) |  | 

## Methods

### NewConfigurationDetail

`func NewConfigurationDetail(id string, createdAt time.Time, updatedAt time.Time, displayName string, type_ ProviderType, details AzureConfigurationDetails, ) *ConfigurationDetail`

NewConfigurationDetail instantiates a new ConfigurationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationDetailWithDefaults

`func NewConfigurationDetailWithDefaults() *ConfigurationDetail`

NewConfigurationDetailWithDefaults instantiates a new ConfigurationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationDetail) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ConfigurationDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConfigurationDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConfigurationDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ConfigurationDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConfigurationDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConfigurationDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDisplayName

`func (o *ConfigurationDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConfigurationDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConfigurationDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *ConfigurationDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ConfigurationDetail) GetType() ProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigurationDetail) GetTypeOk() (*ProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigurationDetail) SetType(v ProviderType)`

SetType sets Type field to given value.


### GetDetails

`func (o *ConfigurationDetail) GetDetails() AzureConfigurationDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConfigurationDetail) GetDetailsOk() (*AzureConfigurationDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConfigurationDetail) SetDetails(v AzureConfigurationDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


