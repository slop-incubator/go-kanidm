# TotpSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accountname** | **string** |  | 
**Issuer** | **string** | User-facing name of the system, issuer of the TOTP | 
**Secret** | **[]int32** |  | 
**Algo** | [**TotpAlgo**](TotpAlgo.md) |  | 
**Step** | **int64** |  | 
**Digits** | **int32** |  | 

## Methods

### NewTotpSecret

`func NewTotpSecret(accountname string, issuer string, secret []int32, algo TotpAlgo, step int64, digits int32, ) *TotpSecret`

NewTotpSecret instantiates a new TotpSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotpSecretWithDefaults

`func NewTotpSecretWithDefaults() *TotpSecret`

NewTotpSecretWithDefaults instantiates a new TotpSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountname

`func (o *TotpSecret) GetAccountname() string`

GetAccountname returns the Accountname field if non-nil, zero value otherwise.

### GetAccountnameOk

`func (o *TotpSecret) GetAccountnameOk() (*string, bool)`

GetAccountnameOk returns a tuple with the Accountname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountname

`func (o *TotpSecret) SetAccountname(v string)`

SetAccountname sets Accountname field to given value.


### GetIssuer

`func (o *TotpSecret) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *TotpSecret) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *TotpSecret) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetSecret

`func (o *TotpSecret) GetSecret() []int32`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TotpSecret) GetSecretOk() (*[]int32, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TotpSecret) SetSecret(v []int32)`

SetSecret sets Secret field to given value.


### GetAlgo

`func (o *TotpSecret) GetAlgo() TotpAlgo`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *TotpSecret) GetAlgoOk() (*TotpAlgo, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *TotpSecret) SetAlgo(v TotpAlgo)`

SetAlgo sets Algo field to given value.


### GetStep

`func (o *TotpSecret) GetStep() int64`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *TotpSecret) GetStepOk() (*int64, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *TotpSecret) SetStep(v int64)`

SetStep sets Step field to given value.


### GetDigits

`func (o *TotpSecret) GetDigits() int32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *TotpSecret) GetDigitsOk() (*int32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *TotpSecret) SetDigits(v int32)`

SetDigits sets Digits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


