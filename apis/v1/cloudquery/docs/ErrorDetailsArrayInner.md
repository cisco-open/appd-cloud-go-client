# ErrorDetailsArrayInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A short description on error cause. | 
**FixSuggestion** | Pointer to **string** | A hint to resolve the error. | [optional] 
**FixPossibilities** | [**[]ErrorDetailItemFixPossibilitiesInner**](ErrorDetailItemFixPossibilitiesInner.md) | A list of fix possibilities to resolve the error. | 
**ErrorType** | **string** | The type of the error. | 
**ErrorFrom** | **string** | The start position of the error in format &#39;lineNum:charIdx&#39;. | 
**ErrorTo** | **string** | The end position of the error in format &#39;lineNum:charIdx&#39;. | 

## Methods

### NewErrorDetailsArrayInner

`func NewErrorDetailsArrayInner(message string, fixPossibilities []ErrorDetailItemFixPossibilitiesInner, errorType string, errorFrom string, errorTo string, ) *ErrorDetailsArrayInner`

NewErrorDetailsArrayInner instantiates a new ErrorDetailsArrayInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailsArrayInnerWithDefaults

`func NewErrorDetailsArrayInnerWithDefaults() *ErrorDetailsArrayInner`

NewErrorDetailsArrayInnerWithDefaults instantiates a new ErrorDetailsArrayInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorDetailsArrayInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetailsArrayInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetailsArrayInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetFixSuggestion

`func (o *ErrorDetailsArrayInner) GetFixSuggestion() string`

GetFixSuggestion returns the FixSuggestion field if non-nil, zero value otherwise.

### GetFixSuggestionOk

`func (o *ErrorDetailsArrayInner) GetFixSuggestionOk() (*string, bool)`

GetFixSuggestionOk returns a tuple with the FixSuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixSuggestion

`func (o *ErrorDetailsArrayInner) SetFixSuggestion(v string)`

SetFixSuggestion sets FixSuggestion field to given value.

### HasFixSuggestion

`func (o *ErrorDetailsArrayInner) HasFixSuggestion() bool`

HasFixSuggestion returns a boolean if a field has been set.

### GetFixPossibilities

`func (o *ErrorDetailsArrayInner) GetFixPossibilities() []ErrorDetailItemFixPossibilitiesInner`

GetFixPossibilities returns the FixPossibilities field if non-nil, zero value otherwise.

### GetFixPossibilitiesOk

`func (o *ErrorDetailsArrayInner) GetFixPossibilitiesOk() (*[]ErrorDetailItemFixPossibilitiesInner, bool)`

GetFixPossibilitiesOk returns a tuple with the FixPossibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixPossibilities

`func (o *ErrorDetailsArrayInner) SetFixPossibilities(v []ErrorDetailItemFixPossibilitiesInner)`

SetFixPossibilities sets FixPossibilities field to given value.


### GetErrorType

`func (o *ErrorDetailsArrayInner) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ErrorDetailsArrayInner) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ErrorDetailsArrayInner) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetErrorFrom

`func (o *ErrorDetailsArrayInner) GetErrorFrom() string`

GetErrorFrom returns the ErrorFrom field if non-nil, zero value otherwise.

### GetErrorFromOk

`func (o *ErrorDetailsArrayInner) GetErrorFromOk() (*string, bool)`

GetErrorFromOk returns a tuple with the ErrorFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFrom

`func (o *ErrorDetailsArrayInner) SetErrorFrom(v string)`

SetErrorFrom sets ErrorFrom field to given value.


### GetErrorTo

`func (o *ErrorDetailsArrayInner) GetErrorTo() string`

GetErrorTo returns the ErrorTo field if non-nil, zero value otherwise.

### GetErrorToOk

`func (o *ErrorDetailsArrayInner) GetErrorToOk() (*string, bool)`

GetErrorToOk returns a tuple with the ErrorTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorTo

`func (o *ErrorDetailsArrayInner) SetErrorTo(v string)`

SetErrorTo sets ErrorTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


