package wxpay

import "encoding/xml"

const (
	// Success status success
	Success = "SUCCESS"
	// Fail status fail
	Fail = "FAIL"
)

type (
	// Sign sign
	Sign struct {
		XMLName        xml.Name `xml:"xml" url:"-"`                                       // 是 公共头部
		AppID          string   `xml:"appid" url:"appid,omitempty"`                       // 是 公众账号ID
		MchID          string   `xml:"mch_id" url:"mch_id,omitempty"`                     // 是 商户号
		NonceStr       string   `xml:"nonce_str" url:"nonce_str,omitempty"`               // 是 随机字符串
		Sign           string   `xml:"sign" url:"-"`                                      // 是 签名
		TradeType      string   `xml:"trade_type" url:"trade_type,omitempty"`             // 是 交易类型
		SpbillCreateIP string   `xml:"spbill_create_ip" url:"spbill_create_ip,omitempty"` // 是 终端IP
		NotifyURL      string   `xml:"notify_url" url:"notify_url,omitempty"`             // 是 通知地址
		OpenID         string   `xml:"openid" url:"openid,omitempty"`                     // 是 openid(公众号支付)
		Request
	}
	// Request request
	Request struct {
		Body       string `xml:"body" url:"body,omitempty"`                 // 是 商品描述
		Detail     string `xml:"detail" url:"detail,omitempty"`             // 否 商品详情
		ProductID  string `xml:"product_id" url:"product_id,omitempty"`     // 否 商品ID
		OutTradeNo string `xml:"out_trade_no" url:"out_trade_no,omitempty"` // 是 商户订单号
		TotalFee   int    `xml:"total_fee" url:"total_fee,omitempty"`       // 是 总金额
	}
	// Response response
	Response struct {
		XMLName    xml.Name `xml:"xml" url:"-"`          // 是 公共头部
		ReturnCode string   `xml:"return_code" url:"-"`  // 是 返回状态码
		ReturnMsg  string   `xml:"return_msg" url:"-"`   // 否 返回信息
		ResultCode string   `xml:"result_code" url:"-"`  // 是 业务结果
		ErrCode    string   `xml:"err_code" url:"-"`     // 否 错误代码
		ErrCodeDes string   `xml:"err_code_des" url:"-"` // 否 错误代码
	}
	// Prepay prepay
	Prepay struct {
		Response
		AppID     string `xml:"appid"`      // 是 公众账号ID
		MchID     string `xml:"mch_id"`     // 是 商户号
		NonceStr  string `xml:"nonce_str"`  // 是 随机字符串
		Sign      string `xml:"sign"`       // 是 签名
		TradeType string `xml:"trade_type"` // 是 交易类型
		PrepayID  string `xml:"prepay_id"`  // 是 预支付交易会话标识
		CodeURL   string `xml:"code_url"`   // 否 二维码链接
	}
	// Defray defray
	Defray struct {
		AppID     string `json:"appId" url:"appId,omitempty"`                   // 是 公众账号ID
		Package   string `json:"package" url:"package,omitempty"`               // 是 app:Sign=WXPay,公众号:prepay_id=xxx
		NonceStr  string `json:"nonceStr" url:"nonceStr,omitempty"`             // 是 随机字符串
		TimeStamp string `json:"timeStamp" url:"timeStamp,omitempty"`           // 是 时间戳
		SignType  string `json:"signType,omitempty" url:"signType,omitempty"`   // 是 签名类型（公众号）
		PaySign   string `json:"paySign,omitempty" url:"paySign,omitempty"`     // 是 支付类型（公众号）
		PartnerID string `json:"partnerid,omitempty" url:"partnerid,omitempty"` // 是 商家ID （app）
		PrepayID  string `json:"prepayid,omitempty" url:"prepayid,omitempty"`   // 是 预支付ID（app）
		Sign      string `json:"sign,omitempty" url:"sign,omitempty"`           // 是 签名 (app)
	}
	// Notice notice
	Notice struct {
		Response
		AppID         string `xml:"appid" url:"appid,omitempty"`                   // 是 公众账号ID
		MchID         string `xml:"mch_id" url:"mch_id,omitempty"`                 // 是 商户号
		DeviceInfo    string `xml:"device_info" url:"device_info,omitempty"`       // 否 设备号
		NonceStr      string `xml:"nonce_str" url:"nonce_str,omitempty"`           // 是 随机字符串
		Sign          string `xml:"sign" url:"-"`                                  // 是 签名
		OpenID        string `xml:"openid" url:"openid,omitempty"`                 // 是 用户标识
		IsSubscribe   string `xml:"is_subscribe" url:"is_subscribe,omitempty"`     // 否 是否关注公众账号
		TradeType     string `xml:"trade_type" url:"trade_type,omitempty"`         // 是 交易类型
		BankType      string `xml:"bank_type" url:"bank_type,omitempty"`           // 是 付款银行
		TotalFee      int    `xml:"total_fee" url:"total_fee,omitempty"`           // 是 总金额
		FeeType       string `xml:"fee_type" url:"fee_type,omitempty"`             // 否 货币类型
		CashFee       int    `xml:"cash_fee" url:"cash_fee,omitempty"`             // 是 现金支付金额
		CashFeeType   string `xml:"cash_fee_type" url:"cash_fee_type,omitempty"`   // 否 现金支付货币类型
		CouponFee     int    `xml:"coupon_fee" url:"coupon_fee,omitempty"`         // 否 代金券或立减优惠金额
		CouponCount   int    `xml:"coupon_count" url:"coupon_count,omitempty"`     // 否 代金券或立减优惠使用数量
		TransactionID string `xml:"transaction_id" url:"transaction_id,omitempty"` // 是 微信支付订单号
		OutTradeNo    string `xml:"out_trade_no" url:"out_trade_no,omitempty"`     // 是 商户订单号
		Attach        string `xml:"attach" url:"attach,omitempty"`                 // 否 商家数据包
		TimeEnd       string `xml:"time_end" url:"time_end,omitempty"`             // 否 支付完成时间
	}
)
