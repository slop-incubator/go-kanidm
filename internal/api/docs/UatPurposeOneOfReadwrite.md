# UatPurposeOneOfReadwrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **NullableTime** | If none, there is no expiry, and this is always rw. If there is an expiry, check that the current time &lt; expiry. | [optional] 

## Methods

### NewUatPurposeOneOfReadwrite

`func NewUatPurposeOneOfReadwrite() *UatPurposeOneOfReadwrite`

NewUatPurposeOneOfReadwrite instantiates a new UatPurposeOneOfReadwrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUatPurposeOneOfReadwriteWithDefaults

`func NewUatPurposeOneOfReadwriteWithDefaults() *UatPurposeOneOfReadwrite`

NewUatPurposeOneOfReadwriteWithDefaults instantiates a new UatPurposeOneOfReadwrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *UatPurposeOneOfReadwrite) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *UatPurposeOneOfReadwrite) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *UatPurposeOneOfReadwrite) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *UatPurposeOneOfReadwrite) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### SetExpiryNil

`func (o *UatPurposeOneOfReadwrite) SetExpiryNil(b bool)`

 SetExpiryNil sets the value for Expiry to be an explicit nil

### UnsetExpiry
`func (o *UatPurposeOneOfReadwrite) UnsetExpiry()`

UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


