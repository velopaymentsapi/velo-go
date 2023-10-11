# SourceAccountResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Source Account Id | 
**Balance** | Pointer to **int64** | Decimal implied | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**FundingRef** | **string** |  | 
**PhysicalAccountName** | **string** |  | 
**RailsId** | **string** |  | 
**PayorId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pooled** | **bool** |  | 
**BalanceVisible** | **bool** |  | 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**PhysicalAccountId** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to [**NotificationsV2**](NotificationsV2.md) |  | [optional] 
**FundingAccountId** | Pointer to **NullableString** |  | [optional] 
**AutoTopUpConfig** | Pointer to [**AutoTopUpConfigV2**](AutoTopUpConfigV2.md) |  | [optional] 
**AccountType** | **string** |  | 

## Methods

### NewSourceAccountResponseV2

`func NewSourceAccountResponseV2(id string, fundingRef string, physicalAccountName string, railsId string, pooled bool, balanceVisible bool, accountType string, ) *SourceAccountResponseV2`

NewSourceAccountResponseV2 instantiates a new SourceAccountResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountResponseV2WithDefaults

`func NewSourceAccountResponseV2WithDefaults() *SourceAccountResponseV2`

NewSourceAccountResponseV2WithDefaults instantiates a new SourceAccountResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceAccountResponseV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceAccountResponseV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceAccountResponseV2) SetId(v string)`

SetId sets Id field to given value.


### GetBalance

`func (o *SourceAccountResponseV2) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SourceAccountResponseV2) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SourceAccountResponseV2) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *SourceAccountResponseV2) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *SourceAccountResponseV2) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SourceAccountResponseV2) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SourceAccountResponseV2) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SourceAccountResponseV2) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFundingRef

`func (o *SourceAccountResponseV2) GetFundingRef() string`

GetFundingRef returns the FundingRef field if non-nil, zero value otherwise.

### GetFundingRefOk

`func (o *SourceAccountResponseV2) GetFundingRefOk() (*string, bool)`

GetFundingRefOk returns a tuple with the FundingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRef

`func (o *SourceAccountResponseV2) SetFundingRef(v string)`

SetFundingRef sets FundingRef field to given value.


### GetPhysicalAccountName

`func (o *SourceAccountResponseV2) GetPhysicalAccountName() string`

GetPhysicalAccountName returns the PhysicalAccountName field if non-nil, zero value otherwise.

### GetPhysicalAccountNameOk

`func (o *SourceAccountResponseV2) GetPhysicalAccountNameOk() (*string, bool)`

GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountName

`func (o *SourceAccountResponseV2) SetPhysicalAccountName(v string)`

SetPhysicalAccountName sets PhysicalAccountName field to given value.


### GetRailsId

`func (o *SourceAccountResponseV2) GetRailsId() string`

GetRailsId returns the RailsId field if non-nil, zero value otherwise.

### GetRailsIdOk

`func (o *SourceAccountResponseV2) GetRailsIdOk() (*string, bool)`

GetRailsIdOk returns a tuple with the RailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsId

`func (o *SourceAccountResponseV2) SetRailsId(v string)`

SetRailsId sets RailsId field to given value.


### GetPayorId

