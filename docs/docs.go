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
        "/companies/{ticker}": {
            "get": {
                "description": "get company info by Ticker",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "companies"
                ],
                "summary": "Get Company info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ticker",
                        "name": "ticker",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1Companies.GetCompanyInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/company-stocks": {
            "get": {
                "description": "get the daily information about company stocks",
                "tags": [
                    "company_stock"
                ],
                "summary": "Get Daily Company Stock",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/company-stocks/buy/{user_id}": {
            "post": {
                "description": "buy the selected amount of company stock",
                "tags": [
                    "company_stock"
                ],
                "summary": "BuyCompanyStock",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1companystock.BuyCompanyStockRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserTransactions"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/company-stocks/sell/{user_id}": {
            "post": {
                "description": "sell the selected amount of company stock",
                "tags": [
                    "company_stock"
                ],
                "summary": "SellCompanyStock",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1companystock.SellCompanyStockRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserTransactions"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/company-stocks/{user_id}": {
            "post": {
                "description": "buy or sell the selected amount of company stock",
                "tags": [
                    "company_stock"
                ],
                "summary": "BuyOrSellCompanyStock",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1companystock.BuyOrSellCompanyStocksRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserTransactions"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user-assets/profits/{user_id}": {
            "get": {
                "description": "get user assets profits",
                "tags": [
                    "user_assets"
                ],
                "summary": "GetUserAssetsProfits",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1userassets.UserAssetsProfits"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user-assets/value/{user_id}": {
            "get": {
                "description": "get user assets values",
                "tags": [
                    "user_assets"
                ],
                "summary": "GetUserAssetsValue",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1userassets.UserAssetsValues"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user-assets/{user_id}": {
            "get": {
                "description": "get user assets for current game",
                "tags": [
                    "user_assets"
                ],
                "summary": "GetUserAssets",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
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
                                "$ref": "#/definitions/models.UserAssets"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/": {
            "post": {
                "description": "Login user",
                "tags": [
                    "users"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1Users.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register user",
                "tags": [
                    "users"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1Users.UserRegistrationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1Users.UserRegistrationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Companies": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "currency_name": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "homepage_url": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "locale": {
                    "type": "string"
                },
                "logo_url": {
                    "type": "string"
                },
                "market": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "primary_exchange": {
                    "type": "string"
                },
                "ticker": {
                    "type": "string"
                },
                "total_employees": {
                    "type": "number"
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CompanyStock": {
            "type": "object",
            "properties": {
                "close_price": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "highest_price": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "lowest_price": {
                    "type": "number"
                },
                "number_of_transactions": {
                    "type": "integer"
                },
                "open_price": {
                    "type": "number"
                },
                "otc": {
                    "type": "boolean"
                },
                "ticker": {
                    "type": "string"
                },
                "trading_volume": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "volume_weighted_average_price": {
                    "type": "number"
                }
            }
        },
        "models.UserAssets": {
            "type": "object",
            "properties": {
                "company_stock_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "price_per_unit": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                },
                "ticker": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.UserTransactions": {
            "type": "object",
            "properties": {
                "buy_price": {
                    "type": "number"
                },
                "company_stock_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "sell_price": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "v1Companies.GetCompanyInfoResponse": {
            "type": "object",
            "properties": {
                "company": {
                    "$ref": "#/definitions/models.Companies"
                },
                "company_stock": {
                    "$ref": "#/definitions/models.CompanyStock"
                }
            }
        },
        "v1Users.UserLoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "v1Users.UserRegistrationRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "v1Users.UserRegistrationResponse": {
            "type": "object",
            "properties": {
                "current_balance": {
                    "type": "number"
                },
                "email": {
                    "type": "string"
                },
                "starting_balance": {
                    "type": "number"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "v1companystock.BuyCompanyStockRequest": {
            "type": "object",
            "properties": {
                "company_stock_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "v1companystock.BuyOrSellCompanyStocksRequest": {
            "type": "object",
            "properties": {
                "buy_or_sell": {
                    "type": "string"
                },
                "company_stock_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "v1companystock.SellCompanyStockRequest": {
            "type": "object",
            "properties": {
                "company_stock_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "v1userassets.UserAssetsProfits": {
            "type": "object",
            "properties": {
                "company_stock": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1userassets.UserAssetsProfitsPerStock"
                    }
                },
                "total_current_value": {
                    "type": "number"
                },
                "total_profit_": {
                    "type": "number"
                },
                "total_profit_margin": {
                    "type": "string"
                },
                "total_spent": {
                    "type": "number"
                }
            }
        },
        "v1userassets.UserAssetsProfitsPerStock": {
            "type": "object",
            "properties": {
                "buy_price": {
                    "type": "number"
                },
                "company_stock": {
                    "$ref": "#/definitions/models.CompanyStock"
                },
                "current_price": {
                    "type": "number"
                },
                "profit_margin": {
                    "type": "string"
                },
                "profit_per_unit": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                },
                "total_buy_price": {
                    "type": "number"
                },
                "total_current_price": {
                    "type": "number"
                },
                "total_profit": {
                    "type": "number"
                }
            }
        },
        "v1userassets.UserAssetsValues": {
            "type": "object",
            "properties": {
                "quantity": {
                    "type": "integer"
                },
                "ticker": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                },
                "value_per_unit": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Capital-Challenge API",
	Description:      "Capital-Challenge backend API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
