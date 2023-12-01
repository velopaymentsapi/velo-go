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

// checks if the PaymentResponseV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentResponseV4{}

// PaymentResponseV4 struct for PaymentResponseV4
type PaymentResponseV4 struct {
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
	// The velo id of the remote system orchestrating the payment. Not populated for normal Velo payments.
	RemoteSystemId *string `json:"remoteSystemId,omitempty"`
	// The id of the payment in the remote system. Not populated for normal Velo payments.
	RemoteSystemPaymentId *string `json:"remoteSystemPaymentId,omitempty"`
	// The source amount for the payment (amount debited to make the payment)
	SourceAmount *int32 `json:"sourceAmount,omitempty"`
	// ISO-4217 3 character currency code
	SourceCurrency *string `json:"sourceCurrency,omitempty"`
	// The amount which the payee will receive
	PaymentAmount int32 `json:"paymentAmount"`
	// ISO-4217 3 character currency code
	PaymentCurrency *string `json:"paymentCurrency,omitempty"`
	// The FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed
	Rate *float64 `json:"rate,omitempty"`
	// The inverted FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed
	InvertedRate *float64 `json:"invertedRate,omitempty"`
	IsPaymentCcyBaseCcy *bool `json:"isPaymentCcyBaseCcy,omitempty"`
	SubmittedDateTime time.Time `json:"submittedDateTime"`
	// Current status of the payment. One of the following values: ACCEPTED, AWAITING_FUNDS, FUNDED, UNFUNDED, BANK_PAYMENT_REQUESTED, REJECTED, ACCEPTED_BY_RAILS, CONFIRMED, RETURNED, WITHDRAWN
	Status string `json:"status"`
	// Current funding status of the payment. One of the following values: FUNDED, INSTRUCTED, UNFUNDED
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
	// The country code of the payee's address.
	PayeeAddressCountryCode *string `json:"payeeAddressCountryCode,omitempty"`
	Events []PaymentEventResponse `json:"events"`
	// The return cost if a returned payment.
	ReturnCost *int32 `json:"returnCost,omitempty"`
	ReturnReason *string `json:"returnReason,omitempty"`
	RailsPaymentId *string `json:"railsPaymentId,omitempty"`
	RailsBatchId *string `json:"railsBatchId,omitempty"`
	RailsAccountId *string `json:"railsAccountId,omitempty"`
	PaymentScheme *string `json:"paymentScheme,omitempty"`
	RejectionReason *string `json:"rejectionReason,omitempty"`
	// The original reason that the payment was rejected. This can be third party rails specific if rejected by the underlying third party rails logic.
	RailsRejectionInformation *string `json:"railsRejectionInformation,omitempty"`
	WithdrawnReason *string `json:"withdrawnReason,omitempty"`
	Withdrawable *bool `json:"withdrawable,omitempty"`
	// Populated with rejection reason code if the payment was withdrawn automatically at instruct time
	AutoWithdrawnReasonCode *string `json:"autoWithdrawnReasonCode,omitempty"`
	// The transmission type of the payment, e.g. ACH, SAME_DAY_ACH, WIRE, GACHO
	TransmissionType *string `json:"transmissionType,omitempty"`
	// The transmission type of the payment requested by the payor
	TransmissionTypeRequested *string `json:"transmissionTypeRequested,omitempty"`
	PaymentTrackingReference *string `json:"paymentTrackingReference,omitempty"`
	// Metadata for the payment
	PaymentMetadata *string `json:"paymentMetadata,omitempty"`
	TransactionId *string `json:"transactionId,omitempty"`
	TransactionReference *string `json:"transactionReference,omitempty"`
	Schedule *PayoutSchedule `json:"schedule,omitempty"`
	PostInstructFxInfo *PostInstructFxInfo `json:"postInstructFxInfo,omitempty"`
	Payout *PaymentResponseV4Payout `json:"payout,omitempty"`
}

// NewPaymentResponseV4 instantiates a new PaymentResponseV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentResponseV4(paymentId string, payeeId string, payorId string, quoteId string, sourceAccountId string, paymentAmount int32, submittedDateTime time.Time, status string, fundingStatus string, railsId string, events []PaymentEventResponse) *PaymentResponseV4 {
	this := PaymentResponseV4{}
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

// NewPaymentResponseV4WithDefaults instantiates a new PaymentResponseV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentResponseV4WithDefaults() *PaymentResponseV4 {
	this := PaymentResponseV4{}
	var railsId string = "RAILS ID UNAVAILABLE"
	this.RailsId = railsId
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentResponseV4) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentResponseV4) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetPayeeId returns the PayeeId field value
func (o *PaymentResponseV4) GetPayeeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPayeeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayeeId, true
}

// SetPayeeId sets field value
func (o *PaymentResponseV4) SetPayeeId(v string) {
	o.PayeeId = v
}

// GetPayorId returns the PayorId field value
func (o *PaymentResponseV4) GetPayorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPayorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayorId, true
}

