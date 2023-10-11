# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | The API version of the notification schema | 
**SequenceNumber** | **int64** | This is a payor specific sequence number starting at 1 for the first notification sent | 
**Category** | **string** | The category that the notification relates to. One of \&quot;payment\&quot;, \&quot;payee\&quot;, \&quot;debit\&quot; or \&quot;system\&quot; | 
**EventName** | **string** | The name of event that led to this notification | 
**Source** | Pointer to [**NotificationSource**](NotificationSource.md) |  | [optional] 

## Methods

### NewNotification

`func NewNotification(apiVersion string, sequenceNumber int64, category string, eventName string, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Notification) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Notification) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Notification) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetSequenceNumber

`func (o *Notification) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Notification) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Notification) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetCategory

`func (o *Notification) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Notification) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Notification) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetEventName

`func (o *Notification) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *Notification) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *Notification) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetSource

`func (o *Notification) GetSource() NotificationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Notification) GetSourceOk() (*NotificationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Notification) SetSource(v NotificationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Notification) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


