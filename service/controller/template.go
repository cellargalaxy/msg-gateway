package controller

import (
	"context"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/service/wx"
)

//获取全部模板
func ListAllTemplate(ctx context.Context, request model.ListAllTemplateRequest) (*model.ListAllTemplateResponse, error) {
	list, err := wx.ListAllTemplate(ctx)
	return &model.ListAllTemplateResponse{List: list}, err
}

//给标签用户发送模板消息
func SendTemplateToTag(ctx context.Context, request model.SendTemplateToTagRequest) (*model.SendTemplateToTagResponse, error) {
	failOpenIds, err := wx.SendTemplateToTag(ctx, request.TemplateId, request.TagId, request.Url, request.Data)
	return &model.SendTemplateToTagResponse{FailOpenIds: failOpenIds}, err
}

//给通用标签用户发送模板消息
func SendTemplateToCommonTag(ctx context.Context, request model.SendTemplateToCommonTagRequest) (*model.SendTemplateToTagResponse, error) {
	failOpenIds, err := wx.SendTemplateToCommonTag(ctx, request.Text)
	return &model.SendTemplateToTagResponse{FailOpenIds: failOpenIds}, err
}
