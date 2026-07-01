# SchemaError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invalidclass** | **[]string** |  | 
**Missingmustattribute** | [**[]Attribute**](Attribute.md) |  | 
**Invalidattribute** | **string** |  | 
**Invalidattributesyntax** | **string** |  | 
**Attributenotvalidforclass** | **string** |  | 
**Supplementsnotsatisfied** | **[]string** |  | 
**Excludesnotsatisfied** | **[]string** |  | 
**Phantomattribute** | **string** |  | 

## Methods

### NewSchemaError

`func NewSchemaError(invalidclass []string, missingmustattribute []Attribute, invalidattribute string, invalidattributesyntax string, attributenotvalidforclass string, supplementsnotsatisfied []string, excludesnotsatisfied []string, phantomattribute string, ) *SchemaError`

NewSchemaError instantiates a new SchemaError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaErrorWithDefaults

`func NewSchemaErrorWithDefaults() *SchemaError`

NewSchemaErrorWithDefaults instantiates a new SchemaError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidclass

`func (o *SchemaError) GetInvalidclass() []string`

GetInvalidclass returns the Invalidclass field if non-nil, zero value otherwise.

### GetInvalidclassOk

`func (o *SchemaError) GetInvalidclassOk() (*[]string, bool)`

GetInvalidclassOk returns a tuple with the Invalidclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidclass

`func (o *SchemaError) SetInvalidclass(v []string)`

SetInvalidclass sets Invalidclass field to given value.


### GetMissingmustattribute

`func (o *SchemaError) GetMissingmustattribute() []Attribute`

GetMissingmustattribute returns the Missingmustattribute field if non-nil, zero value otherwise.

### GetMissingmustattributeOk

`func (o *SchemaError) GetMissingmustattributeOk() (*[]Attribute, bool)`

GetMissingmustattributeOk returns a tuple with the Missingmustattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingmustattribute

`func (o *SchemaError) SetMissingmustattribute(v []Attribute)`

SetMissingmustattribute sets Missingmustattribute field to given value.


### GetInvalidattribute

`func (o *SchemaError) GetInvalidattribute() string`

GetInvalidattribute returns the Invalidattribute field if non-nil, zero value otherwise.

### GetInvalidattributeOk

`func (o *SchemaError) GetInvalidattributeOk() (*string, bool)`

GetInvalidattributeOk returns a tuple with the Invalidattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidattribute

`func (o *SchemaError) SetInvalidattribute(v string)`

SetInvalidattribute sets Invalidattribute field to given value.


### GetInvalidattributesyntax

`func (o *SchemaError) GetInvalidattributesyntax() string`

GetInvalidattributesyntax returns the Invalidattributesyntax field if non-nil, zero value otherwise.

### GetInvalidattributesyntaxOk

`func (o *SchemaError) GetInvalidattributesyntaxOk() (*string, bool)`

GetInvalidattributesyntaxOk returns a tuple with the Invalidattributesyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidattributesyntax

`func (o *SchemaError) SetInvalidattributesyntax(v string)`

SetInvalidattributesyntax sets Invalidattributesyntax field to given value.


### GetAttributenotvalidforclass

`func (o *SchemaError) GetAttributenotvalidforclass() string`

GetAttributenotvalidforclass returns the Attributenotvalidforclass field if non-nil, zero value otherwise.

### GetAttributenotvalidforclassOk

`func (o *SchemaError) GetAttributenotvalidforclassOk() (*string, bool)`

GetAttributenotvalidforclassOk returns a tuple with the Attributenotvalidforclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributenotvalidforclass

`func (o *SchemaError) SetAttributenotvalidforclass(v string)`

SetAttributenotvalidforclass sets Attributenotvalidforclass field to given value.


### GetSupplementsnotsatisfied

`func (o *SchemaError) GetSupplementsnotsatisfied() []string`

GetSupplementsnotsatisfied returns the Supplementsnotsatisfied field if non-nil, zero value otherwise.

### GetSupplementsnotsatisfiedOk

`func (o *SchemaError) GetSupplementsnotsatisfiedOk() (*[]string, bool)`

GetSupplementsnotsatisfiedOk returns a tuple with the Supplementsnotsatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementsnotsatisfied

`func (o *SchemaError) SetSupplementsnotsatisfied(v []string)`

SetSupplementsnotsatisfied sets Supplementsnotsatisfied field to given value.


### GetExcludesnotsatisfied

`func (o *SchemaError) GetExcludesnotsatisfied() []string`

GetExcludesnotsatisfied returns the Excludesnotsatisfied field if non-nil, zero value otherwise.

### GetExcludesnotsatisfiedOk

`func (o *SchemaError) GetExcludesnotsatisfiedOk() (*[]string, bool)`

GetExcludesnotsatisfiedOk returns a tuple with the Excludesnotsatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludesnotsatisfied

`func (o *SchemaError) SetExcludesnotsatisfied(v []string)`

SetExcludesnotsatisfied sets Excludesnotsatisfied field to given value.


### GetPhantomattribute

`func (o *SchemaError) GetPhantomattribute() string`

GetPhantomattribute returns the Phantomattribute field if non-nil, zero value otherwise.

### GetPhantomattributeOk

`func (o *SchemaError) GetPhantomattributeOk() (*string, bool)`

GetPhantomattributeOk returns a tuple with the Phantomattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhantomattribute

`func (o *SchemaError) SetPhantomattribute(v string)`

SetPhantomattribute sets Phantomattribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


