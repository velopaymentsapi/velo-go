# CreatePayeesCSVRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PayeeTypeEnum**](PayeeTypeEnum.md) |  | 
**RemoteId** | **string** |  | 
**Email** | **string** |  | 
**AddressLine1** | **string** |  | 
**AddressLine2** | Pointer to **string** |  | [optional] 
**AddressLine3** | Pointer to **string** |  | [optional] 
**AddressLine4** | Pointer to **string** |  | [optional] 
**AddressCity** | **string** |  | 
**AddressCountyOrProvince** | Pointer to **string** |  | [optional] 
**AddressZipOrPostcode** | **string** |  | 
**AddressCountry** | **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**IndividualNationalIdentification** | Pointer to **string** |  | [optional] 
**IndividualDateOfBirth** | Pointer to **string** | Must not be date in future. Example - 1970-05-20 | [optional] 
**IndividualTitle** | Pointer to **string** |  | [optional] 
**IndividualFirstName** | Pointer to **string** |  | [optional] 
**IndividualOtherNames** | Pointer to **string** |  | [optional] 
**IndividualLastName** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CompanyEIN** | Pointer to **string** |  | [optional] 
**CompanyOperatingName** | Pointer to **string** |  | [optional] 
**PaymentChannelAccountNumber** | Pointer to **string** | Either routing number and account number or only iban must be set | [optional] 
**PaymentChannelRoutingNumber** | Pointer to **string** | Either routing number and account number or only iban must be set | [optional] 
**PaymentChannelAccountName** | Pointer to **string** |  | [optional] 
**PaymentChannelIban** | Pointer to **string** | Must match the regular expression &#x60;&#x60;&#x60;^[A-Za-z0-9]+$&#x60;&#x60;&#x60;. | [optional] 
**PaymentChannelCountryCode** | Pointer to **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**PaymentChannelCurrency** | Pointer to **string** |  | [optional] 
**ChallengeDescription** | Pointer to **string** |  | [optional] 
**ChallengeValue** | Pointer to **string** |  | [optional] 
**PayeeLanguage** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 

## Methods

### NewCreatePayeesCSVRequestV3

`func NewCreatePayeesCSVRequestV3(type_ PayeeTypeEnum, remoteId string, email string, addressLine1 string, addressCity string, addressZipOrPostcode string, addressCountry string, ) *CreatePayeesCSVRequestV3`

NewCreatePayeesCSVRequestV3 instantiates a new CreatePayeesCSVRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesCSVRequestV3WithDefaults

`func NewCreatePayeesCSVRequestV3WithDefaults() *CreatePayeesCSVRequestV3`

NewCreatePayeesCSVRequestV3WithDefaults instantiates a new CreatePayeesCSVRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreatePayeesCSVRequestV3) GetType() PayeeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePayeesCSVRequestV3) GetTypeOk() (*PayeeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePayeesCSVRequestV3) SetType(v PayeeTypeEnum)`

SetType sets Type field to given value.


### GetRemoteId

`func (o *CreatePayeesCSVRequestV3) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreatePayeesCSVRequestV3) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreatePayeesCSVRequestV3) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetEmail

`func (o *CreatePayeesCSVRequestV3) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreatePayeesCSVRequestV3) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreatePayeesCSVRequestV3) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAddressLine1

`func (o *CreatePayeesCSVRequestV3) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CreatePayeesCSVRequestV3) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CreatePayeesCSVRequestV3) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *CreatePayeesCSVRequestV3) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CreatePayeesCSVRequestV3) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CreatePayeesCSVRequestV3) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CreatePayeesCSVRequestV3) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *CreatePayeesCSVRequestV3) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *CreatePayeesCSVRequestV3) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *CreatePayeesCSVRequestV3) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *CreatePayeesCSVRequestV3) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetAddressLine4

`func (o *CreatePayeesCSVRequestV3) GetAddressLine4() string`

GetAddressLine4 returns the AddressLine4 field if non-nil, zero value otherwise.

### GetAddressLine4Ok

