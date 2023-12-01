# PaymentChannelResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PayeeId** | Pointer to **string** |  | [optional] 
**PaymentChannelName** | **string** |  | 
**AccountName** | **string** |  | 
**ChannelType** | **string** | Payment channel type. One of the following values: CHANNEL_BANK | 
**CountryCode** | **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**RoutingNumber** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**Iban** | Pointer to **string** | Must match the regular expression &#x60;&#x60;&#x60;^[A-Za-z0-9]+$&#x60;&#x60;&#x60;. | [optional] 
**Currency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**PayorIds** | Pointer to **[]string** |  | [optional] 
**PayeeName** | Pointer to **string** |  | [optional] 
**BankName** | Pointer to **string** |  | [optional] 
**BankSwiftBic** | Pointer to **string** |  | [optional] 
**BankAddress** | Pointer to [**AddressV4**](AddressV4.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DisabledReasonCode** | Pointer to **string** |  | [optional] 
**DisabledReason** | Pointer to **string** |  | [optional] 
**Payable** | Pointer to **bool** | Whether this payment channel is payable | [optional] 

## Methods

### NewPaymentChannelResponseV4

`func NewPaymentChannelResponseV4(id string, paymentChannelName string, accountName string, channelType string, countryCode string, currency string, ) *PaymentChannelResponseV4`

NewPaymentChannelResponseV4 instantiates a new PaymentChannelResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChannelResponseV4WithDefaults

`func NewPaymentChannelResponseV4WithDefaults() *PaymentChannelResponseV4`

NewPaymentChannelResponseV4WithDefaults instantiates a new PaymentChannelResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentChannelResponseV4) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentChannelResponseV4) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentChannelResponseV4) SetId(v string)`

SetId sets Id field to given value.


### GetPayeeId

`func (o *PaymentChannelResponseV4) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PaymentChannelResponseV4) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PaymentChannelResponseV4) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *PaymentChannelResponseV4) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPaymentChannelName

`func (o *PaymentChannelResponseV4) GetPaymentChannelName() string`

GetPaymentChannelName returns the PaymentChannelName field if non-nil, zero value otherwise.

### GetPaymentChannelNameOk

`func (o *PaymentChannelResponseV4) GetPaymentChannelNameOk() (*string, bool)`

GetPaymentChannelNameOk returns a tuple with the PaymentChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelName

`func (o *PaymentChannelResponseV4) SetPaymentChannelName(v string)`

SetPaymentChannelName sets PaymentChannelName field to given value.


### GetAccountName

`func (o *PaymentChannelResponseV4) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentChannelResponseV4) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentChannelResponseV4) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetChannelType

`func (o *PaymentChannelResponseV4) GetChannelType() string`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *PaymentChannelResponseV4) GetChannelTypeOk() (*string, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *PaymentChannelResponseV4) SetChannelType(v string)`

SetChannelType sets ChannelType field to given value.


### GetCountryCode

`func (o *PaymentChannelResponseV4) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PaymentChannelResponseV4) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PaymentChannelResponseV4) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetRoutingNumber

`func (o *PaymentChannelResponseV4) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PaymentChannelResponseV4) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PaymentChannelResponseV4) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *PaymentChannelResponseV4) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PaymentChannelResponseV4) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentChannelResponseV4) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentChannelResponseV4) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PaymentChannelResponseV4) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetIban

