package config

import (
	"context"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	sc_model "github.com/cellargalaxy/server_center/model"
	"github.com/cellargalaxy/server_center/sdk"
	"github.com/sirupsen/logrus"
	"time"
)

var Config = model.Config{}

func init() {
	ctx := context.Background()
	client, err := sdk.NewDefaultServerCenterClient(&ServerCenterHandler{})
	if err != nil {
		panic(err)
	}
	_, err = client.StartConfWithInitConf(ctx)
	if err != nil {
		panic(err)
	}
}

func checkAndResetConfig(ctx context.Context, config model.Config) (model.Config, error) {
	if config.LogLevel <= 0 || config.LogLevel > logrus.TraceLevel {
		config.LogLevel = logrus.InfoLevel
	}
	if Config.Timeout < 0 {
		Config.Timeout = 3 * time.Second
	}
	if Config.Sleep < 0 {
		Config.Sleep = 3 * time.Second
	}
	if config.Secret == "" {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Error("secret为空")
		return config, fmt.Errorf("secret为空")
	}
	return config, nil
}

type ServerCenterHandler struct {
}

func (this *ServerCenterHandler) GetAddress(ctx context.Context) string {
	return sdk.GetEnvServerCenterAddress(ctx)
}
func (this *ServerCenterHandler) GetSecret(ctx context.Context) string {
	return sdk.GetEnvServerCenterSecret(ctx)
}
func (this *ServerCenterHandler) GetInterval(ctx context.Context) time.Duration {
	return 5 * time.Minute
}
func (this *ServerCenterHandler) ParseConf(ctx context.Context, object sc_model.ServerConfModel) error {
	var config model.Config
	err := util.UnmarshalYamlString(object.ConfText, &config)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("反序列化配置异常")
		return err
	}
	config, err = checkAndResetConfig(ctx, config)
	if err != nil {
		return err
	}
	Config = config
	logrus.SetLevel(Config.LogLevel)
	logrus.WithContext(ctx).WithFields(logrus.Fields{"Config": Config}).Info("加载配置")
	return nil
}
