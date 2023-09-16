# UsersUserListSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EntitiesUser**](EntitiesUser.md) |  | [optional] 
**Pagination** | Pointer to [**CommonCommonPaginationSchema**](CommonCommonPaginationSchema.md) |  | [optional] 

## Methods

### NewUsersUserListSchema

`func NewUsersUserListSchema() *UsersUserListSchema`

NewUsersUserListSchema instantiates a new UsersUserListSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUserListSchemaWithDefaults

`func NewUsersUserListSchemaWithDefaults() *UsersUserListSchema`

NewUsersUserListSchemaWithDefaults instantiates a new UsersUserListSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UsersUserListSchema) GetData() []EntitiesUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsersUserListSchema) GetDataOk() (*[]EntitiesUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsersUserListSchema) SetData(v []EntitiesUser)`

SetData sets Data field to given value.

### HasData

`func (o *UsersUserListSchema) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *UsersUserListSchema) GetPagination() CommonCommonPaginationSchema`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *UsersUserListSchema) GetPaginationOk() (*CommonCommonPaginationSchema, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *UsersUserListSchema) SetPagination(v CommonCommonPaginationSchema)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *UsersUserListSchema) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


