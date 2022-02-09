# FailedPayee2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRef**](PayeePayorRef.md) |  | [optional] [readonly] 
**Email** | Pointer to **string** |  | [optional] 
**RemoteId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PayeeType2**](PayeeType2.md) |  | [optional] 
**Address** | Pointer to [**CreatePayeeAddress2**](CreatePayeeAddress2.md) |  | [optional] 
**PaymentChannel** | Pointer to [**CreatePaymentChannel2**](CreatePaymentChannel2.md) |  | [optional] 
**Challenge** | Pointer to [**Challenge2**](Challenge2.md) |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Company** | Pointer to [**NullableCompany2**](Company2.md) |  | [optional] 
**Individual** | Pointer to [**CreateIndividual2**](CreateIndividual2.md) |  | [optional] 

## Methods

### NewFailedPayee2

`func NewFailedPayee2() *FailedPayee2`

NewFailedPayee2 instantiates a new FailedPayee2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedPayee2WithDefaults

`func NewFailedPayee2WithDefaults() *FailedPayee2`

NewFailedPayee2WithDefaults instantiates a new FailedPayee2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *FailedPayee2) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *FailedPayee2) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *FailedPayee2) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *FailedPayee2) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *FailedPayee2) GetPayorRefs() []PayeePayorRef`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *FailedPayee2) GetPayorRefsOk() (*[]PayeePayorRef, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *FailedPayee2) SetPayorRefs(v []PayeePayorRef)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *FailedPayee2) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *FailedPayee2) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *FailedPayee2) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *FailedPayee2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FailedPayee2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FailedPayee2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FailedPayee2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRemoteId

`func (o *FailedPayee2) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *FailedPayee2) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *FailedPayee2) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *FailedPayee2) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetType

`func (o *FailedPayee2) GetType() PayeeType2`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FailedPayee2) GetTypeOk() (*PayeeType2, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FailedPayee2) SetType(v PayeeType2)`

SetType sets Type field to given value.

### HasType

`func (o *FailedPayee2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *FailedPayee2) GetAddress() CreatePayeeAddress2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FailedPayee2) GetAddressOk() (*CreatePayeeAddress2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FailedPayee2) SetAddress(v CreatePayeeAddress2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FailedPayee2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPaymentChannel

`func (o *FailedPayee2) GetPaymentChannel() CreatePaymentChannel2`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *FailedPayee2) GetPaymentChannelOk() (*CreatePaymentChannel2, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *FailedPayee2) SetPaymentChannel(v CreatePaymentChannel2)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *FailedPayee2) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetChallenge

`func (o *FailedPayee2) GetChallenge() Challenge2`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *FailedPayee2) GetChallengeOk() (*Challenge2, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *FailedPayee2) SetChallenge(v Challenge2)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *FailedPayee2) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetLanguage

`func (o *FailedPayee2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FailedPayee2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FailedPayee2) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *FailedPayee2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCompany

`func (o *FailedPayee2) GetCompany() Company2`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FailedPayee2) GetCompanyOk() (*Company2, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FailedPayee2) SetCompany(v Company2)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *FailedPayee2) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *FailedPayee2) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *FailedPayee2) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndividual

`func (o *FailedPayee2) GetIndividual() CreateIndividual2`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *FailedPayee2) GetIndividualOk() (*CreateIndividual2, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *FailedPayee2) SetIndividual(v CreateIndividual2)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *FailedPayee2) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


