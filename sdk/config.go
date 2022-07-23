package sdk

import (
	"context"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	sc_model "github.com/cellargalaxy/server_center/model"
	"github.com/cellargalaxy/server_center/sdk"
	"github.com/sirupsen/logrus"
)

var Config = model.Config{}

func initConfig(ctx context.Context) {
	var err error

	var handler ServerCenterHandler
	client, err := sdk.NewDefaultServerCenterClient(ctx, &handler)
	if err != nil {
		panic(err)
	}
	if client == nil {
		panic("创建ServerCenterClient为空")
	}
	client.StartWithInitConf(ctx)
	logrus.WithContext(ctx).WithFields(logrus.Fields{"Config": Config}).Info("加载配置")
}

type ServerCenterHandler struct {
	sdk.ServerCenterDefaultHandler
}

func (this *ServerCenterHandler) GetServerName(ctx context.Context) string {
	return model.DefaultServerName
}
func (this *ServerCenterHandler) ParseConf(ctx context.Context, object sc_model.ServerConfModel) error {
	var config model.Config
	err := util.UnmarshalYamlString(object.ConfText, &config)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("反序列化配置异常")
		return err
	}
	if len(config.Addresses) == 0 {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("反序列化配置，Addresses为空")
		return fmt.Errorf("反序列化配置，Addresses为空")
	}
	Config = config
	logrus.WithContext(ctx).WithFields(logrus.Fields{"Config": Config}).Info("加载配置")
	return nil
}
func (this *ServerCenterHandler) GetDefaultConf(ctx context.Context) string {
	var config model.Config
	return util.ToYamlString(config)
}
