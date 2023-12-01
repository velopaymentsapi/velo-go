/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response.   ## Http Status Codes Following is a list of Http Status codes that could be returned by the platform      | Status Code            | Description                                                                          |     | -----------------------| -------------------------------------------------------------------------------------|     | 200 OK                 | The request was successfully processed and usually returns a json response           |     | 201 Created            | A resource was created and a Location header is returned linking to the new resource |     | 202 Accepted           | The request has been accepted for processing                                         |     | 204 No Content         | The request has been processed and there is no response (usually deletes and updates)|     | 400 Bad Request        | The request is invalid and should be fixed before retrying                           |     | 401 Unauthorized       | Authentication has failed, usually means the token has expired                       |     | 403 Forbidden          | The user does not have permissions for the request                                   |     | 404 Not Found          | The resource was not found                                                           |     | 409 Conflict           | The resource already exists and there is a conflict                                  |     | 429 Too Many Requests  | The user has submitted too many requests in a given amount of time                   |     | 5xx Server Error       | Platform internal error (should rarely happen)                                       | 

API version: 2.37.150
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// checks if the PayorAmlTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayorAmlTransaction{}

// PayorAmlTransaction struct for PayorAmlTransaction
type PayorAmlTransaction struct {
	TransactionDate *string `json:"transactionDate,omitempty"`
	TransactionTime *string `json:"transactionTime,omitempty"`
	ReportTransactionType *string `json:"reportTransactionType,omitempty"`
	Debit *int64 `json:"debit,omitempty"`
	// ISO 4217 3 character currency code
	DebitCurrency *string `json:"debitCurrency,omitempty"`
	Credit *int64 `json:"credit,omitempty"`
	// ISO 4217 3 character currency code
	CreditCurrency *string `json:"creditCurrency,omitempty"`
	ReturnFee *string `json:"returnFee,omitempty"`
	// ISO 4217 3 character currency code
	ReturnFeeCurrency *string `json:"returnFeeCurrency,omitempty"`
	ReturnFeeDescription *string `json:"returnFeeDescription,omitempty"`
	ReturnCode *string `json:"returnCode,omitempty"`
	ReturnDescription *string `json:"returnDescription,omitempty"`
	FundingType *string `json:"fundingType,omitempty"`
	DateFundingRequested *string `json:"dateFundingRequested,omitempty"`
	PayeeName *string `json:"payeeName,omitempty"`
	// Remote ID of the Payee, set by Payor
	RemoteId *string `json:"remoteId,omitempty"`
	PayeeType *string `json:"payeeType,omitempty"`
	PayeeEmail *string `json:"payeeEmail,omitempty"`
	SourceAccount *string `json:"sourceAccount,omitempty"`
	PaymentAmount *int64 `json:"paymentAmount,omitempty"`
	// ISO 4217 3 character currency code
	PaymentCurrency *string `json:"paymentCurrency,omitempty"`
	PaymentMemo *string `json:"paymentMemo,omitempty"`
	PaymentRails *string `json:"paymentRails,omitempty"`
	PayorPaymentId *string `json:"payorPaymentId,omitempty"`
	PaymentStatus *string `json:"paymentStatus,omitempty"`
	RejectReason *string `json:"rejectReason,omitempty"`
	FxApplied *float64 `json:"fxApplied,omitempty"`
}

// NewPayorAmlTransaction instantiates a new PayorAmlTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayorAmlTransaction() *PayorAmlTransaction {
	this := PayorAmlTransaction{}
	return &this
}

// NewPayorAmlTransactionWithDefaults instantiates a new PayorAmlTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayorAmlTransactionWithDefaults() *PayorAmlTransaction {
	this := PayorAmlTransaction{}
	return &this
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *PayorAmlTransaction) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetTransactionTime() string {
	if o == nil || IsNil(o.TransactionTime) {
		var ret string
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetTransactionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasTransactionTime() bool {
	if o != nil && !IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given string and assigns it to the TransactionTime field.
func (o *PayorAmlTransaction) SetTransactionTime(v string) {
	o.TransactionTime = &v
}

// GetReportTransactionType returns the ReportTransactionType field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetReportTransactionType() string {
	if o == nil || IsNil(o.ReportTransactionType) {
		var ret string
		return ret
	}
	return *o.ReportTransactionType
}

// GetReportTransactionTypeOk returns a tuple with the ReportTransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetReportTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReportTransactionType) {
		return nil, false
	}
	return o.ReportTransactionType, true
}

