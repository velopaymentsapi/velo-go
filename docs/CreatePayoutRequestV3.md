# CreatePayoutRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | Pointer to **string** | Deprecated in v2.16. Any value supplied here will be ignored. | [optional] 
**PayoutFromPayorId** | Pointer to **string** | The id of the payor whose source account(s) will be debited. payoutFromPayorId and payoutToPayorId must be both supplied or both omitted. | [optional] 
**PayoutToPayorId** | Pointer to **string** | The id of the payor whose payees will be paid. payoutFromPayorId and payoutToPayorId must be both supplied or both omitted. | [optional] 
**PayoutMemo** | Pointer to **string** | Text applied to all payment memos unless specified explicitly on a payment. This should be the reference field on the statement seen by the payee (but not via ACH) | [optional] 
**Payments** | [**[]PaymentInstructionV3**](PaymentInstructionV3.md) |  | 

## Methods

### NewCreatePayoutRequestV3

`func NewCreatePayoutRequestV3(payments []PaymentInstructionV3, ) *CreatePayoutRequestV3`

NewCreatePayoutRequestV3 instantiates a new CreatePayoutRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayoutRequestV3WithDefaults

`func NewCreatePayoutRequestV3WithDefaults() *CreatePayoutRequestV3`

NewCreatePayoutRequestV3WithDefaults instantiates a new CreatePayoutRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *CreatePayoutRequestV3) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *CreatePayoutRequestV3) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *CreatePayoutRequestV3) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *CreatePayoutRequestV3) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetPayoutFromPayorId

`func (o *CreatePayoutRequestV3) GetPayoutFromPayorId() string`

GetPayoutFromPayorId returns the PayoutFromPayorId field if non-nil, zero value otherwise.

### GetPayoutFromPayorIdOk

`func (o *CreatePayoutRequestV3) GetPayoutFromPayorIdOk() (*string, bool)`

GetPayoutFromPayorIdOk returns a tuple with the PayoutFromPayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutFromPayorId

`func (o *CreatePayoutRequestV3) SetPayoutFromPayorId(v string)`

SetPayoutFromPayorId sets PayoutFromPayorId field to given value.

### HasPayoutFromPayorId

`func (o *CreatePayoutRequestV3) HasPayoutFromPayorId() bool`

HasPayoutFromPayorId returns a boolean if a field has been set.

### GetPayoutToPayorId

`func (o *CreatePayoutRequestV3) GetPayoutToPayorId() string`

GetPayoutToPayorId returns the PayoutToPayorId field if non-nil, zero value otherwise.

### GetPayoutToPayorIdOk

`func (o *CreatePayoutRequestV3) GetPayoutToPayorIdOk() (*string, bool)`

GetPayoutToPayorIdOk returns a tuple with the PayoutToPayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutToPayorId

`func (o *CreatePayoutRequestV3) SetPayoutToPayorId(v string)`

SetPayoutToPayorId sets PayoutToPayorId field to given value.

### HasPayoutToPayorId

`func (o *CreatePayoutRequestV3) HasPayoutToPayorId() bool`

HasPayoutToPayorId returns a boolean if a field has been set.

### GetPayoutMemo

`func (o *CreatePayoutRequestV3) GetPayoutMemo() string`

GetPayoutMemo returns the PayoutMemo field if non-nil, zero value otherwise.

### GetPayoutMemoOk

`func (o *CreatePayoutRequestV3) GetPayoutMemoOk() (*string, bool)`

GetPayoutMemoOk returns a tuple with the PayoutMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMemo

`func (o *CreatePayoutRequestV3) SetPayoutMemo(v string)`

SetPayoutMemo sets PayoutMemo field to given value.

### HasPayoutMemo

`func (o *CreatePayoutRequestV3) HasPayoutMemo() bool`

HasPayoutMemo returns a boolean if a field has been set.

### GetPayments

`func (o *CreatePayoutRequestV3) GetPayments() []PaymentInstructionV3`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *CreatePayoutRequestV3) GetPaymentsOk() (*[]PaymentInstructionV3, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *CreatePayoutRequestV3) SetPayments(v []PaymentInstructionV3)`

SetPayments sets Payments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


