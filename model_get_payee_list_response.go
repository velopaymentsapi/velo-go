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
	"time"
)

// GetPayeeListResponse struct for GetPayeeListResponse
type GetPayeeListResponse struct {
	PayeeId *string `json:"payeeId,omitempty"`
	PayorRefs []PayeePayorRefV3 `json:"payorRefs,omitempty"`
	Email NullableString `json:"email,omitempty"`
	OnboardedStatus *OnboardedStatus2 `json:"onboardedStatus,omitempty"`
	WatchlistStatus *WatchlistStatus `json:"watchlistStatus,omitempty"`
	WatchlistStatusUpdatedTimestamp NullableString `json:"watchlistStatusUpdatedTimestamp,omitempty"`
	WatchlistOverrideComment NullableString `json:"watchlistOverrideComment,omitempty"`
	// An IETF BCP 47 language code which has been configured for use within this Velo environment.<BR> See the /v1/supportedLanguages endpoint to list the available codes for an environment. 
	Language *string `json:"language,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Country *string `json:"country,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	PayeeType *PayeeType `json:"payeeType,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	DisabledComment *string `json:"disabledComment,omitempty"`
	DisabledUpdatedTimestamp *time.Time `json:"disabledUpdatedTimestamp,omitempty"`
	Individual *GetPayeeListResponseIndividual `json:"individual,omitempty"`
	Company *GetPayeeListResponseCompany `json:"company,omitempty"`
}

// NewGetPayeeListResponse instantiates a new GetPayeeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayeeListResponse() *GetPayeeListResponse {
	this := GetPayeeListResponse{}
	return &this
}

// NewGetPayeeListResponseWithDefaults instantiates a new GetPayeeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayeeListResponseWithDefaults() *GetPayeeListResponse {
	this := GetPayeeListResponse{}
	return &this
}

// GetPayeeId returns the PayeeId field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetPayeeId() string {
	if o == nil || o.PayeeId == nil {
		var ret string
		return ret
	}
	return *o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetPayeeIdOk() (*string, bool) {
	if o == nil || o.PayeeId == nil {
		return nil, false
	}
	return o.PayeeId, true
}

// HasPayeeId returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasPayeeId() bool {
	if o != nil && o.PayeeId != nil {
		return true
	}

	return false
}

// SetPayeeId gets a reference to the given string and assigns it to the PayeeId field.
func (o *GetPayeeListResponse) SetPayeeId(v string) {
	o.PayeeId = &v
}

// GetPayorRefs returns the PayorRefs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPayeeListResponse) GetPayorRefs() []PayeePayorRefV3 {
	if o == nil  {
		var ret []PayeePayorRefV3
		return ret
	}
	return o.PayorRefs
}

// GetPayorRefsOk returns a tuple with the PayorRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPayeeListResponse) GetPayorRefsOk() (*[]PayeePayorRefV3, bool) {
	if o == nil || o.PayorRefs == nil {
		return nil, false
	}
	return &o.PayorRefs, true
}

// HasPayorRefs returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasPayorRefs() bool {
	if o != nil && o.PayorRefs != nil {
		return true
	}

	return false
}

// SetPayorRefs gets a reference to the given []PayeePayorRefV3 and assigns it to the PayorRefs field.
func (o *GetPayeeListResponse) SetPayorRefs(v []PayeePayorRefV3) {
	o.PayorRefs = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPayeeListResponse) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPayeeListResponse) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *GetPayeeListResponse) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *GetPayeeListResponse) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *GetPayeeListResponse) UnsetEmail() {
	o.Email.Unset()
}

// GetOnboardedStatus returns the OnboardedStatus field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetOnboardedStatus() OnboardedStatus2 {
	if o == nil || o.OnboardedStatus == nil {
		var ret OnboardedStatus2
		return ret
	}
	return *o.OnboardedStatus
}

// GetOnboardedStatusOk returns a tuple with the OnboardedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetOnboardedStatusOk() (*OnboardedStatus2, bool) {
	if o == nil || o.OnboardedStatus == nil {
		return nil, false
	}
	return o.OnboardedStatus, true
}

// HasOnboardedStatus returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasOnboardedStatus() bool {
	if o != nil && o.OnboardedStatus != nil {
		return true
	}

	return false
}

// SetOnboardedStatus gets a reference to the given OnboardedStatus2 and assigns it to the OnboardedStatus field.
func (o *GetPayeeListResponse) SetOnboardedStatus(v OnboardedStatus2) {
	o.OnboardedStatus = &v
}

