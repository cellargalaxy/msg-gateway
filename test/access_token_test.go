package test

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/service/wx"
	"testing"
)

func TestGetAccessToken(test *testing.T) {
	ctx := util.GenCtx()
	accessToken := wx.GetAccessToken(ctx)
	test.Logf("accessToken: %+v\r\n", accessToken)
}
