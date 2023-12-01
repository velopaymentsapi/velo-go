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

// checks if the PayoutSummaryResponseV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayoutSummaryResponseV3{}

// PayoutSummaryResponseV3 struct for PayoutSummaryResponseV3
type PayoutSummaryResponseV3 struct {
	// The id of the payout
	PayoutId *string `json:"payoutId,omitempty"`
	// The status of the payout (one of SUBMITTED, ACCEPTED, REJECTED, QUOTED, INSTRUCTED, COMPLETED, INCOMPLETE, WITHDRAWN)
	Status *string `json:"status,omitempty"`
	// The number of payments that were submitted in the payout
	PaymentsSubmitted *int32 `json:"paymentsSubmitted,omitempty"`
	// The number of payments that were accepted in the payout
	PaymentsAccepted *int32 `json:"paymentsAccepted,omitempty"`
	// The number of payments that were rejected in the payout
	PaymentsRejected *int32 `json:"paymentsRejected,omitempty"`
	// The number of payments that were withdrawn in the payout
	PaymentsWithdrawn int32 `json:"paymentsWithdrawn"`
	FxSummaries []QuoteFxSummaryV3 `json:"fxSummaries"`
	Accounts []SourceAccountV3 `json:"accounts"`
	AcceptedPayments []AcceptedPaymentV3 `json:"acceptedPayments"`
	RejectedPayments []RejectedPaymentV3 `json:"rejectedPayments"`
	Schedule *PayoutScheduleV3 `json:"schedule,omitempty"`
}

// NewPayoutSummaryResponseV3 instantiates a new PayoutSummaryResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutSummaryResponseV3(paymentsWithdrawn int32, fxSummaries []QuoteFxSummaryV3, accounts []SourceAccountV3, acceptedPayments []AcceptedPaymentV3, rejectedPayments []RejectedPaymentV3) *PayoutSummaryResponseV3 {
	this := PayoutSummaryResponseV3{}
	this.PaymentsWithdrawn = paymentsWithdrawn
	this.FxSummaries = fxSummaries
	this.Accounts = accounts
	this.AcceptedPayments = acceptedPayments
	this.RejectedPayments = rejectedPayments
	return &this
}

// NewPayoutSummaryResponseV3WithDefaults instantiates a new PayoutSummaryResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutSummaryResponseV3WithDefaults() *PayoutSummaryResponseV3 {
	this := PayoutSummaryResponseV3{}
	return &this
}

// GetPayoutId returns the PayoutId field value if set, zero value otherwise.
func (o *PayoutSummaryResponseV3) GetPayoutId() string {
	if o == nil || IsNil(o.PayoutId) {
		var ret string
		return ret
	}
	return *o.PayoutId
}

// GetPayoutIdOk returns a tuple with the PayoutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetPayoutIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayoutId) {
		return nil, false
	}
	return o.PayoutId, true
}

// HasPayoutId returns a boolean if a field has been set.
func (o *PayoutSummaryResponseV3) HasPayoutId() bool {
	if o != nil && !IsNil(o.PayoutId) {
		return true
	}

	return false
}

// SetPayoutId gets a reference to the given string and assigns it to the PayoutId field.
func (o *PayoutSummaryResponseV3) SetPayoutId(v string) {
	o.PayoutId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PayoutSummaryResponseV3) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PayoutSummaryResponseV3) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PayoutSummaryResponseV3) SetStatus(v string) {
	o.Status = &v
}

// GetPaymentsSubmitted returns the PaymentsSubmitted field value if set, zero value otherwise.
func (o *PayoutSummaryResponseV3) GetPaymentsSubmitted() int32 {
	if o == nil || IsNil(o.PaymentsSubmitted) {
		var ret int32
		return ret
	}
	return *o.PaymentsSubmitted
}

// GetPaymentsSubmittedOk returns a tuple with the PaymentsSubmitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetPaymentsSubmittedOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentsSubmitted) {
		return nil, false
	}
	return o.PaymentsSubmitted, true
}

// HasPaymentsSubmitted returns a boolean if a field has been set.
func (o *PayoutSummaryResponseV3) HasPaymentsSubmitted() bool {
	if o != nil && !IsNil(o.PaymentsSubmitted) {
		return true
	}

	return false
}

// SetPaymentsSubmitted gets a reference to the given int32 and assigns it to the PaymentsSubmitted field.
func (o *PayoutSummaryResponseV3) SetPaymentsSubmitted(v int32) {
	o.PaymentsSubmitted = &v
}

// GetPaymentsAccepted returns the PaymentsAccepted field value if set, zero value otherwise.
func (o *PayoutSummaryResponseV3) GetPaymentsAccepted() int32 {
	if o == nil || IsNil(o.PaymentsAccepted) {
		var ret int32
		return ret
	}
	return *o.PaymentsAccepted
}

// GetPaymentsAcceptedOk returns a tuple with the PaymentsAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetPaymentsAcceptedOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentsAccepted) {
		return nil, false
	}
	return o.PaymentsAccepted, true
}

// HasPaymentsAccepted returns a boolean if a field has been set.
func (o *PayoutSummaryResponseV3) HasPaymentsAccepted() bool {
	if o != nil && !IsNil(o.PaymentsAccepted) {
		return true
	}

	return false
}

// SetPaymentsAccepted gets a reference to the given int32 and assigns it to the PaymentsAccepted field.
func (o *PayoutSummaryResponseV3) SetPaymentsAccepted(v int32) {
	o.PaymentsAccepted = &v
}

// GetPaymentsRejected returns the PaymentsRejected field value if set, zero value otherwise.
func (o *PayoutSummaryResponseV3) GetPaymentsRejected() int32 {
	if o == nil || IsNil(o.PaymentsRejected) {
		var ret int32
		return ret
	}
	return *o.PaymentsRejected
}

