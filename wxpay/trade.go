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
<<<<<<< HEAD
func (t *Trade) Sign(args *url.Values, key string) (string, error) {
	args.Add("key", key)
	query, err := url.QueryUnescape(args.Encode())
=======
func (t *Trade) Sign(args interface{}, key string) (string, error) {
	params, err := query.Values(args)
>>>>>>> fca77165ffb6a22d9a341d286b64cfb966b500e1
	if err != nil {
		return "", err
	}
	sign := rsae.NewMD5().Encode(query)
	return strings.ToUpper(sign), nil
}

// Prepay trade perpay
<<<<<<< HEAD
func (t *Trade) Prepay(args *Sign) (*Prepay, error) {
=======
func (t *Trade) Prepay(args Sign) (Prepay, error) {
>>>>>>> fca77165ffb6a22d9a341d286b64cfb966b500e1
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
<<<<<<< HEAD
func (t *Trade) Verify(args *Notice, key string) error {
	tmp, err := query.Values(args)
	if err != nil {
		return err
	}
	sign, err := t.Sign(&tmp, key)
=======
func (t *Trade) Verify(args Notice, key string) error {
	sign, err := t.Sign(args, key)
>>>>>>> fca77165ffb6a22d9a341d286b64cfb966b500e1
	if err != nil {
		return err
	}
	if args.Sign != sign {
		return errors.New("签名错误")
	}
	return nil
}

// Query query
<<<<<<< HEAD
func (t *Trade) Query(args *Sign) (*Query, error) {
=======
func (t *Trade) Query(args Sign) (Query, error) {
>>>>>>> fca77165ffb6a22d9a341d286b64cfb966b500e1
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
