# SetNotificationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumBalance** | **int64** | Amount to set as minimum balance in minor units | 

## Methods

### NewSetNotificationsRequest

`func NewSetNotificationsRequest(minimumBalance int64, ) *SetNotificationsRequest`

NewSetNotificationsRequest instantiates a new SetNotificationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNotificationsRequestWithDefaults

`func NewSetNotificationsRequestWithDefaults() *SetNotificationsRequest`

NewSetNotificationsRequestWithDefaults instantiates a new SetNotificationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumBalance

`func (o *SetNotificationsRequest) GetMinimumBalance() int64`

GetMinimumBalance returns the MinimumBalance field if non-nil, zero value otherwise.

### GetMinimumBalanceOk

`func (o *SetNotificationsRequest) GetMinimumBalanceOk() (*int64, bool)`

GetMinimumBalanceOk returns a tuple with the MinimumBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBalance

`func (o *SetNotificationsRequest) SetMinimumBalance(v int64)`

SetMinimumBalance sets MinimumBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


