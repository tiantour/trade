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
		XMLName        xml.Name `xml:"xml" url:"-"`                                                 // 是 公共头部
		AppID          string   `xml:"appid,omitempty" url:"appid,omitempty"`                       // 是 公众账号ID
		MchID          string   `xml:"mch_id,omitempty" url:"mch_id,omitempty"`                     // 是 商户号
		DeviceInfo     string   `xml:"device_info,omitempty" url:"device_info,omitempty"`           // 否 设备号
		NonceStr       string   `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`               // 是 随机字符串
		Sign           string   `xml:"sign,omitempty" url:"-"`                                      // 是 签名
		SignType       string   `xml:"sign_type,omitempty" url:"sign_type,omitempty"`               // 否 签名类型
		FeeType        string   `xml:"fee_type,omitempty" url:"fee_type,omitempty"`                 // 否 币种
		SpbillCreateIP string   `xml:"spbill_create_ip,omitempty" url:"spbill_create_ip,omitempty"` // 是 终端IP
		TimeStart      string   `xml:"time_start,omitempty" url:"time_start,omitempty"`             // 否 开始时间
		TimeExpire     string   `xml:"time_expire,omitempty" url:"time_expire,omitempty"`           // 否 结束时间
		GoodsTag       string   `xml:"goods_tag,omitempty" url:"goods_tag,omitempty"`               // 否 优惠标记
		NotifyURL      string   `xml:"notify_url,omitempty" url:"notify_url,omitempty"`             // 是 通知地址
		TradeType      string   `xml:"trade_type,omitempty" url:"trade_type,omitempty"`             // 是 交易类型 JSAPI || NATIVE || APP
		LimitPay       string   `xml:"limit_pay,omitempty" url:"limit_pay,omitempty"`               // 否 限制支付方式
		OpenID         string   `xml:"openid,omitempty" url:"openid,omitempty"`                     // 是 openid(JSAPI)
		Request
	}
	// Request request
	Request struct {
		Body          string `xml:"body,omitempty" url:"body,omitempty"`                       // 是 商品描述
		Detail        string `xml:"detail,omitempty" url:"detail,omitempty"`                   // 否 商品详情
		Attach        string `xml:"attach,omitempty" url:"attach,omitempty"`                   // 否 附加数据
		TransactionID string `xml:"transaction_id,omitempty" url:"transaction_id,omitempty"`   // 是 微信支付订单号
		OutTradeNo    string `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`       // 是 商户订单号
		TotalFee      int    `xml:"total_fee,omitempty" url:"total_fee,omitempty"`             // 是 总金额
		ProductID     string `xml:"product_id,omitempty" url:"product_id,omitempty"`           // 否 商品ID NATIVE
		OutRefundNo   string `xml:"out_refund_no,omitempty" url:"out_refund_no,omitempty"`     // 是 商户退款单号
		RefundFee     int    `xml:"refund_fee,omitempty" url:"refund_fee,omitempty"`           // 是 退款金额
		RefundFeeType string `xml:"refund_fee_type,omitempty" url:"refund_fee_type,omitempty"` // 否 货币类型
		RefundFeeDesc string `xml:"refund_desc,omitempty" url:"refund_desc,omitempty"`         // 否 退款原因
		RefundAccount string `xml:"refund_account,omitempty" url:"refund_account,omitempty"`   // 否 资金来源
	}
	// Response response
	Response struct {
		XMLName    xml.Name `xml:"xml" url:"-"`                                         // 是 公共头部
		ReturnCode string   `xml:"return_code,omitempty" url:"return_code,omitempty"`   // 是 返回状态码
		ReturnMsg  string   `xml:"return_msg,omitempty" url:"return_msg,omitempty"`     // 否 返回信息
		ResultCode string   `xml:"result_code,omitempty" url:"result_code,omitempty"`   // 是 业务结果
		ErrCode    string   `xml:"err_code,omitempty" url:"err_code,omitempty"`         // 否 错误代码
		ErrCodeDes string   `xml:"err_code_des,omitempty" url:"err_code_des,omitempty"` // 否 错误代码
	}
	// Prepay prepay
	Prepay struct {
		Response
		AppID     string `xml:"appid,omitempty"`      // 是 公众账号ID
		MchID     string `xml:"mch_id,omitempty"`     // 是 商户号
		NonceStr  string `xml:"nonce_str,omitempty"`  // 是 随机字符串
		Sign      string `xml:"sign,omitempty"`       // 是 签名
		TradeType string `xml:"trade_type,omitempty"` // 是 交易类型
		PrepayID  string `xml:"prepay_id,omitempty"`  // 是 预支付交易会话标识
		CodeURL   string `xml:"code_url,omitempty"`   // 否 二维码链接
	}

	// Defray defray
	Defray struct {
		AppID     string `json:"appId,omitempty" url:"appId,omitempty"`         // 是 公众账号ID
		Package   string `json:"package,omitempty" url:"package,omitempty"`     // 是 app:Sign=WXPay,公众号:prepay_id=xxx
		NonceStr  string `json:"nonceStr,omitempty" url:"nonceStr,omitempty"`   // 是 随机字符串
		TimeStamp string `json:"timeStamp,omitempty" url:"timeStamp,omitempty"` // 是 时间戳
		SignType  string `json:"signType,omitempty" url:"signType,omitempty"`   // 是 签名类型（公众号）
		PaySign   string `json:"paySign,omitempty" url:"paySign,omitempty"`     // 是 支付类型（公众号）
		PartnerID string `json:"partnerid,omitempty" url:"partnerid,omitempty"` // 是 商家ID （app）
		PrepayID  string `json:"prepayid,omitempty" url:"prepayid,omitempty"`   // 是 预支付ID（app）
		Sign      string `json:"sign,omitempty" url:"sign,omitempty"`           // 是 签名 (app)
	}
	// Notice notice
	Notice struct {
		Response
		AppID              string `xml:"appid,omitempty" url:"appid,omitempty"`                               // 是 公众账号ID
		MchID              string `xml:"mch_id,omitempty" url:"mch_id,omitempty"`                             // 是 商户号
		DeviceInfo         string `xml:"device_info,omitempty" url:"device_info,omitempty"`                   // 否 设备号
		NonceStr           string `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`                       // 是 随机字符串
		Sign               string `xml:"sign,omitempty" url:"-"`                                              // 是 签名
		SignType           string `xml:"sign_type,omitempty" url:"sign_type,omitempty"`                       // 签名类型
		OpenID             string `xml:"openid,omitempty" url:"openid,omitempty"`                             // 是 用户标识
		IsSubscribe        string `xml:"is_subscribe,omitempty" url:"is_subscribe,omitempty"`                 // 否 是否关注公众账号
		TradeType          string `xml:"trade_type,omitempty" url:"trade_type,omitempty"`                     // 是 交易类型
		BankType           string `xml:"bank_type,omitempty" url:"bank_type,omitempty"`                       // 是 付款银行
		TotalFee           int    `xml:"total_fee,omitempty" url:"total_fee,omitempty"`                       // 是 总金额
		SettlementTotalFee int    `xml:"settlement_total_fee,omitempty" url:"settlement_total_fee,omitempty"` // 是 应付金额
		FeeType            string `xml:"fee_type,omitempty" url:"fee_type,omitempty"`                         // 否 货币类型
		CashFee            int    `xml:"cash_fee,omitempty" url:"cash_fee,omitempty"`                         // 是 现金支付金额
		CashFeeType        string `xml:"cash_fee_type,omitempty" url:"cash_fee_type,omitempty"`               // 否 现金支付货币类型
		CouponFee          int    `xml:"coupon_fee,omitempty" url:"coupon_fee,omitempty"`                     // 否 代金券或立减优惠金额
		CouponCount        int    `xml:"coupon_count,omitempty" url:"coupon_count,omitempty"`                 // 否 代金券或立减优惠使用数量
		CouponTypeN        string `xml:"coupon_type_$n,omitempty" url:"coupon_type_$n,omitempty"`             // 否 代金券类型
		CouponIDN          string `xml:"coupon_id_$n,omitempty" url:"coupon_id_$n,omitempty"`                 // 否 代金券id
		CouponFeeN         int    `xml:"coupon_fee_$n,omitempty" url:"coupon_fee_$n,omitempty"`               // 否 代金券金额
		TransactionID      string `xml:"transaction_id,omitempty" url:"transaction_id,omitempty"`             // 是 微信支付订单号
		OutTradeNo         string `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                 // 是 商户订单号
		Attach             string `xml:"attach,omitempty" url:"attach,omitempty"`                             // 否 商家数据包
		TimeEnd            string `xml:"time_end,omitempty" url:"time_end,omitempty"`                         // 否 支付完成时间
	}
	// Query query
	Query struct {
		Response
		AppID              string `xml:"appid,omitempty" url:"appid,omitempty"`                               // 是 公众账号ID
		MchID              string `xml:"mch_id,omitempty" url:"mch_id,omitempty"`                             // 是 商户号
		NonceStr           string `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`                       // 是 随机字符串
		Sign               string `xml:"sign,omitempty" url:"-"`                                              // 是 签名
		SignType           string `xml:"sign_type,omitempty" url:"sign_type,omitempty"`                       // 签名类型
		DeviceInfo         string `xml:"device_info,omitempty" url:"device_info,omitempty"`                   // 否 设备号
		OpenID             string `xml:"openid,omitempty" url:"openid,omitempty"`                             // 是 用户标识
		IsSubscribe        string `xml:"is_subscribe,omitempty" url:"is_subscribe,omitempty"`                 // 否 是否关注公众账号
		TradeType          string `xml:"trade_type,omitempty" url:"trade_type,omitempty"`                     // 是 交易类型
		TradeState         string `xml:"trade_state,omitempty" url:"trade_state,omitempty"`                   // 是 交易状态
		BankType           string `xml:"bank_type,omitempty" url:"bank_type,omitempty"`                       // 是 付款银行
		TotalFee           int    `xml:"total_fee,omitempty" url:"total_fee,omitempty"`                       // 是 总金额
		SettlementTotalFee int    `xml:"settlement_total_fee,omitempty" url:"settlement_total_fee,omitempty"` // 是 应付金额
		FeeType            string `xml:"fee_type,omitempty" url:"fee_type,omitempty"`                         // 否 货币类型
		CashFee            int    `xml:"cash_fee,omitempty" url:"cash_fee,omitempty"`                         // 是 现金支付金额
		CashFeeType        string `xml:"cash_fee_type,omitempty" url:"cash_fee_type,omitempty"`               // 否 现金支付货币类型
		CouponFee          int    `xml:"coupon_fee,omitempty" url:"coupon_fee,omitempty"`                     // 否 代金券或立减优惠金额
		CouponCount        int    `xml:"coupon_count,omitempty" url:"coupon_count,omitempty"`                 // 否 代金券或立减优惠使用数量
		CouponTypeN        string `xml:"coupon_type_$n,omitempty" url:"coupon_type_$n,omitempty"`             // 否 代金券类型
		CouponIDN          string `xml:"coupon_id_$n,omitempty" url:"coupon_id_$n,omitempty"`                 // 否 代金券id
		CouponFeeN         int    `xml:"coupon_fee_$n,omitempty" url:"coupon_fee_$n,omitempty"`               // 否 代金券金额
		TransactionID      string `xml:"transaction_id,omitempty" url:"transaction_id,omitempty"`             // 是 微信支付订单号
		OutTradeNo         string `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                 // 是 商户订单号
		Attach             string `xml:"attach,omitempty" url:"attach,omitempty"`                             // 否 商家数据包
		TimeEnd            string `xml:"time_end,omitempty" url:"time_end,omitempty"`                         // 否 支付完成时间
		TradeStateDesc     string `xml:"trade_state_desc,omitempty" url:"trade_state_desc,omitempty"`         // 是 交易状态描述
	}
	// Refund refund
	Refund struct {
		Response
		AppID               string `xml:"appid,omitempty" url:"appid,omitempty"`                                 // 是 公众账号ID
		MchID               string `xml:"mch_id,omitempty" url:"mch_id,omitempty"`                               // 是 商户号
		NonceStr            string `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`                         // 是 随机字符串
		Sign                string `xml:"sign,omitempty" url:"-"`                                                // 是 签名
		TransactionID       string `xml:"transaction_id,omitempty" url:"transaction_id,omitempty"`               // 是 微信支付订单号
		OutTradeNo          string `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                   // 是 商户订单号
		OutRefundNo         string `xml:"out_refund_no,omitempty" url:"out_refund_no,omitempty"`                 // 是 商户退款单号
		RefundID            string `xml:"refund_id,omitempty" url:"refund_id,omitempty"`                         // 是 微信退款订单号
		RefundFee           int    `xml:"refund_fee,omitempty" url:"refund_fee,omitempty"`                       // 是 退款金额
		SettlementRefundFee int    `xml:"settlement_refund_fee,omitempty" url:"settlement_refund_fee,omitempty"` // 否 应结退款金额
		TotalFee            int    `xml:"total_fee,omitempty" url:"total_fee,omitempty"`                         // 是 标价金额
		SettlementTotalFee  int    `xml:"settlement_total_fee,omitempty" url:"settlement_total_fee,omitempty"`   // 是 应结订单金额
		FeeType             string `xml:"fee_type,omitempty" url:"fee_type,omitempty"`                           // 否 标价币种
		CashFee             int    `xml:"cash_fee,omitempty" url:"cash_fee,omitempty"`                           // 是 现金支付金额
		CashFeeType         string `xml:"cash_fee_type,omitempty" url:"cash_fee_type,omitempty"`                 // 否 现金支付币种
		CashRefundFee       int    `xml:"cash_refund_fee,omitempty" url:"cash_refund_fee,omitempty"`             // 否 现金退款金额
		CouponTypeN         string `xml:"coupon_type_$n,omitempty" url:"coupon_type_$n,omitempty"`               // 否 代金券类型
		CouponRefundFee     int    `xml:"coupon_refund_fee,omitempty" url:"coupon_refund_fee,omitempty"`         // 否 代金券退款总金额
		CouponRefundFeeN    int    `xml:"coupon_refund_fee_$n,omitempty" url:"coupon_refund_fee_$n,omitempty"`   // 否 单个代金券退款金额
		CouponRefundCount   int    `xml:"coupon_refund_count,omitempty" url:"coupon_refund_count,omitempty"`     // 否 退款代金券使用数量
		CouponRefundIDN     string `xml:"coupon_refund_id_$n,omitempty" url:"coupon_refund_id_$n,omitempty"`     // 否 退款代金券ID
	}
)
