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

// checks if the UpdatePayeeDetailsRequestV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePayeeDetailsRequestV3{}

// UpdatePayeeDetailsRequestV3 struct for UpdatePayeeDetailsRequestV3
type UpdatePayeeDetailsRequestV3 struct {
	Address *PayeeAddressV3 `json:"address,omitempty"`
	Individual *IndividualV3 `json:"individual,omitempty"`
	Company NullableCompanyV3 `json:"company,omitempty"`
	// An IETF BCP 47 language code which has been configured for use within this Velo environment.<BR> See the /v1/supportedLanguages endpoint to list the available codes for an environment. 
	Language *string `json:"language,omitempty"`
	PayeeType *PayeeTypeEnum `json:"payeeType,omitempty"`
	Challenge *ChallengeV3 `json:"challenge,omitempty"`
	Email NullableString `json:"email,omitempty"`
	// The phone number of a device that the payee wishes to receive sms messages on 
	ContactSmsNumber *string `json:"contactSmsNumber,omitempty"`
}

// NewUpdatePayeeDetailsRequestV3 instantiates a new UpdatePayeeDetailsRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePayeeDetailsRequestV3() *UpdatePayeeDetailsRequestV3 {
	this := UpdatePayeeDetailsRequestV3{}
	return &this
}

// NewUpdatePayeeDetailsRequestV3WithDefaults instantiates a new UpdatePayeeDetailsRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePayeeDetailsRequestV3WithDefaults() *UpdatePayeeDetailsRequestV3 {
	this := UpdatePayeeDetailsRequestV3{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdatePayeeDetailsRequestV3) GetAddress() PayeeAddressV3 {
	if o == nil || IsNil(o.Address) {
		var ret PayeeAddressV3
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePayeeDetailsRequestV3) GetAddressOk() (*PayeeAddressV3, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdatePayeeDetailsRequestV3) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given PayeeAddressV3 and assigns it to the Address field.
func (o *UpdatePayeeDetailsRequestV3) SetAddress(v PayeeAddressV3) {
	o.Address = &v
}

// GetIndividual returns the Individual field value if set, zero value otherwise.
func (o *UpdatePayeeDetailsRequestV3) GetIndividual() IndividualV3 {
	if o == nil || IsNil(o.Individual) {
		var ret IndividualV3
		return ret
	}
	return *o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePayeeDetailsRequestV3) GetIndividualOk() (*IndividualV3, bool) {
	if o == nil || IsNil(o.Individual) {
		return nil, false
	}
	return o.Individual, true
}

// HasIndividual returns a boolean if a field has been set.
func (o *UpdatePayeeDetailsRequestV3) HasIndividual() bool {
	if o != nil && !IsNil(o.Individual) {
		return true
	}

	return false
}

// SetIndividual gets a reference to the given IndividualV3 and assigns it to the Individual field.
func (o *UpdatePayeeDetailsRequestV3) SetIndividual(v IndividualV3) {
	o.Individual = &v
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePayeeDetailsRequestV3) GetCompany() CompanyV3 {
	if o == nil || IsNil(o.Company.Get()) {
		var ret CompanyV3
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePayeeDetailsRequestV3) GetCompanyOk() (*CompanyV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *UpdatePayeeDetailsRequestV3) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableCompanyV3 and assigns it to the Company field.
func (o *UpdatePayeeDetailsRequestV3) SetCompany(v CompanyV3) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *UpdatePayeeDetailsRequestV3) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *UpdatePayeeDetailsRequestV3) UnsetCompany() {
	o.Company.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *UpdatePayeeDetailsRequestV3) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePayeeDetailsRequestV3) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *UpdatePayeeDetailsRequestV3) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *UpdatePayeeDetailsRequestV3) SetLanguage(v string) {
	o.Language = &v
}

// GetPayeeType returns the PayeeType field value if set, zero value otherwise.
func (o *UpdatePayeeDetailsRequestV3) GetPayeeType() PayeeTypeEnum {
	if o == nil || IsNil(o.PayeeType) {
		var ret PayeeTypeEnum
		return ret
	}
	return *o.PayeeType
}

