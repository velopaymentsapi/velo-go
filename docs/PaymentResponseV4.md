# PaymentResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | The id of the payment | 
**PayeeId** | **string** | The id of the paymeee | 
**PayorId** | **string** | The id of the payor | 
**PayorName** | Pointer to **string** | The name of the payor | [optional] 
**QuoteId** | **string** | The quote Id used for the FX | 
**SourceAccountId** | **string** | The id of the source account from which the payment was taken | 
**SourceAccountName** | Pointer to **string** | The name of the source account from which the payment was taken | [optional] 
**RemoteId** | Pointer to **string** | The remote id by which the payor refers to the payee. Only populated once payment is confirmed | [optional] 
**RemoteSystemId** | Pointer to **string** | The velo id of the remote system orchestrating the payment. Not populated for normal Velo payments. | [optional] 
**RemoteSystemPaymentId** | Pointer to **string** | The id of the payment in the remote system. Not populated for normal Velo payments. | [optional] 
**SourceAmount** | Pointer to **int32** | The source amount for the payment (amount debited to make the payment) | [optional] 
**SourceCurrency** | Pointer to **string** | ISO-4217 3 character currency code | [optional] 
**PaymentAmount** | **int32** | The amount which the payee will receive | 
**PaymentCurrency** | Pointer to **string** | ISO-4217 3 character currency code | [optional] 
**Rate** | Pointer to **float64** | The FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed | [optional] 
**InvertedRate** | Pointer to **float64** | The inverted FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed | [optional] 
**IsPaymentCcyBaseCcy** | Pointer to **bool** |  | [optional] 
**SubmittedDateTime** | **time.Time** |  | 
**Status** | **string** | Current status of the payment. One of the following values: ACCEPTED, AWAITING_FUNDS, FUNDED, UNFUNDED, BANK_PAYMENT_REQUESTED, REJECTED, ACCEPTED_BY_RAILS, CONFIRMED, RETURNED, WITHDRAWN | 
**FundingStatus** | **string** | Current funding status of the payment. One of the following values: FUNDED, INSTRUCTED, UNFUNDED | 
**RoutingNumber** | Pointer to **string** | The routing number for the payment. | [optional] 
**AccountNumber** | Pointer to **string** | The account number for the account which will receive the payment. | [optional] 
**Iban** | Pointer to **string** | The iban for the payment. | [optional] 
**PaymentMemo** | Pointer to **string** | The payment memo set by the payor | [optional] 
**FilenameReference** | Pointer to **string** | ACH file payment was submitted in, if applicable | [optional] 
**IndividualIdentificationNumber** | Pointer to **string** | Individual Identification Number assigned to the payment in the ACH file, if applicable | [optional] 
**TraceNumber** | Pointer to **string** | Trace Number assigned to the payment in the ACH file, if applicable | [optional] 
**PayorPaymentId** | Pointer to **string** |  | [optional] 
**PaymentChannelId** | Pointer to **string** |  | [optional] 
**PaymentChannelName** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**RailsId** | **string** | The rails ID. Default value is RAILS ID UNAVAILABLE when not populated. | [default to "RAILS ID UNAVAILABLE"]
**CountryCode** | Pointer to **string** | The country code of the payment channel. | [optional] 
**PayeeAddressCountryCode** | Pointer to **string** | The country code of the payee&#39;s address. | [optional] 
**Events** | [**[]PaymentEventResponse**](PaymentEventResponse.md) |  | 
**ReturnCost** | Pointer to **int32** | The return cost if a returned payment. | [optional] 
**ReturnReason** | Pointer to **string** |  | [optional] 
**RailsPaymentId** | Pointer to **string** |  | [optional] 
**RailsBatchId** | Pointer to **string** |  | [optional] 
**PaymentScheme** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **string** |  | [optional] 
**WithdrawnReason** | Pointer to **string** |  | [optional] 
**Withdrawable** | Pointer to **bool** |  | [optional] 
**AutoWithdrawnReasonCode** | Pointer to **string** | Populated with rejection reason code if the payment was withdrawn automatically at instruct time | [optional] 
**TransmissionType** | Pointer to **string** | The transmission type of the payment, e.g. ACH, SAME_DAY_ACH, WIRE, GACHO | [optional] 
**PaymentTrackingReference** | Pointer to **string** |  | [optional] 
**PaymentMetadata** | Pointer to **string** | Metadata for the payment | [optional] 
**Schedule** | Pointer to [**PayoutSchedule**](PayoutSchedule.md) |  | [optional] 
**PostInstructFxInfo** | Pointer to [**PostInstructFxInfo**](PostInstructFxInfo.md) |  | [optional] 
**Payout** | Pointer to [**PaymentResponseV4Payout**](PaymentResponseV4Payout.md) |  | [optional] 

