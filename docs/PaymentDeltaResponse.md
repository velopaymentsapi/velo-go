# PaymentDeltaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedPayeeInvitationStatusResponseV3Page**](PagedPayeeInvitationStatusResponseV3Page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseV3Links**](PagedPayeeResponseV3Links.md) |  | [optional] 
**Content** | Pointer to [**[]PaymentDelta**](PaymentDelta.md) |  | [optional] 

## Methods

### NewPaymentDeltaResponse

`func NewPaymentDeltaResponse() *PaymentDeltaResponse`

NewPaymentDeltaResponse instantiates a new PaymentDeltaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDeltaResponseWithDefaults

`func NewPaymentDeltaResponseWithDefaults() *PaymentDeltaResponse`

NewPaymentDeltaResponseWithDefaults instantiates a new PaymentDeltaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PaymentDeltaResponse) GetPage() PagedPayeeInvitationStatusResponseV3Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaymentDeltaResponse) GetPageOk() (*PagedPayeeInvitationStatusResponseV3Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaymentDeltaResponse) SetPage(v PagedPayeeInvitationStatusResponseV3Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaymentDeltaResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PaymentDeltaResponse) GetLinks() []PagedPayeeResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaymentDeltaResponse) GetLinksOk() (*[]PagedPayeeResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaymentDeltaResponse) SetLinks(v []PagedPayeeResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaymentDeltaResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PaymentDeltaResponse) GetContent() []PaymentDelta`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PaymentDeltaResponse) GetContentOk() (*[]PaymentDelta, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PaymentDeltaResponse) SetContent(v []PaymentDelta)`

SetContent sets Content field to given value.

### HasContent

`func (o *PaymentDeltaResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


