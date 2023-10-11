# PaymentResponseV3

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
**SourceAmount** | Pointer to **int32** | The source amount for the payment (amount debited to make the payment) | [optional] 
**SourceCurrency** | Pointer to **string** | ISO 3 character currency code | [optional] 
**PaymentAmount** | **int32** | The amount which the payee will receive | 
**PaymentCurrency** | Pointer to **string** | ISO 3 character currency code | [optional] 
**Rate** | Pointer to **float32** | The FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed | [optional] 
**InvertedRate** | Pointer to **float32** | The inverted FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed | [optional] 
**SubmittedDateTime** | **time.Time** |  | 
**Status** | **string** | Current status of the payment. One of the following values: ACCEPTED, AWAITING_FUNDS, FUNDED, UNFUNDED, BANK_PAYMENT_REQUESTED, REJECTED, ACCEPTED_BY_RAILS, CONFIRMED, FAILED, WITHDRAWN | 
**FundingStatus** | **string** | The funding status of the payment. One of the following values: [FUNDED, INSTRUCTED, UNFUNDED | 
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
**Events** | [**[]PaymentEventResponseV3**](PaymentEventResponseV3.md) |  | 
**ReturnCost** | Pointer to **int32** | The return cost if a returned payment. | [optional] 
**ReturnReason** | Pointer to **string** |  | [optional] 
**RailsPaymentId** | Pointer to **string** |  | [optional] 
**RailsBatchId** | Pointer to **string** |  | [optional] 
**PaymentScheme** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentResponseV3

`func NewPaymentResponseV3(paymentId string, payeeId string, payorId string, quoteId string, sourceAccountId string, paymentAmount int32, submittedDateTime time.Time, status string, fundingStatus string, railsId string, events []PaymentEventResponseV3, ) *PaymentResponseV3`

NewPaymentResponseV3 instantiates a new PaymentResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseV3WithDefaults

`func NewPaymentResponseV3WithDefaults() *PaymentResponseV3`

NewPaymentResponseV3WithDefaults instantiates a new PaymentResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentResponseV3) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentResponseV3) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentResponseV3) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayeeId

`func (o *PaymentResponseV3) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PaymentResponseV3) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PaymentResponseV3) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetPayorId

`func (o *PaymentResponseV3) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PaymentResponseV3) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PaymentResponseV3) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetPayorName

`func (o *PaymentResponseV3) GetPayorName() string`

GetPayorName returns the PayorName field if non-nil, zero value otherwise.

### GetPayorNameOk

`func (o *PaymentResponseV3) GetPayorNameOk() (*string, bool)`

GetPayorNameOk returns a tuple with the PayorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorName

`func (o *PaymentResponseV3) SetPayorName(v string)`

SetPayorName sets PayorName field to given value.

### HasPayorName

`func (o *PaymentResponseV3) HasPayorName() bool`

HasPayorName returns a boolean if a field has been set.

### GetQuoteId

`func (o *PaymentResponseV3) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *PaymentResponseV3) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *PaymentResponseV3) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetSourceAccountId

`func (o *PaymentResponseV3) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *PaymentResponseV3) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *PaymentResponseV3) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.


### GetSourceAccountName

`func (o *PaymentResponseV3) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *PaymentResponseV3) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *PaymentResponseV3) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *PaymentResponseV3) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetRemoteId

`func (o *PaymentResponseV3) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PaymentResponseV3) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PaymentResponseV3) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PaymentResponseV3) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetSourceAmount

`func (o *PaymentResponseV3) GetSourceAmount() int32`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *PaymentResponseV3) GetSourceAmountOk() (*int32, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *PaymentResponseV3) SetSourceAmount(v int32)`

SetSourceAmount sets SourceAmount field to given value.

### HasSourceAmount

`func (o *PaymentResponseV3) HasSourceAmount() bool`

HasSourceAmount returns a boolean if a field has been set.

### GetSourceCurrency

`func (o *PaymentResponseV3) GetSourceCurrency() string`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *PaymentResponseV3) GetSourceCurrencyOk() (*string, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *PaymentResponseV3) SetSourceCurrency(v string)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *PaymentResponseV3) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *PaymentResponseV3) GetPaymentAmount() int32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentResponseV3) GetPaymentAmountOk() (*int32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentResponseV3) SetPaymentAmount(v int32)`

SetPaymentAmount sets PaymentAmount field to given value.


### GetPaymentCurrency

`func (o *PaymentResponseV3) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PaymentResponseV3) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PaymentResponseV3) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PaymentResponseV3) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetRate

`func (o *PaymentResponseV3) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *PaymentResponseV3) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *PaymentResponseV3) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *PaymentResponseV3) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetInvertedRate

`func (o *PaymentResponseV3) GetInvertedRate() float32`

GetInvertedRate returns the InvertedRate field if non-nil, zero value otherwise.

### GetInvertedRateOk

