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

// checks if the PayoutSummaryAudit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayoutSummaryAudit{}

// PayoutSummaryAudit struct for PayoutSummaryAudit
type PayoutSummaryAudit struct {
	PayoutId *string `json:"payoutId,omitempty"`
	PayorId *string `json:"payorId,omitempty"`
	// Current status of the Payout. One of the following values: ACCEPTED, REJECTED, SUBMITTED, QUOTED, INSTRUCTED, COMPLETED, INCOMPLETE, CONFIRMED, WITHDRAWN
	Status string `json:"status"`
	DateTime *time.Time `json:"dateTime,omitempty"`
	SubmittedDateTime string `json:"submittedDateTime"`
	InstructedDateTime *string `json:"instructedDateTime,omitempty"`
	WithdrawnDateTime *time.Time `json:"withdrawnDateTime,omitempty"`
	TotalPayments *int32 `json:"totalPayments,omitempty"`
	TotalIncompletePayments *int32 `json:"totalIncompletePayments,omitempty"`
	TotalReturnedPayments *int32 `json:"totalReturnedPayments,omitempty"`
	TotalWithdrawnPayments *int32 `json:"totalWithdrawnPayments,omitempty"`
	SourceAccountSummary []SourceAccountSummary `json:"sourceAccountSummary,omitempty"`
	FxSummaries []FxSummary `json:"fxSummaries,omitempty"`
	PayoutMemo *string `json:"payoutMemo,omitempty"`
	// The type of payout. One of the following values: STANDARD, AS, ON_BEHALF_OF
	PayoutType string `json:"payoutType"`
	PayorName string `json:"payorName"`
	Schedule *PayoutSchedule `json:"schedule,omitempty"`
}

// NewPayoutSummaryAudit instantiates a new PayoutSummaryAudit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutSummaryAudit(status string, submittedDateTime string, payoutType string, payorName string) *PayoutSummaryAudit {
	this := PayoutSummaryAudit{}
	this.Status = status
	this.SubmittedDateTime = submittedDateTime
	this.PayoutType = payoutType
	this.PayorName = payorName
	return &this
}

// NewPayoutSummaryAuditWithDefaults instantiates a new PayoutSummaryAudit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutSummaryAuditWithDefaults() *PayoutSummaryAudit {
	this := PayoutSummaryAudit{}
	return &this
}

// GetPayoutId returns the PayoutId field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetPayoutId() string {
	if o == nil || IsNil(o.PayoutId) {
		var ret string
		return ret
	}
	return *o.PayoutId
}

// GetPayoutIdOk returns a tuple with the PayoutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetPayoutIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayoutId) {
		return nil, false
	}
	return o.PayoutId, true
}

// HasPayoutId returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasPayoutId() bool {
	if o != nil && !IsNil(o.PayoutId) {
		return true
	}

	return false
}

// SetPayoutId gets a reference to the given string and assigns it to the PayoutId field.
func (o *PayoutSummaryAudit) SetPayoutId(v string) {
	o.PayoutId = &v
}

// GetPayorId returns the PayorId field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetPayorId() string {
	if o == nil || IsNil(o.PayorId) {
		var ret string
		return ret
	}
	return *o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetPayorIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayorId) {
		return nil, false
	}
	return o.PayorId, true
}

// HasPayorId returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasPayorId() bool {
	if o != nil && !IsNil(o.PayorId) {
		return true
	}

	return false
}

// SetPayorId gets a reference to the given string and assigns it to the PayorId field.
func (o *PayoutSummaryAudit) SetPayorId(v string) {
	o.PayorId = &v
}

