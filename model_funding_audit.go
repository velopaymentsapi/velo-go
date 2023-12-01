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

// checks if the FundingAudit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FundingAudit{}

// FundingAudit struct for FundingAudit
type FundingAudit struct {
	// The id of the payor associated with the funding.
	PayorId *string `json:"payorId,omitempty"`
	// The amount funded
	Amount *float64 `json:"amount,omitempty"`
	// The currency of the funding
	Currency *string `json:"currency,omitempty"`
	DateTime *time.Time `json:"dateTime,omitempty"`
	// Status of the funding. One of the following values: PENDING, FAILED, CREDIT, DEBIT
	Status *string `json:"status,omitempty"`
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	FundingAccountName *string `json:"fundingAccountName,omitempty"`
	// Funding type. One of the following values: ACH, WIRE, EMBEDDED, BANK_TRANSFER
	FundingType *string `json:"fundingType,omitempty"`
	Events []FundingEvent2 `json:"events,omitempty"`
	// Type of top up. One of the following values: AUTOMATIC, MANUAL
	TopupType *string `json:"topupType,omitempty"`
	// The id of the transaction associated with the funding if there was one
	TransactionId *string `json:"transactionId,omitempty"`
	// The payors reference for the transaction associated with the funding if there was one
	TransactionReference *string `json:"transactionReference,omitempty"`
}

// NewFundingAudit instantiates a new FundingAudit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundingAudit() *FundingAudit {
	this := FundingAudit{}
	return &this
}

// NewFundingAuditWithDefaults instantiates a new FundingAudit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingAuditWithDefaults() *FundingAudit {
	this := FundingAudit{}
	return &this
}

// GetPayorId returns the PayorId field value if set, zero value otherwise.
func (o *FundingAudit) GetPayorId() string {
	if o == nil || IsNil(o.PayorId) {
		var ret string
		return ret
	}
	return *o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetPayorIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayorId) {
		return nil, false
	}
	return o.PayorId, true
}

// HasPayorId returns a boolean if a field has been set.
func (o *FundingAudit) HasPayorId() bool {
	if o != nil && !IsNil(o.PayorId) {
		return true
	}

	return false
}

// SetPayorId gets a reference to the given string and assigns it to the PayorId field.
func (o *FundingAudit) SetPayorId(v string) {
	o.PayorId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *FundingAudit) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *FundingAudit) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *FundingAudit) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *FundingAudit) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *FundingAudit) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *FundingAudit) SetCurrency(v string) {
	o.Currency = &v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *FundingAudit) GetDateTime() time.Time {
	if o == nil || IsNil(o.DateTime) {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTime) {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *FundingAudit) HasDateTime() bool {
	if o != nil && !IsNil(o.DateTime) {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *FundingAudit) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FundingAudit) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FundingAudit) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FundingAudit) SetStatus(v string) {
	o.Status = &v
}

// GetSourceAccountName returns the SourceAccountName field value if set, zero value otherwise.
func (o *FundingAudit) GetSourceAccountName() string {
	if o == nil || IsNil(o.SourceAccountName) {
		var ret string
		return ret
	}
	return *o.SourceAccountName
}

// GetSourceAccountNameOk returns a tuple with the SourceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetSourceAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAccountName) {
		return nil, false
	}
	return o.SourceAccountName, true
}

// HasSourceAccountName returns a boolean if a field has been set.
func (o *FundingAudit) HasSourceAccountName() bool {
	if o != nil && !IsNil(o.SourceAccountName) {
		return true
	}

	return false
}

// SetSourceAccountName gets a reference to the given string and assigns it to the SourceAccountName field.
func (o *FundingAudit) SetSourceAccountName(v string) {
	o.SourceAccountName = &v
}

// GetFundingAccountName returns the FundingAccountName field value if set, zero value otherwise.
func (o *FundingAudit) GetFundingAccountName() string {
	if o == nil || IsNil(o.FundingAccountName) {
		var ret string
		return ret
	}
	return *o.FundingAccountName
}

// GetFundingAccountNameOk returns a tuple with the FundingAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetFundingAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.FundingAccountName) {
		return nil, false
	}
	return o.FundingAccountName, true
}