// SetPayorId sets field value
func (o *PaymentResponseV4) SetPayorId(v string) {
	o.PayorId = v
}

// GetPayorName returns the PayorName field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPayorName() string {
	if o == nil || IsNil(o.PayorName) {
		var ret string
		return ret
	}
	return *o.PayorName
}

// GetPayorNameOk returns a tuple with the PayorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPayorNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayorName) {
		return nil, false
	}
	return o.PayorName, true
}

// HasPayorName returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPayorName() bool {
	if o != nil && !IsNil(o.PayorName) {
		return true
	}

	return false
}

// SetPayorName gets a reference to the given string and assigns it to the PayorName field.
func (o *PaymentResponseV4) SetPayorName(v string) {
	o.PayorName = &v
}

// GetQuoteId returns the QuoteId field value
func (o *PaymentResponseV4) GetQuoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetQuoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteId, true
}

// SetQuoteId sets field value
func (o *PaymentResponseV4) SetQuoteId(v string) {
	o.QuoteId = v
}

// GetSourceAccountId returns the SourceAccountId field value
func (o *PaymentResponseV4) GetSourceAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAccountId
}

// GetSourceAccountIdOk returns a tuple with the SourceAccountId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetSourceAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAccountId, true
}

// SetSourceAccountId sets field value
func (o *PaymentResponseV4) SetSourceAccountId(v string) {
	o.SourceAccountId = v
}

// GetSourceAccountName returns the SourceAccountName field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetSourceAccountName() string {
	if o == nil || IsNil(o.SourceAccountName) {
		var ret string
		return ret
	}
	return *o.SourceAccountName
}

// GetSourceAccountNameOk returns a tuple with the SourceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetSourceAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAccountName) {
		return nil, false
	}
	return o.SourceAccountName, true
}

// HasSourceAccountName returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasSourceAccountName() bool {
	if o != nil && !IsNil(o.SourceAccountName) {
		return true
	}

	return false
}

// SetSourceAccountName gets a reference to the given string and assigns it to the SourceAccountName field.
func (o *PaymentResponseV4) SetSourceAccountName(v string) {
	o.SourceAccountName = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRemoteId() string {
	if o == nil || IsNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteId) {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRemoteId() bool {
	if o != nil && !IsNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *PaymentResponseV4) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetRemoteSystemId returns the RemoteSystemId field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRemoteSystemId() string {
	if o == nil || IsNil(o.RemoteSystemId) {
		var ret string
		return ret
	}
	return *o.RemoteSystemId
}

// GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRemoteSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteSystemId) {
		return nil, false
	}
	return o.RemoteSystemId, true
}

// HasRemoteSystemId returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRemoteSystemId() bool {
	if o != nil && !IsNil(o.RemoteSystemId) {
		return true
	}

	return false
}

// SetRemoteSystemId gets a reference to the given string and assigns it to the RemoteSystemId field.
func (o *PaymentResponseV4) SetRemoteSystemId(v string) {
	o.RemoteSystemId = &v
}

// GetRemoteSystemPaymentId returns the RemoteSystemPaymentId field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRemoteSystemPaymentId() string {
	if o == nil || IsNil(o.RemoteSystemPaymentId) {
		var ret string
		return ret
	}
	return *o.RemoteSystemPaymentId
}

// GetRemoteSystemPaymentIdOk returns a tuple with the RemoteSystemPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRemoteSystemPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteSystemPaymentId) {
		return nil, false
	}
	return o.RemoteSystemPaymentId, true
}

// HasRemoteSystemPaymentId returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRemoteSystemPaymentId() bool {
	if o != nil && !IsNil(o.RemoteSystemPaymentId) {
		return true
	}

	return false
}

// SetRemoteSystemPaymentId gets a reference to the given string and assigns it to the RemoteSystemPaymentId field.
func (o *PaymentResponseV4) SetRemoteSystemPaymentId(v string) {
	o.RemoteSystemPaymentId = &v
}

// GetSourceAmount returns the SourceAmount field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetSourceAmount() int32 {
	if o == nil || IsNil(o.SourceAmount) {
		var ret int32
		return ret
	}
	return *o.SourceAmount
}

// GetSourceAmountOk returns a tuple with the SourceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetSourceAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceAmount) {
		return nil, false
	}
	return o.SourceAmount, true
}

// HasSourceAmount returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasSourceAmount() bool {
	if o != nil && !IsNil(o.SourceAmount) {
		return true
	}

	return false
}

// SetSourceAmount gets a reference to the given int32 and assigns it to the SourceAmount field.
func (o *PaymentResponseV4) SetSourceAmount(v int32) {
	o.SourceAmount = &v
}

// GetSourceCurrency returns the SourceCurrency field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetSourceCurrency() string {
	if o == nil || IsNil(o.SourceCurrency) {
		var ret string
		return ret
	}
	return *o.SourceCurrency
}

// GetSourceCurrencyOk returns a tuple with the SourceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetSourceCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCurrency) {
		return nil, false
	}
	return o.SourceCurrency, true
}

