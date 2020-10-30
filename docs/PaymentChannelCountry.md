# PaymentChannelCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoCountryCode** | **string** | The ISO code for the country | 
**Rules** | [**[]PaymentChannelRule**](PaymentChannelRule.md) | The rules for the given country | 

## Methods

### NewPaymentChannelCountry

`func NewPaymentChannelCountry(isoCountryCode string, rules []PaymentChannelRule, ) *PaymentChannelCountry`

NewPaymentChannelCountry instantiates a new PaymentChannelCountry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChannelCountryWithDefaults

`func NewPaymentChannelCountryWithDefaults() *PaymentChannelCountry`

NewPaymentChannelCountryWithDefaults instantiates a new PaymentChannelCountry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsoCountryCode

`func (o *PaymentChannelCountry) GetIsoCountryCode() string`

GetIsoCountryCode returns the IsoCountryCode field if non-nil, zero value otherwise.

### GetIsoCountryCodeOk

`func (o *PaymentChannelCountry) GetIsoCountryCodeOk() (*string, bool)`

GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountryCode

`func (o *PaymentChannelCountry) SetIsoCountryCode(v string)`

SetIsoCountryCode sets IsoCountryCode field to given value.


### GetRules

`func (o *PaymentChannelCountry) GetRules() []PaymentChannelRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PaymentChannelCountry) GetRulesOk() (*[]PaymentChannelRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PaymentChannelCountry) SetRules(v []PaymentChannelRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


