# Description

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Name of the connection or configuration | 
**Description** | Pointer to **string** | Description for this connection or configuration | [optional] 

## Methods

### NewDescription

`func NewDescription(displayName string, ) *Description`

NewDescription instantiates a new Description object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescriptionWithDefaults

`func NewDescriptionWithDefaults() *Description`

NewDescriptionWithDefaults instantiates a new Description object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *Description) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Description) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Description) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *Description) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Description) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Description) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Description) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


