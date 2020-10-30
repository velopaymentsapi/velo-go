# PayorAmlTransactionV4

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
**PaymentType** | Pointer to **string** |  | [optional] 
**PaymentRails** | Pointer to **string** |  | [optional] 
**PayorPaymentId** | Pointer to **string** |  | [optional] 
**PaymentStatus** | Pointer to **string** |  | [optional] 
**RejectReason** | Pointer to **string** |  | [optional] 
**FxApplied** | Pointer to **float64** |  | [optional] 

## Methods

### NewPayorAmlTransactionV4

`func NewPayorAmlTransactionV4() *PayorAmlTransactionV4`

NewPayorAmlTransactionV4 instantiates a new PayorAmlTransactionV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorAmlTransactionV4WithDefaults

`func NewPayorAmlTransactionV4WithDefaults() *PayorAmlTransactionV4`

NewPayorAmlTransactionV4WithDefaults instantiates a new PayorAmlTransactionV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionDate

`func (o *PayorAmlTransactionV4) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *PayorAmlTransactionV4) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *PayorAmlTransactionV4) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *PayorAmlTransactionV4) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionTime

`func (o *PayorAmlTransactionV4) GetTransactionTime() string`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PayorAmlTransactionV4) GetTransactionTimeOk() (*string, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PayorAmlTransactionV4) SetTransactionTime(v string)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *PayorAmlTransactionV4) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetReportTransactionType

`func (o *PayorAmlTransactionV4) GetReportTransactionType() string`

GetReportTransactionType returns the ReportTransactionType field if non-nil, zero value otherwise.

### GetReportTransactionTypeOk

`func (o *PayorAmlTransactionV4) GetReportTransactionTypeOk() (*string, bool)`

GetReportTransactionTypeOk returns a tuple with the ReportTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTransactionType

`func (o *PayorAmlTransactionV4) SetReportTransactionType(v string)`

SetReportTransactionType sets ReportTransactionType field to given value.

### HasReportTransactionType

`func (o *PayorAmlTransactionV4) HasReportTransactionType() bool`

HasReportTransactionType returns a boolean if a field has been set.

### GetDebit

`func (o *PayorAmlTransactionV4) GetDebit() int64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *PayorAmlTransactionV4) GetDebitOk() (*int64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *PayorAmlTransactionV4) SetDebit(v int64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *PayorAmlTransactionV4) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetDebitCurrency

`func (o *PayorAmlTransactionV4) GetDebitCurrency() string`

GetDebitCurrency returns the DebitCurrency field if non-nil, zero value otherwise.

### GetDebitCurrencyOk

`func (o *PayorAmlTransactionV4) GetDebitCurrencyOk() (*string, bool)`

GetDebitCurrencyOk returns a tuple with the DebitCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitCurrency

`func (o *PayorAmlTransactionV4) SetDebitCurrency(v string)`

SetDebitCurrency sets DebitCurrency field to given value.

### HasDebitCurrency

`func (o *PayorAmlTransactionV4) HasDebitCurrency() bool`

HasDebitCurrency returns a boolean if a field has been set.

### GetCredit

