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

// FailedPayee2 struct for FailedPayee2
type FailedPayee2 struct {
	PayeeId *string `json:"payeeId,omitempty"`
	PayorRefs []PayeePayorRef `json:"payorRefs,omitempty"`
	Email *string `json:"email,omitempty"`
	RemoteId *string `json:"remoteId,omitempty"`
	Type *PayeeType `json:"type,omitempty"`
	Address *CreatePayeeAddress2 `json:"address,omitempty"`
	PaymentChannel *CreatePaymentChannel2 `json:"paymentChannel,omitempty"`
	Challenge *Challenge2 `json:"challenge,omitempty"`
	// An IETF BCP 47 language code which has been configured for use within this Velo environment.<BR> See the /v1/supportedLanguages endpoint to list the available codes for an environment. 
	Language *string `json:"language,omitempty"`
	Company NullableCompany2 `json:"company,omitempty"`
	Individual *CreateIndividual2 `json:"individual,omitempty"`
}

// NewFailedPayee2 instantiates a new FailedPayee2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailedPayee2() *FailedPayee2 {
	this := FailedPayee2{}
	return &this
}

// NewFailedPayee2WithDefaults instantiates a new FailedPayee2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailedPayee2WithDefaults() *FailedPayee2 {
	this := FailedPayee2{}
	return &this
}

// GetPayeeId returns the PayeeId field value if set, zero value otherwise.
func (o *FailedPayee2) GetPayeeId() string {
	if o == nil || o.PayeeId == nil {
		var ret string
		return ret
	}
	return *o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedPayee2) GetPayeeIdOk() (*string, bool) {
	if o == nil || o.PayeeId == nil {
		return nil, false
	}
	return o.PayeeId, true
}

// HasPayeeId returns a boolean if a field has been set.
func (o *FailedPayee2) HasPayeeId() bool {
	if o != nil && o.PayeeId != nil {
		return true
	}

	return false
}

// SetPayeeId gets a reference to the given string and assigns it to the PayeeId field.
func (o *FailedPayee2) SetPayeeId(v string) {
	o.PayeeId = &v
}

// GetPayorRefs returns the PayorRefs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailedPayee2) GetPayorRefs() []PayeePayorRef {
	if o == nil  {
		var ret []PayeePayorRef
		return ret
	}
	return o.PayorRefs
}

// GetPayorRefsOk returns a tuple with the PayorRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailedPayee2) GetPayorRefsOk() (*[]PayeePayorRef, bool) {
	if o == nil || o.PayorRefs == nil {
		return nil, false
	}
	return &o.PayorRefs, true
}

// HasPayorRefs returns a boolean if a field has been set.
func (o *FailedPayee2) HasPayorRefs() bool {
	if o != nil && o.PayorRefs != nil {
		return true
	}

	return false
}

