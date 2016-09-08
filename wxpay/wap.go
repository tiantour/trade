package wxpay

import (
	"encoding/xml"
	"strconv"

	"github.com/tiantour/conf"
	"github.com/tiantour/imago"
)

// UnifiedOrder 统一下单
func (w Wap) UnifiedOrder(packageName, optionName, optionID, orderID, tradeIP, openID, notifyURL string, totalFee float64) (map[string]interface{}, error) {
	nonceStr := imago.G.String(16)
	orderData, err := w.orderData(
		nonceStr,
		packageName,
		optionName,
		optionID,
		orderID,
		tradeIP,
		openID,
		notifyURL,
		int(totalFee*100),
	)
	if err != nil {
		return nil, err
	}
	prepayID, err := UnifiedOrder.prepayID(orderData)
	if err != nil {
		return nil, err
	}
	return w.defraySign(prepayID), nil
}

// orderSign 生成签名
func (w Wap) orderSign(nonceStr, body, detail, productID, outTradeNo, spbillCreateIP, openID, notifyURL string, totalFee int) string {
	orderMap := UnifiedOrder.orderMap(
		nonceStr,
		body,
		detail,
		productID,
		outTradeNo,
		spbillCreateIP,
		notifyURL,
		totalFee,
	)
	orderMap["trade_type"] = "JSAPI"
	orderMap["openid"] = openID
	return UnifiedOrder.signString(orderMap)

}

// orderData XML
func (w Wap) orderData(nonceStr, body, detail, productID, outTradeNo, spbillCreateIP, openID, notifyURL string, totalFee int) ([]byte, error) {
	orderData := UnifiedOrder.orderXML(
		nonceStr,
		body,
		detail,
		productID,
		outTradeNo,
		spbillCreateIP,
		notifyURL,
		totalFee,
	)
	orderData.OpenID = openID
	orderData.TradeType = "JSAPI"
	orderData.Sign = w.orderSign(
		nonceStr,
		body,
		detail,
		productID,
		outTradeNo,
		spbillCreateIP,
		openID,
		notifyURL,
		totalFee,
	)
	return xml.Marshal(orderData)
}

// defraySign 调起签名
func (w Wap) defraySign(prepayID string) map[string]interface{} {
	defrayMap := map[string]interface{}{
		"appId":     conf.Options.Wxpay.AppID,
		"timeStamp": strconv.FormatInt(imago.Time.NowToUnix(), 10),
		"nonceStr":  imago.G.String(16),
		"package":   "prepay_id=" + prepayID,
		"signType":  "MD5",
	}
	defrayMap["paySign"] = UnifiedOrder.signString(defrayMap)
	return defrayMap
}
