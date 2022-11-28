# AWSAccessKeyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | AWS Access keys are long-term credentials for an AWS IAM user, or account root user. The access key ID is one of two access keys needed to authenticate to AWS. The other is a secret access key. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell. | 
**SecretAccessKey** | **string** | The secret access key is one of two access keys needed to authenticate to AWS. The other is an access key ID. The secret access key is only available once, when you create it. Download the generated secret access key and save in a secure location. If the secret access key is lost or deleted, you must create a new one. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell. | 

## Methods

### NewAWSAccessKeyDetails

`func NewAWSAccessKeyDetails(accessKeyId string, secretAccessKey string, ) *AWSAccessKeyDetails`

NewAWSAccessKeyDetails instantiates a new AWSAccessKeyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSAccessKeyDetailsWithDefaults

`func NewAWSAccessKeyDetailsWithDefaults() *AWSAccessKeyDetails`

NewAWSAccessKeyDetailsWithDefaults instantiates a new AWSAccessKeyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AWSAccessKeyDetails) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AWSAccessKeyDetails) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AWSAccessKeyDetails) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetSecretAccessKey

`func (o *AWSAccessKeyDetails) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AWSAccessKeyDetails) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AWSAccessKeyDetails) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


