# PayeeResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV2**](PayeePayorRefV2.md) |  | [optional] [readonly] 
**Email** | Pointer to **string** |  | [optional] 
**OnboardedStatus** | [**OnboardedStatus2**](OnboardedStatus_2.md) |  | [optional] 
**OfacStatus** | [**OfacStatus2**](OfacStatus_2.md) |  | [optional] 
**Language** | [**Language2**](Language_2.md) |  | [optional] 
**Created** | [**time.Time**](time.Time.md) |  | [optional] 
**Country** | **string** |  | [optional] 
**DisplayName** | **string** |  | [optional] 
**PayeeType** | [**PayeeType**](PayeeType.md) |  | [optional] 
**Disabled** | **bool** |  | [optional] 
**DisabledComment** | **string** |  | [optional] 
**DisabledUpdatedTimestamp** | [**time.Time**](time.Time.md) |  | [optional] 
**Address** | [**PayeeAddress2**](PayeeAddress_2.md) |  | [optional] 
**Individual** | [**Individual**](Individual.md) |  | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**CellphoneNumber** | **string** |  | [optional] 
**LastOfacCheckTimestamp** | Pointer to **string** |  | [optional] [readonly] 
**OfacOverride** | **bool** |  | [optional] 
**OfacOverrideReason** | **string** |  | [optional] 
**OfacOverrideTimestamp** | Pointer to **string** |  | [optional] 
**GracePeriodEndDate** | Pointer to **string** |  | [optional] [readonly] 
**EnhancedKycCompleted** | **bool** |  | [optional] 
**KycCompletedTimestamp** | Pointer to **string** |  | [optional] 
**PausePayment** | **bool** |  | [optional] 
**PausePaymentTimestamp** | Pointer to **string** |  | [optional] 
**MarketingOptInDecision** | **bool** |  | [optional] 
**MarketingOptInTimestamp** | Pointer to **string** |  | [optional] 
**AcceptTermsAndConditionsTimestamp** | Pointer to [**time.Time**](time.Time.md) | The timestamp when the payee last accepted T&amp;Cs | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


