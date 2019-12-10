# Go API client for velopayments
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![npm version](https://badge.fury.io/go/github.com%2Fvelopaymentsapi%2Fvelo-go.svg)](https://badge.fury.io/go/github.com%2Fvelopaymentsapi%2Fvelo-go) [![CircleCI](https://circleci.com/gh/velopaymentsapi/velo-go.svg?style=svg)](https://circleci.com/gh/velopaymentsapi/velo-go)\
## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.17.8
- Package version: 2.17.8
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./velopayments"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.sandbox.velopayments.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthApi* | [**VeloAuth**](docs/AuthApi.md#veloauth) | **Post** /v1/authenticate | Authentication endpoint
*CountriesApi* | [**ListSupportedCountries**](docs/CountriesApi.md#listsupportedcountries) | **Get** /v1/supportedCountries | List Supported Countries
*CountriesApi* | [**V1PaymentChannelRulesGet**](docs/CountriesApi.md#v1paymentchannelrulesget) | **Get** /v1/paymentChannelRules | List Payment Channel Country Rules
*CurrenciesApi* | [**ListSupportedCurrencies**](docs/CurrenciesApi.md#listsupportedcurrencies) | **Get** /v2/currencies | List Supported Currencies
*FundingManagerApi* | [**CreateAchFundingRequest**](docs/FundingManagerApi.md#createachfundingrequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/achFundingRequest | Create Funding Request
*FundingManagerApi* | [**CreateFundingRequest**](docs/FundingManagerApi.md#createfundingrequest) | **Post** /v2/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
*FundingManagerApi* | [**GetFundings**](docs/FundingManagerApi.md#getfundings) | **Get** /v1/paymentaudit/fundings | Get Fundings for Payor
*FundingManagerApi* | [**GetSourceAccount**](docs/FundingManagerApi.md#getsourceaccount) | **Get** /v1/sourceAccounts/{sourceAccountId} | Get details about given source account.
*FundingManagerApi* | [**GetSourceAccountV2**](docs/FundingManagerApi.md#getsourceaccountv2) | **Get** /v2/sourceAccounts/{sourceAccountId} | Get details about given source account.
*FundingManagerApi* | [**GetSourceAccounts**](docs/FundingManagerApi.md#getsourceaccounts) | **Get** /v1/sourceAccounts | Get list of source accounts
*FundingManagerApi* | [**GetSourceAccountsV2**](docs/FundingManagerApi.md#getsourceaccountsv2) | **Get** /v2/sourceAccounts | Get list of source accounts
*FundingManagerApi* | [**ListFundingAuditDeltas**](docs/FundingManagerApi.md#listfundingauditdeltas) | **Get** /v1/deltas/fundings | List Funding changes
*FundingManagerApi* | [**SetNotificationsRequest**](docs/FundingManagerApi.md#setnotificationsrequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/notifications | Set notifications
*GetPayoutApi* | [**V3PayoutsPayoutIdGet**](docs/GetPayoutApi.md#v3payoutspayoutidget) | **Get** /v3/payouts/{payoutId} | Get Payout Summary
*InstructPayoutApi* | [**V3PayoutsPayoutIdPost**](docs/InstructPayoutApi.md#v3payoutspayoutidpost) | **Post** /v3/payouts/{payoutId} | Instruct Payout
*PayeeInvitationApi* | [**GetPayeesInvitationStatus**](docs/PayeeInvitationApi.md#getpayeesinvitationstatus) | **Get** /v1/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationApi* | [**GetPayeesInvitationStatusV2**](docs/PayeeInvitationApi.md#getpayeesinvitationstatusv2) | **Get** /v2/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationApi* | [**ResendPayeeInvite**](docs/PayeeInvitationApi.md#resendpayeeinvite) | **Post** /v1/payees/{payeeId}/invite | Resend Payee Invite
*PayeeInvitationApi* | [**V2CreatePayee**](docs/PayeeInvitationApi.md#v2createpayee) | **Post** /v2/payees | Intiate Payee Creation V2
*PayeeInvitationApi* | [**V2QueryBatchStatus**](docs/PayeeInvitationApi.md#v2querybatchstatus) | **Get** /v2/payees/batch/{batchId} | Query Batch Status
*PayeeInvitationApi* | [**V3CreatePayee**](docs/PayeeInvitationApi.md#v3createpayee) | **Post** /v3/payees | Intiate Payee Creation V3
*PayeeInvitationApi* | [**V3QueryBatchStatus**](docs/PayeeInvitationApi.md#v3querybatchstatus) | **Get** /v3/payees/batch/{batchId} | Query Batch Status
*PayeesApi* | [**DeletePayeeById**](docs/PayeesApi.md#deletepayeebyid) | **Delete** /v1/payees/{payeeId} | Delete Payee by Id
*PayeesApi* | [**GetPayeeById**](docs/PayeesApi.md#getpayeebyid) | **Get** /v1/payees/{payeeId} | Get Payee by Id
*PayeesApi* | [**ListPayeeChanges**](docs/PayeesApi.md#listpayeechanges) | **Get** /v1/deltas/payees | List Payee Changes
*PayeesApi* | [**ListPayees**](docs/PayeesApi.md#listpayees) | **Get** /v1/payees | List Payees
*PayeesApi* | [**ListPayeesV3**](docs/PayeesApi.md#listpayeesv3) | **Get** /v3/payees | List Payees
*PayeesApi* | [**V1PayeesPayeeIdRemoteIdUpdatePost**](docs/PayeesApi.md#v1payeespayeeidremoteidupdatepost) | **Post** /v1/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
*PaymentAuditServiceApi* | [**ExportTransactionsCSV**](docs/PaymentAuditServiceApi.md#exporttransactionscsv) | **Get** /v4/paymentaudit/transactions | Export Transactions
*PaymentAuditServiceApi* | [**GetFundings**](docs/PaymentAuditServiceApi.md#getfundings) | **Get** /v1/paymentaudit/fundings | Get Fundings for Payor
*PaymentAuditServiceApi* | [**GetPaymentDetails**](docs/PaymentAuditServiceApi.md#getpaymentdetails) | **Get** /v3/paymentaudit/payments/{paymentId} | Get Payment
*PaymentAuditServiceApi* | [**GetPaymentDetailsV4**](docs/PaymentAuditServiceApi.md#getpaymentdetailsv4) | **Get** /v4/paymentaudit/payments/{paymentId} | Get Payment
*PaymentAuditServiceApi* | [**GetPaymentsForPayout**](docs/PaymentAuditServiceApi.md#getpaymentsforpayout) | **Get** /v3/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PaymentAuditServiceApi* | [**GetPaymentsForPayoutV4**](docs/PaymentAuditServiceApi.md#getpaymentsforpayoutv4) | **Get** /v4/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PaymentAuditServiceApi* | [**GetPayoutsForPayor**](docs/PaymentAuditServiceApi.md#getpayoutsforpayor) | **Get** /v3/paymentaudit/payouts | Get Payouts for Payor
*PaymentAuditServiceApi* | [**GetPayoutsForPayorV4**](docs/PaymentAuditServiceApi.md#getpayoutsforpayorv4) | **Get** /v4/paymentaudit/payouts | Get Payouts for Payor
*PaymentAuditServiceApi* | [**ListPaymentChanges**](docs/PaymentAuditServiceApi.md#listpaymentchanges) | **Get** /v1/deltas/payments | List Payment Changes
*PaymentAuditServiceApi* | [**ListPaymentsAudit**](docs/PaymentAuditServiceApi.md#listpaymentsaudit) | **Get** /v3/paymentaudit/payments | Get List of Payments
*PaymentAuditServiceApi* | [**ListPaymentsAuditV4**](docs/PaymentAuditServiceApi.md#listpaymentsauditv4) | **Get** /v4/paymentaudit/payments | Get List of Payments
*PayorApplicationsApi* | [**PayorCreateApiKeyRequest**](docs/PayorApplicationsApi.md#payorcreateapikeyrequest) | **Post** /v1/payors/{payorId}/applications/{applicationId}/keys | Create API Key
*PayorApplicationsApi* | [**PayorCreateApplicationRequest**](docs/PayorApplicationsApi.md#payorcreateapplicationrequest) | **Post** /v1/payors/{payorId}/applications | Create Application
*PayorsApi* | [**GetPayorById**](docs/PayorsApi.md#getpayorbyid) | **Get** /v1/payors/{payorId} | Get Payor
*PayorsApi* | [**GetPayorByIdV2**](docs/PayorsApi.md#getpayorbyidv2) | **Get** /v2/payors/{payorId} | Get Payor
*PayorsApi* | [**PayorAddPayorLogo**](docs/PayorsApi.md#payoraddpayorlogo) | **Post** /v1/payors/{payorId}/branding/logos | Add Logo
*PayorsApi* | [**PayorEmailOptOut**](docs/PayorsApi.md#payoremailoptout) | **Post** /v1/payors/{payorId}/reminderEmailsUpdate | Reminder Email Opt-Out
*PayorsApi* | [**PayorGetBranding**](docs/PayorsApi.md#payorgetbranding) | **Get** /v1/payors/{payorId}/branding | Get Branding
*PayorsApi* | [**PayorLinks**](docs/PayorsApi.md#payorlinks) | **Get** /v1/payorLinks | List Payor Links
*PayoutHistoryApi* | [**GetPaymentsForPayout**](docs/PayoutHistoryApi.md#getpaymentsforpayout) | **Get** /v3/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PayoutHistoryApi* | [**GetPaymentsForPayoutV4**](docs/PayoutHistoryApi.md#getpaymentsforpayoutv4) | **Get** /v4/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PayoutHistoryApi* | [**GetPayoutStats**](docs/PayoutHistoryApi.md#getpayoutstats) | **Get** /v1/paymentaudit/payoutStatistics | Get Payout Statistics
*QuotePayoutApi* | [**V3PayoutsPayoutIdQuotePost**](docs/QuotePayoutApi.md#v3payoutspayoutidquotepost) | **Post** /v3/payouts/{payoutId}/quote | Create a quote for the payout
*SubmitPayoutApi* | [**SubmitPayout**](docs/SubmitPayoutApi.md#submitpayout) | **Post** /v3/payouts | Submit Payout
*WithdrawPayoutApi* | [**V3PayoutsPayoutIdDelete**](docs/WithdrawPayoutApi.md#v3payoutspayoutiddelete) | **Delete** /v3/payouts/{payoutId} | Withdraw Payout


## Documentation For Models

 - [AuthResponse](docs/AuthResponse.md)
 - [AutoTopUpConfig](docs/AutoTopUpConfig.md)
 - [Challenge](docs/Challenge.md)
 - [CompanyResponse](docs/CompanyResponse.md)
 - [CompanyV1](docs/CompanyV1.md)
 - [CreateIndividual](docs/CreateIndividual.md)
 - [CreateIndividual2](docs/CreateIndividual2.md)
 - [CreatePayee](docs/CreatePayee.md)
 - [CreatePayee2](docs/CreatePayee2.md)
 - [CreatePayeeAddress](docs/CreatePayeeAddress.md)
 - [CreatePayeeAddress2](docs/CreatePayeeAddress2.md)
 - [CreatePayeesCsvRequest](docs/CreatePayeesCsvRequest.md)
 - [CreatePayeesCsvRequest2](docs/CreatePayeesCsvRequest2.md)
 - [CreatePayeesCsvResponse](docs/CreatePayeesCsvResponse.md)
 - [CreatePayeesCsvResponse2](docs/CreatePayeesCsvResponse2.md)
 - [CreatePayeesRequest](docs/CreatePayeesRequest.md)
 - [CreatePayeesRequest2](docs/CreatePayeesRequest2.md)
 - [CreatePaymentChannel](docs/CreatePaymentChannel.md)
 - [CreatePaymentChannel2](docs/CreatePaymentChannel2.md)
 - [CreatePayoutRequest](docs/CreatePayoutRequest.md)
 - [Error](docs/Error.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [FailedSubmission](docs/FailedSubmission.md)
 - [FundingAudit](docs/FundingAudit.md)
 - [FundingDelta](docs/FundingDelta.md)
 - [FundingDeltaResponse](docs/FundingDeltaResponse.md)
 - [FundingDeltaResponseLinks](docs/FundingDeltaResponseLinks.md)
 - [FundingEvent](docs/FundingEvent.md)
 - [FundingEventType](docs/FundingEventType.md)
 - [FundingRequestV1](docs/FundingRequestV1.md)
 - [FundingRequestV2](docs/FundingRequestV2.md)
 - [FxSummaryV3](docs/FxSummaryV3.md)
 - [FxSummaryV4](docs/FxSummaryV4.md)
 - [GetFundingsResponse](docs/GetFundingsResponse.md)
 - [GetFundingsResponseAllOf](docs/GetFundingsResponseAllOf.md)
 - [GetPaymentsForPayoutResponseV3](docs/GetPaymentsForPayoutResponseV3.md)
 - [GetPaymentsForPayoutResponseV3Page](docs/GetPaymentsForPayoutResponseV3Page.md)
 - [GetPaymentsForPayoutResponseV3Summary](docs/GetPaymentsForPayoutResponseV3Summary.md)
 - [GetPaymentsForPayoutResponseV4](docs/GetPaymentsForPayoutResponseV4.md)
 - [GetPaymentsForPayoutResponseV4Summary](docs/GetPaymentsForPayoutResponseV4Summary.md)
 - [GetPayoutStatistics](docs/GetPayoutStatistics.md)
 - [GetPayoutsResponseV3](docs/GetPayoutsResponseV3.md)
 - [GetPayoutsResponseV3Links](docs/GetPayoutsResponseV3Links.md)
 - [GetPayoutsResponseV3Page](docs/GetPayoutsResponseV3Page.md)
 - [GetPayoutsResponseV3Summary](docs/GetPayoutsResponseV3Summary.md)
 - [GetPayoutsResponseV4](docs/GetPayoutsResponseV4.md)
 - [IndividualResponse](docs/IndividualResponse.md)
 - [IndividualV1](docs/IndividualV1.md)
 - [IndividualV1Name](docs/IndividualV1Name.md)
 - [InvitationStatus](docs/InvitationStatus.md)
 - [InvitationStatusResponse](docs/InvitationStatusResponse.md)
 - [InvitePayeeRequest](docs/InvitePayeeRequest.md)
 - [Language](docs/Language.md)
 - [ListPaymentsResponse](docs/ListPaymentsResponse.md)
 - [ListPaymentsResponsePage](docs/ListPaymentsResponsePage.md)
 - [ListPaymentsResponseSummary](docs/ListPaymentsResponseSummary.md)
 - [ListSourceAccountResponse](docs/ListSourceAccountResponse.md)
 - [ListSourceAccountResponseLinks](docs/ListSourceAccountResponseLinks.md)
 - [ListSourceAccountResponsePage](docs/ListSourceAccountResponsePage.md)
 - [ListSourceAccountResponseV2](docs/ListSourceAccountResponseV2.md)
 - [ListSourceAccountResponseV2Page](docs/ListSourceAccountResponseV2Page.md)
 - [MarketingOptIn](docs/MarketingOptIn.md)
 - [Notifications](docs/Notifications.md)
 - [OfacStatus](docs/OfacStatus.md)
 - [OnboardedStatus](docs/OnboardedStatus.md)
 - [PagedPayeeInvitationStatusResponse](docs/PagedPayeeInvitationStatusResponse.md)
 - [PagedPayeeResponse](docs/PagedPayeeResponse.md)
 - [PagedPayeeResponse2](docs/PagedPayeeResponse2.md)
 - [PagedPayeeResponse2Summary](docs/PagedPayeeResponse2Summary.md)
 - [PagedPayeeResponseLinks](docs/PagedPayeeResponseLinks.md)
 - [PagedPayeeResponsePage](docs/PagedPayeeResponsePage.md)
 - [PagedPayeeResponseSummary](docs/PagedPayeeResponseSummary.md)
 - [PagedResponse](docs/PagedResponse.md)
 - [PagedResponsePage](docs/PagedResponsePage.md)
 - [Payee](docs/Payee.md)
 - [PayeeAddress](docs/PayeeAddress.md)
 - [PayeeDelta](docs/PayeeDelta.md)
 - [PayeeDeltaResponse](docs/PayeeDeltaResponse.md)
 - [PayeeDeltaResponseLinks](docs/PayeeDeltaResponseLinks.md)
 - [PayeeDeltaResponsePage](docs/PayeeDeltaResponsePage.md)
 - [PayeeInvitationStatus](docs/PayeeInvitationStatus.md)
 - [PayeeInvitationStatusResponse](docs/PayeeInvitationStatusResponse.md)
 - [PayeePaymentChannel](docs/PayeePaymentChannel.md)
 - [PayeePayorRef](docs/PayeePayorRef.md)
 - [PayeePayorRef2](docs/PayeePayorRef2.md)
 - [PayeeResponse](docs/PayeeResponse.md)
 - [PayeeResponse2](docs/PayeeResponse2.md)
 - [PayeeType](docs/PayeeType.md)
 - [PaymentAuditCurrencyV3](docs/PaymentAuditCurrencyV3.md)
 - [PaymentAuditCurrencyV4](docs/PaymentAuditCurrencyV4.md)
 - [PaymentChannelCountry](docs/PaymentChannelCountry.md)
 - [PaymentChannelRule](docs/PaymentChannelRule.md)
 - [PaymentChannelRulesResponse](docs/PaymentChannelRulesResponse.md)
 - [PaymentDelta](docs/PaymentDelta.md)
 - [PaymentDeltaResponse](docs/PaymentDeltaResponse.md)
 - [PaymentEventResponseV3](docs/PaymentEventResponseV3.md)
 - [PaymentEventResponseV4](docs/PaymentEventResponseV4.md)
 - [PaymentInstruction](docs/PaymentInstruction.md)
 - [PaymentResponseV3](docs/PaymentResponseV3.md)
 - [PaymentResponseV4](docs/PaymentResponseV4.md)
 - [PaymentResponseV4Payout](docs/PaymentResponseV4Payout.md)
 - [PaymentStatus](docs/PaymentStatus.md)
 - [PayorAddress](docs/PayorAddress.md)
 - [PayorAddressV2](docs/PayorAddressV2.md)
 - [PayorBrandingResponse](docs/PayorBrandingResponse.md)
 - [PayorCreateApiKeyRequest](docs/PayorCreateApiKeyRequest.md)
 - [PayorCreateApiKeyResponse](docs/PayorCreateApiKeyResponse.md)
 - [PayorCreateApplicationRequest](docs/PayorCreateApplicationRequest.md)
 - [PayorEmailOptOutRequest](docs/PayorEmailOptOutRequest.md)
 - [PayorLinksResponse](docs/PayorLinksResponse.md)
 - [PayorLinksResponseLinks](docs/PayorLinksResponseLinks.md)
 - [PayorLinksResponsePayors](docs/PayorLinksResponsePayors.md)
 - [PayorLogoRequest](docs/PayorLogoRequest.md)
 - [PayorV1](docs/PayorV1.md)
 - [PayorV2](docs/PayorV2.md)
 - [PayoutPayorV4](docs/PayoutPayorV4.md)
 - [PayoutPrincipalV4](docs/PayoutPrincipalV4.md)
 - [PayoutStatusV3](docs/PayoutStatusV3.md)
 - [PayoutStatusV4](docs/PayoutStatusV4.md)
 - [PayoutSummaryAuditV3](docs/PayoutSummaryAuditV3.md)
 - [PayoutSummaryAuditV4](docs/PayoutSummaryAuditV4.md)
 - [PayoutSummaryResponse](docs/PayoutSummaryResponse.md)
 - [PayoutTypeV4](docs/PayoutTypeV4.md)
 - [QueryBatchResponse](docs/QueryBatchResponse.md)
 - [QuoteFxSummary](docs/QuoteFxSummary.md)
 - [QuoteResponse](docs/QuoteResponse.md)
 - [RejectedPayment](docs/RejectedPayment.md)
 - [SetNotificationsRequest](docs/SetNotificationsRequest.md)
 - [SourceAccount](docs/SourceAccount.md)
 - [SourceAccountResponse](docs/SourceAccountResponse.md)
 - [SourceAccountResponseV2](docs/SourceAccountResponseV2.md)
 - [SourceAccountSummaryV3](docs/SourceAccountSummaryV3.md)
 - [SourceAccountSummaryV4](docs/SourceAccountSummaryV4.md)
 - [SupportedCountriesResponse](docs/SupportedCountriesResponse.md)
 - [SupportedCountry](docs/SupportedCountry.md)
 - [SupportedCurrency](docs/SupportedCurrency.md)
 - [SupportedCurrencyResponse](docs/SupportedCurrencyResponse.md)
 - [UpdateRemoteIdRequest](docs/UpdateRemoteIdRequest.md)
 - [WatchlistStatus](docs/WatchlistStatus.md)


## Documentation For Authorization



## OAuth2


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


## basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Author



