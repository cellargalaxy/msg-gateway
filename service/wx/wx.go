package wx

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/go-resty/resty/v2"
	"time"
)

var httpClient *resty.Client

func init() {
	httpClient = util.GetHttpClient()
	ctx := util.GenCtx()
	flushAccessToken(ctx)
	go func() {
		for {
			time.Sleep(30 * time.Minute)
			ctx := util.GenCtx()
			flushAccessToken(ctx)
		}
	}()
}
