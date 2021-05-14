# ListFundingAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ListSourceAccountResponsePage**](ListSourceAccountResponsePage.md) |  | [optional] 
**Links** | Pointer to [**[]ListSourceAccountResponseLinks**](ListSourceAccountResponseLinks.md) |  | [optional] 
**Content** | Pointer to [**[]FundingAccountResponse**](FundingAccountResponse.md) |  | [optional] 

## Methods

### NewListFundingAccountsResponse

`func NewListFundingAccountsResponse() *ListFundingAccountsResponse`

NewListFundingAccountsResponse instantiates a new ListFundingAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFundingAccountsResponseWithDefaults

`func NewListFundingAccountsResponseWithDefaults() *ListFundingAccountsResponse`

NewListFundingAccountsResponseWithDefaults instantiates a new ListFundingAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListFundingAccountsResponse) GetPage() ListSourceAccountResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListFundingAccountsResponse) GetPageOk() (*ListSourceAccountResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListFundingAccountsResponse) SetPage(v ListSourceAccountResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListFundingAccountsResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *ListFundingAccountsResponse) GetLinks() []ListSourceAccountResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListFundingAccountsResponse) GetLinksOk() (*[]ListSourceAccountResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListFundingAccountsResponse) SetLinks(v []ListSourceAccountResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListFundingAccountsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *ListFundingAccountsResponse) GetContent() []FundingAccountResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListFundingAccountsResponse) GetContentOk() (*[]FundingAccountResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListFundingAccountsResponse) SetContent(v []FundingAccountResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListFundingAccountsResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


