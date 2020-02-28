# GetPaymentsForPayoutResponseV3Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutStatus** | **string** | The current status of the payout. | [optional] 
**SubmittedDateTime** | [**time.Time**](time.Time.md) | The date/time at which the payout was submitted. | [optional] 
**InstructedDateTime** | [**time.Time**](time.Time.md) | The date/time at which the payout was instructed. | [optional] 
**WithdrawnDateTime** | [**time.Time**](time.Time.md) | The date/time at which the payout was withdrawn. | [optional] 
**PayoutMemo** | **string** | The memo attached to the payout. | [optional] 
**TotalPayments** | **int32** | The count of payments within the payout. | [optional] 
**ConfirmedPayments** | **int32** | The count of payments within the payout which have been confirmed. | [optional] 
**ReleasedPayments** | **int32** | The count of payments within the payout which have been released. | [optional] 
**IncompletePayments** | **int32** | The count of payments within the payout which are incomplete. | [optional] 
**FailedPayments** | **int32** | The count of payments within the payout which have failed or been returned. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


