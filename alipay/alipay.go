package alipay

var (
	// Sign Sign
	Sign = &sign{}
	// Notify Notify
	Notify = &notify{}
)

type (
	// App App
	App struct{}
	// Wap Wap
	Wap struct{}
	// Notify Notify
	notify struct{}
	// Sign Sign
	sign struct {
		App
		Wap
	}
)
