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

// checks if the PayorV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayorV2{}

// PayorV2 struct for PayorV2
type PayorV2 struct {
	// The Payor Id
	PayorId string `json:"payorId"`
	// The name of the payor
	PayorName string `json:"payorName"`
	// A unique identifier that an external system uses to reference the payor in their system
	PayorXid *string `json:"payorXid,omitempty"`
	// The source of the payorXid, default is null which means Velo
	Provider *string `json:"provider,omitempty"`
	Address *PayorAddressV2 `json:"address,omitempty"`
	// Name of primary contact for the payor.
	PrimaryContactName *string `json:"primaryContactName,omitempty"`
	// Primary contact phone number for the payor.
	PrimaryContactPhone *string `json:"primaryContactPhone,omitempty"`
	// Primary contact email for the payor.
	PrimaryContactEmail *string `json:"primaryContactEmail,omitempty"`
	// The kyc state of the payor. One of the following values: FAILED_KYC, PASSED_KYC, REQUIRES_KYC
	KycState *string `json:"kycState,omitempty"`
	// Whether or not the payor has been manually locked by the backoffice.
	ManualLockout *bool `json:"manualLockout,omitempty"`
	// Is Open Banking supported for this payor
	OpenBankingEnabled *bool `json:"openBankingEnabled,omitempty"`
	// Whether grace period processing is enabled.
	PayeeGracePeriodProcessingEnabled *bool `json:"payeeGracePeriodProcessingEnabled,omitempty"`
	// The grace period for paying payees in days before the payee must be onboarded.
	PayeeGracePeriodDays *int32 `json:"payeeGracePeriodDays,omitempty"`
	// How the payor has chosen to refer to payees.
	CollectiveAlias *string `json:"collectiveAlias,omitempty"`
	// The payor’s support contact email address.
	SupportContact *string `json:"supportContact,omitempty"`
	// The payor’s 'Doing Business As' name.
	DbaName *string `json:"dbaName,omitempty"`
	// Whether or not the payor allows language choice in the UI.
	AllowsLanguageChoice *bool `json:"allowsLanguageChoice,omitempty"`
	// Whether or not the payor has opted-out of reminder emails being sent.
	ReminderEmailsOptOut *bool `json:"reminderEmailsOptOut,omitempty"`
	// The payor’s language preference. Must be one of [EN, FR]
	Language *string `json:"language,omitempty"`
	// For internal use only (will be removed in a later version)
	// Deprecated
	IncludesReports *bool `json:"includesReports,omitempty"`
	// For internal use only (will be removed in a later version)
	// Deprecated
	WuCustomerId *string `json:"wuCustomerId,omitempty"`
	// The maximum number of payor users with the master admin role
	MaxMasterPayorAdmins *int32 `json:"maxMasterPayorAdmins,omitempty"`
	// For internal use only (will be removed in a later version)
	// Deprecated
	PaymentRails *string `json:"paymentRails,omitempty"`
	// For internal use only (will be removed in a later version)
	// Deprecated
	RemoteSystemIds []string `json:"remoteSystemIds,omitempty"`
	// USD in minor units. For internal use only (will be removed in a later version)
	// Deprecated
	UsdTxnValueReportingThreshold *int32 `json:"usdTxnValueReportingThreshold,omitempty"`
	// Does this payor manage their own payees (payees are not invited but managed by the payor)
	ManagingPayees *bool `json:"managingPayees,omitempty"`
	// The date of creation of the payor
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewPayorV2 instantiates a new PayorV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayorV2(payorId string, payorName string) *PayorV2 {
	this := PayorV2{}
	this.PayorId = payorId
	this.PayorName = payorName
	return &this
}

// NewPayorV2WithDefaults instantiates a new PayorV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayorV2WithDefaults() *PayorV2 {
	this := PayorV2{}
	return &this
}

// GetPayorId returns the PayorId field value
func (o *PayorV2) GetPayorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value
// and a boolean to check if the value has been set.
func (o *PayorV2) GetPayorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayorId, true
}

// SetPayorId sets field value
func (o *PayorV2) SetPayorId(v string) {
	o.PayorId = v
}

