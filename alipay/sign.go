package alipay

import (
	"net/url"

	"github.com/tiantour/conf"
	"github.com/tiantour/imago"
	"github.com/tiantour/rsae"
)

// signMap
func (s *sign) signMap(orderID, packageName, optionName, notifyURL string, totalFee float64) map[string]interface{} {
	return map[string]interface{}{
		"partner":        conf.Options.Alipay.Partner, //合作者身份ID，必填
		"_input_charset": "utf-8",                     //参数编码字符集，必填
		"notify_url":     notifyURL,                   //服务器异步通知页面路径，必填
		"out_trade_no":   orderID,                     //商户网站唯一订单号,必填
		"subject":        packageName,                 //商品名称,必填
		"payment_type":   "1",                         //支付类型,必填
		"seller_id":      conf.Options.Alipay.Partner, //卖家支付宝账号,必填
		"total_fee":      totalFee,                    //总金额,必填
		"body":           optionName,                  //商品详情,必填
	}
}

// signString
func (s *sign) signString(data map[string]interface{}) (string, error) {
	signURL := imago.Convert.MapToURL(data)
	sign, err := rsae.RSA.Sign(signURL)
	if err != nil {
		return "", err
	}
	signString := signURL + "&sign=" + url.QueryEscape(sign) + "&sign_type=RSA"
	return signString, err
}
