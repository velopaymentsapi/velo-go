# SourceAccountSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAccountId** | **string** |  | 
**TotalCost** | **int64** |  | 
**Currency** | Pointer to **string** | ISO-4217 3 character currency code | [optional] 

## Methods

### NewSourceAccountSummary

`func NewSourceAccountSummary(sourceAccountId string, totalCost int64, ) *SourceAccountSummary`

NewSourceAccountSummary instantiates a new SourceAccountSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountSummaryWithDefaults

`func NewSourceAccountSummaryWithDefaults() *SourceAccountSummary`

NewSourceAccountSummaryWithDefaults instantiates a new SourceAccountSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAccountId

`func (o *SourceAccountSummary) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *SourceAccountSummary) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *SourceAccountSummary) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.


### GetTotalCost

`func (o *SourceAccountSummary) GetTotalCost() int64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *SourceAccountSummary) GetTotalCostOk() (*int64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *SourceAccountSummary) SetTotalCost(v int64)`

SetTotalCost sets TotalCost field to given value.


### GetCurrency

`func (o *SourceAccountSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SourceAccountSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SourceAccountSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SourceAccountSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


