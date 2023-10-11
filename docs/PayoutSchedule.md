# PayoutSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleStatus** | **string** | Current status of the payout schedule. One of the following values: SCHEDULED, EXECUTED, FAILED | 
**ScheduledAt** | **time.Time** |  | 
**ScheduledFor** | **time.Time** |  | 
**ScheduledByPrincipalId** | **string** | ID of the user or application that scheduled the payout | 
**NotificationsEnabled** | **bool** |  | 
**ScheduledBy** | Pointer to **string** | Optional display name as a hint for who scheduled the payout. Not populated if payout was scheduled by an application. | [optional] 

## Methods

### NewPayoutSchedule

`func NewPayoutSchedule(scheduleStatus string, scheduledAt time.Time, scheduledFor time.Time, scheduledByPrincipalId string, notificationsEnabled bool, ) *PayoutSchedule`

NewPayoutSchedule instantiates a new PayoutSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutScheduleWithDefaults

`func NewPayoutScheduleWithDefaults() *PayoutSchedule`

NewPayoutScheduleWithDefaults instantiates a new PayoutSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleStatus

`func (o *PayoutSchedule) GetScheduleStatus() string`

GetScheduleStatus returns the ScheduleStatus field if non-nil, zero value otherwise.

### GetScheduleStatusOk

`func (o *PayoutSchedule) GetScheduleStatusOk() (*string, bool)`

GetScheduleStatusOk returns a tuple with the ScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStatus

`func (o *PayoutSchedule) SetScheduleStatus(v string)`

SetScheduleStatus sets ScheduleStatus field to given value.


### GetScheduledAt

`func (o *PayoutSchedule) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *PayoutSchedule) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *PayoutSchedule) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.


### GetScheduledFor

`func (o *PayoutSchedule) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *PayoutSchedule) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *PayoutSchedule) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.


### GetScheduledByPrincipalId

`func (o *PayoutSchedule) GetScheduledByPrincipalId() string`

GetScheduledByPrincipalId returns the ScheduledByPrincipalId field if non-nil, zero value otherwise.

### GetScheduledByPrincipalIdOk

`func (o *PayoutSchedule) GetScheduledByPrincipalIdOk() (*string, bool)`

GetScheduledByPrincipalIdOk returns a tuple with the ScheduledByPrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledByPrincipalId

`func (o *PayoutSchedule) SetScheduledByPrincipalId(v string)`

SetScheduledByPrincipalId sets ScheduledByPrincipalId field to given value.


### GetNotificationsEnabled

`func (o *PayoutSchedule) GetNotificationsEnabled() bool`

GetNotificationsEnabled returns the NotificationsEnabled field if non-nil, zero value otherwise.

### GetNotificationsEnabledOk

`func (o *PayoutSchedule) GetNotificationsEnabledOk() (*bool, bool)`

GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEnabled

`func (o *PayoutSchedule) SetNotificationsEnabled(v bool)`

SetNotificationsEnabled sets NotificationsEnabled field to given value.


### GetScheduledBy

`func (o *PayoutSchedule) GetScheduledBy() string`

GetScheduledBy returns the ScheduledBy field if non-nil, zero value otherwise.

### GetScheduledByOk

`func (o *PayoutSchedule) GetScheduledByOk() (*string, bool)`

GetScheduledByOk returns a tuple with the ScheduledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledBy

`func (o *PayoutSchedule) SetScheduledBy(v string)`

SetScheduledBy sets ScheduledBy field to given value.

### HasScheduledBy

`func (o *PayoutSchedule) HasScheduledBy() bool`

HasScheduledBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


