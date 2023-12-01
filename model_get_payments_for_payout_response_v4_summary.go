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

// checks if the GetPaymentsForPayoutResponseV4Summary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPaymentsForPayoutResponseV4Summary{}

// GetPaymentsForPayoutResponseV4Summary struct for GetPaymentsForPayoutResponseV4Summary
type GetPaymentsForPayoutResponseV4Summary struct {
	// Current status of the Payout. One of the following values: ACCEPTED, REJECTED, SUBMITTED, QUOTED, INSTRUCTED, COMPLETED, INCOMPLETE, CONFIRMED, WITHDRAWN
	PayoutStatus *string `json:"payoutStatus,omitempty"`
	// The date/time at which the payout was submitted.
	SubmittedDateTime *time.Time `json:"submittedDateTime,omitempty"`
	// The date/time at which the payout was instructed.
	InstructedDateTime *time.Time `json:"instructedDateTime,omitempty"`
	WithdrawnDateTime *time.Time `json:"withdrawnDateTime,omitempty"`
	// The date/time at which the payout was quoted.
	QuotedDateTime *time.Time `json:"quotedDateTime,omitempty"`
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
	// The count of payments within the payout which have been returned.
	ReturnedPayments *int32 `json:"returnedPayments,omitempty"`
	// The count of payments within the payout which have been withdrawn.
	WithdrawnPayments *int32 `json:"withdrawnPayments,omitempty"`
	// The type of payout. One of the following values: STANDARD, AS, ON_BEHALF_OF
	PayoutType *string `json:"payoutType,omitempty"`
	Submitting *PayoutPayor `json:"submitting,omitempty"`
	PayoutFrom *PayoutPayor `json:"payoutFrom,omitempty"`
	PayoutTo *PayoutPayor `json:"payoutTo,omitempty"`
	Quoted *PayoutPrincipal `json:"quoted,omitempty"`
	Instructed *PayoutPrincipal `json:"instructed,omitempty"`
	Withdrawn *PayoutPrincipal `json:"withdrawn,omitempty"`
	Schedule *PayoutSchedule `json:"schedule,omitempty"`
}

// NewGetPaymentsForPayoutResponseV4Summary instantiates a new GetPaymentsForPayoutResponseV4Summary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPaymentsForPayoutResponseV4Summary() *GetPaymentsForPayoutResponseV4Summary {
	this := GetPaymentsForPayoutResponseV4Summary{}
	return &this
}

// NewGetPaymentsForPayoutResponseV4SummaryWithDefaults instantiates a new GetPaymentsForPayoutResponseV4Summary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPaymentsForPayoutResponseV4SummaryWithDefaults() *GetPaymentsForPayoutResponseV4Summary {
	this := GetPaymentsForPayoutResponseV4Summary{}
	return &this
}

// GetPayoutStatus returns the PayoutStatus field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutStatus() string {
	if o == nil || IsNil(o.PayoutStatus) {
		var ret string
		return ret
	}
	return *o.PayoutStatus
}

// GetPayoutStatusOk returns a tuple with the PayoutStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PayoutStatus) {
		return nil, false
	}
	return o.PayoutStatus, true
}

// HasPayoutStatus returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutStatus() bool {
	if o != nil && !IsNil(o.PayoutStatus) {
		return true
	}

	return false
}

// SetPayoutStatus gets a reference to the given string and assigns it to the PayoutStatus field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutStatus(v string) {
	o.PayoutStatus = &v
}

// GetSubmittedDateTime returns the SubmittedDateTime field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetSubmittedDateTime() time.Time {
	if o == nil || IsNil(o.SubmittedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedDateTime
}

// GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetSubmittedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedDateTime) {
		return nil, false
	}
	return o.SubmittedDateTime, true
}

// HasSubmittedDateTime returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasSubmittedDateTime() bool {
	if o != nil && !IsNil(o.SubmittedDateTime) {
		return true
	}

	return false
}

