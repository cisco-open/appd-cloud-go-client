# ScheduleInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **int32** | The polling interval is five (5) minutes | [default to 5]
**Unit** | **string** |  | [default to "minute"]

## Methods

### NewScheduleInterval

`func NewScheduleInterval(interval int32, unit string, ) *ScheduleInterval`

NewScheduleInterval instantiates a new ScheduleInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleIntervalWithDefaults

`func NewScheduleIntervalWithDefaults() *ScheduleInterval`

NewScheduleIntervalWithDefaults instantiates a new ScheduleInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *ScheduleInterval) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ScheduleInterval) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ScheduleInterval) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetUnit

`func (o *ScheduleInterval) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ScheduleInterval) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ScheduleInterval) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


