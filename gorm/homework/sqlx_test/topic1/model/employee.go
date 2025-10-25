package model

import "github.com/shopspring/decimal"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     decimal.Decimal
}
