package test

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/service/controller"
	"testing"
)

func TestListAllTemplate(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	request := model.ListAllTemplateRequest{}
	response, err := controller.ListAllTemplate(ctx, request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestSendTemplateToTag(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	request := model.SendTemplateToTagRequest{
		TemplateId: "7ub0o1jXJGfar5Zaj-imwwoisFiH6xW6CsS4pKWjnKc",
		TagId:      108,
		Url:        "https://baidu.com",
		Data:       map[string]interface{}{},
	}
	response, err := controller.SendTemplateToTag(ctx, request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
