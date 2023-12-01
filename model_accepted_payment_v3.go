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

// checks if the AcceptedPaymentV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcceptedPaymentV3{}

// AcceptedPaymentV3 struct for AcceptedPaymentV3
type AcceptedPaymentV3 struct {
	// Your identifier for the payee
	RemoteId string `json:"remoteId"`
	// Valid ISO 4217 3 letter currency code. See the <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	CurrencyType string `json:"currencyType"`
	// The amount of the payment in minor units
	Amount int32 `json:"amount"`
	// The identifier of the source account to debit the payment from
	SourceAccountName string `json:"sourceAccountName"`
	// A reference identifier for the payor for the given payee payment
	PayorPaymentId string `json:"payorPaymentId"`
	// <p>Any value here will override the memo value in the parent payout</p> <p>This should be the reference field on the statement seen by the payee (but not via ACH)</p> 
	PaymentMemo *string `json:"paymentMemo,omitempty"`
	// <p>The identifier for the remote payments system if not Velo</p> 
	RemoteSystemId *string `json:"remoteSystemId,omitempty"`
	// <p>Metadata about the payment that may be relevant to the specific rails or remote system making the payout</p> <p>The structure of the data will be dictated by the requirements of the payment rails</p> 
	PaymentMetadata *string `json:"paymentMetadata,omitempty"`
	// Indicates the 3rd party system involved in making this payment
	RailsId string `json:"railsId"`
}

// NewAcceptedPaymentV3 instantiates a new AcceptedPaymentV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptedPaymentV3(remoteId string, currencyType string, amount int32, sourceAccountName string, payorPaymentId string, railsId string) *AcceptedPaymentV3 {
	this := AcceptedPaymentV3{}
	this.RemoteId = remoteId
	this.CurrencyType = currencyType
	this.Amount = amount
	this.SourceAccountName = sourceAccountName
	this.PayorPaymentId = payorPaymentId
	this.RailsId = railsId
	return &this
}

// NewAcceptedPaymentV3WithDefaults instantiates a new AcceptedPaymentV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptedPaymentV3WithDefaults() *AcceptedPaymentV3 {
	this := AcceptedPaymentV3{}
	return &this
}

// GetRemoteId returns the RemoteId field value
func (o *AcceptedPaymentV3) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *AcceptedPaymentV3) GetRemoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *AcceptedPaymentV3) SetRemoteId(v string) {
	o.RemoteId = v
}

// GetCurrencyType returns the CurrencyType field value
func (o *AcceptedPaymentV3) GetCurrencyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyType
}

// GetCurrencyTypeOk returns a tuple with the CurrencyType field value
// and a boolean to check if the value has been set.
func (o *AcceptedPaymentV3) GetCurrencyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyType, true
}

// SetCurrencyType sets field value
func (o *AcceptedPaymentV3) SetCurrencyType(v string) {
	o.CurrencyType = v
}

// GetAmount returns the Amount field value
func (o *AcceptedPaymentV3) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AcceptedPaymentV3) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AcceptedPaymentV3) SetAmount(v int32) {
	o.Amount = v
}

// GetSourceAccountName returns the SourceAccountName field value
func (o *AcceptedPaymentV3) GetSourceAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAccountName
}

// GetSourceAccountNameOk returns a tuple with the SourceAccountName field value
// and a boolean to check if the value has been set.
func (o *AcceptedPaymentV3) GetSourceAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAccountName, true
}

// SetSourceAccountName sets field value
func (o *AcceptedPaymentV3) SetSourceAccountName(v string) {
	o.SourceAccountName = v
}

// GetPayorPaymentId returns the PayorPaymentId field value
func (o *AcceptedPaymentV3) GetPayorPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayorPaymentId
}

// GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field value
// and a boolean to check if the value has been set.
func (o *AcceptedPaymentV3) GetPayorPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayorPaymentId, true
}

// SetPayorPaymentId sets field value
func (o *AcceptedPaymentV3) SetPayorPaymentId(v string) {
	o.PayorPaymentId = v
}

// GetPaymentMemo returns the PaymentMemo field value if set, zero value otherwise.
func (o *AcceptedPaymentV3) GetPaymentMemo() string {
	if o == nil || IsNil(o.PaymentMemo) {
		var ret string
		return ret
	}
	return *o.PaymentMemo
}

