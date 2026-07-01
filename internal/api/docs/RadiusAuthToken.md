# RadiusAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Displayname** | **string** |  | 
**Uuid** | **string** |  | 
**Secret** | **string** |  | 
**Groups** | [**[]Group**](Group.md) |  | 

## Methods

### NewRadiusAuthToken

`func NewRadiusAuthToken(name string, displayname string, uuid string, secret string, groups []Group, ) *RadiusAuthToken`

NewRadiusAuthToken instantiates a new RadiusAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusAuthTokenWithDefaults

`func NewRadiusAuthTokenWithDefaults() *RadiusAuthToken`

NewRadiusAuthTokenWithDefaults instantiates a new RadiusAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RadiusAuthToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadiusAuthToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadiusAuthToken) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayname

`func (o *RadiusAuthToken) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *RadiusAuthToken) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *RadiusAuthToken) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.


### GetUuid

`func (o *RadiusAuthToken) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RadiusAuthToken) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RadiusAuthToken) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetSecret

`func (o *RadiusAuthToken) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RadiusAuthToken) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RadiusAuthToken) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetGroups

`func (o *RadiusAuthToken) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RadiusAuthToken) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RadiusAuthToken) SetGroups(v []Group)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


