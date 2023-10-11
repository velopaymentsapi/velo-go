# InstructPayoutRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FxRateDegredationThresholdPercentage** | Pointer to **float32** | Halt instruction if the FX rates have become worse since the last quote | [optional] 

## Methods

### NewInstructPayoutRequestV3

`func NewInstructPayoutRequestV3() *InstructPayoutRequestV3`

NewInstructPayoutRequestV3 instantiates a new InstructPayoutRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstructPayoutRequestV3WithDefaults

`func NewInstructPayoutRequestV3WithDefaults() *InstructPayoutRequestV3`

NewInstructPayoutRequestV3WithDefaults instantiates a new InstructPayoutRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFxRateDegredationThresholdPercentage

`func (o *InstructPayoutRequestV3) GetFxRateDegredationThresholdPercentage() float32`

GetFxRateDegredationThresholdPercentage returns the FxRateDegredationThresholdPercentage field if non-nil, zero value otherwise.

### GetFxRateDegredationThresholdPercentageOk

`func (o *InstructPayoutRequestV3) GetFxRateDegredationThresholdPercentageOk() (*float32, bool)`

GetFxRateDegredationThresholdPercentageOk returns a tuple with the FxRateDegredationThresholdPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxRateDegredationThresholdPercentage

`func (o *InstructPayoutRequestV3) SetFxRateDegredationThresholdPercentage(v float32)`

SetFxRateDegredationThresholdPercentage sets FxRateDegredationThresholdPercentage field to given value.

### HasFxRateDegredationThresholdPercentage

`func (o *InstructPayoutRequestV3) HasFxRateDegredationThresholdPercentage() bool`

HasFxRateDegredationThresholdPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


