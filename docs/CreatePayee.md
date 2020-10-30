# CreatePayee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV2**](PayeePayorRefV2.md) |  | [optional] [readonly] 
**Email** | **string** |  | 
**RemoteId** | **string** |  | 
**Type** | [**PayeeType**](PayeeType.md) |  | 
**Address** | [**CreatePayeeAddress**](CreatePayeeAddress.md) |  | 
**PaymentChannel** | Pointer to [**CreatePaymentChannel**](CreatePaymentChannel.md) |  | [optional] 
**Challenge** | Pointer to [**Challenge**](Challenge.md) |  | [optional] 
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 
**Company** | Pointer to [**NullableCompanyV1**](CompanyV1.md) |  | [optional] 
**Individual** | Pointer to [**CreateIndividual**](CreateIndividual.md) |  | [optional] 

## Methods

### NewCreatePayee

`func NewCreatePayee(email string, remoteId string, type_ PayeeType, address CreatePayeeAddress, ) *CreatePayee`

NewCreatePayee instantiates a new CreatePayee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeeWithDefaults

`func NewCreatePayeeWithDefaults() *CreatePayee`

NewCreatePayeeWithDefaults instantiates a new CreatePayee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *CreatePayee) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *CreatePayee) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *CreatePayee) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *CreatePayee) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *CreatePayee) GetPayorRefs() []PayeePayorRefV2`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *CreatePayee) GetPayorRefsOk() (*[]PayeePayorRefV2, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *CreatePayee) SetPayorRefs(v []PayeePayorRefV2)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *CreatePayee) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *CreatePayee) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *CreatePayee) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *CreatePayee) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreatePayee) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreatePayee) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRemoteId

`func (o *CreatePayee) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreatePayee) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreatePayee) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetType

`func (o *CreatePayee) GetType() PayeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePayee) GetTypeOk() (*PayeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePayee) SetType(v PayeeType)`

SetType sets Type field to given value.


### GetAddress

`func (o *CreatePayee) GetAddress() CreatePayeeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreatePayee) GetAddressOk() (*CreatePayeeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreatePayee) SetAddress(v CreatePayeeAddress)`

SetAddress sets Address field to given value.


### GetPaymentChannel

`func (o *CreatePayee) GetPaymentChannel() CreatePaymentChannel`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *CreatePayee) GetPaymentChannelOk() (*CreatePaymentChannel, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *CreatePayee) SetPaymentChannel(v CreatePaymentChannel)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *CreatePayee) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetChallenge

`func (o *CreatePayee) GetChallenge() Challenge`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *CreatePayee) GetChallengeOk() (*Challenge, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *CreatePayee) SetChallenge(v Challenge)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *CreatePayee) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetLanguage

`func (o *CreatePayee) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreatePayee) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreatePayee) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CreatePayee) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCompany

`func (o *CreatePayee) GetCompany() CompanyV1`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreatePayee) GetCompanyOk() (*CompanyV1, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreatePayee) SetCompany(v CompanyV1)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreatePayee) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *CreatePayee) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *CreatePayee) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndividual

`func (o *CreatePayee) GetIndividual() CreateIndividual`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *CreatePayee) GetIndividualOk() (*CreateIndividual, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *CreatePayee) SetIndividual(v CreateIndividual)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *CreatePayee) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


