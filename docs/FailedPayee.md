# FailedPayee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV3**](PayeePayorRefV3.md) |  | [optional] [readonly] 
**Email** | Pointer to **string** |  | [optional] 
**RemoteId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PayeeType**](PayeeType.md) |  | [optional] 
**Address** | Pointer to [**CreatePayeeAddress**](CreatePayeeAddress.md) |  | [optional] 
**PaymentChannel** | Pointer to [**CreatePaymentChannel**](CreatePaymentChannel.md) |  | [optional] 
**Challenge** | Pointer to [**Challenge**](Challenge.md) |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Company** | Pointer to [**NullableCompany**](Company.md) |  | [optional] 
**Individual** | Pointer to [**CreateIndividual**](CreateIndividual.md) |  | [optional] 

## Methods

### NewFailedPayee

`func NewFailedPayee() *FailedPayee`

NewFailedPayee instantiates a new FailedPayee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedPayeeWithDefaults

`func NewFailedPayeeWithDefaults() *FailedPayee`

NewFailedPayeeWithDefaults instantiates a new FailedPayee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *FailedPayee) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *FailedPayee) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *FailedPayee) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *FailedPayee) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *FailedPayee) GetPayorRefs() []PayeePayorRefV3`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *FailedPayee) GetPayorRefsOk() (*[]PayeePayorRefV3, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *FailedPayee) SetPayorRefs(v []PayeePayorRefV3)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *FailedPayee) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *FailedPayee) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *FailedPayee) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *FailedPayee) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FailedPayee) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FailedPayee) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FailedPayee) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRemoteId

`func (o *FailedPayee) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *FailedPayee) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *FailedPayee) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *FailedPayee) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetType

`func (o *FailedPayee) GetType() PayeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FailedPayee) GetTypeOk() (*PayeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FailedPayee) SetType(v PayeeType)`

SetType sets Type field to given value.

### HasType

`func (o *FailedPayee) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *FailedPayee) GetAddress() CreatePayeeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FailedPayee) GetAddressOk() (*CreatePayeeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FailedPayee) SetAddress(v CreatePayeeAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FailedPayee) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPaymentChannel

`func (o *FailedPayee) GetPaymentChannel() CreatePaymentChannel`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *FailedPayee) GetPaymentChannelOk() (*CreatePaymentChannel, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *FailedPayee) SetPaymentChannel(v CreatePaymentChannel)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *FailedPayee) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetChallenge

`func (o *FailedPayee) GetChallenge() Challenge`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *FailedPayee) GetChallengeOk() (*Challenge, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *FailedPayee) SetChallenge(v Challenge)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *FailedPayee) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetLanguage

`func (o *FailedPayee) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FailedPayee) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FailedPayee) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *FailedPayee) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCompany

`func (o *FailedPayee) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FailedPayee) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FailedPayee) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *FailedPayee) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *FailedPayee) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *FailedPayee) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndividual

`func (o *FailedPayee) GetIndividual() CreateIndividual`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *FailedPayee) GetIndividualOk() (*CreateIndividual, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *FailedPayee) SetIndividual(v CreateIndividual)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *FailedPayee) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