// GetPaymentsRejectedOk returns a tuple with the PaymentsRejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetPaymentsRejectedOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentsRejected) {
		return nil, false
	}
	return o.PaymentsRejected, true
}

// HasPaymentsRejected returns a boolean if a field has been set.
func (o *PayoutSummaryResponseV3) HasPaymentsRejected() bool {
	if o != nil && !IsNil(o.PaymentsRejected) {
		return true
	}

	return false
}

// SetPaymentsRejected gets a reference to the given int32 and assigns it to the PaymentsRejected field.
func (o *PayoutSummaryResponseV3) SetPaymentsRejected(v int32) {
	o.PaymentsRejected = &v
}

// GetPaymentsWithdrawn returns the PaymentsWithdrawn field value
func (o *PayoutSummaryResponseV3) GetPaymentsWithdrawn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PaymentsWithdrawn
}

// GetPaymentsWithdrawnOk returns a tuple with the PaymentsWithdrawn field value
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetPaymentsWithdrawnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentsWithdrawn, true
}

// SetPaymentsWithdrawn sets field value
func (o *PayoutSummaryResponseV3) SetPaymentsWithdrawn(v int32) {
	o.PaymentsWithdrawn = v
}

// GetFxSummaries returns the FxSummaries field value
func (o *PayoutSummaryResponseV3) GetFxSummaries() []QuoteFxSummaryV3 {
	if o == nil {
		var ret []QuoteFxSummaryV3
		return ret
	}

	return o.FxSummaries
}

// GetFxSummariesOk returns a tuple with the FxSummaries field value
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetFxSummariesOk() ([]QuoteFxSummaryV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.FxSummaries, true
}

// SetFxSummaries sets field value
func (o *PayoutSummaryResponseV3) SetFxSummaries(v []QuoteFxSummaryV3) {
	o.FxSummaries = v
}

// GetAccounts returns the Accounts field value
func (o *PayoutSummaryResponseV3) GetAccounts() []SourceAccountV3 {
	if o == nil {
		var ret []SourceAccountV3
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetAccountsOk() ([]SourceAccountV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *PayoutSummaryResponseV3) SetAccounts(v []SourceAccountV3) {
	o.Accounts = v
}

// GetAcceptedPayments returns the AcceptedPayments field value
func (o *PayoutSummaryResponseV3) GetAcceptedPayments() []AcceptedPaymentV3 {
	if o == nil {
		var ret []AcceptedPaymentV3
		return ret
	}

	return o.AcceptedPayments
}

// GetAcceptedPaymentsOk returns a tuple with the AcceptedPayments field value
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetAcceptedPaymentsOk() ([]AcceptedPaymentV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedPayments, true
}

// SetAcceptedPayments sets field value
func (o *PayoutSummaryResponseV3) SetAcceptedPayments(v []AcceptedPaymentV3) {
	o.AcceptedPayments = v
}

// GetRejectedPayments returns the RejectedPayments field value
func (o *PayoutSummaryResponseV3) GetRejectedPayments() []RejectedPaymentV3 {
	if o == nil {
		var ret []RejectedPaymentV3
		return ret
	}

	return o.RejectedPayments
}

// GetRejectedPaymentsOk returns a tuple with the RejectedPayments field value
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetRejectedPaymentsOk() ([]RejectedPaymentV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.RejectedPayments, true
}

// SetRejectedPayments sets field value
func (o *PayoutSummaryResponseV3) SetRejectedPayments(v []RejectedPaymentV3) {
	o.RejectedPayments = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *PayoutSummaryResponseV3) GetSchedule() PayoutScheduleV3 {
	if o == nil || IsNil(o.Schedule) {
		var ret PayoutScheduleV3
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryResponseV3) GetScheduleOk() (*PayoutScheduleV3, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *PayoutSummaryResponseV3) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given PayoutScheduleV3 and assigns it to the Schedule field.
func (o *PayoutSummaryResponseV3) SetSchedule(v PayoutScheduleV3) {
	o.Schedule = &v
}

func (o PayoutSummaryResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutSummaryResponseV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayoutId) {
		toSerialize["payoutId"] = o.PayoutId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.PaymentsSubmitted) {
		toSerialize["paymentsSubmitted"] = o.PaymentsSubmitted
	}
	if !IsNil(o.PaymentsAccepted) {
		toSerialize["paymentsAccepted"] = o.PaymentsAccepted
	}
	if !IsNil(o.PaymentsRejected) {
		toSerialize["paymentsRejected"] = o.PaymentsRejected
	}
	toSerialize["paymentsWithdrawn"] = o.PaymentsWithdrawn
	toSerialize["fxSummaries"] = o.FxSummaries
	toSerialize["accounts"] = o.Accounts
	toSerialize["acceptedPayments"] = o.AcceptedPayments
	toSerialize["rejectedPayments"] = o.RejectedPayments
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	return toSerialize, nil
}

type NullablePayoutSummaryResponseV3 struct {
	value *PayoutSummaryResponseV3
	isSet bool
}

func (v NullablePayoutSummaryResponseV3) Get() *PayoutSummaryResponseV3 {
	return v.value
}

func (v *NullablePayoutSummaryResponseV3) Set(val *PayoutSummaryResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutSummaryResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutSummaryResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutSummaryResponseV3(val *PayoutSummaryResponseV3) *NullablePayoutSummaryResponseV3 {
	return &NullablePayoutSummaryResponseV3{value: val, isSet: true}
}

func (v NullablePayoutSummaryResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutSummaryResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


