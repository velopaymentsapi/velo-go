# CreatePaymentChannelV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentChannelName** | Pointer to **string** |  | [optional] 
**Iban** | Pointer to **string** | Must match the regular expression &#x60;&#x60;&#x60;^[A-Za-z0-9]+$&#x60;&#x60;&#x60;. | [optional] 
**AccountNumber** | Pointer to **string** | Either routing number and account number or only iban must be set | [optional] 
**RoutingNumber** | Pointer to **string** | Either routing number and account number or only iban must be set | [optional] 
**CountryCode** | **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Currency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**AccountName** | **string** |  | 

## Methods

### NewCreatePaymentChannelV3

`func NewCreatePaymentChannelV3(countryCode string, currency string, accountName string, ) *CreatePaymentChannelV3`

NewCreatePaymentChannelV3 instantiates a new CreatePaymentChannelV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentChannelV3WithDefaults

`func NewCreatePaymentChannelV3WithDefaults() *CreatePaymentChannelV3`

NewCreatePaymentChannelV3WithDefaults instantiates a new CreatePaymentChannelV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentChannelName

`func (o *CreatePaymentChannelV3) GetPaymentChannelName() string`

GetPaymentChannelName returns the PaymentChannelName field if non-nil, zero value otherwise.

### GetPaymentChannelNameOk

`func (o *CreatePaymentChannelV3) GetPaymentChannelNameOk() (*string, bool)`

GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelName

`func (o *CreatePaymentChannelV3) SetPaymentChannelName(v string)`

SetPaymentChannelName sets PaymentChannelName field to given value.

### HasPaymentChannelName

`func (o *CreatePaymentChannelV3) HasPaymentChannelName() bool`

HasPaymentChannelName returns a boolean if a field has been set.

### GetIban

`func (o *CreatePaymentChannelV3) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *CreatePaymentChannelV3) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *CreatePaymentChannelV3) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *CreatePaymentChannelV3) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetAccountNumber

`func (o *CreatePaymentChannelV3) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreatePaymentChannelV3) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreatePaymentChannelV3) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CreatePaymentChannelV3) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *CreatePaymentChannelV3) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *CreatePaymentChannelV3) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *CreatePaymentChannelV3) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *CreatePaymentChannelV3) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetCountryCode

`func (o *CreatePaymentChannelV3) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreatePaymentChannelV3) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreatePaymentChannelV3) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCurrency

`func (o *CreatePaymentChannelV3) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePaymentChannelV3) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePaymentChannelV3) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAccountName

`func (o *CreatePaymentChannelV3) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CreatePaymentChannelV3) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CreatePaymentChannelV3) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


