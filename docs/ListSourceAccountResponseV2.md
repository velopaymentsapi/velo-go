# ListSourceAccountResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedUserResponsePage**](PagedUserResponsePage.md) |  | [optional] 
**Links** | Pointer to [**[]ListSourceAccountResponseV2Links**](ListSourceAccountResponseV2Links.md) |  | [optional] 
**Content** | Pointer to [**[]SourceAccountResponseV2**](SourceAccountResponseV2.md) |  | [optional] 

## Methods

### NewListSourceAccountResponseV2

`func NewListSourceAccountResponseV2() *ListSourceAccountResponseV2`

NewListSourceAccountResponseV2 instantiates a new ListSourceAccountResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourceAccountResponseV2WithDefaults

`func NewListSourceAccountResponseV2WithDefaults() *ListSourceAccountResponseV2`

NewListSourceAccountResponseV2WithDefaults instantiates a new ListSourceAccountResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListSourceAccountResponseV2) GetPage() PagedUserResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListSourceAccountResponseV2) GetPageOk() (*PagedUserResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListSourceAccountResponseV2) SetPage(v PagedUserResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListSourceAccountResponseV2) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *ListSourceAccountResponseV2) GetLinks() []ListSourceAccountResponseV2Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListSourceAccountResponseV2) GetLinksOk() (*[]ListSourceAccountResponseV2Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListSourceAccountResponseV2) SetLinks(v []ListSourceAccountResponseV2Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListSourceAccountResponseV2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *ListSourceAccountResponseV2) GetContent() []SourceAccountResponseV2`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListSourceAccountResponseV2) GetContentOk() (*[]SourceAccountResponseV2, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListSourceAccountResponseV2) SetContent(v []SourceAccountResponseV2)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListSourceAccountResponseV2) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


