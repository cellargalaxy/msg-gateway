package sdk

import (
	"context"
	"fmt"
	common_model "github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var Client *MsgClient

func init() {
	ctx := util.GenCtx()
	var err error
	initConfig(ctx)
	Client, err = NewDefaultMsgClient(ctx)
	if err != nil {
		panic(err)
	}
	if Client == nil {
		panic("创建MsgClient为空")
	}
}

type MsgHandlerInter interface {
	ListAddress(ctx context.Context) []string
	GetSecret(ctx context.Context) string
}

type MsgHandler struct {
}

func (this MsgHandler) ListAddress(ctx context.Context) []string {
	return Config.Addresses
}
func (this MsgHandler) GetSecret(ctx context.Context) string {
	return Config.Secret
}

type MsgClient struct {
	timeout    time.Duration
	retry      int
	httpClient *resty.Client
	handler    MsgHandlerInter
}

func NewDefaultMsgClient(ctx context.Context) (*MsgClient, error) {
	return NewMsgClient(ctx, util.TimeoutDefault, util.RetryDefault, util.GetHttpClient(), &MsgHandler{})
}

func NewMsgClient(ctx context.Context, timeout time.Duration, retry int, httpClient *resty.Client, handler MsgHandlerInter) (*MsgClient, error) {
	if handler == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Error("创建MsgClient，handler为空")
		return nil, fmt.Errorf("MsgHandlerInter为空")
	}
	return &MsgClient{timeout: timeout, retry: retry, handler: handler, httpClient: httpClient}, nil
}

//给配置chatId发送tg信息
func (this *MsgClient) SendTgMsg2ConfigChatId(ctx context.Context, text string) error {
	ctx = util.SetReqId(ctx)
	type Response struct {
		common_model.HttpResponse
		Data model.SendTgMsg2ConfigChatIdRequest `json:"data"`
	}
	var response Response
	err := util.HttpApiRetry(ctx, "给配置chatId发送tg信息", this.retry, []time.Duration{time.Microsecond}, &response, func() (*resty.Response, error) {
		response, err := this.httpClient.R().SetContext(ctx).
			SetHeader(this.genJWT(ctx)).
			SetBody(map[string]interface{}{
				"text": text,
			}).
			Post(this.GetUrl(ctx, model.SendTgMsg2ConfigChatIdPath))
		return response, err
	})
	return err
}

//发送微信模板信息
func (this *MsgClient) SendWxTemplateToTag(ctx context.Context, templateId string, tagId int, url string, data map[string]interface{}) error {
	ctx = util.SetReqId(ctx)
	type Response struct {
		common_model.HttpResponse
		Data model.SendTemplateToTagResponse `json:"data"`
	}
	var response Response
	err := util.HttpApiRetry(ctx, "发送微信模板信息", this.retry, []time.Duration{time.Microsecond}, &response, func() (*resty.Response, error) {
		response, err := this.httpClient.R().SetContext(ctx).
			SetHeader(this.genJWT(ctx)).
			SetBody(map[string]interface{}{
				"template_id": templateId,
				"tag_id":      tagId,
				"url":         url,
				"data":        data,
			}).
			Post(this.GetUrl(ctx, model.SendTemplateToTagPath))
		return response, err
	})
	return err
}

//发送微信通用模板信息
func (this *MsgClient) SendTemplateToCommonTag(ctx context.Context, text string) error {
	ctx = util.SetReqId(ctx)
	type Response struct {
		common_model.HttpResponse
		Data model.SendTemplateToTagResponse `json:"data"`
	}
	var response Response
	err := util.HttpApiRetry(ctx, "发送微信通用模板信息", this.retry, []time.Duration{time.Microsecond}, &response, func() (*resty.Response, error) {
		response, err := this.httpClient.R().SetContext(ctx).
			SetHeader(this.genJWT(ctx)).
			SetBody(map[string]interface{}{
				"text": text,
			}).
			Post(this.GetUrl(ctx, model.SendTemplateToCommonTagPath))
		return response, err
	})
	return err
}

func (this *MsgClient) GetUrl(ctx context.Context, path string) string {
	return this.getUrl(ctx, this.getAddress(ctx), path)
}
func (this *MsgClient) getUrl(ctx context.Context, address, path string) string {
	if strings.HasSuffix(address, "/") && strings.HasPrefix(path, "/") && len(path) > 0 {
		path = path[1:]
	}
	return address + path
}
func (this *MsgClient) getAddress(ctx context.Context) string {
	list := this.handler.ListAddress(ctx)
	if len(list) == 0 {
		return ""
	}
	logId := util.GetLogId(ctx)
	index := int(logId) % len(list)
	return list[index]
}
func (this *MsgClient) genJWT(ctx context.Context) (string, string) {
	return util.GenAuthorizationJWT(ctx, this.timeout, this.handler.GetSecret(ctx))
}
