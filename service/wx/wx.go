package wx

import (
	"crypto/tls"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/config"
	"github.com/go-resty/resty/v2"
	"time"
)

var httpClient *resty.Client

func init() {
	httpClient = resty.New().
		SetTimeout(config.Config.Timeout).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	ctx := util.CreateLogCtx()
	flushAccessToken(ctx)
	go func() {
		for {
			time.Sleep(30 * time.Minute)
			ctx := util.CreateLogCtx()
			flushAccessToken(ctx)
		}
	}()
}
