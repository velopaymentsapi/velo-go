# OnboardingStatusChanged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly | 
**EventId** | **string** | UUID id of the source event in the Velo platform | 
**CreatedAt** | **time.Time** | ISO8601 timestamp indicating when the source event was created | 
**PayeeId** | **string** | ID of this payee within the Velo platform | 
**Reasons** | Pointer to [**[]PayeeEventAllOfReasons**](PayeeEventAllOfReasons.md) | The reasons for the event notification. | [optional] 

## Methods

### NewOnboardingStatusChanged

`func NewOnboardingStatusChanged(sourceType string, eventId string, createdAt time.Time, payeeId string, ) *OnboardingStatusChanged`

NewOnboardingStatusChanged instantiates a new OnboardingStatusChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingStatusChangedWithDefaults

`func NewOnboardingStatusChangedWithDefaults() *OnboardingStatusChanged`

NewOnboardingStatusChangedWithDefaults instantiates a new OnboardingStatusChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *OnboardingStatusChanged) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *OnboardingStatusChanged) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *OnboardingStatusChanged) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetEventId

`func (o *OnboardingStatusChanged) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *OnboardingStatusChanged) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *OnboardingStatusChanged) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCreatedAt

`func (o *OnboardingStatusChanged) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OnboardingStatusChanged) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OnboardingStatusChanged) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPayeeId

`func (o *OnboardingStatusChanged) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *OnboardingStatusChanged) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *OnboardingStatusChanged) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetReasons

`func (o *OnboardingStatusChanged) GetReasons() []PayeeEventAllOfReasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *OnboardingStatusChanged) GetReasonsOk() (*[]PayeeEventAllOfReasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *OnboardingStatusChanged) SetReasons(v []PayeeEventAllOfReasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *OnboardingStatusChanged) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


