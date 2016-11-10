package alipayx

// Sign APP
func (a *App) Sign(args TradeIn) (string, error) {
	args.Service = "mobile.securitypay.pay"
	return Sign.signStr(args)
}
