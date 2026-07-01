# UatStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**SessionId** | **string** |  | 
**State** | [**UatStatusState**](UatStatusState.md) |  | 
**IssuedAt** | **time.Time** |  | 
**Purpose** | [**UatPurposeStatus**](UatPurposeStatus.md) |  | 

## Methods

### NewUatStatus

`func NewUatStatus(accountId string, sessionId string, state UatStatusState, issuedAt time.Time, purpose UatPurposeStatus, ) *UatStatus`

NewUatStatus instantiates a new UatStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUatStatusWithDefaults

`func NewUatStatusWithDefaults() *UatStatus`

NewUatStatusWithDefaults instantiates a new UatStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UatStatus) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UatStatus) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UatStatus) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSessionId

`func (o *UatStatus) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *UatStatus) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *UatStatus) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetState

`func (o *UatStatus) GetState() UatStatusState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UatStatus) GetStateOk() (*UatStatusState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UatStatus) SetState(v UatStatusState)`

SetState sets State field to given value.


### GetIssuedAt

`func (o *UatStatus) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *UatStatus) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *UatStatus) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### GetPurpose

`func (o *UatStatus) GetPurpose() UatPurposeStatus`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *UatStatus) GetPurposeOk() (*UatPurposeStatus, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *UatStatus) SetPurpose(v UatPurposeStatus)`

SetPurpose sets Purpose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