// GetPayorName returns the PayorName field value
func (o *PayorV2) GetPayorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayorName
}

// GetPayorNameOk returns a tuple with the PayorName field value
// and a boolean to check if the value has been set.
func (o *PayorV2) GetPayorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayorName, true
}

// SetPayorName sets field value
func (o *PayorV2) SetPayorName(v string) {
	o.PayorName = v
}

// GetPayorXid returns the PayorXid field value if set, zero value otherwise.
func (o *PayorV2) GetPayorXid() string {
	if o == nil || IsNil(o.PayorXid) {
		var ret string
		return ret
	}
	return *o.PayorXid
}

// GetPayorXidOk returns a tuple with the PayorXid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetPayorXidOk() (*string, bool) {
	if o == nil || IsNil(o.PayorXid) {
		return nil, false
	}
	return o.PayorXid, true
}

// HasPayorXid returns a boolean if a field has been set.
func (o *PayorV2) HasPayorXid() bool {
	if o != nil && !IsNil(o.PayorXid) {
		return true
	}

	return false
}

// SetPayorXid gets a reference to the given string and assigns it to the PayorXid field.
func (o *PayorV2) SetPayorXid(v string) {
	o.PayorXid = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PayorV2) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PayorV2) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *PayorV2) SetProvider(v string) {
	o.Provider = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PayorV2) GetAddress() PayorAddressV2 {
	if o == nil || IsNil(o.Address) {
		var ret PayorAddressV2
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetAddressOk() (*PayorAddressV2, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PayorV2) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given PayorAddressV2 and assigns it to the Address field.
func (o *PayorV2) SetAddress(v PayorAddressV2) {
	o.Address = &v
}

// GetPrimaryContactName returns the PrimaryContactName field value if set, zero value otherwise.
func (o *PayorV2) GetPrimaryContactName() string {
	if o == nil || IsNil(o.PrimaryContactName) {
		var ret string
		return ret
	}
	return *o.PrimaryContactName
}

// GetPrimaryContactNameOk returns a tuple with the PrimaryContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetPrimaryContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactName) {
		return nil, false
	}
	return o.PrimaryContactName, true
}

// HasPrimaryContactName returns a boolean if a field has been set.
func (o *PayorV2) HasPrimaryContactName() bool {
	if o != nil && !IsNil(o.PrimaryContactName) {
		return true
	}

	return false
}

// SetPrimaryContactName gets a reference to the given string and assigns it to the PrimaryContactName field.
func (o *PayorV2) SetPrimaryContactName(v string) {
	o.PrimaryContactName = &v
}

// GetPrimaryContactPhone returns the PrimaryContactPhone field value if set, zero value otherwise.
func (o *PayorV2) GetPrimaryContactPhone() string {
	if o == nil || IsNil(o.PrimaryContactPhone) {
		var ret string
		return ret
	}
	return *o.PrimaryContactPhone
}

// GetPrimaryContactPhoneOk returns a tuple with the PrimaryContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetPrimaryContactPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactPhone) {
		return nil, false
	}
	return o.PrimaryContactPhone, true
}

// HasPrimaryContactPhone returns a boolean if a field has been set.
func (o *PayorV2) HasPrimaryContactPhone() bool {
	if o != nil && !IsNil(o.PrimaryContactPhone) {
		return true
	}

	return false
}

// SetPrimaryContactPhone gets a reference to the given string and assigns it to the PrimaryContactPhone field.
func (o *PayorV2) SetPrimaryContactPhone(v string) {
	o.PrimaryContactPhone = &v
}

// GetPrimaryContactEmail returns the PrimaryContactEmail field value if set, zero value otherwise.
func (o *PayorV2) GetPrimaryContactEmail() string {
	if o == nil || IsNil(o.PrimaryContactEmail) {
		var ret string
		return ret
	}
	return *o.PrimaryContactEmail
}

// GetPrimaryContactEmailOk returns a tuple with the PrimaryContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetPrimaryContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactEmail) {
		return nil, false
	}
	return o.PrimaryContactEmail, true
}

// HasPrimaryContactEmail returns a boolean if a field has been set.
func (o *PayorV2) HasPrimaryContactEmail() bool {
	if o != nil && !IsNil(o.PrimaryContactEmail) {
		return true
	}

	return false
}

