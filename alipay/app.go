package alipay

// Sign APP
func (a App) Sign(orderID, packageName, optionName, notifyURL string, totalFee float64) (string, error) {
	signMap := Sign.signMap(orderID, packageName, optionName, notifyURL, totalFee)
	signMap["service"] = "mobile.securitypay.pay"
	return Sign.signString(signMap)
}
