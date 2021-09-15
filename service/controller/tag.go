package controller

import (
	"context"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/service/wx"
)

//创建标签
func CreateTag(ctx context.Context, request model.CreateTagRequest) (*model.CreateTagResponse, error) {
	result, err := wx.CreateTag(ctx, request.Tag)
	return &model.CreateTagResponse{Result: result}, err
}

//删除标签
func DeleteTag(ctx context.Context, request model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	result, err := wx.DeleteTag(ctx, request.TagId)
	return &model.DeleteTagResponse{Result: result}, err
}

//获取所有标签
func ListAllTag(ctx context.Context, request model.ListAllTagRequest) (*model.ListAllTagResponse, error) {
	list, err := wx.ListAllTag(ctx)
	return &model.ListAllTagResponse{List: list}, err
}

//为用户加标签
func AddTagToUser(ctx context.Context, request model.AddTagToUserRequest) (*model.AddTagToUserResponse, error) {
	result, err := wx.AddTagToUser(ctx, request.TagId, []string{request.OpenId})
	return &model.AddTagToUserResponse{Result: result}, err
}

//为用户删标签
func DeleteTagFromUser(ctx context.Context, request model.DeleteTagFromUserRequest) (*model.DeleteTagFromUserResponse, error) {
	result, err := wx.DeleteTagFromUser(ctx, request.TagId, []string{request.OpenId})
	return &model.DeleteTagFromUserResponse{Result: result}, err
}
