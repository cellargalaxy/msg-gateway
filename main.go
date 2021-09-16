package main

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/config"
	"github.com/cellargalaxy/msg_gateway/controller"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(config.Config.LogLevel)
	util.InitLog("msg_gateway")
}

func main() {
	err := controller.Controller()
	if err != nil {
		panic(err)
	}
}
