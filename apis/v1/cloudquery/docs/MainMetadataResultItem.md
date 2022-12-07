# MainMetadataResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Since** | Pointer to **time.Time** | Specifies the start of the time range which was used for querying. | [optional] 
**Until** | Pointer to **time.Time** | Specifies the end of the time range which was used for querying. | [optional] 

## Methods

### NewMainMetadataResultItem

`func NewMainMetadataResultItem() *MainMetadataResultItem`

NewMainMetadataResultItem instantiates a new MainMetadataResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainMetadataResultItemWithDefaults

`func NewMainMetadataResultItemWithDefaults() *MainMetadataResultItem`

NewMainMetadataResultItemWithDefaults instantiates a new MainMetadataResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSince

`func (o *MainMetadataResultItem) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *MainMetadataResultItem) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *MainMetadataResultItem) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *MainMetadataResultItem) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetUntil

`func (o *MainMetadataResultItem) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *MainMetadataResultItem) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *MainMetadataResultItem) SetUntil(v time.Time)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *MainMetadataResultItem) HasUntil() bool`

HasUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