// HasSourceCurrency returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasSourceCurrency() bool {
	if o != nil && !IsNil(o.SourceCurrency) {
		return true
	}

	return false
}

// SetSourceCurrency gets a reference to the given string and assigns it to the SourceCurrency field.
func (o *PaymentResponseV4) SetSourceCurrency(v string) {
	o.SourceCurrency = &v
}

// GetPaymentAmount returns the PaymentAmount field value
func (o *PaymentResponseV4) GetPaymentAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PaymentAmount
}

// GetPaymentAmountOk returns a tuple with the PaymentAmount field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPaymentAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentAmount, true
}

// SetPaymentAmount sets field value
func (o *PaymentResponseV4) SetPaymentAmount(v int32) {
	o.PaymentAmount = v
}

// GetPaymentCurrency returns the PaymentCurrency field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPaymentCurrency() string {
	if o == nil || IsNil(o.PaymentCurrency) {
		var ret string
		return ret
	}
	return *o.PaymentCurrency
}

// GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPaymentCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentCurrency) {
		return nil, false
	}
	return o.PaymentCurrency, true
}

// HasPaymentCurrency returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPaymentCurrency() bool {
	if o != nil && !IsNil(o.PaymentCurrency) {
		return true
	}

	return false
}

// SetPaymentCurrency gets a reference to the given string and assigns it to the PaymentCurrency field.
func (o *PaymentResponseV4) SetPaymentCurrency(v string) {
	o.PaymentCurrency = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRate() float64 {
	if o == nil || IsNil(o.Rate) {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRateOk() (*float64, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *PaymentResponseV4) SetRate(v float64) {
	o.Rate = &v
}

// GetInvertedRate returns the InvertedRate field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetInvertedRate() float64 {
	if o == nil || IsNil(o.InvertedRate) {
		var ret float64
		return ret
	}
	return *o.InvertedRate
}

// GetInvertedRateOk returns a tuple with the InvertedRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetInvertedRateOk() (*float64, bool) {
	if o == nil || IsNil(o.InvertedRate) {
		return nil, false
	}
	return o.InvertedRate, true
}

// HasInvertedRate returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasInvertedRate() bool {
	if o != nil && !IsNil(o.InvertedRate) {
		return true
	}

	return false
}

// SetInvertedRate gets a reference to the given float64 and assigns it to the InvertedRate field.
func (o *PaymentResponseV4) SetInvertedRate(v float64) {
	o.InvertedRate = &v
}

// GetIsPaymentCcyBaseCcy returns the IsPaymentCcyBaseCcy field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetIsPaymentCcyBaseCcy() bool {
	if o == nil || IsNil(o.IsPaymentCcyBaseCcy) {
		var ret bool
		return ret
	}
	return *o.IsPaymentCcyBaseCcy
}

// GetIsPaymentCcyBaseCcyOk returns a tuple with the IsPaymentCcyBaseCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetIsPaymentCcyBaseCcyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaymentCcyBaseCcy) {
		return nil, false
	}
	return o.IsPaymentCcyBaseCcy, true
}

// HasIsPaymentCcyBaseCcy returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasIsPaymentCcyBaseCcy() bool {
	if o != nil && !IsNil(o.IsPaymentCcyBaseCcy) {
		return true
	}

	return false
}

// SetIsPaymentCcyBaseCcy gets a reference to the given bool and assigns it to the IsPaymentCcyBaseCcy field.
func (o *PaymentResponseV4) SetIsPaymentCcyBaseCcy(v bool) {
	o.IsPaymentCcyBaseCcy = &v
}

// GetSubmittedDateTime returns the SubmittedDateTime field value
func (o *PaymentResponseV4) GetSubmittedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SubmittedDateTime
}

// GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetSubmittedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubmittedDateTime, true
}

// SetSubmittedDateTime sets field value
func (o *PaymentResponseV4) SetSubmittedDateTime(v time.Time) {
	o.SubmittedDateTime = v
}

// GetStatus returns the Status field value
func (o *PaymentResponseV4) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentResponseV4) SetStatus(v string) {
	o.Status = v
}

// GetFundingStatus returns the FundingStatus field value
func (o *PaymentResponseV4) GetFundingStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingStatus
}

// GetFundingStatusOk returns a tuple with the FundingStatus field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetFundingStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingStatus, true
}

// SetFundingStatus sets field value
func (o *PaymentResponseV4) SetFundingStatus(v string) {
	o.FundingStatus = v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRoutingNumber() string {
	if o == nil || IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRoutingNumber() bool {
	if o != nil && !IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *PaymentResponseV4) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *PaymentResponseV4) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetIban() string {
	if o == nil || IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetIbanOk() (*string, bool) {
	if o == nil || IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasIban() bool {
	if o != nil && !IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *PaymentResponseV4) SetIban(v string) {
	o.Iban = &v
}

// GetPaymentMemo returns the PaymentMemo field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPaymentMemo() string {
	if o == nil || IsNil(o.PaymentMemo) {
		var ret string
		return ret
	}
	return *o.PaymentMemo
}

// GetPaymentMemoOk returns a tuple with the PaymentMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPaymentMemoOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMemo) {
		return nil, false
	}
	return o.PaymentMemo, true
}

