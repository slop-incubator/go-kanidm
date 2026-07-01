# ApiTokenGenerate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Expiry** | Pointer to **NullableTime** |  | [optional] 
**ReadWrite** | **bool** |  | 
**Compact** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiTokenGenerate

`func NewApiTokenGenerate(label string, readWrite bool, ) *ApiTokenGenerate`

NewApiTokenGenerate instantiates a new ApiTokenGenerate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenGenerateWithDefaults

`func NewApiTokenGenerateWithDefaults() *ApiTokenGenerate`

NewApiTokenGenerateWithDefaults instantiates a new ApiTokenGenerate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ApiTokenGenerate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiTokenGenerate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiTokenGenerate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetExpiry

`func (o *ApiTokenGenerate) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ApiTokenGenerate) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ApiTokenGenerate) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ApiTokenGenerate) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### SetExpiryNil

`func (o *ApiTokenGenerate) SetExpiryNil(b bool)`

 SetExpiryNil sets the value for Expiry to be an explicit nil

### UnsetExpiry
`func (o *ApiTokenGenerate) UnsetExpiry()`

UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
### GetReadWrite

`func (o *ApiTokenGenerate) GetReadWrite() bool`

GetReadWrite returns the ReadWrite field if non-nil, zero value otherwise.

### GetReadWriteOk

`func (o *ApiTokenGenerate) GetReadWriteOk() (*bool, bool)`

GetReadWriteOk returns a tuple with the ReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWrite

`func (o *ApiTokenGenerate) SetReadWrite(v bool)`

SetReadWrite sets ReadWrite field to given value.


### GetCompact

`func (o *ApiTokenGenerate) GetCompact() bool`

GetCompact returns the Compact field if non-nil, zero value otherwise.

### GetCompactOk

`func (o *ApiTokenGenerate) GetCompactOk() (*bool, bool)`

GetCompactOk returns a tuple with the Compact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompact

`func (o *ApiTokenGenerate) SetCompact(v bool)`

SetCompact sets Compact field to given value.

### HasCompact

`func (o *ApiTokenGenerate) HasCompact() bool`

HasCompact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


