// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2021-07-16 10:10:17.73764308 +0800 CST m=+0.133036334

package docs

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
        "/api/v1/login": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "cmd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "channel",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "open_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/login.loginResponseInfo"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/login.err"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/login.err"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "login.err": {
            "$ref": "#/definitions/login.error"
        },
        "login.loginRequestInfo": {
            "type": "object",
            "properties": {
                "cmd": {
                    "type": "integer"
                },
                "params": {
                    "type": "object",
                    "$ref": "#/definitions/login.params"
                }
            }
        },
        "login.loginResponseInfo": {
            "type": "object",
            "properties": {
                "money": {
                    "type": "integer"
                },
                "open_id": {
                    "type": "string"
                },
                "sign": {
                    "type": "integer"
                }
            }
        },
        "login.params": {
            "type": "object",
            "properties": {
                "channel": {
                    "type": "string"
                },
                "open_id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "login.wxloginRequestInfo": {
            "type": "object",
            "properties": {
                "cmd": {
                    "type": "integer"
                },
                "params": {
                    "type": "object",
                    "$ref": "#/definitions/login.params"
                }
            }
        },
        "login.wxloginResponseInfo": {
            "type": "object",
            "properties": {
                "money": {
                    "type": "integer"
                },
                "open_id": {
                    "type": "string"
                },
                "sign": {
                    "type": "integer"
                }
            }
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "突突大挑战",
	Description: "Golang gin",
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
