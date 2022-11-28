# BaseConfigurationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique identifier for this resource | 
**DisplayName** | **string** | display name for this resource | 
**Description** | Pointer to **string** | a description of this resource | [optional] 

## Methods

### NewBaseConfigurationResource

`func NewBaseConfigurationResource(id string, displayName string, ) *BaseConfigurationResource`

NewBaseConfigurationResource instantiates a new BaseConfigurationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseConfigurationResourceWithDefaults

`func NewBaseConfigurationResourceWithDefaults() *BaseConfigurationResource`

NewBaseConfigurationResourceWithDefaults instantiates a new BaseConfigurationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseConfigurationResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseConfigurationResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseConfigurationResource) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *BaseConfigurationResource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BaseConfigurationResource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BaseConfigurationResource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *BaseConfigurationResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseConfigurationResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseConfigurationResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseConfigurationResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


