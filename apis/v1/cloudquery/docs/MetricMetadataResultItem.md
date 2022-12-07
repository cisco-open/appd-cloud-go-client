# MetricMetadataResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GranularitySeconds** | Pointer to **float32** | Actual granularity of metric data points. The value is the length of the time interval of a single data point. | [optional] 

## Methods

### NewMetricMetadataResultItem

`func NewMetricMetadataResultItem() *MetricMetadataResultItem`

NewMetricMetadataResultItem instantiates a new MetricMetadataResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricMetadataResultItemWithDefaults

`func NewMetricMetadataResultItemWithDefaults() *MetricMetadataResultItem`

NewMetricMetadataResultItemWithDefaults instantiates a new MetricMetadataResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranularitySeconds

`func (o *MetricMetadataResultItem) GetGranularitySeconds() float32`

GetGranularitySeconds returns the GranularitySeconds field if non-nil, zero value otherwise.

### GetGranularitySecondsOk

`func (o *MetricMetadataResultItem) GetGranularitySecondsOk() (*float32, bool)`

GetGranularitySecondsOk returns a tuple with the GranularitySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularitySeconds

`func (o *MetricMetadataResultItem) SetGranularitySeconds(v float32)`

SetGranularitySeconds sets GranularitySeconds field to given value.

### HasGranularitySeconds

`func (o *MetricMetadataResultItem) HasGranularitySeconds() bool`

HasGranularitySeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