// HasPaymentMemo returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPaymentMemo() bool {
	if o != nil && !IsNil(o.PaymentMemo) {
		return true
	}

	return false
}

// SetPaymentMemo gets a reference to the given string and assigns it to the PaymentMemo field.
func (o *PaymentResponseV4) SetPaymentMemo(v string) {
	o.PaymentMemo = &v
}

// GetFilenameReference returns the FilenameReference field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetFilenameReference() string {
	if o == nil || IsNil(o.FilenameReference) {
		var ret string
		return ret
	}
	return *o.FilenameReference
}

// GetFilenameReferenceOk returns a tuple with the FilenameReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetFilenameReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.FilenameReference) {
		return nil, false
	}
	return o.FilenameReference, true
}

// HasFilenameReference returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasFilenameReference() bool {
	if o != nil && !IsNil(o.FilenameReference) {
		return true
	}

	return false
}

// SetFilenameReference gets a reference to the given string and assigns it to the FilenameReference field.
func (o *PaymentResponseV4) SetFilenameReference(v string) {
	o.FilenameReference = &v
}

// GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetIndividualIdentificationNumber() string {
	if o == nil || IsNil(o.IndividualIdentificationNumber) {
		var ret string
		return ret
	}
	return *o.IndividualIdentificationNumber
}

// GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetIndividualIdentificationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualIdentificationNumber) {
		return nil, false
	}
	return o.IndividualIdentificationNumber, true
}

// HasIndividualIdentificationNumber returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasIndividualIdentificationNumber() bool {
	if o != nil && !IsNil(o.IndividualIdentificationNumber) {
		return true
	}

	return false
}

// SetIndividualIdentificationNumber gets a reference to the given string and assigns it to the IndividualIdentificationNumber field.
func (o *PaymentResponseV4) SetIndividualIdentificationNumber(v string) {
	o.IndividualIdentificationNumber = &v
}

// GetTraceNumber returns the TraceNumber field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetTraceNumber() string {
	if o == nil || IsNil(o.TraceNumber) {
		var ret string
		return ret
	}
	return *o.TraceNumber
}

// GetTraceNumberOk returns a tuple with the TraceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetTraceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TraceNumber) {
		return nil, false
	}
	return o.TraceNumber, true
}

// HasTraceNumber returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasTraceNumber() bool {
	if o != nil && !IsNil(o.TraceNumber) {
		return true
	}

	return false
}

// SetTraceNumber gets a reference to the given string and assigns it to the TraceNumber field.
func (o *PaymentResponseV4) SetTraceNumber(v string) {
	o.TraceNumber = &v
}

// GetPayorPaymentId returns the PayorPaymentId field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPayorPaymentId() string {
	if o == nil || IsNil(o.PayorPaymentId) {
		var ret string
		return ret
	}
	return *o.PayorPaymentId
}

// GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPayorPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayorPaymentId) {
		return nil, false
	}
	return o.PayorPaymentId, true
}

// HasPayorPaymentId returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPayorPaymentId() bool {
	if o != nil && !IsNil(o.PayorPaymentId) {
		return true
	}

	return false
}

// SetPayorPaymentId gets a reference to the given string and assigns it to the PayorPaymentId field.
func (o *PaymentResponseV4) SetPayorPaymentId(v string) {
	o.PayorPaymentId = &v
}

// GetPaymentChannelId returns the PaymentChannelId field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPaymentChannelId() string {
	if o == nil || IsNil(o.PaymentChannelId) {
		var ret string
		return ret
	}
	return *o.PaymentChannelId
}

// GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPaymentChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelId) {
		return nil, false
	}
	return o.PaymentChannelId, true
}

// HasPaymentChannelId returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPaymentChannelId() bool {
	if o != nil && !IsNil(o.PaymentChannelId) {
		return true
	}

	return false
}

// SetPaymentChannelId gets a reference to the given string and assigns it to the PaymentChannelId field.
func (o *PaymentResponseV4) SetPaymentChannelId(v string) {
	o.PaymentChannelId = &v
}

// GetPaymentChannelName returns the PaymentChannelName field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPaymentChannelName() string {
	if o == nil || IsNil(o.PaymentChannelName) {
		var ret string
		return ret
	}
	return *o.PaymentChannelName
}

// GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPaymentChannelNameOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelName) {
		return nil, false
	}
	return o.PaymentChannelName, true
}

// HasPaymentChannelName returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPaymentChannelName() bool {
	if o != nil && !IsNil(o.PaymentChannelName) {
		return true
	}

	return false
}

// SetPaymentChannelName gets a reference to the given string and assigns it to the PaymentChannelName field.
func (o *PaymentResponseV4) SetPaymentChannelName(v string) {
	o.PaymentChannelName = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *PaymentResponseV4) SetAccountName(v string) {
	o.AccountName = &v
}

