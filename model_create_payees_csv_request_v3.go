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

// checks if the CreatePayeesCSVRequestV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePayeesCSVRequestV3{}

// CreatePayeesCSVRequestV3 struct for CreatePayeesCSVRequestV3
type CreatePayeesCSVRequestV3 struct {
	Type PayeeTypeEnum `json:"type"`
	RemoteId string `json:"remoteId"`
	Email string `json:"email"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	AddressLine3 *string `json:"addressLine3,omitempty"`
	AddressLine4 *string `json:"addressLine4,omitempty"`
	AddressCity string `json:"addressCity"`
	AddressCountyOrProvince *string `json:"addressCountyOrProvince,omitempty"`
	AddressZipOrPostcode string `json:"addressZipOrPostcode"`
	// Valid ISO 3166 2 character country code. See the <a href=\"https://www.iso.org/iso-3166-country-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	AddressCountry string `json:"addressCountry"`
	IndividualNationalIdentification *string `json:"individualNationalIdentification,omitempty"`
	// Must not be date in future. Example - 1970-05-20
	IndividualDateOfBirth *string `json:"individualDateOfBirth,omitempty"`
	IndividualTitle *string `json:"individualTitle,omitempty"`
	IndividualFirstName *string `json:"individualFirstName,omitempty"`
	IndividualOtherNames *string `json:"individualOtherNames,omitempty"`
	IndividualLastName *string `json:"individualLastName,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	CompanyEIN *string `json:"companyEIN,omitempty"`
	CompanyOperatingName *string `json:"companyOperatingName,omitempty"`
	// Either routing number and account number or only iban must be set
	PaymentChannelAccountNumber *string `json:"paymentChannelAccountNumber,omitempty"`
	// Either routing number and account number or only iban must be set
	PaymentChannelRoutingNumber *string `json:"paymentChannelRoutingNumber,omitempty"`
	PaymentChannelAccountName *string `json:"paymentChannelAccountName,omitempty"`
	// Must match the regular expression ```^[A-Za-z0-9]+$```.
	PaymentChannelIban *string `json:"paymentChannelIban,omitempty"`
	// Valid ISO 3166 2 character country code. See the <a href=\"https://www.iso.org/iso-3166-country-codes.html\" target=\"_blank\" a>ISO specification</a> for details.
	PaymentChannelCountryCode *string `json:"paymentChannelCountryCode,omitempty"`
	PaymentChannelCurrency *string `json:"paymentChannelCurrency,omitempty"`
	ChallengeDescription *string `json:"challengeDescription,omitempty"`
	ChallengeValue *string `json:"challengeValue,omitempty"`
	// An IETF BCP 47 language code which has been configured for use within this Velo environment.<BR> See the /v1/supportedLanguages endpoint to list the available codes for an environment. 
	PayeeLanguage *string `json:"payeeLanguage,omitempty"`
}

// NewCreatePayeesCSVRequestV3 instantiates a new CreatePayeesCSVRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePayeesCSVRequestV3(type_ PayeeTypeEnum, remoteId string, email string, addressLine1 string, addressCity string, addressZipOrPostcode string, addressCountry string) *CreatePayeesCSVRequestV3 {
	this := CreatePayeesCSVRequestV3{}
	this.Type = type_
	this.RemoteId = remoteId
	this.Email = email
	this.AddressLine1 = addressLine1
	this.AddressCity = addressCity
	this.AddressZipOrPostcode = addressZipOrPostcode
	this.AddressCountry = addressCountry
	return &this
}

// NewCreatePayeesCSVRequestV3WithDefaults instantiates a new CreatePayeesCSVRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePayeesCSVRequestV3WithDefaults() *CreatePayeesCSVRequestV3 {
	this := CreatePayeesCSVRequestV3{}
	return &this
}

// GetType returns the Type field value
func (o *CreatePayeesCSVRequestV3) GetType() PayeeTypeEnum {
	if o == nil {
		var ret PayeeTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetTypeOk() (*PayeeTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreatePayeesCSVRequestV3) SetType(v PayeeTypeEnum) {
	o.Type = v
}

// GetRemoteId returns the RemoteId field value
func (o *CreatePayeesCSVRequestV3) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetRemoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *CreatePayeesCSVRequestV3) SetRemoteId(v string) {
	o.RemoteId = v
}

// GetEmail returns the Email field value
func (o *CreatePayeesCSVRequestV3) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreatePayeesCSVRequestV3) SetEmail(v string) {
	o.Email = v
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *CreatePayeesCSVRequestV3) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *CreatePayeesCSVRequestV3) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *CreatePayeesCSVRequestV3) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *CreatePayeesCSVRequestV3) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetAddressLine4 returns the AddressLine4 field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetAddressLine4() string {
	if o == nil || IsNil(o.AddressLine4) {
		var ret string
		return ret
	}
	return *o.AddressLine4
}

