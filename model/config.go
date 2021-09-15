package model

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	ListenAddress = ":8990"
)

type Config struct {
	LogLevel       logrus.Level  `ini:"log_level" json:"log_level"`
	Retry          int           `ini:"retry" json:"retry"`
	Timeout        time.Duration `ini:"timeout" json:"timeout"`
	Sleep          time.Duration `ini:"sleep" json:"sleep"`
	Secret         string        `ini:"secret" json:"secret"`
	WxAppId        string        `ini:"wx_app_id" json:"wx_app_id"`
	WxAppSecret    string        `ini:"wx_app_secret" json:"wx_app_secret"`
	WxCommonTempId string        `ini:"wx_common_temp_id" json:"wx_common_temp_id"`
	WxCommonTagId  int           `ini:"wx_common_tag_id" json:"wx_common_tag_id"`
	TgToken        string        `ini:"tg_token" json:"tg_token"`
	TgChatId       int64         `ini:"tg_chat_id" json:"tg_chat_id"`
}

func (this Config) String() string {
	return util.ToJsonString(this)
}
