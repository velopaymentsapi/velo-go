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

// SourceAccountResponse struct for SourceAccountResponse
type SourceAccountResponse struct {
	// Source Account Id
	Id *string `json:"id,omitempty"`
	// Decimal implied
	Balance *int64 `json:"balance,omitempty"`
	Currency *string `json:"currency,omitempty"`
	FundingRef *string `json:"fundingRef,omitempty"`
	PhysicalAccountName *string `json:"physicalAccountName,omitempty"`
	RailsId *string `json:"railsId,omitempty"`
	PayorId *string `json:"payorId,omitempty"`
	Name *string `json:"name,omitempty"`
	Pooled *bool `json:"pooled,omitempty"`
	BalanceVisible *bool `json:"balanceVisible,omitempty"`
	CustomerId NullableString `json:"customerId,omitempty"`
	PhysicalAccountId *string `json:"physicalAccountId,omitempty"`
	FundingAccountId NullableString `json:"fundingAccountId,omitempty"`
	AccountType *string `json:"accountType,omitempty"`
}

// NewSourceAccountResponse instantiates a new SourceAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAccountResponse() *SourceAccountResponse {
	this := SourceAccountResponse{}
	return &this
}

// NewSourceAccountResponseWithDefaults instantiates a new SourceAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAccountResponseWithDefaults() *SourceAccountResponse {
	this := SourceAccountResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceAccountResponse) SetId(v string) {
	o.Id = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetBalance() int64 {
	if o == nil || o.Balance == nil {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetBalanceOk() (*int64, bool) {
	if o == nil || o.Balance == nil {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasBalance() bool {
	if o != nil && o.Balance != nil {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *SourceAccountResponse) SetBalance(v int64) {
	o.Balance = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SourceAccountResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetFundingRef returns the FundingRef field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetFundingRef() string {
	if o == nil || o.FundingRef == nil {
		var ret string
		return ret
	}
	return *o.FundingRef
}

// GetFundingRefOk returns a tuple with the FundingRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetFundingRefOk() (*string, bool) {
	if o == nil || o.FundingRef == nil {
		return nil, false
	}
	return o.FundingRef, true
}

// HasFundingRef returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasFundingRef() bool {
	if o != nil && o.FundingRef != nil {
		return true
	}

	return false
}

// SetFundingRef gets a reference to the given string and assigns it to the FundingRef field.
func (o *SourceAccountResponse) SetFundingRef(v string) {
	o.FundingRef = &v
}

// GetPhysicalAccountName returns the PhysicalAccountName field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetPhysicalAccountName() string {
	if o == nil || o.PhysicalAccountName == nil {
		var ret string
		return ret
	}
	return *o.PhysicalAccountName
}

// GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetPhysicalAccountNameOk() (*string, bool) {
	if o == nil || o.PhysicalAccountName == nil {
		return nil, false
	}
	return o.PhysicalAccountName, true
}

// HasPhysicalAccountName returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasPhysicalAccountName() bool {
	if o != nil && o.PhysicalAccountName != nil {
		return true
	}

	return false
}

// SetPhysicalAccountName gets a reference to the given string and assigns it to the PhysicalAccountName field.
func (o *SourceAccountResponse) SetPhysicalAccountName(v string) {
	o.PhysicalAccountName = &v
}

// GetRailsId returns the RailsId field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetRailsId() string {
	if o == nil || o.RailsId == nil {
		var ret string
		return ret
	}
	return *o.RailsId
}

// GetRailsIdOk returns a tuple with the RailsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetRailsIdOk() (*string, bool) {
	if o == nil || o.RailsId == nil {
		return nil, false
	}
	return o.RailsId, true
}

// HasRailsId returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasRailsId() bool {
	if o != nil && o.RailsId != nil {
		return true
	}

	return false
}

// SetRailsId gets a reference to the given string and assigns it to the RailsId field.
func (o *SourceAccountResponse) SetRailsId(v string) {
	o.RailsId = &v
}

// GetPayorId returns the PayorId field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetPayorId() string {
	if o == nil || o.PayorId == nil {
		var ret string
		return ret
	}
	return *o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetPayorIdOk() (*string, bool) {
	if o == nil || o.PayorId == nil {
		return nil, false
	}
	return o.PayorId, true
}

// HasPayorId returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasPayorId() bool {
	if o != nil && o.PayorId != nil {
		return true
	}

	return false
}

// SetPayorId gets a reference to the given string and assigns it to the PayorId field.
func (o *SourceAccountResponse) SetPayorId(v string) {
	o.PayorId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceAccountResponse) SetName(v string) {
	o.Name = &v
}

// GetPooled returns the Pooled field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetPooled() bool {
	if o == nil || o.Pooled == nil {
		var ret bool
		return ret
	}
	return *o.Pooled
}

// GetPooledOk returns a tuple with the Pooled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetPooledOk() (*bool, bool) {
	if o == nil || o.Pooled == nil {
		return nil, false
	}
	return o.Pooled, true
}

