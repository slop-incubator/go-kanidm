# ScimSyncRetentionMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retain** | **[]string** | All entries that have their uuid present in this set are retained. Anything not present will be deleted. | 
**Delete** | **[]string** | Any entry with its UUID in this set will be deleted. Anything not present will be retained. | 

## Methods

### NewScimSyncRetentionMode

`func NewScimSyncRetentionMode(retain []string, delete []string, ) *ScimSyncRetentionMode`

NewScimSyncRetentionMode instantiates a new ScimSyncRetentionMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSyncRetentionModeWithDefaults

`func NewScimSyncRetentionModeWithDefaults() *ScimSyncRetentionMode`

NewScimSyncRetentionModeWithDefaults instantiates a new ScimSyncRetentionMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetain

`func (o *ScimSyncRetentionMode) GetRetain() []string`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *ScimSyncRetentionMode) GetRetainOk() (*[]string, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *ScimSyncRetentionMode) SetRetain(v []string)`

SetRetain sets Retain field to given value.


### GetDelete

`func (o *ScimSyncRetentionMode) GetDelete() []string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ScimSyncRetentionMode) GetDeleteOk() (*[]string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ScimSyncRetentionMode) SetDelete(v []string)`

SetDelete sets Delete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


