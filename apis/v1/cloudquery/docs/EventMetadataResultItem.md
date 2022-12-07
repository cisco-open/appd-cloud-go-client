# EventMetadataResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **map[string]interface{}** | Arbitrary map of keys and values associated with the event data. | [optional] 
**Statistics** | Pointer to **map[string]interface{}** | Arbitrary map of statistics associated with the event data. | [optional] 

## Methods

### NewEventMetadataResultItem

`func NewEventMetadataResultItem() *EventMetadataResultItem`

NewEventMetadataResultItem instantiates a new EventMetadataResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMetadataResultItemWithDefaults

`func NewEventMetadataResultItemWithDefaults() *EventMetadataResultItem`

NewEventMetadataResultItemWithDefaults instantiates a new EventMetadataResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *EventMetadataResultItem) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *EventMetadataResultItem) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *EventMetadataResultItem) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *EventMetadataResultItem) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetStatistics

`func (o *EventMetadataResultItem) GetStatistics() map[string]interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *EventMetadataResultItem) GetStatisticsOk() (*map[string]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *EventMetadataResultItem) SetStatistics(v map[string]interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *EventMetadataResultItem) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


