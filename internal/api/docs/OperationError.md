# OperationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Corruptedentry** | **int64** |  | 
**Corruptedindex** | **string** |  | 
**Consistencyerror** | [**[]ConsistencyError**](ConsistencyError.md) |  | 
**Schemaviolation** | [**SchemaError**](SchemaError.md) |  | 
**Plugin** | [**PluginError**](PluginError.md) |  | 
**Invalidattributename** | **string** |  | 
**Invalidattribute** | **string** |  | 
**Invalidacpstate** | **string** |  | 
**Invalidschemastate** | **string** |  | 
**Invalidaccountstate** | **string** |  | 
**Missingclass** | **string** |  | 
**Missingattribute** | [**Attribute**](Attribute.md) |  | 
**Attributeuniqueness** | [**[]Attribute**](Attribute.md) |  | 
**Invalidauthstate** | **string** |  | 
**Passwordquality** | [**[]PasswordFeedback**](PasswordFeedback.md) |  | 
**Wait** | **time.Time** |  | 

## Methods

### NewOperationError

`func NewOperationError(corruptedentry int64, corruptedindex string, consistencyerror []ConsistencyError, schemaviolation SchemaError, plugin PluginError, invalidattributename string, invalidattribute string, invalidacpstate string, invalidschemastate string, invalidaccountstate string, missingclass string, missingattribute Attribute, attributeuniqueness []Attribute, invalidauthstate string, passwordquality []PasswordFeedback, wait time.Time, ) *OperationError`

NewOperationError instantiates a new OperationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationErrorWithDefaults

`func NewOperationErrorWithDefaults() *OperationError`

NewOperationErrorWithDefaults instantiates a new OperationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorruptedentry

`func (o *OperationError) GetCorruptedentry() int64`

GetCorruptedentry returns the Corruptedentry field if non-nil, zero value otherwise.

### GetCorruptedentryOk

`func (o *OperationError) GetCorruptedentryOk() (*int64, bool)`

GetCorruptedentryOk returns a tuple with the Corruptedentry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorruptedentry

`func (o *OperationError) SetCorruptedentry(v int64)`

SetCorruptedentry sets Corruptedentry field to given value.


### GetCorruptedindex

`func (o *OperationError) GetCorruptedindex() string`

GetCorruptedindex returns the Corruptedindex field if non-nil, zero value otherwise.

### GetCorruptedindexOk

`func (o *OperationError) GetCorruptedindexOk() (*string, bool)`

GetCorruptedindexOk returns a tuple with the Corruptedindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorruptedindex

`func (o *OperationError) SetCorruptedindex(v string)`

SetCorruptedindex sets Corruptedindex field to given value.


### GetConsistencyerror

`func (o *OperationError) GetConsistencyerror() []ConsistencyError`

GetConsistencyerror returns the Consistencyerror field if non-nil, zero value otherwise.

### GetConsistencyerrorOk

`func (o *OperationError) GetConsistencyerrorOk() (*[]ConsistencyError, bool)`

GetConsistencyerrorOk returns a tuple with the Consistencyerror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyerror

`func (o *OperationError) SetConsistencyerror(v []ConsistencyError)`

SetConsistencyerror sets Consistencyerror field to given value.


### GetSchemaviolation

`func (o *OperationError) GetSchemaviolation() SchemaError`

GetSchemaviolation returns the Schemaviolation field if non-nil, zero value otherwise.

### GetSchemaviolationOk

`func (o *OperationError) GetSchemaviolationOk() (*SchemaError, bool)`

GetSchemaviolationOk returns a tuple with the Schemaviolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaviolation

`func (o *OperationError) SetSchemaviolation(v SchemaError)`

SetSchemaviolation sets Schemaviolation field to given value.


### GetPlugin

