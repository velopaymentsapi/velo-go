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
)

// AuthResponse struct for AuthResponse
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn *float32 `json:"expires_in,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	Scope *string `json:"scope,omitempty"`
	EntityIds *[]string `json:"entityIds,omitempty"`
}

// NewAuthResponse instantiates a new AuthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthResponse(accessToken string, tokenType string) *AuthResponse {
	this := AuthResponse{}
	this.AccessToken = accessToken
	this.TokenType = tokenType
	return &this
}

// NewAuthResponseWithDefaults instantiates a new AuthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthResponseWithDefaults() *AuthResponse {
	this := AuthResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *AuthResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AuthResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetTokenType returns the TokenType field value
func (o *AuthResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *AuthResponse) SetTokenType(v string) {
	o.TokenType = v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *AuthResponse) GetExpiresIn() float32 {
	if o == nil || o.ExpiresIn == nil {
		var ret float32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetExpiresInOk() (*float32, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *AuthResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given float32 and assigns it to the ExpiresIn field.
func (o *AuthResponse) SetExpiresIn(v float32) {
	o.ExpiresIn = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *AuthResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *AuthResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *AuthResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AuthResponse) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AuthResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AuthResponse) SetScope(v string) {
	o.Scope = &v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *AuthResponse) GetEntityIds() []string {
	if o == nil || o.EntityIds == nil {
		var ret []string
		return ret
	}
	return *o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetEntityIdsOk() (*[]string, bool) {
	if o == nil || o.EntityIds == nil {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *AuthResponse) HasEntityIds() bool {
	if o != nil && o.EntityIds != nil {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *AuthResponse) SetEntityIds(v []string) {
	o.EntityIds = &v
}

func (o AuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.RefreshToken != nil {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.EntityIds != nil {
		toSerialize["entityIds"] = o.EntityIds
	}
	return json.Marshal(toSerialize)
}

type NullableAuthResponse struct {
	value *AuthResponse
	isSet bool
}

func (v NullableAuthResponse) Get() *AuthResponse {
	return v.value
}

func (v *NullableAuthResponse) Set(val *AuthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthResponse(val *AuthResponse) *NullableAuthResponse {
	return &NullableAuthResponse{value: val, isSet: true}
}

func (v NullableAuthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


