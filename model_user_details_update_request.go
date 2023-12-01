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

// checks if the UserDetailsUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDetailsUpdateRequest{}

// UserDetailsUpdateRequest <p>All properties are optional</p> <p>Only provided properties will be updated</p> <p>Use null to null out a property that is allowed to be nullable</p> 
type UserDetailsUpdateRequest struct {
	// The main contact number for the user 
	PrimaryContactNumber NullableString `json:"primaryContactNumber,omitempty"`
	// The secondary contact number for the user 
	SecondaryContactNumber NullableString `json:"secondaryContactNumber,omitempty"`
	FirstName NullableString `json:"firstName,omitempty"`
	LastName NullableString `json:"lastName,omitempty"`
	// the email address of the user
	Email NullableString `json:"email,omitempty"`
	// The phone number of a device that the user can receive sms messages on 
	SmsNumber NullableString `json:"smsNumber,omitempty"`
	MfaType NullableMFAType `json:"mfaType,omitempty"`
	// <p>Optional property that MUST be suppied when manually verifying a user</p> <p>The user's smsNumber is registered via a separate endpoint and an OTP sent to them</p> 
	VerificationCode NullableString `json:"verificationCode,omitempty"`
}

// NewUserDetailsUpdateRequest instantiates a new UserDetailsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDetailsUpdateRequest() *UserDetailsUpdateRequest {
	this := UserDetailsUpdateRequest{}
	return &this
}

// NewUserDetailsUpdateRequestWithDefaults instantiates a new UserDetailsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDetailsUpdateRequestWithDefaults() *UserDetailsUpdateRequest {
	this := UserDetailsUpdateRequest{}
	return &this
}

// GetPrimaryContactNumber returns the PrimaryContactNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetailsUpdateRequest) GetPrimaryContactNumber() string {
	if o == nil || IsNil(o.PrimaryContactNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PrimaryContactNumber.Get()
}

// GetPrimaryContactNumberOk returns a tuple with the PrimaryContactNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetailsUpdateRequest) GetPrimaryContactNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryContactNumber.Get(), o.PrimaryContactNumber.IsSet()
}

// HasPrimaryContactNumber returns a boolean if a field has been set.
func (o *UserDetailsUpdateRequest) HasPrimaryContactNumber() bool {
	if o != nil && o.PrimaryContactNumber.IsSet() {
		return true
	}

	return false
}

// SetPrimaryContactNumber gets a reference to the given NullableString and assigns it to the PrimaryContactNumber field.
func (o *UserDetailsUpdateRequest) SetPrimaryContactNumber(v string) {
	o.PrimaryContactNumber.Set(&v)
}
// SetPrimaryContactNumberNil sets the value for PrimaryContactNumber to be an explicit nil
func (o *UserDetailsUpdateRequest) SetPrimaryContactNumberNil() {
	o.PrimaryContactNumber.Set(nil)
}

// UnsetPrimaryContactNumber ensures that no value is present for PrimaryContactNumber, not even an explicit nil
func (o *UserDetailsUpdateRequest) UnsetPrimaryContactNumber() {
	o.PrimaryContactNumber.Unset()
}

// GetSecondaryContactNumber returns the SecondaryContactNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetailsUpdateRequest) GetSecondaryContactNumber() string {
	if o == nil || IsNil(o.SecondaryContactNumber.Get()) {
		var ret string
		return ret
	}
	return *o.SecondaryContactNumber.Get()
}

// GetSecondaryContactNumberOk returns a tuple with the SecondaryContactNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetailsUpdateRequest) GetSecondaryContactNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondaryContactNumber.Get(), o.SecondaryContactNumber.IsSet()
}

// HasSecondaryContactNumber returns a boolean if a field has been set.
func (o *UserDetailsUpdateRequest) HasSecondaryContactNumber() bool {
	if o != nil && o.SecondaryContactNumber.IsSet() {
		return true
	}

	return false
}

// SetSecondaryContactNumber gets a reference to the given NullableString and assigns it to the SecondaryContactNumber field.
func (o *UserDetailsUpdateRequest) SetSecondaryContactNumber(v string) {
	o.SecondaryContactNumber.Set(&v)
}
// SetSecondaryContactNumberNil sets the value for SecondaryContactNumber to be an explicit nil
func (o *UserDetailsUpdateRequest) SetSecondaryContactNumberNil() {
	o.SecondaryContactNumber.Set(nil)
}

// UnsetSecondaryContactNumber ensures that no value is present for SecondaryContactNumber, not even an explicit nil
func (o *UserDetailsUpdateRequest) UnsetSecondaryContactNumber() {
	o.SecondaryContactNumber.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetailsUpdateRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetailsUpdateRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserDetailsUpdateRequest) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *UserDetailsUpdateRequest) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *UserDetailsUpdateRequest) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *UserDetailsUpdateRequest) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetailsUpdateRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetailsUpdateRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *UserDetailsUpdateRequest) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *UserDetailsUpdateRequest) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *UserDetailsUpdateRequest) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *UserDetailsUpdateRequest) UnsetLastName() {
	o.LastName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetailsUpdateRequest) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetailsUpdateRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserDetailsUpdateRequest) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserDetailsUpdateRequest) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserDetailsUpdateRequest) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserDetailsUpdateRequest) UnsetEmail() {
	o.Email.Unset()
}

