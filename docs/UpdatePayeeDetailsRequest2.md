# UpdatePayeeDetailsRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**PayeeAddress2**](PayeeAddress2.md) |  | [optional] 
**Individual** | Pointer to [**Individual2**](Individual2.md) |  | [optional] 
**Company** | Pointer to [**NullableCompany2**](Company2.md) |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**PayeeType** | Pointer to [**PayeeType**](PayeeType.md) |  | [optional] 
**Challenge** | Pointer to [**Challenge2**](Challenge2.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdatePayeeDetailsRequest2

`func NewUpdatePayeeDetailsRequest2() *UpdatePayeeDetailsRequest2`

NewUpdatePayeeDetailsRequest2 instantiates a new UpdatePayeeDetailsRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePayeeDetailsRequest2WithDefaults

`func NewUpdatePayeeDetailsRequest2WithDefaults() *UpdatePayeeDetailsRequest2`

NewUpdatePayeeDetailsRequest2WithDefaults instantiates a new UpdatePayeeDetailsRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdatePayeeDetailsRequest2) GetAddress() PayeeAddress2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdatePayeeDetailsRequest2) GetAddressOk() (*PayeeAddress2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdatePayeeDetailsRequest2) SetAddress(v PayeeAddress2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdatePayeeDetailsRequest2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIndividual

`func (o *UpdatePayeeDetailsRequest2) GetIndividual() Individual2`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *UpdatePayeeDetailsRequest2) GetIndividualOk() (*Individual2, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *UpdatePayeeDetailsRequest2) SetIndividual(v Individual2)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *UpdatePayeeDetailsRequest2) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *UpdatePayeeDetailsRequest2) GetCompany() Company2`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdatePayeeDetailsRequest2) GetCompanyOk() (*Company2, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdatePayeeDetailsRequest2) SetCompany(v Company2)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdatePayeeDetailsRequest2) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *UpdatePayeeDetailsRequest2) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UpdatePayeeDetailsRequest2) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetLanguage

`func (o *UpdatePayeeDetailsRequest2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdatePayeeDetailsRequest2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdatePayeeDetailsRequest2) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdatePayeeDetailsRequest2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPayeeType

`func (o *UpdatePayeeDetailsRequest2) GetPayeeType() PayeeType`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *UpdatePayeeDetailsRequest2) GetPayeeTypeOk() (*PayeeType, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *UpdatePayeeDetailsRequest2) SetPayeeType(v PayeeType)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *UpdatePayeeDetailsRequest2) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetChallenge

`func (o *UpdatePayeeDetailsRequest2) GetChallenge() Challenge2`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *UpdatePayeeDetailsRequest2) GetChallengeOk() (*Challenge2, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *UpdatePayeeDetailsRequest2) SetChallenge(v Challenge2)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *UpdatePayeeDetailsRequest2) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetEmail

`func (o *UpdatePayeeDetailsRequest2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdatePayeeDetailsRequest2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdatePayeeDetailsRequest2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdatePayeeDetailsRequest2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdatePayeeDetailsRequest2) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdatePayeeDetailsRequest2) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


