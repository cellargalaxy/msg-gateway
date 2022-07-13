package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/http"
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
	handler    MsgHandlerInter
	httpClient *resty.Client
}

func NewDefaultMsgClient(ctx context.Context) (*MsgClient, error) {
	return NewMsgClient(ctx, util.TimeoutDefault, util.RetryDefault, &MsgHandler{})
}

func NewMsgClient(ctx context.Context, timeout time.Duration, retry int, handler MsgHandlerInter) (*MsgClient, error) {
	if handler == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Error("创建MsgClient，handler为空")
		return nil, fmt.Errorf("MsgHandlerInter为空")
	}
	httpClient := util.GetHttpClient()
	return &MsgClient{timeout: timeout, retry: retry, handler: handler, httpClient: httpClient}, nil
}

//给配置chatId发送tg信息
func (this *MsgClient) SendTgMsg2ConfigChatId(ctx context.Context, text string) (bool, error) {
	ctx = util.SetReqId(ctx)
	var jsonString string
	var object bool
	var err error
	for i := 0; i < this.retry; i++ {
		jsonString, err = this.requestSendTgMsg2ConfigChatId(ctx, text)
		if err == nil {
			object, err = this.analysisSendTgMsg2ConfigChatId(ctx, jsonString)
			if err == nil {
				return object, err
			}
		}
	}
	return object, err
}
func (this *MsgClient) analysisSendTgMsg2ConfigChatId(ctx context.Context, jsonString string) (bool, error) {
	type Response struct {
		Code int                                 `json:"code"`
		Msg  string                              `json:"msg"`
		Data model.SendTgMsg2ConfigChatIdRequest `json:"data"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("给配置chatId发送tg信息，解析响应异常")
		return false, fmt.Errorf("给配置chatId发送tg信息，解析响应异常")
	}
	if response.Code != util.HttpSuccessCode {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"jsonString": jsonString}).Error("给配置chatId发送tg信息，失败")
		return false, fmt.Errorf("给配置chatId发送tg信息，失败: %+v", jsonString)
	}
	return true, nil
}
func (this *MsgClient) requestSendTgMsg2ConfigChatId(ctx context.Context, text string) (string, error) {
	response, err := this.httpClient.R().SetContext(ctx).
		SetHeader(this.genJWT(ctx)).
		SetBody(map[string]interface{}{
			"text": text,
		}).
		Post(this.GetUrl(ctx, model.SendTgMsg2ConfigChatIdPath))

	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("给配置chatId发送tg信息，请求异常")
		return "", fmt.Errorf("给配置chatId发送tg信息，请求异常")
	}
	if response == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("给配置chatId发送tg信息，响应为空")
		return "", fmt.Errorf("给配置chatId发送tg信息，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithContext(ctx).WithFields(logrus.Fields{"statusCode": statusCode, "body": body}).Info("给配置chatId发送tg信息，响应")
	if statusCode != http.StatusOK {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"StatusCode": statusCode}).Error("给配置chatId发送tg信息，响应码失败")
		return "", fmt.Errorf("给配置chatId发送tg信息，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//发送微信模板信息
func (this *MsgClient) SendWxTemplateToTag(ctx context.Context, templateId string, tagId int, url string, data map[string]interface{}) (bool, error) {
	ctx = util.SetReqId(ctx)
	var jsonString string
	var object bool
	var err error
	for i := 0; i < this.retry; i++ {
		jsonString, err = this.requestSendWxTemplateToTag(ctx, templateId, tagId, url, data)
		if err == nil {
			object, err = this.analysisSendWxTemplateToTag(ctx, jsonString)
			if err == nil {
				return object, err
			}
		}
	}
	return object, err
}
func (this *MsgClient) analysisSendWxTemplateToTag(ctx context.Context, jsonString string) (bool, error) {
	type Response struct {
		Code int                             `json:"code"`
		Msg  string                          `json:"msg"`
		Data model.SendTemplateToTagResponse `json:"data"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("发送微信模板信息，解析响应异常")
		return false, fmt.Errorf("发送微信模板信息，解析响应异常")
	}
	if response.Code != util.HttpSuccessCode {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"jsonString": jsonString}).Error("发送微信模板信息，失败")
		return false, fmt.Errorf("发送微信模板信息，失败: %+v", jsonString)
	}
	return true, nil
}
func (this *MsgClient) requestSendWxTemplateToTag(ctx context.Context, templateId string, tagId int, url string, data map[string]interface{}) (string, error) {
	response, err := this.httpClient.R().SetContext(ctx).
		SetHeader(this.genJWT(ctx)).
		SetBody(map[string]interface{}{
			"template_id": templateId,
			"tag_id":      tagId,
			"url":         url,
			"data":        data,
		}).
		Post(this.GetUrl(ctx, model.SendTemplateToTagPath))

	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("发送微信模板信息，请求异常")
		return "", fmt.Errorf("发送微信模板信息，请求异常")
	}
	if response == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("发送微信模板信息，响应为空")
		return "", fmt.Errorf("发送微信模板信息，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithContext(ctx).WithFields(logrus.Fields{"statusCode": statusCode, "body": body}).Info("发送微信模板信息，响应")
	if statusCode != http.StatusOK {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"StatusCode": statusCode}).Error("发送微信模板信息，响应码失败")
		return "", fmt.Errorf("发送微信模板信息，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//发送微信通用模板信息
func (this *MsgClient) SendTemplateToCommonTag(ctx context.Context, text string) (bool, error) {
	ctx = util.SetReqId(ctx)
	var jsonString string
	var object bool
	var err error
	for i := 0; i < this.retry; i++ {
		jsonString, err = this.requestSendTemplateToCommonTag(ctx, text)
		if err == nil {
			object, err = this.analysisSendTemplateToCommonTag(ctx, jsonString)
			if err == nil {
				return object, err
			}
		}
	}
	return object, err
}
func (this *MsgClient) analysisSendTemplateToCommonTag(ctx context.Context, jsonString string) (bool, error) {
	return this.analysisSendWxTemplateToTag(ctx, jsonString)
}
func (this *MsgClient) requestSendTemplateToCommonTag(ctx context.Context, text string) (string, error) {
	response, err := this.httpClient.R().SetContext(ctx).
		SetHeader(this.genJWT(ctx)).
		SetBody(map[string]interface{}{
			"text": text,
		}).
		Post(this.GetUrl(ctx, model.SendTemplateToCommonTagPath))

	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("发送微信通用模板信息，请求异常")
		return "", fmt.Errorf("发送微信通用模板信息，请求异常")
	}
	if response == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("发送微信通用模板信息，响应为空")
		return "", fmt.Errorf("发送微信通用模板信息，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithContext(ctx).WithFields(logrus.Fields{"statusCode": statusCode, "body": body}).Info("发送微信通用模板信息，响应")
	if statusCode != http.StatusOK {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"StatusCode": statusCode}).Error("发送微信通用模板信息，响应码失败")
		return "", fmt.Errorf("发送微信通用模板信息，响应码失败: %+v", statusCode)
	}
	return body, nil
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
