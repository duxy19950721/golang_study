package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model    `gorm:"embedded"`
	FromAccountId string `转出账户ID`
	ToAccountId   string `转入账户ID`
	Amount        int    `转账金额`
}
