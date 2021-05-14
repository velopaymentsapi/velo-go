# PayeeDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** |  | 
**PayeeId** | **string** |  | [readonly] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DbaName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**PayeeCountry** | Pointer to **string** |  | [optional] 
**OnboardedStatus** | Pointer to [**OnboardedStatus2**](OnboardedStatus2.md) |  | [optional] 

## Methods

### NewPayeeDelta

`func NewPayeeDelta(remoteId string, payeeId string, ) *PayeeDelta`

NewPayeeDelta instantiates a new PayeeDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeDeltaWithDefaults

`func NewPayeeDeltaWithDefaults() *PayeeDelta`

NewPayeeDeltaWithDefaults instantiates a new PayeeDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *PayeeDelta) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayeeDelta) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayeeDelta) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetPayeeId

`func (o *PayeeDelta) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeDelta) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeDelta) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetDisplayName

`func (o *PayeeDelta) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PayeeDelta) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PayeeDelta) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PayeeDelta) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDbaName

`func (o *PayeeDelta) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *PayeeDelta) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *PayeeDelta) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *PayeeDelta) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### GetEmail

`func (o *PayeeDelta) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayeeDelta) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayeeDelta) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayeeDelta) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PayeeDelta) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PayeeDelta) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPayeeCountry

`func (o *PayeeDelta) GetPayeeCountry() string`

GetPayeeCountry returns the PayeeCountry field if non-nil, zero value otherwise.

### GetPayeeCountryOk

`func (o *PayeeDelta) GetPayeeCountryOk() (*string, bool)`

GetPayeeCountryOk returns a tuple with the PayeeCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeCountry

`func (o *PayeeDelta) SetPayeeCountry(v string)`

SetPayeeCountry sets PayeeCountry field to given value.

### HasPayeeCountry

`func (o *PayeeDelta) HasPayeeCountry() bool`

HasPayeeCountry returns a boolean if a field has been set.

### GetOnboardedStatus

`func (o *PayeeDelta) GetOnboardedStatus() OnboardedStatus2`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *PayeeDelta) GetOnboardedStatusOk() (*OnboardedStatus2, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *PayeeDelta) SetOnboardedStatus(v OnboardedStatus2)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *PayeeDelta) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


