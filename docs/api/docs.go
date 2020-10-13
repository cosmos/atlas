// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package api

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/keywords": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "keywords"
                ],
                "summary": "Return a paginated set of all keywords",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "pagination cursor",
                        "name": "cursor",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "pagination limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputil.PaginationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/me": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get the current authenticated user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserJSON"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/me/tokens": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all API tokens by user ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserTokenJSON"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a user API token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserTokenJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/me/tokens/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Revoke a user API token by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserTokenJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/modules": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Return a paginated set of all Cosmos SDK modules",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "pagination cursor",
                        "name": "cursor",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "pagination limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputil.PaginationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "APIKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Publish a Cosmos SDK module",
                "parameters": [
                    {
                        "description": "module manifest",
                        "name": "manifest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.Manifest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ModuleJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/modules/search": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Search for Cosmos SDK modules by name, team, description and keywords",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "pagination cursor",
                        "name": "cursor",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "pagination limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "search criteria",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputil.PaginationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/modules/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Get a Cosmos SDK module by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "module ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ModuleJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/modules/{id}/authors": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Get all authors for a Cosmos SDK module by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "module ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserJSON"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/modules/{id}/keywords": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Get all keywords for a Cosmos SDK module by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "module ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.KeywordJSON"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/modules/{id}/versions": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Get all versions for a Cosmos SDK module by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "module ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ModuleVersionJSON"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Return a paginated set of all users",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "pagination cursor",
                        "name": "cursor",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "pagination limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputil.PaginationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get a user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}/modules": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Return a paginated set of all Cosmos SDK modules by user ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ModuleJSON"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.ErrResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "httputil.PaginationResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "cursor": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "results": {
                    "type": "object"
                }
            }
        },
        "models.BugTracker": {
            "type": "object",
            "properties": {
                "contact": {
                    "type": "string"
                },
                "module_id": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.Keyword": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.KeywordJSON": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.ModuleJSON": {
            "type": "object",
            "properties": {
                "authors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "bug_tracker": {
                    "type": "object",
                    "$ref": "#/definitions/models.BugTracker"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "documentation": {
                    "type": "string"
                },
                "homepage": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "keywords": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Keyword"
                    }
                },
                "name": {
                    "type": "string"
                },
                "owners": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "repo": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "versions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ModuleVersion"
                    }
                }
            }
        },
        "models.ModuleVersion": {
            "type": "object",
            "properties": {
                "module_id": {
                    "type": "integer"
                },
                "sdk_compat": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.ModuleVersionJSON": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "module_id": {
                    "type": "integer"
                },
                "sdk_compat": {
                    "type": "object"
                },
                "updated_at": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "avatar_url": {
                    "type": "string"
                },
                "gravatar_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.UserJSON": {
            "type": "object",
            "properties": {
                "avatar_url": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "gravatar_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.UserTokenJSON": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "revoked": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "v1.AuthorsManifest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "v1.BugTackerManifest": {
            "type": "object",
            "properties": {
                "contact": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "v1.Manifest": {
            "type": "object",
            "required": [
                "authors",
                "version"
            ],
            "properties": {
                "authors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.AuthorsManifest"
                    }
                },
                "bug_tracker": {
                    "type": "object",
                    "$ref": "#/definitions/v1.BugTackerManifest"
                },
                "module": {
                    "type": "object",
                    "$ref": "#/definitions/v1.ModuleManifest"
                },
                "version": {
                    "type": "object",
                    "$ref": "#/definitions/v1.VersionManifest"
                }
            }
        },
        "v1.ModuleManifest": {
            "type": "object",
            "required": [
                "name",
                "repo",
                "team"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "documentation": {
                    "type": "string"
                },
                "homepage": {
                    "type": "string"
                },
                "keywords": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "repo": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                }
            }
        },
        "v1.VersionManifest": {
            "type": "object",
            "required": [
                "version"
            ],
            "properties": {
                "sdk_compat": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "APIKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