// HasFundingAccountName returns a boolean if a field has been set.
func (o *FundingAudit) HasFundingAccountName() bool {
	if o != nil && !IsNil(o.FundingAccountName) {
		return true
	}

	return false
}

// SetFundingAccountName gets a reference to the given string and assigns it to the FundingAccountName field.
func (o *FundingAudit) SetFundingAccountName(v string) {
	o.FundingAccountName = &v
}

// GetFundingType returns the FundingType field value if set, zero value otherwise.
func (o *FundingAudit) GetFundingType() string {
	if o == nil || IsNil(o.FundingType) {
		var ret string
		return ret
	}
	return *o.FundingType
}

// GetFundingTypeOk returns a tuple with the FundingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetFundingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FundingType) {
		return nil, false
	}
	return o.FundingType, true
}

// HasFundingType returns a boolean if a field has been set.
func (o *FundingAudit) HasFundingType() bool {
	if o != nil && !IsNil(o.FundingType) {
		return true
	}

	return false
}

// SetFundingType gets a reference to the given string and assigns it to the FundingType field.
func (o *FundingAudit) SetFundingType(v string) {
	o.FundingType = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *FundingAudit) GetEvents() []FundingEvent2 {
	if o == nil || IsNil(o.Events) {
		var ret []FundingEvent2
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetEventsOk() ([]FundingEvent2, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *FundingAudit) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []FundingEvent2 and assigns it to the Events field.
func (o *FundingAudit) SetEvents(v []FundingEvent2) {
	o.Events = v
}

// GetTopupType returns the TopupType field value if set, zero value otherwise.
func (o *FundingAudit) GetTopupType() string {
	if o == nil || IsNil(o.TopupType) {
		var ret string
		return ret
	}
	return *o.TopupType
}

// GetTopupTypeOk returns a tuple with the TopupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetTopupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TopupType) {
		return nil, false
	}
	return o.TopupType, true
}

// HasTopupType returns a boolean if a field has been set.
func (o *FundingAudit) HasTopupType() bool {
	if o != nil && !IsNil(o.TopupType) {
		return true
	}

	return false
}

// SetTopupType gets a reference to the given string and assigns it to the TopupType field.
func (o *FundingAudit) SetTopupType(v string) {
	o.TopupType = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *FundingAudit) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *FundingAudit) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *FundingAudit) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionReference returns the TransactionReference field value if set, zero value otherwise.
func (o *FundingAudit) GetTransactionReference() string {
	if o == nil || IsNil(o.TransactionReference) {
		var ret string
		return ret
	}
	return *o.TransactionReference
}

// GetTransactionReferenceOk returns a tuple with the TransactionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAudit) GetTransactionReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionReference) {
		return nil, false
	}
	return o.TransactionReference, true
}

// HasTransactionReference returns a boolean if a field has been set.
func (o *FundingAudit) HasTransactionReference() bool {
	if o != nil && !IsNil(o.TransactionReference) {
		return true
	}

	return false
}

// SetTransactionReference gets a reference to the given string and assigns it to the TransactionReference field.
func (o *FundingAudit) SetTransactionReference(v string) {
	o.TransactionReference = &v
}

func (o FundingAudit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundingAudit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayorId) {
		toSerialize["payorId"] = o.PayorId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.DateTime) {
		toSerialize["dateTime"] = o.DateTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SourceAccountName) {
		toSerialize["sourceAccountName"] = o.SourceAccountName
	}
	if !IsNil(o.FundingAccountName) {
		toSerialize["fundingAccountName"] = o.FundingAccountName
	}
	if !IsNil(o.FundingType) {
		toSerialize["fundingType"] = o.FundingType
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.TopupType) {
		toSerialize["topupType"] = o.TopupType
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !IsNil(o.TransactionReference) {
		toSerialize["transactionReference"] = o.TransactionReference
	}
	return toSerialize, nil
}

type NullableFundingAudit struct {
	value *FundingAudit
	isSet bool
}

func (v NullableFundingAudit) Get() *FundingAudit {
	return v.value
}

func (v *NullableFundingAudit) Set(val *FundingAudit) {
	v.value = val
	v.isSet = true
}

func (v NullableFundingAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableFundingAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundingAudit(val *FundingAudit) *NullableFundingAudit {
	return &NullableFundingAudit{value: val, isSet: true}
}

func (v NullableFundingAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundingAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


