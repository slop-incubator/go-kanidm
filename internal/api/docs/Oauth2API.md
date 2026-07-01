# \Oauth2API

All URIs are relative to *https://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2BasicPost**](Oauth2API.md#Oauth2BasicPost) | **Post** /v1/oauth2/_basic | Create a new Confidential OAuth2 client that authenticates with Http Basic.
[**Oauth2Get**](Oauth2API.md#Oauth2Get) | **Get** /v1/oauth2 | Lists all the OAuth2 Resource Servers
[**Oauth2IdClaimmapDelete**](Oauth2API.md#Oauth2IdClaimmapDelete) | **Delete** /v1/oauth2/{rs_name}/_claimmap/{claim_name}/{group} | 
[**Oauth2IdClaimmapJoinPost**](Oauth2API.md#Oauth2IdClaimmapJoinPost) | **Post** /v1/oauth2/{rs_name}/_claimmap/{claim_name} | Modify the claim map join strategy for a given OAuth2 Resource Server
[**Oauth2IdClaimmapPost**](Oauth2API.md#Oauth2IdClaimmapPost) | **Post** /v1/oauth2/{rs_name}/_claimmap/{claim_name}/{group} | Modify the claim map for a given OAuth2 Resource Server
[**Oauth2IdDelete**](Oauth2API.md#Oauth2IdDelete) | **Delete** /v1/oauth2/{rs_name} | Delete an OAuth2 Resource Server
[**Oauth2IdGet**](Oauth2API.md#Oauth2IdGet) | **Get** /v1/oauth2/{rs_name} | Get the details of a given OAuth2 Resource Server.
[**Oauth2IdGetBasicSecret**](Oauth2API.md#Oauth2IdGetBasicSecret) | **Get** /v1/oauth2/{rs_name}/_basic_secret | Get the basic secret for a given OAuth2 Resource Server. This is used for authentication.
[**Oauth2IdImageDelete**](Oauth2API.md#Oauth2IdImageDelete) | **Delete** /v1/oauth2/{rs_name}/_image | 
[**Oauth2IdImagePost**](Oauth2API.md#Oauth2IdImagePost) | **Post** /v1/oauth2/{rs_name}/_image | API endpoint for creating/replacing the image associated with an OAuth2 Resource Server.
[**Oauth2IdPatch**](Oauth2API.md#Oauth2IdPatch) | **Patch** /v1/oauth2/{rs_name} | Modify an OAuth2 Resource Server
[**Oauth2IdScopemapDelete**](Oauth2API.md#Oauth2IdScopemapDelete) | **Delete** /v1/oauth2/{rs_name}/_scopemap/{group} | 
[**Oauth2IdScopemapPost**](Oauth2API.md#Oauth2IdScopemapPost) | **Post** /v1/oauth2/{rs_name}/_scopemap/{group} | Modify the scope map for a given OAuth2 Resource Server
[**Oauth2IdSupScopemapDelete**](Oauth2API.md#Oauth2IdSupScopemapDelete) | **Delete** /v1/oauth2/{rs_name}/_sup_scopemap/{group} | 
[**Oauth2IdSupScopemapPost**](Oauth2API.md#Oauth2IdSupScopemapPost) | **Post** /v1/oauth2/{rs_name}/_sup_scopemap/{group} | Create a supplemental scope map for a given OAuth2 Resource Server
[**Oauth2PublicPost**](Oauth2API.md#Oauth2PublicPost) | **Post** /v1/oauth2/_public | Create a new Public OAuth2 client



## Oauth2BasicPost

> Oauth2BasicPost(ctx).Entry(entry).Execute()

Create a new Confidential OAuth2 client that authenticates with Http Basic.

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
	r, err := apiClient.Oauth2API.Oauth2BasicPost(context.Background()).Entry(entry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2BasicPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2BasicPostRequest struct via the builder pattern


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


## Oauth2Get

> []Entry Oauth2Get(ctx).Execute()

Lists all the OAuth2 Resource Servers

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
	resp, r, err := apiClient.Oauth2API.Oauth2Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Oauth2Get`: []Entry
	fmt.Fprintf(os.Stdout, "Response from `Oauth2API.Oauth2Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2GetRequest struct via the builder pattern


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


## Oauth2IdClaimmapDelete

> Oauth2IdClaimmapDelete(ctx, rsName, claimName, group).Execute()



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
	rsName := "rsName_example" // string | 
	claimName := "claimName_example" // string | 
	group := "group_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdClaimmapDelete(context.Background(), rsName, claimName, group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdClaimmapDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 
**claimName** | **string** |  | 
**group** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdClaimmapDeleteRequest struct via the builder pattern


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


## Oauth2IdClaimmapJoinPost

> Oauth2IdClaimmapJoinPost(ctx, rsName, claimName).Body(body).Execute()

Modify the claim map join strategy for a given OAuth2 Resource Server

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
	rsName := "rsName_example" // string | 
	claimName := "claimName_example" // string | 
	body := string(987) // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdClaimmapJoinPost(context.Background(), rsName, claimName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdClaimmapJoinPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 
**claimName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdClaimmapJoinPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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


## Oauth2IdClaimmapPost

> Oauth2IdClaimmapPost(ctx, rsName, claimName, group).RequestBody(requestBody).Execute()

Modify the claim map for a given OAuth2 Resource Server

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
	rsName := "rsName_example" // string | 
	claimName := "claimName_example" // string | 
	group := "group_example" // string | 
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdClaimmapPost(context.Background(), rsName, claimName, group).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdClaimmapPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 
**claimName** | **string** |  | 
**group** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdClaimmapPostRequest struct via the builder pattern


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


## Oauth2IdDelete

> Oauth2IdDelete(ctx, rsName).Execute()

Delete an OAuth2 Resource Server

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
	rsName := "rsName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdDelete(context.Background(), rsName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdDeleteRequest struct via the builder pattern


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


## Oauth2IdGet

> Entry Oauth2IdGet(ctx, rsName).Execute()

Get the details of a given OAuth2 Resource Server.

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
	rsName := "rsName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Oauth2API.Oauth2IdGet(context.Background(), rsName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Oauth2IdGet`: Entry
	fmt.Fprintf(os.Stdout, "Response from `Oauth2API.Oauth2IdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdGetRequest struct via the builder pattern


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


## Oauth2IdGetBasicSecret

> string Oauth2IdGetBasicSecret(ctx, rsName).Execute()

Get the basic secret for a given OAuth2 Resource Server. This is used for authentication.

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
	rsName := "rsName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Oauth2API.Oauth2IdGetBasicSecret(context.Background(), rsName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdGetBasicSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Oauth2IdGetBasicSecret`: string
	fmt.Fprintf(os.Stdout, "Response from `Oauth2API.Oauth2IdGetBasicSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdGetBasicSecretRequest struct via the builder pattern


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


## Oauth2IdImageDelete

> Oauth2IdImageDelete(ctx, rsName).Execute()



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
	rsName := "rsName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdImageDelete(context.Background(), rsName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdImageDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdImageDeleteRequest struct via the builder pattern


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


## Oauth2IdImagePost

> Oauth2IdImagePost(ctx, rsName).Execute()

API endpoint for creating/replacing the image associated with an OAuth2 Resource Server.



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
	rsName := "rsName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdImagePost(context.Background(), rsName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdImagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdImagePostRequest struct via the builder pattern


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


## Oauth2IdPatch

> Oauth2IdPatch(ctx, rsName).Entry(entry).Execute()

Modify an OAuth2 Resource Server

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
	rsName := "rsName_example" // string | 
	entry := *openapiclient.NewEntry(map[string][]string{"key": []string{"Inner_example"}}) // Entry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdPatch(context.Background(), rsName).Entry(entry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdPatchRequest struct via the builder pattern


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


## Oauth2IdScopemapDelete

> Oauth2IdScopemapDelete(ctx, rsName, group).Execute()



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
	rsName := "rsName_example" // string | 
	group := "group_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdScopemapDelete(context.Background(), rsName, group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdScopemapDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 
**group** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdScopemapDeleteRequest struct via the builder pattern


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


## Oauth2IdScopemapPost

> Oauth2IdScopemapPost(ctx, rsName, group).RequestBody(requestBody).Execute()

Modify the scope map for a given OAuth2 Resource Server

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
	rsName := "rsName_example" // string | 
	group := "group_example" // string | 
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdScopemapPost(context.Background(), rsName, group).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdScopemapPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 
**group** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdScopemapPostRequest struct via the builder pattern


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


## Oauth2IdSupScopemapDelete

> Oauth2IdSupScopemapDelete(ctx, rsName, group).Execute()



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
	rsName := "rsName_example" // string | 
	group := "group_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdSupScopemapDelete(context.Background(), rsName, group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdSupScopemapDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 
**group** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdSupScopemapDeleteRequest struct via the builder pattern


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


## Oauth2IdSupScopemapPost

> Oauth2IdSupScopemapPost(ctx, rsName, group).RequestBody(requestBody).Execute()

Create a supplemental scope map for a given OAuth2 Resource Server

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
	rsName := "rsName_example" // string | 
	group := "group_example" // string | 
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.Oauth2IdSupScopemapPost(context.Background(), rsName, group).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2IdSupScopemapPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rsName** | **string** |  | 
**group** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2IdSupScopemapPostRequest struct via the builder pattern


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


## Oauth2PublicPost

> Oauth2PublicPost(ctx).Entry(entry).Execute()

Create a new Public OAuth2 client

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
	r, err := apiClient.Oauth2API.Oauth2PublicPost(context.Background()).Entry(entry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.Oauth2PublicPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2PublicPostRequest struct via the builder pattern


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

