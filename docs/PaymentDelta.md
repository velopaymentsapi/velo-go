# PaymentDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  | [readonly] 
**PayoutId** | **string** |  | [readonly] 
**PayorPaymentId** | Pointer to **NullableString** |  | [optional] 
**PaymentCurrency** | Pointer to **NullableString** |  | [optional] 
**PaymentAmount** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**SourceCurrency** | Pointer to **NullableString** |  | [optional] 
**SourceAmount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPaymentDelta

`func NewPaymentDelta(paymentId string, payoutId string, ) *PaymentDelta`

NewPaymentDelta instantiates a new PaymentDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDeltaWithDefaults

`func NewPaymentDeltaWithDefaults() *PaymentDelta`

NewPaymentDeltaWithDefaults instantiates a new PaymentDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentDelta) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentDelta) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentDelta) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayoutId

`func (o *PaymentDelta) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PaymentDelta) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PaymentDelta) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetPayorPaymentId

`func (o *PaymentDelta) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentDelta) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentDelta) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentDelta) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### SetPayorPaymentIdNil

`func (o *PaymentDelta) SetPayorPaymentIdNil(b bool)`

 SetPayorPaymentIdNil sets the value for PayorPaymentId to be an explicit nil

### UnsetPayorPaymentId
`func (o *PaymentDelta) UnsetPayorPaymentId()`

UnsetPayorPaymentId ensures that no value is present for PayorPaymentId, not even an explicit nil
### GetPaymentCurrency

`func (o *PaymentDelta) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PaymentDelta) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PaymentDelta) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PaymentDelta) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### SetPaymentCurrencyNil

`func (o *PaymentDelta) SetPaymentCurrencyNil(b bool)`

 SetPaymentCurrencyNil sets the value for PaymentCurrency to be an explicit nil

### UnsetPaymentCurrency
`func (o *PaymentDelta) UnsetPaymentCurrency()`

UnsetPaymentCurrency ensures that no value is present for PaymentCurrency, not even an explicit nil
### GetPaymentAmount

`func (o *PaymentDelta) GetPaymentAmount() int32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentDelta) GetPaymentAmountOk() (*int32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentDelta) SetPaymentAmount(v int32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PaymentDelta) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### SetPaymentAmountNil

`func (o *PaymentDelta) SetPaymentAmountNil(b bool)`

 SetPaymentAmountNil sets the value for PaymentAmount to be an explicit nil

### UnsetPaymentAmount
`func (o *PaymentDelta) UnsetPaymentAmount()`

UnsetPaymentAmount ensures that no value is present for PaymentAmount, not even an explicit nil
### GetStatus

`func (o *PaymentDelta) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentDelta) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentDelta) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentDelta) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PaymentDelta) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PaymentDelta) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSourceCurrency

`func (o *PaymentDelta) GetSourceCurrency() string`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *PaymentDelta) GetSourceCurrencyOk() (*string, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *PaymentDelta) SetSourceCurrency(v string)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *PaymentDelta) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### SetSourceCurrencyNil

`func (o *PaymentDelta) SetSourceCurrencyNil(b bool)`

 SetSourceCurrencyNil sets the value for SourceCurrency to be an explicit nil

### UnsetSourceCurrency
`func (o *PaymentDelta) UnsetSourceCurrency()`

UnsetSourceCurrency ensures that no value is present for SourceCurrency, not even an explicit nil
### GetSourceAmount

`func (o *PaymentDelta) GetSourceAmount() int32`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *PaymentDelta) GetSourceAmountOk() (*int32, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *PaymentDelta) SetSourceAmount(v int32)`

SetSourceAmount sets SourceAmount field to given value.

### HasSourceAmount

`func (o *PaymentDelta) HasSourceAmount() bool`

HasSourceAmount returns a boolean if a field has been set.

### SetSourceAmountNil

`func (o *PaymentDelta) SetSourceAmountNil(b bool)`

 SetSourceAmountNil sets the value for SourceAmount to be an explicit nil

### UnsetSourceAmount
`func (o *PaymentDelta) UnsetSourceAmount()`

UnsetSourceAmount ensures that no value is present for SourceAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


