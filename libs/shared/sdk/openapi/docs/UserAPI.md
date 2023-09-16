# \UserAPI

All URIs are relative to *http://localhost:3333*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UsersGet**](UserAPI.md#V1UsersGet) | **Get** /v1/users | Get User List
[**V1UsersIdPut**](UserAPI.md#V1UsersIdPut) | **Put** /v1/users/{id} | Update User
[**V1UsersPost**](UserAPI.md#V1UsersPost) | **Post** /v1/users | Create User
[**V1UsersUserIdDelete**](UserAPI.md#V1UsersUserIdDelete) | **Delete** /v1/users/{userId} | Delete user by ID
[**V1UsersUserIdGet**](UserAPI.md#V1UsersUserIdGet) | **Get** /v1/users/{userId} | Get user by ID



## V1UsersGet

> UsersUserListSchema V1UsersGet(ctx).Page(page).PageSize(pageSize).GetBy(getBy).GetByValue(getByValue).Execute()

Get User List

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/sdk"
)

func main() {
    page := int32(56) // int32 | Page number (default: 1) (optional)
    pageSize := int32(56) // int32 | Number of items per page (default: 10) (optional)
    getBy := "getBy_example" // string | Get users by field (optional) (optional)
    getByValue := int32(56) // int32 | Get users by field value (optional) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.V1UsersGet(context.Background()).Page(page).PageSize(pageSize).GetBy(getBy).GetByValue(getByValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UsersGet`: UsersUserListSchema
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (default: 1) | 
 **pageSize** | **int32** | Number of items per page (default: 10) | 
 **getBy** | **string** | Get users by field (optional) | 
 **getByValue** | **int32** | Get users by field value (optional) | 

### Return type

[**UsersUserListSchema**](UsersUserListSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersIdPut

> EntitiesUser V1UsersIdPut(ctx, id).Request(request).Execute()

Update User

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/sdk"
)

func main() {
    id := "id_example" // string | User ID
    request := *openapiclient.NewUsersUpdateUserSchema() // UsersUpdateUserSchema | Request body for updating user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.V1UsersIdPut(context.Background(), id).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UsersIdPut`: EntitiesUser
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**UsersUpdateUserSchema**](UsersUpdateUserSchema.md) | Request body for updating user | 

### Return type

[**EntitiesUser**](EntitiesUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersPost

> EntitiesUser V1UsersPost(ctx).Request(request).Execute()

Create User

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/sdk"
)

func main() {
    request := *openapiclient.NewUsersUserSchema("Email_example", "Username_example") // UsersUserSchema | query params

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.V1UsersPost(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UsersPost`: EntitiesUser
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**UsersUserSchema**](UsersUserSchema.md) | query params | 

### Return type

[**EntitiesUser**](EntitiesUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersUserIdDelete

> string V1UsersUserIdDelete(ctx, userId).Execute()

Delete user by ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/sdk"
)

func main() {
    userId := int32(56) // int32 | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.V1UsersUserIdDelete(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersUserIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UsersUserIdDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersUserIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersUserIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersUserIdGet

> UsersUserSchema V1UsersUserIdGet(ctx, userId).Execute()

Get user by ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/sdk"
)

func main() {
    userId := int32(56) // int32 | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.V1UsersUserIdGet(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UsersUserIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UsersUserIdGet`: UsersUserSchema
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsersUserSchema**](UsersUserSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

