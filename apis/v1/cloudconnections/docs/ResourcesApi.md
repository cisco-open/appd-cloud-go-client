# \ResourcesApi

All URIs are relative to *http://cloudmonconnectionservice.cloudmon.svc.cluster.local:7778/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegions**](ResourcesApi.md#GetRegions) | **Get** /regions | Get all supported hosting regions
[**GetServices**](ResourcesApi.md#GetServices) | **Get** /services | Get all supported hosting services



## GetRegions

> ConfigurationResourceList GetRegions(ctx).Type_(type_).Execute()

Get all supported hosting regions



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
    type_ := openapiclient.ProviderType("aws") // ProviderType | Cloud hosting provider type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetRegions(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegions`: ConfigurationResourceList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ProviderType**](ProviderType.md) | Cloud hosting provider type | 

### Return type

[**ConfigurationResourceList**](ConfigurationResourceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> ConfigurationResourceList GetServices(ctx).Type_(type_).Execute()

Get all supported hosting services



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
    type_ := openapiclient.ProviderType("aws") // ProviderType | Hosting provider type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.GetServices(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServices`: ConfigurationResourceList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ProviderType**](ProviderType.md) | Hosting provider type | 

### Return type

[**ConfigurationResourceList**](ConfigurationResourceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

