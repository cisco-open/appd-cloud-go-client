# AWSConnectionResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**ConnectionAccessType**](ConnectionAccessType.md) |  | [default to ROLE_DELEGATION]
**AccountId** | **string** |  | 
**RoleName** | **string** |  | 
**AppDynamicsAwsAccountId** | Pointer to **string** | AppDynamics AWS Account ID. Delegates a user to an Identity Access Management (IAM) role in AWS. The AWS IAM role provides AppDynamics access to resources. | [optional] 
**ExternalId** | Pointer to **string** | Returns an external ID for AWS role delegation connections | [optional] 

## Methods

### NewAWSConnectionResponseDetails

`func NewAWSConnectionResponseDetails(accessType ConnectionAccessType, accountId string, roleName string, ) *AWSConnectionResponseDetails`

NewAWSConnectionResponseDetails instantiates a new AWSConnectionResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSConnectionResponseDetailsWithDefaults

`func NewAWSConnectionResponseDetailsWithDefaults() *AWSConnectionResponseDetails`

NewAWSConnectionResponseDetailsWithDefaults instantiates a new AWSConnectionResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *AWSConnectionResponseDetails) GetAccessType() ConnectionAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AWSConnectionResponseDetails) GetAccessTypeOk() (*ConnectionAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AWSConnectionResponseDetails) SetAccessType(v ConnectionAccessType)`

SetAccessType sets AccessType field to given value.


### GetAccountId

`func (o *AWSConnectionResponseDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSConnectionResponseDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSConnectionResponseDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetRoleName

`func (o *AWSConnectionResponseDetails) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AWSConnectionResponseDetails) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AWSConnectionResponseDetails) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetAppDynamicsAwsAccountId

`func (o *AWSConnectionResponseDetails) GetAppDynamicsAwsAccountId() string`

GetAppDynamicsAwsAccountId returns the AppDynamicsAwsAccountId field if non-nil, zero value otherwise.

### GetAppDynamicsAwsAccountIdOk

`func (o *AWSConnectionResponseDetails) GetAppDynamicsAwsAccountIdOk() (*string, bool)`

GetAppDynamicsAwsAccountIdOk returns a tuple with the AppDynamicsAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDynamicsAwsAccountId

`func (o *AWSConnectionResponseDetails) SetAppDynamicsAwsAccountId(v string)`

SetAppDynamicsAwsAccountId sets AppDynamicsAwsAccountId field to given value.

### HasAppDynamicsAwsAccountId

`func (o *AWSConnectionResponseDetails) HasAppDynamicsAwsAccountId() bool`

HasAppDynamicsAwsAccountId returns a boolean if a field has been set.

### GetExternalId

`func (o *AWSConnectionResponseDetails) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AWSConnectionResponseDetails) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AWSConnectionResponseDetails) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AWSConnectionResponseDetails) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


