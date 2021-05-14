# FxSummaryV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** |  | 
**CreationDateTime** | **time.Time** |  | 
**Rate** | **float32** |  | 
**InvertedRate** | **float32** |  | 
**TotalCost** | **int32** |  | 
**TotalPaymentAmount** | **int32** |  | 
**SourceCurrency** | Pointer to [**PaymentAuditCurrencyV3**](PaymentAuditCurrencyV3.md) |  | [optional] 
**PaymentCurrency** | Pointer to [**PaymentAuditCurrencyV3**](PaymentAuditCurrencyV3.md) |  | [optional] 
**Status** | **string** |  | 
**FundingStatus** | **string** |  | 

## Methods

### NewFxSummaryV3

`func NewFxSummaryV3(quoteId string, creationDateTime time.Time, rate float32, invertedRate float32, totalCost int32, totalPaymentAmount int32, status string, fundingStatus string, ) *FxSummaryV3`

NewFxSummaryV3 instantiates a new FxSummaryV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxSummaryV3WithDefaults

`func NewFxSummaryV3WithDefaults() *FxSummaryV3`

NewFxSummaryV3WithDefaults instantiates a new FxSummaryV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *FxSummaryV3) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *FxSummaryV3) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *FxSummaryV3) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetCreationDateTime

`func (o *FxSummaryV3) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *FxSummaryV3) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *FxSummaryV3) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.


### GetRate

`func (o *FxSummaryV3) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *FxSummaryV3) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *FxSummaryV3) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetInvertedRate

`func (o *FxSummaryV3) GetInvertedRate() float32`

GetInvertedRate returns the InvertedRate field if non-nil, zero value otherwise.

### GetInvertedRateOk

`func (o *FxSummaryV3) GetInvertedRateOk() (*float32, bool)`

GetInvertedRateOk returns a tuple with the InvertedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertedRate

`func (o *FxSummaryV3) SetInvertedRate(v float32)`

SetInvertedRate sets InvertedRate field to given value.


### GetTotalCost

`func (o *FxSummaryV3) GetTotalCost() int32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *FxSummaryV3) GetTotalCostOk() (*int32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *FxSummaryV3) SetTotalCost(v int32)`

SetTotalCost sets TotalCost field to given value.


### GetTotalPaymentAmount

`func (o *FxSummaryV3) GetTotalPaymentAmount() int32`

GetTotalPaymentAmount returns the TotalPaymentAmount field if non-nil, zero value otherwise.

### GetTotalPaymentAmountOk

`func (o *FxSummaryV3) GetTotalPaymentAmountOk() (*int32, bool)`

GetTotalPaymentAmountOk returns a tuple with the TotalPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaymentAmount

`func (o *FxSummaryV3) SetTotalPaymentAmount(v int32)`

SetTotalPaymentAmount sets TotalPaymentAmount field to given value.


### GetSourceCurrency

`func (o *FxSummaryV3) GetSourceCurrency() PaymentAuditCurrencyV3`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *FxSummaryV3) GetSourceCurrencyOk() (*PaymentAuditCurrencyV3, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *FxSummaryV3) SetSourceCurrency(v PaymentAuditCurrencyV3)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *FxSummaryV3) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *FxSummaryV3) GetPaymentCurrency() PaymentAuditCurrencyV3`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *FxSummaryV3) GetPaymentCurrencyOk() (*PaymentAuditCurrencyV3, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *FxSummaryV3) SetPaymentCurrency(v PaymentAuditCurrencyV3)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *FxSummaryV3) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *FxSummaryV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FxSummaryV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FxSummaryV3) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFundingStatus

`func (o *FxSummaryV3) GetFundingStatus() string`

GetFundingStatus returns the FundingStatus field if non-nil, zero value otherwise.

### GetFundingStatusOk

`func (o *FxSummaryV3) GetFundingStatusOk() (*string, bool)`

GetFundingStatusOk returns a tuple with the FundingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingStatus

`func (o *FxSummaryV3) SetFundingStatus(v string)`

SetFundingStatus sets FundingStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


