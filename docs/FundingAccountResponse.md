# FundingAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Funding Account Id | [optional] 
**PayorId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** | name on the bank account | [optional] 
**AccountNumber** | Pointer to **string** | bank account number | [optional] 
**RoutingNumber** | Pointer to **string** | bank account routing number | [optional] 
**SourceAccountIds** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | name of funding account | [optional] 
**Currency** | Pointer to **string** | ISO 4217 currency code | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 2 letter country code (upper case) | [optional] 
**Type** | Pointer to **string** | Funding account type | [optional] 

## Methods

### NewFundingAccountResponse

`func NewFundingAccountResponse() *FundingAccountResponse`

NewFundingAccountResponse instantiates a new FundingAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingAccountResponseWithDefaults

`func NewFundingAccountResponseWithDefaults() *FundingAccountResponse`

NewFundingAccountResponseWithDefaults instantiates a new FundingAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FundingAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FundingAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FundingAccountResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FundingAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayorId

`func (o *FundingAccountResponse) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *FundingAccountResponse) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *FundingAccountResponse) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *FundingAccountResponse) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetAccountName

`func (o *FundingAccountResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *FundingAccountResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *FundingAccountResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *FundingAccountResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *FundingAccountResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *FundingAccountResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *FundingAccountResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *FundingAccountResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *FundingAccountResponse) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *FundingAccountResponse) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *FundingAccountResponse) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *FundingAccountResponse) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetSourceAccountIds

`func (o *FundingAccountResponse) GetSourceAccountIds() []string`

GetSourceAccountIds returns the SourceAccountIds field if non-nil, zero value otherwise.

### GetSourceAccountIdsOk

`func (o *FundingAccountResponse) GetSourceAccountIdsOk() (*[]string, bool)`

GetSourceAccountIdsOk returns a tuple with the SourceAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountIds

`func (o *FundingAccountResponse) SetSourceAccountIds(v []string)`

SetSourceAccountIds sets SourceAccountIds field to given value.

### HasSourceAccountIds

`func (o *FundingAccountResponse) HasSourceAccountIds() bool`

HasSourceAccountIds returns a boolean if a field has been set.

### GetName

`func (o *FundingAccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FundingAccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FundingAccountResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FundingAccountResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrency

`func (o *FundingAccountResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FundingAccountResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FundingAccountResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FundingAccountResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCountry

`func (o *FundingAccountResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FundingAccountResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FundingAccountResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FundingAccountResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetType

`func (o *FundingAccountResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FundingAccountResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FundingAccountResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FundingAccountResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


