# UpdatePayeeDetailsRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**PayeeAddressV3**](PayeeAddressV3.md) |  | [optional] 
**Individual** | Pointer to [**IndividualV3**](IndividualV3.md) |  | [optional] 
**Company** | Pointer to [**NullableCompanyV3**](CompanyV3.md) |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**PayeeType** | Pointer to [**PayeeTypeEnum**](PayeeTypeEnum.md) |  | [optional] 
**Challenge** | Pointer to [**ChallengeV3**](ChallengeV3.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**ContactSmsNumber** | Pointer to **string** | The phone number of a device that the payee wishes to receive sms messages on  | [optional] 

## Methods

### NewUpdatePayeeDetailsRequestV3

`func NewUpdatePayeeDetailsRequestV3() *UpdatePayeeDetailsRequestV3`

NewUpdatePayeeDetailsRequestV3 instantiates a new UpdatePayeeDetailsRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePayeeDetailsRequestV3WithDefaults

`func NewUpdatePayeeDetailsRequestV3WithDefaults() *UpdatePayeeDetailsRequestV3`

NewUpdatePayeeDetailsRequestV3WithDefaults instantiates a new UpdatePayeeDetailsRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdatePayeeDetailsRequestV3) GetAddress() PayeeAddressV3`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdatePayeeDetailsRequestV3) GetAddressOk() (*PayeeAddressV3, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdatePayeeDetailsRequestV3) SetAddress(v PayeeAddressV3)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdatePayeeDetailsRequestV3) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIndividual

`func (o *UpdatePayeeDetailsRequestV3) GetIndividual() IndividualV3`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *UpdatePayeeDetailsRequestV3) GetIndividualOk() (*IndividualV3, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *UpdatePayeeDetailsRequestV3) SetIndividual(v IndividualV3)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *UpdatePayeeDetailsRequestV3) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *UpdatePayeeDetailsRequestV3) GetCompany() CompanyV3`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdatePayeeDetailsRequestV3) GetCompanyOk() (*CompanyV3, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdatePayeeDetailsRequestV3) SetCompany(v CompanyV3)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdatePayeeDetailsRequestV3) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *UpdatePayeeDetailsRequestV3) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UpdatePayeeDetailsRequestV3) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetLanguage

`func (o *UpdatePayeeDetailsRequestV3) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdatePayeeDetailsRequestV3) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdatePayeeDetailsRequestV3) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdatePayeeDetailsRequestV3) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPayeeType

`func (o *UpdatePayeeDetailsRequestV3) GetPayeeType() PayeeTypeEnum`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *UpdatePayeeDetailsRequestV3) GetPayeeTypeOk() (*PayeeTypeEnum, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *UpdatePayeeDetailsRequestV3) SetPayeeType(v PayeeTypeEnum)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *UpdatePayeeDetailsRequestV3) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetChallenge

`func (o *UpdatePayeeDetailsRequestV3) GetChallenge() ChallengeV3`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *UpdatePayeeDetailsRequestV3) GetChallengeOk() (*ChallengeV3, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *UpdatePayeeDetailsRequestV3) SetChallenge(v ChallengeV3)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *UpdatePayeeDetailsRequestV3) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetEmail

`func (o *UpdatePayeeDetailsRequestV3) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdatePayeeDetailsRequestV3) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdatePayeeDetailsRequestV3) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdatePayeeDetailsRequestV3) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdatePayeeDetailsRequestV3) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdatePayeeDetailsRequestV3) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetContactSmsNumber

`func (o *UpdatePayeeDetailsRequestV3) GetContactSmsNumber() string`

GetContactSmsNumber returns the ContactSmsNumber field if non-nil, zero value otherwise.

### GetContactSmsNumberOk

`func (o *UpdatePayeeDetailsRequestV3) GetContactSmsNumberOk() (*string, bool)`

GetContactSmsNumberOk returns a tuple with the ContactSmsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSmsNumber

`func (o *UpdatePayeeDetailsRequestV3) SetContactSmsNumber(v string)`

SetContactSmsNumber sets ContactSmsNumber field to given value.

### HasContactSmsNumber

`func (o *UpdatePayeeDetailsRequestV3) HasContactSmsNumber() bool`

HasContactSmsNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