// GetWatchlistStatus returns the WatchlistStatus field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetWatchlistStatus() WatchlistStatus {
	if o == nil || o.WatchlistStatus == nil {
		var ret WatchlistStatus
		return ret
	}
	return *o.WatchlistStatus
}

// GetWatchlistStatusOk returns a tuple with the WatchlistStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetWatchlistStatusOk() (*WatchlistStatus, bool) {
	if o == nil || o.WatchlistStatus == nil {
		return nil, false
	}
	return o.WatchlistStatus, true
}

// HasWatchlistStatus returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasWatchlistStatus() bool {
	if o != nil && o.WatchlistStatus != nil {
		return true
	}

	return false
}

// SetWatchlistStatus gets a reference to the given WatchlistStatus and assigns it to the WatchlistStatus field.
func (o *GetPayeeListResponse) SetWatchlistStatus(v WatchlistStatus) {
	o.WatchlistStatus = &v
}

// GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPayeeListResponse) GetWatchlistStatusUpdatedTimestamp() string {
	if o == nil || o.WatchlistStatusUpdatedTimestamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.WatchlistStatusUpdatedTimestamp.Get()
}

// GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPayeeListResponse) GetWatchlistStatusUpdatedTimestampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WatchlistStatusUpdatedTimestamp.Get(), o.WatchlistStatusUpdatedTimestamp.IsSet()
}

// HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasWatchlistStatusUpdatedTimestamp() bool {
	if o != nil && o.WatchlistStatusUpdatedTimestamp.IsSet() {
		return true
	}

	return false
}

// SetWatchlistStatusUpdatedTimestamp gets a reference to the given NullableString and assigns it to the WatchlistStatusUpdatedTimestamp field.
func (o *GetPayeeListResponse) SetWatchlistStatusUpdatedTimestamp(v string) {
	o.WatchlistStatusUpdatedTimestamp.Set(&v)
}
// SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil
func (o *GetPayeeListResponse) SetWatchlistStatusUpdatedTimestampNil() {
	o.WatchlistStatusUpdatedTimestamp.Set(nil)
}

// UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
func (o *GetPayeeListResponse) UnsetWatchlistStatusUpdatedTimestamp() {
	o.WatchlistStatusUpdatedTimestamp.Unset()
}

// GetWatchlistOverrideComment returns the WatchlistOverrideComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPayeeListResponse) GetWatchlistOverrideComment() string {
	if o == nil || o.WatchlistOverrideComment.Get() == nil {
		var ret string
		return ret
	}
	return *o.WatchlistOverrideComment.Get()
}

// GetWatchlistOverrideCommentOk returns a tuple with the WatchlistOverrideComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPayeeListResponse) GetWatchlistOverrideCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WatchlistOverrideComment.Get(), o.WatchlistOverrideComment.IsSet()
}

// HasWatchlistOverrideComment returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasWatchlistOverrideComment() bool {
	if o != nil && o.WatchlistOverrideComment.IsSet() {
		return true
	}

	return false
}

// SetWatchlistOverrideComment gets a reference to the given NullableString and assigns it to the WatchlistOverrideComment field.
func (o *GetPayeeListResponse) SetWatchlistOverrideComment(v string) {
	o.WatchlistOverrideComment.Set(&v)
}
// SetWatchlistOverrideCommentNil sets the value for WatchlistOverrideComment to be an explicit nil
func (o *GetPayeeListResponse) SetWatchlistOverrideCommentNil() {
	o.WatchlistOverrideComment.Set(nil)
}

// UnsetWatchlistOverrideComment ensures that no value is present for WatchlistOverrideComment, not even an explicit nil
func (o *GetPayeeListResponse) UnsetWatchlistOverrideComment() {
	o.WatchlistOverrideComment.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *GetPayeeListResponse) SetLanguage(v string) {
	o.Language = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetPayeeListResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GetPayeeListResponse) SetCountry(v string) {
	o.Country = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GetPayeeListResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPayeeType returns the PayeeType field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetPayeeType() PayeeType {
	if o == nil || o.PayeeType == nil {
		var ret PayeeType
		return ret
	}
	return *o.PayeeType
}

// GetPayeeTypeOk returns a tuple with the PayeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetPayeeTypeOk() (*PayeeType, bool) {
	if o == nil || o.PayeeType == nil {
		return nil, false
	}
	return o.PayeeType, true
}

// HasPayeeType returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasPayeeType() bool {
	if o != nil && o.PayeeType != nil {
		return true
	}

	return false
}

// SetPayeeType gets a reference to the given PayeeType and assigns it to the PayeeType field.
func (o *GetPayeeListResponse) SetPayeeType(v PayeeType) {
	o.PayeeType = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *GetPayeeListResponse) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDisabledComment returns the DisabledComment field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetDisabledComment() string {
	if o == nil || o.DisabledComment == nil {
		var ret string
		return ret
	}
	return *o.DisabledComment
}