// GetStatus returns the Status field value
func (o *PayoutSummaryAudit) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PayoutSummaryAudit) SetStatus(v string) {
	o.Status = v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetDateTime() time.Time {
	if o == nil || IsNil(o.DateTime) {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTime) {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasDateTime() bool {
	if o != nil && !IsNil(o.DateTime) {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *PayoutSummaryAudit) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetSubmittedDateTime returns the SubmittedDateTime field value
func (o *PayoutSummaryAudit) GetSubmittedDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubmittedDateTime
}

// GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field value
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetSubmittedDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubmittedDateTime, true
}

// SetSubmittedDateTime sets field value
func (o *PayoutSummaryAudit) SetSubmittedDateTime(v string) {
	o.SubmittedDateTime = v
}

// GetInstructedDateTime returns the InstructedDateTime field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetInstructedDateTime() string {
	if o == nil || IsNil(o.InstructedDateTime) {
		var ret string
		return ret
	}
	return *o.InstructedDateTime
}

// GetInstructedDateTimeOk returns a tuple with the InstructedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetInstructedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.InstructedDateTime) {
		return nil, false
	}
	return o.InstructedDateTime, true
}

// HasInstructedDateTime returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasInstructedDateTime() bool {
	if o != nil && !IsNil(o.InstructedDateTime) {
		return true
	}

	return false
}

// SetInstructedDateTime gets a reference to the given string and assigns it to the InstructedDateTime field.
func (o *PayoutSummaryAudit) SetInstructedDateTime(v string) {
	o.InstructedDateTime = &v
}

// GetWithdrawnDateTime returns the WithdrawnDateTime field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetWithdrawnDateTime() time.Time {
	if o == nil || IsNil(o.WithdrawnDateTime) {
		var ret time.Time
		return ret
	}
	return *o.WithdrawnDateTime
}

// GetWithdrawnDateTimeOk returns a tuple with the WithdrawnDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetWithdrawnDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.WithdrawnDateTime) {
		return nil, false
	}
	return o.WithdrawnDateTime, true
}

// HasWithdrawnDateTime returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasWithdrawnDateTime() bool {
	if o != nil && !IsNil(o.WithdrawnDateTime) {
		return true
	}

	return false
}

// SetWithdrawnDateTime gets a reference to the given time.Time and assigns it to the WithdrawnDateTime field.
func (o *PayoutSummaryAudit) SetWithdrawnDateTime(v time.Time) {
	o.WithdrawnDateTime = &v
}

// GetTotalPayments returns the TotalPayments field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetTotalPayments() int32 {
	if o == nil || IsNil(o.TotalPayments) {
		var ret int32
		return ret
	}
	return *o.TotalPayments
}

// GetTotalPaymentsOk returns a tuple with the TotalPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetTotalPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPayments) {
		return nil, false
	}
	return o.TotalPayments, true
}

// HasTotalPayments returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasTotalPayments() bool {
	if o != nil && !IsNil(o.TotalPayments) {
		return true
	}

	return false
}

// SetTotalPayments gets a reference to the given int32 and assigns it to the TotalPayments field.
func (o *PayoutSummaryAudit) SetTotalPayments(v int32) {
	o.TotalPayments = &v
}

// GetTotalIncompletePayments returns the TotalIncompletePayments field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetTotalIncompletePayments() int32 {
	if o == nil || IsNil(o.TotalIncompletePayments) {
		var ret int32
		return ret
	}
	return *o.TotalIncompletePayments
}

// GetTotalIncompletePaymentsOk returns a tuple with the TotalIncompletePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetTotalIncompletePaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalIncompletePayments) {
		return nil, false
	}
	return o.TotalIncompletePayments, true
}

// HasTotalIncompletePayments returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasTotalIncompletePayments() bool {
	if o != nil && !IsNil(o.TotalIncompletePayments) {
		return true
	}

	return false
}

// SetTotalIncompletePayments gets a reference to the given int32 and assigns it to the TotalIncompletePayments field.
func (o *PayoutSummaryAudit) SetTotalIncompletePayments(v int32) {
	o.TotalIncompletePayments = &v
}

// GetTotalReturnedPayments returns the TotalReturnedPayments field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetTotalReturnedPayments() int32 {
	if o == nil || IsNil(o.TotalReturnedPayments) {
		var ret int32
		return ret
	}
	return *o.TotalReturnedPayments
}

