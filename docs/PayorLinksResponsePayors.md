# PayorLinksResponsePayors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** |  | 
**PayorName** | **string** |  | 
**PrimaryContactEmail** | Pointer to **string** |  | [optional] 
**KycState** | Pointer to **string** | Current KYC state. One of the following values: FAILED_KYC, PASSED_KYC, REQUIRES_KYC | [optional] 

## Methods

### NewPayorLinksResponsePayors

`func NewPayorLinksResponsePayors(payorId string, payorName string, ) *PayorLinksResponsePayors`

NewPayorLinksResponsePayors instantiates a new PayorLinksResponsePayors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorLinksResponsePayorsWithDefaults

`func NewPayorLinksResponsePayorsWithDefaults() *PayorLinksResponsePayors`

NewPayorLinksResponsePayorsWithDefaults instantiates a new PayorLinksResponsePayors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *PayorLinksResponsePayors) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayorLinksResponsePayors) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayorLinksResponsePayors) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetPayorName

`func (o *PayorLinksResponsePayors) GetPayorName() string`

GetPayorName returns the PayorName field if non-nil, zero value otherwise.

### GetPayorNameOk

`func (o *PayorLinksResponsePayors) GetPayorNameOk() (*string, bool)`

GetPayorNameOk returns a tuple with the PayorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorName

`func (o *PayorLinksResponsePayors) SetPayorName(v string)`

SetPayorName sets PayorName field to given value.


### GetPrimaryContactEmail

`func (o *PayorLinksResponsePayors) GetPrimaryContactEmail() string`

GetPrimaryContactEmail returns the PrimaryContactEmail field if non-nil, zero value otherwise.

### GetPrimaryContactEmailOk

`func (o *PayorLinksResponsePayors) GetPrimaryContactEmailOk() (*string, bool)`

GetPrimaryContactEmailOk returns a tuple with the PrimaryContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactEmail

`func (o *PayorLinksResponsePayors) SetPrimaryContactEmail(v string)`

SetPrimaryContactEmail sets PrimaryContactEmail field to given value.

### HasPrimaryContactEmail

`func (o *PayorLinksResponsePayors) HasPrimaryContactEmail() bool`

HasPrimaryContactEmail returns a boolean if a field has been set.

### GetKycState

`func (o *PayorLinksResponsePayors) GetKycState() string`

GetKycState returns the KycState field if non-nil, zero value otherwise.

### GetKycStateOk

`func (o *PayorLinksResponsePayors) GetKycStateOk() (*string, bool)`

GetKycStateOk returns a tuple with the KycState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycState

`func (o *PayorLinksResponsePayors) SetKycState(v string)`

SetKycState sets KycState field to given value.

### HasKycState

`func (o *PayorLinksResponsePayors) HasKycState() bool`

HasKycState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


