# Hints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Type of observed component from AppD MELT Model. Optional, if applicable | [optional] 
**Field** | Pointer to **string** | Name of the MELT model property mentioned in the &#39;kind&#39; hint. See AppD MELT Model  Optional, if applicable (for example not for expression) | [optional] 
**Type** | Pointer to **string** | Full name with namespace of the type. | [optional] 

## Methods

### NewHints

`func NewHints() *Hints`

NewHints instantiates a new Hints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHintsWithDefaults

`func NewHintsWithDefaults() *Hints`

NewHintsWithDefaults instantiates a new Hints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *Hints) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Hints) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Hints) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Hints) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetField

`func (o *Hints) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Hints) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Hints) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *Hints) HasField() bool`

HasField returns a boolean if a field has been set.

### GetType

`func (o *Hints) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Hints) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Hints) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Hints) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