// SetSubmittedDateTime gets a reference to the given time.Time and assigns it to the SubmittedDateTime field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetSubmittedDateTime(v time.Time) {
	o.SubmittedDateTime = &v
}

// GetInstructedDateTime returns the InstructedDateTime field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetInstructedDateTime() time.Time {
	if o == nil || IsNil(o.InstructedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.InstructedDateTime
}

// GetInstructedDateTimeOk returns a tuple with the InstructedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetInstructedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InstructedDateTime) {
		return nil, false
	}
	return o.InstructedDateTime, true
}

// HasInstructedDateTime returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasInstructedDateTime() bool {
	if o != nil && !IsNil(o.InstructedDateTime) {
		return true
	}

	return false
}

// SetInstructedDateTime gets a reference to the given time.Time and assigns it to the InstructedDateTime field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetInstructedDateTime(v time.Time) {
	o.InstructedDateTime = &v
}

// GetWithdrawnDateTime returns the WithdrawnDateTime field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnDateTime() time.Time {
	if o == nil || IsNil(o.WithdrawnDateTime) {
		var ret time.Time
		return ret
	}
	return *o.WithdrawnDateTime
}

// GetWithdrawnDateTimeOk returns a tuple with the WithdrawnDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.WithdrawnDateTime) {
		return nil, false
	}
	return o.WithdrawnDateTime, true
}

// HasWithdrawnDateTime returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasWithdrawnDateTime() bool {
	if o != nil && !IsNil(o.WithdrawnDateTime) {
		return true
	}

	return false
}

// SetWithdrawnDateTime gets a reference to the given time.Time and assigns it to the WithdrawnDateTime field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetWithdrawnDateTime(v time.Time) {
	o.WithdrawnDateTime = &v
}

// GetQuotedDateTime returns the QuotedDateTime field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetQuotedDateTime() time.Time {
	if o == nil || IsNil(o.QuotedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.QuotedDateTime
}

// GetQuotedDateTimeOk returns a tuple with the QuotedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetQuotedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.QuotedDateTime) {
		return nil, false
	}
	return o.QuotedDateTime, true
}

// HasQuotedDateTime returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasQuotedDateTime() bool {
	if o != nil && !IsNil(o.QuotedDateTime) {
		return true
	}

	return false
}

// SetQuotedDateTime gets a reference to the given time.Time and assigns it to the QuotedDateTime field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetQuotedDateTime(v time.Time) {
	o.QuotedDateTime = &v
}

// GetPayoutMemo returns the PayoutMemo field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutMemo() string {
	if o == nil || IsNil(o.PayoutMemo) {
		var ret string
		return ret
	}
	return *o.PayoutMemo
}

// GetPayoutMemoOk returns a tuple with the PayoutMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutMemoOk() (*string, bool) {
	if o == nil || IsNil(o.PayoutMemo) {
		return nil, false
	}
	return o.PayoutMemo, true
}

// HasPayoutMemo returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutMemo() bool {
	if o != nil && !IsNil(o.PayoutMemo) {
		return true
	}

	return false
}

// SetPayoutMemo gets a reference to the given string and assigns it to the PayoutMemo field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutMemo(v string) {
	o.PayoutMemo = &v
}

// GetTotalPayments returns the TotalPayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetTotalPayments() int32 {
	if o == nil || IsNil(o.TotalPayments) {
		var ret int32
		return ret
	}
	return *o.TotalPayments
}

// GetTotalPaymentsOk returns a tuple with the TotalPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetTotalPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPayments) {
		return nil, false
	}
	return o.TotalPayments, true
}

// HasTotalPayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasTotalPayments() bool {
	if o != nil && !IsNil(o.TotalPayments) {
		return true
	}

	return false
}

// SetTotalPayments gets a reference to the given int32 and assigns it to the TotalPayments field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetTotalPayments(v int32) {
	o.TotalPayments = &v
}

