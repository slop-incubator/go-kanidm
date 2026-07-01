# ScimSyncRetentionModeOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retain** | **[]string** | All entries that have their uuid present in this set are retained. Anything not present will be deleted. | 

## Methods

### NewScimSyncRetentionModeOneOf

`func NewScimSyncRetentionModeOneOf(retain []string, ) *ScimSyncRetentionModeOneOf`

NewScimSyncRetentionModeOneOf instantiates a new ScimSyncRetentionModeOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSyncRetentionModeOneOfWithDefaults

`func NewScimSyncRetentionModeOneOfWithDefaults() *ScimSyncRetentionModeOneOf`

NewScimSyncRetentionModeOneOfWithDefaults instantiates a new ScimSyncRetentionModeOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetain

`func (o *ScimSyncRetentionModeOneOf) GetRetain() []string`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *ScimSyncRetentionModeOneOf) GetRetainOk() (*[]string, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *ScimSyncRetentionModeOneOf) SetRetain(v []string)`

SetRetain sets Retain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


