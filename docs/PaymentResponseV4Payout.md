# PaymentResponseV4Payout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutId** | Pointer to **string** |  | [optional] 
**PayoutFrom** | Pointer to [**PayoutPayorV4**](PayoutPayorV4.md) |  | [optional] 
**PayoutTo** | Pointer to [**PayoutPayorV4**](PayoutPayorV4.md) |  | [optional] 

## Methods

### NewPaymentResponseV4Payout

`func NewPaymentResponseV4Payout() *PaymentResponseV4Payout`

NewPaymentResponseV4Payout instantiates a new PaymentResponseV4Payout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseV4PayoutWithDefaults

`func NewPaymentResponseV4PayoutWithDefaults() *PaymentResponseV4Payout`

NewPaymentResponseV4PayoutWithDefaults instantiates a new PaymentResponseV4Payout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutId

`func (o *PaymentResponseV4Payout) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PaymentResponseV4Payout) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PaymentResponseV4Payout) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.

### HasPayoutId

`func (o *PaymentResponseV4Payout) HasPayoutId() bool`

HasPayoutId returns a boolean if a field has been set.

### GetPayoutFrom

`func (o *PaymentResponseV4Payout) GetPayoutFrom() PayoutPayorV4`

GetPayoutFrom returns the PayoutFrom field if non-nil, zero value otherwise.

### GetPayoutFromOk

`func (o *PaymentResponseV4Payout) GetPayoutFromOk() (*PayoutPayorV4, bool)`

GetPayoutFromOk returns a tuple with the PayoutFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutFrom

`func (o *PaymentResponseV4Payout) SetPayoutFrom(v PayoutPayorV4)`

SetPayoutFrom sets PayoutFrom field to given value.

### HasPayoutFrom

`func (o *PaymentResponseV4Payout) HasPayoutFrom() bool`

HasPayoutFrom returns a boolean if a field has been set.

### GetPayoutTo

`func (o *PaymentResponseV4Payout) GetPayoutTo() PayoutPayorV4`

GetPayoutTo returns the PayoutTo field if non-nil, zero value otherwise.

### GetPayoutToOk

`func (o *PaymentResponseV4Payout) GetPayoutToOk() (*PayoutPayorV4, bool)`

GetPayoutToOk returns a tuple with the PayoutTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutTo

`func (o *PaymentResponseV4Payout) SetPayoutTo(v PayoutPayorV4)`

SetPayoutTo sets PayoutTo field to given value.

### HasPayoutTo

`func (o *PaymentResponseV4Payout) HasPayoutTo() bool`

HasPayoutTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


