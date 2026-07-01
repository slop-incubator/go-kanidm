# \RawAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RawCreate**](RawAPI.md#RawCreate) | **Post** /v1/raw/create | Raw request to the system, be warned this can be dangerous!
[**RawDelete**](RawAPI.md#RawDelete) | **Post** /v1/raw/delete | Raw request to the system, be warned this can be dangerous!
[**RawModify**](RawAPI.md#RawModify) | **Post** /v1/raw/modify | Raw request to the system, be warned this can be dangerous!
[**RawSearch**](RawAPI.md#RawSearch) | **Post** /v1/raw/search | Raw request to the system, be warned this can be dangerous!



## RawCreate

> RawCreate(ctx).CreateRequest(createRequest).Execute()

Raw request to the system, be warned this can be dangerous!

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
	createRequest := *openapiclient.NewCreateRequest([]openapiclient.Entry{*openapiclient.NewEntry(map[string][]string{"key": []string{"Inner_example"}})}) // CreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RawAPI.RawCreate(context.Background()).CreateRequest(createRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RawAPI.RawCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRawCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequest** | [**CreateRequest**](CreateRequest.md) |  | 

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


## RawDelete

> RawDelete(ctx).DeleteRequest(deleteRequest).Execute()

Raw request to the system, be warned this can be dangerous!

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
	deleteRequest := *openapiclient.NewDeleteRequest(openapiclient.Filter{FilterOneOf: openapiclient.NewFilterOneOf([]string{"Eq_example"})}) // DeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RawAPI.RawDelete(context.Background()).DeleteRequest(deleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RawAPI.RawDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRawDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRequest** | [**DeleteRequest**](DeleteRequest.md) |  | 

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


## RawModify

> RawModify(ctx).ModifyRequest(modifyRequest).Execute()

Raw request to the system, be warned this can be dangerous!

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
	modifyRequest := *openapiclient.NewModifyRequest(openapiclient.Filter{FilterOneOf: openapiclient.NewFilterOneOf([]string{"Eq_example"})}, *openapiclient.NewModifyList([]openapiclient.Modify{openapiclient.Modify{ModifyOneOf: openapiclient.NewModifyOneOf([]string{"Present_example"})}})) // ModifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RawAPI.RawModify(context.Background()).ModifyRequest(modifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RawAPI.RawModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRawModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modifyRequest** | [**ModifyRequest**](ModifyRequest.md) |  | 

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


## RawSearch

> SearchResponse RawSearch(ctx).SearchRequest(searchRequest).Execute()

Raw request to the system, be warned this can be dangerous!

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
	searchRequest := *openapiclient.NewSearchRequest(openapiclient.Filter{FilterOneOf: openapiclient.NewFilterOneOf([]string{"Eq_example"})}) // SearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RawAPI.RawSearch(context.Background()).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RawAPI.RawSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RawSearch`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `RawAPI.RawSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRawSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md) |  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

