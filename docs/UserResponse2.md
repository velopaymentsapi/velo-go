# UserResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the user | [optional] 
**Status** | **string** | The status of the user when the user has been invited but not yet enrolled they will have a PENDING status  | [optional] 
**Email** | **string** | the email address of the user | [optional] 
**SmsNumber** | **string** | The phone number of a device that the user can receive sms messages on  | [optional] 
**PrimaryContactNumber** | **string** | The main contact number for the user  | [optional] 
**SecondaryContactNumber** | **string** | The secondary contact number for the user  | [optional] 
**FirstName** | **string** |  | [optional] 
**LastName** | **string** |  | [optional] 
**EntityId** | **string** | The payorId or payeeId or null if the user is not a payor or payee user  | [optional] 
**Roles** | [**[]UserResponse2Roles**](UserResponse_2_roles.md) | The role(s) for the user  | [optional] 
**MfaType** | **string** | The type of the MFA device | [optional] 
**MfaVerified** | **bool** | Will be true if the user has logged in successfully using the MFA Device  | [optional] 
**MfaStatus** | **string** | The status of the MFA device | [optional] 
**LockedOut** | **bool** | If true the user is currently locked out and unable to log in | [optional] 
**LockedOutTimestamp** | Pointer to [**time.Time**](time.Time.md) | A timestamp showing when the user was locked out If null then the user is not currently locked out  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


