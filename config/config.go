package config

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	configFilePath = "resources/config.ini"
)

var Config = model.Config{}

func init() {
	ctx := context.Background()
	exist, _ := util.ExistAndIsFile(ctx, configFilePath)
	if exist {
		cfg, err := ini.Load(configFilePath)
		if err != nil {
			panic(err)
		}
		err = cfg.MapTo(&Config)
		if err != nil {
			panic(err)
		}
	}

	if Config.LogLevel <= 0 || Config.LogLevel > logrus.TraceLevel {
		Config.LogLevel = logrus.InfoLevel
	}
	if Config.Timeout < 0 {
		Config.Timeout = 3 * time.Second
	}
	if Config.Sleep < 0 {
		Config.Sleep = 3 * time.Second
	}
	if Config.Secret == "" {
		panic("Secret配置为空")
	}
}
