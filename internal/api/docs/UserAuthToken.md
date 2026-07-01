# UserAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** |  | 
**IssuedAt** | **time.Time** |  | 
**Expiry** | Pointer to **NullableTime** | If none, there is no expiry, and this is always valid. If there is an expiry, check that the current time &lt; expiry. | [optional] 
**Purpose** | [**UatPurpose**](UatPurpose.md) |  | 
**Uuid** | **string** |  | 
**Displayname** | **string** |  | 
**Spn** | **string** |  | 
**MailPrimary** | Pointer to **NullableString** |  | [optional] 
**UiHints** | [**[]UiHint**](UiHint.md) |  | 
**LimitSearchMaxResults** | Pointer to **NullableInt64** |  | [optional] 
**LimitSearchMaxFilterTest** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewUserAuthToken

`func NewUserAuthToken(sessionId string, issuedAt time.Time, purpose UatPurpose, uuid string, displayname string, spn string, uiHints []UiHint, ) *UserAuthToken`

NewUserAuthToken instantiates a new UserAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAuthTokenWithDefaults

`func NewUserAuthTokenWithDefaults() *UserAuthToken`

NewUserAuthTokenWithDefaults instantiates a new UserAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *UserAuthToken) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *UserAuthToken) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *UserAuthToken) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetIssuedAt

`func (o *UserAuthToken) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *UserAuthToken) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *UserAuthToken) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### GetExpiry

`func (o *UserAuthToken) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *UserAuthToken) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *UserAuthToken) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *UserAuthToken) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### SetExpiryNil

`func (o *UserAuthToken) SetExpiryNil(b bool)`

 SetExpiryNil sets the value for Expiry to be an explicit nil

### UnsetExpiry
`func (o *UserAuthToken) UnsetExpiry()`

UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
### GetPurpose

`func (o *UserAuthToken) GetPurpose() UatPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *UserAuthToken) GetPurposeOk() (*UatPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *UserAuthToken) SetPurpose(v UatPurpose)`

SetPurpose sets Purpose field to given value.


### GetUuid

`func (o *UserAuthToken) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserAuthToken) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserAuthToken) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetDisplayname

`func (o *UserAuthToken) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *UserAuthToken) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *UserAuthToken) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.


### GetSpn

`func (o *UserAuthToken) GetSpn() string`

GetSpn returns the Spn field if non-nil, zero value otherwise.

### GetSpnOk

`func (o *UserAuthToken) GetSpnOk() (*string, bool)`

GetSpnOk returns a tuple with the Spn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpn

`func (o *UserAuthToken) SetSpn(v string)`

SetSpn sets Spn field to given value.


### GetMailPrimary

`func (o *UserAuthToken) GetMailPrimary() string`

GetMailPrimary returns the MailPrimary field if non-nil, zero value otherwise.

### GetMailPrimaryOk

`func (o *UserAuthToken) GetMailPrimaryOk() (*string, bool)`

GetMailPrimaryOk returns a tuple with the MailPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailPrimary

`func (o *UserAuthToken) SetMailPrimary(v string)`

SetMailPrimary sets MailPrimary field to given value.

### HasMailPrimary

`func (o *UserAuthToken) HasMailPrimary() bool`

HasMailPrimary returns a boolean if a field has been set.

### SetMailPrimaryNil

`func (o *UserAuthToken) SetMailPrimaryNil(b bool)`

 SetMailPrimaryNil sets the value for MailPrimary to be an explicit nil

### UnsetMailPrimary
`func (o *UserAuthToken) UnsetMailPrimary()`

UnsetMailPrimary ensures that no value is present for MailPrimary, not even an explicit nil
### GetUiHints

`func (o *UserAuthToken) GetUiHints() []UiHint`

GetUiHints returns the UiHints field if non-nil, zero value otherwise.

### GetUiHintsOk

`func (o *UserAuthToken) GetUiHintsOk() (*[]UiHint, bool)`

GetUiHintsOk returns a tuple with the UiHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHints

`func (o *UserAuthToken) SetUiHints(v []UiHint)`

SetUiHints sets UiHints field to given value.


### GetLimitSearchMaxResults

`func (o *UserAuthToken) GetLimitSearchMaxResults() int64`

GetLimitSearchMaxResults returns the LimitSearchMaxResults field if non-nil, zero value otherwise.

### GetLimitSearchMaxResultsOk

`func (o *UserAuthToken) GetLimitSearchMaxResultsOk() (*int64, bool)`

GetLimitSearchMaxResultsOk returns a tuple with the LimitSearchMaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSearchMaxResults

`func (o *UserAuthToken) SetLimitSearchMaxResults(v int64)`

SetLimitSearchMaxResults sets LimitSearchMaxResults field to given value.

### HasLimitSearchMaxResults

`func (o *UserAuthToken) HasLimitSearchMaxResults() bool`

HasLimitSearchMaxResults returns a boolean if a field has been set.

### SetLimitSearchMaxResultsNil

`func (o *UserAuthToken) SetLimitSearchMaxResultsNil(b bool)`

 SetLimitSearchMaxResultsNil sets the value for LimitSearchMaxResults to be an explicit nil

### UnsetLimitSearchMaxResults
`func (o *UserAuthToken) UnsetLimitSearchMaxResults()`

UnsetLimitSearchMaxResults ensures that no value is present for LimitSearchMaxResults, not even an explicit nil
### GetLimitSearchMaxFilterTest

`func (o *UserAuthToken) GetLimitSearchMaxFilterTest() int64`

GetLimitSearchMaxFilterTest returns the LimitSearchMaxFilterTest field if non-nil, zero value otherwise.

### GetLimitSearchMaxFilterTestOk

`func (o *UserAuthToken) GetLimitSearchMaxFilterTestOk() (*int64, bool)`

GetLimitSearchMaxFilterTestOk returns a tuple with the LimitSearchMaxFilterTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSearchMaxFilterTest

`func (o *UserAuthToken) SetLimitSearchMaxFilterTest(v int64)`

SetLimitSearchMaxFilterTest sets LimitSearchMaxFilterTest field to given value.

### HasLimitSearchMaxFilterTest

`func (o *UserAuthToken) HasLimitSearchMaxFilterTest() bool`

HasLimitSearchMaxFilterTest returns a boolean if a field has been set.

### SetLimitSearchMaxFilterTestNil

`func (o *UserAuthToken) SetLimitSearchMaxFilterTestNil(b bool)`

 SetLimitSearchMaxFilterTestNil sets the value for LimitSearchMaxFilterTest to be an explicit nil

### UnsetLimitSearchMaxFilterTest
`func (o *UserAuthToken) UnsetLimitSearchMaxFilterTest()`

UnsetLimitSearchMaxFilterTest ensures that no value is present for LimitSearchMaxFilterTest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


