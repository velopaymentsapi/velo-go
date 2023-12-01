/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response.   ## Http Status Codes Following is a list of Http Status codes that could be returned by the platform      | Status Code            | Description                                                                          |     | -----------------------| -------------------------------------------------------------------------------------|     | 200 OK                 | The request was successfully processed and usually returns a json response           |     | 201 Created            | A resource was created and a Location header is returned linking to the new resource |     | 202 Accepted           | The request has been accepted for processing                                         |     | 204 No Content         | The request has been processed and there is no response (usually deletes and updates)|     | 400 Bad Request        | The request is invalid and should be fixed before retrying                           |     | 401 Unauthorized       | Authentication has failed, usually means the token has expired                       |     | 403 Forbidden          | The user does not have permissions for the request                                   |     | 404 Not Found          | The resource was not found                                                           |     | 409 Conflict           | The resource already exists and there is a conflict                                  |     | 429 Too Many Requests  | The user has submitted too many requests in a given amount of time                   |     | 5xx Server Error       | Platform internal error (should rarely happen)                                       | 

API version: 2.37.150
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// PaymentAuditServiceDeprecatedAPIService PaymentAuditServiceDeprecatedAPI service
type PaymentAuditServiceDeprecatedAPIService service

type ApiExportTransactionsCSVV3Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceDeprecatedAPIService
	payorId *string
	startDate *string
	endDate *string
}

// The Payor ID for whom you wish to run the report. For a Payor requesting the report, this could be their exact Payor, or it could be a child/descendant Payor. 
func (r ApiExportTransactionsCSVV3Request) PayorId(payorId string) ApiExportTransactionsCSVV3Request {
	r.payorId = &payorId
	return r
}

// Start date, inclusive. Format is YYYY-MM-DD
func (r ApiExportTransactionsCSVV3Request) StartDate(startDate string) ApiExportTransactionsCSVV3Request {
	r.startDate = &startDate
	return r
}

// End date, inclusive. Format is YYYY-MM-DD
func (r ApiExportTransactionsCSVV3Request) EndDate(endDate string) ApiExportTransactionsCSVV3Request {
	r.endDate = &endDate
	return r
}

func (r ApiExportTransactionsCSVV3Request) Execute() (*PayorAmlTransactionV3, *http.Response, error) {
	return r.ApiService.ExportTransactionsCSVV3Execute(r)
}