// GetRailsId returns the RailsId field value
func (o *PaymentResponseV4) GetRailsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RailsId
}

// GetRailsIdOk returns a tuple with the RailsId field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRailsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RailsId, true
}

// SetRailsId sets field value
func (o *PaymentResponseV4) SetRailsId(v string) {
	o.RailsId = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PaymentResponseV4) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPayeeAddressCountryCode returns the PayeeAddressCountryCode field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPayeeAddressCountryCode() string {
	if o == nil || IsNil(o.PayeeAddressCountryCode) {
		var ret string
		return ret
	}
	return *o.PayeeAddressCountryCode
}

// GetPayeeAddressCountryCodeOk returns a tuple with the PayeeAddressCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPayeeAddressCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeAddressCountryCode) {
		return nil, false
	}
	return o.PayeeAddressCountryCode, true
}

// HasPayeeAddressCountryCode returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPayeeAddressCountryCode() bool {
	if o != nil && !IsNil(o.PayeeAddressCountryCode) {
		return true
	}

	return false
}

// SetPayeeAddressCountryCode gets a reference to the given string and assigns it to the PayeeAddressCountryCode field.
func (o *PaymentResponseV4) SetPayeeAddressCountryCode(v string) {
	o.PayeeAddressCountryCode = &v
}

// GetEvents returns the Events field value
func (o *PaymentResponseV4) GetEvents() []PaymentEventResponse {
	if o == nil {
		var ret []PaymentEventResponse
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetEventsOk() ([]PaymentEventResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *PaymentResponseV4) SetEvents(v []PaymentEventResponse) {
	o.Events = v
}

// GetReturnCost returns the ReturnCost field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetReturnCost() int32 {
	if o == nil || IsNil(o.ReturnCost) {
		var ret int32
		return ret
	}
	return *o.ReturnCost
}

// GetReturnCostOk returns a tuple with the ReturnCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetReturnCostOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnCost) {
		return nil, false
	}
	return o.ReturnCost, true
}

// HasReturnCost returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasReturnCost() bool {
	if o != nil && !IsNil(o.ReturnCost) {
		return true
	}

	return false
}

// SetReturnCost gets a reference to the given int32 and assigns it to the ReturnCost field.
func (o *PaymentResponseV4) SetReturnCost(v int32) {
	o.ReturnCost = &v
}

// GetReturnReason returns the ReturnReason field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetReturnReason() string {
	if o == nil || IsNil(o.ReturnReason) {
		var ret string
		return ret
	}
	return *o.ReturnReason
}

// GetReturnReasonOk returns a tuple with the ReturnReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetReturnReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnReason) {
		return nil, false
	}
	return o.ReturnReason, true
}

// HasReturnReason returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasReturnReason() bool {
	if o != nil && !IsNil(o.ReturnReason) {
		return true
	}

	return false
}

// SetReturnReason gets a reference to the given string and assigns it to the ReturnReason field.
func (o *PaymentResponseV4) SetReturnReason(v string) {
	o.ReturnReason = &v
}

// GetRailsPaymentId returns the RailsPaymentId field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRailsPaymentId() string {
	if o == nil || IsNil(o.RailsPaymentId) {
		var ret string
		return ret
	}
	return *o.RailsPaymentId
}

// GetRailsPaymentIdOk returns a tuple with the RailsPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRailsPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.RailsPaymentId) {
		return nil, false
	}
	return o.RailsPaymentId, true
}

// HasRailsPaymentId returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRailsPaymentId() bool {
	if o != nil && !IsNil(o.RailsPaymentId) {
		return true
	}

	return false
}

// SetRailsPaymentId gets a reference to the given string and assigns it to the RailsPaymentId field.
func (o *PaymentResponseV4) SetRailsPaymentId(v string) {
	o.RailsPaymentId = &v
}

// GetRailsBatchId returns the RailsBatchId field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRailsBatchId() string {
	if o == nil || IsNil(o.RailsBatchId) {
		var ret string
		return ret
	}
	return *o.RailsBatchId
}

// GetRailsBatchIdOk returns a tuple with the RailsBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRailsBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.RailsBatchId) {
		return nil, false
	}
	return o.RailsBatchId, true
}

// HasRailsBatchId returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRailsBatchId() bool {
	if o != nil && !IsNil(o.RailsBatchId) {
		return true
	}

	return false
}

// SetRailsBatchId gets a reference to the given string and assigns it to the RailsBatchId field.
func (o *PaymentResponseV4) SetRailsBatchId(v string) {
	o.RailsBatchId = &v
}

// GetRailsAccountId returns the RailsAccountId field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRailsAccountId() string {
	if o == nil || IsNil(o.RailsAccountId) {
		var ret string
		return ret
	}
	return *o.RailsAccountId
}