## Methods

### NewPaymentResponseV4

`func NewPaymentResponseV4(paymentId string, payeeId string, payorId string, quoteId string, sourceAccountId string, paymentAmount int32, submittedDateTime time.Time, status string, fundingStatus string, railsId string, events []PaymentEventResponse, ) *PaymentResponseV4`

NewPaymentResponseV4 instantiates a new PaymentResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseV4WithDefaults

`func NewPaymentResponseV4WithDefaults() *PaymentResponseV4`

NewPaymentResponseV4WithDefaults instantiates a new PaymentResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentResponseV4) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentResponseV4) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentResponseV4) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayeeId

`func (o *PaymentResponseV4) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PaymentResponseV4) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PaymentResponseV4) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetPayorId

`func (o *PaymentResponseV4) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PaymentResponseV4) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PaymentResponseV4) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetPayorName

`func (o *PaymentResponseV4) GetPayorName() string`

GetPayorName returns the PayorName field if non-nil, zero value otherwise.

### GetPayorNameOk

`func (o *PaymentResponseV4) GetPayorNameOk() (*string, bool)`

GetPayorNameOk returns a tuple with the PayorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorName

`func (o *PaymentResponseV4) SetPayorName(v string)`

SetPayorName sets PayorName field to given value.

### HasPayorName

`func (o *PaymentResponseV4) HasPayorName() bool`

HasPayorName returns a boolean if a field has been set.

### GetQuoteId

`func (o *PaymentResponseV4) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *PaymentResponseV4) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *PaymentResponseV4) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetSourceAccountId

`func (o *PaymentResponseV4) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *PaymentResponseV4) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *PaymentResponseV4) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.


### GetSourceAccountName

`func (o *PaymentResponseV4) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *PaymentResponseV4) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *PaymentResponseV4) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *PaymentResponseV4) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetRemoteId

`func (o *PaymentResponseV4) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PaymentResponseV4) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PaymentResponseV4) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PaymentResponseV4) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetRemoteSystemId

`func (o *PaymentResponseV4) GetRemoteSystemId() string`

GetRemoteSystemId returns the RemoteSystemId field if non-nil, zero value otherwise.

### GetRemoteSystemIdOk

`func (o *PaymentResponseV4) GetRemoteSystemIdOk() (*string, bool)`

GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSystemId

`func (o *PaymentResponseV4) SetRemoteSystemId(v string)`

SetRemoteSystemId sets RemoteSystemId field to given value.

### HasRemoteSystemId

`func (o *PaymentResponseV4) HasRemoteSystemId() bool`

HasRemoteSystemId returns a boolean if a field has been set.

### GetRemoteSystemPaymentId

`func (o *PaymentResponseV4) GetRemoteSystemPaymentId() string`

GetRemoteSystemPaymentId returns the RemoteSystemPaymentId field if non-nil, zero value otherwise.

### GetRemoteSystemPaymentIdOk

`func (o *PaymentResponseV4) GetRemoteSystemPaymentIdOk() (*string, bool)`

GetRemoteSystemPaymentIdOk returns a tuple with the RemoteSystemPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSystemPaymentId

