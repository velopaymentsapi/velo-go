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

// FundingAccountResponse2 struct for FundingAccountResponse2
type FundingAccountResponse2 struct {
	// Funding Account Id
	Id *string `json:"id,omitempty"`
	PayorId *string `json:"payorId,omitempty"`
	// name on the bank account
	AccountName *string `json:"accountName,omitempty"`
	// bank account number
	AccountNumber *string `json:"accountNumber,omitempty"`
	// bank account routing number
	RoutingNumber *string `json:"routingNumber,omitempty"`
	// name of funding account
	Name *string `json:"name,omitempty"`
	// ISO 4217 currency code
	Currency *string `json:"currency,omitempty"`
	// ISO 3166-1 2 letter country code (upper case)
	Country *string `json:"country,omitempty"`
	// Funding account type
	Type *string `json:"type,omitempty"`
	// A flag for whether the funding account has been archived.  Only present in the response if true.
	Archived *bool `json:"archived,omitempty"`
}

// NewFundingAccountResponse2 instantiates a new FundingAccountResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundingAccountResponse2() *FundingAccountResponse2 {
	this := FundingAccountResponse2{}
	return &this
}

// NewFundingAccountResponse2WithDefaults instantiates a new FundingAccountResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingAccountResponse2WithDefaults() *FundingAccountResponse2 {
	this := FundingAccountResponse2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FundingAccountResponse2) SetId(v string) {
	o.Id = &v
}

// GetPayorId returns the PayorId field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetPayorId() string {
	if o == nil || o.PayorId == nil {
		var ret string
		return ret
	}
	return *o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetPayorIdOk() (*string, bool) {
	if o == nil || o.PayorId == nil {
		return nil, false
	}
	return o.PayorId, true
}

// HasPayorId returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasPayorId() bool {
	if o != nil && o.PayorId != nil {
		return true
	}

	return false
}

// SetPayorId gets a reference to the given string and assigns it to the PayorId field.
func (o *FundingAccountResponse2) SetPayorId(v string) {
	o.PayorId = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *FundingAccountResponse2) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *FundingAccountResponse2) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetRoutingNumber() string {
	if o == nil || o.RoutingNumber == nil {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetRoutingNumberOk() (*string, bool) {
	if o == nil || o.RoutingNumber == nil {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasRoutingNumber() bool {
	if o != nil && o.RoutingNumber != nil {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *FundingAccountResponse2) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FundingAccountResponse2) SetName(v string) {
	o.Name = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *FundingAccountResponse2) SetCurrency(v string) {
	o.Currency = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *FundingAccountResponse2) SetCountry(v string) {
	o.Country = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FundingAccountResponse2) SetType(v string) {
	o.Type = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *FundingAccountResponse2) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponse2) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *FundingAccountResponse2) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *FundingAccountResponse2) SetArchived(v bool) {
	o.Archived = &v
}

func (o FundingAccountResponse2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PayorId != nil {
		toSerialize["payorId"] = o.PayorId
	}
	if o.AccountName != nil {
		toSerialize["accountName"] = o.AccountName
	}
	if o.AccountNumber != nil {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if o.RoutingNumber != nil {
		toSerialize["routingNumber"] = o.RoutingNumber
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	return json.Marshal(toSerialize)
}

type NullableFundingAccountResponse2 struct {
	value *FundingAccountResponse2
	isSet bool
}

func (v NullableFundingAccountResponse2) Get() *FundingAccountResponse2 {
	return v.value
}

func (v *NullableFundingAccountResponse2) Set(val *FundingAccountResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableFundingAccountResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableFundingAccountResponse2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundingAccountResponse2(val *FundingAccountResponse2) *NullableFundingAccountResponse2 {
	return &NullableFundingAccountResponse2{value: val, isSet: true}
}

func (v NullableFundingAccountResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundingAccountResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


