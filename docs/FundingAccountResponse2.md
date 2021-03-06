# FundingAccountResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Funding Account Id | [optional] 
**PayorId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** | name on the bank account | [optional] 
**AccountNumber** | Pointer to **string** | bank account number | [optional] 
**RoutingNumber** | Pointer to **string** | bank account routing number | [optional] 
**Name** | Pointer to **string** | name of funding account | [optional] 
**Currency** | Pointer to **string** | ISO 4217 currency code | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 2 letter country code (upper case) | [optional] 
**Type** | Pointer to **string** | Funding account type | [optional] 
**Archived** | Pointer to **bool** | A flag for whether the funding account has been archived.  Only present in the response if true. | [optional] 

## Methods

### NewFundingAccountResponse2

`func NewFundingAccountResponse2() *FundingAccountResponse2`

NewFundingAccountResponse2 instantiates a new FundingAccountResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingAccountResponse2WithDefaults

`func NewFundingAccountResponse2WithDefaults() *FundingAccountResponse2`

NewFundingAccountResponse2WithDefaults instantiates a new FundingAccountResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FundingAccountResponse2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FundingAccountResponse2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FundingAccountResponse2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FundingAccountResponse2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayorId

`func (o *FundingAccountResponse2) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *FundingAccountResponse2) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *FundingAccountResponse2) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *FundingAccountResponse2) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetAccountName

`func (o *FundingAccountResponse2) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *FundingAccountResponse2) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *FundingAccountResponse2) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *FundingAccountResponse2) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *FundingAccountResponse2) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *FundingAccountResponse2) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *FundingAccountResponse2) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *FundingAccountResponse2) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *FundingAccountResponse2) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *FundingAccountResponse2) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *FundingAccountResponse2) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *FundingAccountResponse2) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetName

`func (o *FundingAccountResponse2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FundingAccountResponse2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FundingAccountResponse2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FundingAccountResponse2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrency

`func (o *FundingAccountResponse2) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FundingAccountResponse2) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FundingAccountResponse2) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FundingAccountResponse2) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCountry

`func (o *FundingAccountResponse2) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FundingAccountResponse2) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FundingAccountResponse2) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FundingAccountResponse2) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetType

`func (o *FundingAccountResponse2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FundingAccountResponse2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FundingAccountResponse2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FundingAccountResponse2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetArchived

`func (o *FundingAccountResponse2) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FundingAccountResponse2) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FundingAccountResponse2) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *FundingAccountResponse2) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


