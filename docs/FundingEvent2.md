# FundingEvent2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** |  | [optional] 
**EventDateTime** | Pointer to **time.Time** |  | [optional] 
**FundingEventType** | Pointer to **string** | Funding event type. One of the following values: PAYOR_FUNDING_DETECTED, PAYOR_FUNDING_REQUESTED, PAYOR_FUNDING_RETURN_RECEIVED, FUNDING_RETURN_DETECTED, PAYOR_FUNDING_REQUEST_SUBMITTED, PAYOR_FUNDING_ENTRY_DETAIL_RECEIVED, FUNDING_DEALLOCATED | [optional] 
**Principal** | Pointer to **string** |  | [optional] 

## Methods

### NewFundingEvent2

`func NewFundingEvent2() *FundingEvent2`

NewFundingEvent2 instantiates a new FundingEvent2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingEvent2WithDefaults

`func NewFundingEvent2WithDefaults() *FundingEvent2`

NewFundingEvent2WithDefaults instantiates a new FundingEvent2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *FundingEvent2) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *FundingEvent2) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *FundingEvent2) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *FundingEvent2) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventDateTime

`func (o *FundingEvent2) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *FundingEvent2) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *FundingEvent2) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.

### HasEventDateTime

`func (o *FundingEvent2) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.

### GetFundingEventType

`func (o *FundingEvent2) GetFundingEventType() string`

GetFundingEventType returns the FundingEventType field if non-nil, zero value otherwise.

### GetFundingEventTypeOk

`func (o *FundingEvent2) GetFundingEventTypeOk() (*string, bool)`

GetFundingEventTypeOk returns a tuple with the FundingEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingEventType

`func (o *FundingEvent2) SetFundingEventType(v string)`

SetFundingEventType sets FundingEventType field to given value.

### HasFundingEventType

`func (o *FundingEvent2) HasFundingEventType() bool`

HasFundingEventType returns a boolean if a field has been set.

### GetPrincipal

`func (o *FundingEvent2) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *FundingEvent2) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *FundingEvent2) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *FundingEvent2) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