// HasReportTransactionType returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasReportTransactionType() bool {
	if o != nil && !IsNil(o.ReportTransactionType) {
		return true
	}

	return false
}

// SetReportTransactionType gets a reference to the given string and assigns it to the ReportTransactionType field.
func (o *PayorAmlTransaction) SetReportTransactionType(v string) {
	o.ReportTransactionType = &v
}

// GetDebit returns the Debit field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetDebit() int64 {
	if o == nil || IsNil(o.Debit) {
		var ret int64
		return ret
	}
	return *o.Debit
}

// GetDebitOk returns a tuple with the Debit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetDebitOk() (*int64, bool) {
	if o == nil || IsNil(o.Debit) {
		return nil, false
	}
	return o.Debit, true
}

// HasDebit returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasDebit() bool {
	if o != nil && !IsNil(o.Debit) {
		return true
	}

	return false
}

// SetDebit gets a reference to the given int64 and assigns it to the Debit field.
func (o *PayorAmlTransaction) SetDebit(v int64) {
	o.Debit = &v
}

// GetDebitCurrency returns the DebitCurrency field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetDebitCurrency() string {
	if o == nil || IsNil(o.DebitCurrency) {
		var ret string
		return ret
	}
	return *o.DebitCurrency
}

// GetDebitCurrencyOk returns a tuple with the DebitCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetDebitCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.DebitCurrency) {
		return nil, false
	}
	return o.DebitCurrency, true
}

// HasDebitCurrency returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasDebitCurrency() bool {
	if o != nil && !IsNil(o.DebitCurrency) {
		return true
	}

	return false
}

// SetDebitCurrency gets a reference to the given string and assigns it to the DebitCurrency field.
func (o *PayorAmlTransaction) SetDebitCurrency(v string) {
	o.DebitCurrency = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetCredit() int64 {
	if o == nil || IsNil(o.Credit) {
		var ret int64
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetCreditOk() (*int64, bool) {
	if o == nil || IsNil(o.Credit) {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasCredit() bool {
	if o != nil && !IsNil(o.Credit) {
		return true
	}

	return false
}

// SetCredit gets a reference to the given int64 and assigns it to the Credit field.
func (o *PayorAmlTransaction) SetCredit(v int64) {
	o.Credit = &v
}

// GetCreditCurrency returns the CreditCurrency field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetCreditCurrency() string {
	if o == nil || IsNil(o.CreditCurrency) {
		var ret string
		return ret
	}
	return *o.CreditCurrency
}

// GetCreditCurrencyOk returns a tuple with the CreditCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetCreditCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.CreditCurrency) {
		return nil, false
	}
	return o.CreditCurrency, true
}

// HasCreditCurrency returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasCreditCurrency() bool {
	if o != nil && !IsNil(o.CreditCurrency) {
		return true
	}

	return false
}

// SetCreditCurrency gets a reference to the given string and assigns it to the CreditCurrency field.
func (o *PayorAmlTransaction) SetCreditCurrency(v string) {
	o.CreditCurrency = &v
}

// GetReturnFee returns the ReturnFee field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetReturnFee() string {
	if o == nil || IsNil(o.ReturnFee) {
		var ret string
		return ret
	}
	return *o.ReturnFee
}

// GetReturnFeeOk returns a tuple with the ReturnFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetReturnFeeOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnFee) {
		return nil, false
	}
	return o.ReturnFee, true
}

// HasReturnFee returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasReturnFee() bool {
	if o != nil && !IsNil(o.ReturnFee) {
		return true
	}

	return false
}

// SetReturnFee gets a reference to the given string and assigns it to the ReturnFee field.
func (o *PayorAmlTransaction) SetReturnFee(v string) {
	o.ReturnFee = &v
}

// GetReturnFeeCurrency returns the ReturnFeeCurrency field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetReturnFeeCurrency() string {
	if o == nil || IsNil(o.ReturnFeeCurrency) {
		var ret string
		return ret
	}
	return *o.ReturnFeeCurrency
}

