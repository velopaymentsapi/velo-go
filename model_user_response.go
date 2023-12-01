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

// checks if the UserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponse{}

// UserResponse struct for UserResponse
type UserResponse struct {
	// The id of the user
	Id *string `json:"id,omitempty"`
	// The status of the user when the user has been invited but not yet enrolled they will have a PENDING status 
	Status *string `json:"status,omitempty"`
	// the email address of the user
	Email *string `json:"email,omitempty"`
	// The phone number of a device that the user can receive sms messages on 
	SmsNumber *string `json:"smsNumber,omitempty"`
	// The main contact number for the user 
	PrimaryContactNumber *string `json:"primaryContactNumber,omitempty"`
	// The secondary contact number for the user 
	SecondaryContactNumber *string `json:"secondaryContactNumber,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	// The payorId or payeeId or null if the user is not a payor or payee user 
	EntityId *string `json:"entityId,omitempty"`
	// The payor or payee company name or null if the user is not a payor or payee user 
	CompanyName *string `json:"companyName,omitempty"`
	// The role(s) for the user 
	Roles []Role `json:"roles,omitempty"`
	// Indicates the type of user. Could be BACKOFFICE, PAYOR or PAYEE.
	UserType *string `json:"userType,omitempty"`
	// The type of the MFA device
	MfaType *string `json:"mfaType,omitempty"`
	// The status of the MFA device
	MfaStatus *string `json:"mfaStatus,omitempty"`
	// If true the user is currently locked out and unable to log in
	LockedOut *bool `json:"lockedOut,omitempty"`
	// A timestamp showing when the user was locked out If null then the user is not currently locked out 
	LockedOutTimestamp NullableTime `json:"lockedOutTimestamp,omitempty"`
}

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse() *UserResponse {
	this := UserResponse{}
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserResponse) SetStatus(v string) {
	o.Status = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserResponse) SetEmail(v string) {
	o.Email = &v
}

// GetSmsNumber returns the SmsNumber field value if set, zero value otherwise.
func (o *UserResponse) GetSmsNumber() string {
	if o == nil || IsNil(o.SmsNumber) {
		var ret string
		return ret
	}
	return *o.SmsNumber
}

// GetSmsNumberOk returns a tuple with the SmsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSmsNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SmsNumber) {
		return nil, false
	}
	return o.SmsNumber, true
}

// HasSmsNumber returns a boolean if a field has been set.
func (o *UserResponse) HasSmsNumber() bool {
	if o != nil && !IsNil(o.SmsNumber) {
		return true
	}

	return false
}

// SetSmsNumber gets a reference to the given string and assigns it to the SmsNumber field.
func (o *UserResponse) SetSmsNumber(v string) {
	o.SmsNumber = &v
}

// GetPrimaryContactNumber returns the PrimaryContactNumber field value if set, zero value otherwise.
func (o *UserResponse) GetPrimaryContactNumber() string {
	if o == nil || IsNil(o.PrimaryContactNumber) {
		var ret string
		return ret
	}
	return *o.PrimaryContactNumber
}

// GetPrimaryContactNumberOk returns a tuple with the PrimaryContactNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetPrimaryContactNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactNumber) {
		return nil, false
	}
	return o.PrimaryContactNumber, true
}

// HasPrimaryContactNumber returns a boolean if a field has been set.
func (o *UserResponse) HasPrimaryContactNumber() bool {
	if o != nil && !IsNil(o.PrimaryContactNumber) {
		return true
	}

	return false
}

// SetPrimaryContactNumber gets a reference to the given string and assigns it to the PrimaryContactNumber field.
func (o *UserResponse) SetPrimaryContactNumber(v string) {
	o.PrimaryContactNumber = &v
}

// GetSecondaryContactNumber returns the SecondaryContactNumber field value if set, zero value otherwise.
func (o *UserResponse) GetSecondaryContactNumber() string {
	if o == nil || IsNil(o.SecondaryContactNumber) {
		var ret string
		return ret
	}
	return *o.SecondaryContactNumber
}

// GetSecondaryContactNumberOk returns a tuple with the SecondaryContactNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSecondaryContactNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryContactNumber) {
		return nil, false
	}
	return o.SecondaryContactNumber, true
}

// HasSecondaryContactNumber returns a boolean if a field has been set.
func (o *UserResponse) HasSecondaryContactNumber() bool {
	if o != nil && !IsNil(o.SecondaryContactNumber) {
		return true
	}

	return false
}

// SetSecondaryContactNumber gets a reference to the given string and assigns it to the SecondaryContactNumber field.
func (o *UserResponse) SetSecondaryContactNumber(v string) {
	o.SecondaryContactNumber = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserResponse) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserResponse) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserResponse) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *UserResponse) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *UserResponse) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *UserResponse) SetEntityId(v string) {
	o.EntityId = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UserResponse) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UserResponse) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UserResponse) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserResponse) GetRoles() []Role {
	if o == nil || IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetRolesOk() ([]Role, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserResponse) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *UserResponse) SetRoles(v []Role) {
	o.Roles = v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *UserResponse) GetUserType() string {
	if o == nil || IsNil(o.UserType) {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *UserResponse) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *UserResponse) SetUserType(v string) {
	o.UserType = &v
}

// GetMfaType returns the MfaType field value if set, zero value otherwise.
func (o *UserResponse) GetMfaType() string {
	if o == nil || IsNil(o.MfaType) {
		var ret string
		return ret
	}
	return *o.MfaType
}

// GetMfaTypeOk returns a tuple with the MfaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetMfaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MfaType) {
		return nil, false
	}
	return o.MfaType, true
}

// HasMfaType returns a boolean if a field has been set.
func (o *UserResponse) HasMfaType() bool {
	if o != nil && !IsNil(o.MfaType) {
		return true
	}

	return false
}

// SetMfaType gets a reference to the given string and assigns it to the MfaType field.
func (o *UserResponse) SetMfaType(v string) {
	o.MfaType = &v
}

// GetMfaStatus returns the MfaStatus field value if set, zero value otherwise.
func (o *UserResponse) GetMfaStatus() string {
	if o == nil || IsNil(o.MfaStatus) {
		var ret string
		return ret
	}
	return *o.MfaStatus
}

// GetMfaStatusOk returns a tuple with the MfaStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetMfaStatusOk() (*string, bool) {
	if o == nil || IsNil(o.MfaStatus) {
		return nil, false
	}
	return o.MfaStatus, true
}

// HasMfaStatus returns a boolean if a field has been set.
func (o *UserResponse) HasMfaStatus() bool {
	if o != nil && !IsNil(o.MfaStatus) {
		return true
	}

	return false
}

// SetMfaStatus gets a reference to the given string and assigns it to the MfaStatus field.
func (o *UserResponse) SetMfaStatus(v string) {
	o.MfaStatus = &v
}

// GetLockedOut returns the LockedOut field value if set, zero value otherwise.
func (o *UserResponse) GetLockedOut() bool {
	if o == nil || IsNil(o.LockedOut) {
		var ret bool
		return ret
	}
	return *o.LockedOut
}

// GetLockedOutOk returns a tuple with the LockedOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLockedOutOk() (*bool, bool) {
	if o == nil || IsNil(o.LockedOut) {
		return nil, false
	}
	return o.LockedOut, true
}

// HasLockedOut returns a boolean if a field has been set.
func (o *UserResponse) HasLockedOut() bool {
	if o != nil && !IsNil(o.LockedOut) {
		return true
	}

	return false
}

// SetLockedOut gets a reference to the given bool and assigns it to the LockedOut field.
func (o *UserResponse) SetLockedOut(v bool) {
	o.LockedOut = &v
}

// GetLockedOutTimestamp returns the LockedOutTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetLockedOutTimestamp() time.Time {
	if o == nil || IsNil(o.LockedOutTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LockedOutTimestamp.Get()
}

// GetLockedOutTimestampOk returns a tuple with the LockedOutTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetLockedOutTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedOutTimestamp.Get(), o.LockedOutTimestamp.IsSet()
}

// HasLockedOutTimestamp returns a boolean if a field has been set.
func (o *UserResponse) HasLockedOutTimestamp() bool {
	if o != nil && o.LockedOutTimestamp.IsSet() {
		return true
	}

	return false
}

// SetLockedOutTimestamp gets a reference to the given NullableTime and assigns it to the LockedOutTimestamp field.
func (o *UserResponse) SetLockedOutTimestamp(v time.Time) {
	o.LockedOutTimestamp.Set(&v)
}
// SetLockedOutTimestampNil sets the value for LockedOutTimestamp to be an explicit nil
func (o *UserResponse) SetLockedOutTimestampNil() {
	o.LockedOutTimestamp.Set(nil)
}

// UnsetLockedOutTimestamp ensures that no value is present for LockedOutTimestamp, not even an explicit nil
func (o *UserResponse) UnsetLockedOutTimestamp() {
	o.LockedOutTimestamp.Unset()
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.SmsNumber) {
		toSerialize["smsNumber"] = o.SmsNumber
	}
	if !IsNil(o.PrimaryContactNumber) {
		toSerialize["primaryContactNumber"] = o.PrimaryContactNumber
	}
	if !IsNil(o.SecondaryContactNumber) {
		toSerialize["secondaryContactNumber"] = o.SecondaryContactNumber
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.UserType) {
		toSerialize["userType"] = o.UserType
	}
	if !IsNil(o.MfaType) {
		toSerialize["mfaType"] = o.MfaType
	}
	if !IsNil(o.MfaStatus) {
		toSerialize["mfaStatus"] = o.MfaStatus
	}
	if !IsNil(o.LockedOut) {
		toSerialize["lockedOut"] = o.LockedOut
	}
	if o.LockedOutTimestamp.IsSet() {
		toSerialize["lockedOutTimestamp"] = o.LockedOutTimestamp.Get()
	}
	return toSerialize, nil
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