`func (o *PaymentResponseV3) GetInvertedRateOk() (*float32, bool)`

GetInvertedRateOk returns a tuple with the InvertedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertedRate

`func (o *PaymentResponseV3) SetInvertedRate(v float32)`

SetInvertedRate sets InvertedRate field to given value.

### HasInvertedRate

`func (o *PaymentResponseV3) HasInvertedRate() bool`

HasInvertedRate returns a boolean if a field has been set.

### GetSubmittedDateTime

`func (o *PaymentResponseV3) GetSubmittedDateTime() time.Time`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *PaymentResponseV3) GetSubmittedDateTimeOk() (*time.Time, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *PaymentResponseV3) SetSubmittedDateTime(v time.Time)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.


### GetStatus

`func (o *PaymentResponseV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentResponseV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentResponseV3) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFundingStatus

`func (o *PaymentResponseV3) GetFundingStatus() string`

GetFundingStatus returns the FundingStatus field if non-nil, zero value otherwise.

### GetFundingStatusOk

`func (o *PaymentResponseV3) GetFundingStatusOk() (*string, bool)`

GetFundingStatusOk returns a tuple with the FundingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingStatus

`func (o *PaymentResponseV3) SetFundingStatus(v string)`

SetFundingStatus sets FundingStatus field to given value.


### GetRoutingNumber

`func (o *PaymentResponseV3) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PaymentResponseV3) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PaymentResponseV3) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *PaymentResponseV3) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PaymentResponseV3) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentResponseV3) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentResponseV3) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PaymentResponseV3) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetIban

`func (o *PaymentResponseV3) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentResponseV3) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentResponseV3) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PaymentResponseV3) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetPaymentMemo

`func (o *PaymentResponseV3) GetPaymentMemo() string`

GetPaymentMemo returns the PaymentMemo field if non-nil, zero value otherwise.

### GetPaymentMemoOk

`func (o *PaymentResponseV3) GetPaymentMemoOk() (*string, bool)`

GetPaymentMemoOk returns a tuple with the PaymentMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMemo

`func (o *PaymentResponseV3) SetPaymentMemo(v string)`

SetPaymentMemo sets PaymentMemo field to given value.

### HasPaymentMemo

`func (o *PaymentResponseV3) HasPaymentMemo() bool`

HasPaymentMemo returns a boolean if a field has been set.

### GetFilenameReference

`func (o *PaymentResponseV3) GetFilenameReference() string`

GetFilenameReference returns the FilenameReference field if non-nil, zero value otherwise.

### GetFilenameReferenceOk

`func (o *PaymentResponseV3) GetFilenameReferenceOk() (*string, bool)`

GetFilenameReferenceOk returns a tuple with the FilenameReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenameReference

`func (o *PaymentResponseV3) SetFilenameReference(v string)`

SetFilenameReference sets FilenameReference field to given value.

### HasFilenameReference

`func (o *PaymentResponseV3) HasFilenameReference() bool`

HasFilenameReference returns a boolean if a field has been set.

### GetIndividualIdentificationNumber

`func (o *PaymentResponseV3) GetIndividualIdentificationNumber() string`

GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field if non-nil, zero value otherwise.

### GetIndividualIdentificationNumberOk

`func (o *PaymentResponseV3) GetIndividualIdentificationNumberOk() (*string, bool)`

GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentificationNumber

`func (o *PaymentResponseV3) SetIndividualIdentificationNumber(v string)`

SetIndividualIdentificationNumber sets IndividualIdentificationNumber field to given value.

### HasIndividualIdentificationNumber

`func (o *PaymentResponseV3) HasIndividualIdentificationNumber() bool`

HasIndividualIdentificationNumber returns a boolean if a field has been set.

### GetTraceNumber

`func (o *PaymentResponseV3) GetTraceNumber() string`

GetTraceNumber returns the TraceNumber field if non-nil, zero value otherwise.

### GetTraceNumberOk

`func (o *PaymentResponseV3) GetTraceNumberOk() (*string, bool)`

GetTraceNumberOk returns a tuple with the TraceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceNumber

`func (o *PaymentResponseV3) SetTraceNumber(v string)`

SetTraceNumber sets TraceNumber field to given value.

### HasTraceNumber

`func (o *PaymentResponseV3) HasTraceNumber() bool`

HasTraceNumber returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *PaymentResponseV3) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentResponseV3) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentResponseV3) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentResponseV3) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetPaymentChannelId

`func (o *PaymentResponseV3) GetPaymentChannelId() string`

GetPaymentChannelId returns the PaymentChannelId field if non-nil, zero value otherwise.

### GetPaymentChannelIdOk

`func (o *PaymentResponseV3) GetPaymentChannelIdOk() (*string, bool)`

GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelId

`func (o *PaymentResponseV3) SetPaymentChannelId(v string)`

SetPaymentChannelId sets PaymentChannelId field to given value.

### HasPaymentChannelId

`func (o *PaymentResponseV3) HasPaymentChannelId() bool`

HasPaymentChannelId returns a boolean if a field has been set.

### GetPaymentChannelName

`func (o *PaymentResponseV3) GetPaymentChannelName() string`

GetPaymentChannelName returns the PaymentChannelName field if non-nil, zero value otherwise.

### GetPaymentChannelNameOk

`func (o *PaymentResponseV3) GetPaymentChannelNameOk() (*string, bool)`

GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelName

`func (o *PaymentResponseV3) SetPaymentChannelName(v string)`

SetPaymentChannelName sets PaymentChannelName field to given value.

### HasPaymentChannelName

`func (o *PaymentResponseV3) HasPaymentChannelName() bool`

HasPaymentChannelName returns a boolean if a field has been set.

### GetAccountName

`func (o *PaymentResponseV3) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentResponseV3) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentResponseV3) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PaymentResponseV3) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetRailsId

