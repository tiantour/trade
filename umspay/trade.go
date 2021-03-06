package umspay

import (
	"encoding/hex"
	"encoding/json"
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

// Alipay trade alipay
func (t *Trade) Alipay(args *Request) (*Response, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("Authorization", fmt.Sprintf("OPEN-ACCESS-TOKEN AccessToken=%s", AccessToken))
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    "https://api-mop.chinaums.com/v1/netpay/trade/create",
		Body:   body,
		Header: header,
	})
	if err != nil {
		return nil, err
	}

	result := Response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != Success {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}

// Wxpay trade wxpay
func (t *Trade) Wxpay(args *Request) (*Response, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("Authorization", fmt.Sprintf("OPEN-ACCESS-TOKEN AccessToken=%s", AccessToken))
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    "https://api-mop.chinaums.com/v1/netpay/wx/unified-order",
		Body:   body,
		Header: header,
	})
	if err != nil {
		return nil, err
	}

	result := Response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != Success {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}

// Verify verify
func (t *Trade) Verify(args *url.Values, key string) error {
	s := args.Get("sign")
	args.Del("sign")

	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return err
	}
	query = fmt.Sprintf("%s%s", query, key)
	sign := strings.ToUpper(
		hex.EncodeToString(
			rsae.NewSHA().SHA256(query),
		),
	)
	if sign != s {
		return errors.New("签名错误")
	}
	return nil
}

// Query query
func (t *Trade) Query(args *Query) (*Response, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("Authorization", fmt.Sprintf("OPEN-ACCESS-TOKEN AccessToken=%s", AccessToken))
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    "https://api-mop.chinaums.com/v1/netpay/query",
		Body:   body,
		Header: header,
	})
	if err != nil {
		return nil, err
	}

	result := Response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != Success {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
