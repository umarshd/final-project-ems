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
        "contact": {
            "name": "\"Umar Sahid\"",
            "email": "\"umarhshid@gmail.com\""
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/login": {
            "post": {
                "description": "Login Admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Login Admin",
                "parameters": [
                    {
                        "description": "Login Admin",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestLoginAdmin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseAdminLogin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/admin/register": {
            "post": {
                "description": "Create Admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Create Admin",
                "parameters": [
                    {
                        "description": "Create Admin",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestAdmin"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseAdmin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/carts": {
            "get": {
                "description": "Get All Cart By User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Get All Cart By User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Cart"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Create Cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Cart"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/carts/items": {
            "post": {
                "description": "Create Cart Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Create Cart Item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Create Cart Item",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestCartItem"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.CartItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/product-category": {
            "get": {
                "description": "Get All Product Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Category"
                ],
                "summary": "Get All Product Category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseProductCategory"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Product Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Category"
                ],
                "summary": "Create Product Category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Create Product Category",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestProductCategory"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseProductCategory"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Get All Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get All Product",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseProduct"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Create Product",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestProduct"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseProduct"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products/filter": {
            "get": {
                "description": "Filter Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Filter Product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Price Min",
                        "name": "priceMin",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Price Max",
                        "name": "priceMax",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseProduct"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get Product By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Product By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseProduct"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/checkout": {
            "post": {
                "description": "Create Checkout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checkout"
                ],
                "summary": "Create Checkout",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Create Checkout",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestCheckout"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseCheckout"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "Login User",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestLoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseUserLogin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Create User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Create User",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Cart": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_checkout": {
                    "type": "boolean"
                },
                "total": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/entity.User"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.CartItem": {
            "type": "object",
            "properties": {
                "cart": {
                    "$ref": "#/definitions/entity.Cart"
                },
                "cart_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "product": {
                    "$ref": "#/definitions/entity.Product"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "entity.Product": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "product_category": {
                    "$ref": "#/definitions/entity.ProductCategory"
                },
                "quantity": {
                    "type": "integer"
                },
                "sku": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.ProductCategory": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
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
        "entity.RequestAdmin": {
            "type": "object",
            "required": [
                "email",
                "nama",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entity.RequestCartItem": {
            "type": "object",
            "required": [
                "cart_id",
                "product_id",
                "quantity"
            ],
            "properties": {
                "cart_id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "entity.RequestCheckout": {
            "type": "object",
            "required": [
                "cart_id",
                "image_url",
                "user_id"
            ],
            "properties": {
                "cart_id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.RequestLoginAdmin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entity.RequestLoginUser": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entity.RequestProduct": {
            "type": "object",
            "required": [
                "category_id",
                "description",
                "image_url",
                "name",
                "price",
                "quantity",
                "sku"
            ],
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "sku": {
                    "type": "string"
                }
            }
        },
        "entity.RequestProductCategory": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.RequestUser": {
            "type": "object",
            "required": [
                "email",
                "gender",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "entity.ResponseAdmin": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "entity.ResponseAdminLogin": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "entity.ResponseCheckout": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "entity.ResponseProduct": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "sku": {
                    "type": "string"
                }
            }
        },
        "entity.ResponseProductCategory": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.ResponseUser": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "entity.ResponseUserLogin": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/entity.UserProfile"
                },
                "profile_id": {
                    "type": "integer"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.UserProfile": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Final Project ECommerce System API Documentation",
	Description:      "This is the documentation for the Final Project ECommerce System API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
