package wechat

import (
	"context"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"time"
)

var httpClient *resty.Client
var accessToken string

func init() {
	httpClient = util.GetHttpClient()
	ctx := util.GenCtx()
	_, err := util.NewForeverSingleGoPool(ctx, "刷新微信accessToken", time.Second, flushAccessToken)
	if err != nil {
		panic(err)
	}
}

func flushAccessToken(ctx context.Context, cancel func()) {
	defer util.Defer(func(err interface{}, stack string) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err, "stack": stack}).Error("刷新微信accessToken，退出")
	})

	for {
		ctx := util.ResetLogId(ctx)
		token, err := HttpAccessToken(ctx)
		if token != "" && err == nil {
			accessToken = token
		}
		util.SleepWare(ctx, time.Minute*10)
	}
}

func GetAccessToken(ctx context.Context) string {
	return accessToken
}

type WechatResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (this *WechatResponse) String() string {
	return util.ToJsonString(this)
}
func (this *WechatResponse) HttpSuccess(ctx context.Context) error {
	if this.ErrCode == 0 {
		return nil
	}
	return fmt.Errorf("微信响应失败: %+v", this)
}
