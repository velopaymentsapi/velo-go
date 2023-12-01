# ListFundingAccountsResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ListFundingAccountsResponseV2Page**](ListFundingAccountsResponseV2Page.md) |  | [optional] 
**Links** | Pointer to [**[]ListSourceAccountResponseV2Links**](ListSourceAccountResponseV2Links.md) |  | [optional] 
**Content** | Pointer to [**[]FundingAccountResponseV2**](FundingAccountResponseV2.md) |  | [optional] 

## Methods

### NewListFundingAccountsResponseV2

`func NewListFundingAccountsResponseV2() *ListFundingAccountsResponseV2`

NewListFundingAccountsResponseV2 instantiates a new ListFundingAccountsResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFundingAccountsResponseV2WithDefaults

`func NewListFundingAccountsResponseV2WithDefaults() *ListFundingAccountsResponseV2`

NewListFundingAccountsResponseV2WithDefaults instantiates a new ListFundingAccountsResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListFundingAccountsResponseV2) GetPage() ListFundingAccountsResponseV2Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListFundingAccountsResponseV2) GetPageOk() (*ListFundingAccountsResponseV2Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListFundingAccountsResponseV2) SetPage(v ListFundingAccountsResponseV2Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListFundingAccountsResponseV2) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *ListFundingAccountsResponseV2) GetLinks() []ListSourceAccountResponseV2Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListFundingAccountsResponseV2) GetLinksOk() (*[]ListSourceAccountResponseV2Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListFundingAccountsResponseV2) SetLinks(v []ListSourceAccountResponseV2Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListFundingAccountsResponseV2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *ListFundingAccountsResponseV2) GetContent() []FundingAccountResponseV2`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListFundingAccountsResponseV2) GetContentOk() (*[]FundingAccountResponseV2, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListFundingAccountsResponseV2) SetContent(v []FundingAccountResponseV2)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListFundingAccountsResponseV2) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


