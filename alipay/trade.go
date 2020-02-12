package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
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
	return args.Encode(), nil
}

// Prepay prepay
func (t *Trade) Prepay(args string) (*Prepay, error) {
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", args),
	})
	if err != nil {
		return nil, err
	}
	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	prepay := result.AlipayTradeCrateResponse
	if prepay.Code != "10000" {
		return nil, errors.New(prepay.Msg)
	}
	return result.AlipayTradeCrateResponse, nil
}

// Verify verify
func (t *Trade) Verify(args *Notice, publicPath string) error {
	tmp, err := query.Values(args)
	if err != nil {
		return err
	}
	str, err := url.QueryUnescape(tmp.Encode())
	if err != nil {
		return err
	}
	publicKey, err := imago.NewFile().Read(publicPath)
	if err != nil {
		return err
	}
	ok, err := rsae.NewRSA().Verify(str, args.Sign, publicKey)
	if !ok {
		return errors.New("签名错误")
	}
	return nil
}

// Query query
func (t *Trade) Query(args string) (*Query, error) {
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", args),
	})
	if err != nil {
		return nil, err
	}
	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	query := result.AlipayTradeQueryResponse
	if query.Code != "10000" {
		return nil, errors.New(query.Msg)
	}
	return query, nil
}
