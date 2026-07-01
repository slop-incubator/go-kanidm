# AuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessionid** | **string** |  | 
**State** | [**AuthState**](AuthState.md) |  | 

## Methods

### NewAuthResponse

`func NewAuthResponse(sessionid string, state AuthState, ) *AuthResponse`

NewAuthResponse instantiates a new AuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthResponseWithDefaults

`func NewAuthResponseWithDefaults() *AuthResponse`

NewAuthResponseWithDefaults instantiates a new AuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionid

`func (o *AuthResponse) GetSessionid() string`

GetSessionid returns the Sessionid field if non-nil, zero value otherwise.

### GetSessionidOk

`func (o *AuthResponse) GetSessionidOk() (*string, bool)`

GetSessionidOk returns a tuple with the Sessionid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionid

`func (o *AuthResponse) SetSessionid(v string)`

SetSessionid sets Sessionid field to given value.


### GetState

`func (o *AuthResponse) GetState() AuthState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AuthResponse) GetStateOk() (*AuthState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AuthResponse) SetState(v AuthState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