// GetAddressLine4Ok returns a tuple with the AddressLine4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetAddressLine4Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine4) {
		return nil, false
	}
	return o.AddressLine4, true
}

// HasAddressLine4 returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasAddressLine4() bool {
	if o != nil && !IsNil(o.AddressLine4) {
		return true
	}

	return false
}

// SetAddressLine4 gets a reference to the given string and assigns it to the AddressLine4 field.
func (o *CreatePayeesCSVRequestV3) SetAddressLine4(v string) {
	o.AddressLine4 = &v
}

// GetAddressCity returns the AddressCity field value
func (o *CreatePayeesCSVRequestV3) GetAddressCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressCity
}

// GetAddressCityOk returns a tuple with the AddressCity field value
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetAddressCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressCity, true
}

// SetAddressCity sets field value
func (o *CreatePayeesCSVRequestV3) SetAddressCity(v string) {
	o.AddressCity = v
}

// GetAddressCountyOrProvince returns the AddressCountyOrProvince field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetAddressCountyOrProvince() string {
	if o == nil || IsNil(o.AddressCountyOrProvince) {
		var ret string
		return ret
	}
	return *o.AddressCountyOrProvince
}

// GetAddressCountyOrProvinceOk returns a tuple with the AddressCountyOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetAddressCountyOrProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.AddressCountyOrProvince) {
		return nil, false
	}
	return o.AddressCountyOrProvince, true
}

// HasAddressCountyOrProvince returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasAddressCountyOrProvince() bool {
	if o != nil && !IsNil(o.AddressCountyOrProvince) {
		return true
	}

	return false
}

// SetAddressCountyOrProvince gets a reference to the given string and assigns it to the AddressCountyOrProvince field.
func (o *CreatePayeesCSVRequestV3) SetAddressCountyOrProvince(v string) {
	o.AddressCountyOrProvince = &v
}

// GetAddressZipOrPostcode returns the AddressZipOrPostcode field value
func (o *CreatePayeesCSVRequestV3) GetAddressZipOrPostcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressZipOrPostcode
}

// GetAddressZipOrPostcodeOk returns a tuple with the AddressZipOrPostcode field value
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetAddressZipOrPostcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressZipOrPostcode, true
}

// SetAddressZipOrPostcode sets field value
func (o *CreatePayeesCSVRequestV3) SetAddressZipOrPostcode(v string) {
	o.AddressZipOrPostcode = v
}

// GetAddressCountry returns the AddressCountry field value
func (o *CreatePayeesCSVRequestV3) GetAddressCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressCountry
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetAddressCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressCountry, true
}

// SetAddressCountry sets field value
func (o *CreatePayeesCSVRequestV3) SetAddressCountry(v string) {
	o.AddressCountry = v
}

// GetIndividualNationalIdentification returns the IndividualNationalIdentification field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetIndividualNationalIdentification() string {
	if o == nil || IsNil(o.IndividualNationalIdentification) {
		var ret string
		return ret
	}
	return *o.IndividualNationalIdentification
}

// GetIndividualNationalIdentificationOk returns a tuple with the IndividualNationalIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetIndividualNationalIdentificationOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualNationalIdentification) {
		return nil, false
	}
	return o.IndividualNationalIdentification, true
}

// HasIndividualNationalIdentification returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasIndividualNationalIdentification() bool {
	if o != nil && !IsNil(o.IndividualNationalIdentification) {
		return true
	}

	return false
}

// SetIndividualNationalIdentification gets a reference to the given string and assigns it to the IndividualNationalIdentification field.
func (o *CreatePayeesCSVRequestV3) SetIndividualNationalIdentification(v string) {
	o.IndividualNationalIdentification = &v
}

// GetIndividualDateOfBirth returns the IndividualDateOfBirth field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetIndividualDateOfBirth() string {
	if o == nil || IsNil(o.IndividualDateOfBirth) {
		var ret string
		return ret
	}
	return *o.IndividualDateOfBirth
}

