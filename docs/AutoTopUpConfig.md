# AutoTopUpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Is auto top-up enabled? automatically trigger funding to top-up the source account balance when the balance falls below the configured minimum level. | 
**MinBalance** | Pointer to **int64** | When the payor balance falls below this level then auto top-up will be triggered. Note - This is in minor units. | [optional] 
**TargetBalance** | Pointer to **int64** | When the payor balance falls below the min balance then auto top-up will request funds bring the balance to this level. Note - this is in minor units. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


