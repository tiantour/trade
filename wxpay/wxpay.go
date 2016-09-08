package wxpay

import "encoding/xml"

const (
	success = "SUCCESS"
)

var (
	// Notify 通知
	Notify = &notify{}
	// UnifiedOrder 预支付
	UnifiedOrder = &unifiedorder{}
)

type (
	// App App
	App struct{}
	// Wap Wap
	Wap struct{}
	// notify
	notify struct{}
	// unifiedorder
	unifiedorder struct {
		App
		Wap
	}
	//UnifiedOrderInput 统一下单输入
	UnifiedOrderInput struct {
		XMLName        xml.Name `xml:"xml"`
		AppID          string   `xml:"appid"`            // 公众账号ID，必填
		MchID          string   `xml:"mch_id"`           // 商户号，必填
		NonceStr       string   `xml:"nonce_str"`        // 随机字符串，必填
		Sign           string   `xml:"sign"`             // 签名，必填
		Body           string   `xml:"body"`             // 商品描述，必填
		Detail         string   `xml:"detail"`           // 商品详情，非必填
		ProductID      string   `xml:"product_id"`       // 商品ID,非必填
		OutTradeNo     string   `xml:"out_trade_no"`     // 商户订单号，必填
		TotalFee       int      `xml:"total_fee"`        // 总金额，必填
		SpbillCreateIP string   `xml:"spbill_create_ip"` // 终端IP，必填
		NotifyURL      string   `xml:"notify_url"`       // 通知地址,必填
		TradeType      string   `xml:"trade_type"`       // 交易类型,必填
		OpenID         string   `xml:"openid"`           //openid，公众号支付必填
	}

	//UnifiedOrderOutput 统一下单输出
	UnifiedOrderOutput struct {
		XMLName    xml.Name `xml:"xml"`
		ReturnCode string   `xml:"return_code"`  // 返回状态码，必填
		ReturnMsg  string   `xml:"return_msg"`   //返回信息,非必填
		AppID      string   `xml:"appid"`        // 公众账号ID，必填
		MchID      string   `xml:"mch_id"`       // 商户号，必填
		NonceStr   string   `xml:"nonce_str"`    // 随机字符串，必填
		Sign       string   `xml:"sign"`         // 签名，必填
		ResultCode string   `xml:"result_code"`  // 业务结果,必填
		ErrCode    string   `xml:"err_code"`     // 错误代码，非必填
		ErrCodeDes string   `xml:"err_code_des"` // 错误代码，非必填
		TradeType  string   `xml:"trade_type"`   // 交易类型，必填
		PrepayID   string   `xml:"prepay_id"`    // 预支付交易会话标识
		CodeURL    string   `xml:"code_url"`     // 二维码链接
	}
	//NotifyPayInput  接收微信支付通知
	NotifyPayInput struct {
		XMLName       xml.Name `xml:"xml"`
		ReturnCode    string   `xml:"return_code"`    // 返回状态码，必填
		ReturnMsg     string   `xml:"return_msg"`     //返回信息,非必填
		AppID         string   `xml:"appid"`          // 公众账号ID，必填
		MchID         string   `xml:"mch_id"`         // 商户号，必填
		DeviceInfo    string   `xml:"device_info"`    //设备号，非必填
		NonceStr      string   `xml:"nonce_str"`      // 随机字符串，必填
		Sign          string   `xml:"sign"`           // 签名，必填
		ResultCode    string   `xml:"result_code"`    // 业务结果,必填
		ErrCode       string   `xml:"err_code"`       // 错误代码，非必填
		ErrCodeDes    string   `xml:"err_code_des"`   // 错误代码，非必填
		OpenID        string   `xml:"openid"`         // 用户标识，必填
		IsSubscribe   string   `xml:"is_subscribe"`   // 是否关注公众账号，非必填
		TradeType     string   `xml:"trade_type"`     // 交易类型，必填
		BankType      string   `xml:"bank_type"`      // 付款银行，必填
		TotalFee      int      `xml:"total_fee"`      // 总金额，必填
		FeeType       string   `xml:"fee_type"`       // 货币类型，非必填
		CashFee       int      `xml:"cash_fee"`       // 现金支付金额，必填
		CashFeeType   string   `xml:"cash_fee_type"`  // 现金支付货币类型，非必填
		CouponFee     int      `xml:"coupon_fee"`     // 代金券或立减优惠金额，非必填
		CouponCount   int      `xml:"coupon_count"`   // 代金券或立减优惠使用数量，非必填
		TransactionID string   `xml:"transaction_id"` // 微信支付订单号，必填
		OutTradeNo    string   `xml:"out_trade_no"`   // 商户订单号，必填
		Attach        string   `xml:"attach"`         // 商家数据包，非必填
		TimeEnd       string   `xml:"time_end"`       // 支付完成时间，非必填
	}

	//NotifyPayOutput 回复微信支付通知
	NotifyPayOutput struct {
		XMLName    xml.Name `xml:"xml"`
		ReturnCode string   `xml:"return_code"` // 返回状态码，必填
		ReturnMsg  string   `xml:"return_msg"`  // 返回信息，非必填
	}
)
