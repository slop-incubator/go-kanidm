# AccountUnixExtend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gidnumber** | Pointer to **NullableInt32** |  | [optional] 
**Shell** | Pointer to **NullableString** | The internal attribute is \&quot;loginshell\&quot; but we use shell in the API currently | [optional] 

## Methods

### NewAccountUnixExtend

`func NewAccountUnixExtend() *AccountUnixExtend`

NewAccountUnixExtend instantiates a new AccountUnixExtend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUnixExtendWithDefaults

`func NewAccountUnixExtendWithDefaults() *AccountUnixExtend`

NewAccountUnixExtendWithDefaults instantiates a new AccountUnixExtend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGidnumber

`func (o *AccountUnixExtend) GetGidnumber() int32`

GetGidnumber returns the Gidnumber field if non-nil, zero value otherwise.

### GetGidnumberOk

`func (o *AccountUnixExtend) GetGidnumberOk() (*int32, bool)`

GetGidnumberOk returns a tuple with the Gidnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidnumber

`func (o *AccountUnixExtend) SetGidnumber(v int32)`

SetGidnumber sets Gidnumber field to given value.

### HasGidnumber

`func (o *AccountUnixExtend) HasGidnumber() bool`

HasGidnumber returns a boolean if a field has been set.

### SetGidnumberNil

`func (o *AccountUnixExtend) SetGidnumberNil(b bool)`

 SetGidnumberNil sets the value for Gidnumber to be an explicit nil

### UnsetGidnumber
`func (o *AccountUnixExtend) UnsetGidnumber()`

UnsetGidnumber ensures that no value is present for Gidnumber, not even an explicit nil
### GetShell

`func (o *AccountUnixExtend) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *AccountUnixExtend) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *AccountUnixExtend) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *AccountUnixExtend) HasShell() bool`

HasShell returns a boolean if a field has been set.

### SetShellNil

`func (o *AccountUnixExtend) SetShellNil(b bool)`

 SetShellNil sets the value for Shell to be an explicit nil

### UnsetShell
`func (o *AccountUnixExtend) UnsetShell()`

UnsetShell ensures that no value is present for Shell, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


