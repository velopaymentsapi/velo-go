/*
 * Velo Payments APIs
 *
 * ## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 
 *
 * API version: 2.26.124
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// PayorAmlTransactionV3 struct for PayorAmlTransactionV3
type PayorAmlTransactionV3 struct {
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
	// Remote ID of the Payee, set by Payor
	RemoteId *string `json:"remoteId,omitempty"`
	PayeeType *string `json:"payeeType,omitempty"`
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

// NewPayorAmlTransactionV3 instantiates a new PayorAmlTransactionV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayorAmlTransactionV3() *PayorAmlTransactionV3 {
	this := PayorAmlTransactionV3{}
	return &this
}

// NewPayorAmlTransactionV3WithDefaults instantiates a new PayorAmlTransactionV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayorAmlTransactionV3WithDefaults() *PayorAmlTransactionV3 {
	this := PayorAmlTransactionV3{}
	return &this
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetTransactionDate() string {
	if o == nil || o.TransactionDate == nil {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetTransactionDateOk() (*string, bool) {
	if o == nil || o.TransactionDate == nil {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasTransactionDate() bool {
	if o != nil && o.TransactionDate != nil {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *PayorAmlTransactionV3) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetTransactionTime() string {
	if o == nil || o.TransactionTime == nil {
		var ret string
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetTransactionTimeOk() (*string, bool) {
	if o == nil || o.TransactionTime == nil {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasTransactionTime() bool {
	if o != nil && o.TransactionTime != nil {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given string and assigns it to the TransactionTime field.
func (o *PayorAmlTransactionV3) SetTransactionTime(v string) {
	o.TransactionTime = &v
}

// GetReportTransactionType returns the ReportTransactionType field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetReportTransactionType() string {
	if o == nil || o.ReportTransactionType == nil {
		var ret string
		return ret
	}
	return *o.ReportTransactionType
}

// GetReportTransactionTypeOk returns a tuple with the ReportTransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetReportTransactionTypeOk() (*string, bool) {
	if o == nil || o.ReportTransactionType == nil {
		return nil, false
	}
	return o.ReportTransactionType, true
}

// HasReportTransactionType returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasReportTransactionType() bool {
	if o != nil && o.ReportTransactionType != nil {
		return true
	}

	return false
}

// SetReportTransactionType gets a reference to the given string and assigns it to the ReportTransactionType field.
func (o *PayorAmlTransactionV3) SetReportTransactionType(v string) {
	o.ReportTransactionType = &v
}

// GetDebit returns the Debit field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetDebit() int64 {
	if o == nil || o.Debit == nil {
		var ret int64
		return ret
	}
	return *o.Debit
}

// GetDebitOk returns a tuple with the Debit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetDebitOk() (*int64, bool) {
	if o == nil || o.Debit == nil {
		return nil, false
	}
	return o.Debit, true
}

// HasDebit returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasDebit() bool {
	if o != nil && o.Debit != nil {
		return true
	}

	return false
}

// SetDebit gets a reference to the given int64 and assigns it to the Debit field.
func (o *PayorAmlTransactionV3) SetDebit(v int64) {
	o.Debit = &v
}

// GetDebitCurrency returns the DebitCurrency field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetDebitCurrency() string {
	if o == nil || o.DebitCurrency == nil {
		var ret string
		return ret
	}
	return *o.DebitCurrency
}

// GetDebitCurrencyOk returns a tuple with the DebitCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetDebitCurrencyOk() (*string, bool) {
	if o == nil || o.DebitCurrency == nil {
		return nil, false
	}
	return o.DebitCurrency, true
}

// HasDebitCurrency returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasDebitCurrency() bool {
	if o != nil && o.DebitCurrency != nil {
		return true
	}

	return false
}

// SetDebitCurrency gets a reference to the given string and assigns it to the DebitCurrency field.
func (o *PayorAmlTransactionV3) SetDebitCurrency(v string) {
	o.DebitCurrency = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetCredit() int64 {
	if o == nil || o.Credit == nil {
		var ret int64
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetCreditOk() (*int64, bool) {
	if o == nil || o.Credit == nil {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasCredit() bool {
	if o != nil && o.Credit != nil {
		return true
	}

	return false
}

// SetCredit gets a reference to the given int64 and assigns it to the Credit field.
func (o *PayorAmlTransactionV3) SetCredit(v int64) {
	o.Credit = &v
}

// GetCreditCurrency returns the CreditCurrency field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetCreditCurrency() string {
	if o == nil || o.CreditCurrency == nil {
		var ret string
		return ret
	}
	return *o.CreditCurrency
}

// GetCreditCurrencyOk returns a tuple with the CreditCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetCreditCurrencyOk() (*string, bool) {
	if o == nil || o.CreditCurrency == nil {
		return nil, false
	}
	return o.CreditCurrency, true
}

// HasCreditCurrency returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasCreditCurrency() bool {
	if o != nil && o.CreditCurrency != nil {
		return true
	}

	return false
}

// SetCreditCurrency gets a reference to the given string and assigns it to the CreditCurrency field.
func (o *PayorAmlTransactionV3) SetCreditCurrency(v string) {
	o.CreditCurrency = &v
}

// GetReturnFee returns the ReturnFee field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetReturnFee() string {
	if o == nil || o.ReturnFee == nil {
		var ret string
		return ret
	}
	return *o.ReturnFee
}

// GetReturnFeeOk returns a tuple with the ReturnFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetReturnFeeOk() (*string, bool) {
	if o == nil || o.ReturnFee == nil {
		return nil, false
	}
	return o.ReturnFee, true
}

// HasReturnFee returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasReturnFee() bool {
	if o != nil && o.ReturnFee != nil {
		return true
	}

	return false
}

// SetReturnFee gets a reference to the given string and assigns it to the ReturnFee field.
func (o *PayorAmlTransactionV3) SetReturnFee(v string) {
	o.ReturnFee = &v
}

// GetReturnFeeCurrency returns the ReturnFeeCurrency field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetReturnFeeCurrency() string {
	if o == nil || o.ReturnFeeCurrency == nil {
		var ret string
		return ret
	}
	return *o.ReturnFeeCurrency
}

// GetReturnFeeCurrencyOk returns a tuple with the ReturnFeeCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetReturnFeeCurrencyOk() (*string, bool) {
	if o == nil || o.ReturnFeeCurrency == nil {
		return nil, false
	}
	return o.ReturnFeeCurrency, true
}

// HasReturnFeeCurrency returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasReturnFeeCurrency() bool {
	if o != nil && o.ReturnFeeCurrency != nil {
		return true
	}

	return false
}

// SetReturnFeeCurrency gets a reference to the given string and assigns it to the ReturnFeeCurrency field.
func (o *PayorAmlTransactionV3) SetReturnFeeCurrency(v string) {
	o.ReturnFeeCurrency = &v
}

// GetReturnFeeDescription returns the ReturnFeeDescription field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetReturnFeeDescription() string {
	if o == nil || o.ReturnFeeDescription == nil {
		var ret string
		return ret
	}
	return *o.ReturnFeeDescription
}

// GetReturnFeeDescriptionOk returns a tuple with the ReturnFeeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetReturnFeeDescriptionOk() (*string, bool) {
	if o == nil || o.ReturnFeeDescription == nil {
		return nil, false
	}
	return o.ReturnFeeDescription, true
}

// HasReturnFeeDescription returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasReturnFeeDescription() bool {
	if o != nil && o.ReturnFeeDescription != nil {
		return true
	}

	return false
}

// SetReturnFeeDescription gets a reference to the given string and assigns it to the ReturnFeeDescription field.
func (o *PayorAmlTransactionV3) SetReturnFeeDescription(v string) {
	o.ReturnFeeDescription = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetReturnCode() string {
	if o == nil || o.ReturnCode == nil {
		var ret string
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetReturnCodeOk() (*string, bool) {
	if o == nil || o.ReturnCode == nil {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasReturnCode() bool {
	if o != nil && o.ReturnCode != nil {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given string and assigns it to the ReturnCode field.
func (o *PayorAmlTransactionV3) SetReturnCode(v string) {
	o.ReturnCode = &v
}

// GetReturnDescription returns the ReturnDescription field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetReturnDescription() string {
	if o == nil || o.ReturnDescription == nil {
		var ret string
		return ret
	}
	return *o.ReturnDescription
}

// GetReturnDescriptionOk returns a tuple with the ReturnDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetReturnDescriptionOk() (*string, bool) {
	if o == nil || o.ReturnDescription == nil {
		return nil, false
	}
	return o.ReturnDescription, true
}

// HasReturnDescription returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasReturnDescription() bool {
	if o != nil && o.ReturnDescription != nil {
		return true
	}

	return false
}

// SetReturnDescription gets a reference to the given string and assigns it to the ReturnDescription field.
func (o *PayorAmlTransactionV3) SetReturnDescription(v string) {
	o.ReturnDescription = &v
}

// GetFundingType returns the FundingType field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetFundingType() string {
	if o == nil || o.FundingType == nil {
		var ret string
		return ret
	}
	return *o.FundingType
}

// GetFundingTypeOk returns a tuple with the FundingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetFundingTypeOk() (*string, bool) {
	if o == nil || o.FundingType == nil {
		return nil, false
	}
	return o.FundingType, true
}

// HasFundingType returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasFundingType() bool {
	if o != nil && o.FundingType != nil {
		return true
	}

	return false
}

// SetFundingType gets a reference to the given string and assigns it to the FundingType field.
func (o *PayorAmlTransactionV3) SetFundingType(v string) {
	o.FundingType = &v
}

// GetDateFundingRequested returns the DateFundingRequested field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetDateFundingRequested() string {
	if o == nil || o.DateFundingRequested == nil {
		var ret string
		return ret
	}
	return *o.DateFundingRequested
}

// GetDateFundingRequestedOk returns a tuple with the DateFundingRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetDateFundingRequestedOk() (*string, bool) {
	if o == nil || o.DateFundingRequested == nil {
		return nil, false
	}
	return o.DateFundingRequested, true
}

// HasDateFundingRequested returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasDateFundingRequested() bool {
	if o != nil && o.DateFundingRequested != nil {
		return true
	}

	return false
}

// SetDateFundingRequested gets a reference to the given string and assigns it to the DateFundingRequested field.
func (o *PayorAmlTransactionV3) SetDateFundingRequested(v string) {
	o.DateFundingRequested = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetRemoteId() string {
	if o == nil || o.RemoteId == nil {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetRemoteIdOk() (*string, bool) {
	if o == nil || o.RemoteId == nil {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasRemoteId() bool {
	if o != nil && o.RemoteId != nil {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *PayorAmlTransactionV3) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetPayeeType returns the PayeeType field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetPayeeType() string {
	if o == nil || o.PayeeType == nil {
		var ret string
		return ret
	}
	return *o.PayeeType
}

// GetPayeeTypeOk returns a tuple with the PayeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetPayeeTypeOk() (*string, bool) {
	if o == nil || o.PayeeType == nil {
		return nil, false
	}
	return o.PayeeType, true
}

// HasPayeeType returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasPayeeType() bool {
	if o != nil && o.PayeeType != nil {
		return true
	}

	return false
}

// SetPayeeType gets a reference to the given string and assigns it to the PayeeType field.
func (o *PayorAmlTransactionV3) SetPayeeType(v string) {
	o.PayeeType = &v
}

// GetSourceAccount returns the SourceAccount field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetSourceAccount() string {
	if o == nil || o.SourceAccount == nil {
		var ret string
		return ret
	}
	return *o.SourceAccount
}

// GetSourceAccountOk returns a tuple with the SourceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetSourceAccountOk() (*string, bool) {
	if o == nil || o.SourceAccount == nil {
		return nil, false
	}
	return o.SourceAccount, true
}

// HasSourceAccount returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasSourceAccount() bool {
	if o != nil && o.SourceAccount != nil {
		return true
	}

	return false
}

// SetSourceAccount gets a reference to the given string and assigns it to the SourceAccount field.
func (o *PayorAmlTransactionV3) SetSourceAccount(v string) {
	o.SourceAccount = &v
}

// GetPaymentAmount returns the PaymentAmount field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetPaymentAmount() int64 {
	if o == nil || o.PaymentAmount == nil {
		var ret int64
		return ret
	}
	return *o.PaymentAmount
}

// GetPaymentAmountOk returns a tuple with the PaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetPaymentAmountOk() (*int64, bool) {
	if o == nil || o.PaymentAmount == nil {
		return nil, false
	}
	return o.PaymentAmount, true
}

// HasPaymentAmount returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasPaymentAmount() bool {
	if o != nil && o.PaymentAmount != nil {
		return true
	}

	return false
}

// SetPaymentAmount gets a reference to the given int64 and assigns it to the PaymentAmount field.
func (o *PayorAmlTransactionV3) SetPaymentAmount(v int64) {
	o.PaymentAmount = &v
}

// GetPaymentCurrency returns the PaymentCurrency field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetPaymentCurrency() string {
	if o == nil || o.PaymentCurrency == nil {
		var ret string
		return ret
	}
	return *o.PaymentCurrency
}

// GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetPaymentCurrencyOk() (*string, bool) {
	if o == nil || o.PaymentCurrency == nil {
		return nil, false
	}
	return o.PaymentCurrency, true
}

// HasPaymentCurrency returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasPaymentCurrency() bool {
	if o != nil && o.PaymentCurrency != nil {
		return true
	}

	return false
}

// SetPaymentCurrency gets a reference to the given string and assigns it to the PaymentCurrency field.
func (o *PayorAmlTransactionV3) SetPaymentCurrency(v string) {
	o.PaymentCurrency = &v
}

// GetPaymentMemo returns the PaymentMemo field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetPaymentMemo() string {
	if o == nil || o.PaymentMemo == nil {
		var ret string
		return ret
	}
	return *o.PaymentMemo
}

// GetPaymentMemoOk returns a tuple with the PaymentMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetPaymentMemoOk() (*string, bool) {
	if o == nil || o.PaymentMemo == nil {
		return nil, false
	}
	return o.PaymentMemo, true
}

// HasPaymentMemo returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasPaymentMemo() bool {
	if o != nil && o.PaymentMemo != nil {
		return true
	}

	return false
}

// SetPaymentMemo gets a reference to the given string and assigns it to the PaymentMemo field.
func (o *PayorAmlTransactionV3) SetPaymentMemo(v string) {
	o.PaymentMemo = &v
}

// GetPaymentRails returns the PaymentRails field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetPaymentRails() string {
	if o == nil || o.PaymentRails == nil {
		var ret string
		return ret
	}
	return *o.PaymentRails
}

// GetPaymentRailsOk returns a tuple with the PaymentRails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetPaymentRailsOk() (*string, bool) {
	if o == nil || o.PaymentRails == nil {
		return nil, false
	}
	return o.PaymentRails, true
}

// HasPaymentRails returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasPaymentRails() bool {
	if o != nil && o.PaymentRails != nil {
		return true
	}

	return false
}

// SetPaymentRails gets a reference to the given string and assigns it to the PaymentRails field.
func (o *PayorAmlTransactionV3) SetPaymentRails(v string) {
	o.PaymentRails = &v
}

// GetPayorPaymentId returns the PayorPaymentId field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetPayorPaymentId() string {
	if o == nil || o.PayorPaymentId == nil {
		var ret string
		return ret
	}
	return *o.PayorPaymentId
}

// GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetPayorPaymentIdOk() (*string, bool) {
	if o == nil || o.PayorPaymentId == nil {
		return nil, false
	}
	return o.PayorPaymentId, true
}

// HasPayorPaymentId returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasPayorPaymentId() bool {
	if o != nil && o.PayorPaymentId != nil {
		return true
	}

	return false
}

// SetPayorPaymentId gets a reference to the given string and assigns it to the PayorPaymentId field.
func (o *PayorAmlTransactionV3) SetPayorPaymentId(v string) {
	o.PayorPaymentId = &v
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetPaymentStatus() string {
	if o == nil || o.PaymentStatus == nil {
		var ret string
		return ret
	}
	return *o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetPaymentStatusOk() (*string, bool) {
	if o == nil || o.PaymentStatus == nil {
		return nil, false
	}
	return o.PaymentStatus, true
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasPaymentStatus() bool {
	if o != nil && o.PaymentStatus != nil {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given string and assigns it to the PaymentStatus field.
func (o *PayorAmlTransactionV3) SetPaymentStatus(v string) {
	o.PaymentStatus = &v
}

// GetRejectReason returns the RejectReason field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetRejectReason() string {
	if o == nil || o.RejectReason == nil {
		var ret string
		return ret
	}
	return *o.RejectReason
}

// GetRejectReasonOk returns a tuple with the RejectReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetRejectReasonOk() (*string, bool) {
	if o == nil || o.RejectReason == nil {
		return nil, false
	}
	return o.RejectReason, true
}

// HasRejectReason returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasRejectReason() bool {
	if o != nil && o.RejectReason != nil {
		return true
	}

	return false
}

// SetRejectReason gets a reference to the given string and assigns it to the RejectReason field.
func (o *PayorAmlTransactionV3) SetRejectReason(v string) {
	o.RejectReason = &v
}

// GetFxApplied returns the FxApplied field value if set, zero value otherwise.
func (o *PayorAmlTransactionV3) GetFxApplied() float64 {
	if o == nil || o.FxApplied == nil {
		var ret float64
		return ret
	}
	return *o.FxApplied
}

// GetFxAppliedOk returns a tuple with the FxApplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAmlTransactionV3) GetFxAppliedOk() (*float64, bool) {
	if o == nil || o.FxApplied == nil {
		return nil, false
	}
	return o.FxApplied, true
}

// HasFxApplied returns a boolean if a field has been set.
func (o *PayorAmlTransactionV3) HasFxApplied() bool {
	if o != nil && o.FxApplied != nil {
		return true
	}

	return false
}

// SetFxApplied gets a reference to the given float64 and assigns it to the FxApplied field.
func (o *PayorAmlTransactionV3) SetFxApplied(v float64) {
	o.FxApplied = &v
}

func (o PayorAmlTransactionV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransactionDate != nil {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if o.TransactionTime != nil {
		toSerialize["transactionTime"] = o.TransactionTime
	}
	if o.ReportTransactionType != nil {
		toSerialize["reportTransactionType"] = o.ReportTransactionType
	}
	if o.Debit != nil {
		toSerialize["debit"] = o.Debit
	}
	if o.DebitCurrency != nil {
		toSerialize["debitCurrency"] = o.DebitCurrency
	}
	if o.Credit != nil {
		toSerialize["credit"] = o.Credit
	}
	if o.CreditCurrency != nil {
		toSerialize["creditCurrency"] = o.CreditCurrency
	}
	if o.ReturnFee != nil {
		toSerialize["returnFee"] = o.ReturnFee
	}
	if o.ReturnFeeCurrency != nil {
		toSerialize["returnFeeCurrency"] = o.ReturnFeeCurrency
	}
	if o.ReturnFeeDescription != nil {
		toSerialize["returnFeeDescription"] = o.ReturnFeeDescription
	}
	if o.ReturnCode != nil {
		toSerialize["returnCode"] = o.ReturnCode
	}
	if o.ReturnDescription != nil {
		toSerialize["returnDescription"] = o.ReturnDescription
	}
	if o.FundingType != nil {
		toSerialize["fundingType"] = o.FundingType
	}
	if o.DateFundingRequested != nil {
		toSerialize["dateFundingRequested"] = o.DateFundingRequested
	}
	if o.RemoteId != nil {
		toSerialize["remoteId"] = o.RemoteId
	}
	if o.PayeeType != nil {
		toSerialize["payeeType"] = o.PayeeType
	}
	if o.SourceAccount != nil {
		toSerialize["sourceAccount"] = o.SourceAccount
	}
	if o.PaymentAmount != nil {
		toSerialize["paymentAmount"] = o.PaymentAmount
	}
	if o.PaymentCurrency != nil {
		toSerialize["paymentCurrency"] = o.PaymentCurrency
	}
	if o.PaymentMemo != nil {
		toSerialize["paymentMemo"] = o.PaymentMemo
	}
	if o.PaymentRails != nil {
		toSerialize["paymentRails"] = o.PaymentRails
	}
	if o.PayorPaymentId != nil {
		toSerialize["payorPaymentId"] = o.PayorPaymentId
	}
	if o.PaymentStatus != nil {
		toSerialize["paymentStatus"] = o.PaymentStatus
	}
	if o.RejectReason != nil {
		toSerialize["rejectReason"] = o.RejectReason
	}
	if o.FxApplied != nil {
		toSerialize["fxApplied"] = o.FxApplied
	}
	return json.Marshal(toSerialize)
}

type NullablePayorAmlTransactionV3 struct {
	value *PayorAmlTransactionV3
	isSet bool
}

func (v NullablePayorAmlTransactionV3) Get() *PayorAmlTransactionV3 {
	return v.value
}

func (v *NullablePayorAmlTransactionV3) Set(val *PayorAmlTransactionV3) {
	v.value = val
	v.isSet = true
}

func (v NullablePayorAmlTransactionV3) IsSet() bool {
	return v.isSet
}

func (v *NullablePayorAmlTransactionV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayorAmlTransactionV3(val *PayorAmlTransactionV3) *NullablePayorAmlTransactionV3 {
	return &NullablePayorAmlTransactionV3{value: val, isSet: true}
}

func (v NullablePayorAmlTransactionV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayorAmlTransactionV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


