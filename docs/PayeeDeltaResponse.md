# PayeeDeltaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PayeeDeltaResponsePage**](PayeeDeltaResponse_page.md) |  | [optional] 
**Links** | Pointer to [**[]PayeeDeltaResponseLinks**](PayeeDeltaResponse_links.md) |  | [optional] 
**Content** | Pointer to [**[]PayeeDelta**](PayeeDelta.md) |  | [optional] 

## Methods

### NewPayeeDeltaResponse

`func NewPayeeDeltaResponse() *PayeeDeltaResponse`

NewPayeeDeltaResponse instantiates a new PayeeDeltaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeDeltaResponseWithDefaults

`func NewPayeeDeltaResponseWithDefaults() *PayeeDeltaResponse`

NewPayeeDeltaResponseWithDefaults instantiates a new PayeeDeltaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PayeeDeltaResponse) GetPage() PayeeDeltaResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PayeeDeltaResponse) GetPageOk() (*PayeeDeltaResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PayeeDeltaResponse) SetPage(v PayeeDeltaResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *PayeeDeltaResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PayeeDeltaResponse) GetLinks() []PayeeDeltaResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PayeeDeltaResponse) GetLinksOk() (*[]PayeeDeltaResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PayeeDeltaResponse) SetLinks(v []PayeeDeltaResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PayeeDeltaResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PayeeDeltaResponse) GetContent() []PayeeDelta`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PayeeDeltaResponse) GetContentOk() (*[]PayeeDelta, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PayeeDeltaResponse) SetContent(v []PayeeDelta)`

SetContent sets Content field to given value.

### HasContent

`func (o *PayeeDeltaResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


