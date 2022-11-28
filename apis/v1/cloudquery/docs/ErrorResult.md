# ErrorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | URI to the documentation of the error. | 
**ErrorCode** | Pointer to **string** | The code of a error. | [optional] 
**Timestamp** | **time.Time** | Time and date the error occurred. | 
**TraceId** | **string** | Id to correlate events and logs across dependent services for a single query. | 
**Title** | **string** | Brief description of the error. | 
**Detail** | **string** | Usually a more detailed description of the error. | 
**TenantId** | **string** | The id of a tenant for which the query was executed. | 
**Query** | **string** | The input query. | 
**ErrorDetails** | Pointer to [**[]ErrorDetailsArrayInner**](ErrorDetailsArrayInner.md) | An array containing the details of an error. | [optional] 

## Methods

### NewErrorResult

`func NewErrorResult(type_ string, timestamp time.Time, traceId string, title string, detail string, tenantId string, query string, ) *ErrorResult`

NewErrorResult instantiates a new ErrorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResultWithDefaults

`func NewErrorResultWithDefaults() *ErrorResult`

NewErrorResultWithDefaults instantiates a new ErrorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResult) SetType(v string)`

SetType sets Type field to given value.


### GetErrorCode

`func (o *ErrorResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ErrorResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *ErrorResult) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ErrorResult) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ErrorResult) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTraceId

`func (o *ErrorResult) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ErrorResult) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ErrorResult) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetTitle

`func (o *ErrorResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorResult) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ErrorResult) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorResult) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorResult) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetTenantId

`func (o *ErrorResult) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ErrorResult) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ErrorResult) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetQuery

`func (o *ErrorResult) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ErrorResult) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ErrorResult) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetErrorDetails

`func (o *ErrorResult) GetErrorDetails() []ErrorDetailsArrayInner`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ErrorResult) GetErrorDetailsOk() (*[]ErrorDetailsArrayInner, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ErrorResult) SetErrorDetails(v []ErrorDetailsArrayInner)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ErrorResult) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


