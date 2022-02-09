# PayoutSchedule2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleStatus** | [**ScheduleStatus2**](ScheduleStatus2.md) |  | 
**ScheduledAt** | **time.Time** |  | 
**ScheduledFor** | **time.Time** |  | 
**ScheduledByPrincipalId** | **string** | ID of the user or application that scheduled the payout | 
**NotificationsEnabled** | **bool** |  | 

## Methods

### NewPayoutSchedule2

`func NewPayoutSchedule2(scheduleStatus ScheduleStatus2, scheduledAt time.Time, scheduledFor time.Time, scheduledByPrincipalId string, notificationsEnabled bool, ) *PayoutSchedule2`

NewPayoutSchedule2 instantiates a new PayoutSchedule2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutSchedule2WithDefaults

`func NewPayoutSchedule2WithDefaults() *PayoutSchedule2`

NewPayoutSchedule2WithDefaults instantiates a new PayoutSchedule2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleStatus

`func (o *PayoutSchedule2) GetScheduleStatus() ScheduleStatus2`

GetScheduleStatus returns the ScheduleStatus field if non-nil, zero value otherwise.

### GetScheduleStatusOk

`func (o *PayoutSchedule2) GetScheduleStatusOk() (*ScheduleStatus2, bool)`

GetScheduleStatusOk returns a tuple with the ScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStatus

`func (o *PayoutSchedule2) SetScheduleStatus(v ScheduleStatus2)`

SetScheduleStatus sets ScheduleStatus field to given value.


### GetScheduledAt

`func (o *PayoutSchedule2) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *PayoutSchedule2) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *PayoutSchedule2) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.


### GetScheduledFor

`func (o *PayoutSchedule2) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *PayoutSchedule2) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *PayoutSchedule2) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.


### GetScheduledByPrincipalId

`func (o *PayoutSchedule2) GetScheduledByPrincipalId() string`

GetScheduledByPrincipalId returns the ScheduledByPrincipalId field if non-nil, zero value otherwise.

### GetScheduledByPrincipalIdOk

`func (o *PayoutSchedule2) GetScheduledByPrincipalIdOk() (*string, bool)`

GetScheduledByPrincipalIdOk returns a tuple with the ScheduledByPrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledByPrincipalId

`func (o *PayoutSchedule2) SetScheduledByPrincipalId(v string)`

SetScheduledByPrincipalId sets ScheduledByPrincipalId field to given value.


### GetNotificationsEnabled

`func (o *PayoutSchedule2) GetNotificationsEnabled() bool`

GetNotificationsEnabled returns the NotificationsEnabled field if non-nil, zero value otherwise.

### GetNotificationsEnabledOk

`func (o *PayoutSchedule2) GetNotificationsEnabledOk() (*bool, bool)`

GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEnabled

`func (o *PayoutSchedule2) SetNotificationsEnabled(v bool)`

SetNotificationsEnabled sets NotificationsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


