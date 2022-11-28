# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ImportTags** | Pointer to [**ImportTagConfiguration**](ImportTagConfiguration.md) |  | [optional] 
**TagFilter** | Pointer to **string** | expression for filtering resources to be monitored, based on tags | [optional] 
**Polling** | Pointer to [**ScheduleInterval**](ScheduleInterval.md) |  | [optional] 

## Methods

### NewService

`func NewService(name string, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetImportTags

`func (o *Service) GetImportTags() ImportTagConfiguration`

GetImportTags returns the ImportTags field if non-nil, zero value otherwise.

### GetImportTagsOk

`func (o *Service) GetImportTagsOk() (*ImportTagConfiguration, bool)`

GetImportTagsOk returns a tuple with the ImportTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTags

`func (o *Service) SetImportTags(v ImportTagConfiguration)`

SetImportTags sets ImportTags field to given value.

### HasImportTags

`func (o *Service) HasImportTags() bool`

HasImportTags returns a boolean if a field has been set.

### GetTagFilter

`func (o *Service) GetTagFilter() string`

GetTagFilter returns the TagFilter field if non-nil, zero value otherwise.

### GetTagFilterOk

`func (o *Service) GetTagFilterOk() (*string, bool)`

GetTagFilterOk returns a tuple with the TagFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilter

`func (o *Service) SetTagFilter(v string)`

SetTagFilter sets TagFilter field to given value.

### HasTagFilter

`func (o *Service) HasTagFilter() bool`

HasTagFilter returns a boolean if a field has been set.

### GetPolling

`func (o *Service) GetPolling() ScheduleInterval`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *Service) GetPollingOk() (*ScheduleInterval, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *Service) SetPolling(v ScheduleInterval)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *Service) HasPolling() bool`

HasPolling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


