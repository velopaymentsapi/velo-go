# UpdatePayeeDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**PayeeAddress**](PayeeAddress.md) |  | [optional] 
**Individual** | Pointer to [**Individual**](Individual.md) |  | [optional] 
**Company** | Pointer to [**NullableCompany**](Company.md) |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**PayeeType** | Pointer to [**PayeeType2**](PayeeType2.md) |  | [optional] 
**Challenge** | Pointer to [**Challenge**](Challenge.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdatePayeeDetailsRequest

`func NewUpdatePayeeDetailsRequest() *UpdatePayeeDetailsRequest`

NewUpdatePayeeDetailsRequest instantiates a new UpdatePayeeDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePayeeDetailsRequestWithDefaults

`func NewUpdatePayeeDetailsRequestWithDefaults() *UpdatePayeeDetailsRequest`

NewUpdatePayeeDetailsRequestWithDefaults instantiates a new UpdatePayeeDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdatePayeeDetailsRequest) GetAddress() PayeeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdatePayeeDetailsRequest) GetAddressOk() (*PayeeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdatePayeeDetailsRequest) SetAddress(v PayeeAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdatePayeeDetailsRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIndividual

`func (o *UpdatePayeeDetailsRequest) GetIndividual() Individual`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *UpdatePayeeDetailsRequest) GetIndividualOk() (*Individual, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *UpdatePayeeDetailsRequest) SetIndividual(v Individual)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *UpdatePayeeDetailsRequest) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *UpdatePayeeDetailsRequest) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdatePayeeDetailsRequest) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdatePayeeDetailsRequest) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdatePayeeDetailsRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *UpdatePayeeDetailsRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UpdatePayeeDetailsRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetLanguage

`func (o *UpdatePayeeDetailsRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdatePayeeDetailsRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdatePayeeDetailsRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdatePayeeDetailsRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPayeeType

`func (o *UpdatePayeeDetailsRequest) GetPayeeType() PayeeType2`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *UpdatePayeeDetailsRequest) GetPayeeTypeOk() (*PayeeType2, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *UpdatePayeeDetailsRequest) SetPayeeType(v PayeeType2)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *UpdatePayeeDetailsRequest) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetChallenge

`func (o *UpdatePayeeDetailsRequest) GetChallenge() Challenge`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *UpdatePayeeDetailsRequest) GetChallengeOk() (*Challenge, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *UpdatePayeeDetailsRequest) SetChallenge(v Challenge)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *UpdatePayeeDetailsRequest) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetEmail

`func (o *UpdatePayeeDetailsRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdatePayeeDetailsRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdatePayeeDetailsRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdatePayeeDetailsRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdatePayeeDetailsRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdatePayeeDetailsRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


