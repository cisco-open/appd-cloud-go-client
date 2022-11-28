# ComplexModelResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Specifies a type of the complex value. Individual fields are described in the &#39;model&#39;. | [optional] 
**Hints** | Pointer to [**Hints**](Hints.md) |  | [optional] 
**Form** | Pointer to **string** | Describes where timeseries or complex objects are located in the response. Possible values are: inline, reference, link | [optional] 
**Model** | Pointer to [**Model**](Model.md) |  | [optional] 

## Methods

### NewComplexModelResultItem

`func NewComplexModelResultItem() *ComplexModelResultItem`

NewComplexModelResultItem instantiates a new ComplexModelResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplexModelResultItemWithDefaults

`func NewComplexModelResultItemWithDefaults() *ComplexModelResultItem`

NewComplexModelResultItemWithDefaults instantiates a new ComplexModelResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ComplexModelResultItem) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ComplexModelResultItem) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ComplexModelResultItem) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ComplexModelResultItem) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetType

`func (o *ComplexModelResultItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComplexModelResultItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComplexModelResultItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ComplexModelResultItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHints

`func (o *ComplexModelResultItem) GetHints() Hints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *ComplexModelResultItem) GetHintsOk() (*Hints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *ComplexModelResultItem) SetHints(v Hints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *ComplexModelResultItem) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetForm

`func (o *ComplexModelResultItem) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ComplexModelResultItem) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ComplexModelResultItem) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *ComplexModelResultItem) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetModel

`func (o *ComplexModelResultItem) GetModel() Model`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComplexModelResultItem) GetModelOk() (*Model, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComplexModelResultItem) SetModel(v Model)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComplexModelResultItem) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