// GetTotalReturnedPaymentsOk returns a tuple with the TotalReturnedPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetTotalReturnedPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalReturnedPayments) {
		return nil, false
	}
	return o.TotalReturnedPayments, true
}

// HasTotalReturnedPayments returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasTotalReturnedPayments() bool {
	if o != nil && !IsNil(o.TotalReturnedPayments) {
		return true
	}

	return false
}

// SetTotalReturnedPayments gets a reference to the given int32 and assigns it to the TotalReturnedPayments field.
func (o *PayoutSummaryAudit) SetTotalReturnedPayments(v int32) {
	o.TotalReturnedPayments = &v
}

// GetTotalWithdrawnPayments returns the TotalWithdrawnPayments field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetTotalWithdrawnPayments() int32 {
	if o == nil || IsNil(o.TotalWithdrawnPayments) {
		var ret int32
		return ret
	}
	return *o.TotalWithdrawnPayments
}

// GetTotalWithdrawnPaymentsOk returns a tuple with the TotalWithdrawnPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetTotalWithdrawnPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalWithdrawnPayments) {
		return nil, false
	}
	return o.TotalWithdrawnPayments, true
}

// HasTotalWithdrawnPayments returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasTotalWithdrawnPayments() bool {
	if o != nil && !IsNil(o.TotalWithdrawnPayments) {
		return true
	}

	return false
}

// SetTotalWithdrawnPayments gets a reference to the given int32 and assigns it to the TotalWithdrawnPayments field.
func (o *PayoutSummaryAudit) SetTotalWithdrawnPayments(v int32) {
	o.TotalWithdrawnPayments = &v
}

// GetSourceAccountSummary returns the SourceAccountSummary field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetSourceAccountSummary() []SourceAccountSummary {
	if o == nil || IsNil(o.SourceAccountSummary) {
		var ret []SourceAccountSummary
		return ret
	}
	return o.SourceAccountSummary
}

// GetSourceAccountSummaryOk returns a tuple with the SourceAccountSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetSourceAccountSummaryOk() ([]SourceAccountSummary, bool) {
	if o == nil || IsNil(o.SourceAccountSummary) {
		return nil, false
	}
	return o.SourceAccountSummary, true
}

// HasSourceAccountSummary returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasSourceAccountSummary() bool {
	if o != nil && !IsNil(o.SourceAccountSummary) {
		return true
	}

	return false
}

// SetSourceAccountSummary gets a reference to the given []SourceAccountSummary and assigns it to the SourceAccountSummary field.
func (o *PayoutSummaryAudit) SetSourceAccountSummary(v []SourceAccountSummary) {
	o.SourceAccountSummary = v
}

// GetFxSummaries returns the FxSummaries field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetFxSummaries() []FxSummary {
	if o == nil || IsNil(o.FxSummaries) {
		var ret []FxSummary
		return ret
	}
	return o.FxSummaries
}

// GetFxSummariesOk returns a tuple with the FxSummaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetFxSummariesOk() ([]FxSummary, bool) {
	if o == nil || IsNil(o.FxSummaries) {
		return nil, false
	}
	return o.FxSummaries, true
}

// HasFxSummaries returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasFxSummaries() bool {
	if o != nil && !IsNil(o.FxSummaries) {
		return true
	}

	return false
}

// SetFxSummaries gets a reference to the given []FxSummary and assigns it to the FxSummaries field.
func (o *PayoutSummaryAudit) SetFxSummaries(v []FxSummary) {
	o.FxSummaries = v
}

// GetPayoutMemo returns the PayoutMemo field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetPayoutMemo() string {
	if o == nil || IsNil(o.PayoutMemo) {
		var ret string
		return ret
	}
	return *o.PayoutMemo
}

