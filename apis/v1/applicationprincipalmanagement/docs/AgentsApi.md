# \AgentsApi

All URIs are relative to *http://app-principal-mgmt.iam.svc.cluster.local:80/v1beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgentClient**](AgentsApi.md#CreateAgentClient) | **Post** /clients/agents | 
[**DeleteAgentClient**](AgentsApi.md#DeleteAgentClient) | **Delete** /clients/agents/{clientId} | 
[**GetAgentClient**](AgentsApi.md#GetAgentClient) | **Get** /clients/agents/{clientId} | 
[**ListAllAgentClients**](AgentsApi.md#ListAllAgentClients) | **Get** /clients/agents | 
[**RevokeAgentClientSecrets**](AgentsApi.md#RevokeAgentClientSecrets) | **Post** /clients/agents/{clientId}/secret/revokeRotated | 
[**RotateAgentClientSecret**](AgentsApi.md#RotateAgentClientSecret) | **Post** /clients/agents/{clientId}/secret/rotate | 
[**UpdateAgentClient**](AgentsApi.md#UpdateAgentClient) | **Put** /clients/agents/{clientId} | 



## CreateAgentClient

> AgentClientResponseCreate CreateAgentClient(ctx).AgentClientRequest(agentClientRequest).Execute()





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
    agentClientRequest := *openapiclient.NewAgentClientRequest() // AgentClientRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.CreateAgentClient(context.Background()).AgentClientRequest(agentClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.CreateAgentClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentClient`: AgentClientResponseCreate
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.CreateAgentClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentClientRequest** | [**AgentClientRequest**](AgentClientRequest.md) |  | 

### Return type

[**AgentClientResponseCreate**](AgentClientResponseCreate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentClient

> DeleteAgentClient(ctx, clientId).Execute()





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
    resp, r, err := apiClient.AgentsApi.DeleteAgentClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.DeleteAgentClient``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAgentClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentClient

> AgentClientResponse GetAgentClient(ctx, clientId).Execute()





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
    resp, r, err := apiClient.AgentsApi.GetAgentClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.GetAgentClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentClient`: AgentClientResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.GetAgentClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique client identifer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentClientResponse**](AgentClientResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllAgentClients

> AgentClientsContainer ListAllAgentClients(ctx).Cursor(cursor).Max(max).Execute()





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
    resp, r, err := apiClient.AgentsApi.ListAllAgentClients(context.Background()).Cursor(cursor).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.ListAllAgentClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllAgentClients`: AgentClientsContainer
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.ListAllAgentClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllAgentClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The page identifier used in pagination. | 
 **max** | **int32** | Max number of agent clients to be returned. Defaults to 100. | 

### Return type

[**AgentClientsContainer**](AgentClientsContainer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAgentClientSecrets

> RevokeSecretResponse RevokeAgentClientSecrets(ctx, clientId).Execute()





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
    resp, r, err := apiClient.AgentsApi.RevokeAgentClientSecrets(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.RevokeAgentClientSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeAgentClientSecrets`: RevokeSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.RevokeAgentClientSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique client identifer | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAgentClientSecretsRequest struct via the builder pattern


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


## RotateAgentClientSecret

> CredentialsDto RotateAgentClientSecret(ctx, clientId).RotationRequest(rotationRequest).Execute()





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
    resp, r, err := apiClient.AgentsApi.RotateAgentClientSecret(context.Background(), clientId).RotationRequest(rotationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.RotateAgentClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateAgentClientSecret`: CredentialsDto
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.RotateAgentClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique client identifer | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateAgentClientSecretRequest struct via the builder pattern


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


## UpdateAgentClient

> AgentClientResponse UpdateAgentClient(ctx, clientId).AgentClientRequest(agentClientRequest).Execute()





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
    agentClientRequest := *openapiclient.NewAgentClientRequest() // AgentClientRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.UpdateAgentClient(context.Background(), clientId).AgentClientRequest(agentClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.UpdateAgentClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentClient`: AgentClientResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.UpdateAgentClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique client identifer | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentClientRequest** | [**AgentClientRequest**](AgentClientRequest.md) |  | 

### Return type

[**AgentClientResponse**](AgentClientResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

