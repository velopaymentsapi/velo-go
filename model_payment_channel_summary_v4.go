/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response.   ## Http Status Codes Following is a list of Http Status codes that could be returned by the platform      | Status Code            | Description                                                                          |     | -----------------------| -------------------------------------------------------------------------------------|     | 200 OK                 | The request was successfully processed and usually returns a json response           |     | 201 Created            | A resource was created and a Location header is returned linking to the new resource |     | 202 Accepted           | The request has been accepted for processing                                         |     | 204 No Content         | The request has been processed and there is no response (usually deletes and updates)|     | 400 Bad Request        | The request is invalid and should be fixed before retrying                           |     | 401 Unauthorized       | Authentication has failed, usually means the token has expired                       |     | 403 Forbidden          | The user does not have permissions for the request                                   |     | 404 Not Found          | The resource was not found                                                           |     | 409 Conflict           | The resource already exists and there is a conflict                                  |     | 429 Too Many Requests  | The user has submitted too many requests in a given amount of time                   |     | 5xx Server Error       | Platform internal error (should rarely happen)                                       | 

API version: 2.37.150
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// checks if the PaymentChannelSummaryV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentChannelSummaryV4{}

// PaymentChannelSummaryV4 struct for PaymentChannelSummaryV4
type PaymentChannelSummaryV4 struct {
	PaymentChannelId string `json:"paymentChannelId"`
	// The payment channel name
	Name string `json:"name"`
	// Valid ISO 3166 2 character country code. See the <a href=\"https://www.iso.org/iso-3166-country-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	CountryCode string `json:"countryCode"`
	// Valid ISO 4217 3 letter currency code. See the <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	Currency string `json:"currency"`
	// The last 4 digits of the account number or IBAN
	Last4Digits *string `json:"last4Digits,omitempty"`
	// Usually true. False if an associated payment is returned
	Enabled bool `json:"enabled"`
	// Indicates why the payment channel was disabled
	DisabledReason *string `json:"disabledReason,omitempty"`
	DisabledReasonCode *string `json:"disabledReasonCode,omitempty"`
}

// NewPaymentChannelSummaryV4 instantiates a new PaymentChannelSummaryV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentChannelSummaryV4(paymentChannelId string, name string, countryCode string, currency string, enabled bool) *PaymentChannelSummaryV4 {
	this := PaymentChannelSummaryV4{}
	this.PaymentChannelId = paymentChannelId
	this.Name = name
	this.CountryCode = countryCode
	this.Currency = currency
	this.Enabled = enabled
	return &this
}

// NewPaymentChannelSummaryV4WithDefaults instantiates a new PaymentChannelSummaryV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentChannelSummaryV4WithDefaults() *PaymentChannelSummaryV4 {
	this := PaymentChannelSummaryV4{}
	return &this
}

// GetPaymentChannelId returns the PaymentChannelId field value
func (o *PaymentChannelSummaryV4) GetPaymentChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentChannelId
}

// GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelSummaryV4) GetPaymentChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentChannelId, true
}

// SetPaymentChannelId sets field value
func (o *PaymentChannelSummaryV4) SetPaymentChannelId(v string) {
	o.PaymentChannelId = v
}

// GetName returns the Name field value
func (o *PaymentChannelSummaryV4) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelSummaryV4) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentChannelSummaryV4) SetName(v string) {
	o.Name = v
}

// GetCountryCode returns the CountryCode field value
func (o *PaymentChannelSummaryV4) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelSummaryV4) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *PaymentChannelSummaryV4) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCurrency returns the Currency field value
func (o *PaymentChannelSummaryV4) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelSummaryV4) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentChannelSummaryV4) SetCurrency(v string) {
	o.Currency = v
}

// GetLast4Digits returns the Last4Digits field value if set, zero value otherwise.
func (o *PaymentChannelSummaryV4) GetLast4Digits() string {
	if o == nil || IsNil(o.Last4Digits) {
		var ret string
		return ret
	}
	return *o.Last4Digits
}