`func (o *PaymentResponseV4) SetRemoteSystemPaymentId(v string)`

SetRemoteSystemPaymentId sets RemoteSystemPaymentId field to given value.

### HasRemoteSystemPaymentId

`func (o *PaymentResponseV4) HasRemoteSystemPaymentId() bool`

HasRemoteSystemPaymentId returns a boolean if a field has been set.

### GetSourceAmount

`func (o *PaymentResponseV4) GetSourceAmount() int32`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *PaymentResponseV4) GetSourceAmountOk() (*int32, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *PaymentResponseV4) SetSourceAmount(v int32)`

SetSourceAmount sets SourceAmount field to given value.

### HasSourceAmount

`func (o *PaymentResponseV4) HasSourceAmount() bool`

HasSourceAmount returns a boolean if a field has been set.

### GetSourceCurrency

`func (o *PaymentResponseV4) GetSourceCurrency() string`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *PaymentResponseV4) GetSourceCurrencyOk() (*string, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *PaymentResponseV4) SetSourceCurrency(v string)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *PaymentResponseV4) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *PaymentResponseV4) GetPaymentAmount() int32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentResponseV4) GetPaymentAmountOk() (*int32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentResponseV4) SetPaymentAmount(v int32)`

SetPaymentAmount sets PaymentAmount field to given value.


### GetPaymentCurrency

`func (o *PaymentResponseV4) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PaymentResponseV4) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PaymentResponseV4) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PaymentResponseV4) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetRate

`func (o *PaymentResponseV4) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *PaymentResponseV4) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *PaymentResponseV4) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *PaymentResponseV4) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetInvertedRate

`func (o *PaymentResponseV4) GetInvertedRate() float64`

GetInvertedRate returns the InvertedRate field if non-nil, zero value otherwise.

### GetInvertedRateOk

`func (o *PaymentResponseV4) GetInvertedRateOk() (*float64, bool)`

GetInvertedRateOk returns a tuple with the InvertedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertedRate

`func (o *PaymentResponseV4) SetInvertedRate(v float64)`

SetInvertedRate sets InvertedRate field to given value.

### HasInvertedRate

`func (o *PaymentResponseV4) HasInvertedRate() bool`

HasInvertedRate returns a boolean if a field has been set.

### GetIsPaymentCcyBaseCcy

`func (o *PaymentResponseV4) GetIsPaymentCcyBaseCcy() bool`

GetIsPaymentCcyBaseCcy returns the IsPaymentCcyBaseCcy field if non-nil, zero value otherwise.

### GetIsPaymentCcyBaseCcyOk

`func (o *PaymentResponseV4) GetIsPaymentCcyBaseCcyOk() (*bool, bool)`

GetIsPaymentCcyBaseCcyOk returns a tuple with the IsPaymentCcyBaseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaymentCcyBaseCcy

`func (o *PaymentResponseV4) SetIsPaymentCcyBaseCcy(v bool)`

SetIsPaymentCcyBaseCcy sets IsPaymentCcyBaseCcy field to given value.

### HasIsPaymentCcyBaseCcy

`func (o *PaymentResponseV4) HasIsPaymentCcyBaseCcy() bool`

HasIsPaymentCcyBaseCcy returns a boolean if a field has been set.

### GetSubmittedDateTime

`func (o *PaymentResponseV4) GetSubmittedDateTime() time.Time`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *PaymentResponseV4) GetSubmittedDateTimeOk() (*time.Time, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *PaymentResponseV4) SetSubmittedDateTime(v time.Time)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.


### GetStatus