// SetPayorRefs gets a reference to the given []PayeePayorRef and assigns it to the PayorRefs field.
func (o *FailedPayee2) SetPayorRefs(v []PayeePayorRef) {
	o.PayorRefs = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *FailedPayee2) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedPayee2) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *FailedPayee2) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *FailedPayee2) SetEmail(v string) {
	o.Email = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *FailedPayee2) GetRemoteId() string {
	if o == nil || o.RemoteId == nil {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedPayee2) GetRemoteIdOk() (*string, bool) {
	if o == nil || o.RemoteId == nil {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *FailedPayee2) HasRemoteId() bool {
	if o != nil && o.RemoteId != nil {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *FailedPayee2) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FailedPayee2) GetType() PayeeType {
	if o == nil || o.Type == nil {
		var ret PayeeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedPayee2) GetTypeOk() (*PayeeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FailedPayee2) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PayeeType and assigns it to the Type field.
func (o *FailedPayee2) SetType(v PayeeType) {
	o.Type = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *FailedPayee2) GetAddress() CreatePayeeAddress2 {
	if o == nil || o.Address == nil {
		var ret CreatePayeeAddress2
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedPayee2) GetAddressOk() (*CreatePayeeAddress2, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *FailedPayee2) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CreatePayeeAddress2 and assigns it to the Address field.
func (o *FailedPayee2) SetAddress(v CreatePayeeAddress2) {
	o.Address = &v
}

// GetPaymentChannel returns the PaymentChannel field value if set, zero value otherwise.
func (o *FailedPayee2) GetPaymentChannel() CreatePaymentChannel2 {
	if o == nil || o.PaymentChannel == nil {
		var ret CreatePaymentChannel2
		return ret
	}
	return *o.PaymentChannel
}

// GetPaymentChannelOk returns a tuple with the PaymentChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedPayee2) GetPaymentChannelOk() (*CreatePaymentChannel2, bool) {
	if o == nil || o.PaymentChannel == nil {
		return nil, false
	}
	return o.PaymentChannel, true
}

// HasPaymentChannel returns a boolean if a field has been set.
func (o *FailedPayee2) HasPaymentChannel() bool {
	if o != nil && o.PaymentChannel != nil {
		return true
	}

	return false
}

// SetPaymentChannel gets a reference to the given CreatePaymentChannel2 and assigns it to the PaymentChannel field.
func (o *FailedPayee2) SetPaymentChannel(v CreatePaymentChannel2) {
	o.PaymentChannel = &v
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *FailedPayee2) GetChallenge() Challenge2 {
	if o == nil || o.Challenge == nil {
		var ret Challenge2
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedPayee2) GetChallengeOk() (*Challenge2, bool) {
	if o == nil || o.Challenge == nil {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *FailedPayee2) HasChallenge() bool {
	if o != nil && o.Challenge != nil {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given Challenge2 and assigns it to the Challenge field.
func (o *FailedPayee2) SetChallenge(v Challenge2) {
	o.Challenge = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *FailedPayee2) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedPayee2) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *FailedPayee2) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *FailedPayee2) SetLanguage(v string) {
	o.Language = &v
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailedPayee2) GetCompany() Company2 {
	if o == nil || o.Company.Get() == nil {
		var ret Company2
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailedPayee2) GetCompanyOk() (*Company2, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *FailedPayee2) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableCompany2 and assigns it to the Company field.
func (o *FailedPayee2) SetCompany(v Company2) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *FailedPayee2) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *FailedPayee2) UnsetCompany() {
	o.Company.Unset()
}

// GetIndividual returns the Individual field value if set, zero value otherwise.
func (o *FailedPayee2) GetIndividual() CreateIndividual2 {
	if o == nil || o.Individual == nil {
		var ret CreateIndividual2
		return ret
	}
	return *o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedPayee2) GetIndividualOk() (*CreateIndividual2, bool) {
	if o == nil || o.Individual == nil {
		return nil, false
	}
	return o.Individual, true
}

// HasIndividual returns a boolean if a field has been set.
func (o *FailedPayee2) HasIndividual() bool {
	if o != nil && o.Individual != nil {
		return true
	}

	return false
}

// SetIndividual gets a reference to the given CreateIndividual2 and assigns it to the Individual field.
func (o *FailedPayee2) SetIndividual(v CreateIndividual2) {
	o.Individual = &v
}

func (o FailedPayee2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayeeId != nil {
		toSerialize["payeeId"] = o.PayeeId
	}
	if o.PayorRefs != nil {
		toSerialize["payorRefs"] = o.PayorRefs
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.RemoteId != nil {
		toSerialize["remoteId"] = o.RemoteId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.PaymentChannel != nil {
		toSerialize["paymentChannel"] = o.PaymentChannel
	}
	if o.Challenge != nil {
		toSerialize["challenge"] = o.Challenge
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.Individual != nil {
		toSerialize["individual"] = o.Individual
	}
	return json.Marshal(toSerialize)
}

type NullableFailedPayee2 struct {
	value *FailedPayee2
	isSet bool
}

func (v NullableFailedPayee2) Get() *FailedPayee2 {
	return v.value
}

func (v *NullableFailedPayee2) Set(val *FailedPayee2) {
	v.value = val
	v.isSet = true
}

func (v NullableFailedPayee2) IsSet() bool {
	return v.isSet
}

func (v *NullableFailedPayee2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailedPayee2(val *FailedPayee2) *NullableFailedPayee2 {
	return &NullableFailedPayee2{value: val, isSet: true}
}

func (v NullableFailedPayee2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailedPayee2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


