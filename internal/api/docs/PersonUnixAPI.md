# \PersonUnixAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PersonIdUnixCredentialDelete**](PersonUnixAPI.md#PersonIdUnixCredentialDelete) | **Delete** /v1/person/{id}/_unix/_credential | 
[**PersonIdUnixCredentialPut**](PersonUnixAPI.md#PersonIdUnixCredentialPut) | **Put** /v1/person/{id}/_unix/_credential | 
[**PersonIdUnixPost**](PersonUnixAPI.md#PersonIdUnixPost) | **Post** /v1/person/{id}/_unix | 



## PersonIdUnixCredentialDelete

> PersonIdUnixCredentialDelete(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/slop-incubator/go-kanidm/internal/api"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonUnixAPI.PersonIdUnixCredentialDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonUnixAPI.PersonIdUnixCredentialDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonIdUnixCredentialDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonIdUnixCredentialPut

> PersonIdUnixCredentialPut(ctx, id).SingleStringRequest(singleStringRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/slop-incubator/go-kanidm/internal/api"
)

func main() {
	id := "id_example" // string | 
	singleStringRequest := *openapiclient.NewSingleStringRequest("Value_example") // SingleStringRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonUnixAPI.PersonIdUnixCredentialPut(context.Background(), id).SingleStringRequest(singleStringRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonUnixAPI.PersonIdUnixCredentialPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonIdUnixCredentialPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **singleStringRequest** | [**SingleStringRequest**](SingleStringRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonIdUnixPost

> PersonIdUnixPost(ctx, id).AccountUnixExtend(accountUnixExtend).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/slop-incubator/go-kanidm/internal/api"
)

func main() {
	id := "id_example" // string | 
	accountUnixExtend := *openapiclient.NewAccountUnixExtend() // AccountUnixExtend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonUnixAPI.PersonIdUnixPost(context.Background(), id).AccountUnixExtend(accountUnixExtend).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonUnixAPI.PersonIdUnixPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonIdUnixPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountUnixExtend** | [**AccountUnixExtend**](AccountUnixExtend.md) |  | 

### Return type

 (empty response body)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

