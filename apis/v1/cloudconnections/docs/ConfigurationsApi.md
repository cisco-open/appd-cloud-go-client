# \ConfigurationsApi

All URIs are relative to *http://cloudmonconnectionservice.cloudmon.svc.cluster.local:7778/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](ConfigurationsApi.md#CreateConfiguration) | **Post** /configurations | Create a configuration
[**DeleteConfiguration**](ConfigurationsApi.md#DeleteConfiguration) | **Delete** /configurations/{configurationId} | Delete configuration
[**GetConfiguration**](ConfigurationsApi.md#GetConfiguration) | **Get** /configurations/{configurationId} | Get a configuration
[**GetConfigurations**](ConfigurationsApi.md#GetConfigurations) | **Get** /configurations | Query configurations
[**UpdateConfiguration**](ConfigurationsApi.md#UpdateConfiguration) | **Patch** /configurations/{configurationId} | Update a configuration



## CreateConfiguration

> ConfigurationDetail CreateConfiguration(ctx).Configuration(configuration).Execute()

Create a configuration



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
    configuration := openapiclient.Configuration{AWSConfiguration: openapiclient.NewAWSConfiguration(*openapiclient.NewAWSConfigurationDetails(), "AWS dev", openapiclient.ProviderType("aws"))} // Configuration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationsApi.CreateConfiguration(context.Background()).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.CreateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfiguration`: ConfigurationDetail
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.CreateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | [**Configuration**](Configuration.md) |  | 

### Return type

[**ConfigurationDetail**](ConfigurationDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration

> DeleteConfiguration(ctx, configurationId).Execute()

Delete configuration



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
    configurationId := "7c8c01e3-dc9a-4e7e-a536-7240c62af801" // string | Configuration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationsApi.DeleteConfiguration(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.DeleteConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | Configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> ConfigurationDetail GetConfiguration(ctx, configurationId).Execute()

Get a configuration



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
    configurationId := "7c8c01e3-dc9a-4e7e-a536-7240c62af801" // string | Configuration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationsApi.GetConfiguration(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.GetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration`: ConfigurationDetail
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | Configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationDetail**](ConfigurationDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurations

> Configurations GetConfigurations(ctx).Filter(filter).Sort(sort).Order(order).Max(max).Cursor(cursor).Execute()

Query configurations



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
    resp, r, err := apiClient.ConfigurationsApi.GetConfigurations(context.Background()).Filter(filter).Sort(sort).Order(order).Max(max).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.GetConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurations`: Configurations
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.GetConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter expression | 
 **sort** | **string** | Sort by | 
 **order** | **string** | Order by | 
 **max** | **int32** | Maximum number of results to return | [default to 10]
 **cursor** | **string** | Cursor for the paginated requests | 

### Return type

[**Configurations**](Configurations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> ConfigurationDetail UpdateConfiguration(ctx, configurationId).ConfigurationUpdate(configurationUpdate).Execute()

Update a configuration



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
    configurationId := "7c8c01e3-dc9a-4e7e-a536-7240c62af801" // string | Configuration ID
    configurationUpdate := *openapiclient.NewConfigurationUpdate() // ConfigurationUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationsApi.UpdateConfiguration(context.Background(), configurationId).ConfigurationUpdate(configurationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: ConfigurationDetail
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | Configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationUpdate** | [**ConfigurationUpdate**](ConfigurationUpdate.md) |  | 

### Return type

[**ConfigurationDetail**](ConfigurationDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

