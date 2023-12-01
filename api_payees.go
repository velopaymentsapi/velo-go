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


// PayeesAPIService PayeesAPI service
type PayeesAPIService service

type ApiDeletePayeeByIdV3Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payeeId string
}

func (r ApiDeletePayeeByIdV3Request) Execute() (*http.Response, error) {
	return r.ApiService.DeletePayeeByIdV3Execute(r)
}

/*
DeletePayeeByIdV3 Delete Payee by Id

<p>Use v4 instead</p>
<p>This API will delete Payee by Id (UUID). Deletion by ID is not allowed if:</p>
<p>* Payee ID is not found</p>
<p>* If Payee has not been on-boarded</p>
<p>* If Payee is in grace period</p>
<p>* If Payee has existing payments</p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payeeId The UUID of the payee.
 @return ApiDeletePayeeByIdV3Request

Deprecated
*/
func (a *PayeesAPIService) DeletePayeeByIdV3(ctx context.Context, payeeId string) ApiDeletePayeeByIdV3Request {
	return ApiDeletePayeeByIdV3Request{
		ApiService: a,
		ctx: ctx,
		payeeId: payeeId,
	}
}

// Execute executes the request
// Deprecated
func (a *PayeesAPIService) DeletePayeeByIdV3Execute(r ApiDeletePayeeByIdV3Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.DeletePayeeByIdV3")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/payees/{payeeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"payeeId"+"}", url.PathEscape(parameterValueToString(r.payeeId, "payeeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeletePayeeByIdV4Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payeeId string
}

func (r ApiDeletePayeeByIdV4Request) Execute() (*http.Response, error) {
	return r.ApiService.DeletePayeeByIdV4Execute(r)
}

/*
DeletePayeeByIdV4 Delete Payee by Id

<p>This API will delete Payee by Id (UUID). Deletion by ID is not allowed if:</p>
<p>* Payee ID is not found</p>
<p>* If Payee has not been on-boarded</p>
<p>* If Payee is in grace period</p>
<p>* If Payee has existing payments</p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payeeId The UUID of the payee.
 @return ApiDeletePayeeByIdV4Request
*/
func (a *PayeesAPIService) DeletePayeeByIdV4(ctx context.Context, payeeId string) ApiDeletePayeeByIdV4Request {
	return ApiDeletePayeeByIdV4Request{
		ApiService: a,
		ctx: ctx,
		payeeId: payeeId,
	}
}

// Execute executes the request
func (a *PayeesAPIService) DeletePayeeByIdV4Execute(r ApiDeletePayeeByIdV4Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.DeletePayeeByIdV4")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/payees/{payeeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"payeeId"+"}", url.PathEscape(parameterValueToString(r.payeeId, "payeeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPayeeByIdV3Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payeeId string
	sensitive *bool
}

// Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values. 
func (r ApiGetPayeeByIdV3Request) Sensitive(sensitive bool) ApiGetPayeeByIdV3Request {
	r.sensitive = &sensitive
	return r
}

func (r ApiGetPayeeByIdV3Request) Execute() (*PayeeDetailResponseV3, *http.Response, error) {
	return r.ApiService.GetPayeeByIdV3Execute(r)
}

/*
GetPayeeByIdV3 Get Payee by Id

<p>Use v4 instead</p>
<p>Get Payee by Id</p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payeeId The UUID of the payee.
 @return ApiGetPayeeByIdV3Request

Deprecated
*/
func (a *PayeesAPIService) GetPayeeByIdV3(ctx context.Context, payeeId string) ApiGetPayeeByIdV3Request {
	return ApiGetPayeeByIdV3Request{
		ApiService: a,
		ctx: ctx,
		payeeId: payeeId,
	}
}

// Execute executes the request
//  @return PayeeDetailResponseV3
// Deprecated
func (a *PayeesAPIService) GetPayeeByIdV3Execute(r ApiGetPayeeByIdV3Request) (*PayeeDetailResponseV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayeeDetailResponseV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.GetPayeeByIdV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/payees/{payeeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"payeeId"+"}", url.PathEscape(parameterValueToString(r.payeeId, "payeeId")), -1)

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

type ApiGetPayeeByIdV4Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payeeId string
	sensitive *bool
}

// Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values. 
func (r ApiGetPayeeByIdV4Request) Sensitive(sensitive bool) ApiGetPayeeByIdV4Request {
	r.sensitive = &sensitive
	return r
}

func (r ApiGetPayeeByIdV4Request) Execute() (*PayeeDetailResponseV4, *http.Response, error) {
	return r.ApiService.GetPayeeByIdV4Execute(r)
}

/*
GetPayeeByIdV4 Get Payee by Id

Get Payee by Id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payeeId The UUID of the payee.
 @return ApiGetPayeeByIdV4Request
*/
func (a *PayeesAPIService) GetPayeeByIdV4(ctx context.Context, payeeId string) ApiGetPayeeByIdV4Request {
	return ApiGetPayeeByIdV4Request{
		ApiService: a,
		ctx: ctx,
		payeeId: payeeId,
	}
}

// Execute executes the request
//  @return PayeeDetailResponseV4
func (a *PayeesAPIService) GetPayeeByIdV4Execute(r ApiGetPayeeByIdV4Request) (*PayeeDetailResponseV4, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayeeDetailResponseV4
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.GetPayeeByIdV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/payees/{payeeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"payeeId"+"}", url.PathEscape(parameterValueToString(r.payeeId, "payeeId")), -1)

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

type ApiListPayeeChangesV3Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payorId *string
	updatedSince *time.Time
	page *int32
	pageSize *int32
}

// The Payor ID to find associated Payees
func (r ApiListPayeeChangesV3Request) PayorId(payorId string) ApiListPayeeChangesV3Request {
	r.payorId = &payorId
	return r
}

// The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
func (r ApiListPayeeChangesV3Request) UpdatedSince(updatedSince time.Time) ApiListPayeeChangesV3Request {
	r.updatedSince = &updatedSince
	return r
}

// Page number. Default is 1.
func (r ApiListPayeeChangesV3Request) Page(page int32) ApiListPayeeChangesV3Request {
	r.page = &page
	return r
}

// Page size. Default is 100. Max allowable is 1000.
func (r ApiListPayeeChangesV3Request) PageSize(pageSize int32) ApiListPayeeChangesV3Request {
	r.pageSize = &pageSize
	return r
}

func (r ApiListPayeeChangesV3Request) Execute() (*PayeeDeltaResponseV3, *http.Response, error) {
	return r.ApiService.ListPayeeChangesV3Execute(r)
}