// GetReturnFeeCurrencyOk returns a tuple with the ReturnFeeCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetReturnFeeCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnFeeCurrency) {
		return nil, false
	}
	return o.ReturnFeeCurrency, true
}

// HasReturnFeeCurrency returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasReturnFeeCurrency() bool {
	if o != nil && !IsNil(o.ReturnFeeCurrency) {
		return true
	}

	return false
}

// SetReturnFeeCurrency gets a reference to the given string and assigns it to the ReturnFeeCurrency field.
func (o *PayorAmlTransaction) SetReturnFeeCurrency(v string) {
	o.ReturnFeeCurrency = &v
}

// GetReturnFeeDescription returns the ReturnFeeDescription field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetReturnFeeDescription() string {
	if o == nil || IsNil(o.ReturnFeeDescription) {
		var ret string
		return ret
	}
	return *o.ReturnFeeDescription
}

// GetReturnFeeDescriptionOk returns a tuple with the ReturnFeeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetReturnFeeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnFeeDescription) {
		return nil, false
	}
	return o.ReturnFeeDescription, true
}

// HasReturnFeeDescription returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasReturnFeeDescription() bool {
	if o != nil && !IsNil(o.ReturnFeeDescription) {
		return true
	}

	return false
}

// SetReturnFeeDescription gets a reference to the given string and assigns it to the ReturnFeeDescription field.
func (o *PayorAmlTransaction) SetReturnFeeDescription(v string) {
	o.ReturnFeeDescription = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetReturnCode() string {
	if o == nil || IsNil(o.ReturnCode) {
		var ret string
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetReturnCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnCode) {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasReturnCode() bool {
	if o != nil && !IsNil(o.ReturnCode) {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given string and assigns it to the ReturnCode field.
func (o *PayorAmlTransaction) SetReturnCode(v string) {
	o.ReturnCode = &v
}

// GetReturnDescription returns the ReturnDescription field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetReturnDescription() string {
	if o == nil || IsNil(o.ReturnDescription) {
		var ret string
		return ret
	}
	return *o.ReturnDescription
}

// GetReturnDescriptionOk returns a tuple with the ReturnDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetReturnDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnDescription) {
		return nil, false
	}
	return o.ReturnDescription, true
}

// HasReturnDescription returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasReturnDescription() bool {
	if o != nil && !IsNil(o.ReturnDescription) {
		return true
	}

	return false
}

// SetReturnDescription gets a reference to the given string and assigns it to the ReturnDescription field.
func (o *PayorAmlTransaction) SetReturnDescription(v string) {
	o.ReturnDescription = &v
}

// GetFundingType returns the FundingType field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetFundingType() string {
	if o == nil || IsNil(o.FundingType) {
		var ret string
		return ret
	}
	return *o.FundingType
}

// GetFundingTypeOk returns a tuple with the FundingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetFundingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FundingType) {
		return nil, false
	}
	return o.FundingType, true
}

// HasFundingType returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasFundingType() bool {
	if o != nil && !IsNil(o.FundingType) {
		return true
	}

	return false
}

// SetFundingType gets a reference to the given string and assigns it to the FundingType field.
func (o *PayorAmlTransaction) SetFundingType(v string) {
	o.FundingType = &v
}

// GetDateFundingRequested returns the DateFundingRequested field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetDateFundingRequested() string {
	if o == nil || IsNil(o.DateFundingRequested) {
		var ret string
		return ret
	}
	return *o.DateFundingRequested
}

// GetDateFundingRequestedOk returns a tuple with the DateFundingRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetDateFundingRequestedOk() (*string, bool) {
	if o == nil || IsNil(o.DateFundingRequested) {
		return nil, false
	}
	return o.DateFundingRequested, true
}

// HasDateFundingRequested returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasDateFundingRequested() bool {
	if o != nil && !IsNil(o.DateFundingRequested) {
		return true
	}

	return false
}

// SetDateFundingRequested gets a reference to the given string and assigns it to the DateFundingRequested field.
func (o *PayorAmlTransaction) SetDateFundingRequested(v string) {
	o.DateFundingRequested = &v
}

