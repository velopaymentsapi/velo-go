# FxSummaryV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** |  | 
**CreationDateTime** | [**time.Time**](time.Time.md) |  | 
**Rate** | **float64** |  | 
**InvertedRate** | **float64** |  | 
**TotalCost** | **int32** |  | 
**TotalPaymentAmount** | **int32** |  | 
**SourceCurrency** | Pointer to [**PaymentAuditCurrencyV4**](PaymentAuditCurrencyV4.md) |  | [optional] 
**PaymentCurrency** | Pointer to [**PaymentAuditCurrencyV4**](PaymentAuditCurrencyV4.md) |  | [optional] 
**Status** | **string** |  | 
**FundingStatus** | **string** |  | 

## Methods

### NewFxSummaryV4

`func NewFxSummaryV4(quoteId string, creationDateTime time.Time, rate float64, invertedRate float64, totalCost int32, totalPaymentAmount int32, status string, fundingStatus string, ) *FxSummaryV4`

NewFxSummaryV4 instantiates a new FxSummaryV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSummaryV4WithDefaults

`func NewFxSummaryV4WithDefaults() *FxSummaryV4`

NewFxSummaryV4WithDefaults instantiates a new FxSummaryV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *FxSummaryV4) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *FxSummaryV4) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *FxSummaryV4) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetCreationDateTime

`func (o *FxSummaryV4) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *FxSummaryV4) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *FxSummaryV4) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.


### GetRate

`func (o *FxSummaryV4) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *FxSummaryV4) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *FxSummaryV4) SetRate(v float64)`

SetRate sets Rate field to given value.


### GetInvertedRate

`func (o *FxSummaryV4) GetInvertedRate() float64`

GetInvertedRate returns the InvertedRate field if non-nil, zero value otherwise.

### GetInvertedRateOk

`func (o *FxSummaryV4) GetInvertedRateOk() (*float64, bool)`

GetInvertedRateOk returns a tuple with the InvertedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertedRate

`func (o *FxSummaryV4) SetInvertedRate(v float64)`

SetInvertedRate sets InvertedRate field to given value.


### GetTotalCost

`func (o *FxSummaryV4) GetTotalCost() int32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *FxSummaryV4) GetTotalCostOk() (*int32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *FxSummaryV4) SetTotalCost(v int32)`

SetTotalCost sets TotalCost field to given value.


### GetTotalPaymentAmount

`func (o *FxSummaryV4) GetTotalPaymentAmount() int32`

GetTotalPaymentAmount returns the TotalPaymentAmount field if non-nil, zero value otherwise.

### GetTotalPaymentAmountOk

`func (o *FxSummaryV4) GetTotalPaymentAmountOk() (*int32, bool)`

GetTotalPaymentAmountOk returns a tuple with the TotalPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaymentAmount

`func (o *FxSummaryV4) SetTotalPaymentAmount(v int32)`

SetTotalPaymentAmount sets TotalPaymentAmount field to given value.


### GetSourceCurrency

`func (o *FxSummaryV4) GetSourceCurrency() PaymentAuditCurrencyV4`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *FxSummaryV4) GetSourceCurrencyOk() (*PaymentAuditCurrencyV4, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *FxSummaryV4) SetSourceCurrency(v PaymentAuditCurrencyV4)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *FxSummaryV4) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *FxSummaryV4) GetPaymentCurrency() PaymentAuditCurrencyV4`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *FxSummaryV4) GetPaymentCurrencyOk() (*PaymentAuditCurrencyV4, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *FxSummaryV4) SetPaymentCurrency(v PaymentAuditCurrencyV4)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *FxSummaryV4) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *FxSummaryV4) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FxSummaryV4) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FxSummaryV4) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFundingStatus

`func (o *FxSummaryV4) GetFundingStatus() string`

GetFundingStatus returns the FundingStatus field if non-nil, zero value otherwise.

### GetFundingStatusOk

`func (o *FxSummaryV4) GetFundingStatusOk() (*string, bool)`

GetFundingStatusOk returns a tuple with the FundingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingStatus

`func (o *FxSummaryV4) SetFundingStatus(v string)`

SetFundingStatus sets FundingStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