`func (o *CreatePayeesCSVRequestV3) GetAddressLine4Ok() (*string, bool)`

GetAddressLine4Ok returns a tuple with the AddressLine4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine4

`func (o *CreatePayeesCSVRequestV3) SetAddressLine4(v string)`

SetAddressLine4 sets AddressLine4 field to given value.

### HasAddressLine4

`func (o *CreatePayeesCSVRequestV3) HasAddressLine4() bool`

HasAddressLine4 returns a boolean if a field has been set.

### GetAddressCity

`func (o *CreatePayeesCSVRequestV3) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *CreatePayeesCSVRequestV3) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *CreatePayeesCSVRequestV3) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.


### GetAddressCountyOrProvince

`func (o *CreatePayeesCSVRequestV3) GetAddressCountyOrProvince() string`

GetAddressCountyOrProvince returns the AddressCountyOrProvince field if non-nil, zero value otherwise.

### GetAddressCountyOrProvinceOk

`func (o *CreatePayeesCSVRequestV3) GetAddressCountyOrProvinceOk() (*string, bool)`

GetAddressCountyOrProvinceOk returns a tuple with the AddressCountyOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountyOrProvince

`func (o *CreatePayeesCSVRequestV3) SetAddressCountyOrProvince(v string)`

SetAddressCountyOrProvince sets AddressCountyOrProvince field to given value.

### HasAddressCountyOrProvince

`func (o *CreatePayeesCSVRequestV3) HasAddressCountyOrProvince() bool`

HasAddressCountyOrProvince returns a boolean if a field has been set.

### GetAddressZipOrPostcode

`func (o *CreatePayeesCSVRequestV3) GetAddressZipOrPostcode() string`

GetAddressZipOrPostcode returns the AddressZipOrPostcode field if non-nil, zero value otherwise.

### GetAddressZipOrPostcodeOk

`func (o *CreatePayeesCSVRequestV3) GetAddressZipOrPostcodeOk() (*string, bool)`

GetAddressZipOrPostcodeOk returns a tuple with the AddressZipOrPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZipOrPostcode

`func (o *CreatePayeesCSVRequestV3) SetAddressZipOrPostcode(v string)`

SetAddressZipOrPostcode sets AddressZipOrPostcode field to given value.


### GetAddressCountry

`func (o *CreatePayeesCSVRequestV3) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *CreatePayeesCSVRequestV3) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *CreatePayeesCSVRequestV3) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.


### GetIndividualNationalIdentification

`func (o *CreatePayeesCSVRequestV3) GetIndividualNationalIdentification() string`

GetIndividualNationalIdentification returns the IndividualNationalIdentification field if non-nil, zero value otherwise.

### GetIndividualNationalIdentificationOk

`func (o *CreatePayeesCSVRequestV3) GetIndividualNationalIdentificationOk() (*string, bool)`

GetIndividualNationalIdentificationOk returns a tuple with the IndividualNationalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualNationalIdentification

`func (o *CreatePayeesCSVRequestV3) SetIndividualNationalIdentification(v string)`

SetIndividualNationalIdentification sets IndividualNationalIdentification field to given value.

### HasIndividualNationalIdentification

`func (o *CreatePayeesCSVRequestV3) HasIndividualNationalIdentification() bool`

HasIndividualNationalIdentification returns a boolean if a field has been set.

### GetIndividualDateOfBirth

`func (o *CreatePayeesCSVRequestV3) GetIndividualDateOfBirth() string`

GetIndividualDateOfBirth returns the IndividualDateOfBirth field if non-nil, zero value otherwise.

### GetIndividualDateOfBirthOk

`func (o *CreatePayeesCSVRequestV3) GetIndividualDateOfBirthOk() (*string, bool)`

GetIndividualDateOfBirthOk returns a tuple with the IndividualDateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualDateOfBirth

`func (o *CreatePayeesCSVRequestV3) SetIndividualDateOfBirth(v string)`

SetIndividualDateOfBirth sets IndividualDateOfBirth field to given value.

