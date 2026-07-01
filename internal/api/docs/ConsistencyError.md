# ConsistencyError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemaclassmissingattribute** | **[]string** |  | 
**Schemaclassphantomattribute** | **[]string** |  | 
**Schemauuidnotunique** | **string** |  | 
**Entryuuidcorrupt** | **int64** |  | 
**Uuidindexcorrupt** | **string** |  | 
**Uuidnotunique** | **string** |  | 
**Refintnotupheld** | **int64** |  | 
**Memberofinvalid** | **int64** |  | 
**Invalidattributetype** | **string** |  | 
**Invalidspn** | **int64** |  | 
**Changelogdesynchronised** | **int64** |  | 
**Changestatedesynchronised** | **int64** |  | 
**Ruvinconsistent** | **string** |  | 
**Deniedname** | **string** |  | 
**Keyprovideruuidmissing** | [**ConsistencyErrorOneOf14Keyprovideruuidmissing**](ConsistencyErrorOneOf14Keyprovideruuidmissing.md) |  | 
**Keyprovidernokeys** | [**ConsistencyErrorOneOf14Keyprovideruuidmissing**](ConsistencyErrorOneOf14Keyprovideruuidmissing.md) |  | 
**Keyprovidernotfound** | [**ConsistencyErrorOneOf16Keyprovidernotfound**](ConsistencyErrorOneOf16Keyprovidernotfound.md) |  | 

## Methods

### NewConsistencyError

`func NewConsistencyError(schemaclassmissingattribute []string, schemaclassphantomattribute []string, schemauuidnotunique string, entryuuidcorrupt int64, uuidindexcorrupt string, uuidnotunique string, refintnotupheld int64, memberofinvalid int64, invalidattributetype string, invalidspn int64, changelogdesynchronised int64, changestatedesynchronised int64, ruvinconsistent string, deniedname string, keyprovideruuidmissing ConsistencyErrorOneOf14Keyprovideruuidmissing, keyprovidernokeys ConsistencyErrorOneOf14Keyprovideruuidmissing, keyprovidernotfound ConsistencyErrorOneOf16Keyprovidernotfound, ) *ConsistencyError`

NewConsistencyError instantiates a new ConsistencyError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsistencyErrorWithDefaults

`func NewConsistencyErrorWithDefaults() *ConsistencyError`

NewConsistencyErrorWithDefaults instantiates a new ConsistencyError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaclassmissingattribute

`func (o *ConsistencyError) GetSchemaclassmissingattribute() []string`

GetSchemaclassmissingattribute returns the Schemaclassmissingattribute field if non-nil, zero value otherwise.

### GetSchemaclassmissingattributeOk

`func (o *ConsistencyError) GetSchemaclassmissingattributeOk() (*[]string, bool)`

GetSchemaclassmissingattributeOk returns a tuple with the Schemaclassmissingattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaclassmissingattribute

`func (o *ConsistencyError) SetSchemaclassmissingattribute(v []string)`

SetSchemaclassmissingattribute sets Schemaclassmissingattribute field to given value.


### GetSchemaclassphantomattribute

`func (o *ConsistencyError) GetSchemaclassphantomattribute() []string`

GetSchemaclassphantomattribute returns the Schemaclassphantomattribute field if non-nil, zero value otherwise.

### GetSchemaclassphantomattributeOk

`func (o *ConsistencyError) GetSchemaclassphantomattributeOk() (*[]string, bool)`

GetSchemaclassphantomattributeOk returns a tuple with the Schemaclassphantomattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaclassphantomattribute

`func (o *ConsistencyError) SetSchemaclassphantomattribute(v []string)`

SetSchemaclassphantomattribute sets Schemaclassphantomattribute field to given value.


### GetSchemauuidnotunique

`func (o *ConsistencyError) GetSchemauuidnotunique() string`

GetSchemauuidnotunique returns the Schemauuidnotunique field if non-nil, zero value otherwise.

### GetSchemauuidnotuniqueOk

`func (o *ConsistencyError) GetSchemauuidnotuniqueOk() (*string, bool)`

GetSchemauuidnotuniqueOk returns a tuple with the Schemauuidnotunique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemauuidnotunique

`func (o *ConsistencyError) SetSchemauuidnotunique(v string)`

SetSchemauuidnotunique sets Schemauuidnotunique field to given value.


