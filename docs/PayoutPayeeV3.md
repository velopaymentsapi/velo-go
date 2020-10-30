# PayoutPayeeV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | **string** |  | 
**Individual** | Pointer to [**PayoutIndividualV3**](PayoutIndividualV3.md) |  | [optional] 
**Company** | Pointer to [**PayoutCompanyV3**](PayoutCompanyV3.md) |  | [optional] 

## Methods

### NewPayoutPayeeV3

`func NewPayoutPayeeV3(payeeId string, ) *PayoutPayeeV3`

NewPayoutPayeeV3 instantiates a new PayoutPayeeV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutPayeeV3WithDefaults

`func NewPayoutPayeeV3WithDefaults() *PayoutPayeeV3`

NewPayoutPayeeV3WithDefaults instantiates a new PayoutPayeeV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *PayoutPayeeV3) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayoutPayeeV3) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayoutPayeeV3) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetIndividual

`func (o *PayoutPayeeV3) GetIndividual() PayoutIndividualV3`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *PayoutPayeeV3) GetIndividualOk() (*PayoutIndividualV3, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *PayoutPayeeV3) SetIndividual(v PayoutIndividualV3)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *PayoutPayeeV3) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *PayoutPayeeV3) GetCompany() PayoutCompanyV3`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PayoutPayeeV3) GetCompanyOk() (*PayoutCompanyV3, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PayoutPayeeV3) SetCompany(v PayoutCompanyV3)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PayoutPayeeV3) HasCompany() bool`

HasCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


