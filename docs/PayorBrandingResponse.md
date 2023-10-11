# PayorBrandingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorName** | **string** | The name of the payor | 
**LogoUrl** | **string** | &lt;p&gt;The URL to use for this payor’s logo&lt;/p&gt; &lt;p&gt;This will be an immutable system-generated URL&lt;/p&gt;  | 
**CollectiveAlias** | Pointer to **NullableString** | How the payor has chosen to refer to payees | [optional] 
**SupportContact** | Pointer to **NullableString** | The payor’s support contact address | [optional] 
**DbaName** | Pointer to **NullableString** | The payor’s &#39;Doing Business As&#39; name | [optional] 

## Methods

### NewPayorBrandingResponse

`func NewPayorBrandingResponse(payorName string, logoUrl string, ) *PayorBrandingResponse`

NewPayorBrandingResponse instantiates a new PayorBrandingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorBrandingResponseWithDefaults

`func NewPayorBrandingResponseWithDefaults() *PayorBrandingResponse`

NewPayorBrandingResponseWithDefaults instantiates a new PayorBrandingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorName

`func (o *PayorBrandingResponse) GetPayorName() string`

GetPayorName returns the PayorName field if non-nil, zero value otherwise.

### GetPayorNameOk

`func (o *PayorBrandingResponse) GetPayorNameOk() (*string, bool)`

GetPayorNameOk returns a tuple with the PayorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorName

`func (o *PayorBrandingResponse) SetPayorName(v string)`

SetPayorName sets PayorName field to given value.


### GetLogoUrl

`func (o *PayorBrandingResponse) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *PayorBrandingResponse) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *PayorBrandingResponse) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetCollectiveAlias

`func (o *PayorBrandingResponse) GetCollectiveAlias() string`

GetCollectiveAlias returns the CollectiveAlias field if non-nil, zero value otherwise.

### GetCollectiveAliasOk

`func (o *PayorBrandingResponse) GetCollectiveAliasOk() (*string, bool)`

GetCollectiveAliasOk returns a tuple with the CollectiveAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectiveAlias

`func (o *PayorBrandingResponse) SetCollectiveAlias(v string)`

SetCollectiveAlias sets CollectiveAlias field to given value.

### HasCollectiveAlias

`func (o *PayorBrandingResponse) HasCollectiveAlias() bool`

HasCollectiveAlias returns a boolean if a field has been set.

### SetCollectiveAliasNil

`func (o *PayorBrandingResponse) SetCollectiveAliasNil(b bool)`

 SetCollectiveAliasNil sets the value for CollectiveAlias to be an explicit nil

### UnsetCollectiveAlias
`func (o *PayorBrandingResponse) UnsetCollectiveAlias()`

UnsetCollectiveAlias ensures that no value is present for CollectiveAlias, not even an explicit nil
### GetSupportContact

`func (o *PayorBrandingResponse) GetSupportContact() string`

GetSupportContact returns the SupportContact field if non-nil, zero value otherwise.

### GetSupportContactOk

`func (o *PayorBrandingResponse) GetSupportContactOk() (*string, bool)`

GetSupportContactOk returns a tuple with the SupportContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportContact

`func (o *PayorBrandingResponse) SetSupportContact(v string)`

SetSupportContact sets SupportContact field to given value.

### HasSupportContact

`func (o *PayorBrandingResponse) HasSupportContact() bool`

HasSupportContact returns a boolean if a field has been set.

### SetSupportContactNil

`func (o *PayorBrandingResponse) SetSupportContactNil(b bool)`

 SetSupportContactNil sets the value for SupportContact to be an explicit nil

### UnsetSupportContact
`func (o *PayorBrandingResponse) UnsetSupportContact()`

UnsetSupportContact ensures that no value is present for SupportContact, not even an explicit nil
### GetDbaName

`func (o *PayorBrandingResponse) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *PayorBrandingResponse) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *PayorBrandingResponse) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *PayorBrandingResponse) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### SetDbaNameNil

`func (o *PayorBrandingResponse) SetDbaNameNil(b bool)`

 SetDbaNameNil sets the value for DbaName to be an explicit nil

### UnsetDbaName
`func (o *PayorBrandingResponse) UnsetDbaName()`

UnsetDbaName ensures that no value is present for DbaName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


