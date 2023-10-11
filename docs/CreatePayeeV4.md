# CreatePayeeV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RemoteId** | **string** |  | 
**Type** | [**PayeeTypeEnum**](PayeeTypeEnum.md) |  | 
**Address** | [**CreatePayeeAddressV4**](CreatePayeeAddressV4.md) |  | 
**PaymentChannel** | Pointer to [**CreatePaymentChannelV4**](CreatePaymentChannelV4.md) |  | [optional] 
**Challenge** | Pointer to [**ChallengeV4**](ChallengeV4.md) |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Company** | Pointer to [**NullableCompanyV4**](CompanyV4.md) |  | [optional] 
**Individual** | Pointer to [**CreateIndividualV4**](CreateIndividualV4.md) |  | [optional] 

## Methods

### NewCreatePayeeV4

`func NewCreatePayeeV4(email string, remoteId string, type_ PayeeTypeEnum, address CreatePayeeAddressV4, ) *CreatePayeeV4`

NewCreatePayeeV4 instantiates a new CreatePayeeV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeeV4WithDefaults

`func NewCreatePayeeV4WithDefaults() *CreatePayeeV4`

NewCreatePayeeV4WithDefaults instantiates a new CreatePayeeV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreatePayeeV4) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreatePayeeV4) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreatePayeeV4) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRemoteId

`func (o *CreatePayeeV4) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreatePayeeV4) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreatePayeeV4) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetType

`func (o *CreatePayeeV4) GetType() PayeeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePayeeV4) GetTypeOk() (*PayeeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePayeeV4) SetType(v PayeeTypeEnum)`

SetType sets Type field to given value.


### GetAddress

`func (o *CreatePayeeV4) GetAddress() CreatePayeeAddressV4`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreatePayeeV4) GetAddressOk() (*CreatePayeeAddressV4, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreatePayeeV4) SetAddress(v CreatePayeeAddressV4)`

SetAddress sets Address field to given value.


### GetPaymentChannel

`func (o *CreatePayeeV4) GetPaymentChannel() CreatePaymentChannelV4`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *CreatePayeeV4) GetPaymentChannelOk() (*CreatePaymentChannelV4, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *CreatePayeeV4) SetPaymentChannel(v CreatePaymentChannelV4)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *CreatePayeeV4) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetChallenge

`func (o *CreatePayeeV4) GetChallenge() ChallengeV4`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *CreatePayeeV4) GetChallengeOk() (*ChallengeV4, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *CreatePayeeV4) SetChallenge(v ChallengeV4)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *CreatePayeeV4) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetLanguage

`func (o *CreatePayeeV4) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreatePayeeV4) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreatePayeeV4) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CreatePayeeV4) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCompany

`func (o *CreatePayeeV4) GetCompany() CompanyV4`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreatePayeeV4) GetCompanyOk() (*CompanyV4, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreatePayeeV4) SetCompany(v CompanyV4)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreatePayeeV4) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *CreatePayeeV4) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *CreatePayeeV4) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndividual

`func (o *CreatePayeeV4) GetIndividual() CreateIndividualV4`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *CreatePayeeV4) GetIndividualOk() (*CreateIndividualV4, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *CreatePayeeV4) SetIndividual(v CreateIndividualV4)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *CreatePayeeV4) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