`func (o *PayorAmlTransactionV4) GetCredit() int64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *PayorAmlTransactionV4) GetCreditOk() (*int64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *PayorAmlTransactionV4) SetCredit(v int64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *PayorAmlTransactionV4) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCreditCurrency

`func (o *PayorAmlTransactionV4) GetCreditCurrency() string`

GetCreditCurrency returns the CreditCurrency field if non-nil, zero value otherwise.

### GetCreditCurrencyOk

`func (o *PayorAmlTransactionV4) GetCreditCurrencyOk() (*string, bool)`

GetCreditCurrencyOk returns a tuple with the CreditCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCurrency

`func (o *PayorAmlTransactionV4) SetCreditCurrency(v string)`

SetCreditCurrency sets CreditCurrency field to given value.

### HasCreditCurrency

`func (o *PayorAmlTransactionV4) HasCreditCurrency() bool`

HasCreditCurrency returns a boolean if a field has been set.

### GetReturnFee

`func (o *PayorAmlTransactionV4) GetReturnFee() string`

GetReturnFee returns the ReturnFee field if non-nil, zero value otherwise.

### GetReturnFeeOk

`func (o *PayorAmlTransactionV4) GetReturnFeeOk() (*string, bool)`

GetReturnFeeOk returns a tuple with the ReturnFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnFee

`func (o *PayorAmlTransactionV4) SetReturnFee(v string)`

SetReturnFee sets ReturnFee field to given value.

### HasReturnFee

`func (o *PayorAmlTransactionV4) HasReturnFee() bool`

HasReturnFee returns a boolean if a field has been set.

### GetReturnFeeCurrency

`func (o *PayorAmlTransactionV4) GetReturnFeeCurrency() string`

GetReturnFeeCurrency returns the ReturnFeeCurrency field if non-nil, zero value otherwise.

### GetReturnFeeCurrencyOk

`func (o *PayorAmlTransactionV4) GetReturnFeeCurrencyOk() (*string, bool)`

GetReturnFeeCurrencyOk returns a tuple with the ReturnFeeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnFeeCurrency

`func (o *PayorAmlTransactionV4) SetReturnFeeCurrency(v string)`

SetReturnFeeCurrency sets ReturnFeeCurrency field to given value.

### HasReturnFeeCurrency

`func (o *PayorAmlTransactionV4) HasReturnFeeCurrency() bool`

HasReturnFeeCurrency returns a boolean if a field has been set.

### GetReturnFeeDescription

`func (o *PayorAmlTransactionV4) GetReturnFeeDescription() string`

GetReturnFeeDescription returns the ReturnFeeDescription field if non-nil, zero value otherwise.

### GetReturnFeeDescriptionOk

`func (o *PayorAmlTransactionV4) GetReturnFeeDescriptionOk() (*string, bool)`

GetReturnFeeDescriptionOk returns a tuple with the ReturnFeeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnFeeDescription

`func (o *PayorAmlTransactionV4) SetReturnFeeDescription(v string)`

SetReturnFeeDescription sets ReturnFeeDescription field to given value.

### HasReturnFeeDescription

`func (o *PayorAmlTransactionV4) HasReturnFeeDescription() bool`

HasReturnFeeDescription returns a boolean if a field has been set.

### GetReturnCode

`func (o *PayorAmlTransactionV4) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *PayorAmlTransactionV4) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *PayorAmlTransactionV4) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *PayorAmlTransactionV4) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnDescription

`func (o *PayorAmlTransactionV4) GetReturnDescription() string`

GetReturnDescription returns the ReturnDescription field if non-nil, zero value otherwise.

### GetReturnDescriptionOk

`func (o *PayorAmlTransactionV4) GetReturnDescriptionOk() (*string, bool)`

GetReturnDescriptionOk returns a tuple with the ReturnDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDescription

`func (o *PayorAmlTransactionV4) SetReturnDescription(v string)`

SetReturnDescription sets ReturnDescription field to given value.

### HasReturnDescription

`func (o *PayorAmlTransactionV4) HasReturnDescription() bool`

HasReturnDescription returns a boolean if a field has been set.

### GetFundingType

`func (o *PayorAmlTransactionV4) GetFundingType() string`

GetFundingType returns the FundingType field if non-nil, zero value otherwise.

### GetFundingTypeOk

`func (o *PayorAmlTransactionV4) GetFundingTypeOk() (*string, bool)`

GetFundingTypeOk returns a tuple with the FundingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingType

`func (o *PayorAmlTransactionV4) SetFundingType(v string)`

SetFundingType sets FundingType field to given value.

### HasFundingType

`func (o *PayorAmlTransactionV4) HasFundingType() bool`

HasFundingType returns a boolean if a field has been set.

### GetDateFundingRequested

`func (o *PayorAmlTransactionV4) GetDateFundingRequested() string`

GetDateFundingRequested returns the DateFundingRequested field if non-nil, zero value otherwise.

### GetDateFundingRequestedOk

`func (o *PayorAmlTransactionV4) GetDateFundingRequestedOk() (*string, bool)`

GetDateFundingRequestedOk returns a tuple with the DateFundingRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFundingRequested

`func (o *PayorAmlTransactionV4) SetDateFundingRequested(v string)`

SetDateFundingRequested sets DateFundingRequested field to given value.

### HasDateFundingRequested

`func (o *PayorAmlTransactionV4) HasDateFundingRequested() bool`

HasDateFundingRequested returns a boolean if a field has been set.

### GetPayeeName

`func (o *PayorAmlTransactionV4) GetPayeeName() string`

GetPayeeName returns the PayeeName field if non-nil, zero value otherwise.

### GetPayeeNameOk

`func (o *PayorAmlTransactionV4) GetPayeeNameOk() (*string, bool)`

GetPayeeNameOk returns a tuple with the PayeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeName

`func (o *PayorAmlTransactionV4) SetPayeeName(v string)`

SetPayeeName sets PayeeName field to given value.

### HasPayeeName

