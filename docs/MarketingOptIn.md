# MarketingOptIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptedIn** | Pointer to **bool** |  | [optional] 
**Timestamp** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewMarketingOptIn

`func NewMarketingOptIn() *MarketingOptIn`

NewMarketingOptIn instantiates a new MarketingOptIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingOptInWithDefaults

`func NewMarketingOptInWithDefaults() *MarketingOptIn`

NewMarketingOptInWithDefaults instantiates a new MarketingOptIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptedIn

`func (o *MarketingOptIn) GetOptedIn() bool`

GetOptedIn returns the OptedIn field if non-nil, zero value otherwise.

### GetOptedInOk

`func (o *MarketingOptIn) GetOptedInOk() (*bool, bool)`

GetOptedInOk returns a tuple with the OptedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedIn

`func (o *MarketingOptIn) SetOptedIn(v bool)`

SetOptedIn sets OptedIn field to given value.

### HasOptedIn

`func (o *MarketingOptIn) HasOptedIn() bool`

HasOptedIn returns a boolean if a field has been set.

### GetTimestamp

`func (o *MarketingOptIn) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MarketingOptIn) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MarketingOptIn) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MarketingOptIn) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


