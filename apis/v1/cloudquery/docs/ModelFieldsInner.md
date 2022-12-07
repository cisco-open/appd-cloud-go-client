# ModelFieldsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Specifies a type of the complex value. Individual fields are described in the &#39;model&#39;. | [optional] 
**Hints** | Pointer to [**Hints**](Hints.md) |  | [optional] 
**Form** | Pointer to **string** | Describes where timeseries or complex objects are located in the response. Possible values are: inline, reference, link | [optional] 
**Model** | Pointer to [**Model**](Model.md) |  | [optional] 

## Methods

### NewModelFieldsInner

`func NewModelFieldsInner() *ModelFieldsInner`

NewModelFieldsInner instantiates a new ModelFieldsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelFieldsInnerWithDefaults

`func NewModelFieldsInnerWithDefaults() *ModelFieldsInner`

NewModelFieldsInnerWithDefaults instantiates a new ModelFieldsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ModelFieldsInner) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelFieldsInner) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelFieldsInner) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelFieldsInner) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetType

`func (o *ModelFieldsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelFieldsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelFieldsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelFieldsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHints

`func (o *ModelFieldsInner) GetHints() Hints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *ModelFieldsInner) GetHintsOk() (*Hints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *ModelFieldsInner) SetHints(v Hints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *ModelFieldsInner) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetForm

`func (o *ModelFieldsInner) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ModelFieldsInner) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ModelFieldsInner) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *ModelFieldsInner) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetModel

`func (o *ModelFieldsInner) GetModel() Model`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModelFieldsInner) GetModelOk() (*Model, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModelFieldsInner) SetModel(v Model)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModelFieldsInner) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


