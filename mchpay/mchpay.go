package mchpay

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
		MchAppID       string   `xml:"mch_appid,omitempty" url:"mch_appid,omitempty"`               // 是 商户账号appid
		MchID          string   `xml:"mchid,omitempty" url:"mchid,omitempty"`                       // 是 商户号
		DeviceInfo     string   `xml:"device_info,omitempty" url:"device_info,omitempty"`           // 否 设备号
		NonceStr       string   `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`               // 是 随机字符串
		Sign           string   `xml:"sign,omitempty" url:"-"`                                      // 是 签名
		SpbillCreateIP string   `xml:"spbill_create_ip,omitempty" url:"spbill_create_ip,omitempty"` // 是 终端IP
		*Request
	}
	// Request request
	Request struct {
		PartnerTradeNo string `xml:"partner_trade_no,omitempty" url:"partner_trade_no,omitempty"` // 是 商户订单号
		OpenID         string `xml:"openid,omitempty" url:"openid,omitempty"`                     // 是 用户openid
		CheckName      string `xml:"check_name,omitempty" url:"check_name,omitempty"`             // 是 校验用户姓名选项（NO_CHECK：不校验真实姓名，FORCE_CHECK：强校验真实姓名）
		ReUserName     string `xml:"re_user_name,omitempty" url:"re_user_name,omitempty"`         // 否 收款用户姓名
		Amount         int    `xml:"amount,omitempty" url:"amount,omitempty"`                     // 是 金额
		Desc           string `xml:"desc,omitempty" url:"desc,omitempty"`                         // 是 描述
	}
	// Response response
	Response struct {
		XMLName    xml.Name `xml:"xml" url:"-"`                                         // 是 公共头部
		ReturnCode string   `xml:"return_code,omitempty" url:"return_code,omitempty"`   // 是 返回状态码
		ReturnMsg  string   `xml:"return_msg,omitempty" url:"return_msg,omitempty"`     // 否 返回信息
		ResultCode string   `xml:"result_code,omitempty" url:"result_code,omitempty"`   // 是 业务结果
		ErrCode    string   `xml:"err_code,omitempty" url:"err_code,omitempty"`         // 否 错误代码
		ErrCodeDes string   `xml:"err_code_des,omitempty" url:"err_code_des,omitempty"` // 否 错误代码

		MchAppID   string `xml:"mch_appid,omitempty" url:"mch_appid,omitempty"`     // 是 商户账号appid
		MchID      string `xml:"mchid,omitempty" url:"mchid,omitempty"`             // 是 商户号
		DeviceInfo string `xml:"device_info,omitempty" url:"device_info,omitempty"` // 否 设备号
		NonceStr   string `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`     // 是 随机字符串

		PartnerTradeNo string `xml:"partner_trade_no,omitempty" url:"partner_trade_no,omitempty"` // 是 商户订单号
		PaymentNo      string `xml:"payment_no,omitempty" url:"payment_no,omitempty"`             // 是 微信订单号
		PaymentTime    string `xml:"payment_time,omitempty" url:"payment_time,omitempty"`         // 是 微信支付成功时间
		DetailID       string `xml:"detail_id,omitempty" url:"detail_id,omitempty"`               // 是 付款单号
		Status         string `xml:"status,omitempty" url:"status,omitempty"`                     // 是 转账状态
		Reason         string `xml:"reason,omitempty" url:"reason,omitempty"`                     // 否 失败原因
		OpenID         string `xml:"openid,omitempty" url:"openid,omitempty"`                     // 是 用户标识
		TransferName   string `xml:"transfer_name,omitempty" url:"transfer_name,omitempty"`       // 否 收款用户姓名
		PaymentAmount  int    `xml:"payment_amount,omitempty" url:"trade_state,omitempty"`        // 是 付款金额
		TransferTime   string `xml:"transfer_time,omitempty" url:"transfer_time,omitempty"`       // 是 转账时间
		Desc           string `xml:"desc,omitempty" url:"desc,omitempty"`                         // 是 付款描述
	}
	// Query query
	Query struct {
		NonceStr       string `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`               // 是 随机字符串
		Sign           string `xml:"sign,omitempty" url:"-"`                                      // 是 签名
		PartnerTradeNo string `xml:"partner_trade_no,omitempty" url:"partner_trade_no,omitempty"` // 是 商户订单号
		MchID          string `xml:"mch_id,omitempty" url:"mch_id,omitempty"`                     // 是 商户号
		AppID          string `xml:"appid,omitempty" url:"appid,omitempty"`                       // 是 公众账号ID
	}
)
