# SourceAccountSummaryV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAccountId** | **string** |  | 
**TotalCost** | **int64** |  | 
**Currency** | Pointer to **string** | ISO 3 character currency code | [optional] 

## Methods

### NewSourceAccountSummaryV3

`func NewSourceAccountSummaryV3(sourceAccountId string, totalCost int64, ) *SourceAccountSummaryV3`

NewSourceAccountSummaryV3 instantiates a new SourceAccountSummaryV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountSummaryV3WithDefaults

`func NewSourceAccountSummaryV3WithDefaults() *SourceAccountSummaryV3`

NewSourceAccountSummaryV3WithDefaults instantiates a new SourceAccountSummaryV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAccountId

`func (o *SourceAccountSummaryV3) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *SourceAccountSummaryV3) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *SourceAccountSummaryV3) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.


### GetTotalCost

`func (o *SourceAccountSummaryV3) GetTotalCost() int64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *SourceAccountSummaryV3) GetTotalCostOk() (*int64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *SourceAccountSummaryV3) SetTotalCost(v int64)`

SetTotalCost sets TotalCost field to given value.


### GetCurrency

`func (o *SourceAccountSummaryV3) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SourceAccountSummaryV3) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SourceAccountSummaryV3) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SourceAccountSummaryV3) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


