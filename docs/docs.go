// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/create": {
            "post": {
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-types_User"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "description": "查询用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查询用户",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "skip",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyList-types_User"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "post": {
                "description": "修改用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-types_User"
                        }
                    }
                }
            }
        },
        "/user/{id}/disable": {
            "get": {
                "description": "禁用用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "禁用用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-types_User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "curd.ReplyData-types_User": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/types.User"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "curd.ReplyList-types_User": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.User"
                    }
                },
                "error": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "types.User": {
            "type": "object",
            "properties": {
                "admin": {
                    "type": "boolean"
                },
                "cellPhone": {
                    "type": "string"
                },
                "create": {
                    "type": "string"
                },
                "disable": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	//LeftDelim:        "{{",
	//RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}