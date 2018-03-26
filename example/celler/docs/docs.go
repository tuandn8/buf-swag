// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-03-21 13:17:42.608622013 +0900 JST m=+0.058384190

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server celler server.",
        "title": "Swagger Example API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/api/v1",
    "paths": {
        "/accounts": {
            "get": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List accounts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name search by q",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Account"
                            }
                        }
                    },
                    "400": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "404": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "add by json account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a account",
                "parameters": [
                    {
                        "description": "Add account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.AddAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Account"
                        }
                    },
                    "400": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "404": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show a account",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Account"
                        }
                    },
                    "400": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "404": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a account",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Account"
                        }
                    },
                    "400": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "404": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update by json account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a account",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UpdateAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Account"
                        }
                    },
                    "400": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "404": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            }
        },
        "/accounts/{id}/images": {
            "post": {
                "description": "Upload file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Upload account image",
                "operationId": "account id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "account image",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    },
                    "400": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "404": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            }
        },
        "/admin/auth": {
            "post": {
                "description": "get admin info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Auth admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Admin"
                        }
                    },
                    "400": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "401": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "404": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            }
        },
        "/bottles": {
            "get": {
                "description": "get bottles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List bottles",
                "responses": {
                    "200": {
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Bottle"
                            }
                        }
                    },
                    "400": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "404": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            }
        },
        "/bottles/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show a bottle",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bottle ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Bottle"
                        }
                    },
                    "400": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "404": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            }
        },
        "/examples/calc": {
            "get": {
                "description": "plus",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "calc example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "used for calc",
                        "name": "val1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "used for calc",
                        "name": "val2",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "answer",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/examples/groups/{group_id}/accounts/{account_id}": {
            "get": {
                "description": "path params",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "path params example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group ID",
                        "name": "group_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "answer",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/examples/ping": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "controller.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "model.Account": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "account name"
                }
            }
        },
        "model.AddAccount": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "account name"
                }
            }
        },
        "model.Admin": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "admin name"
                }
            }
        },
        "model.Bottle": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "object",
                    "$ref": "#/definitions/model.Account"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "bottle_name"
                }
            }
        },
        "model.UpdateAccount": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "account name"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}