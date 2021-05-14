# PayorAmlTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionDate** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **string** |  | [optional] 
**ReportTransactionType** | Pointer to **string** |  | [optional] 
**Debit** | Pointer to **int64** |  | [optional] 
**DebitCurrency** | Pointer to **string** | ISO 4217 3 character currency code | [optional] 
**Credit** | Pointer to **int64** |  | [optional] 
**CreditCurrency** | Pointer to **string** | ISO 4217 3 character currency code | [optional] 
**ReturnFee** | Pointer to **string** |  | [optional] 
**ReturnFeeCurrency** | Pointer to **string** | ISO 4217 3 character currency code | [optional] 
**ReturnFeeDescription** | Pointer to **string** |  | [optional] 
**ReturnCode** | Pointer to **string** |  | [optional] 
**ReturnDescription** | Pointer to **string** |  | [optional] 
**FundingType** | Pointer to **string** |  | [optional] 
**DateFundingRequested** | Pointer to **string** |  | [optional] 
**PayeeName** | Pointer to **string** |  | [optional] 
**RemoteId** | Pointer to **string** | Remote ID of the Payee, set by Payor | [optional] 
**PayeeType** | Pointer to **string** |  | [optional] 
**PayeeEmail** | Pointer to **string** |  | [optional] 
**SourceAccount** | Pointer to **string** |  | [optional] 
**PaymentAmount** | Pointer to **int64** |  | [optional] 
**PaymentCurrency** | Pointer to **string** | ISO 4217 3 character currency code | [optional] 
**PaymentMemo** | Pointer to **string** |  | [optional] 
**PaymentRails** | Pointer to **string** |  | [optional] 
**PayorPaymentId** | Pointer to **string** |  | [optional] 
**PaymentStatus** | Pointer to **string** |  | [optional] 
**RejectReason** | Pointer to **string** |  | [optional] 
**FxApplied** | Pointer to **float64** |  | [optional] 

## Methods

### NewPayorAmlTransaction

`func NewPayorAmlTransaction() *PayorAmlTransaction`

NewPayorAmlTransaction instantiates a new PayorAmlTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorAmlTransactionWithDefaults

`func NewPayorAmlTransactionWithDefaults() *PayorAmlTransaction`

NewPayorAmlTransactionWithDefaults instantiates a new PayorAmlTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionDate

`func (o *PayorAmlTransaction) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *PayorAmlTransaction) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *PayorAmlTransaction) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *PayorAmlTransaction) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionTime

`func (o *PayorAmlTransaction) GetTransactionTime() string`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PayorAmlTransaction) GetTransactionTimeOk() (*string, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PayorAmlTransaction) SetTransactionTime(v string)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *PayorAmlTransaction) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetReportTransactionType

`func (o *PayorAmlTransaction) GetReportTransactionType() string`

GetReportTransactionType returns the ReportTransactionType field if non-nil, zero value otherwise.

### GetReportTransactionTypeOk

`func (o *PayorAmlTransaction) GetReportTransactionTypeOk() (*string, bool)`

GetReportTransactionTypeOk returns a tuple with the ReportTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTransactionType

`func (o *PayorAmlTransaction) SetReportTransactionType(v string)`

SetReportTransactionType sets ReportTransactionType field to given value.

### HasReportTransactionType

`func (o *PayorAmlTransaction) HasReportTransactionType() bool`

HasReportTransactionType returns a boolean if a field has been set.

### GetDebit

`func (o *PayorAmlTransaction) GetDebit() int64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *PayorAmlTransaction) GetDebitOk() (*int64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *PayorAmlTransaction) SetDebit(v int64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *PayorAmlTransaction) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetDebitCurrency

`func (o *PayorAmlTransaction) GetDebitCurrency() string`

GetDebitCurrency returns the DebitCurrency field if non-nil, zero value otherwise.

### GetDebitCurrencyOk

`func (o *PayorAmlTransaction) GetDebitCurrencyOk() (*string, bool)`

GetDebitCurrencyOk returns a tuple with the DebitCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitCurrency

`func (o *PayorAmlTransaction) SetDebitCurrency(v string)`

SetDebitCurrency sets DebitCurrency field to given value.

### HasDebitCurrency

`func (o *PayorAmlTransaction) HasDebitCurrency() bool`

HasDebitCurrency returns a boolean if a field has been set.

### GetCredit

