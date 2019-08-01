package wxpay

import (
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/fetch"
	"github.com/tiantour/rsae"
)

// Trade trade
type Trade struct{}

// NewTrade new trade
func NewTrade() *Trade {
	return &Trade{}
}

// Sign trade sign
func (t *Trade) Sign(args *url.Values, key string) (string, error) {
	args.Add("key", key)
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return "", err
	}
	sign := rsae.NewMD5().Encode(query)
	return strings.ToUpper(sign), nil
}

// Prepay trade perpay
func (t *Trade) Prepay(args *Sign) (*Prepay, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/xml")
	header.Add("Content-Type", "application/xml;charset=utf-8")
	body, err = fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    "https://api.mch.weixin.qq.com/pay/unifiedorder",
		Body:   body,
		Header: header,
	})
	if err != nil {
		return nil, err
	}
	result := Prepay{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ReturnCode != Success {
		return nil, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != Success {
		return nil, errors.New(result.ErrCodeDes)
	}
	return &result, nil
}

// Verify verify
func (t *Trade) Verify(args *Notice, key string) error {
	tmp, err := query.Values(args)
	if err != nil {
		return err
	}
	sign, err := t.Sign(&tmp, key)
	if err != nil {
		return err
	}
	if args.Sign != sign {
		return errors.New("签名错误")
	}
	return nil
}

// Query query
func (t *Trade) Query(args *Sign) (*Query, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/xml")
	header.Add("Content-Type", "application/xml;charset=utf-8")
	body, err = fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    "https://api.mch.weixin.qq.com/pay/orderquery",
		Body:   body,
		Header: header,
	})
	if err != nil {
		return nil, err
	}
	result := Query{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ReturnCode != Success {
		return nil, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != Success {
		return nil, errors.New(result.ErrCodeDes)
	}
	return &result, nil
}

// Refund refund
func (t Trade) Refund(args *Sign) (*Refund, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/xml")
	header.Add("Content-Type", "application/xml;charset=utf-8")
	body, err = fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    "https://api.mch.weixin.qq.com/pay/refund",
		Body:   body,
		Header: header,
	})
	if err != nil {
		return nil, err
	}
	result := Refund{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ReturnCode != Success {
		return nil, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != Success {
		return nil, errors.New(result.ErrCodeDes)
	}
	return &result, nil
}
