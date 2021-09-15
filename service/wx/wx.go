package wx

import (
	"context"
	"crypto/tls"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/config"
	"github.com/go-resty/resty/v2"
	"time"
)

var httpClient *resty.Client

func init() {
	httpClient = resty.New().
		SetTimeout(config.Config.Timeout).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	cxt := context.Background()
	cxt = util.SetLogId(cxt)
	flushAccessToken(cxt)
	go func() {
		for {
			time.Sleep(30 * time.Minute)
			cxt := context.Background()
			cxt = util.SetLogId(cxt)
			flushAccessToken(cxt)
		}
	}()
}
