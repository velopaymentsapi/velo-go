/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 

API version: 2.35.58
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// checks if the PagedPayeeResponseV3Summary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedPayeeResponseV3Summary{}

// PagedPayeeResponseV3Summary struct for PagedPayeeResponseV3Summary
type PagedPayeeResponseV3Summary struct {
	TotalPayeesCount *int32 `json:"totalPayeesCount,omitempty"`
	TotalInvitedCount *int32 `json:"totalInvitedCount,omitempty"`
	TotalRegisteredCount *int32 `json:"totalRegisteredCount,omitempty"`
	TotalOnboardedCount *int32 `json:"totalOnboardedCount,omitempty"`
	TotalWatchlistFailedCount *int32 `json:"totalWatchlistFailedCount,omitempty"`
}

// NewPagedPayeeResponseV3Summary instantiates a new PagedPayeeResponseV3Summary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedPayeeResponseV3Summary() *PagedPayeeResponseV3Summary {
	this := PagedPayeeResponseV3Summary{}
	return &this
}

// NewPagedPayeeResponseV3SummaryWithDefaults instantiates a new PagedPayeeResponseV3Summary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedPayeeResponseV3SummaryWithDefaults() *PagedPayeeResponseV3Summary {
	this := PagedPayeeResponseV3Summary{}
	return &this
}

// GetTotalPayeesCount returns the TotalPayeesCount field value if set, zero value otherwise.
func (o *PagedPayeeResponseV3Summary) GetTotalPayeesCount() int32 {
	if o == nil || IsNil(o.TotalPayeesCount) {
		var ret int32
		return ret
	}
	return *o.TotalPayeesCount
}

// GetTotalPayeesCountOk returns a tuple with the TotalPayeesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponseV3Summary) GetTotalPayeesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPayeesCount) {
		return nil, false
	}
	return o.TotalPayeesCount, true
}

// HasTotalPayeesCount returns a boolean if a field has been set.
func (o *PagedPayeeResponseV3Summary) HasTotalPayeesCount() bool {
	if o != nil && !IsNil(o.TotalPayeesCount) {
		return true
	}

	return false
}

// SetTotalPayeesCount gets a reference to the given int32 and assigns it to the TotalPayeesCount field.
func (o *PagedPayeeResponseV3Summary) SetTotalPayeesCount(v int32) {
	o.TotalPayeesCount = &v
}

// GetTotalInvitedCount returns the TotalInvitedCount field value if set, zero value otherwise.
func (o *PagedPayeeResponseV3Summary) GetTotalInvitedCount() int32 {
	if o == nil || IsNil(o.TotalInvitedCount) {
		var ret int32
		return ret
	}
	return *o.TotalInvitedCount
}

// GetTotalInvitedCountOk returns a tuple with the TotalInvitedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponseV3Summary) GetTotalInvitedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalInvitedCount) {
		return nil, false
	}
	return o.TotalInvitedCount, true
}

// HasTotalInvitedCount returns a boolean if a field has been set.
func (o *PagedPayeeResponseV3Summary) HasTotalInvitedCount() bool {
	if o != nil && !IsNil(o.TotalInvitedCount) {
		return true
	}

	return false
}

// SetTotalInvitedCount gets a reference to the given int32 and assigns it to the TotalInvitedCount field.
func (o *PagedPayeeResponseV3Summary) SetTotalInvitedCount(v int32) {
	o.TotalInvitedCount = &v
}

// GetTotalRegisteredCount returns the TotalRegisteredCount field value if set, zero value otherwise.
func (o *PagedPayeeResponseV3Summary) GetTotalRegisteredCount() int32 {
	if o == nil || IsNil(o.TotalRegisteredCount) {
		var ret int32
		return ret
	}
	return *o.TotalRegisteredCount
}

// GetTotalRegisteredCountOk returns a tuple with the TotalRegisteredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponseV3Summary) GetTotalRegisteredCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRegisteredCount) {
		return nil, false
	}
	return o.TotalRegisteredCount, true
}

// HasTotalRegisteredCount returns a boolean if a field has been set.
func (o *PagedPayeeResponseV3Summary) HasTotalRegisteredCount() bool {
	if o != nil && !IsNil(o.TotalRegisteredCount) {
		return true
	}

	return false
}