// GetSmsNumber returns the SmsNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetailsUpdateRequest) GetSmsNumber() string {
	if o == nil || IsNil(o.SmsNumber.Get()) {
		var ret string
		return ret
	}
	return *o.SmsNumber.Get()
}

// GetSmsNumberOk returns a tuple with the SmsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetailsUpdateRequest) GetSmsNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmsNumber.Get(), o.SmsNumber.IsSet()
}

// HasSmsNumber returns a boolean if a field has been set.
func (o *UserDetailsUpdateRequest) HasSmsNumber() bool {
	if o != nil && o.SmsNumber.IsSet() {
		return true
	}

	return false
}

// SetSmsNumber gets a reference to the given NullableString and assigns it to the SmsNumber field.
func (o *UserDetailsUpdateRequest) SetSmsNumber(v string) {
	o.SmsNumber.Set(&v)
}
// SetSmsNumberNil sets the value for SmsNumber to be an explicit nil
func (o *UserDetailsUpdateRequest) SetSmsNumberNil() {
	o.SmsNumber.Set(nil)
}

// UnsetSmsNumber ensures that no value is present for SmsNumber, not even an explicit nil
func (o *UserDetailsUpdateRequest) UnsetSmsNumber() {
	o.SmsNumber.Unset()
}

// GetMfaType returns the MfaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetailsUpdateRequest) GetMfaType() MFAType {
	if o == nil || IsNil(o.MfaType.Get()) {
		var ret MFAType
		return ret
	}
	return *o.MfaType.Get()
}

// GetMfaTypeOk returns a tuple with the MfaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetailsUpdateRequest) GetMfaTypeOk() (*MFAType, bool) {
	if o == nil {
		return nil, false
	}
	return o.MfaType.Get(), o.MfaType.IsSet()
}

// HasMfaType returns a boolean if a field has been set.
func (o *UserDetailsUpdateRequest) HasMfaType() bool {
	if o != nil && o.MfaType.IsSet() {
		return true
	}

	return false
}

// SetMfaType gets a reference to the given NullableMFAType and assigns it to the MfaType field.
func (o *UserDetailsUpdateRequest) SetMfaType(v MFAType) {
	o.MfaType.Set(&v)
}
// SetMfaTypeNil sets the value for MfaType to be an explicit nil
func (o *UserDetailsUpdateRequest) SetMfaTypeNil() {
	o.MfaType.Set(nil)
}

// UnsetMfaType ensures that no value is present for MfaType, not even an explicit nil
func (o *UserDetailsUpdateRequest) UnsetMfaType() {
	o.MfaType.Unset()
}

// GetVerificationCode returns the VerificationCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetailsUpdateRequest) GetVerificationCode() string {
	if o == nil || IsNil(o.VerificationCode.Get()) {
		var ret string
		return ret
	}
	return *o.VerificationCode.Get()
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetailsUpdateRequest) GetVerificationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerificationCode.Get(), o.VerificationCode.IsSet()
}

// HasVerificationCode returns a boolean if a field has been set.
func (o *UserDetailsUpdateRequest) HasVerificationCode() bool {
	if o != nil && o.VerificationCode.IsSet() {
		return true
	}

	return false
}

// SetVerificationCode gets a reference to the given NullableString and assigns it to the VerificationCode field.
func (o *UserDetailsUpdateRequest) SetVerificationCode(v string) {
	o.VerificationCode.Set(&v)
}
// SetVerificationCodeNil sets the value for VerificationCode to be an explicit nil
func (o *UserDetailsUpdateRequest) SetVerificationCodeNil() {
	o.VerificationCode.Set(nil)
}

// UnsetVerificationCode ensures that no value is present for VerificationCode, not even an explicit nil
func (o *UserDetailsUpdateRequest) UnsetVerificationCode() {
	o.VerificationCode.Unset()
}

func (o UserDetailsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDetailsUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PrimaryContactNumber.IsSet() {
		toSerialize["primaryContactNumber"] = o.PrimaryContactNumber.Get()
	}
	if o.SecondaryContactNumber.IsSet() {
		toSerialize["secondaryContactNumber"] = o.SecondaryContactNumber.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.SmsNumber.IsSet() {
		toSerialize["smsNumber"] = o.SmsNumber.Get()
	}
	if o.MfaType.IsSet() {
		toSerialize["mfaType"] = o.MfaType.Get()
	}
	if o.VerificationCode.IsSet() {
		toSerialize["verificationCode"] = o.VerificationCode.Get()
	}
	return toSerialize, nil
}

type NullableUserDetailsUpdateRequest struct {
	value *UserDetailsUpdateRequest
	isSet bool
}

func (v NullableUserDetailsUpdateRequest) Get() *UserDetailsUpdateRequest {
	return v.value
}

func (v *NullableUserDetailsUpdateRequest) Set(val *UserDetailsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDetailsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDetailsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDetailsUpdateRequest(val *UserDetailsUpdateRequest) *NullableUserDetailsUpdateRequest {
	return &NullableUserDetailsUpdateRequest{value: val, isSet: true}
}

func (v NullableUserDetailsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDetailsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


