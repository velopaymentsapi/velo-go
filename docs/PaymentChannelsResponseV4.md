# PaymentChannelsResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentChannels** | Pointer to [**[]PaymentChannelResponseV4**](PaymentChannelResponseV4.md) |  | [optional] 
**PayorToPaymentChannelMappings** | Pointer to [**[]PayorToPaymentChannelMappingV4**](PayorToPaymentChannelMappingV4.md) |  | [optional] 

## Methods

### NewPaymentChannelsResponseV4

`func NewPaymentChannelsResponseV4() *PaymentChannelsResponseV4`

NewPaymentChannelsResponseV4 instantiates a new PaymentChannelsResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChannelsResponseV4WithDefaults

`func NewPaymentChannelsResponseV4WithDefaults() *PaymentChannelsResponseV4`

NewPaymentChannelsResponseV4WithDefaults instantiates a new PaymentChannelsResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentChannels

`func (o *PaymentChannelsResponseV4) GetPaymentChannels() []PaymentChannelResponseV4`

GetPaymentChannels returns the PaymentChannels field if non-nil, zero value otherwise.

### GetPaymentChannelsOk

`func (o *PaymentChannelsResponseV4) GetPaymentChannelsOk() (*[]PaymentChannelResponseV4, bool)`

GetPaymentChannelsOk returns a tuple with the PaymentChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannels

`func (o *PaymentChannelsResponseV4) SetPaymentChannels(v []PaymentChannelResponseV4)`

SetPaymentChannels sets PaymentChannels field to given value.

### HasPaymentChannels

`func (o *PaymentChannelsResponseV4) HasPaymentChannels() bool`

HasPaymentChannels returns a boolean if a field has been set.

### GetPayorToPaymentChannelMappings

`func (o *PaymentChannelsResponseV4) GetPayorToPaymentChannelMappings() []PayorToPaymentChannelMappingV4`

GetPayorToPaymentChannelMappings returns the PayorToPaymentChannelMappings field if non-nil, zero value otherwise.

### GetPayorToPaymentChannelMappingsOk

`func (o *PaymentChannelsResponseV4) GetPayorToPaymentChannelMappingsOk() (*[]PayorToPaymentChannelMappingV4, bool)`

GetPayorToPaymentChannelMappingsOk returns a tuple with the PayorToPaymentChannelMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorToPaymentChannelMappings

`func (o *PaymentChannelsResponseV4) SetPayorToPaymentChannelMappings(v []PayorToPaymentChannelMappingV4)`

SetPayorToPaymentChannelMappings sets PayorToPaymentChannelMappings field to given value.

### HasPayorToPaymentChannelMappings

`func (o *PaymentChannelsResponseV4) HasPayorToPaymentChannelMappings() bool`

HasPayorToPaymentChannelMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