// GetConfirmedPayments returns the ConfirmedPayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetConfirmedPayments() int32 {
	if o == nil || IsNil(o.ConfirmedPayments) {
		var ret int32
		return ret
	}
	return *o.ConfirmedPayments
}

// GetConfirmedPaymentsOk returns a tuple with the ConfirmedPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetConfirmedPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfirmedPayments) {
		return nil, false
	}
	return o.ConfirmedPayments, true
}

// HasConfirmedPayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasConfirmedPayments() bool {
	if o != nil && !IsNil(o.ConfirmedPayments) {
		return true
	}

	return false
}

// SetConfirmedPayments gets a reference to the given int32 and assigns it to the ConfirmedPayments field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetConfirmedPayments(v int32) {
	o.ConfirmedPayments = &v
}

// GetReleasedPayments returns the ReleasedPayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetReleasedPayments() int32 {
	if o == nil || IsNil(o.ReleasedPayments) {
		var ret int32
		return ret
	}
	return *o.ReleasedPayments
}

// GetReleasedPaymentsOk returns a tuple with the ReleasedPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetReleasedPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.ReleasedPayments) {
		return nil, false
	}
	return o.ReleasedPayments, true
}

// HasReleasedPayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasReleasedPayments() bool {
	if o != nil && !IsNil(o.ReleasedPayments) {
		return true
	}

	return false
}

// SetReleasedPayments gets a reference to the given int32 and assigns it to the ReleasedPayments field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetReleasedPayments(v int32) {
	o.ReleasedPayments = &v
}

// GetIncompletePayments returns the IncompletePayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetIncompletePayments() int32 {
	if o == nil || IsNil(o.IncompletePayments) {
		var ret int32
		return ret
	}
	return *o.IncompletePayments
}

// GetIncompletePaymentsOk returns a tuple with the IncompletePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetIncompletePaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.IncompletePayments) {
		return nil, false
	}
	return o.IncompletePayments, true
}

// HasIncompletePayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasIncompletePayments() bool {
	if o != nil && !IsNil(o.IncompletePayments) {
		return true
	}

	return false
}

// SetIncompletePayments gets a reference to the given int32 and assigns it to the IncompletePayments field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetIncompletePayments(v int32) {
	o.IncompletePayments = &v
}

// GetReturnedPayments returns the ReturnedPayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetReturnedPayments() int32 {
	if o == nil || IsNil(o.ReturnedPayments) {
		var ret int32
		return ret
	}
	return *o.ReturnedPayments
}

// GetReturnedPaymentsOk returns a tuple with the ReturnedPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetReturnedPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnedPayments) {
		return nil, false
	}
	return o.ReturnedPayments, true
}

// HasReturnedPayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasReturnedPayments() bool {
	if o != nil && !IsNil(o.ReturnedPayments) {
		return true
	}

	return false
}

// SetReturnedPayments gets a reference to the given int32 and assigns it to the ReturnedPayments field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetReturnedPayments(v int32) {
	o.ReturnedPayments = &v
}

// GetWithdrawnPayments returns the WithdrawnPayments field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnPayments() int32 {
	if o == nil || IsNil(o.WithdrawnPayments) {
		var ret int32
		return ret
	}
	return *o.WithdrawnPayments
}

// GetWithdrawnPaymentsOk returns a tuple with the WithdrawnPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnPaymentsOk() (*int32, bool) {
	if o == nil || IsNil(o.WithdrawnPayments) {
		return nil, false
	}
	return o.WithdrawnPayments, true
}

// HasWithdrawnPayments returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasWithdrawnPayments() bool {
	if o != nil && !IsNil(o.WithdrawnPayments) {
		return true
	}

	return false
}

// SetWithdrawnPayments gets a reference to the given int32 and assigns it to the WithdrawnPayments field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetWithdrawnPayments(v int32) {
	o.WithdrawnPayments = &v
}

