package mchpay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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

// Pay trade pay
func (t *Trade) Pay(args *Sign, certPath, keyPath string) (*Response, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/xml")
	header.Add("Content-Type", "application/xml;charset=utf-8")
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers",
		Body:   body,
		Header: header,
		Cert:   certPath,
		Key:    keyPath,
	})
	if err != nil {
		return nil, err
	}
	result := Response{}
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

// Query query
func (t *Trade) Query(args *Query, certPath, keyPath string) (*Response, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/xml")
	header.Add("Content-Type", "application/xml;charset=utf-8")
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo",
		Body:   body,
		Header: header,
		Cert:   certPath,
		Key:    keyPath,
	})
	if err != nil {
		return nil, err
	}
	result := Response{}
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
