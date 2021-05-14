# PayoutPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | **string** | Email address if principal is a user or ID if application. | 
**PrincipalId** | **string** | The id of the principal. | 

## Methods

### NewPayoutPrincipal

`func NewPayoutPrincipal(principal string, principalId string, ) *PayoutPrincipal`

NewPayoutPrincipal instantiates a new PayoutPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutPrincipalWithDefaults

`func NewPayoutPrincipalWithDefaults() *PayoutPrincipal`

NewPayoutPrincipalWithDefaults instantiates a new PayoutPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *PayoutPrincipal) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PayoutPrincipal) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PayoutPrincipal) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetPrincipalId

`func (o *PayoutPrincipal) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *PayoutPrincipal) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *PayoutPrincipal) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


