# \SelfAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SelfApplinksGet**](SelfAPI.md#SelfApplinksGet) | **Get** /v1/self/_applinks | Returns your OAuth2 app links for the Web UI
[**Whoami**](SelfAPI.md#Whoami) | **Get** /v1/self | 
[**WhoamiUat**](SelfAPI.md#WhoamiUat) | **Get** /v1/self/_uat | 



## SelfApplinksGet

> []AppLink SelfApplinksGet(ctx).Execute()

Returns your OAuth2 app links for the Web UI

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfAPI.SelfApplinksGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfAPI.SelfApplinksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelfApplinksGet`: []AppLink
	fmt.Fprintf(os.Stdout, "Response from `SelfAPI.SelfApplinksGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSelfApplinksGetRequest struct via the builder pattern


### Return type

[**[]AppLink**](AppLink.md)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Whoami

> WhoamiResponse Whoami(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfAPI.Whoami(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfAPI.Whoami``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Whoami`: WhoamiResponse
	fmt.Fprintf(os.Stdout, "Response from `SelfAPI.Whoami`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWhoamiRequest struct via the builder pattern


### Return type

[**WhoamiResponse**](WhoamiResponse.md)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WhoamiUat

> UserAuthToken WhoamiUat(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfAPI.WhoamiUat(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfAPI.WhoamiUat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WhoamiUat`: UserAuthToken
	fmt.Fprintf(os.Stdout, "Response from `SelfAPI.WhoamiUat`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWhoamiUatRequest struct via the builder pattern


### Return type

[**UserAuthToken**](UserAuthToken.md)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

