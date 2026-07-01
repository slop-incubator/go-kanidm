# AuthStateOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Continue** | [**[]AuthAllowed**](AuthAllowed.md) | Continue to auth, allowed mechanisms/challenges listed. | 

## Methods

### NewAuthStateOneOf1

`func NewAuthStateOneOf1(continue_ []AuthAllowed, ) *AuthStateOneOf1`

NewAuthStateOneOf1 instantiates a new AuthStateOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthStateOneOf1WithDefaults

`func NewAuthStateOneOf1WithDefaults() *AuthStateOneOf1`

NewAuthStateOneOf1WithDefaults instantiates a new AuthStateOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinue

`func (o *AuthStateOneOf1) GetContinue() []AuthAllowed`

GetContinue returns the Continue field if non-nil, zero value otherwise.

### GetContinueOk

`func (o *AuthStateOneOf1) GetContinueOk() (*[]AuthAllowed, bool)`

GetContinueOk returns a tuple with the Continue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinue

`func (o *AuthStateOneOf1) SetContinue(v []AuthAllowed)`

SetContinue sets Continue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


