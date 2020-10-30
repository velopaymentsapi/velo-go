# FundingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** |  | [optional] 
**EventDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**FundingEventType** | Pointer to [**FundingEventType**](FundingEventType.md) |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 

## Methods

### NewFundingEvent

`func NewFundingEvent() *FundingEvent`

NewFundingEvent instantiates a new FundingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingEventWithDefaults

`func NewFundingEventWithDefaults() *FundingEvent`

NewFundingEventWithDefaults instantiates a new FundingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *FundingEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *FundingEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *FundingEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *FundingEvent) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventDateTime

`func (o *FundingEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *FundingEvent) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *FundingEvent) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.

### HasEventDateTime

`func (o *FundingEvent) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.

### GetFundingEventType

`func (o *FundingEvent) GetFundingEventType() FundingEventType`

GetFundingEventType returns the FundingEventType field if non-nil, zero value otherwise.

### GetFundingEventTypeOk

`func (o *FundingEvent) GetFundingEventTypeOk() (*FundingEventType, bool)`

GetFundingEventTypeOk returns a tuple with the FundingEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingEventType

`func (o *FundingEvent) SetFundingEventType(v FundingEventType)`

SetFundingEventType sets FundingEventType field to given value.

### HasFundingEventType

`func (o *FundingEvent) HasFundingEventType() bool`

HasFundingEventType returns a boolean if a field has been set.

### GetPrincipal

`func (o *FundingEvent) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *FundingEvent) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *FundingEvent) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *FundingEvent) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