// GetPayeeName returns the PayeeName field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetPayeeName() string {
	if o == nil || IsNil(o.PayeeName) {
		var ret string
		return ret
	}
	return *o.PayeeName
}

// GetPayeeNameOk returns a tuple with the PayeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetPayeeNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeName) {
		return nil, false
	}
	return o.PayeeName, true
}

// HasPayeeName returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasPayeeName() bool {
	if o != nil && !IsNil(o.PayeeName) {
		return true
	}

	return false
}

// SetPayeeName gets a reference to the given string and assigns it to the PayeeName field.
func (o *PayorAmlTransaction) SetPayeeName(v string) {
	o.PayeeName = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetRemoteId() string {
	if o == nil || IsNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteId) {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasRemoteId() bool {
	if o != nil && !IsNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *PayorAmlTransaction) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetPayeeType returns the PayeeType field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetPayeeType() string {
	if o == nil || IsNil(o.PayeeType) {
		var ret string
		return ret
	}
	return *o.PayeeType
}

// GetPayeeTypeOk returns a tuple with the PayeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetPayeeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeType) {
		return nil, false
	}
	return o.PayeeType, true
}

// HasPayeeType returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasPayeeType() bool {
	if o != nil && !IsNil(o.PayeeType) {
		return true
	}

	return false
}

// SetPayeeType gets a reference to the given string and assigns it to the PayeeType field.
func (o *PayorAmlTransaction) SetPayeeType(v string) {
	o.PayeeType = &v
}

// GetPayeeEmail returns the PayeeEmail field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetPayeeEmail() string {
	if o == nil || IsNil(o.PayeeEmail) {
		var ret string
		return ret
	}
	return *o.PayeeEmail
}

// GetPayeeEmailOk returns a tuple with the PayeeEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetPayeeEmailOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeEmail) {
		return nil, false
	}
	return o.PayeeEmail, true
}

// HasPayeeEmail returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasPayeeEmail() bool {
	if o != nil && !IsNil(o.PayeeEmail) {
		return true
	}

	return false
}

// SetPayeeEmail gets a reference to the given string and assigns it to the PayeeEmail field.
func (o *PayorAmlTransaction) SetPayeeEmail(v string) {
	o.PayeeEmail = &v
}

// GetSourceAccount returns the SourceAccount field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetSourceAccount() string {
	if o == nil || IsNil(o.SourceAccount) {
		var ret string
		return ret
	}
	return *o.SourceAccount
}

// GetSourceAccountOk returns a tuple with the SourceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetSourceAccountOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAccount) {
		return nil, false
	}
	return o.SourceAccount, true
}

// HasSourceAccount returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasSourceAccount() bool {
	if o != nil && !IsNil(o.SourceAccount) {
		return true
	}

	return false
}

// SetSourceAccount gets a reference to the given string and assigns it to the SourceAccount field.
func (o *PayorAmlTransaction) SetSourceAccount(v string) {
	o.SourceAccount = &v
}

// GetPaymentAmount returns the PaymentAmount field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetPaymentAmount() int64 {
	if o == nil || IsNil(o.PaymentAmount) {
		var ret int64
		return ret
	}
	return *o.PaymentAmount
}

// GetPaymentAmountOk returns a tuple with the PaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetPaymentAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.PaymentAmount) {
		return nil, false
	}
	return o.PaymentAmount, true
}

// HasPaymentAmount returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasPaymentAmount() bool {
	if o != nil && !IsNil(o.PaymentAmount) {
		return true
	}

	return false
}

// SetPaymentAmount gets a reference to the given int64 and assigns it to the PaymentAmount field.
func (o *PayorAmlTransaction) SetPaymentAmount(v int64) {
	o.PaymentAmount = &v
}

// GetPaymentCurrency returns the PaymentCurrency field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetPaymentCurrency() string {
	if o == nil || IsNil(o.PaymentCurrency) {
		var ret string
		return ret
	}
	return *o.PaymentCurrency
}

// GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetPaymentCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentCurrency) {
		return nil, false
	}
	return o.PaymentCurrency, true
}

