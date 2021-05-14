# FxSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** |  | 
**CreationDateTime** | **time.Time** |  | 
**Rate** | **float64** |  | 
**InvertedRate** | **float64** |  | 
**TotalCost** | **int32** |  | 
**TotalPaymentAmount** | **int32** |  | 
**SourceCurrency** | Pointer to [**PaymentAuditCurrency**](PaymentAuditCurrency.md) |  | [optional] 
**PaymentCurrency** | Pointer to [**PaymentAuditCurrency**](PaymentAuditCurrency.md) |  | [optional] 
**Status** | **string** |  | 
**FundingStatus** | **string** |  | 

## Methods

### NewFxSummary

`func NewFxSummary(quoteId string, creationDateTime time.Time, rate float64, invertedRate float64, totalCost int32, totalPaymentAmount int32, status string, fundingStatus string, ) *FxSummary`

NewFxSummary instantiates a new FxSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSummaryWithDefaults

`func NewFxSummaryWithDefaults() *FxSummary`

NewFxSummaryWithDefaults instantiates a new FxSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *FxSummary) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *FxSummary) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *FxSummary) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetCreationDateTime

`func (o *FxSummary) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *FxSummary) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *FxSummary) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.


### GetRate

`func (o *FxSummary) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *FxSummary) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *FxSummary) SetRate(v float64)`

SetRate sets Rate field to given value.


### GetInvertedRate

`func (o *FxSummary) GetInvertedRate() float64`

GetInvertedRate returns the InvertedRate field if non-nil, zero value otherwise.

### GetInvertedRateOk

`func (o *FxSummary) GetInvertedRateOk() (*float64, bool)`

GetInvertedRateOk returns a tuple with the InvertedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertedRate

`func (o *FxSummary) SetInvertedRate(v float64)`

SetInvertedRate sets InvertedRate field to given value.


### GetTotalCost

`func (o *FxSummary) GetTotalCost() int32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *FxSummary) GetTotalCostOk() (*int32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *FxSummary) SetTotalCost(v int32)`

SetTotalCost sets TotalCost field to given value.


### GetTotalPaymentAmount

`func (o *FxSummary) GetTotalPaymentAmount() int32`

GetTotalPaymentAmount returns the TotalPaymentAmount field if non-nil, zero value otherwise.

### GetTotalPaymentAmountOk

`func (o *FxSummary) GetTotalPaymentAmountOk() (*int32, bool)`

GetTotalPaymentAmountOk returns a tuple with the TotalPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaymentAmount

`func (o *FxSummary) SetTotalPaymentAmount(v int32)`

SetTotalPaymentAmount sets TotalPaymentAmount field to given value.


### GetSourceCurrency

`func (o *FxSummary) GetSourceCurrency() PaymentAuditCurrency`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *FxSummary) GetSourceCurrencyOk() (*PaymentAuditCurrency, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *FxSummary) SetSourceCurrency(v PaymentAuditCurrency)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *FxSummary) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *FxSummary) GetPaymentCurrency() PaymentAuditCurrency`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *FxSummary) GetPaymentCurrencyOk() (*PaymentAuditCurrency, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *FxSummary) SetPaymentCurrency(v PaymentAuditCurrency)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *FxSummary) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *FxSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FxSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FxSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFundingStatus

`func (o *FxSummary) GetFundingStatus() string`

GetFundingStatus returns the FundingStatus field if non-nil, zero value otherwise.

### GetFundingStatusOk

`func (o *FxSummary) GetFundingStatusOk() (*string, bool)`

GetFundingStatusOk returns a tuple with the FundingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingStatus

`func (o *FxSummary) SetFundingStatus(v string)`

SetFundingStatus sets FundingStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


