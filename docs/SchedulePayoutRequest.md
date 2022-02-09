# SchedulePayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledFor** | **time.Time** | UTC timestamp for instructing the payout. Format is ISO-8601. | 
**NotificationsEnabled** | **bool** | Flag to indicate whether to receive notifications when scheduled payout is processed | 

## Methods

### NewSchedulePayoutRequest

`func NewSchedulePayoutRequest(scheduledFor time.Time, notificationsEnabled bool, ) *SchedulePayoutRequest`

NewSchedulePayoutRequest instantiates a new SchedulePayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePayoutRequestWithDefaults

`func NewSchedulePayoutRequestWithDefaults() *SchedulePayoutRequest`

NewSchedulePayoutRequestWithDefaults instantiates a new SchedulePayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledFor

`func (o *SchedulePayoutRequest) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *SchedulePayoutRequest) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *SchedulePayoutRequest) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.


### GetNotificationsEnabled

`func (o *SchedulePayoutRequest) GetNotificationsEnabled() bool`

GetNotificationsEnabled returns the NotificationsEnabled field if non-nil, zero value otherwise.

### GetNotificationsEnabledOk

`func (o *SchedulePayoutRequest) GetNotificationsEnabledOk() (*bool, bool)`

GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEnabled

`func (o *SchedulePayoutRequest) SetNotificationsEnabled(v bool)`

SetNotificationsEnabled sets NotificationsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


