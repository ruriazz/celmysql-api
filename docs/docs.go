// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all the existing auths",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auths"
                ],
                "summary": "List existing auths",
                "parameters": [
                    {
                        "description": "auth",
                        "name": "authDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PayloadLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/common.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "create user",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PayloadRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bank/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create new bank",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "banks"
                ],
                "summary": "create bank",
                "parameters": [
                    {
                        "description": "bank",
                        "name": "bank",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBankDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bank/delete/{oid}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete bank by oid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "banks"
                ],
                "summary": "delete bank",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oid",
                        "name": "oid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bank/q": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all the existing banks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "banks"
                ],
                "summary": "List existing banks",
                "parameters": [
                    {
                        "description": "bank",
                        "name": "bankDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBankDto"
                        }
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "page Index",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "page Size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/common.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/bank/{oid}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update new bank",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "banks"
                ],
                "summary": "find bank by oid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oid",
                        "name": "oid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/common.DefaultResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update existing bank",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "banks"
                ],
                "summary": "update bank by oid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oid",
                        "name": "oid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "bank",
                        "name": "bank",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateBankDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/image-file/delete/{oid}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete imageFile by oid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "imageFiles"
                ],
                "summary": "delete imageFile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oid",
                        "name": "oid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/image-file/q": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all the existing imageFiles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "imageFiles"
                ],
                "summary": "List existing imageFiles",
                "parameters": [
                    {
                        "description": "imageFile",
                        "name": "imageFileDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateImageFileDto"
                        }
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "page Index",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "page Size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/common.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/image-file/upload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create new imageFile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "imageFiles"
                ],
                "summary": "create imageFile",
                "parameters": [
                    {
                        "description": "imageFile",
                        "name": "imageFile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateImageFileDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/image-file/{oid}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update new imageFile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "imageFiles"
                ],
                "summary": "find imageFile by oid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oid",
                        "name": "oid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/common.DefaultResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update existing imageFile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "imageFiles"
                ],
                "summary": "update imageFile by oid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oid",
                        "name": "oid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "imageFile",
                        "name": "imageFile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateImageFileDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/raja-ongkir": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all the existing rajaOngkirs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rajaOngkirs"
                ],
                "summary": "List existing rajaOngkirs",
                "parameters": [
                    {
                        "description": "rajaOngkir",
                        "name": "rajaOngkirDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateRajaOngkirDto"
                        }
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "page Index",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "description": "page Size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/common.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/send-email": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all the existing Send Email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Send Email"
                ],
                "summary": "List existing Send Email",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/common.DefaultResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.DefaultResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "developerMessage": {
                    "type": "string"
                },
                "http_status": {
                    "type": "string"
                },
                "resultCode": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateBankDto": {
            "type": "object",
            "required": [
                "bankCode",
                "bankName"
            ],
            "properties": {
                "bankCode": {
                    "type": "string"
                },
                "bankName": {
                    "type": "string"
                },
                "userInserted": {
                    "type": "string"
                }
            }
        },
        "dto.CreateImageFileDto": {
            "type": "object",
            "properties": {
                "fileName": {
                    "type": "string"
                },
                "fileUrl": {
                    "type": "string"
                },
                "userInserted": {
                    "type": "string"
                }
            }
        },
        "dto.CreateRajaOngkirDto": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "dto.PayloadLogin": {
            "type": "object",
            "properties": {
                "emailName": {
                    "description": "ID        uuid.UUID ` + "`" + `json:\"id\"` + "`" + `",
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.PayloadRegister": {
            "type": "object",
            "properties": {
                "emailName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateBankDto": {
            "type": "object",
            "required": [
                "bankCode",
                "bankName"
            ],
            "properties": {
                "bankCode": {
                    "type": "string"
                },
                "bankName": {
                    "type": "string"
                },
                "lastUserId": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateImageFileDto": {
            "type": "object",
            "properties": {
                "fileName": {
                    "type": "string"
                },
                "fileUrl": {
                    "type": "string"
                },
                "lastUserId": {
                    "type": "string"
                },
                "oid": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
