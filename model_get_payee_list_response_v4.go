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

// checks if the GetPayeeListResponseV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPayeeListResponseV4{}

// GetPayeeListResponseV4 struct for GetPayeeListResponseV4
type GetPayeeListResponseV4 struct {
	PayeeId *string `json:"payeeId,omitempty"`
	PayorRefs []PayeePayorRefV4 `json:"payorRefs,omitempty"`
	Email NullableString `json:"email,omitempty"`
	// Payee onboarded status. One of the following values: CREATED, INVITED, REGISTERED, ONBOARDED
	OnboardedStatus *string `json:"onboardedStatus,omitempty"`
	// Current watchlist status. One of the following values: NONE, PENDING, REVIEW, PASSED, FAILED
	WatchlistStatus *string `json:"watchlistStatus,omitempty"`
	WatchlistStatusUpdatedTimestamp NullableString `json:"watchlistStatusUpdatedTimestamp,omitempty"`
	WatchlistOverrideComment NullableString `json:"watchlistOverrideComment,omitempty"`
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
	// The id of the payor if the payee is managed
	ManagedByPayorId *string `json:"managedByPayorId,omitempty"`
	Individual *GetPayeeListResponseIndividualV4 `json:"individual,omitempty"`
	Company *GetPayeeListResponseCompanyV4 `json:"company,omitempty"`
}

// NewGetPayeeListResponseV4 instantiates a new GetPayeeListResponseV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayeeListResponseV4() *GetPayeeListResponseV4 {
	this := GetPayeeListResponseV4{}
	return &this
}

// NewGetPayeeListResponseV4WithDefaults instantiates a new GetPayeeListResponseV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayeeListResponseV4WithDefaults() *GetPayeeListResponseV4 {
	this := GetPayeeListResponseV4{}
	return &this
}

// GetPayeeId returns the PayeeId field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetPayeeId() string {
	if o == nil || IsNil(o.PayeeId) {
		var ret string
		return ret
	}
	return *o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetPayeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeId) {
		return nil, false
	}
	return o.PayeeId, true
}

// HasPayeeId returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasPayeeId() bool {
	if o != nil && !IsNil(o.PayeeId) {
		return true
	}

	return false
}

// SetPayeeId gets a reference to the given string and assigns it to the PayeeId field.
func (o *GetPayeeListResponseV4) SetPayeeId(v string) {
	o.PayeeId = &v
}

// GetPayorRefs returns the PayorRefs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPayeeListResponseV4) GetPayorRefs() []PayeePayorRefV4 {
	if o == nil {
		var ret []PayeePayorRefV4
		return ret
	}
	return o.PayorRefs
}

// GetPayorRefsOk returns a tuple with the PayorRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPayeeListResponseV4) GetPayorRefsOk() ([]PayeePayorRefV4, bool) {
	if o == nil || IsNil(o.PayorRefs) {
		return nil, false
	}
	return o.PayorRefs, true
}

// HasPayorRefs returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasPayorRefs() bool {
	if o != nil && IsNil(o.PayorRefs) {
		return true
	}

	return false
}

// SetPayorRefs gets a reference to the given []PayeePayorRefV4 and assigns it to the PayorRefs field.
func (o *GetPayeeListResponseV4) SetPayorRefs(v []PayeePayorRefV4) {
	o.PayorRefs = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPayeeListResponseV4) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPayeeListResponseV4) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *GetPayeeListResponseV4) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *GetPayeeListResponseV4) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *GetPayeeListResponseV4) UnsetEmail() {
	o.Email.Unset()
}

// GetOnboardedStatus returns the OnboardedStatus field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetOnboardedStatus() string {
	if o == nil || IsNil(o.OnboardedStatus) {
		var ret string
		return ret
	}
	return *o.OnboardedStatus
}

// GetOnboardedStatusOk returns a tuple with the OnboardedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetOnboardedStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OnboardedStatus) {
		return nil, false
	}
	return o.OnboardedStatus, true
}

// HasOnboardedStatus returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasOnboardedStatus() bool {
	if o != nil && !IsNil(o.OnboardedStatus) {
		return true
	}

	return false
}

// SetOnboardedStatus gets a reference to the given string and assigns it to the OnboardedStatus field.
func (o *GetPayeeListResponseV4) SetOnboardedStatus(v string) {
	o.OnboardedStatus = &v
}

// GetWatchlistStatus returns the WatchlistStatus field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetWatchlistStatus() string {
	if o == nil || IsNil(o.WatchlistStatus) {
		var ret string
		return ret
	}
	return *o.WatchlistStatus
}

// GetWatchlistStatusOk returns a tuple with the WatchlistStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetWatchlistStatusOk() (*string, bool) {
	if o == nil || IsNil(o.WatchlistStatus) {
		return nil, false
	}
	return o.WatchlistStatus, true
}

