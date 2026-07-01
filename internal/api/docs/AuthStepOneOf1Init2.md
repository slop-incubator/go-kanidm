# AuthStepOneOf1Init2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Issue** | [**AuthIssueSession**](AuthIssueSession.md) |  | 
**Privileged** | Pointer to **bool** | If true, the session will have r/w access. | [optional] 

## Methods

### NewAuthStepOneOf1Init2

`func NewAuthStepOneOf1Init2(username string, issue AuthIssueSession, ) *AuthStepOneOf1Init2`

NewAuthStepOneOf1Init2 instantiates a new AuthStepOneOf1Init2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthStepOneOf1Init2WithDefaults

`func NewAuthStepOneOf1Init2WithDefaults() *AuthStepOneOf1Init2`

NewAuthStepOneOf1Init2WithDefaults instantiates a new AuthStepOneOf1Init2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AuthStepOneOf1Init2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthStepOneOf1Init2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthStepOneOf1Init2) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetIssue

`func (o *AuthStepOneOf1Init2) GetIssue() AuthIssueSession`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *AuthStepOneOf1Init2) GetIssueOk() (*AuthIssueSession, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *AuthStepOneOf1Init2) SetIssue(v AuthIssueSession)`

SetIssue sets Issue field to given value.


### GetPrivileged

`func (o *AuthStepOneOf1Init2) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *AuthStepOneOf1Init2) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *AuthStepOneOf1Init2) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *AuthStepOneOf1Init2) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


