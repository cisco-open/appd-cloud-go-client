# BadRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewBadRequestError

`func NewBadRequestError() *BadRequestError`

NewBadRequestError instantiates a new BadRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestErrorWithDefaults

`func NewBadRequestErrorWithDefaults() *BadRequestError`

NewBadRequestErrorWithDefaults instantiates a new BadRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BadRequestError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BadRequestError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BadRequestError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BadRequestError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *BadRequestError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BadRequestError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BadRequestError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BadRequestError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *BadRequestError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BadRequestError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BadRequestError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BadRequestError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *BadRequestError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BadRequestError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BadRequestError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *BadRequestError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


