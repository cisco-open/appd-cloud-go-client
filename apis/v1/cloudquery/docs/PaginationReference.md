# PaginationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**PaginationReferenceNext**](PaginationReferenceNext.md) |  | [optional] 

## Methods

### NewPaginationReference

`func NewPaginationReference() *PaginationReference`

NewPaginationReference instantiates a new PaginationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationReferenceWithDefaults

`func NewPaginationReferenceWithDefaults() *PaginationReference`

NewPaginationReferenceWithDefaults instantiates a new PaginationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginationReference) GetNext() PaginationReferenceNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginationReference) GetNextOk() (*PaginationReferenceNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginationReference) SetNext(v PaginationReferenceNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginationReference) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


