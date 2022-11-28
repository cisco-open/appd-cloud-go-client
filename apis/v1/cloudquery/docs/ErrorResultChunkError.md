# ErrorResultChunkError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | URI to the documentation of the error. | [optional] 
**ErrorCode** | Pointer to **string** | The code of the error. | [optional] 
**Timestamp** | Pointer to **time.Time** | Time and date the &#x60;ErrorResultChunk&#x60; was created. | [optional] 
**TraceId** | Pointer to **string** | It correlates events and logs messages across dependent services for single business transaction. | [optional] 
**Title** | Pointer to **string** | Brief description of the error. | [optional] 
**Detail** | Pointer to **string** | Usually more detailed description of the error. | [optional] 
**Query** | Pointer to **string** | The input query. | [optional] 

## Methods

### NewErrorResultChunkError

`func NewErrorResultChunkError() *ErrorResultChunkError`

NewErrorResultChunkError instantiates a new ErrorResultChunkError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResultChunkErrorWithDefaults

`func NewErrorResultChunkErrorWithDefaults() *ErrorResultChunkError`

NewErrorResultChunkErrorWithDefaults instantiates a new ErrorResultChunkError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorResultChunkError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResultChunkError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResultChunkError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorResultChunkError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetErrorCode

`func (o *ErrorResultChunkError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorResultChunkError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorResultChunkError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ErrorResultChunkError) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *ErrorResultChunkError) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ErrorResultChunkError) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ErrorResultChunkError) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ErrorResultChunkError) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *ErrorResultChunkError) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ErrorResultChunkError) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ErrorResultChunkError) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ErrorResultChunkError) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetTitle

`func (o *ErrorResultChunkError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorResultChunkError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorResultChunkError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ErrorResultChunkError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *ErrorResultChunkError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorResultChunkError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorResultChunkError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorResultChunkError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetQuery

`func (o *ErrorResultChunkError) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ErrorResultChunkError) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ErrorResultChunkError) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ErrorResultChunkError) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


