# PaymentDeltaV1

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

### NewPaymentDeltaV1

`func NewPaymentDeltaV1(paymentId string, payoutId string, ) *PaymentDeltaV1`

NewPaymentDeltaV1 instantiates a new PaymentDeltaV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDeltaV1WithDefaults

`func NewPaymentDeltaV1WithDefaults() *PaymentDeltaV1`

NewPaymentDeltaV1WithDefaults instantiates a new PaymentDeltaV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentDeltaV1) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentDeltaV1) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentDeltaV1) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayoutId

`func (o *PaymentDeltaV1) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PaymentDeltaV1) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PaymentDeltaV1) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetPayorPaymentId

`func (o *PaymentDeltaV1) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentDeltaV1) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentDeltaV1) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentDeltaV1) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### SetPayorPaymentIdNil

`func (o *PaymentDeltaV1) SetPayorPaymentIdNil(b bool)`

 SetPayorPaymentIdNil sets the value for PayorPaymentId to be an explicit nil

### UnsetPayorPaymentId
`func (o *PaymentDeltaV1) UnsetPayorPaymentId()`

UnsetPayorPaymentId ensures that no value is present for PayorPaymentId, not even an explicit nil
### GetPaymentCurrency

`func (o *PaymentDeltaV1) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PaymentDeltaV1) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PaymentDeltaV1) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PaymentDeltaV1) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### SetPaymentCurrencyNil

`func (o *PaymentDeltaV1) SetPaymentCurrencyNil(b bool)`

 SetPaymentCurrencyNil sets the value for PaymentCurrency to be an explicit nil

### UnsetPaymentCurrency
`func (o *PaymentDeltaV1) UnsetPaymentCurrency()`

UnsetPaymentCurrency ensures that no value is present for PaymentCurrency, not even an explicit nil
### GetPaymentAmount

`func (o *PaymentDeltaV1) GetPaymentAmount() int32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentDeltaV1) GetPaymentAmountOk() (*int32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentDeltaV1) SetPaymentAmount(v int32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PaymentDeltaV1) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### SetPaymentAmountNil

`func (o *PaymentDeltaV1) SetPaymentAmountNil(b bool)`

 SetPaymentAmountNil sets the value for PaymentAmount to be an explicit nil

### UnsetPaymentAmount
`func (o *PaymentDeltaV1) UnsetPaymentAmount()`

UnsetPaymentAmount ensures that no value is present for PaymentAmount, not even an explicit nil
### GetStatus

`func (o *PaymentDeltaV1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentDeltaV1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentDeltaV1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentDeltaV1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PaymentDeltaV1) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PaymentDeltaV1) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSourceCurrency

`func (o *PaymentDeltaV1) GetSourceCurrency() string`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *PaymentDeltaV1) GetSourceCurrencyOk() (*string, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *PaymentDeltaV1) SetSourceCurrency(v string)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *PaymentDeltaV1) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### SetSourceCurrencyNil

`func (o *PaymentDeltaV1) SetSourceCurrencyNil(b bool)`

 SetSourceCurrencyNil sets the value for SourceCurrency to be an explicit nil

### UnsetSourceCurrency
`func (o *PaymentDeltaV1) UnsetSourceCurrency()`

UnsetSourceCurrency ensures that no value is present for SourceCurrency, not even an explicit nil
### GetSourceAmount

`func (o *PaymentDeltaV1) GetSourceAmount() int32`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *PaymentDeltaV1) GetSourceAmountOk() (*int32, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *PaymentDeltaV1) SetSourceAmount(v int32)`

SetSourceAmount sets SourceAmount field to given value.

### HasSourceAmount

`func (o *PaymentDeltaV1) HasSourceAmount() bool`

HasSourceAmount returns a boolean if a field has been set.

### SetSourceAmountNil

`func (o *PaymentDeltaV1) SetSourceAmountNil(b bool)`

 SetSourceAmountNil sets the value for SourceAmount to be an explicit nil

### UnsetSourceAmount
`func (o *PaymentDeltaV1) UnsetSourceAmount()`

UnsetSourceAmount ensures that no value is present for SourceAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


