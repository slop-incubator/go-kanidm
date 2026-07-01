# UnixGroupToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Spn** | **string** |  | 
**Uuid** | **string** |  | 
**Gidnumber** | **int32** |  | 

## Methods

### NewUnixGroupToken

`func NewUnixGroupToken(name string, spn string, uuid string, gidnumber int32, ) *UnixGroupToken`

NewUnixGroupToken instantiates a new UnixGroupToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnixGroupTokenWithDefaults

`func NewUnixGroupTokenWithDefaults() *UnixGroupToken`

NewUnixGroupTokenWithDefaults instantiates a new UnixGroupToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UnixGroupToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnixGroupToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnixGroupToken) SetName(v string)`

SetName sets Name field to given value.


### GetSpn

`func (o *UnixGroupToken) GetSpn() string`

GetSpn returns the Spn field if non-nil, zero value otherwise.

### GetSpnOk

`func (o *UnixGroupToken) GetSpnOk() (*string, bool)`

GetSpnOk returns a tuple with the Spn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpn

`func (o *UnixGroupToken) SetSpn(v string)`

SetSpn sets Spn field to given value.


### GetUuid

`func (o *UnixGroupToken) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UnixGroupToken) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UnixGroupToken) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetGidnumber

`func (o *UnixGroupToken) GetGidnumber() int32`

GetGidnumber returns the Gidnumber field if non-nil, zero value otherwise.

### GetGidnumberOk

`func (o *UnixGroupToken) GetGidnumberOk() (*int32, bool)`

GetGidnumberOk returns a tuple with the Gidnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidnumber

`func (o *UnixGroupToken) SetGidnumber(v int32)`

SetGidnumber sets Gidnumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


