# GetPayeeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV3**](PayeePayorRefV3.md) |  | [optional] [readonly] 
**Email** | Pointer to **string** |  | [optional] 
**OnboardedStatus** | [**OnboardedStatus2**](OnboardedStatus_2.md) |  | [optional] 
**WatchlistStatus** | [**WatchlistStatus**](WatchlistStatus.md) |  | [optional] 
**WatchlistStatusUpdatedTimestamp** | Pointer to **string** |  | [optional] [readonly] 
**WatchlistOverrideComment** | Pointer to **string** |  | [optional] 
**Language** | [**Language2**](Language_2.md) |  | [optional] 
**Created** | [**time.Time**](time.Time.md) |  | [optional] 
**Country** | **string** |  | [optional] 
**DisplayName** | **string** |  | [optional] 
**PayeeType** | [**PayeeType**](PayeeType.md) |  | [optional] 
**Disabled** | **bool** |  | [optional] 
**DisabledComment** | **string** |  | [optional] 
**DisabledUpdatedTimestamp** | [**time.Time**](time.Time.md) |  | [optional] 
**Individual** | [**GetPayeeListResponseIndividual**](GetPayeeListResponseIndividual.md) |  | [optional] 
**Company** | [**GetPayeeListResponseCompany**](GetPayeeListResponseCompany.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