// HasPooled returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasPooled() bool {
	if o != nil && o.Pooled != nil {
		return true
	}

	return false
}

// SetPooled gets a reference to the given bool and assigns it to the Pooled field.
func (o *SourceAccountResponse) SetPooled(v bool) {
	o.Pooled = &v
}

// GetBalanceVisible returns the BalanceVisible field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetBalanceVisible() bool {
	if o == nil || o.BalanceVisible == nil {
		var ret bool
		return ret
	}
	return *o.BalanceVisible
}

// GetBalanceVisibleOk returns a tuple with the BalanceVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetBalanceVisibleOk() (*bool, bool) {
	if o == nil || o.BalanceVisible == nil {
		return nil, false
	}
	return o.BalanceVisible, true
}

// HasBalanceVisible returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasBalanceVisible() bool {
	if o != nil && o.BalanceVisible != nil {
		return true
	}

	return false
}

// SetBalanceVisible gets a reference to the given bool and assigns it to the BalanceVisible field.
func (o *SourceAccountResponse) SetBalanceVisible(v bool) {
	o.BalanceVisible = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceAccountResponse) GetCustomerId() string {
	if o == nil || o.CustomerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceAccountResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *SourceAccountResponse) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *SourceAccountResponse) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *SourceAccountResponse) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetPhysicalAccountId returns the PhysicalAccountId field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetPhysicalAccountId() string {
	if o == nil || o.PhysicalAccountId == nil {
		var ret string
		return ret
	}
	return *o.PhysicalAccountId
}

// GetPhysicalAccountIdOk returns a tuple with the PhysicalAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetPhysicalAccountIdOk() (*string, bool) {
	if o == nil || o.PhysicalAccountId == nil {
		return nil, false
	}
	return o.PhysicalAccountId, true
}

// HasPhysicalAccountId returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasPhysicalAccountId() bool {
	if o != nil && o.PhysicalAccountId != nil {
		return true
	}

	return false
}

// SetPhysicalAccountId gets a reference to the given string and assigns it to the PhysicalAccountId field.
func (o *SourceAccountResponse) SetPhysicalAccountId(v string) {
	o.PhysicalAccountId = &v
}

// GetFundingAccountId returns the FundingAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceAccountResponse) GetFundingAccountId() string {
	if o == nil || o.FundingAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.FundingAccountId.Get()
}

// GetFundingAccountIdOk returns a tuple with the FundingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceAccountResponse) GetFundingAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FundingAccountId.Get(), o.FundingAccountId.IsSet()
}

// HasFundingAccountId returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasFundingAccountId() bool {
	if o != nil && o.FundingAccountId.IsSet() {
		return true
	}

	return false
}

// SetFundingAccountId gets a reference to the given NullableString and assigns it to the FundingAccountId field.
func (o *SourceAccountResponse) SetFundingAccountId(v string) {
	o.FundingAccountId.Set(&v)
}
// SetFundingAccountIdNil sets the value for FundingAccountId to be an explicit nil
func (o *SourceAccountResponse) SetFundingAccountIdNil() {
	o.FundingAccountId.Set(nil)
}

// UnsetFundingAccountId ensures that no value is present for FundingAccountId, not even an explicit nil
func (o *SourceAccountResponse) UnsetFundingAccountId() {
	o.FundingAccountId.Unset()
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *SourceAccountResponse) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponse) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *SourceAccountResponse) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *SourceAccountResponse) SetAccountType(v string) {
	o.AccountType = &v
}

func (o SourceAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Balance != nil {
		toSerialize["balance"] = o.Balance
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.FundingRef != nil {
		toSerialize["fundingRef"] = o.FundingRef
	}
	if o.PhysicalAccountName != nil {
		toSerialize["physicalAccountName"] = o.PhysicalAccountName
	}
	if o.RailsId != nil {
		toSerialize["railsId"] = o.RailsId
	}
	if o.PayorId != nil {
		toSerialize["payorId"] = o.PayorId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Pooled != nil {
		toSerialize["pooled"] = o.Pooled
	}
	if o.BalanceVisible != nil {
		toSerialize["balanceVisible"] = o.BalanceVisible
	}
	if o.CustomerId.IsSet() {
		toSerialize["customerId"] = o.CustomerId.Get()
	}
	if o.PhysicalAccountId != nil {
		toSerialize["physicalAccountId"] = o.PhysicalAccountId
	}
	if o.FundingAccountId.IsSet() {
		toSerialize["fundingAccountId"] = o.FundingAccountId.Get()
	}
	if o.AccountType != nil {
		toSerialize["accountType"] = o.AccountType
	}
	return json.Marshal(toSerialize)
}

type NullableSourceAccountResponse struct {
	value *SourceAccountResponse
	isSet bool
}

func (v NullableSourceAccountResponse) Get() *SourceAccountResponse {
	return v.value
}

func (v *NullableSourceAccountResponse) Set(val *SourceAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAccountResponse(val *SourceAccountResponse) *NullableSourceAccountResponse {
	return &NullableSourceAccountResponse{value: val, isSet: true}
}

func (v NullableSourceAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