/*
ExportTransactionsCSVV3 V3 Export Transactions

Deprecated (use /v4/paymentaudit/transactions instead)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExportTransactionsCSVV3Request

Deprecated
*/
func (a *PaymentAuditServiceDeprecatedAPIService) ExportTransactionsCSVV3(ctx context.Context) ApiExportTransactionsCSVV3Request {
	return ApiExportTransactionsCSVV3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PayorAmlTransactionV3
// Deprecated
func (a *PaymentAuditServiceDeprecatedAPIService) ExportTransactionsCSVV3Execute(r ApiExportTransactionsCSVV3Request) (*PayorAmlTransactionV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayorAmlTransactionV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceDeprecatedAPIService.ExportTransactionsCSVV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/paymentaudit/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.payorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/csv", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFundingsV1Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceDeprecatedAPIService
	payorId *string
	page *int32
	pageSize *int32
	sort *string
}

// The account owner Payor ID
func (r ApiGetFundingsV1Request) PayorId(payorId string) ApiGetFundingsV1Request {
	r.payorId = &payorId
	return r
}

// Page number. Default is 1.
func (r ApiGetFundingsV1Request) Page(page int32) ApiGetFundingsV1Request {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiGetFundingsV1Request) PageSize(pageSize int32) ApiGetFundingsV1Request {
	r.pageSize = &pageSize
	return r
}

// List of sort fields. Example: &#x60;&#x60;&#x60;?sort&#x3D;destinationCurrency:asc,destinationAmount:asc&#x60;&#x60;&#x60; Default is no sort. The supported sort fields are: dateTime and amount. 
func (r ApiGetFundingsV1Request) Sort(sort string) ApiGetFundingsV1Request {
	r.sort = &sort
	return r
}

func (r ApiGetFundingsV1Request) Execute() (*GetFundingsResponse, *http.Response, error) {
	return r.ApiService.GetFundingsV1Execute(r)
}

/*
GetFundingsV1 V1 Get Fundings for Payor

Deprecated (use /v4/paymentaudit/fundings)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFundingsV1Request

Deprecated
*/
func (a *PaymentAuditServiceDeprecatedAPIService) GetFundingsV1(ctx context.Context) ApiGetFundingsV1Request {
	return ApiGetFundingsV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFundingsResponse
// Deprecated
func (a *PaymentAuditServiceDeprecatedAPIService) GetFundingsV1Execute(r ApiGetFundingsV1Request) (*GetFundingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFundingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceDeprecatedAPIService.GetFundingsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/paymentaudit/fundings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payorId == nil {
		return localVarReturnValue, nil, reportError("payorId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 25
		r.pageSize = &defaultValue
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentDetailsV3Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceDeprecatedAPIService
	paymentId string
	sensitive *bool
}

// Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values. 
func (r ApiGetPaymentDetailsV3Request) Sensitive(sensitive bool) ApiGetPaymentDetailsV3Request {
	r.sensitive = &sensitive
	return r
}

func (r ApiGetPaymentDetailsV3Request) Execute() (*PaymentResponseV3, *http.Response, error) {
	return r.ApiService.GetPaymentDetailsV3Execute(r)
}

/*
GetPaymentDetailsV3 V3 Get Payment

Deprecated (use /v4/paymentaudit/payments/<paymentId> instead)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentId Payment Id
 @return ApiGetPaymentDetailsV3Request

Deprecated
*/
func (a *PaymentAuditServiceDeprecatedAPIService) GetPaymentDetailsV3(ctx context.Context, paymentId string) ApiGetPaymentDetailsV3Request {
	return ApiGetPaymentDetailsV3Request{
		ApiService: a,
		ctx: ctx,
		paymentId: paymentId,
	}
}

// Execute executes the request
//  @return PaymentResponseV3
// Deprecated
func (a *PaymentAuditServiceDeprecatedAPIService) GetPaymentDetailsV3Execute(r ApiGetPaymentDetailsV3Request) (*PaymentResponseV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentResponseV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceDeprecatedAPIService.GetPaymentDetailsV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/paymentaudit/payments/{paymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentId"+"}", url.PathEscape(parameterValueToString(r.paymentId, "paymentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sensitive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sensitive", r.sensitive, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentsForPayoutPAV3Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceDeprecatedAPIService
	payoutId string
	remoteId *string
	status *string
	sourceAmountFrom *int32
	sourceAmountTo *int32
	paymentAmountFrom *int32
	paymentAmountTo *int32
	submittedDateFrom *string
	submittedDateTo *string
	page *int32
	pageSize *int32
	sort *string
	sensitive *bool
}

// The remote id of the payees.
func (r ApiGetPaymentsForPayoutPAV3Request) RemoteId(remoteId string) ApiGetPaymentsForPayoutPAV3Request {
	r.remoteId = &remoteId
	return r
}

// Payment Status
func (r ApiGetPaymentsForPayoutPAV3Request) Status(status string) ApiGetPaymentsForPayoutPAV3Request {
	r.status = &status
	return r
}

// The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom
func (r ApiGetPaymentsForPayoutPAV3Request) SourceAmountFrom(sourceAmountFrom int32) ApiGetPaymentsForPayoutPAV3Request {
	r.sourceAmountFrom = &sourceAmountFrom
	return r
}

// The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo
func (r ApiGetPaymentsForPayoutPAV3Request) SourceAmountTo(sourceAmountTo int32) ApiGetPaymentsForPayoutPAV3Request {
	r.sourceAmountTo = &sourceAmountTo
	return r
}

// The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom
func (r ApiGetPaymentsForPayoutPAV3Request) PaymentAmountFrom(paymentAmountFrom int32) ApiGetPaymentsForPayoutPAV3Request {
	r.paymentAmountFrom = &paymentAmountFrom
	return r
}

// The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo
func (r ApiGetPaymentsForPayoutPAV3Request) PaymentAmountTo(paymentAmountTo int32) ApiGetPaymentsForPayoutPAV3Request {
	r.paymentAmountTo = &paymentAmountTo
	return r
}

// The submitted date from range filter. Format is yyyy-MM-dd.
func (r ApiGetPaymentsForPayoutPAV3Request) SubmittedDateFrom(submittedDateFrom string) ApiGetPaymentsForPayoutPAV3Request {
	r.submittedDateFrom = &submittedDateFrom
	return r
}

// The submitted date to range filter. Format is yyyy-MM-dd.
func (r ApiGetPaymentsForPayoutPAV3Request) SubmittedDateTo(submittedDateTo string) ApiGetPaymentsForPayoutPAV3Request {
	r.submittedDateTo = &submittedDateTo
	return r
}

// Page number. Default is 1.
func (r ApiGetPaymentsForPayoutPAV3Request) Page(page int32) ApiGetPaymentsForPayoutPAV3Request {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiGetPaymentsForPayoutPAV3Request) PageSize(pageSize int32) ApiGetPaymentsForPayoutPAV3Request {
	r.pageSize = &pageSize
	return r
}

// &lt;p&gt;List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId&lt;/p&gt; &lt;p&gt;The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status&lt;/p&gt; 
func (r ApiGetPaymentsForPayoutPAV3Request) Sort(sort string) ApiGetPaymentsForPayoutPAV3Request {
	r.sort = &sort
	return r
}

// Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values. 
func (r ApiGetPaymentsForPayoutPAV3Request) Sensitive(sensitive bool) ApiGetPaymentsForPayoutPAV3Request {
	r.sensitive = &sensitive
	return r
}

func (r ApiGetPaymentsForPayoutPAV3Request) Execute() (*GetPaymentsForPayoutResponseV3, *http.Response, error) {
	return r.ApiService.GetPaymentsForPayoutPAV3Execute(r)
}

/*
GetPaymentsForPayoutPAV3 V3 Get Payments for Payout

Deprecated (use /v4/paymentaudit/payouts/<payoutId> instead)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payoutId The id (UUID) of the payout.
 @return ApiGetPaymentsForPayoutPAV3Request

Deprecated
*/
func (a *PaymentAuditServiceDeprecatedAPIService) GetPaymentsForPayoutPAV3(ctx context.Context, payoutId string) ApiGetPaymentsForPayoutPAV3Request {
	return ApiGetPaymentsForPayoutPAV3Request{
		ApiService: a,
		ctx: ctx,
		payoutId: payoutId,
	}
}

// Execute executes the request
//  @return GetPaymentsForPayoutResponseV3
// Deprecated
func (a *PaymentAuditServiceDeprecatedAPIService) GetPaymentsForPayoutPAV3Execute(r ApiGetPaymentsForPayoutPAV3Request) (*GetPaymentsForPayoutResponseV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPaymentsForPayoutResponseV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceDeprecatedAPIService.GetPaymentsForPayoutPAV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/paymentaudit/payouts/{payoutId}"
	localVarPath = strings.Replace(localVarPath, "{"+"payoutId"+"}", url.PathEscape(parameterValueToString(r.payoutId, "payoutId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.remoteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remoteId", r.remoteId, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.sourceAmountFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sourceAmountFrom", r.sourceAmountFrom, "")
	}
	if r.sourceAmountTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sourceAmountTo", r.sourceAmountTo, "")
	}
	if r.paymentAmountFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paymentAmountFrom", r.paymentAmountFrom, "")
	}
	if r.paymentAmountTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paymentAmountTo", r.paymentAmountTo, "")
	}
	if r.submittedDateFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "submittedDateFrom", r.submittedDateFrom, "")
	}
	if r.submittedDateTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "submittedDateTo", r.submittedDateTo, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 25
		r.pageSize = &defaultValue
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.sensitive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sensitive", r.sensitive, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPayoutStatsV1Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceDeprecatedAPIService
	payorId *string
}

// The account owner Payor ID. Required for external users.
func (r ApiGetPayoutStatsV1Request) PayorId(payorId string) ApiGetPayoutStatsV1Request {
	r.payorId = &payorId
	return r
}

func (r ApiGetPayoutStatsV1Request) Execute() (*GetPayoutStatistics, *http.Response, error) {
	return r.ApiService.GetPayoutStatsV1Execute(r)
}

/*
GetPayoutStatsV1 V1 Get Payout Statistics

Deprecated (Use /v4/paymentaudit/payoutStatistics)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPayoutStatsV1Request

Deprecated
*/
func (a *PaymentAuditServiceDeprecatedAPIService) GetPayoutStatsV1(ctx context.Context) ApiGetPayoutStatsV1Request {
	return ApiGetPayoutStatsV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPayoutStatistics
// Deprecated
func (a *PaymentAuditServiceDeprecatedAPIService) GetPayoutStatsV1Execute(r ApiGetPayoutStatsV1Request) (*GetPayoutStatistics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPayoutStatistics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceDeprecatedAPIService.GetPayoutStatsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/paymentaudit/payoutStatistics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.payorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPayoutsForPayorV3Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceDeprecatedAPIService
	payorId *string
	payoutMemo *string
	status *string
	submittedDateFrom *string
	submittedDateTo *string
	page *int32
	pageSize *int32
	sort *string
}

// The account owner Payor ID
func (r ApiGetPayoutsForPayorV3Request) PayorId(payorId string) ApiGetPayoutsForPayorV3Request {
	r.payorId = &payorId
	return r
}

// Payout Memo filter - case insensitive sub-string match
func (r ApiGetPayoutsForPayorV3Request) PayoutMemo(payoutMemo string) ApiGetPayoutsForPayorV3Request {
	r.payoutMemo = &payoutMemo
	return r
}

// Payout Status
func (r ApiGetPayoutsForPayorV3Request) Status(status string) ApiGetPayoutsForPayorV3Request {
	r.status = &status
	return r
}

// The submitted date from range filter. Format is yyyy-MM-dd.
func (r ApiGetPayoutsForPayorV3Request) SubmittedDateFrom(submittedDateFrom string) ApiGetPayoutsForPayorV3Request {
	r.submittedDateFrom = &submittedDateFrom
	return r
}

// The submitted date to range filter. Format is yyyy-MM-dd.
func (r ApiGetPayoutsForPayorV3Request) SubmittedDateTo(submittedDateTo string) ApiGetPayoutsForPayorV3Request {
	r.submittedDateTo = &submittedDateTo
	return r
}

// Page number. Default is 1.
func (r ApiGetPayoutsForPayorV3Request) Page(page int32) ApiGetPayoutsForPayorV3Request {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiGetPayoutsForPayorV3Request) PageSize(pageSize int32) ApiGetPayoutsForPayorV3Request {
	r.pageSize = &pageSize
	return r
}

// List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,instructedDateTime:asc,status:asc) Default is submittedDateTime:asc The supported sort fields are: submittedDateTime, instructedDateTime, status. 
func (r ApiGetPayoutsForPayorV3Request) Sort(sort string) ApiGetPayoutsForPayorV3Request {
	r.sort = &sort
	return r
}

func (r ApiGetPayoutsForPayorV3Request) Execute() (*GetPayoutsResponseV3, *http.Response, error) {
	return r.ApiService.GetPayoutsForPayorV3Execute(r)
}

/*
GetPayoutsForPayorV3 V3 Get Payouts for Payor

Deprecated (use /v4/paymentaudit/payouts instead)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPayoutsForPayorV3Request

Deprecated
*/
func (a *PaymentAuditServiceDeprecatedAPIService) GetPayoutsForPayorV3(ctx context.Context) ApiGetPayoutsForPayorV3Request {
	return ApiGetPayoutsForPayorV3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPayoutsResponseV3
// Deprecated
func (a *PaymentAuditServiceDeprecatedAPIService) GetPayoutsForPayorV3Execute(r ApiGetPayoutsForPayorV3Request) (*GetPayoutsResponseV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPayoutsResponseV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceDeprecatedAPIService.GetPayoutsForPayorV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/paymentaudit/payouts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payorId == nil {
		return localVarReturnValue, nil, reportError("payorId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	if r.payoutMemo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payoutMemo", r.payoutMemo, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.submittedDateFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "submittedDateFrom", r.submittedDateFrom, "")
	}
	if r.submittedDateTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "submittedDateTo", r.submittedDateTo, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 25
		r.pageSize = &defaultValue
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPaymentChangesRequest struct {
	ctx context.Context
	ApiService *PaymentAuditServiceDeprecatedAPIService
	payorId *string
	updatedSince *time.Time
	page *int32
	pageSize *int32
}

// The Payor ID to find associated Payments
func (r ApiListPaymentChangesRequest) PayorId(payorId string) ApiListPaymentChangesRequest {
	r.payorId = &payorId
	return r
}

// The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
func (r ApiListPaymentChangesRequest) UpdatedSince(updatedSince time.Time) ApiListPaymentChangesRequest {
	r.updatedSince = &updatedSince
	return r
}

// Page number. Default is 1.
func (r ApiListPaymentChangesRequest) Page(page int32) ApiListPaymentChangesRequest {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiListPaymentChangesRequest) PageSize(pageSize int32) ApiListPaymentChangesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiListPaymentChangesRequest) Execute() (*PaymentDeltaResponseV1, *http.Response, error) {
	return r.ApiService.ListPaymentChangesExecute(r)
}

/*
ListPaymentChanges V1 List Payment Changes

Deprecated (use /v4/payments/deltas instead)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPaymentChangesRequest

Deprecated
*/
func (a *PaymentAuditServiceDeprecatedAPIService) ListPaymentChanges(ctx context.Context) ApiListPaymentChangesRequest {
	return ApiListPaymentChangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentDeltaResponseV1
// Deprecated
func (a *PaymentAuditServiceDeprecatedAPIService) ListPaymentChangesExecute(r ApiListPaymentChangesRequest) (*PaymentDeltaResponseV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentDeltaResponseV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceDeprecatedAPIService.ListPaymentChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/deltas/payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payorId == nil {
		return localVarReturnValue, nil, reportError("payorId is required and must be specified")
	}
	if r.updatedSince == nil {
		return localVarReturnValue, nil, reportError("updatedSince is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "updatedSince", r.updatedSince, "")
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 100
		r.pageSize = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPaymentsAuditV3Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceDeprecatedAPIService
	payeeId *string
	payorId *string
	payorName *string
	remoteId *string
	status *string
	sourceAccountName *string
	sourceAmountFrom *int32
	sourceAmountTo *int32
	sourceCurrency *string
	paymentAmountFrom *int32
	paymentAmountTo *int32
	paymentCurrency *string
	submittedDateFrom *string
	submittedDateTo *string
	paymentMemo *string
	page *int32
	pageSize *int32
	sort *string
	sensitive *bool
}

// The UUID of the payee.
func (r ApiListPaymentsAuditV3Request) PayeeId(payeeId string) ApiListPaymentsAuditV3Request {
	r.payeeId = &payeeId
	return r
}

// The account owner Payor Id. Required for external users.
func (r ApiListPaymentsAuditV3Request) PayorId(payorId string) ApiListPaymentsAuditV3Request {
	r.payorId = &payorId
	return r
}

// The payor’s name. This filters via a case insensitive substring match.
func (r ApiListPaymentsAuditV3Request) PayorName(payorName string) ApiListPaymentsAuditV3Request {
	r.payorName = &payorName
	return r
}

// The remote id of the payees.
func (r ApiListPaymentsAuditV3Request) RemoteId(remoteId string) ApiListPaymentsAuditV3Request {
	r.remoteId = &remoteId
	return r
}

// Payment Status
func (r ApiListPaymentsAuditV3Request) Status(status string) ApiListPaymentsAuditV3Request {
	r.status = &status
	return r
}

// The source account name filter. This filters via a case insensitive substring match.
func (r ApiListPaymentsAuditV3Request) SourceAccountName(sourceAccountName string) ApiListPaymentsAuditV3Request {
	r.sourceAccountName = &sourceAccountName
	return r
}

// The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom
func (r ApiListPaymentsAuditV3Request) SourceAmountFrom(sourceAmountFrom int32) ApiListPaymentsAuditV3Request {
	r.sourceAmountFrom = &sourceAmountFrom
	return r
}

// The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo
func (r ApiListPaymentsAuditV3Request) SourceAmountTo(sourceAmountTo int32) ApiListPaymentsAuditV3Request {
	r.sourceAmountTo = &sourceAmountTo
	return r
}

// The source currency filter. Filters based on an exact match on the currency.
func (r ApiListPaymentsAuditV3Request) SourceCurrency(sourceCurrency string) ApiListPaymentsAuditV3Request {
	r.sourceCurrency = &sourceCurrency
	return r
}

// The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom
func (r ApiListPaymentsAuditV3Request) PaymentAmountFrom(paymentAmountFrom int32) ApiListPaymentsAuditV3Request {
	r.paymentAmountFrom = &paymentAmountFrom
	return r
}

// The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo
func (r ApiListPaymentsAuditV3Request) PaymentAmountTo(paymentAmountTo int32) ApiListPaymentsAuditV3Request {
	r.paymentAmountTo = &paymentAmountTo
	return r
}

// The payment currency filter. Filters based on an exact match on the currency.
func (r ApiListPaymentsAuditV3Request) PaymentCurrency(paymentCurrency string) ApiListPaymentsAuditV3Request {
	r.paymentCurrency = &paymentCurrency
	return r
}

// The submitted date from range filter. Format is yyyy-MM-dd.
func (r ApiListPaymentsAuditV3Request) SubmittedDateFrom(submittedDateFrom string) ApiListPaymentsAuditV3Request {
	r.submittedDateFrom = &submittedDateFrom
	return r
}

// The submitted date to range filter. Format is yyyy-MM-dd.
func (r ApiListPaymentsAuditV3Request) SubmittedDateTo(submittedDateTo string) ApiListPaymentsAuditV3Request {
	r.submittedDateTo = &submittedDateTo
	return r
}

// The payment memo filter. This filters via a case insensitive substring match.
func (r ApiListPaymentsAuditV3Request) PaymentMemo(paymentMemo string) ApiListPaymentsAuditV3Request {
	r.paymentMemo = &paymentMemo
	return r
}

// Page number. Default is 1.
func (r ApiListPaymentsAuditV3Request) Page(page int32) ApiListPaymentsAuditV3Request {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiListPaymentsAuditV3Request) PageSize(pageSize int32) ApiListPaymentsAuditV3Request {
	r.pageSize = &pageSize
	return r
}

// List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status 
func (r ApiListPaymentsAuditV3Request) Sort(sort string) ApiListPaymentsAuditV3Request {
	r.sort = &sort
	return r
}

// Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values. 
func (r ApiListPaymentsAuditV3Request) Sensitive(sensitive bool) ApiListPaymentsAuditV3Request {
	r.sensitive = &sensitive
	return r
}

func (r ApiListPaymentsAuditV3Request) Execute() (*ListPaymentsResponseV3, *http.Response, error) {
	return r.ApiService.ListPaymentsAuditV3Execute(r)
}

/*
ListPaymentsAuditV3 V3 Get List of Payments

Deprecated (use /v4/paymentaudit/payments instead)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPaymentsAuditV3Request

Deprecated
*/
func (a *PaymentAuditServiceDeprecatedAPIService) ListPaymentsAuditV3(ctx context.Context) ApiListPaymentsAuditV3Request {
	return ApiListPaymentsAuditV3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListPaymentsResponseV3
// Deprecated
func (a *PaymentAuditServiceDeprecatedAPIService) ListPaymentsAuditV3Execute(r ApiListPaymentsAuditV3Request) (*ListPaymentsResponseV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListPaymentsResponseV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceDeprecatedAPIService.ListPaymentsAuditV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/paymentaudit/payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.payeeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payeeId", r.payeeId, "")
	}
	if r.payorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	}
	if r.payorName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payorName", r.payorName, "")
	}
	if r.remoteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remoteId", r.remoteId, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.sourceAccountName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sourceAccountName", r.sourceAccountName, "")
	}
	if r.sourceAmountFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sourceAmountFrom", r.sourceAmountFrom, "")
	}
	if r.sourceAmountTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sourceAmountTo", r.sourceAmountTo, "")
	}
	if r.sourceCurrency != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sourceCurrency", r.sourceCurrency, "")
	}
	if r.paymentAmountFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paymentAmountFrom", r.paymentAmountFrom, "")
	}
	if r.paymentAmountTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paymentAmountTo", r.paymentAmountTo, "")
	}
	if r.paymentCurrency != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paymentCurrency", r.paymentCurrency, "")
	}
	if r.submittedDateFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "submittedDateFrom", r.submittedDateFrom, "")
	}
	if r.submittedDateTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "submittedDateTo", r.submittedDateTo, "")
	}
	if r.paymentMemo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paymentMemo", r.paymentMemo, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 25
		r.pageSize = &defaultValue
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.sensitive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sensitive", r.sensitive, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