/*
ListPayeeChangesV3 List Payee Changes

<p>Use v4 instead</p>
<p>Get a paginated response listing payee changes.</p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPayeeChangesV3Request

Deprecated
*/
func (a *PayeesAPIService) ListPayeeChangesV3(ctx context.Context) ApiListPayeeChangesV3Request {
	return ApiListPayeeChangesV3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PayeeDeltaResponseV3
// Deprecated
func (a *PayeesAPIService) ListPayeeChangesV3Execute(r ApiListPayeeChangesV3Request) (*PayeeDeltaResponseV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayeeDeltaResponseV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.ListPayeeChangesV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/payees/deltas"

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

type ApiListPayeeChangesV4Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payorId *string
	updatedSince *time.Time
	page *int32
	pageSize *int32
}

// The Payor ID to find associated Payees
func (r ApiListPayeeChangesV4Request) PayorId(payorId string) ApiListPayeeChangesV4Request {
	r.payorId = &payorId
	return r
}

// The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
func (r ApiListPayeeChangesV4Request) UpdatedSince(updatedSince time.Time) ApiListPayeeChangesV4Request {
	r.updatedSince = &updatedSince
	return r
}

// Page number. Default is 1.
func (r ApiListPayeeChangesV4Request) Page(page int32) ApiListPayeeChangesV4Request {
	r.page = &page
	return r
}

// Page size. Default is 100. Max allowable is 1000.
func (r ApiListPayeeChangesV4Request) PageSize(pageSize int32) ApiListPayeeChangesV4Request {
	r.pageSize = &pageSize
	return r
}

func (r ApiListPayeeChangesV4Request) Execute() (*PayeeDeltaResponseV4, *http.Response, error) {
	return r.ApiService.ListPayeeChangesV4Execute(r)
}

/*
ListPayeeChangesV4 List Payee Changes

Get a paginated response listing payee changes (updated since a particular time) to a limited set of fields:
- dbaName
- displayName
- email
- onboardedStatus
- payeeCountry
- payeeId
- remoteId


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPayeeChangesV4Request
*/
func (a *PayeesAPIService) ListPayeeChangesV4(ctx context.Context) ApiListPayeeChangesV4Request {
	return ApiListPayeeChangesV4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PayeeDeltaResponseV4
func (a *PayeesAPIService) ListPayeeChangesV4Execute(r ApiListPayeeChangesV4Request) (*PayeeDeltaResponseV4, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayeeDeltaResponseV4
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.ListPayeeChangesV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/payees/deltas"

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

type ApiListPayeesV3Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payorId *string
	watchlistStatus *string
	disabled *bool
	onboardedStatus *string
	email *string
	displayName *string
	remoteId *string
	payeeType *string
	payeeCountry *string
	page *int32
	pageSize *int32
	sort *string
}

// The account owner Payor ID
func (r ApiListPayeesV3Request) PayorId(payorId string) ApiListPayeesV3Request {
	r.payorId = &payorId
	return r
}

// The watchlistStatus of the payees.
func (r ApiListPayeesV3Request) WatchlistStatus(watchlistStatus string) ApiListPayeesV3Request {
	r.watchlistStatus = &watchlistStatus
	return r
}

// Payee disabled
func (r ApiListPayeesV3Request) Disabled(disabled bool) ApiListPayeesV3Request {
	r.disabled = &disabled
	return r
}

// The onboarded status of the payees.
func (r ApiListPayeesV3Request) OnboardedStatus(onboardedStatus string) ApiListPayeesV3Request {
	r.onboardedStatus = &onboardedStatus
	return r
}

// Email address
func (r ApiListPayeesV3Request) Email(email string) ApiListPayeesV3Request {
	r.email = &email
	return r
}

// The display name of the payees.
func (r ApiListPayeesV3Request) DisplayName(displayName string) ApiListPayeesV3Request {
	r.displayName = &displayName
	return r
}

// The remote id of the payees.
func (r ApiListPayeesV3Request) RemoteId(remoteId string) ApiListPayeesV3Request {
	r.remoteId = &remoteId
	return r
}

// The onboarded status of the payees.
func (r ApiListPayeesV3Request) PayeeType(payeeType string) ApiListPayeesV3Request {
	r.payeeType = &payeeType
	return r
}

// The country of the payee - 2 letter ISO 3166-1 country code (upper case)
func (r ApiListPayeesV3Request) PayeeCountry(payeeCountry string) ApiListPayeesV3Request {
	r.payeeCountry = &payeeCountry
	return r
}

// Page number. Default is 1.
func (r ApiListPayeesV3Request) Page(page int32) ApiListPayeesV3Request {
	r.page = &page
	return r
}

// Page size. Default is 25. Max allowable is 100.
func (r ApiListPayeesV3Request) PageSize(pageSize int32) ApiListPayeesV3Request {
	r.pageSize = &pageSize
	return r
}

// List of sort fields (e.g. ?sort&#x3D;onboardedStatus:asc,name:asc) Default is name:asc &#39;name&#39; is treated as company name for companies - last name + &#39;,&#39; + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus. 
func (r ApiListPayeesV3Request) Sort(sort string) ApiListPayeesV3Request {
	r.sort = &sort
	return r
}

func (r ApiListPayeesV3Request) Execute() (*PagedPayeeResponseV3, *http.Response, error) {
	return r.ApiService.ListPayeesV3Execute(r)
}

/*
ListPayeesV3 List Payees

<p>Use v4 instead</p>
Get a paginated response listing the payees for a payor.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPayeesV3Request

Deprecated
*/
func (a *PayeesAPIService) ListPayeesV3(ctx context.Context) ApiListPayeesV3Request {
	return ApiListPayeesV3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PagedPayeeResponseV3
// Deprecated
func (a *PayeesAPIService) ListPayeesV3Execute(r ApiListPayeesV3Request) (*PagedPayeeResponseV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PagedPayeeResponseV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.ListPayeesV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/payees"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payorId == nil {
		return localVarReturnValue, nil, reportError("payorId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	if r.watchlistStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "watchlistStatus", r.watchlistStatus, "")
	}
	if r.disabled != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "disabled", r.disabled, "")
	}
	if r.onboardedStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "onboardedStatus", r.onboardedStatus, "")
	}
	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "")
	}
	if r.displayName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "displayName", r.displayName, "")
	}
	if r.remoteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remoteId", r.remoteId, "")
	}
	if r.payeeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payeeType", r.payeeType, "")
	}
	if r.payeeCountry != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payeeCountry", r.payeeCountry, "")
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
	} else {
		var defaultValue string = "displayName:asc"
		r.sort = &defaultValue
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiListPayeesV4Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payorId *string
	watchlistStatus *string
	disabled *bool
	onboardedStatus *string
	email *string
	displayName *string
	remoteId *string
	payeeType *string
	payeeCountry *string
	ofacStatus *string
	page *int32
	pageSize *int32
	sort *string
}

