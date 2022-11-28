# ConfigurationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name of the configuration | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Details** | Pointer to [**ConfigurationUpdateDetails**](ConfigurationUpdateDetails.md) |  | [optional] 

## Methods

### NewConfigurationUpdate

`func NewConfigurationUpdate() *ConfigurationUpdate`

NewConfigurationUpdate instantiates a new ConfigurationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationUpdateWithDefaults

`func NewConfigurationUpdateWithDefaults() *ConfigurationUpdate`

NewConfigurationUpdateWithDefaults instantiates a new ConfigurationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ConfigurationUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConfigurationUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConfigurationUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConfigurationUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigurationUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigurationUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigurationUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDetails

`func (o *ConfigurationUpdate) GetDetails() ConfigurationUpdateDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConfigurationUpdate) GetDetailsOk() (*ConfigurationUpdateDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConfigurationUpdate) SetDetails(v ConfigurationUpdateDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ConfigurationUpdate) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