// GetPayoutType returns the PayoutType field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutType() string {
	if o == nil || IsNil(o.PayoutType) {
		var ret string
		return ret
	}
	return *o.PayoutType
}

// GetPayoutTypeOk returns a tuple with the PayoutType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PayoutType) {
		return nil, false
	}
	return o.PayoutType, true
}

// HasPayoutType returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutType() bool {
	if o != nil && !IsNil(o.PayoutType) {
		return true
	}

	return false
}

// SetPayoutType gets a reference to the given string and assigns it to the PayoutType field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutType(v string) {
	o.PayoutType = &v
}

// GetSubmitting returns the Submitting field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetSubmitting() PayoutPayor {
	if o == nil || IsNil(o.Submitting) {
		var ret PayoutPayor
		return ret
	}
	return *o.Submitting
}

// GetSubmittingOk returns a tuple with the Submitting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetSubmittingOk() (*PayoutPayor, bool) {
	if o == nil || IsNil(o.Submitting) {
		return nil, false
	}
	return o.Submitting, true
}

// HasSubmitting returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasSubmitting() bool {
	if o != nil && !IsNil(o.Submitting) {
		return true
	}

	return false
}

// SetSubmitting gets a reference to the given PayoutPayor and assigns it to the Submitting field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetSubmitting(v PayoutPayor) {
	o.Submitting = &v
}

// GetPayoutFrom returns the PayoutFrom field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutFrom() PayoutPayor {
	if o == nil || IsNil(o.PayoutFrom) {
		var ret PayoutPayor
		return ret
	}
	return *o.PayoutFrom
}

// GetPayoutFromOk returns a tuple with the PayoutFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutFromOk() (*PayoutPayor, bool) {
	if o == nil || IsNil(o.PayoutFrom) {
		return nil, false
	}
	return o.PayoutFrom, true
}

// HasPayoutFrom returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutFrom() bool {
	if o != nil && !IsNil(o.PayoutFrom) {
		return true
	}

	return false
}

// SetPayoutFrom gets a reference to the given PayoutPayor and assigns it to the PayoutFrom field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutFrom(v PayoutPayor) {
	o.PayoutFrom = &v
}

// GetPayoutTo returns the PayoutTo field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutTo() PayoutPayor {
	if o == nil || IsNil(o.PayoutTo) {
		var ret PayoutPayor
		return ret
	}
	return *o.PayoutTo
}

// GetPayoutToOk returns a tuple with the PayoutTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutToOk() (*PayoutPayor, bool) {
	if o == nil || IsNil(o.PayoutTo) {
		return nil, false
	}
	return o.PayoutTo, true
}

// HasPayoutTo returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutTo() bool {
	if o != nil && !IsNil(o.PayoutTo) {
		return true
	}

	return false
}

// SetPayoutTo gets a reference to the given PayoutPayor and assigns it to the PayoutTo field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutTo(v PayoutPayor) {
	o.PayoutTo = &v
}

// GetQuoted returns the Quoted field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetQuoted() PayoutPrincipal {
	if o == nil || IsNil(o.Quoted) {
		var ret PayoutPrincipal
		return ret
	}
	return *o.Quoted
}

// GetQuotedOk returns a tuple with the Quoted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetQuotedOk() (*PayoutPrincipal, bool) {
	if o == nil || IsNil(o.Quoted) {
		return nil, false
	}
	return o.Quoted, true
}

// HasQuoted returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasQuoted() bool {
	if o != nil && !IsNil(o.Quoted) {
		return true
	}

	return false
}

// SetQuoted gets a reference to the given PayoutPrincipal and assigns it to the Quoted field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetQuoted(v PayoutPrincipal) {
	o.Quoted = &v
}

// GetInstructed returns the Instructed field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetInstructed() PayoutPrincipal {
	if o == nil || IsNil(o.Instructed) {
		var ret PayoutPrincipal
		return ret
	}
	return *o.Instructed
}

