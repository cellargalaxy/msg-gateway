package wechat

import (
	"context"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/config"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"time"
)

//给微信通用标签用户发送模板消息
func SendTemplateToCommonTag(ctx context.Context, text string) ([]string, error) {
	dataMap := make(map[string]interface{})
	dataMap["logid"] = util.GetLogId(ctx)
	dataMap["text"] = text
	return SendTemplateToTag(ctx, config.Config.WxCommonTempId, config.Config.WxCommonTagId, fmt.Sprintf("https://wx2.qq.com?logid=%+v&text=%+v", dataMap["logid"], dataMap["text"]), dataMap)
}

//给微信标签用户发送模板消息
func SendTemplateToTag(ctx context.Context, templateId string, tagId int, url string, dataMap map[string]interface{}) ([]string, error) {
	data := map[string]model.TemplateData{}
	for key, value := range dataMap {
		data[key] = model.TemplateData{Value: value}
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"data": data}).Info("给微信标签用户发送模板消息")
	openIds, err := ListOpenIdByTagId(ctx, tagId)
	if err != nil {
		return nil, err
	}
	if len(openIds) == 0 {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Warn("给微信标签用户发送模板消息，openIds为空")
		return nil, nil
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

//获取微信所有模板
func ListAllTemplate(ctx context.Context) ([]model.Template, error) {
	type Response struct {
		WechatResponse
		TemplateList []model.Template `json:"template_list"`
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "获取微信所有模板", util.TryDefault, []time.Duration{0}, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetQueryParam("access_token", GetAccessToken(ctx)).
			Get("https://api.weixin.qq.com/cgi-bin/template/get_all_private_template")
		return response, err
	})
	if err != nil {
		return nil, err
	}

	return response.TemplateList, nil
}

//发送微信模板信息
func SendTemplate(ctx context.Context, openId string, templateId string, url string, data map[string]model.TemplateData) (bool, error) {
	type Response struct {
		WechatResponse
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "发送微信模板信息", util.TryDefault, []time.Duration{0}, &response, func() (*resty.Response, error) {
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
		return response, err
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