// GetLast4DigitsOk returns a tuple with the Last4Digits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelSummaryV4) GetLast4DigitsOk() (*string, bool) {
	if o == nil || IsNil(o.Last4Digits) {
		return nil, false
	}
	return o.Last4Digits, true
}

// HasLast4Digits returns a boolean if a field has been set.
func (o *PaymentChannelSummaryV4) HasLast4Digits() bool {
	if o != nil && !IsNil(o.Last4Digits) {
		return true
	}

	return false
}

// SetLast4Digits gets a reference to the given string and assigns it to the Last4Digits field.
func (o *PaymentChannelSummaryV4) SetLast4Digits(v string) {
	o.Last4Digits = &v
}

// GetEnabled returns the Enabled field value
func (o *PaymentChannelSummaryV4) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelSummaryV4) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PaymentChannelSummaryV4) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDisabledReason returns the DisabledReason field value if set, zero value otherwise.
func (o *PaymentChannelSummaryV4) GetDisabledReason() string {
	if o == nil || IsNil(o.DisabledReason) {
		var ret string
		return ret
	}
	return *o.DisabledReason
}

// GetDisabledReasonOk returns a tuple with the DisabledReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelSummaryV4) GetDisabledReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DisabledReason) {
		return nil, false
	}
	return o.DisabledReason, true
}

// HasDisabledReason returns a boolean if a field has been set.
func (o *PaymentChannelSummaryV4) HasDisabledReason() bool {
	if o != nil && !IsNil(o.DisabledReason) {
		return true
	}

	return false
}

// SetDisabledReason gets a reference to the given string and assigns it to the DisabledReason field.
func (o *PaymentChannelSummaryV4) SetDisabledReason(v string) {
	o.DisabledReason = &v
}

// GetDisabledReasonCode returns the DisabledReasonCode field value if set, zero value otherwise.
func (o *PaymentChannelSummaryV4) GetDisabledReasonCode() string {
	if o == nil || IsNil(o.DisabledReasonCode) {
		var ret string
		return ret
	}
	return *o.DisabledReasonCode
}

// GetDisabledReasonCodeOk returns a tuple with the DisabledReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelSummaryV4) GetDisabledReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DisabledReasonCode) {
		return nil, false
	}
	return o.DisabledReasonCode, true
}

// HasDisabledReasonCode returns a boolean if a field has been set.
func (o *PaymentChannelSummaryV4) HasDisabledReasonCode() bool {
	if o != nil && !IsNil(o.DisabledReasonCode) {
		return true
	}

	return false
}

// SetDisabledReasonCode gets a reference to the given string and assigns it to the DisabledReasonCode field.
func (o *PaymentChannelSummaryV4) SetDisabledReasonCode(v string) {
	o.DisabledReasonCode = &v
}

func (o PaymentChannelSummaryV4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentChannelSummaryV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentChannelId"] = o.PaymentChannelId
	toSerialize["name"] = o.Name
	toSerialize["countryCode"] = o.CountryCode
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Last4Digits) {
		toSerialize["last4Digits"] = o.Last4Digits
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.DisabledReason) {
		toSerialize["disabledReason"] = o.DisabledReason
	}
	if !IsNil(o.DisabledReasonCode) {
		toSerialize["disabledReasonCode"] = o.DisabledReasonCode
	}
	return toSerialize, nil
}

type NullablePaymentChannelSummaryV4 struct {
	value *PaymentChannelSummaryV4
	isSet bool
}

func (v NullablePaymentChannelSummaryV4) Get() *PaymentChannelSummaryV4 {
	return v.value
}

func (v *NullablePaymentChannelSummaryV4) Set(val *PaymentChannelSummaryV4) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentChannelSummaryV4) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentChannelSummaryV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentChannelSummaryV4(val *PaymentChannelSummaryV4) *NullablePaymentChannelSummaryV4 {
	return &NullablePaymentChannelSummaryV4{value: val, isSet: true}
}

func (v NullablePaymentChannelSummaryV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentChannelSummaryV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


