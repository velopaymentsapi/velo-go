# SourceAccountV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAccountName** | **string** | The name of the source account as referencec in the payout | 
**SourceAccountId** | **string** | The id of the payout | 
**Currency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**TotalPayoutCost** | **int32** | The total amount (in mnor units) that will be debited from the source account for the payout | 

## Methods

### NewSourceAccountV3

`func NewSourceAccountV3(sourceAccountName string, sourceAccountId string, currency string, totalPayoutCost int32, ) *SourceAccountV3`

NewSourceAccountV3 instantiates a new SourceAccountV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountV3WithDefaults

`func NewSourceAccountV3WithDefaults() *SourceAccountV3`

NewSourceAccountV3WithDefaults instantiates a new SourceAccountV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAccountName

`func (o *SourceAccountV3) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *SourceAccountV3) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *SourceAccountV3) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.


### GetSourceAccountId

`func (o *SourceAccountV3) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *SourceAccountV3) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *SourceAccountV3) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.


### GetCurrency

`func (o *SourceAccountV3) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SourceAccountV3) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SourceAccountV3) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTotalPayoutCost

`func (o *SourceAccountV3) GetTotalPayoutCost() int32`

GetTotalPayoutCost returns the TotalPayoutCost field if non-nil, zero value otherwise.

### GetTotalPayoutCostOk

`func (o *SourceAccountV3) GetTotalPayoutCostOk() (*int32, bool)`

GetTotalPayoutCostOk returns a tuple with the TotalPayoutCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPayoutCost

`func (o *SourceAccountV3) SetTotalPayoutCost(v int32)`

SetTotalPayoutCost sets TotalPayoutCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