// GetIndividualDateOfBirthOk returns a tuple with the IndividualDateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetIndividualDateOfBirthOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualDateOfBirth) {
		return nil, false
	}
	return o.IndividualDateOfBirth, true
}

// HasIndividualDateOfBirth returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasIndividualDateOfBirth() bool {
	if o != nil && !IsNil(o.IndividualDateOfBirth) {
		return true
	}

	return false
}

// SetIndividualDateOfBirth gets a reference to the given string and assigns it to the IndividualDateOfBirth field.
func (o *CreatePayeesCSVRequestV3) SetIndividualDateOfBirth(v string) {
	o.IndividualDateOfBirth = &v
}

// GetIndividualTitle returns the IndividualTitle field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetIndividualTitle() string {
	if o == nil || IsNil(o.IndividualTitle) {
		var ret string
		return ret
	}
	return *o.IndividualTitle
}

// GetIndividualTitleOk returns a tuple with the IndividualTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetIndividualTitleOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualTitle) {
		return nil, false
	}
	return o.IndividualTitle, true
}

// HasIndividualTitle returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasIndividualTitle() bool {
	if o != nil && !IsNil(o.IndividualTitle) {
		return true
	}

	return false
}

// SetIndividualTitle gets a reference to the given string and assigns it to the IndividualTitle field.
func (o *CreatePayeesCSVRequestV3) SetIndividualTitle(v string) {
	o.IndividualTitle = &v
}

// GetIndividualFirstName returns the IndividualFirstName field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetIndividualFirstName() string {
	if o == nil || IsNil(o.IndividualFirstName) {
		var ret string
		return ret
	}
	return *o.IndividualFirstName
}

// GetIndividualFirstNameOk returns a tuple with the IndividualFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetIndividualFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualFirstName) {
		return nil, false
	}
	return o.IndividualFirstName, true
}

// HasIndividualFirstName returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasIndividualFirstName() bool {
	if o != nil && !IsNil(o.IndividualFirstName) {
		return true
	}

	return false
}

// SetIndividualFirstName gets a reference to the given string and assigns it to the IndividualFirstName field.
func (o *CreatePayeesCSVRequestV3) SetIndividualFirstName(v string) {
	o.IndividualFirstName = &v
}

// GetIndividualOtherNames returns the IndividualOtherNames field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetIndividualOtherNames() string {
	if o == nil || IsNil(o.IndividualOtherNames) {
		var ret string
		return ret
	}
	return *o.IndividualOtherNames
}

// GetIndividualOtherNamesOk returns a tuple with the IndividualOtherNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetIndividualOtherNamesOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualOtherNames) {
		return nil, false
	}
	return o.IndividualOtherNames, true
}

// HasIndividualOtherNames returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasIndividualOtherNames() bool {
	if o != nil && !IsNil(o.IndividualOtherNames) {
		return true
	}

	return false
}

// SetIndividualOtherNames gets a reference to the given string and assigns it to the IndividualOtherNames field.
func (o *CreatePayeesCSVRequestV3) SetIndividualOtherNames(v string) {
	o.IndividualOtherNames = &v
}

// GetIndividualLastName returns the IndividualLastName field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetIndividualLastName() string {
	if o == nil || IsNil(o.IndividualLastName) {
		var ret string
		return ret
	}
	return *o.IndividualLastName
}

// GetIndividualLastNameOk returns a tuple with the IndividualLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetIndividualLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualLastName) {
		return nil, false
	}
	return o.IndividualLastName, true
}

// HasIndividualLastName returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasIndividualLastName() bool {
	if o != nil && !IsNil(o.IndividualLastName) {
		return true
	}

	return false
}

// SetIndividualLastName gets a reference to the given string and assigns it to the IndividualLastName field.
func (o *CreatePayeesCSVRequestV3) SetIndividualLastName(v string) {
	o.IndividualLastName = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *CreatePayeesCSVRequestV3) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCompanyEIN returns the CompanyEIN field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetCompanyEIN() string {
	if o == nil || IsNil(o.CompanyEIN) {
		var ret string
		return ret
	}
	return *o.CompanyEIN
}

// GetCompanyEINOk returns a tuple with the CompanyEIN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetCompanyEINOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyEIN) {
		return nil, false
	}
	return o.CompanyEIN, true
}

