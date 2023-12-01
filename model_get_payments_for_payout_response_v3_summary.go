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

// checks if the GetPaymentsForPayoutResponseV3Summary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPaymentsForPayoutResponseV3Summary{}

// GetPaymentsForPayoutResponseV3Summary struct for GetPaymentsForPayoutResponseV3Summary
type GetPaymentsForPayoutResponseV3Summary struct {
	// The current status of the payout. One of the following values: ACCEPTED, REJECTED, SUBMITTED, QUOTED, INSTRUCTED, COMPLETED, INCOMPLETE, CONFIRMED, WITHDRAWN
	PayoutStatus *string `json:"payoutStatus,omitempty"`
	// The date/time at which the payout was submitted.
	SubmittedDateTime *time.Time `json:"submittedDateTime,omitempty"`
	// The date/time at which the payout was instructed.
	InstructedDateTime *time.Time `json:"instructedDateTime,omitempty"`
	// The date/time at which the payout was withdrawn.
	WithdrawnDateTime *time.Time `json:"withdrawnDateTime,omitempty"`
	// The memo attached to the payout.
	PayoutMemo *string `json:"payoutMemo,omitempty"`
	// The count of payments within the payout.
	TotalPayments *int32 `json:"totalPayments,omitempty"`
	// The count of payments within the payout which have been confirmed.
	ConfirmedPayments *int32 `json:"confirmedPayments,omitempty"`
	// The count of payments within the payout which have been released.
	ReleasedPayments *int32 `json:"releasedPayments,omitempty"`
	// The count of payments within the payout which are incomplete.
	IncompletePayments *int32 `json:"incompletePayments,omitempty"`
	// The count of payments within the payout which have failed or been returned.
	FailedPayments *int32 `json:"failedPayments,omitempty"`
}

// NewGetPaymentsForPayoutResponseV3Summary instantiates a new GetPaymentsForPayoutResponseV3Summary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPaymentsForPayoutResponseV3Summary() *GetPaymentsForPayoutResponseV3Summary {
	this := GetPaymentsForPayoutResponseV3Summary{}
	return &this
}

// NewGetPaymentsForPayoutResponseV3SummaryWithDefaults instantiates a new GetPaymentsForPayoutResponseV3Summary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPaymentsForPayoutResponseV3SummaryWithDefaults() *GetPaymentsForPayoutResponseV3Summary {
	this := GetPaymentsForPayoutResponseV3Summary{}
	return &this
}

// GetPayoutStatus returns the PayoutStatus field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetPayoutStatus() string {
	if o == nil || IsNil(o.PayoutStatus) {
		var ret string
		return ret
	}
	return *o.PayoutStatus
}

// GetPayoutStatusOk returns a tuple with the PayoutStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetPayoutStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PayoutStatus) {
		return nil, false
	}
	return o.PayoutStatus, true
}

// HasPayoutStatus returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasPayoutStatus() bool {
	if o != nil && !IsNil(o.PayoutStatus) {
		return true
	}

	return false
}

// SetPayoutStatus gets a reference to the given string and assigns it to the PayoutStatus field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetPayoutStatus(v string) {
	o.PayoutStatus = &v
}

// GetSubmittedDateTime returns the SubmittedDateTime field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetSubmittedDateTime() time.Time {
	if o == nil || IsNil(o.SubmittedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedDateTime
}

// GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetSubmittedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedDateTime) {
		return nil, false
	}
	return o.SubmittedDateTime, true
}

// HasSubmittedDateTime returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasSubmittedDateTime() bool {
	if o != nil && !IsNil(o.SubmittedDateTime) {
		return true
	}

	return false
}

// SetSubmittedDateTime gets a reference to the given time.Time and assigns it to the SubmittedDateTime field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetSubmittedDateTime(v time.Time) {
	o.SubmittedDateTime = &v
}

// GetInstructedDateTime returns the InstructedDateTime field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetInstructedDateTime() time.Time {
	if o == nil || IsNil(o.InstructedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.InstructedDateTime
}

