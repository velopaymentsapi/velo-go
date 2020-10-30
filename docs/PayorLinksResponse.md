# PayorLinksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]PayorLinksResponseLinks**](PayorLinksResponse_links.md) |  | [optional] 
**Payors** | Pointer to [**[]PayorLinksResponsePayors**](PayorLinksResponse_payors.md) |  | [optional] 

## Methods

### NewPayorLinksResponse

`func NewPayorLinksResponse() *PayorLinksResponse`

NewPayorLinksResponse instantiates a new PayorLinksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorLinksResponseWithDefaults

`func NewPayorLinksResponseWithDefaults() *PayorLinksResponse`

NewPayorLinksResponseWithDefaults instantiates a new PayorLinksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PayorLinksResponse) GetLinks() []PayorLinksResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PayorLinksResponse) GetLinksOk() (*[]PayorLinksResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PayorLinksResponse) SetLinks(v []PayorLinksResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PayorLinksResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPayors

`func (o *PayorLinksResponse) GetPayors() []PayorLinksResponsePayors`

GetPayors returns the Payors field if non-nil, zero value otherwise.

### GetPayorsOk

`func (o *PayorLinksResponse) GetPayorsOk() (*[]PayorLinksResponsePayors, bool)`

GetPayorsOk returns a tuple with the Payors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayors

`func (o *PayorLinksResponse) SetPayors(v []PayorLinksResponsePayors)`

SetPayors sets Payors field to given value.

### HasPayors

`func (o *PayorLinksResponse) HasPayors() bool`

HasPayors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


