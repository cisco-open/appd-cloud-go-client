# AWSRoleDelegationCreationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**ConnectionAccessType**](ConnectionAccessType.md) |  | [default to ROLE_DELEGATION]
**AccountId** | **string** |  | 

## Methods

### NewAWSRoleDelegationCreationDetails

`func NewAWSRoleDelegationCreationDetails(accessType ConnectionAccessType, accountId string, ) *AWSRoleDelegationCreationDetails`

NewAWSRoleDelegationCreationDetails instantiates a new AWSRoleDelegationCreationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSRoleDelegationCreationDetailsWithDefaults

`func NewAWSRoleDelegationCreationDetailsWithDefaults() *AWSRoleDelegationCreationDetails`

NewAWSRoleDelegationCreationDetailsWithDefaults instantiates a new AWSRoleDelegationCreationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *AWSRoleDelegationCreationDetails) GetAccessType() ConnectionAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AWSRoleDelegationCreationDetails) GetAccessTypeOk() (*ConnectionAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AWSRoleDelegationCreationDetails) SetAccessType(v ConnectionAccessType)`

SetAccessType sets AccessType field to given value.


### GetAccountId

`func (o *AWSRoleDelegationCreationDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSRoleDelegationCreationDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSRoleDelegationCreationDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