// SetPrimaryContactEmail gets a reference to the given string and assigns it to the PrimaryContactEmail field.
func (o *PayorV2) SetPrimaryContactEmail(v string) {
	o.PrimaryContactEmail = &v
}

// GetKycState returns the KycState field value if set, zero value otherwise.
func (o *PayorV2) GetKycState() string {
	if o == nil || IsNil(o.KycState) {
		var ret string
		return ret
	}
	return *o.KycState
}

// GetKycStateOk returns a tuple with the KycState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetKycStateOk() (*string, bool) {
	if o == nil || IsNil(o.KycState) {
		return nil, false
	}
	return o.KycState, true
}

// HasKycState returns a boolean if a field has been set.
func (o *PayorV2) HasKycState() bool {
	if o != nil && !IsNil(o.KycState) {
		return true
	}

	return false
}

// SetKycState gets a reference to the given string and assigns it to the KycState field.
func (o *PayorV2) SetKycState(v string) {
	o.KycState = &v
}

// GetManualLockout returns the ManualLockout field value if set, zero value otherwise.
func (o *PayorV2) GetManualLockout() bool {
	if o == nil || IsNil(o.ManualLockout) {
		var ret bool
		return ret
	}
	return *o.ManualLockout
}

// GetManualLockoutOk returns a tuple with the ManualLockout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetManualLockoutOk() (*bool, bool) {
	if o == nil || IsNil(o.ManualLockout) {
		return nil, false
	}
	return o.ManualLockout, true
}

// HasManualLockout returns a boolean if a field has been set.
func (o *PayorV2) HasManualLockout() bool {
	if o != nil && !IsNil(o.ManualLockout) {
		return true
	}

	return false
}

// SetManualLockout gets a reference to the given bool and assigns it to the ManualLockout field.
func (o *PayorV2) SetManualLockout(v bool) {
	o.ManualLockout = &v
}

// GetOpenBankingEnabled returns the OpenBankingEnabled field value if set, zero value otherwise.
func (o *PayorV2) GetOpenBankingEnabled() bool {
	if o == nil || IsNil(o.OpenBankingEnabled) {
		var ret bool
		return ret
	}
	return *o.OpenBankingEnabled
}

// GetOpenBankingEnabledOk returns a tuple with the OpenBankingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetOpenBankingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenBankingEnabled) {
		return nil, false
	}
	return o.OpenBankingEnabled, true
}

// HasOpenBankingEnabled returns a boolean if a field has been set.
func (o *PayorV2) HasOpenBankingEnabled() bool {
	if o != nil && !IsNil(o.OpenBankingEnabled) {
		return true
	}

	return false
}

// SetOpenBankingEnabled gets a reference to the given bool and assigns it to the OpenBankingEnabled field.
func (o *PayorV2) SetOpenBankingEnabled(v bool) {
	o.OpenBankingEnabled = &v
}

// GetPayeeGracePeriodProcessingEnabled returns the PayeeGracePeriodProcessingEnabled field value if set, zero value otherwise.
func (o *PayorV2) GetPayeeGracePeriodProcessingEnabled() bool {
	if o == nil || IsNil(o.PayeeGracePeriodProcessingEnabled) {
		var ret bool
		return ret
	}
	return *o.PayeeGracePeriodProcessingEnabled
}

// GetPayeeGracePeriodProcessingEnabledOk returns a tuple with the PayeeGracePeriodProcessingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetPayeeGracePeriodProcessingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PayeeGracePeriodProcessingEnabled) {
		return nil, false
	}
	return o.PayeeGracePeriodProcessingEnabled, true
}

// HasPayeeGracePeriodProcessingEnabled returns a boolean if a field has been set.
func (o *PayorV2) HasPayeeGracePeriodProcessingEnabled() bool {
	if o != nil && !IsNil(o.PayeeGracePeriodProcessingEnabled) {
		return true
	}

	return false
}

// SetPayeeGracePeriodProcessingEnabled gets a reference to the given bool and assigns it to the PayeeGracePeriodProcessingEnabled field.
func (o *PayorV2) SetPayeeGracePeriodProcessingEnabled(v bool) {
	o.PayeeGracePeriodProcessingEnabled = &v
}