// HasCompanyEIN returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasCompanyEIN() bool {
	if o != nil && !IsNil(o.CompanyEIN) {
		return true
	}

	return false
}

// SetCompanyEIN gets a reference to the given string and assigns it to the CompanyEIN field.
func (o *CreatePayeesCSVRequestV3) SetCompanyEIN(v string) {
	o.CompanyEIN = &v
}

// GetCompanyOperatingName returns the CompanyOperatingName field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetCompanyOperatingName() string {
	if o == nil || IsNil(o.CompanyOperatingName) {
		var ret string
		return ret
	}
	return *o.CompanyOperatingName
}

// GetCompanyOperatingNameOk returns a tuple with the CompanyOperatingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetCompanyOperatingNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyOperatingName) {
		return nil, false
	}
	return o.CompanyOperatingName, true
}

// HasCompanyOperatingName returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasCompanyOperatingName() bool {
	if o != nil && !IsNil(o.CompanyOperatingName) {
		return true
	}

	return false
}

// SetCompanyOperatingName gets a reference to the given string and assigns it to the CompanyOperatingName field.
func (o *CreatePayeesCSVRequestV3) SetCompanyOperatingName(v string) {
	o.CompanyOperatingName = &v
}

// GetPaymentChannelAccountNumber returns the PaymentChannelAccountNumber field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelAccountNumber() string {
	if o == nil || IsNil(o.PaymentChannelAccountNumber) {
		var ret string
		return ret
	}
	return *o.PaymentChannelAccountNumber
}

// GetPaymentChannelAccountNumberOk returns a tuple with the PaymentChannelAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelAccountNumber) {
		return nil, false
	}
	return o.PaymentChannelAccountNumber, true
}

// HasPaymentChannelAccountNumber returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasPaymentChannelAccountNumber() bool {
	if o != nil && !IsNil(o.PaymentChannelAccountNumber) {
		return true
	}

	return false
}

// SetPaymentChannelAccountNumber gets a reference to the given string and assigns it to the PaymentChannelAccountNumber field.
func (o *CreatePayeesCSVRequestV3) SetPaymentChannelAccountNumber(v string) {
	o.PaymentChannelAccountNumber = &v
}

// GetPaymentChannelRoutingNumber returns the PaymentChannelRoutingNumber field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelRoutingNumber() string {
	if o == nil || IsNil(o.PaymentChannelRoutingNumber) {
		var ret string
		return ret
	}
	return *o.PaymentChannelRoutingNumber
}

// GetPaymentChannelRoutingNumberOk returns a tuple with the PaymentChannelRoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelRoutingNumber) {
		return nil, false
	}
	return o.PaymentChannelRoutingNumber, true
}

// HasPaymentChannelRoutingNumber returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasPaymentChannelRoutingNumber() bool {
	if o != nil && !IsNil(o.PaymentChannelRoutingNumber) {
		return true
	}

	return false
}

// SetPaymentChannelRoutingNumber gets a reference to the given string and assigns it to the PaymentChannelRoutingNumber field.
func (o *CreatePayeesCSVRequestV3) SetPaymentChannelRoutingNumber(v string) {
	o.PaymentChannelRoutingNumber = &v
}

// GetPaymentChannelAccountName returns the PaymentChannelAccountName field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelAccountName() string {
	if o == nil || IsNil(o.PaymentChannelAccountName) {
		var ret string
		return ret
	}
	return *o.PaymentChannelAccountName
}

// GetPaymentChannelAccountNameOk returns a tuple with the PaymentChannelAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelAccountName) {
		return nil, false
	}
	return o.PaymentChannelAccountName, true
}

// HasPaymentChannelAccountName returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasPaymentChannelAccountName() bool {
	if o != nil && !IsNil(o.PaymentChannelAccountName) {
		return true
	}

	return false
}

// SetPaymentChannelAccountName gets a reference to the given string and assigns it to the PaymentChannelAccountName field.
func (o *CreatePayeesCSVRequestV3) SetPaymentChannelAccountName(v string) {
	o.PaymentChannelAccountName = &v
}

// GetPaymentChannelIban returns the PaymentChannelIban field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelIban() string {
	if o == nil || IsNil(o.PaymentChannelIban) {
		var ret string
		return ret
	}
	return *o.PaymentChannelIban
}

// GetPaymentChannelIbanOk returns a tuple with the PaymentChannelIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelIbanOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelIban) {
		return nil, false
	}
	return o.PaymentChannelIban, true
}

