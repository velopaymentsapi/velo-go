# PaymentRejectedOrReturned

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly | 
**EventId** | **string** | UUID id of the source event in the Velo platform | 
**CreatedAt** | **time.Time** | ISO8601 timestamp indicating when the source event was created | 
**PaymentId** | **string** | ID of this payment within the Velo platform | 
**PayoutPayorIds** | Pointer to [**PayoutPayorIds**](PayoutPayorIds.md) |  | [optional] 
**PayorPaymentId** | Pointer to **string** | ID of this payment in the payors system | [optional] 
**Status** | **string** | The new status of the payment. One of \&quot;SUBMITTED\&quot; \&quot;ACCEPTED\&quot; \&quot;REJECTED\&quot; \&quot;ACCEPTED_BY_RAILS\&quot; \&quot;CONFIRMED\&quot; \&quot;RETURNED\&quot; \&quot;WITHDRAWN\&quot; | 
**ReasonCode** | **string** | The Velo code that indicates why the payment was rejected or returned | 
**ReasonMessage** | **string** | The description of why the payment was rejected or returned | 

## Methods

### NewPaymentRejectedOrReturned

`func NewPaymentRejectedOrReturned(sourceType string, eventId string, createdAt time.Time, paymentId string, status string, reasonCode string, reasonMessage string, ) *PaymentRejectedOrReturned`

NewPaymentRejectedOrReturned instantiates a new PaymentRejectedOrReturned object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRejectedOrReturnedWithDefaults

`func NewPaymentRejectedOrReturnedWithDefaults() *PaymentRejectedOrReturned`

NewPaymentRejectedOrReturnedWithDefaults instantiates a new PaymentRejectedOrReturned object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *PaymentRejectedOrReturned) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PaymentRejectedOrReturned) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PaymentRejectedOrReturned) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetEventId

`func (o *PaymentRejectedOrReturned) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PaymentRejectedOrReturned) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PaymentRejectedOrReturned) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCreatedAt

`func (o *PaymentRejectedOrReturned) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentRejectedOrReturned) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentRejectedOrReturned) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPaymentId

`func (o *PaymentRejectedOrReturned) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentRejectedOrReturned) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentRejectedOrReturned) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayoutPayorIds

`func (o *PaymentRejectedOrReturned) GetPayoutPayorIds() PayoutPayorIds`

GetPayoutPayorIds returns the PayoutPayorIds field if non-nil, zero value otherwise.

### GetPayoutPayorIdsOk

`func (o *PaymentRejectedOrReturned) GetPayoutPayorIdsOk() (*PayoutPayorIds, bool)`

GetPayoutPayorIdsOk returns a tuple with the PayoutPayorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutPayorIds

`func (o *PaymentRejectedOrReturned) SetPayoutPayorIds(v PayoutPayorIds)`

SetPayoutPayorIds sets PayoutPayorIds field to given value.

### HasPayoutPayorIds

`func (o *PaymentRejectedOrReturned) HasPayoutPayorIds() bool`

HasPayoutPayorIds returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *PaymentRejectedOrReturned) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentRejectedOrReturned) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentRejectedOrReturned) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentRejectedOrReturned) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentRejectedOrReturned) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRejectedOrReturned) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRejectedOrReturned) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReasonCode

`func (o *PaymentRejectedOrReturned) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *PaymentRejectedOrReturned) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *PaymentRejectedOrReturned) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetReasonMessage

`func (o *PaymentRejectedOrReturned) GetReasonMessage() string`

GetReasonMessage returns the ReasonMessage field if non-nil, zero value otherwise.

### GetReasonMessageOk

`func (o *PaymentRejectedOrReturned) GetReasonMessageOk() (*string, bool)`

GetReasonMessageOk returns a tuple with the ReasonMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonMessage

`func (o *PaymentRejectedOrReturned) SetReasonMessage(v string)`

SetReasonMessage sets ReasonMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


