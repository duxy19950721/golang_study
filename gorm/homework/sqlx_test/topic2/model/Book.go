package model

import "github.com/shopspring/decimal"

type Book struct {
	ID     int
	Name   string
	Author string
	Price  decimal.Decimal
}
