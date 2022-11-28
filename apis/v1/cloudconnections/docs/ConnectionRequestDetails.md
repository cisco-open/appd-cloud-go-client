# ConnectionRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client IDs, also known as Application IDs, are long-term credentials for an Azure user, or account root user. The Client ID is one of three properties needed to authenticate to Azure, the other two being Client Secret and Tenant (Directory) ID | 
**ClientSecret** | **string** | A Client Secret allows an Azure application to provide its identity when requesting an access token. The Client Secret is one of three properties needed to authenticate to Azure, the other two being Client ID (Application ID) and Tenant (Directory) ID | 
**TenantId** | **string** | The Azure AD Tenant (Directory) IDis one of three properties needed to authenticate to Azure. The other two are Client Secret and Client ID (Application ID). | 
**SubscriptionId** | **string** | Specify a GUID Subscription ID to monitor. If monitoring all subscriptions, do not specify a Subscription ID. | 

## Methods

### NewConnectionRequestDetails

`func NewConnectionRequestDetails(clientId string, clientSecret string, tenantId string, subscriptionId string, ) *ConnectionRequestDetails`

NewConnectionRequestDetails instantiates a new ConnectionRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRequestDetailsWithDefaults

`func NewConnectionRequestDetailsWithDefaults() *ConnectionRequestDetails`

NewConnectionRequestDetailsWithDefaults instantiates a new ConnectionRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ConnectionRequestDetails) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConnectionRequestDetails) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConnectionRequestDetails) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *ConnectionRequestDetails) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ConnectionRequestDetails) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ConnectionRequestDetails) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantId

`func (o *ConnectionRequestDetails) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ConnectionRequestDetails) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ConnectionRequestDetails) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *ConnectionRequestDetails) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ConnectionRequestDetails) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ConnectionRequestDetails) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