// GetInstructedDateTimeOk returns a tuple with the InstructedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetInstructedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InstructedDateTime) {
		return nil, false
	}
	return o.InstructedDateTime, true
}

// HasInstructedDateTime returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasInstructedDateTime() bool {
	if o != nil && !IsNil(o.InstructedDateTime) {
		return true
	}

	return false
}

// SetInstructedDateTime gets a reference to the given time.Time and assigns it to the InstructedDateTime field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetInstructedDateTime(v time.Time) {
	o.InstructedDateTime = &v
}

// GetWithdrawnDateTime returns the WithdrawnDateTime field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetWithdrawnDateTime() time.Time {
	if o == nil || IsNil(o.WithdrawnDateTime) {
		var ret time.Time
		return ret
	}
	return *o.WithdrawnDateTime
}

// GetWithdrawnDateTimeOk returns a tuple with the WithdrawnDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetWithdrawnDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.WithdrawnDateTime) {
		return nil, false
	}
	return o.WithdrawnDateTime, true
}

// HasWithdrawnDateTime returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasWithdrawnDateTime() bool {
	if o != nil && !IsNil(o.WithdrawnDateTime) {
		return true
	}

	return false
}

// SetWithdrawnDateTime gets a reference to the given time.Time and assigns it to the WithdrawnDateTime field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetWithdrawnDateTime(v time.Time) {
	o.WithdrawnDateTime = &v
}

// GetPayoutMemo returns the PayoutMemo field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetPayoutMemo() string {
	if o == nil || IsNil(o.PayoutMemo) {
		var ret string
		return ret
	}
	return *o.PayoutMemo
}

// GetPayoutMemoOk returns a tuple with the PayoutMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetPayoutMemoOk() (*string, bool) {
	if o == nil || IsNil(o.PayoutMemo) {
		return nil, false
	}
	return o.PayoutMemo, true
}

// HasPayoutMemo returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasPayoutMemo() bool {
	if o != nil && !IsNil(o.PayoutMemo) {
		return true
	}

	return false
}

// SetPayoutMemo gets a reference to the given string and assigns it to the PayoutMemo field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetPayoutMemo(v string) {
	o.PayoutMemo = &v
}

// GetTotalPayments returns the TotalPayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetTotalPayments() int32 {
	if o == nil || IsNil(o.TotalPayments) {
		var ret int32
		return ret
	}
	return *o.TotalPayments
}

// GetTotalPaymentsOk returns a tuple with the TotalPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetTotalPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPayments) {
		return nil, false
	}
	return o.TotalPayments, true
}

// HasTotalPayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasTotalPayments() bool {
	if o != nil && !IsNil(o.TotalPayments) {
		return true
	}

	return false
}

// SetTotalPayments gets a reference to the given int32 and assigns it to the TotalPayments field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetTotalPayments(v int32) {
	o.TotalPayments = &v
}

// GetConfirmedPayments returns the ConfirmedPayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetConfirmedPayments() int32 {
	if o == nil || IsNil(o.ConfirmedPayments) {
		var ret int32
		return ret
	}
	return *o.ConfirmedPayments
}

// GetConfirmedPaymentsOk returns a tuple with the ConfirmedPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetConfirmedPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfirmedPayments) {
		return nil, false
	}
	return o.ConfirmedPayments, true
}

// HasConfirmedPayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasConfirmedPayments() bool {
	if o != nil && !IsNil(o.ConfirmedPayments) {
		return true
	}

	return false
}

// SetConfirmedPayments gets a reference to the given int32 and assigns it to the ConfirmedPayments field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetConfirmedPayments(v int32) {
	o.ConfirmedPayments = &v
}

// GetReleasedPayments returns the ReleasedPayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetReleasedPayments() int32 {
	if o == nil || IsNil(o.ReleasedPayments) {
		var ret int32
		return ret
	}
	return *o.ReleasedPayments
}

// GetReleasedPaymentsOk returns a tuple with the ReleasedPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetReleasedPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.ReleasedPayments) {
		return nil, false
	}
	return o.ReleasedPayments, true
}

// HasReleasedPayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasReleasedPayments() bool {
	if o != nil && !IsNil(o.ReleasedPayments) {
		return true
	}

	return false
}

