# RotationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevokeRotatedAfter** | Pointer to **interface{}** | An ISO-8601 duration for how long the rotated secret will last before expiring. Input can not exceed 30 days. If not specified, defaults to 30 days. | [optional] 

## Methods

### NewRotationRequest

`func NewRotationRequest() *RotationRequest`

NewRotationRequest instantiates a new RotationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotationRequestWithDefaults

`func NewRotationRequestWithDefaults() *RotationRequest`

NewRotationRequestWithDefaults instantiates a new RotationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevokeRotatedAfter

`func (o *RotationRequest) GetRevokeRotatedAfter() interface{}`

GetRevokeRotatedAfter returns the RevokeRotatedAfter field if non-nil, zero value otherwise.

### GetRevokeRotatedAfterOk

`func (o *RotationRequest) GetRevokeRotatedAfterOk() (*interface{}, bool)`

GetRevokeRotatedAfterOk returns a tuple with the RevokeRotatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeRotatedAfter

`func (o *RotationRequest) SetRevokeRotatedAfter(v interface{})`

SetRevokeRotatedAfter sets RevokeRotatedAfter field to given value.

### HasRevokeRotatedAfter

`func (o *RotationRequest) HasRevokeRotatedAfter() bool`

HasRevokeRotatedAfter returns a boolean if a field has been set.

### SetRevokeRotatedAfterNil

`func (o *RotationRequest) SetRevokeRotatedAfterNil(b bool)`

 SetRevokeRotatedAfterNil sets the value for RevokeRotatedAfter to be an explicit nil

### UnsetRevokeRotatedAfter
`func (o *RotationRequest) UnsetRevokeRotatedAfter()`

UnsetRevokeRotatedAfter ensures that no value is present for RevokeRotatedAfter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


