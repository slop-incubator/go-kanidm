# ScimSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromState** | [**ScimSyncState**](ScimSyncState.md) |  | 
**ToState** | [**ScimSyncState**](ScimSyncState.md) |  | 
**Entries** | [**[]ScimEntry**](ScimEntry.md) |  | 
**Retain** | [**ScimSyncRetentionMode**](ScimSyncRetentionMode.md) |  | 

## Methods

### NewScimSyncRequest

`func NewScimSyncRequest(fromState ScimSyncState, toState ScimSyncState, entries []ScimEntry, retain ScimSyncRetentionMode, ) *ScimSyncRequest`

NewScimSyncRequest instantiates a new ScimSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSyncRequestWithDefaults

`func NewScimSyncRequestWithDefaults() *ScimSyncRequest`

NewScimSyncRequestWithDefaults instantiates a new ScimSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromState

`func (o *ScimSyncRequest) GetFromState() ScimSyncState`

GetFromState returns the FromState field if non-nil, zero value otherwise.

### GetFromStateOk

`func (o *ScimSyncRequest) GetFromStateOk() (*ScimSyncState, bool)`

GetFromStateOk returns a tuple with the FromState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromState

`func (o *ScimSyncRequest) SetFromState(v ScimSyncState)`

SetFromState sets FromState field to given value.


### GetToState

`func (o *ScimSyncRequest) GetToState() ScimSyncState`

GetToState returns the ToState field if non-nil, zero value otherwise.

### GetToStateOk

`func (o *ScimSyncRequest) GetToStateOk() (*ScimSyncState, bool)`

GetToStateOk returns a tuple with the ToState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToState

`func (o *ScimSyncRequest) SetToState(v ScimSyncState)`

SetToState sets ToState field to given value.


### GetEntries

`func (o *ScimSyncRequest) GetEntries() []ScimEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ScimSyncRequest) GetEntriesOk() (*[]ScimEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ScimSyncRequest) SetEntries(v []ScimEntry)`

SetEntries sets Entries field to given value.


### GetRetain

`func (o *ScimSyncRequest) GetRetain() ScimSyncRetentionMode`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *ScimSyncRequest) GetRetainOk() (*ScimSyncRetentionMode, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *ScimSyncRequest) SetRetain(v ScimSyncRetentionMode)`

SetRetain sets Retain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