// The account owner Payor ID
func (r ApiListPayeesV4Request) PayorId(payorId string) ApiListPayeesV4Request {
	r.payorId = &payorId
	return r
}

// The watchlistStatus of the payees.
func (r ApiListPayeesV4Request) WatchlistStatus(watchlistStatus string) ApiListPayeesV4Request {
	r.watchlistStatus = &watchlistStatus
	return r
}

// Payee disabled
func (r ApiListPayeesV4Request) Disabled(disabled bool) ApiListPayeesV4Request {
	r.disabled = &disabled
	return r
}

// The onboarded status of the payees.
func (r ApiListPayeesV4Request) OnboardedStatus(onboardedStatus string) ApiListPayeesV4Request {
	r.onboardedStatus = &onboardedStatus
	return r
}

// Email address
func (r ApiListPayeesV4Request) Email(email string) ApiListPayeesV4Request {
	r.email = &email
	return r
}

// The display name of the payees.
func (r ApiListPayeesV4Request) DisplayName(displayName string) ApiListPayeesV4Request {
	r.displayName = &displayName
	return r
}

// The remote id of the payees.
func (r ApiListPayeesV4Request) RemoteId(remoteId string) ApiListPayeesV4Request {
	r.remoteId = &remoteId
	return r
}

// The onboarded status of the payees.
func (r ApiListPayeesV4Request) PayeeType(payeeType string) ApiListPayeesV4Request {
	r.payeeType = &payeeType
	return r
}

// The country of the payee - 2 letter ISO 3166-1 country code (upper case)
func (r ApiListPayeesV4Request) PayeeCountry(payeeCountry string) ApiListPayeesV4Request {
	r.payeeCountry = &payeeCountry
	return r
}

// The ofacStatus of the payees.
func (r ApiListPayeesV4Request) OfacStatus(ofacStatus string) ApiListPayeesV4Request {
	r.ofacStatus = &ofacStatus
	return r
}

// Page number. Default is 1.
func (r ApiListPayeesV4Request) Page(page int32) ApiListPayeesV4Request {
	r.page = &page
	return r
}

// Page size. Default is 25. Max allowable is 100.
func (r ApiListPayeesV4Request) PageSize(pageSize int32) ApiListPayeesV4Request {
	r.pageSize = &pageSize
	return r
}

// List of sort fields (e.g. ?sort&#x3D;onboardedStatus:asc,name:asc) Default is name:asc &#39;name&#39; is treated as company name for companies - last name + &#39;,&#39; + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus. 
func (r ApiListPayeesV4Request) Sort(sort string) ApiListPayeesV4Request {
	r.sort = &sort
	return r
}

func (r ApiListPayeesV4Request) Execute() (*PagedPayeeResponseV4, *http.Response, error) {
	return r.ApiService.ListPayeesV4Execute(r)
}

