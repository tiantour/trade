package alipay

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/imago"
	"github.com/tiantour/rsae"
)

// Trade trade
type Trade struct{}

// NewTrade new trade
func NewTrade() *Trade {
	return &Trade{}
}

// Sign trade sign
func (t Trade) Sign(args interface{}, privatePath string) (string, error) {
	params, err := query.Values(args)
	if err != nil {
		return "", err
	}
	query, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return "", err
	}
	privateKey, err := imago.NewFile().Read(privatePath)
	if err != nil {
		return "", err
	}
	sign, err := rsae.NewRsae().Sign(query, privateKey)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s&sign=%s",
		query,
		url.QueryEscape(sign),
	), nil
}

// Verify verify
func (t Trade) Verify(args url.Values, publicPath string) error {
	sign := args.Get("sign")
	args.Del("sign")
	args.Del("sign_type")
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return err
	}
	publicKey, err := imago.NewFile().Read(publicPath)
	if err != nil {
		return err
	}
	ok, err := rsae.NewRsae().Verify(query, sign, publicKey)
	if !ok {
		return errors.New("签名错误")
	}
	return nil
}