// HasWatchlistStatus returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasWatchlistStatus() bool {
	if o != nil && !IsNil(o.WatchlistStatus) {
		return true
	}

	return false
}

// SetWatchlistStatus gets a reference to the given string and assigns it to the WatchlistStatus field.
func (o *GetPayeeListResponseV4) SetWatchlistStatus(v string) {
	o.WatchlistStatus = &v
}

// GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPayeeListResponseV4) GetWatchlistStatusUpdatedTimestamp() string {
	if o == nil || IsNil(o.WatchlistStatusUpdatedTimestamp.Get()) {
		var ret string
		return ret
	}
	return *o.WatchlistStatusUpdatedTimestamp.Get()
}

// GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPayeeListResponseV4) GetWatchlistStatusUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WatchlistStatusUpdatedTimestamp.Get(), o.WatchlistStatusUpdatedTimestamp.IsSet()
}

// HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasWatchlistStatusUpdatedTimestamp() bool {
	if o != nil && o.WatchlistStatusUpdatedTimestamp.IsSet() {
		return true
	}

	return false
}

// SetWatchlistStatusUpdatedTimestamp gets a reference to the given NullableString and assigns it to the WatchlistStatusUpdatedTimestamp field.
func (o *GetPayeeListResponseV4) SetWatchlistStatusUpdatedTimestamp(v string) {
	o.WatchlistStatusUpdatedTimestamp.Set(&v)
}
// SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil
func (o *GetPayeeListResponseV4) SetWatchlistStatusUpdatedTimestampNil() {
	o.WatchlistStatusUpdatedTimestamp.Set(nil)
}

// UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
func (o *GetPayeeListResponseV4) UnsetWatchlistStatusUpdatedTimestamp() {
	o.WatchlistStatusUpdatedTimestamp.Unset()
}

// GetWatchlistOverrideComment returns the WatchlistOverrideComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPayeeListResponseV4) GetWatchlistOverrideComment() string {
	if o == nil || IsNil(o.WatchlistOverrideComment.Get()) {
		var ret string
		return ret
	}
	return *o.WatchlistOverrideComment.Get()
}

// GetWatchlistOverrideCommentOk returns a tuple with the WatchlistOverrideComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPayeeListResponseV4) GetWatchlistOverrideCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WatchlistOverrideComment.Get(), o.WatchlistOverrideComment.IsSet()
}

// HasWatchlistOverrideComment returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasWatchlistOverrideComment() bool {
	if o != nil && o.WatchlistOverrideComment.IsSet() {
		return true
	}

	return false
}

// SetWatchlistOverrideComment gets a reference to the given NullableString and assigns it to the WatchlistOverrideComment field.
func (o *GetPayeeListResponseV4) SetWatchlistOverrideComment(v string) {
	o.WatchlistOverrideComment.Set(&v)
}
// SetWatchlistOverrideCommentNil sets the value for WatchlistOverrideComment to be an explicit nil
func (o *GetPayeeListResponseV4) SetWatchlistOverrideCommentNil() {
	o.WatchlistOverrideComment.Set(nil)
}

// UnsetWatchlistOverrideComment ensures that no value is present for WatchlistOverrideComment, not even an explicit nil
func (o *GetPayeeListResponseV4) UnsetWatchlistOverrideComment() {
	o.WatchlistOverrideComment.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *GetPayeeListResponseV4) SetLanguage(v string) {
	o.Language = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetPayeeListResponseV4) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GetPayeeListResponseV4) SetCountry(v string) {
	o.Country = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GetPayeeListResponseV4) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPayeeType returns the PayeeType field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetPayeeType() string {
	if o == nil || IsNil(o.PayeeType) {
		var ret string
		return ret
	}
	return *o.PayeeType
}

// GetPayeeTypeOk returns a tuple with the PayeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetPayeeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeType) {
		return nil, false
	}
	return o.PayeeType, true
}

// HasPayeeType returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasPayeeType() bool {
	if o != nil && !IsNil(o.PayeeType) {
		return true
	}

	return false
}

// SetPayeeType gets a reference to the given string and assigns it to the PayeeType field.
func (o *GetPayeeListResponseV4) SetPayeeType(v string) {
	o.PayeeType = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *GetPayeeListResponseV4) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDisabledComment returns the DisabledComment field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetDisabledComment() string {
	if o == nil || IsNil(o.DisabledComment) {
		var ret string
		return ret
	}
	return *o.DisabledComment
}

// GetDisabledCommentOk returns a tuple with the DisabledComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetDisabledCommentOk() (*string, bool) {
	if o == nil || IsNil(o.DisabledComment) {
		return nil, false
	}
	return o.DisabledComment, true
}