`func (o *SourceAccountResponseV2) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *SourceAccountResponseV2) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *SourceAccountResponseV2) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *SourceAccountResponseV2) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetName

`func (o *SourceAccountResponseV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceAccountResponseV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceAccountResponseV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceAccountResponseV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPooled

`func (o *SourceAccountResponseV2) GetPooled() bool`

GetPooled returns the Pooled field if non-nil, zero value otherwise.

### GetPooledOk

`func (o *SourceAccountResponseV2) GetPooledOk() (*bool, bool)`

GetPooledOk returns a tuple with the Pooled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooled

`func (o *SourceAccountResponseV2) SetPooled(v bool)`

SetPooled sets Pooled field to given value.


### GetBalanceVisible

`func (o *SourceAccountResponseV2) GetBalanceVisible() bool`

GetBalanceVisible returns the BalanceVisible field if non-nil, zero value otherwise.

### GetBalanceVisibleOk

`func (o *SourceAccountResponseV2) GetBalanceVisibleOk() (*bool, bool)`

GetBalanceVisibleOk returns a tuple with the BalanceVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceVisible

`func (o *SourceAccountResponseV2) SetBalanceVisible(v bool)`

SetBalanceVisible sets BalanceVisible field to given value.


### GetCustomerId

`func (o *SourceAccountResponseV2) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SourceAccountResponseV2) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SourceAccountResponseV2) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SourceAccountResponseV2) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *SourceAccountResponseV2) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *SourceAccountResponseV2) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetPhysicalAccountId

`func (o *SourceAccountResponseV2) GetPhysicalAccountId() string`

GetPhysicalAccountId returns the PhysicalAccountId field if non-nil, zero value otherwise.

### GetPhysicalAccountIdOk

`func (o *SourceAccountResponseV2) GetPhysicalAccountIdOk() (*string, bool)`

GetPhysicalAccountIdOk returns a tuple with the PhysicalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountId

`func (o *SourceAccountResponseV2) SetPhysicalAccountId(v string)`

SetPhysicalAccountId sets PhysicalAccountId field to given value.

### HasPhysicalAccountId

`func (o *SourceAccountResponseV2) HasPhysicalAccountId() bool`

HasPhysicalAccountId returns a boolean if a field has been set.

### GetNotifications

`func (o *SourceAccountResponseV2) GetNotifications() NotificationsV2`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SourceAccountResponseV2) GetNotificationsOk() (*NotificationsV2, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SourceAccountResponseV2) SetNotifications(v NotificationsV2)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SourceAccountResponseV2) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetFundingAccountId

`func (o *SourceAccountResponseV2) GetFundingAccountId() string`

GetFundingAccountId returns the FundingAccountId field if non-nil, zero value otherwise.

### GetFundingAccountIdOk

`func (o *SourceAccountResponseV2) GetFundingAccountIdOk() (*string, bool)`

GetFundingAccountIdOk returns a tuple with the FundingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountId

`func (o *SourceAccountResponseV2) SetFundingAccountId(v string)`

SetFundingAccountId sets FundingAccountId field to given value.

### HasFundingAccountId

`func (o *SourceAccountResponseV2) HasFundingAccountId() bool`

HasFundingAccountId returns a boolean if a field has been set.

### SetFundingAccountIdNil

`func (o *SourceAccountResponseV2) SetFundingAccountIdNil(b bool)`

 SetFundingAccountIdNil sets the value for FundingAccountId to be an explicit nil

### UnsetFundingAccountId
`func (o *SourceAccountResponseV2) UnsetFundingAccountId()`

UnsetFundingAccountId ensures that no value is present for FundingAccountId, not even an explicit nil
### GetAutoTopUpConfig

`func (o *SourceAccountResponseV2) GetAutoTopUpConfig() AutoTopUpConfigV2`

GetAutoTopUpConfig returns the AutoTopUpConfig field if non-nil, zero value otherwise.

### GetAutoTopUpConfigOk

`func (o *SourceAccountResponseV2) GetAutoTopUpConfigOk() (*AutoTopUpConfigV2, bool)`

GetAutoTopUpConfigOk returns a tuple with the AutoTopUpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopUpConfig

`func (o *SourceAccountResponseV2) SetAutoTopUpConfig(v AutoTopUpConfigV2)`

SetAutoTopUpConfig sets AutoTopUpConfig field to given value.

### HasAutoTopUpConfig

`func (o *SourceAccountResponseV2) HasAutoTopUpConfig() bool`

HasAutoTopUpConfig returns a boolean if a field has been set.

### GetAccountType

`func (o *SourceAccountResponseV2) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *SourceAccountResponseV2) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *SourceAccountResponseV2) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


