# AuthStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Init** | **string** | Initialise a new authentication session | 
**Init2** | [**AuthStepOneOf1Init2**](AuthStepOneOf1Init2.md) |  | 
**Begin** | [**AuthMech**](AuthMech.md) | Request the named authentication mechanism to proceed | 
**Cred** | [**AuthCredential**](AuthCredential.md) | Provide a credential in response to a challenge | 

## Methods

### NewAuthStep

`func NewAuthStep(init string, init2 AuthStepOneOf1Init2, begin AuthMech, cred AuthCredential, ) *AuthStep`

NewAuthStep instantiates a new AuthStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthStepWithDefaults

`func NewAuthStepWithDefaults() *AuthStep`

NewAuthStepWithDefaults instantiates a new AuthStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInit

`func (o *AuthStep) GetInit() string`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *AuthStep) GetInitOk() (*string, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *AuthStep) SetInit(v string)`

SetInit sets Init field to given value.


### GetInit2

`func (o *AuthStep) GetInit2() AuthStepOneOf1Init2`

GetInit2 returns the Init2 field if non-nil, zero value otherwise.

### GetInit2Ok

`func (o *AuthStep) GetInit2Ok() (*AuthStepOneOf1Init2, bool)`

GetInit2Ok returns a tuple with the Init2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit2

`func (o *AuthStep) SetInit2(v AuthStepOneOf1Init2)`

SetInit2 sets Init2 field to given value.


### GetBegin

`func (o *AuthStep) GetBegin() AuthMech`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *AuthStep) GetBeginOk() (*AuthMech, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *AuthStep) SetBegin(v AuthMech)`

SetBegin sets Begin field to given value.


### GetCred

`func (o *AuthStep) GetCred() AuthCredential`

GetCred returns the Cred field if non-nil, zero value otherwise.

### GetCredOk

`func (o *AuthStep) GetCredOk() (*AuthCredential, bool)`

GetCredOk returns a tuple with the Cred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCred

`func (o *AuthStep) SetCred(v AuthCredential)`

SetCred sets Cred field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


