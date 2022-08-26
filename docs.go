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
            "url": "https://walletchat.fun",
            "email": "walletchatextension@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/get_inbox/{address}": {
            "get": {
                "description": "Get Each 1-on-1 Conversation, NFT and Community Chat For Display in Inbox",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get Inbox Summary With Last Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet Address",
                        "name": "address",
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
                                "$ref": "#/definitions/entity.Chatiteminbox"
                            }
                        }
                    }
                }
            }
        },
        "/get_unread_cnt/{address}": {
            "get": {
                "description": "Get Each 1-on-1 Conversation, NFT and Community Chat For Display in Inbox",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get all unread messages TO a specific user, used for total count notification at top notification bar",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet Address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/get_unread_cnt/{address}/{nftaddr}/{nftid}": {
            "get": {
                "description": "Get Unread count for specifc NFT context given a wallet address and specific NFT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get all unread messages for a specific NFT context",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet Address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NFT Contract Address",
                        "name": "nftaddr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NFT ID",
                        "name": "nftid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/get_unread_cnt/{fromaddr}/{toaddr}": {
            "get": {
                "description": "Get Unread count for DMs (\u003e_\u003c)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get all unread messages between two addresses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TO: Wallet Address",
                        "name": "toaddr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "FROM: Wallet Address",
                        "name": "from",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/get_unread_cnt_by_type/{address}/{type}": {
            "get": {
                "description": "Get Each 1-on-1 Conversation, NFT and Community Chat For Display in Inbox (\u003e_\u003c)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get all unread messages TO a specific user, used for total count notification at top notification bar",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet Address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Message Type - nft|community|dm|all",
                        "name": "type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/get_unread_cnt_nft/{address}": {
            "get": {
                "description": "Get Unread count for all NFT contexts given a wallet address (\u003e_\u003c)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get all unread messages for all NFT related chats for given user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet Address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/getall_chatitems": {
            "get": {
                "description": "Get Entire Chatitems table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get All Chat Items (legacy - not used currently)",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Chatitem"
                            }
                        }
                    }
                }
            }
        },
        "/getall_chatitems/{address}": {
            "get": {
                "description": "Get all Chat Items for DMs for a given wallet address (\u003e_\u003c)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get Chat Item For Given Wallet Address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet Address",
                        "name": "toaddr",
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
                                "$ref": "#/definitions/entity.Chatitem"
                            }
                        }
                    }
                }
            }
        },
        "/getall_chatitems/{fromaddr}/{toaddr}": {
            "get": {
                "description": "Get chat data between the given two addresses, TO and FROM and interchangable here",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get Chat Data Between Two Addresses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TO: Wallet Address",
                        "name": "toaddr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "FROM: Wallet Address",
                        "name": "from",
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
                                "$ref": "#/definitions/entity.Chatitem"
                            }
                        }
                    }
                }
            }
        },
        "/getnft_chatitems/{address}": {
            "get": {
                "description": "Get ALL NFT context items for a given wallet address (\u003e_\u003c)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get NFT Related Chat Items For Given Wallet Address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet Address",
                        "name": "toaddr",
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
                                "$ref": "#/definitions/entity.Chatitem"
                            }
                        }
                    }
                }
            }
        },
        "/getnft_chatitems/{address}/{nftaddr}/{nftid}": {
            "get": {
                "description": "Get all specified NFT contract and ID items for a given wallet address (\u003e_\u003c)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get NFT Related Chat Items For Given NFT Contract and ID, relating to one wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet Address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NFT Contract Address",
                        "name": "nftaddr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NFT ID",
                        "name": "nftid",
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
                                "$ref": "#/definitions/entity.Chatitem"
                            }
                        }
                    }
                }
            }
        },
        "/getnft_chatitems/{fromaddr}/{toaddr}/{nftaddr}/{nftid}": {
            "get": {
                "description": "Get ALL NFT context items for a specifc NFT context convo between two wallets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get NFT Related Chat Items For Given NFT Contract and ID, between two wallet addresses (TO and FROM are interchangable)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NFT Contract Address",
                        "name": "nftaddr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NFT ID",
                        "name": "nftid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "TO: Wallet Address",
                        "name": "toaddr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "FROM: Wallet Address",
                        "name": "from",
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
                                "$ref": "#/definitions/entity.Chatitem"
                            }
                        }
                    }
                }
            }
        },
        "/getnft_chatitems/{nftaddr}/{nftid}": {
            "get": {
                "description": "Get ALL NFT context items for a given wallet address (\u003e_\u003c)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get NFT Related Chat Items For Given NFT Contract and ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NFT Contract Address",
                        "name": "nftaddr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NFT ID",
                        "name": "nftid",
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
                                "$ref": "#/definitions/entity.Chatitem"
                            }
                        }
                    }
                }
            }
        },
        "/unreadcount/{address}": {
            "get": {
                "description": "Get Unread count just given an address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inbox"
                ],
                "summary": "Get all unread messages TO a specific user, used for total count notification at top notification bar",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet Address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Chatitem": {
            "type": "object",
            "properties": {
                "fromaddr": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "nftaddr": {
                    "type": "string"
                },
                "nftid": {
                    "type": "string"
                },
                "read": {
                    "type": "boolean"
                },
                "sender_name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "timestamp_dtm": {
                    "type": "string"
                },
                "toaddr": {
                    "type": "string"
                }
            }
        },
        "entity.Chatiteminbox": {
            "type": "object",
            "properties": {
                "chain": {
                    "type": "string"
                },
                "context_type": {
                    "type": "string"
                },
                "fromaddr": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logo": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nftaddr": {
                    "type": "string"
                },
                "nftid": {
                    "type": "string"
                },
                "read": {
                    "type": "boolean"
                },
                "sender_name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "timestamp_dtm": {
                    "type": "string"
                },
                "toaddr": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "unread": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "restwalletchat-app-sey3k.ondigitalocean.app:25060",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "WalletChat API",
	Description:      "This is the WalletChat messagez",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
