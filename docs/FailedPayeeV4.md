# FailedPayeeV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV4**](PayeePayorRefV4.md) |  | [optional] [readonly] 
**Email** | Pointer to **string** |  | [optional] 
**RemoteId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type of Payee. One of the following values: Individual, Company | [optional] 
**Address** | Pointer to [**CreatePayeeAddressV4**](CreatePayeeAddressV4.md) |  | [optional] 
**PaymentChannel** | Pointer to [**CreatePaymentChannelV4**](CreatePaymentChannelV4.md) |  | [optional] 
**Challenge** | Pointer to [**ChallengeV4**](ChallengeV4.md) |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Company** | Pointer to [**NullableCompanyV4**](CompanyV4.md) |  | [optional] 
**Individual** | Pointer to [**CreateIndividualV4**](CreateIndividualV4.md) |  | [optional] 

## Methods

### NewFailedPayeeV4

`func NewFailedPayeeV4() *FailedPayeeV4`

NewFailedPayeeV4 instantiates a new FailedPayeeV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedPayeeV4WithDefaults

`func NewFailedPayeeV4WithDefaults() *FailedPayeeV4`

NewFailedPayeeV4WithDefaults instantiates a new FailedPayeeV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *FailedPayeeV4) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *FailedPayeeV4) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *FailedPayeeV4) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *FailedPayeeV4) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *FailedPayeeV4) GetPayorRefs() []PayeePayorRefV4`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *FailedPayeeV4) GetPayorRefsOk() (*[]PayeePayorRefV4, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *FailedPayeeV4) SetPayorRefs(v []PayeePayorRefV4)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *FailedPayeeV4) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *FailedPayeeV4) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *FailedPayeeV4) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *FailedPayeeV4) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FailedPayeeV4) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FailedPayeeV4) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FailedPayeeV4) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRemoteId

`func (o *FailedPayeeV4) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *FailedPayeeV4) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *FailedPayeeV4) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *FailedPayeeV4) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetType

`func (o *FailedPayeeV4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FailedPayeeV4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FailedPayeeV4) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FailedPayeeV4) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *FailedPayeeV4) GetAddress() CreatePayeeAddressV4`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FailedPayeeV4) GetAddressOk() (*CreatePayeeAddressV4, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FailedPayeeV4) SetAddress(v CreatePayeeAddressV4)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FailedPayeeV4) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPaymentChannel

`func (o *FailedPayeeV4) GetPaymentChannel() CreatePaymentChannelV4`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *FailedPayeeV4) GetPaymentChannelOk() (*CreatePaymentChannelV4, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *FailedPayeeV4) SetPaymentChannel(v CreatePaymentChannelV4)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *FailedPayeeV4) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetChallenge

`func (o *FailedPayeeV4) GetChallenge() ChallengeV4`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *FailedPayeeV4) GetChallengeOk() (*ChallengeV4, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *FailedPayeeV4) SetChallenge(v ChallengeV4)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *FailedPayeeV4) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetLanguage

`func (o *FailedPayeeV4) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FailedPayeeV4) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FailedPayeeV4) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *FailedPayeeV4) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCompany

`func (o *FailedPayeeV4) GetCompany() CompanyV4`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FailedPayeeV4) GetCompanyOk() (*CompanyV4, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FailedPayeeV4) SetCompany(v CompanyV4)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *FailedPayeeV4) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *FailedPayeeV4) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *FailedPayeeV4) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndividual

`func (o *FailedPayeeV4) GetIndividual() CreateIndividualV4`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *FailedPayeeV4) GetIndividualOk() (*CreateIndividualV4, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *FailedPayeeV4) SetIndividual(v CreateIndividualV4)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *FailedPayeeV4) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


