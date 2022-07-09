package model

import (
	"github.com/cellargalaxy/go_common/util"
)

const (
	DefaultServerName           = "msg_gateway"
	ListAllTemplatePath         = "/api/listAllTemplate"
	SendTemplateToTagPath       = "/api/sendTemplateToTag"
	SendTemplateToCommonTagPath = "/api/sendTemplateToCommonTag"
	CreateTagPath               = "/api/createTag"
	DeleteTagPath               = "/api/deleteTag"
	ListAllTagPath              = "/api/listAllTag"
	AddTagToUserPath            = "/api/addTagToUser"
	DeleteTagFromUserPath       = "/api/deleteTagFromUser"
	ListAllUserInfoPath         = "/api/listAllUserInfo"
	SendTgMsg2ConfigChatIdPath  = "/api/sendTgMsg2ConfigChatId"
	ListenAddress               = ":8990"
)

type Config struct {
	Retry     int      `yaml:"retry" json:"retry"`
	Addresses []string `yaml:"addresses" json:"addresses"`
	Secret    string   `yaml:"secret" json:"-"`

	WxAppId        string `yaml:"wx_app_id" json:"wx_app_id"`
	WxAppSecret    string `yaml:"wx_app_secret" json:"-"`
	WxCommonTempId string `yaml:"wx_common_temp_id" json:"wx_common_temp_id"`
	WxCommonTagId  int    `yaml:"wx_common_tag_id" json:"wx_common_tag_id"`

	TgToken  string `yaml:"tg_token" json:"-"`
	TgChatId int64  `yaml:"tg_chat_id" json:"tg_chat_id"`
}

func (this Config) String() string {
	return util.ToJsonString(this)
}