// GetRailsAccountIdOk returns a tuple with the RailsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRailsAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.RailsAccountId) {
		return nil, false
	}
	return o.RailsAccountId, true
}

// HasRailsAccountId returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRailsAccountId() bool {
	if o != nil && !IsNil(o.RailsAccountId) {
		return true
	}

	return false
}

// SetRailsAccountId gets a reference to the given string and assigns it to the RailsAccountId field.
func (o *PaymentResponseV4) SetRailsAccountId(v string) {
	o.RailsAccountId = &v
}

// GetPaymentScheme returns the PaymentScheme field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPaymentScheme() string {
	if o == nil || IsNil(o.PaymentScheme) {
		var ret string
		return ret
	}
	return *o.PaymentScheme
}

// GetPaymentSchemeOk returns a tuple with the PaymentScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPaymentSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentScheme) {
		return nil, false
	}
	return o.PaymentScheme, true
}

// HasPaymentScheme returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPaymentScheme() bool {
	if o != nil && !IsNil(o.PaymentScheme) {
		return true
	}

	return false
}

// SetPaymentScheme gets a reference to the given string and assigns it to the PaymentScheme field.
func (o *PaymentResponseV4) SetPaymentScheme(v string) {
	o.PaymentScheme = &v
}

// GetRejectionReason returns the RejectionReason field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRejectionReason() string {
	if o == nil || IsNil(o.RejectionReason) {
		var ret string
		return ret
	}
	return *o.RejectionReason
}

// GetRejectionReasonOk returns a tuple with the RejectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRejectionReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RejectionReason) {
		return nil, false
	}
	return o.RejectionReason, true
}

// HasRejectionReason returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRejectionReason() bool {
	if o != nil && !IsNil(o.RejectionReason) {
		return true
	}

	return false
}

// SetRejectionReason gets a reference to the given string and assigns it to the RejectionReason field.
func (o *PaymentResponseV4) SetRejectionReason(v string) {
	o.RejectionReason = &v
}

// GetRailsRejectionInformation returns the RailsRejectionInformation field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetRailsRejectionInformation() string {
	if o == nil || IsNil(o.RailsRejectionInformation) {
		var ret string
		return ret
	}
	return *o.RailsRejectionInformation
}

// GetRailsRejectionInformationOk returns a tuple with the RailsRejectionInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetRailsRejectionInformationOk() (*string, bool) {
	if o == nil || IsNil(o.RailsRejectionInformation) {
		return nil, false
	}
	return o.RailsRejectionInformation, true
}

// HasRailsRejectionInformation returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasRailsRejectionInformation() bool {
	if o != nil && !IsNil(o.RailsRejectionInformation) {
		return true
	}

	return false
}

// SetRailsRejectionInformation gets a reference to the given string and assigns it to the RailsRejectionInformation field.
func (o *PaymentResponseV4) SetRailsRejectionInformation(v string) {
	o.RailsRejectionInformation = &v
}

// GetWithdrawnReason returns the WithdrawnReason field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetWithdrawnReason() string {
	if o == nil || IsNil(o.WithdrawnReason) {
		var ret string
		return ret
	}
	return *o.WithdrawnReason
}

// GetWithdrawnReasonOk returns a tuple with the WithdrawnReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetWithdrawnReasonOk() (*string, bool) {
	if o == nil || IsNil(o.WithdrawnReason) {
		return nil, false
	}
	return o.WithdrawnReason, true
}

// HasWithdrawnReason returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasWithdrawnReason() bool {
	if o != nil && !IsNil(o.WithdrawnReason) {
		return true
	}

	return false
}

// SetWithdrawnReason gets a reference to the given string and assigns it to the WithdrawnReason field.
func (o *PaymentResponseV4) SetWithdrawnReason(v string) {
	o.WithdrawnReason = &v
}

// GetWithdrawable returns the Withdrawable field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetWithdrawable() bool {
	if o == nil || IsNil(o.Withdrawable) {
		var ret bool
		return ret
	}
	return *o.Withdrawable
}

// GetWithdrawableOk returns a tuple with the Withdrawable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetWithdrawableOk() (*bool, bool) {
	if o == nil || IsNil(o.Withdrawable) {
		return nil, false
	}
	return o.Withdrawable, true
}

// HasWithdrawable returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasWithdrawable() bool {
	if o != nil && !IsNil(o.Withdrawable) {
		return true
	}

	return false
}

// SetWithdrawable gets a reference to the given bool and assigns it to the Withdrawable field.
func (o *PaymentResponseV4) SetWithdrawable(v bool) {
	o.Withdrawable = &v
}

// GetAutoWithdrawnReasonCode returns the AutoWithdrawnReasonCode field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetAutoWithdrawnReasonCode() string {
	if o == nil || IsNil(o.AutoWithdrawnReasonCode) {
		var ret string
		return ret
	}
	return *o.AutoWithdrawnReasonCode
}

// GetAutoWithdrawnReasonCodeOk returns a tuple with the AutoWithdrawnReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetAutoWithdrawnReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AutoWithdrawnReasonCode) {
		return nil, false
	}
	return o.AutoWithdrawnReasonCode, true
}

