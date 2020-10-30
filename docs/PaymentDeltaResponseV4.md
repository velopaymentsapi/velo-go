# PaymentDeltaResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedPayeeInvitationStatusResponsePage**](PagedPayeeInvitationStatusResponse_page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseLinks**](PagedPayeeResponse_links.md) |  | [optional] 
**Content** | Pointer to [**[]PaymentDeltaV4**](PaymentDeltaV4.md) |  | [optional] 

## Methods

### NewPaymentDeltaResponseV4

`func NewPaymentDeltaResponseV4() *PaymentDeltaResponseV4`

NewPaymentDeltaResponseV4 instantiates a new PaymentDeltaResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDeltaResponseV4WithDefaults

`func NewPaymentDeltaResponseV4WithDefaults() *PaymentDeltaResponseV4`

NewPaymentDeltaResponseV4WithDefaults instantiates a new PaymentDeltaResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PaymentDeltaResponseV4) GetPage() PagedPayeeInvitationStatusResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaymentDeltaResponseV4) GetPageOk() (*PagedPayeeInvitationStatusResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaymentDeltaResponseV4) SetPage(v PagedPayeeInvitationStatusResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaymentDeltaResponseV4) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PaymentDeltaResponseV4) GetLinks() []PagedPayeeResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaymentDeltaResponseV4) GetLinksOk() (*[]PagedPayeeResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaymentDeltaResponseV4) SetLinks(v []PagedPayeeResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaymentDeltaResponseV4) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PaymentDeltaResponseV4) GetContent() []PaymentDeltaV4`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PaymentDeltaResponseV4) GetContentOk() (*[]PaymentDeltaV4, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PaymentDeltaResponseV4) SetContent(v []PaymentDeltaV4)`

SetContent sets Content field to given value.

### HasContent

`func (o *PaymentDeltaResponseV4) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


