# \ResultPaginationApi

All URIs are relative to *https://customer1.observe.appdynamics.com/monitoring/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResultPagination**](ResultPaginationApi.md#ResultPagination) | **Get** /query/continue | Fetch the next page of results.



## ResultPagination

> []QueryResponseArrayBodyInner ResultPagination(ctx).Cursor(cursor).Execute()

Fetch the next page of results.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cursor := "ewogICJza2lwIiA6IDEwMDAsCiAgInF1ZXJ5IiA6ICJGRVRDSFxuICAgIGlkLFxuICAgIGhlYWx0aF9hc19" // string | An opaque string which will allow the retrieval of the next page of results. The cursor is obtained from a previous query response.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResultPaginationApi.ResultPagination(context.Background()).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultPaginationApi.ResultPagination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResultPagination`: []QueryResponseArrayBodyInner
    fmt.Fprintf(os.Stdout, "Response from `ResultPaginationApi.ResultPagination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResultPaginationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | An opaque string which will allow the retrieval of the next page of results. The cursor is obtained from a previous query response. | 

### Return type

[**[]QueryResponseArrayBodyInner**](QueryResponseArrayBodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

