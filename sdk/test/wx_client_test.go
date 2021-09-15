package test

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/sdk"
	"testing"
	"time"
)

func TestSendWxTemplateToTag(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	util.InitLog("msg_gateway.log")
	wxClient, err := sdk.NewMsgClient(time.Second*3, time.Second*3, 3, "http://127.0.0.1:8990", "secret")
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
	response, err := wxClient.SendWxTemplateToTag(ctx, "7ub0o1jXJGfar5Zaj-imwwoisFiH6xW6CsS4pKWjnKc", 109, "", map[string]interface{}{"zhi": 111})
	test.Logf("response: %+v\r\n", util.ToJsonIndentString(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestSendTgMsg2ConfigChatId(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	util.InitLog("msg_gateway.log")
	wxClient, err := sdk.NewMsgClient(time.Second*3, time.Second*3, 3, "http://127.0.0.1:8990", "secret")
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
	response, err := wxClient.SendTgMsg2ConfigChatId(ctx, `{{date.DATA}}，{{fund_count.DATA}}个基金，{{gain_count.DATA}}个盈，{{loss_count.DATA}}个亏
盈收益{{gain_income_rate.DATA}}%，亏收益{{loss_income_rate.DATA}}%，总收益{{income_rate.DATA}}%
{{aip_buy_count.DATA}}个定投买{{aip_buy_amount.DATA}}￥，{{aip_ransom_count.DATA}}个定投卖{{aip_ransom_amount.DATA}}￥
k年利率{{k_annual.DATA}}%，定投年利率{{aip_annual.DATA}}%
斜率: {{k7.DATA}}‰; {{k15.DATA}}‰; {{k30.DATA}}‰; {{k60.DATA}}‰
持有升{{hold_pos_rate.DATA}}%，市场升{{market_pos_rate.DATA}}%`)
	test.Logf("response: %+v\r\n", util.ToJsonIndentString(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestSendTemplateToCommonTag(test *testing.T) {
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	util.InitLog("msg_gateway.log")
	wxClient, err := sdk.NewMsgClient(time.Second*3, time.Second*3, 3, "http://127.0.0.1:8990", "secret")
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
	response, err := wxClient.SendTemplateToCommonTag(ctx, `{{date.DATA}}，{{fund_count.DATA}}个基金，{{gain_count.DATA}}个盈，{{loss_count.DATA}}个亏
盈收益{{gain_income_rate.DATA}}%，亏收益{{loss_income_rate.DATA}}%，总收益{{income_rate.DATA}}%
{{aip_buy_count.DATA}}个定投买{{aip_buy_amount.DATA}}￥，{{aip_ransom_count.DATA}}个定投卖{{aip_ransom_amount.DATA}}￥
k年利率{{k_annual.DATA}}%，定投年利率{{aip_annual.DATA}}%
斜率: {{k7.DATA}}‰; {{k15.DATA}}‰; {{k30.DATA}}‰; {{k60.DATA}}‰
持有升{{hold_pos_rate.DATA}}%，市场升{{market_pos_rate.DATA}}%`)
	test.Logf("response: %+v\r\n", util.ToJsonIndentString(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
