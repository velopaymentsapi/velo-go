# CreatePaymentChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentChannelName** | Pointer to **string** |  | [optional] 
**Iban** | Pointer to **string** | Must match the regular expression &#x60;&#x60;&#x60;^[A-Za-z0-9]+$&#x60;&#x60;&#x60;. Either routing number and account number or only iban must be set | [optional] 
**AccountNumber** | Pointer to **string** | Either routing number and account number or only iban must be set | [optional] 
**RoutingNumber** | Pointer to **string** | Either routing number and account number or only iban must be set | [optional] 
**CountryCode** | **string** | Two character country code | 
**Currency** | **string** |  | 
**AccountName** | **string** |  | 

## Methods

### NewCreatePaymentChannel

`func NewCreatePaymentChannel(countryCode string, currency string, accountName string, ) *CreatePaymentChannel`

NewCreatePaymentChannel instantiates a new CreatePaymentChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentChannelWithDefaults

`func NewCreatePaymentChannelWithDefaults() *CreatePaymentChannel`

NewCreatePaymentChannelWithDefaults instantiates a new CreatePaymentChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentChannelName

`func (o *CreatePaymentChannel) GetPaymentChannelName() string`

GetPaymentChannelName returns the PaymentChannelName field if non-nil, zero value otherwise.

### GetPaymentChannelNameOk

`func (o *CreatePaymentChannel) GetPaymentChannelNameOk() (*string, bool)`

GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelName

`func (o *CreatePaymentChannel) SetPaymentChannelName(v string)`

SetPaymentChannelName sets PaymentChannelName field to given value.

### HasPaymentChannelName

`func (o *CreatePaymentChannel) HasPaymentChannelName() bool`

HasPaymentChannelName returns a boolean if a field has been set.

### GetIban

`func (o *CreatePaymentChannel) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *CreatePaymentChannel) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *CreatePaymentChannel) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *CreatePaymentChannel) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetAccountNumber

`func (o *CreatePaymentChannel) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreatePaymentChannel) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreatePaymentChannel) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CreatePaymentChannel) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *CreatePaymentChannel) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *CreatePaymentChannel) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *CreatePaymentChannel) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *CreatePaymentChannel) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetCountryCode

`func (o *CreatePaymentChannel) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreatePaymentChannel) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreatePaymentChannel) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCurrency

`func (o *CreatePaymentChannel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePaymentChannel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePaymentChannel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAccountName

`func (o *CreatePaymentChannel) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CreatePaymentChannel) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CreatePaymentChannel) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