### HasIndividualDateOfBirth

`func (o *CreatePayeesCSVRequestV3) HasIndividualDateOfBirth() bool`

HasIndividualDateOfBirth returns a boolean if a field has been set.

### GetIndividualTitle

`func (o *CreatePayeesCSVRequestV3) GetIndividualTitle() string`

GetIndividualTitle returns the IndividualTitle field if non-nil, zero value otherwise.

### GetIndividualTitleOk

`func (o *CreatePayeesCSVRequestV3) GetIndividualTitleOk() (*string, bool)`

GetIndividualTitleOk returns a tuple with the IndividualTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualTitle

`func (o *CreatePayeesCSVRequestV3) SetIndividualTitle(v string)`

SetIndividualTitle sets IndividualTitle field to given value.

### HasIndividualTitle

`func (o *CreatePayeesCSVRequestV3) HasIndividualTitle() bool`

HasIndividualTitle returns a boolean if a field has been set.

### GetIndividualFirstName

`func (o *CreatePayeesCSVRequestV3) GetIndividualFirstName() string`

GetIndividualFirstName returns the IndividualFirstName field if non-nil, zero value otherwise.

### GetIndividualFirstNameOk

`func (o *CreatePayeesCSVRequestV3) GetIndividualFirstNameOk() (*string, bool)`

GetIndividualFirstNameOk returns a tuple with the IndividualFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualFirstName

`func (o *CreatePayeesCSVRequestV3) SetIndividualFirstName(v string)`

SetIndividualFirstName sets IndividualFirstName field to given value.

### HasIndividualFirstName

`func (o *CreatePayeesCSVRequestV3) HasIndividualFirstName() bool`

HasIndividualFirstName returns a boolean if a field has been set.

### GetIndividualOtherNames

`func (o *CreatePayeesCSVRequestV3) GetIndividualOtherNames() string`

GetIndividualOtherNames returns the IndividualOtherNames field if non-nil, zero value otherwise.

### GetIndividualOtherNamesOk

`func (o *CreatePayeesCSVRequestV3) GetIndividualOtherNamesOk() (*string, bool)`

GetIndividualOtherNamesOk returns a tuple with the IndividualOtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualOtherNames

`func (o *CreatePayeesCSVRequestV3) SetIndividualOtherNames(v string)`

SetIndividualOtherNames sets IndividualOtherNames field to given value.

### HasIndividualOtherNames

`func (o *CreatePayeesCSVRequestV3) HasIndividualOtherNames() bool`

HasIndividualOtherNames returns a boolean if a field has been set.

### GetIndividualLastName

`func (o *CreatePayeesCSVRequestV3) GetIndividualLastName() string`

GetIndividualLastName returns the IndividualLastName field if non-nil, zero value otherwise.

### GetIndividualLastNameOk

`func (o *CreatePayeesCSVRequestV3) GetIndividualLastNameOk() (*string, bool)`

GetIndividualLastNameOk returns a tuple with the IndividualLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualLastName

`func (o *CreatePayeesCSVRequestV3) SetIndividualLastName(v string)`

SetIndividualLastName sets IndividualLastName field to given value.

### HasIndividualLastName

`func (o *CreatePayeesCSVRequestV3) HasIndividualLastName() bool`

HasIndividualLastName returns a boolean if a field has been set.

### GetCompanyName