`func (o *PayorAmlTransaction) GetCredit() int64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *PayorAmlTransaction) GetCreditOk() (*int64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *PayorAmlTransaction) SetCredit(v int64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *PayorAmlTransaction) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCreditCurrency

`func (o *PayorAmlTransaction) GetCreditCurrency() string`

GetCreditCurrency returns the CreditCurrency field if non-nil, zero value otherwise.

### GetCreditCurrencyOk

`func (o *PayorAmlTransaction) GetCreditCurrencyOk() (*string, bool)`

GetCreditCurrencyOk returns a tuple with the CreditCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCurrency

`func (o *PayorAmlTransaction) SetCreditCurrency(v string)`

SetCreditCurrency sets CreditCurrency field to given value.

### HasCreditCurrency

`func (o *PayorAmlTransaction) HasCreditCurrency() bool`

HasCreditCurrency returns a boolean if a field has been set.

### GetReturnFee

`func (o *PayorAmlTransaction) GetReturnFee() string`

GetReturnFee returns the ReturnFee field if non-nil, zero value otherwise.

### GetReturnFeeOk

`func (o *PayorAmlTransaction) GetReturnFeeOk() (*string, bool)`

GetReturnFeeOk returns a tuple with the ReturnFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnFee

`func (o *PayorAmlTransaction) SetReturnFee(v string)`

SetReturnFee sets ReturnFee field to given value.

### HasReturnFee

`func (o *PayorAmlTransaction) HasReturnFee() bool`

HasReturnFee returns a boolean if a field has been set.

### GetReturnFeeCurrency

`func (o *PayorAmlTransaction) GetReturnFeeCurrency() string`

GetReturnFeeCurrency returns the ReturnFeeCurrency field if non-nil, zero value otherwise.

### GetReturnFeeCurrencyOk

`func (o *PayorAmlTransaction) GetReturnFeeCurrencyOk() (*string, bool)`

GetReturnFeeCurrencyOk returns a tuple with the ReturnFeeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnFeeCurrency

`func (o *PayorAmlTransaction) SetReturnFeeCurrency(v string)`

SetReturnFeeCurrency sets ReturnFeeCurrency field to given value.

### HasReturnFeeCurrency

`func (o *PayorAmlTransaction) HasReturnFeeCurrency() bool`

HasReturnFeeCurrency returns a boolean if a field has been set.

### GetReturnFeeDescription

`func (o *PayorAmlTransaction) GetReturnFeeDescription() string`

GetReturnFeeDescription returns the ReturnFeeDescription field if non-nil, zero value otherwise.

### GetReturnFeeDescriptionOk

`func (o *PayorAmlTransaction) GetReturnFeeDescriptionOk() (*string, bool)`

GetReturnFeeDescriptionOk returns a tuple with the ReturnFeeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnFeeDescription

`func (o *PayorAmlTransaction) SetReturnFeeDescription(v string)`

SetReturnFeeDescription sets ReturnFeeDescription field to given value.

### HasReturnFeeDescription

`func (o *PayorAmlTransaction) HasReturnFeeDescription() bool`

HasReturnFeeDescription returns a boolean if a field has been set.

### GetReturnCode

`func (o *PayorAmlTransaction) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *PayorAmlTransaction) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *PayorAmlTransaction) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *PayorAmlTransaction) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnDescription

`func (o *PayorAmlTransaction) GetReturnDescription() string`

GetReturnDescription returns the ReturnDescription field if non-nil, zero value otherwise.

### GetReturnDescriptionOk

`func (o *PayorAmlTransaction) GetReturnDescriptionOk() (*string, bool)`

GetReturnDescriptionOk returns a tuple with the ReturnDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDescription

`func (o *PayorAmlTransaction) SetReturnDescription(v string)`

SetReturnDescription sets ReturnDescription field to given value.

### HasReturnDescription

`func (o *PayorAmlTransaction) HasReturnDescription() bool`

HasReturnDescription returns a boolean if a field has been set.

### GetFundingType

`func (o *PayorAmlTransaction) GetFundingType() string`

GetFundingType returns the FundingType field if non-nil, zero value otherwise.

### GetFundingTypeOk

`func (o *PayorAmlTransaction) GetFundingTypeOk() (*string, bool)`

GetFundingTypeOk returns a tuple with the FundingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingType

`func (o *PayorAmlTransaction) SetFundingType(v string)`

SetFundingType sets FundingType field to given value.

### HasFundingType

`func (o *PayorAmlTransaction) HasFundingType() bool`

HasFundingType returns a boolean if a field has been set.

### GetDateFundingRequested

`func (o *PayorAmlTransaction) GetDateFundingRequested() string`

GetDateFundingRequested returns the DateFundingRequested field if non-nil, zero value otherwise.

### GetDateFundingRequestedOk

`func (o *PayorAmlTransaction) GetDateFundingRequestedOk() (*string, bool)`

GetDateFundingRequestedOk returns a tuple with the DateFundingRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFundingRequested

`func (o *PayorAmlTransaction) SetDateFundingRequested(v string)`

SetDateFundingRequested sets DateFundingRequested field to given value.

### HasDateFundingRequested

