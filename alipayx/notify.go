package alipayx

import (
	"net/url"

	"github.com/tiantour/conf"
	"github.com/tiantour/requests"
	"github.com/tiantour/rsae"
)

// Notify 通知
func (n *notify) Notify(data url.Values) bool {
	// 验证签名
	signString := data.Get("sign")
	data.Del("sign")
	data.Del("sign_type")
	signStatus := n.verifySign(data, signString)
	// 验证通知
	notifyID := data.Get("notify_id")
	notifyStatus := n.verifyNotify(notifyID)
	// 验证商户
	sellerID := data.Get("seller_id")
	sellerStatus := n.verifySeller(sellerID)
	//签名不正确,通知不正确,商家不正确,返回错误
	if !signStatus || !notifyStatus || !sellerStatus {
		return false
	}
	return true
}

// sign 签名
func (n *notify) verifySign(data url.Values, signString string) bool {
	signURL, err := url.QueryUnescape(data.Encode())
	if err != nil {
		return false
	}
	signStatus, err := rsae.RSA.Verify(signURL, signString)
	if err != nil {
		return false
	}
	return signStatus
}

// notify 通知
func (n *notify) verifyNotify(notifyID string) bool {
	requestURL, requestData, requestHeader := requests.Options()
	requestURL = "https://mapi.alipay.com/gateway.do?service=notify_verify&partner=" + conf.Options.Alipay.Partner + "&notify_id=" + notifyID
	body, err := requests.Get(requestURL, requestData, requestHeader)
	if err != nil || string(body) != "true" {
		return false
	}
	return true
}

// seller 商家
func (n *notify) verifySeller(sellerID string) bool {
	if sellerID != conf.Options.Alipay.Partner {
		return false
	}
	return true
}
