# AzureDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client IDs, also known as Application IDs, are long-term credentials for an Azure user, or account root user. The Client ID is one of three properties needed to authenticate to Azure, the other two being Client Secret and Tenant (Directory) ID | 
**ClientSecret** | **string** | A Client Secret allows an Azure application to provide its identity when requesting an access token. The Client Secret is one of three properties needed to authenticate to Azure, the other two being Client ID (Application ID) and Tenant (Directory) ID | 
**TenantId** | **string** | The Azure AD Tenant (Directory) IDis one of three properties needed to authenticate to Azure. The other two are Client Secret and Client ID (Application ID). | 
**SubscriptionId** | **string** | Specify a GUID Subscription ID to monitor. If monitoring all subscriptions, do not specify a Subscription ID. | 

## Methods

### NewAzureDetails

`func NewAzureDetails(clientId string, clientSecret string, tenantId string, subscriptionId string, ) *AzureDetails`

NewAzureDetails instantiates a new AzureDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDetailsWithDefaults

`func NewAzureDetailsWithDefaults() *AzureDetails`

NewAzureDetailsWithDefaults instantiates a new AzureDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AzureDetails) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureDetails) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureDetails) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AzureDetails) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureDetails) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureDetails) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantId

`func (o *AzureDetails) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureDetails) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureDetails) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *AzureDetails) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureDetails) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureDetails) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


