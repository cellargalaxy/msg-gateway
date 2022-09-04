package telegram

import (
	"context"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/config"
	"github.com/go-resty/resty/v2"
)

var httpClient *resty.Client

func init() {
	httpClient = util.GetHttpClient()
}

type TgResponse struct {
	Ok bool `json:"ok"`
}

func (this *TgResponse) String() string {
	return util.ToJsonString(this)
}
func (this *TgResponse) HttpSuccess(ctx context.Context) error {
	if this.Ok {
		return nil
	}
	return fmt.Errorf("TG响应失败: %+v", this)
}

//https://www.cnblogs.com/kainhuck/p/13576012.html

//给配置chatId发送tg信息
func SendTgMsg2ConfigChatId(ctx context.Context, serverName, text string) (bool, error) {
	return SendMsg(ctx, config.Config.TgChatId, serverName, text)
}
