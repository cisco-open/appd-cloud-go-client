# appdynamicscloud-go-client

<!-- TODO: Update url -->
This repository contains the golang client SDK to interact with AppDynamics Cloud using REST API calls. This SDK is used by [terraform-provider-appdynamicscloud](https://github.com/AniketK-Crest/terraform-provider-appdynamicscloud).

## Installation

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or project's Go module dependencies:

<!-- TODO: Update url -->
```shell
go get github.com/aniketk-crest/appdynamicscloud-go-client
```

There are no additional dependencies that need to be installed.

## Overview
  
* <strong>apis</strong> :- This package contains the client package of a specific API version generated by the OpenAPI client generator.

* <strong>custom-templates</strong> :- This folder contains the custom mustache template files.

* <strong>.openapi-generator-ignore</strong> :- This file contains the information about individual files or directories that can be ignored.

* <strong>apis.txt</strong> :- This file contains the URLs of the Open API Spec for the AppDynamics Cloud.

* <strong>config.yaml</strong> :- This file contains all the additional configuration settings that are needed to generate Go Client SDK.

* <strong>configuration.go</strong> :- This file contains all configuration of the API Client.

* <strong>script.sh</strong> :- This file contains the script which leads to the generation of the Go Client SDK.

## How to Use

Get the specific API Client through which REST call can be made to access the resource. Below is the example of fetching a Connection for a specific `CONNECTION_ID` using Cloud Connection API Client.

<!-- TODO: Update url -->
```golang
import (
    "fmt"
    "context"
    "strings"
    client "github.com/aniketk-crest/appdynamicscloud-go-client"
    cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
)

// Get the configuration object
configuration := client.NewConfiguration()

// ContextServerVariables override the server configuration variables
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
    "tenant-name": tenantName,
})
// Set Access Token in the context required for authentication for the request
ctx = context.WithValue(ctx, client.ContextAccessToken, "TOKEN")

// Create a new API Client for accessing the Cloud Connection API.
apiClient := cloudconnectionapi.NewAPIClient(configuration)

// Set ContextServerIndex in the context which uses a server configuration for the index of Cloud Connection.
ctx := context.WithValue(ctx, client.ContextServerIndex, client.SERVER_INDEX_CLOUD_CONNECTION)

// GET Request using API Client to get the connection information for a specific connection id.
resp, httpResp, err := apiClient.ConnectionsApi.GetConnection(ctx, "CONNECTION_ID").Execute()
if err != nil {
    fmt.Printf("Err: %v", err)
}

```

## Adding New Functionality

In case of adding the new functionality support for new API, you will need to update `apis.txt` file with the API Spec link in the file. [OpenAPI Generator](https://openapi-generator.tech/) is responsible for generation of Go Client SDK as per the changes in the `apis.txt` file. 
