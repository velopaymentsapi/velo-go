# ListSourceAccountResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedUserResponsePage**](PagedUserResponse_page.md) |  | [optional] 
**Links** | Pointer to [**[]ListSourceAccountResponseV3Links**](ListSourceAccountResponseV3_links.md) |  | [optional] 
**Content** | Pointer to [**[]SourceAccountResponseV3**](SourceAccountResponseV3.md) |  | [optional] 

## Methods

### NewListSourceAccountResponseV3

`func NewListSourceAccountResponseV3() *ListSourceAccountResponseV3`

NewListSourceAccountResponseV3 instantiates a new ListSourceAccountResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourceAccountResponseV3WithDefaults

`func NewListSourceAccountResponseV3WithDefaults() *ListSourceAccountResponseV3`

NewListSourceAccountResponseV3WithDefaults instantiates a new ListSourceAccountResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListSourceAccountResponseV3) GetPage() PagedUserResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListSourceAccountResponseV3) GetPageOk() (*PagedUserResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListSourceAccountResponseV3) SetPage(v PagedUserResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListSourceAccountResponseV3) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *ListSourceAccountResponseV3) GetLinks() []ListSourceAccountResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListSourceAccountResponseV3) GetLinksOk() (*[]ListSourceAccountResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListSourceAccountResponseV3) SetLinks(v []ListSourceAccountResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListSourceAccountResponseV3) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *ListSourceAccountResponseV3) GetContent() []SourceAccountResponseV3`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListSourceAccountResponseV3) GetContentOk() (*[]SourceAccountResponseV3, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListSourceAccountResponseV3) SetContent(v []SourceAccountResponseV3)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListSourceAccountResponseV3) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


