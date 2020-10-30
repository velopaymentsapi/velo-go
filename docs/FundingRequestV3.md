# FundingRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundingAccountId** | **string** | The funding account id | 
**Amount** | **int64** | Amount to fund in minor units | 

## Methods

### NewFundingRequestV3

`func NewFundingRequestV3(fundingAccountId string, amount int64, ) *FundingRequestV3`

NewFundingRequestV3 instantiates a new FundingRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingRequestV3WithDefaults

`func NewFundingRequestV3WithDefaults() *FundingRequestV3`

NewFundingRequestV3WithDefaults instantiates a new FundingRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundingAccountId

`func (o *FundingRequestV3) GetFundingAccountId() string`

GetFundingAccountId returns the FundingAccountId field if non-nil, zero value otherwise.

### GetFundingAccountIdOk

`func (o *FundingRequestV3) GetFundingAccountIdOk() (*string, bool)`

GetFundingAccountIdOk returns a tuple with the FundingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountId

`func (o *FundingRequestV3) SetFundingAccountId(v string)`

SetFundingAccountId sets FundingAccountId field to given value.


### GetAmount

`func (o *FundingRequestV3) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FundingRequestV3) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FundingRequestV3) SetAmount(v int64)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


