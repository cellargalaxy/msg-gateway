package sdk

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/go_common/consd"
	common_model "github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type MsgClient struct {
	address    string
	secret     string
	retry      int
	httpClient *resty.Client
}

func NewMsgClient(timeout, sleep time.Duration, retry int, address, secret string) (*MsgClient, error) {
	if address == "" {
		return nil, fmt.Errorf("address为空")
	}
	if secret == "" {
		return nil, fmt.Errorf("secret为空")
	}
	httpClient := createHttpClient(timeout, sleep, retry)
	return &MsgClient{address: address, secret: secret, retry: retry, httpClient: httpClient}, nil
}

func createHttpClient(timeout, sleep time.Duration, retry int) *resty.Client {
	httpClient := resty.New().
		SetTimeout(timeout).
		SetRetryCount(retry).
		SetRetryWaitTime(sleep).
		SetRetryMaxWaitTime(5 * time.Minute).
		AddRetryCondition(func(response *resty.Response, err error) bool {
			ctx := util.CreateLogCtx()
			if response != nil && response.Request != nil {
				ctx = response.Request.Context()
			}
			var statusCode int
			if response != nil {
				statusCode = response.StatusCode()
			}
			retry := statusCode != http.StatusOK || err != nil
			if retry {
				logrus.WithContext(ctx).WithFields(logrus.Fields{"statusCode": statusCode, "err": err}).Warn("HTTP请求异常，进行重试")
			}
			return retry
		}).
		SetRetryAfter(func(client *resty.Client, response *resty.Response) (time.Duration, error) {
			ctx := util.CreateLogCtx()
			if response != nil && response.Request != nil {
				ctx = response.Request.Context()
			}
			var attempt int
			if response != nil && response.Request != nil {
				attempt = response.Request.Attempt
			}
			if attempt > retry {
				logrus.WithContext(ctx).WithFields(logrus.Fields{"attempt": attempt}).Error("HTTP请求异常，超过最大重试次数")
				return 0, fmt.Errorf("HTTP请求异常，超过最大重试次数")
			}
			duration := util.WareDuration(sleep)
			for i := 0; i < attempt-1; i++ {
				duration *= 10
			}
			logrus.WithContext(ctx).WithFields(logrus.Fields{"attempt": attempt, "duration": duration}).Warn("HTTP请求异常，休眠重试")
			return duration, nil
		}).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return httpClient
}

//给配置chatId发送tg信息
func (this MsgClient) SendTgMsg2ConfigChatId(ctx context.Context, text string) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < this.retry; i++ {
		jwtToken, err := this.genJWT(ctx)
		if err != nil {
			return false, err
		}
		jsonString, err = this.requestSendTgMsg2ConfigChatId(ctx, jwtToken, text)
		if err == nil {
			object, err = this.analysisSendTgMsg2ConfigChatId(ctx, jsonString)
			if err == nil {
				return object, err
			}
		}
	}
	return object, err
}

//给配置chatId发送tg信息
func (this MsgClient) analysisSendTgMsg2ConfigChatId(ctx context.Context, jsonString string) (bool, error) {
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
	if response.Code != consd.HttpSuccessCode {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"jsonString": jsonString}).Error("给配置chatId发送tg信息，失败")
		return false, fmt.Errorf("给配置chatId发送tg信息，失败: %+v", jsonString)
	}
	return true, nil
}

//给配置chatId发送tg信息
func (this MsgClient) requestSendTgMsg2ConfigChatId(ctx context.Context, jwtToken string, text string) (string, error) {
	response, err := this.httpClient.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetHeader("Authorization", "Bearer "+jwtToken).
		SetHeader(util.LogIdKey, fmt.Sprint(util.GetLogId(ctx))).
		SetBody(map[string]interface{}{
			"text": text,
		}).
		Post(this.address + "/api/sendTgMsg2ConfigChatId")

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
func (this MsgClient) SendWxTemplateToTag(ctx context.Context, templateId string, tagId int, url string, data map[string]interface{}) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < this.retry; i++ {
		jwtToken, err := this.genJWT(ctx)
		if err != nil {
			return false, err
		}
		jsonString, err = this.requestSendWxTemplateToTag(ctx, jwtToken, templateId, tagId, url, data)
		if err == nil {
			object, err = this.analysisSendWxTemplateToTag(ctx, jsonString)
			if err == nil {
				return object, err
			}
		}
	}
	return object, err
}

//发送微信模板信息
func (this MsgClient) analysisSendWxTemplateToTag(ctx context.Context, jsonString string) (bool, error) {
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
	if response.Code != consd.HttpSuccessCode {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"jsonString": jsonString}).Error("发送微信模板信息，失败")
		return false, fmt.Errorf("发送微信模板信息，失败: %+v", jsonString)
	}
	return true, nil
}

//发送微信模板信息
func (this MsgClient) requestSendWxTemplateToTag(ctx context.Context, jwtToken string, templateId string, tagId int, url string, data map[string]interface{}) (string, error) {
	response, err := this.httpClient.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetHeader("Authorization", "Bearer "+jwtToken).
		SetHeader(util.LogIdKey, fmt.Sprint(util.GetLogId(ctx))).
		SetBody(map[string]interface{}{
			"template_id": templateId,
			"tag_id":      tagId,
			"url":         url,
			"data":        data,
		}).
		Post(this.address + "/api/sendTemplateToTag")

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
func (this MsgClient) SendTemplateToCommonTag(ctx context.Context, text string) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < this.retry; i++ {
		jwtToken, err := this.genJWT(ctx)
		if err != nil {
			return false, err
		}
		jsonString, err = this.requestSendTemplateToCommonTag(ctx, jwtToken, text)
		if err == nil {
			object, err = this.analysisSendTemplateToCommonTag(ctx, jsonString)
			if err == nil {
				return object, err
			}
		}
	}
	return object, err
}

//发送微信通用模板信息
func (this MsgClient) analysisSendTemplateToCommonTag(ctx context.Context, jsonString string) (bool, error) {
	return this.analysisSendWxTemplateToTag(ctx, jsonString)
}

//发送微信通用模板信息
func (this MsgClient) requestSendTemplateToCommonTag(ctx context.Context, jwtToken string, text string) (string, error) {
	response, err := this.httpClient.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetHeader("Authorization", "Bearer "+jwtToken).
		SetHeader(util.LogIdKey, fmt.Sprint(util.GetLogId(ctx))).
		SetBody(map[string]interface{}{
			"text": text,
		}).
		Post(this.address + "/api/sendTemplateToCommonTag")

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

func (this MsgClient) genJWT(ctx context.Context) (string, error) {
	now := time.Now()
	var claims common_model.Claims
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = now.Unix() + int64(this.retry*3)
	claims.RequestId = fmt.Sprint(util.GenId())
	jwtToken, err := util.GenJWT(ctx, this.secret, claims)
	return jwtToken, err
}
