# SourceAccountResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Source Account Id | 
**Balance** | Pointer to **int64** | Decimal implied | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**FundingRef** | Pointer to **string** | The funding reference (will not be set for DECOUPLED accounts). | [optional] 
**PhysicalAccountName** | Pointer to **string** | The physical account name (will not be set for DECOUPLED accounts). | [optional] 
**RailsId** | **string** |  | 
**PayorId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pooled** | Pointer to **bool** | The pooled account flag (will not be set for DECOUPLED accounts). | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**PhysicalAccountId** | Pointer to **string** | The physical account id (will not be set for DECOUPLED accounts). | [optional] 
**Notifications** | Pointer to [**Notifications2**](Notifications_2.md) |  | [optional] 
**AutoTopUpConfig** | Pointer to [**AutoTopUpConfig2**](AutoTopUpConfig_2.md) |  | [optional] 
**Type** | **string** |  | 
**Country** | Pointer to **string** | The two character ISO country code for the associated account | [optional] 
**Archived** | Pointer to **bool** | A flag for whether the source account has been archived.  Only present in the response if true. | [optional] 

## Methods

### NewSourceAccountResponseV3

`func NewSourceAccountResponseV3(id string, railsId string, type_ string, ) *SourceAccountResponseV3`

NewSourceAccountResponseV3 instantiates a new SourceAccountResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountResponseV3WithDefaults

`func NewSourceAccountResponseV3WithDefaults() *SourceAccountResponseV3`

NewSourceAccountResponseV3WithDefaults instantiates a new SourceAccountResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceAccountResponseV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceAccountResponseV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceAccountResponseV3) SetId(v string)`

SetId sets Id field to given value.


### GetBalance

`func (o *SourceAccountResponseV3) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SourceAccountResponseV3) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SourceAccountResponseV3) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *SourceAccountResponseV3) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *SourceAccountResponseV3) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SourceAccountResponseV3) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SourceAccountResponseV3) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SourceAccountResponseV3) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFundingRef

`func (o *SourceAccountResponseV3) GetFundingRef() string`

GetFundingRef returns the FundingRef field if non-nil, zero value otherwise.

### GetFundingRefOk

`func (o *SourceAccountResponseV3) GetFundingRefOk() (*string, bool)`

GetFundingRefOk returns a tuple with the FundingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRef

`func (o *SourceAccountResponseV3) SetFundingRef(v string)`

SetFundingRef sets FundingRef field to given value.

### HasFundingRef

`func (o *SourceAccountResponseV3) HasFundingRef() bool`

HasFundingRef returns a boolean if a field has been set.

### GetPhysicalAccountName

`func (o *SourceAccountResponseV3) GetPhysicalAccountName() string`

GetPhysicalAccountName returns the PhysicalAccountName field if non-nil, zero value otherwise.

### GetPhysicalAccountNameOk

`func (o *SourceAccountResponseV3) GetPhysicalAccountNameOk() (*string, bool)`

GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountName

`func (o *SourceAccountResponseV3) SetPhysicalAccountName(v string)`

SetPhysicalAccountName sets PhysicalAccountName field to given value.

### HasPhysicalAccountName

`func (o *SourceAccountResponseV3) HasPhysicalAccountName() bool`

HasPhysicalAccountName returns a boolean if a field has been set.

### GetRailsId

`func (o *SourceAccountResponseV3) GetRailsId() string`

GetRailsId returns the RailsId field if non-nil, zero value otherwise.

### GetRailsIdOk

`func (o *SourceAccountResponseV3) GetRailsIdOk() (*string, bool)`

GetRailsIdOk returns a tuple with the RailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsId

`func (o *SourceAccountResponseV3) SetRailsId(v string)`

SetRailsId sets RailsId field to given value.


### GetPayorId

`func (o *SourceAccountResponseV3) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *SourceAccountResponseV3) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *SourceAccountResponseV3) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *SourceAccountResponseV3) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetName

`func (o *SourceAccountResponseV3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceAccountResponseV3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceAccountResponseV3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceAccountResponseV3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPooled

`func (o *SourceAccountResponseV3) GetPooled() bool`

GetPooled returns the Pooled field if non-nil, zero value otherwise.

### GetPooledOk

`func (o *SourceAccountResponseV3) GetPooledOk() (*bool, bool)`

GetPooledOk returns a tuple with the Pooled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooled

`func (o *SourceAccountResponseV3) SetPooled(v bool)`

SetPooled sets Pooled field to given value.

### HasPooled

`func (o *SourceAccountResponseV3) HasPooled() bool`

HasPooled returns a boolean if a field has been set.

### GetCustomerId

`func (o *SourceAccountResponseV3) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SourceAccountResponseV3) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SourceAccountResponseV3) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SourceAccountResponseV3) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *SourceAccountResponseV3) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *SourceAccountResponseV3) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetPhysicalAccountId

`func (o *SourceAccountResponseV3) GetPhysicalAccountId() string`

GetPhysicalAccountId returns the PhysicalAccountId field if non-nil, zero value otherwise.

### GetPhysicalAccountIdOk

`func (o *SourceAccountResponseV3) GetPhysicalAccountIdOk() (*string, bool)`

GetPhysicalAccountIdOk returns a tuple with the PhysicalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountId

`func (o *SourceAccountResponseV3) SetPhysicalAccountId(v string)`

SetPhysicalAccountId sets PhysicalAccountId field to given value.

### HasPhysicalAccountId

`func (o *SourceAccountResponseV3) HasPhysicalAccountId() bool`

HasPhysicalAccountId returns a boolean if a field has been set.

### GetNotifications

`func (o *SourceAccountResponseV3) GetNotifications() Notifications2`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SourceAccountResponseV3) GetNotificationsOk() (*Notifications2, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SourceAccountResponseV3) SetNotifications(v Notifications2)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SourceAccountResponseV3) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetAutoTopUpConfig

`func (o *SourceAccountResponseV3) GetAutoTopUpConfig() AutoTopUpConfig2`

GetAutoTopUpConfig returns the AutoTopUpConfig field if non-nil, zero value otherwise.

### GetAutoTopUpConfigOk

`func (o *SourceAccountResponseV3) GetAutoTopUpConfigOk() (*AutoTopUpConfig2, bool)`

GetAutoTopUpConfigOk returns a tuple with the AutoTopUpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopUpConfig

`func (o *SourceAccountResponseV3) SetAutoTopUpConfig(v AutoTopUpConfig2)`

SetAutoTopUpConfig sets AutoTopUpConfig field to given value.

### HasAutoTopUpConfig

`func (o *SourceAccountResponseV3) HasAutoTopUpConfig() bool`

HasAutoTopUpConfig returns a boolean if a field has been set.

### GetType

`func (o *SourceAccountResponseV3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceAccountResponseV3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceAccountResponseV3) SetType(v string)`

SetType sets Type field to given value.


### GetCountry

`func (o *SourceAccountResponseV3) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SourceAccountResponseV3) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SourceAccountResponseV3) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SourceAccountResponseV3) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetArchived

`func (o *SourceAccountResponseV3) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *SourceAccountResponseV3) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *SourceAccountResponseV3) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *SourceAccountResponseV3) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


