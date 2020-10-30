# SourceAccountSummaryV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAccountId** | **string** |  | 
**TotalCost** | **int64** |  | 
**Currency** | Pointer to [**PaymentAuditCurrencyV4**](PaymentAuditCurrencyV4.md) |  | [optional] 

## Methods

### NewSourceAccountSummaryV4

`func NewSourceAccountSummaryV4(sourceAccountId string, totalCost int64, ) *SourceAccountSummaryV4`

NewSourceAccountSummaryV4 instantiates a new SourceAccountSummaryV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountSummaryV4WithDefaults

`func NewSourceAccountSummaryV4WithDefaults() *SourceAccountSummaryV4`

NewSourceAccountSummaryV4WithDefaults instantiates a new SourceAccountSummaryV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAccountId

`func (o *SourceAccountSummaryV4) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *SourceAccountSummaryV4) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *SourceAccountSummaryV4) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.


### GetTotalCost

`func (o *SourceAccountSummaryV4) GetTotalCost() int64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *SourceAccountSummaryV4) GetTotalCostOk() (*int64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *SourceAccountSummaryV4) SetTotalCost(v int64)`

SetTotalCost sets TotalCost field to given value.


### GetCurrency

`func (o *SourceAccountSummaryV4) GetCurrency() PaymentAuditCurrencyV4`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SourceAccountSummaryV4) GetCurrencyOk() (*PaymentAuditCurrencyV4, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SourceAccountSummaryV4) SetCurrency(v PaymentAuditCurrencyV4)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SourceAccountSummaryV4) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


