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
func (t *Trade) Sign(args *url.Values, key string) (string, error) {
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return "", err
	}
	query = fmt.Sprintf("%s&key=%s", query, key)
	return strings.ToUpper(
		rsae.NewMD5().Encode(query),
	), nil
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
	body, err = fetch.Cmd(&fetch.Request{
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
	data, err := query.Values(args)
	if err != nil {
		return err
	}
	sign, err := t.Sign(&data, key)
	if err != nil {
		return err
	}
	if sign != args.Sign {
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
	body, err = fetch.Cmd(&fetch.Request{
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
