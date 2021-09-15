package controller

import (
	"context"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/service/wx"
)

//获取全部用户信息
func ListAllUserInfo(ctx context.Context, request model.ListAllUserInfoRequest) (*model.ListAllUserInfoResponse, error) {
	list, err := wx.ListAllUserInfo(ctx)
	return &model.ListAllUserInfoResponse{List: list}, err
}
