# ScimSyncState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | [**ScimSyncStateOneOfActive**](ScimSyncStateOneOfActive.md) |  | 

## Methods

### NewScimSyncState

`func NewScimSyncState(active ScimSyncStateOneOfActive, ) *ScimSyncState`

NewScimSyncState instantiates a new ScimSyncState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSyncStateWithDefaults

`func NewScimSyncStateWithDefaults() *ScimSyncState`

NewScimSyncStateWithDefaults instantiates a new ScimSyncState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ScimSyncState) GetActive() ScimSyncStateOneOfActive`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimSyncState) GetActiveOk() (*ScimSyncStateOneOfActive, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ScimSyncState) SetActive(v ScimSyncStateOneOfActive)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


