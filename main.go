package main

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/controller"
	"github.com/cellargalaxy/msg_gateway/model"
)

func init() {
	util.Init(model.DefaultServerName)
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
