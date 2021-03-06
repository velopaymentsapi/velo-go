/*
 * Velo Payments APIs
 *
 * ## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 
 *
 * API version: 2.26.124
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
	"time"
)

// PaymentEvent Base type for all Payment Events
type PaymentEvent struct {
	// OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly
	SourceType string `json:"sourceType"`
	// UUID id of the source event in the Velo platform
	EventId string `json:"eventId"`
	// ISO8601 timestamp indicating when the source event was created
	CreatedAt time.Time `json:"createdAt"`
	// ID of this payment within the Velo platform
	PaymentId string `json:"paymentId"`
	PayoutPayorIds *PayoutPayorIds `json:"payoutPayorIds,omitempty"`
	// ID of this payment in the payors system
	PayorPaymentId *string `json:"payorPaymentId,omitempty"`
}

// NewPaymentEvent instantiates a new PaymentEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentEvent(sourceType string, eventId string, createdAt time.Time, paymentId string) *PaymentEvent {
	this := PaymentEvent{}
	this.SourceType = sourceType
	this.EventId = eventId
	this.CreatedAt = createdAt
	this.PaymentId = paymentId
	return &this
}

// NewPaymentEventWithDefaults instantiates a new PaymentEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentEventWithDefaults() *PaymentEvent {
	this := PaymentEvent{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *PaymentEvent) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *PaymentEvent) GetSourceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *PaymentEvent) SetSourceType(v string) {
	o.SourceType = v
}

// GetEventId returns the EventId field value
func (o *PaymentEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *PaymentEvent) GetEventIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *PaymentEvent) SetEventId(v string) {
	o.EventId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentEvent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentEvent) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentEvent) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentEvent) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetPayoutPayorIds returns the PayoutPayorIds field value if set, zero value otherwise.
func (o *PaymentEvent) GetPayoutPayorIds() PayoutPayorIds {
	if o == nil || o.PayoutPayorIds == nil {
		var ret PayoutPayorIds
		return ret
	}
	return *o.PayoutPayorIds
}

// GetPayoutPayorIdsOk returns a tuple with the PayoutPayorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentEvent) GetPayoutPayorIdsOk() (*PayoutPayorIds, bool) {
	if o == nil || o.PayoutPayorIds == nil {
		return nil, false
	}
	return o.PayoutPayorIds, true
}

// HasPayoutPayorIds returns a boolean if a field has been set.
func (o *PaymentEvent) HasPayoutPayorIds() bool {
	if o != nil && o.PayoutPayorIds != nil {
		return true
	}

	return false
}

// SetPayoutPayorIds gets a reference to the given PayoutPayorIds and assigns it to the PayoutPayorIds field.
func (o *PaymentEvent) SetPayoutPayorIds(v PayoutPayorIds) {
	o.PayoutPayorIds = &v
}

// GetPayorPaymentId returns the PayorPaymentId field value if set, zero value otherwise.
func (o *PaymentEvent) GetPayorPaymentId() string {
	if o == nil || o.PayorPaymentId == nil {
		var ret string
		return ret
	}
	return *o.PayorPaymentId
}

// GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentEvent) GetPayorPaymentIdOk() (*string, bool) {
	if o == nil || o.PayorPaymentId == nil {
		return nil, false
	}
	return o.PayorPaymentId, true
}

// HasPayorPaymentId returns a boolean if a field has been set.
func (o *PaymentEvent) HasPayorPaymentId() bool {
	if o != nil && o.PayorPaymentId != nil {
		return true
	}

	return false
}

// SetPayorPaymentId gets a reference to the given string and assigns it to the PayorPaymentId field.
func (o *PaymentEvent) SetPayorPaymentId(v string) {
	o.PayorPaymentId = &v
}

func (o PaymentEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceType"] = o.SourceType
	}
	if true {
		toSerialize["eventId"] = o.EventId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["paymentId"] = o.PaymentId
	}
	if o.PayoutPayorIds != nil {
		toSerialize["payoutPayorIds"] = o.PayoutPayorIds
	}
	if o.PayorPaymentId != nil {
		toSerialize["payorPaymentId"] = o.PayorPaymentId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentEvent struct {
	value *PaymentEvent
	isSet bool
}

func (v NullablePaymentEvent) Get() *PaymentEvent {
	return v.value
}

func (v *NullablePaymentEvent) Set(val *PaymentEvent) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentEvent) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentEvent(val *PaymentEvent) *NullablePaymentEvent {
	return &NullablePaymentEvent{value: val, isSet: true}
}

func (v NullablePaymentEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


