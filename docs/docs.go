// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/club/introduction": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改社团简介",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌 例:Bearer fbhraewifvg43uwerfaewobf",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "社团简介",
                        "name": "introduction",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/club.Response"
                        }
                    }
                }
            }
        },
        "/club/post": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "创建活动或动态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌 例:Bearer fbhraewifvg43uwerfaewobf",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "minLength": 2,
                        "type": "string",
                        "description": "标题,min=2",
                        "name": "article",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 10,
                        "type": "string",
                        "description": "内容,min=10",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "是否为通知",
                        "name": "is_notice",
                        "in": "formData"
                    },
                    {
                        "maxItems": 9,
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "description": "包涵要上传的帖子图片的id的数组,最多9张图",
                        "name": "picture_ids",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/club.CreateResponse"
                        }
                    }
                }
            }
        },
        "/club/signed_url": {
            "get": {
                "description": "获取上传图片的url及回调字符串,图片Id",
                "produces": [
                    "application/json"
                ],
                "summary": "获取上传图片或头像的签名",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌 例:Bearer fbhraewifvg43uwerfaewobf",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "要发送的图片名",
                        "name": "picture_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "\"post_picture\"或\"avatar\"",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/club.ResUrl"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "登录社团账号",
                "parameters": [
                    {
                        "minLength": 2,
                        "type": "string",
                        "description": "社团名或社团账号,至少2位",
                        "name": "club_name_or_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 6,
                        "type": "string",
                        "description": "密码,至少6位",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/club.ResLogin"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "注册社团账号",
                "parameters": [
                    {
                        "maximum": 999999999999,
                        "minimum": 99999,
                        "type": "integer",
                        "description": "社团登录账号7-12位",
                        "name": "club_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 2,
                        "type": "string",
                        "description": "社团名称",
                        "name": "club_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "邀请码",
                        "name": "invitation_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 32,
                        "minLength": 6,
                        "type": "string",
                        "description": "密码6-32位",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "password2",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/club.RspRegister"
                        }
                    }
                }
            }
        },
        "/user/club_info": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "查看社团信息,可用于展示社团的页面",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "社团id",
                        "name": "club_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.ResponseClubInfo"
                        }
                    }
                }
            }
        },
        "/user/post": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取单条通知/动态(详情)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "帖子的id,min=1",
                        "name": "post_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.ResponsePost"
                        }
                    }
                }
            }
        },
        "/user/posts": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取最新的多条通知/动态",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "\"要查询的是否为通知,是则为true,否则为false\"",
                        "name": "is_notice",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "description": "页码,最小为1",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maximum": 40,
                        "minimum": 1,
                        "type": "integer",
                        "example": 10,
                        "description": "每页数据量,最大为40",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data内有多条post",
                        "schema": {
                            "$ref": "#/definitions/user.ResponsePosts"
                        }
                    }
                }
            }
        },
        "/user/posts_from_clubs_user_fellow": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取用户关注的社团发布的最新的多条通知和动态",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "description": "页码,最小为1",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maximum": 40,
                        "minimum": 1,
                        "type": "integer",
                        "example": 10,
                        "description": "每页数据量,最大为40",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "student_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data内有多条post",
                        "schema": {
                            "$ref": "#/definitions/user.ResponsePosts"
                        }
                    }
                }
            }
        },
        "/user/subscribe": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "关注社团",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "社团Id",
                        "name": "club_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "学号",
                        "name": "student_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.ResponseSubscribe"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "取消关注社团",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "社团Id",
                        "name": "club_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "学号",
                        "name": "student_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.ResponseSubscribe"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "club.CreateResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "description": "返回的信息",
                    "type": "string"
                }
            }
        },
        "club.ResLogin": {
            "type": "object",
            "properties": {
                "msg": {
                    "description": "信息",
                    "type": "string"
                },
                "token": {
                    "description": "token",
                    "type": "string"
                }
            }
        },
        "club.ResUrl": {
            "type": "object",
            "properties": {
                "callback_str": {
                    "description": "回调的字符串",
                    "type": "string"
                },
                "url": {
                    "description": "签名url",
                    "type": "string"
                }
            }
        },
        "club.Response": {
            "type": "object",
            "properties": {
                "msg": {
                    "description": "返回的信息",
                    "type": "string"
                }
            }
        },
        "club.RspRegister": {
            "type": "object",
            "properties": {
                "club_id": {
                    "description": "社团Id",
                    "type": "integer"
                },
                "club_name": {
                    "type": "string"
                },
                "msg": {
                    "description": "信息",
                    "type": "string"
                },
                "password": {
                    "description": "用户密码",
                    "type": "string"
                }
            }
        },
        "user.ResponseClubInfo": {
            "type": "object",
            "properties": {
                "avatar_addr": {
                    "description": "社团头像url地址",
                    "type": "string"
                },
                "club_name": {
                    "description": "社团名称",
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "description": "社团简介",
                    "type": "string"
                },
                "num_of_fans": {
                    "description": "粉丝数",
                    "type": "integer"
                },
                "num_of_favorites": {
                    "description": "活动被收藏的总次数",
                    "type": "integer"
                }
            }
        },
        "user.ResponsePost": {
            "type": "object",
            "properties": {
                "article": {
                    "description": "标题",
                    "type": "string"
                },
                "avatar_addr": {
                    "description": "头像url地址",
                    "type": "string"
                },
                "club_id": {
                    "description": "社团Id",
                    "type": "integer"
                },
                "club_name": {
                    "description": "社团名称",
                    "type": "string"
                },
                "content": {
                    "description": "内容",
                    "type": "string"
                },
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "Msg         string ` + "`" + `json:\"msg\"` + "`" + `   //信息 如\"获取成功\"",
                    "type": "integer"
                },
                "is_notice": {
                    "description": "是否为通知",
                    "type": "boolean"
                },
                "picture_addr": {
                    "description": "帖子图片的多个可访问地址",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updated_at": {
                    "description": "更新时间(刚创建时为空)",
                    "type": "string"
                }
            }
        },
        "user.ResponsePosts": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "data内包涵多条帖子数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.ResponsePost"
                    }
                }
            }
        },
        "user.ResponseSubscribe": {
            "type": "object",
            "properties": {
                "msg": {
                    "description": "信息",
                    "type": "string"
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