// GetPayeeGracePeriodDays returns the PayeeGracePeriodDays field value if set, zero value otherwise.
func (o *PayorV2) GetPayeeGracePeriodDays() int32 {
	if o == nil || IsNil(o.PayeeGracePeriodDays) {
		var ret int32
		return ret
	}
	return *o.PayeeGracePeriodDays
}

// GetPayeeGracePeriodDaysOk returns a tuple with the PayeeGracePeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetPayeeGracePeriodDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.PayeeGracePeriodDays) {
		return nil, false
	}
	return o.PayeeGracePeriodDays, true
}

// HasPayeeGracePeriodDays returns a boolean if a field has been set.
func (o *PayorV2) HasPayeeGracePeriodDays() bool {
	if o != nil && !IsNil(o.PayeeGracePeriodDays) {
		return true
	}

	return false
}

// SetPayeeGracePeriodDays gets a reference to the given int32 and assigns it to the PayeeGracePeriodDays field.
func (o *PayorV2) SetPayeeGracePeriodDays(v int32) {
	o.PayeeGracePeriodDays = &v
}

// GetCollectiveAlias returns the CollectiveAlias field value if set, zero value otherwise.
func (o *PayorV2) GetCollectiveAlias() string {
	if o == nil || IsNil(o.CollectiveAlias) {
		var ret string
		return ret
	}
	return *o.CollectiveAlias
}

// GetCollectiveAliasOk returns a tuple with the CollectiveAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetCollectiveAliasOk() (*string, bool) {
	if o == nil || IsNil(o.CollectiveAlias) {
		return nil, false
	}
	return o.CollectiveAlias, true
}

// HasCollectiveAlias returns a boolean if a field has been set.
func (o *PayorV2) HasCollectiveAlias() bool {
	if o != nil && !IsNil(o.CollectiveAlias) {
		return true
	}

	return false
}

// SetCollectiveAlias gets a reference to the given string and assigns it to the CollectiveAlias field.
func (o *PayorV2) SetCollectiveAlias(v string) {
	o.CollectiveAlias = &v
}

// GetSupportContact returns the SupportContact field value if set, zero value otherwise.
func (o *PayorV2) GetSupportContact() string {
	if o == nil || IsNil(o.SupportContact) {
		var ret string
		return ret
	}
	return *o.SupportContact
}

// GetSupportContactOk returns a tuple with the SupportContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetSupportContactOk() (*string, bool) {
	if o == nil || IsNil(o.SupportContact) {
		return nil, false
	}
	return o.SupportContact, true
}

// HasSupportContact returns a boolean if a field has been set.
func (o *PayorV2) HasSupportContact() bool {
	if o != nil && !IsNil(o.SupportContact) {
		return true
	}

	return false
}

// SetSupportContact gets a reference to the given string and assigns it to the SupportContact field.
func (o *PayorV2) SetSupportContact(v string) {
	o.SupportContact = &v
}

// GetDbaName returns the DbaName field value if set, zero value otherwise.
func (o *PayorV2) GetDbaName() string {
	if o == nil || IsNil(o.DbaName) {
		var ret string
		return ret
	}
	return *o.DbaName
}

// GetDbaNameOk returns a tuple with the DbaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetDbaNameOk() (*string, bool) {
	if o == nil || IsNil(o.DbaName) {
		return nil, false
	}
	return o.DbaName, true
}

// HasDbaName returns a boolean if a field has been set.
func (o *PayorV2) HasDbaName() bool {
	if o != nil && !IsNil(o.DbaName) {
		return true
	}

	return false
}

// SetDbaName gets a reference to the given string and assigns it to the DbaName field.
func (o *PayorV2) SetDbaName(v string) {
	o.DbaName = &v
}

// GetAllowsLanguageChoice returns the AllowsLanguageChoice field value if set, zero value otherwise.
func (o *PayorV2) GetAllowsLanguageChoice() bool {
	if o == nil || IsNil(o.AllowsLanguageChoice) {
		var ret bool
		return ret
	}
	return *o.AllowsLanguageChoice
}

