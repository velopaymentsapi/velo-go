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

// checks if the PaymentChannelResponseV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentChannelResponseV4{}

// PaymentChannelResponseV4 struct for PaymentChannelResponseV4
type PaymentChannelResponseV4 struct {
	Id string `json:"id"`
	PayeeId *string `json:"payeeId,omitempty"`
	PaymentChannelName string `json:"paymentChannelName"`
	AccountName string `json:"accountName"`
	// Payment channel type. One of the following values: CHANNEL_BANK
	ChannelType string `json:"channelType"`
	// Valid ISO 3166 2 character country code. See the <a href=\"https://www.iso.org/iso-3166-country-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	CountryCode string `json:"countryCode"`
	RoutingNumber *string `json:"routingNumber,omitempty"`
	AccountNumber *string `json:"accountNumber,omitempty"`
	// Must match the regular expression ```^[A-Za-z0-9]+$```.
	Iban *string `json:"iban,omitempty"`
	// Valid ISO 4217 3 letter currency code. See the <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	Currency string `json:"currency"`
	// Deprecated
	PayorIds []string `json:"payorIds,omitempty"`
	PayeeName *string `json:"payeeName,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	BankSwiftBic *string `json:"bankSwiftBic,omitempty"`
	BankAddress *AddressV4 `json:"bankAddress,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	DisabledReasonCode *string `json:"disabledReasonCode,omitempty"`
	DisabledReason *string `json:"disabledReason,omitempty"`
	// Whether this payment channel is payable
	Payable *bool `json:"payable,omitempty"`
}

// NewPaymentChannelResponseV4 instantiates a new PaymentChannelResponseV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentChannelResponseV4(id string, paymentChannelName string, accountName string, channelType string, countryCode string, currency string) *PaymentChannelResponseV4 {
	this := PaymentChannelResponseV4{}
	this.Id = id
	this.PaymentChannelName = paymentChannelName
	this.AccountName = accountName
	this.ChannelType = channelType
	this.CountryCode = countryCode
	this.Currency = currency
	return &this
}

// NewPaymentChannelResponseV4WithDefaults instantiates a new PaymentChannelResponseV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentChannelResponseV4WithDefaults() *PaymentChannelResponseV4 {
	this := PaymentChannelResponseV4{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentChannelResponseV4) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentChannelResponseV4) SetId(v string) {
	o.Id = v
}

// GetPayeeId returns the PayeeId field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetPayeeId() string {
	if o == nil || IsNil(o.PayeeId) {
		var ret string
		return ret
	}
	return *o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetPayeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeId) {
		return nil, false
	}
	return o.PayeeId, true
}

// HasPayeeId returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasPayeeId() bool {
	if o != nil && !IsNil(o.PayeeId) {
		return true
	}

	return false
}

// SetPayeeId gets a reference to the given string and assigns it to the PayeeId field.
func (o *PaymentChannelResponseV4) SetPayeeId(v string) {
	o.PayeeId = &v
}

// GetPaymentChannelName returns the PaymentChannelName field value
func (o *PaymentChannelResponseV4) GetPaymentChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentChannelName
}

// GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetPaymentChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentChannelName, true
}

// SetPaymentChannelName sets field value
func (o *PaymentChannelResponseV4) SetPaymentChannelName(v string) {
	o.PaymentChannelName = v
}

// GetAccountName returns the AccountName field value
func (o *PaymentChannelResponseV4) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *PaymentChannelResponseV4) SetAccountName(v string) {
	o.AccountName = v
}

// GetChannelType returns the ChannelType field value
func (o *PaymentChannelResponseV4) GetChannelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelType
}

// GetChannelTypeOk returns a tuple with the ChannelType field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetChannelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelType, true
}

// SetChannelType sets field value
func (o *PaymentChannelResponseV4) SetChannelType(v string) {
	o.ChannelType = v
}

