# CreatePayeeV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV3**](PayeePayorRefV3.md) |  | [optional] [readonly] 
**Email** | **string** |  | 
**RemoteId** | **string** |  | 
**Type** | **string** | Type of Payee. One of the following values: Individual, Company | 
**Address** | [**CreatePayeeAddressV3**](CreatePayeeAddressV3.md) |  | 
**PaymentChannel** | Pointer to [**CreatePaymentChannelV3**](CreatePaymentChannelV3.md) |  | [optional] 
**Challenge** | Pointer to [**ChallengeV3**](ChallengeV3.md) |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Company** | Pointer to [**NullableCompanyV3**](CompanyV3.md) |  | [optional] 
**Individual** | Pointer to [**CreateIndividualV3**](CreateIndividualV3.md) |  | [optional] 

## Methods

### NewCreatePayeeV3

`func NewCreatePayeeV3(email string, remoteId string, type_ string, address CreatePayeeAddressV3, ) *CreatePayeeV3`

NewCreatePayeeV3 instantiates a new CreatePayeeV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeeV3WithDefaults

`func NewCreatePayeeV3WithDefaults() *CreatePayeeV3`

NewCreatePayeeV3WithDefaults instantiates a new CreatePayeeV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *CreatePayeeV3) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *CreatePayeeV3) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *CreatePayeeV3) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *CreatePayeeV3) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *CreatePayeeV3) GetPayorRefs() []PayeePayorRefV3`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *CreatePayeeV3) GetPayorRefsOk() (*[]PayeePayorRefV3, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *CreatePayeeV3) SetPayorRefs(v []PayeePayorRefV3)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *CreatePayeeV3) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *CreatePayeeV3) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *CreatePayeeV3) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *CreatePayeeV3) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreatePayeeV3) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreatePayeeV3) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRemoteId

`func (o *CreatePayeeV3) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreatePayeeV3) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreatePayeeV3) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetType

`func (o *CreatePayeeV3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePayeeV3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePayeeV3) SetType(v string)`

SetType sets Type field to given value.


### GetAddress

`func (o *CreatePayeeV3) GetAddress() CreatePayeeAddressV3`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreatePayeeV3) GetAddressOk() (*CreatePayeeAddressV3, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreatePayeeV3) SetAddress(v CreatePayeeAddressV3)`

SetAddress sets Address field to given value.


### GetPaymentChannel

`func (o *CreatePayeeV3) GetPaymentChannel() CreatePaymentChannelV3`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *CreatePayeeV3) GetPaymentChannelOk() (*CreatePaymentChannelV3, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *CreatePayeeV3) SetPaymentChannel(v CreatePaymentChannelV3)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *CreatePayeeV3) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetChallenge

`func (o *CreatePayeeV3) GetChallenge() ChallengeV3`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *CreatePayeeV3) GetChallengeOk() (*ChallengeV3, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *CreatePayeeV3) SetChallenge(v ChallengeV3)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *CreatePayeeV3) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetLanguage

`func (o *CreatePayeeV3) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreatePayeeV3) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreatePayeeV3) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CreatePayeeV3) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCompany

`func (o *CreatePayeeV3) GetCompany() CompanyV3`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreatePayeeV3) GetCompanyOk() (*CompanyV3, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreatePayeeV3) SetCompany(v CompanyV3)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreatePayeeV3) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *CreatePayeeV3) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *CreatePayeeV3) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndividual

`func (o *CreatePayeeV3) GetIndividual() CreateIndividualV3`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *CreatePayeeV3) GetIndividualOk() (*CreateIndividualV3, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *CreatePayeeV3) SetIndividual(v CreateIndividualV3)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *CreatePayeeV3) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


