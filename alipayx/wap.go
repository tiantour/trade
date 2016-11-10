package alipayx

// Sign Wap
func (w *Wap) Sign(args TradeIn) (string, error) {
	args.Service = "alipay.wap.create.direct.pay.by.user"
	return Sign.signStr(args)
}
