# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spn** | **string** |  | 
**Uuid** | **string** |  | 

## Methods

### NewGroup

`func NewGroup(spn string, uuid string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpn

`func (o *Group) GetSpn() string`

GetSpn returns the Spn field if non-nil, zero value otherwise.

### GetSpnOk

`func (o *Group) GetSpnOk() (*string, bool)`

GetSpnOk returns a tuple with the Spn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpn

`func (o *Group) SetSpn(v string)`

SetSpn sets Spn field to given value.


### GetUuid

`func (o *Group) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Group) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Group) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


