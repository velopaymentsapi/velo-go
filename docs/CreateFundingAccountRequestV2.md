# CreateFundingAccountRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**PayorId** | **string** |  | 
**AccountName** | Pointer to **string** | Required if type is FBO | [optional] 
**AccountNumber** | Pointer to **string** | Required if type is FBO | [optional] 
**RoutingNumber** | Pointer to **string** | Required if type is FBO | [optional] 
**Currency** | Pointer to **string** | ISO 4217 currency code, Required if type is WUBS_DECOUPLED | [optional] 

## Methods

### NewCreateFundingAccountRequestV2

`func NewCreateFundingAccountRequestV2(type_ string, name string, payorId string, ) *CreateFundingAccountRequestV2`

NewCreateFundingAccountRequestV2 instantiates a new CreateFundingAccountRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFundingAccountRequestV2WithDefaults

`func NewCreateFundingAccountRequestV2WithDefaults() *CreateFundingAccountRequestV2`

NewCreateFundingAccountRequestV2WithDefaults instantiates a new CreateFundingAccountRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateFundingAccountRequestV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateFundingAccountRequestV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateFundingAccountRequestV2) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CreateFundingAccountRequestV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFundingAccountRequestV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFundingAccountRequestV2) SetName(v string)`

SetName sets Name field to given value.


### GetPayorId

`func (o *CreateFundingAccountRequestV2) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *CreateFundingAccountRequestV2) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *CreateFundingAccountRequestV2) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetAccountName

`func (o *CreateFundingAccountRequestV2) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CreateFundingAccountRequestV2) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CreateFundingAccountRequestV2) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *CreateFundingAccountRequestV2) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *CreateFundingAccountRequestV2) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreateFundingAccountRequestV2) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreateFundingAccountRequestV2) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CreateFundingAccountRequestV2) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *CreateFundingAccountRequestV2) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *CreateFundingAccountRequestV2) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *CreateFundingAccountRequestV2) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *CreateFundingAccountRequestV2) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateFundingAccountRequestV2) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateFundingAccountRequestV2) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateFundingAccountRequestV2) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateFundingAccountRequestV2) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


