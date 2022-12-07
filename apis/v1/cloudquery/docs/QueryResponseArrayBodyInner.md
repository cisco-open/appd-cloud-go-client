# QueryResponseArrayBodyInner

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

### NewQueryResponseArrayBodyInner

`func NewQueryResponseArrayBodyInner() *QueryResponseArrayBodyInner`

NewQueryResponseArrayBodyInner instantiates a new QueryResponseArrayBodyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseArrayBodyInnerWithDefaults

`func NewQueryResponseArrayBodyInnerWithDefaults() *QueryResponseArrayBodyInner`

NewQueryResponseArrayBodyInnerWithDefaults instantiates a new QueryResponseArrayBodyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueryResponseArrayBodyInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryResponseArrayBodyInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryResponseArrayBodyInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryResponseArrayBodyInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetModel

`func (o *QueryResponseArrayBodyInner) GetModel() ModelReference`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *QueryResponseArrayBodyInner) GetModelOk() (*ModelReference, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *QueryResponseArrayBodyInner) SetModel(v ModelReference)`

SetModel sets Model field to given value.

### HasModel

`func (o *QueryResponseArrayBodyInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetLinks

`func (o *QueryResponseArrayBodyInner) GetLinks() PaginationReference`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *QueryResponseArrayBodyInner) GetLinksOk() (*PaginationReference, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *QueryResponseArrayBodyInner) SetLinks(v PaginationReference)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *QueryResponseArrayBodyInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDataset

`func (o *QueryResponseArrayBodyInner) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *QueryResponseArrayBodyInner) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *QueryResponseArrayBodyInner) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *QueryResponseArrayBodyInner) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetMetadata

`func (o *QueryResponseArrayBodyInner) GetMetadata() MetadataResultItem`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QueryResponseArrayBodyInner) GetMetadataOk() (*MetadataResultItem, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QueryResponseArrayBodyInner) SetMetadata(v MetadataResultItem)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QueryResponseArrayBodyInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetData

`func (o *QueryResponseArrayBodyInner) GetData() [][]DataResultItemInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryResponseArrayBodyInner) GetDataOk() (*[][]DataResultItemInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryResponseArrayBodyInner) SetData(v [][]DataResultItemInner)`

SetData sets Data field to given value.

### HasData

`func (o *QueryResponseArrayBodyInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *QueryResponseArrayBodyInner) GetError() ErrorResultChunkError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryResponseArrayBodyInner) GetErrorOk() (*ErrorResultChunkError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryResponseArrayBodyInner) SetError(v ErrorResultChunkError)`

SetError sets Error field to given value.

### HasError

`func (o *QueryResponseArrayBodyInner) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


