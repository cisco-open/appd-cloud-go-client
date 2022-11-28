# ConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | [**ConnectionRequestDetails**](ConnectionRequestDetails.md) |  | 

## Methods

### NewConnectionRequest

`func NewConnectionRequest(details ConnectionRequestDetails, ) *ConnectionRequest`

NewConnectionRequest instantiates a new ConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRequestWithDefaults

`func NewConnectionRequestWithDefaults() *ConnectionRequest`

NewConnectionRequestWithDefaults instantiates a new ConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ConnectionRequest) GetDetails() ConnectionRequestDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConnectionRequest) GetDetailsOk() (*ConnectionRequestDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConnectionRequest) SetDetails(v ConnectionRequestDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


