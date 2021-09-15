package test

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/service/controller"
	"testing"
)

func TestCreateTag(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	request := model.CreateTagRequest{
		Tag: "",
	}
	response, err := controller.CreateTag(ctx, request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestDeleteTag(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	request := model.DeleteTagRequest{
		TagId: 0,
	}
	response, err := controller.DeleteTag(ctx, request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestListAllTag(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	request := model.ListAllTagRequest{}
	response, err := controller.ListAllTag(ctx, request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestAddTagToUser(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	request := model.AddTagToUserRequest{
		TagId:  0,
		OpenId: "",
	}
	response, err := controller.AddTagToUser(ctx, request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestDeleteTagFromUser(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	request := model.DeleteTagFromUserRequest{
		TagId:  0,
		OpenId: "",
	}
	response, err := controller.DeleteTagFromUser(ctx, request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
