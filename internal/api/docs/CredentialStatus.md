# CredentialStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creds** | [**[]CredentialDetail**](CredentialDetail.md) |  | 

## Methods

### NewCredentialStatus

`func NewCredentialStatus(creds []CredentialDetail, ) *CredentialStatus`

NewCredentialStatus instantiates a new CredentialStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialStatusWithDefaults

`func NewCredentialStatusWithDefaults() *CredentialStatus`

NewCredentialStatusWithDefaults instantiates a new CredentialStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreds

`func (o *CredentialStatus) GetCreds() []CredentialDetail`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *CredentialStatus) GetCredsOk() (*[]CredentialDetail, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *CredentialStatus) SetCreds(v []CredentialDetail)`

SetCreds sets Creds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


