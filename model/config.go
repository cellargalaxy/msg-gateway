package model

import (
	"github.com/cellargalaxy/go_common/util"
	"time"
)

const (
	DefaultServerName = "msg_gateway"
	ListenAddress     = ":8990"
)

type Config struct {
	Retry          int           `yaml:"retry" json:"retry"`
	Timeout        time.Duration `yaml:"timeout" json:"timeout"`
	Sleep          time.Duration `yaml:"sleep" json:"sleep"`
	Secret         string        `yaml:"secret" json:"-"`
	WxAppId        string        `yaml:"wx_app_id" json:"wx_app_id"`
	WxAppSecret    string        `yaml:"wx_app_secret" json:"-"`
	WxCommonTempId string        `yaml:"wx_common_temp_id" json:"wx_common_temp_id"`
	WxCommonTagId  int           `yaml:"wx_common_tag_id" json:"wx_common_tag_id"`
	TgToken        string        `yaml:"tg_token" json:"-"`
	TgChatId       int64         `yaml:"tg_chat_id" json:"tg_chat_id"`
}

func (this Config) String() string {
	return util.ToJsonString(this)
}
