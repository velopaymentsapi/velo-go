/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response.   ## Http Status Codes Following is a list of Http Status codes that could be returned by the platform      | Status Code            | Description                                                                          |     | -----------------------| -------------------------------------------------------------------------------------|     | 200 OK                 | The request was successfully processed and usually returns a json response           |     | 201 Created            | A resource was created and a Location header is returned linking to the new resource |     | 202 Accepted           | The request has been accepted for processing                                         |     | 204 No Content         | The request has been processed and there is no response (usually deletes and updates)|     | 400 Bad Request        | The request is invalid and should be fixed before retrying                           |     | 401 Unauthorized       | Authentication has failed, usually means the token has expired                       |     | 403 Forbidden          | The user does not have permissions for the request                                   |     | 404 Not Found          | The resource was not found                                                           |     | 409 Conflict           | The resource already exists and there is a conflict                                  |     | 429 Too Many Requests  | The user has submitted too many requests in a given amount of time                   |     | 5xx Server Error       | Platform internal error (should rarely happen)                                       | 

API version: 2.37.150
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
	"time"
)

// checks if the PaymentResponseV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentResponseV3{}

// PaymentResponseV3 struct for PaymentResponseV3
type PaymentResponseV3 struct {
	// The id of the payment
	PaymentId string `json:"paymentId"`
	// The id of the paymeee
	PayeeId string `json:"payeeId"`
	// The id of the payor
	PayorId string `json:"payorId"`
	// The name of the payor
	PayorName *string `json:"payorName,omitempty"`
	// The quote Id used for the FX
	QuoteId string `json:"quoteId"`
	// The id of the source account from which the payment was taken
	SourceAccountId string `json:"sourceAccountId"`
	// The name of the source account from which the payment was taken
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	// The remote id by which the payor refers to the payee. Only populated once payment is confirmed
	RemoteId *string `json:"remoteId,omitempty"`
	// The source amount for the payment (amount debited to make the payment)
	SourceAmount *int32 `json:"sourceAmount,omitempty"`
	// ISO 3 character currency code
	SourceCurrency *string `json:"sourceCurrency,omitempty"`
	// The amount which the payee will receive
	PaymentAmount int32 `json:"paymentAmount"`
	// ISO 3 character currency code
	PaymentCurrency *string `json:"paymentCurrency,omitempty"`
	// The FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed
	Rate *float32 `json:"rate,omitempty"`
	// The inverted FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed
	InvertedRate *float32 `json:"invertedRate,omitempty"`
	SubmittedDateTime time.Time `json:"submittedDateTime"`
	// Current status of the payment. One of the following values: ACCEPTED, AWAITING_FUNDS, FUNDED, UNFUNDED, BANK_PAYMENT_REQUESTED, REJECTED, ACCEPTED_BY_RAILS, CONFIRMED, FAILED, WITHDRAWN
	Status string `json:"status"`
	// The funding status of the payment. One of the following values: [FUNDED, INSTRUCTED, UNFUNDED
	FundingStatus string `json:"fundingStatus"`
	// The routing number for the payment.
	RoutingNumber *string `json:"routingNumber,omitempty"`
	// The account number for the account which will receive the payment.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The iban for the payment.
	Iban *string `json:"iban,omitempty"`
	// The payment memo set by the payor
	PaymentMemo *string `json:"paymentMemo,omitempty"`
	// ACH file payment was submitted in, if applicable
	FilenameReference *string `json:"filenameReference,omitempty"`
	// Individual Identification Number assigned to the payment in the ACH file, if applicable
	IndividualIdentificationNumber *string `json:"individualIdentificationNumber,omitempty"`
	// Trace Number assigned to the payment in the ACH file, if applicable
	TraceNumber *string `json:"traceNumber,omitempty"`
	PayorPaymentId *string `json:"payorPaymentId,omitempty"`
	PaymentChannelId *string `json:"paymentChannelId,omitempty"`
	PaymentChannelName *string `json:"paymentChannelName,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	// The rails ID. Default value is RAILS ID UNAVAILABLE when not populated.
	RailsId string `json:"railsId"`
	// The country code of the payment channel.
	CountryCode *string `json:"countryCode,omitempty"`
	Events []PaymentEventResponseV3 `json:"events"`
	// The return cost if a returned payment.
	ReturnCost *int32 `json:"returnCost,omitempty"`
	ReturnReason *string `json:"returnReason,omitempty"`
	RailsPaymentId *string `json:"railsPaymentId,omitempty"`
	RailsBatchId *string `json:"railsBatchId,omitempty"`
	PaymentScheme *string `json:"paymentScheme,omitempty"`
	RejectionReason *string `json:"rejectionReason,omitempty"`
}

// NewPaymentResponseV3 instantiates a new PaymentResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentResponseV3(paymentId string, payeeId string, payorId string, quoteId string, sourceAccountId string, paymentAmount int32, submittedDateTime time.Time, status string, fundingStatus string, railsId string, events []PaymentEventResponseV3) *PaymentResponseV3 {
	this := PaymentResponseV3{}
	this.PaymentId = paymentId
	this.PayeeId = payeeId
	this.PayorId = payorId
	this.QuoteId = quoteId
	this.SourceAccountId = sourceAccountId
	this.PaymentAmount = paymentAmount
	this.SubmittedDateTime = submittedDateTime
	this.Status = status
	this.FundingStatus = fundingStatus
	this.RailsId = railsId
	this.Events = events
	return &this
}

// NewPaymentResponseV3WithDefaults instantiates a new PaymentResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentResponseV3WithDefaults() *PaymentResponseV3 {
	this := PaymentResponseV3{}
	var railsId string = "RAILS ID UNAVAILABLE"
	this.RailsId = railsId
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentResponseV3) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentResponseV3) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetPayeeId returns the PayeeId field value
func (o *PaymentResponseV3) GetPayeeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPayeeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayeeId, true
}

// SetPayeeId sets field value
func (o *PaymentResponseV3) SetPayeeId(v string) {
	o.PayeeId = v
}

// GetPayorId returns the PayorId field value
func (o *PaymentResponseV3) GetPayorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPayorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayorId, true
}

// SetPayorId sets field value
func (o *PaymentResponseV3) SetPayorId(v string) {
	o.PayorId = v
}

// GetPayorName returns the PayorName field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetPayorName() string {
	if o == nil || IsNil(o.PayorName) {
		var ret string
		return ret
	}
	return *o.PayorName
}

// GetPayorNameOk returns a tuple with the PayorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPayorNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayorName) {
		return nil, false
	}
	return o.PayorName, true
}

// HasPayorName returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasPayorName() bool {
	if o != nil && !IsNil(o.PayorName) {
		return true
	}

	return false
}

// SetPayorName gets a reference to the given string and assigns it to the PayorName field.
func (o *PaymentResponseV3) SetPayorName(v string) {
	o.PayorName = &v
}

// GetQuoteId returns the QuoteId field value
func (o *PaymentResponseV3) GetQuoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetQuoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteId, true
}

// SetQuoteId sets field value
func (o *PaymentResponseV3) SetQuoteId(v string) {
	o.QuoteId = v
}

// GetSourceAccountId returns the SourceAccountId field value
func (o *PaymentResponseV3) GetSourceAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAccountId
}

// GetSourceAccountIdOk returns a tuple with the SourceAccountId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetSourceAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAccountId, true
}

// SetSourceAccountId sets field value
func (o *PaymentResponseV3) SetSourceAccountId(v string) {
	o.SourceAccountId = v
}

// GetSourceAccountName returns the SourceAccountName field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetSourceAccountName() string {
	if o == nil || IsNil(o.SourceAccountName) {
		var ret string
		return ret
	}
	return *o.SourceAccountName
}

// GetSourceAccountNameOk returns a tuple with the SourceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetSourceAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAccountName) {
		return nil, false
	}
	return o.SourceAccountName, true
}

// HasSourceAccountName returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasSourceAccountName() bool {
	if o != nil && !IsNil(o.SourceAccountName) {
		return true
	}

	return false
}

// SetSourceAccountName gets a reference to the given string and assigns it to the SourceAccountName field.
func (o *PaymentResponseV3) SetSourceAccountName(v string) {
	o.SourceAccountName = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetRemoteId() string {
	if o == nil || IsNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteId) {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasRemoteId() bool {
	if o != nil && !IsNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *PaymentResponseV3) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetSourceAmount returns the SourceAmount field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetSourceAmount() int32 {
	if o == nil || IsNil(o.SourceAmount) {
		var ret int32
		return ret
	}
	return *o.SourceAmount
}

// GetSourceAmountOk returns a tuple with the SourceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetSourceAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceAmount) {
		return nil, false
	}
	return o.SourceAmount, true
}

// HasSourceAmount returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasSourceAmount() bool {
	if o != nil && !IsNil(o.SourceAmount) {
		return true
	}

	return false
}

// SetSourceAmount gets a reference to the given int32 and assigns it to the SourceAmount field.
func (o *PaymentResponseV3) SetSourceAmount(v int32) {
	o.SourceAmount = &v
}

// GetSourceCurrency returns the SourceCurrency field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetSourceCurrency() string {
	if o == nil || IsNil(o.SourceCurrency) {
		var ret string
		return ret
	}
	return *o.SourceCurrency
}

// GetSourceCurrencyOk returns a tuple with the SourceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetSourceCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCurrency) {
		return nil, false
	}
	return o.SourceCurrency, true
}

// HasSourceCurrency returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasSourceCurrency() bool {
	if o != nil && !IsNil(o.SourceCurrency) {
		return true
	}

	return false
}

// SetSourceCurrency gets a reference to the given string and assigns it to the SourceCurrency field.
func (o *PaymentResponseV3) SetSourceCurrency(v string) {
	o.SourceCurrency = &v
}

// GetPaymentAmount returns the PaymentAmount field value
func (o *PaymentResponseV3) GetPaymentAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PaymentAmount
}

// GetPaymentAmountOk returns a tuple with the PaymentAmount field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPaymentAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentAmount, true
}

// SetPaymentAmount sets field value
func (o *PaymentResponseV3) SetPaymentAmount(v int32) {
	o.PaymentAmount = v
}

// GetPaymentCurrency returns the PaymentCurrency field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetPaymentCurrency() string {
	if o == nil || IsNil(o.PaymentCurrency) {
		var ret string
		return ret
	}
	return *o.PaymentCurrency
}

// GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPaymentCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentCurrency) {
		return nil, false
	}
	return o.PaymentCurrency, true
}

// HasPaymentCurrency returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasPaymentCurrency() bool {
	if o != nil && !IsNil(o.PaymentCurrency) {
		return true
	}

	return false
}

// SetPaymentCurrency gets a reference to the given string and assigns it to the PaymentCurrency field.
func (o *PaymentResponseV3) SetPaymentCurrency(v string) {
	o.PaymentCurrency = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *PaymentResponseV3) SetRate(v float32) {
	o.Rate = &v
}

// GetInvertedRate returns the InvertedRate field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetInvertedRate() float32 {
	if o == nil || IsNil(o.InvertedRate) {
		var ret float32
		return ret
	}
	return *o.InvertedRate
}

// GetInvertedRateOk returns a tuple with the InvertedRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetInvertedRateOk() (*float32, bool) {
	if o == nil || IsNil(o.InvertedRate) {
		return nil, false
	}
	return o.InvertedRate, true
}

// HasInvertedRate returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasInvertedRate() bool {
	if o != nil && !IsNil(o.InvertedRate) {
		return true
	}

	return false
}

// SetInvertedRate gets a reference to the given float32 and assigns it to the InvertedRate field.
func (o *PaymentResponseV3) SetInvertedRate(v float32) {
	o.InvertedRate = &v
}

// GetSubmittedDateTime returns the SubmittedDateTime field value
func (o *PaymentResponseV3) GetSubmittedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SubmittedDateTime
}

// GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetSubmittedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubmittedDateTime, true
}

// SetSubmittedDateTime sets field value
func (o *PaymentResponseV3) SetSubmittedDateTime(v time.Time) {
	o.SubmittedDateTime = v
}

// GetStatus returns the Status field value
func (o *PaymentResponseV3) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentResponseV3) SetStatus(v string) {
	o.Status = v
}

// GetFundingStatus returns the FundingStatus field value
func (o *PaymentResponseV3) GetFundingStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingStatus
}

// GetFundingStatusOk returns a tuple with the FundingStatus field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetFundingStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingStatus, true
}

// SetFundingStatus sets field value
func (o *PaymentResponseV3) SetFundingStatus(v string) {
	o.FundingStatus = v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetRoutingNumber() string {
	if o == nil || IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasRoutingNumber() bool {
	if o != nil && !IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *PaymentResponseV3) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *PaymentResponseV3) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetIban() string {
	if o == nil || IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetIbanOk() (*string, bool) {
	if o == nil || IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasIban() bool {
	if o != nil && !IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *PaymentResponseV3) SetIban(v string) {
	o.Iban = &v
}

// GetPaymentMemo returns the PaymentMemo field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetPaymentMemo() string {
	if o == nil || IsNil(o.PaymentMemo) {
		var ret string
		return ret
	}
	return *o.PaymentMemo
}

// GetPaymentMemoOk returns a tuple with the PaymentMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPaymentMemoOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMemo) {
		return nil, false
	}
	return o.PaymentMemo, true
}

// HasPaymentMemo returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasPaymentMemo() bool {
	if o != nil && !IsNil(o.PaymentMemo) {
		return true
	}

	return false
}

// SetPaymentMemo gets a reference to the given string and assigns it to the PaymentMemo field.
func (o *PaymentResponseV3) SetPaymentMemo(v string) {
	o.PaymentMemo = &v
}

// GetFilenameReference returns the FilenameReference field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetFilenameReference() string {
	if o == nil || IsNil(o.FilenameReference) {
		var ret string
		return ret
	}
	return *o.FilenameReference
}

// GetFilenameReferenceOk returns a tuple with the FilenameReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetFilenameReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.FilenameReference) {
		return nil, false
	}
	return o.FilenameReference, true
}

// HasFilenameReference returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasFilenameReference() bool {
	if o != nil && !IsNil(o.FilenameReference) {
		return true
	}

	return false
}

// SetFilenameReference gets a reference to the given string and assigns it to the FilenameReference field.
func (o *PaymentResponseV3) SetFilenameReference(v string) {
	o.FilenameReference = &v
}

// GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetIndividualIdentificationNumber() string {
	if o == nil || IsNil(o.IndividualIdentificationNumber) {
		var ret string
		return ret
	}
	return *o.IndividualIdentificationNumber
}

// GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetIndividualIdentificationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualIdentificationNumber) {
		return nil, false
	}
	return o.IndividualIdentificationNumber, true
}

// HasIndividualIdentificationNumber returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasIndividualIdentificationNumber() bool {
	if o != nil && !IsNil(o.IndividualIdentificationNumber) {
		return true
	}

	return false
}

// SetIndividualIdentificationNumber gets a reference to the given string and assigns it to the IndividualIdentificationNumber field.
func (o *PaymentResponseV3) SetIndividualIdentificationNumber(v string) {
	o.IndividualIdentificationNumber = &v
}

// GetTraceNumber returns the TraceNumber field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetTraceNumber() string {
	if o == nil || IsNil(o.TraceNumber) {
		var ret string
		return ret
	}
	return *o.TraceNumber
}

// GetTraceNumberOk returns a tuple with the TraceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetTraceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TraceNumber) {
		return nil, false
	}
	return o.TraceNumber, true
}

// HasTraceNumber returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasTraceNumber() bool {
	if o != nil && !IsNil(o.TraceNumber) {
		return true
	}

	return false
}

// SetTraceNumber gets a reference to the given string and assigns it to the TraceNumber field.
func (o *PaymentResponseV3) SetTraceNumber(v string) {
	o.TraceNumber = &v
}

// GetPayorPaymentId returns the PayorPaymentId field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetPayorPaymentId() string {
	if o == nil || IsNil(o.PayorPaymentId) {
		var ret string
		return ret
	}
	return *o.PayorPaymentId
}

// GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPayorPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayorPaymentId) {
		return nil, false
	}
	return o.PayorPaymentId, true
}

// HasPayorPaymentId returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasPayorPaymentId() bool {
	if o != nil && !IsNil(o.PayorPaymentId) {
		return true
	}

	return false
}

// SetPayorPaymentId gets a reference to the given string and assigns it to the PayorPaymentId field.
func (o *PaymentResponseV3) SetPayorPaymentId(v string) {
	o.PayorPaymentId = &v
}

// GetPaymentChannelId returns the PaymentChannelId field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetPaymentChannelId() string {
	if o == nil || IsNil(o.PaymentChannelId) {
		var ret string
		return ret
	}
	return *o.PaymentChannelId
}

// GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPaymentChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelId) {
		return nil, false
	}
	return o.PaymentChannelId, true
}

// HasPaymentChannelId returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasPaymentChannelId() bool {
	if o != nil && !IsNil(o.PaymentChannelId) {
		return true
	}

	return false
}

// SetPaymentChannelId gets a reference to the given string and assigns it to the PaymentChannelId field.
func (o *PaymentResponseV3) SetPaymentChannelId(v string) {
	o.PaymentChannelId = &v
}

// GetPaymentChannelName returns the PaymentChannelName field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetPaymentChannelName() string {
	if o == nil || IsNil(o.PaymentChannelName) {
		var ret string
		return ret
	}
	return *o.PaymentChannelName
}

// GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPaymentChannelNameOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelName) {
		return nil, false
	}
	return o.PaymentChannelName, true
}

// HasPaymentChannelName returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasPaymentChannelName() bool {
	if o != nil && !IsNil(o.PaymentChannelName) {
		return true
	}

	return false
}

// SetPaymentChannelName gets a reference to the given string and assigns it to the PaymentChannelName field.
func (o *PaymentResponseV3) SetPaymentChannelName(v string) {
	o.PaymentChannelName = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *PaymentResponseV3) SetAccountName(v string) {
	o.AccountName = &v
}

// GetRailsId returns the RailsId field value
func (o *PaymentResponseV3) GetRailsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RailsId
}

// GetRailsIdOk returns a tuple with the RailsId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetRailsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RailsId, true
}

// SetRailsId sets field value
func (o *PaymentResponseV3) SetRailsId(v string) {
	o.RailsId = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PaymentResponseV3) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEvents returns the Events field value
func (o *PaymentResponseV3) GetEvents() []PaymentEventResponseV3 {
	if o == nil {
		var ret []PaymentEventResponseV3
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetEventsOk() ([]PaymentEventResponseV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *PaymentResponseV3) SetEvents(v []PaymentEventResponseV3) {
	o.Events = v
}

// GetReturnCost returns the ReturnCost field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetReturnCost() int32 {
	if o == nil || IsNil(o.ReturnCost) {
		var ret int32
		return ret
	}
	return *o.ReturnCost
}

// GetReturnCostOk returns a tuple with the ReturnCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetReturnCostOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnCost) {
		return nil, false
	}
	return o.ReturnCost, true
}

// HasReturnCost returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasReturnCost() bool {
	if o != nil && !IsNil(o.ReturnCost) {
		return true
	}

	return false
}

// SetReturnCost gets a reference to the given int32 and assigns it to the ReturnCost field.
func (o *PaymentResponseV3) SetReturnCost(v int32) {
	o.ReturnCost = &v
}

// GetReturnReason returns the ReturnReason field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetReturnReason() string {
	if o == nil || IsNil(o.ReturnReason) {
		var ret string
		return ret
	}
	return *o.ReturnReason
}

// GetReturnReasonOk returns a tuple with the ReturnReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetReturnReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnReason) {
		return nil, false
	}
	return o.ReturnReason, true
}

// HasReturnReason returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasReturnReason() bool {
	if o != nil && !IsNil(o.ReturnReason) {
		return true
	}

	return false
}

// SetReturnReason gets a reference to the given string and assigns it to the ReturnReason field.
func (o *PaymentResponseV3) SetReturnReason(v string) {
	o.ReturnReason = &v
}

// GetRailsPaymentId returns the RailsPaymentId field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetRailsPaymentId() string {
	if o == nil || IsNil(o.RailsPaymentId) {
		var ret string
		return ret
	}
	return *o.RailsPaymentId
}

// GetRailsPaymentIdOk returns a tuple with the RailsPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetRailsPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.RailsPaymentId) {
		return nil, false
	}
	return o.RailsPaymentId, true
}

// HasRailsPaymentId returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasRailsPaymentId() bool {
	if o != nil && !IsNil(o.RailsPaymentId) {
		return true
	}

	return false
}

// SetRailsPaymentId gets a reference to the given string and assigns it to the RailsPaymentId field.
func (o *PaymentResponseV3) SetRailsPaymentId(v string) {
	o.RailsPaymentId = &v
}

// GetRailsBatchId returns the RailsBatchId field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetRailsBatchId() string {
	if o == nil || IsNil(o.RailsBatchId) {
		var ret string
		return ret
	}
	return *o.RailsBatchId
}

// GetRailsBatchIdOk returns a tuple with the RailsBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetRailsBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.RailsBatchId) {
		return nil, false
	}
	return o.RailsBatchId, true
}

// HasRailsBatchId returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasRailsBatchId() bool {
	if o != nil && !IsNil(o.RailsBatchId) {
		return true
	}

	return false
}

// SetRailsBatchId gets a reference to the given string and assigns it to the RailsBatchId field.
func (o *PaymentResponseV3) SetRailsBatchId(v string) {
	o.RailsBatchId = &v
}

// GetPaymentScheme returns the PaymentScheme field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetPaymentScheme() string {
	if o == nil || IsNil(o.PaymentScheme) {
		var ret string
		return ret
	}
	return *o.PaymentScheme
}

// GetPaymentSchemeOk returns a tuple with the PaymentScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetPaymentSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentScheme) {
		return nil, false
	}
	return o.PaymentScheme, true
}

// HasPaymentScheme returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasPaymentScheme() bool {
	if o != nil && !IsNil(o.PaymentScheme) {
		return true
	}

	return false
}

// SetPaymentScheme gets a reference to the given string and assigns it to the PaymentScheme field.
func (o *PaymentResponseV3) SetPaymentScheme(v string) {
	o.PaymentScheme = &v
}

// GetRejectionReason returns the RejectionReason field value if set, zero value otherwise.
func (o *PaymentResponseV3) GetRejectionReason() string {
	if o == nil || IsNil(o.RejectionReason) {
		var ret string
		return ret
	}
	return *o.RejectionReason
}

// GetRejectionReasonOk returns a tuple with the RejectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV3) GetRejectionReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RejectionReason) {
		return nil, false
	}
	return o.RejectionReason, true
}

// HasRejectionReason returns a boolean if a field has been set.
func (o *PaymentResponseV3) HasRejectionReason() bool {
	if o != nil && !IsNil(o.RejectionReason) {
		return true
	}

	return false
}

// SetRejectionReason gets a reference to the given string and assigns it to the RejectionReason field.
func (o *PaymentResponseV3) SetRejectionReason(v string) {
	o.RejectionReason = &v
}

func (o PaymentResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentResponseV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentId"] = o.PaymentId
	toSerialize["payeeId"] = o.PayeeId
	toSerialize["payorId"] = o.PayorId
	if !IsNil(o.PayorName) {
		toSerialize["payorName"] = o.PayorName
	}
	toSerialize["quoteId"] = o.QuoteId
	toSerialize["sourceAccountId"] = o.SourceAccountId
	if !IsNil(o.SourceAccountName) {
		toSerialize["sourceAccountName"] = o.SourceAccountName
	}
	if !IsNil(o.RemoteId) {
		toSerialize["remoteId"] = o.RemoteId
	}
	if !IsNil(o.SourceAmount) {
		toSerialize["sourceAmount"] = o.SourceAmount
	}
	if !IsNil(o.SourceCurrency) {
		toSerialize["sourceCurrency"] = o.SourceCurrency
	}
	toSerialize["paymentAmount"] = o.PaymentAmount
	if !IsNil(o.PaymentCurrency) {
		toSerialize["paymentCurrency"] = o.PaymentCurrency
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.InvertedRate) {
		toSerialize["invertedRate"] = o.InvertedRate
	}
	toSerialize["submittedDateTime"] = o.SubmittedDateTime
	toSerialize["status"] = o.Status
	toSerialize["fundingStatus"] = o.FundingStatus
	if !IsNil(o.RoutingNumber) {
		toSerialize["routingNumber"] = o.RoutingNumber
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	if !IsNil(o.PaymentMemo) {
		toSerialize["paymentMemo"] = o.PaymentMemo
	}
	if !IsNil(o.FilenameReference) {
		toSerialize["filenameReference"] = o.FilenameReference
	}
	if !IsNil(o.IndividualIdentificationNumber) {
		toSerialize["individualIdentificationNumber"] = o.IndividualIdentificationNumber
	}
	if !IsNil(o.TraceNumber) {
		toSerialize["traceNumber"] = o.TraceNumber
	}
	if !IsNil(o.PayorPaymentId) {
		toSerialize["payorPaymentId"] = o.PayorPaymentId
	}
	if !IsNil(o.PaymentChannelId) {
		toSerialize["paymentChannelId"] = o.PaymentChannelId
	}
	if !IsNil(o.PaymentChannelName) {
		toSerialize["paymentChannelName"] = o.PaymentChannelName
	}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	toSerialize["railsId"] = o.RailsId
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	toSerialize["events"] = o.Events
	if !IsNil(o.ReturnCost) {
		toSerialize["returnCost"] = o.ReturnCost
	}
	if !IsNil(o.ReturnReason) {
		toSerialize["returnReason"] = o.ReturnReason
	}
	if !IsNil(o.RailsPaymentId) {
		toSerialize["railsPaymentId"] = o.RailsPaymentId
	}
	if !IsNil(o.RailsBatchId) {
		toSerialize["railsBatchId"] = o.RailsBatchId
	}
	if !IsNil(o.PaymentScheme) {
		toSerialize["paymentScheme"] = o.PaymentScheme
	}
	if !IsNil(o.RejectionReason) {
		toSerialize["rejectionReason"] = o.RejectionReason
	}
	return toSerialize, nil
}

type NullablePaymentResponseV3 struct {
	value *PaymentResponseV3
	isSet bool
}

func (v NullablePaymentResponseV3) Get() *PaymentResponseV3 {
	return v.value
}

func (v *NullablePaymentResponseV3) Set(val *PaymentResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentResponseV3(val *PaymentResponseV3) *NullablePaymentResponseV3 {
	return &NullablePaymentResponseV3{value: val, isSet: true}
}

func (v NullablePaymentResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