// HasPaymentCurrency returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasPaymentCurrency() bool {
	if o != nil && !IsNil(o.PaymentCurrency) {
		return true
	}

	return false
}

// SetPaymentCurrency gets a reference to the given string and assigns it to the PaymentCurrency field.
func (o *PayorAmlTransaction) SetPaymentCurrency(v string) {
	o.PaymentCurrency = &v
}

// GetPaymentMemo returns the PaymentMemo field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetPaymentMemo() string {
	if o == nil || IsNil(o.PaymentMemo) {
		var ret string
		return ret
	}
	return *o.PaymentMemo
}

// GetPaymentMemoOk returns a tuple with the PaymentMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetPaymentMemoOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMemo) {
		return nil, false
	}
	return o.PaymentMemo, true
}

// HasPaymentMemo returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasPaymentMemo() bool {
	if o != nil && !IsNil(o.PaymentMemo) {
		return true
	}

	return false
}

// SetPaymentMemo gets a reference to the given string and assigns it to the PaymentMemo field.
func (o *PayorAmlTransaction) SetPaymentMemo(v string) {
	o.PaymentMemo = &v
}

// GetPaymentRails returns the PaymentRails field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetPaymentRails() string {
	if o == nil || IsNil(o.PaymentRails) {
		var ret string
		return ret
	}
	return *o.PaymentRails
}

// GetPaymentRailsOk returns a tuple with the PaymentRails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetPaymentRailsOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentRails) {
		return nil, false
	}
	return o.PaymentRails, true
}

// HasPaymentRails returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasPaymentRails() bool {
	if o != nil && !IsNil(o.PaymentRails) {
		return true
	}

	return false
}

// SetPaymentRails gets a reference to the given string and assigns it to the PaymentRails field.
func (o *PayorAmlTransaction) SetPaymentRails(v string) {
	o.PaymentRails = &v
}

// GetPayorPaymentId returns the PayorPaymentId field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetPayorPaymentId() string {
	if o == nil || IsNil(o.PayorPaymentId) {
		var ret string
		return ret
	}
	return *o.PayorPaymentId
}

// GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetPayorPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayorPaymentId) {
		return nil, false
	}
	return o.PayorPaymentId, true
}

// HasPayorPaymentId returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasPayorPaymentId() bool {
	if o != nil && !IsNil(o.PayorPaymentId) {
		return true
	}

	return false
}

// SetPayorPaymentId gets a reference to the given string and assigns it to the PayorPaymentId field.
func (o *PayorAmlTransaction) SetPayorPaymentId(v string) {
	o.PayorPaymentId = &v
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetPaymentStatus() string {
	if o == nil || IsNil(o.PaymentStatus) {
		var ret string
		return ret
	}
	return *o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetPaymentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentStatus) {
		return nil, false
	}
	return o.PaymentStatus, true
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasPaymentStatus() bool {
	if o != nil && !IsNil(o.PaymentStatus) {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given string and assigns it to the PaymentStatus field.
func (o *PayorAmlTransaction) SetPaymentStatus(v string) {
	o.PaymentStatus = &v
}

// GetRejectReason returns the RejectReason field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetRejectReason() string {
	if o == nil || IsNil(o.RejectReason) {
		var ret string
		return ret
	}
	return *o.RejectReason
}

// GetRejectReasonOk returns a tuple with the RejectReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetRejectReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RejectReason) {
		return nil, false
	}
	return o.RejectReason, true
}

// HasRejectReason returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasRejectReason() bool {
	if o != nil && !IsNil(o.RejectReason) {
		return true
	}

	return false
}

// SetRejectReason gets a reference to the given string and assigns it to the RejectReason field.
func (o *PayorAmlTransaction) SetRejectReason(v string) {
	o.RejectReason = &v
}

// GetFxApplied returns the FxApplied field value if set, zero value otherwise.
func (o *PayorAmlTransaction) GetFxApplied() float64 {
	if o == nil || IsNil(o.FxApplied) {
		var ret float64
		return ret
	}
	return *o.FxApplied
}

// GetFxAppliedOk returns a tuple with the FxApplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransaction) GetFxAppliedOk() (*float64, bool) {
	if o == nil || IsNil(o.FxApplied) {
		return nil, false
	}
	return o.FxApplied, true
}