`func (o *PaymentResponseV4) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentResponseV4) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentResponseV4) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFundingStatus

`func (o *PaymentResponseV4) GetFundingStatus() string`

GetFundingStatus returns the FundingStatus field if non-nil, zero value otherwise.

### GetFundingStatusOk

`func (o *PaymentResponseV4) GetFundingStatusOk() (*string, bool)`

GetFundingStatusOk returns a tuple with the FundingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingStatus

`func (o *PaymentResponseV4) SetFundingStatus(v string)`

SetFundingStatus sets FundingStatus field to given value.


### GetRoutingNumber

`func (o *PaymentResponseV4) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PaymentResponseV4) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PaymentResponseV4) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *PaymentResponseV4) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PaymentResponseV4) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentResponseV4) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentResponseV4) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PaymentResponseV4) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetIban

`func (o *PaymentResponseV4) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentResponseV4) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentResponseV4) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PaymentResponseV4) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetPaymentMemo

`func (o *PaymentResponseV4) GetPaymentMemo() string`

GetPaymentMemo returns the PaymentMemo field if non-nil, zero value otherwise.

### GetPaymentMemoOk

`func (o *PaymentResponseV4) GetPaymentMemoOk() (*string, bool)`

GetPaymentMemoOk returns a tuple with the PaymentMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMemo

`func (o *PaymentResponseV4) SetPaymentMemo(v string)`

SetPaymentMemo sets PaymentMemo field to given value.

### HasPaymentMemo

`func (o *PaymentResponseV4) HasPaymentMemo() bool`

HasPaymentMemo returns a boolean if a field has been set.

### GetFilenameReference

`func (o *PaymentResponseV4) GetFilenameReference() string`

GetFilenameReference returns the FilenameReference field if non-nil, zero value otherwise.

### GetFilenameReferenceOk

`func (o *PaymentResponseV4) GetFilenameReferenceOk() (*string, bool)`

GetFilenameReferenceOk returns a tuple with the FilenameReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenameReference

`func (o *PaymentResponseV4) SetFilenameReference(v string)`

SetFilenameReference sets FilenameReference field to given value.

### HasFilenameReference

`func (o *PaymentResponseV4) HasFilenameReference() bool`

HasFilenameReference returns a boolean if a field has been set.

### GetIndividualIdentificationNumber

`func (o *PaymentResponseV4) GetIndividualIdentificationNumber() string`

GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field if non-nil, zero value otherwise.

### GetIndividualIdentificationNumberOk

`func (o *PaymentResponseV4) GetIndividualIdentificationNumberOk() (*string, bool)`

GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentificationNumber

`func (o *PaymentResponseV4) SetIndividualIdentificationNumber(v string)`

SetIndividualIdentificationNumber sets IndividualIdentificationNumber field to given value.

### HasIndividualIdentificationNumber

`func (o *PaymentResponseV4) HasIndividualIdentificationNumber() bool`

HasIndividualIdentificationNumber returns a boolean if a field has been set.

### GetTraceNumber

`func (o *PaymentResponseV4) GetTraceNumber() string`

GetTraceNumber returns the TraceNumber field if non-nil, zero value otherwise.

### GetTraceNumberOk

`func (o *PaymentResponseV4) GetTraceNumberOk() (*string, bool)`

GetTraceNumberOk returns a tuple with the TraceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceNumber

`func (o *PaymentResponseV4) SetTraceNumber(v string)`

SetTraceNumber sets TraceNumber field to given value.

### HasTraceNumber

`func (o *PaymentResponseV4) HasTraceNumber() bool`

HasTraceNumber returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *PaymentResponseV4) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentResponseV4) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentResponseV4) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentResponseV4) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetPaymentChannelId

`func (o *PaymentResponseV4) GetPaymentChannelId() string`

GetPaymentChannelId returns the PaymentChannelId field if non-nil, zero value otherwise.

### GetPaymentChannelIdOk

`func (o *PaymentResponseV4) GetPaymentChannelIdOk() (*string, bool)`

GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelId

`func (o *PaymentResponseV4) SetPaymentChannelId(v string)`

SetPaymentChannelId sets PaymentChannelId field to given value.

### HasPaymentChannelId

`func (o *PaymentResponseV4) HasPaymentChannelId() bool`

HasPaymentChannelId returns a boolean if a field has been set.

### GetPaymentChannelName

`func (o *PaymentResponseV4) GetPaymentChannelName() string`

GetPaymentChannelName returns the PaymentChannelName field if non-nil, zero value otherwise.

### GetPaymentChannelNameOk

`func (o *PaymentResponseV4) GetPaymentChannelNameOk() (*string, bool)`

GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelName

`func (o *PaymentResponseV4) SetPaymentChannelName(v string)`

SetPaymentChannelName sets PaymentChannelName field to given value.

### HasPaymentChannelName

`func (o *PaymentResponseV4) HasPaymentChannelName() bool`

HasPaymentChannelName returns a boolean if a field has been set.

### GetAccountName

`func (o *PaymentResponseV4) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentResponseV4) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentResponseV4) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PaymentResponseV4) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetRailsId

