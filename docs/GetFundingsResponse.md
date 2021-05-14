# GetFundingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedUserResponsePage**](PagedUserResponsePage.md) |  | [optional] 
**Links** | Pointer to [**[]GetFundingsResponseLinks**](GetFundingsResponseLinks.md) |  | [optional] 
**Content** | Pointer to [**[]FundingAudit**](FundingAudit.md) |  | [optional] 

## Methods

### NewGetFundingsResponse

`func NewGetFundingsResponse() *GetFundingsResponse`

NewGetFundingsResponse instantiates a new GetFundingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFundingsResponseWithDefaults

`func NewGetFundingsResponseWithDefaults() *GetFundingsResponse`

NewGetFundingsResponseWithDefaults instantiates a new GetFundingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetFundingsResponse) GetPage() PagedUserResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetFundingsResponse) GetPageOk() (*PagedUserResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetFundingsResponse) SetPage(v PagedUserResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetFundingsResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *GetFundingsResponse) GetLinks() []GetFundingsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetFundingsResponse) GetLinksOk() (*[]GetFundingsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetFundingsResponse) SetLinks(v []GetFundingsResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetFundingsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *GetFundingsResponse) GetContent() []FundingAudit`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetFundingsResponse) GetContentOk() (*[]FundingAudit, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetFundingsResponse) SetContent(v []FundingAudit)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetFundingsResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


