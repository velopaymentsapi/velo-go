# RegisterSmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmsNumber** | **string** | The phone number of a device that the user can receive sms messages on  | 

## Methods

### NewRegisterSmsRequest

`func NewRegisterSmsRequest(smsNumber string, ) *RegisterSmsRequest`

NewRegisterSmsRequest instantiates a new RegisterSmsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterSmsRequestWithDefaults

`func NewRegisterSmsRequestWithDefaults() *RegisterSmsRequest`

NewRegisterSmsRequestWithDefaults instantiates a new RegisterSmsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmsNumber

`func (o *RegisterSmsRequest) GetSmsNumber() string`

GetSmsNumber returns the SmsNumber field if non-nil, zero value otherwise.

### GetSmsNumberOk

`func (o *RegisterSmsRequest) GetSmsNumberOk() (*string, bool)`

GetSmsNumberOk returns a tuple with the SmsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumber

`func (o *RegisterSmsRequest) SetSmsNumber(v string)`

SetSmsNumber sets SmsNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