// SetReleasedPayments gets a reference to the given int32 and assigns it to the ReleasedPayments field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetReleasedPayments(v int32) {
	o.ReleasedPayments = &v
}

// GetIncompletePayments returns the IncompletePayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetIncompletePayments() int32 {
	if o == nil || IsNil(o.IncompletePayments) {
		var ret int32
		return ret
	}
	return *o.IncompletePayments
}

// GetIncompletePaymentsOk returns a tuple with the IncompletePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetIncompletePaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.IncompletePayments) {
		return nil, false
	}
	return o.IncompletePayments, true
}

// HasIncompletePayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasIncompletePayments() bool {
	if o != nil && !IsNil(o.IncompletePayments) {
		return true
	}

	return false
}

// SetIncompletePayments gets a reference to the given int32 and assigns it to the IncompletePayments field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetIncompletePayments(v int32) {
	o.IncompletePayments = &v
}

// GetFailedPayments returns the FailedPayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV3Summary) GetFailedPayments() int32 {
	if o == nil || IsNil(o.FailedPayments) {
		var ret int32
		return ret
	}
	return *o.FailedPayments
}

// GetFailedPaymentsOk returns a tuple with the FailedPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) GetFailedPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.FailedPayments) {
		return nil, false
	}
	return o.FailedPayments, true
}

// HasFailedPayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV3Summary) HasFailedPayments() bool {
	if o != nil && !IsNil(o.FailedPayments) {
		return true
	}

	return false
}

// SetFailedPayments gets a reference to the given int32 and assigns it to the FailedPayments field.
func (o *GetPaymentsForPayoutResponseV3Summary) SetFailedPayments(v int32) {
	o.FailedPayments = &v
}

func (o GetPaymentsForPayoutResponseV3Summary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPaymentsForPayoutResponseV3Summary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayoutStatus) {
		toSerialize["payoutStatus"] = o.PayoutStatus
	}
	if !IsNil(o.SubmittedDateTime) {
		toSerialize["submittedDateTime"] = o.SubmittedDateTime
	}
	if !IsNil(o.InstructedDateTime) {
		toSerialize["instructedDateTime"] = o.InstructedDateTime
	}
	if !IsNil(o.WithdrawnDateTime) {
		toSerialize["withdrawnDateTime"] = o.WithdrawnDateTime
	}
	if !IsNil(o.PayoutMemo) {
		toSerialize["payoutMemo"] = o.PayoutMemo
	}
	if !IsNil(o.TotalPayments) {
		toSerialize["totalPayments"] = o.TotalPayments
	}
	if !IsNil(o.ConfirmedPayments) {
		toSerialize["confirmedPayments"] = o.ConfirmedPayments
	}
	if !IsNil(o.ReleasedPayments) {
		toSerialize["releasedPayments"] = o.ReleasedPayments
	}
	if !IsNil(o.IncompletePayments) {
		toSerialize["incompletePayments"] = o.IncompletePayments
	}
	if !IsNil(o.FailedPayments) {
		toSerialize["failedPayments"] = o.FailedPayments
	}
	return toSerialize, nil
}

type NullableGetPaymentsForPayoutResponseV3Summary struct {
	value *GetPaymentsForPayoutResponseV3Summary
	isSet bool
}

func (v NullableGetPaymentsForPayoutResponseV3Summary) Get() *GetPaymentsForPayoutResponseV3Summary {
	return v.value
}

func (v *NullableGetPaymentsForPayoutResponseV3Summary) Set(val *GetPaymentsForPayoutResponseV3Summary) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPaymentsForPayoutResponseV3Summary) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPaymentsForPayoutResponseV3Summary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPaymentsForPayoutResponseV3Summary(val *GetPaymentsForPayoutResponseV3Summary) *NullableGetPaymentsForPayoutResponseV3Summary {
	return &NullableGetPaymentsForPayoutResponseV3Summary{value: val, isSet: true}
}

func (v NullableGetPaymentsForPayoutResponseV3Summary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPaymentsForPayoutResponseV3Summary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


