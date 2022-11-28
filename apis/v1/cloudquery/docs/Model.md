# Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | This name will be used by the following DataResultChunks as a reference to the description of their structure. | [optional] 
**Fields** | Pointer to [**[]ModelFieldsInner**](ModelFieldsInner.md) | An array of type descriptors for each fetched expression in the query. The order of the array matches the order of values in the DataResultChunks with this schema model.&#39; | [optional] 

## Methods

### NewModel

`func NewModel() *Model`

NewModel instantiates a new Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelWithDefaults

`func NewModelWithDefaults() *Model`

NewModelWithDefaults instantiates a new Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Model) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Model) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Model) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Model) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFields

`func (o *Model) GetFields() []ModelFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Model) GetFieldsOk() (*[]ModelFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Model) SetFields(v []ModelFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Model) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


