# DataResultChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Attribute identifying a response chunk holding data. | [optional] 
**Links** | Pointer to [**PaginationReference**](PaginationReference.md) |  | [optional] 
**Dataset** | Pointer to **string** | Name of the dataset. May be used as a reference in other datasets when returning multi dimensional data. | [optional] 
**Model** | Pointer to [**ModelReference**](ModelReference.md) |  | [optional] 
**Metadata** | Pointer to [**MetadataResultItem**](MetadataResultItem.md) |  | [optional] 
**Data** | Pointer to [**[][]DataResultItemInner**]([]DataResultItemInner.md) |  | [optional] 

## Methods

### NewDataResultChunk

`func NewDataResultChunk() *DataResultChunk`

NewDataResultChunk instantiates a new DataResultChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataResultChunkWithDefaults

`func NewDataResultChunkWithDefaults() *DataResultChunk`

NewDataResultChunkWithDefaults instantiates a new DataResultChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataResultChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataResultChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataResultChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DataResultChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *DataResultChunk) GetLinks() PaginationReference`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DataResultChunk) GetLinksOk() (*PaginationReference, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DataResultChunk) SetLinks(v PaginationReference)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DataResultChunk) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDataset

`func (o *DataResultChunk) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *DataResultChunk) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *DataResultChunk) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *DataResultChunk) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetModel

`func (o *DataResultChunk) GetModel() ModelReference`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DataResultChunk) GetModelOk() (*ModelReference, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DataResultChunk) SetModel(v ModelReference)`

SetModel sets Model field to given value.

### HasModel

`func (o *DataResultChunk) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMetadata

`func (o *DataResultChunk) GetMetadata() MetadataResultItem`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DataResultChunk) GetMetadataOk() (*MetadataResultItem, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DataResultChunk) SetMetadata(v MetadataResultItem)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DataResultChunk) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetData

`func (o *DataResultChunk) GetData() [][]DataResultItemInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DataResultChunk) GetDataOk() (*[][]DataResultItemInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DataResultChunk) SetData(v [][]DataResultItemInner)`

SetData sets Data field to given value.

### HasData

`func (o *DataResultChunk) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


