# CredentialDetailType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passkey** | **[]string** |  | 
**PasswordMfa** | **[]map[string]interface{}** | totp, webauthn | 

## Methods

### NewCredentialDetailType

`func NewCredentialDetailType(passkey []string, passwordMfa []map[string]interface{}, ) *CredentialDetailType`

NewCredentialDetailType instantiates a new CredentialDetailType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialDetailTypeWithDefaults

`func NewCredentialDetailTypeWithDefaults() *CredentialDetailType`

NewCredentialDetailTypeWithDefaults instantiates a new CredentialDetailType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasskey

`func (o *CredentialDetailType) GetPasskey() []string`

GetPasskey returns the Passkey field if non-nil, zero value otherwise.

### GetPasskeyOk

`func (o *CredentialDetailType) GetPasskeyOk() (*[]string, bool)`

GetPasskeyOk returns a tuple with the Passkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasskey

`func (o *CredentialDetailType) SetPasskey(v []string)`

SetPasskey sets Passkey field to given value.


### GetPasswordMfa

`func (o *CredentialDetailType) GetPasswordMfa() []map[string]interface{}`

GetPasswordMfa returns the PasswordMfa field if non-nil, zero value otherwise.

### GetPasswordMfaOk

`func (o *CredentialDetailType) GetPasswordMfaOk() (*[]map[string]interface{}, bool)`

GetPasswordMfaOk returns a tuple with the PasswordMfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMfa

`func (o *CredentialDetailType) SetPasswordMfa(v []map[string]interface{})`

SetPasswordMfa sets PasswordMfa field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


