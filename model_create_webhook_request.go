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

// checks if the CreateWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebhookRequest{}

// CreateWebhookRequest struct for CreateWebhookRequest
type CreateWebhookRequest struct {
	PayorId string `json:"payorId"`
	// the webhook URL to use.
	WebhookUrl string `json:"webhookUrl"`
	// the authorization header to include with the notification.
	AuthorizationHeader *string `json:"authorizationHeader,omitempty"`
	// whether the webhook is enabled.
	Enabled bool `json:"enabled"`
	// the categories to enable.
	Categories []Category `json:"categories,omitempty"`
}

// NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhookRequest(payorId string, webhookUrl string, enabled bool) *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	this.PayorId = payorId
	this.WebhookUrl = webhookUrl
	this.Enabled = enabled
	return &this
}

// NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	return &this
}

// GetPayorId returns the PayorId field value
func (o *CreateWebhookRequest) GetPayorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetPayorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayorId, true
}

// SetPayorId sets field value
func (o *CreateWebhookRequest) SetPayorId(v string) {
	o.PayorId = v
}

// GetWebhookUrl returns the WebhookUrl field value
func (o *CreateWebhookRequest) GetWebhookUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetWebhookUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookUrl, true
}

// SetWebhookUrl sets field value
func (o *CreateWebhookRequest) SetWebhookUrl(v string) {
	o.WebhookUrl = v
}

// GetAuthorizationHeader returns the AuthorizationHeader field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetAuthorizationHeader() string {
	if o == nil || IsNil(o.AuthorizationHeader) {
		var ret string
		return ret
	}
	return *o.AuthorizationHeader
}

// GetAuthorizationHeaderOk returns a tuple with the AuthorizationHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetAuthorizationHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationHeader) {
		return nil, false
	}
	return o.AuthorizationHeader, true
}

// HasAuthorizationHeader returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasAuthorizationHeader() bool {
	if o != nil && !IsNil(o.AuthorizationHeader) {
		return true
	}

	return false
}

// SetAuthorizationHeader gets a reference to the given string and assigns it to the AuthorizationHeader field.
func (o *CreateWebhookRequest) SetAuthorizationHeader(v string) {
	o.AuthorizationHeader = &v
}

// GetEnabled returns the Enabled field value
func (o *CreateWebhookRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateWebhookRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetCategories() []Category {
	if o == nil || IsNil(o.Categories) {
		var ret []Category
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetCategoriesOk() ([]Category, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []Category and assigns it to the Categories field.
func (o *CreateWebhookRequest) SetCategories(v []Category) {
	o.Categories = v
}

func (o CreateWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payorId"] = o.PayorId
	toSerialize["webhookUrl"] = o.WebhookUrl
	if !IsNil(o.AuthorizationHeader) {
		toSerialize["authorizationHeader"] = o.AuthorizationHeader
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableCreateWebhookRequest struct {
	value *CreateWebhookRequest
	isSet bool
}

func (v NullableCreateWebhookRequest) Get() *CreateWebhookRequest {
	return v.value
}

func (v *NullableCreateWebhookRequest) Set(val *CreateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhookRequest(val *CreateWebhookRequest) *NullableCreateWebhookRequest {
	return &NullableCreateWebhookRequest{value: val, isSet: true}
}

func (v NullableCreateWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