// GetAllowsLanguageChoiceOk returns a tuple with the AllowsLanguageChoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetAllowsLanguageChoiceOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowsLanguageChoice) {
		return nil, false
	}
	return o.AllowsLanguageChoice, true
}

// HasAllowsLanguageChoice returns a boolean if a field has been set.
func (o *PayorV2) HasAllowsLanguageChoice() bool {
	if o != nil && !IsNil(o.AllowsLanguageChoice) {
		return true
	}

	return false
}

// SetAllowsLanguageChoice gets a reference to the given bool and assigns it to the AllowsLanguageChoice field.
func (o *PayorV2) SetAllowsLanguageChoice(v bool) {
	o.AllowsLanguageChoice = &v
}

// GetReminderEmailsOptOut returns the ReminderEmailsOptOut field value if set, zero value otherwise.
func (o *PayorV2) GetReminderEmailsOptOut() bool {
	if o == nil || IsNil(o.ReminderEmailsOptOut) {
		var ret bool
		return ret
	}
	return *o.ReminderEmailsOptOut
}

// GetReminderEmailsOptOutOk returns a tuple with the ReminderEmailsOptOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetReminderEmailsOptOutOk() (*bool, bool) {
	if o == nil || IsNil(o.ReminderEmailsOptOut) {
		return nil, false
	}
	return o.ReminderEmailsOptOut, true
}

// HasReminderEmailsOptOut returns a boolean if a field has been set.
func (o *PayorV2) HasReminderEmailsOptOut() bool {
	if o != nil && !IsNil(o.ReminderEmailsOptOut) {
		return true
	}

	return false
}

// SetReminderEmailsOptOut gets a reference to the given bool and assigns it to the ReminderEmailsOptOut field.
func (o *PayorV2) SetReminderEmailsOptOut(v bool) {
	o.ReminderEmailsOptOut = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PayorV2) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PayorV2) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *PayorV2) SetLanguage(v string) {
	o.Language = &v
}

// GetIncludesReports returns the IncludesReports field value if set, zero value otherwise.
// Deprecated
func (o *PayorV2) GetIncludesReports() bool {
	if o == nil || IsNil(o.IncludesReports) {
		var ret bool
		return ret
	}
	return *o.IncludesReports
}

// GetIncludesReportsOk returns a tuple with the IncludesReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PayorV2) GetIncludesReportsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludesReports) {
		return nil, false
	}
	return o.IncludesReports, true
}

// HasIncludesReports returns a boolean if a field has been set.
func (o *PayorV2) HasIncludesReports() bool {
	if o != nil && !IsNil(o.IncludesReports) {
		return true
	}

	return false
}

// SetIncludesReports gets a reference to the given bool and assigns it to the IncludesReports field.
// Deprecated
func (o *PayorV2) SetIncludesReports(v bool) {
	o.IncludesReports = &v
}

// GetWuCustomerId returns the WuCustomerId field value if set, zero value otherwise.
// Deprecated
func (o *PayorV2) GetWuCustomerId() string {
	if o == nil || IsNil(o.WuCustomerId) {
		var ret string
		return ret
	}
	return *o.WuCustomerId
}

// GetWuCustomerIdOk returns a tuple with the WuCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PayorV2) GetWuCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.WuCustomerId) {
		return nil, false
	}
	return o.WuCustomerId, true
}

// HasWuCustomerId returns a boolean if a field has been set.
func (o *PayorV2) HasWuCustomerId() bool {
	if o != nil && !IsNil(o.WuCustomerId) {
		return true
	}

	return false
}

// SetWuCustomerId gets a reference to the given string and assigns it to the WuCustomerId field.
// Deprecated
func (o *PayorV2) SetWuCustomerId(v string) {
	o.WuCustomerId = &v
}

// GetMaxMasterPayorAdmins returns the MaxMasterPayorAdmins field value if set, zero value otherwise.
func (o *PayorV2) GetMaxMasterPayorAdmins() int32 {
	if o == nil || IsNil(o.MaxMasterPayorAdmins) {
		var ret int32
		return ret
	}
	return *o.MaxMasterPayorAdmins
}

