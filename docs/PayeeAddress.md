# PayeeAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | **string** |  | 
**Line2** | Pointer to **NullableString** |  | [optional] 
**Line3** | Pointer to **NullableString** |  | [optional] 
**Line4** | Pointer to **NullableString** |  | [optional] 
**City** | **string** |  | 
**CountyOrProvince** | Pointer to **NullableString** |  | [optional] 
**ZipOrPostcode** | Pointer to **NullableString** |  | [optional] 
**Country** | **string** |  | 

## Methods

### NewPayeeAddress

`func NewPayeeAddress(line1 string, city string, country string, ) *PayeeAddress`

NewPayeeAddress instantiates a new PayeeAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeAddressWithDefaults

`func NewPayeeAddressWithDefaults() *PayeeAddress`

NewPayeeAddressWithDefaults instantiates a new PayeeAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *PayeeAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *PayeeAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *PayeeAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *PayeeAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *PayeeAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *PayeeAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *PayeeAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *PayeeAddress) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *PayeeAddress) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *PayeeAddress) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *PayeeAddress) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *PayeeAddress) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *PayeeAddress) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### SetLine3Nil

`func (o *PayeeAddress) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *PayeeAddress) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
### GetLine4

`func (o *PayeeAddress) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *PayeeAddress) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *PayeeAddress) SetLine4(v string)`

SetLine4 sets Line4 field to given value.

### HasLine4

`func (o *PayeeAddress) HasLine4() bool`

HasLine4 returns a boolean if a field has been set.

### SetLine4Nil

`func (o *PayeeAddress) SetLine4Nil(b bool)`

 SetLine4Nil sets the value for Line4 to be an explicit nil

### UnsetLine4
`func (o *PayeeAddress) UnsetLine4()`

UnsetLine4 ensures that no value is present for Line4, not even an explicit nil
### GetCity

`func (o *PayeeAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PayeeAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PayeeAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetCountyOrProvince

`func (o *PayeeAddress) GetCountyOrProvince() string`

GetCountyOrProvince returns the CountyOrProvince field if non-nil, zero value otherwise.

### GetCountyOrProvinceOk

`func (o *PayeeAddress) GetCountyOrProvinceOk() (*string, bool)`

GetCountyOrProvinceOk returns a tuple with the CountyOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyOrProvince

`func (o *PayeeAddress) SetCountyOrProvince(v string)`

SetCountyOrProvince sets CountyOrProvince field to given value.

### HasCountyOrProvince

`func (o *PayeeAddress) HasCountyOrProvince() bool`

HasCountyOrProvince returns a boolean if a field has been set.

### SetCountyOrProvinceNil

`func (o *PayeeAddress) SetCountyOrProvinceNil(b bool)`

 SetCountyOrProvinceNil sets the value for CountyOrProvince to be an explicit nil

### UnsetCountyOrProvince
`func (o *PayeeAddress) UnsetCountyOrProvince()`

UnsetCountyOrProvince ensures that no value is present for CountyOrProvince, not even an explicit nil
### GetZipOrPostcode

`func (o *PayeeAddress) GetZipOrPostcode() string`

GetZipOrPostcode returns the ZipOrPostcode field if non-nil, zero value otherwise.

### GetZipOrPostcodeOk

`func (o *PayeeAddress) GetZipOrPostcodeOk() (*string, bool)`

GetZipOrPostcodeOk returns a tuple with the ZipOrPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipOrPostcode

`func (o *PayeeAddress) SetZipOrPostcode(v string)`

SetZipOrPostcode sets ZipOrPostcode field to given value.

### HasZipOrPostcode

`func (o *PayeeAddress) HasZipOrPostcode() bool`

HasZipOrPostcode returns a boolean if a field has been set.

### SetZipOrPostcodeNil

`func (o *PayeeAddress) SetZipOrPostcodeNil(b bool)`

 SetZipOrPostcodeNil sets the value for ZipOrPostcode to be an explicit nil

### UnsetZipOrPostcode
`func (o *PayeeAddress) UnsetZipOrPostcode()`

UnsetZipOrPostcode ensures that no value is present for ZipOrPostcode, not even an explicit nil
### GetCountry

`func (o *PayeeAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PayeeAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PayeeAddress) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


