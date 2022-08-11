package telegram

import (
	"context"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/config"
	"github.com/go-resty/resty/v2"
)

func SendMsg(ctx context.Context, chatId int64, text string) (bool, error) {
	text = fmt.Sprintf("```\n%+v\n```\nlogid: ```%+v```", text, util.GetLogId(ctx))

	type Response struct {
		TgResponse
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "发送tg信息", util.TryDefault, nil, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetHeader("Content-Type", "application/json;CHARSET=utf-8").
			SetQueryParam("parse_mode", "MarkdownV2").
			SetQueryParam("chat_id", fmt.Sprint(chatId)).
			SetQueryParam("text", text).
			Get(fmt.Sprintf("https://api.telegram.org/bot%+v/sendMessage", config.Config.TgToken))
		return response, err
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
