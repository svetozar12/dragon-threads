# EntitiesUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | **int32** |  | 
**SubDragonId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewEntitiesUser

`func NewEntitiesUser(id int32, ) *EntitiesUser`

NewEntitiesUser instantiates a new EntitiesUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitiesUserWithDefaults

`func NewEntitiesUserWithDefaults() *EntitiesUser`

NewEntitiesUserWithDefaults instantiates a new EntitiesUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *EntitiesUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EntitiesUser) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EntitiesUser) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EntitiesUser) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAvatar

`func (o *EntitiesUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *EntitiesUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *EntitiesUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *EntitiesUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBio

`func (o *EntitiesUser) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *EntitiesUser) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *EntitiesUser) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *EntitiesUser) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntitiesUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntitiesUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntitiesUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntitiesUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EntitiesUser) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntitiesUser) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntitiesUser) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntitiesUser) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEmail

`func (o *EntitiesUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EntitiesUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EntitiesUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EntitiesUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *EntitiesUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitiesUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitiesUser) SetId(v int32)`

SetId sets Id field to given value.


### GetSubDragonId

`func (o *EntitiesUser) GetSubDragonId() int32`

GetSubDragonId returns the SubDragonId field if non-nil, zero value otherwise.

### GetSubDragonIdOk

`func (o *EntitiesUser) GetSubDragonIdOk() (*int32, bool)`

GetSubDragonIdOk returns a tuple with the SubDragonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDragonId

`func (o *EntitiesUser) SetSubDragonId(v int32)`

SetSubDragonId sets SubDragonId field to given value.

### HasSubDragonId

`func (o *EntitiesUser) HasSubDragonId() bool`

HasSubDragonId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EntitiesUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntitiesUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntitiesUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntitiesUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *EntitiesUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EntitiesUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EntitiesUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EntitiesUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


