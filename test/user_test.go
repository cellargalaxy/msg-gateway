package test

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/service/controller"
	"testing"
)

func TestListAllUserInfo(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	request := model.ListAllUserInfoRequest{}
	response, err := controller.ListAllUserInfo(ctx, request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
