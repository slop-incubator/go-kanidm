# \SyncAccountAPI

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncAccountGet**](SyncAccountAPI.md#SyncAccountGet) | **Get** /v1/sync_account | Get all? the sync accounts.
[**SyncAccountIdAttrGet**](SyncAccountAPI.md#SyncAccountIdAttrGet) | **Get** /v1/sync_account/{id}/_attr/{attr} | 
[**SyncAccountIdAttrPut**](SyncAccountAPI.md#SyncAccountIdAttrPut) | **Post** /v1/sync_account/{id}/_attr/{attr} | 
[**SyncAccountIdFinaliseGet**](SyncAccountAPI.md#SyncAccountIdFinaliseGet) | **Get** /v1/sync_account/{id}/_finalise | 
[**SyncAccountIdGet**](SyncAccountAPI.md#SyncAccountIdGet) | **Get** /v1/sync_account/{id} | Get the details of a sync account
[**SyncAccountIdPatch**](SyncAccountAPI.md#SyncAccountIdPatch) | **Patch** /v1/sync_account/{id} | Modify a sync account in-place
[**SyncAccountIdTerminateGet**](SyncAccountAPI.md#SyncAccountIdTerminateGet) | **Get** /v1/sync_account/{id}/_terminate | 
[**SyncAccountPost**](SyncAccountAPI.md#SyncAccountPost) | **Post** /v1/sync_account | 
[**SyncAccountTokenDelete**](SyncAccountAPI.md#SyncAccountTokenDelete) | **Delete** /v1/sync_account/{id}/_sync_token | 
[**SyncAccountTokenPost**](SyncAccountAPI.md#SyncAccountTokenPost) | **Post** /v1/sync_account/{id}/_sync_token | 



## SyncAccountGet

> []Entry SyncAccountGet(ctx).Execute()

Get all? the sync accounts.

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
	resp, r, err := apiClient.SyncAccountAPI.SyncAccountGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncAccountGet`: []Entry
	fmt.Fprintf(os.Stdout, "Response from `SyncAccountAPI.SyncAccountGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAccountGetRequest struct via the builder pattern


### Return type

[**[]Entry**](Entry.md)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAccountIdAttrGet

> []string SyncAccountIdAttrGet(ctx, id, attr).Execute()



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
	attr := "attr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAccountAPI.SyncAccountIdAttrGet(context.Background(), id, attr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountIdAttrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncAccountIdAttrGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `SyncAccountAPI.SyncAccountIdAttrGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAccountIdAttrGetRequest struct via the builder pattern


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


## SyncAccountIdAttrPut

> SyncAccountIdAttrPut(ctx, id, attr).RequestBody(requestBody).Execute()



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
	attr := "attr_example" // string | 
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncAccountAPI.SyncAccountIdAttrPut(context.Background(), id, attr).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountIdAttrPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAccountIdAttrPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** |  | 

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


## SyncAccountIdFinaliseGet

> SyncAccountIdFinaliseGet(ctx, id).Execute()



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
	r, err := apiClient.SyncAccountAPI.SyncAccountIdFinaliseGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountIdFinaliseGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSyncAccountIdFinaliseGetRequest struct via the builder pattern


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


## SyncAccountIdGet

> Entry SyncAccountIdGet(ctx, id).Execute()

Get the details of a sync account

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
	resp, r, err := apiClient.SyncAccountAPI.SyncAccountIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncAccountIdGet`: Entry
	fmt.Fprintf(os.Stdout, "Response from `SyncAccountAPI.SyncAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Entry**](Entry.md)

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAccountIdPatch

> SyncAccountIdPatch(ctx, id).Entry(entry).Execute()

Modify a sync account in-place

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
	entry := *openapiclient.NewEntry(map[string][]string{"key": []string{"Inner_example"}}) // Entry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncAccountAPI.SyncAccountIdPatch(context.Background(), id).Entry(entry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountIdPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSyncAccountIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entry** | [**Entry**](Entry.md) |  | 

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


## SyncAccountIdTerminateGet

> SyncAccountIdTerminateGet(ctx, id).Execute()



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
	r, err := apiClient.SyncAccountAPI.SyncAccountIdTerminateGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountIdTerminateGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSyncAccountIdTerminateGetRequest struct via the builder pattern


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


## SyncAccountPost

> SyncAccountPost(ctx).Entry(entry).Execute()



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
	entry := *openapiclient.NewEntry(map[string][]string{"key": []string{"Inner_example"}}) // Entry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncAccountAPI.SyncAccountPost(context.Background()).Entry(entry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entry** | [**Entry**](Entry.md) |  | 

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


## SyncAccountTokenDelete

> SyncAccountTokenDelete(ctx, id).Execute()



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
	r, err := apiClient.SyncAccountAPI.SyncAccountTokenDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountTokenDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSyncAccountTokenDeleteRequest struct via the builder pattern


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


## SyncAccountTokenPost

> string SyncAccountTokenPost(ctx, id).Body(body).Execute()



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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAccountAPI.SyncAccountTokenPost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAccountAPI.SyncAccountTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncAccountTokenPost`: string
	fmt.Fprintf(os.Stdout, "Response from `SyncAccountAPI.SyncAccountTokenPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAccountTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

**string**

### Authorization

[token_jwt](../README.md#token_jwt)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

