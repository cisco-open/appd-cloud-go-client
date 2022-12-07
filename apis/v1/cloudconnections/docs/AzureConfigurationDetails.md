# AzureConfigurationDetails

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

### NewAzureConfigurationDetails

`func NewAzureConfigurationDetails() *AzureConfigurationDetails`

NewAzureConfigurationDetails instantiates a new AzureConfigurationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConfigurationDetailsWithDefaults

`func NewAzureConfigurationDetailsWithDefaults() *AzureConfigurationDetails`

NewAzureConfigurationDetailsWithDefaults instantiates a new AzureConfigurationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *AzureConfigurationDetails) GetRegions() interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AzureConfigurationDetails) GetRegionsOk() (*interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AzureConfigurationDetails) SetRegions(v interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AzureConfigurationDetails) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *AzureConfigurationDetails) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *AzureConfigurationDetails) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetPolling

`func (o *AzureConfigurationDetails) GetPolling() ScheduleInterval`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *AzureConfigurationDetails) GetPollingOk() (*ScheduleInterval, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *AzureConfigurationDetails) SetPolling(v ScheduleInterval)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *AzureConfigurationDetails) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetImportTags

`func (o *AzureConfigurationDetails) GetImportTags() ImportTagConfiguration`

GetImportTags returns the ImportTags field if non-nil, zero value otherwise.

### GetImportTagsOk

`func (o *AzureConfigurationDetails) GetImportTagsOk() (*ImportTagConfiguration, bool)`

GetImportTagsOk returns a tuple with the ImportTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTags

`func (o *AzureConfigurationDetails) SetImportTags(v ImportTagConfiguration)`

SetImportTags sets ImportTags field to given value.

### HasImportTags

`func (o *AzureConfigurationDetails) HasImportTags() bool`

HasImportTags returns a boolean if a field has been set.

### GetTagFilter

`func (o *AzureConfigurationDetails) GetTagFilter() string`

GetTagFilter returns the TagFilter field if non-nil, zero value otherwise.

### GetTagFilterOk

`func (o *AzureConfigurationDetails) GetTagFilterOk() (*string, bool)`

GetTagFilterOk returns a tuple with the TagFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilter

`func (o *AzureConfigurationDetails) SetTagFilter(v string)`

SetTagFilter sets TagFilter field to given value.

### HasTagFilter

`func (o *AzureConfigurationDetails) HasTagFilter() bool`

HasTagFilter returns a boolean if a field has been set.

### GetServices

`func (o *AzureConfigurationDetails) GetServices() interface{}`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AzureConfigurationDetails) GetServicesOk() (*interface{}, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AzureConfigurationDetails) SetServices(v interface{})`

SetServices sets Services field to given value.

### HasServices

`func (o *AzureConfigurationDetails) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *AzureConfigurationDetails) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *AzureConfigurationDetails) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetResourceGroups

`func (o *AzureConfigurationDetails) GetResourceGroups() []string`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *AzureConfigurationDetails) GetResourceGroupsOk() (*[]string, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *AzureConfigurationDetails) SetResourceGroups(v []string)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *AzureConfigurationDetails) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


