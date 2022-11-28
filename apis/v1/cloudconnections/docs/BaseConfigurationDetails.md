# BaseConfigurationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to **[]string** | Geographic locations used to fetch metrics | [optional] 
**Polling** | Pointer to [**ScheduleInterval**](ScheduleInterval.md) |  | [optional] 
**ImportTags** | Pointer to [**ImportTagConfiguration**](ImportTagConfiguration.md) |  | [optional] 
**TagFilter** | Pointer to **string** | expression for filtering resources to be monitored, based on tags | [optional] 
**Services** | Pointer to [**[]Service**](Service.md) | services for which we will fetch metrics | [optional] 

## Methods

### NewBaseConfigurationDetails

`func NewBaseConfigurationDetails() *BaseConfigurationDetails`

NewBaseConfigurationDetails instantiates a new BaseConfigurationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseConfigurationDetailsWithDefaults

`func NewBaseConfigurationDetailsWithDefaults() *BaseConfigurationDetails`

NewBaseConfigurationDetailsWithDefaults instantiates a new BaseConfigurationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *BaseConfigurationDetails) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *BaseConfigurationDetails) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *BaseConfigurationDetails) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *BaseConfigurationDetails) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetPolling

`func (o *BaseConfigurationDetails) GetPolling() ScheduleInterval`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *BaseConfigurationDetails) GetPollingOk() (*ScheduleInterval, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *BaseConfigurationDetails) SetPolling(v ScheduleInterval)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *BaseConfigurationDetails) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetImportTags

`func (o *BaseConfigurationDetails) GetImportTags() ImportTagConfiguration`

GetImportTags returns the ImportTags field if non-nil, zero value otherwise.

### GetImportTagsOk

`func (o *BaseConfigurationDetails) GetImportTagsOk() (*ImportTagConfiguration, bool)`

GetImportTagsOk returns a tuple with the ImportTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTags

`func (o *BaseConfigurationDetails) SetImportTags(v ImportTagConfiguration)`

SetImportTags sets ImportTags field to given value.

### HasImportTags

`func (o *BaseConfigurationDetails) HasImportTags() bool`

HasImportTags returns a boolean if a field has been set.

### GetTagFilter

`func (o *BaseConfigurationDetails) GetTagFilter() string`

GetTagFilter returns the TagFilter field if non-nil, zero value otherwise.

### GetTagFilterOk

`func (o *BaseConfigurationDetails) GetTagFilterOk() (*string, bool)`

GetTagFilterOk returns a tuple with the TagFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilter

`func (o *BaseConfigurationDetails) SetTagFilter(v string)`

SetTagFilter sets TagFilter field to given value.

### HasTagFilter

`func (o *BaseConfigurationDetails) HasTagFilter() bool`

HasTagFilter returns a boolean if a field has been set.

### GetServices

`func (o *BaseConfigurationDetails) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BaseConfigurationDetails) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BaseConfigurationDetails) SetServices(v []Service)`

SetServices sets Services field to given value.

### HasServices

`func (o *BaseConfigurationDetails) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


