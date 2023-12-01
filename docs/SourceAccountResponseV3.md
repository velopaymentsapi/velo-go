# SourceAccountResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Source Account Id | 
**Balance** | Pointer to **int64** | Decimal implied | [optional] 
**Currency** | Pointer to **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**FundingRef** | Pointer to **string** | The funding reference (will not be set for DECOUPLED accounts). | [optional] 
**PhysicalAccountName** | Pointer to **string** | The physical account name (will not be set for DECOUPLED accounts). | [optional] 
**RailsId** | **string** |  | 
**PayorId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pooled** | Pointer to **bool** | The pooled account flag (will not be set for DECOUPLED accounts). | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**PhysicalAccountId** | Pointer to **string** | The physical account id (will not be set for DECOUPLED accounts). | [optional] 
**Notifications** | Pointer to [**NotificationsV3**](NotificationsV3.md) |  | [optional] 
**AutoTopUpConfig** | Pointer to [**AutoTopUpConfigV3**](AutoTopUpConfigV3.md) |  | [optional] 
**Type** | **string** |  | 
**Country** | Pointer to **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**Deleted** | Pointer to **bool** | An optional flag for whether the source account has been deleted. Only present in the response if true. | [optional] 
**UserDeleted** | Pointer to **bool** | An optional flag for whether the source account has been deleted by a user. Only present in the response if true. | [optional] 
**DeletedAt** | Pointer to **time.Time** | An optional timestamp when the source account has been deleted. Only present in the response if deleted. | [optional] 
**TransmissionTypes** | Pointer to **[]string** | List of supported transmission types. | [optional] 

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

`func (o *SourceAccountResponseV3) GetNotifications() NotificationsV3`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SourceAccountResponseV3) GetNotificationsOk() (*NotificationsV3, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SourceAccountResponseV3) SetNotifications(v NotificationsV3)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SourceAccountResponseV3) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetAutoTopUpConfig

`func (o *SourceAccountResponseV3) GetAutoTopUpConfig() AutoTopUpConfigV3`

GetAutoTopUpConfig returns the AutoTopUpConfig field if non-nil, zero value otherwise.

### GetAutoTopUpConfigOk

`func (o *SourceAccountResponseV3) GetAutoTopUpConfigOk() (*AutoTopUpConfigV3, bool)`

GetAutoTopUpConfigOk returns a tuple with the AutoTopUpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopUpConfig

`func (o *SourceAccountResponseV3) SetAutoTopUpConfig(v AutoTopUpConfigV3)`

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

### GetDeleted

`func (o *SourceAccountResponseV3) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SourceAccountResponseV3) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SourceAccountResponseV3) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *SourceAccountResponseV3) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUserDeleted

`func (o *SourceAccountResponseV3) GetUserDeleted() bool`

GetUserDeleted returns the UserDeleted field if non-nil, zero value otherwise.

### GetUserDeletedOk

`func (o *SourceAccountResponseV3) GetUserDeletedOk() (*bool, bool)`

GetUserDeletedOk returns a tuple with the UserDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeleted

`func (o *SourceAccountResponseV3) SetUserDeleted(v bool)`

SetUserDeleted sets UserDeleted field to given value.

### HasUserDeleted

`func (o *SourceAccountResponseV3) HasUserDeleted() bool`

HasUserDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SourceAccountResponseV3) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SourceAccountResponseV3) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SourceAccountResponseV3) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SourceAccountResponseV3) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetTransmissionTypes

`func (o *SourceAccountResponseV3) GetTransmissionTypes() []string`

GetTransmissionTypes returns the TransmissionTypes field if non-nil, zero value otherwise.

### GetTransmissionTypesOk

`func (o *SourceAccountResponseV3) GetTransmissionTypesOk() (*[]string, bool)`

GetTransmissionTypesOk returns a tuple with the TransmissionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionTypes

`func (o *SourceAccountResponseV3) SetTransmissionTypes(v []string)`

SetTransmissionTypes sets TransmissionTypes field to given value.

### HasTransmissionTypes

`func (o *SourceAccountResponseV3) HasTransmissionTypes() bool`

HasTransmissionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