`func (o *OperationError) GetPlugin() PluginError`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *OperationError) GetPluginOk() (*PluginError, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *OperationError) SetPlugin(v PluginError)`

SetPlugin sets Plugin field to given value.


### GetInvalidattributename

`func (o *OperationError) GetInvalidattributename() string`

GetInvalidattributename returns the Invalidattributename field if non-nil, zero value otherwise.

### GetInvalidattributenameOk

`func (o *OperationError) GetInvalidattributenameOk() (*string, bool)`

GetInvalidattributenameOk returns a tuple with the Invalidattributename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidattributename

`func (o *OperationError) SetInvalidattributename(v string)`

SetInvalidattributename sets Invalidattributename field to given value.


### GetInvalidattribute

`func (o *OperationError) GetInvalidattribute() string`

GetInvalidattribute returns the Invalidattribute field if non-nil, zero value otherwise.

### GetInvalidattributeOk

`func (o *OperationError) GetInvalidattributeOk() (*string, bool)`

GetInvalidattributeOk returns a tuple with the Invalidattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidattribute

`func (o *OperationError) SetInvalidattribute(v string)`

SetInvalidattribute sets Invalidattribute field to given value.


### GetInvalidacpstate

`func (o *OperationError) GetInvalidacpstate() string`

GetInvalidacpstate returns the Invalidacpstate field if non-nil, zero value otherwise.

### GetInvalidacpstateOk

`func (o *OperationError) GetInvalidacpstateOk() (*string, bool)`

GetInvalidacpstateOk returns a tuple with the Invalidacpstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidacpstate

`func (o *OperationError) SetInvalidacpstate(v string)`

SetInvalidacpstate sets Invalidacpstate field to given value.


### GetInvalidschemastate

`func (o *OperationError) GetInvalidschemastate() string`

GetInvalidschemastate returns the Invalidschemastate field if non-nil, zero value otherwise.

### GetInvalidschemastateOk

`func (o *OperationError) GetInvalidschemastateOk() (*string, bool)`

GetInvalidschemastateOk returns a tuple with the Invalidschemastate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidschemastate

`func (o *OperationError) SetInvalidschemastate(v string)`

SetInvalidschemastate sets Invalidschemastate field to given value.


### GetInvalidaccountstate

`func (o *OperationError) GetInvalidaccountstate() string`

GetInvalidaccountstate returns the Invalidaccountstate field if non-nil, zero value otherwise.

### GetInvalidaccountstateOk

`func (o *OperationError) GetInvalidaccountstateOk() (*string, bool)`

GetInvalidaccountstateOk returns a tuple with the Invalidaccountstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidaccountstate

`func (o *OperationError) SetInvalidaccountstate(v string)`

SetInvalidaccountstate sets Invalidaccountstate field to given value.


### GetMissingclass

`func (o *OperationError) GetMissingclass() string`

GetMissingclass returns the Missingclass field if non-nil, zero value otherwise.

### GetMissingclassOk

`func (o *OperationError) GetMissingclassOk() (*string, bool)`

GetMissingclassOk returns a tuple with the Missingclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingclass

`func (o *OperationError) SetMissingclass(v string)`

SetMissingclass sets Missingclass field to given value.


### GetMissingattribute

`func (o *OperationError) GetMissingattribute() Attribute`

GetMissingattribute returns the Missingattribute field if non-nil, zero value otherwise.

### GetMissingattributeOk

`func (o *OperationError) GetMissingattributeOk() (*Attribute, bool)`

GetMissingattributeOk returns a tuple with the Missingattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingattribute

`func (o *OperationError) SetMissingattribute(v Attribute)`

SetMissingattribute sets Missingattribute field to given value.


### GetAttributeuniqueness

`func (o *OperationError) GetAttributeuniqueness() []Attribute`

GetAttributeuniqueness returns the Attributeuniqueness field if non-nil, zero value otherwise.

### GetAttributeuniquenessOk

`func (o *OperationError) GetAttributeuniquenessOk() (*[]Attribute, bool)`

GetAttributeuniquenessOk returns a tuple with the Attributeuniqueness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeuniqueness

`func (o *OperationError) SetAttributeuniqueness(v []Attribute)`

SetAttributeuniqueness sets Attributeuniqueness field to given value.


### GetInvalidauthstate

`func (o *OperationError) GetInvalidauthstate() string`

GetInvalidauthstate returns the Invalidauthstate field if non-nil, zero value otherwise.

### GetInvalidauthstateOk

`func (o *OperationError) GetInvalidauthstateOk() (*string, bool)`

GetInvalidauthstateOk returns a tuple with the Invalidauthstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidauthstate

`func (o *OperationError) SetInvalidauthstate(v string)`

SetInvalidauthstate sets Invalidauthstate field to given value.


### GetPasswordquality

`func (o *OperationError) GetPasswordquality() []PasswordFeedback`

GetPasswordquality returns the Passwordquality field if non-nil, zero value otherwise.

### GetPasswordqualityOk

`func (o *OperationError) GetPasswordqualityOk() (*[]PasswordFeedback, bool)`

GetPasswordqualityOk returns a tuple with the Passwordquality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordquality

`func (o *OperationError) SetPasswordquality(v []PasswordFeedback)`

SetPasswordquality sets Passwordquality field to given value.


### GetWait

`func (o *OperationError) GetWait() time.Time`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *OperationError) GetWaitOk() (*time.Time, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *OperationError) SetWait(v time.Time)`

SetWait sets Wait field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


