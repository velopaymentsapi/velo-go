# PaymentChannelSummaryV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentChannelId** | **string** |  | 
**Name** | **string** | The payment channel name | 
**CountryCode** | **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Currency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Last4Digits** | Pointer to **string** | The last 4 digits of the account number or IBAN | [optional] 
**Enabled** | **bool** | Usually true. False if an associated payment is returned | 
**DisabledReason** | Pointer to **string** | Indicates why the payment channel was disabled | [optional] 
**DisabledReasonCode** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentChannelSummaryV4

`func NewPaymentChannelSummaryV4(paymentChannelId string, name string, countryCode string, currency string, enabled bool, ) *PaymentChannelSummaryV4`

NewPaymentChannelSummaryV4 instantiates a new PaymentChannelSummaryV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChannelSummaryV4WithDefaults

`func NewPaymentChannelSummaryV4WithDefaults() *PaymentChannelSummaryV4`

NewPaymentChannelSummaryV4WithDefaults instantiates a new PaymentChannelSummaryV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentChannelId

`func (o *PaymentChannelSummaryV4) GetPaymentChannelId() string`

GetPaymentChannelId returns the PaymentChannelId field if non-nil, zero value otherwise.

### GetPaymentChannelIdOk

`func (o *PaymentChannelSummaryV4) GetPaymentChannelIdOk() (*string, bool)`

GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelId

`func (o *PaymentChannelSummaryV4) SetPaymentChannelId(v string)`

SetPaymentChannelId sets PaymentChannelId field to given value.


### GetName

`func (o *PaymentChannelSummaryV4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentChannelSummaryV4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentChannelSummaryV4) SetName(v string)`

SetName sets Name field to given value.


### GetCountryCode

`func (o *PaymentChannelSummaryV4) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PaymentChannelSummaryV4) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PaymentChannelSummaryV4) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCurrency

`func (o *PaymentChannelSummaryV4) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentChannelSummaryV4) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentChannelSummaryV4) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetLast4Digits

`func (o *PaymentChannelSummaryV4) GetLast4Digits() string`

GetLast4Digits returns the Last4Digits field if non-nil, zero value otherwise.

### GetLast4DigitsOk

`func (o *PaymentChannelSummaryV4) GetLast4DigitsOk() (*string, bool)`

GetLast4DigitsOk returns a tuple with the Last4Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4Digits

`func (o *PaymentChannelSummaryV4) SetLast4Digits(v string)`

SetLast4Digits sets Last4Digits field to given value.

### HasLast4Digits

`func (o *PaymentChannelSummaryV4) HasLast4Digits() bool`

HasLast4Digits returns a boolean if a field has been set.

### GetEnabled

`func (o *PaymentChannelSummaryV4) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PaymentChannelSummaryV4) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PaymentChannelSummaryV4) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDisabledReason

`func (o *PaymentChannelSummaryV4) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *PaymentChannelSummaryV4) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *PaymentChannelSummaryV4) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.

### HasDisabledReason

`func (o *PaymentChannelSummaryV4) HasDisabledReason() bool`

HasDisabledReason returns a boolean if a field has been set.

### GetDisabledReasonCode

`func (o *PaymentChannelSummaryV4) GetDisabledReasonCode() string`

GetDisabledReasonCode returns the DisabledReasonCode field if non-nil, zero value otherwise.

### GetDisabledReasonCodeOk

`func (o *PaymentChannelSummaryV4) GetDisabledReasonCodeOk() (*string, bool)`

GetDisabledReasonCodeOk returns a tuple with the DisabledReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReasonCode

`func (o *PaymentChannelSummaryV4) SetDisabledReasonCode(v string)`

SetDisabledReasonCode sets DisabledReasonCode field to given value.

### HasDisabledReasonCode

`func (o *PaymentChannelSummaryV4) HasDisabledReasonCode() bool`

HasDisabledReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


