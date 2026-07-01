# AuthCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**Totp** | **int32** |  | 
**Securitykey** | **map[string]interface{}** |  | 
**Backupcode** | **string** |  | 
**Passkey** | **string** |  | 

## Methods

### NewAuthCredential

`func NewAuthCredential(password string, totp int32, securitykey map[string]interface{}, backupcode string, passkey string, ) *AuthCredential`

NewAuthCredential instantiates a new AuthCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthCredentialWithDefaults

`func NewAuthCredentialWithDefaults() *AuthCredential`

NewAuthCredentialWithDefaults instantiates a new AuthCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *AuthCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthCredential) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTotp

`func (o *AuthCredential) GetTotp() int32`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *AuthCredential) GetTotpOk() (*int32, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *AuthCredential) SetTotp(v int32)`

SetTotp sets Totp field to given value.


### GetSecuritykey

`func (o *AuthCredential) GetSecuritykey() map[string]interface{}`

GetSecuritykey returns the Securitykey field if non-nil, zero value otherwise.

### GetSecuritykeyOk

`func (o *AuthCredential) GetSecuritykeyOk() (*map[string]interface{}, bool)`

GetSecuritykeyOk returns a tuple with the Securitykey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritykey

`func (o *AuthCredential) SetSecuritykey(v map[string]interface{})`

SetSecuritykey sets Securitykey field to given value.


### GetBackupcode

`func (o *AuthCredential) GetBackupcode() string`

GetBackupcode returns the Backupcode field if non-nil, zero value otherwise.

### GetBackupcodeOk

`func (o *AuthCredential) GetBackupcodeOk() (*string, bool)`

GetBackupcodeOk returns a tuple with the Backupcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupcode

`func (o *AuthCredential) SetBackupcode(v string)`

SetBackupcode sets Backupcode field to given value.


### GetPasskey

`func (o *AuthCredential) GetPasskey() string`

GetPasskey returns the Passkey field if non-nil, zero value otherwise.

### GetPasskeyOk

`func (o *AuthCredential) GetPasskeyOk() (*string, bool)`

GetPasskeyOk returns a tuple with the Passkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasskey

`func (o *AuthCredential) SetPasskey(v string)`

SetPasskey sets Passkey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