// GetCountryCode returns the CountryCode field value
func (o *PaymentChannelResponseV4) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *PaymentChannelResponseV4) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetRoutingNumber() string {
	if o == nil || IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasRoutingNumber() bool {
	if o != nil && !IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *PaymentChannelResponseV4) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *PaymentChannelResponseV4) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetIban() string {
	if o == nil || IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetIbanOk() (*string, bool) {
	if o == nil || IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasIban() bool {
	if o != nil && !IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *PaymentChannelResponseV4) SetIban(v string) {
	o.Iban = &v
}

// GetCurrency returns the Currency field value
func (o *PaymentChannelResponseV4) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentChannelResponseV4) SetCurrency(v string) {
	o.Currency = v
}

// GetPayorIds returns the PayorIds field value if set, zero value otherwise.
// Deprecated
func (o *PaymentChannelResponseV4) GetPayorIds() []string {
	if o == nil || IsNil(o.PayorIds) {
		var ret []string
		return ret
	}
	return o.PayorIds
}

// GetPayorIdsOk returns a tuple with the PayorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PaymentChannelResponseV4) GetPayorIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PayorIds) {
		return nil, false
	}
	return o.PayorIds, true
}

// HasPayorIds returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasPayorIds() bool {
	if o != nil && !IsNil(o.PayorIds) {
		return true
	}

	return false
}

// SetPayorIds gets a reference to the given []string and assigns it to the PayorIds field.
// Deprecated
func (o *PaymentChannelResponseV4) SetPayorIds(v []string) {
	o.PayorIds = v
}

// GetPayeeName returns the PayeeName field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetPayeeName() string {
	if o == nil || IsNil(o.PayeeName) {
		var ret string
		return ret
	}
	return *o.PayeeName
}

// GetPayeeNameOk returns a tuple with the PayeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetPayeeNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeName) {
		return nil, false
	}
	return o.PayeeName, true
}

// HasPayeeName returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasPayeeName() bool {
	if o != nil && !IsNil(o.PayeeName) {
		return true
	}

	return false
}

// SetPayeeName gets a reference to the given string and assigns it to the PayeeName field.
func (o *PaymentChannelResponseV4) SetPayeeName(v string) {
	o.PayeeName = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *PaymentChannelResponseV4) SetBankName(v string) {
	o.BankName = &v
}

// GetBankSwiftBic returns the BankSwiftBic field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetBankSwiftBic() string {
	if o == nil || IsNil(o.BankSwiftBic) {
		var ret string
		return ret
	}
	return *o.BankSwiftBic
}

// GetBankSwiftBicOk returns a tuple with the BankSwiftBic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetBankSwiftBicOk() (*string, bool) {
	if o == nil || IsNil(o.BankSwiftBic) {
		return nil, false
	}
	return o.BankSwiftBic, true
}

// HasBankSwiftBic returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasBankSwiftBic() bool {
	if o != nil && !IsNil(o.BankSwiftBic) {
		return true
	}

	return false
}

// SetBankSwiftBic gets a reference to the given string and assigns it to the BankSwiftBic field.
func (o *PaymentChannelResponseV4) SetBankSwiftBic(v string) {
	o.BankSwiftBic = &v
}

// GetBankAddress returns the BankAddress field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetBankAddress() AddressV4 {
	if o == nil || IsNil(o.BankAddress) {
		var ret AddressV4
		return ret
	}
	return *o.BankAddress
}

// GetBankAddressOk returns a tuple with the BankAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetBankAddressOk() (*AddressV4, bool) {
	if o == nil || IsNil(o.BankAddress) {
		return nil, false
	}
	return o.BankAddress, true
}

// HasBankAddress returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasBankAddress() bool {
	if o != nil && !IsNil(o.BankAddress) {
		return true
	}

	return false
}

// SetBankAddress gets a reference to the given AddressV4 and assigns it to the BankAddress field.
func (o *PaymentChannelResponseV4) SetBankAddress(v AddressV4) {
	o.BankAddress = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *PaymentChannelResponseV4) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PaymentChannelResponseV4) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDisabledReasonCode returns the DisabledReasonCode field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetDisabledReasonCode() string {
	if o == nil || IsNil(o.DisabledReasonCode) {
		var ret string
		return ret
	}
	return *o.DisabledReasonCode
}

