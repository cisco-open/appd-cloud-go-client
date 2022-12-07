# MetadataResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Since** | Pointer to **time.Time** | Specifies the start of the time range which was used for querying. | [optional] 
**Until** | Pointer to **time.Time** | Specifies the end of the time range which was used for querying. | [optional] 
**GranularitySeconds** | Pointer to **float32** | Actual granularity of metric data points. The value is the length of the time interval of a single data point. | [optional] 
**Schema** | Pointer to **map[string]interface{}** | Arbitrary map of keys and values associated with the event data. | [optional] 
**Statistics** | Pointer to **map[string]interface{}** | Arbitrary map of statistics associated with the event data. | [optional] 

## Methods

### NewMetadataResultItem

`func NewMetadataResultItem() *MetadataResultItem`

NewMetadataResultItem instantiates a new MetadataResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataResultItemWithDefaults

`func NewMetadataResultItemWithDefaults() *MetadataResultItem`

NewMetadataResultItemWithDefaults instantiates a new MetadataResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSince

`func (o *MetadataResultItem) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *MetadataResultItem) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *MetadataResultItem) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *MetadataResultItem) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetUntil

`func (o *MetadataResultItem) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *MetadataResultItem) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *MetadataResultItem) SetUntil(v time.Time)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *MetadataResultItem) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetGranularitySeconds

`func (o *MetadataResultItem) GetGranularitySeconds() float32`

GetGranularitySeconds returns the GranularitySeconds field if non-nil, zero value otherwise.

### GetGranularitySecondsOk

`func (o *MetadataResultItem) GetGranularitySecondsOk() (*float32, bool)`

GetGranularitySecondsOk returns a tuple with the GranularitySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularitySeconds

`func (o *MetadataResultItem) SetGranularitySeconds(v float32)`

SetGranularitySeconds sets GranularitySeconds field to given value.

### HasGranularitySeconds

`func (o *MetadataResultItem) HasGranularitySeconds() bool`

HasGranularitySeconds returns a boolean if a field has been set.

### GetSchema

`func (o *MetadataResultItem) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *MetadataResultItem) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *MetadataResultItem) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *MetadataResultItem) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetStatistics

`func (o *MetadataResultItem) GetStatistics() map[string]interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *MetadataResultItem) GetStatisticsOk() (*map[string]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *MetadataResultItem) SetStatistics(v map[string]interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *MetadataResultItem) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


