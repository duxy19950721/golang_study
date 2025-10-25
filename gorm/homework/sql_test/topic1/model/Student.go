package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model `gorm:"embedded"`
	Name       string
	Age        uint
	Grade      string
}
