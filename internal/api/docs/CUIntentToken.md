# CUIntentToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**ExpiryTime** | **time.Time** |  | 

## Methods

### NewCUIntentToken

`func NewCUIntentToken(token string, expiryTime time.Time, ) *CUIntentToken`

NewCUIntentToken instantiates a new CUIntentToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCUIntentTokenWithDefaults

`func NewCUIntentTokenWithDefaults() *CUIntentToken`

NewCUIntentTokenWithDefaults instantiates a new CUIntentToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CUIntentToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CUIntentToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CUIntentToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiryTime

`func (o *CUIntentToken) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *CUIntentToken) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *CUIntentToken) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


