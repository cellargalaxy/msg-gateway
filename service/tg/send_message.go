package tg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

//发送tg信息
func SendMsg(ctx context.Context, chatId int64, text string) (bool, error) {
	content := fmt.Sprintf("```\n%+v\n```\nlogid: ```%+v```", text, util.GetLogId(ctx))
	var jsonString string
	var object bool
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestSendMsg(ctx, chatId, content)
		if err == nil {
			object, err = analysisSendMsg(ctx, jsonString)
			if err == nil {
				return object, err
			}
		}
	}
	return object, err
}

//发送tg信息
func analysisSendMsg(ctx context.Context, jsonString string) (bool, error) {
	type Response struct {
		Ok bool `json:"ok"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("发送tg信息，解析响应异常")
		return false, fmt.Errorf("发送tg信息，解析响应异常")
	}
	if !response.Ok {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"jsonString": jsonString}).Error("发送tg信息，失败")
		return false, fmt.Errorf("发送tg信息，失败")
	}
	return true, nil
}

//发送tg信息
func requestSendMsg(ctx context.Context, chatId int64, text string) (string, error) {
	response, err := httpClient.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetQueryParam("parse_mode", "MarkdownV2").
		SetQueryParam("chat_id", fmt.Sprint(chatId)).
		SetQueryParam("text", text).
		Get(fmt.Sprintf("https://api.telegram.org/bot%+v/sendMessage", config.Config.TgToken))

	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("发送tg信息，请求异常")
		return "", fmt.Errorf("发送tg信息，请求异常")
	}
	if response == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("发送tg信息，响应为空")
		return "", fmt.Errorf("发送tg信息，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithContext(ctx).WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("发送tg信息，响应")
	if statusCode != http.StatusOK {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"StatusCode": statusCode}).Error("发送tg信息，响应码失败")
		return "", fmt.Errorf("发送tg信息，响应码失败: %+v", statusCode)
	}
	return body, nil
}
