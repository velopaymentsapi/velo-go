# ListSourceAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ListSourceAccountResponsePage**](ListSourceAccountResponsePage.md) |  | [optional] 
**Links** | Pointer to [**[]ListSourceAccountResponseLinks**](ListSourceAccountResponseLinks.md) |  | [optional] 
**Content** | Pointer to [**[]SourceAccountResponse**](SourceAccountResponse.md) |  | [optional] 

## Methods

### NewListSourceAccountResponse

`func NewListSourceAccountResponse() *ListSourceAccountResponse`

NewListSourceAccountResponse instantiates a new ListSourceAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourceAccountResponseWithDefaults

`func NewListSourceAccountResponseWithDefaults() *ListSourceAccountResponse`

NewListSourceAccountResponseWithDefaults instantiates a new ListSourceAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListSourceAccountResponse) GetPage() ListSourceAccountResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListSourceAccountResponse) GetPageOk() (*ListSourceAccountResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListSourceAccountResponse) SetPage(v ListSourceAccountResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListSourceAccountResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *ListSourceAccountResponse) GetLinks() []ListSourceAccountResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListSourceAccountResponse) GetLinksOk() (*[]ListSourceAccountResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListSourceAccountResponse) SetLinks(v []ListSourceAccountResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListSourceAccountResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *ListSourceAccountResponse) GetContent() []SourceAccountResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListSourceAccountResponse) GetContentOk() (*[]SourceAccountResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListSourceAccountResponse) SetContent(v []SourceAccountResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListSourceAccountResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


