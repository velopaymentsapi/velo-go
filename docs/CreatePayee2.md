# CreatePayee2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV3**](PayeePayorRefV3.md) |  | [optional] [readonly] 
**Email** | **string** |  | 
**RemoteId** | **string** |  | 
**Type** | [**PayeeType**](PayeeType.md) |  | 
**Address** | [**CreatePayeeAddress2**](CreatePayeeAddress_2.md) |  | 
**PaymentChannel** | Pointer to [**CreatePaymentChannel2**](CreatePaymentChannel_2.md) |  | [optional] 
**Challenge** | Pointer to [**Challenge2**](Challenge_2.md) |  | [optional] 
**Language** | Pointer to [**Language2**](Language_2.md) |  | [optional] 
**Company** | Pointer to [**NullableCompany2**](Company_2.md) |  | [optional] 
**Individual** | Pointer to [**CreateIndividual2**](CreateIndividual_2.md) |  | [optional] 

## Methods

### NewCreatePayee2

`func NewCreatePayee2(email string, remoteId string, type_ PayeeType, address CreatePayeeAddress2, ) *CreatePayee2`

NewCreatePayee2 instantiates a new CreatePayee2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayee2WithDefaults

`func NewCreatePayee2WithDefaults() *CreatePayee2`

NewCreatePayee2WithDefaults instantiates a new CreatePayee2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *CreatePayee2) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *CreatePayee2) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *CreatePayee2) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *CreatePayee2) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *CreatePayee2) GetPayorRefs() []PayeePayorRefV3`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *CreatePayee2) GetPayorRefsOk() (*[]PayeePayorRefV3, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *CreatePayee2) SetPayorRefs(v []PayeePayorRefV3)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *CreatePayee2) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *CreatePayee2) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *CreatePayee2) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *CreatePayee2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreatePayee2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreatePayee2) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRemoteId

`func (o *CreatePayee2) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreatePayee2) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreatePayee2) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetType

`func (o *CreatePayee2) GetType() PayeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePayee2) GetTypeOk() (*PayeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePayee2) SetType(v PayeeType)`

SetType sets Type field to given value.


### GetAddress

`func (o *CreatePayee2) GetAddress() CreatePayeeAddress2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreatePayee2) GetAddressOk() (*CreatePayeeAddress2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreatePayee2) SetAddress(v CreatePayeeAddress2)`

SetAddress sets Address field to given value.


### GetPaymentChannel

`func (o *CreatePayee2) GetPaymentChannel() CreatePaymentChannel2`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *CreatePayee2) GetPaymentChannelOk() (*CreatePaymentChannel2, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *CreatePayee2) SetPaymentChannel(v CreatePaymentChannel2)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *CreatePayee2) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetChallenge

`func (o *CreatePayee2) GetChallenge() Challenge2`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *CreatePayee2) GetChallengeOk() (*Challenge2, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *CreatePayee2) SetChallenge(v Challenge2)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *CreatePayee2) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetLanguage

`func (o *CreatePayee2) GetLanguage() Language2`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreatePayee2) GetLanguageOk() (*Language2, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreatePayee2) SetLanguage(v Language2)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CreatePayee2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCompany

`func (o *CreatePayee2) GetCompany() Company2`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreatePayee2) GetCompanyOk() (*Company2, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreatePayee2) SetCompany(v Company2)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreatePayee2) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *CreatePayee2) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *CreatePayee2) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndividual

`func (o *CreatePayee2) GetIndividual() CreateIndividual2`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *CreatePayee2) GetIndividualOk() (*CreateIndividual2, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *CreatePayee2) SetIndividual(v CreateIndividual2)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *CreatePayee2) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


