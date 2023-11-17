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

// checks if the PagedPayeeResponseV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedPayeeResponseV4{}

// PagedPayeeResponseV4 List Payees Response Object
type PagedPayeeResponseV4 struct {
	Summary *PagedPayeeResponseV3Summary `json:"summary,omitempty"`
	Page *PagedPayeeResponseV3Page `json:"page,omitempty"`
	Links []PagedPayeeResponseV3Links `json:"links,omitempty"`
	Content []GetPayeeListResponseV4 `json:"content,omitempty"`
}

// NewPagedPayeeResponseV4 instantiates a new PagedPayeeResponseV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedPayeeResponseV4() *PagedPayeeResponseV4 {
	this := PagedPayeeResponseV4{}
	return &this
}

// NewPagedPayeeResponseV4WithDefaults instantiates a new PagedPayeeResponseV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedPayeeResponseV4WithDefaults() *PagedPayeeResponseV4 {
	this := PagedPayeeResponseV4{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PagedPayeeResponseV4) GetSummary() PagedPayeeResponseV3Summary {
	if o == nil || IsNil(o.Summary) {
		var ret PagedPayeeResponseV3Summary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponseV4) GetSummaryOk() (*PagedPayeeResponseV3Summary, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PagedPayeeResponseV4) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given PagedPayeeResponseV3Summary and assigns it to the Summary field.
func (o *PagedPayeeResponseV4) SetSummary(v PagedPayeeResponseV3Summary) {
	o.Summary = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PagedPayeeResponseV4) GetPage() PagedPayeeResponseV3Page {
	if o == nil || IsNil(o.Page) {
		var ret PagedPayeeResponseV3Page
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponseV4) GetPageOk() (*PagedPayeeResponseV3Page, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PagedPayeeResponseV4) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given PagedPayeeResponseV3Page and assigns it to the Page field.
func (o *PagedPayeeResponseV4) SetPage(v PagedPayeeResponseV3Page) {
	o.Page = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PagedPayeeResponseV4) GetLinks() []PagedPayeeResponseV3Links {
	if o == nil || IsNil(o.Links) {
		var ret []PagedPayeeResponseV3Links
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponseV4) GetLinksOk() ([]PagedPayeeResponseV3Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PagedPayeeResponseV4) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []PagedPayeeResponseV3Links and assigns it to the Links field.
func (o *PagedPayeeResponseV4) SetLinks(v []PagedPayeeResponseV3Links) {
	o.Links = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PagedPayeeResponseV4) GetContent() []GetPayeeListResponseV4 {
	if o == nil || IsNil(o.Content) {
		var ret []GetPayeeListResponseV4
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponseV4) GetContentOk() ([]GetPayeeListResponseV4, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PagedPayeeResponseV4) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []GetPayeeListResponseV4 and assigns it to the Content field.
func (o *PagedPayeeResponseV4) SetContent(v []GetPayeeListResponseV4) {
	o.Content = v
}

func (o PagedPayeeResponseV4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedPayeeResponseV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullablePagedPayeeResponseV4 struct {
	value *PagedPayeeResponseV4
	isSet bool
}

func (v NullablePagedPayeeResponseV4) Get() *PagedPayeeResponseV4 {
	return v.value
}

func (v *NullablePagedPayeeResponseV4) Set(val *PagedPayeeResponseV4) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedPayeeResponseV4) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedPayeeResponseV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedPayeeResponseV4(val *PagedPayeeResponseV4) *NullablePagedPayeeResponseV4 {
	return &NullablePagedPayeeResponseV4{value: val, isSet: true}
}

func (v NullablePagedPayeeResponseV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedPayeeResponseV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

