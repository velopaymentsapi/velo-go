# Ping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly | 
**EventId** | **string** | UUID id of the source event in the Velo platform | 
**CreatedAt** | **time.Time** | ISO8601 timestamp indicating when the source event was created | 

## Methods

### NewPing

`func NewPing(sourceType string, eventId string, createdAt time.Time, ) *Ping`

NewPing instantiates a new Ping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingWithDefaults

`func NewPingWithDefaults() *Ping`

NewPingWithDefaults instantiates a new Ping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *Ping) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Ping) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Ping) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetEventId

`func (o *Ping) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Ping) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Ping) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCreatedAt

`func (o *Ping) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Ping) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Ping) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


