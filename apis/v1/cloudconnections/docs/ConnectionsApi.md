# \ConnectionsApi

All URIs are relative to *http://cloudmonconnectionservice.cloudmon.svc.cluster.local:7778/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](ConnectionsApi.md#CreateConnection) | **Post** /connections | Create a connection
[**DeleteConnection**](ConnectionsApi.md#DeleteConnection) | **Delete** /connections/{connectionId} | Delete connection
[**GetConnection**](ConnectionsApi.md#GetConnection) | **Get** /connections/{connectionId} | Get a connection
[**GetConnections**](ConnectionsApi.md#GetConnections) | **Get** /connections | Query connections
[**UpdateConnection**](ConnectionsApi.md#UpdateConnection) | **Patch** /connections/{connectionId} | Update a connection



## CreateConnection

> ConnectionResponse CreateConnection(ctx).ConnectionRequest(connectionRequest).Execute()

Create a connection



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
    connectionRequest := *openapiclient.NewConnectionRequest(openapiclient.ConnectionRequestDetails{AWSConnectionRequestDetails: penapiclient.AWSConnectionRequestDetails{AWSAccessKeyDetails: openapiclient.NewAWSAccessKeyDetails("AKBATYTCJY7XPKGABY5L", "7ccw68gBKMMXUdXBut+7qC9CCr1brotxD5ClcGGE", openapiclient.ConnectionAccessType("role_delegation"))}}, "AWS dev", openapiclient.ProviderType("aws")) // ConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.CreateConnection(context.Background()).ConnectionRequest(connectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnection`: ConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionRequest** | [**ConnectionRequest**](ConnectionRequest.md) |  | 

### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> ConnectionResponse DeleteConnection(ctx, connectionId).Execute()

Delete connection



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
    connectionId := "7c8c01e3-dc9a-4e7e-a536-7240c62af801" // string | Connection ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.DeleteConnection(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.DeleteConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConnection`: ConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.DeleteConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnection

> ConnectionResponse GetConnection(ctx, connectionId).Execute()

Get a connection



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
    connectionId := "7c8c01e3-dc9a-4e7e-a536-7240c62af801" // string | Connection ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetConnection(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnection`: ConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnections

> Connections GetConnections(ctx).Filter(filter).Sort(sort).Order(order).Max(max).Cursor(cursor).Execute()

Query connections



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
    filter := "displayName sw "dev"" // string | Filter expression (optional)
    sort := "sort_example" // string | Sort by (optional)
    order := "order_example" // string | Order by (optional)
    max := int32(56) // int32 | Maximum number of results to return (optional) (default to 10)
    cursor := "cursor_example" // string | Cursor for the paginated requests (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetConnections(context.Background()).Filter(filter).Sort(sort).Order(order).Max(max).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnections`: Connections
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter expression | 
 **sort** | **string** | Sort by | 
 **order** | **string** | Order by | 
 **max** | **int32** | Maximum number of results to return | [default to 10]
 **cursor** | **string** | Cursor for the paginated requests | 

### Return type

[**Connections**](Connections.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnection

> ConnectionResponse UpdateConnection(ctx, connectionId).ConnectionUpdate(connectionUpdate).Execute()

Update a connection



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
    connectionId := "7c8c01e3-dc9a-4e7e-a536-7240c62af801" // string | Connection ID
    connectionUpdate := *openapiclient.NewConnectionUpdate() // ConnectionUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.UpdateConnection(context.Background(), connectionId).ConnectionUpdate(connectionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.UpdateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnection`: ConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.UpdateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionUpdate** | [**ConnectionUpdate**](ConnectionUpdate.md) |  | 

### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

