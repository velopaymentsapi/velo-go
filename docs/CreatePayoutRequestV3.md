# CreatePayoutRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** | Deprecated in v2.16. Any value supplied here will be ignored. | [optional] 
**PayoutFromPayorId** | **string** | The id of the payor whose source account(s) will be debited. payoutFromPayorId and payoutToPayorId must be both supplied or both omitted. | [optional] 
**PayoutToPayorId** | **string** | The id of the payor whose payees will be paid. payoutFromPayorId and payoutToPayorId must be both supplied or both omitted. | [optional] 
**PayoutMemo** | **string** | Text applied to all payment memos unless specified explicitly on a payment. This should be the reference field on the statement seen by the payee (but not via ACH) | [optional] 
**Payments** | [**[]PaymentInstructionV3**](PaymentInstructionV3.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


