package wxpayx

import (
	"strconv"

	"github.com/tiantour/conf"
	"github.com/tiantour/imago"
	"github.com/tiantour/tempo"
)

// Sign 统一下单
func (w *Wap) Sign(args TradeIn) (TradeOut, error) {
	args.AppID = conf.Data.Wxpay.AppID // appid
	args.MchID = conf.Data.Wxpay.MchID // 商户id
	trade, err := w.tradeSign(args)
	if err != nil {
		return TradeOut{}, err
	}
	prepayID, err := Sign.prepayID(trade)
	if err != nil {
		return TradeOut{}, err
	}
	return w.defraySign(prepayID)
}

// tradeSign
func (w *Wap) tradeSign(args TradeIn) (signIn, error) {
	result := signIn{
		TradeIn:   args,
		TradeType: "JSAPI",
	}
	sign, err := Sign.signStr(result, conf.Data.Wxpay.Key)
	if err != nil {
		return signIn{}, err
	}
	result.Sign = sign
	return result, err
}

// defraySign 调起签名
func (w *Wap) defraySign(prepayID string) (TradeOut, error) {
	result := TradeOut{
		AppID:     conf.Data.Wxpay.AppID,
		Package:   "prepay_id=" + prepayID,
		NonceStr:  imago.Gen.String(16),
		TimeStamp: strconv.FormatInt(tempo.Now.Unix(), 10),
		SignType:  "MD5",
	}
	sign, err := Sign.signStr(result, conf.Data.Wxpay.Key)
	if err != nil {
		return TradeOut{}, err
	}
	result.PaySign = sign
	return result, nil
}
