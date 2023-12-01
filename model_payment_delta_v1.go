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

// checks if the PaymentDeltaV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentDeltaV1{}

// PaymentDeltaV1 struct for PaymentDeltaV1
type PaymentDeltaV1 struct {
	PaymentId string `json:"paymentId"`
	PayoutId string `json:"payoutId"`
	PayorPaymentId NullableString `json:"payorPaymentId,omitempty"`
	PaymentCurrency NullableString `json:"paymentCurrency,omitempty"`
	PaymentAmount NullableInt32 `json:"paymentAmount,omitempty"`
	Status NullableString `json:"status,omitempty"`
	SourceCurrency NullableString `json:"sourceCurrency,omitempty"`
	SourceAmount NullableInt32 `json:"sourceAmount,omitempty"`
}

// NewPaymentDeltaV1 instantiates a new PaymentDeltaV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentDeltaV1(paymentId string, payoutId string) *PaymentDeltaV1 {
	this := PaymentDeltaV1{}
	this.PaymentId = paymentId
	this.PayoutId = payoutId
	return &this
}

// NewPaymentDeltaV1WithDefaults instantiates a new PaymentDeltaV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentDeltaV1WithDefaults() *PaymentDeltaV1 {
	this := PaymentDeltaV1{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentDeltaV1) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentDeltaV1) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentDeltaV1) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetPayoutId returns the PayoutId field value
func (o *PaymentDeltaV1) GetPayoutId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayoutId
}

// GetPayoutIdOk returns a tuple with the PayoutId field value
// and a boolean to check if the value has been set.
func (o *PaymentDeltaV1) GetPayoutIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayoutId, true
}

// SetPayoutId sets field value
func (o *PaymentDeltaV1) SetPayoutId(v string) {
	o.PayoutId = v
}

// GetPayorPaymentId returns the PayorPaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentDeltaV1) GetPayorPaymentId() string {
	if o == nil || IsNil(o.PayorPaymentId.Get()) {
		var ret string
		return ret
	}
	return *o.PayorPaymentId.Get()
}

// GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentDeltaV1) GetPayorPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayorPaymentId.Get(), o.PayorPaymentId.IsSet()
}

// HasPayorPaymentId returns a boolean if a field has been set.
func (o *PaymentDeltaV1) HasPayorPaymentId() bool {
	if o != nil && o.PayorPaymentId.IsSet() {
		return true
	}

	return false
}

// SetPayorPaymentId gets a reference to the given NullableString and assigns it to the PayorPaymentId field.
func (o *PaymentDeltaV1) SetPayorPaymentId(v string) {
	o.PayorPaymentId.Set(&v)
}
// SetPayorPaymentIdNil sets the value for PayorPaymentId to be an explicit nil
func (o *PaymentDeltaV1) SetPayorPaymentIdNil() {
	o.PayorPaymentId.Set(nil)
}

// UnsetPayorPaymentId ensures that no value is present for PayorPaymentId, not even an explicit nil
func (o *PaymentDeltaV1) UnsetPayorPaymentId() {
	o.PayorPaymentId.Unset()
}

// GetPaymentCurrency returns the PaymentCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentDeltaV1) GetPaymentCurrency() string {
	if o == nil || IsNil(o.PaymentCurrency.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentCurrency.Get()
}

// GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentDeltaV1) GetPaymentCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentCurrency.Get(), o.PaymentCurrency.IsSet()
}

// HasPaymentCurrency returns a boolean if a field has been set.
func (o *PaymentDeltaV1) HasPaymentCurrency() bool {
	if o != nil && o.PaymentCurrency.IsSet() {
		return true
	}

	return false
}

// SetPaymentCurrency gets a reference to the given NullableString and assigns it to the PaymentCurrency field.
func (o *PaymentDeltaV1) SetPaymentCurrency(v string) {
	o.PaymentCurrency.Set(&v)
}
// SetPaymentCurrencyNil sets the value for PaymentCurrency to be an explicit nil
func (o *PaymentDeltaV1) SetPaymentCurrencyNil() {
	o.PaymentCurrency.Set(nil)
}

// UnsetPaymentCurrency ensures that no value is present for PaymentCurrency, not even an explicit nil
func (o *PaymentDeltaV1) UnsetPaymentCurrency() {
	o.PaymentCurrency.Unset()
}

// GetPaymentAmount returns the PaymentAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentDeltaV1) GetPaymentAmount() int32 {
	if o == nil || IsNil(o.PaymentAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.PaymentAmount.Get()
}

// GetPaymentAmountOk returns a tuple with the PaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentDeltaV1) GetPaymentAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentAmount.Get(), o.PaymentAmount.IsSet()
}

// HasPaymentAmount returns a boolean if a field has been set.
func (o *PaymentDeltaV1) HasPaymentAmount() bool {
	if o != nil && o.PaymentAmount.IsSet() {
		return true
	}

	return false
}

