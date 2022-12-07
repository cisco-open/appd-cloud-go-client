# ImportTagConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | It is true by default. Tags will be imported for all the resources that are being monitored by default | [default to true]
**ExcludedKeys** | Pointer to **[]string** | Array of that need to be excluded from being imported. It can be set only when enabled is true | [optional] 

## Methods

### NewImportTagConfiguration

`func NewImportTagConfiguration(enabled bool, ) *ImportTagConfiguration`

NewImportTagConfiguration instantiates a new ImportTagConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTagConfigurationWithDefaults

`func NewImportTagConfigurationWithDefaults() *ImportTagConfiguration`

NewImportTagConfigurationWithDefaults instantiates a new ImportTagConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ImportTagConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ImportTagConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ImportTagConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExcludedKeys

`func (o *ImportTagConfiguration) GetExcludedKeys() []string`

GetExcludedKeys returns the ExcludedKeys field if non-nil, zero value otherwise.

### GetExcludedKeysOk

`func (o *ImportTagConfiguration) GetExcludedKeysOk() (*[]string, bool)`

GetExcludedKeysOk returns a tuple with the ExcludedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedKeys

`func (o *ImportTagConfiguration) SetExcludedKeys(v []string)`

SetExcludedKeys sets ExcludedKeys field to given value.

### HasExcludedKeys

`func (o *ImportTagConfiguration) HasExcludedKeys() bool`

HasExcludedKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


