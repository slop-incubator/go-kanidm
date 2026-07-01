# ScimMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | 
**Created** | **time.Time** |  | 
**LastModified** | **time.Time** |  | 
**Location** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewScimMeta

`func NewScimMeta(resourceType string, created time.Time, lastModified time.Time, location string, version string, ) *ScimMeta`

NewScimMeta instantiates a new ScimMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimMetaWithDefaults

`func NewScimMetaWithDefaults() *ScimMeta`

NewScimMetaWithDefaults instantiates a new ScimMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *ScimMeta) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ScimMeta) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ScimMeta) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetCreated

`func (o *ScimMeta) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ScimMeta) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ScimMeta) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastModified

`func (o *ScimMeta) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ScimMeta) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ScimMeta) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetLocation

`func (o *ScimMeta) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ScimMeta) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ScimMeta) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetVersion

`func (o *ScimMeta) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ScimMeta) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ScimMeta) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


