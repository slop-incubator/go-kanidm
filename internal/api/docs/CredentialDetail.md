# CredentialDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Type** | [**CredentialDetailType**](CredentialDetailType.md) |  | 

## Methods

### NewCredentialDetail

`func NewCredentialDetail(uuid string, type_ CredentialDetailType, ) *CredentialDetail`

NewCredentialDetail instantiates a new CredentialDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialDetailWithDefaults

`func NewCredentialDetailWithDefaults() *CredentialDetail`

NewCredentialDetailWithDefaults instantiates a new CredentialDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *CredentialDetail) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CredentialDetail) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CredentialDetail) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *CredentialDetail) GetType() CredentialDetailType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialDetail) GetTypeOk() (*CredentialDetailType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialDetail) SetType(v CredentialDetailType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


