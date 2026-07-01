# ScimListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** |  | 
**TotalResults** | **int64** |  | 
**ItemsPerPage** | Pointer to **int64** |  | [optional] 
**StartIndex** | Pointer to **int64** |  | [optional] 
**Resources** | [**[]ScimEntry**](ScimEntry.md) |  | 

## Methods

### NewScimListResponse

`func NewScimListResponse(schemas []string, totalResults int64, resources []ScimEntry, ) *ScimListResponse`

NewScimListResponse instantiates a new ScimListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimListResponseWithDefaults

`func NewScimListResponseWithDefaults() *ScimListResponse`

NewScimListResponseWithDefaults instantiates a new ScimListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetTotalResults

`func (o *ScimListResponse) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ScimListResponse) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ScimListResponse) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.


### GetItemsPerPage

`func (o *ScimListResponse) GetItemsPerPage() int64`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ScimListResponse) GetItemsPerPageOk() (*int64, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ScimListResponse) SetItemsPerPage(v int64)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ScimListResponse) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetStartIndex

`func (o *ScimListResponse) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ScimListResponse) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ScimListResponse) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *ScimListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetResources

`func (o *ScimListResponse) GetResources() []ScimEntry`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ScimListResponse) GetResourcesOk() (*[]ScimEntry, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ScimListResponse) SetResources(v []ScimEntry)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


