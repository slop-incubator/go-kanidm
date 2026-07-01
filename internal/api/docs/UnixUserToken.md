# UnixUserToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Spn** | **string** |  | 
**Displayname** | **string** |  | 
**Gidnumber** | **int32** |  | 
**Uuid** | **string** |  | 
**Shell** | Pointer to **NullableString** |  | [optional] 
**Groups** | [**[]UnixGroupToken**](UnixGroupToken.md) |  | 
**Sshkeys** | **[]string** |  | 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewUnixUserToken

`func NewUnixUserToken(name string, spn string, displayname string, gidnumber int32, uuid string, groups []UnixGroupToken, sshkeys []string, ) *UnixUserToken`

NewUnixUserToken instantiates a new UnixUserToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnixUserTokenWithDefaults

`func NewUnixUserTokenWithDefaults() *UnixUserToken`

NewUnixUserTokenWithDefaults instantiates a new UnixUserToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UnixUserToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnixUserToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnixUserToken) SetName(v string)`

SetName sets Name field to given value.


### GetSpn

`func (o *UnixUserToken) GetSpn() string`

GetSpn returns the Spn field if non-nil, zero value otherwise.

### GetSpnOk

`func (o *UnixUserToken) GetSpnOk() (*string, bool)`

GetSpnOk returns a tuple with the Spn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpn

`func (o *UnixUserToken) SetSpn(v string)`

SetSpn sets Spn field to given value.


### GetDisplayname

`func (o *UnixUserToken) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *UnixUserToken) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *UnixUserToken) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.


### GetGidnumber

`func (o *UnixUserToken) GetGidnumber() int32`

GetGidnumber returns the Gidnumber field if non-nil, zero value otherwise.

### GetGidnumberOk

`func (o *UnixUserToken) GetGidnumberOk() (*int32, bool)`

GetGidnumberOk returns a tuple with the Gidnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidnumber

`func (o *UnixUserToken) SetGidnumber(v int32)`

SetGidnumber sets Gidnumber field to given value.


### GetUuid

`func (o *UnixUserToken) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UnixUserToken) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UnixUserToken) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetShell

`func (o *UnixUserToken) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *UnixUserToken) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *UnixUserToken) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *UnixUserToken) HasShell() bool`

HasShell returns a boolean if a field has been set.

### SetShellNil

`func (o *UnixUserToken) SetShellNil(b bool)`

 SetShellNil sets the value for Shell to be an explicit nil

### UnsetShell
`func (o *UnixUserToken) UnsetShell()`

UnsetShell ensures that no value is present for Shell, not even an explicit nil
### GetGroups

`func (o *UnixUserToken) GetGroups() []UnixGroupToken`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UnixUserToken) GetGroupsOk() (*[]UnixGroupToken, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UnixUserToken) SetGroups(v []UnixGroupToken)`

SetGroups sets Groups field to given value.


### GetSshkeys

`func (o *UnixUserToken) GetSshkeys() []string`

GetSshkeys returns the Sshkeys field if non-nil, zero value otherwise.

### GetSshkeysOk

`func (o *UnixUserToken) GetSshkeysOk() (*[]string, bool)`

GetSshkeysOk returns a tuple with the Sshkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkeys

`func (o *UnixUserToken) SetSshkeys(v []string)`

SetSshkeys sets Sshkeys field to given value.


### GetValid

`func (o *UnixUserToken) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *UnixUserToken) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *UnixUserToken) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *UnixUserToken) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


