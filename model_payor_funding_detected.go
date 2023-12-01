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

// checks if the PayorFundingDetected type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayorFundingDetected{}

// PayorFundingDetected Notification of funding received on behalf of payor.
type PayorFundingDetected struct {
	// OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly
	SourceType string `json:"sourceType"`
	// UUID id of the source event in the Velo platform
	EventId string `json:"eventId"`
	// ISO8601 timestamp indicating when the source event was created
	CreatedAt time.Time `json:"createdAt"`
	// the identifier of the payment rail from which funding was received
	RailsId *string `json:"railsId,omitempty"`
	// ID of the payor within the Velo platform
	PayorId string `json:"payorId"`
	// ID of this funding transaction within the Velo platform
	FundingRequestId string `json:"fundingRequestId"`
	// the external identity reference for this funding transaction
	FundingRef *string `json:"fundingRef,omitempty"`
	// the ISO-4217 code for the currency in which the funding was made
	Currency *string `json:"currency,omitempty"`
	// the received funding amount in currency minor units
	Amount *int64 `json:"amount,omitempty"`
	// the name of the account as registered with the payment rail
	PhysicalAccountName *string `json:"physicalAccountName,omitempty"`
	// the name of the account as registered with the Velo platform
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	// the ID of the account as registered with the Velo platform
	SourceAccountId *string `json:"sourceAccountId,omitempty"`
	// any additional information received from the payment rail
	AdditionalInformation *string `json:"additionalInformation,omitempty"`
	// The Id of the related transaction
	TransactionId *string `json:"transactionId,omitempty"`
	// The payors own reference for the related transaction
	TransactionReference *string `json:"transactionReference,omitempty"`
}

// NewPayorFundingDetected instantiates a new PayorFundingDetected object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayorFundingDetected(sourceType string, eventId string, createdAt time.Time, payorId string, fundingRequestId string) *PayorFundingDetected {
	this := PayorFundingDetected{}
	this.SourceType = sourceType
	this.EventId = eventId
	this.CreatedAt = createdAt
	this.PayorId = payorId
	this.FundingRequestId = fundingRequestId
	return &this
}

// NewPayorFundingDetectedWithDefaults instantiates a new PayorFundingDetected object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayorFundingDetectedWithDefaults() *PayorFundingDetected {
	this := PayorFundingDetected{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *PayorFundingDetected) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *PayorFundingDetected) SetSourceType(v string) {
	o.SourceType = v
}

// GetEventId returns the EventId field value
func (o *PayorFundingDetected) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *PayorFundingDetected) SetEventId(v string) {
	o.EventId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PayorFundingDetected) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PayorFundingDetected) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetRailsId returns the RailsId field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetRailsId() string {
	if o == nil || IsNil(o.RailsId) {
		var ret string
		return ret
	}
	return *o.RailsId
}

// GetRailsIdOk returns a tuple with the RailsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetRailsIdOk() (*string, bool) {
	if o == nil || IsNil(o.RailsId) {
		return nil, false
	}
	return o.RailsId, true
}

// HasRailsId returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasRailsId() bool {
	if o != nil && !IsNil(o.RailsId) {
		return true
	}

	return false
}

// SetRailsId gets a reference to the given string and assigns it to the RailsId field.
func (o *PayorFundingDetected) SetRailsId(v string) {
	o.RailsId = &v
}

// GetPayorId returns the PayorId field value
func (o *PayorFundingDetected) GetPayorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetPayorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayorId, true
}

// SetPayorId sets field value
func (o *PayorFundingDetected) SetPayorId(v string) {
	o.PayorId = v
}

// GetFundingRequestId returns the FundingRequestId field value
func (o *PayorFundingDetected) GetFundingRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingRequestId
}

// GetFundingRequestIdOk returns a tuple with the FundingRequestId field value
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetFundingRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingRequestId, true
}

// SetFundingRequestId sets field value
func (o *PayorFundingDetected) SetFundingRequestId(v string) {
	o.FundingRequestId = v
}

// GetFundingRef returns the FundingRef field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetFundingRef() string {
	if o == nil || IsNil(o.FundingRef) {
		var ret string
		return ret
	}
	return *o.FundingRef
}

// GetFundingRefOk returns a tuple with the FundingRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetFundingRefOk() (*string, bool) {
	if o == nil || IsNil(o.FundingRef) {
		return nil, false
	}
	return o.FundingRef, true
}

// HasFundingRef returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasFundingRef() bool {
	if o != nil && !IsNil(o.FundingRef) {
		return true
	}

	return false
}

// SetFundingRef gets a reference to the given string and assigns it to the FundingRef field.
func (o *PayorFundingDetected) SetFundingRef(v string) {
	o.FundingRef = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PayorFundingDetected) SetCurrency(v string) {
	o.Currency = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *PayorFundingDetected) SetAmount(v int64) {
	o.Amount = &v
}

// GetPhysicalAccountName returns the PhysicalAccountName field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetPhysicalAccountName() string {
	if o == nil || IsNil(o.PhysicalAccountName) {
		var ret string
		return ret
	}
	return *o.PhysicalAccountName
}

// GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetPhysicalAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalAccountName) {
		return nil, false
	}
	return o.PhysicalAccountName, true
}

// HasPhysicalAccountName returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasPhysicalAccountName() bool {
	if o != nil && !IsNil(o.PhysicalAccountName) {
		return true
	}

	return false
}

// SetPhysicalAccountName gets a reference to the given string and assigns it to the PhysicalAccountName field.
func (o *PayorFundingDetected) SetPhysicalAccountName(v string) {
	o.PhysicalAccountName = &v
}

// GetSourceAccountName returns the SourceAccountName field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetSourceAccountName() string {
	if o == nil || IsNil(o.SourceAccountName) {
		var ret string
		return ret
	}
	return *o.SourceAccountName
}

// GetSourceAccountNameOk returns a tuple with the SourceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetSourceAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAccountName) {
		return nil, false
	}
	return o.SourceAccountName, true
}

// HasSourceAccountName returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasSourceAccountName() bool {
	if o != nil && !IsNil(o.SourceAccountName) {
		return true
	}

	return false
}

// SetSourceAccountName gets a reference to the given string and assigns it to the SourceAccountName field.
func (o *PayorFundingDetected) SetSourceAccountName(v string) {
	o.SourceAccountName = &v
}

// GetSourceAccountId returns the SourceAccountId field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetSourceAccountId() string {
	if o == nil || IsNil(o.SourceAccountId) {
		var ret string
		return ret
	}
	return *o.SourceAccountId
}

// GetSourceAccountIdOk returns a tuple with the SourceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetSourceAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAccountId) {
		return nil, false
	}
	return o.SourceAccountId, true
}

// HasSourceAccountId returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasSourceAccountId() bool {
	if o != nil && !IsNil(o.SourceAccountId) {
		return true
	}

	return false
}

// SetSourceAccountId gets a reference to the given string and assigns it to the SourceAccountId field.
func (o *PayorFundingDetected) SetSourceAccountId(v string) {
	o.SourceAccountId = &v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetAdditionalInformation() string {
	if o == nil || IsNil(o.AdditionalInformation) {
		var ret string
		return ret
	}
	return *o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetAdditionalInformationOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInformation) {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasAdditionalInformation() bool {
	if o != nil && !IsNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given string and assigns it to the AdditionalInformation field.
func (o *PayorFundingDetected) SetAdditionalInformation(v string) {
	o.AdditionalInformation = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *PayorFundingDetected) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionReference returns the TransactionReference field value if set, zero value otherwise.
func (o *PayorFundingDetected) GetTransactionReference() string {
	if o == nil || IsNil(o.TransactionReference) {
		var ret string
		return ret
	}
	return *o.TransactionReference
}

// GetTransactionReferenceOk returns a tuple with the TransactionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorFundingDetected) GetTransactionReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionReference) {
		return nil, false
	}
	return o.TransactionReference, true
}

// HasTransactionReference returns a boolean if a field has been set.
func (o *PayorFundingDetected) HasTransactionReference() bool {
	if o != nil && !IsNil(o.TransactionReference) {
		return true
	}

	return false
}

// SetTransactionReference gets a reference to the given string and assigns it to the TransactionReference field.
func (o *PayorFundingDetected) SetTransactionReference(v string) {
	o.TransactionReference = &v
}

func (o PayorFundingDetected) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayorFundingDetected) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceType"] = o.SourceType
	toSerialize["eventId"] = o.EventId
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.RailsId) {
		toSerialize["railsId"] = o.RailsId
	}
	toSerialize["payorId"] = o.PayorId
	toSerialize["fundingRequestId"] = o.FundingRequestId
	if !IsNil(o.FundingRef) {
		toSerialize["fundingRef"] = o.FundingRef
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.PhysicalAccountName) {
		toSerialize["physicalAccountName"] = o.PhysicalAccountName
	}
	if !IsNil(o.SourceAccountName) {
		toSerialize["sourceAccountName"] = o.SourceAccountName
	}
	if !IsNil(o.SourceAccountId) {
		toSerialize["sourceAccountId"] = o.SourceAccountId
	}
	if !IsNil(o.AdditionalInformation) {
		toSerialize["additionalInformation"] = o.AdditionalInformation
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !IsNil(o.TransactionReference) {
		toSerialize["transactionReference"] = o.TransactionReference
	}
	return toSerialize, nil
}

type NullablePayorFundingDetected struct {
	value *PayorFundingDetected
	isSet bool
}

func (v NullablePayorFundingDetected) Get() *PayorFundingDetected {
	return v.value
}

func (v *NullablePayorFundingDetected) Set(val *PayorFundingDetected) {
	v.value = val
	v.isSet = true
}

func (v NullablePayorFundingDetected) IsSet() bool {
	return v.isSet
}

func (v *NullablePayorFundingDetected) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayorFundingDetected(val *PayorFundingDetected) *NullablePayorFundingDetected {
	return &NullablePayorFundingDetected{value: val, isSet: true}
}

func (v NullablePayorFundingDetected) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayorFundingDetected) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