// GetPayoutMemoOk returns a tuple with the PayoutMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetPayoutMemoOk() (*string, bool) {
	if o == nil || IsNil(o.PayoutMemo) {
		return nil, false
	}
	return o.PayoutMemo, true
}

// HasPayoutMemo returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasPayoutMemo() bool {
	if o != nil && !IsNil(o.PayoutMemo) {
		return true
	}

	return false
}

// SetPayoutMemo gets a reference to the given string and assigns it to the PayoutMemo field.
func (o *PayoutSummaryAudit) SetPayoutMemo(v string) {
	o.PayoutMemo = &v
}

// GetPayoutType returns the PayoutType field value
func (o *PayoutSummaryAudit) GetPayoutType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayoutType
}

// GetPayoutTypeOk returns a tuple with the PayoutType field value
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetPayoutTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayoutType, true
}

// SetPayoutType sets field value
func (o *PayoutSummaryAudit) SetPayoutType(v string) {
	o.PayoutType = v
}

// GetPayorName returns the PayorName field value
func (o *PayoutSummaryAudit) GetPayorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayorName
}

// GetPayorNameOk returns a tuple with the PayorName field value
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetPayorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayorName, true
}

// SetPayorName sets field value
func (o *PayoutSummaryAudit) SetPayorName(v string) {
	o.PayorName = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *PayoutSummaryAudit) GetSchedule() PayoutSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret PayoutSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSummaryAudit) GetScheduleOk() (*PayoutSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *PayoutSummaryAudit) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given PayoutSchedule and assigns it to the Schedule field.
func (o *PayoutSummaryAudit) SetSchedule(v PayoutSchedule) {
	o.Schedule = &v
}

func (o PayoutSummaryAudit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutSummaryAudit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayoutId) {
		toSerialize["payoutId"] = o.PayoutId
	}
	if !IsNil(o.PayorId) {
		toSerialize["payorId"] = o.PayorId
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.DateTime) {
		toSerialize["dateTime"] = o.DateTime
	}
	toSerialize["submittedDateTime"] = o.SubmittedDateTime
	if !IsNil(o.InstructedDateTime) {
		toSerialize["instructedDateTime"] = o.InstructedDateTime
	}
	if !IsNil(o.WithdrawnDateTime) {
		toSerialize["withdrawnDateTime"] = o.WithdrawnDateTime
	}
	if !IsNil(o.TotalPayments) {
		toSerialize["totalPayments"] = o.TotalPayments
	}
	if !IsNil(o.TotalIncompletePayments) {
		toSerialize["totalIncompletePayments"] = o.TotalIncompletePayments
	}
	if !IsNil(o.TotalReturnedPayments) {
		toSerialize["totalReturnedPayments"] = o.TotalReturnedPayments
	}
	if !IsNil(o.TotalWithdrawnPayments) {
		toSerialize["totalWithdrawnPayments"] = o.TotalWithdrawnPayments
	}
	if !IsNil(o.SourceAccountSummary) {
		toSerialize["sourceAccountSummary"] = o.SourceAccountSummary
	}
	if !IsNil(o.FxSummaries) {
		toSerialize["fxSummaries"] = o.FxSummaries
	}
	if !IsNil(o.PayoutMemo) {
		toSerialize["payoutMemo"] = o.PayoutMemo
	}
	toSerialize["payoutType"] = o.PayoutType
	toSerialize["payorName"] = o.PayorName
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	return toSerialize, nil
}

type NullablePayoutSummaryAudit struct {
	value *PayoutSummaryAudit
	isSet bool
}

func (v NullablePayoutSummaryAudit) Get() *PayoutSummaryAudit {
	return v.value
}

func (v *NullablePayoutSummaryAudit) Set(val *PayoutSummaryAudit) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutSummaryAudit) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutSummaryAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutSummaryAudit(val *PayoutSummaryAudit) *NullablePayoutSummaryAudit {
	return &NullablePayoutSummaryAudit{value: val, isSet: true}
}

func (v NullablePayoutSummaryAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutSummaryAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


