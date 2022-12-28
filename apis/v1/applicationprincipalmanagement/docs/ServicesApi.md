# \ServicesApi

All URIs are relative to *http://app-principal-mgmt.iam.svc.cluster.local:80/v1beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceClient**](ServicesApi.md#CreateServiceClient) | **Post** /clients/services | 
[**DeleteServiceClient**](ServicesApi.md#DeleteServiceClient) | **Delete** /clients/services/{clientId} | 
[**GetServiceClientById**](ServicesApi.md#GetServiceClientById) | **Get** /clients/services/{clientId} | 
[**ListAllServiceClients**](ServicesApi.md#ListAllServiceClients) | **Get** /clients/services | 
[**RevokeServiceClientSecret**](ServicesApi.md#RevokeServiceClientSecret) | **Post** /clients/services/{clientId}/secret/revokeRotated | 
[**RotateServiceClientSecret**](ServicesApi.md#RotateServiceClientSecret) | **Post** /clients/services/{clientId}/secret/rotate | 
[**UpdateServiceClient**](ServicesApi.md#UpdateServiceClient) | **Put** /clients/services/{clientId} | 



## CreateServiceClient

> ServiceClientResponseCreate CreateServiceClient(ctx).ServiceClientRequest(serviceClientRequest).Execute()





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
    serviceClientRequest := *openapiclient.NewServiceClientRequest() // ServiceClientRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.CreateServiceClient(context.Background()).ServiceClientRequest(serviceClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.CreateServiceClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceClient`: ServiceClientResponseCreate
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.CreateServiceClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceClientRequest** | [**ServiceClientRequest**](ServiceClientRequest.md) |  | 

### Return type

[**ServiceClientResponseCreate**](ServiceClientResponseCreate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceClient

> DeleteServiceClient(ctx, clientId).Execute()





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
    clientId := "clientId_example" // string | The unique client identifer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.DeleteServiceClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeleteServiceClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique client identifer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceClientById

> ServiceClientResponse GetServiceClientById(ctx, clientId).Execute()





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
    clientId := "clientId_example" // string | The unique client identifer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.GetServiceClientById(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServiceClientById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceClientById`: ServiceClientResponse
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetServiceClientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique client identifer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceClientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceClientResponse**](ServiceClientResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllServiceClients

> ServiceClientsContainer ListAllServiceClients(ctx).Cursor(cursor).Max(max).Execute()





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
    cursor := "cursor_example" // string | The page identifier used in pagination. (optional)
    max := int32(56) // int32 | Max number of agent clients to be returned. Defaults to 100. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.ListAllServiceClients(context.Background()).Cursor(cursor).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ListAllServiceClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllServiceClients`: ServiceClientsContainer
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ListAllServiceClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllServiceClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The page identifier used in pagination. | 
 **max** | **int32** | Max number of agent clients to be returned. Defaults to 100. | 

### Return type

[**ServiceClientsContainer**](ServiceClientsContainer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeServiceClientSecret

> RevokeSecretResponse RevokeServiceClientSecret(ctx, clientId).Execute()





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
    clientId := "clientId_example" // string | The unique client identifer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.RevokeServiceClientSecret(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.RevokeServiceClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeServiceClientSecret`: RevokeSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.RevokeServiceClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique client identifer | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeServiceClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RevokeSecretResponse**](RevokeSecretResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateServiceClientSecret

> CredentialsDto RotateServiceClientSecret(ctx, clientId).RotationRequest(rotationRequest).Execute()





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
    clientId := "clientId_example" // string | The unique client identifer
    rotationRequest := *openapiclient.NewRotationRequest() // RotationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.RotateServiceClientSecret(context.Background(), clientId).RotationRequest(rotationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.RotateServiceClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateServiceClientSecret`: CredentialsDto
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.RotateServiceClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique client identifer | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateServiceClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rotationRequest** | [**RotationRequest**](RotationRequest.md) |  | 

### Return type

[**CredentialsDto**](CredentialsDto.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceClient

> ServiceClientResponse UpdateServiceClient(ctx, clientId).ServiceClientRequest(serviceClientRequest).Execute()





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
    clientId := "clientId_example" // string | The unique client identifer
    serviceClientRequest := *openapiclient.NewServiceClientRequest() // ServiceClientRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.UpdateServiceClient(context.Background(), clientId).ServiceClientRequest(serviceClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateServiceClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceClient`: ServiceClientResponse
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.UpdateServiceClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique client identifer | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceClientRequest** | [**ServiceClientRequest**](ServiceClientRequest.md) |  | 

### Return type

[**ServiceClientResponse**](ServiceClientResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

