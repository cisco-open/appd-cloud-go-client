# ErrorResultChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Attribute identifying a response chunk as holding information of an error. | [optional] 
**Error** | Pointer to [**ErrorResultChunkError**](ErrorResultChunkError.md) |  | [optional] 

## Methods

### NewErrorResultChunk

`func NewErrorResultChunk() *ErrorResultChunk`

NewErrorResultChunk instantiates a new ErrorResultChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResultChunkWithDefaults

`func NewErrorResultChunkWithDefaults() *ErrorResultChunk`

NewErrorResultChunkWithDefaults instantiates a new ErrorResultChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorResultChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResultChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResultChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorResultChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetError

`func (o *ErrorResultChunk) GetError() ErrorResultChunkError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorResultChunk) GetErrorOk() (*ErrorResultChunkError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorResultChunk) SetError(v ErrorResultChunkError)`

SetError sets Error field to given value.

### HasError

`func (o *ErrorResultChunk) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