// GetPayeeTypeOk returns a tuple with the PayeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePayeeDetailsRequestV3) GetPayeeTypeOk() (*PayeeTypeEnum, bool) {
	if o == nil || IsNil(o.PayeeType) {
		return nil, false
	}
	return o.PayeeType, true
}

// HasPayeeType returns a boolean if a field has been set.
func (o *UpdatePayeeDetailsRequestV3) HasPayeeType() bool {
	if o != nil && !IsNil(o.PayeeType) {
		return true
	}

	return false
}

// SetPayeeType gets a reference to the given PayeeTypeEnum and assigns it to the PayeeType field.
func (o *UpdatePayeeDetailsRequestV3) SetPayeeType(v PayeeTypeEnum) {
	o.PayeeType = &v
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *UpdatePayeeDetailsRequestV3) GetChallenge() ChallengeV3 {
	if o == nil || IsNil(o.Challenge) {
		var ret ChallengeV3
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePayeeDetailsRequestV3) GetChallengeOk() (*ChallengeV3, bool) {
	if o == nil || IsNil(o.Challenge) {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *UpdatePayeeDetailsRequestV3) HasChallenge() bool {
	if o != nil && !IsNil(o.Challenge) {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given ChallengeV3 and assigns it to the Challenge field.
func (o *UpdatePayeeDetailsRequestV3) SetChallenge(v ChallengeV3) {
	o.Challenge = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePayeeDetailsRequestV3) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePayeeDetailsRequestV3) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdatePayeeDetailsRequestV3) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UpdatePayeeDetailsRequestV3) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UpdatePayeeDetailsRequestV3) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UpdatePayeeDetailsRequestV3) UnsetEmail() {
	o.Email.Unset()
}

// GetContactSmsNumber returns the ContactSmsNumber field value if set, zero value otherwise.
func (o *UpdatePayeeDetailsRequestV3) GetContactSmsNumber() string {
	if o == nil || IsNil(o.ContactSmsNumber) {
		var ret string
		return ret
	}
	return *o.ContactSmsNumber
}

// GetContactSmsNumberOk returns a tuple with the ContactSmsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePayeeDetailsRequestV3) GetContactSmsNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ContactSmsNumber) {
		return nil, false
	}
	return o.ContactSmsNumber, true
}

// HasContactSmsNumber returns a boolean if a field has been set.
func (o *UpdatePayeeDetailsRequestV3) HasContactSmsNumber() bool {
	if o != nil && !IsNil(o.ContactSmsNumber) {
		return true
	}

	return false
}

// SetContactSmsNumber gets a reference to the given string and assigns it to the ContactSmsNumber field.
func (o *UpdatePayeeDetailsRequestV3) SetContactSmsNumber(v string) {
	o.ContactSmsNumber = &v
}

func (o UpdatePayeeDetailsRequestV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePayeeDetailsRequestV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Individual) {
		toSerialize["individual"] = o.Individual
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.PayeeType) {
		toSerialize["payeeType"] = o.PayeeType
	}
	if !IsNil(o.Challenge) {
		toSerialize["challenge"] = o.Challenge
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.ContactSmsNumber) {
		toSerialize["contactSmsNumber"] = o.ContactSmsNumber
	}
	return toSerialize, nil
}

type NullableUpdatePayeeDetailsRequestV3 struct {
	value *UpdatePayeeDetailsRequestV3
	isSet bool
}

func (v NullableUpdatePayeeDetailsRequestV3) Get() *UpdatePayeeDetailsRequestV3 {
	return v.value
}

func (v *NullableUpdatePayeeDetailsRequestV3) Set(val *UpdatePayeeDetailsRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePayeeDetailsRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePayeeDetailsRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePayeeDetailsRequestV3(val *UpdatePayeeDetailsRequestV3) *NullableUpdatePayeeDetailsRequestV3 {
	return &NullableUpdatePayeeDetailsRequestV3{value: val, isSet: true}
}

func (v NullableUpdatePayeeDetailsRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePayeeDetailsRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