`func (o *CreatePayeesCSVRequestV3) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CreatePayeesCSVRequestV3) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CreatePayeesCSVRequestV3) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CreatePayeesCSVRequestV3) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyEIN

`func (o *CreatePayeesCSVRequestV3) GetCompanyEIN() string`

GetCompanyEIN returns the CompanyEIN field if non-nil, zero value otherwise.

### GetCompanyEINOk

`func (o *CreatePayeesCSVRequestV3) GetCompanyEINOk() (*string, bool)`

GetCompanyEINOk returns a tuple with the CompanyEIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEIN

`func (o *CreatePayeesCSVRequestV3) SetCompanyEIN(v string)`

SetCompanyEIN sets CompanyEIN field to given value.

### HasCompanyEIN

`func (o *CreatePayeesCSVRequestV3) HasCompanyEIN() bool`

HasCompanyEIN returns a boolean if a field has been set.

### GetCompanyOperatingName

`func (o *CreatePayeesCSVRequestV3) GetCompanyOperatingName() string`

GetCompanyOperatingName returns the CompanyOperatingName field if non-nil, zero value otherwise.

### GetCompanyOperatingNameOk

`func (o *CreatePayeesCSVRequestV3) GetCompanyOperatingNameOk() (*string, bool)`

GetCompanyOperatingNameOk returns a tuple with the CompanyOperatingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyOperatingName

`func (o *CreatePayeesCSVRequestV3) SetCompanyOperatingName(v string)`

SetCompanyOperatingName sets CompanyOperatingName field to given value.

### HasCompanyOperatingName

`func (o *CreatePayeesCSVRequestV3) HasCompanyOperatingName() bool`

HasCompanyOperatingName returns a boolean if a field has been set.

### GetPaymentChannelAccountNumber

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelAccountNumber() string`

GetPaymentChannelAccountNumber returns the PaymentChannelAccountNumber field if non-nil, zero value otherwise.

### GetPaymentChannelAccountNumberOk

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelAccountNumberOk() (*string, bool)`

GetPaymentChannelAccountNumberOk returns a tuple with the PaymentChannelAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelAccountNumber

`func (o *CreatePayeesCSVRequestV3) SetPaymentChannelAccountNumber(v string)`

SetPaymentChannelAccountNumber sets PaymentChannelAccountNumber field to given value.

### HasPaymentChannelAccountNumber

`func (o *CreatePayeesCSVRequestV3) HasPaymentChannelAccountNumber() bool`

HasPaymentChannelAccountNumber returns a boolean if a field has been set.

### GetPaymentChannelRoutingNumber

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelRoutingNumber() string`

GetPaymentChannelRoutingNumber returns the PaymentChannelRoutingNumber field if non-nil, zero value otherwise.

### GetPaymentChannelRoutingNumberOk

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelRoutingNumberOk() (*string, bool)`

GetPaymentChannelRoutingNumberOk returns a tuple with the PaymentChannelRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelRoutingNumber

`func (o *CreatePayeesCSVRequestV3) SetPaymentChannelRoutingNumber(v string)`

SetPaymentChannelRoutingNumber sets PaymentChannelRoutingNumber field to given value.

### HasPaymentChannelRoutingNumber

`func (o *CreatePayeesCSVRequestV3) HasPaymentChannelRoutingNumber() bool`

HasPaymentChannelRoutingNumber returns a boolean if a field has been set.

### GetPaymentChannelAccountName

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelAccountName() string`

GetPaymentChannelAccountName returns the PaymentChannelAccountName field if non-nil, zero value otherwise.

### GetPaymentChannelAccountNameOk

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelAccountNameOk() (*string, bool)`

GetPaymentChannelAccountNameOk returns a tuple with the PaymentChannelAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelAccountName

`func (o *CreatePayeesCSVRequestV3) SetPaymentChannelAccountName(v string)`

SetPaymentChannelAccountName sets PaymentChannelAccountName field to given value.

### HasPaymentChannelAccountName

`func (o *CreatePayeesCSVRequestV3) HasPaymentChannelAccountName() bool`

HasPaymentChannelAccountName returns a boolean if a field has been set.

### GetPaymentChannelIban

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelIban() string`

GetPaymentChannelIban returns the PaymentChannelIban field if non-nil, zero value otherwise.

### GetPaymentChannelIbanOk

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelIbanOk() (*string, bool)`

GetPaymentChannelIbanOk returns a tuple with the PaymentChannelIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelIban

`func (o *CreatePayeesCSVRequestV3) SetPaymentChannelIban(v string)`

SetPaymentChannelIban sets PaymentChannelIban field to given value.

### HasPaymentChannelIban

`func (o *CreatePayeesCSVRequestV3) HasPaymentChannelIban() bool`

HasPaymentChannelIban returns a boolean if a field has been set.

### GetPaymentChannelCountryCode

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelCountryCode() string`

