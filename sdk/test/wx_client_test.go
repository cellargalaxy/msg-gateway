package test

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/cellargalaxy/msg_gateway/sdk"
	"testing"
)

/**
export server_name=msg_gateway
export server_center_address=http://127.0.0.1:7557
export server_center_secret=secret_secret

server_name=msg_gateway;server_center_address=http://127.0.0.1:7557;server_center_secret=secret_secret
*/

func init() {
	util.Init(model.DefaultServerName)
}

func TestSendWxTemplateToTag(test *testing.T) {
	ctx := util.GenCtx()
	err := sdk.Client.SendWxTemplateToTag(ctx, "7ub0o1jXJGfar5Zaj-imwwoisFiH6xW6CsS4pKWjnKc", 109, "", map[string]interface{}{"zhi": 111})
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestSendTgMsg2ConfigChatId(test *testing.T) {
	ctx := util.GenCtx()
	err := sdk.Client.SendTgMsg2ConfigChatId(ctx, `啊啊啊`)
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestSendTemplateToCommonTag(test *testing.T) {
	ctx := util.GenCtx()
	err := sdk.Client.SendTemplateToCommonTag(ctx, `啊啊啊`)
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
