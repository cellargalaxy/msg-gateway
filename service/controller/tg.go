package controller

import (
	"context"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/cellargalaxy/msg_gateway/service/telegram"
)

//给配置chatId发送tg信息
func SendTgMsg2ConfigChatId(ctx context.Context, request model.SendTgMsg2ConfigChatIdRequest) (*model.SendTgMsg2ConfigChatIdResponse, error) {
	result, err := telegram.SendTgMsg2ConfigChatId(ctx, request.Text)
	return &model.SendTgMsg2ConfigChatIdResponse{Result: result}, err
}
