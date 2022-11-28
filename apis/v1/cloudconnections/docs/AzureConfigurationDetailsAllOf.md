# AzureConfigurationDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroups** | Pointer to **[]string** | Azure Resource groups used to fetch metrics | [optional] 
**Services** | Pointer to **interface{}** |  | [optional] 
**Regions** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAzureConfigurationDetailsAllOf

`func NewAzureConfigurationDetailsAllOf() *AzureConfigurationDetailsAllOf`

NewAzureConfigurationDetailsAllOf instantiates a new AzureConfigurationDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConfigurationDetailsAllOfWithDefaults

`func NewAzureConfigurationDetailsAllOfWithDefaults() *AzureConfigurationDetailsAllOf`

NewAzureConfigurationDetailsAllOfWithDefaults instantiates a new AzureConfigurationDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroups

`func (o *AzureConfigurationDetailsAllOf) GetResourceGroups() []string`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *AzureConfigurationDetailsAllOf) GetResourceGroupsOk() (*[]string, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *AzureConfigurationDetailsAllOf) SetResourceGroups(v []string)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *AzureConfigurationDetailsAllOf) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.

### GetServices

`func (o *AzureConfigurationDetailsAllOf) GetServices() interface{}`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AzureConfigurationDetailsAllOf) GetServicesOk() (*interface{}, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AzureConfigurationDetailsAllOf) SetServices(v interface{})`

SetServices sets Services field to given value.

### HasServices

`func (o *AzureConfigurationDetailsAllOf) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *AzureConfigurationDetailsAllOf) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *AzureConfigurationDetailsAllOf) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetRegions

`func (o *AzureConfigurationDetailsAllOf) GetRegions() interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AzureConfigurationDetailsAllOf) GetRegionsOk() (*interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AzureConfigurationDetailsAllOf) SetRegions(v interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AzureConfigurationDetailsAllOf) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *AzureConfigurationDetailsAllOf) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *AzureConfigurationDetailsAllOf) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


