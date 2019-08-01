package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/tiantour/fetch"
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
func (t *Trade) Sign(args *url.Values, privatePath string) (string, error) {
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return "", err
	}
	privateKey, err := imago.NewFile().Read(privatePath)
	if err != nil {
		return "", err
	}
	sign, err := rsae.NewRSA().Sign(query, privateKey)
	if err != nil {
		return "", err
	}
	args.Add("sign", sign)
	return url.QueryEscape(args.Encode()), nil
}

// Verify verify
func (t *Trade) Verify(args *url.Values, publicPath string) error {
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
	ok, err := rsae.NewRSA().Verify(query, sign, publicKey)
	if !ok {
		return errors.New("签名错误")
	}
	return nil
}

// Query query
func (t *Trade) Query(str string) (*Query, error) {
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", str),
	})
	if err != nil {
		return nil, err
	}
	result := Query{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != "10000" {
		return nil, errors.New(result.Msg)
	}
	return &result, nil
}

// Refund refund
func (t *Trade) Refund(str string) (*Query, error) {
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", str),
	})
	if err != nil {
		return nil, err
	}
	result := Query{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != "10000" {
		return nil, errors.New(result.Msg)
	}
	return &result, nil
}
