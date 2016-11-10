package trade

import (
	"github.com/tiantour/trade/alipayx"
	"github.com/tiantour/trade/wxpayx"
)

// trade trade
var (
	Alipay = &alipayX{}
	Wxpay  = &wxpayX{}
)

type (
	alipayX struct {
		App alipayx.App // 支付宝app支付
		Wap alipayx.Wap // 支付宝h5支付
	}
	wxpayX struct {
		App wxpayx.App // 微信app支付
		Wap wxpayx.Wap // 微信公众号支付
	}
)
