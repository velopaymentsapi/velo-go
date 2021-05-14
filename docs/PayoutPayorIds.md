# PayoutPayorIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubmittingPayorId** | **string** | The ID of the Payor that is submitting the payout | 
**PayoutFromPayorId** | **string** | The ID of the Payor providing the source account for the payout | 
**PayoutToPayorId** | **string** | The ID of the Payor that owns the Payee (on behalf of) | 

## Methods

### NewPayoutPayorIds

`func NewPayoutPayorIds(submittingPayorId string, payoutFromPayorId string, payoutToPayorId string, ) *PayoutPayorIds`

NewPayoutPayorIds instantiates a new PayoutPayorIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutPayorIdsWithDefaults

`func NewPayoutPayorIdsWithDefaults() *PayoutPayorIds`

NewPayoutPayorIdsWithDefaults instantiates a new PayoutPayorIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubmittingPayorId

`func (o *PayoutPayorIds) GetSubmittingPayorId() string`

GetSubmittingPayorId returns the SubmittingPayorId field if non-nil, zero value otherwise.

### GetSubmittingPayorIdOk

`func (o *PayoutPayorIds) GetSubmittingPayorIdOk() (*string, bool)`

GetSubmittingPayorIdOk returns a tuple with the SubmittingPayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittingPayorId

`func (o *PayoutPayorIds) SetSubmittingPayorId(v string)`

SetSubmittingPayorId sets SubmittingPayorId field to given value.


### GetPayoutFromPayorId

`func (o *PayoutPayorIds) GetPayoutFromPayorId() string`

GetPayoutFromPayorId returns the PayoutFromPayorId field if non-nil, zero value otherwise.

### GetPayoutFromPayorIdOk

`func (o *PayoutPayorIds) GetPayoutFromPayorIdOk() (*string, bool)`

GetPayoutFromPayorIdOk returns a tuple with the PayoutFromPayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutFromPayorId

`func (o *PayoutPayorIds) SetPayoutFromPayorId(v string)`

SetPayoutFromPayorId sets PayoutFromPayorId field to given value.


### GetPayoutToPayorId

`func (o *PayoutPayorIds) GetPayoutToPayorId() string`

GetPayoutToPayorId returns the PayoutToPayorId field if non-nil, zero value otherwise.

### GetPayoutToPayorIdOk

`func (o *PayoutPayorIds) GetPayoutToPayorIdOk() (*string, bool)`

GetPayoutToPayorIdOk returns a tuple with the PayoutToPayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutToPayorId

`func (o *PayoutPayorIds) SetPayoutToPayorId(v string)`

SetPayoutToPayorId sets PayoutToPayorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


