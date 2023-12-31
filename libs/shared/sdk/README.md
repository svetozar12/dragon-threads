# Go API client for openapi

This is a sample swagger for Fiber

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:3333*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*UserAPI* | [**V1UsersGet**](docs/UserAPI.md#v1usersget) | **Get** /v1/users | Get User List
*UserAPI* | [**V1UsersIdPut**](docs/UserAPI.md#v1usersidput) | **Put** /v1/users/{id} | Update User
*UserAPI* | [**V1UsersPost**](docs/UserAPI.md#v1userspost) | **Post** /v1/users | Create User
*UserAPI* | [**V1UsersUserIdDelete**](docs/UserAPI.md#v1usersuseriddelete) | **Delete** /v1/users/{userId} | Delete user by ID
*UserAPI* | [**V1UsersUserIdGet**](docs/UserAPI.md#v1usersuseridget) | **Get** /v1/users/{userId} | Get user by ID


## Documentation For Models

 - [CommonCommonErrorSchema](docs/CommonCommonErrorSchema.md)
 - [CommonCommonPaginationSchema](docs/CommonCommonPaginationSchema.md)
 - [EntitiesUser](docs/EntitiesUser.md)
 - [GormDeletedAt](docs/GormDeletedAt.md)
 - [UsersUpdateUserSchema](docs/UsersUpdateUserSchema.md)
 - [UsersUserListSchema](docs/UsersUserListSchema.md)
 - [UsersUserSchema](docs/UsersUserSchema.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

fiber@swagger.io

