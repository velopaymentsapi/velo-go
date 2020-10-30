# QuoteFxSummaryV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | **float32** |  | 
**InvertedRate** | Pointer to **float32** |  | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) |  | 
**ExpiryTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**QuoteId** | **string** |  | 
**TotalSourceAmount** | **int32** |  | 
**TotalPaymentAmount** | **int32** |  | 
**SourceCurrency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**PaymentCurrency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**FundingStatus** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewQuoteFxSummaryV3

`func NewQuoteFxSummaryV3(rate float32, creationTime time.Time, quoteId string, totalSourceAmount int32, totalPaymentAmount int32, sourceCurrency string, paymentCurrency string, fundingStatus string, status string, ) *QuoteFxSummaryV3`

NewQuoteFxSummaryV3 instantiates a new QuoteFxSummaryV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteFxSummaryV3WithDefaults

`func NewQuoteFxSummaryV3WithDefaults() *QuoteFxSummaryV3`

NewQuoteFxSummaryV3WithDefaults instantiates a new QuoteFxSummaryV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *QuoteFxSummaryV3) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QuoteFxSummaryV3) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QuoteFxSummaryV3) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetInvertedRate

`func (o *QuoteFxSummaryV3) GetInvertedRate() float32`

GetInvertedRate returns the InvertedRate field if non-nil, zero value otherwise.

### GetInvertedRateOk

`func (o *QuoteFxSummaryV3) GetInvertedRateOk() (*float32, bool)`

GetInvertedRateOk returns a tuple with the InvertedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertedRate

`func (o *QuoteFxSummaryV3) SetInvertedRate(v float32)`

SetInvertedRate sets InvertedRate field to given value.

### HasInvertedRate

`func (o *QuoteFxSummaryV3) HasInvertedRate() bool`

HasInvertedRate returns a boolean if a field has been set.

### GetCreationTime

`func (o *QuoteFxSummaryV3) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *QuoteFxSummaryV3) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *QuoteFxSummaryV3) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetExpiryTime

`func (o *QuoteFxSummaryV3) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *QuoteFxSummaryV3) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *QuoteFxSummaryV3) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *QuoteFxSummaryV3) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetQuoteId

`func (o *QuoteFxSummaryV3) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QuoteFxSummaryV3) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QuoteFxSummaryV3) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetTotalSourceAmount

`func (o *QuoteFxSummaryV3) GetTotalSourceAmount() int32`

GetTotalSourceAmount returns the TotalSourceAmount field if non-nil, zero value otherwise.

### GetTotalSourceAmountOk

`func (o *QuoteFxSummaryV3) GetTotalSourceAmountOk() (*int32, bool)`

GetTotalSourceAmountOk returns a tuple with the TotalSourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSourceAmount

`func (o *QuoteFxSummaryV3) SetTotalSourceAmount(v int32)`

SetTotalSourceAmount sets TotalSourceAmount field to given value.


### GetTotalPaymentAmount

`func (o *QuoteFxSummaryV3) GetTotalPaymentAmount() int32`

GetTotalPaymentAmount returns the TotalPaymentAmount field if non-nil, zero value otherwise.

### GetTotalPaymentAmountOk

`func (o *QuoteFxSummaryV3) GetTotalPaymentAmountOk() (*int32, bool)`

GetTotalPaymentAmountOk returns a tuple with the TotalPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaymentAmount

`func (o *QuoteFxSummaryV3) SetTotalPaymentAmount(v int32)`

SetTotalPaymentAmount sets TotalPaymentAmount field to given value.


### GetSourceCurrency

`func (o *QuoteFxSummaryV3) GetSourceCurrency() string`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *QuoteFxSummaryV3) GetSourceCurrencyOk() (*string, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *QuoteFxSummaryV3) SetSourceCurrency(v string)`

SetSourceCurrency sets SourceCurrency field to given value.


### GetPaymentCurrency

`func (o *QuoteFxSummaryV3) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *QuoteFxSummaryV3) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *QuoteFxSummaryV3) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.


### GetFundingStatus

`func (o *QuoteFxSummaryV3) GetFundingStatus() string`

GetFundingStatus returns the FundingStatus field if non-nil, zero value otherwise.

### GetFundingStatusOk

`func (o *QuoteFxSummaryV3) GetFundingStatusOk() (*string, bool)`

GetFundingStatusOk returns a tuple with the FundingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingStatus

`func (o *QuoteFxSummaryV3) SetFundingStatus(v string)`

SetFundingStatus sets FundingStatus field to given value.


### GetStatus

`func (o *QuoteFxSummaryV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuoteFxSummaryV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuoteFxSummaryV3) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


