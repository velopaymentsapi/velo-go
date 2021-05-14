# DebitStatusChangedAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The new status of the debit. One of \&quot;PENDING\&quot; \&quot;PROCESSING\&quot; \&quot;REJECTED\&quot; \&quot;RELEASED\&quot; | 

## Methods

### NewDebitStatusChangedAllOf

`func NewDebitStatusChangedAllOf(status string, ) *DebitStatusChangedAllOf`

NewDebitStatusChangedAllOf instantiates a new DebitStatusChangedAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebitStatusChangedAllOfWithDefaults

`func NewDebitStatusChangedAllOfWithDefaults() *DebitStatusChangedAllOf`

NewDebitStatusChangedAllOfWithDefaults instantiates a new DebitStatusChangedAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DebitStatusChangedAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DebitStatusChangedAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DebitStatusChangedAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


