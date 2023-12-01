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

// checks if the SourceAccountResponseV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceAccountResponseV3{}

// SourceAccountResponseV3 struct for SourceAccountResponseV3
type SourceAccountResponseV3 struct {
	// Source Account Id
	Id string `json:"id"`
	// Decimal implied
	Balance *int64 `json:"balance,omitempty"`
	// Valid ISO 4217 3 letter currency code. See the <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	Currency *string `json:"currency,omitempty"`
	// The funding reference (will not be set for DECOUPLED accounts).
	FundingRef *string `json:"fundingRef,omitempty"`
	// The physical account name (will not be set for DECOUPLED accounts).
	PhysicalAccountName *string `json:"physicalAccountName,omitempty"`
	RailsId string `json:"railsId"`
	PayorId *string `json:"payorId,omitempty"`
	Name *string `json:"name,omitempty"`
	// The pooled account flag (will not be set for DECOUPLED accounts).
	Pooled *bool `json:"pooled,omitempty"`
	CustomerId NullableString `json:"customerId,omitempty"`
	// The physical account id (will not be set for DECOUPLED accounts).
	PhysicalAccountId *string `json:"physicalAccountId,omitempty"`
	Notifications *NotificationsV3 `json:"notifications,omitempty"`
	AutoTopUpConfig *AutoTopUpConfigV3 `json:"autoTopUpConfig,omitempty"`
	Type string `json:"type"`
	// Valid ISO 3166 2 character country code. See the <a href=\"https://www.iso.org/iso-3166-country-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	Country *string `json:"country,omitempty"`
	// An optional flag for whether the source account has been deleted. Only present in the response if true.
	Deleted *bool `json:"deleted,omitempty"`
	// An optional flag for whether the source account has been deleted by a user. Only present in the response if true.
	UserDeleted *bool `json:"userDeleted,omitempty"`
	// An optional timestamp when the source account has been deleted. Only present in the response if deleted.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// List of supported transmission types.
	TransmissionTypes []string `json:"transmissionTypes,omitempty"`
}

// NewSourceAccountResponseV3 instantiates a new SourceAccountResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAccountResponseV3(id string, railsId string, type_ string) *SourceAccountResponseV3 {
	this := SourceAccountResponseV3{}
	this.Id = id
	this.RailsId = railsId
	this.Type = type_
	return &this
}

// NewSourceAccountResponseV3WithDefaults instantiates a new SourceAccountResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAccountResponseV3WithDefaults() *SourceAccountResponseV3 {
	this := SourceAccountResponseV3{}
	return &this
}

// GetId returns the Id field value
func (o *SourceAccountResponseV3) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceAccountResponseV3) SetId(v string) {
	o.Id = v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetBalance() int64 {
	if o == nil || IsNil(o.Balance) {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetBalanceOk() (*int64, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *SourceAccountResponseV3) SetBalance(v int64) {
	o.Balance = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SourceAccountResponseV3) SetCurrency(v string) {
	o.Currency = &v
}

// GetFundingRef returns the FundingRef field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetFundingRef() string {
	if o == nil || IsNil(o.FundingRef) {
		var ret string
		return ret
	}
	return *o.FundingRef
}

// GetFundingRefOk returns a tuple with the FundingRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetFundingRefOk() (*string, bool) {
	if o == nil || IsNil(o.FundingRef) {
		return nil, false
	}
	return o.FundingRef, true
}

// HasFundingRef returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasFundingRef() bool {
	if o != nil && !IsNil(o.FundingRef) {
		return true
	}

	return false
}

// SetFundingRef gets a reference to the given string and assigns it to the FundingRef field.
func (o *SourceAccountResponseV3) SetFundingRef(v string) {
	o.FundingRef = &v
}

// GetPhysicalAccountName returns the PhysicalAccountName field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetPhysicalAccountName() string {
	if o == nil || IsNil(o.PhysicalAccountName) {
		var ret string
		return ret
	}
	return *o.PhysicalAccountName
}

// GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetPhysicalAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalAccountName) {
		return nil, false
	}
	return o.PhysicalAccountName, true
}

