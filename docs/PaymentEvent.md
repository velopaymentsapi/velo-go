# PaymentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly | 
**EventId** | **string** | UUID id of the source event in the Velo platform | 
**CreatedAt** | **time.Time** | ISO8601 timestamp indicating when the source event was created | 
**PaymentId** | **string** | ID of this payment within the Velo platform | 
**PayoutPayorIds** | Pointer to [**PayoutPayorIds**](PayoutPayorIds.md) |  | [optional] 
**PayorPaymentId** | Pointer to **string** | ID of this payment in the payors system | [optional] 

## Methods

### NewPaymentEvent

`func NewPaymentEvent(sourceType string, eventId string, createdAt time.Time, paymentId string, ) *PaymentEvent`

NewPaymentEvent instantiates a new PaymentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentEventWithDefaults

`func NewPaymentEventWithDefaults() *PaymentEvent`

NewPaymentEventWithDefaults instantiates a new PaymentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *PaymentEvent) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PaymentEvent) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PaymentEvent) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetEventId

`func (o *PaymentEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PaymentEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PaymentEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCreatedAt

`func (o *PaymentEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPaymentId

`func (o *PaymentEvent) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentEvent) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentEvent) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayoutPayorIds

`func (o *PaymentEvent) GetPayoutPayorIds() PayoutPayorIds`

GetPayoutPayorIds returns the PayoutPayorIds field if non-nil, zero value otherwise.

### GetPayoutPayorIdsOk

`func (o *PaymentEvent) GetPayoutPayorIdsOk() (*PayoutPayorIds, bool)`

GetPayoutPayorIdsOk returns a tuple with the PayoutPayorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutPayorIds

`func (o *PaymentEvent) SetPayoutPayorIds(v PayoutPayorIds)`

SetPayoutPayorIds sets PayoutPayorIds field to given value.

### HasPayoutPayorIds

`func (o *PaymentEvent) HasPayoutPayorIds() bool`

HasPayoutPayorIds returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *PaymentEvent) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentEvent) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentEvent) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentEvent) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


