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

// checks if the PayeePayorRefV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayeePayorRefV3{}

// PayeePayorRefV3 struct for PayeePayorRefV3
type PayeePayorRefV3 struct {
	PayorId *string `json:"payorId,omitempty"`
	RemoteId *string `json:"remoteId,omitempty"`
	InvitationStatus *string `json:"invitationStatus,omitempty"`
	// The timestamp when the invitation status is updated
	InvitationStatusTimestamp NullableTime `json:"invitationStatusTimestamp,omitempty"`
	PaymentChannelId *string `json:"paymentChannelId,omitempty"`
	// Indicates if the payee is payable for this payor
	PayableStatus *bool `json:"payableStatus,omitempty"`
	// Indicates any conditions which prevent the payee from being payable for this payor
	PayableIssues []PayableIssueV3 `json:"payableIssues,omitempty"`
}

// NewPayeePayorRefV3 instantiates a new PayeePayorRefV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayeePayorRefV3() *PayeePayorRefV3 {
	this := PayeePayorRefV3{}
	return &this
}

// NewPayeePayorRefV3WithDefaults instantiates a new PayeePayorRefV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayeePayorRefV3WithDefaults() *PayeePayorRefV3 {
	this := PayeePayorRefV3{}
	return &this
}

// GetPayorId returns the PayorId field value if set, zero value otherwise.
func (o *PayeePayorRefV3) GetPayorId() string {
	if o == nil || IsNil(o.PayorId) {
		var ret string
		return ret
	}
	return *o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeePayorRefV3) GetPayorIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayorId) {
		return nil, false
	}
	return o.PayorId, true
}

// HasPayorId returns a boolean if a field has been set.
func (o *PayeePayorRefV3) HasPayorId() bool {
	if o != nil && !IsNil(o.PayorId) {
		return true
	}

	return false
}

// SetPayorId gets a reference to the given string and assigns it to the PayorId field.
func (o *PayeePayorRefV3) SetPayorId(v string) {
	o.PayorId = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *PayeePayorRefV3) GetRemoteId() string {
	if o == nil || IsNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeePayorRefV3) GetRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteId) {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *PayeePayorRefV3) HasRemoteId() bool {
	if o != nil && !IsNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *PayeePayorRefV3) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetInvitationStatus returns the InvitationStatus field value if set, zero value otherwise.
func (o *PayeePayorRefV3) GetInvitationStatus() string {
	if o == nil || IsNil(o.InvitationStatus) {
		var ret string
		return ret
	}
	return *o.InvitationStatus
}

// GetInvitationStatusOk returns a tuple with the InvitationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeePayorRefV3) GetInvitationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.InvitationStatus) {
		return nil, false
	}
	return o.InvitationStatus, true
}

// HasInvitationStatus returns a boolean if a field has been set.
func (o *PayeePayorRefV3) HasInvitationStatus() bool {
	if o != nil && !IsNil(o.InvitationStatus) {
		return true
	}

	return false
}

// SetInvitationStatus gets a reference to the given string and assigns it to the InvitationStatus field.
func (o *PayeePayorRefV3) SetInvitationStatus(v string) {
	o.InvitationStatus = &v
}

// GetInvitationStatusTimestamp returns the InvitationStatusTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeePayorRefV3) GetInvitationStatusTimestamp() time.Time {
	if o == nil || IsNil(o.InvitationStatusTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.InvitationStatusTimestamp.Get()
}

// GetInvitationStatusTimestampOk returns a tuple with the InvitationStatusTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeePayorRefV3) GetInvitationStatusTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvitationStatusTimestamp.Get(), o.InvitationStatusTimestamp.IsSet()
}

// HasInvitationStatusTimestamp returns a boolean if a field has been set.
func (o *PayeePayorRefV3) HasInvitationStatusTimestamp() bool {
	if o != nil && o.InvitationStatusTimestamp.IsSet() {
		return true
	}

	return false
}

// SetInvitationStatusTimestamp gets a reference to the given NullableTime and assigns it to the InvitationStatusTimestamp field.
func (o *PayeePayorRefV3) SetInvitationStatusTimestamp(v time.Time) {
	o.InvitationStatusTimestamp.Set(&v)
}
// SetInvitationStatusTimestampNil sets the value for InvitationStatusTimestamp to be an explicit nil
func (o *PayeePayorRefV3) SetInvitationStatusTimestampNil() {
	o.InvitationStatusTimestamp.Set(nil)
}

// UnsetInvitationStatusTimestamp ensures that no value is present for InvitationStatusTimestamp, not even an explicit nil
func (o *PayeePayorRefV3) UnsetInvitationStatusTimestamp() {
	o.InvitationStatusTimestamp.Unset()
}

