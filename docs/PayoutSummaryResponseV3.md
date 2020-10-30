# PayoutSummaryResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PaymentsSubmitted** | Pointer to **int32** |  | [optional] 
**PaymentsAccepted** | Pointer to **int32** |  | [optional] 
**PaymentsRejected** | Pointer to **int32** |  | [optional] 
**PaymentsWithdrawn** | **int32** |  | 
**FxSummaries** | [**[]QuoteFxSummaryV3**](QuoteFxSummaryV3.md) |  | 
**Accounts** | [**[]SourceAccountV3**](SourceAccountV3.md) |  | 
**AcceptedPayments** | [**[]AcceptedPaymentV3**](AcceptedPaymentV3.md) |  | 
**RejectedPayments** | [**[]RejectedPaymentV3**](RejectedPaymentV3.md) |  | 

## Methods

### NewPayoutSummaryResponseV3

`func NewPayoutSummaryResponseV3(paymentsWithdrawn int32, fxSummaries []QuoteFxSummaryV3, accounts []SourceAccountV3, acceptedPayments []AcceptedPaymentV3, rejectedPayments []RejectedPaymentV3, ) *PayoutSummaryResponseV3`

NewPayoutSummaryResponseV3 instantiates a new PayoutSummaryResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutSummaryResponseV3WithDefaults

`func NewPayoutSummaryResponseV3WithDefaults() *PayoutSummaryResponseV3`

NewPayoutSummaryResponseV3WithDefaults instantiates a new PayoutSummaryResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutId

`func (o *PayoutSummaryResponseV3) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PayoutSummaryResponseV3) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PayoutSummaryResponseV3) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.

### HasPayoutId

`func (o *PayoutSummaryResponseV3) HasPayoutId() bool`

HasPayoutId returns a boolean if a field has been set.

### GetStatus

`func (o *PayoutSummaryResponseV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutSummaryResponseV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutSummaryResponseV3) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PayoutSummaryResponseV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentsSubmitted

`func (o *PayoutSummaryResponseV3) GetPaymentsSubmitted() int32`

GetPaymentsSubmitted returns the PaymentsSubmitted field if non-nil, zero value otherwise.

### GetPaymentsSubmittedOk

`func (o *PayoutSummaryResponseV3) GetPaymentsSubmittedOk() (*int32, bool)`

GetPaymentsSubmittedOk returns a tuple with the PaymentsSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsSubmitted

`func (o *PayoutSummaryResponseV3) SetPaymentsSubmitted(v int32)`

SetPaymentsSubmitted sets PaymentsSubmitted field to given value.

### HasPaymentsSubmitted

`func (o *PayoutSummaryResponseV3) HasPaymentsSubmitted() bool`

HasPaymentsSubmitted returns a boolean if a field has been set.

### GetPaymentsAccepted

`func (o *PayoutSummaryResponseV3) GetPaymentsAccepted() int32`

GetPaymentsAccepted returns the PaymentsAccepted field if non-nil, zero value otherwise.

### GetPaymentsAcceptedOk

`func (o *PayoutSummaryResponseV3) GetPaymentsAcceptedOk() (*int32, bool)`

GetPaymentsAcceptedOk returns a tuple with the PaymentsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsAccepted

`func (o *PayoutSummaryResponseV3) SetPaymentsAccepted(v int32)`

SetPaymentsAccepted sets PaymentsAccepted field to given value.

### HasPaymentsAccepted

`func (o *PayoutSummaryResponseV3) HasPaymentsAccepted() bool`

HasPaymentsAccepted returns a boolean if a field has been set.

### GetPaymentsRejected

`func (o *PayoutSummaryResponseV3) GetPaymentsRejected() int32`

GetPaymentsRejected returns the PaymentsRejected field if non-nil, zero value otherwise.

### GetPaymentsRejectedOk

`func (o *PayoutSummaryResponseV3) GetPaymentsRejectedOk() (*int32, bool)`

GetPaymentsRejectedOk returns a tuple with the PaymentsRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsRejected

`func (o *PayoutSummaryResponseV3) SetPaymentsRejected(v int32)`

SetPaymentsRejected sets PaymentsRejected field to given value.

### HasPaymentsRejected

`func (o *PayoutSummaryResponseV3) HasPaymentsRejected() bool`

HasPaymentsRejected returns a boolean if a field has been set.

### GetPaymentsWithdrawn

`func (o *PayoutSummaryResponseV3) GetPaymentsWithdrawn() int32`

GetPaymentsWithdrawn returns the PaymentsWithdrawn field if non-nil, zero value otherwise.

### GetPaymentsWithdrawnOk

`func (o *PayoutSummaryResponseV3) GetPaymentsWithdrawnOk() (*int32, bool)`

GetPaymentsWithdrawnOk returns a tuple with the PaymentsWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsWithdrawn

`func (o *PayoutSummaryResponseV3) SetPaymentsWithdrawn(v int32)`

SetPaymentsWithdrawn sets PaymentsWithdrawn field to given value.


### GetFxSummaries

`func (o *PayoutSummaryResponseV3) GetFxSummaries() []QuoteFxSummaryV3`

GetFxSummaries returns the FxSummaries field if non-nil, zero value otherwise.

### GetFxSummariesOk

`func (o *PayoutSummaryResponseV3) GetFxSummariesOk() (*[]QuoteFxSummaryV3, bool)`

GetFxSummariesOk returns a tuple with the FxSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxSummaries

`func (o *PayoutSummaryResponseV3) SetFxSummaries(v []QuoteFxSummaryV3)`

SetFxSummaries sets FxSummaries field to given value.


### GetAccounts

`func (o *PayoutSummaryResponseV3) GetAccounts() []SourceAccountV3`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PayoutSummaryResponseV3) GetAccountsOk() (*[]SourceAccountV3, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PayoutSummaryResponseV3) SetAccounts(v []SourceAccountV3)`

SetAccounts sets Accounts field to given value.


### GetAcceptedPayments

`func (o *PayoutSummaryResponseV3) GetAcceptedPayments() []AcceptedPaymentV3`

GetAcceptedPayments returns the AcceptedPayments field if non-nil, zero value otherwise.

### GetAcceptedPaymentsOk

`func (o *PayoutSummaryResponseV3) GetAcceptedPaymentsOk() (*[]AcceptedPaymentV3, bool)`

GetAcceptedPaymentsOk returns a tuple with the AcceptedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedPayments

`func (o *PayoutSummaryResponseV3) SetAcceptedPayments(v []AcceptedPaymentV3)`

SetAcceptedPayments sets AcceptedPayments field to given value.


### GetRejectedPayments

`func (o *PayoutSummaryResponseV3) GetRejectedPayments() []RejectedPaymentV3`

GetRejectedPayments returns the RejectedPayments field if non-nil, zero value otherwise.

### GetRejectedPaymentsOk

`func (o *PayoutSummaryResponseV3) GetRejectedPaymentsOk() (*[]RejectedPaymentV3, bool)`

GetRejectedPaymentsOk returns a tuple with the RejectedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedPayments

`func (o *PayoutSummaryResponseV3) SetRejectedPayments(v []RejectedPaymentV3)`

SetRejectedPayments sets RejectedPayments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


