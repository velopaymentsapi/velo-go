# PayorV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** |  | [optional] [readonly] 
**PayorName** | **string** | The name of the payor. | 
**Address** | [**PayorAddress**](PayorAddress.md) |  | [optional] 
**PrimaryContactName** | **string** | Name of primary contact for the payor. | [optional] 
**PrimaryContactPhone** | **string** | Primary contact phone number for the payor. | [optional] 
**PrimaryContactEmail** | **string** | Primary contact email for the payor. | [optional] 
**FundingAccountRoutingNumber** | **string** | The funding account routing number to be used for the payor. | [optional] 
**FundingAccountAccountNumber** | **string** | The funding account number to be used for the payor. | [optional] 
**FundingAccountAccountName** | **string** | The funding account name to be used for the payor. | [optional] 
**KycState** | [**KycState**](KycState.md) |  | [optional] 
**ManualLockout** | **bool** | Whether or not the payor has been manually locked by the backoffice. | [optional] 
**PayeeGracePeriodProcessingEnabled** | **bool** | Whether grace period processing is enabled. | [optional] [readonly] 
**PayeeGracePeriodDays** | **int32** | The grace period for paying payees in days. | [optional] [readonly] 
**CollectiveAlias** | **string** | How the payor has chosen to refer to payees. | [optional] 
**SupportContact** | **string** | The payor’s support contact email address. | [optional] 
**DbaName** | **string** | The payor’s &#39;Doing Business As&#39; name. | [optional] 
**AllowsLanguageChoice** | **bool** | Whether or not the payor allows language choice in the UI. | [optional] 
**ReminderEmailsOptOut** | **bool** | Whether or not the payor has opted-out of reminder emails being sent. | [optional] [readonly] 
**Language** | **string** | The payor’s language preference. Must be one of [EN, FR]. | [optional] 
**IncludesReports** | **bool** |  | [optional] 
**MaxMasterPayorAdmins** | **int32** |  | [optional] 
**TransmissionTypes** | [**TransmissionTypes**](TransmissionTypes.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


