# ListFundingAccountsResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ListSourceAccountResponsePage**](ListSourceAccountResponsePage.md) |  | [optional] 
**Links** | Pointer to [**[]ListSourceAccountResponseLinks**](ListSourceAccountResponseLinks.md) |  | [optional] 
**Content** | Pointer to [**[]FundingAccountResponse2**](FundingAccountResponse2.md) |  | [optional] 

## Methods

### NewListFundingAccountsResponse2

`func NewListFundingAccountsResponse2() *ListFundingAccountsResponse2`

NewListFundingAccountsResponse2 instantiates a new ListFundingAccountsResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFundingAccountsResponse2WithDefaults

`func NewListFundingAccountsResponse2WithDefaults() *ListFundingAccountsResponse2`

NewListFundingAccountsResponse2WithDefaults instantiates a new ListFundingAccountsResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListFundingAccountsResponse2) GetPage() ListSourceAccountResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListFundingAccountsResponse2) GetPageOk() (*ListSourceAccountResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListFundingAccountsResponse2) SetPage(v ListSourceAccountResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListFundingAccountsResponse2) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *ListFundingAccountsResponse2) GetLinks() []ListSourceAccountResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListFundingAccountsResponse2) GetLinksOk() (*[]ListSourceAccountResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListFundingAccountsResponse2) SetLinks(v []ListSourceAccountResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListFundingAccountsResponse2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *ListFundingAccountsResponse2) GetContent() []FundingAccountResponse2`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListFundingAccountsResponse2) GetContentOk() (*[]FundingAccountResponse2, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListFundingAccountsResponse2) SetContent(v []FundingAccountResponse2)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListFundingAccountsResponse2) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


