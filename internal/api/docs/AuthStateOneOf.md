# AuthStateOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choose** | [**[]AuthMech**](AuthMech.md) | You need to select how you want to proceed. | 

## Methods

### NewAuthStateOneOf

`func NewAuthStateOneOf(choose []AuthMech, ) *AuthStateOneOf`

NewAuthStateOneOf instantiates a new AuthStateOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthStateOneOfWithDefaults

`func NewAuthStateOneOfWithDefaults() *AuthStateOneOf`

NewAuthStateOneOfWithDefaults instantiates a new AuthStateOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoose

`func (o *AuthStateOneOf) GetChoose() []AuthMech`

GetChoose returns the Choose field if non-nil, zero value otherwise.

### GetChooseOk

`func (o *AuthStateOneOf) GetChooseOk() (*[]AuthMech, bool)`

GetChooseOk returns a tuple with the Choose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoose

`func (o *AuthStateOneOf) SetChoose(v []AuthMech)`

SetChoose sets Choose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