`func (o *PayorAmlTransaction) HasDateFundingRequested() bool`

HasDateFundingRequested returns a boolean if a field has been set.

### GetPayeeName

`func (o *PayorAmlTransaction) GetPayeeName() string`

GetPayeeName returns the PayeeName field if non-nil, zero value otherwise.

### GetPayeeNameOk

`func (o *PayorAmlTransaction) GetPayeeNameOk() (*string, bool)`

GetPayeeNameOk returns a tuple with the PayeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeName

`func (o *PayorAmlTransaction) SetPayeeName(v string)`

SetPayeeName sets PayeeName field to given value.

### HasPayeeName

`func (o *PayorAmlTransaction) HasPayeeName() bool`

HasPayeeName returns a boolean if a field has been set.

### GetRemoteId

`func (o *PayorAmlTransaction) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayorAmlTransaction) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayorAmlTransaction) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PayorAmlTransaction) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetPayeeType

`func (o *PayorAmlTransaction) GetPayeeType() string`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *PayorAmlTransaction) GetPayeeTypeOk() (*string, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *PayorAmlTransaction) SetPayeeType(v string)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *PayorAmlTransaction) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetPayeeEmail

`func (o *PayorAmlTransaction) GetPayeeEmail() string`

GetPayeeEmail returns the PayeeEmail field if non-nil, zero value otherwise.

### GetPayeeEmailOk

`func (o *PayorAmlTransaction) GetPayeeEmailOk() (*string, bool)`

GetPayeeEmailOk returns a tuple with the PayeeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeEmail

`func (o *PayorAmlTransaction) SetPayeeEmail(v string)`

SetPayeeEmail sets PayeeEmail field to given value.

### HasPayeeEmail

`func (o *PayorAmlTransaction) HasPayeeEmail() bool`

HasPayeeEmail returns a boolean if a field has been set.

### GetSourceAccount

`func (o *PayorAmlTransaction) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *PayorAmlTransaction) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *PayorAmlTransaction) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.

### HasSourceAccount

`func (o *PayorAmlTransaction) HasSourceAccount() bool`

HasSourceAccount returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *PayorAmlTransaction) GetPaymentAmount() int64`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PayorAmlTransaction) GetPaymentAmountOk() (*int64, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PayorAmlTransaction) SetPaymentAmount(v int64)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PayorAmlTransaction) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *PayorAmlTransaction) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PayorAmlTransaction) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PayorAmlTransaction) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PayorAmlTransaction) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetPaymentMemo

`func (o *PayorAmlTransaction) GetPaymentMemo() string`

GetPaymentMemo returns the PaymentMemo field if non-nil, zero value otherwise.

### GetPaymentMemoOk

`func (o *PayorAmlTransaction) GetPaymentMemoOk() (*string, bool)`

GetPaymentMemoOk returns a tuple with the PaymentMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMemo

`func (o *PayorAmlTransaction) SetPaymentMemo(v string)`

SetPaymentMemo sets PaymentMemo field to given value.

### HasPaymentMemo

`func (o *PayorAmlTransaction) HasPaymentMemo() bool`

HasPaymentMemo returns a boolean if a field has been set.

### GetPaymentRails

`func (o *PayorAmlTransaction) GetPaymentRails() string`

GetPaymentRails returns the PaymentRails field if non-nil, zero value otherwise.

### GetPaymentRailsOk

`func (o *PayorAmlTransaction) GetPaymentRailsOk() (*string, bool)`

GetPaymentRailsOk returns a tuple with the PaymentRails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRails

`func (o *PayorAmlTransaction) SetPaymentRails(v string)`

SetPaymentRails sets PaymentRails field to given value.

### HasPaymentRails

`func (o *PayorAmlTransaction) HasPaymentRails() bool`

HasPaymentRails returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *PayorAmlTransaction) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PayorAmlTransaction) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PayorAmlTransaction) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PayorAmlTransaction) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *PayorAmlTransaction) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *PayorAmlTransaction) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *PayorAmlTransaction) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *PayorAmlTransaction) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetRejectReason

`func (o *PayorAmlTransaction) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *PayorAmlTransaction) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *PayorAmlTransaction) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *PayorAmlTransaction) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetFxApplied

`func (o *PayorAmlTransaction) GetFxApplied() float64`

GetFxApplied returns the FxApplied field if non-nil, zero value otherwise.

### GetFxAppliedOk

`func (o *PayorAmlTransaction) GetFxAppliedOk() (*float64, bool)`

GetFxAppliedOk returns a tuple with the FxApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxApplied

`func (o *PayorAmlTransaction) SetFxApplied(v float64)`

SetFxApplied sets FxApplied field to given value.

### HasFxApplied

`func (o *PayorAmlTransaction) HasFxApplied() bool`

HasFxApplied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


