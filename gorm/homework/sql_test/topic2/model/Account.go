package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model `gorm:"embedded"`
	ID         string
	Balance    int
}