// HasAutoWithdrawnReasonCode returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasAutoWithdrawnReasonCode() bool {
	if o != nil && !IsNil(o.AutoWithdrawnReasonCode) {
		return true
	}

	return false
}

// SetAutoWithdrawnReasonCode gets a reference to the given string and assigns it to the AutoWithdrawnReasonCode field.
func (o *PaymentResponseV4) SetAutoWithdrawnReasonCode(v string) {
	o.AutoWithdrawnReasonCode = &v
}

// GetTransmissionType returns the TransmissionType field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetTransmissionType() string {
	if o == nil || IsNil(o.TransmissionType) {
		var ret string
		return ret
	}
	return *o.TransmissionType
}

// GetTransmissionTypeOk returns a tuple with the TransmissionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetTransmissionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransmissionType) {
		return nil, false
	}
	return o.TransmissionType, true
}

// HasTransmissionType returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasTransmissionType() bool {
	if o != nil && !IsNil(o.TransmissionType) {
		return true
	}

	return false
}

// SetTransmissionType gets a reference to the given string and assigns it to the TransmissionType field.
func (o *PaymentResponseV4) SetTransmissionType(v string) {
	o.TransmissionType = &v
}

// GetTransmissionTypeRequested returns the TransmissionTypeRequested field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetTransmissionTypeRequested() string {
	if o == nil || IsNil(o.TransmissionTypeRequested) {
		var ret string
		return ret
	}
	return *o.TransmissionTypeRequested
}

// GetTransmissionTypeRequestedOk returns a tuple with the TransmissionTypeRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetTransmissionTypeRequestedOk() (*string, bool) {
	if o == nil || IsNil(o.TransmissionTypeRequested) {
		return nil, false
	}
	return o.TransmissionTypeRequested, true
}

// HasTransmissionTypeRequested returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasTransmissionTypeRequested() bool {
	if o != nil && !IsNil(o.TransmissionTypeRequested) {
		return true
	}

	return false
}

// SetTransmissionTypeRequested gets a reference to the given string and assigns it to the TransmissionTypeRequested field.
func (o *PaymentResponseV4) SetTransmissionTypeRequested(v string) {
	o.TransmissionTypeRequested = &v
}

// GetPaymentTrackingReference returns the PaymentTrackingReference field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPaymentTrackingReference() string {
	if o == nil || IsNil(o.PaymentTrackingReference) {
		var ret string
		return ret
	}
	return *o.PaymentTrackingReference
}

// GetPaymentTrackingReferenceOk returns a tuple with the PaymentTrackingReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPaymentTrackingReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTrackingReference) {
		return nil, false
	}
	return o.PaymentTrackingReference, true
}

// HasPaymentTrackingReference returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPaymentTrackingReference() bool {
	if o != nil && !IsNil(o.PaymentTrackingReference) {
		return true
	}

	return false
}

// SetPaymentTrackingReference gets a reference to the given string and assigns it to the PaymentTrackingReference field.
func (o *PaymentResponseV4) SetPaymentTrackingReference(v string) {
	o.PaymentTrackingReference = &v
}

// GetPaymentMetadata returns the PaymentMetadata field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPaymentMetadata() string {
	if o == nil || IsNil(o.PaymentMetadata) {
		var ret string
		return ret
	}
	return *o.PaymentMetadata
}

// GetPaymentMetadataOk returns a tuple with the PaymentMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPaymentMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMetadata) {
		return nil, false
	}
	return o.PaymentMetadata, true
}

// HasPaymentMetadata returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPaymentMetadata() bool {
	if o != nil && !IsNil(o.PaymentMetadata) {
		return true
	}

	return false
}

// SetPaymentMetadata gets a reference to the given string and assigns it to the PaymentMetadata field.
func (o *PaymentResponseV4) SetPaymentMetadata(v string) {
	o.PaymentMetadata = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *PaymentResponseV4) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionReference returns the TransactionReference field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetTransactionReference() string {
	if o == nil || IsNil(o.TransactionReference) {
		var ret string
		return ret
	}
	return *o.TransactionReference
}

// GetTransactionReferenceOk returns a tuple with the TransactionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetTransactionReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionReference) {
		return nil, false
	}
	return o.TransactionReference, true
}

// HasTransactionReference returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasTransactionReference() bool {
	if o != nil && !IsNil(o.TransactionReference) {
		return true
	}

	return false
}

