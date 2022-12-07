# AWSConfigurationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to **interface{}** |  | [optional] 
**Polling** | Pointer to [**ScheduleInterval**](ScheduleInterval.md) |  | [optional] 
**ImportTags** | Pointer to [**ImportTagConfiguration**](ImportTagConfiguration.md) |  | [optional] 
**TagFilter** | Pointer to **string** | expression for filtering resources to be monitored, based on tags | [optional] 
**Services** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAWSConfigurationDetails

`func NewAWSConfigurationDetails() *AWSConfigurationDetails`

NewAWSConfigurationDetails instantiates a new AWSConfigurationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSConfigurationDetailsWithDefaults

`func NewAWSConfigurationDetailsWithDefaults() *AWSConfigurationDetails`

NewAWSConfigurationDetailsWithDefaults instantiates a new AWSConfigurationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *AWSConfigurationDetails) GetRegions() interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AWSConfigurationDetails) GetRegionsOk() (*interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AWSConfigurationDetails) SetRegions(v interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AWSConfigurationDetails) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *AWSConfigurationDetails) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *AWSConfigurationDetails) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetPolling

`func (o *AWSConfigurationDetails) GetPolling() ScheduleInterval`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *AWSConfigurationDetails) GetPollingOk() (*ScheduleInterval, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *AWSConfigurationDetails) SetPolling(v ScheduleInterval)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *AWSConfigurationDetails) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetImportTags

`func (o *AWSConfigurationDetails) GetImportTags() ImportTagConfiguration`

GetImportTags returns the ImportTags field if non-nil, zero value otherwise.

### GetImportTagsOk

`func (o *AWSConfigurationDetails) GetImportTagsOk() (*ImportTagConfiguration, bool)`

GetImportTagsOk returns a tuple with the ImportTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTags

`func (o *AWSConfigurationDetails) SetImportTags(v ImportTagConfiguration)`

SetImportTags sets ImportTags field to given value.

### HasImportTags

`func (o *AWSConfigurationDetails) HasImportTags() bool`

HasImportTags returns a boolean if a field has been set.

### GetTagFilter

`func (o *AWSConfigurationDetails) GetTagFilter() string`

GetTagFilter returns the TagFilter field if non-nil, zero value otherwise.

### GetTagFilterOk

`func (o *AWSConfigurationDetails) GetTagFilterOk() (*string, bool)`

GetTagFilterOk returns a tuple with the TagFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilter

`func (o *AWSConfigurationDetails) SetTagFilter(v string)`

SetTagFilter sets TagFilter field to given value.

### HasTagFilter

`func (o *AWSConfigurationDetails) HasTagFilter() bool`

HasTagFilter returns a boolean if a field has been set.

### GetServices

`func (o *AWSConfigurationDetails) GetServices() interface{}`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AWSConfigurationDetails) GetServicesOk() (*interface{}, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AWSConfigurationDetails) SetServices(v interface{})`

SetServices sets Services field to given value.

### HasServices

`func (o *AWSConfigurationDetails) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *AWSConfigurationDetails) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *AWSConfigurationDetails) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


