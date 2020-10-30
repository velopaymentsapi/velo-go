# SourceAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Source Account Id | [optional] 
**Balance** | Pointer to **int64** | Decimal implied | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**FundingRef** | Pointer to **string** |  | [optional] 
**PhysicalAccountName** | Pointer to **string** |  | [optional] 
**RailsId** | Pointer to **string** |  | [optional] 
**PayorId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pooled** | Pointer to **bool** |  | [optional] 
**BalanceVisible** | Pointer to **bool** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**PhysicalAccountId** | Pointer to **string** |  | [optional] 
**FundingAccountId** | Pointer to **NullableString** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 

## Methods

### NewSourceAccountResponse

`func NewSourceAccountResponse() *SourceAccountResponse`

NewSourceAccountResponse instantiates a new SourceAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountResponseWithDefaults

`func NewSourceAccountResponseWithDefaults() *SourceAccountResponse`

NewSourceAccountResponseWithDefaults instantiates a new SourceAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceAccountResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SourceAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBalance

`func (o *SourceAccountResponse) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SourceAccountResponse) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SourceAccountResponse) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *SourceAccountResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *SourceAccountResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SourceAccountResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SourceAccountResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SourceAccountResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFundingRef

`func (o *SourceAccountResponse) GetFundingRef() string`

GetFundingRef returns the FundingRef field if non-nil, zero value otherwise.

### GetFundingRefOk

`func (o *SourceAccountResponse) GetFundingRefOk() (*string, bool)`

GetFundingRefOk returns a tuple with the FundingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRef

`func (o *SourceAccountResponse) SetFundingRef(v string)`

SetFundingRef sets FundingRef field to given value.

### HasFundingRef

`func (o *SourceAccountResponse) HasFundingRef() bool`

HasFundingRef returns a boolean if a field has been set.

### GetPhysicalAccountName

`func (o *SourceAccountResponse) GetPhysicalAccountName() string`

GetPhysicalAccountName returns the PhysicalAccountName field if non-nil, zero value otherwise.

### GetPhysicalAccountNameOk

`func (o *SourceAccountResponse) GetPhysicalAccountNameOk() (*string, bool)`

GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountName

`func (o *SourceAccountResponse) SetPhysicalAccountName(v string)`

SetPhysicalAccountName sets PhysicalAccountName field to given value.

### HasPhysicalAccountName

`func (o *SourceAccountResponse) HasPhysicalAccountName() bool`

HasPhysicalAccountName returns a boolean if a field has been set.

### GetRailsId

`func (o *SourceAccountResponse) GetRailsId() string`

GetRailsId returns the RailsId field if non-nil, zero value otherwise.

### GetRailsIdOk

`func (o *SourceAccountResponse) GetRailsIdOk() (*string, bool)`

GetRailsIdOk returns a tuple with the RailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsId

`func (o *SourceAccountResponse) SetRailsId(v string)`

SetRailsId sets RailsId field to given value.

### HasRailsId

`func (o *SourceAccountResponse) HasRailsId() bool`

HasRailsId returns a boolean if a field has been set.

### GetPayorId

`func (o *SourceAccountResponse) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *SourceAccountResponse) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *SourceAccountResponse) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *SourceAccountResponse) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetName

`func (o *SourceAccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceAccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceAccountResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceAccountResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPooled

`func (o *SourceAccountResponse) GetPooled() bool`

GetPooled returns the Pooled field if non-nil, zero value otherwise.

### GetPooledOk

`func (o *SourceAccountResponse) GetPooledOk() (*bool, bool)`

GetPooledOk returns a tuple with the Pooled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooled

`func (o *SourceAccountResponse) SetPooled(v bool)`

SetPooled sets Pooled field to given value.

### HasPooled

`func (o *SourceAccountResponse) HasPooled() bool`

HasPooled returns a boolean if a field has been set.

### GetBalanceVisible

`func (o *SourceAccountResponse) GetBalanceVisible() bool`

GetBalanceVisible returns the BalanceVisible field if non-nil, zero value otherwise.

### GetBalanceVisibleOk

`func (o *SourceAccountResponse) GetBalanceVisibleOk() (*bool, bool)`

GetBalanceVisibleOk returns a tuple with the BalanceVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceVisible

`func (o *SourceAccountResponse) SetBalanceVisible(v bool)`

SetBalanceVisible sets BalanceVisible field to given value.

### HasBalanceVisible

`func (o *SourceAccountResponse) HasBalanceVisible() bool`

HasBalanceVisible returns a boolean if a field has been set.

### GetCustomerId

`func (o *SourceAccountResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SourceAccountResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SourceAccountResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SourceAccountResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *SourceAccountResponse) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *SourceAccountResponse) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetPhysicalAccountId

`func (o *SourceAccountResponse) GetPhysicalAccountId() string`

GetPhysicalAccountId returns the PhysicalAccountId field if non-nil, zero value otherwise.

### GetPhysicalAccountIdOk

`func (o *SourceAccountResponse) GetPhysicalAccountIdOk() (*string, bool)`

GetPhysicalAccountIdOk returns a tuple with the PhysicalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountId

`func (o *SourceAccountResponse) SetPhysicalAccountId(v string)`

SetPhysicalAccountId sets PhysicalAccountId field to given value.

### HasPhysicalAccountId

`func (o *SourceAccountResponse) HasPhysicalAccountId() bool`

HasPhysicalAccountId returns a boolean if a field has been set.

### GetFundingAccountId

`func (o *SourceAccountResponse) GetFundingAccountId() string`

GetFundingAccountId returns the FundingAccountId field if non-nil, zero value otherwise.

### GetFundingAccountIdOk

`func (o *SourceAccountResponse) GetFundingAccountIdOk() (*string, bool)`

GetFundingAccountIdOk returns a tuple with the FundingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountId

`func (o *SourceAccountResponse) SetFundingAccountId(v string)`

SetFundingAccountId sets FundingAccountId field to given value.

### HasFundingAccountId

`func (o *SourceAccountResponse) HasFundingAccountId() bool`

HasFundingAccountId returns a boolean if a field has been set.

### SetFundingAccountIdNil

`func (o *SourceAccountResponse) SetFundingAccountIdNil(b bool)`

 SetFundingAccountIdNil sets the value for FundingAccountId to be an explicit nil

### UnsetFundingAccountId
`func (o *SourceAccountResponse) UnsetFundingAccountId()`

UnsetFundingAccountId ensures that no value is present for FundingAccountId, not even an explicit nil
### GetAccountType

`func (o *SourceAccountResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *SourceAccountResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *SourceAccountResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *SourceAccountResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


