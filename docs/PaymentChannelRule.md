# PaymentChannelRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Element** | **string** | &lt;p&gt;the rule element&lt;/p&gt; &lt;p&gt;will match a given element name for a payment channel configuration  | 
**Required** | **bool** | is this element required | 
**DisplayName** | **string** | User friendly name | 
**MinLength** | Pointer to **int32** | mininum length of the element data | [optional] 
**MaxLength** | Pointer to **int32** | maximum length of the element data | [optional] 
**Validation** | **string** | a regex to validate the element data | 
**DisplayOrder** | Pointer to **int32** |  | [optional] 

## Methods

### NewPaymentChannelRule

`func NewPaymentChannelRule(element string, required bool, displayName string, validation string, ) *PaymentChannelRule`

NewPaymentChannelRule instantiates a new PaymentChannelRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChannelRuleWithDefaults

`func NewPaymentChannelRuleWithDefaults() *PaymentChannelRule`

NewPaymentChannelRuleWithDefaults instantiates a new PaymentChannelRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElement

`func (o *PaymentChannelRule) GetElement() string`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *PaymentChannelRule) GetElementOk() (*string, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *PaymentChannelRule) SetElement(v string)`

SetElement sets Element field to given value.


### GetRequired

`func (o *PaymentChannelRule) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PaymentChannelRule) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PaymentChannelRule) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDisplayName

`func (o *PaymentChannelRule) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentChannelRule) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentChannelRule) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMinLength

`func (o *PaymentChannelRule) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PaymentChannelRule) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PaymentChannelRule) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *PaymentChannelRule) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *PaymentChannelRule) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *PaymentChannelRule) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *PaymentChannelRule) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *PaymentChannelRule) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetValidation

`func (o *PaymentChannelRule) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *PaymentChannelRule) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *PaymentChannelRule) SetValidation(v string)`

SetValidation sets Validation field to given value.


### GetDisplayOrder

`func (o *PaymentChannelRule) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PaymentChannelRule) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PaymentChannelRule) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *PaymentChannelRule) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


