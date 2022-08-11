package wechat

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/config"
	"github.com/go-resty/resty/v2"
	"time"
)

//获取微信accessToken
func HttpAccessToken(ctx context.Context) (string, error) {
	type Response struct {
		WechatResponse
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	var response Response
	err := util.HttpApiWithTry(ctx, "获取微信accessToken", util.TryDefault, []time.Duration{0}, &response, func() (*resty.Response, error) {
		response, err := httpClient.R().SetContext(ctx).
			SetQueryParam("appid", config.Config.WxAppId).
			SetQueryParam("secret", config.Config.WxAppSecret).
			SetQueryParam("grant_type", "client_credential").
			Get("https://api.weixin.qq.com/cgi-bin/token")
		return response, err
	})
	if err != nil {
		return "", err
	}

	return response.AccessToken, nil
}