// HasFxApplied returns a boolean if a field has been set.
func (o *PayorAmlTransaction) HasFxApplied() bool {
	if o != nil && !IsNil(o.FxApplied) {
		return true
	}

	return false
}

// SetFxApplied gets a reference to the given float64 and assigns it to the FxApplied field.
func (o *PayorAmlTransaction) SetFxApplied(v float64) {
	o.FxApplied = &v
}

func (o PayorAmlTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayorAmlTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionDate) {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if !IsNil(o.TransactionTime) {
		toSerialize["transactionTime"] = o.TransactionTime
	}
	if !IsNil(o.ReportTransactionType) {
		toSerialize["reportTransactionType"] = o.ReportTransactionType
	}
	if !IsNil(o.Debit) {
		toSerialize["debit"] = o.Debit
	}
	if !IsNil(o.DebitCurrency) {
		toSerialize["debitCurrency"] = o.DebitCurrency
	}
	if !IsNil(o.Credit) {
		toSerialize["credit"] = o.Credit
	}
	if !IsNil(o.CreditCurrency) {
		toSerialize["creditCurrency"] = o.CreditCurrency
	}
	if !IsNil(o.ReturnFee) {
		toSerialize["returnFee"] = o.ReturnFee
	}
	if !IsNil(o.ReturnFeeCurrency) {
		toSerialize["returnFeeCurrency"] = o.ReturnFeeCurrency
	}
	if !IsNil(o.ReturnFeeDescription) {
		toSerialize["returnFeeDescription"] = o.ReturnFeeDescription
	}
	if !IsNil(o.ReturnCode) {
		toSerialize["returnCode"] = o.ReturnCode
	}
	if !IsNil(o.ReturnDescription) {
		toSerialize["returnDescription"] = o.ReturnDescription
	}
	if !IsNil(o.FundingType) {
		toSerialize["fundingType"] = o.FundingType
	}
	if !IsNil(o.DateFundingRequested) {
		toSerialize["dateFundingRequested"] = o.DateFundingRequested
	}
	if !IsNil(o.PayeeName) {
		toSerialize["payeeName"] = o.PayeeName
	}
	if !IsNil(o.RemoteId) {
		toSerialize["remoteId"] = o.RemoteId
	}
	if !IsNil(o.PayeeType) {
		toSerialize["payeeType"] = o.PayeeType
	}
	if !IsNil(o.PayeeEmail) {
		toSerialize["payeeEmail"] = o.PayeeEmail
	}
	if !IsNil(o.SourceAccount) {
		toSerialize["sourceAccount"] = o.SourceAccount
	}
	if !IsNil(o.PaymentAmount) {
		toSerialize["paymentAmount"] = o.PaymentAmount
	}
	if !IsNil(o.PaymentCurrency) {
		toSerialize["paymentCurrency"] = o.PaymentCurrency
	}
	if !IsNil(o.PaymentMemo) {
		toSerialize["paymentMemo"] = o.PaymentMemo
	}
	if !IsNil(o.PaymentRails) {
		toSerialize["paymentRails"] = o.PaymentRails
	}
	if !IsNil(o.PayorPaymentId) {
		toSerialize["payorPaymentId"] = o.PayorPaymentId
	}
	if !IsNil(o.PaymentStatus) {
		toSerialize["paymentStatus"] = o.PaymentStatus
	}
	if !IsNil(o.RejectReason) {
		toSerialize["rejectReason"] = o.RejectReason
	}
	if !IsNil(o.FxApplied) {
		toSerialize["fxApplied"] = o.FxApplied
	}
	return toSerialize, nil
}

type NullablePayorAmlTransaction struct {
	value *PayorAmlTransaction
	isSet bool
}

func (v NullablePayorAmlTransaction) Get() *PayorAmlTransaction {
	return v.value
}

func (v *NullablePayorAmlTransaction) Set(val *PayorAmlTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePayorAmlTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePayorAmlTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayorAmlTransaction(val *PayorAmlTransaction) *NullablePayorAmlTransaction {
	return &NullablePayorAmlTransaction{value: val, isSet: true}
}

func (v NullablePayorAmlTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayorAmlTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


