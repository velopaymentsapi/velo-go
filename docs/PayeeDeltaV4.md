# PayeeDeltaV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** |  | 
**PayeeId** | **string** |  | [readonly] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DbaName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**PayeeCountry** | Pointer to **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**OnboardedStatus** | Pointer to **string** | Payee onboarded status. One of the following value: CREATED, INVITED, REGISTERED, ONBOARDED | [optional] 

## Methods

### NewPayeeDeltaV4

`func NewPayeeDeltaV4(remoteId string, payeeId string, ) *PayeeDeltaV4`

NewPayeeDeltaV4 instantiates a new PayeeDeltaV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeDeltaV4WithDefaults

`func NewPayeeDeltaV4WithDefaults() *PayeeDeltaV4`

NewPayeeDeltaV4WithDefaults instantiates a new PayeeDeltaV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *PayeeDeltaV4) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayeeDeltaV4) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayeeDeltaV4) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetPayeeId

`func (o *PayeeDeltaV4) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeDeltaV4) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeDeltaV4) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetDisplayName

`func (o *PayeeDeltaV4) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PayeeDeltaV4) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PayeeDeltaV4) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PayeeDeltaV4) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDbaName

`func (o *PayeeDeltaV4) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *PayeeDeltaV4) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *PayeeDeltaV4) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *PayeeDeltaV4) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### GetEmail

`func (o *PayeeDeltaV4) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayeeDeltaV4) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayeeDeltaV4) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayeeDeltaV4) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PayeeDeltaV4) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PayeeDeltaV4) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPayeeCountry

`func (o *PayeeDeltaV4) GetPayeeCountry() string`

GetPayeeCountry returns the PayeeCountry field if non-nil, zero value otherwise.

### GetPayeeCountryOk

`func (o *PayeeDeltaV4) GetPayeeCountryOk() (*string, bool)`

GetPayeeCountryOk returns a tuple with the PayeeCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeCountry

`func (o *PayeeDeltaV4) SetPayeeCountry(v string)`

SetPayeeCountry sets PayeeCountry field to given value.

### HasPayeeCountry

`func (o *PayeeDeltaV4) HasPayeeCountry() bool`

HasPayeeCountry returns a boolean if a field has been set.

### GetOnboardedStatus

`func (o *PayeeDeltaV4) GetOnboardedStatus() string`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *PayeeDeltaV4) GetOnboardedStatusOk() (*string, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *PayeeDeltaV4) SetOnboardedStatus(v string)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *PayeeDeltaV4) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