/*
ListPayeesV4 List Payees

Get a paginated response listing the payees for a payor.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPayeesV4Request
*/
func (a *PayeesAPIService) ListPayeesV4(ctx context.Context) ApiListPayeesV4Request {
	return ApiListPayeesV4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PagedPayeeResponseV4
func (a *PayeesAPIService) ListPayeesV4Execute(r ApiListPayeesV4Request) (*PagedPayeeResponseV4, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PagedPayeeResponseV4
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.ListPayeesV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/payees"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payorId == nil {
		return localVarReturnValue, nil, reportError("payorId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	if r.watchlistStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "watchlistStatus", r.watchlistStatus, "")
	}
	if r.disabled != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "disabled", r.disabled, "")
	}
	if r.onboardedStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "onboardedStatus", r.onboardedStatus, "")
	}
	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "")
	}
	if r.displayName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "displayName", r.displayName, "")
	}
	if r.remoteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remoteId", r.remoteId, "")
	}
	if r.payeeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payeeType", r.payeeType, "")
	}
	if r.payeeCountry != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payeeCountry", r.payeeCountry, "")
	}
	if r.ofacStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ofacStatus", r.ofacStatus, "")
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
	} else {
		var defaultValue string = "displayName:asc"
		r.sort = &defaultValue
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiPayeeDetailsUpdateV3Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payeeId string
	updatePayeeDetailsRequestV3 *UpdatePayeeDetailsRequestV3
}

// Request to update payee details
func (r ApiPayeeDetailsUpdateV3Request) UpdatePayeeDetailsRequestV3(updatePayeeDetailsRequestV3 UpdatePayeeDetailsRequestV3) ApiPayeeDetailsUpdateV3Request {
	r.updatePayeeDetailsRequestV3 = &updatePayeeDetailsRequestV3
	return r
}

func (r ApiPayeeDetailsUpdateV3Request) Execute() (*http.Response, error) {
	return r.ApiService.PayeeDetailsUpdateV3Execute(r)
}

/*
PayeeDetailsUpdateV3 Update Payee Details

<p>Use v4 instead</p>
<p>Update payee details for the given Payee Id.<p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payeeId The UUID of the payee.
 @return ApiPayeeDetailsUpdateV3Request

Deprecated
*/
func (a *PayeesAPIService) PayeeDetailsUpdateV3(ctx context.Context, payeeId string) ApiPayeeDetailsUpdateV3Request {
	return ApiPayeeDetailsUpdateV3Request{
		ApiService: a,
		ctx: ctx,
		payeeId: payeeId,
	}
}

