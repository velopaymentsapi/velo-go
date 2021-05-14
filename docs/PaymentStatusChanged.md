# PaymentStatusChanged

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

## Methods

### NewPaymentStatusChanged

`func NewPaymentStatusChanged(sourceType string, eventId string, createdAt time.Time, paymentId string, status string, ) *PaymentStatusChanged`

NewPaymentStatusChanged instantiates a new PaymentStatusChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentStatusChangedWithDefaults

`func NewPaymentStatusChangedWithDefaults() *PaymentStatusChanged`

NewPaymentStatusChangedWithDefaults instantiates a new PaymentStatusChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *PaymentStatusChanged) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PaymentStatusChanged) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PaymentStatusChanged) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetEventId

`func (o *PaymentStatusChanged) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PaymentStatusChanged) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PaymentStatusChanged) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCreatedAt

`func (o *PaymentStatusChanged) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentStatusChanged) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentStatusChanged) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPaymentId

`func (o *PaymentStatusChanged) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentStatusChanged) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentStatusChanged) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayoutPayorIds

`func (o *PaymentStatusChanged) GetPayoutPayorIds() PayoutPayorIds`

GetPayoutPayorIds returns the PayoutPayorIds field if non-nil, zero value otherwise.

### GetPayoutPayorIdsOk

`func (o *PaymentStatusChanged) GetPayoutPayorIdsOk() (*PayoutPayorIds, bool)`

GetPayoutPayorIdsOk returns a tuple with the PayoutPayorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutPayorIds

`func (o *PaymentStatusChanged) SetPayoutPayorIds(v PayoutPayorIds)`

SetPayoutPayorIds sets PayoutPayorIds field to given value.

### HasPayoutPayorIds

`func (o *PaymentStatusChanged) HasPayoutPayorIds() bool`

HasPayoutPayorIds returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *PaymentStatusChanged) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentStatusChanged) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentStatusChanged) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentStatusChanged) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentStatusChanged) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentStatusChanged) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentStatusChanged) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


