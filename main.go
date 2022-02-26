package main

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/config"
	"github.com/cellargalaxy/msg_gateway/controller"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(config.Config.LogLevel)
	util.InitDefaultLog(model.DefaultServerName)
}

/**
export server_name=msg_gateway
export server_center_address=http://127.0.0.1:7557
export server_center_secret=secret_secret

server_name=msg_gateway;server_center_address=http://127.0.0.1:7557;server_center_secret=secret_secret
*/
func main() {
	err := controller.Controller()
	if err != nil {
		panic(err)
	}
}