// Execute executes the request
// Deprecated
func (a *PayeesAPIService) PayeeDetailsUpdateV3Execute(r ApiPayeeDetailsUpdateV3Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.PayeeDetailsUpdateV3")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/payees/{payeeId}/payeeDetailsUpdate"
	localVarPath = strings.Replace(localVarPath, "{"+"payeeId"+"}", url.PathEscape(parameterValueToString(r.payeeId, "payeeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updatePayeeDetailsRequestV3 == nil {
		return nil, reportError("updatePayeeDetailsRequestV3 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updatePayeeDetailsRequestV3
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPayeeDetailsUpdateV4Request struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payeeId string
	updatePayeeDetailsRequestV4 *UpdatePayeeDetailsRequestV4
}

// Request to update payee details
func (r ApiPayeeDetailsUpdateV4Request) UpdatePayeeDetailsRequestV4(updatePayeeDetailsRequestV4 UpdatePayeeDetailsRequestV4) ApiPayeeDetailsUpdateV4Request {
	r.updatePayeeDetailsRequestV4 = &updatePayeeDetailsRequestV4
	return r
}

func (r ApiPayeeDetailsUpdateV4Request) Execute() (*http.Response, error) {
	return r.ApiService.PayeeDetailsUpdateV4Execute(r)
}

/*
PayeeDetailsUpdateV4 Update Payee Details

<p>Update payee details for the given Payee Id.</p>
<p>Payors may only update the payee details if the payee has not yet onboarded</p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payeeId The UUID of the payee.
 @return ApiPayeeDetailsUpdateV4Request
*/
func (a *PayeesAPIService) PayeeDetailsUpdateV4(ctx context.Context, payeeId string) ApiPayeeDetailsUpdateV4Request {
	return ApiPayeeDetailsUpdateV4Request{
		ApiService: a,
		ctx: ctx,
		payeeId: payeeId,
	}
}

// Execute executes the request
func (a *PayeesAPIService) PayeeDetailsUpdateV4Execute(r ApiPayeeDetailsUpdateV4Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.PayeeDetailsUpdateV4")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/payees/{payeeId}/payeeDetailsUpdate"
	localVarPath = strings.Replace(localVarPath, "{"+"payeeId"+"}", url.PathEscape(parameterValueToString(r.payeeId, "payeeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updatePayeeDetailsRequestV4 == nil {
		return nil, reportError("updatePayeeDetailsRequestV4 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updatePayeeDetailsRequestV4
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV3PayeesPayeeIdRemoteIdUpdatePostRequest struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payeeId string
	updateRemoteIdRequestV3 *UpdateRemoteIdRequestV3
}

// Request to update payee remote id v3
func (r ApiV3PayeesPayeeIdRemoteIdUpdatePostRequest) UpdateRemoteIdRequestV3(updateRemoteIdRequestV3 UpdateRemoteIdRequestV3) ApiV3PayeesPayeeIdRemoteIdUpdatePostRequest {
	r.updateRemoteIdRequestV3 = &updateRemoteIdRequestV3
	return r
}

func (r ApiV3PayeesPayeeIdRemoteIdUpdatePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V3PayeesPayeeIdRemoteIdUpdatePostExecute(r)
}

/*
V3PayeesPayeeIdRemoteIdUpdatePost Update Payee Remote Id

<p>Use v4 instead</p>
<p>Update the remote Id for the given Payee Id.</p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payeeId The UUID of the payee.
 @return ApiV3PayeesPayeeIdRemoteIdUpdatePostRequest

Deprecated
*/
func (a *PayeesAPIService) V3PayeesPayeeIdRemoteIdUpdatePost(ctx context.Context, payeeId string) ApiV3PayeesPayeeIdRemoteIdUpdatePostRequest {
	return ApiV3PayeesPayeeIdRemoteIdUpdatePostRequest{
		ApiService: a,
		ctx: ctx,
		payeeId: payeeId,
	}
}

// Execute executes the request
// Deprecated
func (a *PayeesAPIService) V3PayeesPayeeIdRemoteIdUpdatePostExecute(r ApiV3PayeesPayeeIdRemoteIdUpdatePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.V3PayeesPayeeIdRemoteIdUpdatePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/payees/{payeeId}/remoteIdUpdate"
	localVarPath = strings.Replace(localVarPath, "{"+"payeeId"+"}", url.PathEscape(parameterValueToString(r.payeeId, "payeeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRemoteIdRequestV3 == nil {
		return nil, reportError("updateRemoteIdRequestV3 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateRemoteIdRequestV3
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v InlineResponse409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV4PayeesPayeeIdRemoteIdUpdatePostRequest struct {
	ctx context.Context
	ApiService *PayeesAPIService
	payeeId string
	updateRemoteIdRequestV4 *UpdateRemoteIdRequestV4
}

// Request to update payee remote id v4
func (r ApiV4PayeesPayeeIdRemoteIdUpdatePostRequest) UpdateRemoteIdRequestV4(updateRemoteIdRequestV4 UpdateRemoteIdRequestV4) ApiV4PayeesPayeeIdRemoteIdUpdatePostRequest {
	r.updateRemoteIdRequestV4 = &updateRemoteIdRequestV4
	return r
}

func (r ApiV4PayeesPayeeIdRemoteIdUpdatePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V4PayeesPayeeIdRemoteIdUpdatePostExecute(r)
}

/*
V4PayeesPayeeIdRemoteIdUpdatePost Update Payee Remote Id

<p>Update the remote Id for the given Payee Id.</p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payeeId The UUID of the payee.
 @return ApiV4PayeesPayeeIdRemoteIdUpdatePostRequest
*/
func (a *PayeesAPIService) V4PayeesPayeeIdRemoteIdUpdatePost(ctx context.Context, payeeId string) ApiV4PayeesPayeeIdRemoteIdUpdatePostRequest {
	return ApiV4PayeesPayeeIdRemoteIdUpdatePostRequest{
		ApiService: a,
		ctx: ctx,
		payeeId: payeeId,
	}
}

// Execute executes the request
func (a *PayeesAPIService) V4PayeesPayeeIdRemoteIdUpdatePostExecute(r ApiV4PayeesPayeeIdRemoteIdUpdatePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayeesAPIService.V4PayeesPayeeIdRemoteIdUpdatePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/payees/{payeeId}/remoteIdUpdate"
	localVarPath = strings.Replace(localVarPath, "{"+"payeeId"+"}", url.PathEscape(parameterValueToString(r.payeeId, "payeeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRemoteIdRequestV4 == nil {
		return nil, reportError("updateRemoteIdRequestV4 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateRemoteIdRequestV4
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v InlineResponse409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