`func (o *PaymentResponseV4) GetRailsId() string`

GetRailsId returns the RailsId field if non-nil, zero value otherwise.

### GetRailsIdOk

`func (o *PaymentResponseV4) GetRailsIdOk() (*string, bool)`

GetRailsIdOk returns a tuple with the RailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsId

`func (o *PaymentResponseV4) SetRailsId(v string)`

SetRailsId sets RailsId field to given value.


### GetCountryCode

`func (o *PaymentResponseV4) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PaymentResponseV4) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PaymentResponseV4) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PaymentResponseV4) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPayeeAddressCountryCode

`func (o *PaymentResponseV4) GetPayeeAddressCountryCode() string`

GetPayeeAddressCountryCode returns the PayeeAddressCountryCode field if non-nil, zero value otherwise.

### GetPayeeAddressCountryCodeOk

`func (o *PaymentResponseV4) GetPayeeAddressCountryCodeOk() (*string, bool)`

GetPayeeAddressCountryCodeOk returns a tuple with the PayeeAddressCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeAddressCountryCode

`func (o *PaymentResponseV4) SetPayeeAddressCountryCode(v string)`

SetPayeeAddressCountryCode sets PayeeAddressCountryCode field to given value.

### HasPayeeAddressCountryCode

`func (o *PaymentResponseV4) HasPayeeAddressCountryCode() bool`

HasPayeeAddressCountryCode returns a boolean if a field has been set.

### GetEvents

`func (o *PaymentResponseV4) GetEvents() []PaymentEventResponse`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PaymentResponseV4) GetEventsOk() (*[]PaymentEventResponse, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PaymentResponseV4) SetEvents(v []PaymentEventResponse)`

SetEvents sets Events field to given value.


### GetReturnCost

`func (o *PaymentResponseV4) GetReturnCost() int32`

GetReturnCost returns the ReturnCost field if non-nil, zero value otherwise.

### GetReturnCostOk

`func (o *PaymentResponseV4) GetReturnCostOk() (*int32, bool)`

GetReturnCostOk returns a tuple with the ReturnCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCost

`func (o *PaymentResponseV4) SetReturnCost(v int32)`

SetReturnCost sets ReturnCost field to given value.

### HasReturnCost

`func (o *PaymentResponseV4) HasReturnCost() bool`

HasReturnCost returns a boolean if a field has been set.

### GetReturnReason

