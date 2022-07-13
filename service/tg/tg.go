package tg

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/config"
	"github.com/go-resty/resty/v2"
)

var httpClient *resty.Client

func init() {
	httpClient = util.GetHttpClient()
}

//https://www.cnblogs.com/kainhuck/p/13576012.html

//给配置chatId发送tg信息
func SendTgMsg2ConfigChatId(ctx context.Context, text string) (bool, error) {
	return SendMsg(ctx, config.Config.TgChatId, text)
}
