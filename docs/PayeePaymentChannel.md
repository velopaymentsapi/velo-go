# PayeePaymentChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentChannelName** | Pointer to **string** |  | [optional] 
**Iban** | Pointer to **string** | Must match the regular expression &#x60;&#x60;&#x60;^[A-Za-z0-9]+$&#x60;&#x60;&#x60;. | [optional] 
**AccountNumber** | **string** |  | 
**RoutingNumber** | **string** |  | 
**CountryCode** | **string** | Country Code must be a valid 2 letter ISO 3166-1 country code. | 
**Currency** | **string** |  | 
**AccountName** | **string** |  | 

## Methods

### NewPayeePaymentChannel

`func NewPayeePaymentChannel(accountNumber string, routingNumber string, countryCode string, currency string, accountName string, ) *PayeePaymentChannel`

NewPayeePaymentChannel instantiates a new PayeePaymentChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeePaymentChannelWithDefaults

`func NewPayeePaymentChannelWithDefaults() *PayeePaymentChannel`

NewPayeePaymentChannelWithDefaults instantiates a new PayeePaymentChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentChannelName

`func (o *PayeePaymentChannel) GetPaymentChannelName() string`

GetPaymentChannelName returns the PaymentChannelName field if non-nil, zero value otherwise.

### GetPaymentChannelNameOk

`func (o *PayeePaymentChannel) GetPaymentChannelNameOk() (*string, bool)`

GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelName

`func (o *PayeePaymentChannel) SetPaymentChannelName(v string)`

SetPaymentChannelName sets PaymentChannelName field to given value.

### HasPaymentChannelName

`func (o *PayeePaymentChannel) HasPaymentChannelName() bool`

HasPaymentChannelName returns a boolean if a field has been set.

### GetIban

`func (o *PayeePaymentChannel) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PayeePaymentChannel) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PayeePaymentChannel) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PayeePaymentChannel) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PayeePaymentChannel) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PayeePaymentChannel) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PayeePaymentChannel) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetRoutingNumber

`func (o *PayeePaymentChannel) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PayeePaymentChannel) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PayeePaymentChannel) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetCountryCode

`func (o *PayeePaymentChannel) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PayeePaymentChannel) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PayeePaymentChannel) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCurrency

`func (o *PayeePaymentChannel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayeePaymentChannel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayeePaymentChannel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAccountName

`func (o *PayeePaymentChannel) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PayeePaymentChannel) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PayeePaymentChannel) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


