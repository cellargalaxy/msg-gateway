package wechat

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

//获取微信全部用户信息
func ListAllUserInfo(ctx context.Context) ([]model.UserInfo, error) {
	openIds, err := ListAllOpenId(ctx)
	if err != nil {
		return nil, err
	}
	return ListUserInfo(ctx, openIds)
}

//获取微信全部openId
func ListAllOpenId(ctx context.Context) ([]string, error) {
	type Data struct {
		OpenIds []string `json:"openid"`
	}
	type Response struct {
		WechatResponse
		Data Data `json:"data"`
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "获取微信全部openId", util.TryDefault, nil, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetQueryParam("access_token", GetAccessToken(ctx)).
			Get("https://api.weixin.qq.com/cgi-bin/user/get")
		return response, err
	})
	if err != nil {
		return nil, err
	}

	return response.Data.OpenIds, nil
}

//获取微信用户信息
func ListUserInfo(ctx context.Context, openIds []string) ([]model.UserInfo, error) {
	var userList []map[string]interface{}
	for i := range openIds {
		userList = append(userList, map[string]interface{}{"openid": openIds[i], "lang": "zh_CN"})
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"userList": userList}).Info("获取微信用户信息")

	type Response struct {
		WechatResponse
		UserInfoList []model.UserInfo `json:"user_info_list"`
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "获取微信用户信息", util.TryDefault, nil, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetHeader("Content-Type", "application/json;CHARSET=utf-8").
			SetQueryParam("access_token", GetAccessToken(ctx)).
			SetBody(map[string]interface{}{
				"user_list": userList,
			}).
			Post("https://api.weixin.qq.com/cgi-bin/user/info/batchget")
		return response, err
	})
	if err != nil {
		return nil, err
	}

	return response.UserInfoList, nil
}

//获取微信标签下的openId
func ListOpenIdByTagId(ctx context.Context, tagId int) ([]string, error) {
	type Data struct {
		OpenIds []string `json:"openid"`
	}
	type Response struct {
		WechatResponse
		Data Data `json:"data"`
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "获取微信标签下的openId", util.TryDefault, nil, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetHeader("Content-Type", "application/json;CHARSET=utf-8").
			SetQueryParam("access_token", GetAccessToken(ctx)).
			SetBody(map[string]interface{}{
				"tagid":       tagId,
				"next_openid": "",
			}).
			Post("https://api.weixin.qq.com/cgi-bin/user/tag/get")
		return response, err
	})
	if err != nil {
		return nil, err
	}

	return response.Data.OpenIds, nil
}
