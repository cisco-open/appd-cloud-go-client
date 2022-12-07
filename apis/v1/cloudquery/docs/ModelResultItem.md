# ModelResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Specifies a type of the value in the JSON response. | [optional] 
**Hints** | Pointer to [**Hints**](Hints.md) |  | [optional] 

## Methods

### NewModelResultItem

`func NewModelResultItem() *ModelResultItem`

NewModelResultItem instantiates a new ModelResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResultItemWithDefaults

`func NewModelResultItemWithDefaults() *ModelResultItem`

NewModelResultItemWithDefaults instantiates a new ModelResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ModelResultItem) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelResultItem) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelResultItem) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelResultItem) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetType

`func (o *ModelResultItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelResultItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelResultItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelResultItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHints

`func (o *ModelResultItem) GetHints() Hints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *ModelResultItem) GetHintsOk() (*Hints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *ModelResultItem) SetHints(v Hints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *ModelResultItem) HasHints() bool`

HasHints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