// SetTransactionReference gets a reference to the given string and assigns it to the TransactionReference field.
func (o *PaymentResponseV4) SetTransactionReference(v string) {
	o.TransactionReference = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetSchedule() PayoutSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret PayoutSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetScheduleOk() (*PayoutSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given PayoutSchedule and assigns it to the Schedule field.
func (o *PaymentResponseV4) SetSchedule(v PayoutSchedule) {
	o.Schedule = &v
}

// GetPostInstructFxInfo returns the PostInstructFxInfo field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPostInstructFxInfo() PostInstructFxInfo {
	if o == nil || IsNil(o.PostInstructFxInfo) {
		var ret PostInstructFxInfo
		return ret
	}
	return *o.PostInstructFxInfo
}

// GetPostInstructFxInfoOk returns a tuple with the PostInstructFxInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPostInstructFxInfoOk() (*PostInstructFxInfo, bool) {
	if o == nil || IsNil(o.PostInstructFxInfo) {
		return nil, false
	}
	return o.PostInstructFxInfo, true
}

// HasPostInstructFxInfo returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPostInstructFxInfo() bool {
	if o != nil && !IsNil(o.PostInstructFxInfo) {
		return true
	}

	return false
}

// SetPostInstructFxInfo gets a reference to the given PostInstructFxInfo and assigns it to the PostInstructFxInfo field.
func (o *PaymentResponseV4) SetPostInstructFxInfo(v PostInstructFxInfo) {
	o.PostInstructFxInfo = &v
}

// GetPayout returns the Payout field value if set, zero value otherwise.
func (o *PaymentResponseV4) GetPayout() PaymentResponseV4Payout {
	if o == nil || IsNil(o.Payout) {
		var ret PaymentResponseV4Payout
		return ret
	}
	return *o.Payout
}

// GetPayoutOk returns a tuple with the Payout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponseV4) GetPayoutOk() (*PaymentResponseV4Payout, bool) {
	if o == nil || IsNil(o.Payout) {
		return nil, false
	}
	return o.Payout, true
}

// HasPayout returns a boolean if a field has been set.
func (o *PaymentResponseV4) HasPayout() bool {
	if o != nil && !IsNil(o.Payout) {
		return true
	}

	return false
}

// SetPayout gets a reference to the given PaymentResponseV4Payout and assigns it to the Payout field.
func (o *PaymentResponseV4) SetPayout(v PaymentResponseV4Payout) {
	o.Payout = &v
}

func (o PaymentResponseV4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentResponseV4) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RemoteSystemId) {
		toSerialize["remoteSystemId"] = o.RemoteSystemId
	}
	if !IsNil(o.RemoteSystemPaymentId) {
		toSerialize["remoteSystemPaymentId"] = o.RemoteSystemPaymentId
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
	if !IsNil(o.IsPaymentCcyBaseCcy) {
		toSerialize["isPaymentCcyBaseCcy"] = o.IsPaymentCcyBaseCcy
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
	if !IsNil(o.PayeeAddressCountryCode) {
		toSerialize["payeeAddressCountryCode"] = o.PayeeAddressCountryCode
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
	if !IsNil(o.RailsAccountId) {
		toSerialize["railsAccountId"] = o.RailsAccountId
	}
	if !IsNil(o.PaymentScheme) {
		toSerialize["paymentScheme"] = o.PaymentScheme
	}
	if !IsNil(o.RejectionReason) {
		toSerialize["rejectionReason"] = o.RejectionReason
	}
	if !IsNil(o.RailsRejectionInformation) {
		toSerialize["railsRejectionInformation"] = o.RailsRejectionInformation
	}
	if !IsNil(o.WithdrawnReason) {
		toSerialize["withdrawnReason"] = o.WithdrawnReason
	}
	if !IsNil(o.Withdrawable) {
		toSerialize["withdrawable"] = o.Withdrawable
	}
	if !IsNil(o.AutoWithdrawnReasonCode) {
		toSerialize["autoWithdrawnReasonCode"] = o.AutoWithdrawnReasonCode
	}
	if !IsNil(o.TransmissionType) {
		toSerialize["transmissionType"] = o.TransmissionType
	}
	if !IsNil(o.TransmissionTypeRequested) {
		toSerialize["transmissionTypeRequested"] = o.TransmissionTypeRequested
	}
	if !IsNil(o.PaymentTrackingReference) {
		toSerialize["paymentTrackingReference"] = o.PaymentTrackingReference
	}
	if !IsNil(o.PaymentMetadata) {
		toSerialize["paymentMetadata"] = o.PaymentMetadata
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !IsNil(o.TransactionReference) {
		toSerialize["transactionReference"] = o.TransactionReference
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.PostInstructFxInfo) {
		toSerialize["postInstructFxInfo"] = o.PostInstructFxInfo
	}
	if !IsNil(o.Payout) {
		toSerialize["payout"] = o.Payout
	}
	return toSerialize, nil
}

type NullablePaymentResponseV4 struct {
	value *PaymentResponseV4
	isSet bool
}

func (v NullablePaymentResponseV4) Get() *PaymentResponseV4 {
	return v.value
}

func (v *NullablePaymentResponseV4) Set(val *PaymentResponseV4) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentResponseV4) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentResponseV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentResponseV4(val *PaymentResponseV4) *NullablePaymentResponseV4 {
	return &NullablePaymentResponseV4{value: val, isSet: true}
}

func (v NullablePaymentResponseV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentResponseV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


