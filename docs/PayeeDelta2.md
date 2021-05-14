# PayeeDelta2

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

### NewPayeeDelta2

`func NewPayeeDelta2(remoteId string, payeeId string, ) *PayeeDelta2`

NewPayeeDelta2 instantiates a new PayeeDelta2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeDelta2WithDefaults

`func NewPayeeDelta2WithDefaults() *PayeeDelta2`

NewPayeeDelta2WithDefaults instantiates a new PayeeDelta2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *PayeeDelta2) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayeeDelta2) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayeeDelta2) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetPayeeId

`func (o *PayeeDelta2) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeDelta2) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeDelta2) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetDisplayName

`func (o *PayeeDelta2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PayeeDelta2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PayeeDelta2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PayeeDelta2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDbaName

`func (o *PayeeDelta2) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *PayeeDelta2) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *PayeeDelta2) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *PayeeDelta2) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### GetEmail

`func (o *PayeeDelta2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayeeDelta2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayeeDelta2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayeeDelta2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PayeeDelta2) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PayeeDelta2) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPayeeCountry

`func (o *PayeeDelta2) GetPayeeCountry() string`

GetPayeeCountry returns the PayeeCountry field if non-nil, zero value otherwise.

### GetPayeeCountryOk

`func (o *PayeeDelta2) GetPayeeCountryOk() (*string, bool)`

GetPayeeCountryOk returns a tuple with the PayeeCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeCountry

`func (o *PayeeDelta2) SetPayeeCountry(v string)`

SetPayeeCountry sets PayeeCountry field to given value.

### HasPayeeCountry

`func (o *PayeeDelta2) HasPayeeCountry() bool`

HasPayeeCountry returns a boolean if a field has been set.

### GetOnboardedStatus

`func (o *PayeeDelta2) GetOnboardedStatus() OnboardedStatus2`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *PayeeDelta2) GetOnboardedStatusOk() (*OnboardedStatus2, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *PayeeDelta2) SetOnboardedStatus(v OnboardedStatus2)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *PayeeDelta2) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


