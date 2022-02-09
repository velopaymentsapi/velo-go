# CreatePayeesCSVRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PayeeType2**](PayeeType2.md) |  | 
**RemoteId** | **string** |  | 
**Email** | **string** |  | 
**AddressLine1** | **string** |  | 
**AddressLine2** | Pointer to **string** |  | [optional] 
**AddressLine3** | Pointer to **string** |  | [optional] 
**AddressLine4** | Pointer to **string** |  | [optional] 
**AddressCity** | **string** |  | 
**AddressCountyOrProvince** | Pointer to **string** |  | [optional] 
**AddressZipOrPostcode** | **string** |  | 
**AddressCountry** | **string** | Must be a 2 character country code - per ISO 3166-1 | 
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
**PaymentChannelCountryCode** | Pointer to **string** | Must be a 2 character country code - per ISO 3166-1 | [optional] 
**PaymentChannelCurrency** | Pointer to **string** |  | [optional] 
**ChallengeDescription** | Pointer to **string** |  | [optional] 
**ChallengeValue** | Pointer to **string** |  | [optional] 
**PayeeLanguage** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 

## Methods

### NewCreatePayeesCSVRequest

`func NewCreatePayeesCSVRequest(type_ PayeeType2, remoteId string, email string, addressLine1 string, addressCity string, addressZipOrPostcode string, addressCountry string, ) *CreatePayeesCSVRequest`

NewCreatePayeesCSVRequest instantiates a new CreatePayeesCSVRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesCSVRequestWithDefaults

`func NewCreatePayeesCSVRequestWithDefaults() *CreatePayeesCSVRequest`

NewCreatePayeesCSVRequestWithDefaults instantiates a new CreatePayeesCSVRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreatePayeesCSVRequest) GetType() PayeeType2`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePayeesCSVRequest) GetTypeOk() (*PayeeType2, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePayeesCSVRequest) SetType(v PayeeType2)`

SetType sets Type field to given value.


### GetRemoteId

`func (o *CreatePayeesCSVRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreatePayeesCSVRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreatePayeesCSVRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetEmail

`func (o *CreatePayeesCSVRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreatePayeesCSVRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreatePayeesCSVRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAddressLine1

`func (o *CreatePayeesCSVRequest) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CreatePayeesCSVRequest) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CreatePayeesCSVRequest) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *CreatePayeesCSVRequest) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CreatePayeesCSVRequest) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CreatePayeesCSVRequest) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CreatePayeesCSVRequest) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *CreatePayeesCSVRequest) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *CreatePayeesCSVRequest) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *CreatePayeesCSVRequest) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *CreatePayeesCSVRequest) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetAddressLine4

`func (o *CreatePayeesCSVRequest) GetAddressLine4() string`

GetAddressLine4 returns the AddressLine4 field if non-nil, zero value otherwise.

### GetAddressLine4Ok

`func (o *CreatePayeesCSVRequest) GetAddressLine4Ok() (*string, bool)`

GetAddressLine4Ok returns a tuple with the AddressLine4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine4

`func (o *CreatePayeesCSVRequest) SetAddressLine4(v string)`

SetAddressLine4 sets AddressLine4 field to given value.

### HasAddressLine4

`func (o *CreatePayeesCSVRequest) HasAddressLine4() bool`

HasAddressLine4 returns a boolean if a field has been set.

### GetAddressCity

`func (o *CreatePayeesCSVRequest) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *CreatePayeesCSVRequest) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *CreatePayeesCSVRequest) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.


### GetAddressCountyOrProvince

`func (o *CreatePayeesCSVRequest) GetAddressCountyOrProvince() string`

GetAddressCountyOrProvince returns the AddressCountyOrProvince field if non-nil, zero value otherwise.

### GetAddressCountyOrProvinceOk

`func (o *CreatePayeesCSVRequest) GetAddressCountyOrProvinceOk() (*string, bool)`

GetAddressCountyOrProvinceOk returns a tuple with the AddressCountyOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountyOrProvince

`func (o *CreatePayeesCSVRequest) SetAddressCountyOrProvince(v string)`

SetAddressCountyOrProvince sets AddressCountyOrProvince field to given value.

### HasAddressCountyOrProvince

`func (o *CreatePayeesCSVRequest) HasAddressCountyOrProvince() bool`

HasAddressCountyOrProvince returns a boolean if a field has been set.

### GetAddressZipOrPostcode

`func (o *CreatePayeesCSVRequest) GetAddressZipOrPostcode() string`

GetAddressZipOrPostcode returns the AddressZipOrPostcode field if non-nil, zero value otherwise.

### GetAddressZipOrPostcodeOk

`func (o *CreatePayeesCSVRequest) GetAddressZipOrPostcodeOk() (*string, bool)`

