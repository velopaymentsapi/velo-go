# FundingPayorStatusAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundingId** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewFundingPayorStatusAuditResponse

`func NewFundingPayorStatusAuditResponse() *FundingPayorStatusAuditResponse`

NewFundingPayorStatusAuditResponse instantiates a new FundingPayorStatusAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingPayorStatusAuditResponseWithDefaults

`func NewFundingPayorStatusAuditResponseWithDefaults() *FundingPayorStatusAuditResponse`

NewFundingPayorStatusAuditResponseWithDefaults instantiates a new FundingPayorStatusAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundingId

`func (o *FundingPayorStatusAuditResponse) GetFundingId() string`

GetFundingId returns the FundingId field if non-nil, zero value otherwise.

### GetFundingIdOk

`func (o *FundingPayorStatusAuditResponse) GetFundingIdOk() (*string, bool)`

GetFundingIdOk returns a tuple with the FundingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingId

`func (o *FundingPayorStatusAuditResponse) SetFundingId(v string)`

SetFundingId sets FundingId field to given value.

### HasFundingId

`func (o *FundingPayorStatusAuditResponse) HasFundingId() bool`

HasFundingId returns a boolean if a field has been set.

### GetAmount

`func (o *FundingPayorStatusAuditResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FundingPayorStatusAuditResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FundingPayorStatusAuditResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FundingPayorStatusAuditResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *FundingPayorStatusAuditResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FundingPayorStatusAuditResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FundingPayorStatusAuditResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FundingPayorStatusAuditResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *FundingPayorStatusAuditResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FundingPayorStatusAuditResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FundingPayorStatusAuditResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FundingPayorStatusAuditResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