// GetMaxMasterPayorAdminsOk returns a tuple with the MaxMasterPayorAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetMaxMasterPayorAdminsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxMasterPayorAdmins) {
		return nil, false
	}
	return o.MaxMasterPayorAdmins, true
}

// HasMaxMasterPayorAdmins returns a boolean if a field has been set.
func (o *PayorV2) HasMaxMasterPayorAdmins() bool {
	if o != nil && !IsNil(o.MaxMasterPayorAdmins) {
		return true
	}

	return false
}

// SetMaxMasterPayorAdmins gets a reference to the given int32 and assigns it to the MaxMasterPayorAdmins field.
func (o *PayorV2) SetMaxMasterPayorAdmins(v int32) {
	o.MaxMasterPayorAdmins = &v
}

// GetPaymentRails returns the PaymentRails field value if set, zero value otherwise.
// Deprecated
func (o *PayorV2) GetPaymentRails() string {
	if o == nil || IsNil(o.PaymentRails) {
		var ret string
		return ret
	}
	return *o.PaymentRails
}

// GetPaymentRailsOk returns a tuple with the PaymentRails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PayorV2) GetPaymentRailsOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentRails) {
		return nil, false
	}
	return o.PaymentRails, true
}

// HasPaymentRails returns a boolean if a field has been set.
func (o *PayorV2) HasPaymentRails() bool {
	if o != nil && !IsNil(o.PaymentRails) {
		return true
	}

	return false
}

// SetPaymentRails gets a reference to the given string and assigns it to the PaymentRails field.
// Deprecated
func (o *PayorV2) SetPaymentRails(v string) {
	o.PaymentRails = &v
}

// GetRemoteSystemIds returns the RemoteSystemIds field value if set, zero value otherwise.
// Deprecated
func (o *PayorV2) GetRemoteSystemIds() []string {
	if o == nil || IsNil(o.RemoteSystemIds) {
		var ret []string
		return ret
	}
	return o.RemoteSystemIds
}

// GetRemoteSystemIdsOk returns a tuple with the RemoteSystemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PayorV2) GetRemoteSystemIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoteSystemIds) {
		return nil, false
	}
	return o.RemoteSystemIds, true
}

// HasRemoteSystemIds returns a boolean if a field has been set.
func (o *PayorV2) HasRemoteSystemIds() bool {
	if o != nil && !IsNil(o.RemoteSystemIds) {
		return true
	}

	return false
}

// SetRemoteSystemIds gets a reference to the given []string and assigns it to the RemoteSystemIds field.
// Deprecated
func (o *PayorV2) SetRemoteSystemIds(v []string) {
	o.RemoteSystemIds = v
}

// GetUsdTxnValueReportingThreshold returns the UsdTxnValueReportingThreshold field value if set, zero value otherwise.
// Deprecated
func (o *PayorV2) GetUsdTxnValueReportingThreshold() int32 {
	if o == nil || IsNil(o.UsdTxnValueReportingThreshold) {
		var ret int32
		return ret
	}
	return *o.UsdTxnValueReportingThreshold
}

// GetUsdTxnValueReportingThresholdOk returns a tuple with the UsdTxnValueReportingThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PayorV2) GetUsdTxnValueReportingThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.UsdTxnValueReportingThreshold) {
		return nil, false
	}
	return o.UsdTxnValueReportingThreshold, true
}

// HasUsdTxnValueReportingThreshold returns a boolean if a field has been set.
func (o *PayorV2) HasUsdTxnValueReportingThreshold() bool {
	if o != nil && !IsNil(o.UsdTxnValueReportingThreshold) {
		return true
	}

	return false
}

// SetUsdTxnValueReportingThreshold gets a reference to the given int32 and assigns it to the UsdTxnValueReportingThreshold field.
// Deprecated
func (o *PayorV2) SetUsdTxnValueReportingThreshold(v int32) {
	o.UsdTxnValueReportingThreshold = &v
}

// GetManagingPayees returns the ManagingPayees field value if set, zero value otherwise.
func (o *PayorV2) GetManagingPayees() bool {
	if o == nil || IsNil(o.ManagingPayees) {
		var ret bool
		return ret
	}
	return *o.ManagingPayees
}

