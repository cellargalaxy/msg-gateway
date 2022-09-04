package controller

import (
	"context"
	common_model "github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/cellargalaxy/msg_gateway/service/telegram"
)

//给配置chatId发送tg信息
func SendTgMsg2ConfigChatId(ctx context.Context, claims *common_model.Claims, request model.SendTgMsg2ConfigChatIdRequest) (*model.SendTgMsg2ConfigChatIdResponse, error) {
	var serverName string
	if claims != nil {
		serverName = claims.ServerName
	}
	result, err := telegram.SendTgMsg2ConfigChatId(ctx, serverName, request.Text)
	return &model.SendTgMsg2ConfigChatIdResponse{Result: result}, err
}
