package wx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/config"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

//给通用标签用户发送模板消息
func SendTemplateToCommonTag(ctx context.Context, text string) ([]string, error) {
	dataMap := make(map[string]interface{})
	dataMap["logid"] = util.GetLogId(ctx)
	dataMap["text"] = text
	return SendTemplateToTag(ctx, config.Config.WxCommonTempId, config.Config.WxCommonTagId, fmt.Sprintf("https://wx2.qq.com?logid=%+v&text=%+v", dataMap["logid"], dataMap["text"]), dataMap)
}

//给标签用户发送模板消息
func SendTemplateToTag(ctx context.Context, templateId string, tagId int, url string, dataMap map[string]interface{}) ([]string, error) {
	data := map[string]model.TemplateData{}
	for key, value := range dataMap {
		data[key] = model.TemplateData{Value: value}
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"data": data}).Info("给标签用户发送模板消息")
	openIds, err := ListOpenIdByTagId(ctx, tagId)
	if err != nil {
		return nil, err
	}
	var failOpenIds []string
	for i := range openIds {
		success, _ := SendTemplate(ctx, openIds[i], templateId, url, data)
		if !success {
			failOpenIds = append(failOpenIds, openIds[i])
		}
	}
	return failOpenIds, nil
}

//获取全部模板
func ListAllTemplate(ctx context.Context) ([]model.Template, error) {
	var jsonString string
	var object []model.Template
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestListAllTemplate(ctx)
		if err == nil {
			object, err = analysisListAllTemplate(ctx, jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken(ctx)
	}
	return object, err
}

//获取所有模板
func analysisListAllTemplate(ctx context.Context, jsonString string) ([]model.Template, error) {
	type Response struct {
		ErrCode      int              `json:"errcode"`
		ErrMsg       string           `json:"errmsg"`
		TemplateList []model.Template `json:"template_list"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取所有模板，解析响应异常")
		return nil, fmt.Errorf("获取所有模板，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"jsonString": jsonString}).Error("获取所有模板，失败")
		return nil, fmt.Errorf("获取所有模板，失败")
	}
	return response.TemplateList, nil
}

//获取所有模板
func requestListAllTemplate(ctx context.Context) (string, error) {
	response, err := httpClient.R().SetContext(ctx).
		SetQueryParam("access_token", GetAccessToken(ctx)).
		Get("https://api.weixin.qq.com/cgi-bin/template/get_all_private_template")

	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("获取所有模板，请求异常")
		return "", fmt.Errorf("获取所有模板，请求异常")
	}
	if response == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("获取所有模板，响应为空")
		return "", fmt.Errorf("获取所有模板，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithContext(ctx).WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("获取所有模板，响应")
	if statusCode != http.StatusOK {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"StatusCode": statusCode}).Error("获取所有模板，响应码失败")
		return "", fmt.Errorf("获取所有模板，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//发送模板信息
func SendTemplate(ctx context.Context, openId string, templateId string, url string, data map[string]model.TemplateData) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestSendTemplate(ctx, openId, templateId, url, data)
		if err == nil {
			object, err = analysisSendTemplate(ctx, jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken(ctx)
	}
	return object, err
}

//发送模板信息
func analysisSendTemplate(ctx context.Context, jsonString string) (bool, error) {
	type Response struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("发送模板信息，解析响应异常")
		return false, fmt.Errorf("发送模板信息，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"jsonString": jsonString}).Error("发送模板信息，失败")
		return false, fmt.Errorf("发送模板信息，失败")
	}
	return true, nil
}

//发送模板信息
func requestSendTemplate(ctx context.Context, openId string, templateId string, url string, data map[string]model.TemplateData) (string, error) {
	response, err := httpClient.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetQueryParam("access_token", GetAccessToken(ctx)).
		SetBody(map[string]interface{}{
			"touser":      openId,
			"template_id": templateId,
			"url":         url,
			"data":        data,
		}).
		Post("https://api.weixin.qq.com/cgi-bin/message/template/send")

	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("发送模板信息，请求异常")
		return "", fmt.Errorf("发送模板信息，请求异常")
	}
	if response == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("发送模板信息，响应为空")
		return "", fmt.Errorf("发送模板信息，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithContext(ctx).WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("发送模板信息，响应")
	if statusCode != http.StatusOK {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"StatusCode": statusCode}).Error("发送模板信息，响应码失败")
		return "", fmt.Errorf("发送模板信息，响应码失败: %+v", statusCode)
	}
	return body, nil
}
