# SshPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyType** | **string** |  | 
**Kind** | **string** |  | 
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSshPublicKey

`func NewSshPublicKey(keyType string, kind string, ) *SshPublicKey`

NewSshPublicKey instantiates a new SshPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshPublicKeyWithDefaults

`func NewSshPublicKeyWithDefaults() *SshPublicKey`

NewSshPublicKeyWithDefaults instantiates a new SshPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyType

`func (o *SshPublicKey) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SshPublicKey) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SshPublicKey) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKind

`func (o *SshPublicKey) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SshPublicKey) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SshPublicKey) SetKind(v string)`

SetKind sets Kind field to given value.


### GetComment

`func (o *SshPublicKey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SshPublicKey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SshPublicKey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SshPublicKey) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *SshPublicKey) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *SshPublicKey) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


