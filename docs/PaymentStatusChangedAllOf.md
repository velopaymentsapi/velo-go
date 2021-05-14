# PaymentStatusChangedAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The new status of the payment. One of \&quot;SUBMITTED\&quot; \&quot;ACCEPTED\&quot; \&quot;REJECTED\&quot; \&quot;ACCEPTED_BY_RAILS\&quot; \&quot;CONFIRMED\&quot; \&quot;RETURNED\&quot; \&quot;WITHDRAWN\&quot; | 

## Methods

### NewPaymentStatusChangedAllOf

`func NewPaymentStatusChangedAllOf(status string, ) *PaymentStatusChangedAllOf`

NewPaymentStatusChangedAllOf instantiates a new PaymentStatusChangedAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentStatusChangedAllOfWithDefaults

`func NewPaymentStatusChangedAllOfWithDefaults() *PaymentStatusChangedAllOf`

NewPaymentStatusChangedAllOfWithDefaults instantiates a new PaymentStatusChangedAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaymentStatusChangedAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentStatusChangedAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentStatusChangedAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


