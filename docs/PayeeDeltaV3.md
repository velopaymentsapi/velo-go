# PayeeDeltaV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** |  | 
**PayeeId** | **string** |  | [readonly] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DbaName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**PayeeCountry** | Pointer to **string** |  | [optional] 
**OnboardedStatus** | Pointer to **string** | Onboarded status. One of the following values: CREATED, INVITED, REGISTERED, ONBOARDED | [optional] 

## Methods

### NewPayeeDeltaV3

`func NewPayeeDeltaV3(remoteId string, payeeId string, ) *PayeeDeltaV3`

NewPayeeDeltaV3 instantiates a new PayeeDeltaV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeDeltaV3WithDefaults

`func NewPayeeDeltaV3WithDefaults() *PayeeDeltaV3`

NewPayeeDeltaV3WithDefaults instantiates a new PayeeDeltaV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *PayeeDeltaV3) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayeeDeltaV3) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayeeDeltaV3) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetPayeeId

`func (o *PayeeDeltaV3) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeDeltaV3) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeDeltaV3) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetDisplayName

`func (o *PayeeDeltaV3) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PayeeDeltaV3) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PayeeDeltaV3) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PayeeDeltaV3) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDbaName

`func (o *PayeeDeltaV3) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *PayeeDeltaV3) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *PayeeDeltaV3) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *PayeeDeltaV3) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### GetEmail

`func (o *PayeeDeltaV3) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayeeDeltaV3) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayeeDeltaV3) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayeeDeltaV3) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PayeeDeltaV3) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PayeeDeltaV3) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPayeeCountry

`func (o *PayeeDeltaV3) GetPayeeCountry() string`

GetPayeeCountry returns the PayeeCountry field if non-nil, zero value otherwise.

### GetPayeeCountryOk

`func (o *PayeeDeltaV3) GetPayeeCountryOk() (*string, bool)`

GetPayeeCountryOk returns a tuple with the PayeeCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeCountry

`func (o *PayeeDeltaV3) SetPayeeCountry(v string)`

SetPayeeCountry sets PayeeCountry field to given value.

### HasPayeeCountry

`func (o *PayeeDeltaV3) HasPayeeCountry() bool`

HasPayeeCountry returns a boolean if a field has been set.

### GetOnboardedStatus

`func (o *PayeeDeltaV3) GetOnboardedStatus() string`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *PayeeDeltaV3) GetOnboardedStatusOk() (*string, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *PayeeDeltaV3) SetOnboardedStatus(v string)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *PayeeDeltaV3) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