`func (o *PayorAmlTransactionV4) HasPayeeName() bool`

HasPayeeName returns a boolean if a field has been set.

### GetRemoteId

`func (o *PayorAmlTransactionV4) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayorAmlTransactionV4) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayorAmlTransactionV4) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PayorAmlTransactionV4) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetPayeeType

`func (o *PayorAmlTransactionV4) GetPayeeType() string`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *PayorAmlTransactionV4) GetPayeeTypeOk() (*string, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *PayorAmlTransactionV4) SetPayeeType(v string)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *PayorAmlTransactionV4) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetPayeeEmail

`func (o *PayorAmlTransactionV4) GetPayeeEmail() string`

GetPayeeEmail returns the PayeeEmail field if non-nil, zero value otherwise.

### GetPayeeEmailOk

`func (o *PayorAmlTransactionV4) GetPayeeEmailOk() (*string, bool)`

GetPayeeEmailOk returns a tuple with the PayeeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeEmail

`func (o *PayorAmlTransactionV4) SetPayeeEmail(v string)`

SetPayeeEmail sets PayeeEmail field to given value.

### HasPayeeEmail

`func (o *PayorAmlTransactionV4) HasPayeeEmail() bool`

HasPayeeEmail returns a boolean if a field has been set.

### GetSourceAccount

`func (o *PayorAmlTransactionV4) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *PayorAmlTransactionV4) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *PayorAmlTransactionV4) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.

### HasSourceAccount

`func (o *PayorAmlTransactionV4) HasSourceAccount() bool`

HasSourceAccount returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *PayorAmlTransactionV4) GetPaymentAmount() int64`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PayorAmlTransactionV4) GetPaymentAmountOk() (*int64, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PayorAmlTransactionV4) SetPaymentAmount(v int64)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PayorAmlTransactionV4) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *PayorAmlTransactionV4) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PayorAmlTransactionV4) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PayorAmlTransactionV4) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PayorAmlTransactionV4) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetPaymentMemo

`func (o *PayorAmlTransactionV4) GetPaymentMemo() string`

GetPaymentMemo returns the PaymentMemo field if non-nil, zero value otherwise.

### GetPaymentMemoOk

`func (o *PayorAmlTransactionV4) GetPaymentMemoOk() (*string, bool)`

GetPaymentMemoOk returns a tuple with the PaymentMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMemo

`func (o *PayorAmlTransactionV4) SetPaymentMemo(v string)`

SetPaymentMemo sets PaymentMemo field to given value.

### HasPaymentMemo

`func (o *PayorAmlTransactionV4) HasPaymentMemo() bool`

HasPaymentMemo returns a boolean if a field has been set.

### GetPaymentType

`func (o *PayorAmlTransactionV4) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PayorAmlTransactionV4) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PayorAmlTransactionV4) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PayorAmlTransactionV4) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentRails

`func (o *PayorAmlTransactionV4) GetPaymentRails() string`

GetPaymentRails returns the PaymentRails field if non-nil, zero value otherwise.

### GetPaymentRailsOk

`func (o *PayorAmlTransactionV4) GetPaymentRailsOk() (*string, bool)`

GetPaymentRailsOk returns a tuple with the PaymentRails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRails

`func (o *PayorAmlTransactionV4) SetPaymentRails(v string)`

SetPaymentRails sets PaymentRails field to given value.

### HasPaymentRails

`func (o *PayorAmlTransactionV4) HasPaymentRails() bool`

HasPaymentRails returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *PayorAmlTransactionV4) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PayorAmlTransactionV4) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PayorAmlTransactionV4) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PayorAmlTransactionV4) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *PayorAmlTransactionV4) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *PayorAmlTransactionV4) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *PayorAmlTransactionV4) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *PayorAmlTransactionV4) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetRejectReason

`func (o *PayorAmlTransactionV4) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *PayorAmlTransactionV4) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *PayorAmlTransactionV4) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *PayorAmlTransactionV4) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetFxApplied

`func (o *PayorAmlTransactionV4) GetFxApplied() float64`

GetFxApplied returns the FxApplied field if non-nil, zero value otherwise.

### GetFxAppliedOk

`func (o *PayorAmlTransactionV4) GetFxAppliedOk() (*float64, bool)`

GetFxAppliedOk returns a tuple with the FxApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxApplied

`func (o *PayorAmlTransactionV4) SetFxApplied(v float64)`

SetFxApplied sets FxApplied field to given value.

### HasFxApplied

`func (o *PayorAmlTransactionV4) HasFxApplied() bool`

HasFxApplied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


