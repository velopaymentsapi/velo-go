# PayoutPrincipalV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | **string** | Email address if principal is a user or ID if application. | 
**PrincipalId** | **string** | The id of the principal. | 

## Methods

### NewPayoutPrincipalV4

`func NewPayoutPrincipalV4(principal string, principalId string, ) *PayoutPrincipalV4`

NewPayoutPrincipalV4 instantiates a new PayoutPrincipalV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutPrincipalV4WithDefaults

`func NewPayoutPrincipalV4WithDefaults() *PayoutPrincipalV4`

NewPayoutPrincipalV4WithDefaults instantiates a new PayoutPrincipalV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *PayoutPrincipalV4) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PayoutPrincipalV4) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PayoutPrincipalV4) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetPrincipalId

`func (o *PayoutPrincipalV4) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *PayoutPrincipalV4) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *PayoutPrincipalV4) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


