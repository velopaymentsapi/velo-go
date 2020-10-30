# PaymentChannelRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | Pointer to [**[]PaymentChannelCountry**](PaymentChannelCountry.md) |  | [optional] 

## Methods

### NewPaymentChannelRulesResponse

`func NewPaymentChannelRulesResponse() *PaymentChannelRulesResponse`

NewPaymentChannelRulesResponse instantiates a new PaymentChannelRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChannelRulesResponseWithDefaults

`func NewPaymentChannelRulesResponseWithDefaults() *PaymentChannelRulesResponse`

NewPaymentChannelRulesResponseWithDefaults instantiates a new PaymentChannelRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *PaymentChannelRulesResponse) GetBank() []PaymentChannelCountry`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *PaymentChannelRulesResponse) GetBankOk() (*[]PaymentChannelCountry, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *PaymentChannelRulesResponse) SetBank(v []PaymentChannelCountry)`

SetBank sets Bank field to given value.

### HasBank

`func (o *PaymentChannelRulesResponse) HasBank() bool`

HasBank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


