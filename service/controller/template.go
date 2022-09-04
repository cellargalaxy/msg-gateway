package controller

import (
	"context"
	common_model "github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/cellargalaxy/msg_gateway/service/wechat"
)

//获取全部模板
func ListAllTemplate(ctx context.Context, request model.ListAllTemplateRequest) (*model.ListAllTemplateResponse, error) {
	list, err := wechat.ListAllTemplate(ctx)
	return &model.ListAllTemplateResponse{List: list}, err
}

//给标签用户发送模板消息
func SendTemplateToTag(ctx context.Context, request model.SendTemplateToTagRequest) (*model.SendTemplateToTagResponse, error) {
	failOpenIds, err := wechat.SendTemplateToTag(ctx, request.TemplateId, request.TagId, request.Url, request.Data)
	return &model.SendTemplateToTagResponse{FailOpenIds: failOpenIds}, err
}

//给通用标签用户发送模板消息
func SendTemplateToCommonTag(ctx context.Context, claims *common_model.Claims, request model.SendTemplateToCommonTagRequest) (*model.SendTemplateToTagResponse, error) {
	var serverName string
	if claims != nil {
		serverName = claims.ServerName
	}
	failOpenIds, err := wechat.SendTemplateToCommonTag(ctx, serverName, request.Text)
	return &model.SendTemplateToTagResponse{FailOpenIds: failOpenIds}, err
}