`func (o *PaymentResponseV3) GetRailsId() string`

GetRailsId returns the RailsId field if non-nil, zero value otherwise.

### GetRailsIdOk

`func (o *PaymentResponseV3) GetRailsIdOk() (*string, bool)`

GetRailsIdOk returns a tuple with the RailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsId

`func (o *PaymentResponseV3) SetRailsId(v string)`

SetRailsId sets RailsId field to given value.


### GetCountryCode

`func (o *PaymentResponseV3) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PaymentResponseV3) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PaymentResponseV3) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PaymentResponseV3) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEvents

`func (o *PaymentResponseV3) GetEvents() []PaymentEventResponseV3`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PaymentResponseV3) GetEventsOk() (*[]PaymentEventResponseV3, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PaymentResponseV3) SetEvents(v []PaymentEventResponseV3)`

SetEvents sets Events field to given value.


### GetReturnCost

`func (o *PaymentResponseV3) GetReturnCost() int32`

GetReturnCost returns the ReturnCost field if non-nil, zero value otherwise.

### GetReturnCostOk

`func (o *PaymentResponseV3) GetReturnCostOk() (*int32, bool)`

GetReturnCostOk returns a tuple with the ReturnCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCost

`func (o *PaymentResponseV3) SetReturnCost(v int32)`

SetReturnCost sets ReturnCost field to given value.

### HasReturnCost

`func (o *PaymentResponseV3) HasReturnCost() bool`

HasReturnCost returns a boolean if a field has been set.

### GetReturnReason

`func (o *PaymentResponseV3) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *PaymentResponseV3) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *PaymentResponseV3) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *PaymentResponseV3) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetRailsPaymentId

`func (o *PaymentResponseV3) GetRailsPaymentId() string`

GetRailsPaymentId returns the RailsPaymentId field if non-nil, zero value otherwise.

### GetRailsPaymentIdOk

`func (o *PaymentResponseV3) GetRailsPaymentIdOk() (*string, bool)`

GetRailsPaymentIdOk returns a tuple with the RailsPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsPaymentId

`func (o *PaymentResponseV3) SetRailsPaymentId(v string)`

SetRailsPaymentId sets RailsPaymentId field to given value.

### HasRailsPaymentId

`func (o *PaymentResponseV3) HasRailsPaymentId() bool`

HasRailsPaymentId returns a boolean if a field has been set.

### GetRailsBatchId

`func (o *PaymentResponseV3) GetRailsBatchId() string`

GetRailsBatchId returns the RailsBatchId field if non-nil, zero value otherwise.

### GetRailsBatchIdOk

`func (o *PaymentResponseV3) GetRailsBatchIdOk() (*string, bool)`

GetRailsBatchIdOk returns a tuple with the RailsBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsBatchId

`func (o *PaymentResponseV3) SetRailsBatchId(v string)`

SetRailsBatchId sets RailsBatchId field to given value.

### HasRailsBatchId

`func (o *PaymentResponseV3) HasRailsBatchId() bool`

HasRailsBatchId returns a boolean if a field has been set.

### GetPaymentScheme

`func (o *PaymentResponseV3) GetPaymentScheme() string`

GetPaymentScheme returns the PaymentScheme field if non-nil, zero value otherwise.

### GetPaymentSchemeOk

`func (o *PaymentResponseV3) GetPaymentSchemeOk() (*string, bool)`

GetPaymentSchemeOk returns a tuple with the PaymentScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentScheme

`func (o *PaymentResponseV3) SetPaymentScheme(v string)`

SetPaymentScheme sets PaymentScheme field to given value.

### HasPaymentScheme

`func (o *PaymentResponseV3) HasPaymentScheme() bool`

HasPaymentScheme returns a boolean if a field has been set.

### GetRejectionReason

`func (o *PaymentResponseV3) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *PaymentResponseV3) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *PaymentResponseV3) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *PaymentResponseV3) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


