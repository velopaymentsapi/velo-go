# SupportedCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoCountryCode** | Pointer to **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**Currencies** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSupportedCountry

`func NewSupportedCountry() *SupportedCountry`

NewSupportedCountry instantiates a new SupportedCountry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedCountryWithDefaults

`func NewSupportedCountryWithDefaults() *SupportedCountry`

NewSupportedCountryWithDefaults instantiates a new SupportedCountry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsoCountryCode

`func (o *SupportedCountry) GetIsoCountryCode() string`

GetIsoCountryCode returns the IsoCountryCode field if non-nil, zero value otherwise.

### GetIsoCountryCodeOk

`func (o *SupportedCountry) GetIsoCountryCodeOk() (*string, bool)`

GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountryCode

`func (o *SupportedCountry) SetIsoCountryCode(v string)`

SetIsoCountryCode sets IsoCountryCode field to given value.

### HasIsoCountryCode

`func (o *SupportedCountry) HasIsoCountryCode() bool`

HasIsoCountryCode returns a boolean if a field has been set.

### GetCurrencies

`func (o *SupportedCountry) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *SupportedCountry) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *SupportedCountry) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *SupportedCountry) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


