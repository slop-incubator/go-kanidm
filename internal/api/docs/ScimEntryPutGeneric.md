# ScimEntryPutGeneric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]Attribute**](Attribute.md) |  | [optional] 
**ExtAccessCheck** | Pointer to **bool** |  | [optional] 
**SortBy** | Pointer to [**NullableAttribute**](Attribute.md) |  | [optional] 
**SortOrder** | Pointer to [**NullableScimSortOrder**](ScimSortOrder.md) |  | [optional] 
**StartIndex** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Filter** | Pointer to **interface{}** |  | [optional] 
**Id** | **string** |  | 

## Methods

### NewScimEntryPutGeneric

`func NewScimEntryPutGeneric(id string, ) *ScimEntryPutGeneric`

NewScimEntryPutGeneric instantiates a new ScimEntryPutGeneric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimEntryPutGenericWithDefaults

`func NewScimEntryPutGenericWithDefaults() *ScimEntryPutGeneric`

NewScimEntryPutGenericWithDefaults instantiates a new ScimEntryPutGeneric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ScimEntryPutGeneric) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ScimEntryPutGeneric) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ScimEntryPutGeneric) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ScimEntryPutGeneric) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ScimEntryPutGeneric) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ScimEntryPutGeneric) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetExtAccessCheck

`func (o *ScimEntryPutGeneric) GetExtAccessCheck() bool`

GetExtAccessCheck returns the ExtAccessCheck field if non-nil, zero value otherwise.

### GetExtAccessCheckOk

`func (o *ScimEntryPutGeneric) GetExtAccessCheckOk() (*bool, bool)`

GetExtAccessCheckOk returns a tuple with the ExtAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAccessCheck

`func (o *ScimEntryPutGeneric) SetExtAccessCheck(v bool)`

SetExtAccessCheck sets ExtAccessCheck field to given value.

### HasExtAccessCheck

`func (o *ScimEntryPutGeneric) HasExtAccessCheck() bool`

HasExtAccessCheck returns a boolean if a field has been set.

### GetSortBy

`func (o *ScimEntryPutGeneric) GetSortBy() Attribute`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ScimEntryPutGeneric) GetSortByOk() (*Attribute, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ScimEntryPutGeneric) SetSortBy(v Attribute)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ScimEntryPutGeneric) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *ScimEntryPutGeneric) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *ScimEntryPutGeneric) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *ScimEntryPutGeneric) GetSortOrder() ScimSortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ScimEntryPutGeneric) GetSortOrderOk() (*ScimSortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ScimEntryPutGeneric) SetSortOrder(v ScimSortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ScimEntryPutGeneric) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *ScimEntryPutGeneric) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *ScimEntryPutGeneric) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetStartIndex

`func (o *ScimEntryPutGeneric) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ScimEntryPutGeneric) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ScimEntryPutGeneric) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *ScimEntryPutGeneric) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetCount

`func (o *ScimEntryPutGeneric) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ScimEntryPutGeneric) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ScimEntryPutGeneric) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ScimEntryPutGeneric) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFilter

`func (o *ScimEntryPutGeneric) GetFilter() interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ScimEntryPutGeneric) GetFilterOk() (*interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ScimEntryPutGeneric) SetFilter(v interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ScimEntryPutGeneric) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ScimEntryPutGeneric) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ScimEntryPutGeneric) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetId

`func (o *ScimEntryPutGeneric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimEntryPutGeneric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimEntryPutGeneric) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


