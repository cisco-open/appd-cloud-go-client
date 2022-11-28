# AWSRoleDelegationUpdateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**ConnectionAccessType**](ConnectionAccessType.md) |  | [default to ROLE_DELEGATION]
**RoleName** | **string** |  | 

## Methods

### NewAWSRoleDelegationUpdateDetails

`func NewAWSRoleDelegationUpdateDetails(accessType ConnectionAccessType, roleName string, ) *AWSRoleDelegationUpdateDetails`

NewAWSRoleDelegationUpdateDetails instantiates a new AWSRoleDelegationUpdateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSRoleDelegationUpdateDetailsWithDefaults

`func NewAWSRoleDelegationUpdateDetailsWithDefaults() *AWSRoleDelegationUpdateDetails`

NewAWSRoleDelegationUpdateDetailsWithDefaults instantiates a new AWSRoleDelegationUpdateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *AWSRoleDelegationUpdateDetails) GetAccessType() ConnectionAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AWSRoleDelegationUpdateDetails) GetAccessTypeOk() (*ConnectionAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AWSRoleDelegationUpdateDetails) SetAccessType(v ConnectionAccessType)`

SetAccessType sets AccessType field to given value.


### GetRoleName

`func (o *AWSRoleDelegationUpdateDetails) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AWSRoleDelegationUpdateDetails) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AWSRoleDelegationUpdateDetails) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


