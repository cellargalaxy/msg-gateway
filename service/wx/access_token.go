package wx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/msg-gateway/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

var accessToken string

func GetAccessToken(ctx context.Context) string {
	return accessToken
}

func flushAccessToken(ctx context.Context) {
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err := requestAccessToken(ctx)
		if err != nil {
			continue
		}
		token, err := analysisAccessToken(ctx, jsonString)
		if err != nil {
			continue
		}
		accessToken = token
		return
	}
}

//获取accessToken
func analysisAccessToken(ctx context.Context, jsonString string) (string, error) {
	type Response struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取accessToken，解析响应异常")
		return "", fmt.Errorf("获取accessToken，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取accessToken，失败")
		return "", fmt.Errorf("获取accessToken，失败")
	}
	if response.AccessToken == "" {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取accessToken，accessToken为空")
		return "", fmt.Errorf("获取accessToken，accessToken为空")
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"accessToken": len(response.AccessToken)}).Info("accessToken长度")
	return response.AccessToken, nil
}

//获取accessToken
func requestAccessToken(ctx context.Context) (string, error) {
	response, err := httpClient.R().SetContext(ctx).
		SetQueryParam("appid", config.Config.WxAppId).
		SetQueryParam("secret", config.Config.WxAppSecret).
		SetQueryParam("grant_type", "client_credential").
		Get("https://api.weixin.qq.com/cgi-bin/token")

	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("获取accessToken，请求异常")
		return "", fmt.Errorf("获取accessToken，请求异常")
	}
	if response == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("获取accessToken，响应为空")
		return "", fmt.Errorf("获取accessToken，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithContext(ctx).WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("获取accessToken，响应")
	if statusCode != http.StatusOK {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"StatusCode": statusCode}).Error("获取accessToken，响应码失败")
		return "", fmt.Errorf("获取accessToken，响应码失败: %+v", statusCode)
	}
	return body, nil
}
