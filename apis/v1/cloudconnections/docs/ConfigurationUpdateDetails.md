# ConfigurationUpdateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to **interface{}** |  | [optional] 
**Polling** | Pointer to [**ScheduleInterval**](ScheduleInterval.md) |  | [optional] 
**ImportTags** | Pointer to [**ImportTagConfiguration**](ImportTagConfiguration.md) |  | [optional] 
**TagFilter** | Pointer to **string** | expression for filtering resources to be monitored, based on tags | [optional] 
**Services** | Pointer to **interface{}** |  | [optional] 
**ResourceGroups** | Pointer to **[]string** | Azure Resource groups used to fetch metrics | [optional] 

## Methods

### NewConfigurationUpdateDetails

`func NewConfigurationUpdateDetails() *ConfigurationUpdateDetails`

NewConfigurationUpdateDetails instantiates a new ConfigurationUpdateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationUpdateDetailsWithDefaults

`func NewConfigurationUpdateDetailsWithDefaults() *ConfigurationUpdateDetails`

NewConfigurationUpdateDetailsWithDefaults instantiates a new ConfigurationUpdateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *ConfigurationUpdateDetails) GetRegions() interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ConfigurationUpdateDetails) GetRegionsOk() (*interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ConfigurationUpdateDetails) SetRegions(v interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ConfigurationUpdateDetails) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *ConfigurationUpdateDetails) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *ConfigurationUpdateDetails) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetPolling

`func (o *ConfigurationUpdateDetails) GetPolling() ScheduleInterval`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *ConfigurationUpdateDetails) GetPollingOk() (*ScheduleInterval, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *ConfigurationUpdateDetails) SetPolling(v ScheduleInterval)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *ConfigurationUpdateDetails) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetImportTags

`func (o *ConfigurationUpdateDetails) GetImportTags() ImportTagConfiguration`

GetImportTags returns the ImportTags field if non-nil, zero value otherwise.

### GetImportTagsOk

`func (o *ConfigurationUpdateDetails) GetImportTagsOk() (*ImportTagConfiguration, bool)`

GetImportTagsOk returns a tuple with the ImportTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTags

`func (o *ConfigurationUpdateDetails) SetImportTags(v ImportTagConfiguration)`

SetImportTags sets ImportTags field to given value.

### HasImportTags

`func (o *ConfigurationUpdateDetails) HasImportTags() bool`

HasImportTags returns a boolean if a field has been set.

### GetTagFilter

`func (o *ConfigurationUpdateDetails) GetTagFilter() string`

GetTagFilter returns the TagFilter field if non-nil, zero value otherwise.

### GetTagFilterOk

`func (o *ConfigurationUpdateDetails) GetTagFilterOk() (*string, bool)`

GetTagFilterOk returns a tuple with the TagFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilter

`func (o *ConfigurationUpdateDetails) SetTagFilter(v string)`

SetTagFilter sets TagFilter field to given value.

### HasTagFilter

`func (o *ConfigurationUpdateDetails) HasTagFilter() bool`

HasTagFilter returns a boolean if a field has been set.

### GetServices

`func (o *ConfigurationUpdateDetails) GetServices() interface{}`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ConfigurationUpdateDetails) GetServicesOk() (*interface{}, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ConfigurationUpdateDetails) SetServices(v interface{})`

SetServices sets Services field to given value.

### HasServices

`func (o *ConfigurationUpdateDetails) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *ConfigurationUpdateDetails) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *ConfigurationUpdateDetails) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetResourceGroups

`func (o *ConfigurationUpdateDetails) GetResourceGroups() []string`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *ConfigurationUpdateDetails) GetResourceGroupsOk() (*[]string, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *ConfigurationUpdateDetails) SetResourceGroups(v []string)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *ConfigurationUpdateDetails) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


