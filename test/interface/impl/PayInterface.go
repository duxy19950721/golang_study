package impl

// 定义支付接口
type PayInterface interface {
	Pay(amount int)
	Print()
}
