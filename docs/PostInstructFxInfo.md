# PostInstructFxInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FxMode** | **string** | The mode by which the FX rate is to be determined (MANUAL or AUTO) | 
**FxStatus** | **string** | The state to which the Post-Instruct FX process has reached (INITIATED or COMPLETED) | 
**FxStatusUpdatedAt** | **time.Time** | The date-time at which the most recent fxStatus was determined. | 
**FxTransactionReference** | Pointer to **string** | The reference assigned to the FX funding that will fulfil this payment. | [optional] 

## Methods

### NewPostInstructFxInfo

`func NewPostInstructFxInfo(fxMode string, fxStatus string, fxStatusUpdatedAt time.Time, ) *PostInstructFxInfo`

NewPostInstructFxInfo instantiates a new PostInstructFxInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstructFxInfoWithDefaults

`func NewPostInstructFxInfoWithDefaults() *PostInstructFxInfo`

NewPostInstructFxInfoWithDefaults instantiates a new PostInstructFxInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFxMode

`func (o *PostInstructFxInfo) GetFxMode() string`

GetFxMode returns the FxMode field if non-nil, zero value otherwise.

### GetFxModeOk

`func (o *PostInstructFxInfo) GetFxModeOk() (*string, bool)`

GetFxModeOk returns a tuple with the FxMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxMode

`func (o *PostInstructFxInfo) SetFxMode(v string)`

SetFxMode sets FxMode field to given value.


### GetFxStatus

`func (o *PostInstructFxInfo) GetFxStatus() string`

GetFxStatus returns the FxStatus field if non-nil, zero value otherwise.

### GetFxStatusOk

`func (o *PostInstructFxInfo) GetFxStatusOk() (*string, bool)`

GetFxStatusOk returns a tuple with the FxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxStatus

`func (o *PostInstructFxInfo) SetFxStatus(v string)`

SetFxStatus sets FxStatus field to given value.


### GetFxStatusUpdatedAt

`func (o *PostInstructFxInfo) GetFxStatusUpdatedAt() time.Time`

GetFxStatusUpdatedAt returns the FxStatusUpdatedAt field if non-nil, zero value otherwise.

### GetFxStatusUpdatedAtOk

`func (o *PostInstructFxInfo) GetFxStatusUpdatedAtOk() (*time.Time, bool)`

GetFxStatusUpdatedAtOk returns a tuple with the FxStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxStatusUpdatedAt

`func (o *PostInstructFxInfo) SetFxStatusUpdatedAt(v time.Time)`

SetFxStatusUpdatedAt sets FxStatusUpdatedAt field to given value.


### GetFxTransactionReference

`func (o *PostInstructFxInfo) GetFxTransactionReference() string`

GetFxTransactionReference returns the FxTransactionReference field if non-nil, zero value otherwise.

### GetFxTransactionReferenceOk

`func (o *PostInstructFxInfo) GetFxTransactionReferenceOk() (*string, bool)`

GetFxTransactionReferenceOk returns a tuple with the FxTransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxTransactionReference

`func (o *PostInstructFxInfo) SetFxTransactionReference(v string)`

SetFxTransactionReference sets FxTransactionReference field to given value.

### HasFxTransactionReference

`func (o *PostInstructFxInfo) HasFxTransactionReference() bool`

HasFxTransactionReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


