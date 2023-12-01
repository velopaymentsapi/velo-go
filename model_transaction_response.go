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

// checks if the TransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionResponse{}

// TransactionResponse struct for TransactionResponse
type TransactionResponse struct {
	// The Id of the transaction 
	TransactionId string `json:"transactionId"`
	// The short code for the transaction that can be used when sending funds into the platform 
	TransactionShortCode string `json:"transactionShortCode"`
	// Indicates the Payor of the Transaction and which matches the payorId on the provided source account
	PayorId *string `json:"payorId,omitempty"`
	// The name of the source account associated with the Transaction 
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	// The Id of the source account associated with the Transaction 
	SourceAccountId *string `json:"sourceAccountId,omitempty"`
	// The payors own reference for the transaction 
	TransactionReference *string `json:"transactionReference,omitempty"`
	// Optional metadata attached to the transaction at creation time. 
	TransactionMetadata map[string]string `json:"transactionMetadata,omitempty"`
	// The total amount of transaction in minor units 
	Balance *int64 `json:"balance,omitempty"`
	// Valid ISO 4217 3 letter currency code. See the <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	Currency *string `json:"currency,omitempty"`
	// A timestamp when the transaction has been created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewTransactionResponse instantiates a new TransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResponse(transactionId string, transactionShortCode string) *TransactionResponse {
	this := TransactionResponse{}
	this.TransactionId = transactionId
	this.TransactionShortCode = transactionShortCode
	return &this
}

// NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResponseWithDefaults() *TransactionResponse {
	this := TransactionResponse{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *TransactionResponse) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TransactionResponse) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetTransactionShortCode returns the TransactionShortCode field value
func (o *TransactionResponse) GetTransactionShortCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionShortCode
}

// GetTransactionShortCodeOk returns a tuple with the TransactionShortCode field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetTransactionShortCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionShortCode, true
}

// SetTransactionShortCode sets field value
func (o *TransactionResponse) SetTransactionShortCode(v string) {
	o.TransactionShortCode = v
}

// GetPayorId returns the PayorId field value if set, zero value otherwise.
func (o *TransactionResponse) GetPayorId() string {
	if o == nil || IsNil(o.PayorId) {
		var ret string
		return ret
	}
	return *o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetPayorIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayorId) {
		return nil, false
	}
	return o.PayorId, true
}

// HasPayorId returns a boolean if a field has been set.
func (o *TransactionResponse) HasPayorId() bool {
	if o != nil && !IsNil(o.PayorId) {
		return true
	}

	return false
}

// SetPayorId gets a reference to the given string and assigns it to the PayorId field.
func (o *TransactionResponse) SetPayorId(v string) {
	o.PayorId = &v
}

// GetSourceAccountName returns the SourceAccountName field value if set, zero value otherwise.
func (o *TransactionResponse) GetSourceAccountName() string {
	if o == nil || IsNil(o.SourceAccountName) {
		var ret string
		return ret
	}
	return *o.SourceAccountName
}

// GetSourceAccountNameOk returns a tuple with the SourceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetSourceAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAccountName) {
		return nil, false
	}
	return o.SourceAccountName, true
}

// HasSourceAccountName returns a boolean if a field has been set.
func (o *TransactionResponse) HasSourceAccountName() bool {
	if o != nil && !IsNil(o.SourceAccountName) {
		return true
	}

	return false
}

// SetSourceAccountName gets a reference to the given string and assigns it to the SourceAccountName field.
func (o *TransactionResponse) SetSourceAccountName(v string) {
	o.SourceAccountName = &v
}

// GetSourceAccountId returns the SourceAccountId field value if set, zero value otherwise.
func (o *TransactionResponse) GetSourceAccountId() string {
	if o == nil || IsNil(o.SourceAccountId) {
		var ret string
		return ret
	}
	return *o.SourceAccountId
}

// GetSourceAccountIdOk returns a tuple with the SourceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetSourceAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAccountId) {
		return nil, false
	}
	return o.SourceAccountId, true
}

// HasSourceAccountId returns a boolean if a field has been set.
func (o *TransactionResponse) HasSourceAccountId() bool {
	if o != nil && !IsNil(o.SourceAccountId) {
		return true
	}

	return false
}

// SetSourceAccountId gets a reference to the given string and assigns it to the SourceAccountId field.
func (o *TransactionResponse) SetSourceAccountId(v string) {
	o.SourceAccountId = &v
}

// GetTransactionReference returns the TransactionReference field value if set, zero value otherwise.
func (o *TransactionResponse) GetTransactionReference() string {
	if o == nil || IsNil(o.TransactionReference) {
		var ret string
		return ret
	}
	return *o.TransactionReference
}

// GetTransactionReferenceOk returns a tuple with the TransactionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetTransactionReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionReference) {
		return nil, false
	}
	return o.TransactionReference, true
}

// HasTransactionReference returns a boolean if a field has been set.
func (o *TransactionResponse) HasTransactionReference() bool {
	if o != nil && !IsNil(o.TransactionReference) {
		return true
	}

	return false
}

// SetTransactionReference gets a reference to the given string and assigns it to the TransactionReference field.
func (o *TransactionResponse) SetTransactionReference(v string) {
	o.TransactionReference = &v
}

// GetTransactionMetadata returns the TransactionMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionResponse) GetTransactionMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.TransactionMetadata
}

// GetTransactionMetadataOk returns a tuple with the TransactionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionResponse) GetTransactionMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.TransactionMetadata) {
		return nil, false
	}
	return &o.TransactionMetadata, true
}

// HasTransactionMetadata returns a boolean if a field has been set.
func (o *TransactionResponse) HasTransactionMetadata() bool {
	if o != nil && IsNil(o.TransactionMetadata) {
		return true
	}

	return false
}

// SetTransactionMetadata gets a reference to the given map[string]string and assigns it to the TransactionMetadata field.
func (o *TransactionResponse) SetTransactionMetadata(v map[string]string) {
	o.TransactionMetadata = v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *TransactionResponse) GetBalance() int64 {
	if o == nil || IsNil(o.Balance) {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetBalanceOk() (*int64, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *TransactionResponse) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *TransactionResponse) SetBalance(v int64) {
	o.Balance = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *TransactionResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *TransactionResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *TransactionResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransactionResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransactionResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TransactionResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o TransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionId"] = o.TransactionId
	toSerialize["transactionShortCode"] = o.TransactionShortCode
	if !IsNil(o.PayorId) {
		toSerialize["payorId"] = o.PayorId
	}
	if !IsNil(o.SourceAccountName) {
		toSerialize["sourceAccountName"] = o.SourceAccountName
	}
	if !IsNil(o.SourceAccountId) {
		toSerialize["sourceAccountId"] = o.SourceAccountId
	}
	if !IsNil(o.TransactionReference) {
		toSerialize["transactionReference"] = o.TransactionReference
	}
	if o.TransactionMetadata != nil {
		toSerialize["transactionMetadata"] = o.TransactionMetadata
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableTransactionResponse struct {
	value *TransactionResponse
	isSet bool
}

func (v NullableTransactionResponse) Get() *TransactionResponse {
	return v.value
}

func (v *NullableTransactionResponse) Set(val *TransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResponse(val *TransactionResponse) *NullableTransactionResponse {
	return &NullableTransactionResponse{value: val, isSet: true}
}

func (v NullableTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


