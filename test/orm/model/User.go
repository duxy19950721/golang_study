package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Model    gorm.Model `gorm:"embedded"`
	Age      uint8
	Birthday *time.Time
}