// GetDisabledCommentOk returns a tuple with the DisabledComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetDisabledCommentOk() (*string, bool) {
	if o == nil || o.DisabledComment == nil {
		return nil, false
	}
	return o.DisabledComment, true
}

// HasDisabledComment returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasDisabledComment() bool {
	if o != nil && o.DisabledComment != nil {
		return true
	}

	return false
}

// SetDisabledComment gets a reference to the given string and assigns it to the DisabledComment field.
func (o *GetPayeeListResponse) SetDisabledComment(v string) {
	o.DisabledComment = &v
}

// GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetDisabledUpdatedTimestamp() time.Time {
	if o == nil || o.DisabledUpdatedTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.DisabledUpdatedTimestamp
}

// GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetDisabledUpdatedTimestampOk() (*time.Time, bool) {
	if o == nil || o.DisabledUpdatedTimestamp == nil {
		return nil, false
	}
	return o.DisabledUpdatedTimestamp, true
}

// HasDisabledUpdatedTimestamp returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasDisabledUpdatedTimestamp() bool {
	if o != nil && o.DisabledUpdatedTimestamp != nil {
		return true
	}

	return false
}

// SetDisabledUpdatedTimestamp gets a reference to the given time.Time and assigns it to the DisabledUpdatedTimestamp field.
func (o *GetPayeeListResponse) SetDisabledUpdatedTimestamp(v time.Time) {
	o.DisabledUpdatedTimestamp = &v
}

// GetIndividual returns the Individual field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetIndividual() GetPayeeListResponseIndividual {
	if o == nil || o.Individual == nil {
		var ret GetPayeeListResponseIndividual
		return ret
	}
	return *o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetIndividualOk() (*GetPayeeListResponseIndividual, bool) {
	if o == nil || o.Individual == nil {
		return nil, false
	}
	return o.Individual, true
}

// HasIndividual returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasIndividual() bool {
	if o != nil && o.Individual != nil {
		return true
	}

	return false
}

// SetIndividual gets a reference to the given GetPayeeListResponseIndividual and assigns it to the Individual field.
func (o *GetPayeeListResponse) SetIndividual(v GetPayeeListResponseIndividual) {
	o.Individual = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *GetPayeeListResponse) GetCompany() GetPayeeListResponseCompany {
	if o == nil || o.Company == nil {
		var ret GetPayeeListResponseCompany
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponse) GetCompanyOk() (*GetPayeeListResponseCompany, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *GetPayeeListResponse) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given GetPayeeListResponseCompany and assigns it to the Company field.
func (o *GetPayeeListResponse) SetCompany(v GetPayeeListResponseCompany) {
	o.Company = &v
}

func (o GetPayeeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayeeId != nil {
		toSerialize["payeeId"] = o.PayeeId
	}
	if o.PayorRefs != nil {
		toSerialize["payorRefs"] = o.PayorRefs
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.OnboardedStatus != nil {
		toSerialize["onboardedStatus"] = o.OnboardedStatus
	}
	if o.WatchlistStatus != nil {
		toSerialize["watchlistStatus"] = o.WatchlistStatus
	}
	if o.WatchlistStatusUpdatedTimestamp.IsSet() {
		toSerialize["watchlistStatusUpdatedTimestamp"] = o.WatchlistStatusUpdatedTimestamp.Get()
	}
	if o.WatchlistOverrideComment.IsSet() {
		toSerialize["watchlistOverrideComment"] = o.WatchlistOverrideComment.Get()
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.PayeeType != nil {
		toSerialize["payeeType"] = o.PayeeType
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.DisabledComment != nil {
		toSerialize["disabledComment"] = o.DisabledComment
	}
	if o.DisabledUpdatedTimestamp != nil {
		toSerialize["disabledUpdatedTimestamp"] = o.DisabledUpdatedTimestamp
	}
	if o.Individual != nil {
		toSerialize["individual"] = o.Individual
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	return json.Marshal(toSerialize)
}

type NullableGetPayeeListResponse struct {
	value *GetPayeeListResponse
	isSet bool
}

func (v NullableGetPayeeListResponse) Get() *GetPayeeListResponse {
	return v.value
}

func (v *NullableGetPayeeListResponse) Set(val *GetPayeeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayeeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayeeListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayeeListResponse(val *GetPayeeListResponse) *NullableGetPayeeListResponse {
	return &NullableGetPayeeListResponse{value: val, isSet: true}
}

func (v NullableGetPayeeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayeeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