// HasPaymentChannelIban returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasPaymentChannelIban() bool {
	if o != nil && !IsNil(o.PaymentChannelIban) {
		return true
	}

	return false
}

// SetPaymentChannelIban gets a reference to the given string and assigns it to the PaymentChannelIban field.
func (o *CreatePayeesCSVRequestV3) SetPaymentChannelIban(v string) {
	o.PaymentChannelIban = &v
}

// GetPaymentChannelCountryCode returns the PaymentChannelCountryCode field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelCountryCode() string {
	if o == nil || IsNil(o.PaymentChannelCountryCode) {
		var ret string
		return ret
	}
	return *o.PaymentChannelCountryCode
}

// GetPaymentChannelCountryCodeOk returns a tuple with the PaymentChannelCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelCountryCode) {
		return nil, false
	}
	return o.PaymentChannelCountryCode, true
}

// HasPaymentChannelCountryCode returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasPaymentChannelCountryCode() bool {
	if o != nil && !IsNil(o.PaymentChannelCountryCode) {
		return true
	}

	return false
}

// SetPaymentChannelCountryCode gets a reference to the given string and assigns it to the PaymentChannelCountryCode field.
func (o *CreatePayeesCSVRequestV3) SetPaymentChannelCountryCode(v string) {
	o.PaymentChannelCountryCode = &v
}

// GetPaymentChannelCurrency returns the PaymentChannelCurrency field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelCurrency() string {
	if o == nil || IsNil(o.PaymentChannelCurrency) {
		var ret string
		return ret
	}
	return *o.PaymentChannelCurrency
}

// GetPaymentChannelCurrencyOk returns a tuple with the PaymentChannelCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetPaymentChannelCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentChannelCurrency) {
		return nil, false
	}
	return o.PaymentChannelCurrency, true
}

// HasPaymentChannelCurrency returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasPaymentChannelCurrency() bool {
	if o != nil && !IsNil(o.PaymentChannelCurrency) {
		return true
	}

	return false
}

// SetPaymentChannelCurrency gets a reference to the given string and assigns it to the PaymentChannelCurrency field.
func (o *CreatePayeesCSVRequestV3) SetPaymentChannelCurrency(v string) {
	o.PaymentChannelCurrency = &v
}

// GetChallengeDescription returns the ChallengeDescription field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetChallengeDescription() string {
	if o == nil || IsNil(o.ChallengeDescription) {
		var ret string
		return ret
	}
	return *o.ChallengeDescription
}

// GetChallengeDescriptionOk returns a tuple with the ChallengeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetChallengeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ChallengeDescription) {
		return nil, false
	}
	return o.ChallengeDescription, true
}

// HasChallengeDescription returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasChallengeDescription() bool {
	if o != nil && !IsNil(o.ChallengeDescription) {
		return true
	}

	return false
}

// SetChallengeDescription gets a reference to the given string and assigns it to the ChallengeDescription field.
func (o *CreatePayeesCSVRequestV3) SetChallengeDescription(v string) {
	o.ChallengeDescription = &v
}

// GetChallengeValue returns the ChallengeValue field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetChallengeValue() string {
	if o == nil || IsNil(o.ChallengeValue) {
		var ret string
		return ret
	}
	return *o.ChallengeValue
}

// GetChallengeValueOk returns a tuple with the ChallengeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetChallengeValueOk() (*string, bool) {
	if o == nil || IsNil(o.ChallengeValue) {
		return nil, false
	}
	return o.ChallengeValue, true
}

// HasChallengeValue returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasChallengeValue() bool {
	if o != nil && !IsNil(o.ChallengeValue) {
		return true
	}

	return false
}

// SetChallengeValue gets a reference to the given string and assigns it to the ChallengeValue field.
func (o *CreatePayeesCSVRequestV3) SetChallengeValue(v string) {
	o.ChallengeValue = &v
}

// GetPayeeLanguage returns the PayeeLanguage field value if set, zero value otherwise.
func (o *CreatePayeesCSVRequestV3) GetPayeeLanguage() string {
	if o == nil || IsNil(o.PayeeLanguage) {
		var ret string
		return ret
	}
	return *o.PayeeLanguage
}

// GetPayeeLanguageOk returns a tuple with the PayeeLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayeesCSVRequestV3) GetPayeeLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeLanguage) {
		return nil, false
	}
	return o.PayeeLanguage, true
}

