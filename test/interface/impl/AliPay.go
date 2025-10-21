package impl

import "fmt"

type AlipayPay struct {
	Balance int
}

func (a *AlipayPay) Pay(amount int) {
	a.Balance -= amount
}

func (a *AlipayPay) Print() {
	fmt.Println("AlipayPay balance:", a.Balance)
}
