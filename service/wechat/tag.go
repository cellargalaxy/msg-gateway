package wechat

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/go-resty/resty/v2"
	"time"
)

//为微信用户删标签
func DeleteTagFromUser(ctx context.Context, tagId int, openIds []string) (bool, error) {
	type Response struct {
		WechatResponse
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "为微信用户删标签", util.TryDefault, []time.Duration{0}, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetHeader("Content-Type", "application/json;CHARSET=utf-8").
			SetQueryParam("access_token", GetAccessToken(ctx)).
			SetBody(map[string]interface{}{
				"tagid":       tagId,
				"openid_list": openIds,
			}).
			Post("https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging")
		return response, err
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

//给微信用户加标签
func AddTagToUser(ctx context.Context, tagId int, openIds []string) (bool, error) {
	type Response struct {
		WechatResponse
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "给微信用户加标签", util.TryDefault, []time.Duration{0}, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetHeader("Content-Type", "application/json;CHARSET=utf-8").
			SetQueryParam("access_token", GetAccessToken(ctx)).
			SetBody(map[string]interface{}{
				"tagid":       tagId,
				"openid_list": openIds,
			}).
			Post("https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging")
		return response, err
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

//删除微信标签
func DeleteTag(ctx context.Context, tagId int) (bool, error) {
	type Response struct {
		WechatResponse
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "删除微信标签", util.TryDefault, []time.Duration{0}, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetHeader("Content-Type", "application/json;CHARSET=utf-8").
			SetQueryParam("access_token", GetAccessToken(ctx)).
			SetBody(map[string]interface{}{
				"tag": map[string]interface{}{"id": tagId},
			}).
			Post("https://api.weixin.qq.com/cgi-bin/tags/delete")
		return response, err
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

//获取微信所有标签
func ListAllTag(ctx context.Context) ([]model.Tag, error) {
	type Response struct {
		WechatResponse
		Tags []model.Tag `json:"tags"`
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "获取微信所有标签", util.TryDefault, []time.Duration{0}, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetQueryParam("access_token", GetAccessToken(ctx)).
			Get("https://api.weixin.qq.com/cgi-bin/tags/get")
		return response, err
	})
	if err != nil {
		return nil, err
	}

	return response.Tags, nil
}

//创建微信标签
func CreateTag(ctx context.Context, tag string) (bool, error) {
	type Response struct {
		WechatResponse
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "创建微信标签", util.TryDefault, []time.Duration{0}, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetHeader("Content-Type", "application/json;CHARSET=utf-8").
			SetQueryParam("access_token", GetAccessToken(ctx)).
			SetBody(map[string]interface{}{
				"tag": map[string]string{"name": tag},
			}).
			Post("https://api.weixin.qq.com/cgi-bin/tags/create")
		return response, err
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
