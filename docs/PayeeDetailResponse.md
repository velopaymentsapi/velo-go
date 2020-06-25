# PayeeDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV3**](PayeePayorRefV3.md) |  | [optional] [readonly] 
**Email** | Pointer to **string** |  | [optional] 
**OnboardedStatus** | [**OnboardedStatus2**](OnboardedStatus_2.md) |  | [optional] 
**WatchlistStatus** | [**WatchlistStatus**](WatchlistStatus.md) |  | [optional] 
**WatchlistOverrideExpiresAtTimestamp** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**WatchlistOverrideComment** | **string** |  | [optional] 
**Language** | [**Language2**](Language_2.md) |  | [optional] 
**Created** | [**time.Time**](time.Time.md) |  | [optional] 
**Country** | **string** |  | [optional] 
**DisplayName** | **string** |  | [optional] 
**PayeeType** | [**PayeeType**](PayeeType.md) |  | [optional] 
**Disabled** | **bool** |  | [optional] 
**DisabledComment** | **string** |  | [optional] 
**DisabledUpdatedTimestamp** | [**time.Time**](time.Time.md) |  | [optional] 
**Address** | [**PayeeAddress2**](PayeeAddress_2.md) |  | [optional] 
**Individual** | [**Individual2**](Individual_2.md) |  | [optional] 
**Company** | Pointer to [**Company2**](Company_2.md) |  | [optional] 
**CellphoneNumber** | **string** |  | [optional] 
**WatchlistStatusUpdatedTimestamp** | Pointer to **string** |  | [optional] [readonly] 
**GracePeriodEndDate** | Pointer to **string** |  | [optional] [readonly] 
**EnhancedKycCompleted** | **bool** |  | [optional] 
**KycCompletedTimestamp** | Pointer to **string** |  | [optional] 
**PausePayment** | **bool** |  | [optional] 
**PausePaymentTimestamp** | Pointer to **string** |  | [optional] 
**MarketingOptInDecision** | **bool** |  | [optional] 
**MarketingOptInTimestamp** | Pointer to **string** |  | [optional] 
**AcceptTermsAndConditionsTimestamp** | Pointer to [**time.Time**](time.Time.md) | The timestamp when the payee last accepted T&amp;Cs | [optional] [readonly] 
**Challenge** | [**Challenge2**](Challenge_2.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


