/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 

API version: 2.29.128
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// CreatePaymentChannel struct for CreatePaymentChannel
type CreatePaymentChannel struct {
	PaymentChannelName *string `json:"paymentChannelName,omitempty"`
	// Must match the regular expression ```^[A-Za-z0-9]+$```. Either routing number and account number or only iban must be set
	Iban *string `json:"iban,omitempty"`
	// Either routing number and account number or only iban must be set
	AccountNumber *string `json:"accountNumber,omitempty"`
	// Either routing number and account number or only iban must be set
	RoutingNumber *string `json:"routingNumber,omitempty"`
	// Two character country code
	CountryCode string `json:"countryCode"`
	Currency string `json:"currency"`
	AccountName string `json:"accountName"`
}

// NewCreatePaymentChannel instantiates a new CreatePaymentChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePaymentChannel(countryCode string, currency string, accountName string) *CreatePaymentChannel {
	this := CreatePaymentChannel{}
	this.CountryCode = countryCode
	this.Currency = currency
	this.AccountName = accountName
	return &this
}

// NewCreatePaymentChannelWithDefaults instantiates a new CreatePaymentChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePaymentChannelWithDefaults() *CreatePaymentChannel {
	this := CreatePaymentChannel{}
	return &this
}

// GetPaymentChannelName returns the PaymentChannelName field value if set, zero value otherwise.
func (o *CreatePaymentChannel) GetPaymentChannelName() string {
	if o == nil || o.PaymentChannelName == nil {
		var ret string
		return ret
	}
	return *o.PaymentChannelName
}

// GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentChannel) GetPaymentChannelNameOk() (*string, bool) {
	if o == nil || o.PaymentChannelName == nil {
		return nil, false
	}
	return o.PaymentChannelName, true
}

// HasPaymentChannelName returns a boolean if a field has been set.
func (o *CreatePaymentChannel) HasPaymentChannelName() bool {
	if o != nil && o.PaymentChannelName != nil {
		return true
	}

	return false
}

// SetPaymentChannelName gets a reference to the given string and assigns it to the PaymentChannelName field.
func (o *CreatePaymentChannel) SetPaymentChannelName(v string) {
	o.PaymentChannelName = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *CreatePaymentChannel) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentChannel) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *CreatePaymentChannel) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *CreatePaymentChannel) SetIban(v string) {
	o.Iban = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *CreatePaymentChannel) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentChannel) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *CreatePaymentChannel) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *CreatePaymentChannel) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *CreatePaymentChannel) GetRoutingNumber() string {
	if o == nil || o.RoutingNumber == nil {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentChannel) GetRoutingNumberOk() (*string, bool) {
	if o == nil || o.RoutingNumber == nil {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *CreatePaymentChannel) HasRoutingNumber() bool {
	if o != nil && o.RoutingNumber != nil {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *CreatePaymentChannel) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetCountryCode returns the CountryCode field value
func (o *CreatePaymentChannel) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *CreatePaymentChannel) GetCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *CreatePaymentChannel) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCurrency returns the Currency field value
func (o *CreatePaymentChannel) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreatePaymentChannel) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreatePaymentChannel) SetCurrency(v string) {
	o.Currency = v
}

// GetAccountName returns the AccountName field value
func (o *CreatePaymentChannel) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *CreatePaymentChannel) GetAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *CreatePaymentChannel) SetAccountName(v string) {
	o.AccountName = v
}

func (o CreatePaymentChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentChannelName != nil {
		toSerialize["paymentChannelName"] = o.PaymentChannelName
	}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	if o.AccountNumber != nil {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if o.RoutingNumber != nil {
		toSerialize["routingNumber"] = o.RoutingNumber
	}
	if true {
		toSerialize["countryCode"] = o.CountryCode
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["accountName"] = o.AccountName
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePaymentChannel struct {
	value *CreatePaymentChannel
	isSet bool
}

func (v NullableCreatePaymentChannel) Get() *CreatePaymentChannel {
	return v.value
}

func (v *NullableCreatePaymentChannel) Set(val *CreatePaymentChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePaymentChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePaymentChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePaymentChannel(val *CreatePaymentChannel) *NullableCreatePaymentChannel {
	return &NullableCreatePaymentChannel{value: val, isSet: true}
}

func (v NullableCreatePaymentChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePaymentChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


