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

// checks if the FundingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FundingResponse{}

// FundingResponse struct for FundingResponse
type FundingResponse struct {
	FundingId string `json:"fundingId"`
	PayorId string `json:"payorId"`
	// The date and time the funding was created
	CreatedAt time.Time `json:"createdAt"`
	DetectedFundingRef *string `json:"detectedFundingRef,omitempty"`
	Amount int64 `json:"amount"`
	// Valid ISO 4217 3 letter currency code. See the <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	Currency string `json:"currency"`
	Text *string `json:"text,omitempty"`
	PhysicalAccountName *string `json:"physicalAccountName,omitempty"`
	SourceAccountId *string `json:"sourceAccountId,omitempty"`
	// Funding Allocation Type. One of the following values: AUTOMATIC, MANUAL
	AllocationType *string `json:"allocationType,omitempty"`
	// Populated only if the funding has been allocated. The date and time the funding was allocated.
	AllocatedAt *time.Time `json:"allocatedAt,omitempty"`
	// Populated with allocatedAt if allocated otherwise createdAt. Deprecated in v2.36 - will be removed in the future.
	// Deprecated
	AllocationDate *time.Time `json:"allocationDate,omitempty"`
	Reason *string `json:"reason,omitempty"`
	HiddenDate *time.Time `json:"hiddenDate,omitempty"`
	// Funding Account Type. One of the following values: FBO, PRIVATE
	FundingAccountType string `json:"fundingAccountType"`
	// Current status of the funding. One of the follwing values: PENDING, UNALLOCATED, ALLOCATED, HIDDEN, RETURNED, RETURNING, BULK_RETURN, OTHER
	Status string `json:"status"`
}

// NewFundingResponse instantiates a new FundingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundingResponse(fundingId string, payorId string, createdAt time.Time, amount int64, currency string, fundingAccountType string, status string) *FundingResponse {
	this := FundingResponse{}
	this.FundingId = fundingId
	this.PayorId = payorId
	this.CreatedAt = createdAt
	this.Amount = amount
	this.Currency = currency
	this.FundingAccountType = fundingAccountType
	this.Status = status
	return &this
}

// NewFundingResponseWithDefaults instantiates a new FundingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingResponseWithDefaults() *FundingResponse {
	this := FundingResponse{}
	return &this
}

// GetFundingId returns the FundingId field value
func (o *FundingResponse) GetFundingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingId
}

// GetFundingIdOk returns a tuple with the FundingId field value
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetFundingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingId, true
}

// SetFundingId sets field value
func (o *FundingResponse) SetFundingId(v string) {
	o.FundingId = v
}

// GetPayorId returns the PayorId field value
func (o *FundingResponse) GetPayorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetPayorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayorId, true
}

// SetPayorId sets field value
func (o *FundingResponse) SetPayorId(v string) {
	o.PayorId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FundingResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FundingResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDetectedFundingRef returns the DetectedFundingRef field value if set, zero value otherwise.
func (o *FundingResponse) GetDetectedFundingRef() string {
	if o == nil || IsNil(o.DetectedFundingRef) {
		var ret string
		return ret
	}
	return *o.DetectedFundingRef
}

// GetDetectedFundingRefOk returns a tuple with the DetectedFundingRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetDetectedFundingRefOk() (*string, bool) {
	if o == nil || IsNil(o.DetectedFundingRef) {
		return nil, false
	}
	return o.DetectedFundingRef, true
}

// HasDetectedFundingRef returns a boolean if a field has been set.
func (o *FundingResponse) HasDetectedFundingRef() bool {
	if o != nil && !IsNil(o.DetectedFundingRef) {
		return true
	}

	return false
}

// SetDetectedFundingRef gets a reference to the given string and assigns it to the DetectedFundingRef field.
func (o *FundingResponse) SetDetectedFundingRef(v string) {
	o.DetectedFundingRef = &v
}

// GetAmount returns the Amount field value
func (o *FundingResponse) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *FundingResponse) SetAmount(v int64) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *FundingResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *FundingResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *FundingResponse) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *FundingResponse) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *FundingResponse) SetText(v string) {
	o.Text = &v
}

// GetPhysicalAccountName returns the PhysicalAccountName field value if set, zero value otherwise.
func (o *FundingResponse) GetPhysicalAccountName() string {
	if o == nil || IsNil(o.PhysicalAccountName) {
		var ret string
		return ret
	}
	return *o.PhysicalAccountName
}

// GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetPhysicalAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalAccountName) {
		return nil, false
	}
	return o.PhysicalAccountName, true
}

// HasPhysicalAccountName returns a boolean if a field has been set.
func (o *FundingResponse) HasPhysicalAccountName() bool {
	if o != nil && !IsNil(o.PhysicalAccountName) {
		return true
	}

	return false
}

// SetPhysicalAccountName gets a reference to the given string and assigns it to the PhysicalAccountName field.
func (o *FundingResponse) SetPhysicalAccountName(v string) {
	o.PhysicalAccountName = &v
}

// GetSourceAccountId returns the SourceAccountId field value if set, zero value otherwise.
func (o *FundingResponse) GetSourceAccountId() string {
	if o == nil || IsNil(o.SourceAccountId) {
		var ret string
		return ret
	}
	return *o.SourceAccountId
}

// GetSourceAccountIdOk returns a tuple with the SourceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetSourceAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAccountId) {
		return nil, false
	}
	return o.SourceAccountId, true
}

// HasSourceAccountId returns a boolean if a field has been set.
func (o *FundingResponse) HasSourceAccountId() bool {
	if o != nil && !IsNil(o.SourceAccountId) {
		return true
	}

	return false
}

