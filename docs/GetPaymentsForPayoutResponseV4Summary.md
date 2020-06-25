# GetPaymentsForPayoutResponseV4Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutStatus** | **string** | The current status of the payout. | [optional] 
**SubmittedDateTime** | [**time.Time**](time.Time.md) | The date/time at which the payout was submitted. | [optional] 
**InstructedDateTime** | [**time.Time**](time.Time.md) | The date/time at which the payout was instructed. | [optional] 
**WithdrawnDateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**QuotedDateTime** | [**time.Time**](time.Time.md) | The date/time at which the payout was quoted. | [optional] 
**PayoutMemo** | **string** | The memo attached to the payout. | [optional] 
**TotalPayments** | **int32** | The count of payments within the payout. | [optional] 
**ConfirmedPayments** | **int32** | The count of payments within the payout which have been confirmed. | [optional] 
**ReleasedPayments** | **int32** | The count of payments within the payout which have been released. | [optional] 
**IncompletePayments** | **int32** | The count of payments within the payout which are incomplete. | [optional] 
**ReturnedPayments** | **int32** | The count of payments within the payout which have been returned. | [optional] 
**WithdrawnPayments** | **int32** | The count of payments within the payout which have been withdrawn. | [optional] 
**PayoutType** | [**PayoutTypeV4**](PayoutTypeV4.md) |  | [optional] 
**Submitting** | [**PayoutPayorV4**](PayoutPayorV4.md) |  | [optional] 
**PayoutFrom** | [**PayoutPayorV4**](PayoutPayorV4.md) |  | [optional] 
**PayoutTo** | [**PayoutPayorV4**](PayoutPayorV4.md) |  | [optional] 
**Quoted** | [**PayoutPrincipalV4**](PayoutPrincipalV4.md) |  | [optional] 
**Instructed** | [**PayoutPrincipalV4**](PayoutPrincipalV4.md) |  | [optional] 
**Withdrawn** | [**PayoutPrincipalV4**](PayoutPrincipalV4.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