// GetPaymentMemoOk returns a tuple with the PaymentMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptedPaymentV3) GetPaymentMemoOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMemo) {
		return nil, false
	}
	return o.PaymentMemo, true
}

// HasPaymentMemo returns a boolean if a field has been set.
func (o *AcceptedPaymentV3) HasPaymentMemo() bool {
	if o != nil && !IsNil(o.PaymentMemo) {
		return true
	}

	return false
}

// SetPaymentMemo gets a reference to the given string and assigns it to the PaymentMemo field.
func (o *AcceptedPaymentV3) SetPaymentMemo(v string) {
	o.PaymentMemo = &v
}

// GetRemoteSystemId returns the RemoteSystemId field value if set, zero value otherwise.
func (o *AcceptedPaymentV3) GetRemoteSystemId() string {
	if o == nil || IsNil(o.RemoteSystemId) {
		var ret string
		return ret
	}
	return *o.RemoteSystemId
}

// GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptedPaymentV3) GetRemoteSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteSystemId) {
		return nil, false
	}
	return o.RemoteSystemId, true
}

// HasRemoteSystemId returns a boolean if a field has been set.
func (o *AcceptedPaymentV3) HasRemoteSystemId() bool {
	if o != nil && !IsNil(o.RemoteSystemId) {
		return true
	}

	return false
}

// SetRemoteSystemId gets a reference to the given string and assigns it to the RemoteSystemId field.
func (o *AcceptedPaymentV3) SetRemoteSystemId(v string) {
	o.RemoteSystemId = &v
}

// GetPaymentMetadata returns the PaymentMetadata field value if set, zero value otherwise.
func (o *AcceptedPaymentV3) GetPaymentMetadata() string {
	if o == nil || IsNil(o.PaymentMetadata) {
		var ret string
		return ret
	}
	return *o.PaymentMetadata
}

// GetPaymentMetadataOk returns a tuple with the PaymentMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptedPaymentV3) GetPaymentMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMetadata) {
		return nil, false
	}
	return o.PaymentMetadata, true
}

// HasPaymentMetadata returns a boolean if a field has been set.
func (o *AcceptedPaymentV3) HasPaymentMetadata() bool {
	if o != nil && !IsNil(o.PaymentMetadata) {
		return true
	}

	return false
}

// SetPaymentMetadata gets a reference to the given string and assigns it to the PaymentMetadata field.
func (o *AcceptedPaymentV3) SetPaymentMetadata(v string) {
	o.PaymentMetadata = &v
}

// GetRailsId returns the RailsId field value
func (o *AcceptedPaymentV3) GetRailsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RailsId
}

// GetRailsIdOk returns a tuple with the RailsId field value
// and a boolean to check if the value has been set.
func (o *AcceptedPaymentV3) GetRailsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RailsId, true
}

// SetRailsId sets field value
func (o *AcceptedPaymentV3) SetRailsId(v string) {
	o.RailsId = v
}

func (o AcceptedPaymentV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptedPaymentV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["remoteId"] = o.RemoteId
	toSerialize["currencyType"] = o.CurrencyType
	toSerialize["amount"] = o.Amount
	toSerialize["sourceAccountName"] = o.SourceAccountName
	toSerialize["payorPaymentId"] = o.PayorPaymentId
	if !IsNil(o.PaymentMemo) {
		toSerialize["paymentMemo"] = o.PaymentMemo
	}
	if !IsNil(o.RemoteSystemId) {
		toSerialize["remoteSystemId"] = o.RemoteSystemId
	}
	if !IsNil(o.PaymentMetadata) {
		toSerialize["paymentMetadata"] = o.PaymentMetadata
	}
	toSerialize["railsId"] = o.RailsId
	return toSerialize, nil
}

type NullableAcceptedPaymentV3 struct {
	value *AcceptedPaymentV3
	isSet bool
}

func (v NullableAcceptedPaymentV3) Get() *AcceptedPaymentV3 {
	return v.value
}

func (v *NullableAcceptedPaymentV3) Set(val *AcceptedPaymentV3) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptedPaymentV3) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptedPaymentV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptedPaymentV3(val *AcceptedPaymentV3) *NullableAcceptedPaymentV3 {
	return &NullableAcceptedPaymentV3{value: val, isSet: true}
}

func (v NullableAcceptedPaymentV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptedPaymentV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


