// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number (default: 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page (default: 10)",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "sub_dragon_id"
                        ],
                        "type": "string",
                        "description": "Get users by field (optional)",
                        "name": "getBy",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Get users by field value (optional)",
                        "name": "getByValue",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/users.UserListSchema"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.CommonErrorSchema"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.UserSchema"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.CommonErrorSchema"
                        }
                    }
                }
            }
        },
        "/v1/users/{userId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/users.UserSchema"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.CommonErrorSchema"
                        }
                    },
                    "404": {
                        "description": "Not Found Request",
                        "schema": {
                            "$ref": "#/definitions/common.CommonErrorSchema"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.CommonErrorSchema"
                        }
                    },
                    "404": {
                        "description": "Not Found Request",
                        "schema": {
                            "$ref": "#/definitions/common.CommonErrorSchema"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.CommonErrorSchema": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "common.CommonPaginationSchema": {
            "type": "object",
            "properties": {
                "hasNext": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "totalCount": {
                    "type": "integer"
                }
            }
        },
        "entities.User": {
            "type": "object",
            "required": [
                "email",
                "id",
                "username"
            ],
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "avatar": {
                    "type": "string"
                },
                "bio": {
                    "type": "string",
                    "maxLength": 250
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "subDragonId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "users.UserListSchema": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.User"
                    }
                },
                "pagination": {
                    "$ref": "#/definitions/common.CommonPaginationSchema"
                }
            }
        },
        "users.UserSchema": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "bio": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "subDragonId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3333",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
