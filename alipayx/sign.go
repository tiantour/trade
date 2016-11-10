package alipayx

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/rsae"
)

// signString
func (s *sign) signStr(args TradeIn) (string, error) {
	args.InputCharset = "utf-8"
	args.PaymentType = "1"
	signURL, err := query.Values(args)
	if err != nil {
		return "", err
	}
	sign, err := rsae.RSA.Sign(signURL.Encode())
	if err != nil {
		return "", err
	}
	return signURL.Encode() + "&sign=" + url.QueryEscape(sign) + "&sign_type=RSA", nil
}
