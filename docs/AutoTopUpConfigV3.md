# AutoTopUpConfigV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Is auto top-up enabled? automatically trigger funding to top-up the source account balance when the balance falls below the configured minimum level. | 
**MinBalance** | Pointer to **NullableInt64** | When the payor balance falls below this level then auto top-up will be triggered. Note - This is in minor units. | [optional] 
**TargetBalance** | Pointer to **NullableInt64** | When the payor balance falls below the min balance then auto top-up will request funds bring the balance to this level. Note - this is in minor units. | [optional] 
**FundingAccountId** | Pointer to **string** | Id of funding account from which to pull funds when auto top-up is triggered.  Note - if this is not set then auto top-up is effectively disabled. | [optional] 

## Methods

### NewAutoTopUpConfigV3

`func NewAutoTopUpConfigV3(enabled bool, ) *AutoTopUpConfigV3`

NewAutoTopUpConfigV3 instantiates a new AutoTopUpConfigV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTopUpConfigV3WithDefaults

`func NewAutoTopUpConfigV3WithDefaults() *AutoTopUpConfigV3`

NewAutoTopUpConfigV3WithDefaults instantiates a new AutoTopUpConfigV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutoTopUpConfigV3) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoTopUpConfigV3) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoTopUpConfigV3) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMinBalance

`func (o *AutoTopUpConfigV3) GetMinBalance() int64`

GetMinBalance returns the MinBalance field if non-nil, zero value otherwise.

### GetMinBalanceOk

`func (o *AutoTopUpConfigV3) GetMinBalanceOk() (*int64, bool)`

GetMinBalanceOk returns a tuple with the MinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBalance

`func (o *AutoTopUpConfigV3) SetMinBalance(v int64)`

SetMinBalance sets MinBalance field to given value.

### HasMinBalance

`func (o *AutoTopUpConfigV3) HasMinBalance() bool`

HasMinBalance returns a boolean if a field has been set.

### SetMinBalanceNil

`func (o *AutoTopUpConfigV3) SetMinBalanceNil(b bool)`

 SetMinBalanceNil sets the value for MinBalance to be an explicit nil

### UnsetMinBalance
`func (o *AutoTopUpConfigV3) UnsetMinBalance()`

UnsetMinBalance ensures that no value is present for MinBalance, not even an explicit nil
### GetTargetBalance

`func (o *AutoTopUpConfigV3) GetTargetBalance() int64`

GetTargetBalance returns the TargetBalance field if non-nil, zero value otherwise.

### GetTargetBalanceOk

`func (o *AutoTopUpConfigV3) GetTargetBalanceOk() (*int64, bool)`

GetTargetBalanceOk returns a tuple with the TargetBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBalance

`func (o *AutoTopUpConfigV3) SetTargetBalance(v int64)`

SetTargetBalance sets TargetBalance field to given value.

### HasTargetBalance

`func (o *AutoTopUpConfigV3) HasTargetBalance() bool`

HasTargetBalance returns a boolean if a field has been set.

### SetTargetBalanceNil

`func (o *AutoTopUpConfigV3) SetTargetBalanceNil(b bool)`

 SetTargetBalanceNil sets the value for TargetBalance to be an explicit nil

### UnsetTargetBalance
`func (o *AutoTopUpConfigV3) UnsetTargetBalance()`

UnsetTargetBalance ensures that no value is present for TargetBalance, not even an explicit nil
### GetFundingAccountId

`func (o *AutoTopUpConfigV3) GetFundingAccountId() string`

GetFundingAccountId returns the FundingAccountId field if non-nil, zero value otherwise.

### GetFundingAccountIdOk

`func (o *AutoTopUpConfigV3) GetFundingAccountIdOk() (*string, bool)`

GetFundingAccountIdOk returns a tuple with the FundingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountId

`func (o *AutoTopUpConfigV3) SetFundingAccountId(v string)`

SetFundingAccountId sets FundingAccountId field to given value.

### HasFundingAccountId

`func (o *AutoTopUpConfigV3) HasFundingAccountId() bool`

HasFundingAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


