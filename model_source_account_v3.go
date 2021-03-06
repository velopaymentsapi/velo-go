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

// SourceAccountV3 struct for SourceAccountV3
type SourceAccountV3 struct {
	SourceAccountName string `json:"sourceAccountName"`
	SourceAccountId string `json:"sourceAccountId"`
	// Valid ISO 4217 3 letter currency code. See the <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	Currency string `json:"currency"`
	TotalPayoutCost int32 `json:"totalPayoutCost"`
}

// NewSourceAccountV3 instantiates a new SourceAccountV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAccountV3(sourceAccountName string, sourceAccountId string, currency string, totalPayoutCost int32) *SourceAccountV3 {
	this := SourceAccountV3{}
	this.SourceAccountName = sourceAccountName
	this.SourceAccountId = sourceAccountId
	this.Currency = currency
	this.TotalPayoutCost = totalPayoutCost
	return &this
}

// NewSourceAccountV3WithDefaults instantiates a new SourceAccountV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAccountV3WithDefaults() *SourceAccountV3 {
	this := SourceAccountV3{}
	return &this
}

// GetSourceAccountName returns the SourceAccountName field value
func (o *SourceAccountV3) GetSourceAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAccountName
}

// GetSourceAccountNameOk returns a tuple with the SourceAccountName field value
// and a boolean to check if the value has been set.
func (o *SourceAccountV3) GetSourceAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceAccountName, true
}

// SetSourceAccountName sets field value
func (o *SourceAccountV3) SetSourceAccountName(v string) {
	o.SourceAccountName = v
}

// GetSourceAccountId returns the SourceAccountId field value
func (o *SourceAccountV3) GetSourceAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAccountId
}

// GetSourceAccountIdOk returns a tuple with the SourceAccountId field value
// and a boolean to check if the value has been set.
func (o *SourceAccountV3) GetSourceAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceAccountId, true
}

// SetSourceAccountId sets field value
func (o *SourceAccountV3) SetSourceAccountId(v string) {
	o.SourceAccountId = v
}

// GetCurrency returns the Currency field value
func (o *SourceAccountV3) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SourceAccountV3) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SourceAccountV3) SetCurrency(v string) {
	o.Currency = v
}

// GetTotalPayoutCost returns the TotalPayoutCost field value
func (o *SourceAccountV3) GetTotalPayoutCost() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPayoutCost
}

// GetTotalPayoutCostOk returns a tuple with the TotalPayoutCost field value
// and a boolean to check if the value has been set.
func (o *SourceAccountV3) GetTotalPayoutCostOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalPayoutCost, true
}

// SetTotalPayoutCost sets field value
func (o *SourceAccountV3) SetTotalPayoutCost(v int32) {
	o.TotalPayoutCost = v
}

func (o SourceAccountV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceAccountName"] = o.SourceAccountName
	}
	if true {
		toSerialize["sourceAccountId"] = o.SourceAccountId
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["totalPayoutCost"] = o.TotalPayoutCost
	}
	return json.Marshal(toSerialize)
}

type NullableSourceAccountV3 struct {
	value *SourceAccountV3
	isSet bool
}

func (v NullableSourceAccountV3) Get() *SourceAccountV3 {
	return v.value
}

func (v *NullableSourceAccountV3) Set(val *SourceAccountV3) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAccountV3) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAccountV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAccountV3(val *SourceAccountV3) *NullableSourceAccountV3 {
	return &NullableSourceAccountV3{value: val, isSet: true}
}

func (v NullableSourceAccountV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAccountV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


