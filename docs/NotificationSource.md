# NotificationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly | 
**EventId** | **string** | UUID id of the source event in the Velo platform | 
**CreatedAt** | **time.Time** | ISO8601 timestamp indicating when the source event was created | 
**PaymentId** | **string** | ID of this payment within the Velo platform | 
**PayoutPayorIds** | Pointer to [**PayoutPayorIds**](PayoutPayorIds.md) |  | [optional] 
**PayorPaymentId** | Pointer to **string** | ID of this payment in the payors system | [optional] 
**Status** | **string** | The new status of the debit. One of \&quot;PENDING\&quot; \&quot;PROCESSING\&quot; \&quot;REJECTED\&quot; \&quot;RELEASED\&quot; | 
**ReasonCode** | **string** | The Velo code that indicates why the payment was rejected or returned | 
**ReasonMessage** | **string** | The description of why the payment was rejected or returned | 
**PayeeId** | **string** | ID of this payee within the Velo platform | 
**Reasons** | Pointer to [**[]PayeeEventAllOfReasons**](PayeeEventAllOfReasons.md) | The reasons for the event notification. | [optional] 
**DebitTransactionId** | **string** | ID of this debit transaction within the Velo platform | 
**RailsId** | Pointer to **string** | the identifier of the payment rail from which funding was received | [optional] 
**PayorId** | **string** | ID of the payor within the Velo platform | 
**FundingRequestId** | **string** | ID of this funding transaction within the Velo platform | 
**FundingRef** | Pointer to **string** | the external identity reference for this funding transaction | [optional] 
**Currency** | Pointer to **string** | the ISO-4217 code for the currency in which the funding was made | [optional] 
**Amount** | Pointer to **int64** | the received funding amount in currency minor units | [optional] 
**PhysicalAccountName** | Pointer to **string** | the name of the account as registered with the payment rail | [optional] 
**SourceAccountName** | Pointer to **string** | the name of the account as registered with the Velo platform | [optional] 
**SourceAccountId** | Pointer to **string** | the ID of the account as registered with the Velo platform | [optional] 
**AdditionalInformation** | Pointer to **string** | any additional information received from the payment rail | [optional] 
**TransactionId** | Pointer to **string** | The Id of the related transaction | [optional] 
**TransactionReference** | Pointer to **string** | The payors own reference for the related transaction | [optional] 

## Methods

### NewNotificationSource

`func NewNotificationSource(sourceType string, eventId string, createdAt time.Time, paymentId string, status string, reasonCode string, reasonMessage string, payeeId string, debitTransactionId string, payorId string, fundingRequestId string, ) *NotificationSource`

NewNotificationSource instantiates a new NotificationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSourceWithDefaults

`func NewNotificationSourceWithDefaults() *NotificationSource`

NewNotificationSourceWithDefaults instantiates a new NotificationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *NotificationSource) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *NotificationSource) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *NotificationSource) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetEventId

`func (o *NotificationSource) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *NotificationSource) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *NotificationSource) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCreatedAt

`func (o *NotificationSource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationSource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationSource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPaymentId

`func (o *NotificationSource) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *NotificationSource) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *NotificationSource) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPayoutPayorIds

`func (o *NotificationSource) GetPayoutPayorIds() PayoutPayorIds`

GetPayoutPayorIds returns the PayoutPayorIds field if non-nil, zero value otherwise.

### GetPayoutPayorIdsOk

`func (o *NotificationSource) GetPayoutPayorIdsOk() (*PayoutPayorIds, bool)`

GetPayoutPayorIdsOk returns a tuple with the PayoutPayorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutPayorIds

`func (o *NotificationSource) SetPayoutPayorIds(v PayoutPayorIds)`

SetPayoutPayorIds sets PayoutPayorIds field to given value.

### HasPayoutPayorIds

`func (o *NotificationSource) HasPayoutPayorIds() bool`

HasPayoutPayorIds returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *NotificationSource) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *NotificationSource) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *NotificationSource) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *NotificationSource) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *NotificationSource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationSource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationSource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReasonCode

`func (o *NotificationSource) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *NotificationSource) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *NotificationSource) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetReasonMessage

`func (o *NotificationSource) GetReasonMessage() string`

GetReasonMessage returns the ReasonMessage field if non-nil, zero value otherwise.

### GetReasonMessageOk

