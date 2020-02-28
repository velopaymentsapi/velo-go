# Payee2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [optional] 
**PayeeId** | **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV2**](PayeePayorRefV2.md) |  | [optional] [readonly] 
**Email** | Pointer to **string** |  | [optional] 
**Address** | [**PayeeAddress**](PayeeAddress.md) |  | [optional] 
**Country** | **string** |  | [optional] 
**DisplayName** | **string** |  | [optional] 
**PaymentChannel** | [**PayeePaymentChannel2**](PayeePaymentChannel_2.md) |  | [optional] 
**Challenge** | [**Challenge**](Challenge.md) |  | [optional] 
**Language** | [**Language2**](Language_2.md) |  | [optional] 
**AcceptTermsAndConditionsTimestamp** | Pointer to [**time.Time**](time.Time.md) | The timestamp when the payee last accepted T&amp;Cs | [optional] [readonly] 
**CellphoneNumber** | **string** |  | [optional] 
**PayeeType** | [**PayeeType**](PayeeType.md) |  | [optional] 
**Company** | Pointer to [**CompanyV1**](CompanyV1.md) |  | [optional] 
**Individual** | [**IndividualV1**](IndividualV1.md) |  | [optional] 
**Created** | [**time.Time**](time.Time.md) |  | [optional] 
**GracePeriodEndDate** | Pointer to **string** |  | [optional] [readonly] 
**WatchlistStatusUpdatedTimestamp** | Pointer to **string** |  | [optional] [readonly] 
**MarketingOptIns** | [**[]MarketingOptIn**](MarketingOptIn.md) |  | [optional] 
**OnboardedStatus** | [**OnboardedStatus**](OnboardedStatus.md) |  | [optional] 
**RemoteId** | **string** | Remote Id must be between 1 and 100 characters long | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