// HasPhysicalAccountName returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasPhysicalAccountName() bool {
	if o != nil && !IsNil(o.PhysicalAccountName) {
		return true
	}

	return false
}

// SetPhysicalAccountName gets a reference to the given string and assigns it to the PhysicalAccountName field.
func (o *SourceAccountResponseV3) SetPhysicalAccountName(v string) {
	o.PhysicalAccountName = &v
}

// GetRailsId returns the RailsId field value
func (o *SourceAccountResponseV3) GetRailsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RailsId
}

// GetRailsIdOk returns a tuple with the RailsId field value
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetRailsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RailsId, true
}

// SetRailsId sets field value
func (o *SourceAccountResponseV3) SetRailsId(v string) {
	o.RailsId = v
}

// GetPayorId returns the PayorId field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetPayorId() string {
	if o == nil || IsNil(o.PayorId) {
		var ret string
		return ret
	}
	return *o.PayorId
}

// GetPayorIdOk returns a tuple with the PayorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetPayorIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayorId) {
		return nil, false
	}
	return o.PayorId, true
}

// HasPayorId returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasPayorId() bool {
	if o != nil && !IsNil(o.PayorId) {
		return true
	}

	return false
}

// SetPayorId gets a reference to the given string and assigns it to the PayorId field.
func (o *SourceAccountResponseV3) SetPayorId(v string) {
	o.PayorId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceAccountResponseV3) SetName(v string) {
	o.Name = &v
}

// GetPooled returns the Pooled field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetPooled() bool {
	if o == nil || IsNil(o.Pooled) {
		var ret bool
		return ret
	}
	return *o.Pooled
}

// GetPooledOk returns a tuple with the Pooled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetPooledOk() (*bool, bool) {
	if o == nil || IsNil(o.Pooled) {
		return nil, false
	}
	return o.Pooled, true
}

// HasPooled returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasPooled() bool {
	if o != nil && !IsNil(o.Pooled) {
		return true
	}

	return false
}

// SetPooled gets a reference to the given bool and assigns it to the Pooled field.
func (o *SourceAccountResponseV3) SetPooled(v bool) {
	o.Pooled = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceAccountResponseV3) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceAccountResponseV3) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *SourceAccountResponseV3) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *SourceAccountResponseV3) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *SourceAccountResponseV3) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetPhysicalAccountId returns the PhysicalAccountId field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetPhysicalAccountId() string {
	if o == nil || IsNil(o.PhysicalAccountId) {
		var ret string
		return ret
	}
	return *o.PhysicalAccountId
}

// GetPhysicalAccountIdOk returns a tuple with the PhysicalAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetPhysicalAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalAccountId) {
		return nil, false
	}
	return o.PhysicalAccountId, true
}

// HasPhysicalAccountId returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasPhysicalAccountId() bool {
	if o != nil && !IsNil(o.PhysicalAccountId) {
		return true
	}

	return false
}

// SetPhysicalAccountId gets a reference to the given string and assigns it to the PhysicalAccountId field.
func (o *SourceAccountResponseV3) SetPhysicalAccountId(v string) {
	o.PhysicalAccountId = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetNotifications() NotificationsV3 {
	if o == nil || IsNil(o.Notifications) {
		var ret NotificationsV3
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetNotificationsOk() (*NotificationsV3, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given NotificationsV3 and assigns it to the Notifications field.
func (o *SourceAccountResponseV3) SetNotifications(v NotificationsV3) {
	o.Notifications = &v
}

// GetAutoTopUpConfig returns the AutoTopUpConfig field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetAutoTopUpConfig() AutoTopUpConfigV3 {
	if o == nil || IsNil(o.AutoTopUpConfig) {
		var ret AutoTopUpConfigV3
		return ret
	}
	return *o.AutoTopUpConfig
}

// GetAutoTopUpConfigOk returns a tuple with the AutoTopUpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetAutoTopUpConfigOk() (*AutoTopUpConfigV3, bool) {
	if o == nil || IsNil(o.AutoTopUpConfig) {
		return nil, false
	}
	return o.AutoTopUpConfig, true
}

// HasAutoTopUpConfig returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasAutoTopUpConfig() bool {
	if o != nil && !IsNil(o.AutoTopUpConfig) {
		return true
	}

	return false
}

// SetAutoTopUpConfig gets a reference to the given AutoTopUpConfigV3 and assigns it to the AutoTopUpConfig field.
func (o *SourceAccountResponseV3) SetAutoTopUpConfig(v AutoTopUpConfigV3) {
	o.AutoTopUpConfig = &v
}

// GetType returns the Type field value
func (o *SourceAccountResponseV3) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SourceAccountResponseV3) SetType(v string) {
	o.Type = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *SourceAccountResponseV3) SetCountry(v string) {
	o.Country = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *SourceAccountResponseV3) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetUserDeleted returns the UserDeleted field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetUserDeleted() bool {
	if o == nil || IsNil(o.UserDeleted) {
		var ret bool
		return ret
	}
	return *o.UserDeleted
}

