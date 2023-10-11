# UpdatePayeeDetailsRequestV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**PayeeAddressV4**](PayeeAddressV4.md) |  | [optional] 
**Individual** | Pointer to [**IndividualV4**](IndividualV4.md) |  | [optional] 
**Company** | Pointer to [**NullableCompanyV4**](CompanyV4.md) |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**PayeeType** | Pointer to [**PayeeTypeEnum**](PayeeTypeEnum.md) |  | [optional] 
**Challenge** | Pointer to [**ChallengeV4**](ChallengeV4.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**ContactSmsNumber** | Pointer to **string** | The phone number of a device that the payee wishes to receive sms messages on  | [optional] 

## Methods

### NewUpdatePayeeDetailsRequestV4

`func NewUpdatePayeeDetailsRequestV4() *UpdatePayeeDetailsRequestV4`

NewUpdatePayeeDetailsRequestV4 instantiates a new UpdatePayeeDetailsRequestV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePayeeDetailsRequestV4WithDefaults

`func NewUpdatePayeeDetailsRequestV4WithDefaults() *UpdatePayeeDetailsRequestV4`

NewUpdatePayeeDetailsRequestV4WithDefaults instantiates a new UpdatePayeeDetailsRequestV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdatePayeeDetailsRequestV4) GetAddress() PayeeAddressV4`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdatePayeeDetailsRequestV4) GetAddressOk() (*PayeeAddressV4, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdatePayeeDetailsRequestV4) SetAddress(v PayeeAddressV4)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdatePayeeDetailsRequestV4) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIndividual

`func (o *UpdatePayeeDetailsRequestV4) GetIndividual() IndividualV4`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *UpdatePayeeDetailsRequestV4) GetIndividualOk() (*IndividualV4, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *UpdatePayeeDetailsRequestV4) SetIndividual(v IndividualV4)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *UpdatePayeeDetailsRequestV4) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *UpdatePayeeDetailsRequestV4) GetCompany() CompanyV4`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdatePayeeDetailsRequestV4) GetCompanyOk() (*CompanyV4, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdatePayeeDetailsRequestV4) SetCompany(v CompanyV4)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdatePayeeDetailsRequestV4) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *UpdatePayeeDetailsRequestV4) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UpdatePayeeDetailsRequestV4) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetLanguage

`func (o *UpdatePayeeDetailsRequestV4) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdatePayeeDetailsRequestV4) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdatePayeeDetailsRequestV4) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdatePayeeDetailsRequestV4) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPayeeType

`func (o *UpdatePayeeDetailsRequestV4) GetPayeeType() PayeeTypeEnum`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *UpdatePayeeDetailsRequestV4) GetPayeeTypeOk() (*PayeeTypeEnum, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *UpdatePayeeDetailsRequestV4) SetPayeeType(v PayeeTypeEnum)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *UpdatePayeeDetailsRequestV4) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetChallenge

`func (o *UpdatePayeeDetailsRequestV4) GetChallenge() ChallengeV4`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *UpdatePayeeDetailsRequestV4) GetChallengeOk() (*ChallengeV4, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *UpdatePayeeDetailsRequestV4) SetChallenge(v ChallengeV4)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *UpdatePayeeDetailsRequestV4) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetEmail

`func (o *UpdatePayeeDetailsRequestV4) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdatePayeeDetailsRequestV4) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdatePayeeDetailsRequestV4) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdatePayeeDetailsRequestV4) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdatePayeeDetailsRequestV4) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdatePayeeDetailsRequestV4) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetContactSmsNumber

`func (o *UpdatePayeeDetailsRequestV4) GetContactSmsNumber() string`

GetContactSmsNumber returns the ContactSmsNumber field if non-nil, zero value otherwise.

### GetContactSmsNumberOk

`func (o *UpdatePayeeDetailsRequestV4) GetContactSmsNumberOk() (*string, bool)`

GetContactSmsNumberOk returns a tuple with the ContactSmsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSmsNumber

`func (o *UpdatePayeeDetailsRequestV4) SetContactSmsNumber(v string)`

SetContactSmsNumber sets ContactSmsNumber field to given value.

### HasContactSmsNumber

`func (o *UpdatePayeeDetailsRequestV4) HasContactSmsNumber() bool`

HasContactSmsNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


