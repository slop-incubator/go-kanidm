# ScimEntryGetQuery

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

## Methods

### NewScimEntryGetQuery

`func NewScimEntryGetQuery() *ScimEntryGetQuery`

NewScimEntryGetQuery instantiates a new ScimEntryGetQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimEntryGetQueryWithDefaults

`func NewScimEntryGetQueryWithDefaults() *ScimEntryGetQuery`

NewScimEntryGetQueryWithDefaults instantiates a new ScimEntryGetQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ScimEntryGetQuery) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ScimEntryGetQuery) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ScimEntryGetQuery) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ScimEntryGetQuery) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ScimEntryGetQuery) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ScimEntryGetQuery) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetExtAccessCheck

`func (o *ScimEntryGetQuery) GetExtAccessCheck() bool`

GetExtAccessCheck returns the ExtAccessCheck field if non-nil, zero value otherwise.

### GetExtAccessCheckOk

`func (o *ScimEntryGetQuery) GetExtAccessCheckOk() (*bool, bool)`

GetExtAccessCheckOk returns a tuple with the ExtAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAccessCheck

`func (o *ScimEntryGetQuery) SetExtAccessCheck(v bool)`

SetExtAccessCheck sets ExtAccessCheck field to given value.

### HasExtAccessCheck

`func (o *ScimEntryGetQuery) HasExtAccessCheck() bool`

HasExtAccessCheck returns a boolean if a field has been set.

### GetSortBy

`func (o *ScimEntryGetQuery) GetSortBy() Attribute`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ScimEntryGetQuery) GetSortByOk() (*Attribute, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ScimEntryGetQuery) SetSortBy(v Attribute)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ScimEntryGetQuery) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *ScimEntryGetQuery) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *ScimEntryGetQuery) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *ScimEntryGetQuery) GetSortOrder() ScimSortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ScimEntryGetQuery) GetSortOrderOk() (*ScimSortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ScimEntryGetQuery) SetSortOrder(v ScimSortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ScimEntryGetQuery) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *ScimEntryGetQuery) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *ScimEntryGetQuery) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetStartIndex

`func (o *ScimEntryGetQuery) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ScimEntryGetQuery) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ScimEntryGetQuery) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *ScimEntryGetQuery) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetCount

`func (o *ScimEntryGetQuery) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ScimEntryGetQuery) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ScimEntryGetQuery) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ScimEntryGetQuery) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFilter

`func (o *ScimEntryGetQuery) GetFilter() interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ScimEntryGetQuery) GetFilterOk() (*interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ScimEntryGetQuery) SetFilter(v interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ScimEntryGetQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ScimEntryGetQuery) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ScimEntryGetQuery) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


