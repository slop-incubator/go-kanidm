# CURegState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotpCheck** | [**TotpSecret**](TotpSecret.md) |  | 
**TotpNameTryAgain** | **string** |  | 
**BackupCodes** | **[]string** |  | 
**Passkey** | **map[string]interface{}** |  | 
**AttestedPasskey** | **map[string]interface{}** |  | 

## Methods

### NewCURegState

`func NewCURegState(totpCheck TotpSecret, totpNameTryAgain string, backupCodes []string, passkey map[string]interface{}, attestedPasskey map[string]interface{}, ) *CURegState`

NewCURegState instantiates a new CURegState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCURegStateWithDefaults

`func NewCURegStateWithDefaults() *CURegState`

NewCURegStateWithDefaults instantiates a new CURegState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotpCheck

`func (o *CURegState) GetTotpCheck() TotpSecret`

GetTotpCheck returns the TotpCheck field if non-nil, zero value otherwise.

### GetTotpCheckOk

`func (o *CURegState) GetTotpCheckOk() (*TotpSecret, bool)`

GetTotpCheckOk returns a tuple with the TotpCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCheck

`func (o *CURegState) SetTotpCheck(v TotpSecret)`

SetTotpCheck sets TotpCheck field to given value.


### GetTotpNameTryAgain

`func (o *CURegState) GetTotpNameTryAgain() string`

GetTotpNameTryAgain returns the TotpNameTryAgain field if non-nil, zero value otherwise.

### GetTotpNameTryAgainOk

`func (o *CURegState) GetTotpNameTryAgainOk() (*string, bool)`

GetTotpNameTryAgainOk returns a tuple with the TotpNameTryAgain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpNameTryAgain

`func (o *CURegState) SetTotpNameTryAgain(v string)`

SetTotpNameTryAgain sets TotpNameTryAgain field to given value.


### GetBackupCodes

`func (o *CURegState) GetBackupCodes() []string`

GetBackupCodes returns the BackupCodes field if non-nil, zero value otherwise.

### GetBackupCodesOk

`func (o *CURegState) GetBackupCodesOk() (*[]string, bool)`

GetBackupCodesOk returns a tuple with the BackupCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCodes

`func (o *CURegState) SetBackupCodes(v []string)`

SetBackupCodes sets BackupCodes field to given value.


### GetPasskey

`func (o *CURegState) GetPasskey() map[string]interface{}`

GetPasskey returns the Passkey field if non-nil, zero value otherwise.

### GetPasskeyOk

`func (o *CURegState) GetPasskeyOk() (*map[string]interface{}, bool)`

GetPasskeyOk returns a tuple with the Passkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasskey

`func (o *CURegState) SetPasskey(v map[string]interface{})`

SetPasskey sets Passkey field to given value.


### GetAttestedPasskey

`func (o *CURegState) GetAttestedPasskey() map[string]interface{}`

GetAttestedPasskey returns the AttestedPasskey field if non-nil, zero value otherwise.

### GetAttestedPasskeyOk

`func (o *CURegState) GetAttestedPasskeyOk() (*map[string]interface{}, bool)`

GetAttestedPasskeyOk returns a tuple with the AttestedPasskey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestedPasskey

`func (o *CURegState) SetAttestedPasskey(v map[string]interface{})`

SetAttestedPasskey sets AttestedPasskey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


