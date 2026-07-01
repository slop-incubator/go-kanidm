# \PersonSshPubkeysAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PersonIdSshPubkeysGet**](PersonSshPubkeysAPI.md#PersonIdSshPubkeysGet) | **Get** /v1/person/{id}/_ssh_pubkeys | 
[**PersonIdSshPubkeysPost**](PersonSshPubkeysAPI.md#PersonIdSshPubkeysPost) | **Post** /v1/person/{id}/_ssh_pubkeys | 
[**PersonIdSshPubkeysTagDelete**](PersonSshPubkeysAPI.md#PersonIdSshPubkeysTagDelete) | **Delete** /v1/person/{id}/_ssh_pubkeys/{tag} | 
[**PersonIdSshPubkeysTagGet**](PersonSshPubkeysAPI.md#PersonIdSshPubkeysTagGet) | **Get** /v1/person/{id}/_ssh_pubkeys/{tag} | 



## PersonIdSshPubkeysGet

> []string PersonIdSshPubkeysGet(ctx, id).Execute()



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
	resp, r, err := apiClient.PersonSshPubkeysAPI.PersonIdSshPubkeysGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonSshPubkeysAPI.PersonIdSshPubkeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PersonIdSshPubkeysGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `PersonSshPubkeysAPI.PersonIdSshPubkeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonIdSshPubkeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonIdSshPubkeysPost

> PersonIdSshPubkeysPost(ctx, id).RequestBody(requestBody).Execute()



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
	requestBody := []interface{}{interface{}(123)} // []interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonSshPubkeysAPI.PersonIdSshPubkeysPost(context.Background(), id).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonSshPubkeysAPI.PersonIdSshPubkeysPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPersonIdSshPubkeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonIdSshPubkeysTagDelete

> PersonIdSshPubkeysTagDelete(ctx, tag, id).Execute()



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
	tag := "tag_example" // string | The tag of the SSH key
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonSshPubkeysAPI.PersonIdSshPubkeysTagDelete(context.Background(), tag, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonSshPubkeysAPI.PersonIdSshPubkeysTagDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** | The tag of the SSH key | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonIdSshPubkeysTagDeleteRequest struct via the builder pattern


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


## PersonIdSshPubkeysTagGet

> string PersonIdSshPubkeysTagGet(ctx, id, tag).Execute()



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
	tag := "tag_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonSshPubkeysAPI.PersonIdSshPubkeysTagGet(context.Background(), id, tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonSshPubkeysAPI.PersonIdSshPubkeysTagGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PersonIdSshPubkeysTagGet`: string
	fmt.Fprintf(os.Stdout, "Response from `PersonSshPubkeysAPI.PersonIdSshPubkeysTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonIdSshPubkeysTagGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

