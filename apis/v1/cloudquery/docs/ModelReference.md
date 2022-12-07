# ModelReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | A unique model name. | [optional] 
**JsonPath** | Pointer to **string** | JSON Path that filters the correct model from the ModelResultChunk. | [optional] 

## Methods

### NewModelReference

`func NewModelReference() *ModelReference`

NewModelReference instantiates a new ModelReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelReferenceWithDefaults

`func NewModelReferenceWithDefaults() *ModelReference`

NewModelReferenceWithDefaults instantiates a new ModelReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ModelReference) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModelReference) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModelReference) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModelReference) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetJsonPath

`func (o *ModelReference) GetJsonPath() string`

GetJsonPath returns the JsonPath field if non-nil, zero value otherwise.

### GetJsonPathOk

`func (o *ModelReference) GetJsonPathOk() (*string, bool)`

GetJsonPathOk returns a tuple with the JsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPath

`func (o *ModelReference) SetJsonPath(v string)`

SetJsonPath sets JsonPath field to given value.

### HasJsonPath

`func (o *ModelReference) HasJsonPath() bool`

HasJsonPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