// SetTotalRegisteredCount gets a reference to the given int32 and assigns it to the TotalRegisteredCount field.
func (o *PagedPayeeResponseV3Summary) SetTotalRegisteredCount(v int32) {
	o.TotalRegisteredCount = &v
}

// GetTotalOnboardedCount returns the TotalOnboardedCount field value if set, zero value otherwise.
func (o *PagedPayeeResponseV3Summary) GetTotalOnboardedCount() int32 {
	if o == nil || IsNil(o.TotalOnboardedCount) {
		var ret int32
		return ret
	}
	return *o.TotalOnboardedCount
}

// GetTotalOnboardedCountOk returns a tuple with the TotalOnboardedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponseV3Summary) GetTotalOnboardedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalOnboardedCount) {
		return nil, false
	}
	return o.TotalOnboardedCount, true
}

// HasTotalOnboardedCount returns a boolean if a field has been set.
func (o *PagedPayeeResponseV3Summary) HasTotalOnboardedCount() bool {
	if o != nil && !IsNil(o.TotalOnboardedCount) {
		return true
	}

	return false
}

// SetTotalOnboardedCount gets a reference to the given int32 and assigns it to the TotalOnboardedCount field.
func (o *PagedPayeeResponseV3Summary) SetTotalOnboardedCount(v int32) {
	o.TotalOnboardedCount = &v
}

// GetTotalWatchlistFailedCount returns the TotalWatchlistFailedCount field value if set, zero value otherwise.
func (o *PagedPayeeResponseV3Summary) GetTotalWatchlistFailedCount() int32 {
	if o == nil || IsNil(o.TotalWatchlistFailedCount) {
		var ret int32
		return ret
	}
	return *o.TotalWatchlistFailedCount
}

// GetTotalWatchlistFailedCountOk returns a tuple with the TotalWatchlistFailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponseV3Summary) GetTotalWatchlistFailedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalWatchlistFailedCount) {
		return nil, false
	}
	return o.TotalWatchlistFailedCount, true
}

// HasTotalWatchlistFailedCount returns a boolean if a field has been set.
func (o *PagedPayeeResponseV3Summary) HasTotalWatchlistFailedCount() bool {
	if o != nil && !IsNil(o.TotalWatchlistFailedCount) {
		return true
	}

	return false
}

// SetTotalWatchlistFailedCount gets a reference to the given int32 and assigns it to the TotalWatchlistFailedCount field.
func (o *PagedPayeeResponseV3Summary) SetTotalWatchlistFailedCount(v int32) {
	o.TotalWatchlistFailedCount = &v
}

func (o PagedPayeeResponseV3Summary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedPayeeResponseV3Summary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalPayeesCount) {
		toSerialize["totalPayeesCount"] = o.TotalPayeesCount
	}
	if !IsNil(o.TotalInvitedCount) {
		toSerialize["totalInvitedCount"] = o.TotalInvitedCount
	}
	if !IsNil(o.TotalRegisteredCount) {
		toSerialize["totalRegisteredCount"] = o.TotalRegisteredCount
	}
	if !IsNil(o.TotalOnboardedCount) {
		toSerialize["totalOnboardedCount"] = o.TotalOnboardedCount
	}
	if !IsNil(o.TotalWatchlistFailedCount) {
		toSerialize["totalWatchlistFailedCount"] = o.TotalWatchlistFailedCount
	}
	return toSerialize, nil
}

type NullablePagedPayeeResponseV3Summary struct {
	value *PagedPayeeResponseV3Summary
	isSet bool
}

func (v NullablePagedPayeeResponseV3Summary) Get() *PagedPayeeResponseV3Summary {
	return v.value
}

func (v *NullablePagedPayeeResponseV3Summary) Set(val *PagedPayeeResponseV3Summary) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedPayeeResponseV3Summary) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedPayeeResponseV3Summary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedPayeeResponseV3Summary(val *PagedPayeeResponseV3Summary) *NullablePagedPayeeResponseV3Summary {
	return &NullablePagedPayeeResponseV3Summary{value: val, isSet: true}
}

func (v NullablePagedPayeeResponseV3Summary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedPayeeResponseV3Summary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

