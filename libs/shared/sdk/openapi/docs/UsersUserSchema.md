# UsersUserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**SubDragonId** | Pointer to **int32** |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewUsersUserSchema

`func NewUsersUserSchema(email string, username string, ) *UsersUserSchema`

NewUsersUserSchema instantiates a new UsersUserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUserSchemaWithDefaults

`func NewUsersUserSchemaWithDefaults() *UsersUserSchema`

NewUsersUserSchemaWithDefaults instantiates a new UsersUserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *UsersUserSchema) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UsersUserSchema) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UsersUserSchema) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UsersUserSchema) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBio

`func (o *UsersUserSchema) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UsersUserSchema) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UsersUserSchema) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UsersUserSchema) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetEmail

`func (o *UsersUserSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersUserSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersUserSchema) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetSubDragonId

`func (o *UsersUserSchema) GetSubDragonId() int32`

GetSubDragonId returns the SubDragonId field if non-nil, zero value otherwise.

### GetSubDragonIdOk

`func (o *UsersUserSchema) GetSubDragonIdOk() (*int32, bool)`

GetSubDragonIdOk returns a tuple with the SubDragonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDragonId

`func (o *UsersUserSchema) SetSubDragonId(v int32)`

SetSubDragonId sets SubDragonId field to given value.

### HasSubDragonId

`func (o *UsersUserSchema) HasSubDragonId() bool`

HasSubDragonId returns a boolean if a field has been set.

### GetUsername

`func (o *UsersUserSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UsersUserSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UsersUserSchema) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