### GetEntryuuidcorrupt

`func (o *ConsistencyError) GetEntryuuidcorrupt() int64`

GetEntryuuidcorrupt returns the Entryuuidcorrupt field if non-nil, zero value otherwise.

### GetEntryuuidcorruptOk

`func (o *ConsistencyError) GetEntryuuidcorruptOk() (*int64, bool)`

GetEntryuuidcorruptOk returns a tuple with the Entryuuidcorrupt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryuuidcorrupt

`func (o *ConsistencyError) SetEntryuuidcorrupt(v int64)`

SetEntryuuidcorrupt sets Entryuuidcorrupt field to given value.


### GetUuidindexcorrupt

`func (o *ConsistencyError) GetUuidindexcorrupt() string`

GetUuidindexcorrupt returns the Uuidindexcorrupt field if non-nil, zero value otherwise.

### GetUuidindexcorruptOk

`func (o *ConsistencyError) GetUuidindexcorruptOk() (*string, bool)`

GetUuidindexcorruptOk returns a tuple with the Uuidindexcorrupt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidindexcorrupt

`func (o *ConsistencyError) SetUuidindexcorrupt(v string)`

SetUuidindexcorrupt sets Uuidindexcorrupt field to given value.


### GetUuidnotunique

`func (o *ConsistencyError) GetUuidnotunique() string`

GetUuidnotunique returns the Uuidnotunique field if non-nil, zero value otherwise.

### GetUuidnotuniqueOk

`func (o *ConsistencyError) GetUuidnotuniqueOk() (*string, bool)`

GetUuidnotuniqueOk returns a tuple with the Uuidnotunique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidnotunique

`func (o *ConsistencyError) SetUuidnotunique(v string)`

SetUuidnotunique sets Uuidnotunique field to given value.


### GetRefintnotupheld

`func (o *ConsistencyError) GetRefintnotupheld() int64`

GetRefintnotupheld returns the Refintnotupheld field if non-nil, zero value otherwise.

### GetRefintnotupheldOk

`func (o *ConsistencyError) GetRefintnotupheldOk() (*int64, bool)`

GetRefintnotupheldOk returns a tuple with the Refintnotupheld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefintnotupheld

`func (o *ConsistencyError) SetRefintnotupheld(v int64)`

SetRefintnotupheld sets Refintnotupheld field to given value.


### GetMemberofinvalid

`func (o *ConsistencyError) GetMemberofinvalid() int64`

GetMemberofinvalid returns the Memberofinvalid field if non-nil, zero value otherwise.

### GetMemberofinvalidOk

`func (o *ConsistencyError) GetMemberofinvalidOk() (*int64, bool)`

GetMemberofinvalidOk returns a tuple with the Memberofinvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberofinvalid

`func (o *ConsistencyError) SetMemberofinvalid(v int64)`

SetMemberofinvalid sets Memberofinvalid field to given value.


### GetInvalidattributetype

`func (o *ConsistencyError) GetInvalidattributetype() string`

GetInvalidattributetype returns the Invalidattributetype field if non-nil, zero value otherwise.

### GetInvalidattributetypeOk

`func (o *ConsistencyError) GetInvalidattributetypeOk() (*string, bool)`

GetInvalidattributetypeOk returns a tuple with the Invalidattributetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidattributetype

`func (o *ConsistencyError) SetInvalidattributetype(v string)`

SetInvalidattributetype sets Invalidattributetype field to given value.


### GetInvalidspn

`func (o *ConsistencyError) GetInvalidspn() int64`

GetInvalidspn returns the Invalidspn field if non-nil, zero value otherwise.

### GetInvalidspnOk

`func (o *ConsistencyError) GetInvalidspnOk() (*int64, bool)`

GetInvalidspnOk returns a tuple with the Invalidspn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidspn

`func (o *ConsistencyError) SetInvalidspn(v int64)`

SetInvalidspn sets Invalidspn field to given value.


### GetChangelogdesynchronised

`func (o *ConsistencyError) GetChangelogdesynchronised() int64`

GetChangelogdesynchronised returns the Changelogdesynchronised field if non-nil, zero value otherwise.

### GetChangelogdesynchronisedOk

`func (o *ConsistencyError) GetChangelogdesynchronisedOk() (*int64, bool)`

