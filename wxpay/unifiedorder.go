package wxpay

import (
	"encoding/xml"
	"errors"
	"strings"

	"github.com/tiantour/conf"
	"github.com/tiantour/imago"
	"github.com/tiantour/requests"
)

// orderMap
func (u *unifiedorder) orderMap(nonceStr, body, detail, productID, outTradeNo, spbillCreateIP, notifyURL string, totalFee int) map[string]interface{} {
	return map[string]interface{}{
		"appid":            conf.Options.Wxpay.AppID,
		"mch_id":           conf.Options.Wxpay.MchID,
		"nonce_str":        nonceStr,
		"body":             body,
		"detail":           detail,
		"product_id":       productID,
		"out_trade_no":     outTradeNo,
		"total_fee":        totalFee,
		"spbill_create_ip": spbillCreateIP,
		"notify_url":       notifyURL,
	}
}

// orderXML
func (u *unifiedorder) orderXML(nonceStr, body, detail, productID, outTradeNo, spbillCreateIP, notifyURL string, totalFee int) UnifiedOrderInput {
	return UnifiedOrderInput{
		AppID:          conf.Options.Wxpay.AppID, // 公众账号ID，必填
		MchID:          conf.Options.Wxpay.MchID, // 商户号，必填
		NonceStr:       nonceStr,                 // 随机字符串，必填
		Body:           body,                     // 商品描述，必填
		Detail:         detail,                   // 商品详情，非必填
		ProductID:      productID,                // 商品ID,非必填
		OutTradeNo:     outTradeNo,               // 商户订单号，必填
		TotalFee:       totalFee,                 // 总金额，必填
		SpbillCreateIP: spbillCreateIP,           // 终端IP，必填
		NotifyURL:      notifyURL,                // 通知地址,必填
	}
}

// prepayID
func (u *unifiedorder) prepayID(orderXML []byte) (string, error) {
	requestURL, requestBody, requestHeader := requests.Options()
	requestURL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	requestHeader.Add("Accept", "application/xml")
	requestHeader.Add("Content-Type", "application/xml;charset=utf-8")
	requestBody, err := requests.Post(requestURL, orderXML, requestHeader)
	if err != nil {
		return "", err
	}
	orderMap := UnifiedOrderOutput{}
	err = xml.Unmarshal(requestBody, &orderMap)
	if err != nil {
		return "", nil
	}
	if orderMap.ReturnCode != success {
		return "", errors.New(orderMap.ReturnMsg)
	}
	if orderMap.ResultCode != success {
		return "", errors.New(orderMap.ErrCodeDes)
	}
	return orderMap.PrepayID, nil
}

// signString
func (u *unifiedorder) signString(data map[string]interface{}) string {
	signURL := imago.Convert.MapToURL(data)
	signString := imago.Crypto.Md532(signURL + "&key=" + conf.Options.Wxpay.Key)
	return strings.ToUpper(signString)
}
