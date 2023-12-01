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

// checks if the FundingEvent2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FundingEvent2{}

// FundingEvent2 struct for FundingEvent2
type FundingEvent2 struct {
	EventId *string `json:"eventId,omitempty"`
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`
	// Funding event type. One of the following values: PAYOR_FUNDING_DETECTED, PAYOR_FUNDING_REQUESTED, PAYOR_FUNDING_RETURN_RECEIVED, FUNDING_RETURN_DETECTED, PAYOR_FUNDING_REQUEST_SUBMITTED, PAYOR_FUNDING_ENTRY_DETAIL_RECEIVED, FUNDING_DEALLOCATED
	FundingEventType *string `json:"fundingEventType,omitempty"`
	Principal *string `json:"principal,omitempty"`
}

// NewFundingEvent2 instantiates a new FundingEvent2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundingEvent2() *FundingEvent2 {
	this := FundingEvent2{}
	return &this
}

// NewFundingEvent2WithDefaults instantiates a new FundingEvent2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingEvent2WithDefaults() *FundingEvent2 {
	this := FundingEvent2{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *FundingEvent2) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingEvent2) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *FundingEvent2) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *FundingEvent2) SetEventId(v string) {
	o.EventId = &v
}

// GetEventDateTime returns the EventDateTime field value if set, zero value otherwise.
func (o *FundingEvent2) GetEventDateTime() time.Time {
	if o == nil || IsNil(o.EventDateTime) {
		var ret time.Time
		return ret
	}
	return *o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingEvent2) GetEventDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventDateTime) {
		return nil, false
	}
	return o.EventDateTime, true
}

// HasEventDateTime returns a boolean if a field has been set.
func (o *FundingEvent2) HasEventDateTime() bool {
	if o != nil && !IsNil(o.EventDateTime) {
		return true
	}

	return false
}

// SetEventDateTime gets a reference to the given time.Time and assigns it to the EventDateTime field.
func (o *FundingEvent2) SetEventDateTime(v time.Time) {
	o.EventDateTime = &v
}

// GetFundingEventType returns the FundingEventType field value if set, zero value otherwise.
func (o *FundingEvent2) GetFundingEventType() string {
	if o == nil || IsNil(o.FundingEventType) {
		var ret string
		return ret
	}
	return *o.FundingEventType
}

// GetFundingEventTypeOk returns a tuple with the FundingEventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingEvent2) GetFundingEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FundingEventType) {
		return nil, false
	}
	return o.FundingEventType, true
}

// HasFundingEventType returns a boolean if a field has been set.
func (o *FundingEvent2) HasFundingEventType() bool {
	if o != nil && !IsNil(o.FundingEventType) {
		return true
	}

	return false
}

// SetFundingEventType gets a reference to the given string and assigns it to the FundingEventType field.
func (o *FundingEvent2) SetFundingEventType(v string) {
	o.FundingEventType = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *FundingEvent2) GetPrincipal() string {
	if o == nil || IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingEvent2) GetPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *FundingEvent2) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *FundingEvent2) SetPrincipal(v string) {
	o.Principal = &v
}

func (o FundingEvent2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundingEvent2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !IsNil(o.EventDateTime) {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	if !IsNil(o.FundingEventType) {
		toSerialize["fundingEventType"] = o.FundingEventType
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	return toSerialize, nil
}

type NullableFundingEvent2 struct {
	value *FundingEvent2
	isSet bool
}

func (v NullableFundingEvent2) Get() *FundingEvent2 {
	return v.value
}

func (v *NullableFundingEvent2) Set(val *FundingEvent2) {
	v.value = val
	v.isSet = true
}

func (v NullableFundingEvent2) IsSet() bool {
	return v.isSet
}

func (v *NullableFundingEvent2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundingEvent2(val *FundingEvent2) *NullableFundingEvent2 {
	return &NullableFundingEvent2{value: val, isSet: true}
}

func (v NullableFundingEvent2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundingEvent2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