// GetInstructedOk returns a tuple with the Instructed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetInstructedOk() (*PayoutPrincipal, bool) {
	if o == nil || IsNil(o.Instructed) {
		return nil, false
	}
	return o.Instructed, true
}

// HasInstructed returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasInstructed() bool {
	if o != nil && !IsNil(o.Instructed) {
		return true
	}

	return false
}

// SetInstructed gets a reference to the given PayoutPrincipal and assigns it to the Instructed field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetInstructed(v PayoutPrincipal) {
	o.Instructed = &v
}

// GetWithdrawn returns the Withdrawn field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawn() PayoutPrincipal {
	if o == nil || IsNil(o.Withdrawn) {
		var ret PayoutPrincipal
		return ret
	}
	return *o.Withdrawn
}

// GetWithdrawnOk returns a tuple with the Withdrawn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnOk() (*PayoutPrincipal, bool) {
	if o == nil || IsNil(o.Withdrawn) {
		return nil, false
	}
	return o.Withdrawn, true
}

// HasWithdrawn returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasWithdrawn() bool {
	if o != nil && !IsNil(o.Withdrawn) {
		return true
	}

	return false
}

// SetWithdrawn gets a reference to the given PayoutPrincipal and assigns it to the Withdrawn field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetWithdrawn(v PayoutPrincipal) {
	o.Withdrawn = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *GetPaymentsForPayoutResponseV4Summary) GetSchedule() PayoutSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret PayoutSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) GetScheduleOk() (*PayoutSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *GetPaymentsForPayoutResponseV4Summary) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given PayoutSchedule and assigns it to the Schedule field.
func (o *GetPaymentsForPayoutResponseV4Summary) SetSchedule(v PayoutSchedule) {
	o.Schedule = &v
}

func (o GetPaymentsForPayoutResponseV4Summary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPaymentsForPayoutResponseV4Summary) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.QuotedDateTime) {
		toSerialize["quotedDateTime"] = o.QuotedDateTime
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
	if !IsNil(o.ReturnedPayments) {
		toSerialize["returnedPayments"] = o.ReturnedPayments
	}
	if !IsNil(o.WithdrawnPayments) {
		toSerialize["withdrawnPayments"] = o.WithdrawnPayments
	}
	if !IsNil(o.PayoutType) {
		toSerialize["payoutType"] = o.PayoutType
	}
	if !IsNil(o.Submitting) {
		toSerialize["submitting"] = o.Submitting
	}
	if !IsNil(o.PayoutFrom) {
		toSerialize["payoutFrom"] = o.PayoutFrom
	}
	if !IsNil(o.PayoutTo) {
		toSerialize["payoutTo"] = o.PayoutTo
	}
	if !IsNil(o.Quoted) {
		toSerialize["quoted"] = o.Quoted
	}
	if !IsNil(o.Instructed) {
		toSerialize["instructed"] = o.Instructed
	}
	if !IsNil(o.Withdrawn) {
		toSerialize["withdrawn"] = o.Withdrawn
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	return toSerialize, nil
}

type NullableGetPaymentsForPayoutResponseV4Summary struct {
	value *GetPaymentsForPayoutResponseV4Summary
	isSet bool
}

func (v NullableGetPaymentsForPayoutResponseV4Summary) Get() *GetPaymentsForPayoutResponseV4Summary {
	return v.value
}

func (v *NullableGetPaymentsForPayoutResponseV4Summary) Set(val *GetPaymentsForPayoutResponseV4Summary) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPaymentsForPayoutResponseV4Summary) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPaymentsForPayoutResponseV4Summary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPaymentsForPayoutResponseV4Summary(val *GetPaymentsForPayoutResponseV4Summary) *NullableGetPaymentsForPayoutResponseV4Summary {
	return &NullableGetPaymentsForPayoutResponseV4Summary{value: val, isSet: true}
}

func (v NullableGetPaymentsForPayoutResponseV4Summary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPaymentsForPayoutResponseV4Summary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


