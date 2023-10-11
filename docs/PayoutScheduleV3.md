# PayoutScheduleV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleStatus** | **string** | The current status of the payout schedule. One of the following values: SCHEDULED, EXECUTED, FAILED | 
**ScheduledAt** | **time.Time** |  | 
**ScheduledFor** | **time.Time** |  | 
**ScheduledByPrincipalId** | **string** | ID of the user or application that scheduled the payout | 
**NotificationsEnabled** | **bool** |  | 

## Methods

### NewPayoutScheduleV3

`func NewPayoutScheduleV3(scheduleStatus string, scheduledAt time.Time, scheduledFor time.Time, scheduledByPrincipalId string, notificationsEnabled bool, ) *PayoutScheduleV3`

NewPayoutScheduleV3 instantiates a new PayoutScheduleV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutScheduleV3WithDefaults

`func NewPayoutScheduleV3WithDefaults() *PayoutScheduleV3`

NewPayoutScheduleV3WithDefaults instantiates a new PayoutScheduleV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleStatus

`func (o *PayoutScheduleV3) GetScheduleStatus() string`

GetScheduleStatus returns the ScheduleStatus field if non-nil, zero value otherwise.

### GetScheduleStatusOk

`func (o *PayoutScheduleV3) GetScheduleStatusOk() (*string, bool)`

GetScheduleStatusOk returns a tuple with the ScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStatus

`func (o *PayoutScheduleV3) SetScheduleStatus(v string)`

SetScheduleStatus sets ScheduleStatus field to given value.


### GetScheduledAt

`func (o *PayoutScheduleV3) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *PayoutScheduleV3) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *PayoutScheduleV3) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.


### GetScheduledFor

`func (o *PayoutScheduleV3) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *PayoutScheduleV3) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *PayoutScheduleV3) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.


### GetScheduledByPrincipalId

`func (o *PayoutScheduleV3) GetScheduledByPrincipalId() string`

GetScheduledByPrincipalId returns the ScheduledByPrincipalId field if non-nil, zero value otherwise.

### GetScheduledByPrincipalIdOk

`func (o *PayoutScheduleV3) GetScheduledByPrincipalIdOk() (*string, bool)`

GetScheduledByPrincipalIdOk returns a tuple with the ScheduledByPrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledByPrincipalId

`func (o *PayoutScheduleV3) SetScheduledByPrincipalId(v string)`

SetScheduledByPrincipalId sets ScheduledByPrincipalId field to given value.


### GetNotificationsEnabled

`func (o *PayoutScheduleV3) GetNotificationsEnabled() bool`

GetNotificationsEnabled returns the NotificationsEnabled field if non-nil, zero value otherwise.

### GetNotificationsEnabledOk

`func (o *PayoutScheduleV3) GetNotificationsEnabledOk() (*bool, bool)`

GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEnabled

`func (o *PayoutScheduleV3) SetNotificationsEnabled(v bool)`

SetNotificationsEnabled sets NotificationsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


