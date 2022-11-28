# AWSRoleDelegationCreationDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**ConnectionAccessType**](ConnectionAccessType.md) |  | [default to ROLE_DELEGATION]

## Methods

### NewAWSRoleDelegationCreationDetailsAllOf

`func NewAWSRoleDelegationCreationDetailsAllOf(accessType ConnectionAccessType, ) *AWSRoleDelegationCreationDetailsAllOf`

NewAWSRoleDelegationCreationDetailsAllOf instantiates a new AWSRoleDelegationCreationDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSRoleDelegationCreationDetailsAllOfWithDefaults

`func NewAWSRoleDelegationCreationDetailsAllOfWithDefaults() *AWSRoleDelegationCreationDetailsAllOf`

NewAWSRoleDelegationCreationDetailsAllOfWithDefaults instantiates a new AWSRoleDelegationCreationDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *AWSRoleDelegationCreationDetailsAllOf) GetAccessType() ConnectionAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AWSRoleDelegationCreationDetailsAllOf) GetAccessTypeOk() (*ConnectionAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AWSRoleDelegationCreationDetailsAllOf) SetAccessType(v ConnectionAccessType)`

SetAccessType sets AccessType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


