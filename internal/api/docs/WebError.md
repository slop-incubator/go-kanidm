# WebError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationError** | [**OperationError**](OperationError.md) | Something went wrong when doing things. | 
**InternalServerError** | **string** |  | 
**OAuth2** | **map[string]interface{}** |  | 

## Methods

### NewWebError

`func NewWebError(operationError OperationError, internalServerError string, oAuth2 map[string]interface{}, ) *WebError`

NewWebError instantiates a new WebError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebErrorWithDefaults

`func NewWebErrorWithDefaults() *WebError`

NewWebErrorWithDefaults instantiates a new WebError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationError

`func (o *WebError) GetOperationError() OperationError`

GetOperationError returns the OperationError field if non-nil, zero value otherwise.

### GetOperationErrorOk

`func (o *WebError) GetOperationErrorOk() (*OperationError, bool)`

GetOperationErrorOk returns a tuple with the OperationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationError

`func (o *WebError) SetOperationError(v OperationError)`

SetOperationError sets OperationError field to given value.


### GetInternalServerError

`func (o *WebError) GetInternalServerError() string`

GetInternalServerError returns the InternalServerError field if non-nil, zero value otherwise.

### GetInternalServerErrorOk

`func (o *WebError) GetInternalServerErrorOk() (*string, bool)`

GetInternalServerErrorOk returns a tuple with the InternalServerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalServerError

`func (o *WebError) SetInternalServerError(v string)`

SetInternalServerError sets InternalServerError field to given value.


### GetOAuth2

`func (o *WebError) GetOAuth2() map[string]interface{}`

GetOAuth2 returns the OAuth2 field if non-nil, zero value otherwise.

### GetOAuth2Ok

`func (o *WebError) GetOAuth2Ok() (*map[string]interface{}, bool)`

GetOAuth2Ok returns a tuple with the OAuth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuth2

`func (o *WebError) SetOAuth2(v map[string]interface{})`

SetOAuth2 sets OAuth2 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


