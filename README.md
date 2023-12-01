# Go client for Velo
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![npm version](https://badge.fury.io/go/github.com%2Fvelopaymentsapi%2Fvelo-go.svg)](https://badge.fury.io/go/github.com%2Fvelopaymentsapi%2Fvelo-go) [![CircleCI](https://circleci.com/gh/velopaymentsapi/velo-go.svg?style=svg)](https://circleci.com/gh/velopaymentsapi/velo-go)\
## Terms and Definitions

Throughout this document and the Velo platform the following terms are used:

* **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout.
* **Payee.** The recipient of funds paid out by a payor.
* **Payment.** A single transfer of funds from a payor to a payee.
* **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee.
* **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.

## Overview

The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:

* Authenticate with the Velo platform
* Maintain a collection of payees
* Query the payor’s current balance of funds within the platform and perform additional funding
* Issue payments to payees
* Query the platform for a history of those payments

This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.

## API Considerations

The Velo Payments API is REST based and uses the JSON format for requests and responses.

Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.

Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).

Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.

## Authenticating with the Velo Platform

Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.

You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:

create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529

base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==

create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==

perform the Velo authentication REST call using the HTTP header created above e.g. via curl:

```
  curl -X POST \\
  -H \"Content-Type: application/json\" \\
  -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\
  'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials'
```

If successful, this call will result in a **200** HTTP status code and a response body such as:

```
  {
    \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",
    \"token_type\":\"bearer\",
    \"expires_in\":1799,
    \"scope\":\"...\"
  }
```
## API access following authentication
Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.

This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:

```
  -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \"
```

If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response.


