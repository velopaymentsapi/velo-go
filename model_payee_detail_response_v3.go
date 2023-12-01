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

// checks if the PayeeDetailResponseV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayeeDetailResponseV3{}

// PayeeDetailResponseV3 struct for PayeeDetailResponseV3
type PayeeDetailResponseV3 struct {
	PayeeId *string `json:"payeeId,omitempty"`
	PayorRefs []PayeePayorRefV3 `json:"payorRefs,omitempty"`
	Email NullableString `json:"email,omitempty"`
	// Onboarded status. One of the following values: CREATED, INVITED, REGISTERED, ONBOARDED
	OnboardedStatus *string `json:"onboardedStatus,omitempty"`
	// Current watchlist status. One of the following values: NONE, PENDING, REVIEW, PASSED, FAILED
	WatchlistStatus *string `json:"watchlistStatus,omitempty"`
	WatchlistOverrideExpiresAtTimestamp NullableTime `json:"watchlistOverrideExpiresAtTimestamp,omitempty"`
	WatchlistOverrideComment *string `json:"watchlistOverrideComment,omitempty"`
	// An IETF BCP 47 language code which has been configured for use within this Velo environment.<BR> See the /v1/supportedLanguages endpoint to list the available codes for an environment. 
	Language *string `json:"language,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	// Valid ISO 3166 2 character country code. See the <a href=\"https://www.iso.org/iso-3166-country-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	Country *string `json:"country,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	// Type of Payee. One of the following values: Individual, Company
	PayeeType *string `json:"payeeType,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	DisabledComment *string `json:"disabledComment,omitempty"`
	DisabledUpdatedTimestamp *time.Time `json:"disabledUpdatedTimestamp,omitempty"`
	Address *PayeeAddressV3 `json:"address,omitempty"`
	Individual *IndividualV3 `json:"individual,omitempty"`
	Company NullableCompanyV3 `json:"company,omitempty"`
	CellphoneNumber *string `json:"cellphoneNumber,omitempty"`
	WatchlistStatusUpdatedTimestamp NullableString `json:"watchlistStatusUpdatedTimestamp,omitempty"`
	GracePeriodEndDate NullableString `json:"gracePeriodEndDate,omitempty"`
	EnhancedKycCompleted *bool `json:"enhancedKycCompleted,omitempty"`
	KycCompletedTimestamp NullableString `json:"kycCompletedTimestamp,omitempty"`
	PausePayment *bool `json:"pausePayment,omitempty"`
	PausePaymentTimestamp NullableString `json:"pausePaymentTimestamp,omitempty"`
	MarketingOptInDecision *bool `json:"marketingOptInDecision,omitempty"`
	MarketingOptInTimestamp NullableString `json:"marketingOptInTimestamp,omitempty"`
	// The timestamp when the payee last accepted T&Cs
	AcceptTermsAndConditionsTimestamp NullableTime `json:"acceptTermsAndConditionsTimestamp,omitempty"`
	Challenge *ChallengeV3 `json:"challenge,omitempty"`
}

// NewPayeeDetailResponseV3 instantiates a new PayeeDetailResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayeeDetailResponseV3() *PayeeDetailResponseV3 {
	this := PayeeDetailResponseV3{}
	return &this
}

// NewPayeeDetailResponseV3WithDefaults instantiates a new PayeeDetailResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayeeDetailResponseV3WithDefaults() *PayeeDetailResponseV3 {
	this := PayeeDetailResponseV3{}
	return &this
}

// GetPayeeId returns the PayeeId field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetPayeeId() string {
	if o == nil || IsNil(o.PayeeId) {
		var ret string
		return ret
	}
	return *o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetPayeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeId) {
		return nil, false
	}
	return o.PayeeId, true
}

// HasPayeeId returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasPayeeId() bool {
	if o != nil && !IsNil(o.PayeeId) {
		return true
	}

	return false
}

// SetPayeeId gets a reference to the given string and assigns it to the PayeeId field.
func (o *PayeeDetailResponseV3) SetPayeeId(v string) {
	o.PayeeId = &v
}

// GetPayorRefs returns the PayorRefs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetPayorRefs() []PayeePayorRefV3 {
	if o == nil {
		var ret []PayeePayorRefV3
		return ret
	}
	return o.PayorRefs
}

// GetPayorRefsOk returns a tuple with the PayorRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetPayorRefsOk() ([]PayeePayorRefV3, bool) {
	if o == nil || IsNil(o.PayorRefs) {
		return nil, false
	}
	return o.PayorRefs, true
}

// HasPayorRefs returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasPayorRefs() bool {
	if o != nil && IsNil(o.PayorRefs) {
		return true
	}

	return false
}

// SetPayorRefs gets a reference to the given []PayeePayorRefV3 and assigns it to the PayorRefs field.
func (o *PayeeDetailResponseV3) SetPayorRefs(v []PayeePayorRefV3) {
	o.PayorRefs = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *PayeeDetailResponseV3) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *PayeeDetailResponseV3) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *PayeeDetailResponseV3) UnsetEmail() {
	o.Email.Unset()
}

// GetOnboardedStatus returns the OnboardedStatus field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetOnboardedStatus() string {
	if o == nil || IsNil(o.OnboardedStatus) {
		var ret string
		return ret
	}
	return *o.OnboardedStatus
}

// GetOnboardedStatusOk returns a tuple with the OnboardedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetOnboardedStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OnboardedStatus) {
		return nil, false
	}
	return o.OnboardedStatus, true
}

// HasOnboardedStatus returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasOnboardedStatus() bool {
	if o != nil && !IsNil(o.OnboardedStatus) {
		return true
	}

	return false
}

// SetOnboardedStatus gets a reference to the given string and assigns it to the OnboardedStatus field.
func (o *PayeeDetailResponseV3) SetOnboardedStatus(v string) {
	o.OnboardedStatus = &v
}

// GetWatchlistStatus returns the WatchlistStatus field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetWatchlistStatus() string {
	if o == nil || IsNil(o.WatchlistStatus) {
		var ret string
		return ret
	}
	return *o.WatchlistStatus
}

// GetWatchlistStatusOk returns a tuple with the WatchlistStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetWatchlistStatusOk() (*string, bool) {
	if o == nil || IsNil(o.WatchlistStatus) {
		return nil, false
	}
	return o.WatchlistStatus, true
}

// HasWatchlistStatus returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasWatchlistStatus() bool {
	if o != nil && !IsNil(o.WatchlistStatus) {
		return true
	}

	return false
}

// SetWatchlistStatus gets a reference to the given string and assigns it to the WatchlistStatus field.
func (o *PayeeDetailResponseV3) SetWatchlistStatus(v string) {
	o.WatchlistStatus = &v
}

// GetWatchlistOverrideExpiresAtTimestamp returns the WatchlistOverrideExpiresAtTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetWatchlistOverrideExpiresAtTimestamp() time.Time {
	if o == nil || IsNil(o.WatchlistOverrideExpiresAtTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.WatchlistOverrideExpiresAtTimestamp.Get()
}

// GetWatchlistOverrideExpiresAtTimestampOk returns a tuple with the WatchlistOverrideExpiresAtTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetWatchlistOverrideExpiresAtTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.WatchlistOverrideExpiresAtTimestamp.Get(), o.WatchlistOverrideExpiresAtTimestamp.IsSet()
}

// HasWatchlistOverrideExpiresAtTimestamp returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasWatchlistOverrideExpiresAtTimestamp() bool {
	if o != nil && o.WatchlistOverrideExpiresAtTimestamp.IsSet() {
		return true
	}

	return false
}

// SetWatchlistOverrideExpiresAtTimestamp gets a reference to the given NullableTime and assigns it to the WatchlistOverrideExpiresAtTimestamp field.
func (o *PayeeDetailResponseV3) SetWatchlistOverrideExpiresAtTimestamp(v time.Time) {
	o.WatchlistOverrideExpiresAtTimestamp.Set(&v)
}
// SetWatchlistOverrideExpiresAtTimestampNil sets the value for WatchlistOverrideExpiresAtTimestamp to be an explicit nil
func (o *PayeeDetailResponseV3) SetWatchlistOverrideExpiresAtTimestampNil() {
	o.WatchlistOverrideExpiresAtTimestamp.Set(nil)
}

// UnsetWatchlistOverrideExpiresAtTimestamp ensures that no value is present for WatchlistOverrideExpiresAtTimestamp, not even an explicit nil
func (o *PayeeDetailResponseV3) UnsetWatchlistOverrideExpiresAtTimestamp() {
	o.WatchlistOverrideExpiresAtTimestamp.Unset()
}

// GetWatchlistOverrideComment returns the WatchlistOverrideComment field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetWatchlistOverrideComment() string {
	if o == nil || IsNil(o.WatchlistOverrideComment) {
		var ret string
		return ret
	}
	return *o.WatchlistOverrideComment
}

// GetWatchlistOverrideCommentOk returns a tuple with the WatchlistOverrideComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetWatchlistOverrideCommentOk() (*string, bool) {
	if o == nil || IsNil(o.WatchlistOverrideComment) {
		return nil, false
	}
	return o.WatchlistOverrideComment, true
}

// HasWatchlistOverrideComment returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasWatchlistOverrideComment() bool {
	if o != nil && !IsNil(o.WatchlistOverrideComment) {
		return true
	}

	return false
}

// SetWatchlistOverrideComment gets a reference to the given string and assigns it to the WatchlistOverrideComment field.
func (o *PayeeDetailResponseV3) SetWatchlistOverrideComment(v string) {
	o.WatchlistOverrideComment = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *PayeeDetailResponseV3) SetLanguage(v string) {
	o.Language = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PayeeDetailResponseV3) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PayeeDetailResponseV3) SetCountry(v string) {
	o.Country = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PayeeDetailResponseV3) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPayeeType returns the PayeeType field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetPayeeType() string {
	if o == nil || IsNil(o.PayeeType) {
		var ret string
		return ret
	}
	return *o.PayeeType
}

// GetPayeeTypeOk returns a tuple with the PayeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetPayeeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeType) {
		return nil, false
	}
	return o.PayeeType, true
}

// HasPayeeType returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasPayeeType() bool {
	if o != nil && !IsNil(o.PayeeType) {
		return true
	}

	return false
}

// SetPayeeType gets a reference to the given string and assigns it to the PayeeType field.
func (o *PayeeDetailResponseV3) SetPayeeType(v string) {
	o.PayeeType = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *PayeeDetailResponseV3) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDisabledComment returns the DisabledComment field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetDisabledComment() string {
	if o == nil || IsNil(o.DisabledComment) {
		var ret string
		return ret
	}
	return *o.DisabledComment
}

// GetDisabledCommentOk returns a tuple with the DisabledComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetDisabledCommentOk() (*string, bool) {
	if o == nil || IsNil(o.DisabledComment) {
		return nil, false
	}
	return o.DisabledComment, true
}

// HasDisabledComment returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasDisabledComment() bool {
	if o != nil && !IsNil(o.DisabledComment) {
		return true
	}

	return false
}

// SetDisabledComment gets a reference to the given string and assigns it to the DisabledComment field.
func (o *PayeeDetailResponseV3) SetDisabledComment(v string) {
	o.DisabledComment = &v
}

// GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetDisabledUpdatedTimestamp() time.Time {
	if o == nil || IsNil(o.DisabledUpdatedTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.DisabledUpdatedTimestamp
}

// GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetDisabledUpdatedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DisabledUpdatedTimestamp) {
		return nil, false
	}
	return o.DisabledUpdatedTimestamp, true
}

// HasDisabledUpdatedTimestamp returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasDisabledUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.DisabledUpdatedTimestamp) {
		return true
	}

	return false
}

// SetDisabledUpdatedTimestamp gets a reference to the given time.Time and assigns it to the DisabledUpdatedTimestamp field.
func (o *PayeeDetailResponseV3) SetDisabledUpdatedTimestamp(v time.Time) {
	o.DisabledUpdatedTimestamp = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetAddress() PayeeAddressV3 {
	if o == nil || IsNil(o.Address) {
		var ret PayeeAddressV3
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetAddressOk() (*PayeeAddressV3, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given PayeeAddressV3 and assigns it to the Address field.
func (o *PayeeDetailResponseV3) SetAddress(v PayeeAddressV3) {
	o.Address = &v
}

// GetIndividual returns the Individual field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetIndividual() IndividualV3 {
	if o == nil || IsNil(o.Individual) {
		var ret IndividualV3
		return ret
	}
	return *o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetIndividualOk() (*IndividualV3, bool) {
	if o == nil || IsNil(o.Individual) {
		return nil, false
	}
	return o.Individual, true
}

// HasIndividual returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasIndividual() bool {
	if o != nil && !IsNil(o.Individual) {
		return true
	}

	return false
}

// SetIndividual gets a reference to the given IndividualV3 and assigns it to the Individual field.
func (o *PayeeDetailResponseV3) SetIndividual(v IndividualV3) {
	o.Individual = &v
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetCompany() CompanyV3 {
	if o == nil || IsNil(o.Company.Get()) {
		var ret CompanyV3
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetCompanyOk() (*CompanyV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableCompanyV3 and assigns it to the Company field.
func (o *PayeeDetailResponseV3) SetCompany(v CompanyV3) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *PayeeDetailResponseV3) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *PayeeDetailResponseV3) UnsetCompany() {
	o.Company.Unset()
}

// GetCellphoneNumber returns the CellphoneNumber field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetCellphoneNumber() string {
	if o == nil || IsNil(o.CellphoneNumber) {
		var ret string
		return ret
	}
	return *o.CellphoneNumber
}

// GetCellphoneNumberOk returns a tuple with the CellphoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetCellphoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CellphoneNumber) {
		return nil, false
	}
	return o.CellphoneNumber, true
}

// HasCellphoneNumber returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasCellphoneNumber() bool {
	if o != nil && !IsNil(o.CellphoneNumber) {
		return true
	}

	return false
}

// SetCellphoneNumber gets a reference to the given string and assigns it to the CellphoneNumber field.
func (o *PayeeDetailResponseV3) SetCellphoneNumber(v string) {
	o.CellphoneNumber = &v
}

// GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetWatchlistStatusUpdatedTimestamp() string {
	if o == nil || IsNil(o.WatchlistStatusUpdatedTimestamp.Get()) {
		var ret string
		return ret
	}
	return *o.WatchlistStatusUpdatedTimestamp.Get()
}

// GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetWatchlistStatusUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WatchlistStatusUpdatedTimestamp.Get(), o.WatchlistStatusUpdatedTimestamp.IsSet()
}

// HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasWatchlistStatusUpdatedTimestamp() bool {
	if o != nil && o.WatchlistStatusUpdatedTimestamp.IsSet() {
		return true
	}

	return false
}

// SetWatchlistStatusUpdatedTimestamp gets a reference to the given NullableString and assigns it to the WatchlistStatusUpdatedTimestamp field.
func (o *PayeeDetailResponseV3) SetWatchlistStatusUpdatedTimestamp(v string) {
	o.WatchlistStatusUpdatedTimestamp.Set(&v)
}
// SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil
func (o *PayeeDetailResponseV3) SetWatchlistStatusUpdatedTimestampNil() {
	o.WatchlistStatusUpdatedTimestamp.Set(nil)
}

// UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
func (o *PayeeDetailResponseV3) UnsetWatchlistStatusUpdatedTimestamp() {
	o.WatchlistStatusUpdatedTimestamp.Unset()
}

// GetGracePeriodEndDate returns the GracePeriodEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetGracePeriodEndDate() string {
	if o == nil || IsNil(o.GracePeriodEndDate.Get()) {
		var ret string
		return ret
	}
	return *o.GracePeriodEndDate.Get()
}

// GetGracePeriodEndDateOk returns a tuple with the GracePeriodEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetGracePeriodEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GracePeriodEndDate.Get(), o.GracePeriodEndDate.IsSet()
}

// HasGracePeriodEndDate returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasGracePeriodEndDate() bool {
	if o != nil && o.GracePeriodEndDate.IsSet() {
		return true
	}

	return false
}

// SetGracePeriodEndDate gets a reference to the given NullableString and assigns it to the GracePeriodEndDate field.
func (o *PayeeDetailResponseV3) SetGracePeriodEndDate(v string) {
	o.GracePeriodEndDate.Set(&v)
}
// SetGracePeriodEndDateNil sets the value for GracePeriodEndDate to be an explicit nil
func (o *PayeeDetailResponseV3) SetGracePeriodEndDateNil() {
	o.GracePeriodEndDate.Set(nil)
}

// UnsetGracePeriodEndDate ensures that no value is present for GracePeriodEndDate, not even an explicit nil
func (o *PayeeDetailResponseV3) UnsetGracePeriodEndDate() {
	o.GracePeriodEndDate.Unset()
}

// GetEnhancedKycCompleted returns the EnhancedKycCompleted field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetEnhancedKycCompleted() bool {
	if o == nil || IsNil(o.EnhancedKycCompleted) {
		var ret bool
		return ret
	}
	return *o.EnhancedKycCompleted
}

// GetEnhancedKycCompletedOk returns a tuple with the EnhancedKycCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetEnhancedKycCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.EnhancedKycCompleted) {
		return nil, false
	}
	return o.EnhancedKycCompleted, true
}

// HasEnhancedKycCompleted returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasEnhancedKycCompleted() bool {
	if o != nil && !IsNil(o.EnhancedKycCompleted) {
		return true
	}

	return false
}

// SetEnhancedKycCompleted gets a reference to the given bool and assigns it to the EnhancedKycCompleted field.
func (o *PayeeDetailResponseV3) SetEnhancedKycCompleted(v bool) {
	o.EnhancedKycCompleted = &v
}

// GetKycCompletedTimestamp returns the KycCompletedTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetKycCompletedTimestamp() string {
	if o == nil || IsNil(o.KycCompletedTimestamp.Get()) {
		var ret string
		return ret
	}
	return *o.KycCompletedTimestamp.Get()
}

// GetKycCompletedTimestampOk returns a tuple with the KycCompletedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetKycCompletedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KycCompletedTimestamp.Get(), o.KycCompletedTimestamp.IsSet()
}

// HasKycCompletedTimestamp returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasKycCompletedTimestamp() bool {
	if o != nil && o.KycCompletedTimestamp.IsSet() {
		return true
	}

	return false
}

// SetKycCompletedTimestamp gets a reference to the given NullableString and assigns it to the KycCompletedTimestamp field.
func (o *PayeeDetailResponseV3) SetKycCompletedTimestamp(v string) {
	o.KycCompletedTimestamp.Set(&v)
}
// SetKycCompletedTimestampNil sets the value for KycCompletedTimestamp to be an explicit nil
func (o *PayeeDetailResponseV3) SetKycCompletedTimestampNil() {
	o.KycCompletedTimestamp.Set(nil)
}

// UnsetKycCompletedTimestamp ensures that no value is present for KycCompletedTimestamp, not even an explicit nil
func (o *PayeeDetailResponseV3) UnsetKycCompletedTimestamp() {
	o.KycCompletedTimestamp.Unset()
}

// GetPausePayment returns the PausePayment field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetPausePayment() bool {
	if o == nil || IsNil(o.PausePayment) {
		var ret bool
		return ret
	}
	return *o.PausePayment
}

// GetPausePaymentOk returns a tuple with the PausePayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetPausePaymentOk() (*bool, bool) {
	if o == nil || IsNil(o.PausePayment) {
		return nil, false
	}
	return o.PausePayment, true
}

// HasPausePayment returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasPausePayment() bool {
	if o != nil && !IsNil(o.PausePayment) {
		return true
	}

	return false
}

// SetPausePayment gets a reference to the given bool and assigns it to the PausePayment field.
func (o *PayeeDetailResponseV3) SetPausePayment(v bool) {
	o.PausePayment = &v
}

// GetPausePaymentTimestamp returns the PausePaymentTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetPausePaymentTimestamp() string {
	if o == nil || IsNil(o.PausePaymentTimestamp.Get()) {
		var ret string
		return ret
	}
	return *o.PausePaymentTimestamp.Get()
}

// GetPausePaymentTimestampOk returns a tuple with the PausePaymentTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetPausePaymentTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PausePaymentTimestamp.Get(), o.PausePaymentTimestamp.IsSet()
}

// HasPausePaymentTimestamp returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasPausePaymentTimestamp() bool {
	if o != nil && o.PausePaymentTimestamp.IsSet() {
		return true
	}

	return false
}

// SetPausePaymentTimestamp gets a reference to the given NullableString and assigns it to the PausePaymentTimestamp field.
func (o *PayeeDetailResponseV3) SetPausePaymentTimestamp(v string) {
	o.PausePaymentTimestamp.Set(&v)
}
// SetPausePaymentTimestampNil sets the value for PausePaymentTimestamp to be an explicit nil
func (o *PayeeDetailResponseV3) SetPausePaymentTimestampNil() {
	o.PausePaymentTimestamp.Set(nil)
}

// UnsetPausePaymentTimestamp ensures that no value is present for PausePaymentTimestamp, not even an explicit nil
func (o *PayeeDetailResponseV3) UnsetPausePaymentTimestamp() {
	o.PausePaymentTimestamp.Unset()
}

// GetMarketingOptInDecision returns the MarketingOptInDecision field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetMarketingOptInDecision() bool {
	if o == nil || IsNil(o.MarketingOptInDecision) {
		var ret bool
		return ret
	}
	return *o.MarketingOptInDecision
}

// GetMarketingOptInDecisionOk returns a tuple with the MarketingOptInDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetMarketingOptInDecisionOk() (*bool, bool) {
	if o == nil || IsNil(o.MarketingOptInDecision) {
		return nil, false
	}
	return o.MarketingOptInDecision, true
}

// HasMarketingOptInDecision returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasMarketingOptInDecision() bool {
	if o != nil && !IsNil(o.MarketingOptInDecision) {
		return true
	}

	return false
}

// SetMarketingOptInDecision gets a reference to the given bool and assigns it to the MarketingOptInDecision field.
func (o *PayeeDetailResponseV3) SetMarketingOptInDecision(v bool) {
	o.MarketingOptInDecision = &v
}

// GetMarketingOptInTimestamp returns the MarketingOptInTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetMarketingOptInTimestamp() string {
	if o == nil || IsNil(o.MarketingOptInTimestamp.Get()) {
		var ret string
		return ret
	}
	return *o.MarketingOptInTimestamp.Get()
}

// GetMarketingOptInTimestampOk returns a tuple with the MarketingOptInTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetMarketingOptInTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketingOptInTimestamp.Get(), o.MarketingOptInTimestamp.IsSet()
}

// HasMarketingOptInTimestamp returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasMarketingOptInTimestamp() bool {
	if o != nil && o.MarketingOptInTimestamp.IsSet() {
		return true
	}

	return false
}

// SetMarketingOptInTimestamp gets a reference to the given NullableString and assigns it to the MarketingOptInTimestamp field.
func (o *PayeeDetailResponseV3) SetMarketingOptInTimestamp(v string) {
	o.MarketingOptInTimestamp.Set(&v)
}
// SetMarketingOptInTimestampNil sets the value for MarketingOptInTimestamp to be an explicit nil
func (o *PayeeDetailResponseV3) SetMarketingOptInTimestampNil() {
	o.MarketingOptInTimestamp.Set(nil)
}

// UnsetMarketingOptInTimestamp ensures that no value is present for MarketingOptInTimestamp, not even an explicit nil
func (o *PayeeDetailResponseV3) UnsetMarketingOptInTimestamp() {
	o.MarketingOptInTimestamp.Unset()
}

// GetAcceptTermsAndConditionsTimestamp returns the AcceptTermsAndConditionsTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayeeDetailResponseV3) GetAcceptTermsAndConditionsTimestamp() time.Time {
	if o == nil || IsNil(o.AcceptTermsAndConditionsTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AcceptTermsAndConditionsTimestamp.Get()
}

// GetAcceptTermsAndConditionsTimestampOk returns a tuple with the AcceptTermsAndConditionsTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayeeDetailResponseV3) GetAcceptTermsAndConditionsTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptTermsAndConditionsTimestamp.Get(), o.AcceptTermsAndConditionsTimestamp.IsSet()
}

// HasAcceptTermsAndConditionsTimestamp returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasAcceptTermsAndConditionsTimestamp() bool {
	if o != nil && o.AcceptTermsAndConditionsTimestamp.IsSet() {
		return true
	}

	return false
}

// SetAcceptTermsAndConditionsTimestamp gets a reference to the given NullableTime and assigns it to the AcceptTermsAndConditionsTimestamp field.
func (o *PayeeDetailResponseV3) SetAcceptTermsAndConditionsTimestamp(v time.Time) {
	o.AcceptTermsAndConditionsTimestamp.Set(&v)
}
// SetAcceptTermsAndConditionsTimestampNil sets the value for AcceptTermsAndConditionsTimestamp to be an explicit nil
func (o *PayeeDetailResponseV3) SetAcceptTermsAndConditionsTimestampNil() {
	o.AcceptTermsAndConditionsTimestamp.Set(nil)
}

// UnsetAcceptTermsAndConditionsTimestamp ensures that no value is present for AcceptTermsAndConditionsTimestamp, not even an explicit nil
func (o *PayeeDetailResponseV3) UnsetAcceptTermsAndConditionsTimestamp() {
	o.AcceptTermsAndConditionsTimestamp.Unset()
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *PayeeDetailResponseV3) GetChallenge() ChallengeV3 {
	if o == nil || IsNil(o.Challenge) {
		var ret ChallengeV3
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayeeDetailResponseV3) GetChallengeOk() (*ChallengeV3, bool) {
	if o == nil || IsNil(o.Challenge) {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *PayeeDetailResponseV3) HasChallenge() bool {
	if o != nil && !IsNil(o.Challenge) {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given ChallengeV3 and assigns it to the Challenge field.
func (o *PayeeDetailResponseV3) SetChallenge(v ChallengeV3) {
	o.Challenge = &v
}

func (o PayeeDetailResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayeeDetailResponseV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayeeId) {
		toSerialize["payeeId"] = o.PayeeId
	}
	if o.PayorRefs != nil {
		toSerialize["payorRefs"] = o.PayorRefs
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.OnboardedStatus) {
		toSerialize["onboardedStatus"] = o.OnboardedStatus
	}
	if !IsNil(o.WatchlistStatus) {
		toSerialize["watchlistStatus"] = o.WatchlistStatus
	}
	if o.WatchlistOverrideExpiresAtTimestamp.IsSet() {
		toSerialize["watchlistOverrideExpiresAtTimestamp"] = o.WatchlistOverrideExpiresAtTimestamp.Get()
	}
	if !IsNil(o.WatchlistOverrideComment) {
		toSerialize["watchlistOverrideComment"] = o.WatchlistOverrideComment
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.PayeeType) {
		toSerialize["payeeType"] = o.PayeeType
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.DisabledComment) {
		toSerialize["disabledComment"] = o.DisabledComment
	}
	if !IsNil(o.DisabledUpdatedTimestamp) {
		toSerialize["disabledUpdatedTimestamp"] = o.DisabledUpdatedTimestamp
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Individual) {
		toSerialize["individual"] = o.Individual
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if !IsNil(o.CellphoneNumber) {
		toSerialize["cellphoneNumber"] = o.CellphoneNumber
	}
	if o.WatchlistStatusUpdatedTimestamp.IsSet() {
		toSerialize["watchlistStatusUpdatedTimestamp"] = o.WatchlistStatusUpdatedTimestamp.Get()
	}
	if o.GracePeriodEndDate.IsSet() {
		toSerialize["gracePeriodEndDate"] = o.GracePeriodEndDate.Get()
	}
	if !IsNil(o.EnhancedKycCompleted) {
		toSerialize["enhancedKycCompleted"] = o.EnhancedKycCompleted
	}
	if o.KycCompletedTimestamp.IsSet() {
		toSerialize["kycCompletedTimestamp"] = o.KycCompletedTimestamp.Get()
	}
	if !IsNil(o.PausePayment) {
		toSerialize["pausePayment"] = o.PausePayment
	}
	if o.PausePaymentTimestamp.IsSet() {
		toSerialize["pausePaymentTimestamp"] = o.PausePaymentTimestamp.Get()
	}
	if !IsNil(o.MarketingOptInDecision) {
		toSerialize["marketingOptInDecision"] = o.MarketingOptInDecision
	}
	if o.MarketingOptInTimestamp.IsSet() {
		toSerialize["marketingOptInTimestamp"] = o.MarketingOptInTimestamp.Get()
	}
	if o.AcceptTermsAndConditionsTimestamp.IsSet() {
		toSerialize["acceptTermsAndConditionsTimestamp"] = o.AcceptTermsAndConditionsTimestamp.Get()
	}
	if !IsNil(o.Challenge) {
		toSerialize["challenge"] = o.Challenge
	}
	return toSerialize, nil
}

type NullablePayeeDetailResponseV3 struct {
	value *PayeeDetailResponseV3
	isSet bool
}

func (v NullablePayeeDetailResponseV3) Get() *PayeeDetailResponseV3 {
	return v.value
}

func (v *NullablePayeeDetailResponseV3) Set(val *PayeeDetailResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullablePayeeDetailResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullablePayeeDetailResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayeeDetailResponseV3(val *PayeeDetailResponseV3) *NullablePayeeDetailResponseV3 {
	return &NullablePayeeDetailResponseV3{value: val, isSet: true}
}

func (v NullablePayeeDetailResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayeeDetailResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


