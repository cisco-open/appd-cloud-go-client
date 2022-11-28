# ConnectionDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**State** | **string** | Connection state | 
**StateMessage** | **map[string]interface{}** | Connection state message | 

## Methods

### NewConnectionDetailAllOf

`func NewConnectionDetailAllOf(id string, createdAt time.Time, updatedAt time.Time, state string, stateMessage map[string]interface{}, ) *ConnectionDetailAllOf`

NewConnectionDetailAllOf instantiates a new ConnectionDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionDetailAllOfWithDefaults

`func NewConnectionDetailAllOfWithDefaults() *ConnectionDetailAllOf`

NewConnectionDetailAllOfWithDefaults instantiates a new ConnectionDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionDetailAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionDetailAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionDetailAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ConnectionDetailAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectionDetailAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectionDetailAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ConnectionDetailAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConnectionDetailAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConnectionDetailAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetState

`func (o *ConnectionDetailAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectionDetailAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectionDetailAllOf) SetState(v string)`

SetState sets State field to given value.


### GetStateMessage

`func (o *ConnectionDetailAllOf) GetStateMessage() map[string]interface{}`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *ConnectionDetailAllOf) GetStateMessageOk() (*map[string]interface{}, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *ConnectionDetailAllOf) SetStateMessage(v map[string]interface{})`

SetStateMessage sets StateMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


