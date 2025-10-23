package model

import (
	"gorm.io/gorm"
)

type User struct {
	Model gorm.Model
	Name  string // A regular string field
	Age   uint8
}
