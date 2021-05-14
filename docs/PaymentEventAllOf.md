# PaymentEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | ID of this payment within the Velo platform | 
**PayoutPayorIds** | Pointer to [**PayoutPayorIds**](PayoutPayorIds.md) |  | [optional] 
**PayorPaymentId** | Pointer to **string** | ID of this payment in the payors system | [optional] 

## Methods

### NewPaymentEventAllOf

`func NewPaymentEventAllOf(paymentId string, ) *PaymentEventAllOf`

NewPaymentEventAllOf instantiates a new PaymentEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentEventAllOfWithDefaults

`func NewPaymentEventAllOfWithDefaults() *PaymentEventAllOf`

NewPaymentEventAllOfWithDefaults instantiates a new PaymentEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentEventAllOf) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentEventAllOf) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentEventAllOf) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayoutPayorIds

`func (o *PaymentEventAllOf) GetPayoutPayorIds() PayoutPayorIds`

GetPayoutPayorIds returns the PayoutPayorIds field if non-nil, zero value otherwise.

### GetPayoutPayorIdsOk

`func (o *PaymentEventAllOf) GetPayoutPayorIdsOk() (*PayoutPayorIds, bool)`

GetPayoutPayorIdsOk returns a tuple with the PayoutPayorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutPayorIds

`func (o *PaymentEventAllOf) SetPayoutPayorIds(v PayoutPayorIds)`

SetPayoutPayorIds sets PayoutPayorIds field to given value.

### HasPayoutPayorIds

`func (o *PaymentEventAllOf) HasPayoutPayorIds() bool`

HasPayoutPayorIds returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *PaymentEventAllOf) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentEventAllOf) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentEventAllOf) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentEventAllOf) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