GetChangelogdesynchronisedOk returns a tuple with the Changelogdesynchronised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogdesynchronised

`func (o *ConsistencyError) SetChangelogdesynchronised(v int64)`

SetChangelogdesynchronised sets Changelogdesynchronised field to given value.


### GetChangestatedesynchronised

`func (o *ConsistencyError) GetChangestatedesynchronised() int64`

GetChangestatedesynchronised returns the Changestatedesynchronised field if non-nil, zero value otherwise.

### GetChangestatedesynchronisedOk

`func (o *ConsistencyError) GetChangestatedesynchronisedOk() (*int64, bool)`

GetChangestatedesynchronisedOk returns a tuple with the Changestatedesynchronised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangestatedesynchronised

`func (o *ConsistencyError) SetChangestatedesynchronised(v int64)`

SetChangestatedesynchronised sets Changestatedesynchronised field to given value.


### GetRuvinconsistent

`func (o *ConsistencyError) GetRuvinconsistent() string`

GetRuvinconsistent returns the Ruvinconsistent field if non-nil, zero value otherwise.

### GetRuvinconsistentOk

`func (o *ConsistencyError) GetRuvinconsistentOk() (*string, bool)`

GetRuvinconsistentOk returns a tuple with the Ruvinconsistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuvinconsistent

`func (o *ConsistencyError) SetRuvinconsistent(v string)`

SetRuvinconsistent sets Ruvinconsistent field to given value.


### GetDeniedname

`func (o *ConsistencyError) GetDeniedname() string`

GetDeniedname returns the Deniedname field if non-nil, zero value otherwise.

### GetDeniednameOk

`func (o *ConsistencyError) GetDeniednameOk() (*string, bool)`

GetDeniednameOk returns a tuple with the Deniedname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedname

`func (o *ConsistencyError) SetDeniedname(v string)`

SetDeniedname sets Deniedname field to given value.


### GetKeyprovideruuidmissing

`func (o *ConsistencyError) GetKeyprovideruuidmissing() ConsistencyErrorOneOf14Keyprovideruuidmissing`

GetKeyprovideruuidmissing returns the Keyprovideruuidmissing field if non-nil, zero value otherwise.

### GetKeyprovideruuidmissingOk

`func (o *ConsistencyError) GetKeyprovideruuidmissingOk() (*ConsistencyErrorOneOf14Keyprovideruuidmissing, bool)`

GetKeyprovideruuidmissingOk returns a tuple with the Keyprovideruuidmissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyprovideruuidmissing

`func (o *ConsistencyError) SetKeyprovideruuidmissing(v ConsistencyErrorOneOf14Keyprovideruuidmissing)`

SetKeyprovideruuidmissing sets Keyprovideruuidmissing field to given value.


### GetKeyprovidernokeys

`func (o *ConsistencyError) GetKeyprovidernokeys() ConsistencyErrorOneOf14Keyprovideruuidmissing`

GetKeyprovidernokeys returns the Keyprovidernokeys field if non-nil, zero value otherwise.

### GetKeyprovidernokeysOk

`func (o *ConsistencyError) GetKeyprovidernokeysOk() (*ConsistencyErrorOneOf14Keyprovideruuidmissing, bool)`

GetKeyprovidernokeysOk returns a tuple with the Keyprovidernokeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyprovidernokeys

`func (o *ConsistencyError) SetKeyprovidernokeys(v ConsistencyErrorOneOf14Keyprovideruuidmissing)`

SetKeyprovidernokeys sets Keyprovidernokeys field to given value.


### GetKeyprovidernotfound

`func (o *ConsistencyError) GetKeyprovidernotfound() ConsistencyErrorOneOf16Keyprovidernotfound`

GetKeyprovidernotfound returns the Keyprovidernotfound field if non-nil, zero value otherwise.

### GetKeyprovidernotfoundOk

`func (o *ConsistencyError) GetKeyprovidernotfoundOk() (*ConsistencyErrorOneOf16Keyprovidernotfound, bool)`

GetKeyprovidernotfoundOk returns a tuple with the Keyprovidernotfound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyprovidernotfound

`func (o *ConsistencyError) SetKeyprovidernotfound(v ConsistencyErrorOneOf16Keyprovidernotfound)`

SetKeyprovidernotfound sets Keyprovidernotfound field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


