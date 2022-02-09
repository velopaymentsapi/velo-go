# InstructPayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FxRateDegredationThresholdPercentage** | Pointer to **float32** | Halt instruction if the FX rates have become worse since the last quote | [optional] 

## Methods

### NewInstructPayoutRequest

`func NewInstructPayoutRequest() *InstructPayoutRequest`

NewInstructPayoutRequest instantiates a new InstructPayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstructPayoutRequestWithDefaults

`func NewInstructPayoutRequestWithDefaults() *InstructPayoutRequest`

NewInstructPayoutRequestWithDefaults instantiates a new InstructPayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFxRateDegredationThresholdPercentage

`func (o *InstructPayoutRequest) GetFxRateDegredationThresholdPercentage() float32`

GetFxRateDegredationThresholdPercentage returns the FxRateDegredationThresholdPercentage field if non-nil, zero value otherwise.

### GetFxRateDegredationThresholdPercentageOk

`func (o *InstructPayoutRequest) GetFxRateDegredationThresholdPercentageOk() (*float32, bool)`

GetFxRateDegredationThresholdPercentageOk returns a tuple with the FxRateDegredationThresholdPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxRateDegredationThresholdPercentage

`func (o *InstructPayoutRequest) SetFxRateDegredationThresholdPercentage(v float32)`

SetFxRateDegredationThresholdPercentage sets FxRateDegredationThresholdPercentage field to given value.

### HasFxRateDegredationThresholdPercentage

`func (o *InstructPayoutRequest) HasFxRateDegredationThresholdPercentage() bool`

HasFxRateDegredationThresholdPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


