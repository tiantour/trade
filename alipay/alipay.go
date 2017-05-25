package alipay

type (
	// Sign sign
	Sign struct {
		AppID      string `json:"app_id,omitempty" url:"app_id,omitempty"`           // 是 应用ID
		Method     string `json:"method,omitempty" url:"method,omitempty"`           // 是 接口名称
		Format     string `json:"format,omitempty" url:"format,omitempty"`           // 否 JSON
		ReturnURL  string `json:"return_url,omitempty" url:"return_url,omitempty"`   // 否 跳转地址
		Charset    string `json:"charset,omitempty" url:"charset,omitempty"`         // 是 utf-8
		SignType   string `json:"sign_type,omitempty" url:"sign_type,omitempty"`     // 是 RSA2
		Sign       string `json:"sign,omitempty" url:"sign,omitempty"`               // 是 签名
		TimeStamp  string `json:"timestamp,omitempty" url:"timestamp,omitempty"`     // 是 时间
		Version    string `json:"version,omitempty" url:"version,omitempty"`         // 是 1.0
		NotifyURL  string `json:"notify_url,omitempty" url:"notify_url,omitempty"`   // 否 通知地址
		BizContent string `json:"biz_content,omitempty" url:"biz_content,omitempty"` // 是 非公共参数
	}
	// Request request
	Request struct {
		Body               string  `json:"body,omitempty" url:"body,omitempty"`                                 // 否 描述
		Subject            string  `json:"subject,omitempty" url:"subject,omitempty"`                           // 是 标题
		OutTradeNo         string  `json:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                 // 是 单号
		TimeExpress        string  `json:"timeout_express,omitempty" url:"timeout_express,omitempty"`           // 否 过期
		TotalAmount        float64 `json:"total_amount,omitempty" url:"total_amount,omitempty"`                 // 是 金额
		SellerID           string  `json:"seller_id,omitempty" url:"seller_id,omitempty"`                       // 否 卖家ID
		AuthToken          string  `json:"auth_token,omitempty" url:"auth_token,omitempty"`                     // 否 授权
		ProductCode        string  `json:"product_code,omitempty" url:"product_code,omitempty"`                 // 是 QUICK_WAP_PAY
		GoodsType          string  `json:"goods_type,omitempty" url:"goods_type,omitempty"`                     // 否 0 虚拟商品 1 实物商品
		PassbackParams     string  `json:"passback_params,omitempty" url:"passback_params,omitempty"`           // 否 回传参数
		PromoParams        string  `json:"promo_params,omitempty" url:"promo_params,omitempty"`                 // 否 优惠参数
		ExtendParams       string  `json:"extend_params,omitempty" url:"extend_params,omitempty"`               // 否 拓展参数
		EnablePayChannels  string  `json:"enable_pay_channels,omitempty" url:"enable_pay_channels,omitempty"`   // 否 可用渠道
		DisablePayChannels string  `json:"disable_pay_channels,omitempty" url:"disable_pay_channels,omitempty"` // 否 禁用渠道
		StoreID            string  `json:"store_id,omitempty" url:"store_id"`                                   // 否 门店编号
	}
	// Notice notice
	Notice struct {
		NotifyTime        string  `json:"notify_time,omitempty" url:"notify_time,omitempty"`                 // 是 时间
		NotifyType        string  `json:"notify_type,omitempty" url:"notify_type,omitempty"`                 // 是 类型
		NotifyID          string  `json:"notify_id,omitempty" url:"notify_id,omitempty"`                     // 是 校验ID
		APPID             string  `json:"appid,omitempty" url:"appid,omitempty"`                             // 是 应用ID
		Charset           string  `json:"charset,omitempty" url:"charset,omitempty"`                         // 是 utf-8
		Version           string  `json:"version,omitempty" url:"version,omitempty"`                         // 是 1.0
		SignType          string  `json:"sign_type,omitempty" url:"sign_type,omitempty"`                     // 是 RSA2
		Sign              string  `json:"sign,omitempty" url:"sign,omitempty"`                               // 是 签名
		TradeNo           string  `json:"trade_no,omitempty" url:"trade_no,omitempty"`                       // 是 交易号
		OutTradeNo        string  `json:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`               // 是 单号
		OutBizNo          string  `json:"out_biz_no,omitempty" url:"out_biz_no,omitempty"`                   // 否 业务ID
		BuyerID           string  `json:"buyer_id,omitempty" url:"buyer_id,omitempty"`                       // 否 买家帐号
		SellerID          string  `json:"seller_id,omitempty" url:"seller_id,omitempty"`                     // 否 卖家帐号
		SellerEmail       string  `json:"seller_email,omitempty" url:"seller_email,omitempty"`               // 否 卖家邮箱
		TradeStatus       string  `json:"trade_status,omitempty" url:"trade_status,omitempty"`               // 否 交易状态
		TotalAmount       float64 `json:"total_amount,omitempty" url:"total_amount,omitempty"`               // 否 订单金额
		ReceiptAmount     float64 `json:"receipt_amount,omitempty" url:"receipt_amount,omitempty"`           // 否 实收金额
		InvoiceAmount     float64 `json:"invoice_amount,omitempty" url:"invoice_amount,omitempty"`           // 否 开票金额
		BuyerPayAmount    float64 `json:"buyer_pay_amount,omitempty" url:"buyer_pay_amount,omitempty"`       // 否 付款金额
		PointAmount       float64 `json:"point_amount,omitempty" url:"point_amount,omitempty"`               // 否 积分金额
		RefundFee         float64 `json:"refund_fee,omitempty" url:"refund_fee,omitempty"`                   // 否 退款金额
		Subject           string  `json:"subject,omitempty" url:"subject,omitempty"`                         // 是 标题
		Body              string  `json:"body,omitempty" url:"body,omitempty"`                               // 否 描述
		GmtCreate         string  `json:"gmt_create,omitempty" url:"gmt_create,omitempty"`                   // 否 创建时间
		GmtPayment        string  `json:"gmt_payment,omitempty" url:"gmt_payment,omitempty"`                 // 否 付款时间
		GmtRefund         string  `json:"gmt_refund,omitempty" url:"gmt_refund,omitempty"`                   // 否 退款时间
		GmtClose          string  `json:"gmt_close,omitempty" url:"gmt_close,omitempty"`                     // 否 结束时间
		FundBillList      string  `json:"fund_bill_list,omitempty" url:"fund_bill_list,omitempty"`           // 否 渠道金额
		PassbackParams    string  `json:"passback_params,omitempty" url:"passback_params,omitempty"`         // 否 回传参数
		VoucherDetailList string  `json:"voucher_detail_list,omitempty" url:"voucher_detail_list,omitempty"` // 否 优惠券
	}
)
