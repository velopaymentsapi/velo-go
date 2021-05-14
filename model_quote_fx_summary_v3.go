/*
 * Velo Payments APIs
 *
 * ## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 
 *
 * API version: 2.26.127
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
	"time"
)

// QuoteFxSummaryV3 struct for QuoteFxSummaryV3
type QuoteFxSummaryV3 struct {
	Rate float32 `json:"rate"`
	InvertedRate *float32 `json:"invertedRate,omitempty"`
	CreationTime time.Time `json:"creationTime"`
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	QuoteId string `json:"quoteId"`
	TotalSourceAmount int32 `json:"totalSourceAmount"`
	TotalPaymentAmount int32 `json:"totalPaymentAmount"`
	// Valid ISO 4217 3 letter currency code. See the <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	SourceCurrency string `json:"sourceCurrency"`
	// Valid ISO 4217 3 letter currency code. See the <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	PaymentCurrency string `json:"paymentCurrency"`
	FundingStatus string `json:"fundingStatus"`
	Status string `json:"status"`
}

// NewQuoteFxSummaryV3 instantiates a new QuoteFxSummaryV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteFxSummaryV3(rate float32, creationTime time.Time, quoteId string, totalSourceAmount int32, totalPaymentAmount int32, sourceCurrency string, paymentCurrency string, fundingStatus string, status string) *QuoteFxSummaryV3 {
	this := QuoteFxSummaryV3{}
	this.Rate = rate
	this.CreationTime = creationTime
	this.QuoteId = quoteId
	this.TotalSourceAmount = totalSourceAmount
	this.TotalPaymentAmount = totalPaymentAmount
	this.SourceCurrency = sourceCurrency
	this.PaymentCurrency = paymentCurrency
	this.FundingStatus = fundingStatus
	this.Status = status
	return &this
}

// NewQuoteFxSummaryV3WithDefaults instantiates a new QuoteFxSummaryV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteFxSummaryV3WithDefaults() *QuoteFxSummaryV3 {
	this := QuoteFxSummaryV3{}
	return &this
}

// GetRate returns the Rate field value
func (o *QuoteFxSummaryV3) GetRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetRateOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *QuoteFxSummaryV3) SetRate(v float32) {
	o.Rate = v
}

// GetInvertedRate returns the InvertedRate field value if set, zero value otherwise.
func (o *QuoteFxSummaryV3) GetInvertedRate() float32 {
	if o == nil || o.InvertedRate == nil {
		var ret float32
		return ret
	}
	return *o.InvertedRate
}

// GetInvertedRateOk returns a tuple with the InvertedRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetInvertedRateOk() (*float32, bool) {
	if o == nil || o.InvertedRate == nil {
		return nil, false
	}
	return o.InvertedRate, true
}

// HasInvertedRate returns a boolean if a field has been set.
func (o *QuoteFxSummaryV3) HasInvertedRate() bool {
	if o != nil && o.InvertedRate != nil {
		return true
	}

	return false
}

// SetInvertedRate gets a reference to the given float32 and assigns it to the InvertedRate field.
func (o *QuoteFxSummaryV3) SetInvertedRate(v float32) {
	o.InvertedRate = &v
}

// GetCreationTime returns the CreationTime field value
func (o *QuoteFxSummaryV3) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *QuoteFxSummaryV3) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *QuoteFxSummaryV3) GetExpiryTime() time.Time {
	if o == nil || o.ExpiryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetExpiryTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpiryTime == nil {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *QuoteFxSummaryV3) HasExpiryTime() bool {
	if o != nil && o.ExpiryTime != nil {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given time.Time and assigns it to the ExpiryTime field.
func (o *QuoteFxSummaryV3) SetExpiryTime(v time.Time) {
	o.ExpiryTime = &v
}

// GetQuoteId returns the QuoteId field value
func (o *QuoteFxSummaryV3) GetQuoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetQuoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QuoteId, true
}

// SetQuoteId sets field value
func (o *QuoteFxSummaryV3) SetQuoteId(v string) {
	o.QuoteId = v
}

// GetTotalSourceAmount returns the TotalSourceAmount field value
func (o *QuoteFxSummaryV3) GetTotalSourceAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalSourceAmount
}

// GetTotalSourceAmountOk returns a tuple with the TotalSourceAmount field value
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetTotalSourceAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalSourceAmount, true
}

// SetTotalSourceAmount sets field value
func (o *QuoteFxSummaryV3) SetTotalSourceAmount(v int32) {
	o.TotalSourceAmount = v
}

// GetTotalPaymentAmount returns the TotalPaymentAmount field value
func (o *QuoteFxSummaryV3) GetTotalPaymentAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPaymentAmount
}

// GetTotalPaymentAmountOk returns a tuple with the TotalPaymentAmount field value
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetTotalPaymentAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalPaymentAmount, true
}

// SetTotalPaymentAmount sets field value
func (o *QuoteFxSummaryV3) SetTotalPaymentAmount(v int32) {
	o.TotalPaymentAmount = v
}

// GetSourceCurrency returns the SourceCurrency field value
func (o *QuoteFxSummaryV3) GetSourceCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceCurrency
}

// GetSourceCurrencyOk returns a tuple with the SourceCurrency field value
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetSourceCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceCurrency, true
}

// SetSourceCurrency sets field value
func (o *QuoteFxSummaryV3) SetSourceCurrency(v string) {
	o.SourceCurrency = v
}

// GetPaymentCurrency returns the PaymentCurrency field value
func (o *QuoteFxSummaryV3) GetPaymentCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentCurrency
}

// GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field value
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetPaymentCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentCurrency, true
}

// SetPaymentCurrency sets field value
func (o *QuoteFxSummaryV3) SetPaymentCurrency(v string) {
	o.PaymentCurrency = v
}

// GetFundingStatus returns the FundingStatus field value
func (o *QuoteFxSummaryV3) GetFundingStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingStatus
}

// GetFundingStatusOk returns a tuple with the FundingStatus field value
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetFundingStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FundingStatus, true
}

// SetFundingStatus sets field value
func (o *QuoteFxSummaryV3) SetFundingStatus(v string) {
	o.FundingStatus = v
}

// GetStatus returns the Status field value
func (o *QuoteFxSummaryV3) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *QuoteFxSummaryV3) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *QuoteFxSummaryV3) SetStatus(v string) {
	o.Status = v
}

func (o QuoteFxSummaryV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rate"] = o.Rate
	}
	if o.InvertedRate != nil {
		toSerialize["invertedRate"] = o.InvertedRate
	}
	if true {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.ExpiryTime != nil {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if true {
		toSerialize["quoteId"] = o.QuoteId
	}
	if true {
		toSerialize["totalSourceAmount"] = o.TotalSourceAmount
	}
	if true {
		toSerialize["totalPaymentAmount"] = o.TotalPaymentAmount
	}
	if true {
		toSerialize["sourceCurrency"] = o.SourceCurrency
	}
	if true {
		toSerialize["paymentCurrency"] = o.PaymentCurrency
	}
	if true {
		toSerialize["fundingStatus"] = o.FundingStatus
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteFxSummaryV3 struct {
	value *QuoteFxSummaryV3
	isSet bool
}

func (v NullableQuoteFxSummaryV3) Get() *QuoteFxSummaryV3 {
	return v.value
}

func (v *NullableQuoteFxSummaryV3) Set(val *QuoteFxSummaryV3) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteFxSummaryV3) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteFxSummaryV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteFxSummaryV3(val *QuoteFxSummaryV3) *NullableQuoteFxSummaryV3 {
	return &NullableQuoteFxSummaryV3{value: val, isSet: true}
}

func (v NullableQuoteFxSummaryV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteFxSummaryV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


