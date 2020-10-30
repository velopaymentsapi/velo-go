# SupportedCountryV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoCountryCode** | Pointer to **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**Currencies** | Pointer to **[]string** |  | [optional] 
**Regions** | Pointer to [**[]RegionV2**](RegionV2.md) |  | [optional] 

## Methods

### NewSupportedCountryV2

`func NewSupportedCountryV2() *SupportedCountryV2`

NewSupportedCountryV2 instantiates a new SupportedCountryV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedCountryV2WithDefaults

`func NewSupportedCountryV2WithDefaults() *SupportedCountryV2`

NewSupportedCountryV2WithDefaults instantiates a new SupportedCountryV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsoCountryCode

`func (o *SupportedCountryV2) GetIsoCountryCode() string`

GetIsoCountryCode returns the IsoCountryCode field if non-nil, zero value otherwise.

### GetIsoCountryCodeOk

`func (o *SupportedCountryV2) GetIsoCountryCodeOk() (*string, bool)`

GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountryCode

`func (o *SupportedCountryV2) SetIsoCountryCode(v string)`

SetIsoCountryCode sets IsoCountryCode field to given value.

### HasIsoCountryCode

`func (o *SupportedCountryV2) HasIsoCountryCode() bool`

HasIsoCountryCode returns a boolean if a field has been set.

### GetCurrencies

`func (o *SupportedCountryV2) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *SupportedCountryV2) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *SupportedCountryV2) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *SupportedCountryV2) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetRegions

`func (o *SupportedCountryV2) GetRegions() []RegionV2`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SupportedCountryV2) GetRegionsOk() (*[]RegionV2, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SupportedCountryV2) SetRegions(v []RegionV2)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SupportedCountryV2) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


