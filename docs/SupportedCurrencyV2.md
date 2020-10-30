# SupportedCurrencyV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**MaxPaymentAmount** | Pointer to **int32** | The max amount allowed in this currency | [optional] 

## Methods

### NewSupportedCurrencyV2

`func NewSupportedCurrencyV2() *SupportedCurrencyV2`

NewSupportedCurrencyV2 instantiates a new SupportedCurrencyV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedCurrencyV2WithDefaults

`func NewSupportedCurrencyV2WithDefaults() *SupportedCurrencyV2`

NewSupportedCurrencyV2WithDefaults instantiates a new SupportedCurrencyV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *SupportedCurrencyV2) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SupportedCurrencyV2) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SupportedCurrencyV2) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SupportedCurrencyV2) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMaxPaymentAmount

`func (o *SupportedCurrencyV2) GetMaxPaymentAmount() int32`

GetMaxPaymentAmount returns the MaxPaymentAmount field if non-nil, zero value otherwise.

### GetMaxPaymentAmountOk

`func (o *SupportedCurrencyV2) GetMaxPaymentAmountOk() (*int32, bool)`

GetMaxPaymentAmountOk returns a tuple with the MaxPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPaymentAmount

`func (o *SupportedCurrencyV2) SetMaxPaymentAmount(v int32)`

SetMaxPaymentAmount sets MaxPaymentAmount field to given value.

### HasMaxPaymentAmount

`func (o *SupportedCurrencyV2) HasMaxPaymentAmount() bool`

HasMaxPaymentAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


