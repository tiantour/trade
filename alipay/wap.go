package alipay

// Sign Wap
func (w Wap) Sign(orderID, packageName, optionName, notifyURL, showURL, returnURL string, totalFee float64) (string, error) {
	signMap := Sign.signMap(orderID, packageName, optionName, notifyURL, totalFee)
	signMap["service"] = "alipay.wap.create.direct.pay.by.user"
	signMap["show_url"] = showURL
	signMap["return_url"] = returnURL
	return Sign.signString(signMap)
}
