# DataResultItemInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** | Unique name of the dataset. | [optional] 
**JsonPath** | Pointer to **string** | JSON Path that filters the correct dataset from the list of all datasets. | [optional] 

## Methods

### NewDataResultItemInner

`func NewDataResultItemInner() *DataResultItemInner`

NewDataResultItemInner instantiates a new DataResultItemInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataResultItemInnerWithDefaults

`func NewDataResultItemInnerWithDefaults() *DataResultItemInner`

NewDataResultItemInnerWithDefaults instantiates a new DataResultItemInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *DataResultItemInner) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *DataResultItemInner) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *DataResultItemInner) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *DataResultItemInner) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetJsonPath

`func (o *DataResultItemInner) GetJsonPath() string`

GetJsonPath returns the JsonPath field if non-nil, zero value otherwise.

### GetJsonPathOk

`func (o *DataResultItemInner) GetJsonPathOk() (*string, bool)`

GetJsonPathOk returns a tuple with the JsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPath

`func (o *DataResultItemInner) SetJsonPath(v string)`

SetJsonPath sets JsonPath field to given value.

### HasJsonPath

`func (o *DataResultItemInner) HasJsonPath() bool`

HasJsonPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


