# SourceAccountResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Source Account Id | 
**Balance** | **int64** | Decimal implied | [optional] 
**Currency** | **string** |  | [optional] 
**FundingRef** | **string** | The funding reference (will not be set for DECOUPLED accounts). | [optional] 
**PhysicalAccountName** | **string** | The physical account name (will not be set for DECOUPLED accounts). | [optional] 
**RailsId** | **string** |  | 
**PayorId** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Pooled** | **bool** | The pooled account flag (will not be set for DECOUPLED accounts). | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**PhysicalAccountId** | **string** | The physical account id (will not be set for DECOUPLED accounts). | [optional] 
**Notifications** | [**Notifications2**](Notifications_2.md) |  | [optional] 
**AutoTopUpConfig** | [**AutoTopUpConfig2**](AutoTopUpConfig_2.md) |  | [optional] 
**Type** | **string** |  | 
**Country** | **string** | The two character ISO country code for the associated account | [optional] 
**Archived** | **bool** | A flag for whether the source account has been archived.  Only present in the response if true. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


