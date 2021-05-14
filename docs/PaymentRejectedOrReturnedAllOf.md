# PaymentRejectedOrReturnedAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonCode** | **string** | The Velo code that indicates why the payment was rejected or returned | 
**ReasonMessage** | **string** | The description of why the payment was rejected or returned | 

## Methods

### NewPaymentRejectedOrReturnedAllOf

`func NewPaymentRejectedOrReturnedAllOf(reasonCode string, reasonMessage string, ) *PaymentRejectedOrReturnedAllOf`

NewPaymentRejectedOrReturnedAllOf instantiates a new PaymentRejectedOrReturnedAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRejectedOrReturnedAllOfWithDefaults

`func NewPaymentRejectedOrReturnedAllOfWithDefaults() *PaymentRejectedOrReturnedAllOf`

NewPaymentRejectedOrReturnedAllOfWithDefaults instantiates a new PaymentRejectedOrReturnedAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonCode

`func (o *PaymentRejectedOrReturnedAllOf) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *PaymentRejectedOrReturnedAllOf) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *PaymentRejectedOrReturnedAllOf) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetReasonMessage

`func (o *PaymentRejectedOrReturnedAllOf) GetReasonMessage() string`

GetReasonMessage returns the ReasonMessage field if non-nil, zero value otherwise.

### GetReasonMessageOk

`func (o *PaymentRejectedOrReturnedAllOf) GetReasonMessageOk() (*string, bool)`

GetReasonMessageOk returns a tuple with the ReasonMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonMessage

`func (o *PaymentRejectedOrReturnedAllOf) SetReasonMessage(v string)`

SetReasonMessage sets ReasonMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


