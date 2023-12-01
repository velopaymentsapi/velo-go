/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response.   ## Http Status Codes Following is a list of Http Status codes that could be returned by the platform      | Status Code            | Description                                                                          |     | -----------------------| -------------------------------------------------------------------------------------|     | 200 OK                 | The request was successfully processed and usually returns a json response           |     | 201 Created            | A resource was created and a Location header is returned linking to the new resource |     | 202 Accepted           | The request has been accepted for processing                                         |     | 204 No Content         | The request has been processed and there is no response (usually deletes and updates)|     | 400 Bad Request        | The request is invalid and should be fixed before retrying                           |     | 401 Unauthorized       | Authentication has failed, usually means the token has expired                       |     | 403 Forbidden          | The user does not have permissions for the request                                   |     | 404 Not Found          | The resource was not found                                                           |     | 409 Conflict           | The resource already exists and there is a conflict                                  |     | 429 Too Many Requests  | The user has submitted too many requests in a given amount of time                   |     | 5xx Server Error       | Platform internal error (should rarely happen)                                       | 

API version: 2.37.150
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// checks if the AccessTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenResponse{}

// AccessTokenResponse struct for AccessTokenResponse
type AccessTokenResponse struct {
	// Bearer token used in headers to access secure endpoints 
	AccessToken *string `json:"access_token,omitempty"`
	// the type of the token
	TokenType *string `json:"token_type,omitempty"`
	// can be used to obtain a new access token
	RefreshToken *string `json:"refresh_token,omitempty"`
	// The lifetime in seconds of the access token
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	// the scope of the access token
	Scope *string `json:"scope,omitempty"`
	UserInfo *UserInfo `json:"user_info,omitempty"`
	// If the user is a payee then the payeeId<P> If the user is a payor then the payorId 
	EntityIds []string `json:"entityIds,omitempty"`
}

// NewAccessTokenResponse instantiates a new AccessTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenResponse() *AccessTokenResponse {
	this := AccessTokenResponse{}
	var tokenType string = "bearer"
	this.TokenType = &tokenType
	return &this
}

// NewAccessTokenResponseWithDefaults instantiates a new AccessTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenResponseWithDefaults() *AccessTokenResponse {
	this := AccessTokenResponse{}
	var tokenType string = "bearer"
	this.TokenType = &tokenType
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AccessTokenResponse) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AccessTokenResponse) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AccessTokenResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *AccessTokenResponse) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *AccessTokenResponse) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *AccessTokenResponse) SetTokenType(v string) {
	o.TokenType = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *AccessTokenResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *AccessTokenResponse) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *AccessTokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *AccessTokenResponse) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *AccessTokenResponse) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *AccessTokenResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AccessTokenResponse) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenResponse) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AccessTokenResponse) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AccessTokenResponse) SetScope(v string) {
	o.Scope = &v
}

// GetUserInfo returns the UserInfo field value if set, zero value otherwise.
func (o *AccessTokenResponse) GetUserInfo() UserInfo {
	if o == nil || IsNil(o.UserInfo) {
		var ret UserInfo
		return ret
	}
	return *o.UserInfo
}

// GetUserInfoOk returns a tuple with the UserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenResponse) GetUserInfoOk() (*UserInfo, bool) {
	if o == nil || IsNil(o.UserInfo) {
		return nil, false
	}
	return o.UserInfo, true
}

// HasUserInfo returns a boolean if a field has been set.
func (o *AccessTokenResponse) HasUserInfo() bool {
	if o != nil && !IsNil(o.UserInfo) {
		return true
	}

	return false
}

// SetUserInfo gets a reference to the given UserInfo and assigns it to the UserInfo field.
func (o *AccessTokenResponse) SetUserInfo(v UserInfo) {
	o.UserInfo = &v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *AccessTokenResponse) GetEntityIds() []string {
	if o == nil || IsNil(o.EntityIds) {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenResponse) GetEntityIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EntityIds) {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *AccessTokenResponse) HasEntityIds() bool {
	if o != nil && !IsNil(o.EntityIds) {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *AccessTokenResponse) SetEntityIds(v []string) {
	o.EntityIds = v
}

func (o AccessTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.TokenType) {
		toSerialize["token_type"] = o.TokenType
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.UserInfo) {
		toSerialize["user_info"] = o.UserInfo
	}
	if !IsNil(o.EntityIds) {
		toSerialize["entityIds"] = o.EntityIds
	}
	return toSerialize, nil
}

type NullableAccessTokenResponse struct {
	value *AccessTokenResponse
	isSet bool
}

func (v NullableAccessTokenResponse) Get() *AccessTokenResponse {
	return v.value
}

func (v *NullableAccessTokenResponse) Set(val *AccessTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenResponse(val *AccessTokenResponse) *NullableAccessTokenResponse {
	return &NullableAccessTokenResponse{value: val, isSet: true}
}

func (v NullableAccessTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


