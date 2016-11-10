package alipayx

// alipay alipay
var (
	Sign   = &sign{}
	Notify = &notify{}
)

type (
	// App app
	App struct {
		notify
	}
	// Wap wap
	Wap struct {
		notify
	}
	// TradeIn tradein
	TradeIn struct {
		Partner      string  `json:"partner" url:"partner"`                 //合作者身份ID，必填
		InputCharset string  `json:"_input_charset",url:"_input_charset"`   //参数编码字符集，必填
		NotifyURL    string  `json:"notify_url",url:"notify_url"`           //服务器异步通知页面路径，必填
		OutTradeNo   string  `json:"out_trade_no" url:"out_trade_no"`       //商户网站唯一订单号,必填
		Subject      string  `json:"subject" url:"subject"`                 //商品名称,必填
		PaymentType  string  `json:"payment_type" url:"payment_type"`       //支付类型,必填
		SellerID     string  `json:"seller_id" url:"seller_id"`             //卖家支付宝账号,必填
		TotalFee     float64 `json:"total_fee" url:"total_fee"`             //总金额,必填
		Body         string  `json:"body" url:"body"`                       //商品详情,必填
		Service      string  `json:"service" url:"service"`                 //支付渠道，app或h5
		ShowURL      string  `json:"show_url" url:"show_url,omitempty"`     // 展示地址，h5必填
		ReturnURL    string  `json:"return_url" url:"return_url,omitempty"` // 回调地址，h5必填
	}
	notify struct{}
	sign   struct{}
)
