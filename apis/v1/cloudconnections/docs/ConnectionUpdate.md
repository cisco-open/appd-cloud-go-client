# ConnectionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Updated name for the connection | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ConfigurationId** | Pointer to **string** | Configuration ID assigned to the connection | [optional] 
**State** | Pointer to **string** | Set the state of the connection | [optional] 
**Details** | Pointer to [**ConnectionUpdateDetails**](ConnectionUpdateDetails.md) |  | [optional] 

## Methods

### NewConnectionUpdate

`func NewConnectionUpdate() *ConnectionUpdate`

NewConnectionUpdate instantiates a new ConnectionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionUpdateWithDefaults

`func NewConnectionUpdateWithDefaults() *ConnectionUpdate`

NewConnectionUpdateWithDefaults instantiates a new ConnectionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ConnectionUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectionUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectionUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectionUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectionUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConnectionUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConnectionUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConfigurationId

`func (o *ConnectionUpdate) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *ConnectionUpdate) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *ConnectionUpdate) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *ConnectionUpdate) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetState

`func (o *ConnectionUpdate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectionUpdate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectionUpdate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ConnectionUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDetails

`func (o *ConnectionUpdate) GetDetails() ConnectionUpdateDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConnectionUpdate) GetDetailsOk() (*ConnectionUpdateDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConnectionUpdate) SetDetails(v ConnectionUpdateDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ConnectionUpdate) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