// GetUserDeletedOk returns a tuple with the UserDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetUserDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.UserDeleted) {
		return nil, false
	}
	return o.UserDeleted, true
}

// HasUserDeleted returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasUserDeleted() bool {
	if o != nil && !IsNil(o.UserDeleted) {
		return true
	}

	return false
}

// SetUserDeleted gets a reference to the given bool and assigns it to the UserDeleted field.
func (o *SourceAccountResponseV3) SetUserDeleted(v bool) {
	o.UserDeleted = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *SourceAccountResponseV3) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetTransmissionTypes returns the TransmissionTypes field value if set, zero value otherwise.
func (o *SourceAccountResponseV3) GetTransmissionTypes() []string {
	if o == nil || IsNil(o.TransmissionTypes) {
		var ret []string
		return ret
	}
	return o.TransmissionTypes
}

// GetTransmissionTypesOk returns a tuple with the TransmissionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccountResponseV3) GetTransmissionTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.TransmissionTypes) {
		return nil, false
	}
	return o.TransmissionTypes, true
}

// HasTransmissionTypes returns a boolean if a field has been set.
func (o *SourceAccountResponseV3) HasTransmissionTypes() bool {
	if o != nil && !IsNil(o.TransmissionTypes) {
		return true
	}

	return false
}

// SetTransmissionTypes gets a reference to the given []string and assigns it to the TransmissionTypes field.
func (o *SourceAccountResponseV3) SetTransmissionTypes(v []string) {
	o.TransmissionTypes = v
}

func (o SourceAccountResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceAccountResponseV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.FundingRef) {
		toSerialize["fundingRef"] = o.FundingRef
	}
	if !IsNil(o.PhysicalAccountName) {
		toSerialize["physicalAccountName"] = o.PhysicalAccountName
	}
	toSerialize["railsId"] = o.RailsId
	if !IsNil(o.PayorId) {
		toSerialize["payorId"] = o.PayorId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Pooled) {
		toSerialize["pooled"] = o.Pooled
	}
	if o.CustomerId.IsSet() {
		toSerialize["customerId"] = o.CustomerId.Get()
	}
	if !IsNil(o.PhysicalAccountId) {
		toSerialize["physicalAccountId"] = o.PhysicalAccountId
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.AutoTopUpConfig) {
		toSerialize["autoTopUpConfig"] = o.AutoTopUpConfig
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.UserDeleted) {
		toSerialize["userDeleted"] = o.UserDeleted
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if !IsNil(o.TransmissionTypes) {
		toSerialize["transmissionTypes"] = o.TransmissionTypes
	}
	return toSerialize, nil
}

type NullableSourceAccountResponseV3 struct {
	value *SourceAccountResponseV3
	isSet bool
}

func (v NullableSourceAccountResponseV3) Get() *SourceAccountResponseV3 {
	return v.value
}

func (v *NullableSourceAccountResponseV3) Set(val *SourceAccountResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAccountResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAccountResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAccountResponseV3(val *SourceAccountResponseV3) *NullableSourceAccountResponseV3 {
	return &NullableSourceAccountResponseV3{value: val, isSet: true}
}

func (v NullableSourceAccountResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAccountResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