GetAddressZipOrPostcodeOk returns a tuple with the AddressZipOrPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZipOrPostcode

`func (o *CreatePayeesCSVRequest) SetAddressZipOrPostcode(v string)`

SetAddressZipOrPostcode sets AddressZipOrPostcode field to given value.


### GetAddressCountry

`func (o *CreatePayeesCSVRequest) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *CreatePayeesCSVRequest) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *CreatePayeesCSVRequest) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.


### GetIndividualNationalIdentification

`func (o *CreatePayeesCSVRequest) GetIndividualNationalIdentification() string`

GetIndividualNationalIdentification returns the IndividualNationalIdentification field if non-nil, zero value otherwise.

### GetIndividualNationalIdentificationOk

`func (o *CreatePayeesCSVRequest) GetIndividualNationalIdentificationOk() (*string, bool)`

GetIndividualNationalIdentificationOk returns a tuple with the IndividualNationalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualNationalIdentification

`func (o *CreatePayeesCSVRequest) SetIndividualNationalIdentification(v string)`

SetIndividualNationalIdentification sets IndividualNationalIdentification field to given value.

### HasIndividualNationalIdentification

`func (o *CreatePayeesCSVRequest) HasIndividualNationalIdentification() bool`

HasIndividualNationalIdentification returns a boolean if a field has been set.

### GetIndividualDateOfBirth

`func (o *CreatePayeesCSVRequest) GetIndividualDateOfBirth() string`

GetIndividualDateOfBirth returns the IndividualDateOfBirth field if non-nil, zero value otherwise.

### GetIndividualDateOfBirthOk

`func (o *CreatePayeesCSVRequest) GetIndividualDateOfBirthOk() (*string, bool)`

GetIndividualDateOfBirthOk returns a tuple with the IndividualDateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualDateOfBirth

`func (o *CreatePayeesCSVRequest) SetIndividualDateOfBirth(v string)`

SetIndividualDateOfBirth sets IndividualDateOfBirth field to given value.

### HasIndividualDateOfBirth

`func (o *CreatePayeesCSVRequest) HasIndividualDateOfBirth() bool`

HasIndividualDateOfBirth returns a boolean if a field has been set.

### GetIndividualTitle

`func (o *CreatePayeesCSVRequest) GetIndividualTitle() string`

GetIndividualTitle returns the IndividualTitle field if non-nil, zero value otherwise.

### GetIndividualTitleOk

`func (o *CreatePayeesCSVRequest) GetIndividualTitleOk() (*string, bool)`

GetIndividualTitleOk returns a tuple with the IndividualTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualTitle

`func (o *CreatePayeesCSVRequest) SetIndividualTitle(v string)`

SetIndividualTitle sets IndividualTitle field to given value.

### HasIndividualTitle

`func (o *CreatePayeesCSVRequest) HasIndividualTitle() bool`

HasIndividualTitle returns a boolean if a field has been set.

### GetIndividualFirstName

`func (o *CreatePayeesCSVRequest) GetIndividualFirstName() string`

GetIndividualFirstName returns the IndividualFirstName field if non-nil, zero value otherwise.

### GetIndividualFirstNameOk

`func (o *CreatePayeesCSVRequest) GetIndividualFirstNameOk() (*string, bool)`

GetIndividualFirstNameOk returns a tuple with the IndividualFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualFirstName

`func (o *CreatePayeesCSVRequest) SetIndividualFirstName(v string)`

SetIndividualFirstName sets IndividualFirstName field to given value.

### HasIndividualFirstName

`func (o *CreatePayeesCSVRequest) HasIndividualFirstName() bool`

HasIndividualFirstName returns a boolean if a field has been set.

### GetIndividualOtherNames

`func (o *CreatePayeesCSVRequest) GetIndividualOtherNames() string`

GetIndividualOtherNames returns the IndividualOtherNames field if non-nil, zero value otherwise.

### GetIndividualOtherNamesOk

`func (o *CreatePayeesCSVRequest) GetIndividualOtherNamesOk() (*string, bool)`

GetIndividualOtherNamesOk returns a tuple with the IndividualOtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualOtherNames

`func (o *CreatePayeesCSVRequest) SetIndividualOtherNames(v string)`

SetIndividualOtherNames sets IndividualOtherNames field to given value.

### HasIndividualOtherNames

`func (o *CreatePayeesCSVRequest) HasIndividualOtherNames() bool`

HasIndividualOtherNames returns a boolean if a field has been set.

### GetIndividualLastName

`func (o *CreatePayeesCSVRequest) GetIndividualLastName() string`

GetIndividualLastName returns the IndividualLastName field if non-nil, zero value otherwise.

### GetIndividualLastNameOk

`func (o *CreatePayeesCSVRequest) GetIndividualLastNameOk() (*string, bool)`

GetIndividualLastNameOk returns a tuple with the IndividualLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualLastName

`func (o *CreatePayeesCSVRequest) SetIndividualLastName(v string)`

SetIndividualLastName sets IndividualLastName field to given value.

### HasIndividualLastName

`func (o *CreatePayeesCSVRequest) HasIndividualLastName() bool`

HasIndividualLastName returns a boolean if a field has been set.

### GetCompanyName

`func (o *CreatePayeesCSVRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CreatePayeesCSVRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CreatePayeesCSVRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CreatePayeesCSVRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyEIN

`func (o *CreatePayeesCSVRequest) GetCompanyEIN() string`

GetCompanyEIN returns the CompanyEIN field if non-nil, zero value otherwise.

### GetCompanyEINOk

`func (o *CreatePayeesCSVRequest) GetCompanyEINOk() (*string, bool)`

GetCompanyEINOk returns a tuple with the CompanyEIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEIN

`func (o *CreatePayeesCSVRequest) SetCompanyEIN(v string)`

SetCompanyEIN sets CompanyEIN field to given value.

### HasCompanyEIN

`func (o *CreatePayeesCSVRequest) HasCompanyEIN() bool`

HasCompanyEIN returns a boolean if a field has been set.

### GetCompanyOperatingName

`func (o *CreatePayeesCSVRequest) GetCompanyOperatingName() string`

GetCompanyOperatingName returns the CompanyOperatingName field if non-nil, zero value otherwise.

### GetCompanyOperatingNameOk

`func (o *CreatePayeesCSVRequest) GetCompanyOperatingNameOk() (*string, bool)`

GetCompanyOperatingNameOk returns a tuple with the CompanyOperatingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyOperatingName

`func (o *CreatePayeesCSVRequest) SetCompanyOperatingName(v string)`

SetCompanyOperatingName sets CompanyOperatingName field to given value.

### HasCompanyOperatingName

`func (o *CreatePayeesCSVRequest) HasCompanyOperatingName() bool`

HasCompanyOperatingName returns a boolean if a field has been set.

### GetPaymentChannelAccountNumber

`func (o *CreatePayeesCSVRequest) GetPaymentChannelAccountNumber() string`

GetPaymentChannelAccountNumber returns the PaymentChannelAccountNumber field if non-nil, zero value otherwise.

### GetPaymentChannelAccountNumberOk

`func (o *CreatePayeesCSVRequest) GetPaymentChannelAccountNumberOk() (*string, bool)`

GetPaymentChannelAccountNumberOk returns a tuple with the PaymentChannelAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelAccountNumber

`func (o *CreatePayeesCSVRequest) SetPaymentChannelAccountNumber(v string)`

SetPaymentChannelAccountNumber sets PaymentChannelAccountNumber field to given value.

### HasPaymentChannelAccountNumber

`func (o *CreatePayeesCSVRequest) HasPaymentChannelAccountNumber() bool`

HasPaymentChannelAccountNumber returns a boolean if a field has been set.

### GetPaymentChannelRoutingNumber

`func (o *CreatePayeesCSVRequest) GetPaymentChannelRoutingNumber() string`

GetPaymentChannelRoutingNumber returns the PaymentChannelRoutingNumber field if non-nil, zero value otherwise.

### GetPaymentChannelRoutingNumberOk

`func (o *CreatePayeesCSVRequest) GetPaymentChannelRoutingNumberOk() (*string, bool)`

GetPaymentChannelRoutingNumberOk returns a tuple with the PaymentChannelRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelRoutingNumber

`func (o *CreatePayeesCSVRequest) SetPaymentChannelRoutingNumber(v string)`

SetPaymentChannelRoutingNumber sets PaymentChannelRoutingNumber field to given value.

### HasPaymentChannelRoutingNumber

`func (o *CreatePayeesCSVRequest) HasPaymentChannelRoutingNumber() bool`

HasPaymentChannelRoutingNumber returns a boolean if a field has been set.

### GetPaymentChannelAccountName

`func (o *CreatePayeesCSVRequest) GetPaymentChannelAccountName() string`

GetPaymentChannelAccountName returns the PaymentChannelAccountName field if non-nil, zero value otherwise.

### GetPaymentChannelAccountNameOk

`func (o *CreatePayeesCSVRequest) GetPaymentChannelAccountNameOk() (*string, bool)`

GetPaymentChannelAccountNameOk returns a tuple with the PaymentChannelAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelAccountName

`func (o *CreatePayeesCSVRequest) SetPaymentChannelAccountName(v string)`

SetPaymentChannelAccountName sets PaymentChannelAccountName field to given value.

### HasPaymentChannelAccountName

`func (o *CreatePayeesCSVRequest) HasPaymentChannelAccountName() bool`

HasPaymentChannelAccountName returns a boolean if a field has been set.

### GetPaymentChannelIban

`func (o *CreatePayeesCSVRequest) GetPaymentChannelIban() string`

GetPaymentChannelIban returns the PaymentChannelIban field if non-nil, zero value otherwise.

### GetPaymentChannelIbanOk

`func (o *CreatePayeesCSVRequest) GetPaymentChannelIbanOk() (*string, bool)`

GetPaymentChannelIbanOk returns a tuple with the PaymentChannelIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelIban

`func (o *CreatePayeesCSVRequest) SetPaymentChannelIban(v string)`

SetPaymentChannelIban sets PaymentChannelIban field to given value.

### HasPaymentChannelIban

`func (o *CreatePayeesCSVRequest) HasPaymentChannelIban() bool`

HasPaymentChannelIban returns a boolean if a field has been set.

### GetPaymentChannelCountryCode

`func (o *CreatePayeesCSVRequest) GetPaymentChannelCountryCode() string`

GetPaymentChannelCountryCode returns the PaymentChannelCountryCode field if non-nil, zero value otherwise.

### GetPaymentChannelCountryCodeOk

`func (o *CreatePayeesCSVRequest) GetPaymentChannelCountryCodeOk() (*string, bool)`

GetPaymentChannelCountryCodeOk returns a tuple with the PaymentChannelCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelCountryCode

`func (o *CreatePayeesCSVRequest) SetPaymentChannelCountryCode(v string)`

SetPaymentChannelCountryCode sets PaymentChannelCountryCode field to given value.

### HasPaymentChannelCountryCode

`func (o *CreatePayeesCSVRequest) HasPaymentChannelCountryCode() bool`

HasPaymentChannelCountryCode returns a boolean if a field has been set.

### GetPaymentChannelCurrency

`func (o *CreatePayeesCSVRequest) GetPaymentChannelCurrency() string`

GetPaymentChannelCurrency returns the PaymentChannelCurrency field if non-nil, zero value otherwise.

### GetPaymentChannelCurrencyOk

`func (o *CreatePayeesCSVRequest) GetPaymentChannelCurrencyOk() (*string, bool)`

GetPaymentChannelCurrencyOk returns a tuple with the PaymentChannelCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelCurrency

`func (o *CreatePayeesCSVRequest) SetPaymentChannelCurrency(v string)`

SetPaymentChannelCurrency sets PaymentChannelCurrency field to given value.

### HasPaymentChannelCurrency

`func (o *CreatePayeesCSVRequest) HasPaymentChannelCurrency() bool`

HasPaymentChannelCurrency returns a boolean if a field has been set.

### GetChallengeDescription

`func (o *CreatePayeesCSVRequest) GetChallengeDescription() string`

GetChallengeDescription returns the ChallengeDescription field if non-nil, zero value otherwise.

### GetChallengeDescriptionOk

`func (o *CreatePayeesCSVRequest) GetChallengeDescriptionOk() (*string, bool)`

GetChallengeDescriptionOk returns a tuple with the ChallengeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeDescription

`func (o *CreatePayeesCSVRequest) SetChallengeDescription(v string)`

SetChallengeDescription sets ChallengeDescription field to given value.

### HasChallengeDescription

`func (o *CreatePayeesCSVRequest) HasChallengeDescription() bool`

HasChallengeDescription returns a boolean if a field has been set.

### GetChallengeValue

`func (o *CreatePayeesCSVRequest) GetChallengeValue() string`

GetChallengeValue returns the ChallengeValue field if non-nil, zero value otherwise.

### GetChallengeValueOk

`func (o *CreatePayeesCSVRequest) GetChallengeValueOk() (*string, bool)`

GetChallengeValueOk returns a tuple with the ChallengeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeValue

`func (o *CreatePayeesCSVRequest) SetChallengeValue(v string)`

SetChallengeValue sets ChallengeValue field to given value.

### HasChallengeValue

`func (o *CreatePayeesCSVRequest) HasChallengeValue() bool`

HasChallengeValue returns a boolean if a field has been set.

### GetPayeeLanguage

`func (o *CreatePayeesCSVRequest) GetPayeeLanguage() string`

GetPayeeLanguage returns the PayeeLanguage field if non-nil, zero value otherwise.

### GetPayeeLanguageOk

`func (o *CreatePayeesCSVRequest) GetPayeeLanguageOk() (*string, bool)`

GetPayeeLanguageOk returns a tuple with the PayeeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeLanguage

`func (o *CreatePayeesCSVRequest) SetPayeeLanguage(v string)`

SetPayeeLanguage sets PayeeLanguage field to given value.

### HasPayeeLanguage

`func (o *CreatePayeesCSVRequest) HasPayeeLanguage() bool`

HasPayeeLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


