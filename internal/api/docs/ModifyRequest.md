# ModifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**Filter**](Filter.md) |  | 
**Modlist** | [**ModifyList**](ModifyList.md) |  | 

## Methods

### NewModifyRequest

`func NewModifyRequest(filter Filter, modlist ModifyList, ) *ModifyRequest`

NewModifyRequest instantiates a new ModifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyRequestWithDefaults

`func NewModifyRequestWithDefaults() *ModifyRequest`

NewModifyRequestWithDefaults instantiates a new ModifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ModifyRequest) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ModifyRequest) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ModifyRequest) SetFilter(v Filter)`

SetFilter sets Filter field to given value.


### GetModlist

`func (o *ModifyRequest) GetModlist() ModifyList`

GetModlist returns the Modlist field if non-nil, zero value otherwise.

### GetModlistOk

`func (o *ModifyRequest) GetModlistOk() (*ModifyList, bool)`

GetModlistOk returns a tuple with the Modlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModlist

`func (o *ModifyRequest) SetModlist(v ModifyList)`

SetModlist sets Modlist field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


