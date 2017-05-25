# trade
trade for alipay,wxpay with go

### Alipay

sign

```
package main

import (
	"encoding/json"
	"fmt"

	"github.com/tiantour/tempo"
	"github.com/tiantour/trade/alipay"
)

func main() {
	sign, err := sign(alipay.Request{
		Body:        "测试商品详情",
		Subject:     "测试商品名称",
		OutTradeNo:  "1234567812345678",
		TotalAmount: 0.02,
		ProductCode: "QUICK_WAP_PAY",
	})
	fmt.Println(sign, err)
}

func sign(args alipay.Request) (string, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return "", err
	}
	// 签名
	sign, err := alipay.NewTrade().Sign(alipay.Sign{
		AppID:      "your appid",
		Method:     "your method",
		ReturnURL:  "your return url",
		Charset:    "utf-8",
		SignType:   "RSA2",
		TimeStamp:  tempo.NewNow().String(),
		Version:    "1.0",
		NotifyURL:  "your notify url",
		BizContent: string(body),
	}, "your private key")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s?%s",
		"your gateway",
		sign,
	), nil
}
```

notify

```
package main

import (
	"fmt"

	"github.com/tiantour/trade/alipay"
)

func main() {
	args := alipay.Notice{}
	err := alipay.NewTrade().Verify(args, "your public key")
	fmt.Println(err)
}
```

### Wxpay

sign 

```
package main

import (
	"fmt"
	"strconv"

	"github.com/tiantour/imago"
	"github.com/tiantour/tempo"
	"github.com/tiantour/trade/wxpay"
)

func main() {
	args := wxpay.Request{
		Body:       "测试商品详情",
		Detail:     "商品名称",
		ProductID:  "商品编号",
		OutTradeNo: "1234567812345678",
		TotalFee:   1,
	}
	sign, err := sign(args)
	fmt.Println(sign, err)
}

func sign(args wxpay.Request) (wxpay.Defray, error) {
	sign := wxpay.Sign{
		AppID:          "your appid",
		MchID:          "your machid",
		NonceStr:       imago.NewRandom().String(16),
		TradeType:      "JSAPI",
		SpbillCreateIP: "user ip address",
		NotifyURL:      "your notify url",
		OpenID:         "user openid",
		Request:        args,
	}
	str, err := wxpay.NewTrade().Sign(sign, "your api key")
	if err != nil {
		return wxpay.Defray{}, err
	}
	sign.Sign = str
	prepay, err := wxpay.NewTrade().Prepay(sign)
	if err != nil {
		return wxpay.Defray{}, err
	}
	defray := wxpay.Defray{
		AppID:     "your app id",
		Package:   "prepay_id=" + prepay,
		NonceStr:  imago.NewRandom().String(16),
		TimeStamp: strconv.FormatInt(tempo.NewNow().Unix(), 10),
		SignType:  "MD5",
	}
	str, err = wxpay.NewTrade().Sign(defray, "your api key")
	if err != nil {
		return wxpay.Defray{}, err
	}
	defray.PaySign = str
	return defray, nil
}
```

notify

```
package main

import (
	"fmt"

	"github.com/tiantour/trade/wxpay"
)

func main() {
	args := wxpay.Notice{}
	err := wxpay.NewTrade().Verify(args, "your api key")
	fmt.Println(err)
}
```