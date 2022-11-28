# ExecuteQuery200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Attribute identifying a response chunk as a heartbeat. | [optional] 
**Model** | Pointer to [**ModelReference**](ModelReference.md) |  | [optional] 
**Links** | Pointer to [**PaginationReference**](PaginationReference.md) |  | [optional] 
**Dataset** | Pointer to **string** | Name of the dataset. May be used as a reference in other datasets when returning multi dimensional data. | [optional] 
**Metadata** | Pointer to [**MetadataResultItem**](MetadataResultItem.md) |  | [optional] 
**Data** | Pointer to [**[][]DataResultItemInner**]([]DataResultItemInner.md) |  | [optional] 
**Error** | Pointer to [**ErrorResultChunkError**](ErrorResultChunkError.md) |  | [optional] 

## Methods

### NewExecuteQuery200Response

`func NewExecuteQuery200Response() *ExecuteQuery200Response`

NewExecuteQuery200Response instantiates a new ExecuteQuery200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteQuery200ResponseWithDefaults

`func NewExecuteQuery200ResponseWithDefaults() *ExecuteQuery200Response`

NewExecuteQuery200ResponseWithDefaults instantiates a new ExecuteQuery200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExecuteQuery200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExecuteQuery200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExecuteQuery200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExecuteQuery200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetModel

`func (o *ExecuteQuery200Response) GetModel() ModelReference`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ExecuteQuery200Response) GetModelOk() (*ModelReference, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ExecuteQuery200Response) SetModel(v ModelReference)`

SetModel sets Model field to given value.

### HasModel

`func (o *ExecuteQuery200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetLinks

`func (o *ExecuteQuery200Response) GetLinks() PaginationReference`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExecuteQuery200Response) GetLinksOk() (*PaginationReference, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExecuteQuery200Response) SetLinks(v PaginationReference)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExecuteQuery200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDataset

`func (o *ExecuteQuery200Response) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *ExecuteQuery200Response) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *ExecuteQuery200Response) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *ExecuteQuery200Response) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetMetadata

`func (o *ExecuteQuery200Response) GetMetadata() MetadataResultItem`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExecuteQuery200Response) GetMetadataOk() (*MetadataResultItem, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExecuteQuery200Response) SetMetadata(v MetadataResultItem)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExecuteQuery200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetData

`func (o *ExecuteQuery200Response) GetData() [][]DataResultItemInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExecuteQuery200Response) GetDataOk() (*[][]DataResultItemInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExecuteQuery200Response) SetData(v [][]DataResultItemInner)`

SetData sets Data field to given value.

### HasData

`func (o *ExecuteQuery200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ExecuteQuery200Response) GetError() ErrorResultChunkError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExecuteQuery200Response) GetErrorOk() (*ErrorResultChunkError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExecuteQuery200Response) SetError(v ErrorResultChunkError)`

SetError sets Error field to given value.

### HasError

`func (o *ExecuteQuery200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