// SetSourceAccountId gets a reference to the given string and assigns it to the SourceAccountId field.
func (o *FundingResponse) SetSourceAccountId(v string) {
	o.SourceAccountId = &v
}

// GetAllocationType returns the AllocationType field value if set, zero value otherwise.
func (o *FundingResponse) GetAllocationType() string {
	if o == nil || IsNil(o.AllocationType) {
		var ret string
		return ret
	}
	return *o.AllocationType
}

// GetAllocationTypeOk returns a tuple with the AllocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetAllocationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AllocationType) {
		return nil, false
	}
	return o.AllocationType, true
}

// HasAllocationType returns a boolean if a field has been set.
func (o *FundingResponse) HasAllocationType() bool {
	if o != nil && !IsNil(o.AllocationType) {
		return true
	}

	return false
}

// SetAllocationType gets a reference to the given string and assigns it to the AllocationType field.
func (o *FundingResponse) SetAllocationType(v string) {
	o.AllocationType = &v
}

// GetAllocatedAt returns the AllocatedAt field value if set, zero value otherwise.
func (o *FundingResponse) GetAllocatedAt() time.Time {
	if o == nil || IsNil(o.AllocatedAt) {
		var ret time.Time
		return ret
	}
	return *o.AllocatedAt
}

// GetAllocatedAtOk returns a tuple with the AllocatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetAllocatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AllocatedAt) {
		return nil, false
	}
	return o.AllocatedAt, true
}

// HasAllocatedAt returns a boolean if a field has been set.
func (o *FundingResponse) HasAllocatedAt() bool {
	if o != nil && !IsNil(o.AllocatedAt) {
		return true
	}

	return false
}

// SetAllocatedAt gets a reference to the given time.Time and assigns it to the AllocatedAt field.
func (o *FundingResponse) SetAllocatedAt(v time.Time) {
	o.AllocatedAt = &v
}

// GetAllocationDate returns the AllocationDate field value if set, zero value otherwise.
// Deprecated
func (o *FundingResponse) GetAllocationDate() time.Time {
	if o == nil || IsNil(o.AllocationDate) {
		var ret time.Time
		return ret
	}
	return *o.AllocationDate
}

// GetAllocationDateOk returns a tuple with the AllocationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FundingResponse) GetAllocationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AllocationDate) {
		return nil, false
	}
	return o.AllocationDate, true
}

// HasAllocationDate returns a boolean if a field has been set.
func (o *FundingResponse) HasAllocationDate() bool {
	if o != nil && !IsNil(o.AllocationDate) {
		return true
	}

	return false
}

// SetAllocationDate gets a reference to the given time.Time and assigns it to the AllocationDate field.
// Deprecated
func (o *FundingResponse) SetAllocationDate(v time.Time) {
	o.AllocationDate = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *FundingResponse) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *FundingResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *FundingResponse) SetReason(v string) {
	o.Reason = &v
}

// GetHiddenDate returns the HiddenDate field value if set, zero value otherwise.
func (o *FundingResponse) GetHiddenDate() time.Time {
	if o == nil || IsNil(o.HiddenDate) {
		var ret time.Time
		return ret
	}
	return *o.HiddenDate
}

// GetHiddenDateOk returns a tuple with the HiddenDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetHiddenDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.HiddenDate) {
		return nil, false
	}
	return o.HiddenDate, true
}

// HasHiddenDate returns a boolean if a field has been set.
func (o *FundingResponse) HasHiddenDate() bool {
	if o != nil && !IsNil(o.HiddenDate) {
		return true
	}

	return false
}

// SetHiddenDate gets a reference to the given time.Time and assigns it to the HiddenDate field.
func (o *FundingResponse) SetHiddenDate(v time.Time) {
	o.HiddenDate = &v
}

// GetFundingAccountType returns the FundingAccountType field value
func (o *FundingResponse) GetFundingAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingAccountType
}

// GetFundingAccountTypeOk returns a tuple with the FundingAccountType field value
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetFundingAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingAccountType, true
}

// SetFundingAccountType sets field value
func (o *FundingResponse) SetFundingAccountType(v string) {
	o.FundingAccountType = v
}

// GetStatus returns the Status field value
func (o *FundingResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FundingResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FundingResponse) SetStatus(v string) {
	o.Status = v
}

func (o FundingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fundingId"] = o.FundingId
	toSerialize["payorId"] = o.PayorId
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.DetectedFundingRef) {
		toSerialize["detectedFundingRef"] = o.DetectedFundingRef
	}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.PhysicalAccountName) {
		toSerialize["physicalAccountName"] = o.PhysicalAccountName
	}
	if !IsNil(o.SourceAccountId) {
		toSerialize["sourceAccountId"] = o.SourceAccountId
	}
	if !IsNil(o.AllocationType) {
		toSerialize["allocationType"] = o.AllocationType
	}
	if !IsNil(o.AllocatedAt) {
		toSerialize["allocatedAt"] = o.AllocatedAt
	}
	if !IsNil(o.AllocationDate) {
		toSerialize["allocationDate"] = o.AllocationDate
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.HiddenDate) {
		toSerialize["hiddenDate"] = o.HiddenDate
	}
	toSerialize["fundingAccountType"] = o.FundingAccountType
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableFundingResponse struct {
	value *FundingResponse
	isSet bool
}

func (v NullableFundingResponse) Get() *FundingResponse {
	return v.value
}

func (v *NullableFundingResponse) Set(val *FundingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFundingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFundingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundingResponse(val *FundingResponse) *NullableFundingResponse {
	return &NullableFundingResponse{value: val, isSet: true}
}

func (v NullableFundingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