`func (o *PaymentResponseV4) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *PaymentResponseV4) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *PaymentResponseV4) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *PaymentResponseV4) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetRailsPaymentId

`func (o *PaymentResponseV4) GetRailsPaymentId() string`

GetRailsPaymentId returns the RailsPaymentId field if non-nil, zero value otherwise.

### GetRailsPaymentIdOk

`func (o *PaymentResponseV4) GetRailsPaymentIdOk() (*string, bool)`

GetRailsPaymentIdOk returns a tuple with the RailsPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsPaymentId

`func (o *PaymentResponseV4) SetRailsPaymentId(v string)`

SetRailsPaymentId sets RailsPaymentId field to given value.

### HasRailsPaymentId

`func (o *PaymentResponseV4) HasRailsPaymentId() bool`

HasRailsPaymentId returns a boolean if a field has been set.

### GetRailsBatchId

`func (o *PaymentResponseV4) GetRailsBatchId() string`

GetRailsBatchId returns the RailsBatchId field if non-nil, zero value otherwise.

### GetRailsBatchIdOk

`func (o *PaymentResponseV4) GetRailsBatchIdOk() (*string, bool)`

GetRailsBatchIdOk returns a tuple with the RailsBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsBatchId

`func (o *PaymentResponseV4) SetRailsBatchId(v string)`

SetRailsBatchId sets RailsBatchId field to given value.

### HasRailsBatchId

`func (o *PaymentResponseV4) HasRailsBatchId() bool`

HasRailsBatchId returns a boolean if a field has been set.

### GetPaymentScheme

`func (o *PaymentResponseV4) GetPaymentScheme() string`

GetPaymentScheme returns the PaymentScheme field if non-nil, zero value otherwise.

### GetPaymentSchemeOk

`func (o *PaymentResponseV4) GetPaymentSchemeOk() (*string, bool)`

GetPaymentSchemeOk returns a tuple with the PaymentScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentScheme

`func (o *PaymentResponseV4) SetPaymentScheme(v string)`

SetPaymentScheme sets PaymentScheme field to given value.

### HasPaymentScheme

`func (o *PaymentResponseV4) HasPaymentScheme() bool`

HasPaymentScheme returns a boolean if a field has been set.

### GetRejectionReason

`func (o *PaymentResponseV4) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *PaymentResponseV4) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *PaymentResponseV4) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *PaymentResponseV4) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetWithdrawnReason

`func (o *PaymentResponseV4) GetWithdrawnReason() string`

GetWithdrawnReason returns the WithdrawnReason field if non-nil, zero value otherwise.

### GetWithdrawnReasonOk

`func (o *PaymentResponseV4) GetWithdrawnReasonOk() (*string, bool)`

GetWithdrawnReasonOk returns a tuple with the WithdrawnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnReason

`func (o *PaymentResponseV4) SetWithdrawnReason(v string)`

SetWithdrawnReason sets WithdrawnReason field to given value.

### HasWithdrawnReason

`func (o *PaymentResponseV4) HasWithdrawnReason() bool`

HasWithdrawnReason returns a boolean if a field has been set.

### GetWithdrawable

`func (o *PaymentResponseV4) GetWithdrawable() bool`

GetWithdrawable returns the Withdrawable field if non-nil, zero value otherwise.

### GetWithdrawableOk

`func (o *PaymentResponseV4) GetWithdrawableOk() (*bool, bool)`

GetWithdrawableOk returns a tuple with the Withdrawable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawable

`func (o *PaymentResponseV4) SetWithdrawable(v bool)`

SetWithdrawable sets Withdrawable field to given value.

### HasWithdrawable

`func (o *PaymentResponseV4) HasWithdrawable() bool`

HasWithdrawable returns a boolean if a field has been set.

### GetAutoWithdrawnReasonCode

`func (o *PaymentResponseV4) GetAutoWithdrawnReasonCode() string`

GetAutoWithdrawnReasonCode returns the AutoWithdrawnReasonCode field if non-nil, zero value otherwise.

### GetAutoWithdrawnReasonCodeOk

`func (o *PaymentResponseV4) GetAutoWithdrawnReasonCodeOk() (*string, bool)`

GetAutoWithdrawnReasonCodeOk returns a tuple with the AutoWithdrawnReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWithdrawnReasonCode

`func (o *PaymentResponseV4) SetAutoWithdrawnReasonCode(v string)`

SetAutoWithdrawnReasonCode sets AutoWithdrawnReasonCode field to given value.

### HasAutoWithdrawnReasonCode

`func (o *PaymentResponseV4) HasAutoWithdrawnReasonCode() bool`