GetPaymentChannelCountryCode returns the PaymentChannelCountryCode field if non-nil, zero value otherwise.

### GetPaymentChannelCountryCodeOk

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelCountryCodeOk() (*string, bool)`

GetPaymentChannelCountryCodeOk returns a tuple with the PaymentChannelCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelCountryCode

`func (o *CreatePayeesCSVRequestV3) SetPaymentChannelCountryCode(v string)`

SetPaymentChannelCountryCode sets PaymentChannelCountryCode field to given value.

### HasPaymentChannelCountryCode

`func (o *CreatePayeesCSVRequestV3) HasPaymentChannelCountryCode() bool`

HasPaymentChannelCountryCode returns a boolean if a field has been set.

### GetPaymentChannelCurrency

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelCurrency() string`

GetPaymentChannelCurrency returns the PaymentChannelCurrency field if non-nil, zero value otherwise.

### GetPaymentChannelCurrencyOk

`func (o *CreatePayeesCSVRequestV3) GetPaymentChannelCurrencyOk() (*string, bool)`

GetPaymentChannelCurrencyOk returns a tuple with the PaymentChannelCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelCurrency

`func (o *CreatePayeesCSVRequestV3) SetPaymentChannelCurrency(v string)`

SetPaymentChannelCurrency sets PaymentChannelCurrency field to given value.

### HasPaymentChannelCurrency

`func (o *CreatePayeesCSVRequestV3) HasPaymentChannelCurrency() bool`

HasPaymentChannelCurrency returns a boolean if a field has been set.

### GetChallengeDescription

`func (o *CreatePayeesCSVRequestV3) GetChallengeDescription() string`

GetChallengeDescription returns the ChallengeDescription field if non-nil, zero value otherwise.

### GetChallengeDescriptionOk

`func (o *CreatePayeesCSVRequestV3) GetChallengeDescriptionOk() (*string, bool)`

GetChallengeDescriptionOk returns a tuple with the ChallengeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeDescription

`func (o *CreatePayeesCSVRequestV3) SetChallengeDescription(v string)`

SetChallengeDescription sets ChallengeDescription field to given value.

### HasChallengeDescription

`func (o *CreatePayeesCSVRequestV3) HasChallengeDescription() bool`

HasChallengeDescription returns a boolean if a field has been set.

### GetChallengeValue

`func (o *CreatePayeesCSVRequestV3) GetChallengeValue() string`

GetChallengeValue returns the ChallengeValue field if non-nil, zero value otherwise.

### GetChallengeValueOk

`func (o *CreatePayeesCSVRequestV3) GetChallengeValueOk() (*string, bool)`

GetChallengeValueOk returns a tuple with the ChallengeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeValue

`func (o *CreatePayeesCSVRequestV3) SetChallengeValue(v string)`

SetChallengeValue sets ChallengeValue field to given value.

### HasChallengeValue

`func (o *CreatePayeesCSVRequestV3) HasChallengeValue() bool`

HasChallengeValue returns a boolean if a field has been set.

### GetPayeeLanguage

`func (o *CreatePayeesCSVRequestV3) GetPayeeLanguage() string`

GetPayeeLanguage returns the PayeeLanguage field if non-nil, zero value otherwise.

### GetPayeeLanguageOk

`func (o *CreatePayeesCSVRequestV3) GetPayeeLanguageOk() (*string, bool)`

GetPayeeLanguageOk returns a tuple with the PayeeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeLanguage

`func (o *CreatePayeesCSVRequestV3) SetPayeeLanguage(v string)`

SetPayeeLanguage sets PayeeLanguage field to given value.

### HasPayeeLanguage

`func (o *CreatePayeesCSVRequestV3) HasPayeeLanguage() bool`

HasPayeeLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


