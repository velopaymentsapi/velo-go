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

// AutoTopUpConfig struct for AutoTopUpConfig
type AutoTopUpConfig struct {
	// Is auto top-up enabled? automatically trigger funding to top-up the source account balance when the balance falls below the configured minimum level.
	Enabled bool `json:"enabled"`
	// When the payor balance falls below this level then auto top-up will be triggered. Note - This is in minor units.
	MinBalance NullableInt64 `json:"minBalance,omitempty"`
	// When the payor balance falls below the min balance then auto top-up will request funds bring the balance to this level. Note - this is in minor units.
	TargetBalance NullableInt64 `json:"targetBalance,omitempty"`
}

// NewAutoTopUpConfig instantiates a new AutoTopUpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTopUpConfig(enabled bool) *AutoTopUpConfig {
	this := AutoTopUpConfig{}
	this.Enabled = enabled
	return &this
}

// NewAutoTopUpConfigWithDefaults instantiates a new AutoTopUpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTopUpConfigWithDefaults() *AutoTopUpConfig {
	this := AutoTopUpConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *AutoTopUpConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AutoTopUpConfig) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AutoTopUpConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMinBalance returns the MinBalance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTopUpConfig) GetMinBalance() int64 {
	if o == nil || o.MinBalance.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MinBalance.Get()
}

// GetMinBalanceOk returns a tuple with the MinBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTopUpConfig) GetMinBalanceOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MinBalance.Get(), o.MinBalance.IsSet()
}

// HasMinBalance returns a boolean if a field has been set.
func (o *AutoTopUpConfig) HasMinBalance() bool {
	if o != nil && o.MinBalance.IsSet() {
		return true
	}

	return false
}

// SetMinBalance gets a reference to the given NullableInt64 and assigns it to the MinBalance field.
func (o *AutoTopUpConfig) SetMinBalance(v int64) {
	o.MinBalance.Set(&v)
}
// SetMinBalanceNil sets the value for MinBalance to be an explicit nil
func (o *AutoTopUpConfig) SetMinBalanceNil() {
	o.MinBalance.Set(nil)
}

// UnsetMinBalance ensures that no value is present for MinBalance, not even an explicit nil
func (o *AutoTopUpConfig) UnsetMinBalance() {
	o.MinBalance.Unset()
}

// GetTargetBalance returns the TargetBalance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTopUpConfig) GetTargetBalance() int64 {
	if o == nil || o.TargetBalance.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TargetBalance.Get()
}

// GetTargetBalanceOk returns a tuple with the TargetBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTopUpConfig) GetTargetBalanceOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetBalance.Get(), o.TargetBalance.IsSet()
}

// HasTargetBalance returns a boolean if a field has been set.
func (o *AutoTopUpConfig) HasTargetBalance() bool {
	if o != nil && o.TargetBalance.IsSet() {
		return true
	}

	return false
}

// SetTargetBalance gets a reference to the given NullableInt64 and assigns it to the TargetBalance field.
func (o *AutoTopUpConfig) SetTargetBalance(v int64) {
	o.TargetBalance.Set(&v)
}
// SetTargetBalanceNil sets the value for TargetBalance to be an explicit nil
func (o *AutoTopUpConfig) SetTargetBalanceNil() {
	o.TargetBalance.Set(nil)
}

// UnsetTargetBalance ensures that no value is present for TargetBalance, not even an explicit nil
func (o *AutoTopUpConfig) UnsetTargetBalance() {
	o.TargetBalance.Unset()
}

func (o AutoTopUpConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MinBalance.IsSet() {
		toSerialize["minBalance"] = o.MinBalance.Get()
	}
	if o.TargetBalance.IsSet() {
		toSerialize["targetBalance"] = o.TargetBalance.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAutoTopUpConfig struct {
	value *AutoTopUpConfig
	isSet bool
}

func (v NullableAutoTopUpConfig) Get() *AutoTopUpConfig {
	return v.value
}

func (v *NullableAutoTopUpConfig) Set(val *AutoTopUpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTopUpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTopUpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTopUpConfig(val *AutoTopUpConfig) *NullableAutoTopUpConfig {
	return &NullableAutoTopUpConfig{value: val, isSet: true}
}

func (v NullableAutoTopUpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTopUpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


