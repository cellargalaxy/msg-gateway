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

var config = model.ClientConfig{}

func InitConfig() {
	ctx := util.CreateLogCtx()
	client, err := sdk.NewDefaultServerCenterClient(&ServerCenterHandler{})
	if err != nil {
		panic(err)
	}
	_, err = client.StartConfWithInitConf(ctx)
	if err != nil {
		panic(err)
	}
}

func checkAndResetConfig(ctx context.Context, config model.ClientConfig) (model.ClientConfig, error) {
	if config.Timeout < 0 {
		config.Timeout = 3 * time.Second
	}
	if config.Sleep < 0 {
		config.Sleep = 3 * time.Second
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
	return sdk.GetEnvServerName(ctx) + "_sdk"
}
func (this *ServerCenterHandler) GetInterval(ctx context.Context) time.Duration {
	return 5 * time.Minute
}
func (this *ServerCenterHandler) ParseConf(ctx context.Context, object sc_model.ServerConfModel) error {
	var conf model.ClientConfig
	err := util.UnmarshalYamlString(object.ConfText, &config)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("反序列化配置异常")
		return err
	}
	config, err = checkAndResetConfig(ctx, conf)
	if err != nil {
		return err
	}
	config = conf
	logrus.WithContext(ctx).WithFields(logrus.Fields{"config": config}).Info("加载配置")
	return nil
}
