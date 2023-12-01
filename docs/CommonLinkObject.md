# CommonLinkObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rel** | Pointer to **string** | One of: first, last, or self | [optional] 
**Href** | Pointer to **string** | the resource URI | [optional] 

## Methods

### NewCommonLinkObject

`func NewCommonLinkObject() *CommonLinkObject`

NewCommonLinkObject instantiates a new CommonLinkObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonLinkObjectWithDefaults

`func NewCommonLinkObjectWithDefaults() *CommonLinkObject`

NewCommonLinkObjectWithDefaults instantiates a new CommonLinkObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRel

`func (o *CommonLinkObject) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *CommonLinkObject) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *CommonLinkObject) SetRel(v string)`

SetRel sets Rel field to given value.

### HasRel

`func (o *CommonLinkObject) HasRel() bool`

HasRel returns a boolean if a field has been set.

### GetHref

`func (o *CommonLinkObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CommonLinkObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CommonLinkObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CommonLinkObject) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


