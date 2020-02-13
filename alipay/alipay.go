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
		OutTradeNo         string  `json:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                 // 是 单号
		SellerID           string  `json:"seller_id,omitempty" url:"seller_id,omitempty"`                       // 否 卖家编号
		TotalAmount        float64 `json:"total_amount,omitempty" url:"total_amount,omitempty"`                 // 是 订单金额
		DiscountableAmount float64 `json:"discountable_amount,omitempty" url:"discountable_amount,omitempty"`   // 是 打折金额
		Subject            string  `json:"subject,omitempty" url:"subject,omitempty"`                           // 是 标题
		Body               string  `json:"body,omitempty" url:"body,omitempty"`                                 // 否 描述
		BuyerID            string  `json:"buyer_id,omitempty" url:"buyer_id,omitempty"`                         // 否 用户号
		ProductCode        string  `json:"product_code,omitempty" url:"product_code,omitempty"`                 // 是 QUICK_WAP_PAY || FAST_INSTANT_TRADE_PAY
		OperatorID         string  `json:"operator_id,omitempty" url:"operator_id,omitempty"`                   // 否 操作员编号
		StoreID            string  `json:"store_id,omitempty" url:"store_id,omitempty"`                         // 否 门店编号
		TerminalID         string  `json:"terminal_id,omitempty" url:"terminal_id,omitempty"`                   // 否 终端编号
		TimeExpress        string  `json:"timeout_express,omitempty" url:"timeout_express,omitempty"`           // 否 过期
		AuthToken          string  `json:"auth_token,omitempty" url:"auth_token,omitempty"`                     // 否 授权
		GoodsType          string  `json:"goods_type,omitempty" url:"goods_type,omitempty"`                     // 否 0 虚拟商品 1 实物商品
		PassbackParams     string  `json:"passback_params,omitempty" url:"passback_params,omitempty"`           // 否 回传参数
		PromoParams        string  `json:"promo_params,omitempty" url:"promo_params,omitempty"`                 // 否 优惠参数
		EnablePayChannels  string  `json:"enable_pay_channels,omitempty" url:"enable_pay_channels,omitempty"`   // 否 可用渠道
		DisablePayChannels string  `json:"disable_pay_channels,omitempty" url:"disable_pay_channels,omitempty"` // 否 禁用渠道
		QrPayMode          string  `json:"qr_pay_mode,omitempty" url:"qr_pay_mode,omitempty"`                   // 否 二维码模式
		QrcodeWidth        string  `json:"qrcode_width,omitempty" url:"qrcode_width,omitempty"`                 // 否 二维码宽度
	}
	// Result result
	Result struct {
		AlipayTradeCrateResponse *Prepay `json:"alipay_trade_create_response,omitempty"` // 内容
		AlipayTradeQueryResponse *Query  `json:"alipay_trade_query_response,omitempty"`  // 内容
		Sign                     string  `json:"sign,omitempty"`                         // 签名
	}
	// Response response
	Response struct {
		Code    string `json:"code,omitempty"`     // 网关返回码
		Msg     string `json:"msg,omitempty"`      // 网关返回码描述
		SubCode string `json:"sub_code,omitempty"` // 业务返回码
		SumbMsg string `json:"sub_msg,omitempty"`  // 业务描述
	}
	// Prepay prepay
	Prepay struct {
		Response
		OutTradeNo string `json:"out_trade_no,omitempty"` // 是 单号
		TradeNo    string `json:"trade_no,omitempty"`     // 是 单号
	}
	// Notice notice
	Notice struct {
		NotifyTime     string `json:"notify_time,omitempty" url:"notify_time,omitempty" schema:"notify_time,omitempty"`                // 是 时间
		NotifyType     string `json:"notify_type,omitempty" url:"notify_type,omitempty" schema:"notify_type,omitempty"`                // 是 类型
		NotifyID       string `json:"notify_id,omitempty" url:"notify_id,omitempty" schema:"notify_id,omitempty"`                      // 是 校验ID
		SignType       string `json:"sign_type,omitempty" url:"-" schema:"-"`                                                          // 是 RSA2
		Sign           string `json:"sign,omitempty" url:"-" schema:"-"`                                                               // 是 签名
		TradeNo        string `json:"trade_no,omitempty" url:"trade_no,omitempty" schema:"trade_no,omitempty"`                         // 是 交易号
		APPID          string `json:"app_id,omitempty" url:"app_id,omitempty" schema:"app_id,omitempty"`                               // 是 应用ID
		OutTradeNo     string `json:"out_trade_no,omitempty" url:"out_trade_no,omitempty" schema:"out_trade_no,omitempty"`             // 是 单号
		OutBizNo       string `json:"out_biz_no,omitempty" url:"out_biz_no,omitempty" schema:"out_biz_no,omitempty"`                   // 否 业务ID
		BuyerID        string `json:"buyer_id,omitempty" url:"buyer_id,omitempty" schema:"buyer_id,omitempty"`                         // 否 买家帐号
		BuyerLogonID   string `json:"buyer_logon_id,omitempty" url:"buyer_logon_id,omitempty" schema:"buyer_logon_id,omitempty"`       // 否 买家帐号
		SellerID       string `json:"seller_id,omitempty" url:"seller_id,omitempty" schema:"seller_id,omitempty"`                      // 否 卖家帐号
		SellerEmail    string `json:"seller_email,omitempty" url:"seller_email,omitempty" schema:"seller_email,omitempty"`             // 否 卖家邮箱
		TradeStatus    string `json:"trade_status,omitempty" url:"trade_status,omitempty" schema:"trade_status,omitempty"`             // 否 交易状态
		TotalAmount    string `json:"total_amount,omitempty" url:"total_amount,omitempty" schema:"total_amount,omitempty"`             // 否 订单金额
		ReceiptAmount  string `json:"receipt_amount,omitempty" url:"receipt_amount,omitempty" schema:"receipt_amount,omitempty"`       // 否 实收金额
		InvoiceAmount  string `json:"invoice_amount,omitempty" url:"invoice_amount,omitempty" schema:"invoice_amount,omitempty"`       // 否 开票金额
		BuyerPayAmount string `json:"buyer_pay_amount,omitempty" url:"buyer_pay_amount,omitempty" schema:"buyer_pay_amount,omitempty"` // 否 付款金额
		PointAmount    string `json:"point_amount,omitempty" url:"point_amount,omitempty" schema:"point_amount,omitempty"`             // 否 积分金额
		RefundFee      string `json:"refund_fee,omitempty" url:"refund_fee,omitempty" schema:"refund_fee,omitempty"`                   // 否 退款金额
		SendBackFee    string `json:"send_back_fee,omitempty" url:"send_back_fee,omitempty" schema:"send_back_fee,omitempty"`          // 否 实际退款
		Subject        string `json:"subject,omitempty" url:"subject,omitempty" schema:"subject,omitempty"`                            // 是 标题
		Body           string `json:"body,omitempty" url:"body,omitempty" schema:"body,omitempty"`                                     // 否 描述
		GmtCreate      string `json:"gmt_create,omitempty" url:"gmt_create,omitempty" schema:"gmt_create,omitempty"`                   // 否 创建时间
		GmtPayment     string `json:"gmt_payment,omitempty" url:"gmt_payment,omitempty" schema:"gmt_payment,omitempty"`                // 否 付款时间
		GmtRefund      string `json:"gmt_refund,omitempty" url:"gmt_refund,omitempty" schema:"gmt_refund,omitempty"`                   // 否 退款时间
		GmtClose       string `json:"gmt_close,omitempty" url:"gmt_close,omitempty" schema:"gmt_close,omitempty"`                      // 否 结束时间
	}
	// Query query
	Query struct {
		Response
		TradeNo             string `json:"trade_no,omitempty" url:"trade_no,omitempty"`                             // 是 交易号
		OutTradeNo          string `json:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                     // 是 单号
		BuyerLogonID        string `json:"buyer_logon_id,omitempty" url:"buyer_logon_id,omitempty"`                 // 是 买家帐号
		TradeStatus         string `json:"trade_status,omitempty" url:"trade_status,omitempty"`                     // 是 交易状态
		TotalAmount         string `json:"total_amount,omitempty" url:"total_amount,omitempty"`                     // 是 订单金额
		TransCurrency       string `json:"trans_currency,omitempty" url:"trans_currency,omitempty"`                 // 否 标价币种
		SettleCurrency      string `json:"settle_currency,omitempty" url:"settle_currency,omitempty"`               // 否 结算币种
		SettleAmount        string `json:"settle_amount,omitempty" url:"settle_amount,omitempty"`                   // 否 结算金额
		PayCurrency         string `json:"pay_currency,omitempty" url:"pay_currency,omitempty"`                     // 否 支付币种
		PayAmount           string `json:"pay_amount,omitempty" url:"pay_amount,omitempty"`                         // 否 支付金额
		SettleTransRate     string `json:"settle_trans_rate,omitempty" url:"settle_trans_rate,omitempty"`           // 否 结算汇率
		TransPayRate        string `json:"trans_pay_rate,omitempty" url:"trans_pay_rate,omitempty"`                 // 否 标价汇率
		BuyerPayAmount      string `json:"buyer_pay_amount,omitempty" url:"buyer_pay_amount,omitempty"`             // 否 付款金额
		PointAmount         string `json:"point_amount,omitempty" url:"point_amount,omitempty"`                     // 否 积分金额
		InvoiceAmount       string `json:"invoice_amount,omitempty" url:"invoice_amount,omitempty"`                 // 搜 开票金额
		SendPayDate         string `json:"send_pay_date,omitempty" url:"send_pay_date,omitempty"`                   // 是 打款时间
		ReceiptAmount       string `json:"receipt_amount,omitempty" url:"receipt_amount,omitempty"`                 // 是 实收金额
		StoreID             string `json:"store_id,omitempty" url:"store_id,omitempty"`                             // 否 门店编号
		TerminalID          string `json:"terminal_id,omitempty" url:"terminal_id,omitempty"`                       // 否 终端编号
		StoreName           string `json:"store_name,omitempty" url:"store_name,omitempty"`                         // 是 店铺名称
		BuyerUserID         string `json:"buyer_user_id,omitempty" url:"buyer_user_id,omitempty"`                   // 否 卖家用户
		ChargeAmount        string `json:"charge_amount,omitempty" url:"charge_amount,omitempty"`                   // 否 收费金额
		ChargeFlags         string `json:"charge_flags,omitempty" url:"charge_flags,omitempty"`                     // 否 收费标识
		SettlementID        string `json:"settlement_id,omitempty" url:"settlement_id,omitempty"`                   // 否 清算编号
		AuthTradePayMode    string `json:"auth_trade_pay_mode,omitempty" url:"auth_trade_pay_mode,omitempty"`       // 否 预授权支付
		BuyerUserType       string `json:"buyer_user_type,omitempty" url:"buyer_user_type,omitempty"`               // 否 买家类型
		MdiscountAmount     string `json:"mdiscount_amount,omitempty" url:"mdiscount_amount,omitempty"`             // 否 商家优惠
		DiscountAmount      string `json:"discount_amount,omitempty" url:"discount_amount,omitempty"`               // 否 平台优惠
		BuyerUserName       string `json:"buyer_user_name,omitempty" url:"buyer_user_name,omitempty"`               // 否 买家类型
		Subject             string `json:"subject,omitempty" url:"subject,omitempty"`                               // 否 订单标题
		Body                string `json:"body,omitempty" url:"body,omitempty"`                                     // 否 订单描述
		AlipaySubMerchantID string `json:"alipay_sub_merchant_id,omitempty" url:"alipay_sub_merchant_id,omitempty"` // 是 商家编号
		ExtInfos            string `json:"ext_infos,omitempty" url:"ext_infos,omitempty"`                           // 否 额外信息
	}
)
