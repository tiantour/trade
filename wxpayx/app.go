package wxpayx

import (
	"strconv"

	"github.com/tiantour/conf"
	"github.com/tiantour/imago"
	"github.com/tiantour/tempo"
)

// Sign 统一下单
func (a *App) Sign(args TradeIn) (TradeOut, error) {
	args.AppID = conf.Data.Wxpay.AppID // appid
	args.MchID = conf.Data.Wxpay.MchID // 商户id
	trade, err := a.tradeSign(args)
	if err != nil {
		return TradeOut{}, err
	}
	prepayID, err := Sign.prepayID(trade)
	if err != nil {
		return TradeOut{}, err
	}
	return a.defraySign(prepayID)
}

// tradeSign
func (a *App) tradeSign(args TradeIn) (signIn, error) {
	result := signIn{
		TradeIn:   args,
		TradeType: "APP",
	}
	sign, err := Sign.signStr(result, conf.Data.Wxpay.Key)
	if err != nil {
		return signIn{}, err
	}
	result.Sign = sign
	return result, err
}

// defraySign
func (a *App) defraySign(prepayID string) (TradeOut, error) {
	result := TradeOut{
		AppID:     conf.Data.Wxpay.AppID,
		PartnerID: conf.Data.Wxpay.MchID,
		PrepayID:  prepayID,
		Package:   "Sign=WXPay",
		NonceStr:  imago.Gen.String(16),
		TimeStamp: strconv.FormatInt(tempo.Now.Unix(), 10),
	}
	sign, err := Sign.signStr(result, conf.Data.Wxpay.Key)
	if err != nil {
		return TradeOut{}, err
	}
	result.Sign = sign
	return result, nil
}
