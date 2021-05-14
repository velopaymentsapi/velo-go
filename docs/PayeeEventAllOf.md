# PayeeEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | **string** | ID of this payee within the Velo platform | 
**Reasons** | Pointer to [**[]PayeeEventAllOfReasons**](PayeeEventAllOfReasons.md) | The reasons for the event notification. | [optional] 

## Methods

### NewPayeeEventAllOf

`func NewPayeeEventAllOf(payeeId string, ) *PayeeEventAllOf`

NewPayeeEventAllOf instantiates a new PayeeEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeEventAllOfWithDefaults

`func NewPayeeEventAllOfWithDefaults() *PayeeEventAllOf`

NewPayeeEventAllOfWithDefaults instantiates a new PayeeEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *PayeeEventAllOf) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeEventAllOf) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeEventAllOf) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetReasons

`func (o *PayeeEventAllOf) GetReasons() []PayeeEventAllOfReasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *PayeeEventAllOf) GetReasonsOk() (*[]PayeeEventAllOfReasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *PayeeEventAllOf) SetReasons(v []PayeeEventAllOfReasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *PayeeEventAllOf) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


