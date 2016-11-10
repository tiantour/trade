package wxpayx

import (
	"encoding/xml"
	"errors"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/imago"
	"github.com/tiantour/requests"
)

// prepayID
func (s *sign) prepayID(args signIn) (string, error) {
	requestBody, err := xml.Marshal(args)
	if err != nil {
		return "", err
	}
	requestURL := "https://api.mch.weixin.qq.com/pay/unifiedorder"
	_, _, requestHeader := requests.Options()
	requestHeader.Add("Accept", "application/xml")
	requestHeader.Add("Content-Type", "application/xml;charset=utf-8")
	body, err := requests.Post(requestURL, requestBody, requestHeader)
	if err != nil {
		return "", err
	}
	result := signOut{}
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

// signStr
func (s *sign) signStr(args interface{}, key string) (string, error) {
	// url.Values to string
	signURL, err := query.Values(args)
	if err != nil {
		return "", err
	}
	// url encode string to url unencode string
	str, err := url.QueryUnescape(signURL.Encode())
	if err != nil {
		return "", err
	}
	str = imago.Crypto.Md532(str + "&key=" + key)
	return strings.ToUpper(str), nil
}
