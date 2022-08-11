package controller

import (
	"context"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/cellargalaxy/msg_gateway/service/wechat"
)

//获取全部用户信息
func ListAllUserInfo(ctx context.Context, request model.ListAllUserInfoRequest) (*model.ListAllUserInfoResponse, error) {
	list, err := wechat.ListAllUserInfo(ctx)
	return &model.ListAllUserInfoResponse{List: list}, err
}
