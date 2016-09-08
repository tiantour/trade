package wxpay

import "encoding/xml"

// Notify 通知
func (n *notify) Notify(body []byte) (bool, NotifyPayInput) {
	notifyStruct := NotifyPayInput{}
	err := xml.Unmarshal(body, &notifyStruct)
	if err != nil {
		return false, notifyStruct
	}
	if notifyStruct.ReturnCode != success || notifyStruct.ResultCode != success {
		return false, notifyStruct
	}
	return true, notifyStruct
}
