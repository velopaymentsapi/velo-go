# TransmissionTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ACH** | **bool** | Whether the Payor is allowed to pay via ACH | 
**SAME_DAY_ACH** | **bool** | Whether the Payor is allowed to pay via same day ACH | 
**WIRE** | **bool** | Whether the Payor is allowed to pay via wire | 

## Methods

### NewTransmissionTypes

`func NewTransmissionTypes(aCH bool, sAMEDAYACH bool, wIRE bool, ) *TransmissionTypes`

NewTransmissionTypes instantiates a new TransmissionTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransmissionTypesWithDefaults

`func NewTransmissionTypesWithDefaults() *TransmissionTypes`

NewTransmissionTypesWithDefaults instantiates a new TransmissionTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetACH

`func (o *TransmissionTypes) GetACH() bool`

GetACH returns the ACH field if non-nil, zero value otherwise.

### GetACHOk

`func (o *TransmissionTypes) GetACHOk() (*bool, bool)`

GetACHOk returns a tuple with the ACH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACH

`func (o *TransmissionTypes) SetACH(v bool)`

SetACH sets ACH field to given value.


### GetSAME_DAY_ACH

`func (o *TransmissionTypes) GetSAME_DAY_ACH() bool`

GetSAME_DAY_ACH returns the SAME_DAY_ACH field if non-nil, zero value otherwise.

### GetSAME_DAY_ACHOk

`func (o *TransmissionTypes) GetSAME_DAY_ACHOk() (*bool, bool)`

GetSAME_DAY_ACHOk returns a tuple with the SAME_DAY_ACH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAME_DAY_ACH

`func (o *TransmissionTypes) SetSAME_DAY_ACH(v bool)`

SetSAME_DAY_ACH sets SAME_DAY_ACH field to given value.


### GetWIRE

`func (o *TransmissionTypes) GetWIRE() bool`

GetWIRE returns the WIRE field if non-nil, zero value otherwise.

### GetWIREOk

`func (o *TransmissionTypes) GetWIREOk() (*bool, bool)`

GetWIREOk returns a tuple with the WIRE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWIRE

`func (o *TransmissionTypes) SetWIRE(v bool)`

SetWIRE sets WIRE field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


