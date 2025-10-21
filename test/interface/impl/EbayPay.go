package impl

import "fmt"

type EbayPay struct {
	Balance int
}

func (a *EbayPay) Pay(amount int) {
	a.Balance -= amount
}

func (a *EbayPay) Print() {
	fmt.Println("EbayPay balance:", a.Balance)
}
