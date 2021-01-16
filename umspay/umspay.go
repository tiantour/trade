package umspay

const (
	// Success status success
	Success = "SUCCESS"
	// FAILED status failed
	FAILED = "FAILED"
)

var (
	// AccessToken access token
	AccessToken = ""
)

type (
	// Request request
	Request struct {
		MsgID             string   `json:"msgId,omitempty"`             // 否 消息ID
		RequestTimestamp  string   `json:"requestTimestamp,omitempty"`  // 是 报文请求时间 格式yyyy-MM-dd HH:mm:ss
		MerOrderID        string   `json:"merOrderId,omitempty"`        // 是 商户订单号
		SrcReserve        string   `json:"srcReserve,omitempty"`        // 否 请求系统预留字段
		MID               string   `json:"mid,omitempty"`               // 是 商户号
		TID               string   `json:"tid,omitempty"`               // 是 机构商户号
		Goods             []*Goods `json:"goods,omitempty"`             // 否 商品信息
		AttachedData      string   `json:"attachedData,omitempty"`      // 否 商户附加数据
		ExpireTime        string   `json:"expireTime,omitempty"`        // 否 订单过期时间
		GoodsTag          string   `json:"GoodsTag,omitempty"`          // 否 商品标记
		GoodsTradeNo      string   `json:"goodsTradeNo,omitempty"`      // 否 商品交易单号
		OrderDesc         string   `json:"orderDesc,omitempty"`         // 否 账单描述
		OriginalAmount    int      `json:"originalAmount,omitempty"`    // 否 订单原始金额
		ProductID         string   `json:"productId,omitempty"`         // 否 商品ID
		TotalAmount       int      `json:"totalAmount,omitempty"`       // 否 支付总金额
		DivisionFlag      bool     `json:"divisionFlag,omitempty"`      // 否 分账标记
		PlatformAmount    int      `json:"platformAmount,omitempty"`    // 否 平台商户分账金额
		SubOrders         *Order   `json:"subOrders,omitempty"`         // 否 子订单信息
		NotifyURL         string   `json:"notifyUrl,omitempty"`         // 否 支付结果通知地址
		ReturnURL         string   `json:"returnUrl,omitempty"`         // 否 网页跳转地址
		ShowURL           string   `json:"showUrl,omitempty"`           // 否 订单展示页面
		SecureTransaction string   `json:"secureTransaction,omitempty"` // 否 担保交易标识
		SubAppID          string   `json:"subAppId,omitempty"`          // 是 微信子商户
		SubOpenID         string   `json:"subOpenId,omitempty"`         // 是 用户子标识，微信必传
		UserID            string   `json:"userId,omitempty"`            // 是 用户子标识，支付宝必传
		TradeType         string   `json:"tradeType,omitempty"`         // 是 交易类型，值为MINI
		LimitCreditCard   string   `json:"limitCreditCard,omitempty"`   // 否 是否需要限制信用卡支付，：true或false，默认false
		InstallmentNumber string   `json:"installmentNumber,omitempty"` // 否 花呗分期数，仅支持3、6、12
	}

	// Response response
	Response struct {
		ErrCode          string          `json:"errCode,omitempty"`          // 否 平台错误码
		ErrMsg           string          `json:"errMsg,omitempty"`           // 否 平台错误信息
		ErrInfo          string          `json:"errInfo,omitempty"`          // 否 平台错误信息
		MsgID            string          `json:"msgId,omitempty"`            // 否 消息ID
		SrcReserve       string          `json:"srcReserve,omitempty"`       // 否 请求系统预留字段
		RequestTimestamp string          `json:"requestTimestamp,omitempty"` // 是 报文请求时间 格式yyyy-MM-dd HH:mm:ss
		MerName          string          `json:"merName,omitempty"`          // 否 商户名称
		MerOrderID       string          `json:"merOrderId,omitempty"`       // 否 商户订单号
		MID              string          `json:"mid,omitempty"`              // 是 商户号
		TID              string          `json:"tid,omitempty"`              // 是 机构商户号
		SeqID            string          `json:"seqId,omitempty"`            // 否 平台流水号
		SettleRefID      string          `json:"settleRefId,omitempty"`      // 否 清分ID 字符串
		Status           string          `json:"status,omitempty"`           // 否 交易状态
		TotalAmount      int             `json:"totalAmount,omitempty"`      // 否 支付总金额
		TargetOrderID    string          `json:"targetOrderId,omitempty"`    // 否 第三方订单号
		TargetSys        string          `json:"targetSys,omitempty"`        // 否 目标平台代码
		TargetStatus     string          `json:"targetStatus,omitempty"`     // 否 目标平台的状态
		MiniPayRequest   *MiniPayRequest `json:"miniPayRequest,omitempty"`   // 否 小程序支付用的请求报文，带有签名信息
		TargetMID        string          `json:"targetMid,omitempty"`        // 否 支付渠道商户号
		InstMID          string          `json:"instMid,omitempty"`          // 否 业务类型
		RefID            string          `json:"refId,omitempty"`            // 否 检索参考号
		BuyerID          string          `json:"buyerId,omitempty"`          // 否 买家ID
		BankCardNo       string          `json:"bankCardNo,omitempty"`       // 否 银行卡号
		BankInfo         string          `json:"bankInfo,omitempty"`         // 否 银行信息
		BillFunds        string          `json:"billFunds,omitempty"`        // 否 支付渠道列表
		BillFundsDesc    string          `json:"billFundsDesc,omitempty"`    // 否 支付渠道描述
		BuyerPayAmount   int             `json:"buyerPayAmount,omitempty"`   // 否 买家付款的金额
		BuyerUsername    string          `json:"buyerUsername,omitempty"`    // 否 买家用户名
		CouponAmount     int             `json:"couponAmount,omitempty"`     // 否 网付计算的优惠金额
		InvoiceAmount    int             `json:"invoiceAmount,omitempty"`    // 否 交易中可给用户开具发票的金额
		PayTime          string          `json:"payTime,omitempty"`          // 否 支付时间
		ReceiptAmount    int             `json:"receiptAmount,omitempty"`    // 否 商户实收金额
		SettleDate       string          `json:"settleDate,omitempty"`       // 否 结算日期
		SubBuyerID       string          `json:"subBuyerId,omitempty"`       // 否 子买家ID
		ActivityIDs      string          `json:"activityIds,omitempty"`      // 否 微信活动ID
		YXLMAmount       int             `json:"yxlmAmount,omitempty"`       // 否 营销联盟优惠金额
	}

	// MiniPayRequest minipay request
	MiniPayRequest struct {
		TimeStamp string `json:"timeStamp,omitempty"` // 时间戳
		Package   string `json:"package,omitempty"`   // 预支付
		PaySign   string `json:"paySign,omitempty"`   // 签名
		AppID     string `json:"appId,omitempty"`     // 小程序id
		SignType  string `json:"signType,omitempty"`  // 签名类型
		NonceStr  string `json:"nonceStr,omitempty"`  // 随机字符串
	}

	// Goods goods
	Goods struct {
		GoodsID        string `json:"goodsId,omitempty"`        // 否 商品ID
		GoodsName      string `json:"goodsName,omitempty"`      // 否 商品名称
		Quantity       int    `json:"quantity,omitempty"`       // 否 商品数量 数字型
		Price          int    `json:"price,omitempty"`          // 否  商品单价，单位分
		GoodsCategory  string `json:"goodsCategory,omitempty"`  // 否 商品分类
		Body           string `json:"body,omitempty"`           // 否 商品说明 字符串 ⇐1024
		Unit           string `json:"unit,omitempty"`           // 否 商品单位
		Discount       int    `json:"discount,omitempty"`       // 否 商品折扣
		SubMerchantID  string `json:"subMerchantId,omitempty"`  // 否 子商户号
		MerOrderID     string `json:"merOrderId,omitempty"`     // 否 商户子订单号
		SubOrderAmount int    `json:"subOrderAmount,omitempty"` // 否 子商户商品总额
	}

	// Order suborder
	Order struct {
		MID         string `json:"mid,omitempty"`         // 否 子商户号
		MerOrderID  string `json:"merOrderId,omitempty"`  // 否 商户子订单号
		TotalAmount int    `json:"totalAmount,omitempty"` // 否 子商户分账金额
	}

	// Notice notice
	Notice struct {
		MID                      string `schema:"mid,omitempty" url:"mid,omitempty"`                                           // 是 商户号
		TID                      string `schema:"tid,omitempty" url:"tid,omitempty"`                                           // 是 机构商户号
		InstMID                  string `schema:"instMid,omitempty" url:"instMid,omitempty"`                                   // 否 业务类型
		MerName                  string `schema:"merName,omitempty" url:"merName,omitempty"`                                   // 否 商户名称
		AttachedData             string `schema:"attachedData,omitempty" url:"attachedData,omitempty"`                         // 否 商户附加数据
		BankCardNo               string `schema:"bankCardNo,omitempty" url:"bankCardNo,omitempty"`                             // 否 银行卡号
		BankInfo                 string `schema:"bankInfo,omitempty" url:"bankInfo,omitempty"`                                 // 否 银行信息
		BillFunds                string `schema:"billFunds,omitempty" url:"billFunds,omitempty"`                               // 否 支付渠道列表
		BillFundsDesc            string `schema:"billFundsDesc,omitempty" url:"billFundsDesc,omitempty"`                       // 否 支付渠道描述
		BuyerID                  string `schema:"buyerId,omitempty" url:"buyerId,omitempty"`                                   // 否 买家ID
		BuyerPayAmount           int    `schema:"buyerPayAmount,omitempty" url:"buyerPayAmount,omitempty"`                     // 否 买家付款的金额
		BuyerUsername            string `schema:"buyerUsername,omitempty" url:"buyerUsername,omitempty"`                       // 否 买家用户名
		BuyerCashPayAmt          string `schema:"buyerCashPayAmt,omitempty" url:"buyerCashPayAmt,omitempty"`                   // 否 买家现金金额
		TotalAmount              int    `schema:"totalAmount,omitempty" url:"totalAmount,omitempty"`                           // 否 子商户分账金额
		InvoiceAmount            int    `schema:"invoiceAmount,omitempty" url:"invoiceAmount,omitempty"`                       // 否 交易中可给用户开具发票的金额
		MerOrderID               string `schema:"merOrderId,omitempty" url:"merOrderId,omitempty"`                             // 否 商户订单号
		PayTime                  string `schema:"payTime,omitempty" url:"payTime,omitempty"`                                   // 否 支付时间
		ReceiptAmount            int    `schema:"receiptAmount,omitempty" url:"receiptAmount,omitempty"`                       // 否 商户实收金额
		RefID                    string `schema:"refId,omitempty" url:"refId,omitempty"`                                       // 否 检索参考号
		RefundAmount             int    `schema:"refundAmount,omitempty" url:"refundAmount,omitempty"`                         // 否 退款金额
		RefundDesc               string `schema:"refundDesc,omitempty" url:"refundDesc,omitempty"`                             // 否 退款说明
		SeqID                    string `schema:"seqId,omitempty" url:"seqId,omitempty"`                                       // 否 平台流水号
		SettleDate               string `schema:"settleDate,omitempty" url:"settleDate,omitempty"`                             // 否 结算日期
		Status                   string `schema:"status,omitempty" url:"status,omitempty"`                                     // 否 交易状态
		SubBuyerID               string `schema:"subBuyerId,omitempty" url:"subBuyerId,omitempty"`                             // 否 子买家ID
		TargetOrderID            string `schema:"targetOrderId,omitempty" url:"targetOrderId,omitempty"`                       // 否 第三方订单号
		TargetSys                string `schema:"targetSys,omitempty" url:"targetSys,omitempty"`                               // 否 目标平台代码
		Sign                     string `schema:"sign,omitempty" url:"-"`                                                      // 否 签名
		SignType                 string `schema:"signType,omitempty" url:"signType,omitempty"`                                 // 签名类型
		CouponMerchantContribute int    `schema:"couponMerchantContribute,omitempty" url:"couponMerchantContribute,omitempty"` // 否 商户出资优惠金额
		CouponOtherContribute    int    `schema:"couponOtherContribute,omitempty" url:"couponOtherContribute,omitempty"`       // 否 其他出资优惠金额
		ActivityIDs              string `schema:"activityIds,omitempty" url:"activityIds,omitempty"`                           // 否 微信活动ID
		RefundTargetOrderID      string `schema:"refundTargetOrderId,omitempty" url:"refundTargetOrderId,omitempty"`           // 否 退货渠道订单号
		RefundPayTime            string `schema:"refundPayTime,omitempty" url:"refundPayTime,omitempty"`                       // 否 退货时间
		RefundSettleDate         string `schema:"refundSettleDate,omitempty" url:"refundSettleDate,omitempty"`                 // 否 结算日期
		OrderDesc                string `schema:"orderDesc,omitempty" url:"orderDesc,omitempty"`                               // 否 账单描述
		CreateTime               string `schema:"createTime,omitempty" url:"createTime,omitempty"`                             // 否 订单创建时间
		MchntUUID                string `schema:"mchntUuid,omitempty" url:"mchntUuid,omitempty"`                               // 否 商户UUID
		ConnectSys               string `schema:"connectSys,omitempty" url:"connectSys,omitempty"`                             // 否 转接系统
		SubInst                  string `schema:"subInst,omitempty" url:"subInst,omitempty"`                                   // 否 商户所属分支机构代码
		YXLMAmount               int    `schema:"yxlmAmount,omitempty" url:"yxlmAmount,omitempty"`                             // 否 营销联盟优惠金额
		RefundExtOrderID         string `schema:"refundExtOrderId,omitempty" url:"refundExtOrderId,omitempty"`                 // 否 退货外部订单号
		GoodsTradeNo             string `schema:"goodsTradeNo,omitempty" url:"goodsTradeNo,omitempty"`                         // 否 商品交易单号
		ExtOrderID               string `schema:"ExtOrderId,omitempty" url:"ExtOrderId,omitempty"`                             // 否 外部订单号
		SecureStatus             string `schema:"secureStatus,omitempty" url:"secureStatus,omitempty"`                         // 否 担保交易状态
		CompleteAmount           string `schema:"completeAmount,omitempty" url:"completeAmount,omitempty"`                     // 否 担保完成金额
		RefundOrderID            string `schema:"refundOrderId,omitempty" url:"refundOrderId,omitempty"`                       // 否 退货订单号
		CouponAmount             int    `schema:"couponAmount,omitempty" url:"couponAmount,omitempty"`                         // 否 网付计算的优惠金额
		Bi                       string `schema:"Bi,omitempty" url:"Bi,omitempty"`                                             // 否 Bi
		NotifyID                 string `schema:"notifyId,omitempty" url:"notifyId,omitempty"`                                 // 否 通知编号
	}

	// Query query
	Query struct {
		MsgID            string `json:"msgId,omitempty"`            // 否 消息ID
		RequestTimestamp string `json:"requestTimestamp,omitempty"` // 是 报文请求时间 格式yyyy-MM-dd HH:mm:ss
		SrcReserve       string `json:"srcReserve,omitempty"`       // 否 请求系统预留字段
		MID              string `json:"mid,omitempty"`              // 是 商户号
		TID              string `json:"tid,omitempty"`              // 是 机构商户号
		InstMID          string `json:"instMid,omitempty"`          // 否 业务类型
		MerOrderID       string `json:"merOrderId,omitempty"`       // 否 商户订单号
		TargetOrderID    string `json:"targetOrderId,omitempty"`    // 否 支付订单号
	}
)
