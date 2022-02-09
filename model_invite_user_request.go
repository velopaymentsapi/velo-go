/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 

API version: 2.29.128
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// InviteUserRequest struct for InviteUserRequest
type InviteUserRequest struct {
	// the email address of the invited user
	Email string `json:"email"`
	// <p>The MFA type that the user will use</p> <p>The type may be conditional on the role(s) the user has</p> 
	MfaType string `json:"mfaType"`
	// The phone number of a device that the user can receive sms messages on 
	SmsNumber string `json:"smsNumber"`
	// The main contact number for the user 
	PrimaryContactNumber string `json:"primaryContactNumber"`
	// The secondary contact number for the user 
	SecondaryContactNumber NullableString `json:"secondaryContactNumber,omitempty"`
	// The role(s) for the user The role must exist The role can be a custom role or a system role but the invoker must have the permissions to assign the role System roles are: velo.backoffice.admin, velo.payor.master_admin, velo.payor.admin, velo.payor.support, velo.payee.admin, velo.payee.support 
	Roles []string `json:"roles"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	// The payorId or payeeId or null if the user is a backoffice admin 
	EntityId NullableString `json:"entityId,omitempty"`
	// Will default to PAYOR if not provided but entityId is provided
	UserType *string `json:"userType,omitempty"`
	// Optional property that MUST be suppied when manually verifying a user The user's smsNumber is registered via a separate endpoint and an OTP sent to them 
	VerificationCode NullableString `json:"verificationCode,omitempty"`
}

// NewInviteUserRequest instantiates a new InviteUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteUserRequest(email string, mfaType string, smsNumber string, primaryContactNumber string, roles []string) *InviteUserRequest {
	this := InviteUserRequest{}
	this.Email = email
	this.MfaType = mfaType
	this.SmsNumber = smsNumber
	this.PrimaryContactNumber = primaryContactNumber
	this.Roles = roles
	return &this
}

// NewInviteUserRequestWithDefaults instantiates a new InviteUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteUserRequestWithDefaults() *InviteUserRequest {
	this := InviteUserRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *InviteUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InviteUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetMfaType returns the MfaType field value
func (o *InviteUserRequest) GetMfaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MfaType
}

// GetMfaTypeOk returns a tuple with the MfaType field value
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetMfaTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MfaType, true
}

// SetMfaType sets field value
func (o *InviteUserRequest) SetMfaType(v string) {
	o.MfaType = v
}

// GetSmsNumber returns the SmsNumber field value
func (o *InviteUserRequest) GetSmsNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmsNumber
}

// GetSmsNumberOk returns a tuple with the SmsNumber field value
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetSmsNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SmsNumber, true
}

// SetSmsNumber sets field value
func (o *InviteUserRequest) SetSmsNumber(v string) {
	o.SmsNumber = v
}

// GetPrimaryContactNumber returns the PrimaryContactNumber field value
func (o *InviteUserRequest) GetPrimaryContactNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryContactNumber
}

// GetPrimaryContactNumberOk returns a tuple with the PrimaryContactNumber field value
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetPrimaryContactNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrimaryContactNumber, true
}

// SetPrimaryContactNumber sets field value
func (o *InviteUserRequest) SetPrimaryContactNumber(v string) {
	o.PrimaryContactNumber = v
}

// GetSecondaryContactNumber returns the SecondaryContactNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteUserRequest) GetSecondaryContactNumber() string {
	if o == nil || o.SecondaryContactNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecondaryContactNumber.Get()
}

// GetSecondaryContactNumberOk returns a tuple with the SecondaryContactNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteUserRequest) GetSecondaryContactNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecondaryContactNumber.Get(), o.SecondaryContactNumber.IsSet()
}

// HasSecondaryContactNumber returns a boolean if a field has been set.
func (o *InviteUserRequest) HasSecondaryContactNumber() bool {
	if o != nil && o.SecondaryContactNumber.IsSet() {
		return true
	}

	return false
}

// SetSecondaryContactNumber gets a reference to the given NullableString and assigns it to the SecondaryContactNumber field.
func (o *InviteUserRequest) SetSecondaryContactNumber(v string) {
	o.SecondaryContactNumber.Set(&v)
}
// SetSecondaryContactNumberNil sets the value for SecondaryContactNumber to be an explicit nil
func (o *InviteUserRequest) SetSecondaryContactNumberNil() {
	o.SecondaryContactNumber.Set(nil)
}

// UnsetSecondaryContactNumber ensures that no value is present for SecondaryContactNumber, not even an explicit nil
func (o *InviteUserRequest) UnsetSecondaryContactNumber() {
	o.SecondaryContactNumber.Unset()
}

// GetRoles returns the Roles field value
func (o *InviteUserRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetRolesOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *InviteUserRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *InviteUserRequest) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *InviteUserRequest) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *InviteUserRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *InviteUserRequest) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *InviteUserRequest) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *InviteUserRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteUserRequest) GetEntityId() string {
	if o == nil || o.EntityId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteUserRequest) GetEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field has been set.
func (o *InviteUserRequest) HasEntityId() bool {
	if o != nil && o.EntityId.IsSet() {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableString and assigns it to the EntityId field.
func (o *InviteUserRequest) SetEntityId(v string) {
	o.EntityId.Set(&v)
}
// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *InviteUserRequest) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *InviteUserRequest) UnsetEntityId() {
	o.EntityId.Unset()
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *InviteUserRequest) GetUserType() string {
	if o == nil || o.UserType == nil {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteUserRequest) GetUserTypeOk() (*string, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *InviteUserRequest) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *InviteUserRequest) SetUserType(v string) {
	o.UserType = &v
}

// GetVerificationCode returns the VerificationCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteUserRequest) GetVerificationCode() string {
	if o == nil || o.VerificationCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.VerificationCode.Get()
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteUserRequest) GetVerificationCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VerificationCode.Get(), o.VerificationCode.IsSet()
}

// HasVerificationCode returns a boolean if a field has been set.
func (o *InviteUserRequest) HasVerificationCode() bool {
	if o != nil && o.VerificationCode.IsSet() {
		return true
	}

	return false
}

// SetVerificationCode gets a reference to the given NullableString and assigns it to the VerificationCode field.
func (o *InviteUserRequest) SetVerificationCode(v string) {
	o.VerificationCode.Set(&v)
}
// SetVerificationCodeNil sets the value for VerificationCode to be an explicit nil
func (o *InviteUserRequest) SetVerificationCodeNil() {
	o.VerificationCode.Set(nil)
}

// UnsetVerificationCode ensures that no value is present for VerificationCode, not even an explicit nil
func (o *InviteUserRequest) UnsetVerificationCode() {
	o.VerificationCode.Unset()
}

func (o InviteUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["mfaType"] = o.MfaType
	}
	if true {
		toSerialize["smsNumber"] = o.SmsNumber
	}
	if true {
		toSerialize["primaryContactNumber"] = o.PrimaryContactNumber
	}
	if o.SecondaryContactNumber.IsSet() {
		toSerialize["secondaryContactNumber"] = o.SecondaryContactNumber.Get()
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.EntityId.IsSet() {
		toSerialize["entityId"] = o.EntityId.Get()
	}
	if o.UserType != nil {
		toSerialize["userType"] = o.UserType
	}
	if o.VerificationCode.IsSet() {
		toSerialize["verificationCode"] = o.VerificationCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInviteUserRequest struct {
	value *InviteUserRequest
	isSet bool
}

func (v NullableInviteUserRequest) Get() *InviteUserRequest {
	return v.value
}

func (v *NullableInviteUserRequest) Set(val *InviteUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteUserRequest(val *InviteUserRequest) *NullableInviteUserRequest {
	return &NullableInviteUserRequest{value: val, isSet: true}
}

func (v NullableInviteUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