// GetDisabledReasonCodeOk returns a tuple with the DisabledReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetDisabledReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DisabledReasonCode) {
		return nil, false
	}
	return o.DisabledReasonCode, true
}

// HasDisabledReasonCode returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasDisabledReasonCode() bool {
	if o != nil && !IsNil(o.DisabledReasonCode) {
		return true
	}

	return false
}

// SetDisabledReasonCode gets a reference to the given string and assigns it to the DisabledReasonCode field.
func (o *PaymentChannelResponseV4) SetDisabledReasonCode(v string) {
	o.DisabledReasonCode = &v
}

// GetDisabledReason returns the DisabledReason field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetDisabledReason() string {
	if o == nil || IsNil(o.DisabledReason) {
		var ret string
		return ret
	}
	return *o.DisabledReason
}

// GetDisabledReasonOk returns a tuple with the DisabledReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetDisabledReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DisabledReason) {
		return nil, false
	}
	return o.DisabledReason, true
}

// HasDisabledReason returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasDisabledReason() bool {
	if o != nil && !IsNil(o.DisabledReason) {
		return true
	}

	return false
}

// SetDisabledReason gets a reference to the given string and assigns it to the DisabledReason field.
func (o *PaymentChannelResponseV4) SetDisabledReason(v string) {
	o.DisabledReason = &v
}

// GetPayable returns the Payable field value if set, zero value otherwise.
func (o *PaymentChannelResponseV4) GetPayable() bool {
	if o == nil || IsNil(o.Payable) {
		var ret bool
		return ret
	}
	return *o.Payable
}

// GetPayableOk returns a tuple with the Payable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelResponseV4) GetPayableOk() (*bool, bool) {
	if o == nil || IsNil(o.Payable) {
		return nil, false
	}
	return o.Payable, true
}

// HasPayable returns a boolean if a field has been set.
func (o *PaymentChannelResponseV4) HasPayable() bool {
	if o != nil && !IsNil(o.Payable) {
		return true
	}

	return false
}

// SetPayable gets a reference to the given bool and assigns it to the Payable field.
func (o *PaymentChannelResponseV4) SetPayable(v bool) {
	o.Payable = &v
}

func (o PaymentChannelResponseV4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentChannelResponseV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.PayeeId) {
		toSerialize["payeeId"] = o.PayeeId
	}
	toSerialize["paymentChannelName"] = o.PaymentChannelName
	toSerialize["accountName"] = o.AccountName
	toSerialize["channelType"] = o.ChannelType
	toSerialize["countryCode"] = o.CountryCode
	if !IsNil(o.RoutingNumber) {
		toSerialize["routingNumber"] = o.RoutingNumber
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.PayorIds) {
		toSerialize["payorIds"] = o.PayorIds
	}
	if !IsNil(o.PayeeName) {
		toSerialize["payeeName"] = o.PayeeName
	}
	if !IsNil(o.BankName) {
		toSerialize["bankName"] = o.BankName
	}
	if !IsNil(o.BankSwiftBic) {
		toSerialize["bankSwiftBic"] = o.BankSwiftBic
	}
	if !IsNil(o.BankAddress) {
		toSerialize["bankAddress"] = o.BankAddress
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DisabledReasonCode) {
		toSerialize["disabledReasonCode"] = o.DisabledReasonCode
	}
	if !IsNil(o.DisabledReason) {
		toSerialize["disabledReason"] = o.DisabledReason
	}
	if !IsNil(o.Payable) {
		toSerialize["payable"] = o.Payable
	}
	return toSerialize, nil
}

type NullablePaymentChannelResponseV4 struct {
	value *PaymentChannelResponseV4
	isSet bool
}

func (v NullablePaymentChannelResponseV4) Get() *PaymentChannelResponseV4 {
	return v.value
}

func (v *NullablePaymentChannelResponseV4) Set(val *PaymentChannelResponseV4) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentChannelResponseV4) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentChannelResponseV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentChannelResponseV4(val *PaymentChannelResponseV4) *NullablePaymentChannelResponseV4 {
	return &NullablePaymentChannelResponseV4{value: val, isSet: true}
}

func (v NullablePaymentChannelResponseV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentChannelResponseV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


