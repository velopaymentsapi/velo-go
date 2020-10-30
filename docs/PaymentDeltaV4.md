# PaymentDeltaV4

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

### NewPaymentDeltaV4

`func NewPaymentDeltaV4(paymentId string, payoutId string, ) *PaymentDeltaV4`

NewPaymentDeltaV4 instantiates a new PaymentDeltaV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDeltaV4WithDefaults

`func NewPaymentDeltaV4WithDefaults() *PaymentDeltaV4`

NewPaymentDeltaV4WithDefaults instantiates a new PaymentDeltaV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentDeltaV4) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentDeltaV4) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentDeltaV4) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayoutId

`func (o *PaymentDeltaV4) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PaymentDeltaV4) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PaymentDeltaV4) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetPayorPaymentId

`func (o *PaymentDeltaV4) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentDeltaV4) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentDeltaV4) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentDeltaV4) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### SetPayorPaymentIdNil

`func (o *PaymentDeltaV4) SetPayorPaymentIdNil(b bool)`

 SetPayorPaymentIdNil sets the value for PayorPaymentId to be an explicit nil

### UnsetPayorPaymentId
`func (o *PaymentDeltaV4) UnsetPayorPaymentId()`

UnsetPayorPaymentId ensures that no value is present for PayorPaymentId, not even an explicit nil
### GetPaymentCurrency

`func (o *PaymentDeltaV4) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PaymentDeltaV4) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PaymentDeltaV4) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PaymentDeltaV4) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### SetPaymentCurrencyNil

`func (o *PaymentDeltaV4) SetPaymentCurrencyNil(b bool)`

 SetPaymentCurrencyNil sets the value for PaymentCurrency to be an explicit nil

### UnsetPaymentCurrency
`func (o *PaymentDeltaV4) UnsetPaymentCurrency()`

UnsetPaymentCurrency ensures that no value is present for PaymentCurrency, not even an explicit nil
### GetPaymentAmount

`func (o *PaymentDeltaV4) GetPaymentAmount() int32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentDeltaV4) GetPaymentAmountOk() (*int32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentDeltaV4) SetPaymentAmount(v int32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PaymentDeltaV4) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### SetPaymentAmountNil

`func (o *PaymentDeltaV4) SetPaymentAmountNil(b bool)`

 SetPaymentAmountNil sets the value for PaymentAmount to be an explicit nil

### UnsetPaymentAmount
`func (o *PaymentDeltaV4) UnsetPaymentAmount()`

UnsetPaymentAmount ensures that no value is present for PaymentAmount, not even an explicit nil
### GetStatus

`func (o *PaymentDeltaV4) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentDeltaV4) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentDeltaV4) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentDeltaV4) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PaymentDeltaV4) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PaymentDeltaV4) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSourceCurrency

`func (o *PaymentDeltaV4) GetSourceCurrency() string`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *PaymentDeltaV4) GetSourceCurrencyOk() (*string, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *PaymentDeltaV4) SetSourceCurrency(v string)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *PaymentDeltaV4) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### SetSourceCurrencyNil

`func (o *PaymentDeltaV4) SetSourceCurrencyNil(b bool)`

 SetSourceCurrencyNil sets the value for SourceCurrency to be an explicit nil

### UnsetSourceCurrency
`func (o *PaymentDeltaV4) UnsetSourceCurrency()`

UnsetSourceCurrency ensures that no value is present for SourceCurrency, not even an explicit nil
### GetSourceAmount

`func (o *PaymentDeltaV4) GetSourceAmount() int32`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *PaymentDeltaV4) GetSourceAmountOk() (*int32, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *PaymentDeltaV4) SetSourceAmount(v int32)`

SetSourceAmount sets SourceAmount field to given value.

### HasSourceAmount

`func (o *PaymentDeltaV4) HasSourceAmount() bool`

HasSourceAmount returns a boolean if a field has been set.

### SetSourceAmountNil

`func (o *PaymentDeltaV4) SetSourceAmountNil(b bool)`

 SetSourceAmountNil sets the value for SourceAmount to be an explicit nil

### UnsetSourceAmount
`func (o *PaymentDeltaV4) UnsetSourceAmount()`

UnsetSourceAmount ensures that no value is present for SourceAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


