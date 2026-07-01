# \PersonCredentialAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PersonGetIdCredentialStatus**](PersonCredentialAPI.md#PersonGetIdCredentialStatus) | **Get** /v1/person/{id}/_credential/_status | 
[**PersonIdCredentialUpdateGet**](PersonCredentialAPI.md#PersonIdCredentialUpdateGet) | **Get** /v1/person/{id}/_credential/_update | 
[**PersonIdCredentialUpdateIntentGet**](PersonCredentialAPI.md#PersonIdCredentialUpdateIntentGet) | **Get** /v1/person/{id}/_credential/_update_intent | 
[**PersonIdCredentialUpdateIntentSendPost**](PersonCredentialAPI.md#PersonIdCredentialUpdateIntentSendPost) | **Post** /v1/person/{id}/_credential/_update_intent_send | 
[**PersonIdCredentialUpdateIntentTtlGet**](PersonCredentialAPI.md#PersonIdCredentialUpdateIntentTtlGet) | **Get** /v1/person/{id}/_credential/_update_intent/{ttl} | 



## PersonGetIdCredentialStatus

> PersonGetIdCredentialStatus(ctx, id).Execute()



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
	r, err := apiClient.PersonCredentialAPI.PersonGetIdCredentialStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonCredentialAPI.PersonGetIdCredentialStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPersonGetIdCredentialStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonIdCredentialUpdateGet

> PersonIdCredentialUpdateGet(ctx, id).Execute()



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
	r, err := apiClient.PersonCredentialAPI.PersonIdCredentialUpdateGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonCredentialAPI.PersonIdCredentialUpdateGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPersonIdCredentialUpdateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonIdCredentialUpdateIntentGet

> PersonIdCredentialUpdateIntentGet(ctx, id).Execute()



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
	r, err := apiClient.PersonCredentialAPI.PersonIdCredentialUpdateIntentGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonCredentialAPI.PersonIdCredentialUpdateIntentGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPersonIdCredentialUpdateIntentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonIdCredentialUpdateIntentSendPost

> PersonIdCredentialUpdateIntentSendPost(ctx, id).CUIntentSend(cUIntentSend).Execute()



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
	cUIntentSend := *openapiclient.NewCUIntentSend() // CUIntentSend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonCredentialAPI.PersonIdCredentialUpdateIntentSendPost(context.Background(), id).CUIntentSend(cUIntentSend).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonCredentialAPI.PersonIdCredentialUpdateIntentSendPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPersonIdCredentialUpdateIntentSendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cUIntentSend** | [**CUIntentSend**](CUIntentSend.md) |  | 

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


## PersonIdCredentialUpdateIntentTtlGet

> CUIntentToken PersonIdCredentialUpdateIntentTtlGet(ctx, ttl, id).Execute()



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
	ttl := int64(789) // int64 | The new TTL for the credential?
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonCredentialAPI.PersonIdCredentialUpdateIntentTtlGet(context.Background(), ttl, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonCredentialAPI.PersonIdCredentialUpdateIntentTtlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PersonIdCredentialUpdateIntentTtlGet`: CUIntentToken
	fmt.Fprintf(os.Stdout, "Response from `PersonCredentialAPI.PersonIdCredentialUpdateIntentTtlGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ttl** | **int64** | The new TTL for the credential? | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonIdCredentialUpdateIntentTtlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CUIntentToken**](CUIntentToken.md)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