`func (o *PaymentChannelResponseV4) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentChannelResponseV4) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentChannelResponseV4) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PaymentChannelResponseV4) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentChannelResponseV4) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentChannelResponseV4) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentChannelResponseV4) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPayorIds

`func (o *PaymentChannelResponseV4) GetPayorIds() []string`

GetPayorIds returns the PayorIds field if non-nil, zero value otherwise.

### GetPayorIdsOk

`func (o *PaymentChannelResponseV4) GetPayorIdsOk() (*[]string, bool)`

GetPayorIdsOk returns a tuple with the PayorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorIds

`func (o *PaymentChannelResponseV4) SetPayorIds(v []string)`

SetPayorIds sets PayorIds field to given value.

### HasPayorIds

`func (o *PaymentChannelResponseV4) HasPayorIds() bool`

HasPayorIds returns a boolean if a field has been set.

### GetPayeeName

`func (o *PaymentChannelResponseV4) GetPayeeName() string`

GetPayeeName returns the PayeeName field if non-nil, zero value otherwise.

### GetPayeeNameOk

`func (o *PaymentChannelResponseV4) GetPayeeNameOk() (*string, bool)`

GetPayeeNameOk returns a tuple with the PayeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeName

`func (o *PaymentChannelResponseV4) SetPayeeName(v string)`

SetPayeeName sets PayeeName field to given value.

### HasPayeeName

`func (o *PaymentChannelResponseV4) HasPayeeName() bool`

HasPayeeName returns a boolean if a field has been set.

### GetBankName

`func (o *PaymentChannelResponseV4) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PaymentChannelResponseV4) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PaymentChannelResponseV4) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PaymentChannelResponseV4) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBankSwiftBic

`func (o *PaymentChannelResponseV4) GetBankSwiftBic() string`

GetBankSwiftBic returns the BankSwiftBic field if non-nil, zero value otherwise.

### GetBankSwiftBicOk

`func (o *PaymentChannelResponseV4) GetBankSwiftBicOk() (*string, bool)`

GetBankSwiftBicOk returns a tuple with the BankSwiftBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankSwiftBic

`func (o *PaymentChannelResponseV4) SetBankSwiftBic(v string)`

SetBankSwiftBic sets BankSwiftBic field to given value.

### HasBankSwiftBic

`func (o *PaymentChannelResponseV4) HasBankSwiftBic() bool`

HasBankSwiftBic returns a boolean if a field has been set.

### GetBankAddress

`func (o *PaymentChannelResponseV4) GetBankAddress() AddressV4`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *PaymentChannelResponseV4) GetBankAddressOk() (*AddressV4, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *PaymentChannelResponseV4) SetBankAddress(v AddressV4)`

SetBankAddress sets BankAddress field to given value.

### HasBankAddress

`func (o *PaymentChannelResponseV4) HasBankAddress() bool`

HasBankAddress returns a boolean if a field has been set.

### GetDeleted

`func (o *PaymentChannelResponseV4) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PaymentChannelResponseV4) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PaymentChannelResponseV4) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PaymentChannelResponseV4) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEnabled

`func (o *PaymentChannelResponseV4) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PaymentChannelResponseV4) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PaymentChannelResponseV4) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PaymentChannelResponseV4) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDisabledReasonCode

`func (o *PaymentChannelResponseV4) GetDisabledReasonCode() string`

GetDisabledReasonCode returns the DisabledReasonCode field if non-nil, zero value otherwise.

### GetDisabledReasonCodeOk

`func (o *PaymentChannelResponseV4) GetDisabledReasonCodeOk() (*string, bool)`

GetDisabledReasonCodeOk returns a tuple with the DisabledReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReasonCode

`func (o *PaymentChannelResponseV4) SetDisabledReasonCode(v string)`

SetDisabledReasonCode sets DisabledReasonCode field to given value.

### HasDisabledReasonCode

`func (o *PaymentChannelResponseV4) HasDisabledReasonCode() bool`

HasDisabledReasonCode returns a boolean if a field has been set.

### GetDisabledReason

`func (o *PaymentChannelResponseV4) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *PaymentChannelResponseV4) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *PaymentChannelResponseV4) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.

### HasDisabledReason

`func (o *PaymentChannelResponseV4) HasDisabledReason() bool`

HasDisabledReason returns a boolean if a field has been set.

### GetPayable

`func (o *PaymentChannelResponseV4) GetPayable() bool`

GetPayable returns the Payable field if non-nil, zero value otherwise.

### GetPayableOk

`func (o *PaymentChannelResponseV4) GetPayableOk() (*bool, bool)`

GetPayableOk returns a tuple with the Payable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayable

`func (o *PaymentChannelResponseV4) SetPayable(v bool)`

SetPayable sets Payable field to given value.

### HasPayable

`func (o *PaymentChannelResponseV4) HasPayable() bool`

HasPayable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