// HasDisabledComment returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasDisabledComment() bool {
	if o != nil && !IsNil(o.DisabledComment) {
		return true
	}

	return false
}

// SetDisabledComment gets a reference to the given string and assigns it to the DisabledComment field.
func (o *GetPayeeListResponseV4) SetDisabledComment(v string) {
	o.DisabledComment = &v
}

// GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetDisabledUpdatedTimestamp() time.Time {
	if o == nil || IsNil(o.DisabledUpdatedTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.DisabledUpdatedTimestamp
}

// GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetDisabledUpdatedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DisabledUpdatedTimestamp) {
		return nil, false
	}
	return o.DisabledUpdatedTimestamp, true
}

// HasDisabledUpdatedTimestamp returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasDisabledUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.DisabledUpdatedTimestamp) {
		return true
	}

	return false
}

// SetDisabledUpdatedTimestamp gets a reference to the given time.Time and assigns it to the DisabledUpdatedTimestamp field.
func (o *GetPayeeListResponseV4) SetDisabledUpdatedTimestamp(v time.Time) {
	o.DisabledUpdatedTimestamp = &v
}

// GetManagedByPayorId returns the ManagedByPayorId field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetManagedByPayorId() string {
	if o == nil || IsNil(o.ManagedByPayorId) {
		var ret string
		return ret
	}
	return *o.ManagedByPayorId
}

// GetManagedByPayorIdOk returns a tuple with the ManagedByPayorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetManagedByPayorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedByPayorId) {
		return nil, false
	}
	return o.ManagedByPayorId, true
}

// HasManagedByPayorId returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasManagedByPayorId() bool {
	if o != nil && !IsNil(o.ManagedByPayorId) {
		return true
	}

	return false
}

// SetManagedByPayorId gets a reference to the given string and assigns it to the ManagedByPayorId field.
func (o *GetPayeeListResponseV4) SetManagedByPayorId(v string) {
	o.ManagedByPayorId = &v
}

// GetIndividual returns the Individual field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetIndividual() GetPayeeListResponseIndividualV4 {
	if o == nil || IsNil(o.Individual) {
		var ret GetPayeeListResponseIndividualV4
		return ret
	}
	return *o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetIndividualOk() (*GetPayeeListResponseIndividualV4, bool) {
	if o == nil || IsNil(o.Individual) {
		return nil, false
	}
	return o.Individual, true
}

// HasIndividual returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasIndividual() bool {
	if o != nil && !IsNil(o.Individual) {
		return true
	}

	return false
}

// SetIndividual gets a reference to the given GetPayeeListResponseIndividualV4 and assigns it to the Individual field.
func (o *GetPayeeListResponseV4) SetIndividual(v GetPayeeListResponseIndividualV4) {
	o.Individual = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *GetPayeeListResponseV4) GetCompany() GetPayeeListResponseCompanyV4 {
	if o == nil || IsNil(o.Company) {
		var ret GetPayeeListResponseCompanyV4
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayeeListResponseV4) GetCompanyOk() (*GetPayeeListResponseCompanyV4, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *GetPayeeListResponseV4) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given GetPayeeListResponseCompanyV4 and assigns it to the Company field.
func (o *GetPayeeListResponseV4) SetCompany(v GetPayeeListResponseCompanyV4) {
	o.Company = &v
}

func (o GetPayeeListResponseV4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayeeListResponseV4) ToMap() (map[string]interface{}, error) {
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
	if o.WatchlistStatusUpdatedTimestamp.IsSet() {
		toSerialize["watchlistStatusUpdatedTimestamp"] = o.WatchlistStatusUpdatedTimestamp.Get()
	}
	if o.WatchlistOverrideComment.IsSet() {
		toSerialize["watchlistOverrideComment"] = o.WatchlistOverrideComment.Get()
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
	if !IsNil(o.ManagedByPayorId) {
		toSerialize["managedByPayorId"] = o.ManagedByPayorId
	}
	if !IsNil(o.Individual) {
		toSerialize["individual"] = o.Individual
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	return toSerialize, nil
}

type NullableGetPayeeListResponseV4 struct {
	value *GetPayeeListResponseV4
	isSet bool
}

func (v NullableGetPayeeListResponseV4) Get() *GetPayeeListResponseV4 {
	return v.value
}

func (v *NullableGetPayeeListResponseV4) Set(val *GetPayeeListResponseV4) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayeeListResponseV4) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayeeListResponseV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayeeListResponseV4(val *GetPayeeListResponseV4) *NullableGetPayeeListResponseV4 {
	return &NullableGetPayeeListResponseV4{value: val, isSet: true}
}

func (v NullableGetPayeeListResponseV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayeeListResponseV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