## Http Status Codes
Following is a list of Http Status codes that could be returned by the platform

    | Status Code            | Description                                                                          |
    | -----------------------| -------------------------------------------------------------------------------------|
    | 200 OK                 | The request was successfully processed and usually returns a json response           |
    | 201 Created            | A resource was created and a Location header is returned linking to the new resource |
    | 202 Accepted           | The request has been accepted for processing                                         |
    | 204 No Content         | The request has been processed and there is no response (usually deletes and updates)|
    | 400 Bad Request        | The request is invalid and should be fixed before retrying                           |
    | 401 Unauthorized       | Authentication has failed, usually means the token has expired                       |
    | 403 Forbidden          | The user does not have permissions for the request                                   |
    | 404 Not Found          | The resource was not found                                                           |
    | 409 Conflict           | The resource already exists and there is a conflict                                  |
    | 429 Too Many Requests  | The user has submitted too many requests in a given amount of time                   |
    | 5xx Server Error       | Platform internal error (should rarely happen)                                       |


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.37.150
- Package version: 2.37.150.beta1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import velopayments "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), velopayments.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), velopayments.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), velopayments.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), velopayments.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.sandbox.velopayments.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CountriesAPI* | [**ListPaymentChannelRulesV1**](docs/CountriesAPI.md#listpaymentchannelrulesv1) | **Get** /v1/paymentChannelRules | List Payment Channel Country Rules
*CountriesAPI* | [**ListSupportedCountriesV1**](docs/CountriesAPI.md#listsupportedcountriesv1) | **Get** /v1/supportedCountries | List Supported Countries
*CountriesAPI* | [**ListSupportedCountriesV2**](docs/CountriesAPI.md#listsupportedcountriesv2) | **Get** /v2/supportedCountries | List Supported Countries
*CurrenciesAPI* | [**ListSupportedCurrenciesV2**](docs/CurrenciesAPI.md#listsupportedcurrenciesv2) | **Get** /v2/currencies | List Supported Currencies
*FundingAPI* | [**CreateFundingRequestV2**](docs/FundingAPI.md#createfundingrequestv2) | **Post** /v2/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
*FundingAPI* | [**CreateFundingRequestV3**](docs/FundingAPI.md#createfundingrequestv3) | **Post** /v3/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
*FundingAPI* | [**GetFundingAccountV2**](docs/FundingAPI.md#getfundingaccountv2) | **Get** /v2/fundingAccounts/{fundingAccountId} | Get Funding Account
*FundingAPI* | [**GetFundingAccountsV2**](docs/FundingAPI.md#getfundingaccountsv2) | **Get** /v2/fundingAccounts | Get Funding Accounts
*FundingAPI* | [**GetFundingByIdV1**](docs/FundingAPI.md#getfundingbyidv1) | **Get** /v1/fundings/{fundingId} | Get Funding
*FundingAPI* | [**ListFundingAuditDeltas**](docs/FundingAPI.md#listfundingauditdeltas) | **Get** /v1/deltas/fundings | Get Funding Audit Delta
*FundingManagerPrivateAPI* | [**CreateFundingAccountV2**](docs/FundingManagerPrivateAPI.md#createfundingaccountv2) | **Post** /v2/fundingAccounts | Create Funding Account
*FundingManagerPrivateAPI* | [**DeleteSourceAccountV3**](docs/FundingManagerPrivateAPI.md#deletesourceaccountv3) | **Delete** /v3/sourceAccounts/{sourceAccountId} | Delete a source account by ID
*LoginAPI* | [**Logout**](docs/LoginAPI.md#logout) | **Post** /v1/logout | Logout
*LoginAPI* | [**ResetPassword**](docs/LoginAPI.md#resetpassword) | **Post** /v1/password/reset | Reset password
*LoginAPI* | [**ValidateAccessToken**](docs/LoginAPI.md#validateaccesstoken) | **Post** /v1/validate | validate
*LoginAPI* | [**VeloAuth**](docs/LoginAPI.md#veloauth) | **Post** /v1/authenticate | Authentication endpoint
*PayeeInvitationAPI* | [**CreatePayeeV3**](docs/PayeeInvitationAPI.md#createpayeev3) | **Post** /v3/payees | Initiate Payee Creation
*PayeeInvitationAPI* | [**GetPayeesInvitationStatusV3**](docs/PayeeInvitationAPI.md#getpayeesinvitationstatusv3) | **Get** /v3/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationAPI* | [**GetPayeesInvitationStatusV4**](docs/PayeeInvitationAPI.md#getpayeesinvitationstatusv4) | **Get** /v4/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationAPI* | [**QueryBatchStatusV3**](docs/PayeeInvitationAPI.md#querybatchstatusv3) | **Get** /v3/payees/batch/{batchId} | Query Batch Status
*PayeeInvitationAPI* | [**QueryBatchStatusV4**](docs/PayeeInvitationAPI.md#querybatchstatusv4) | **Get** /v4/payees/batch/{batchId} | Query Batch Status
*PayeeInvitationAPI* | [**ResendPayeeInviteV3**](docs/PayeeInvitationAPI.md#resendpayeeinvitev3) | **Post** /v3/payees/{payeeId}/invite | Resend Payee Invite
*PayeeInvitationAPI* | [**ResendPayeeInviteV4**](docs/PayeeInvitationAPI.md#resendpayeeinvitev4) | **Post** /v4/payees/{payeeId}/invite | Resend Payee Invite
*PayeeInvitationAPI* | [**V4CreatePayee**](docs/PayeeInvitationAPI.md#v4createpayee) | **Post** /v4/payees | Initiate Payee Creation
*PayeePaymentChannelsAPI* | [**CreatePaymentChannelV4**](docs/PayeePaymentChannelsAPI.md#createpaymentchannelv4) | **Post** /v4/payees/{payeeId}/paymentChannels/ | Create Payment Channel
*PayeePaymentChannelsAPI* | [**DeletePaymentChannelV4**](docs/PayeePaymentChannelsAPI.md#deletepaymentchannelv4) | **Delete** /v4/payees/{payeeId}/paymentChannels/{paymentChannelId} | Delete Payment Channel
*PayeePaymentChannelsAPI* | [**EnablePaymentChannelV4**](docs/PayeePaymentChannelsAPI.md#enablepaymentchannelv4) | **Post** /v4/payees/{payeeId}/paymentChannels/{paymentChannelId}/enable | Enable Payment Channel
*PayeePaymentChannelsAPI* | [**GetPaymentChannelV4**](docs/PayeePaymentChannelsAPI.md#getpaymentchannelv4) | **Get** /v4/payees/{payeeId}/paymentChannels/{paymentChannelId} | Get Payment Channel Details
*PayeePaymentChannelsAPI* | [**GetPaymentChannelsV4**](docs/PayeePaymentChannelsAPI.md#getpaymentchannelsv4) | **Get** /v4/payees/{payeeId}/paymentChannels/ | Get All Payment Channels Details
*PayeePaymentChannelsAPI* | [**UpdatePaymentChannelOrderV4**](docs/PayeePaymentChannelsAPI.md#updatepaymentchannelorderv4) | **Put** /v4/payees/{payeeId}/paymentChannels/order | Update Payees preferred Payment Channel order
*PayeePaymentChannelsAPI* | [**UpdatePaymentChannelV4**](docs/PayeePaymentChannelsAPI.md#updatepaymentchannelv4) | **Post** /v4/payees/{payeeId}/paymentChannels/{paymentChannelId} | Update Payment Channel
*PayeesAPI* | [**DeletePayeeByIdV3**](docs/PayeesAPI.md#deletepayeebyidv3) | **Delete** /v3/payees/{payeeId} | Delete Payee by Id
*PayeesAPI* | [**DeletePayeeByIdV4**](docs/PayeesAPI.md#deletepayeebyidv4) | **Delete** /v4/payees/{payeeId} | Delete Payee by Id
*PayeesAPI* | [**GetPayeeByIdV3**](docs/PayeesAPI.md#getpayeebyidv3) | **Get** /v3/payees/{payeeId} | Get Payee by Id
*PayeesAPI* | [**GetPayeeByIdV4**](docs/PayeesAPI.md#getpayeebyidv4) | **Get** /v4/payees/{payeeId} | Get Payee by Id
*PayeesAPI* | [**ListPayeeChangesV3**](docs/PayeesAPI.md#listpayeechangesv3) | **Get** /v3/payees/deltas | List Payee Changes
*PayeesAPI* | [**ListPayeeChangesV4**](docs/PayeesAPI.md#listpayeechangesv4) | **Get** /v4/payees/deltas | List Payee Changes
*PayeesAPI* | [**ListPayeesV3**](docs/PayeesAPI.md#listpayeesv3) | **Get** /v3/payees | List Payees
*PayeesAPI* | [**ListPayeesV4**](docs/PayeesAPI.md#listpayeesv4) | **Get** /v4/payees | List Payees
*PayeesAPI* | [**PayeeDetailsUpdateV3**](docs/PayeesAPI.md#payeedetailsupdatev3) | **Post** /v3/payees/{payeeId}/payeeDetailsUpdate | Update Payee Details
*PayeesAPI* | [**PayeeDetailsUpdateV4**](docs/PayeesAPI.md#payeedetailsupdatev4) | **Post** /v4/payees/{payeeId}/payeeDetailsUpdate | Update Payee Details
*PayeesAPI* | [**V3PayeesPayeeIdRemoteIdUpdatePost**](docs/PayeesAPI.md#v3payeespayeeidremoteidupdatepost) | **Post** /v3/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
*PayeesAPI* | [**V4PayeesPayeeIdRemoteIdUpdatePost**](docs/PayeesAPI.md#v4payeespayeeidremoteidupdatepost) | **Post** /v4/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
*PaymentAuditServiceAPI* | [**ExportTransactionsCSVV4**](docs/PaymentAuditServiceAPI.md#exporttransactionscsvv4) | **Get** /v4/paymentaudit/transactions | Export Transactions
*PaymentAuditServiceAPI* | [**GetFundingsV4**](docs/PaymentAuditServiceAPI.md#getfundingsv4) | **Get** /v4/paymentaudit/fundings | Get Fundings for Payor
*PaymentAuditServiceAPI* | [**GetPaymentDetailsV4**](docs/PaymentAuditServiceAPI.md#getpaymentdetailsv4) | **Get** /v4/paymentaudit/payments/{paymentId} | Get Payment
*PaymentAuditServiceAPI* | [**GetPaymentsForPayoutV4**](docs/PaymentAuditServiceAPI.md#getpaymentsforpayoutv4) | **Get** /v4/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PaymentAuditServiceAPI* | [**GetPayoutStatsV4**](docs/PaymentAuditServiceAPI.md#getpayoutstatsv4) | **Get** /v4/paymentaudit/payoutStatistics | Get Payout Statistics
*PaymentAuditServiceAPI* | [**GetPayoutsForPayorV4**](docs/PaymentAuditServiceAPI.md#getpayoutsforpayorv4) | **Get** /v4/paymentaudit/payouts | Get Payouts for Payor
*PaymentAuditServiceAPI* | [**ListPaymentChangesV4**](docs/PaymentAuditServiceAPI.md#listpaymentchangesv4) | **Get** /v4/payments/deltas | List Payment Changes
*PaymentAuditServiceAPI* | [**ListPaymentsAuditV4**](docs/PaymentAuditServiceAPI.md#listpaymentsauditv4) | **Get** /v4/paymentaudit/payments | Get List of Payments
*PaymentAuditServiceDeprecatedAPI* | [**ExportTransactionsCSVV3**](docs/PaymentAuditServiceDeprecatedAPI.md#exporttransactionscsvv3) | **Get** /v3/paymentaudit/transactions | V3 Export Transactions
*PaymentAuditServiceDeprecatedAPI* | [**GetFundingsV1**](docs/PaymentAuditServiceDeprecatedAPI.md#getfundingsv1) | **Get** /v1/paymentaudit/fundings | V1 Get Fundings for Payor
*PaymentAuditServiceDeprecatedAPI* | [**GetPaymentDetailsV3**](docs/PaymentAuditServiceDeprecatedAPI.md#getpaymentdetailsv3) | **Get** /v3/paymentaudit/payments/{paymentId} | V3 Get Payment
*PaymentAuditServiceDeprecatedAPI* | [**GetPaymentsForPayoutPAV3**](docs/PaymentAuditServiceDeprecatedAPI.md#getpaymentsforpayoutpav3) | **Get** /v3/paymentaudit/payouts/{payoutId} | V3 Get Payments for Payout
*PaymentAuditServiceDeprecatedAPI* | [**GetPayoutStatsV1**](docs/PaymentAuditServiceDeprecatedAPI.md#getpayoutstatsv1) | **Get** /v1/paymentaudit/payoutStatistics | V1 Get Payout Statistics
*PaymentAuditServiceDeprecatedAPI* | [**GetPayoutsForPayorV3**](docs/PaymentAuditServiceDeprecatedAPI.md#getpayoutsforpayorv3) | **Get** /v3/paymentaudit/payouts | V3 Get Payouts for Payor
*PaymentAuditServiceDeprecatedAPI* | [**ListPaymentChanges**](docs/PaymentAuditServiceDeprecatedAPI.md#listpaymentchanges) | **Get** /v1/deltas/payments | V1 List Payment Changes
*PaymentAuditServiceDeprecatedAPI* | [**ListPaymentsAuditV3**](docs/PaymentAuditServiceDeprecatedAPI.md#listpaymentsauditv3) | **Get** /v3/paymentaudit/payments | V3 Get List of Payments
*PayorHierarchyAPI* | [**PayorLinksV1**](docs/PayorHierarchyAPI.md#payorlinksv1) | **Get** /v1/payorLinks | List Payor Links
*PayorsAPI* | [**GetPayorByIdV2**](docs/PayorsAPI.md#getpayorbyidv2) | **Get** /v2/payors/{payorId} | Get Payor
*PayorsAPI* | [**PayorAddPayorLogoV1**](docs/PayorsAPI.md#payoraddpayorlogov1) | **Post** /v1/payors/{payorId}/branding/logos | Add Logo
*PayorsAPI* | [**PayorCreateApiKeyV1**](docs/PayorsAPI.md#payorcreateapikeyv1) | **Post** /v1/payors/{payorId}/applications/{applicationId}/keys | Create API Key
*PayorsAPI* | [**PayorCreateApplicationV1**](docs/PayorsAPI.md#payorcreateapplicationv1) | **Post** /v1/payors/{payorId}/applications | Create Application
*PayorsAPI* | [**PayorEmailOptOut**](docs/PayorsAPI.md#payoremailoptout) | **Post** /v1/payors/{payorId}/reminderEmailsUpdate | Reminder Email Opt-Out
*PayorsAPI* | [**PayorGetBranding**](docs/PayorsAPI.md#payorgetbranding) | **Get** /v1/payors/{payorId}/branding | Get Branding
*PayorsPrivateAPI* | [**CreatePayorLinks**](docs/PayorsPrivateAPI.md#createpayorlinks) | **Post** /v1/payorLinks | Create a Payor Link
*PayoutsAPI* | [**CreateQuoteForPayoutV3**](docs/PayoutsAPI.md#createquoteforpayoutv3) | **Post** /v3/payouts/{payoutId}/quote | Create a quote for the payout
*PayoutsAPI* | [**DeschedulePayout**](docs/PayoutsAPI.md#deschedulepayout) | **Delete** /v3/payouts/{payoutId}/schedule | Deschedule a payout
*PayoutsAPI* | [**GetPaymentsForPayoutV3**](docs/PayoutsAPI.md#getpaymentsforpayoutv3) | **Get** /v3/payouts/{payoutId}/payments | Retrieve payments for a payout
*PayoutsAPI* | [**GetPayoutSummaryV3**](docs/PayoutsAPI.md#getpayoutsummaryv3) | **Get** /v3/payouts/{payoutId} | Get Payout Summary
*PayoutsAPI* | [**InstructPayoutV3**](docs/PayoutsAPI.md#instructpayoutv3) | **Post** /v3/payouts/{payoutId} | Instruct Payout
*PayoutsAPI* | [**ScheduleForPayout**](docs/PayoutsAPI.md#scheduleforpayout) | **Post** /v3/payouts/{payoutId}/schedule | Schedule a payout
*PayoutsAPI* | [**SubmitPayoutV3**](docs/PayoutsAPI.md#submitpayoutv3) | **Post** /v3/payouts | Submit Payout
*PayoutsAPI* | [**WithdrawPayment**](docs/PayoutsAPI.md#withdrawpayment) | **Post** /v1/payments/{paymentId}/withdraw | Withdraw a Payment
*PayoutsAPI* | [**WithdrawPayoutV3**](docs/PayoutsAPI.md#withdrawpayoutv3) | **Delete** /v3/payouts/{payoutId} | Withdraw Payout
*SourceAccountsAPI* | [**GetSourceAccountV2**](docs/SourceAccountsAPI.md#getsourceaccountv2) | **Get** /v2/sourceAccounts/{sourceAccountId} | Get Source Account
*SourceAccountsAPI* | [**GetSourceAccountV3**](docs/SourceAccountsAPI.md#getsourceaccountv3) | **Get** /v3/sourceAccounts/{sourceAccountId} | Get details about given source account.
*SourceAccountsAPI* | [**GetSourceAccountsV2**](docs/SourceAccountsAPI.md#getsourceaccountsv2) | **Get** /v2/sourceAccounts | Get list of source accounts
*SourceAccountsAPI* | [**GetSourceAccountsV3**](docs/SourceAccountsAPI.md#getsourceaccountsv3) | **Get** /v3/sourceAccounts | Get list of source accounts
*SourceAccountsAPI* | [**SetNotificationsRequest**](docs/SourceAccountsAPI.md#setnotificationsrequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/notifications | Set notifications
*SourceAccountsAPI* | [**SetNotificationsRequestV3**](docs/SourceAccountsAPI.md#setnotificationsrequestv3) | **Post** /v3/sourceAccounts/{sourceAccountId}/notifications | Set notifications
*SourceAccountsAPI* | [**TransferFundsV2**](docs/SourceAccountsAPI.md#transferfundsv2) | **Post** /v2/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts
*SourceAccountsAPI* | [**TransferFundsV3**](docs/SourceAccountsAPI.md#transferfundsv3) | **Post** /v3/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts
*TokensAPI* | [**ResendToken**](docs/TokensAPI.md#resendtoken) | **Post** /v2/users/{userId}/tokens | Resend a token
*TransactionsAPI* | [**CreateTransactionV1**](docs/TransactionsAPI.md#createtransactionv1) | **Post** /v1/transactions | Create a Transaction
*TransactionsAPI* | [**GetTransactionByIdV1**](docs/TransactionsAPI.md#gettransactionbyidv1) | **Get** /v1/transactions/{transactionId} | Get Transaction
*TransactionsAPI* | [**GetTransactions**](docs/TransactionsAPI.md#gettransactions) | **Get** /v1/transactions | Get Transactions
*UsersAPI* | [**DeleteUserByIdV2**](docs/UsersAPI.md#deleteuserbyidv2) | **Delete** /v2/users/{userId} | Delete a User
*UsersAPI* | [**DisableUserV2**](docs/UsersAPI.md#disableuserv2) | **Post** /v2/users/{userId}/disable | Disable a User
*UsersAPI* | [**EnableUserV2**](docs/UsersAPI.md#enableuserv2) | **Post** /v2/users/{userId}/enable | Enable a User
*UsersAPI* | [**GetSelf**](docs/UsersAPI.md#getself) | **Get** /v2/users/self | Get Self
*UsersAPI* | [**GetUserByIdV2**](docs/UsersAPI.md#getuserbyidv2) | **Get** /v2/users/{userId} | Get User
*UsersAPI* | [**InviteUser**](docs/UsersAPI.md#inviteuser) | **Post** /v2/users/invite | Invite a User
*UsersAPI* | [**ListUsers**](docs/UsersAPI.md#listusers) | **Get** /v2/users | List Users
*UsersAPI* | [**RegisterSms**](docs/UsersAPI.md#registersms) | **Post** /v2/users/registration/sms | Register SMS Number
*UsersAPI* | [**ResendToken**](docs/UsersAPI.md#resendtoken) | **Post** /v2/users/{userId}/tokens | Resend a token
*UsersAPI* | [**RoleUpdate**](docs/UsersAPI.md#roleupdate) | **Post** /v2/users/{userId}/roleUpdate | Update User Role
*UsersAPI* | [**UnlockUserV2**](docs/UsersAPI.md#unlockuserv2) | **Post** /v2/users/{userId}/unlock | Unlock a User
*UsersAPI* | [**UnregisterMFA**](docs/UsersAPI.md#unregistermfa) | **Post** /v2/users/{userId}/mfa/unregister | Unregister MFA for the user
*UsersAPI* | [**UnregisterMFAForSelf**](docs/UsersAPI.md#unregistermfaforself) | **Post** /v2/users/self/mfa/unregister | Unregister MFA for Self
*UsersAPI* | [**UpdatePasswordSelf**](docs/UsersAPI.md#updatepasswordself) | **Post** /v2/users/self/password | Update Password for self
*UsersAPI* | [**UserDetailsUpdate**](docs/UsersAPI.md#userdetailsupdate) | **Post** /v2/users/{userId}/userDetailsUpdate | Update User Details
*UsersAPI* | [**UserDetailsUpdateForSelf**](docs/UsersAPI.md#userdetailsupdateforself) | **Post** /v2/users/self/userDetailsUpdate | Update User Details for self
*UsersAPI* | [**ValidatePasswordSelf**](docs/UsersAPI.md#validatepasswordself) | **Post** /v2/users/self/password/validate | Validate the proposed password
*WebhooksAPI* | [**CreateWebhookV1**](docs/WebhooksAPI.md#createwebhookv1) | **Post** /v1/webhooks | Create Webhook
*WebhooksAPI* | [**GetWebhookV1**](docs/WebhooksAPI.md#getwebhookv1) | **Get** /v1/webhooks/{webhookId} | Get details about the given webhook.
*WebhooksAPI* | [**ListWebhooksV1**](docs/WebhooksAPI.md#listwebhooksv1) | **Get** /v1/webhooks | List the details about the webhooks for the given payor.
*WebhooksAPI* | [**PingWebhookV1**](docs/WebhooksAPI.md#pingwebhookv1) | **Post** /v1/webhooks/{webhookId}/ping | 
*WebhooksAPI* | [**UpdateWebhookV1**](docs/WebhooksAPI.md#updatewebhookv1) | **Post** /v1/webhooks/{webhookId} | Update Webhook


## Documentation For Models

 - [AcceptedPaymentV3](docs/AcceptedPaymentV3.md)
 - [AccessTokenResponse](docs/AccessTokenResponse.md)
 - [AccessTokenValidationRequest](docs/AccessTokenValidationRequest.md)
 - [AddressV4](docs/AddressV4.md)
 - [AuthResponse](docs/AuthResponse.md)
 - [AutoTopUpConfigV2](docs/AutoTopUpConfigV2.md)
 - [AutoTopUpConfigV3](docs/AutoTopUpConfigV3.md)
 - [Category](docs/Category.md)
 - [ChallengeV3](docs/ChallengeV3.md)
 - [ChallengeV4](docs/ChallengeV4.md)
 - [CommonLinkObject](docs/CommonLinkObject.md)
 - [CommonPageObject](docs/CommonPageObject.md)
 - [CompanyV3](docs/CompanyV3.md)
 - [CompanyV4](docs/CompanyV4.md)
 - [CreateFundingAccountRequestV2](docs/CreateFundingAccountRequestV2.md)
 - [CreateIndividualV3](docs/CreateIndividualV3.md)
 - [CreateIndividualV3Name](docs/CreateIndividualV3Name.md)
 - [CreateIndividualV4](docs/CreateIndividualV4.md)
 - [CreatePayeeAddressV3](docs/CreatePayeeAddressV3.md)
 - [CreatePayeeAddressV4](docs/CreatePayeeAddressV4.md)
 - [CreatePayeeV3](docs/CreatePayeeV3.md)
 - [CreatePayeeV4](docs/CreatePayeeV4.md)
 - [CreatePayeesCSVRequestV3](docs/CreatePayeesCSVRequestV3.md)
 - [CreatePayeesCSVRequestV4](docs/CreatePayeesCSVRequestV4.md)
 - [CreatePayeesCSVResponseV3](docs/CreatePayeesCSVResponseV3.md)
 - [CreatePayeesCSVResponseV3RejectedCsvRows](docs/CreatePayeesCSVResponseV3RejectedCsvRows.md)
 - [CreatePayeesCSVResponseV4](docs/CreatePayeesCSVResponseV4.md)
 - [CreatePayeesRequestV3](docs/CreatePayeesRequestV3.md)
 - [CreatePayeesRequestV4](docs/CreatePayeesRequestV4.md)
 - [CreatePaymentChannelRequestV4](docs/CreatePaymentChannelRequestV4.md)
 - [CreatePaymentChannelV3](docs/CreatePaymentChannelV3.md)
 - [CreatePaymentChannelV4](docs/CreatePaymentChannelV4.md)
 - [CreatePayorLinkRequest](docs/CreatePayorLinkRequest.md)
 - [CreatePayoutRequestV3](docs/CreatePayoutRequestV3.md)
 - [CreateTransactionRequest](docs/CreateTransactionRequest.md)
 - [CreateTransactionResponse](docs/CreateTransactionResponse.md)
 - [CreateWebhookRequest](docs/CreateWebhookRequest.md)
 - [DebitEvent](docs/DebitEvent.md)
 - [DebitEventAllOf](docs/DebitEventAllOf.md)
 - [DebitStatusChanged](docs/DebitStatusChanged.md)
 - [DebitStatusChangedAllOf](docs/DebitStatusChangedAllOf.md)
 - [Error](docs/Error.md)
 - [ErrorData](docs/ErrorData.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [FailedPayeeV3](docs/FailedPayeeV3.md)
 - [FailedPayeeV4](docs/FailedPayeeV4.md)
 - [FailedSubmissionV3](docs/FailedSubmissionV3.md)
 - [FailedSubmissionV4](docs/FailedSubmissionV4.md)
 - [FundingAccountResponseV2](docs/FundingAccountResponseV2.md)
 - [FundingAudit](docs/FundingAudit.md)
 - [FundingEvent](docs/FundingEvent.md)
 - [FundingEvent2](docs/FundingEvent2.md)
 - [FundingPayorStatusAuditResponse](docs/FundingPayorStatusAuditResponse.md)
 - [FundingRequestV2](docs/FundingRequestV2.md)
 - [FundingRequestV3](docs/FundingRequestV3.md)
 - [FundingResponse](docs/FundingResponse.md)
 - [FxSummary](docs/FxSummary.md)
 - [FxSummaryV3](docs/FxSummaryV3.md)
 - [GetFundingsResponse](docs/GetFundingsResponse.md)
 - [GetFundingsResponseLinks](docs/GetFundingsResponseLinks.md)
 - [GetPayeeListResponseCompanyV3](docs/GetPayeeListResponseCompanyV3.md)
 - [GetPayeeListResponseCompanyV4](docs/GetPayeeListResponseCompanyV4.md)
 - [GetPayeeListResponseIndividualV3](docs/GetPayeeListResponseIndividualV3.md)
 - [GetPayeeListResponseIndividualV4](docs/GetPayeeListResponseIndividualV4.md)
 - [GetPayeeListResponseV3](docs/GetPayeeListResponseV3.md)
 - [GetPayeeListResponseV4](docs/GetPayeeListResponseV4.md)
 - [GetPaymentsForPayoutResponseV3](docs/GetPaymentsForPayoutResponseV3.md)
 - [GetPaymentsForPayoutResponseV3Page](docs/GetPaymentsForPayoutResponseV3Page.md)
 - [GetPaymentsForPayoutResponseV3Summary](docs/GetPaymentsForPayoutResponseV3Summary.md)
 - [GetPaymentsForPayoutResponseV4](docs/GetPaymentsForPayoutResponseV4.md)
 - [GetPaymentsForPayoutResponseV4Summary](docs/GetPaymentsForPayoutResponseV4Summary.md)
 - [GetPayoutStatistics](docs/GetPayoutStatistics.md)
 - [GetPayoutsResponse](docs/GetPayoutsResponse.md)
 - [GetPayoutsResponseV3](docs/GetPayoutsResponseV3.md)
 - [GetPayoutsResponseV3Links](docs/GetPayoutsResponseV3Links.md)
 - [GetPayoutsResponseV3Page](docs/GetPayoutsResponseV3Page.md)
 - [IndividualV3](docs/IndividualV3.md)
 - [IndividualV3Name](docs/IndividualV3Name.md)
 - [IndividualV4](docs/IndividualV4.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [InlineResponse401](docs/InlineResponse401.md)
 - [InlineResponse403](docs/InlineResponse403.md)
 - [InlineResponse404](docs/InlineResponse404.md)
 - [InlineResponse409](docs/InlineResponse409.md)
 - [InlineResponse412](docs/InlineResponse412.md)
 - [InstructPayoutRequestV3](docs/InstructPayoutRequestV3.md)
 - [InvitePayeeRequestV3](docs/InvitePayeeRequestV3.md)
 - [InvitePayeeRequestV4](docs/InvitePayeeRequestV4.md)
 - [InviteUserRequest](docs/InviteUserRequest.md)
 - [LinkForResponse](docs/LinkForResponse.md)
 - [ListFundingAccountsResponseV2](docs/ListFundingAccountsResponseV2.md)
 - [ListFundingAccountsResponseV2Page](docs/ListFundingAccountsResponseV2Page.md)
 - [ListPaymentsResponseV3](docs/ListPaymentsResponseV3.md)
 - [ListPaymentsResponseV3Page](docs/ListPaymentsResponseV3Page.md)
 - [ListPaymentsResponseV4](docs/ListPaymentsResponseV4.md)
 - [ListSourceAccountResponseV2](docs/ListSourceAccountResponseV2.md)
 - [ListSourceAccountResponseV2Links](docs/ListSourceAccountResponseV2Links.md)
 - [ListSourceAccountResponseV3](docs/ListSourceAccountResponseV3.md)
 - [ListSourceAccountResponseV3Links](docs/ListSourceAccountResponseV3Links.md)
 - [LocalisationDetails](docs/LocalisationDetails.md)
 - [MFADetails](docs/MFADetails.md)
 - [MFAType](docs/MFAType.md)
 - [NameV3](docs/NameV3.md)
 - [NameV4](docs/NameV4.md)
 - [Notification](docs/Notification.md)
 - [NotificationSource](docs/NotificationSource.md)
 - [NotificationsV2](docs/NotificationsV2.md)
 - [NotificationsV3](docs/NotificationsV3.md)
 - [OnboardingStatusChanged](docs/OnboardingStatusChanged.md)
 - [PageForResponse](docs/PageForResponse.md)
 - [PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse](docs/PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse.md)
 - [PageResourceTransactions](docs/PageResourceTransactions.md)
 - [PagedPayeeInvitationStatusResponseV3](docs/PagedPayeeInvitationStatusResponseV3.md)
 - [PagedPayeeInvitationStatusResponseV3Page](docs/PagedPayeeInvitationStatusResponseV3Page.md)
 - [PagedPayeeInvitationStatusResponseV4](docs/PagedPayeeInvitationStatusResponseV4.md)
 - [PagedPayeeResponseV3](docs/PagedPayeeResponseV3.md)
 - [PagedPayeeResponseV3Links](docs/PagedPayeeResponseV3Links.md)
 - [PagedPayeeResponseV3Page](docs/PagedPayeeResponseV3Page.md)
 - [PagedPayeeResponseV3Summary](docs/PagedPayeeResponseV3Summary.md)
 - [PagedPayeeResponseV4](docs/PagedPayeeResponseV4.md)
 - [PagedPaymentsResponseV3](docs/PagedPaymentsResponseV3.md)
 - [PagedUserResponse](docs/PagedUserResponse.md)
 - [PagedUserResponseLinks](docs/PagedUserResponseLinks.md)
 - [PagedUserResponsePage](docs/PagedUserResponsePage.md)
 - [PasswordRequest](docs/PasswordRequest.md)
 - [PayableIssueV3](docs/PayableIssueV3.md)
 - [PayableIssueV4](docs/PayableIssueV4.md)
 - [PayableStatusChanged](docs/PayableStatusChanged.md)
 - [PayeeAddressV3](docs/PayeeAddressV3.md)
 - [PayeeAddressV4](docs/PayeeAddressV4.md)
 - [PayeeDeltaResponseV3](docs/PayeeDeltaResponseV3.md)
 - [PayeeDeltaResponseV3Links](docs/PayeeDeltaResponseV3Links.md)
 - [PayeeDeltaResponseV3Page](docs/PayeeDeltaResponseV3Page.md)
 - [PayeeDeltaResponseV4](docs/PayeeDeltaResponseV4.md)
 - [PayeeDeltaResponseV4Links](docs/PayeeDeltaResponseV4Links.md)
 - [PayeeDeltaV3](docs/PayeeDeltaV3.md)
 - [PayeeDeltaV4](docs/PayeeDeltaV4.md)
 - [PayeeDetailResponseV3](docs/PayeeDetailResponseV3.md)
 - [PayeeDetailResponseV4](docs/PayeeDetailResponseV4.md)
 - [PayeeDetailsChanged](docs/PayeeDetailsChanged.md)
 - [PayeeEvent](docs/PayeeEvent.md)
 - [PayeeEventAllOf](docs/PayeeEventAllOf.md)
 - [PayeeEventAllOfReasons](docs/PayeeEventAllOfReasons.md)
 - [PayeeInvitationStatusResponseV3](docs/PayeeInvitationStatusResponseV3.md)
 - [PayeeInvitationStatusResponseV4](docs/PayeeInvitationStatusResponseV4.md)
 - [PayeePayorRefV3](docs/PayeePayorRefV3.md)
 - [PayeePayorRefV4](docs/PayeePayorRefV4.md)
 - [PayeeType](docs/PayeeType.md)
 - [PayeeTypeEnum](docs/PayeeTypeEnum.md)
 - [PayeeUserSelfUpdateRequest](docs/PayeeUserSelfUpdateRequest.md)
 - [PaymentChannelCountry](docs/PaymentChannelCountry.md)
 - [PaymentChannelOrderRequestV4](docs/PaymentChannelOrderRequestV4.md)
 - [PaymentChannelResponseV4](docs/PaymentChannelResponseV4.md)
 - [PaymentChannelRule](docs/PaymentChannelRule.md)
 - [PaymentChannelRulesResponse](docs/PaymentChannelRulesResponse.md)
 - [PaymentChannelSummaryV4](docs/PaymentChannelSummaryV4.md)
 - [PaymentChannelsResponseV4](docs/PaymentChannelsResponseV4.md)
 - [PaymentDelta](docs/PaymentDelta.md)
 - [PaymentDeltaResponse](docs/PaymentDeltaResponse.md)
 - [PaymentDeltaResponseV1](docs/PaymentDeltaResponseV1.md)
 - [PaymentDeltaV1](docs/PaymentDeltaV1.md)
 - [PaymentEvent](docs/PaymentEvent.md)
 - [PaymentEventAllOf](docs/PaymentEventAllOf.md)
 - [PaymentEventResponse](docs/PaymentEventResponse.md)
 - [PaymentEventResponseV3](docs/PaymentEventResponseV3.md)
 - [PaymentInstructionV3](docs/PaymentInstructionV3.md)
 - [PaymentRejectedOrReturned](docs/PaymentRejectedOrReturned.md)
 - [PaymentRejectedOrReturnedAllOf](docs/PaymentRejectedOrReturnedAllOf.md)
 - [PaymentResponseV3](docs/PaymentResponseV3.md)
 - [PaymentResponseV4](docs/PaymentResponseV4.md)
 - [PaymentResponseV4Payout](docs/PaymentResponseV4Payout.md)
 - [PaymentStatusChanged](docs/PaymentStatusChanged.md)
 - [PaymentStatusChangedAllOf](docs/PaymentStatusChangedAllOf.md)
 - [PaymentV3](docs/PaymentV3.md)
 - [PayorAddressV2](docs/PayorAddressV2.md)
 - [PayorAmlTransaction](docs/PayorAmlTransaction.md)
 - [PayorAmlTransactionV3](docs/PayorAmlTransactionV3.md)
 - [PayorBrandingResponse](docs/PayorBrandingResponse.md)
 - [PayorCreateApiKeyRequest](docs/PayorCreateApiKeyRequest.md)
 - [PayorCreateApiKeyResponse](docs/PayorCreateApiKeyResponse.md)
 - [PayorCreateApplicationRequest](docs/PayorCreateApplicationRequest.md)
 - [PayorEmailOptOutRequest](docs/PayorEmailOptOutRequest.md)
 - [PayorFundingDetected](docs/PayorFundingDetected.md)
 - [PayorFundingDetectedAllOf](docs/PayorFundingDetectedAllOf.md)
 - [PayorLinksResponse](docs/PayorLinksResponse.md)
 - [PayorLinksResponseLinks](docs/PayorLinksResponseLinks.md)
 - [PayorLinksResponsePayors](docs/PayorLinksResponsePayors.md)
 - [PayorToPaymentChannelMappingV4](docs/PayorToPaymentChannelMappingV4.md)
 - [PayorV2](docs/PayorV2.md)
 - [PayoutCompanyV3](docs/PayoutCompanyV3.md)
 - [PayoutIndividualV3](docs/PayoutIndividualV3.md)
 - [PayoutNameV3](docs/PayoutNameV3.md)
 - [PayoutPayeeV3](docs/PayoutPayeeV3.md)
 - [PayoutPayor](docs/PayoutPayor.md)
 - [PayoutPayorIds](docs/PayoutPayorIds.md)
 - [PayoutPrincipal](docs/PayoutPrincipal.md)
 - [PayoutSchedule](docs/PayoutSchedule.md)
 - [PayoutScheduleV3](docs/PayoutScheduleV3.md)
 - [PayoutSummaryAudit](docs/PayoutSummaryAudit.md)
 - [PayoutSummaryAuditV3](docs/PayoutSummaryAuditV3.md)
 - [PayoutSummaryResponseV3](docs/PayoutSummaryResponseV3.md)
 - [Ping](docs/Ping.md)
 - [PingResponse](docs/PingResponse.md)
 - [PostInstructFxInfo](docs/PostInstructFxInfo.md)
 - [QueryBatchResponseV3](docs/QueryBatchResponseV3.md)
 - [QueryBatchResponseV4](docs/QueryBatchResponseV4.md)
 - [QuoteFxSummaryV3](docs/QuoteFxSummaryV3.md)
 - [QuoteResponseV3](docs/QuoteResponseV3.md)
 - [RegionV2](docs/RegionV2.md)
 - [RegisterSmsRequest](docs/RegisterSmsRequest.md)
 - [RejectedPaymentV3](docs/RejectedPaymentV3.md)
 - [ResendTokenRequest](docs/ResendTokenRequest.md)
 - [ResetPasswordRequest](docs/ResetPasswordRequest.md)
 - [Role](docs/Role.md)
 - [RoleUpdateRequest](docs/RoleUpdateRequest.md)
 - [SchedulePayoutRequestV3](docs/SchedulePayoutRequestV3.md)
 - [SelfMFATypeUnregisterRequest](docs/SelfMFATypeUnregisterRequest.md)
 - [SelfUpdatePasswordRequest](docs/SelfUpdatePasswordRequest.md)
 - [SetNotificationsRequest](docs/SetNotificationsRequest.md)
 - [SetNotificationsRequest2](docs/SetNotificationsRequest2.md)
 - [SourceAccountResponseV2](docs/SourceAccountResponseV2.md)
 - [SourceAccountResponseV3](docs/SourceAccountResponseV3.md)
 - [SourceAccountSummary](docs/SourceAccountSummary.md)
 - [SourceAccountSummaryV3](docs/SourceAccountSummaryV3.md)
 - [SourceAccountV3](docs/SourceAccountV3.md)
 - [SourceEvent](docs/SourceEvent.md)
 - [SupportedCountriesResponse](docs/SupportedCountriesResponse.md)
 - [SupportedCountriesResponseV2](docs/SupportedCountriesResponseV2.md)
 - [SupportedCountry](docs/SupportedCountry.md)
 - [SupportedCountryV2](docs/SupportedCountryV2.md)
 - [SupportedCurrencyResponseV2](docs/SupportedCurrencyResponseV2.md)
 - [SupportedCurrencyV2](docs/SupportedCurrencyV2.md)
 - [TransactionResponse](docs/TransactionResponse.md)
 - [TransferRequestV2](docs/TransferRequestV2.md)
 - [TransferRequestV3](docs/TransferRequestV3.md)
 - [UnregisterMFARequest](docs/UnregisterMFARequest.md)
 - [UpdatePayeeDetailsRequestV3](docs/UpdatePayeeDetailsRequestV3.md)
 - [UpdatePayeeDetailsRequestV4](docs/UpdatePayeeDetailsRequestV4.md)
 - [UpdatePaymentChannelRequestV4](docs/UpdatePaymentChannelRequestV4.md)
 - [UpdateRemoteIdRequestV3](docs/UpdateRemoteIdRequestV3.md)
 - [UpdateRemoteIdRequestV4](docs/UpdateRemoteIdRequestV4.md)
 - [UpdateWebhookRequest](docs/UpdateWebhookRequest.md)
 - [UserDetailsUpdateRequest](docs/UserDetailsUpdateRequest.md)
 - [UserInfo](docs/UserInfo.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserStatus](docs/UserStatus.md)
 - [UserType](docs/UserType.md)
 - [UserType2](docs/UserType2.md)
 - [ValidatePasswordResponse](docs/ValidatePasswordResponse.md)
 - [WebhookResponse](docs/WebhookResponse.md)
 - [WebhooksResponse](docs/WebhooksResponse.md)
 - [WithdrawPaymentRequest](docs/WithdrawPaymentRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### OAuth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - ** **: Scopes not required

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

### oAuthVeloBackOffice


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - ** **: Scopes not required

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



