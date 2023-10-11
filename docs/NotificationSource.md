# NotificationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly | 
**EventId** | **string** | UUID id of the source event in the Velo platform | 
**CreatedAt** | **time.Time** | ISO8601 timestamp indicating when the source event was created | 
**PaymentId** | **string** | ID of this payment within the Velo platform | 
**PayoutPayorIds** | Pointer to [**PayoutPayorIds**](PayoutPayorIds.md) |  | [optional] 
**PayorPaymentId** | Pointer to **string** | ID of this payment in the payors system | [optional] 
**Status** | **string** | The new status of the debit. One of \&quot;PENDING\&quot; \&quot;PROCESSING\&quot; \&quot;REJECTED\&quot; \&quot;RELEASED\&quot; | 
**ReasonCode** | **string** | The Velo code that indicates why the payment was rejected or returned | 
**ReasonMessage** | **string** | The description of why the payment was rejected or returned | 
**PayeeId** | **string** | ID of this payee within the Velo platform | 
**Reasons** | Pointer to [**[]PayeeEventAllOfReasons**](PayeeEventAllOfReasons.md) | The reasons for the event notification. | [optional] 
**DebitTransactionId** | **string** | ID of this debit transaction within the Velo platform | 

## Methods

### NewNotificationSource

`func NewNotificationSource(sourceType string, eventId string, createdAt time.Time, paymentId string, status string, reasonCode string, reasonMessage string, payeeId string, debitTransactionId string, ) *NotificationSource`

NewNotificationSource instantiates a new NotificationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSourceWithDefaults

`func NewNotificationSourceWithDefaults() *NotificationSource`

NewNotificationSourceWithDefaults instantiates a new NotificationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *NotificationSource) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *NotificationSource) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *NotificationSource) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetEventId

`func (o *NotificationSource) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *NotificationSource) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *NotificationSource) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCreatedAt

`func (o *NotificationSource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationSource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationSource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPaymentId

`func (o *NotificationSource) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *NotificationSource) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *NotificationSource) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayoutPayorIds

`func (o *NotificationSource) GetPayoutPayorIds() PayoutPayorIds`

GetPayoutPayorIds returns the PayoutPayorIds field if non-nil, zero value otherwise.

### GetPayoutPayorIdsOk

`func (o *NotificationSource) GetPayoutPayorIdsOk() (*PayoutPayorIds, bool)`

GetPayoutPayorIdsOk returns a tuple with the PayoutPayorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutPayorIds

`func (o *NotificationSource) SetPayoutPayorIds(v PayoutPayorIds)`

SetPayoutPayorIds sets PayoutPayorIds field to given value.

### HasPayoutPayorIds

`func (o *NotificationSource) HasPayoutPayorIds() bool`

HasPayoutPayorIds returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *NotificationSource) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *NotificationSource) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *NotificationSource) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *NotificationSource) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *NotificationSource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationSource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationSource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReasonCode

`func (o *NotificationSource) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *NotificationSource) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *NotificationSource) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetReasonMessage

`func (o *NotificationSource) GetReasonMessage() string`

GetReasonMessage returns the ReasonMessage field if non-nil, zero value otherwise.

### GetReasonMessageOk

`func (o *NotificationSource) GetReasonMessageOk() (*string, bool)`

GetReasonMessageOk returns a tuple with the ReasonMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonMessage

`func (o *NotificationSource) SetReasonMessage(v string)`

SetReasonMessage sets ReasonMessage field to given value.


### GetPayeeId

`func (o *NotificationSource) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *NotificationSource) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *NotificationSource) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetReasons

`func (o *NotificationSource) GetReasons() []PayeeEventAllOfReasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *NotificationSource) GetReasonsOk() (*[]PayeeEventAllOfReasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *NotificationSource) SetReasons(v []PayeeEventAllOfReasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *NotificationSource) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetDebitTransactionId

`func (o *NotificationSource) GetDebitTransactionId() string`

GetDebitTransactionId returns the DebitTransactionId field if non-nil, zero value otherwise.

### GetDebitTransactionIdOk

`func (o *NotificationSource) GetDebitTransactionIdOk() (*string, bool)`

GetDebitTransactionIdOk returns a tuple with the DebitTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitTransactionId

`func (o *NotificationSource) SetDebitTransactionId(v string)`

SetDebitTransactionId sets DebitTransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