`func (o *NotificationSource) GetReasonMessageOk() (*string, bool)`

GetReasonMessageOk returns a tuple with the ReasonMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonMessage

`func (o *NotificationSource) SetReasonMessage(v string)`

SetReasonMessage sets ReasonMessage field to given value.


### GetPayeeId

`func (o *NotificationSource) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *NotificationSource) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *NotificationSource) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetReasons

`func (o *NotificationSource) GetReasons() []PayeeEventAllOfReasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *NotificationSource) GetReasonsOk() (*[]PayeeEventAllOfReasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *NotificationSource) SetReasons(v []PayeeEventAllOfReasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *NotificationSource) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetDebitTransactionId

`func (o *NotificationSource) GetDebitTransactionId() string`

GetDebitTransactionId returns the DebitTransactionId field if non-nil, zero value otherwise.

### GetDebitTransactionIdOk

`func (o *NotificationSource) GetDebitTransactionIdOk() (*string, bool)`

GetDebitTransactionIdOk returns a tuple with the DebitTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitTransactionId

`func (o *NotificationSource) SetDebitTransactionId(v string)`

SetDebitTransactionId sets DebitTransactionId field to given value.


### GetRailsId

`func (o *NotificationSource) GetRailsId() string`

GetRailsId returns the RailsId field if non-nil, zero value otherwise.

### GetRailsIdOk

`func (o *NotificationSource) GetRailsIdOk() (*string, bool)`

GetRailsIdOk returns a tuple with the RailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsId

`func (o *NotificationSource) SetRailsId(v string)`

SetRailsId sets RailsId field to given value.

### HasRailsId

`func (o *NotificationSource) HasRailsId() bool`

HasRailsId returns a boolean if a field has been set.

### GetPayorId

`func (o *NotificationSource) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *NotificationSource) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *NotificationSource) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetFundingRequestId

`func (o *NotificationSource) GetFundingRequestId() string`

GetFundingRequestId returns the FundingRequestId field if non-nil, zero value otherwise.

### GetFundingRequestIdOk

`func (o *NotificationSource) GetFundingRequestIdOk() (*string, bool)`

GetFundingRequestIdOk returns a tuple with the FundingRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRequestId

`func (o *NotificationSource) SetFundingRequestId(v string)`

SetFundingRequestId sets FundingRequestId field to given value.


### GetFundingRef

`func (o *NotificationSource) GetFundingRef() string`

GetFundingRef returns the FundingRef field if non-nil, zero value otherwise.

### GetFundingRefOk

`func (o *NotificationSource) GetFundingRefOk() (*string, bool)`

GetFundingRefOk returns a tuple with the FundingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRef

`func (o *NotificationSource) SetFundingRef(v string)`

SetFundingRef sets FundingRef field to given value.

### HasFundingRef

`func (o *NotificationSource) HasFundingRef() bool`

HasFundingRef returns a boolean if a field has been set.

### GetCurrency

`func (o *NotificationSource) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *NotificationSource) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *NotificationSource) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *NotificationSource) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *NotificationSource) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NotificationSource) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NotificationSource) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *NotificationSource) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPhysicalAccountName

`func (o *NotificationSource) GetPhysicalAccountName() string`

GetPhysicalAccountName returns the PhysicalAccountName field if non-nil, zero value otherwise.

### GetPhysicalAccountNameOk

`func (o *NotificationSource) GetPhysicalAccountNameOk() (*string, bool)`

GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountName

`func (o *NotificationSource) SetPhysicalAccountName(v string)`

SetPhysicalAccountName sets PhysicalAccountName field to given value.

### HasPhysicalAccountName

`func (o *NotificationSource) HasPhysicalAccountName() bool`

HasPhysicalAccountName returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *NotificationSource) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *NotificationSource) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *NotificationSource) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *NotificationSource) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetSourceAccountId

`func (o *NotificationSource) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *NotificationSource) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *NotificationSource) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.

### HasSourceAccountId

`func (o *NotificationSource) HasSourceAccountId() bool`

HasSourceAccountId returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *NotificationSource) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *NotificationSource) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *NotificationSource) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *NotificationSource) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetTransactionId

`func (o *NotificationSource) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *NotificationSource) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *NotificationSource) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *NotificationSource) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionReference

`func (o *NotificationSource) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *NotificationSource) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *NotificationSource) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *NotificationSource) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


