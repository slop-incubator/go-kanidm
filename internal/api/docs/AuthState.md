# AuthState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choose** | [**[]AuthMech**](AuthMech.md) | You need to select how you want to proceed. | 
**Continue** | [**[]AuthAllowed**](AuthAllowed.md) | Continue to auth, allowed mechanisms/challenges listed. | 
**Denied** | **string** | Something was bad, your session is terminated and no cookie. | 
**Success** | **string** | Everything is good, your bearer token has been issued and is within. | 

## Methods

### NewAuthState

`func NewAuthState(choose []AuthMech, continue_ []AuthAllowed, denied string, success string, ) *AuthState`

NewAuthState instantiates a new AuthState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthStateWithDefaults

`func NewAuthStateWithDefaults() *AuthState`

NewAuthStateWithDefaults instantiates a new AuthState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoose

`func (o *AuthState) GetChoose() []AuthMech`

GetChoose returns the Choose field if non-nil, zero value otherwise.

### GetChooseOk

`func (o *AuthState) GetChooseOk() (*[]AuthMech, bool)`

GetChooseOk returns a tuple with the Choose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoose

`func (o *AuthState) SetChoose(v []AuthMech)`

SetChoose sets Choose field to given value.


### GetContinue

`func (o *AuthState) GetContinue() []AuthAllowed`

GetContinue returns the Continue field if non-nil, zero value otherwise.

### GetContinueOk

`func (o *AuthState) GetContinueOk() (*[]AuthAllowed, bool)`

GetContinueOk returns a tuple with the Continue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinue

`func (o *AuthState) SetContinue(v []AuthAllowed)`

SetContinue sets Continue field to given value.


### GetDenied

`func (o *AuthState) GetDenied() string`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *AuthState) GetDeniedOk() (*string, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *AuthState) SetDenied(v string)`

SetDenied sets Denied field to given value.


### GetSuccess

`func (o *AuthState) GetSuccess() string`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AuthState) GetSuccessOk() (*string, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AuthState) SetSuccess(v string)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


