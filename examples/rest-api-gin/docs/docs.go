// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/coditori/javatori/blob/master/LICENSE",
        "contact": {
            "name": "Ario Afrashteh"
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
        "/auth/videos": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Takes nothing. Return list of videos.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "videos"
                ],
                "summary": "Return all videos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Video"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Takes a video JSON and save it in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "videos"
                ],
                "summary": "Save a new video",
                "parameters": [
                    {
                        "description": "Video JSON",
                        "name": "video",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Video"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Video"
                        }
                    }
                }
            }
        },
        "/auth/videos/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Takes a video id and return it from DB.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "videos"
                ],
                "summary": "Find an existing video",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "VideoId",
                        "name": "uint",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Video"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Takes a video JSON and video id and update it in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "videos"
                ],
                "summary": "Update an existing video",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "VideoId",
                        "name": "uint",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Video JSON",
                        "name": "video",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Video"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Video"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Takes a video id and detele it in DB. Return nothing.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "videos"
                ],
                "summary": "Delete an existing video",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "VideoId",
                        "name": "uint",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Video"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Takes nothing. Return list of videos.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Login by username and password",
                "parameters": [
                    {
                        "description": "Login JSON",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Return poing if server is fine.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "Ping the API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Person": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 130,
                    "minimum": 1
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                }
            }
        },
        "models.Video": {
            "type": "object",
            "required": [
                "author",
                "url"
            ],
            "properties": {
                "author": {
                    "$ref": "#/definitions/models.Person"
                },
                "description": {
                    "type": "string",
                    "maxLength": 200
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                },
                "url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Gin Book Service",
	Description:      "A book management service API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
