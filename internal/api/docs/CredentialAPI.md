# \CredentialAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CredentialUpdateCancel**](CredentialAPI.md#CredentialUpdateCancel) | **Post** /v1/credential/_cancel | 
[**CredentialUpdateCommit**](CredentialAPI.md#CredentialUpdateCommit) | **Post** /v1/credential/_commit | 
[**CredentialUpdateExchangeIntent**](CredentialAPI.md#CredentialUpdateExchangeIntent) | **Post** /v1/credential/_exchange_intent | 
[**CredentialUpdateStatus**](CredentialAPI.md#CredentialUpdateStatus) | **Post** /v1/credential/_status | 
[**CredentialUpdateUpdate**](CredentialAPI.md#CredentialUpdateUpdate) | **Post** /v1/credential/_update | 



## CredentialUpdateCancel

> CredentialUpdateCancel(ctx).CUSessionToken(cUSessionToken).Execute()



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
	cUSessionToken := *openapiclient.NewCUSessionToken("Token_example") // CUSessionToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialAPI.CredentialUpdateCancel(context.Background()).CUSessionToken(cUSessionToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialAPI.CredentialUpdateCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialUpdateCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cUSessionToken** | [**CUSessionToken**](CUSessionToken.md) |  | 

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


## CredentialUpdateCommit

> CredentialUpdateCommit(ctx).CUSessionToken(cUSessionToken).Execute()



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
	cUSessionToken := *openapiclient.NewCUSessionToken("Token_example") // CUSessionToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialAPI.CredentialUpdateCommit(context.Background()).CUSessionToken(cUSessionToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialAPI.CredentialUpdateCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialUpdateCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cUSessionToken** | [**CUSessionToken**](CUSessionToken.md) |  | 

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


## CredentialUpdateExchangeIntent

> CredentialUpdateExchangeIntent(ctx).Body(body).Execute()



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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialAPI.CredentialUpdateExchangeIntent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialAPI.CredentialUpdateExchangeIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialUpdateExchangeIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CredentialUpdateStatus

> CredentialUpdateStatus(ctx).CUSessionToken(cUSessionToken).Execute()



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
	cUSessionToken := *openapiclient.NewCUSessionToken("Token_example") // CUSessionToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialAPI.CredentialUpdateStatus(context.Background()).CUSessionToken(cUSessionToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialAPI.CredentialUpdateStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cUSessionToken** | [**CUSessionToken**](CUSessionToken.md) |  | 

### Return type

 (empty response body)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CredentialUpdateUpdate

> CUStatus CredentialUpdateUpdate(ctx).RequestBody(requestBody).Execute()



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
	requestBody := []interface{}{interface{}(123)} // []interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialAPI.CredentialUpdateUpdate(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialAPI.CredentialUpdateUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CredentialUpdateUpdate`: CUStatus
	fmt.Fprintf(os.Stdout, "Response from `CredentialAPI.CredentialUpdateUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialUpdateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]interface{}** |  | 

### Return type

[**CUStatus**](CUStatus.md)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

