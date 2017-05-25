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
	fmt.Println(query)
	sign := rsae.NewRsae().Md532(
		fmt.Sprintf("%s&key=%s",
			query,
			key,
		),
	)
	return strings.ToUpper(sign), nil
}

// Prepay trade perpay
func (t Trade) Prepay(args Sign) (string, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return "", err
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
		return "", err
	}
	result := Prepay{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	if result.ReturnCode != Success {
		return "", errors.New(result.ReturnMsg)
	}
	if result.ResultCode != Success {
		return "", errors.New(result.ErrCodeDes)
	}
	return result.PrepayID, nil
}

// Verify verify
func (t Trade) Verify(args Notice, key string) error {
	sign, err := t.Sign(args, key)
	if err != nil {
		return err
	}
	fmt.Println(args.Sign, sign)
	if args.Sign != sign {
		return errors.New("签名错误")
	}
	return nil
}
