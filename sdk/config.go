package sdk

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

var Config = model.ClientConfig{}

func InitConfig(handler sdk.ServerCenterHandlerInter) {
	ctx := util.CreateLogCtx()
	if handler == nil {
		handler = &ServerCenterHandler{}
	}
	client, err := sdk.NewDefaultServerCenterClient(ctx, handler)
	if err != nil {
		panic(err)
	}
	client.StartConfWithInitConf(ctx)
}

func checkAndResetConfig(ctx context.Context, config model.ClientConfig) (model.ClientConfig, error) {
	if config.Timeout < 0 {
		config.Timeout = 3 * time.Second
	}
	if config.Sleep < 0 {
		config.Sleep = 3 * time.Second
	}
	if config.Address == "" {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Error("address为空")
		return config, fmt.Errorf("address为空")
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
func (this *ServerCenterHandler) GetServerName(ctx context.Context) string {
	return "msg_gateway_sdk"
}
func (this *ServerCenterHandler) GetInterval(ctx context.Context) time.Duration {
	return 5 * time.Minute
}
func (this *ServerCenterHandler) ParseConf(ctx context.Context, object sc_model.ServerConfModel) error {
	var config model.ClientConfig
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
	logrus.WithContext(ctx).WithFields(logrus.Fields{"Config": Config}).Info("加载配置")
	return nil
}
func (this *ServerCenterHandler) GetDefaultConf(ctx context.Context) string {
	var config model.ClientConfig
	config, _ = checkAndResetConfig(ctx, config)
	return util.ToYamlString(config)
}
func (this *ServerCenterHandler) GetLocalFilePath(ctx context.Context) string {
	return ""
}
