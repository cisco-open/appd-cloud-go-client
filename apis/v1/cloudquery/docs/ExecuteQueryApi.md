# \ExecuteQueryApi

All URIs are relative to *https://customer1.observe.appdynamics.com/monitoring/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteQuery**](ExecuteQueryApi.md#ExecuteQuery) | **Post** /query/execute | Execute a query to fetch observation data.



## ExecuteQuery

> []QueryResponseArrayBodyInner ExecuteQuery(ctx).QueryRequestBody(queryRequestBody).Execute()

Execute a query to fetch observation data.



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
    queryRequestBody := *openapiclient.NewQueryRequestBody() // QueryRequestBody | A simple JSON object containing a query string. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExecuteQueryApi.ExecuteQuery(context.Background()).QueryRequestBody(queryRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExecuteQueryApi.ExecuteQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteQuery`: []QueryResponseArrayBodyInner
    fmt.Fprintf(os.Stdout, "Response from `ExecuteQueryApi.ExecuteQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryRequestBody** | [**QueryRequestBody**](QueryRequestBody.md) | A simple JSON object containing a query string. | 

### Return type

[**[]QueryResponseArrayBodyInner**](QueryResponseArrayBodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/x-ndjson, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

