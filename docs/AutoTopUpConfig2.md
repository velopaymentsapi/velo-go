# AutoTopUpConfig2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Is auto top-up enabled? automatically trigger funding to top-up the source account balance when the balance falls below the configured minimum level. | 
**MinBalance** | Pointer to **NullableInt64** | When the payor balance falls below this level then auto top-up will be triggered. Note - This is in minor units. | [optional] 
**TargetBalance** | Pointer to **NullableInt64** | When the payor balance falls below the min balance then auto top-up will request funds bring the balance to this level. Note - this is in minor units. | [optional] 
**FundingAccountId** | Pointer to **string** | Id of funding account from which to pull funds when auto top-up is triggered.  Note - if this is not set then auto top-up is effectively disabled. | [optional] 

## Methods

### NewAutoTopUpConfig2

`func NewAutoTopUpConfig2(enabled bool, ) *AutoTopUpConfig2`

NewAutoTopUpConfig2 instantiates a new AutoTopUpConfig2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTopUpConfig2WithDefaults

`func NewAutoTopUpConfig2WithDefaults() *AutoTopUpConfig2`

NewAutoTopUpConfig2WithDefaults instantiates a new AutoTopUpConfig2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutoTopUpConfig2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoTopUpConfig2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoTopUpConfig2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMinBalance

`func (o *AutoTopUpConfig2) GetMinBalance() int64`

GetMinBalance returns the MinBalance field if non-nil, zero value otherwise.

### GetMinBalanceOk

`func (o *AutoTopUpConfig2) GetMinBalanceOk() (*int64, bool)`

GetMinBalanceOk returns a tuple with the MinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBalance

`func (o *AutoTopUpConfig2) SetMinBalance(v int64)`

SetMinBalance sets MinBalance field to given value.

### HasMinBalance

`func (o *AutoTopUpConfig2) HasMinBalance() bool`

HasMinBalance returns a boolean if a field has been set.

### SetMinBalanceNil

`func (o *AutoTopUpConfig2) SetMinBalanceNil(b bool)`

 SetMinBalanceNil sets the value for MinBalance to be an explicit nil

### UnsetMinBalance
`func (o *AutoTopUpConfig2) UnsetMinBalance()`

UnsetMinBalance ensures that no value is present for MinBalance, not even an explicit nil
### GetTargetBalance

`func (o *AutoTopUpConfig2) GetTargetBalance() int64`

GetTargetBalance returns the TargetBalance field if non-nil, zero value otherwise.

### GetTargetBalanceOk

`func (o *AutoTopUpConfig2) GetTargetBalanceOk() (*int64, bool)`

GetTargetBalanceOk returns a tuple with the TargetBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBalance

`func (o *AutoTopUpConfig2) SetTargetBalance(v int64)`

SetTargetBalance sets TargetBalance field to given value.

### HasTargetBalance

`func (o *AutoTopUpConfig2) HasTargetBalance() bool`

HasTargetBalance returns a boolean if a field has been set.

### SetTargetBalanceNil

`func (o *AutoTopUpConfig2) SetTargetBalanceNil(b bool)`

 SetTargetBalanceNil sets the value for TargetBalance to be an explicit nil

### UnsetTargetBalance
`func (o *AutoTopUpConfig2) UnsetTargetBalance()`

UnsetTargetBalance ensures that no value is present for TargetBalance, not even an explicit nil
### GetFundingAccountId

`func (o *AutoTopUpConfig2) GetFundingAccountId() string`

GetFundingAccountId returns the FundingAccountId field if non-nil, zero value otherwise.

### GetFundingAccountIdOk

`func (o *AutoTopUpConfig2) GetFundingAccountIdOk() (*string, bool)`

GetFundingAccountIdOk returns a tuple with the FundingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountId

`func (o *AutoTopUpConfig2) SetFundingAccountId(v string)`

SetFundingAccountId sets FundingAccountId field to given value.

### HasFundingAccountId

`func (o *AutoTopUpConfig2) HasFundingAccountId() bool`

HasFundingAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