// GetPaymentChannelId returns the PaymentChannelId field value if set, zero value otherwise.
func (o *PayeePayorRefV3) GetPaymentChannelId() string {
	if o == nil || IsNil(o.PaymentChannelId) {
		var ret string
		return ret
	}
	return *o.PaymentChannelId
}

// GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeePayorRefV3) GetPaymentChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelId) {
		return nil, false
	}
	return o.PaymentChannelId, true
}

// HasPaymentChannelId returns a boolean if a field has been set.
func (o *PayeePayorRefV3) HasPaymentChannelId() bool {
	if o != nil && !IsNil(o.PaymentChannelId) {
		return true
	}

	return false
}

// SetPaymentChannelId gets a reference to the given string and assigns it to the PaymentChannelId field.
func (o *PayeePayorRefV3) SetPaymentChannelId(v string) {
	o.PaymentChannelId = &v
}

// GetPayableStatus returns the PayableStatus field value if set, zero value otherwise.
func (o *PayeePayorRefV3) GetPayableStatus() bool {
	if o == nil || IsNil(o.PayableStatus) {
		var ret bool
		return ret
	}
	return *o.PayableStatus
}

// GetPayableStatusOk returns a tuple with the PayableStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeePayorRefV3) GetPayableStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.PayableStatus) {
		return nil, false
	}
	return o.PayableStatus, true
}

// HasPayableStatus returns a boolean if a field has been set.
func (o *PayeePayorRefV3) HasPayableStatus() bool {
	if o != nil && !IsNil(o.PayableStatus) {
		return true
	}

	return false
}

// SetPayableStatus gets a reference to the given bool and assigns it to the PayableStatus field.
func (o *PayeePayorRefV3) SetPayableStatus(v bool) {
	o.PayableStatus = &v
}

// GetPayableIssues returns the PayableIssues field value if set, zero value otherwise.
func (o *PayeePayorRefV3) GetPayableIssues() []PayableIssueV3 {
	if o == nil || IsNil(o.PayableIssues) {
		var ret []PayableIssueV3
		return ret
	}
	return o.PayableIssues
}

// GetPayableIssuesOk returns a tuple with the PayableIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeePayorRefV3) GetPayableIssuesOk() ([]PayableIssueV3, bool) {
	if o == nil || IsNil(o.PayableIssues) {
		return nil, false
	}
	return o.PayableIssues, true
}

// HasPayableIssues returns a boolean if a field has been set.
func (o *PayeePayorRefV3) HasPayableIssues() bool {
	if o != nil && !IsNil(o.PayableIssues) {
		return true
	}

	return false
}

// SetPayableIssues gets a reference to the given []PayableIssueV3 and assigns it to the PayableIssues field.
func (o *PayeePayorRefV3) SetPayableIssues(v []PayableIssueV3) {
	o.PayableIssues = v
}

func (o PayeePayorRefV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayeePayorRefV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayorId) {
		toSerialize["payorId"] = o.PayorId
	}
	if !IsNil(o.RemoteId) {
		toSerialize["remoteId"] = o.RemoteId
	}
	if !IsNil(o.InvitationStatus) {
		toSerialize["invitationStatus"] = o.InvitationStatus
	}
	if o.InvitationStatusTimestamp.IsSet() {
		toSerialize["invitationStatusTimestamp"] = o.InvitationStatusTimestamp.Get()
	}
	if !IsNil(o.PaymentChannelId) {
		toSerialize["paymentChannelId"] = o.PaymentChannelId
	}
	if !IsNil(o.PayableStatus) {
		toSerialize["payableStatus"] = o.PayableStatus
	}
	if !IsNil(o.PayableIssues) {
		toSerialize["payableIssues"] = o.PayableIssues
	}
	return toSerialize, nil
}

type NullablePayeePayorRefV3 struct {
	value *PayeePayorRefV3
	isSet bool
}

func (v NullablePayeePayorRefV3) Get() *PayeePayorRefV3 {
	return v.value
}

func (v *NullablePayeePayorRefV3) Set(val *PayeePayorRefV3) {
	v.value = val
	v.isSet = true
}

func (v NullablePayeePayorRefV3) IsSet() bool {
	return v.isSet
}

func (v *NullablePayeePayorRefV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayeePayorRefV3(val *PayeePayorRefV3) *NullablePayeePayorRefV3 {
	return &NullablePayeePayorRefV3{value: val, isSet: true}
}

func (v NullablePayeePayorRefV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayeePayorRefV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


