# PayoutPayor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** | The id of the payor. | 
**PayorName** | **string** | The name of the payor. | 
**DbaName** | **string** | The alternate name of the payor. | 
**Principal** | **string** | Email address if principal is a user or ID if application. | 
**PrincipalId** | **string** | The id of the principal. | 

## Methods

### NewPayoutPayor

`func NewPayoutPayor(payorId string, payorName string, dbaName string, principal string, principalId string, ) *PayoutPayor`

NewPayoutPayor instantiates a new PayoutPayor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutPayorWithDefaults

`func NewPayoutPayorWithDefaults() *PayoutPayor`

NewPayoutPayorWithDefaults instantiates a new PayoutPayor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *PayoutPayor) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayoutPayor) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayoutPayor) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetPayorName

`func (o *PayoutPayor) GetPayorName() string`

GetPayorName returns the PayorName field if non-nil, zero value otherwise.

### GetPayorNameOk

`func (o *PayoutPayor) GetPayorNameOk() (*string, bool)`

GetPayorNameOk returns a tuple with the PayorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorName

`func (o *PayoutPayor) SetPayorName(v string)`

SetPayorName sets PayorName field to given value.


### GetDbaName

`func (o *PayoutPayor) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *PayoutPayor) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *PayoutPayor) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.


### GetPrincipal

`func (o *PayoutPayor) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PayoutPayor) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PayoutPayor) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetPrincipalId

`func (o *PayoutPayor) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *PayoutPayor) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *PayoutPayor) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


