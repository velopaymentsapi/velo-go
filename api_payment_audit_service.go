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


// PaymentAuditServiceAPIService PaymentAuditServiceAPI service
type PaymentAuditServiceAPIService service

type ApiExportTransactionsCSVV4Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceAPIService
	payorId *string
	startDate *string
	endDate *string
	include *string
}

// &lt;p&gt;The Payor ID for whom you wish to run the report.&lt;/p&gt; &lt;p&gt;For a Payor requesting the report, this could be their exact Payor, or it could be a child/descendant Payor.&lt;/p&gt; 
func (r ApiExportTransactionsCSVV4Request) PayorId(payorId string) ApiExportTransactionsCSVV4Request {
	r.payorId = &payorId
	return r
}

// Start date, inclusive. Format is YYYY-MM-DD
func (r ApiExportTransactionsCSVV4Request) StartDate(startDate string) ApiExportTransactionsCSVV4Request {
	r.startDate = &startDate
	return r
}

// End date, inclusive. Format is YYYY-MM-DD
func (r ApiExportTransactionsCSVV4Request) EndDate(endDate string) ApiExportTransactionsCSVV4Request {
	r.endDate = &endDate
	return r
}

// &lt;p&gt;Mode to determine whether to include other Payor&#39;s data in the results.&lt;/p&gt; &lt;p&gt;May only be used if payorId is specified.&lt;/p&gt; &lt;p&gt;Can be omitted or set to &#39;payorOnly&#39; or &#39;payorAndDescendants&#39;.&lt;/p&gt; &lt;p&gt;payorOnly: Only include results for the specified Payor. This is the default if &#39;include&#39; is omitted.&lt;/p&gt; &lt;p&gt;payorAndDescendants: Aggregate results for all descendant Payors of the specified Payor. Should only be used if the Payor with the specified payorId has at least one child Payor.&lt;/p&gt; &lt;p&gt;Note when a Payor requests the report and include&#x3D;payorAndDescendants is used, the following additional columns are included in the CSV: Payor Name, Payor Id&lt;/p&gt; 
func (r ApiExportTransactionsCSVV4Request) Include(include string) ApiExportTransactionsCSVV4Request {
	r.include = &include
	return r
}

func (r ApiExportTransactionsCSVV4Request) Execute() (*PayorAmlTransaction, *http.Response, error) {
	return r.ApiService.ExportTransactionsCSVV4Execute(r)
}

/*
ExportTransactionsCSVV4 Export Transactions

Download a CSV file containing payments in a date range. Uses Transfer-Encoding - chunked to stream to the client. Date range is inclusive of both the start and end dates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExportTransactionsCSVV4Request
*/
func (a *PaymentAuditServiceAPIService) ExportTransactionsCSVV4(ctx context.Context) ApiExportTransactionsCSVV4Request {
	return ApiExportTransactionsCSVV4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PayorAmlTransaction
func (a *PaymentAuditServiceAPIService) ExportTransactionsCSVV4Execute(r ApiExportTransactionsCSVV4Request) (*PayorAmlTransaction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PayorAmlTransaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceAPIService.ExportTransactionsCSVV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/paymentaudit/transactions"

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
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "")
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

type ApiGetFundingsV4Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceAPIService
	payorId *string
	sourceAccountName *string
	page *int32
	pageSize *int32
	sort *string
}

// The account owner Payor ID
func (r ApiGetFundingsV4Request) PayorId(payorId string) ApiGetFundingsV4Request {
	r.payorId = &payorId
	return r
}

// The source account name
func (r ApiGetFundingsV4Request) SourceAccountName(sourceAccountName string) ApiGetFundingsV4Request {
	r.sourceAccountName = &sourceAccountName
	return r
}

// Page number. Default is 1.
func (r ApiGetFundingsV4Request) Page(page int32) ApiGetFundingsV4Request {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiGetFundingsV4Request) PageSize(pageSize int32) ApiGetFundingsV4Request {
	r.pageSize = &pageSize
	return r
}

