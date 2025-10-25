package main

import "golang_study/test/interface/impl"

// 注意：由于PayInterface.go和main.go在同一个目录(test/interface/)下，并且都声明为package main，
// 所以PayInterface接口应该可以直接使用，不需要额外导入

func main() {
	aliPay := impl.AlipayPay{Balance: 100}
	ebayPay := impl.EbayPay{Balance: 100}

	aliPay.Pay(10)
	aliPay.Print()

	ebayPay.Pay(10)
	ebayPay.Print()

	// 直接使用PayInterface类型，因为它在同一个包中
	var a impl.PayInterface = &aliPay
	a.Pay(10)
	a.Print()
}
