# ModelResultChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Attribute identifying a response chunk as holding the response model. | [optional] 
**Model** | Pointer to [**Model**](Model.md) |  | [optional] 

## Methods

### NewModelResultChunk

`func NewModelResultChunk() *ModelResultChunk`

NewModelResultChunk instantiates a new ModelResultChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResultChunkWithDefaults

`func NewModelResultChunkWithDefaults() *ModelResultChunk`

NewModelResultChunkWithDefaults instantiates a new ModelResultChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelResultChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelResultChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelResultChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelResultChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetModel

`func (o *ModelResultChunk) GetModel() Model`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModelResultChunk) GetModelOk() (*Model, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModelResultChunk) SetModel(v Model)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModelResultChunk) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