// List of sort fields. Example: &#x60;&#x60;&#x60;?sort&#x3D;destinationCurrency:asc,destinationAmount:asc&#x60;&#x60;&#x60; Default is no sort. The supported sort fields are: dateTime and amount. 
func (r ApiGetFundingsV4Request) Sort(sort string) ApiGetFundingsV4Request {
	r.sort = &sort
	return r
}

func (r ApiGetFundingsV4Request) Execute() (*GetFundingsResponse, *http.Response, error) {
	return r.ApiService.GetFundingsV4Execute(r)
}

/*
GetFundingsV4 Get Fundings for Payor

<p>Get a list of Fundings for a payor.</p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFundingsV4Request
*/
func (a *PaymentAuditServiceAPIService) GetFundingsV4(ctx context.Context) ApiGetFundingsV4Request {
	return ApiGetFundingsV4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFundingsResponse
func (a *PaymentAuditServiceAPIService) GetFundingsV4Execute(r ApiGetFundingsV4Request) (*GetFundingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFundingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceAPIService.GetFundingsV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/paymentaudit/fundings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payorId == nil {
		return localVarReturnValue, nil, reportError("payorId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	if r.sourceAccountName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sourceAccountName", r.sourceAccountName, "")
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

type ApiGetPaymentDetailsV4Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceAPIService
	paymentId string
	sensitive *bool
}

// Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values. 
func (r ApiGetPaymentDetailsV4Request) Sensitive(sensitive bool) ApiGetPaymentDetailsV4Request {
	r.sensitive = &sensitive
	return r
}

func (r ApiGetPaymentDetailsV4Request) Execute() (*PaymentResponseV4, *http.Response, error) {
	return r.ApiService.GetPaymentDetailsV4Execute(r)
}

/*
GetPaymentDetailsV4 Get Payment

Get the payment with the given id. This contains the payment history.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentId Payment Id
 @return ApiGetPaymentDetailsV4Request
*/
func (a *PaymentAuditServiceAPIService) GetPaymentDetailsV4(ctx context.Context, paymentId string) ApiGetPaymentDetailsV4Request {
	return ApiGetPaymentDetailsV4Request{
		ApiService: a,
		ctx: ctx,
		paymentId: paymentId,
	}
}

// Execute executes the request
//  @return PaymentResponseV4
func (a *PaymentAuditServiceAPIService) GetPaymentDetailsV4Execute(r ApiGetPaymentDetailsV4Request) (*PaymentResponseV4, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentResponseV4
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceAPIService.GetPaymentDetailsV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/paymentaudit/payments/{paymentId}"
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

type ApiGetPaymentsForPayoutV4Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceAPIService
	payoutId string
	railsId *string
	remoteId *string
	remoteSystemId *string
	status *string
	sourceAmountFrom *int32
	sourceAmountTo *int32
	paymentAmountFrom *int32
	paymentAmountTo *int32
	submittedDateFrom *string
	submittedDateTo *string
	transmissionType *string
	page *int32
	pageSize *int32
	sort *string
	sensitive *bool
}

// Payout Rails ID filter - case insensitive match. Any value is allowed, but you should use one of the supported railsId values. To get this list of values, you should call the &#39;Get Supported Rails&#39; endpoint. 
func (r ApiGetPaymentsForPayoutV4Request) RailsId(railsId string) ApiGetPaymentsForPayoutV4Request {
	r.railsId = &railsId
	return r
}

// The remote id of the payees.
func (r ApiGetPaymentsForPayoutV4Request) RemoteId(remoteId string) ApiGetPaymentsForPayoutV4Request {
	r.remoteId = &remoteId
	return r
}

// The id of the remote system that is orchestrating payments
func (r ApiGetPaymentsForPayoutV4Request) RemoteSystemId(remoteSystemId string) ApiGetPaymentsForPayoutV4Request {
	r.remoteSystemId = &remoteSystemId
	return r
}

// Payment Status
func (r ApiGetPaymentsForPayoutV4Request) Status(status string) ApiGetPaymentsForPayoutV4Request {
	r.status = &status
	return r
}

// The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom
func (r ApiGetPaymentsForPayoutV4Request) SourceAmountFrom(sourceAmountFrom int32) ApiGetPaymentsForPayoutV4Request {
	r.sourceAmountFrom = &sourceAmountFrom
	return r
}

// The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo
func (r ApiGetPaymentsForPayoutV4Request) SourceAmountTo(sourceAmountTo int32) ApiGetPaymentsForPayoutV4Request {
	r.sourceAmountTo = &sourceAmountTo
	return r
}

// The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom
func (r ApiGetPaymentsForPayoutV4Request) PaymentAmountFrom(paymentAmountFrom int32) ApiGetPaymentsForPayoutV4Request {
	r.paymentAmountFrom = &paymentAmountFrom
	return r
}

// The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo
func (r ApiGetPaymentsForPayoutV4Request) PaymentAmountTo(paymentAmountTo int32) ApiGetPaymentsForPayoutV4Request {
	r.paymentAmountTo = &paymentAmountTo
	return r
}

// The submitted date from range filter. Format is yyyy-MM-dd.
func (r ApiGetPaymentsForPayoutV4Request) SubmittedDateFrom(submittedDateFrom string) ApiGetPaymentsForPayoutV4Request {
	r.submittedDateFrom = &submittedDateFrom
	return r
}

// The submitted date to range filter. Format is yyyy-MM-dd.
func (r ApiGetPaymentsForPayoutV4Request) SubmittedDateTo(submittedDateTo string) ApiGetPaymentsForPayoutV4Request {
	r.submittedDateTo = &submittedDateTo
	return r
}

// Transmission Type * ACH * SAME_DAY_ACH * WIRE * GACHO 
func (r ApiGetPaymentsForPayoutV4Request) TransmissionType(transmissionType string) ApiGetPaymentsForPayoutV4Request {
	r.transmissionType = &transmissionType
	return r
}

// Page number. Default is 1.
func (r ApiGetPaymentsForPayoutV4Request) Page(page int32) ApiGetPaymentsForPayoutV4Request {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiGetPaymentsForPayoutV4Request) PageSize(pageSize int32) ApiGetPaymentsForPayoutV4Request {
	r.pageSize = &pageSize
	return r
}

// List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status 
func (r ApiGetPaymentsForPayoutV4Request) Sort(sort string) ApiGetPaymentsForPayoutV4Request {
	r.sort = &sort
	return r
}

// Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values. 
func (r ApiGetPaymentsForPayoutV4Request) Sensitive(sensitive bool) ApiGetPaymentsForPayoutV4Request {
	r.sensitive = &sensitive
	return r
}

func (r ApiGetPaymentsForPayoutV4Request) Execute() (*GetPaymentsForPayoutResponseV4, *http.Response, error) {
	return r.ApiService.GetPaymentsForPayoutV4Execute(r)
}

/*
GetPaymentsForPayoutV4 Get Payments for Payout

Get List of payments for Payout, allowing for RETURNED status


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param payoutId The id (UUID) of the payout.
 @return ApiGetPaymentsForPayoutV4Request
*/
func (a *PaymentAuditServiceAPIService) GetPaymentsForPayoutV4(ctx context.Context, payoutId string) ApiGetPaymentsForPayoutV4Request {
	return ApiGetPaymentsForPayoutV4Request{
		ApiService: a,
		ctx: ctx,
		payoutId: payoutId,
	}
}

// Execute executes the request
//  @return GetPaymentsForPayoutResponseV4
func (a *PaymentAuditServiceAPIService) GetPaymentsForPayoutV4Execute(r ApiGetPaymentsForPayoutV4Request) (*GetPaymentsForPayoutResponseV4, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPaymentsForPayoutResponseV4
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceAPIService.GetPaymentsForPayoutV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/paymentaudit/payouts/{payoutId}"
	localVarPath = strings.Replace(localVarPath, "{"+"payoutId"+"}", url.PathEscape(parameterValueToString(r.payoutId, "payoutId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.railsId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "railsId", r.railsId, "")
	}
	if r.remoteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remoteId", r.remoteId, "")
	}
	if r.remoteSystemId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remoteSystemId", r.remoteSystemId, "")
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
	if r.transmissionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transmissionType", r.transmissionType, "")
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

type ApiGetPayoutStatsV4Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceAPIService
	payorId *string
}

// The account owner Payor ID. Required for external users.
func (r ApiGetPayoutStatsV4Request) PayorId(payorId string) ApiGetPayoutStatsV4Request {
	r.payorId = &payorId
	return r
}

func (r ApiGetPayoutStatsV4Request) Execute() (*GetPayoutStatistics, *http.Response, error) {
	return r.ApiService.GetPayoutStatsV4Execute(r)
}

/*
GetPayoutStatsV4 Get Payout Statistics

<p>Get payout statistics for a payor.</p>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPayoutStatsV4Request
*/
func (a *PaymentAuditServiceAPIService) GetPayoutStatsV4(ctx context.Context) ApiGetPayoutStatsV4Request {
	return ApiGetPayoutStatsV4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPayoutStatistics
func (a *PaymentAuditServiceAPIService) GetPayoutStatsV4Execute(r ApiGetPayoutStatsV4Request) (*GetPayoutStatistics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPayoutStatistics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceAPIService.GetPayoutStatsV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/paymentaudit/payoutStatistics"

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

type ApiGetPayoutsForPayorV4Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceAPIService
	payorId *string
	payoutMemo *string
	status *string
	submittedDateFrom *string
	submittedDateTo *string
	fromPayorName *string
	scheduledForDateFrom *string
	scheduledForDateTo *string
	scheduleStatus *string
	page *int32
	pageSize *int32
	sort *string
}

// The id (UUID) of the payor funding the payout or the payor whose payees are being paid.
func (r ApiGetPayoutsForPayorV4Request) PayorId(payorId string) ApiGetPayoutsForPayorV4Request {
	r.payorId = &payorId
	return r
}

// Payout Memo filter - case insensitive sub-string match
func (r ApiGetPayoutsForPayorV4Request) PayoutMemo(payoutMemo string) ApiGetPayoutsForPayorV4Request {
	r.payoutMemo = &payoutMemo
	return r
}

// Payout Status
func (r ApiGetPayoutsForPayorV4Request) Status(status string) ApiGetPayoutsForPayorV4Request {
	r.status = &status
	return r
}

// The submitted date from range filter. Format is yyyy-MM-dd.
func (r ApiGetPayoutsForPayorV4Request) SubmittedDateFrom(submittedDateFrom string) ApiGetPayoutsForPayorV4Request {
	r.submittedDateFrom = &submittedDateFrom
	return r
}

// The submitted date to range filter. Format is yyyy-MM-dd.
func (r ApiGetPayoutsForPayorV4Request) SubmittedDateTo(submittedDateTo string) ApiGetPayoutsForPayorV4Request {
	r.submittedDateTo = &submittedDateTo
	return r
}

// The name of the payor whose payees are being paid. This filters via a case insensitive substring match.
func (r ApiGetPayoutsForPayorV4Request) FromPayorName(fromPayorName string) ApiGetPayoutsForPayorV4Request {
	r.fromPayorName = &fromPayorName
	return r
}

// Filter payouts scheduled to run on or after the given date. Format is yyyy-MM-dd.
func (r ApiGetPayoutsForPayorV4Request) ScheduledForDateFrom(scheduledForDateFrom string) ApiGetPayoutsForPayorV4Request {
	r.scheduledForDateFrom = &scheduledForDateFrom
	return r
}

// Filter payouts scheduled to run on or before the given date. Format is yyyy-MM-dd.
func (r ApiGetPayoutsForPayorV4Request) ScheduledForDateTo(scheduledForDateTo string) ApiGetPayoutsForPayorV4Request {
	r.scheduledForDateTo = &scheduledForDateTo
	return r
}

// Payout Schedule Status
func (r ApiGetPayoutsForPayorV4Request) ScheduleStatus(scheduleStatus string) ApiGetPayoutsForPayorV4Request {
	r.scheduleStatus = &scheduleStatus
	return r
}

// Page number. Default is 1.
func (r ApiGetPayoutsForPayorV4Request) Page(page int32) ApiGetPayoutsForPayorV4Request {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiGetPayoutsForPayorV4Request) PageSize(pageSize int32) ApiGetPayoutsForPayorV4Request {
	r.pageSize = &pageSize
	return r
}

// List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,instructedDateTime:asc,status:asc) Default is submittedDateTime:asc The supported sort fields are: submittedDateTime, instructedDateTime, status, totalPayments, payoutId, scheduledFor 
func (r ApiGetPayoutsForPayorV4Request) Sort(sort string) ApiGetPayoutsForPayorV4Request {
	r.sort = &sort
	return r
}

func (r ApiGetPayoutsForPayorV4Request) Execute() (*GetPayoutsResponse, *http.Response, error) {
	return r.ApiService.GetPayoutsForPayorV4Execute(r)
}

/*
GetPayoutsForPayorV4 Get Payouts for Payor

Get List of payouts for payor


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPayoutsForPayorV4Request
*/
func (a *PaymentAuditServiceAPIService) GetPayoutsForPayorV4(ctx context.Context) ApiGetPayoutsForPayorV4Request {
	return ApiGetPayoutsForPayorV4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPayoutsResponse
func (a *PaymentAuditServiceAPIService) GetPayoutsForPayorV4Execute(r ApiGetPayoutsForPayorV4Request) (*GetPayoutsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPayoutsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceAPIService.GetPayoutsForPayorV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/paymentaudit/payouts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.payorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payorId", r.payorId, "")
	}
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
	if r.fromPayorName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromPayorName", r.fromPayorName, "")
	}
	if r.scheduledForDateFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scheduledForDateFrom", r.scheduledForDateFrom, "")
	}
	if r.scheduledForDateTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scheduledForDateTo", r.scheduledForDateTo, "")
	}
	if r.scheduleStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scheduleStatus", r.scheduleStatus, "")
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

type ApiListPaymentChangesV4Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceAPIService
	payorId *string
	updatedSince *time.Time
	page *int32
	pageSize *int32
}

// The Payor ID to find associated Payments
func (r ApiListPaymentChangesV4Request) PayorId(payorId string) ApiListPaymentChangesV4Request {
	r.payorId = &payorId
	return r
}

// The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
func (r ApiListPaymentChangesV4Request) UpdatedSince(updatedSince time.Time) ApiListPaymentChangesV4Request {
	r.updatedSince = &updatedSince
	return r
}

// Page number. Default is 1.
func (r ApiListPaymentChangesV4Request) Page(page int32) ApiListPaymentChangesV4Request {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiListPaymentChangesV4Request) PageSize(pageSize int32) ApiListPaymentChangesV4Request {
	r.pageSize = &pageSize
	return r
}

func (r ApiListPaymentChangesV4Request) Execute() (*PaymentDeltaResponse, *http.Response, error) {
	return r.ApiService.ListPaymentChangesV4Execute(r)
}

/*
ListPaymentChangesV4 List Payment Changes

Get a paginated response listing payment changes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPaymentChangesV4Request
*/
func (a *PaymentAuditServiceAPIService) ListPaymentChangesV4(ctx context.Context) ApiListPaymentChangesV4Request {
	return ApiListPaymentChangesV4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentDeltaResponse
func (a *PaymentAuditServiceAPIService) ListPaymentChangesV4Execute(r ApiListPaymentChangesV4Request) (*PaymentDeltaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentDeltaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceAPIService.ListPaymentChangesV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/payments/deltas"

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

type ApiListPaymentsAuditV4Request struct {
	ctx context.Context
	ApiService *PaymentAuditServiceAPIService
	payeeId *string
	payorId *string
	payorName *string
	remoteId *string
	remoteSystemId *string
	status *string
	transmissionType *string
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
	payorPaymentId *string
	railsId *string
	scheduledForDateFrom *string
	scheduledForDateTo *string
	scheduleStatus *string
	postInstructFxStatus *string
	transactionReference *string
	transactionId *string
	page *int32
	pageSize *int32
	sort *string
	sensitive *bool
}

// The UUID of the payee.
func (r ApiListPaymentsAuditV4Request) PayeeId(payeeId string) ApiListPaymentsAuditV4Request {
	r.payeeId = &payeeId
	return r
}

// The account owner Payor Id. Required for external users.
func (r ApiListPaymentsAuditV4Request) PayorId(payorId string) ApiListPaymentsAuditV4Request {
	r.payorId = &payorId
	return r
}

// The payor’s name. This filters via a case insensitive substring match.
func (r ApiListPaymentsAuditV4Request) PayorName(payorName string) ApiListPaymentsAuditV4Request {
	r.payorName = &payorName
	return r
}

// The remote id of the payees.
func (r ApiListPaymentsAuditV4Request) RemoteId(remoteId string) ApiListPaymentsAuditV4Request {
	r.remoteId = &remoteId
	return r
}

// The id of the remote system that is orchestrating payments
func (r ApiListPaymentsAuditV4Request) RemoteSystemId(remoteSystemId string) ApiListPaymentsAuditV4Request {
	r.remoteSystemId = &remoteSystemId
	return r
}

// Payment Status
func (r ApiListPaymentsAuditV4Request) Status(status string) ApiListPaymentsAuditV4Request {
	r.status = &status
	return r
}

// Transmission Type * ACH * SAME_DAY_ACH * WIRE * GACHO 
func (r ApiListPaymentsAuditV4Request) TransmissionType(transmissionType string) ApiListPaymentsAuditV4Request {
	r.transmissionType = &transmissionType
	return r
}

// The source account name filter. This filters via a case insensitive substring match.
func (r ApiListPaymentsAuditV4Request) SourceAccountName(sourceAccountName string) ApiListPaymentsAuditV4Request {
	r.sourceAccountName = &sourceAccountName
	return r
}

// The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom
func (r ApiListPaymentsAuditV4Request) SourceAmountFrom(sourceAmountFrom int32) ApiListPaymentsAuditV4Request {
	r.sourceAmountFrom = &sourceAmountFrom
	return r
}

// The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo
func (r ApiListPaymentsAuditV4Request) SourceAmountTo(sourceAmountTo int32) ApiListPaymentsAuditV4Request {
	r.sourceAmountTo = &sourceAmountTo
	return r
}

// The source currency filter. Filters based on an exact match on the currency.
func (r ApiListPaymentsAuditV4Request) SourceCurrency(sourceCurrency string) ApiListPaymentsAuditV4Request {
	r.sourceCurrency = &sourceCurrency
	return r
}

// The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom
func (r ApiListPaymentsAuditV4Request) PaymentAmountFrom(paymentAmountFrom int32) ApiListPaymentsAuditV4Request {
	r.paymentAmountFrom = &paymentAmountFrom
	return r
}

// The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo
func (r ApiListPaymentsAuditV4Request) PaymentAmountTo(paymentAmountTo int32) ApiListPaymentsAuditV4Request {
	r.paymentAmountTo = &paymentAmountTo
	return r
}

// The payment currency filter. Filters based on an exact match on the currency.
func (r ApiListPaymentsAuditV4Request) PaymentCurrency(paymentCurrency string) ApiListPaymentsAuditV4Request {
	r.paymentCurrency = &paymentCurrency
	return r
}

// The submitted date from range filter. Format is yyyy-MM-dd.
func (r ApiListPaymentsAuditV4Request) SubmittedDateFrom(submittedDateFrom string) ApiListPaymentsAuditV4Request {
	r.submittedDateFrom = &submittedDateFrom
	return r
}

// The submitted date to range filter. Format is yyyy-MM-dd.
func (r ApiListPaymentsAuditV4Request) SubmittedDateTo(submittedDateTo string) ApiListPaymentsAuditV4Request {
	r.submittedDateTo = &submittedDateTo
	return r
}

// The payment memo filter. This filters via a case insensitive substring match.
func (r ApiListPaymentsAuditV4Request) PaymentMemo(paymentMemo string) ApiListPaymentsAuditV4Request {
	r.paymentMemo = &paymentMemo
	return r
}

// Payor&#39;s Id of the Payment
func (r ApiListPaymentsAuditV4Request) PayorPaymentId(payorPaymentId string) ApiListPaymentsAuditV4Request {
	r.payorPaymentId = &payorPaymentId
	return r
}

// Payout Rails ID filter - case insensitive match. Any value is allowed, but you should use one of the supported railsId values. To get this list of values, you should call the &#39;Get Supported Rails&#39; endpoint. 
func (r ApiListPaymentsAuditV4Request) RailsId(railsId string) ApiListPaymentsAuditV4Request {
	r.railsId = &railsId
	return r
}

// Filter payouts scheduled to run on or after the given date. Format is yyyy-MM-dd.
func (r ApiListPaymentsAuditV4Request) ScheduledForDateFrom(scheduledForDateFrom string) ApiListPaymentsAuditV4Request {
	r.scheduledForDateFrom = &scheduledForDateFrom
	return r
}

// Filter payouts scheduled to run on or before the given date. Format is yyyy-MM-dd.
func (r ApiListPaymentsAuditV4Request) ScheduledForDateTo(scheduledForDateTo string) ApiListPaymentsAuditV4Request {
	r.scheduledForDateTo = &scheduledForDateTo
	return r
}

// Payout Schedule Status
func (r ApiListPaymentsAuditV4Request) ScheduleStatus(scheduleStatus string) ApiListPaymentsAuditV4Request {
	r.scheduleStatus = &scheduleStatus
	return r
}

// The status of the post instruct FX step if one was required for the payment
func (r ApiListPaymentsAuditV4Request) PostInstructFxStatus(postInstructFxStatus string) ApiListPaymentsAuditV4Request {
	r.postInstructFxStatus = &postInstructFxStatus
	return r
}

// Query for all payments associated with a specific transaction reference
func (r ApiListPaymentsAuditV4Request) TransactionReference(transactionReference string) ApiListPaymentsAuditV4Request {
	r.transactionReference = &transactionReference
	return r
}

// Query for all payments associated with a specific transaction id
func (r ApiListPaymentsAuditV4Request) TransactionId(transactionId string) ApiListPaymentsAuditV4Request {
	r.transactionId = &transactionId
	return r
}

// Page number. Default is 1.
func (r ApiListPaymentsAuditV4Request) Page(page int32) ApiListPaymentsAuditV4Request {
	r.page = &page
	return r
}

// The number of results to return in a page
func (r ApiListPaymentsAuditV4Request) PageSize(pageSize int32) ApiListPaymentsAuditV4Request {
	r.pageSize = &pageSize
	return r
}

// List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by submittedDateTime:desc,paymentId:asc The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime, status and paymentId 
func (r ApiListPaymentsAuditV4Request) Sort(sort string) ApiListPaymentsAuditV4Request {
	r.sort = &sort
	return r
}

// Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values. 
func (r ApiListPaymentsAuditV4Request) Sensitive(sensitive bool) ApiListPaymentsAuditV4Request {
	r.sensitive = &sensitive
	return r
}

func (r ApiListPaymentsAuditV4Request) Execute() (*ListPaymentsResponseV4, *http.Response, error) {
	return r.ApiService.ListPaymentsAuditV4Execute(r)
}

/*
ListPaymentsAuditV4 Get List of Payments

Get payments for the given payor Id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPaymentsAuditV4Request
*/
func (a *PaymentAuditServiceAPIService) ListPaymentsAuditV4(ctx context.Context) ApiListPaymentsAuditV4Request {
	return ApiListPaymentsAuditV4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListPaymentsResponseV4
func (a *PaymentAuditServiceAPIService) ListPaymentsAuditV4Execute(r ApiListPaymentsAuditV4Request) (*ListPaymentsResponseV4, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListPaymentsResponseV4
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentAuditServiceAPIService.ListPaymentsAuditV4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v4/paymentaudit/payments"

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
	if r.remoteSystemId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remoteSystemId", r.remoteSystemId, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.transmissionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transmissionType", r.transmissionType, "")
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
	if r.payorPaymentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payorPaymentId", r.payorPaymentId, "")
	}
	if r.railsId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "railsId", r.railsId, "")
	}
	if r.scheduledForDateFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scheduledForDateFrom", r.scheduledForDateFrom, "")
	}
	if r.scheduledForDateTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scheduledForDateTo", r.scheduledForDateTo, "")
	}
	if r.scheduleStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scheduleStatus", r.scheduleStatus, "")
	}
	if r.postInstructFxStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "postInstructFxStatus", r.postInstructFxStatus, "")
	}
	if r.transactionReference != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transactionReference", r.transactionReference, "")
	}
	if r.transactionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transactionId", r.transactionId, "")
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
