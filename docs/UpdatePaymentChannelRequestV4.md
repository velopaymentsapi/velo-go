# UpdatePaymentChannelRequestV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingNumber** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**Iban** | Pointer to **string** | Must match the regular expression &#x60;&#x60;&#x60;^[A-Za-z0-9]+$&#x60;&#x60;&#x60;. | [optional] 
**PaymentChannelName** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 

## Methods

### NewUpdatePaymentChannelRequestV4

`func NewUpdatePaymentChannelRequestV4() *UpdatePaymentChannelRequestV4`

NewUpdatePaymentChannelRequestV4 instantiates a new UpdatePaymentChannelRequestV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentChannelRequestV4WithDefaults

`func NewUpdatePaymentChannelRequestV4WithDefaults() *UpdatePaymentChannelRequestV4`

NewUpdatePaymentChannelRequestV4WithDefaults instantiates a new UpdatePaymentChannelRequestV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingNumber

`func (o *UpdatePaymentChannelRequestV4) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *UpdatePaymentChannelRequestV4) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *UpdatePaymentChannelRequestV4) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *UpdatePaymentChannelRequestV4) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *UpdatePaymentChannelRequestV4) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *UpdatePaymentChannelRequestV4) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *UpdatePaymentChannelRequestV4) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *UpdatePaymentChannelRequestV4) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetIban

`func (o *UpdatePaymentChannelRequestV4) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *UpdatePaymentChannelRequestV4) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *UpdatePaymentChannelRequestV4) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *UpdatePaymentChannelRequestV4) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetPaymentChannelName

`func (o *UpdatePaymentChannelRequestV4) GetPaymentChannelName() string`

GetPaymentChannelName returns the PaymentChannelName field if non-nil, zero value otherwise.

### GetPaymentChannelNameOk

`func (o *UpdatePaymentChannelRequestV4) GetPaymentChannelNameOk() (*string, bool)`

GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelName

`func (o *UpdatePaymentChannelRequestV4) SetPaymentChannelName(v string)`

SetPaymentChannelName sets PaymentChannelName field to given value.

### HasPaymentChannelName

`func (o *UpdatePaymentChannelRequestV4) HasPaymentChannelName() bool`

HasPaymentChannelName returns a boolean if a field has been set.

### GetCountryCode

`func (o *UpdatePaymentChannelRequestV4) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UpdatePaymentChannelRequestV4) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UpdatePaymentChannelRequestV4) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UpdatePaymentChannelRequestV4) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetAccountName

`func (o *UpdatePaymentChannelRequestV4) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *UpdatePaymentChannelRequestV4) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *UpdatePaymentChannelRequestV4) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *UpdatePaymentChannelRequestV4) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdatePaymentChannelRequestV4) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdatePaymentChannelRequestV4) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdatePaymentChannelRequestV4) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdatePaymentChannelRequestV4) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