// GetManagingPayeesOk returns a tuple with the ManagingPayees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetManagingPayeesOk() (*bool, bool) {
	if o == nil || IsNil(o.ManagingPayees) {
		return nil, false
	}
	return o.ManagingPayees, true
}

// HasManagingPayees returns a boolean if a field has been set.
func (o *PayorV2) HasManagingPayees() bool {
	if o != nil && !IsNil(o.ManagingPayees) {
		return true
	}

	return false
}

// SetManagingPayees gets a reference to the given bool and assigns it to the ManagingPayees field.
func (o *PayorV2) SetManagingPayees(v bool) {
	o.ManagingPayees = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PayorV2) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorV2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PayorV2) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PayorV2) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o PayorV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayorV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payorId"] = o.PayorId
	toSerialize["payorName"] = o.PayorName
	if !IsNil(o.PayorXid) {
		toSerialize["payorXid"] = o.PayorXid
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.PrimaryContactName) {
		toSerialize["primaryContactName"] = o.PrimaryContactName
	}
	if !IsNil(o.PrimaryContactPhone) {
		toSerialize["primaryContactPhone"] = o.PrimaryContactPhone
	}
	if !IsNil(o.PrimaryContactEmail) {
		toSerialize["primaryContactEmail"] = o.PrimaryContactEmail
	}
	if !IsNil(o.KycState) {
		toSerialize["kycState"] = o.KycState
	}
	if !IsNil(o.ManualLockout) {
		toSerialize["manualLockout"] = o.ManualLockout
	}
	if !IsNil(o.OpenBankingEnabled) {
		toSerialize["openBankingEnabled"] = o.OpenBankingEnabled
	}
	if !IsNil(o.PayeeGracePeriodProcessingEnabled) {
		toSerialize["payeeGracePeriodProcessingEnabled"] = o.PayeeGracePeriodProcessingEnabled
	}
	if !IsNil(o.PayeeGracePeriodDays) {
		toSerialize["payeeGracePeriodDays"] = o.PayeeGracePeriodDays
	}
	if !IsNil(o.CollectiveAlias) {
		toSerialize["collectiveAlias"] = o.CollectiveAlias
	}
	if !IsNil(o.SupportContact) {
		toSerialize["supportContact"] = o.SupportContact
	}
	if !IsNil(o.DbaName) {
		toSerialize["dbaName"] = o.DbaName
	}
	if !IsNil(o.AllowsLanguageChoice) {
		toSerialize["allowsLanguageChoice"] = o.AllowsLanguageChoice
	}
	if !IsNil(o.ReminderEmailsOptOut) {
		toSerialize["reminderEmailsOptOut"] = o.ReminderEmailsOptOut
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.IncludesReports) {
		toSerialize["includesReports"] = o.IncludesReports
	}
	if !IsNil(o.WuCustomerId) {
		toSerialize["wuCustomerId"] = o.WuCustomerId
	}
	if !IsNil(o.MaxMasterPayorAdmins) {
		toSerialize["maxMasterPayorAdmins"] = o.MaxMasterPayorAdmins
	}
	if !IsNil(o.PaymentRails) {
		toSerialize["paymentRails"] = o.PaymentRails
	}
	if !IsNil(o.RemoteSystemIds) {
		toSerialize["remoteSystemIds"] = o.RemoteSystemIds
	}
	if !IsNil(o.UsdTxnValueReportingThreshold) {
		toSerialize["usdTxnValueReportingThreshold"] = o.UsdTxnValueReportingThreshold
	}
	if !IsNil(o.ManagingPayees) {
		toSerialize["managingPayees"] = o.ManagingPayees
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullablePayorV2 struct {
	value *PayorV2
	isSet bool
}

func (v NullablePayorV2) Get() *PayorV2 {
	return v.value
}

func (v *NullablePayorV2) Set(val *PayorV2) {
	v.value = val
	v.isSet = true
}

func (v NullablePayorV2) IsSet() bool {
	return v.isSet
}

func (v *NullablePayorV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayorV2(val *PayorV2) *NullablePayorV2 {
	return &NullablePayorV2{value: val, isSet: true}
}

func (v NullablePayorV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayorV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


