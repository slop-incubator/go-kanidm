# Modify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Present** | **[]string** |  | 
**Removed** | **[]string** |  | 
**Purged** | **string** |  | 

## Methods

### NewModify

`func NewModify(present []string, removed []string, purged string, ) *Modify`

NewModify instantiates a new Modify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyWithDefaults

`func NewModifyWithDefaults() *Modify`

NewModifyWithDefaults instantiates a new Modify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresent

`func (o *Modify) GetPresent() []string`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *Modify) GetPresentOk() (*[]string, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *Modify) SetPresent(v []string)`

SetPresent sets Present field to given value.


### GetRemoved

`func (o *Modify) GetRemoved() []string`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *Modify) GetRemovedOk() (*[]string, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *Modify) SetRemoved(v []string)`

SetRemoved sets Removed field to given value.


### GetPurged

`func (o *Modify) GetPurged() string`

GetPurged returns the Purged field if non-nil, zero value otherwise.

### GetPurgedOk

`func (o *Modify) GetPurgedOk() (*string, bool)`

GetPurgedOk returns a tuple with the Purged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurged

`func (o *Modify) SetPurged(v string)`

SetPurged sets Purged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