// SetPaymentAmount gets a reference to the given NullableInt32 and assigns it to the PaymentAmount field.
func (o *PaymentDeltaV1) SetPaymentAmount(v int32) {
	o.PaymentAmount.Set(&v)
}
// SetPaymentAmountNil sets the value for PaymentAmount to be an explicit nil
func (o *PaymentDeltaV1) SetPaymentAmountNil() {
	o.PaymentAmount.Set(nil)
}

// UnsetPaymentAmount ensures that no value is present for PaymentAmount, not even an explicit nil
func (o *PaymentDeltaV1) UnsetPaymentAmount() {
	o.PaymentAmount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentDeltaV1) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentDeltaV1) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentDeltaV1) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *PaymentDeltaV1) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *PaymentDeltaV1) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PaymentDeltaV1) UnsetStatus() {
	o.Status.Unset()
}

// GetSourceCurrency returns the SourceCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentDeltaV1) GetSourceCurrency() string {
	if o == nil || IsNil(o.SourceCurrency.Get()) {
		var ret string
		return ret
	}
	return *o.SourceCurrency.Get()
}

// GetSourceCurrencyOk returns a tuple with the SourceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentDeltaV1) GetSourceCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceCurrency.Get(), o.SourceCurrency.IsSet()
}

// HasSourceCurrency returns a boolean if a field has been set.
func (o *PaymentDeltaV1) HasSourceCurrency() bool {
	if o != nil && o.SourceCurrency.IsSet() {
		return true
	}

	return false
}

// SetSourceCurrency gets a reference to the given NullableString and assigns it to the SourceCurrency field.
func (o *PaymentDeltaV1) SetSourceCurrency(v string) {
	o.SourceCurrency.Set(&v)
}
// SetSourceCurrencyNil sets the value for SourceCurrency to be an explicit nil
func (o *PaymentDeltaV1) SetSourceCurrencyNil() {
	o.SourceCurrency.Set(nil)
}

// UnsetSourceCurrency ensures that no value is present for SourceCurrency, not even an explicit nil
func (o *PaymentDeltaV1) UnsetSourceCurrency() {
	o.SourceCurrency.Unset()
}

// GetSourceAmount returns the SourceAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentDeltaV1) GetSourceAmount() int32 {
	if o == nil || IsNil(o.SourceAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.SourceAmount.Get()
}

// GetSourceAmountOk returns a tuple with the SourceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentDeltaV1) GetSourceAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceAmount.Get(), o.SourceAmount.IsSet()
}

// HasSourceAmount returns a boolean if a field has been set.
func (o *PaymentDeltaV1) HasSourceAmount() bool {
	if o != nil && o.SourceAmount.IsSet() {
		return true
	}

	return false
}

// SetSourceAmount gets a reference to the given NullableInt32 and assigns it to the SourceAmount field.
func (o *PaymentDeltaV1) SetSourceAmount(v int32) {
	o.SourceAmount.Set(&v)
}
// SetSourceAmountNil sets the value for SourceAmount to be an explicit nil
func (o *PaymentDeltaV1) SetSourceAmountNil() {
	o.SourceAmount.Set(nil)
}

// UnsetSourceAmount ensures that no value is present for SourceAmount, not even an explicit nil
func (o *PaymentDeltaV1) UnsetSourceAmount() {
	o.SourceAmount.Unset()
}

func (o PaymentDeltaV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentDeltaV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentId"] = o.PaymentId
	toSerialize["payoutId"] = o.PayoutId
	if o.PayorPaymentId.IsSet() {
		toSerialize["payorPaymentId"] = o.PayorPaymentId.Get()
	}
	if o.PaymentCurrency.IsSet() {
		toSerialize["paymentCurrency"] = o.PaymentCurrency.Get()
	}
	if o.PaymentAmount.IsSet() {
		toSerialize["paymentAmount"] = o.PaymentAmount.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.SourceCurrency.IsSet() {
		toSerialize["sourceCurrency"] = o.SourceCurrency.Get()
	}
	if o.SourceAmount.IsSet() {
		toSerialize["sourceAmount"] = o.SourceAmount.Get()
	}
	return toSerialize, nil
}

type NullablePaymentDeltaV1 struct {
	value *PaymentDeltaV1
	isSet bool
}

func (v NullablePaymentDeltaV1) Get() *PaymentDeltaV1 {
	return v.value
}

func (v *NullablePaymentDeltaV1) Set(val *PaymentDeltaV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentDeltaV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentDeltaV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentDeltaV1(val *PaymentDeltaV1) *NullablePaymentDeltaV1 {
	return &NullablePaymentDeltaV1{value: val, isSet: true}
}

func (v NullablePaymentDeltaV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentDeltaV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