// HasPayeeLanguage returns a boolean if a field has been set.
func (o *CreatePayeesCSVRequestV3) HasPayeeLanguage() bool {
	if o != nil && !IsNil(o.PayeeLanguage) {
		return true
	}

	return false
}

// SetPayeeLanguage gets a reference to the given string and assigns it to the PayeeLanguage field.
func (o *CreatePayeesCSVRequestV3) SetPayeeLanguage(v string) {
	o.PayeeLanguage = &v
}

func (o CreatePayeesCSVRequestV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePayeesCSVRequestV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["remoteId"] = o.RemoteId
	toSerialize["email"] = o.Email
	toSerialize["addressLine1"] = o.AddressLine1
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["addressLine3"] = o.AddressLine3
	}
	if !IsNil(o.AddressLine4) {
		toSerialize["addressLine4"] = o.AddressLine4
	}
	toSerialize["addressCity"] = o.AddressCity
	if !IsNil(o.AddressCountyOrProvince) {
		toSerialize["addressCountyOrProvince"] = o.AddressCountyOrProvince
	}
	toSerialize["addressZipOrPostcode"] = o.AddressZipOrPostcode
	toSerialize["addressCountry"] = o.AddressCountry
	if !IsNil(o.IndividualNationalIdentification) {
		toSerialize["individualNationalIdentification"] = o.IndividualNationalIdentification
	}
	if !IsNil(o.IndividualDateOfBirth) {
		toSerialize["individualDateOfBirth"] = o.IndividualDateOfBirth
	}
	if !IsNil(o.IndividualTitle) {
		toSerialize["individualTitle"] = o.IndividualTitle
	}
	if !IsNil(o.IndividualFirstName) {
		toSerialize["individualFirstName"] = o.IndividualFirstName
	}
	if !IsNil(o.IndividualOtherNames) {
		toSerialize["individualOtherNames"] = o.IndividualOtherNames
	}
	if !IsNil(o.IndividualLastName) {
		toSerialize["individualLastName"] = o.IndividualLastName
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.CompanyEIN) {
		toSerialize["companyEIN"] = o.CompanyEIN
	}
	if !IsNil(o.CompanyOperatingName) {
		toSerialize["companyOperatingName"] = o.CompanyOperatingName
	}
	if !IsNil(o.PaymentChannelAccountNumber) {
		toSerialize["paymentChannelAccountNumber"] = o.PaymentChannelAccountNumber
	}
	if !IsNil(o.PaymentChannelRoutingNumber) {
		toSerialize["paymentChannelRoutingNumber"] = o.PaymentChannelRoutingNumber
	}
	if !IsNil(o.PaymentChannelAccountName) {
		toSerialize["paymentChannelAccountName"] = o.PaymentChannelAccountName
	}
	if !IsNil(o.PaymentChannelIban) {
		toSerialize["paymentChannelIban"] = o.PaymentChannelIban
	}
	if !IsNil(o.PaymentChannelCountryCode) {
		toSerialize["paymentChannelCountryCode"] = o.PaymentChannelCountryCode
	}
	if !IsNil(o.PaymentChannelCurrency) {
		toSerialize["paymentChannelCurrency"] = o.PaymentChannelCurrency
	}
	if !IsNil(o.ChallengeDescription) {
		toSerialize["challengeDescription"] = o.ChallengeDescription
	}
	if !IsNil(o.ChallengeValue) {
		toSerialize["challengeValue"] = o.ChallengeValue
	}
	if !IsNil(o.PayeeLanguage) {
		toSerialize["payeeLanguage"] = o.PayeeLanguage
	}
	return toSerialize, nil
}

type NullableCreatePayeesCSVRequestV3 struct {
	value *CreatePayeesCSVRequestV3
	isSet bool
}

func (v NullableCreatePayeesCSVRequestV3) Get() *CreatePayeesCSVRequestV3 {
	return v.value
}

func (v *NullableCreatePayeesCSVRequestV3) Set(val *CreatePayeesCSVRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePayeesCSVRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePayeesCSVRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePayeesCSVRequestV3(val *CreatePayeesCSVRequestV3) *NullableCreatePayeesCSVRequestV3 {
	return &NullableCreatePayeesCSVRequestV3{value: val, isSet: true}
}

func (v NullableCreatePayeesCSVRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePayeesCSVRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


