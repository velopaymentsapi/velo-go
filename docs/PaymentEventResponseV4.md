# PaymentEventResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The id of the event. | 
**EventDateTime** | [**time.Time**](time.Time.md) | The date/time at which the event occurred. | 
**EventType** | **string** | The type of the event. | 
**SourceCurrency** | [**PaymentAuditCurrencyV4**](PaymentAuditCurrencyV4.md) |  | [optional] 
**SourceAmount** | **int64** | The source amount exposed by the event. | [optional] 
**PaymentCurrency** | [**PaymentAuditCurrencyV4**](PaymentAuditCurrencyV4.md) |  | [optional] 
**PaymentAmount** | **int64** | The destination amount exposed by the event. | [optional] 
**AccountNumber** | **string** | The account number attached to the event. | [optional] 
**RoutingNumber** | **string** | The routing number attached to the event. | [optional] 
**Iban** | **string** |  | [optional] 
**AccountName** | **string** |  | [optional] 
**Principal** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


