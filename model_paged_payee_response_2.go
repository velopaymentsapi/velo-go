/*
 * Velo Payments APIs
 *
 * ## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 
 *
 * API version: 2.26.127
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// PagedPayeeResponse2 List Payees Response Object
type PagedPayeeResponse2 struct {
	Summary *PagedPayeeResponseSummary `json:"summary,omitempty"`
	Page *PagedPayeeResponsePage `json:"page,omitempty"`
	Links *[]PagedPayeeResponseLinks `json:"links,omitempty"`
	Content *[]GetPayeeListResponse2 `json:"content,omitempty"`
}

// NewPagedPayeeResponse2 instantiates a new PagedPayeeResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedPayeeResponse2() *PagedPayeeResponse2 {
	this := PagedPayeeResponse2{}
	return &this
}

// NewPagedPayeeResponse2WithDefaults instantiates a new PagedPayeeResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedPayeeResponse2WithDefaults() *PagedPayeeResponse2 {
	this := PagedPayeeResponse2{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PagedPayeeResponse2) GetSummary() PagedPayeeResponseSummary {
	if o == nil || o.Summary == nil {
		var ret PagedPayeeResponseSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponse2) GetSummaryOk() (*PagedPayeeResponseSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PagedPayeeResponse2) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given PagedPayeeResponseSummary and assigns it to the Summary field.
func (o *PagedPayeeResponse2) SetSummary(v PagedPayeeResponseSummary) {
	o.Summary = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PagedPayeeResponse2) GetPage() PagedPayeeResponsePage {
	if o == nil || o.Page == nil {
		var ret PagedPayeeResponsePage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponse2) GetPageOk() (*PagedPayeeResponsePage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PagedPayeeResponse2) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given PagedPayeeResponsePage and assigns it to the Page field.
func (o *PagedPayeeResponse2) SetPage(v PagedPayeeResponsePage) {
	o.Page = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PagedPayeeResponse2) GetLinks() []PagedPayeeResponseLinks {
	if o == nil || o.Links == nil {
		var ret []PagedPayeeResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponse2) GetLinksOk() (*[]PagedPayeeResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PagedPayeeResponse2) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []PagedPayeeResponseLinks and assigns it to the Links field.
func (o *PagedPayeeResponse2) SetLinks(v []PagedPayeeResponseLinks) {
	o.Links = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PagedPayeeResponse2) GetContent() []GetPayeeListResponse2 {
	if o == nil || o.Content == nil {
		var ret []GetPayeeListResponse2
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPayeeResponse2) GetContentOk() (*[]GetPayeeListResponse2, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PagedPayeeResponse2) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []GetPayeeListResponse2 and assigns it to the Content field.
func (o *PagedPayeeResponse2) SetContent(v []GetPayeeListResponse2) {
	o.Content = &v
}

func (o PagedPayeeResponse2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullablePagedPayeeResponse2 struct {
	value *PagedPayeeResponse2
	isSet bool
}

func (v NullablePagedPayeeResponse2) Get() *PagedPayeeResponse2 {
	return v.value
}

func (v *NullablePagedPayeeResponse2) Set(val *PagedPayeeResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedPayeeResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedPayeeResponse2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedPayeeResponse2(val *PagedPayeeResponse2) *NullablePagedPayeeResponse2 {
	return &NullablePagedPayeeResponse2{value: val, isSet: true}
}

func (v NullablePagedPayeeResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedPayeeResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


