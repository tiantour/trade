package wxpay

import (
	"encoding/xml"
	"errors"
	"fmt"
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
func (t Trade) Sign(args interface{}, key string) (string, error) {
	params, err := query.Values(args)
	if err != nil {
		return "", err
	}
	query, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return "", err
	}
	sign := rsae.NewMD5().Encode(fmt.Sprintf("%s&key=%s",
		query,
		key,
	))
	return strings.ToUpper(sign), nil
}

// Prepay trade perpay
func (t Trade) Prepay(args Sign) (Prepay, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return Prepay{}, err
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
		return Prepay{}, err
	}
	result := Prepay{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	if result.ReturnCode != Success {
		return result, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != Success {
		return result, errors.New(result.ErrCodeDes)
	}
	return result, nil
}

// Verify verify
func (t Trade) Verify(args Notice, key string) error {
	sign, err := t.Sign(args, key)
	if err != nil {
		return err
	}
	if args.Sign != sign {
		return errors.New("签名错误")
	}
	return nil
}

// Query query
func (t Trade) Query(args Sign) (Query, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return Query{}, err
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
		return Query{}, err
	}
	result := Query{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	if result.ReturnCode != Success {
		return result, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != Success {
		return result, errors.New(result.ErrCodeDes)
	}
	return result, nil
}

// Refund refund
func (t Trade) Refund(args Sign) (Refund, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return Refund{}, err
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
		return Refund{}, err
	}
	result := Refund{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	if result.ReturnCode != Success {
		return result, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != Success {
		return result, errors.New(result.ErrCodeDes)
	}
	return result, nil
}
