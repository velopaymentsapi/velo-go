# \PayorHierarchyAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PayorLinksV1**](PayorHierarchyAPI.md#PayorLinksV1) | **Get** /v1/payorLinks | List Payor Links



## PayorLinksV1

> PayorLinksResponse PayorLinksV1(ctx).DescendantsOfPayor(descendantsOfPayor).ParentOfPayor(parentOfPayor).Fields(fields).Execute()

List Payor Links



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    descendantsOfPayor := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor ID from which to start the query to show all descendants (optional)
    parentOfPayor := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Query for the parent payor details for this payor id (optional)
    fields := "fields_example" // string | <p>List of additional Payor fields to include in the response for each Payor</p> <p>The values of payorId and payorName are always included for each Payor by default</p> <p>You can add fields to the response for each payor by including them in the fields parameter separated by commas</p> <p>The supported fields are any combination of: primaryContactEmail,kycState</p>  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorHierarchyAPI.PayorLinksV1(context.Background()).DescendantsOfPayor(descendantsOfPayor).ParentOfPayor(parentOfPayor).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorHierarchyAPI.PayorLinksV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayorLinksV1`: PayorLinksResponse
    fmt.Fprintf(os.Stdout, "Response from `PayorHierarchyAPI.PayorLinksV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayorLinksV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **descendantsOfPayor** | **string** | The Payor ID from which to start the query to show all descendants | 
 **parentOfPayor** | **string** | Query for the parent payor details for this payor id | 
 **fields** | **string** | &lt;p&gt;List of additional Payor fields to include in the response for each Payor&lt;/p&gt; &lt;p&gt;The values of payorId and payorName are always included for each Payor by default&lt;/p&gt; &lt;p&gt;You can add fields to the response for each payor by including them in the fields parameter separated by commas&lt;/p&gt; &lt;p&gt;The supported fields are any combination of: primaryContactEmail,kycState&lt;/p&gt;  | 

### Return type

[**PayorLinksResponse**](PayorLinksResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

