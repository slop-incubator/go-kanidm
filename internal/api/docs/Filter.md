# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eq** | **[]string** |  | 
**Cnt** | **[]string** |  | 
**Pres** | **string** |  | 
**Or** | [**[]Filter**](Filter.md) |  | 
**And** | [**[]Filter**](Filter.md) |  | 
**Andnot** | [**Filter**](Filter.md) |  | 

## Methods

### NewFilter

`func NewFilter(eq []string, cnt []string, pres string, or []Filter, and []Filter, andnot Filter, ) *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEq

`func (o *Filter) GetEq() []string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *Filter) GetEqOk() (*[]string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *Filter) SetEq(v []string)`

SetEq sets Eq field to given value.


### GetCnt

`func (o *Filter) GetCnt() []string`

GetCnt returns the Cnt field if non-nil, zero value otherwise.

### GetCntOk

`func (o *Filter) GetCntOk() (*[]string, bool)`

GetCntOk returns a tuple with the Cnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnt

`func (o *Filter) SetCnt(v []string)`

SetCnt sets Cnt field to given value.


### GetPres

`func (o *Filter) GetPres() string`

GetPres returns the Pres field if non-nil, zero value otherwise.

### GetPresOk

`func (o *Filter) GetPresOk() (*string, bool)`

GetPresOk returns a tuple with the Pres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPres

`func (o *Filter) SetPres(v string)`

SetPres sets Pres field to given value.


### GetOr

`func (o *Filter) GetOr() []Filter`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *Filter) GetOrOk() (*[]Filter, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *Filter) SetOr(v []Filter)`

SetOr sets Or field to given value.


### GetAnd

`func (o *Filter) GetAnd() []Filter`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *Filter) GetAndOk() (*[]Filter, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *Filter) SetAnd(v []Filter)`

SetAnd sets And field to given value.


### GetAndnot

`func (o *Filter) GetAndnot() Filter`

GetAndnot returns the Andnot field if non-nil, zero value otherwise.

### GetAndnotOk

`func (o *Filter) GetAndnotOk() (*Filter, bool)`

GetAndnotOk returns a tuple with the Andnot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndnot

`func (o *Filter) SetAndnot(v Filter)`

SetAndnot sets Andnot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


