package wxpay

import (
	"encoding/xml"
	"time"

	"github.com/tiantour/conf"
	"github.com/tiantour/imago"
)

// UnifiedOrder 统一下单
func (a App) UnifiedOrder(packageName, optionName, optionID, orderID, tradeIP, notifyURL string, totalFee float64) (map[string]interface{}, error) {
	nonceStr := imago.G.String(16)
	orderData, err := a.orderData(
		nonceStr,
		packageName,
		optionName,
		optionID,
		orderID,
		tradeIP,
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
	return a.defraySign(prepayID), nil
}

// orderSign 生成签名
func (a App) orderSign(nonceStr, body, detail, productID, outTradeNo, spbillCreateIP, notifyURL string, totalFee int) string {
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
	orderMap["trade_type"] = "APP"
	return UnifiedOrder.signString(orderMap)

}

// orderData XML
func (a App) orderData(nonceStr, body, detail, productID, outTradeNo, spbillCreateIP, notifyURL string, totalFee int) ([]byte, error) {
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
	orderData.TradeType = "APP"
	orderData.Sign = a.orderSign(
		nonceStr,
		body,
		detail,
		productID,
		outTradeNo,
		spbillCreateIP,
		notifyURL,
		totalFee,
	)
	return xml.Marshal(orderData)
}

// sign 调起签名
func (a App) defraySign(prepayID string) map[string]interface{} {
	defrayMap := map[string]interface{}{
		"appid":     conf.Options.Wxpay.AppID,
		"partnerid": conf.Options.Wxpay.MchID,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  imago.G.String(16),
		"timestamp": time.Now().Unix(),
	}
	defrayMap["sign"] = UnifiedOrder.signString(defrayMap)
	return defrayMap
}