HasAutoWithdrawnReasonCode returns a boolean if a field has been set.

### GetTransmissionType

`func (o *PaymentResponseV4) GetTransmissionType() string`

GetTransmissionType returns the TransmissionType field if non-nil, zero value otherwise.

### GetTransmissionTypeOk

`func (o *PaymentResponseV4) GetTransmissionTypeOk() (*string, bool)`

GetTransmissionTypeOk returns a tuple with the TransmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionType

`func (o *PaymentResponseV4) SetTransmissionType(v string)`

SetTransmissionType sets TransmissionType field to given value.

### HasTransmissionType

`func (o *PaymentResponseV4) HasTransmissionType() bool`

HasTransmissionType returns a boolean if a field has been set.

### GetPaymentTrackingReference

`func (o *PaymentResponseV4) GetPaymentTrackingReference() string`

GetPaymentTrackingReference returns the PaymentTrackingReference field if non-nil, zero value otherwise.

### GetPaymentTrackingReferenceOk

`func (o *PaymentResponseV4) GetPaymentTrackingReferenceOk() (*string, bool)`

GetPaymentTrackingReferenceOk returns a tuple with the PaymentTrackingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTrackingReference

`func (o *PaymentResponseV4) SetPaymentTrackingReference(v string)`

SetPaymentTrackingReference sets PaymentTrackingReference field to given value.

### HasPaymentTrackingReference

`func (o *PaymentResponseV4) HasPaymentTrackingReference() bool`

HasPaymentTrackingReference returns a boolean if a field has been set.

### GetPaymentMetadata

`func (o *PaymentResponseV4) GetPaymentMetadata() string`

GetPaymentMetadata returns the PaymentMetadata field if non-nil, zero value otherwise.

### GetPaymentMetadataOk

`func (o *PaymentResponseV4) GetPaymentMetadataOk() (*string, bool)`

GetPaymentMetadataOk returns a tuple with the PaymentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMetadata

`func (o *PaymentResponseV4) SetPaymentMetadata(v string)`

SetPaymentMetadata sets PaymentMetadata field to given value.

### HasPaymentMetadata

`func (o *PaymentResponseV4) HasPaymentMetadata() bool`

HasPaymentMetadata returns a boolean if a field has been set.

### GetSchedule

`func (o *PaymentResponseV4) GetSchedule() PayoutSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PaymentResponseV4) GetScheduleOk() (*PayoutSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PaymentResponseV4) SetSchedule(v PayoutSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *PaymentResponseV4) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetPostInstructFxInfo

`func (o *PaymentResponseV4) GetPostInstructFxInfo() PostInstructFxInfo`

GetPostInstructFxInfo returns the PostInstructFxInfo field if non-nil, zero value otherwise.

### GetPostInstructFxInfoOk

`func (o *PaymentResponseV4) GetPostInstructFxInfoOk() (*PostInstructFxInfo, bool)`

GetPostInstructFxInfoOk returns a tuple with the PostInstructFxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostInstructFxInfo

`func (o *PaymentResponseV4) SetPostInstructFxInfo(v PostInstructFxInfo)`

SetPostInstructFxInfo sets PostInstructFxInfo field to given value.

### HasPostInstructFxInfo

`func (o *PaymentResponseV4) HasPostInstructFxInfo() bool`

HasPostInstructFxInfo returns a boolean if a field has been set.

### GetPayout

`func (o *PaymentResponseV4) GetPayout() PaymentResponseV4Payout`

GetPayout returns the Payout field if non-nil, zero value otherwise.

### GetPayoutOk

`func (o *PaymentResponseV4) GetPayoutOk() (*PaymentResponseV4Payout, bool)`

GetPayoutOk returns a tuple with the Payout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayout

`func (o *PaymentResponseV4) SetPayout(v PaymentResponseV4Payout)`

SetPayout sets Payout field to given value.

### HasPayout

`func (o *PaymentResponseV4) HasPayout() bool`

HasPayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


